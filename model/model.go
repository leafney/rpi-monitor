/**
 * @Author:      leafney
 * @Date:        2022-12-24 23:00
 * @Project:     rpi-monitor
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package model

type (
	Monitor struct {
		Common  CommonInfo     `json:"common"`
		Basic   *BasicInfo     `json:"basic,omitempty"`
		OS      *OSInfo        `json:"os,omitempty"`
		CPU     *CpuInfo       `json:"cpu,omitempty"`
		MEM     *MemoryInfo    `json:"mem,omitempty"`
		Swap    *SwapInfo      `json:"swap,omitempty"`
		Drives  *DrivesInfo    `json:"drives,omitempty"`
		Network *[]NetworkInfo `json:"network,omitempty"`
	}

	BasicInfo struct {
		Timestamp     string `json:"timestamp"`
		DateTime      string `json:"date_time"`
		BootTime      string `json:"boot_time"`
		DeviceModel   string `json:"device_model"`
		UpTime        string `json:"up_time"`
		UpTimeDays    string `json:"up_time_days"`
		UpTimeHours   string `json:"up_time_hours"`
		UpTimeMinutes string `json:"up_time_minutes"`
	}

	OSInfo struct {
		Hostname string `json:"hostname"`
		FQDN     string `json:"fqdn"`
		Kernel   string `json:"kernel"`
		Distro   string `json:"distro"`
		Codename string `json:"codename"`
	}

	CpuInfo struct {
		Temperature      float64 `json:"temperature"`
		TempUnit         string  `json:"temp_unit"`
		Load1minPercent  string  `json:"load_1min_percent"`
		Load5minPercent  string  `json:"load_5min_percent"`
		Load15minPercent string  `json:"load_15min_percent"`
		Hardware         string  `json:"hardware"`
		Revision         string  `json:"revision"`
		Serial           string  `json:"serial"`
		CpuCores         int     `json:"cpu_cores"`
		ModelName        string  `json:"model_name"`
		Arch             string  `json:"arch"`
		ArchModel        string  `json:"arch_model"`
	}

	MemoryInfo struct {
		TotalMB     string `json:"total_mb"`
		UsedMB      string `json:"used_mb"`
		UsedPercent string `json:"used_percent"`
		//FreeMB      string `json:"free_mb"`
		//FreePercent string `json:"free_percent"`
	}

	SwapInfo struct {
		TotalMB     string `json:"total_mb"`
		UsedMB      string `json:"used_mb"`
		UsedPercent string `json:"used_percent"`
		//FreeMB  string `json:"free_mb"`
		//FreePercent string `json:"free_percent"`
	}

	DrivesInfo struct {
		Total       string `json:"total"`
		Used        string `json:"used"`
		Free        string `json:"free"`
		UsedPercent string `json:"used_percent"`
	}

	NetworkInfo struct {
		Name string `json:"name"`
		IP   string `json:"ip"`
		MAC  string `json:"mac"`
	}

	CommonInfo struct {
		Name        string `json:"name"`
		Version     string `json:"version"`
		IntervalSec int64  `json:"interval_sec"`
	}
)
