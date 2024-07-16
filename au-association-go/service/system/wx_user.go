package system

import (
	"au-golang/global"
	"au-golang/model/WxBean"
	"au-golang/model/common/request"
	"au-golang/model/system"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"regexp"
	"strconv"
)

type WxUserService struct{}

func (wxUserService *WxUserService) FindByStudentId(id string) WxBean.WxUser {
	var wxUser WxBean.WxUser
	db := global.GVA_DB.Model(WxBean.WxUser{})
	db.Where("student_id = ?", id).First(&wxUser)
	return wxUser
}

func (wxUserService *WxUserService) Save(column string, user WxBean.WxUser, interViewUser WxBean.InterViewUser) {
	db := global.GVA_DB.Model(WxBean.WxUser{})

	db.Where("id=?", user.Id).Update(column, interViewUser.SubmitAssId)
}

func (w *WxUserService) FindByOpenId(openId string) (user WxBean.WxUser) {
	global.GVA_DB.Model(&WxBean.WxUser{}).Where("open_id = ?", openId).Find(&user)
	return user
}

func (w *WxUserService) FindByMpOpenId(MpOpenId string) (user WxBean.WxUser, err error) {
	err = global.GVA_DB.Model(&WxBean.WxUser{}).Where("mp_open_id = ?", MpOpenId).Find(&user).Error
	return
}

// 当数据库中存在用户的student_id并且mp_open_id不为空时，学生已经认证
func (w *WxUserService) FindStudentIsAuth(openid string) (isAuth bool, err error) {
	var user WxBean.WxUser
	err = global.GVA_DB.Model(&WxBean.WxUser{}).Where("mp_open_id = ?", openid).Find(&user).Error
	if user.MpOpenId != "" {
		isAuth = true
		return
	}
	isAuth = false
	return
}

func (w *WxUserService) SaveWxUser(user WxBean.WxUser) (err error) {
	//先查出是否有此学生，有的话更新一下mp_open_id,没有的话之后直接插入
	var u WxBean.WxUser
	err = global.GVA_DB.Model(&WxBean.WxUser{}).Where("student_id=?", user.StudentId).Find(&u).Error
	if u.Id != "" {
		err = global.GVA_DB.Model(&WxBean.WxUser{}).Where("student_id=?", user.StudentId).Update("mp_open_id", user.MpOpenId).Error
		return
	}
	//生成主键
	id := uuid.NewString()
	user.Id = id[:31]

	//使用事务
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		//跳过所有关联
		if err := tx.Omit("AssEntityFirst", "AssEntitySecond").Create(&user).Error; err != nil {
			return nil
		}
		//返回nil提交事务
		return nil
	})
	return
}

func (w *WxUserService) FindAssByAssId(AssId int) (Ass system.Ass) {
	global.GVA_DB.Model(&system.Ass{}).Where("assid = ?", AssId).Find(&Ass)
	return Ass
}

func (w *WxUserService) FindAllByFirstAssOrSecondAss(ass system.Ass) []WxBean.WxUser {
	var wxUser []WxBean.WxUser
	db := global.GVA_DB.Model(WxBean.WxUser{})
	db.Debug().Where("ass_entity_first = ? or ass_entity_second = ?", ass.Assid, ass.Assid).Find(&wxUser)
	return wxUser
}

func (w *WxUserService) SearchAllBy(c *gin.Context, assId int, info request.PageInfo) ([]WxBean.WxUser, int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var total int64

	var wxUser []WxBean.WxUser
	db := global.GVA_DB.Model(WxBean.WxUser{})
	if nickName, isExist := c.GetQuery("nickName"); isExist {
		db.Where("nickname = ?", nickName)
	}
	if realName, isExist := c.GetQuery("realName"); isExist {
		db.Where("real_name = ?", realName)
	}
	if studentId, isExist := c.GetQuery("studentId"); isExist {
		db.Where("student_id = ?", studentId)
	}
	if wechatId, isExist := c.GetQuery("weChatId"); isExist {
		db.Where("we_chat_id = ?", wechatId)
	}
	if phoneNum, isExist := c.GetQuery("phoneNum"); isExist {
		db.Where("phone_num = ?", phoneNum)
	}
	if assId != 0 {
		db.Where("ass_entity_first = ? or ass_entity_second = ?", assId, assId)
	}

	query := c.Query("query")
	isId, _ := regexp.MatchString(`^[0-9]+$`, query)
	if isId {
		result := db.Count(&total)
		result.Debug().Limit(limit).Offset(offset).Find(&wxUser, "student_id like CONCAT('%',?,'%')", query)
		return wxUser, total
	}
	result := db.Count(&total)
	result.Debug().Limit(limit).Offset(offset).Find(&wxUser, "real_name like CONCAT('%',?,'%')", query)
	return wxUser, total
}

func (w *WxUserService) FindAll() (ass []system.Ass) {

	global.GVA_DB.Model(&system.Ass{}).Find(&ass)
	return
}
func (w *WxUserService) FindExcellentAss(excellent int) (ass []system.Ass) {
	global.GVA_DB.Model(&system.Ass{}).Where("is_execllent=?", excellent).Find(&ass)
	return
}

func (w *WxUserService) FindAllByAssidIsNot(id int) (ass []system.Ass) {
	global.GVA_DB.Model(&system.Ass{}).Not("assid", id).Find(&ass)
	return
}

func (w *WxUserService) Login(studentId, password string) bool {
	var user WxBean.WxUser
	global.GVA_DB.Model(&WxBean.WxUser{}).Where("student_id = ? and password = ?", studentId, password).Find(&user)
	if user.Id == "" {
		return false
	}
	return true
}

func (w *WxUserService) SaveOpenIdByStudentId(openid, studentId string) {
	global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&WxBean.WxUser{}).Where("student_id = ?", studentId).Update("open_id", openid).Error; err != nil {
			return err
		}
		return nil
	})
}

// 有待改进
func (w *WxUserService) ExitAssByStudentId(assId, studentId string) (err error) {

	Id, _ := strconv.Atoi(assId)
	var user WxBean.WxUser
	global.GVA_DB.Model(&WxBean.WxUser{}).Where("student_id = ?", studentId).Find(&user)

	if user.AssEntityFirst == Id {
		err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Debug().Model(&user).Where("id=?", user.Id).Update("AssEntityFirst", nil).Error; err != nil {
				return err
			}
			if err := tx.Debug().Model(&WxBean.InterViewUser{}).Where("student_id=? and submit_ass_id=?", user.StudentId, Id).Update("inter_view_status", 5).Error; err != nil {
				return err
			}
			return nil
		})
	} else if user.AssEntitySecond == Id {
		err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Debug().Model(&user).Where("id=?", user.Id).Update("AssEntitySecond", nil).Error; err != nil {
				return err
			}
			if err := tx.Debug().Model(&WxBean.InterViewUser{}).Where("student_id=? and submit_ass_id=?", user.StudentId, Id).Update("inter_view_status", 5).Error; err != nil {
				return err
			}
			return nil
		})
	}
	return

}

func (w *WxUserService) GetAssByType(typ string) (asses []system.Ass, err error) {
	db := global.GVA_DB
	err = db.Model(&system.Ass{}).Where("asstype = ?", typ).Find(&asses).Error
	return
}

func (w *WxUserService) GetMpOpenId(studentId string) (openid string, err error) {
	var user WxBean.WxUser
	err = global.GVA_DB.Model(&WxBean.WxUser{}).Where("student_id", studentId).Find(&user).Error
	openid = user.MpOpenId
	return
}

// 删除社团时批量更新对应社团id字段
func (w *WxUserService) UpdateAssId(assId int) {
	db := global.GVA_DB
	db.Exec("UPDATE wxuser "+
		"SET ass_entity_first = ( CASE WHEN ass_entity_first = ? THEN NULL ELSE ass_entity_first END ),"+
		"ass_entity_second = ( CASE WHEN ass_entity_second = ? THEN NULL ELSE ass_entity_second END )", assId, assId)
}
