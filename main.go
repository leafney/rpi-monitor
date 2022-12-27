package main

import (
	"fmt"
	"github.com/leafney/rose"
	"github.com/leafney/rpi-monitor/config"
	"github.com/leafney/rpi-monitor/pkg/mqt"
	"github.com/spf13/viper"
)

func main() {

	// set default config
	viper.SetDefault("mqtt", map[string]interface{}{
		"host":     "127.0.0.1",
		"port":     "1883",
		"topic":    "",
		"username": "",
		"password": "",
		"qos":      0})

	// load config file
	viper.SetConfigName("monitor")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/monitor/")
	viper.AddConfigPath("$HOME/.monitor")
	viper.AddConfigPath(".")

	// use env config
	viper.AutomaticEnv()
	viper.SetEnvPrefix("monitor")

	// config change
	//viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("Config file changed:", e.Name)
	//})

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var cfg config.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	fmt.Println(rose.JsonMarshalStr(&cfg))

	//	mqtt
	mqt.GetMQTTClient(cfg.MQTT)

	fmt.Println("starting...")
	select {}

}
