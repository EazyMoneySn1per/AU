package system

import (
	"au-golang/config"
	"au-golang/global"
	"au-golang/model/WxBean"
	"au-golang/model/common/request"
	"au-golang/model/common/response"
	"au-golang/model/system"
	"au-golang/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"path"
	"strconv"
)

type AdminApi struct{}

// GetAssociation 获取所有社团
func (AdminApi *AdminApi) GetAssociation(c *gin.Context) {
	associations := assService.FindAll()
	response.OkWithDetailed(associations, "ok", c)
}

func (AdminApi *AdminApi) GetAssociationsByAssid(c *gin.Context) {
	assIdStr := c.Query("assid")
	assId, _ := strconv.Atoi(assIdStr)
	ass := assService.FindByAssid(assId)
	response.OkWithDetailed(ass, "ok", c)
}

// SetAssociationsByAssid 更改社团信息
func (AdminApi *AdminApi) SetAssociationsByAssid(c *gin.Context) {
	assData := system.Ass{}
	c.ShouldBind(&assData)

	//ass := assService.FindByAssid(assData.Assid)

	//assData.Logo = ass.Logo
	assService.UpdateByAssId(&assData)
	response.OkWithMessage("ok", c)
}

// AddAssociations 新增社团
func (AdminApi *AdminApi) AddAssociations(c *gin.Context) {
	assData := system.Ass{}
	c.ShouldBind(&assData)
	assData.Logo = "default.png"
	assId := assService.FindAssIdMax()
	assData.Assid = assId
	if err := assService.Create(assData); err != nil {
		response.FailWithMessage("添加失败", c)
		return
	}

	// 同步建立社团面试表中的部门
	interViewAss := WxBean.InterViewAss{}
	interViewAss.AssId = assId
	interViewAss.AssName = assData.Assname
	interViewAss.Logo = "default.png"
	interViewAss.PresidentName = assData.Presidentname
	interViewAss.PresidentWechat = "无"
	interViewAss.ShowMessage = "无"
	interViewAss.ConfirmJoinMessage = "无"
	interViewAss.CodeUrl = "无"
	interViewAss.Id = uuid.NewString()
	interviewAssService.Create(&interViewAss)

	response.OkWithMessage("添加成功", c)
}

// DeleteAssociations 删除社团，先删除对应的学生再删除社团
func (AdminApi *AdminApi) DeleteAssociations(c *gin.Context) {
	assIdStr := c.Query("assid")
	if assIdStr == "" {
		response.FailWithMessage("社团不存在", c)
		return
	}
	assId, _ := strconv.Atoi(assIdStr)
	go func() {
		// 删除学生
		interviewUserService.DeleteBatchByAssId(assId)
		userService.DeleteBatchUser(assId)
		wxUserService.UpdateAssId(assId)
		// 删除社团
		interviewAssService.Delete(assId)
		assService.Delete(assId)
	}()
	response.OkWithMessage("后台正在删除，请不要重复删除", c)
}

func (AdminApi *AdminApi) GetAssociationsNameMapAssid(c *gin.Context) {

	type responseData struct {
		Assid   int    `json:"assid"`
		Assname string `json:"assname"`
	}

	associations := assService.FindAll()
	var assData []responseData
	for _, data := range associations {
		info := responseData{data.Assid, data.Assname}
		assData = append(assData, info)
	}

	response.OkWithDetailed(assData, "ok", c)
}

func (AdminApi *AdminApi) GetAssociationName(c *gin.Context) {

	type responseData struct {
		Assname string `json:"assname"`
	}

	associations := assService.FindAll()
	var assData []responseData
	for _, data := range associations {
		info := responseData{data.Assname}
		assData = append(assData, info)
	}
	response.OkWithDetailed(assData, "ok", c)
}

func (AdminApi *AdminApi) GetAssociationStudents(c *gin.Context) {

	type responseData struct {
		Assname string `json:"assname"`
	}

	assIdStr := c.Query("assId")
	assId, _ := strconv.Atoi(assIdStr)
	ass := assService.FindByAssid(assId)
	wxUsers := wxUserService.FindAllByFirstAssOrSecondAss(ass)

	response.OkWithDetailed(wxUsers, "ok", c)
}

// GetStudentsList 根据条件筛选查询学生
func (AdminApi *AdminApi) GetStudentsList(c *gin.Context) {
	var pageInfo request.PageInfo
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("limit"))
	info, _ := strconv.Atoi(c.Query("assId"))

	var assId = info
	if assName, isExist := c.GetQuery("assName"); isExist {
		ass, _ := assService.FindByAssname(assName)
		assId = ass.Assid
	}

	type studentsInfo struct {
		Id              string   `json:"id" gorm:"primary_key"`
		Nickname        string   `json:"nickName"`
		OpenId          string   `json:"openId"`
		MpOpenId        string   `json:"MpOpenId"`
		RealName        string   `json:"realName"`
		Avatar          string   `json:"avatar"`
		StudentId       string   `json:"studentId"`
		WeChatId        string   `json:"weChatId"`
		PhoneNum        string   `json:"phoneNum"`
		AssEntityFirst  int      `json:"ass_entity_first"`
		AssEntitySecond int      `json:"ass_entity_second"`
		AssList         []string `json:"assList"`
	}
	wxUser, total := wxUserService.SearchAllBy(c, assId, pageInfo)

	var sInfo []studentsInfo
	for _, v := range wxUser {
		var ass1 system.Ass
		var ass2 system.Ass
		if v.AssEntityFirst != 0 {
			ass1 = wxUserService.FindAssByAssId(v.AssEntityFirst)
		}
		if v.AssEntitySecond != 0 {
			ass2 = wxUserService.FindAssByAssId(v.AssEntitySecond)
		}
		var list = []string{ass1.Assname, ass2.Assname}
		s := studentsInfo{
			Id:              v.Id,
			Nickname:        v.Nickname,
			OpenId:          v.OpenId,
			MpOpenId:        v.MpOpenId,
			RealName:        v.RealName,
			Avatar:          v.Avatar,
			StudentId:       v.StudentId,
			WeChatId:        v.WeChatId,
			PhoneNum:        v.PhoneNum,
			AssEntityFirst:  v.AssEntityFirst,
			AssEntitySecond: v.AssEntitySecond,
			AssList:         list,
		}
		sInfo = append(sInfo, s)
	}

	response.OkWithDetailed(gin.H{"item": sInfo, "total": total}, "ok", c)
}

// GetActivityCompleteList 根据条件筛选查询活动
func (AdminApi *AdminApi) GetActivityCompleteList(c *gin.Context) {
	var pageInfo request.PageInfo
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("limit"))

	assName := c.Query("assName")
	assIdStr := c.Query("assId")

	var assId int
	var activities []system.ActivityComplete
	var total int64

	if assIdStr != "" {
		assId, _ = strconv.Atoi(assIdStr)
		activities, total = activityCompleteService.FindallbyassAssnameorassAssidorderbydate(assName, assId, pageInfo)
	} else {
		activities, total = activityCompleteService.FindAllLimit(pageInfo)
	}

	response.OkWithDetailed(gin.H{"item": activities, "total": total}, "ok", c)

}

func (AdminApi *AdminApi) GetTwitter(c *gin.Context) {
	var pageInfo request.PageInfo
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("limit"))

	var assId = 0
	if assIdStr, isExist := c.GetQuery("assId"); isExist {
		assId, _ = strconv.Atoi(assIdStr)
	}
	var twitters []system.Twitter
	var total int64
	if assId != 0 {
		twitters, total = twitterService.FindAllByAssIdAndStep(FinalStepTwitter, assId, pageInfo)
	} else {
		twitters, total = twitterService.FindAllByStepLimit(FinalStepTwitter, pageInfo)
	}

	response.OkWithDetailed(gin.H{"item": twitters, "total": total}, "ok", c)
}

func (AdminApi *AdminApi) SetActivityTwitter(c *gin.Context) {
	idStr := c.PostForm("id")
	twitterUrl := c.PostForm("twitterUrl")
	id, _ := strconv.Atoi(idStr)

	activityComplete := activityCompleteService.FindById(id)
	activityComplete.Twitter = twitterUrl
	activityCompleteService.Save(&activityComplete)
	response.OkWithMessage("ok", c)
}

// GetOutlayCompleteList 根据条件筛选查询推文
func (AdminApi *AdminApi) GetOutlayCompleteList(c *gin.Context) {
	var pageInfo request.PageInfo
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("limit"))
	//assName := c.Query("assName")
	assIdStr := c.Query("assId")

	if assIdStr != "" {
		assId, _ := strconv.Atoi(assIdStr)
		twitters, total := outlayService.FindByAssidAndStep(FinalStepTwitter, assId, pageInfo)
		response.OkWithDetailed(gin.H{"item": twitters, "total": total}, "ok", c)
		return
	}

	twitters, total := outlayService.FindAllByStepLimit(FinalStepTwitter, pageInfo)
	response.OkWithDetailed(gin.H{"item": twitters, "total": total}, "ok", c)
	return
}

// SetMiniProgramModule 设置小程序的配置 TODO 测试
func (AdminApi *AdminApi) SetMiniProgramModule(c *gin.Context) {
	form := request.MPModule{}
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c)
		return
	}

	fileForm, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		response.FailWithMessage("上传文件失败", c)
		return
	}

	var bt []byte
	fileForm.Read(bt)

	// 判断文件名后缀是否为 .png .jpg .jpeg
	imageName := fileHeader.Filename
	fileSuffix := path.Ext(imageName)
	if fileSuffix != ".png" && fileSuffix != ".jpg" && fileSuffix != ".jpeg" {
		response.FailWithMessage("只支持png，jpg，jpeg格式的图片", c)
		return
	}
	newName := uuid.NewString() + fileSuffix

	err = utils.UploadFile(bt, global.GVA_CONFIG.StoragePath.RootImagesPath, newName)
	if err != nil {
		response.FailWithMessage("保存文件失败", c)
		return
	}

	// 修改全局变量
	serverIP := "http://au.sztu.edu.cn/images/"
	module := global.GVA_CONFIG.MiniProgramConfig.Module

	for i := 0; i < len(module); i++ {
		if module[i].Name == form.Name {
			module[i].Click = form.Click
			module[i].ImagePath = serverIP + newName

			err := utils.SaveConfig()
			if err != nil {
				response.FailWithMessage(err.Error(), c)
				return
			}

			response.OkWithData(module[i], c)
			return
		}
	}

	// 如果没有这个module，新建一个
	m := config.Module{
		Name:      form.Name,
		ImagePath: serverIP + newName,
		Click:     form.Click,
	}
	module = append(module, m)

	err = utils.SaveConfig()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(m, c)
	return
}

// GetMiniProgramModule 获取小程序配置 TODO 测试
func (AdminApi *AdminApi) GetMiniProgramModule(c *gin.Context) {
	form := request.MPModule{}
	err := c.ShouldBind(&c)

	if err != nil {
		response.OkWithData(global.GVA_CONFIG.MiniProgramConfig.Module, c)
		return
	} else {
		for _, v := range global.GVA_CONFIG.MiniProgramConfig.Module {
			if v.Name == form.Name {
				response.OkWithData(v, c)
				return
			}
		}
	}
}
