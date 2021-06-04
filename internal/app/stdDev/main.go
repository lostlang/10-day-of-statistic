package StdDev

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"

	"github.com/lostlang/10-day-of-statistic/pkg/statistic"
	"github.com/lostlang/10-day-of-statistic/pkg/utils"
)

func stdDev(arr []int64) (ans float64) {
	var slice statistic.StatisticSlice
	slice.SetData(arr)
	sort.Sort(slice)
	mean := slice.Mean()
	var sum float64
	for _, value := range slice.GetData() {
		sum += math.Pow(float64(value)-mean, 2)
	}
	ans = math.Sqrt(sum / float64(slice.Len()))
	return ans
}

func StdDev() {
	reader := bufio.NewReader(os.Stdin)
	utils.ReadLine(reader)
	xArr := utils.StringToArray(utils.ReadLine(reader))
	result := stdDev(xArr)
	fmt.Printf("%.1f", result)
}
