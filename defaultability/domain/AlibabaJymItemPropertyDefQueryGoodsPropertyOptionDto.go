package domain


type AlibabaJymItemPropertyDefQueryGoodsPropertyOptionDto struct {
    /*
        选项ID     */
    Id  *int64 `json:"id,omitempty" `

    /*
        选项名称     */
    Name  *string `json:"name,omitempty" `

    /*
        属性ID     */
    AttrId  *int64 `json:"attr_id,omitempty" `

    /*
        属性值     */
    Value  *string `json:"value,omitempty" `

}

func (s *AlibabaJymItemPropertyDefQueryGoodsPropertyOptionDto) SetId(v int64) *AlibabaJymItemPropertyDefQueryGoodsPropertyOptionDto {
    s.Id = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPropertyOptionDto) SetName(v string) *AlibabaJymItemPropertyDefQueryGoodsPropertyOptionDto {
    s.Name = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPropertyOptionDto) SetAttrId(v int64) *AlibabaJymItemPropertyDefQueryGoodsPropertyOptionDto {
    s.AttrId = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPropertyOptionDto) SetValue(v string) *AlibabaJymItemPropertyDefQueryGoodsPropertyOptionDto {
    s.Value = &v
    return s
}
