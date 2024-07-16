package system

import (
	"au-golang/global"
	"au-golang/model/common/response"
	"au-golang/model/system"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

type SynthesizeApi struct{}

func (SynthesizeApi *SynthesizeApi) SynthesizeSubmit(c *gin.Context) {
	assidStr := c.PostForm("assid")
	assId, _ := strconv.Atoi(assidStr)
	typeStr := c.PostForm("type")
	_type, _ := strconv.Atoi(typeStr)
	id := uuid.NewString()
	synthesizeSubmit := system.SynthesizeSubmit{
		Id: id,
		//Ass:         assService.FindByAssid(assId),
		AssEntity:   assId,
		Description: c.PostForm("description"),
		Name:        c.PostForm("name"),
		Type:        _type,
	}
	file, _ := c.FormFile("files")
	fileUrl := UpLoadFile(file, c)
	synthesizeSubmit.FileUrl = fileUrl
	synthesizeSubmitService.Create(&synthesizeSubmit)
	response.Ok(c)
}

func UpLoadFile(files *multipart.FileHeader, c *gin.Context) string {
	// 创建日期
	dataStr := time.Now().Format("2006-01-02")

	//创建时间目录
	var storagePathFather = global.GVA_CONFIG.StoragePath.RootSynthesizePath + "/" + dataStr

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

func (SynthesizeApi *SynthesizeApi) GetList(c *gin.Context) {
	assId := c.Query("assid")

	synthesizeSubmits := synthesizeSubmitService.FindByAssEntityOrderByCreateTimeDesc(assId)
	response.OkWithDetailed(synthesizeSubmits, "ok", c)
}
