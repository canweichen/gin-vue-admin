package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitShopGoodsRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	ShopGoodsRouter := Router.Group("goods").Use(middleware.OperationRecord())
	{
		ShopGoodsRouter.POST("getGoodsList", v1.GetGoodsList)   // 获取商品列表
		ShopGoodsRouter.POST("pullGoods", v1.PostPullGoodsList) //拉取商品
	}
	return ShopGoodsRouter
}
