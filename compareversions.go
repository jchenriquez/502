package compareversions

import (
	"strconv"
	"strings"
)

func checkIfGT1(splitted []string, curr int) (ret int) {

	for curr < len(splitted) {
		v, _ := strconv.Atoi(splitted[curr])

		if v > 0 {
			ret = 1
		}
		curr++
	}

	return
}

// CompareVersion will take two versioning strings and compare them, will return -1, 0, 1
func CompareVersion(version1 string, version2 string) int {
	v1Splitted := strings.Split(version1, ".")
	v2Splitted := strings.Split(version2, ".")

	var v1Curr int
	var v2Curr int

	for v1Curr < len(v1Splitted) && v2Curr < len(v2Splitted) {

		v1, _ := strconv.Atoi(v1Splitted[v1Curr])
		v2, _ := strconv.Atoi(v2Splitted[v2Curr])

		if v1 > v2 {
			return 1
		}

		if v1 < v2 {
			return -1
		}

		v1Curr++
		v2Curr++
	}

	gt1 := checkIfGT1(v1Splitted, v1Curr)

	if gt1 != 0 {
		return 1
	}

	gt1 = checkIfGT1(v2Splitted, v2Curr)

	if gt1 != 0 {
		return -1
	}

	return 0
}
