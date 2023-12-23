package request

import (
        "github.com/lianmengworld/topsdk/defaultability/domain"
        "github.com/lianmengworld/topsdk/util"
    )

type AlibabaJymItemExternalGoodsBatchModifypriceRequest struct {
    /*
        商品批量改价接口入参     */
    GoodsPriceModifyCommand  *domain.AlibabaJymItemExternalGoodsBatchModifypriceGoodsPriceModifyCommandDto `json:"goods_price_modify_command" required:"true" `
}

func (s *AlibabaJymItemExternalGoodsBatchModifypriceRequest) SetGoodsPriceModifyCommand(v domain.AlibabaJymItemExternalGoodsBatchModifypriceGoodsPriceModifyCommandDto) *AlibabaJymItemExternalGoodsBatchModifypriceRequest {
    s.GoodsPriceModifyCommand = &v
    return s
}

func (req *AlibabaJymItemExternalGoodsBatchModifypriceRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GoodsPriceModifyCommand != nil) {
        paramMap["goods_price_modify_command"] = util.ConvertStruct(*req.GoodsPriceModifyCommand)
    }
    return paramMap
}

func (req *AlibabaJymItemExternalGoodsBatchModifypriceRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}