package request

import (
        "github.com/lianmengworld/topsdk/defaultability/domain"
        "github.com/lianmengworld/topsdk/util"
    )

type AlibabaJymItemExternalGoodsBatchtaskQueryRequest struct {
    /*
        查询商品批量任务接口入参     */
    GoodsBatchTaskQuery  *domain.AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchTaskQueryDto `json:"goods_batch_task_query" required:"true" `
}

func (s *AlibabaJymItemExternalGoodsBatchtaskQueryRequest) SetGoodsBatchTaskQuery(v domain.AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchTaskQueryDto) *AlibabaJymItemExternalGoodsBatchtaskQueryRequest {
    s.GoodsBatchTaskQuery = &v
    return s
}

func (req *AlibabaJymItemExternalGoodsBatchtaskQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GoodsBatchTaskQuery != nil) {
        paramMap["goods_batch_task_query"] = util.ConvertStruct(*req.GoodsBatchTaskQuery)
    }
    return paramMap
}

func (req *AlibabaJymItemExternalGoodsBatchtaskQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}