package mqtt

import (
	"encoding/json"
	"fmt"
	"lscdoorbellmqtt/config"
	"lscdoorbellmqtt/gpiohandler"
	"lscdoorbellmqtt/logger"
	"lscdoorbellmqtt/sound"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	manufacturer      = "LSC"
	fullName          = manufacturer + " Connect Video Doorbell"
	deviceClass       = "sound"
	mqttClass         = "binary_sensor"
	topic             = "homeassistant/" + mqttClass + "/lscdoorbell/bell"
	subscribeTopic    = topic + "/#"
	stateTopic        = topic + "/contact"
	configTopic       = topic + "/config"
	availabilityTopic = topic + "/availability"
	onPayload         = "ON"
	offPayload        = "OFF"
	onlinePayload     = "online"
	offlinePayload    = "offline"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	logger.Infof("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	logger.Status.Println("Connected to MQTT broker")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	logger.Error.Println("Connect lost:", err)
}

func discoverHA(client mqtt.Client) {
	discoveryMessage := map[string]interface{}{
		"name":         "Doorbell",
		"device_class": deviceClass,
		"state_topic":  stateTopic,
		"payload_on":   onPayload,
		"payload_off":  offPayload,
		"availability": []map[string]interface{}{
			{
				"topic":                 availabilityTopic,
				"payload_available":     onlinePayload,
				"payload_not_available": offlinePayload,
			},
		},
		"device": map[string]interface{}{
			"identifiers":  []string{"lscdoorbell1"},
			"manufacturer": manufacturer,
			"model":        fullName,
			"sw_version":   "1.0",
			"name":         fullName,
		},
		"unique_id": "doorbell",
	}

	discoveryPayload, err := json.Marshal(discoveryMessage)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to encode discovery message: %v", err))
	}

	publishConfig(client, discoveryPayload)
}

func publishState(client mqtt.Client, state string) {
	token := client.Publish(stateTopic, 0, true, state)
	token.Wait()
}

func publishAvailability(client mqtt.Client, state string) {
	token := client.Publish(availabilityTopic, 0, true, state)
	token.Wait()
}

func publishConfig(client mqtt.Client, discoveryPayload []byte) {
	token := client.Publish(configTopic, 0, true, discoveryPayload)
	token.Wait()
}

func Start() {
	var broker = config.GetString("settings.mqtt_broker")
	var port = config.GetInt64("settings.mqtt_port")
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(config.GetString("settings.mqtt_client_id"))
	opts.SetUsername(config.GetString("settings.mqtt_username"))
	opts.SetPassword(config.GetString("settings.mqtt_password"))
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Fatal(token.Error().Error())
		os.Exit(1)
	}

	subscribe(client)

	discoverHA(client)

	publishState(client, offPayload)

	stateLoop(client)

	client.Disconnect(250)
}

func stateLoop(client mqtt.Client) {
	updateTicker := 50
	sendUpdateTimes := 10

	for {
		bellState := gpiohandler.GetBellState()

		if bellState == 0 {
			handleBellState(client, sendUpdateTimes)
		}

		updateAvailability(client, &updateTicker)
		time.Sleep(100 * time.Millisecond)
	}
}

func handleBellState(client mqtt.Client, sendUpdateTimes int) {
	go sound.PlaySound("dingdong.wav")
	go gpiohandler.BellBlink()

	for i := 1; i < sendUpdateTimes; i++ {
		publishState(client, onPayload)
		time.Sleep(1 * time.Second)
	}

	publishState(client, offPayload)
}

func updateAvailability(client mqtt.Client, updateTicker *int) {
	if *updateTicker == 0 {
		publishAvailability(client, onlinePayload)
		*updateTicker = 50
	}
	*updateTicker--
}

func subscribe(client mqtt.Client) {
	token := client.Subscribe(subscribeTopic, 1, nil)
	token.Wait()
	logger.Infof("Subscribed to topic: %s", subscribeTopic)
}
