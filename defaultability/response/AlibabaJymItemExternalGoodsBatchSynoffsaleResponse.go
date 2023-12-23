package response

import (
    "github.com/lianmengworld/topsdk/defaultability/domain"
)

type AlibabaJymItemExternalGoodsBatchSynoffsaleResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        出参
    */
    Result  domain.AlibabaJymItemExternalGoodsBatchSynoffsaleResultDTO `json:"result,omitempty" `
}
