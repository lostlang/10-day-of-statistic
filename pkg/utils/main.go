package utils

import (
	"bufio"
	"strconv"
	"strings"
)

func ReadLine(reader *bufio.Reader) string {
	bArr, _, _ := reader.ReadLine()
	return string(bArr)
}

func ToInt(str string) int64 {
	const system = 10
	const intDigit = 64
	ans, _ := strconv.ParseInt(str, system, intDigit)
	return ans
}

func StringToArray(str string) []int64 {
	tmp := strings.Split(str, " ")
	intArray := make([]int64, len(tmp))
	for i := range tmp {
		intArray[i] = ToInt(tmp[i])
	}
	return intArray
}
