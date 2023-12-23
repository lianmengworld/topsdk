package request


type AlibabaJymItemGameSeverQueryRequest struct {
    /*
        游戏ID     */
    GameId  *int64 `json:"game_id" required:"true" `
    /*
        客户端ID     */
    ClientId  *int64 `json:"client_id" required:"true" `
}

func (s *AlibabaJymItemGameSeverQueryRequest) SetGameId(v int64) *AlibabaJymItemGameSeverQueryRequest {
    s.GameId = &v
    return s
}
func (s *AlibabaJymItemGameSeverQueryRequest) SetClientId(v int64) *AlibabaJymItemGameSeverQueryRequest {
    s.ClientId = &v
    return s
}

func (req *AlibabaJymItemGameSeverQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GameId != nil) {
        paramMap["game_id"] = *req.GameId
    }
    if(req.ClientId != nil) {
        paramMap["client_id"] = *req.ClientId
    }
    return paramMap
}

func (req *AlibabaJymItemGameSeverQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}