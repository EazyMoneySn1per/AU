package system

import (
	"au-golang/model/common/response"
	"au-golang/model/system"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strconv"
)

type AccountApi struct{}

/*
GetLists 查询所有后台管理员账号，

	@param role 根据权限查询，非必须
	@param page 查询页数，必须
	@param limit 每页条数，必须
*/
func (account *AccountApi) GetLists(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	role := c.Query("role")
	if role != "" {
		err, users, total := userService.FindAllByRole(role, page, limit)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithDetailed(gin.H{"item": users, "total": total}, "ok", c)
		return
	} else {
		users, total, err := userService.FindAll(page, limit)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithDetailed(gin.H{"item": users, "total": total}, "ok", c)
		return
	}

}

type Info struct {
	Name     string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
	AssName  string `form:"assName" json:"assName"`
	Role     string `form:"role" json:"role"`
}

/*
UpdateInfo 更变账号信息

	@param name 用户名
	@param password 密码
	@param assName 所在社团，如果不对则无法更新数据
	@param role 权限
*/
func (account *AccountApi) UpdateInfo(c *gin.Context) {
	userInfo := Info{}
	c.ShouldBind(&userInfo)

	err, user := userService.FindByName(userInfo.Name)
	if err == gorm.ErrRecordNotFound {
		user = system.User{}
		user.Id = uuid.NewString()[:31]
	}
	user.Name = userInfo.Name
	user.Password = userInfo.Password
	user.Role = userInfo.Role

	ass, err := assService.FindByAssname(userInfo.AssName)

	if err == gorm.ErrRecordNotFound {
		response.FailWithMessage("社团不存在，设置失败", c)
		return
	} else {
		user.Ass = ass
		user.Assid = ass.Assid
		err := userService.Save(user)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}
	response.OkWithMessage("设置成功", c)
}

/*
DeleteUser 删除用户

	@param name
	@return
*/
func (account *AccountApi) DeleteUser(c *gin.Context) {
	_, user := userService.FindByName(c.Query("name"))
	userService.DeleteUser(user)
	response.OkWithMessage("Delete success", c)
}
