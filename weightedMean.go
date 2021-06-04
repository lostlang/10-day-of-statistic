package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLine(reader *bufio.Reader) string {
	bArr, _, _ := reader.ReadLine()
	return string(bArr)
}

func toInt(str string) int64 {
	const system = 10
	const intDigit = 64
	ans, _ := strconv.ParseInt(str, system, intDigit)
	return ans
}

func stringToArray(str string) []int64 {
	tmp := strings.Split(str, " ")
	intArray := make([]int64, len(tmp))
	for i := range tmp {
		intArray[i] = toInt(tmp[i])
	}
	return intArray
}

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
	readLine(reader)
	xArr := stringToArray(readLine(reader))
	wArr := stringToArray(readLine(reader))
	result := weightedMean(xArr, wArr)
	fmt.Printf("%f",result)
}
