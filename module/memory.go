/**
 * @Author:      leafney
 * @Date:        2022-12-24 11:52
 * @Project:     rpi-monitor
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package module

import (
	"github.com/leafney/rpi-monitor/utils"
	"strings"
)

// GetMemoryInfo
//	➜ /bin/cat /proc/meminfo | /bin/egrep -i "mem[tfa]|swap[tf]|buffers|cached"
//	MemTotal:         946392 kB
//	MemFree:           96268 kB
//	MemAvailable:     407268 kB
//	Buffers:           53128 kB
//	Cached:           288424 kB
//	SwapCached:         4264 kB
//	SwapTotal:        102396 kB
//	SwapFree:           2612 kB
func GetMemoryInfo() (memTotal, memFree, memAvail, buff, cache, swapTotal, swapFree string) {
	res, _ := utils.RunCommand(`/bin/cat /proc/meminfo | /bin/egrep -i "mem[tfa]|swap[tf]|buffers|cached"`)

	trimmedLines := utils.StrTrimLines(res)

	for _, curLine := range trimmedLines {
		curValue := ""

		lineParts := utils.StrSplitAny(curLine, ": ")
		if len(lineParts) >= 3 {
			curValue = lineParts[1]
		}

		if strings.HasPrefix(curLine, "MemTotal") {
			memTotal = curValue
		} else if strings.HasPrefix(curLine, "MemFree") {
			memFree = curValue
		} else if strings.HasPrefix(curLine, "MemAvailable") {
			memAvail = curValue
		} else if strings.HasPrefix(curLine, "Buffers") {
			buff = curValue
		} else if strings.HasPrefix(curLine, "Cached") {
			cache = curValue
		} else if strings.HasPrefix(curLine, "SwapTotal") {
			swapTotal = curValue
		} else if strings.HasPrefix(curLine, "SwapFree") {
			swapFree = curValue
		}
	}

	return
}
