package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Got from HackerRank
// //Given an array of positive integers, return the number of elements that are strictly
// greater than the average of all previous elements. Skip the first element.

// responseTimes = [100, 200, 150,300]
// - Day 0: 100 (no previous days, skip)
// - Day 1: 200 > average(100) = 100 → count = 1
// - Day 2: 150 vs average(100, 200) = 150 → not greater → count = 1
// - Day 3: 300 > average(100, 200, 150) = 150 → count = 2 Return 2.

func countResponseTimeRegressions(responseTimes []int32) int32 {
	if len(responseTimes) < 2 {
		return 0
	}
	var count int32 = 0
	var sum int64 = int64(responseTimes[0])

	for i, val := range responseTimes {
		if i == 0 {
			continue
		}
		if int64(val)*int64(i) > int64(sum) {
			count++
		}
		sum = int64(sum) + int64(val)
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	responseTimesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var responseTimes []int32

	for i := 0; i < int(responseTimesCount); i++ {
		responseTimesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		responseTimesItem := int32(responseTimesItemTemp)
		responseTimes = append(responseTimes, responseTimesItem)
	}

	result := countResponseTimeRegressions(responseTimes)

	fmt.Printf("%d\n", result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
