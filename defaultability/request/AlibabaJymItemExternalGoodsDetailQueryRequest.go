package request

import (
        "github.com/lianmengworld/topsdk/defaultability/domain"
        "github.com/lianmengworld/topsdk/util"
    )

type AlibabaJymItemExternalGoodsDetailQueryRequest struct {
    /*
        商品详情查询接口入参     */
    GoodsDetailQuery  *domain.AlibabaJymItemExternalGoodsDetailQueryGoodsDetailQueryDto `json:"goods_detail_query" required:"true" `
}

func (s *AlibabaJymItemExternalGoodsDetailQueryRequest) SetGoodsDetailQuery(v domain.AlibabaJymItemExternalGoodsDetailQueryGoodsDetailQueryDto) *AlibabaJymItemExternalGoodsDetailQueryRequest {
    s.GoodsDetailQuery = &v
    return s
}

func (req *AlibabaJymItemExternalGoodsDetailQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GoodsDetailQuery != nil) {
        paramMap["goods_detail_query"] = util.ConvertStruct(*req.GoodsDetailQuery)
    }
    return paramMap
}

func (req *AlibabaJymItemExternalGoodsDetailQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}