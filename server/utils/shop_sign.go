package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	APP_KEY = "607142f3d3e5a"
	APP_SECRET = "539a83d769ad765c8cf9e621d79e241a"
	APP_URL = "https://openapi.dataoke.com/api/goods/get-goods-list"
	APP_VERSION = "v1.2.4"
)

func BindSign() string{
	times := time.Now().UnixNano() / 1e6 //获取时间戳毫秒
	nonce := GetRandBySixByte()//获取6位随机数
	sign := fmt.Sprintf("appKey=%s&timer=%d&nonce=%s&key=%s",APP_KEY,times,nonce,APP_SECRET) //按格式一定格式拼接
	signRand := strings.ToUpper(MD5V([]byte(sign))) //MD5加密 字母转化为大写
	return fmt.Sprintf("%s?appKey=%s&nonce=%s&pageId=1&pageSize=100&signRan=%s&timer=%d&version=%s",APP_URL,APP_KEY,nonce,signRand,times,APP_VERSION)
}

func GetRandBySixByte() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}

func GetGoodsList() {
	
}
