package response

import "gin-vue-admin/model"

type RespShopGoods struct {
	RequestId string    `json:"requestId"`
	Time      uint      `json:"time"`
	Code      int       `json:"code"`
	Msg       string    `json:"msg"`
	Data      ChildList `json:"data"`
}

type ChildList struct {
	List     []model.ShopGoods `json:"list,omitempty"`
	PageId   string            `json:"pageId,omitempty"`
	TotalNum uint              `json:"totalNum,omitempty"`
	GoScroll bool              `json:"goScroll,omitempty"`
	Model    string            `json:"model,omitempty"`
	LongTpwd string            `json:"longTpwd,omitempty"`
}
