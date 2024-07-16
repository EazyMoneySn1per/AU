package system

import (
	"au-golang/global"
	"au-golang/model/common/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

type OutlayApi struct{}

const FinalStepOutlay int = 4

// OutlaySubmit 1.财务报表上传
func (outlayApi *OutlayApi) OutlaySubmit(c *gin.Context) {
	assIdStr := c.PostForm("assid")
	assId, _ := strconv.Atoi(assIdStr)
	stepStr := c.PostForm("step")
	step, _ := strconv.Atoi(stepStr)
	name := c.PostForm("name")
	moneyStr := c.PostForm("money")
	money, _ := strconv.ParseFloat(moneyStr, 64)

	outlay, _ := outlayService.FindByAssidAndStepIsNot(assId, FinalStepOutlay)

	//判断是否是二次提交
	if outlay.OutlayUrl != "" { //删除上一次提交的文件
		preOutlayurl := outlay.OutlayUrl
		temPath := global.GVA_CONFIG.StoragePath.RootOutlayPath + "/" + preOutlayurl
		os.Remove(temPath)
	}

	// 获取文件
	file, _ := c.FormFile("files")
	outlayUrl := UpLoadOutlay(file, c)

	outlay.OutlayUrl = outlayUrl
	outlay.Assid = assId
	outlay.Name = name
	outlay.Step = step + 1
	outlay.Money = money
	outlay.Backmsg = ""

	if outlay.Id != 0 {
		outlayService.Save(&outlay)
	} else {
		outlayService.Create(&outlay)
	}
	response.OkWithMessage("ok", c)
}

func UpLoadOutlay(files *multipart.FileHeader, c *gin.Context) string {
	// 创建日期
	if files == nil {
		return ""
	}
	dataStr := time.Now().Format("2006-01-02")

	//创建时间目录
	var storagePathFather = global.GVA_CONFIG.StoragePath.RootOutlayPath + "/" + dataStr

	_, err := os.Stat(storagePathFather)
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		// 文件不存在 , 创建文件夹
		err := os.Mkdir(storagePathFather, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录异常 -> %v\n", err)
		} else {
			fmt.Println("创建成功!")
		}
	}
	uuid := uuid.NewString()

	////数据库中存储的地址
	dbStoragePath := dataStr + "/" + uuid + "_" + files.Filename
	//实际物理存储地址
	storagePath := storagePathFather + "/" + uuid + "_" + files.Filename
	//// 存储在特定的文件夹里边
	err = c.SaveUploadedFile(files, storagePath)
	if err != nil {
	}
	return dbStoragePath
}

// GetInfo 2.当社长进入页面的时候，获取进度和失败信息
func (outlayApi *OutlayApi) GetInfo(c *gin.Context) {
	assIdStr := c.Query("assid")
	assId, _ := strconv.Atoi(assIdStr)
	outlay, err := outlayService.FindByAssidAndStepIsNot(assId, FinalStepOutlay)
	if err == nil {
		response.OkWithDetailed(outlay, "ok", c)
		return
	}
	response.OkWithMessage(err.Error(), c)
}

// Nextstep 3.审核成功后下一步
func (outlayApi *OutlayApi) Nextstep(c *gin.Context) {
	assIdStr := c.Query("assid")
	assId, _ := strconv.Atoi(assIdStr)
	outlay, _ := outlayService.FindByAssidAndStepIsNot(assId, FinalStepOutlay)
	temstep := outlay.Step
	outlay.Step = temstep + 1
	outlayService.Save(&outlay)
	response.OkWithMessage("ok", c)
}

// SetBackMsg 4.审核失败的返回信息
func (outlayApi *OutlayApi) SetBackMsg(c *gin.Context) {
	assIdStr := c.Query("assid")
	assId, _ := strconv.Atoi(assIdStr)
	backmsg := c.Query("backmsg")
	outlay, _ := outlayService.FindByAssidAndStepIsNot(assId, FinalStepOutlay)
	outlay.Step = 1
	outlay.Backmsg = backmsg
	outlayService.Save(&outlay)
	response.OkWithMessage("ok", c)
}

// GetActivities 5.获取等待审核列表
func (outlayApi *OutlayApi) GetActivities(c *gin.Context) {
	outlay := outlayService.FindAllByStepNot(FinalStepTwitter)
	response.OkWithDetailed(outlay, "ok", c)
}

// DownFile 6.文件下载
func (outlayApi *OutlayApi) DownFile(c *gin.Context) {
	path := c.Query("path")
	resPath := global.GVA_CONFIG.StoragePath.RootOutlayPath + "/" + path
	_, errByOpenFile := os.Open(resPath)
	arr := strings.Split(path, "_")
	if errByOpenFile != nil {
		response.FailWithMessage("fail", c)
		return
	}
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", "attachment; filename="+arr[1])
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(resPath)
	return
}

func (o *OutlayApi) GetRecord(c *gin.Context) {

	page := c.Query("page")
	limit := c.Query("limit")
	assId, exist := c.GetQuery("assId")
	if !exist {
		response.FailWithMessage("传参出错", c)
		return
	}
	outlay, count := outlayService.GetRecord(cast.ToInt(assId), cast.ToInt(page), cast.ToInt(limit))

	response.OkWithDetailed(gin.H{
		"data":  outlay,
		"count": count,
	}, "ok", c)
}
