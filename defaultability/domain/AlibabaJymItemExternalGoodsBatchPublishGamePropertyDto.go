package domain


type AlibabaJymItemExternalGoodsBatchPublishGamePropertyDto struct {
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

func (s *AlibabaJymItemExternalGoodsBatchPublishGamePropertyDto) SetServerId(v int64) *AlibabaJymItemExternalGoodsBatchPublishGamePropertyDto {
    s.ServerId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGamePropertyDto) SetClientId(v int64) *AlibabaJymItemExternalGoodsBatchPublishGamePropertyDto {
    s.ClientId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGamePropertyDto) SetPlatformId(v int64) *AlibabaJymItemExternalGoodsBatchPublishGamePropertyDto {
    s.PlatformId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGamePropertyDto) SetGameId(v int64) *AlibabaJymItemExternalGoodsBatchPublishGamePropertyDto {
    s.GameId = &v
    return s
}
