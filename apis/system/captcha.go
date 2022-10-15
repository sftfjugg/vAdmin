package system

import (
	"github.com/gin-gonic/gin"
	"vAdmin/tools"
	_ "vAdmin/tools"
	"vAdmin/tools/app"
	"vAdmin/tools/captcha"
)

func GenerateCaptchaHandler(c *gin.Context) {
	id, b64s, err := captcha.DriverDigitFunc()
	tools.HasError(err, "验证码获取失败", 500)
	//id, b64s, _ := captcha.DriverDigitFunc()
	app.Custum(c, gin.H{
		"code": 200,
		"data": b64s,
		"id":   id,
		"msg":  "success",
	})
}
