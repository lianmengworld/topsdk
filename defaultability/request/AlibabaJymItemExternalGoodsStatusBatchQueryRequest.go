package request

import (
        "github.com/lianmengworld/topsdk/defaultability/domain"
        "github.com/lianmengworld/topsdk/util"
    )

type AlibabaJymItemExternalGoodsStatusBatchQueryRequest struct {
    /*
        批量查询商品状态接口入参     */
    BatchGoodsStatusQuery  *domain.AlibabaJymItemExternalGoodsStatusBatchQueryBatchGoodsStatusQueryDto `json:"batch_goods_status_query" required:"true" `
}

func (s *AlibabaJymItemExternalGoodsStatusBatchQueryRequest) SetBatchGoodsStatusQuery(v domain.AlibabaJymItemExternalGoodsStatusBatchQueryBatchGoodsStatusQueryDto) *AlibabaJymItemExternalGoodsStatusBatchQueryRequest {
    s.BatchGoodsStatusQuery = &v
    return s
}

func (req *AlibabaJymItemExternalGoodsStatusBatchQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BatchGoodsStatusQuery != nil) {
        paramMap["batch_goods_status_query"] = util.ConvertStruct(*req.BatchGoodsStatusQuery)
    }
    return paramMap
}

func (req *AlibabaJymItemExternalGoodsStatusBatchQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}