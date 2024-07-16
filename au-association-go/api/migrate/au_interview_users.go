package api

import (
	"au-golang/global"
	model "au-golang/model/migrate"
	"au-golang/model/migrate/request"
	"au-golang/model/migrate/response"
	service "au-golang/service/migrate"
	"au-golang/service/system"
	utils "au-golang/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/xuri/excelize/v2"
	"regexp"
	"strconv"
)

type UserApiMigrate struct{}

//func Test(c *gin.Context) {
//	ExportExcelSetInterviewTime(c)
//}

/*
*
添加报名学生
*/
func (userApiMigrate *UserApiMigrate) AddInterviewUser(c *gin.Context) {
	var r request.AddUser
	_ = c.ShouldBindJSON(&r)
	// 根据ID找到所在的部门
	department := service.FindDepartment(r.Department)
	var resp response.StatusReturn
	if !utils.CheckTime(global.GVA_CONFIG.Interview.OpenTime, global.GVA_CONFIG.Interview.EndTime) {
		resp.Status = 0
		response.OkWithDetailed(resp, "报名开始时间为:"+global.GVA_CONFIG.Interview.OpenTime, c)
		return
	}
	// 检测报名部门数量
	_, checkCount := service.FindStudentInterviewLists(r.StudentId)
	if len(checkCount) >= 2 {
		resp.Status = 0
		response.OkWithDetailed(resp, "最多只能报名2个部门", c)
		return
	}
	_, checkUser := service.FindStudentInterviewAndDepartment(r.StudentId, department.ID)
	// 如果没有重复报同一个部门
	if checkUser.StudentId == r.StudentId {
		resp.Status = 0
		response.OkWithDetailed(resp, "已报名此部门", c)
		return
	}
	// 初始化时间表实体
	time := model.AuInterviewTimeTable{Month: -1, Date: -1, Hour: -1, Minute: -1, Location: "-1"}
	bothUUID := uuid.NewV4()
	time.UUID = bothUUID
	// 初始化用户实体
	user := model.AuInterviewUser{Name: r.Name, StudentId: r.StudentId, PhoneNum: r.PhoneNum, WxNum: r.WxNum, Description: r.Description, Sex: r.Sex, Status: "1", AuInterviewDepartmentID: department.ID, AuInterviewTimeTable: time}
	user.UUID = bothUUID
	service.AddUser(user)
	// 设置返回状态
	resp.Status = 1
	response.OkWithDetailed(resp, "报名成功", c)

}

/*
*
查找学生的面试记录
*/
func (userApiMigrate *UserApiMigrate) FindStudentInterviewLists(c *gin.Context) {
	studentId := c.Query("studentId")
	// 找到这个学生所有的报名信息
	err, user := service.FindStudentInterviewLists(studentId)
	// 定义一个返回体数组，只返回需要的字段
	var resps []response.WxUserReturn
	for _, v := range user {
		// 根据面试状态数字获取当前所在的阶段
		statusMsg := utils.GetInterviewStatusMsg(v.Status)
		// 根据部门ID找到报名的部门
		department := service.FindDepartmentById(v.AuInterviewDepartmentID)
		// 根据面试状态数字和部门ID找到对应阶段的返回信息
		backMessage := utils.GetInterviewDepartmentMsg(v.Status, department)
		var timeBack string
		var locationBack string
		// 判读有无面试时间和地点
		if v.AuInterviewTimeTable.Month == -1 && v.AuInterviewTimeTable.Date == -1 {
			timeBack = "暂无"
		} else {
			temMinute := strconv.Itoa(v.AuInterviewTimeTable.Minute)
			if len(temMinute) == 1 {
				temMinute = "0" + temMinute
			}
			timeBack = strconv.Itoa(v.AuInterviewTimeTable.Month) + "月" + strconv.Itoa(v.AuInterviewTimeTable.Date) + "日 " + strconv.Itoa(v.AuInterviewTimeTable.Hour) + ":" + temMinute
		}
		if v.AuInterviewTimeTable.Location == "-1" {
			locationBack = "暂无"
		} else {
			locationBack = v.AuInterviewTimeTable.Location
		}
		// 设置响应返回体
		resp := response.WxUserReturn{
			UUID:           v.UUID,
			Name:           v.Name,
			BackMessage:    backMessage,
			Time:           timeBack,
			Location:       locationBack,
			Status:         v.Status,
			StatusMsg:      statusMsg,
			DepartmentName: department.DepartmentName,
		}
		// 拼接切片
		resps = append(resps, resp)
	}
	if err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(resps, "获取成功", c)
	}
}

/*
*
设置面试成功
*/
func (userApiMigrate *UserApiMigrate) SetInterviewUerStatusSuccess(c *gin.Context) {
	queryUuid := c.Query("uuid")
	_, user := service.FindUserByUUID(queryUuid)

	var s system.WxUserService
	openid, _ := s.GetMpOpenId(user.StudentId)

	department := service.FindDepartmentById(user.AuInterviewDepartmentID)
	if err := utils.SendDepartmentMessage(openid, "通过", user.Status, department); err != nil {
		fmt.Printf("订阅消息推送出错: %v。学生:%v %v，面试状态: %v \n", user.StudentId, user.Name, err.Error(), user.GetStatus())
	}

	switch user.Status {
	case "1":
		//一面成功
		user.Status = "2"
	case "2":
		//二面成功
		user.Status = "3"
	}
	var array []string
	array = append(array, queryUuid)

	err := service.SetInterviewUserStatusInBatch(array, user.Status)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	response.OkWithMessage("设置成功", c)
}

/*
*
批量设置面试状态
*/
func (userApiMigrate *UserApiMigrate) SetInterviewUserStatusInBatch(c *gin.Context) {
	var req request.StatusAndUuids
	_ = c.ShouldBindJSON(&req)
	// 这里找到用户们最初的状态，用于消息推送
	_, oldUsers := service.FindUserByUUIDInBatch(req.Uuids)

	_ = service.SetInterviewUserStatusInBatch(req.Uuids, req.Status)

	msg := ""
	if req.Status == "2" || req.Status == "3" {
		msg = "通过"
	} else {
		msg = "失败"
	}

	// 异步推送消息
	go func() {
		for _, v := range req.Uuids {
			_, user := service.FindUserByUUID(v)
			var s system.WxUserService
			openid, _ := s.GetMpOpenId(user.StudentId)
			department := service.FindDepartmentById(user.AuInterviewDepartmentID)

			oldStatus := ""
			for i := 0; i < len(oldUsers); i++ {
				if oldUsers[i].StudentId == user.StudentId {
					oldStatus = oldUsers[i].Status
				}
			}
			if err := utils.SendDepartmentMessage(openid, msg, oldStatus, department); err != nil {
				fmt.Printf("订阅消息推送出错: %v。学生:%v %v，面试状态: %v \n", user.StudentId, user.Name, err.Error(), user.GetStatus())
			}
		}
	}()

	response.OkWithMessage("yes", c)
}

/*
*
设置面试失败
*/
func (userApiMigrate *UserApiMigrate) SetInterviewUerStatusFailed(c *gin.Context) {
	queryUuid := c.Query("uuid")
	_, user := service.FindUserByUUID(queryUuid)

	var s system.WxUserService
	openid, _ := s.GetMpOpenId(user.StudentId)
	department := service.FindDepartmentById(user.AuInterviewDepartmentID)
	if err := utils.SendDepartmentMessage(openid, "失败", user.Status, department); err != nil {
		fmt.Printf("订阅消息推送出错: %v。学生:%v %v，面试状态: %v \n", user.StudentId, user.Name, err.Error(), user.GetStatus())
	}

	switch user.Status {
	case "1":
		//一面失败
		user.Status = "9"
	case "2":
		//二面失败
		user.Status = "8"
	}

	var array []string
	array = append(array, queryUuid)

	err := service.SetInterviewUserStatusInBatch(array, user.Status)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	response.OkWithMessage("设置成功", c)
}

/*
*
导入文件设置面试时间和地点
*/
func (userApiMigrate *UserApiMigrate) ExportExcelSetInterviewTime(c *gin.Context) {
	//var r request.FileAndDepartment
	file, _ := c.FormFile("file")
	department := c.PostForm("department")
	departmentId := service.FindDepartment(department).ID
	// 存到本地
	var dst = global.GVA_CONFIG.Local.Excel + file.Filename
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		return
	}
	// 打开excel 并且逐行读取
	f, err := excelize.OpenFile(dst)
	if err != nil {
		return
	}
	rows, err := f.GetRows(f.GetSheetName(0))

	var errStudents []response.ErrImport
	for i, row := range rows {
		var tem []string
		if i != 0 { // 跳过第一行
			for _, colCell := range row {
				tem = append(tem, colCell)
				//fmt.Print(colCell, "\t")
			}
			_, student := service.FindStudentInterviewAndDepartment(tem[0], departmentId)
			// 导入失败的
			if student.StudentId != tem[0] {
				errStudents = append(errStudents, response.ErrImport{
					Name:      tem[1],
					StudentId: tem[0],
				})
			}
			err := service.SetInterviewUserTime(student.AuInterviewTimeTable.UUID.String(), tem[2], tem[3], tem[4], tem[5], tem[6])
			if err != nil {
				return
			}

			// 异步推送消息
			go func(uuid string) {
				_, user := service.FindUserByUUID(uuid)
				timeBack, locationBack := user.GetInterviewTime()

				var s system.WxUserService
				openid, _ := s.GetMpOpenId(user.StudentId)
				department := service.FindDepartmentById(user.AuInterviewDepartmentID)

				if err := utils.SendInterviewTime(openid, timeBack, locationBack, user.Name, department.DepartmentName, "请在小程序中查看"); err != nil {
					fmt.Printf("订阅消息-面试时间-推送出错: %v。学生:%v %v，面试状态: %v \n", user.StudentId, user.Name, err.Error(), user.GetStatus())
				}

			}(student.AuInterviewTimeTable.UUID.String())
		}
	}
	response.OkWithDetailed(errStudents, "导入成功", c)
}

/*
*
单个设置学生面试时间和地点
*/
func (userApiMigrate *UserApiMigrate) SetInterviewTime(c *gin.Context) {
	var r request.InterviewTime
	var resp response.StatusReturn
	_ = c.ShouldBindJSON(&r)
	err := service.SetInterviewUserTime(r.Uuid, r.Month, r.Date, r.Hour, r.Minute, r.Location)
	if err != nil {
		resp.Status = 0
		response.OkWithDetailed(resp, "设置失败", c)
		return
	}

	err, user := service.FindUserByUUID(r.Uuid)
	if err != nil {
		response.OkWithDetailed(resp, "查找uuid失败", c)
		return
	}
	timeBack, locationBack := user.GetInterviewTime()

	var s system.WxUserService
	openid, _ := s.GetMpOpenId(user.StudentId)
	department := service.FindDepartmentById(user.AuInterviewDepartmentID)

	if err := utils.SendInterviewTime(openid, timeBack, locationBack, user.Name, department.DepartmentName, "请在小程序中查看"); err != nil {
		fmt.Printf("订阅消息-面试时间-推送出错: %v。学生:%v %v，面试状态: %v \n", user.StudentId, user.Name, err.Error(), user.GetStatus())
	}

	resp.Status = 1
	response.OkWithDetailed(resp, "设置成功", c)
}

/*
*
查找部门的面试学生(分页)
*/
func (userApiMigrate *UserApiMigrate) FindDepartmentUsers(c *gin.Context) {
	var r request.DepartmentAndStatus
	_ = c.ShouldBindQuery(&r)
	var users = model.SortUserList{}
	query := r.Query
	isId, _ := regexp.MatchString(`^[0-9]+$`, query)
	users, count := service.FindDepartmentUsers(r.Department, r.Status, query, r.PageInfo.Page, r.PageInfo.Limit, isId)
	users.Sort()

	res := response.PageResult{
		Item:  users,
		Total: count,
	}
	response.OkWithDetailed(res, "查找成功", c)
}

/*
*
查找部门的面试所有学生
*/
func (userApiMigrate *UserApiMigrate) FindDepartmentAllUsers(c *gin.Context) {
	var r request.DepartmentAndStatus
	_ = c.ShouldBindQuery(&r)
	var users = model.SortUserList{}
	users, _ = service.FindDepartmentAllUsers(r.Department, r.Status)
	users.Sort()
	//var respUsers []response.StudentWithTime
	//for _,user :=  range users {
	//	resp := response.StudentWithTime{
	//		StudentId: user.StudentId,
	//		Name:      user.Name,
	//		Month:     user.AuInterviewTimeTable.Month,
	//		Date:      user.AuInterviewTimeTable.Date,
	//		Hour:      user.AuInterviewTimeTable.Hour,
	//		Minute:    user.AuInterviewTimeTable.Minute,
	//		Location:  user.AuInterviewTimeTable.Location,
	//	}
	//	respUsers = append(respUsers,resp)
	//}
	response.OkWithDetailed(users, "ok", c)

}
