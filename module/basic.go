/**
 * @Author:      leafney
 * @Date:        2022-12-24 02:26
 * @Project:     rpi-monitor
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package module

import (
	"github.com/leafney/rpi-monitor/utils"
	"strings"
	"time"
)

// GetBasicDateTime Get the current time of the system
func GetBasicDateTime() string {
	res, _ := utils.RunCommand(`/usr/bin/date "+%Y-%m-%d %H:%M:%S"`)
	return res
}

// GetBasicBootTime Get system boot time
func GetBasicBootTime() string {
	res, _ := utils.RunCommand("/usr/bin/uptime -s")
	return res
}

// GetBasicDeviceModel Get the system device model
func GetBasicDeviceModel() string {
	res, _ := utils.RunCommand(`/bin/cat /proc/device-tree/model`)

	// Raspberry Pi 3 Model B Plus Rev 1.3 ==> RPi 3 Model B+ r1.3
	res = strings.ReplaceAll(res, "Raspberry ", "R")
	res = strings.ReplaceAll(res, "i Model ", "i 1 Model")
	res = strings.ReplaceAll(res, " Plus", "+")
	res = strings.ReplaceAll(res, "Rev ", "r")

	return res
}

func GetBasicTimestamp() string {
	res := time.Now().Format(time.RFC3339)
	return res
}
