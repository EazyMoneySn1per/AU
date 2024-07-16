package api

import (
	"au-golang/global"
	"au-golang/model/WxBean"
	"au-golang/model/migrate/response"
	service "au-golang/service/migrate"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
)

var AssMap = map[string]int{
	"ACM爱好者协会": 346,
	//"筑梦模拟联合国协会":347,
	"筑梦模拟联合国":  347,
	"ART＋潮创":   348,
	"远行客文学社":   349,
	"DIY创想社":   350,
	"排球社":      351,
	"创客社团":     352,
	"街舞社":      353,
	"定向越野社":    354,
	"广府文化研究协会": 355,
	"数模与算法协会":  356,
	"羽毛球社":     357,
	//"先进制造协会":358,
	"先进制造者协会":     358,
	"烹饪社":         359,
	"外语协会":        360,
	"摄影社":         361,
	"书法协会":        362,
	"电竞社":         364,
	"由你设计社":       365,
	"学生就业与职业发展协会": 366,
	"科幻天文协会":      367,
	"体育舞蹈社":       368,
	"武术协会":        369,
	"鸣镝辩论社":       370,
	"篮球社":         371,
	"足球协会":        381,
	"司南画社":        382,
	"魔术协会":        384,
	"ACG动漫社":      385,
	"播音主持与演讲社":    386,
	"乒乓球社":        387,
	"学长团":         388,
	"创新成果转化协会":    389,
	"音乐社":         390,
	//"UAVA无人机协会":391,
	"无人机协会": 391,
	"弓箭社":   392,
}

// 失败结构体
type FailedStudent struct {
	StudentId   string `json:"studentId"`
	StudentName string `json:"studentName"`
	HandleAss   string `json:"handleAss"`
}

func (departmentApiMigrate *DepartmentApiMigrate) SetWxUserAssInBatch(c *gin.Context) {
	// 已加入两个
	var AlreadyJoinTwo []FailedStudent
	// 已加入这个
	var AlreadyJoinThis []FailedStudent
	// 没有该学生
	var NonExistent []FailedStudent

	//var SuccessStudent []model.Wxuser

	file, _ := c.FormFile("file")
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
	for i := 0; i <= 3; i++ {
		rows, err := f.GetRows(f.GetSheetName(i))
		if err != nil {
			return
		}
		for _, row := range rows {
			var tem []string
			for _, colCell := range row {
				tem = append(tem, colCell)
			}
			fmt.Printf("第%d表\n", i)
			fmt.Printf(tem[2])
			if assId, ok := AssMap[tem[2]]; ok {
				wxuser := service.FindWxUserByStudentId(tem[0])
				if wxuser != (WxBean.WxUser{}) {
					if assId != wxuser.AssEntitySecond && assId != wxuser.AssEntityFirst {
						assIdStringFirst := strconv.Itoa(wxuser.AssEntityFirst)
						assIdStringSecond := strconv.Itoa(wxuser.AssEntitySecond)
						if assIdStringFirst == "0" {
							service.UpdateWxUserFirstAss(wxuser.StudentId, assId)
						} else if assIdStringSecond == "0" {
							service.UpdateWxUserFirstSecond(wxuser.StudentId, assId)
						} else {
							AlreadyJoinTwo = append(AlreadyJoinTwo, FailedStudent{wxuser.StudentId, wxuser.RealName, tem[2]})
							fmt.Printf("该学生社团已满 %s %s,要加入社团:%s\n", wxuser.StudentId, wxuser.RealName, tem[2])
						}
					} else {
						AlreadyJoinThis = append(AlreadyJoinThis, FailedStudent{wxuser.StudentId, wxuser.RealName, tem[2]})
						fmt.Printf("该学生已加入该社团 %s %s,要加入社团:%s\n", wxuser.StudentId, wxuser.RealName, tem[2])
					}
				} else {
					fmt.Printf("该学生不存在 %s", tem[0])
					NonExistent = append(NonExistent, FailedStudent{tem[0], tem[1], tem[2]})
				}
			} else {
				fmt.Printf("出现错误社团名 %s\n", tem[2])
			}

		}
	}
	back := map[string]interface{}{
		"AlreadyJoinTwo":  AlreadyJoinTwo,
		"AlreadyJoinThis": AlreadyJoinThis,
		"NonExistent":     NonExistent,
	}
	response.OkWithDetailed(back, "导入成功", c)
	faildFile1, _ := os.Create("社团名额已满名单.txt")
	faildFile1.WriteString("\n")
	faildFile1.WriteString("两个社团名额已满: \n")
	faildFile1.WriteString("\n")
	for _, v := range AlreadyJoinTwo {
		faildFile1.WriteString("姓名: " + v.StudentName + " 要加入的社团: " + v.HandleAss + " 学号: " + v.StudentId + "\n")
	}
	faildFile2, _ := os.Create("已经加入该社团名单.txt")
	faildFile2.WriteString("\n")
	faildFile2.WriteString("已经加入该社团: \n")
	faildFile2.WriteString("\n")
	for _, v := range AlreadyJoinThis {
		faildFile2.WriteString("姓名: " + v.StudentName + " 要加入的社团: " + v.HandleAss + " 学号: " + v.StudentId + "\n")
	}
	faildFile3, _ := os.Create("没有绑定学号名单.txt")
	faildFile3.WriteString("\n")
	faildFile3.WriteString("没有在社联+系统绑定学号: \n")
	faildFile3.WriteString("\n")
	for _, v := range NonExistent {
		faildFile3.WriteString("姓名: " + v.StudentName + " 要加入的社团: " + v.HandleAss + " 学号: " + v.StudentId + "\n")
	}
}

func (departmentApiMigrate *DepartmentApiMigrate) GetTotalJoinAssStudent(c *gin.Context) {
	count := 0
	wxuser := service.FindWxUser()
	for _, v := range wxuser {
		assIdStringFirst := strconv.Itoa(v.AssEntityFirst)
		assIdStringSecond := strconv.Itoa(v.AssEntitySecond)
		if assIdStringFirst != "0" || assIdStringSecond != "0" {
			count++
		}
	}
	fmt.Println("%d", count)
}
