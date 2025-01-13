package utils

import (
	"fmt"
	"math"
)

func HumanFileSize(bytes int64, si bool, dp int) string {
	var thresh float64
	if si {
		thresh = 1000
	} else {
		thresh = 1024
	}

	if math.Abs(float64(bytes)) < thresh {
		return fmt.Sprintf("%d B", bytes)
	}

	var units []string
	if si {
		units = []string{"kB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	} else {
		units = []string{"KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"}
	}

	u := -1
	r := math.Pow(10, float64(dp))

	bytesFloat := float64(bytes)
	for math.Round(math.Abs(bytesFloat)*r)/r >= thresh && u < len(units)-1 {
		bytesFloat /= thresh
		u++
	}

	return fmt.Sprintf("%.*f %s", dp, bytesFloat, units[u])
}
