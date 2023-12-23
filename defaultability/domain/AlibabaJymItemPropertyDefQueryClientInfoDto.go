package domain


type AlibabaJymItemPropertyDefQueryClientInfoDto struct {
    /*
        平台ID,2:"安卓",3:"苹果",4,:"越狱",1:"其他",5:"PC",6:"XBOX - 主机",7:多端     */
    PlatformId  *int64 `json:"platform_id,omitempty" `

    /*
        客户端ID     */
    Id  *int64 `json:"id,omitempty" `

    /*
        平台ID数组，支持多端互通的客户端必须取该字段中的元素     */
    PlatformIds  *[]int32 `json:"platform_ids,omitempty" `

}

func (s *AlibabaJymItemPropertyDefQueryClientInfoDto) SetPlatformId(v int64) *AlibabaJymItemPropertyDefQueryClientInfoDto {
    s.PlatformId = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryClientInfoDto) SetId(v int64) *AlibabaJymItemPropertyDefQueryClientInfoDto {
    s.Id = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryClientInfoDto) SetPlatformIds(v []int32) *AlibabaJymItemPropertyDefQueryClientInfoDto {
    s.PlatformIds = &v
    return s
}
