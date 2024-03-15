/*
*

	@author:
	@date : 2023/9/27
*/
package rc4Utils

import (
	"crypto/rc4"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
)

// 解密
func DecryptRc4(data, key string) (temp []byte, err error) {
	// decrypt process
	//temp, err = json.Marshal(data)
	temp, err = hex.DecodeString(data)
	if err != nil {
		return temp, errors.New("decrypt Data err")
	}

	c3, err := rc4.NewCipher([]byte(key))
	if err != nil {
		return temp, errors.New("error^26^new cipher key")
	}
	c3.XORKeyStream(temp, temp)
	return temp, err
}

// 加密
func EncryptRc4(EncryptData interface{}, key string) (temp string, err error) {
	//temp, err = hex.DecodeString(EncryptData)
	data, err := json.Marshal(EncryptData)
	if err != nil {
		return temp, errors.New("decrypt Data err")
	}
	c1, err := rc4.NewCipher([]byte(key))
	if err != nil {
		return temp, errors.New("salt err")
	}
	c1.XORKeyStream(data, data)
	temp = fmt.Sprintf("%x", data)
	return temp, err
}
