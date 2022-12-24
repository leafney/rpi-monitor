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

// GetBasicUpTime
// 30 days, 14 hours, 3 mins
// - [linux - Bash format uptime to show days, hours, minutes - Stack Overflow](https://stackoverflow.com/questions/28353409/bash-format-uptime-to-show-days-hours-minutes)
func GetBasicUpTime() string {
	cmdStr := `/usr/bin/uptime | awk -F '[ ,:\t\n]+' '{
    msg = ""
    if ($7 == "day" || $7 == "days") {
        msg = msg $6 " " $7 ", "

        h = $8
        m = $9
    } else {
        h = $6
        m = $7
    }

    if (int(m) == 0) {
        msg = msg int(h)" "m
    } else {
        msg = msg int(h)" hour"
        if (h > 1) { msg = msg "s"}

        msg = msg ", " int(m) " min"
        if (m > 1) { msg = msg "s"}
    }

    print msg
}'`

	res, _ := utils.RunCommand(cmdStr)
	return res
}

// GetBasicUpTimeDHM
// 30 days, 14 hours, 3 mins ==> 30 14 3
func GetBasicUpTimeDHM(uptime string) (day, hour, min string) {
	if uptime == "" {
		return
	}

	uts := strings.Split(uptime, ",")
	for _, ut := range uts {
		if strings.Contains(ut, "day") {
			day = strings.TrimSpace(utils.StrRemoveAny(ut, "days", "day"))
		} else if strings.Contains(ut, "hour") {
			hour = strings.TrimSpace(utils.StrRemoveAny(ut, "hours", "hour"))
		} else if strings.Contains(ut, "min") {
			min = strings.TrimSpace(utils.StrRemoveAny(ut, "mins", "min"))
		}
	}

	return
}
