/**
 * @Author:      leafney
 * @Date:        2022-12-24 03:33
 * @Project:     rpi-monitor
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package module

import (
	"fmt"
	"github.com/leafney/rpi-monitor/utils"
	"strings"
)

// GetCPUInfo
//	➜ cat /proc/cpuinfo | /bin/egrep -i 'processor|model|hardware|revision|serial' | sort | uniq
//	CPU revision    : 4
//	Hardware        : BCM2835
//	model name      : ARMv7 Processor rev 4 (v7l)
//	Model           : Raspberry Pi 3 Model B Plus Rev 1.3
//	processor       : 0
//	processor       : 1
//	processor       : 2
//	processor       : 3
//	Revision        : a020d3
//	Serial          : 000000007a867fb6
func GetCPUInfo() (hardware, modelName, model, rev, serial string, cores int) {
	res, _ := utils.RunCommand(`/bin/cat /proc/cpuinfo | /bin/egrep -i 'processor|model|hardware|revision|serial' | sort | uniq`)

	trimmedLines := utils.StrTrimLines(res)

	for _, curLine := range trimmedLines {
		curValue := ""
		lineParts := strings.Split(curLine, ":")
		if len(lineParts) >= 2 {
			curValue = strings.TrimSpace(lineParts[1])
		}

		if strings.HasPrefix(curLine, "Hardware") {
			hardware = curValue
		} else if strings.HasPrefix(curLine, "model name") {
			modelName = curValue
		} else if strings.HasPrefix(curLine, "processor") {
			cores += 1
		} else if strings.HasPrefix(curLine, "Model") {
			model = curValue
		} else if strings.HasPrefix(curLine, "Revision") {
			rev = curValue
		} else if strings.HasPrefix(curLine, "Serial") {
			serial = curValue
		}
	}

	return
}

// GetCPULoadAvg percent
func GetCPULoadAvg(cores int) (load1, load5, load15 string) {
	res, _ := utils.RunCommand(`/bin/cat /proc/loadavg`)
	cpuLoads := strings.Split(res, " ")

	load1 = fmt.Sprintf("%.2f", utils.StrToFloat64(cpuLoads[0])/float64(cores)*100)
	load5 = fmt.Sprintf("%.2f", utils.StrToFloat64(cpuLoads[1])/float64(cores)*100)
	load15 = fmt.Sprintf("%.2f", utils.StrToFloat64(cpuLoads[2])/float64(cores)*100)

	return
}

// GetCPUTemperatureC
// eg: 59.1 ℃
func GetCPUTemperatureC() float64 {
	res := ""
	cmdGen := getVcGenCmd()
	if cmdGen != "" {
		res, _ = utils.RunCommand(fmt.Sprintf(`%s measure_temp | /usr/bin/awk -F "[=\']" '{print $2}'`, cmdGen))
	} else {
		// get another way
		res = GetCPUSysTemp()
	}
	return utils.StrToFloat64(res)
}

// GetCPUTemperatureF
// eg: 136.4 ℉
func GetCPUTemperatureF() float64 {
	res := ""
	cmdGen := getVcGenCmd()
	if cmdGen != "" {
		res, _ = utils.RunCommand(fmt.Sprintf(`%s measure_temp | /usr/bin/awk -F "[=\']" '{print($2 * 1.8)+32}'`, cmdGen))
		return utils.StrToFloat64(res)
	} else {
		res = GetCPUSysTemp()
		// Convert Celsius to Fahrenheit
		resF := (utils.StrToFloat64(res) * 1.8) + 32.0
		return resF
	}
}

func getVcGenCmd() string {
	cmdVcGen1 := "/usr/bin/vcgencmd"
	cmdVcGen2 := "/opt/vc/bin/vcgencmd"

	desCmd := cmdVcGen1
	if !utils.FIsExist(desCmd) {
		desCmd = cmdVcGen2
		if !utils.FIsExist(desCmd) {
			desCmd = ""
		}
	}

	return desCmd
}

// GetCPUSysTemp
func GetCPUSysTemp() string {
	res, _ := utils.RunCommand(`/bin/cat /sys/class/thermal/thermal_zone0/temp`)
	return fmt.Sprintf("%.1f", utils.StrToFloat64(res)/1000.0)
}

// GetCPUArch
//	➜ /usr/bin/lscpu | /bin/egrep -i 'architecture|vendor|model'
//	Architecture:        armv7l
//	Vendor ID:           ARM
//	Model:               4
//	Model name:          Cortex-A53
func GetCPUArch() (arch, archModel string) {
	res, _ := utils.RunCommand(`/usr/bin/lscpu | /bin/egrep -i 'architecture|vendor|model'`)

	trimmedLines := utils.StrTrimLines(res)

	var (
		vendor = ""
		model  = ""
	)
	for _, curLine := range trimmedLines {
		curValue := ""
		lineParts := strings.Split(curLine, ":")
		if len(lineParts) >= 2 {
			curValue = strings.TrimSpace(lineParts[1])
		}

		if strings.HasPrefix(curLine, "Architecture") {
			arch = curValue
		} else if strings.HasPrefix(curLine, "Vendor") {
			vendor = curValue
		} else if strings.HasPrefix(curLine, "Model name") {
			model = curValue
		}
	}
	archModel = fmt.Sprintf("%s %s", vendor, model)

	return
}
