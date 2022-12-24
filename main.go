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
	//boot_time  2022-11-24 03:13:09
	//date_time  2022-12-24 16:41:34
	//timestamp  2022-12-24T16:41:34+08:00
	//device_mode  RPi 3 Model B+ r1.3
}
