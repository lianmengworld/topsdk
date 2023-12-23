package request

import (
        "github.com/lianmengworld/topsdk/defaultability/domain"
        "github.com/lianmengworld/topsdk/util"
    )

type AlibabaJymItemExternalGoodsBatchOnsaleRequest struct {
    /*
        商品批量上架接口入参     */
    GoodsOnSaleCommand  *domain.AlibabaJymItemExternalGoodsBatchOnsaleGoodsOnSaleCommandDto `json:"goods_on_sale_command" required:"true" `
}

func (s *AlibabaJymItemExternalGoodsBatchOnsaleRequest) SetGoodsOnSaleCommand(v domain.AlibabaJymItemExternalGoodsBatchOnsaleGoodsOnSaleCommandDto) *AlibabaJymItemExternalGoodsBatchOnsaleRequest {
    s.GoodsOnSaleCommand = &v
    return s
}

func (req *AlibabaJymItemExternalGoodsBatchOnsaleRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GoodsOnSaleCommand != nil) {
        paramMap["goods_on_sale_command"] = util.ConvertStruct(*req.GoodsOnSaleCommand)
    }
    return paramMap
}

func (req *AlibabaJymItemExternalGoodsBatchOnsaleRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}