/**
 * @Author:      leafney
 * @Date:        2022-12-24 23:21
 * @Project:     rpi-monitor
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package config

type Config struct {
	MQTT MQTTConfig `mapstructure:"mqtt"`
}

type MQTTConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Topic    string `mapstructure:"topic"`
	ClientId string `mapstructure:"client_id"`
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Qos      byte   `mapstructure:"qos"`
}
