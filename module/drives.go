/**
 * @Author:      leafney
 * @Date:        2022-12-24 13:25
 * @Project:     rpi-monitor
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package module

import (
	"github.com/leafney/rpi-monitor/utils"
	"strings"
)

// GetDrivesFileSystem
//	➜ df -h / | tail -n 1 | awk '{print $2,$3,$4,$5}'
//	15G 5.7G 8.1G 42%
func GetDrivesFileSystem() (total, used, free, usedPercent string) {
	res, _ := utils.RunCommand(`/bin/df -h / | /usr/bin/tail -n 1 | /usr/bin/awk '{print $2,$3,$4,$5}`)

	lineParts := strings.Split(res, " ")
	if len(lineParts) >= 4 {
		total, used, free, usedPercent = lineParts[0], lineParts[1], lineParts[2], strings.ReplaceAll(lineParts[3], "%", "")
	}
	return
}
