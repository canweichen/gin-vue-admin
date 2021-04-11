package v1

import(
	"fmt"
	"gin-vue-admin/global"
	_ "gin-vue-admin/model"
	_ "gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goods/getGoodsList [get]
func GetGoodsList(c *gin.Context) {
	pageS := c.DefaultQuery("page","1")
	limitS := c.DefaultQuery("limit","10")
	page , _ := strconv.Atoi(pageS)
	limit , _ := strconv.Atoi(limitS)
	goodsList , total , err := service.GetGoodsList(page,limit)
	if err != nil {
		global.GVA_LOG.Error("获取商品记录失败",zap.Any("err",err))
		response.FailWithMessage("获取商品记录失败",c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     goodsList,
		Total:    total,
		Page:     page,
		PageSize: limit,
	} ,"获取成功",c)
}