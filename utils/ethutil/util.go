package ethutil

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"hk_storage/common/logger"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
)

// IsValidAddress validate hex address 是否标准地址
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// 是否是零地址
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}
	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

// 将小数转换为wei(整数）。 第二个参数是小数位数。
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)
	wei := new(big.Int)
	wei.SetString(result.String(), 10)
	return wei
}

// 将wei（整数）转换为小数。 第二个参数是小数位数
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)
	return result
}

// 从签名中提取R，S和V值。
func SigRSV(isig interface{}) ([32]byte, [32]byte, uint8) {
	var sig []byte
	switch v := isig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}
	sigstr := common.Bytes2Hex(sig)
	rS := sigstr[0:64]
	sS := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigstr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI + 27)
	return R, S, V
}

func CreateWallet() (address common.Address, priKey string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		logger.Info("生成钱包私钥失败", err)
	}
	//我们可以通过导入golang crypto/ecdsa 包并使用 FromECDSA 方法将其
	//转换为字节。
	privateKeyBytes := crypto.FromECDSA(privateKey)
	//我们现在可以使用go-ethereum hexutil 包将它转换为十六进制字符串，该
	//包提供了一个带有字节切片的 Encode 方法。 然后我们在十六进制编码之
	//后删除“0x”。
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) //用于签署交易的私钥
	//于公钥是从私钥派生的，因此go-ethereum的加密私钥具有一个返回公
	//钥的 Public 方法
	publicKey := privateKey.Public()
	//将其转换为十六进制的过程与我们使用转化私钥的过程类似。 我们剥离
	//了 0x 和前2个字符 04 ，它始终是EC前缀，不是必需的
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.Error("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283f
	//我们拥有公钥，就可以轻松生成你经常看到的公共地址。 为了做到
	//这一点，go-ethereum加密包有一个 PubkeyToAddress 方法，它接受一个
	//ECDSA公钥，并返回公共地址
	addressHex := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(addressHex) // 0x96216849c49358B10257cb55b28eA603c874b05E
	//公共地址其实就是公钥的Keccak-256哈希，然后我们取最后40个字符
	//（20个字节）并用“0x”作为前缀。

	address = common.HexToAddress(addressHex)
	priKey = hexutil.Encode(privateKeyBytes)[2:]
	return address, priKey
}
