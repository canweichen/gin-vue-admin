package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

const (
	APP_KEY     = "607142f3d3e5a"
	APP_SECRET  = "539a83d769ad765c8cf9e621d79e241a"
	APP_VERSION = "v1.2.4"
)

//请求接口Api
var urlMap = map[string]string{
	"GetGoodsList":       "https://openapi.dataoke.com/api/goods/get-goods-list",        //商品列表
	"GetGoodsDetails":    "https://openapi.dataoke.com/api/goods/get-goods-details",     //单品详情
	"PullGoodsByTime":    "https://openapi.dataoke.com/api/goods/pull-goods-by-time",    //定时拉取
	"OpGoodsList":        "https://openapi.dataoke.com/api/goods/nine/op-goods-list",    //9.9元包邮
	"GetBrandList":       "https://openapi.dataoke.com/api/tb-service/get-brand-list",   //品牌库
	"ExplosiveGoodsList": "https://openapi.dataoke.com/api/goods/explosive-goods-list",  //每日爆品推荐
	"CreateTaoKouLing":   "https://openapi.dataoke.com/api/tb-service/creat-taokouling", //淘口令生成
}

//获取签名
func BindSign() map[string]interface{} {
	signMap := make(map[string]interface{})
	signMap["times"] = time.Now().UnixNano() / 1e6 //获取时间戳毫秒
	signMap["nonce"] = GetRandBySixByte()          //获取6位随机数
	sign := fmt.Sprintf("appKey=%s&timer=%d&nonce=%s&key=%s", APP_KEY, signMap["times"], signMap["nonce"], APP_SECRET)
	signMap["sign"] = sign
	//按格式一定格式拼接
	signMap["signRand"] = strings.ToUpper(MD5V([]byte(sign))) //MD5加密 字母转化为大写
	return signMap
}

//拉取商品列表
func BindRequestParams(key string, page interface{}) (string, error) {
	if _, ok := urlMap[key]; !ok {
		return "", errors.New("无效key")
	}
	sign := BindSign()
	return fmt.Sprintf("%s?appKey=%s&nonce=%s&pageId=%v&pageSize=100&signRan=%s&timer=%d&version=%s", urlMap[key], APP_KEY, sign["nonce"], page, sign["signRand"], sign["times"], APP_VERSION), nil
}

//生成淘口令
func CreateTaoKouLing(key string, urls string) (string, error) {
	text := "福建蓝梦科技股份有限公司"
	text = url.QueryEscape(text)
	if _, ok := urlMap[key]; !ok {
		return "", fmt.Errorf("无效key:%s", key)
	}
	sign := BindSign()
	return fmt.Sprintf("%s?appKey=%s&nonce=%s&signRan=%s&timer=%d&version=%s&text=%s&url=%s", urlMap[key], APP_KEY,
		sign["nonce"], sign["signRand"], sign["times"], APP_VERSION, text, urls), nil
}

//生成淘口令
func GetGoodsDetails(key string, id interface{}) (string, error) {
	if _, ok := urlMap[key]; !ok {
		return "", fmt.Errorf("无效key:%s", key)
	}
	sign := BindSign()
	return fmt.Sprintf("%s?appKey=%s&nonce=%s&signRan=%s&timer=%d&version=%s&id=%v", urlMap[key], APP_KEY,
		sign["nonce"], sign["signRand"], sign["times"], APP_VERSION, id), nil
}

/**
 *获取六位随机数
 */
func GetRandBySixByte() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}
