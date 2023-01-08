package main

import (
	"fmt"
	"github.com/leafney/rose"
	"github.com/leafney/rpi-monitor/config"
	"github.com/leafney/rpi-monitor/pkg/mqt"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"runtime"
)

var (
	configFile *string
	v          bool
	h          bool
	Version    = "0.1.0"
	GitBranch  = ""
	GitCommit  = ""
	BuildTime  = "2023-01-08 23:37:48"
)

func main() {
	//configFile = flag.StringP("config", "f", "config.yaml", "the config file")

	flag.BoolVarP(&h, "help", "h", false, "help")
	flag.BoolVarP(&v, "version", "v", false, "version")
	flag.Parse()

	if h {
		flag.PrintDefaults()
	} else if v {
		// 输出版本信息
		fmt.Println("Version:      " + Version)
		fmt.Println("Git branch:   " + GitBranch)
		fmt.Println("Git commit:   " + GitCommit)
		fmt.Println("Built time:   " + BuildTime)
		fmt.Println("Go version:   " + runtime.Version())
		fmt.Println("OS/Arch:      " + runtime.GOOS + "/" + runtime.GOARCH)
	} else {
		start()
	}
}

func start() {

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
