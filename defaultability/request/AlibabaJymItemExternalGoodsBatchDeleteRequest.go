package request

import (
        "github.com/lianmengworld/topsdk/defaultability/domain"
        "github.com/lianmengworld/topsdk/util"
    )

type AlibabaJymItemExternalGoodsBatchDeleteRequest struct {
    /*
        商品批量删除接口入参     */
    GoodsDeleteCommandDto  *domain.AlibabaJymItemExternalGoodsBatchDeleteGoodsDeleteCommandDTO `json:"goods_delete_command_dto" required:"true" `
}

func (s *AlibabaJymItemExternalGoodsBatchDeleteRequest) SetGoodsDeleteCommandDto(v domain.AlibabaJymItemExternalGoodsBatchDeleteGoodsDeleteCommandDTO) *AlibabaJymItemExternalGoodsBatchDeleteRequest {
    s.GoodsDeleteCommandDto = &v
    return s
}

func (req *AlibabaJymItemExternalGoodsBatchDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GoodsDeleteCommandDto != nil) {
        paramMap["goods_delete_command_dto"] = util.ConvertStruct(*req.GoodsDeleteCommandDto)
    }
    return paramMap
}

func (req *AlibabaJymItemExternalGoodsBatchDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}