package system

import (
	"au-golang/global"
	"au-golang/model/common/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"math/rand"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

type TwitterApi struct{}

const FinalStepTwitter int = 4

var rootPicturePath = global.GVA_CONFIG.StoragePath.RootPicturePath

// 推文图片上传+请求提交
func (twitterApi *TwitterApi) TwitterSubmit(c *gin.Context) {

	form, _ := c.MultipartForm()
	files := form.File["files"]
	picUrl := UpLoadPicture(files, c)
	assIdStr := c.PostForm("assid")
	assId, _ := strconv.Atoi(assIdStr)
	name := c.PostForm("name")
	stepStr := c.PostForm("step")
	step, _ := strconv.Atoi(stepStr)

	twitter, _ := twitterService.FindByAssidAndStepIsNot(assId, FinalStepTwitter)

	rootPicturePath := global.GVA_CONFIG.StoragePath.RootPicturePath

	// 判断是否是二次提交
	if string(twitter.PictureUrl) != "" { //删除上一次提交的文件
		prePicUrl := strings.Split(string(twitter.PictureUrl), "`")
		for i := 1; i < len(prePicUrl); i++ {
			temPath := rootPicturePath + "/" + prePicUrl[i]
			os.Remove(temPath)
		}
	}
	twitter.PictureUrl = picUrl
	twitter.Assid = assId
	twitter.Name = name
	twitter.Step = step + 1
	twitter.Backmsg = ""
	if twitter.Id != 0 {
		twitter.UpdatedAt = time.Now()
		twitterService.UpdateInfo(&twitter)
	} else {
		rand.Seed(time.Now().Unix())
		Id := rand.Int31()
		twitter.Id = int(Id)
		twitter.CreatedAt = time.Now()
		twitterService.Save(&twitter)
	}
	response.OkWithMessage("ok", c)
}

// 将图片存入本地服务器
func UpLoadPicture(files []*multipart.FileHeader, c *gin.Context) string {
	splicePath := ""
	for _, file := range files {
		// 使用uuid工具
		uuid := strings.Replace(uuid.NewString(), "", "-", -1)
		// 创建日期
		//dataStr := time.Now().Format("2006-01-02")
		//创建时间目录
		//var storagePathFather = rootPicturePath + "/" + dataStr
		var storagePathFather = global.GVA_CONFIG.StoragePath.RootPicturePath
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
		//数据库中存储的地址
		//dbStoragePath := dataStr + "/" + uuid + "_" + file.Filename
		dbStoragePath := uuid + "_" + file.Filename
		//实际物理存储地址
		storagePath := storagePathFather + "/" + dbStoragePath
		// 存储在特定的文件夹里边
		c.SaveUploadedFile(file, storagePath)
		//将所有文件的地址拼接成一个地址
		if splicePath != "" {
			splicePath = splicePath + "`" + dbStoragePath
		} else {
			splicePath = dbStoragePath
		}
	}
	return splicePath
}

func (twitterApi *TwitterApi) DownFile(c *gin.Context) {
	path := c.Query("path")
	if strings.Contains(path, "../") || strings.Contains(path, "./") {
		response.FailWithMessage("上传路径不正确", c)
		return
	}
	rootPicturePath := global.GVA_CONFIG.StoragePath.RootPicturePath
	if path != "" {
		fileName := rootPicturePath + "/" + path
		file, _ := os.Open(fileName)
		fileInfo, _ := file.Stat()
		fileSize := fileInfo.Size()
		buffer := make([]byte, fileSize)
		_, err := file.Read(buffer)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithDetailed(buffer, "ok", c)
		return
	}
}

// GetInfo 2.当社长进入页面的时候，获取推文进度和失败信息
func (twitterApi *TwitterApi) GetInfo(c *gin.Context) {
	assIdStr := c.Query("assid")
	assId, _ := strconv.Atoi(assIdStr)
	twitter, err := twitterService.FindByAssidAndStepIsNot(assId, FinalStepTwitter)
	if err == gorm.ErrRecordNotFound {
	}
	if twitter.Assid == 0 {
		response.OkWithDetailed(nil, "ok", c)
		return
	}
	response.OkWithDetailed(twitter, "ok", c)
}

// Nextstep 3.审核成功后下一步
func (twitterApi *TwitterApi) Nextstep(c *gin.Context) {
	assIdStr := c.Query("assid")
	assId, _ := strconv.Atoi(assIdStr)
	twitter, _ := twitterService.FindByAssidAndStepIsNot(assId, FinalStepTwitter)
	temstep := twitter.Step
	twitter.Step = temstep + 1

	twitterService.UpdateInfo(&twitter)
	response.OkWithMessage("ok", c)
}

// SetBackMsg 4.审核失败的返回信息
func (twitterApi *TwitterApi) SetBackMsg(c *gin.Context) {
	assIdStr := c.Query("assid")
	assId, _ := strconv.Atoi(assIdStr)
	backmsg := c.Query("backmsg")
	twitter, _ := twitterService.FindByAssidAndStepIsNot(assId, FinalStepTwitter)
	twitter.Step = 1
	twitter.Backmsg = backmsg
	twitterService.UpdateInfo(&twitter)
	response.OkWithMessage("ok", c)
}

// GetActivities 5.获取等待审核列表
func (twitterApi *TwitterApi) GetActivities(c *gin.Context) {
	twitters := twitterService.FindAllByStepNot(FinalStepTwitter)
	response.OkWithDetailed(twitters, "ok", c)
}

func (t *TwitterApi) GetRecord(c *gin.Context) {

	page := c.Query("page")
	limit := c.Query("limit")
	assId, exist := c.GetQuery("assId")
	if !exist {
		response.FailWithMessage("传参出错", c)
		return
	}
	twitter, count := twitterService.GetRecord(cast.ToInt(assId), cast.ToInt(page), cast.ToInt(limit))

	response.OkWithDetailed(gin.H{
		"data":  twitter,
		"count": count,
	}, "ok", c)
}
