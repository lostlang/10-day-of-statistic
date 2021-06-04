package weightedMean

import (
	"bufio"
	"fmt"
	"os"

	"github.com/lostlang/10-day-of-statistic/pkg/utils"
)

func weightedMean(x []int64, w []int64) float64 {
	var mean int64
	var sumW int64
	for i := range x {
		mean += x[i] * w[i]
		sumW += w[i]
	}
	ans := float64(mean) / float64(sumW)
	return ans
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	utils.ReadLine(reader)
	xArr := utils.StringToArray(utils.ReadLine(reader))
	wArr := utils.StringToArray(utils.ReadLine(reader))
	result := weightedMean(xArr, wArr)
	fmt.Printf("%f", result)
}
