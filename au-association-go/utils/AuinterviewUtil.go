package utils

import (
	model "au-golang/model/migrate"
	"au-golang/model/migrate/request"
	service "au-golang/service/migrate"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
)

// AddInterviewFromExcel 从excel中导入面试数据
func AddInterviewFromExcel() {
	// 打开excel 并且逐行读取
	f, err := excelize.OpenFile("./test.xlsx")
	if err != nil {
		return
	}
	rows, err := f.GetRows(f.GetSheetName(0))

	for i, row := range rows {
		var tem []string
		if i != 0 { // 跳过第一行
			for _, colCell := range row {
				tem = append(tem, colCell)
			}

			r := request.AddUser{
				Name:        tem[0],
				StudentId:   tem[1],
				PhoneNum:    tem[2],
				WxNum:       tem[3],
				Description: tem[4],
				Sex:         tem[5],
				Department:  tem[6],
			}

			// 面试部门，做处理
			if strings.Contains(tem[6], "，") {
				departments := strings.Split(tem[6], "，")
				for _, v := range departments {
					r.Department = v
					handle(r)
				}
			} else {
				handle(r)
			}
		}
	}

}

func handle(r request.AddUser) {
	// 根据ID找到所在的部门
	departmentId, err := strconv.Atoi(r.Department)
	if err != nil {
		fmt.Printf("报名失败。学生%v,学号:%v,报名部门:%v,原因：%v \n", r.Name, r.StudentId, r.Department, "最多只能报名2个部门")
		return
	}
	department := service.FindDepartmentById(uint(departmentId))

	err, checkUser := service.FindStudentInterviewAndDepartment(r.StudentId, department.ID)
	if err != nil {
		fmt.Printf("报名失败。学生%v,学号:%v,报名部门:%v,原因：%v \n", r.Name, r.StudentId, department.DepartmentName, err.Error())
		return
	}
	// 如果没有重复报同一个部门
	if checkUser.StudentId == r.StudentId {
		fmt.Printf("报名失败。学生%v,学号:%v,报名部门:%v,原因：%v \n", r.Name, r.StudentId, department.DepartmentName, "已报名此部门")
		return
	}

	// 检测报名部门数量
	_, checkCount := service.FindStudentInterviewLists(r.StudentId)
	if len(checkCount) >= 2 {
		fmt.Printf("报名失败。学生%v,学号:%v,报名部门:%v,原因：%v \n", r.Name, r.StudentId, department.DepartmentName, "最多只能报名2个部门")
		return
	}

	// 初始化时间表实体
	time := model.AuInterviewTimeTable{Month: -1, Date: -1, Hour: -1, Minute: -1, Location: "-1"}
	bothUUID := uuid.NewV4()
	time.UUID = bothUUID
	// 初始化用户实体
	user := model.AuInterviewUser{Name: r.Name, StudentId: r.StudentId, PhoneNum: r.PhoneNum, WxNum: r.WxNum, Description: r.Description, Sex: r.Sex, Status: "1", AuInterviewDepartmentID: department.ID, AuInterviewTimeTable: time}
	user.UUID = bothUUID
	err, _ = service.AddUser(user)
	if err != nil {
		fmt.Printf("报名失败。学生%v,学号:%v,报名部门:%v,原因：%v \n", r.Name, r.StudentId, department.DepartmentName, err.Error())
		return
	}
}
