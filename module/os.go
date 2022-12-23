/**
 * @Author:      leafney
 * @Date:        2022-12-24 03:06
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

// GetOSHostName Get hostname and FQDN, support for custom domain name
func GetOSHostName(defDomain string) (string, string) {
	var resHostname, resFQDN = "", ""
	res, _ := utils.RunCommand(`/bin/hostname -f`)

	resHostname = res
	// check domain exist
	if strings.Contains(res, ".") {
		nameParts := strings.Split(res, ".")
		resFQDN = res
		resHostname = nameParts[0]
	} else {
		if len(defDomain) > 0 {
			resFQDN = fmt.Sprintf("%s.%s", res, defDomain)
		} else {
			resFQDN = resHostname
		}
	}

	return resHostname, resFQDN
}

// GetOSKernel Get kernel version
// eg: Linux 5.10.17-v7+
func GetOSKernel() string {
	res, _ := utils.RunCommand(`/bin/uname -rs`)
	return res
}

// GetOSDistro
// eg: Raspbian GNU/Linux 10 (buster)
func GetOSDistro() string {
	res, _ := utils.RunCommand(`/usr/bin/lsb_release -d | awk -F ':' '{print $2}'`)
	return res
}

// GetOSCodeName
// eg: buster
func GetOSCodeName() string {
	res, _ := utils.RunCommand(`/usr/bin/lsb_release -c | awk -F ':' '{print $2}'`)
	return res
}
