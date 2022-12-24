/**
 * @Author:      leafney
 * @Date:        2022-12-24 14:15
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

// GetNetworkIPs
func GetNetworkIPs() {
	cmdIp := getIPCmd()
	if cmdIp != "" {
		interfaceNames := GetNetworkIFNames(cmdIp)

		for _, ifName := range interfaceNames {
			interfaceNameDetailLines := GetNetworkIFDetails(cmdIp, ifName)
			ipAddr, macAddr := LoadNetworkIFDetailLines(interfaceNameDetailLines)

			fmt.Printf("IF[%s] Mac[%s] Ip[%s]\n", ifName, macAddr, ipAddr)
		}
	}
}

func getIPCmd() string {
	var (
		cmdIp1 = "/bin/ip"
		cmdIp2 = "/usr/bin/ip"
		cmdIp3 = "/sbin/ip"
		cmdIp4 = "/usr/sbin/ip"
		desCmd = ""
	)

	if utils.FIsExist(cmdIp1) {
		desCmd = cmdIp1
	} else if utils.FIsExist(cmdIp2) {
		desCmd = cmdIp2
	} else if utils.FIsExist(cmdIp3) {
		desCmd = cmdIp3
	} else if utils.FIsExist(cmdIp4) {
		desCmd = cmdIp4
	}

	return desCmd
}

// GetNetworkIFNames
//	➜ ip link show | /bin/egrep -v 'link' | /bin/egrep ' eth| wlan'
//	2: eth0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc pfifo_fast state DOWN mode DEFAULT group default qlen 1000
//	3: wlan0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP mode DORMANT group default qlen 1000
// returns [eht0 wlan0]
func GetNetworkIFNames(ipCmd string) []string {
	cmdStr := fmt.Sprintf(`%s link show | /bin/egrep -v "link" | /bin/egrep " eth| wlan"`, ipCmd)
	res, _ := utils.RunCommand(cmdStr)

	interfaceNames := make([]string, 0)
	trimmedLines := utils.StrTrimLines(res)

	for _, line := range trimmedLines {
		lineParts := utils.StrSplitAny(line, " :")
		if len(lineParts) >= 2 {
			interfaceNames = append(interfaceNames, lineParts[1])
		}
	}

	return interfaceNames
}

// GetNetworkIFDetails
//	➜ ip addr show eth0 | egrep "Link|flags|inet |ether | eth0" | /bin/egrep -v -i "lo:|loopback|inet6|\:\:1|127\.0\.0\.1"
//	2: eth0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc pfifo_fast state DOWN group default qlen 1000
//	link/ether b8:27:eb:86:7f:b6 brd ff:ff:ff:ff:ff:ff
func GetNetworkIFDetails(ipCmd, interfaceName string) []string {
	cmdStr := fmt.Sprintf(`%s address show %s | /bin/egrep "Link|flags|inet |ether | %s" | /bin/egrep -v -i "lo:|loopback|inet6|\:\:1|127\.0\.0\.1"`, ipCmd, interfaceName, interfaceName)
	res, _ := utils.RunCommand(cmdStr)

	trimmedLines := utils.StrTrimLines(res)
	return trimmedLines
}

// LoadNetworkIFDetailLines
func LoadNetworkIFDetailLines(detailLines []string) (ipAddr, macAddr string) {
	var (
		haveIF  = false
		isIPCmd = false
	)

	for _, curLine := range detailLines {
		lineParts := strings.Split(curLine, "")
		if len(lineParts) > 0 {

			if strings.Contains(curLine, "flags") {
				// NEWER
				haveIF = true
			} else if utils.StrContainsAny(curLine, "mtu", "qlen") {
				// IP Command
				haveIF = true
				isIPCmd = true
			} else if strings.Contains(curLine, "Link") {
				// OLDER
				haveIF = true

				macAddr = lineParts[4]

			} else if haveIF && !isIPCmd {
				// OLDER & NEWER
				if strings.Contains(curLine, "inet") {
					ipAddr = strings.ReplaceAll(lineParts[1], "addr:", "")
				} else if strings.Contains(curLine, "ether") {
					// NEWER
					macAddr = lineParts[1]

					//	??
					haveIF = false
				}
			} else if haveIF && isIPCmd {
				//	IP Command
				if strings.Contains(curLine, "inet") {
					inets := strings.Split(lineParts[1], "/")
					if len(inets) == 2 {
						ipAddr = inets[0]
					}
					// ??
					haveIF = false
					isIPCmd = false
				} else if strings.Contains(curLine, "link/ether") {
					macAddr = lineParts[1]
				}
			}

		}
	}

	return
}
