package request

import (
        "github.com/lianmengworld/topsdk/defaultability/domain"
        "github.com/lianmengworld/topsdk/util"
    )

type AlibabaJymItemExternalGoodsBatchOffsaleRequest struct {
    /*
        商品批量下架接口入参     */
    GoodsOffSaleCommand  *domain.AlibabaJymItemExternalGoodsBatchOffsaleGoodsOffSaleCommandDto `json:"goods_off_sale_command" required:"true" `
}

func (s *AlibabaJymItemExternalGoodsBatchOffsaleRequest) SetGoodsOffSaleCommand(v domain.AlibabaJymItemExternalGoodsBatchOffsaleGoodsOffSaleCommandDto) *AlibabaJymItemExternalGoodsBatchOffsaleRequest {
    s.GoodsOffSaleCommand = &v
    return s
}

func (req *AlibabaJymItemExternalGoodsBatchOffsaleRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GoodsOffSaleCommand != nil) {
        paramMap["goods_off_sale_command"] = util.ConvertStruct(*req.GoodsOffSaleCommand)
    }
    return paramMap
}

func (req *AlibabaJymItemExternalGoodsBatchOffsaleRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}