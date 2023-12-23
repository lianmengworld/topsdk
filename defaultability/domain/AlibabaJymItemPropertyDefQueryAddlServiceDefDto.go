package domain


type AlibabaJymItemPropertyDefQueryAddlServiceDefDto struct {
    /*
        是否支持议价     */
    Bargain  *bool `json:"bargain,omitempty" `

}

func (s *AlibabaJymItemPropertyDefQueryAddlServiceDefDto) SetBargain(v bool) *AlibabaJymItemPropertyDefQueryAddlServiceDefDto {
    s.Bargain = &v
    return s
}
