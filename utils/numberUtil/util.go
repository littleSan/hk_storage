package numberUtil

import (
	"math/big"
	"regexp"
	"strings"
)

func IsNumber(str string) bool {
	pattern := "^[0-9]+$"
	match, err := regexp.MatchString(pattern, str)
	if err != nil {
		return false
	}
	return match
}

func FloatToBigInt(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)

	result := new(big.Int)
	bigval.Int(result) // store converted number in result
	return result
}

func StringToSlice(str string) (res []string) {
	str = strings.ReplaceAll(str, "[", "")
	str = strings.ReplaceAll(str, "]", "")
	res = strings.Split(str, ",")
	return
}
