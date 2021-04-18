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
func InsertMore(goodsList []model.ShopGoods) error {
	//获取主键值
	var ids []uint
	for _, item := range goodsList {
		ids = append(ids, item.Id)
	}
	var existsIds []uint
	existGoods := []model.ShopGoods{}
	//查询数据库是否存在相应记录
	err := global.GVA_DB.Model(&model.ShopGoods{}).Select("id").Find(&existGoods, ids).Error
	if err != nil {
		return err
	}
	//删除重复的记录
	for _, val := range existGoods {
		existsIds = append(existsIds, val.Id)
	}
	if len(existsIds) > 0 {
		err = global.GVA_DB.Delete(&model.ShopGoods{}, existsIds).Error
		if err != nil {
			return err
		}
	}
	//插入记录
	err = global.GVA_DB.Create(&goodsList).Error
	return err
}
