package system

import (
	"au-golang/global"
	"au-golang/model/common/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
	"os"
	"strconv"
)

type AdministerApi struct{}

func (administerApi *AdministerApi) UploadAssLogo(c *gin.Context) {
	assIdStr := c.PostForm("assId")
	assId, _ := strconv.Atoi(assIdStr)
	ass := assService.FindByAssid(assId)
	interViewAss := interviewAssService.FindByAssId(assId)

	file, _ := c.FormFile("files")
	newName := UpLoadAssLogo(file, c)

	ass.Logo = newName
	assService.Save(&ass)
	interViewAss.Logo = newName
	interviewAssService.Save(&interViewAss)

	response.Ok(c)
}

func UpLoadAssLogo(file *multipart.FileHeader, c *gin.Context) string {
	// 创建日期
	//dataStr := time.Now().Format("2006-01-02")

	//创建时间目录
	var storagePathFather = global.GVA_CONFIG.StoragePath.RootAssLogoPath

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
	//dbStoragePath := dataStr + "/" + uuid + "_" + file.Filename
	dbStoragePath := uuid + "_" + file.Filename

	//实际物理存储地址
	storagePath := storagePathFather + "/" + uuid + "_" + file.Filename
	//// 存储在特定的文件夹里边
	err = c.SaveUploadedFile(file, storagePath)
	if err != nil {
	}
	return dbStoragePath
}

type req struct {
	AssId          int    `json:"assId"`
	AssDescription string `json:"assDescription"`
}

func (administerApi *AdministerApi) SetAssDescription(c *gin.Context) {
	var info req
	c.BindJSON(&info)

	ass := assService.FindByAssid(info.AssId)
	ass.AssDescription = info.AssDescription
	assService.Save(&ass)
	response.Ok(c)
}
