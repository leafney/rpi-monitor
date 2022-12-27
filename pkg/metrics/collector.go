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
)

func ShowBaseInfo() *model.BasicInfo {
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

func ShowMemInfo() (*model.MemoryInfo, *model.SwapInfo) {
	m := &model.MemoryInfo{}
	s := &model.SwapInfo{}

	//mTotal, mFree, avail, buff, cache, sTotal, sFree := module.GetMemoryInfo()

	mTotal, mFree, _, mBuff, mCache, sTotal, sFree := module.GetMemoryInfo()

	mTotalF := rose.StrToFloat64(mTotal)
	mFreeF := rose.StrToFloat64(mFree)
	mBuffF := rose.StrToFloat64(mBuff)
	mCacheF := rose.StrToFloat64(mCache)
	mUsedF := mTotalF - mFreeF - mBuffF - mCacheF
	m.TotalMB = fmt.Sprintf("%.1f", mTotalF/1024.0)
	m.FreeMB = fmt.Sprintf("%.1f", mFreeF/1024.0)
	m.UsedMB = fmt.Sprintf("%.1f", mUsedF/1024.0)
	m.FreePercent = fmt.Sprintf("%.1f", mFreeF/mTotalF*100)
	m.UsedPercent = fmt.Sprintf("%.1f", (mUsedF)/mTotalF*100)

	// swap
	sTotalF := rose.StrToFloat64(sTotal)
	sFreeF := rose.StrToFloat64(sFree)
	sUsedF := sTotalF - sFreeF
	s.TotalMB = fmt.Sprintf("%.1f", sTotalF/1024.0)
	s.FreeMB = fmt.Sprintf("%.1f", sFreeF/1024.0)
	s.UsedMB = fmt.Sprintf("%.1f", sUsedF/1024.0)

	return m, s
}
