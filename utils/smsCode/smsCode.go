/*
*

	@author:
	@date : 2023/9/28
*/
package smsCode

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func generateRandomNumber(length int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < length; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func GetSmsCode() string {
	return generateRandomNumber(6)
}
