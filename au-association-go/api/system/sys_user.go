package system

import (
	"au-golang/model/Vo"
	"au-golang/model/common/response"
	"au-golang/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"net/http"
	"time"
)

type UserApi struct{}

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (b *UserApi) Login(c *gin.Context) {

	decrypt, _ := ioutil.ReadAll(c.Request.Body)

	body, err := utils.Decrypt2(string(decrypt), []byte("RC3!%(a14f*op52E"))
	if err != nil {
		response.FailWithMessage("解密出错", c)
		return
	}
	fmt.Println("res:", string(body))
	var u user

	err = json.Unmarshal(body, &u)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	realUser, err := userService.FindByNameAndPassword(u.Username, u.Password)
	// 判空处理

	if err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}
	//Jwt签名，存入用户姓名
	token, _ := utils.GenToken(realUser.Id, realUser.Name)
	c.Set("user_id", realUser.Id)
	c.Set("username", realUser.Name)
	response.OkWithDetailed(gin.H{"token": token}, "request success", c)
}

func (b *UserApi) Logout(c *gin.Context) {
	//token := c.GetHeader("X-Token")

	response.OkWithDetailed(gin.H{"data": "success", "code": 20000}, "logout success", c)
}

func (b *UserApi) GetInfo(c *gin.Context) {
	username, exist := c.Get("username")

	if !exist {
		response.FailWithMessage("token失效", c)
	}
	_, user := userService.FindByName(username.(string))
	info := Vo.VoUserInfo{}
	info.Avatar = ""
	info.Name = user.Name
	info.Assid = user.Assid
	info.Introduction = "测试用户"
	info.Role = []string{user.Role}
	response.OkWithDetailed(gin.H{
		"roles":        info.Role,
		"introduction": info.Introduction,
		"avatar":       info.Avatar,
		"name":         info.Name,
		"assid":        info.Assid,
	}, "ok", c)
}

// 新增一键导出社团成员名单
func (b *UserApi) ExportUsersList(c *gin.Context) {
	assId := c.Query("assId")
	users := userService.FindAssUsersByAssId(assId)

	//创建一个excel文件
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("社团人员信息")

	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "真实姓名"
	cell = row.AddCell()
	cell.Value = "学号"
	cell = row.AddCell()
	cell.Value = "微信号"
	cell = row.AddCell()
	cell.Value = "联系电话"
	for _, v := range users {
		row = sheet.AddRow()

		cell := row.AddCell()
		cell.Value = v.RealName
		cell = row.AddCell()
		cell.Value = v.StudentId
		cell = row.AddCell()
		cell.Value = v.WeChatId
		cell = row.AddCell()
		cell.Value = v.PhoneNum
	}
	var buffer bytes.Buffer
	_ = file.Write(&buffer)
	content := bytes.NewReader(buffer.Bytes())
	fileName := fmt.Sprintf("%s.xlsx", "社团成员名单")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	c.Writer.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	http.ServeContent(c.Writer, c.Request, fileName, time.Now(), content)
}
