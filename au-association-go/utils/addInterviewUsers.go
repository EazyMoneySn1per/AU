package utils

import (
	"au-golang/global"
	"au-golang/model/WxBean"
	"au-golang/service/system"
	"fmt"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
	"strings"
)

type interviewUsers struct {
	Name        string `json:"name" binding:"required"`
	Sex         string `json:"sex" binding:"required"`
	StudentId   string `json:"studentId" binding:"required"`
	PhoneNum    string `json:"phoneNum" binding:"required"`
	WxNum       string `json:"wxNum" binding:"required"`
	Description string `json:"description"`
	Ass         string `json:"ass" binding:"required"`
}

// 从社团导入学生列表
func AddInterviewUsers() {
	f, err := excelize.OpenFile("./线下报名社团人员名单.xlsx")
	if err != nil {
		return
	}
	rows, err := f.GetRows(f.GetSheetName(0))
	for i, row := range rows {
		var temp []string
		if i != 0 {
			for _, cell := range row {
				temp = append(temp, cell)
			}
			r := interviewUsers{
				Name:        temp[0],
				StudentId:   temp[1],
				PhoneNum:    temp[2],
				WxNum:       temp[3],
				Description: temp[4],
				Sex:         temp[5],
				Ass:         temp[6],
			}
			if strings.Contains(temp[6], " ") {
				asses := strings.Split(temp[6], " ")
				for _, v := range asses {
					r.Ass = v
					handleAss(r)
				}
			} else {
				handleAss(r)
			}
		}
	}
}

func handleAss(r interviewUsers) {

	//判断社团数量
	var userService system.InterViewUserService
	var assService system.InterViewAssService
	count := userService.FindAlByInterviewAssCount(r.StudentId, 0)
	//面试社团中有些社团没有，通过名称查找所在社团
	ass, err := assService.FindByAssname(r.Ass)
	//找到面试社团
	interviewAss := assService.FindByAssId(ass.Assid)
	if err != nil {
		fmt.Printf("报名失败，学生：%v，学号：%v，报名社团：%v，原因：%v", r.Name, r.StudentId, ass.Assname, err.Error())
		return
	}
	if count >= 2 {
		fmt.Printf("报名失败，学生：%v， 学号：%v，报名社团：%v，原因：%v", r.Name, r.StudentId, ass.Assname, "报名社团已经达到2个")
		return
	}

	//判断是否报名了同一个社团
	exist := userService.FindByStudentIdAndInterViewAss(r.StudentId, ass.Assid)
	if exist {
		fmt.Printf("报名失败，学生：%v，学号：%v，报名社团：%v，原因：%v", r.Name, r.StudentId, ass.Assname, "重复报名社团")
		return
	}

	//存入数据库
	user := WxBean.InterViewUser{
		Id:              uuid.NewString(),
		Name:            r.Name,
		StudentId:       r.StudentId,
		Description:     r.Description,
		PhoneNum:        r.PhoneNum,
		WxNum:           r.WxNum,
		Sex:             r.Sex,
		SubmitAssId:     ass.Assid,
		InterViewAss:    interviewAss,
		BackMessage:     global.StageOneWaitB.GetMessageByStage(1),
		ButtonControl:   1,
		InterViewStatus: 1,
	}
	err = userService.AddUser(user)
	if err != nil {
		fmt.Printf("报名失败，学生：%v，学号：%v，报名社团：%v，原因：%v", r.Name, r.StudentId, ass, err.Error())
		return
	}
}
