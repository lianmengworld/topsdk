package domain


type AlibabaJymItemGameSeverQueryGoodsServerDto struct {
    /*
        游戏ID     */
    GameId  *int64 `json:"game_id,omitempty" `

    /*
        服务器名称     */
    Name  *string `json:"name,omitempty" `

    /*
        服务器ID     */
    Id  *int64 `json:"id,omitempty" `

}

func (s *AlibabaJymItemGameSeverQueryGoodsServerDto) SetGameId(v int64) *AlibabaJymItemGameSeverQueryGoodsServerDto {
    s.GameId = &v
    return s
}
func (s *AlibabaJymItemGameSeverQueryGoodsServerDto) SetName(v string) *AlibabaJymItemGameSeverQueryGoodsServerDto {
    s.Name = &v
    return s
}
func (s *AlibabaJymItemGameSeverQueryGoodsServerDto) SetId(v int64) *AlibabaJymItemGameSeverQueryGoodsServerDto {
    s.Id = &v
    return s
}
