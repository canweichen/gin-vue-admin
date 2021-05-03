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
	"net/url"
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
	//图片校验
	for key, item := range goodsList {
		if item.MainPic[:5] != "https" {
			goodsList[key].MainPic = "https:" + item.MainPic
		}
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

// @Tags CreateTaoBaoLing
// @Summary 生成淘口令
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /goods/createTaoLing [post]
func CreateTaoBaoLing(c *gin.Context) {
	//接收生成淘口令的url
	urls := c.DefaultQuery("url", "")
	if urls == "" {
		response.FailWithMessage("链接不为空", c)
		return
	}
	urls = url.QueryEscape(urls)
	var err error
	//生成淘口令
	path, err := utils.CreateTaoKouLing("CreateTaoKouLing", urls)
	//path, err := utils.GetGoodsDetails("GetGoodsDetails", 33235553)
	if err != nil {
		response.FailWithMessage("签名失败："+err.Error(), c)
		return
	}
	resp, err := http.Get(path)
	if err != nil {
		response.FailWithMessage("淘口令成成失败:"+err.Error(), c)
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.FailWithMessage("数据接受失败:"+err.Error(), c)
		return
	}
	//解析数据流
	newData := &response.RespShopGoods{}
	err = json.Unmarshal(data, newData)
	if err != nil {
		response.FailWithMessage("数据解析失败:"+err.Error(), c)
		return
	}
	response.OkWithData(newData, c)
}
