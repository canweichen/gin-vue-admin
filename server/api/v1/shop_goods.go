package v1

import (
	"encoding/json"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags GetGoodsList
// @Summary 获取商品列表
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goods/getGoodsList [post]
func GetGoodsList(c *gin.Context) {
	type ReceiveParams struct {
		Page     int    `json:"page"`
		PageSize int    `json:"pageSize"`
		Title    string `json:"title"`
	}
	var pageInfo ReceiveParams = ReceiveParams{}
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	goodsList, total, err := service.GetGoodsList(pageInfo.Page, pageInfo.PageSize, pageInfo.Title)
	if err != nil {
		global.GVA_LOG.Error("获取商品记录失败", zap.Any("err", err))
		response.FailWithMessage("获取商品记录失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     goodsList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// @Tags PostPullGoodsList
// @Summary 拉取商品列表
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goods/postPullGoodsList [post]
func PostPullGoodsList(c *gin.Context) {
	//临时存储页码
	key := "pull:goods:pid"
	pageId, _ := global.GVA_REDIS.Get(key).Result()
	if pageId == "" {
		pageId = "1"
	}
	fmt.Println(pageId)
	sign, err := utils.BindRequestParams("GetGoodsList", pageId)
	if err != nil {
		global.GVA_LOG.Error("验签失败", zap.Any("err", err))
		response.FailWithMessage("验签失败", c)
		return
	}
	resp, err := http.Get(sign)
	if err != nil {
		global.GVA_LOG.Error("数据抓取失败！", zap.Any("err", err))
		response.FailWithMessage("数据抓取失败", c)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("数据读取失败！", zap.Any("err", err))
		response.FailWithMessage("数据读取失败", c)
		return
	}
	newBody := &response.RespShopGoods{}
	//数据映射
	err = json.Unmarshal(body, newBody)
	if err != nil {
		global.GVA_LOG.Error("json序列化失败！", zap.Any("err", err))
		response.FailWithMessage("json序列化失败", c)
		return
	}
	//页码写入缓存
	timer := time.Duration(time.Now().Unix()) + 30*time.Second
	err = global.GVA_REDIS.Set(key, newBody.Data.PageId, timer).Err()
	if err != nil {
		global.GVA_LOG.Error("写入缓存失败", zap.Any("err", err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err = service.InsertMore(newBody.Data.List); err != nil {
		global.GVA_LOG.Error("批量插入数据失败", zap.Any("err", err))
		response.FailWithMessage("批量插入数据失败:"+err.Error(), c)
		return
	}
	response.Result(0, newBody, "拉取成功", c)
}
