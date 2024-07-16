package system

import (
	"au-golang/global"
	"au-golang/model/WxBean"
	"au-golang/model/common/request"
	"au-golang/model/common/response"
	"au-golang/model/system"
	"au-golang/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/cast"
	"github.com/thedevsaddam/gojsonq/v2"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type WxUserApi struct {
}

// 返回用户信息
type Ass struct {
	AssName string
	Logo    string
}
type wxuser struct {
	Avatar    string
	NickName  string
	AssName   []Ass
	StudentId string
	RealName  string
	Token     string
}

const (
	code2sessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

// @Desciption 登录接口

func (w *WxUserApi) GetUserInfo(c *gin.Context) {
	//获取code
	code := c.Query("code")
	if code == "" {
		response.FailWithMessage("参数不能为空", c)
		return
	}

	//调用auth.code2Session接口获取openid
	url := fmt.Sprintf(code2sessionURL, global.GVA_CONFIG.MpKey.AppId, global.GVA_CONFIG.MpKey.AppSecret, code)

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	json := gojsonq.New().FromString(string(body)).Find("openid")

	if json == nil {
		response.FailWithMessage("get openid fail", c)
		return
	}
	openId := json.(string)

	//通过微信小程序的openid查询学生

	var user WxBean.WxUser
	if user, err = wxUserService.FindByMpOpenId(openId); err != nil {
		response.FailWithMessage("查询openid出错", c)
		return
	}

	if user.Id != "" {
		//如果找到该用户，说明已经认证过了，直接签发token
		var token string
		token, err = utils.GenToken(user.Id, user.Nickname)

		var ass system.Ass
		ass = wxUserService.FindAssByAssId(user.AssEntityFirst)
		assFirst := Ass{
			AssName: ass.Assname,
			Logo:    ass.Logo,
		}
		ass = wxUserService.FindAssByAssId(user.AssEntitySecond)
		assSec := Ass{
			AssName: ass.Assname,
			Logo:    ass.Logo,
		}
		var assArr []Ass
		assArr = append(assArr, assFirst)
		assArr = append(assArr, assSec)
		u := wxuser{
			NickName:  user.Nickname,
			Avatar:    user.Avatar,
			RealName:  user.RealName,
			StudentId: user.StudentId,
			AssName:   assArr,
			Token:     token,
		}
		response.OkWithDetailed(u, "ok", c)
		return
	} else {
		//不存在此学生，返回openid，在学生认证的时候一起存入数据库
		response.FailWithDetailed(gin.H{"MpOpenId": openId}, "请前往用户认证！", c)
	}
}

// 学号认证
func (w *WxUserApi) AddWxUser(c *gin.Context) {
	var err error
	var json map[string]string
	err = c.ShouldBindBodyWith(&json, binding.JSON)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//是否认证成功
	resultData := utils.Veri(json["studentId"], json["password"])
	//认证成功
	var num float64 = 20000
	if resultData["code"] == num {
	} else {
		//认证失败
		response.FailWithMessage("学号认证失败", c)
		return
	}
	var user WxBean.WxUser
	err = c.ShouldBindBodyWith(&user, binding.JSON)
	if err != nil {

		response.FailWithMessage(err.Error(), c)
		return
	}

	//验证之后查看是否已经验证
	var exist bool
	if exist, err = wxUserService.FindStudentIsAuth(user.MpOpenId); err != nil {
		log.Println("查询是否已经验证出错：", err)
	}
	if exist == true {
		response.OkWithDetailed(gin.H{"code": 20000}, "您已经绑定过学号", c)
		return
	}

	//否则存入用户信息
	err = wxUserService.SaveWxUser(user)

	if err != nil {

		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{}, "学号认证成功", c)
}

func (w *WxUserApi) GetTotalAssociation(c *gin.Context) {
	ass := wxUserService.FindAll()
	response.OkWithDetailed(ass, "ok", c)
}

func (w *WxUserApi) GetExecllentAssociation(c *gin.Context) {
	excellentAss := wxUserService.FindExcellentAss(1)
	response.OkWithDetailed(excellentAss, "ok", c)
}

func (w *WxUserApi) Test(c *gin.Context) {

}

func (w *WxUserApi) GetAssociationsNameMapAssid(c *gin.Context) {
	ass := wxUserService.FindAllByAssidIsNot(1)
	type mapInfo struct {
		Assid   int
		Assname string
	}
	var asses []mapInfo
	for _, v := range ass {
		info := mapInfo{
			Assid:   v.Assid,
			Assname: v.Assname,
		}
		asses = append(asses, info)
	}
	response.OkWithDetailed(asses, "ok", c)
}
func (w *WxUserApi) GetAssociationsByAssid(c *gin.Context) {
	assId := c.Query("assId")
	Id, _ := strconv.Atoi(assId)
	ass := wxUserService.FindAssByAssId(Id)
	response.OkWithDetailed(ass, "ok", c)
}

// 待测试
func (w *WxUserApi) GetTwitterPic(c *gin.Context) {
	url := c.Query("url")
	resp, err := http.Get(url)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	result := string(body)
	//解析封面图片地址
	begin := strings.Index(result, "var msg_cdn_url")
	end := strings.Index(result, "var cdn_url_1_1")
	picUrl := result[begin:end]
	pictureUrl := strings.Split(picUrl, "\"")[1]
	//解析推文名称
	title_first := strings.LastIndex(result, "<meta property=\"twitter:title\" content=\"")
	title_last := strings.Index(result, "<meta property=\"twitter:creator\"")
	title := strings.Split(result[title_first:title_last], "\"")[3]

	//解析时间
	time_first := strings.Index(result, "if(window.__second_open__)return;")
	time_last := strings.Index(result, "e(t,n,i,document.getElementById(\"publish_time\"));")
	time := strings.Split(result[time_first:time_last], "\"")[5]
	type res struct {
		Time   string
		PicUrl string
		Title  string
	}
	resData := res{
		Time:   time,
		Title:  title,
		PicUrl: pictureUrl,
	}
	response.OkWithDetailed(resData, "ok", c)
}

/*
*
新增一个社团成员退出功能
通过studentId和assId退出
*/
func (w *WxUserApi) ExitAss(c *gin.Context) {
	assId := c.Query("assId")
	studentId := c.Query("studentId")
	if assId == "" || studentId == "" {
		response.FailWithMessage("参数不能为空", c)
		return
	}
	err := wxUserService.ExitAssByStudentId(assId, studentId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("ok", c)
}

func (w *WxUserApi) GetAssByType(c *gin.Context) {
	typ := c.Query("type")
	if typ == "" {
		response.FailWithMessage("参数不能为空", c)
		return
	}
	data, err := wxUserService.GetAssByType(typ)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(data, "success", c)
}

// 社团查找学生
func (w *WxUserApi) SearchStudent(c *gin.Context) {
	var pageInfo request.PageInfo
	assId, exist := c.GetQuery("assId")
	pageInfo.PageSize = 10
	if !exist {
		response.FailWithMessage("传参出错", c)
		return
	}

	user, count := wxUserService.SearchAllBy(c, cast.ToInt(assId), pageInfo)

	response.OkWithDetailed(gin.H{
		"user":  user,
		"count": count,
	}, "ok", c)
}
