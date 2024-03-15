/*
*

	@author:
	@date : 2023/10/30
*/
package sha256Utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(str string) string {

	h := sha256.New()
	h.Write([]byte(str))

	res := hex.EncodeToString(h.Sum(nil))
	return res
}

func ShaWithKey(str, key string) string {
	str = str + key
	return Sha256(str)
}
