## 测试参数获取

```go
package main

import (
	"fmt"
	"github.com/leafney/rpi-monitor/module"
)

func main() {
	//	test basic info
	/*
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

	*/

	//	test os info
	/*
		fmt.Println(module.GetOSHostName(""))
		fmt.Println(module.GetOSKernel())
		fmt.Println(module.GetOSDistro())
		fmt.Println(module.GetOSCodeName())
		//raspberrypi raspberrypi
		//Linux 5.10.17-v7+
		//Raspbian GNU/Linux 10 (buster)
		//buster

	*/

	//	test cpu info
	/*
		fmt.Println(module.GetCPUArch())
		fmt.Println(module.GetCPUTemperatureC())
		fmt.Println(module.GetCPUTemperatureF())
		fmt.Println(module.GetCPUInfo())
		fmt.Println(module.GetCPULoadAvg(4))
		//armv7l ARM Cortex-A53
		//60.1
		//140.18
		//BCM2835 ARMv7 Processor rev 4 (v7l) Raspberry Pi 3 Model B Plus Rev 1.3 a020d3 000000007a867fb6 4
		//15.75 8.50 7.75

	*/

	//	test memory info
	fmt.Println(module.GetMemoryInfo())
	// 946392 124880 401852 51552 256136 102396 0

	//	test drives info
	fmt.Println(module.GetDrivesFileSystem())
	// 15G 5.7G 8.1G 42

	//	test network info
	fmt.Println(module.GetNetworkIPs())

}

```