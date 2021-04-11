import service from '@/utils/request'

// @Summary 获取Goods列表
// @Produce  application/json
// @Param {
//  page  int
//	limit int
// }
// @Router /goods/getGoodsList [get]
export const getMenuList = (data) => {
    return service({
        url: "/goods/getGoodsList",
        method: 'get',
        data
    })
}


