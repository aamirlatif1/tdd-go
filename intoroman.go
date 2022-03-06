package tdd_go

import "strings"

func intToRoman(num int) string {
	var roman strings.Builder
	romanVals := []byte{'M', 'D', 'C', 'L', 'X', 'V', 'I'}
	numVals := []int{1000, 500, 100, 50, 10, 5, 1}
	n := num
	index := 0
	for n > 0 {
		if n >= numVals[index] {
			roman.WriteByte(romanVals[index])
			n -= numVals[index]
		} else if index%2 == 0 && n >= numVals[index]-numVals[index+2] {
			roman.WriteByte(romanVals[index+2])
			roman.WriteByte(romanVals[index])
			n -= numVals[index] - numVals[index+2]
		} else if index%2 == 1 && n >= numVals[index]-numVals[index+1] {
			roman.WriteByte(romanVals[index+1])
			roman.WriteByte(romanVals[index])
			n -= numVals[index] - numVals[index+1]
		} else {
			index++
		}
	}
	return roman.String()
}
