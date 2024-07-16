package utils

import (
	"au-golang/global"
	"au-golang/model/WxBean"
	model "au-golang/model/migrate"
	"au-golang/service"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	//社联面试
	Interview_code_01 = "wBtDklx451_30mFe3tNFmJ6tTlIt1T8rzUYU_qIeS-s"
	//社团面试
	Interview_code_02 = "u_Dg5h_Fuq-3LniqEWtM08ZjmYElNJXtva5_kUgI1og"
	//社联面试时间
	INTERVIEW_TIME = "PHl68j37-9Imvpf_TK0Qo1Bz0aTnA77QlsNimxDjxNk"

	//获取access_token
	u1 = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	u2 = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s"
)

var (
	ACCESS_TOKEN = ""
)

// 社团面试信息
type assInfo struct {
	Phrase1 value `json:"phrase1"`
	Thing2  value `json:"thing2"`
	Thing3  value `json:"thing3"`
	Thing4  value `json:"thing4"`
}

type AssMessage struct {
	TemplateId string   `json:"template_id"`
	Touser     string   `json:"touser"`
	Page       string   `json:"page" default:"index"`
	Data       *assInfo `json:"data"`
}

// 社联面试信息
type departmentInfo struct {
	Thing1  value `json:"thing1"`
	Thing2  value `json:"thing2"`
	Phrase3 value `json:"phrase3"`
}

// 面试时间信息
type interviewTime struct {
	Name       value `json:"thing1"`
	Time       value `json:"time3"`
	Location   value `json:"thing8"`
	Department value `json:"thing9"`
	Info       value `json:"thing5"`
}

type value struct {
	Value string `json:"value"`
}

type WxResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type DepartmentMessage struct {
	TemplateId string          `json:"template_id"`
	Touser     string          `json:"touser"`
	Page       string          `json:"page" default:"index"`
	Data       *departmentInfo `json:"data"`
}

type InterviewTime struct {
	TemplateId string         `json:"template_id"`
	Touser     string         `json:"touser"`
	Page       string         `json:"page" default:"index"`
	Data       *interviewTime `json:"data"`
}

func SendDepartmentMessage(openid, status string, statusNum string, department model.AuInterviewDepartment) (err error) {
	url := fmt.Sprintf(u2, ACCESS_TOKEN)
	info := model.AuInterviewUser{}

	var data = DepartmentMessage{
		TemplateId: Interview_code_01,
		Touser:     openid,
		Page:       "pages/Asso/myInterview/myInterview?success=true",
		Data: &departmentInfo{
			Thing1:  value{Value: department.DepartmentName},
			Thing2:  value{Value: info.GetStatusByNumber(statusNum)},
			Phrase3: value{Value: status},
		},
	}
	msg, _ := json.Marshal(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(msg))
	if err != nil {
		fmt.Println(err)
		return err
	}
	bt, _ := ioutil.ReadAll(resp.Body)

	var wxResp WxResp
	_ = json.Unmarshal(bt, &wxResp)

	if wxResp.ErrCode != 0 {
		return errors.New(wxResp.ErrMsg)
	}

	return nil
}

func SendInterviewTime(openid, time, location, name, departmentName, departmentMessage string) error {
	url := fmt.Sprintf(u2, ACCESS_TOKEN)

	var data = InterviewTime{
		TemplateId: INTERVIEW_TIME,
		Touser:     openid,
		Page:       "pages/Asso/myInterview/myInterview",
		Data: &interviewTime{
			Name:       value{Value: name},
			Time:       value{Value: time},
			Location:   value{Value: location},
			Department: value{Value: departmentName},
			Info:       value{Value: departmentMessage},
		},
	}
	msg, _ := json.Marshal(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(msg))
	if err != nil {
		fmt.Println(err)
		return err
	}
	bt, _ := ioutil.ReadAll(resp.Body)

	var wxResp WxResp
	_ = json.Unmarshal(bt, &wxResp)

	if wxResp.ErrCode != 0 {
		return errors.New(wxResp.ErrMsg)
	}

	return nil
}

func SendAssMessage(openId, status string, interviewUser WxBean.InterViewUser) (err error) {

	url := fmt.Sprintf(u2, ACCESS_TOKEN)

	//查出社团的信息
	assService := service.ServiceGroupApp.SystemServiceGroup.AssService
	ass := assService.FindByAssid(interviewUser.SubmitAssId)

	var data = AssMessage{
		TemplateId: Interview_code_02,
		Touser:     openId,
		Page:       "pages/Asso/myInterview/myInterview",
		Data: &assInfo{
			Phrase1: value{Value: status},
			Thing2:  value{Value: ass.Assname},
			Thing3:  value{Value: ass.Asstype},
			Thing4:  value{Value: ass.Presidentname},
		},
	}
	msg, _ := json.Marshal(data)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(msg))
	if err != nil {
		return err
	}

	bt, _ := ioutil.ReadAll(resp.Body)

	var wxResp WxResp
	_ = json.Unmarshal(bt, &wxResp)

	if wxResp.ErrCode != 0 {
		return errors.New(wxResp.ErrMsg)
	}

	return
}

// 获取access_token
func getToken() (result map[string]interface{}, err error) {
	url := fmt.Sprintf(u1, global.GVA_CONFIG.MpKey.AppId, global.GVA_CONFIG.MpKey.AppSecret)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &result)

	ACCESS_TOKEN = fmt.Sprintf("%v", result["access_token"])

	return
}

func MPTokenRefresh() error {
	// 启动定时前先调用一次
	_, err := getToken()
	if err != nil {
		return err
	}

	freshTokenTicker := time.NewTicker(7000 * time.Second)
	go func() {

		for range freshTokenTicker.C {

			_, err := getToken()
			if err != nil {
				fmt.Printf("小程序access定时器获取失败 %v，%v \n", err.Error(), time.Now())
			}

			fmt.Printf("小程序token定时刷新 %v，%v \n", ACCESS_TOKEN, time.Now())
		}
	}()

	return nil
}
