package response

import (
    "github.com/lianmengworld/topsdk/defaultability/domain"
)

type AlibabaJymItemExternalGoodsBatchOnsaleResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品批量上架接口返回
    */
    Result  domain.AlibabaJymItemExternalGoodsBatchOnsaleGoodsBatchResultDto `json:"result,omitempty" `
    /*
        是否成功
    */
    Succeed  bool `json:"succeed,omitempty" `
    /*
        状态码
    */
    StateCode  string `json:"state_code,omitempty" `
    /*
        错误信息
    */
    ExtraErrMsg  string `json:"extra_err_msg,omitempty" `
}
