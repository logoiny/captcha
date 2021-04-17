package router

import (
	"github.com/gin-gonic/gin"
	"captcha/work"
	"captcha/modle"
)

func Router() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("v1/captcha")

	v1.POST("/get", Get)
	v1.POST("/verify", Verify)
	v1.POST("/reload", Reload)

	return r

}

func Reload(ctx *gin.Context) {
	var req modle.CaptchaReload

	if err := ctx.BindJSON(&req); err != nil {
		ctx.String(200, "err params")
	}

	rsp := work.CaptchaReload(req.CaptchaId, req.ImgWidth, req.ImgHigh, req.DfWh)
	ctx.JSON(200, rsp)
}

func Verify(ctx *gin.Context) {
	var req modle.CaptchaVerify

	if err := ctx.BindJSON(&req); err != nil {
		ctx.String(200, "err params")
	}

	rsp := work.CaptchaVerify(req.CaptchaId, req.Captcha)
	ctx.JSON(200, rsp)
}

func Get(ctx *gin.Context) {
	var req modle.CaptchaGet

	if err := ctx.BindJSON(&req); err != nil {
		ctx.String(200, "err params")
	}

	rsp := work.CaptchaGet(req.ImgWidth, req.ImgHigh, req.DfWh)
	ctx.JSON(200, rsp)
}
