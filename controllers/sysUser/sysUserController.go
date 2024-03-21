/*
*

	@author:
	@date : 2023/9/27
*/
package sysUser

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"hk_storage/common/logger"
	"hk_storage/common/pages"
	"hk_storage/common/redisCli"
	"hk_storage/common/response"
	"hk_storage/core/configs"
	"hk_storage/models/studyFile"
	"hk_storage/services/StudyFile"
	"hk_storage/services/ipfs"
	"hk_storage/services/sysUser"
	"hk_storage/utils/chainUtil"
	"hk_storage/utils/ethutil"
	"hk_storage/utils/randomUtils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var _ Controller = (*controller)(nil)

type Controller interface {
	i()
	Login(ctx *gin.Context)
	UserFile(ctx *gin.Context)
	SaveUserFile(ctx *gin.Context)
	DeleteUserFile(ctx *gin.Context)
	DrawYlem(ctx *gin.Context)
	WaterTap(ctx *gin.Context)
	GetDrawSurplusQuantity(ctx *gin.Context)
}
type LoginVo struct {
	Account string `json:"account,omitempty" form:"account"`
	Address string `json:"address,omitempty" form:"password"`
	Uid     string `json:"uid,omitempty" form:"code"`
}

type controller struct {
	SysUser      sysUser.Service
	IpfsService  ipfs.Service
	StudyService StudyFile.Service
	RedisCli     redisCli.Repo
}

func New() *controller {
	return &controller{
		SysUser:      sysUser.New(),
		IpfsService:  ipfs.New(),
		StudyService: StudyFile.New(),
		RedisCli:     redisCli.New(),
	}
}

func (c *controller) i() {
}

func (c *controller) Login(ctx *gin.Context) {
	infos := &LoginVo{}
	err := ctx.BindJSON(infos)
	if err != nil {
		logger.Info("登陆参数错误", err)
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.ParamError))
		return
	}
	if infos.Account == "" {
		logger.Info("用户登陆 用户名不能为空")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.AccountError))
		return
	}
	cc, err := c.SysUser.GetUserInfoByUid(infos.Uid)
	if err != nil {
		logger.Info("未查询到用户信息")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.AccountError))
		return
	}
	ctx.JSON(http.StatusOK, response.SUCCESS(ctx, cc))
	return
}

func (c *controller) UserFile(ctx *gin.Context) {
	pagination := pages.GeneratePaginationFromRequest(ctx)
	infos := &studyFile.StudyFile{}
	err := ctx.BindQuery(infos)
	if err != nil {
		logger.Info("参数错误", err)
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.ParamError))
		return
	}
	if infos.Uid == "" {
		logger.Info("账户UID不能为空")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.AccountError))
		return
	}

	cc, err := c.SysUser.GetUserInfoByUid(infos.Uid)
	if err != nil {
		logger.Info("未查询到用户信息")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.AccountError))
		return
	}
	//获取用户上传数据
	infos.Uid = cc.Uid
	res, _ := c.StudyService.List(*infos, &pagination)
	pagination.InitPage(res)
	ctx.JSON(http.StatusOK, response.SUCCESS(ctx, pagination))
	return
}

func (c *controller) SaveUserFile(ctx *gin.Context) {
	infos := &studyFile.StudyFile{}
	err := ctx.BindJSON(infos)
	if err != nil {
		logger.Info("参数错误", err)
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.ParamError))
		return
	}
	if infos.Uid == "" {
		logger.Info("账户UID不能为空")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.AccountError))
		return
	}

	if len(infos.Url) == 0 {
		logger.Info("链接地址不允许为空")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.ParamError))
		return
	}
	var v1 []string
	err = json.Unmarshal([]byte(infos.Url), &v1)
	if err != nil {
		logger.Info("err", err)
	}
	if infos.Amount != len(v1) {
		logger.Info("上链文件数量错误")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.FileAmountErr))
		return
	}
	cc, err := c.SysUser.GetUserInfoByUid(infos.Uid)
	if err != nil {
		logger.Info("未查询到用户信息")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.AccountError))
		return
	}
	infos.Address = cc.Address
	infos.Status = studyFile.StatusToChain
	infos.Token = randomUtils.GenerateDefaultRandomNumber()
	c.StudyService.Save(infos)
	//获取用户上传数据,生成json数据，保存记录
	go func() {
		json := map[string]interface{}{"name": infos.Name, "image": infos.Url, "description": infos.Description,
			"attributes": []map[string]interface{}{{"trait_type": "owner", "value": infos.Address}}}
		jsonStr, _ := c.IpfsService.UploadJson(json, strconv.Itoa(int(time.Now().Unix())))
		infos.JsonUrl = jsonStr

		//c1 := &contract.Contract{}
		//c1.Type = "storage"
		//c1.GetInfo()
		//hash, err := chainUtil.Mint(c1, "mint", float64(infos.Token), jsonStr, cc.Address, cc.PriKey)
		//if err != nil {
		//	logger.Info("上链操作失败", err)
		//	ctx.JSON(http.StatusOK, response.Failure(ctx, response.Fail))
		//	return
		//}
		//infos.Status = studyFile.StatusChaining
		//infos.ContractAddr = c1.Address
		//infos.Hash = hash
		c.StudyService.Update(*infos)
		//sendMq.SendMsg(queueConst.ChainQueue, hash)
	}()
	ctx.JSON(http.StatusOK, response.SUCCESS(ctx, infos))
	return
}

func (c *controller) DeleteUserFile(ctx *gin.Context) {
	infos := &studyFile.StudyFile{}
	err := ctx.BindJSON(infos)
	if err != nil {
		logger.Info("参数错误", err)
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.ParamError))
		return
	}
	if infos.Uid == "" {
		logger.Info("账户UID不能为空")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.AccountError))
		return
	}

	cc, err := c.SysUser.GetUserInfoByUid(infos.Uid)
	if err != nil {
		logger.Info("未查询到用户信息")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.AccountError))
		return
	}
	files, err := c.StudyService.GetFileByToken(infos.Token)
	if err != nil || files == nil {
		logger.Info("未查询到上传数据")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.FileNotExist))
		return
	}
	if strings.Compare(cc.Address, files.Address) != 0 {
		logger.Info("文件地址与用户不匹配")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.AddressFormatErr))
		return
	}
	if files.Status != studyFile.StatusToChain {
		logger.Info("文件状态值错误")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.FileStatusErr))
		return
	}
	files.Status = studyFile.StatusDelete
	c.StudyService.Update(*files)
	ctx.JSON(http.StatusOK, response.SUCCESS(ctx, files))
	return
}

func (c *controller) DrawYlem(ctx *gin.Context) {
	infos := &LoginVo{}
	err := ctx.BindJSON(infos)
	if err != nil {
		logger.Info("登陆参数错误", err)
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.ParamError))
		return
	}
	cc, err := c.SysUser.GetUserInfoByUid(infos.Uid)
	if err != nil {
		logger.Info("未查询到用户信息")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.AccountError))
		return
	}
	if len(cc.Address) == 42 && cc.Address[:2] == "0x" {

	} else {
		logger.Info("地址错误")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.AddressFormatErr))
		return
	}
	//通过redis 控制一天限制领取次数
	b1, _ := c.RedisCli.LimitCount(cc.Uid, 20)
	if !b1 {
		logger.Info("积分已领取")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.YLemMaxLimitErr))
		return
	}
	//生成转让金额
	rand := randomUtils.GetRandomIn1000()
	am1 := ethutil.ToWei(decimal.NewFromInt(int64(rand)), 15)
	logger.Info("打款金额{}", am1)
	hash, err := chainUtil.TransCoin(configs.TomlConfig.Chain.BaseAddress, cc.Address, configs.TomlConfig.Chain.BaseAddressKey, am1)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.Fail))
		return
	}
	ctx.JSON(http.StatusOK, response.SUCCESS(ctx, hash))
	return
}
func (c *controller) GetDrawSurplusQuantity(ctx *gin.Context) {
	infos := &LoginVo{}
	err := ctx.BindJSON(infos)
	if err != nil {
		logger.Info("登陆参数错误", err)
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.ParamError))
		return
	}
	cc, err := c.SysUser.GetUserInfoByUid(infos.Uid)
	if err != nil {
		logger.Info("未查询到用户信息")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.AccountError))
		return
	}
	if len(cc.Address) == 42 && cc.Address[:2] == "0x" {

	} else {
		logger.Info("地址错误")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.AddressFormatErr))
		return
	}
	//通过redis 控制一天限制领取次数
	b1, _ := c.RedisCli.LimitAmount(cc.Uid)
	ctx.JSON(http.StatusOK, response.SUCCESS(ctx, b1))
	return
}

func (c *controller) WaterTap(ctx *gin.Context) {
	infos := &LoginVo{}
	err := ctx.BindJSON(infos)
	if err != nil {
		logger.Info("登陆参数错误", err)
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.ParamError))
		return
	}
	cc, err := c.SysUser.GetUserInfoByUid(infos.Uid)
	if err != nil {
		logger.Info("未查询到用户信息")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.AccountError))
		return
	}

	if len(cc.Address) == 42 && cc.Address[:2] == "0x" {

	} else {
		logger.Info("地址错误")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.AddressFormatErr))
		return
	}
	//通过redis 控制一天只能领一次
	b1, _ := c.RedisCli.LimitCount(cc.Address, 1)
	if !b1 {
		logger.Info("积分已领取")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.YLemReceived))
		return
	}

	//生成转让金额
	rand := randomUtils.GetRandomIn1000()
	am1 := ethutil.ToWei(decimal.NewFromInt(int64(rand)), 15)
	logger.Info("打款金额{}", am1)
	hash, err := chainUtil.TransCoin(configs.TomlConfig.Chain.BaseAddress, cc.Address, configs.TomlConfig.Chain.BaseAddressKey, am1)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.Fail))
		return
	}
	ctx.JSON(http.StatusOK, response.SUCCESS(ctx, hash))
	return
}
