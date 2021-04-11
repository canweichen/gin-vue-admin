package response

import "gin-vue-admin/model"

type RespShopGoods struct{
	RequestId string `json:"requestId"`
	Time uint `json:"time"`
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data ChildList `json:"data"`
}

type ChildList struct{
	List []model.ShopGoods `json:"list"`
	PageId string `json:"pageId"`
	TotalNum uint `json:"totalNum"`
	GoScroll bool `json:"goScroll"`
}