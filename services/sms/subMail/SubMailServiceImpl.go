package subMail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hk_storage/common/logger"
	"hk_storage/core/configs"
	"io/ioutil"
	"net/http"
	"time"
)

type SmsResponse struct {
	Status                  string `json:"status,omitempty"`
	sendId                  string `json:"send_id,omitempty"`
	Fee                     int    `json:"fee,omitempty"`
	smsCredits              string `json:"sms_credits,omitempty"`
	transactionalSmsCredits string `json:"transactional_sms_credits,omitempty"`
	Msg                     string `json:"msg"`
}

type request struct {
	To        string `json:"to,omitempty"`
	Content   string `json:"content,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Appid     int    `json:"appid,omitempty"`
	Signature string `json:"signature,omitempty"`
}

func (i *service) Send(phone, content string) bool {
	req := request{}
	req.Appid = configs.TomlConfig.SmsModel["subMail"].Id
	req.Signature = configs.TomlConfig.SmsModel["subMail"].Key
	req.To = phone
	req.Content = content
	req.Timestamp = time.Now().Unix()

	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}
	data1, err := json.Marshal(req)
	if err != nil {
		logger.Info("发送短信，转换数据出错", err)
		return false
	}

	resp, err := httpClient.Post(configs.TomlConfig.SmsModel["subMail"].EndPoint, "application/json", bytes.NewReader(data1))
	if err != nil {
		fmt.Println("error opening json file", err)
		return false
	}
	defer resp.Body.Close()
	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error opening json file", err)
		return false
	}
	response := SmsResponse{}
	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return false
	}
	if response.Status != "success" {
		logger.Info("发送短信失败", response.Msg)
		return false
	}
	return true
}

func (i *service) SendGlobal(phone, content string) bool {
	req := request{}
	req.Appid = configs.TomlConfig.SmsModel["subMailGlobal"].Id
	req.Signature = configs.TomlConfig.SmsModel["subMailGlobal"].Key
	req.To = phone
	req.Content = content
	req.Timestamp = time.Now().Unix()
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}
	data1, err := json.Marshal(req)
	if err != nil {
		logger.Info("发送短信，转换数据出错", err)
		return false
	}

	resp, err := httpClient.Post(configs.TomlConfig.SmsModel["subMailGlobal"].EndPoint, "application/json", bytes.NewReader(data1))
	if err != nil {
		fmt.Println("error opening json file", err)
		return false
	}
	defer resp.Body.Close()
	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error opening json file", err)
		return false
	}
	response := SmsResponse{}
	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return false
	}
	if response.Status != "success" {
		logger.Info("发送短信失败", response.Msg)
		return false
	}
	return true
}
