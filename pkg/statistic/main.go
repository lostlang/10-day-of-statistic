package statistic

import "github.com/lostlang/10-day-of-statistic/pkg/utils"

type StatisticSlice struct {
	utils.IntSlice64
}

func (s StatisticSlice) Sum() (sum int64) {
	for _, value := range s.GetData() {
		sum += value
	}
	return sum
}

func (s StatisticSlice) Mean() float64 {
	return float64(s.Sum()) / float64(s.Len())
}
