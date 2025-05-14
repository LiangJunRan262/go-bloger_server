package log_service

import (
	"bloger_server/core"
	"bloger_server/global"
	"bloger_server/models"
	"bloger_server/models/enum"
	"bloger_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

func NewLoginSuccess(c *gin.Context, loginType enum.LoginType) {
	ip := c.ClientIP()
	addr := core.GetIPAddr(ip)

	// token := c.GetHeader("token")
	//userID := uint(1)   // 假设从token中获取用户ID，这里简化为1
	//Username := "admin" // 假设从token中获取用户名，这里简化为"admin"
	// Password := "123456" // 假设从token中获取密码，这里简化为"123456"

	claims, err := jwts.ParseTokenByGin(c)
	userID := uint(0)
	Username := ""
	if err == nil && claims != nil {
		userID = claims.UserID
		Username = claims.Username
	}

	loginLog := &models.LogModel{
		LogType:     enum.LoginLogType, // 假设转换方式为 enum.LoginLogType(loginType),
		Title:       "登录成功",
		Content:     "",
		UserID:      userID,
		IP:          ip,
		Address:     addr,
		LoginStatus: 1,         // 登录成功
		Username:    Username,  // 用户名
		Password:    "",        // 密码
		LoginType:   loginType, // 登录类型
	}
	global.DB.Create(loginLog)
}

func NewLoginFail(c *gin.Context, loginType enum.LoginType, msg string, username string, password string) {
	ip := c.ClientIP()
	addr := core.GetIPAddr(ip)

	loginLog := &models.LogModel{
		LogType:     enum.LoginLogType, // 假设转换方式为 enum.LoginLogType(loginType),
		Title:       "用户登录失败",
		Content:     msg,
		IP:          ip,
		Address:     addr,
		LoginStatus: 1,         // 登录成功
		Username:    username,  // 用户名
		Password:    password,  // 密码
		LoginType:   loginType, // 登录类型
	}
	global.DB.Create(loginLog)
}
