package domain


type AlibabaJymItemExternalGoodsBatchPublishChildPropertyKeyValueDTO struct {
    /*
        属性键     */
    Key  *string `json:"key,omitempty" `

    /*
        属性值     */
    Value  *string `json:"value,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchPublishChildPropertyKeyValueDTO) SetKey(v string) *AlibabaJymItemExternalGoodsBatchPublishChildPropertyKeyValueDTO {
    s.Key = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishChildPropertyKeyValueDTO) SetValue(v string) *AlibabaJymItemExternalGoodsBatchPublishChildPropertyKeyValueDTO {
    s.Value = &v
    return s
}
