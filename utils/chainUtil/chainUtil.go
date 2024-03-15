/*
*

	@author:
	@date : 2023/11/9
*/
package chainUtil

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"hk_storage/common/client"
	"hk_storage/common/logger"
	"hk_storage/core/configs"
	"hk_storage/models/contract"
	"hk_storage/utils/ethutil"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"time"
)

func Mint(c1 *contract.Contract, method string, tokenId float64, uri, from, prikey string) (string, error) {
	p := contract.ABIQueryDTO{}
	//c1 := new(contract.Contract)
	//c1.Type = contractType
	//c1.GetInfo()
	p.AbiCode = c1.Abi
	p.Amount = "0"
	p.ContractAddr = c1.Address
	p.From = c1.From
	if strings.Trim(from, " ") != "" {
		p.From = from
	}
	p.PriKey = c1.PriKey
	if strings.Trim(prikey, " ") != "" {
		p.PriKey = prikey
	}
	p.Method = method
	param := make([]interface{}, 2)
	param[0] = tokenId
	param[1] = uri
	p.Params = param
	res, err := AbiControlPost(p)
	if err != nil {
		return "", err
	}
	if res.Code == 200200 {
		return res.Data["hash"].(string), nil
	} else {
		return "", errors.New(res.Message)
	}
}

func CheckBalanceEnough(address string, amount *big.Int) bool {
	balance, err := GetBalance(address)
	if err != nil {
		return false
	}
	//needInt := ethutil.ToWei(configs.TomlConfig.Chain.FileFee, 18)
	logger.Info("账户余额{},{}", balance, amount)
	if balance.Cmp(amount) <= 0 {
		return false
	}
	return true
}

func GetBalance(address string) (*big.Int, error) {

	balance, err := client.GetEthClient().BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		logger.Info("获取当前账户余额出错 address{}", address)
		return nil, err
	}
	return balance, err
}

func TransCoin(from, to, prikey string, amount *big.Int) (hash string, err error) {
	if prikey[:2] == "0x" {
		prikey = prikey[2:]
	}
	//amount 单位为ylem 需要转换一下
	am1 := ethutil.ToDecimal(amount, 18)
	param := map[string]interface{}{"fromAddress": from, "toAddress": to, "priKey": prikey,
		"amount": am1.String()}
	response, err := ContractServerPost(configs.TomlConfig.Chain.TransCoinUrl, param)
	if err != nil {
		return "", err
	}
	if response.Code == 200200 {
		return response.Data["hash"].(string), nil
	} else {
		return "", errors.New(response.Message)
	}
}

func AbiControlPost(p contract.ABIQueryDTO) (contract.ABIResponse, error) {
	result := contract.ABIResponse{}
	cli := http.Client{Timeout: 10 * time.Second}
	data, _ := json.Marshal(p)
	res, err := cli.Post(configs.TomlConfig.Chain.AbiUrl, "application/json", bytes.NewReader(data))
	if err != nil {
		logger.Info("根据高度 获取hash出错", err)
		return result, errors.New("err")
	}
	defer res.Body.Close()
	jsonData, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		logger.Info("解析数据出错", err)
		return result, err
	}
	return result, nil

}

func ContractServerPost(uri string, requestData interface{}) (contract.ABIResponse, error) {
	result := contract.ABIResponse{}
	cli := http.Client{Timeout: 10 * time.Second}
	data, _ := json.Marshal(requestData)
	res, err := cli.Post(uri, "application/json", bytes.NewReader(data))
	if err != nil {
		logger.Info("根据高度 获取hash出错", err)
		return result, errors.New("err")
	}
	defer res.Body.Close()
	jsonData, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		logger.Info("解析数据出错", err)
		return result, err
	}
	return result, nil

}
