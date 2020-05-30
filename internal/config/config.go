package config

import "os"

// MqttTopic topic to subscribe to
func MqttTopic() string {
	return os.Getenv("MQTT_TOPIC")
}

// MqttBroker broker url
func MqttBroker() string {
	return os.Getenv("MQTT_BROKER")
}

// MqttClientID consumer client id
func MqttClientID() string {
	return os.Getenv("MQTT_CLIENTID")
}
