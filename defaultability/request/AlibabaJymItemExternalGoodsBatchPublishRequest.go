package request

import (
        "github.com/lianmengworld/topsdk/defaultability/domain"
        "github.com/lianmengworld/topsdk/util"
    )

type AlibabaJymItemExternalGoodsBatchPublishRequest struct {
    /*
        商品批量发布接口入参     */
    GoodsPublishCommand  *domain.AlibabaJymItemExternalGoodsBatchPublishGoodsPublishCommandDto `json:"goods_publish_command" required:"true" `
}

func (s *AlibabaJymItemExternalGoodsBatchPublishRequest) SetGoodsPublishCommand(v domain.AlibabaJymItemExternalGoodsBatchPublishGoodsPublishCommandDto) *AlibabaJymItemExternalGoodsBatchPublishRequest {
    s.GoodsPublishCommand = &v
    return s
}

func (req *AlibabaJymItemExternalGoodsBatchPublishRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GoodsPublishCommand != nil) {
        paramMap["goods_publish_command"] = util.ConvertStruct(*req.GoodsPublishCommand)
    }
    return paramMap
}

func (req *AlibabaJymItemExternalGoodsBatchPublishRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}