package request

import (
        "github.com/lianmengworld/topsdk/defaultability/domain"
        "github.com/lianmengworld/topsdk/util"
    )

type AlibabaJymItemExternalGoodsBatchSynoffsaleRequest struct {
    /*
        入参     */
    ParamSyncOffSaleCommandDTO  *domain.AlibabaJymItemExternalGoodsBatchSynoffsaleSyncOffSaleCommandDTO `json:"param_sync_off_sale_command_d_t_o" required:"true" `
}

func (s *AlibabaJymItemExternalGoodsBatchSynoffsaleRequest) SetParamSyncOffSaleCommandDTO(v domain.AlibabaJymItemExternalGoodsBatchSynoffsaleSyncOffSaleCommandDTO) *AlibabaJymItemExternalGoodsBatchSynoffsaleRequest {
    s.ParamSyncOffSaleCommandDTO = &v
    return s
}

func (req *AlibabaJymItemExternalGoodsBatchSynoffsaleRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ParamSyncOffSaleCommandDTO != nil) {
        paramMap["param_sync_off_sale_command_d_t_o"] = util.ConvertStruct(*req.ParamSyncOffSaleCommandDTO)
    }
    return paramMap
}

func (req *AlibabaJymItemExternalGoodsBatchSynoffsaleRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}