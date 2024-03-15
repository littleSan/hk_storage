/*
*

	@author:
	@date : 2023/11/7
*/
package randomUtils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GetRandom(base, height int) int {
	//设置种子
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(height) + base
}

// 100 以内的随机数
func GetDefaultRandom() int {
	return GetRandom(1, 1000)
}

func GetRandomIn1000() int {
	return GetRandom(1, 1000)
}

func GenerateRandomNumber(n int) int64 {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 生成N位随机数
	randomNumber := rand.Int63n(10)                  // 生成0-9之间的随机数作为第一位
	randomNumber = randomNumber*10 + rand.Int63n(10) // 生成0-9之间的随机数作为第二位
	// 以此类推，重复n-1次，生成n位随机数
	for i := 2; i < n; i++ {
		randomNumber = randomNumber*10 + rand.Int63n(10)
	}

	return randomNumber
}

func GenerateDefaultRandomNumber() int64 {
	return GenerateRandomNumber(13)
}

func GetRandomCode(length int) string {
	numeric := [62]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S',
		'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w',
		'x', 'y', 'z'}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < length; i++ {
		fmt.Fprintf(&sb, "%c", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func GetRandomDefaultCode() string {
	return GetRandomCode(15)
}

//func main() {
//	fmt.Print(GetDefaultRandom())
//}
