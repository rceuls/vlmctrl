package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"github.com/rceuls/vlmctrl/internal/config"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("TOPIC: %s -MSG: %s\n", msg.Topic(), msg.Payload())
	payload := string(msg.Payload())
	translated := convert(payload)
	for _, cmd := range translated {
		sendCommand(cmd)
	}
}

func cleanup(client mqtt.Client) {
	if token := client.Unsubscribe(config.MqttTopic()); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
		os.Exit(1)
	}

	client.Disconnect(250)
	log.Println("Disconnected client")
}

func setupCloseHandler(client mqtt.Client) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup(client)
		os.Exit(0)
	}()
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	mqtt.ERROR = log.New(os.Stdout, "", 0)

	opts := mqtt.NewClientOptions().AddBroker(config.MqttBroker()).SetClientID(config.MqttClientID())
	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	setupCloseHandler(c)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe(config.MqttTopic(), 0, nil); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
		os.Exit(1)
	}

	log.Print("Starting polling as pid ", os.Getpid())

	for {
		time.Sleep(10 * time.Millisecond)
	}
}
