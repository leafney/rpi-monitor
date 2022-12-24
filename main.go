package main

import (
	"fmt"
	"github.com/leafney/rpi-monitor/module"
)

func main() {
	//	test basic info

	fmt.Println("boot_time ", module.GetBasicBootTime())
	fmt.Println("date_time ", module.GetBasicDateTime())
	fmt.Println("timestamp ", module.GetBasicTimestamp())
	fmt.Println("device_mode ", module.GetBasicDeviceModel())
	upTime := module.GetBasicUpTime()
	fmt.Println("uptime ", upTime)
	utd, uth, utm := module.GetBasicUpTimeDHM(upTime)
	fmt.Printf("uptime_days [%s] uptime_hours [%s] uptime_mins [%s] \n", utd, uth, utm)
	//boot_time  2022-11-24 03:13:09
	//date_time  2022-12-24 18:04:49
	//timestamp  2022-12-24T18:04:49+08:00
	//device_mode  RPi 3 Model B+ r1.3
	//uptime  30 days, 14 hours, 51 mins
	//uptime_days [30] uptime_hours [14] uptime_mins [51]
}
