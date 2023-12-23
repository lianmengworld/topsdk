package response

import (
    "github.com/lianmengworld/topsdk/defaultability/domain"
)

type AlibabaJymItemGameSeverQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        游戏服务器对象
    */
    Result  []domain.AlibabaJymItemGameSeverQueryGoodsServerDto `json:"result,omitempty" `
    /*
        响应码信息描述
    */
    ExtraErrMsg  string `json:"extra_err_msg,omitempty" `
    /*
        是否请求成功
    */
    Succeed  bool `json:"succeed,omitempty" `
    /*
        响应码
    */
    StateCode  string `json:"state_code,omitempty" `
}
