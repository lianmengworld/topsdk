package domain


type AlibabaJymItemExternalGoodsDetailQueryGamePropertyDto struct {
    /*
        服务器id     */
    ServerId  *int64 `json:"server_id,omitempty" `

    /*
        客户端id     */
    ClientId  *int64 `json:"client_id,omitempty" `

    /*
        平台id     */
    PlatformId  *int64 `json:"platform_id,omitempty" `

    /*
        游戏id     */
    GameId  *int64 `json:"game_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsDetailQueryGamePropertyDto) SetServerId(v int64) *AlibabaJymItemExternalGoodsDetailQueryGamePropertyDto {
    s.ServerId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryGamePropertyDto) SetClientId(v int64) *AlibabaJymItemExternalGoodsDetailQueryGamePropertyDto {
    s.ClientId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryGamePropertyDto) SetPlatformId(v int64) *AlibabaJymItemExternalGoodsDetailQueryGamePropertyDto {
    s.PlatformId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryGamePropertyDto) SetGameId(v int64) *AlibabaJymItemExternalGoodsDetailQueryGamePropertyDto {
    s.GameId = &v
    return s
}
