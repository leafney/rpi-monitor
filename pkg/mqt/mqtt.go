/**
 * @Author:      leafney
 * @Date:        2022-12-27 14:39
 * @Project:     rpi-monitor
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package mqt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/leafney/rose"
	"github.com/leafney/rpi-monitor/config"
	"github.com/leafney/rpi-monitor/model"
	"github.com/leafney/rpi-monitor/pkg/metrics"
	"log"
	"os"
	"time"
)

func CrtMQTTClientID() (string, error) {
	host, err := os.Hostname()
	if err != nil {
		return "", err
	}
	pid := os.Getpid()
	return fmt.Sprintf("%s-%d", host, pid), nil
}

func GetMQTTClient(cfg config.MQTTConfig) mqtt.Client {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", cfg.Host, cfg.Port))
	opts.SetCleanSession(true)
	opts.SetAutoReconnect(true)

	if !rose.StrIsEmpty(cfg.UserName) {
		opts.SetUsername(cfg.UserName)
	}
	if !rose.StrIsEmpty(cfg.Password) {
		opts.SetPassword(cfg.Password)
	}

	//	clientID

	// tls

	opts.SetOnConnectHandler(connectHandler)
	opts.SetConnectionLostHandler(connectLostHandler)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("mqtt connected")

	Publish(client, "topic/hello")

}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("mqtt connect lost err: %v\n", err)
}

func Publish(client mqtt.Client, topic string) {

	log.Println("publish ...")

	go func() {
		for {
			log.Println("to load monitor info")

			data := &model.Monitor{}
			data.Basic = metrics.ShowBaseInfo()
			data.MEM, data.Swap = metrics.ShowMemInfo()
			value := rose.JsonMarshalStr(data)
			log.Println(value)

			token := client.Publish(topic, 1, true, value)
			token.Wait()

			time.Sleep(30 * time.Second)
		}
	}()

}

func Listen(client mqtt.Client) {

}
