/**
 * @Author:      leafney
 * @Date:        2022-12-23 20:24
 * @Project:     rpi-monitor
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package utils

import (
	"os"
	"strconv"
)

func StrToFloat64(s string) float64 {
	if s == "" {
		return 0.0
	}
	if i, err := strconv.ParseFloat(s, 64); err != nil {
		return 0.0
	} else {
		return i
	}
}

func StrToFloat64WithDef(s string, def float64) float64 {
	if s == "" {
		return def
	}
	if i, err := strconv.ParseFloat(s, 64); err != nil {
		return def
	} else {
		return i
	}
}

func FIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
