/*
*

	@author:
	@date : 2023/10/8
*/
package aliOss

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"hk_storage/common/logger"
	"hk_storage/core/configs"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"
)

func (s *service) Upload(path string, file *multipart.FileHeader) (rest string, err error) {
	var uploadFileName = path + "/"
	isExist, err := s.Oss.IsBucketExist(configs.TomlConfig.AliOss.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		return "", errors.New("bucket err")
	}
	logger.Info("IsBucketExist result : ", isExist)
	//文件上传，文件上传有简单上传，追加上传，断点续传上传，分片上传
	if !isExist {
		return "", errors.New("bucket not exist")
	}
	//注意此处不要写错，写错的话，err让然是nil，我们应该需要先判断一下是否存在
	bucket, _ := s.Oss.Bucket(configs.TomlConfig.AliOss.Bucket)
	source, err := file.Open()
	if err != nil {
		logger.Info("open file err", err)
		return "", errors.New("open file err")
	}
	defer source.Close()
	newName := getHash(file.Filename + time.Now().String())
	uploadFileName = uploadFileName + newName + filepath.Ext(file.Filename)
	err = bucket.PutObject(uploadFileName, source)
	if err != nil {
		logger.Info("Error:", err)
		logger.Info("path", uploadFileName)
		return "", errors.New("upload file err")
	}
	return getPath(uploadFileName), err
}

func (s *service) UploadUrl(path string, url string) (rest string, err error) {
	var uploadFileName = path + "/"
	isExist, err := s.Oss.IsBucketExist(configs.TomlConfig.AliOss.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		return "", errors.New("bucket err")
	}
	logger.Info("IsBucketExist result : ", isExist)
	//文件上传，文件上传有简单上传，追加上传，断点续传上传，分片上传
	if !isExist {
		return "", errors.New("bucket not exist")
	}
	//注意此处不要写错，写错的话，err让然是nil，我们应该需要先判断一下是否存在
	bucket, _ := s.Oss.Bucket(configs.TomlConfig.AliOss.Bucket)
	//获取URL文件
	resp, err := http.Get(url) // 从URL获取图片内容
	if err != nil {
		fmt.Println("Failed to retrieve image from URL:", err)
		return
	}
	defer resp.Body.Close()
	nowName := getHash(url + time.Now().String())
	uploadFileName = uploadFileName + nowName + filepath.Ext(url)
	bucket.PutObject(uploadFileName, resp.Body)
	return getPath(uploadFileName), nil

}

func getHash(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

func getPath(path string) string {
	return "https://" + configs.TomlConfig.AliOss.Bucket + "." + configs.TomlConfig.AliOss.Endpoint + "/" + path
}
