/**
 * @Author:      leafney
 * @Date:        2022-12-27 18:06
 * @Project:     rpi-monitor
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package metrics

import (
	"fmt"
	"github.com/leafney/rose"
	"github.com/leafney/rpi-monitor/model"
	"github.com/leafney/rpi-monitor/module"
	"strings"
)

// ShowBasicInfo basic
func ShowBasicInfo() *model.BasicInfo {
	basic := &model.BasicInfo{}

	basic.DateTime = module.GetBasicDateTime()
	basic.Timestamp = module.GetBasicTimestamp()
	basic.BootTime = module.GetBasicBootTime()
	upTime := module.GetBasicUpTime()
	basic.UpTime = upTime
	basic.UpTimeDays, basic.UpTimeHours, basic.UpTimeMinutes = module.GetBasicUpTimeDHM(upTime)
	basic.DeviceModel = module.GetBasicDeviceModel()

	return basic
}

// ShowMemInfo memory and swap
func ShowMemInfo() (*model.MemoryInfo, *model.SwapInfo) {
	m := &model.MemoryInfo{}
	s := &model.SwapInfo{}

	mTotal, mAvail, sTotal, sFree := module.GetMemorySimple()

	mTotalF := rose.StrToFloat64(mTotal)
	mAvailF := rose.StrToFloat64(mAvail)
	mUsedF := mTotalF - mAvailF
	m.TotalMB = fmt.Sprintf("%.1f", mTotalF/1024.0)
	m.UsedMB = fmt.Sprintf("%.1f", mUsedF/1024.0)
	m.UsedPercent = fmt.Sprintf("%.1f", mUsedF/mTotalF*100)

	// swap
	sTotalF := rose.StrToFloat64(sTotal)
	sFreeF := rose.StrToFloat64(sFree)
	sUsedF := sTotalF - sFreeF
	s.TotalMB = fmt.Sprintf("%.1f", sTotalF/1024.0)
	s.UsedMB = fmt.Sprintf("%.1f", sUsedF/1024.0)
	s.UsedPercent = fmt.Sprintf("%.1f", sUsedF/sTotalF*100)

	return m, s
}

// ShowCpuInfo cpu
func ShowCpuInfo(unitName string) *model.CpuInfo {
	c := &model.CpuInfo{}

	var (
	//cModel = ""
	)

	cCores := 0

	if strings.ToUpper(unitName) == "F" {
		c.Temperature = module.GetCPUTemperatureF()
		c.TempUnit = "°F"
	} else {
		c.Temperature = module.GetCPUTemperatureC()
		c.TempUnit = "°C"
	}

	c.Hardware, c.ModelName, _, c.Revision, c.Serial, cCores = module.GetCPUInfo()
	c.CpuCores = cCores
	L1Per, L5Per, L15Per := module.GetCPULoadAvg(cCores)
	c.Load1minPercent = L1Per
	c.Load5minPercent = L5Per
	c.Load15minPercent = L15Per

	c.Arch, c.ArchModel = module.GetCPUArch()

	return c
}

func ShowDrivesInfo() *model.DrivesInfo {
	d := &model.DrivesInfo{}

	d.Total, d.Used, d.Free, d.UsedPercent = module.GetDrivesFileSystem()

	return d
}

func ShowNetworkInfo() *[]model.NetworkInfo {
	ns := module.GetNetworkIPs()
	return &ns
}

func ShowOSInfo(domain string) *model.OSInfo {
	o := &model.OSInfo{}

	o.Kernel = module.GetOSKernel()
	o.Distro = module.GetOSDistro()
	o.Codename = module.GetOSCodeName()
	o.Hostname, o.FQDN = module.GetOSHostName(domain)

	return o
}
