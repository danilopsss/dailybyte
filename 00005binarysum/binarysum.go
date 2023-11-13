package main

import (
	"strings"
	"strconv"
)

type IncomingBinaries struct {
	strOne string
	strTwo string
}

func (i *IncomingBinaries) GetBigAndSmallBinaries() (string, string) {
	bigger, smaller := "", ""
	if len(i.strOne) > len(i.strTwo) {
		bigger, smaller = i.strOne, i.strTwo
	} else {
		bigger, smaller = i.strTwo, i.strOne
	}
	smaller = ResizeSmallerBinary(bigger, smaller)
	return bigger, smaller
}

func ResizeSmallerBinary(bigger string, smaller string) string {
	diff := len(bigger) - len(smaller)
	return strings.Repeat("0", diff) + smaller
}

func Sum(s1 string, s2 string, hasCarry int) (int, string) {
	valueOne, _ := strconv.Atoi(s1)
	valueTwo, _ := strconv.Atoi(s2)
	summedValue := valueOne + valueTwo + hasCarry

	if summedValue == 3 {
		return 1, "1"
	}
	if summedValue == 2 {
		return 1, "0"
	}
	if summedValue == 1 {
		return 0, "1"
	}
	return 0, "0"
}


func (i *IncomingBinaries) BinarySum() string {
	BinOne, BinTwo := i.GetBigAndSmallBinaries()
	hasCarry := 0
	summedString, incomingStr := []string{}, ""
	for idx := len(BinOne) - 1; idx >= 0; idx-- {
		hasCarry, incomingStr = Sum(string(BinOne[idx]), string(BinTwo[idx]), hasCarry)
		summedString = append([]string{incomingStr}, summedString...)
	}
	if hasCarry == 1 {
		summedString = append([]string{"1"}, summedString...)
	}
	return strings.Join(summedString, "")
}
