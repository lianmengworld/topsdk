package defaultability

import (
    "log"
    "errors"
    "github.com/lianmengworld/topsdk"
    "github.com/lianmengworld/topsdk/defaultability/request"
    "github.com/lianmengworld/topsdk/defaultability/response"
    "github.com/lianmengworld/topsdk/util"
)

type Defaultability struct {
    Client *topsdk.TopClient
}

func NewDefaultability(client *topsdk.TopClient) *Defaultability{
    return &Defaultability{client}
}

/*
    交易猫外部商家批量商品改价接口
*/
func (ability *Defaultability) AlibabaJymItemExternalGoodsBatchModifyprice(req *request.AlibabaJymItemExternalGoodsBatchModifypriceRequest) (*response.AlibabaJymItemExternalGoodsBatchModifypriceResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.jym.item.external.goods.batch.modifyprice",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaJymItemExternalGoodsBatchModifypriceResponse{}
    if(err != nil){
        log.Println("alibabaJymItemExternalGoodsBatchModifyprice error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    交易猫外部商家批量上架商品接口
*/
func (ability *Defaultability) AlibabaJymItemExternalGoodsBatchOnsale(req *request.AlibabaJymItemExternalGoodsBatchOnsaleRequest) (*response.AlibabaJymItemExternalGoodsBatchOnsaleResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.jym.item.external.goods.batch.onsale",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaJymItemExternalGoodsBatchOnsaleResponse{}
    if(err != nil){
        log.Println("alibabaJymItemExternalGoodsBatchOnsale error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    交易猫外部商家查询商品批量任务接口
*/
func (ability *Defaultability) AlibabaJymItemExternalGoodsBatchtaskQuery(req *request.AlibabaJymItemExternalGoodsBatchtaskQueryRequest) (*response.AlibabaJymItemExternalGoodsBatchtaskQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.jym.item.external.goods.batchtask.query",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaJymItemExternalGoodsBatchtaskQueryResponse{}
    if(err != nil){
        log.Println("alibabaJymItemExternalGoodsBatchtaskQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    关键词过滤匹配
*/
func (ability *Defaultability) TaobaoKfcKeywordSearch(req *request.TaobaoKfcKeywordSearchRequest,session string) (*response.TaobaoKfcKeywordSearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.kfc.keyword.search",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoKfcKeywordSearchResponse{}
    if(err != nil){
        log.Println("taobaoKfcKeywordSearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    交易猫外部商家商品详情查询接口
*/
func (ability *Defaultability) AlibabaJymItemExternalGoodsDetailQuery(req *request.AlibabaJymItemExternalGoodsDetailQueryRequest) (*response.AlibabaJymItemExternalGoodsDetailQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.jym.item.external.goods.detail.query",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaJymItemExternalGoodsDetailQueryResponse{}
    if(err != nil){
        log.Println("alibabaJymItemExternalGoodsDetailQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    交易猫外部商家商品状态批量查询接口
*/
func (ability *Defaultability) AlibabaJymItemExternalGoodsStatusBatchQuery(req *request.AlibabaJymItemExternalGoodsStatusBatchQueryRequest) (*response.AlibabaJymItemExternalGoodsStatusBatchQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.jym.item.external.goods.status.batch.query",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaJymItemExternalGoodsStatusBatchQueryResponse{}
    if(err != nil){
        log.Println("alibabaJymItemExternalGoodsStatusBatchQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    交易猫外部商家批量下架商品接口
*/
func (ability *Defaultability) AlibabaJymItemExternalGoodsBatchOffsale(req *request.AlibabaJymItemExternalGoodsBatchOffsaleRequest) (*response.AlibabaJymItemExternalGoodsBatchOffsaleResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.jym.item.external.goods.batch.offsale",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaJymItemExternalGoodsBatchOffsaleResponse{}
    if(err != nil){
        log.Println("alibabaJymItemExternalGoodsBatchOffsale error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    交易猫外部商家批量发布商品接口
*/
func (ability *Defaultability) AlibabaJymItemExternalGoodsBatchPublish(req *request.AlibabaJymItemExternalGoodsBatchPublishRequest) (*response.AlibabaJymItemExternalGoodsBatchPublishResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.jym.item.external.goods.batch.publish",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaJymItemExternalGoodsBatchPublishResponse{}
    if(err != nil){
        log.Println("alibabaJymItemExternalGoodsBatchPublish error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    同步下架接口
*/
func (ability *Defaultability) AlibabaJymItemExternalGoodsBatchSynoffsale(req *request.AlibabaJymItemExternalGoodsBatchSynoffsaleRequest) (*response.AlibabaJymItemExternalGoodsBatchSynoffsaleResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.jym.item.external.goods.batch.synoffsale",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaJymItemExternalGoodsBatchSynoffsaleResponse{}
    if(err != nil){
        log.Println("alibabaJymItemExternalGoodsBatchSynoffsale error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    交易猫商品属性定义查询
*/
func (ability *Defaultability) AlibabaJymItemPropertyDefQuery(req *request.AlibabaJymItemPropertyDefQueryRequest) (*response.AlibabaJymItemPropertyDefQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.jym.item.property.def.query",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaJymItemPropertyDefQueryResponse{}
    if(err != nil){
        log.Println("alibabaJymItemPropertyDefQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询商品发布客户端下可用服务器列表
*/
func (ability *Defaultability) AlibabaJymItemGameSeverQuery(req *request.AlibabaJymItemGameSeverQueryRequest) (*response.AlibabaJymItemGameSeverQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.jym.item.game.sever.query",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaJymItemGameSeverQueryResponse{}
    if(err != nil){
        log.Println("alibabaJymItemGameSeverQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    交易猫外部商家批量删除商品接口
*/
func (ability *Defaultability) AlibabaJymItemExternalGoodsBatchDelete(req *request.AlibabaJymItemExternalGoodsBatchDeleteRequest) (*response.AlibabaJymItemExternalGoodsBatchDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.jym.item.external.goods.batch.delete",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaJymItemExternalGoodsBatchDeleteResponse{}
    if(err != nil){
        log.Println("alibabaJymItemExternalGoodsBatchDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
