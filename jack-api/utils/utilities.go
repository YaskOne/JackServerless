package utils

import (
	"strings"
	"strconv"
)

func SplitArrayString(stringIds string) (res []uint) {
	stringIds = strings.Replace(stringIds, "[", "", -1)
	stringIds = strings.Replace(stringIds, "]", "", -1)

	stringArray := strings.Split(stringIds, ",")

	var intArray = []uint{}

	for _, i := range stringArray {
		value, err := strconv.ParseUint(i, 10, 64)
		if err != nil {
			panic(err)
			return nil
		}
		intArray = append(intArray, uint(value))
	}
	return intArray
}