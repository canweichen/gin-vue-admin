package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

//@author: [piexlmax](https://learnku.com/docs/gorm/v2/create/9732)
//@function: GetGoodsList
//@description: 获取商品列表
//@param: page int,limit int
//@return: goodsList *[]model.ShopGoods,err error
func GetGoodsList(page, limit int, title string) (goodsList []model.ShopGoods, total int64, err error) {
	offset := (page - 1) * limit
	db := global.GVA_DB.Model(&model.ShopGoods{})
	if title != "" {
		db.Where("title LIKE ?", "%"+title+"%")
	}
	_ = db.Count(&total).Error
	err = db.Order("id DESC").Limit(limit).Offset(offset).Find(&goodsList).Error
	return goodsList, total, err
}

//@author: [piexlmax](https://learnku.com/docs/gorm/v2/create/9732)
//@function: InsertMore
//@description: 批量插入商品记录
//@param: goodsList *[]model.ShopGoods
//@return: err error
func InsertMore(goodsList *[]model.ShopGoods) error {
	err := global.GVA_DB.Create(goodsList).Error
	return err
}
