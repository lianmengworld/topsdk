package domain


type AlibabaJymItemPropertyDefQueryPropertyRuleDto struct {
    /*
        价格规则，最大值     */
    HighestPrice  *int64 `json:"highest_price,omitempty" `

    /*
        价格规则，最小值     */
    LowestPrice  *int64 `json:"lowest_price,omitempty" `

    /*
        数字规则，最小值     */
    Min  *int64 `json:"min,omitempty" `

    /*
        数字规则，最大值     */
    Max  *int64 `json:"max,omitempty" `

    /*
        图片规则，最小长度     */
    MinLength  *int64 `json:"min_length,omitempty" `

    /*
        字符串规则，正则表达式     */
    Pattern  *string `json:"pattern,omitempty" `

    /*
        图片规则，单图片最大体积     */
    MaxSize  *int64 `json:"max_size,omitempty" `

    /*
        正则表达式,校验提示     */
    PatternMsg  *string `json:"pattern_msg,omitempty" `

    /*
        图片规则，最大图片数     */
    MaxCount  *int64 `json:"max_count,omitempty" `

    /*
        通用规则，是否必填     */
    Required  *bool `json:"required,omitempty" `

    /*
        字符串规则，最大长度     */
    MaxLength  *int64 `json:"max_length,omitempty" `

}

func (s *AlibabaJymItemPropertyDefQueryPropertyRuleDto) SetHighestPrice(v int64) *AlibabaJymItemPropertyDefQueryPropertyRuleDto {
    s.HighestPrice = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryPropertyRuleDto) SetLowestPrice(v int64) *AlibabaJymItemPropertyDefQueryPropertyRuleDto {
    s.LowestPrice = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryPropertyRuleDto) SetMin(v int64) *AlibabaJymItemPropertyDefQueryPropertyRuleDto {
    s.Min = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryPropertyRuleDto) SetMax(v int64) *AlibabaJymItemPropertyDefQueryPropertyRuleDto {
    s.Max = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryPropertyRuleDto) SetMinLength(v int64) *AlibabaJymItemPropertyDefQueryPropertyRuleDto {
    s.MinLength = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryPropertyRuleDto) SetPattern(v string) *AlibabaJymItemPropertyDefQueryPropertyRuleDto {
    s.Pattern = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryPropertyRuleDto) SetMaxSize(v int64) *AlibabaJymItemPropertyDefQueryPropertyRuleDto {
    s.MaxSize = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryPropertyRuleDto) SetPatternMsg(v string) *AlibabaJymItemPropertyDefQueryPropertyRuleDto {
    s.PatternMsg = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryPropertyRuleDto) SetMaxCount(v int64) *AlibabaJymItemPropertyDefQueryPropertyRuleDto {
    s.MaxCount = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryPropertyRuleDto) SetRequired(v bool) *AlibabaJymItemPropertyDefQueryPropertyRuleDto {
    s.Required = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryPropertyRuleDto) SetMaxLength(v int64) *AlibabaJymItemPropertyDefQueryPropertyRuleDto {
    s.MaxLength = &v
    return s
}
