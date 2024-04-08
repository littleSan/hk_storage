/*
*

	@author:
	@date : 2023/10/8
*/
package ipfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hk_storage/common/logger"
	"hk_storage/core/configs"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type storageResponse struct {
	Ok    bool                   `json:"ok"`
	Value map[string]interface{} `json:"value"`
}
type storageFile struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (s *service) Upload(header *multipart.FileHeader) (rest string, err error) {
	file, err := header.Open()
	defer file.Close()
	if err != nil {
		logger.Info("上传IPFS文件，打开文件出错")
		return "", err
	}
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// Create a file form
	part, err := writer.CreateFormFile("file", header.Filename) // replace with your field name and file name
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Close the writer and get the content type based on the writer's headers
	writer.Close()
	fmt.Println("数据大小", body.Len())
	contentType := writer.FormDataContentType()
	req, err := http.NewRequest("POST", configs.TomlConfig.Ipfs.UploadUrl, body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", configs.TomlConfig.Ipfs.Authorization)

	client := &http.Client{
		//Transport: &transport,
		Timeout: 0,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 处理响应
	fmt.Println(resp.Status)
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println("response data{}", string(data))
	if err != nil {
		return "", err
	}
	response := storageResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		logger.Info("解析IPfs 返回数据出错", err)
		return string(data), err
	}
	cid := fmt.Sprintf("%v", response.Value["cid"])
	files := []storageFile{}
	//fileStr := fmt.Sprintf("%v", response.Value["files"])
	data, err = json.Marshal(response.Value["files"])
	err = json.Unmarshal(data, &files)
	if err != nil {
		logger.Info("解析file 出错", err)
	}
	var filename string
	for _, s2 := range files {
		logger.Info("ipfs 文件上传名称", s2.Name)
		filename = s2.Name
	}
	if strings.Trim(filename, " ") == "" {
		filename = header.Filename
	}
	rest = configs.TomlConfig.Ipfs.IpfsUrl + cid + "/" + filename
	return rest, err
}

func (s *service) UploadJson(obj interface{}, fileName string) (rest string, err error) {
	if strings.Trim(fileName, " ") == "" {
		fileName = strconv.FormatInt(time.Now().UnixMilli(), 10)
	}
	if !strings.Contains(strings.ToLower(fileName), ".json") {
		fileName = fmt.Sprint(fileName + ".json")
	}
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// Create a file form
	part, err := writer.CreateFormFile("file", fileName) // replace with your field name and file name
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonData, _ := json.Marshal(obj)
	part.Write(jsonData)
	writer.Close()
	contentType := writer.FormDataContentType()
	req, err := http.NewRequest("POST", configs.TomlConfig.Ipfs.UploadUrl, body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", configs.TomlConfig.Ipfs.Authorization)
	// 发送请求并获取响应
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 处理响应
	fmt.Println(resp.Status)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	response := storageResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		logger.Info("解析IPfs 返回数据出错", err)
		return string(data), err
	}
	cid := fmt.Sprintf("%v", response.Value["cid"])
	files := []storageFile{}
	//fileStr := fmt.Sprintf("%v", response.Value["files"])
	data, err = json.Marshal(response.Value["files"])
	err = json.Unmarshal(data, &files)
	if err != nil {
		logger.Info("解析file 出错", err)
		return "", err
	}

	for _, s2 := range files {
		logger.Info("ipfs 文件上传名称", s2.Name)
		fileName = s2.Name
	}
	rest = configs.TomlConfig.Ipfs.IpfsUrl + cid + "/" + fileName
	return rest, err

}
