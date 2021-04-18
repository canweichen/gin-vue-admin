import service from '@/utils/request'

// @tags goods
// @Summary 获取商品列表
// @Produce  application/json
// @Param {
//  page  int
//	pageSize int
// }
// @Router /goods/getGoodsList [post]
export const getGoodsList = (data) => {
    return service({
        url: "/goods/getGoodsList",
        method: 'post',
        data
    })
}

// @tags goods
// @Summary 拉取商品
// @Produce  application/json
// @Param {
//  page  int
// }
// @Router /goods/pullGoods [post]
export const pullGoods = (data) => {
    return service({
        url: "/goods/pullGoods",
        method: 'post',
        data
    })
}


