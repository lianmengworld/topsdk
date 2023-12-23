package domain


type AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto struct {
    /*
        属性内容类型,1:"单选",2:"多选",3:"字符",4:"长字符",5:"数字",6:"日期"     */
    PropertyContentType  *int64 `json:"property_content_type,omitempty" `

    /*
        是否可见     */
    Visible  *bool `json:"visible,omitempty" `

    /*
        属性类型     */
    CategoryAttrType  *int64 `json:"category_attr_type,omitempty" `

    /*
        属性规则对象     */
    PropertyRule  *AlibabaJymItemPropertyDefQueryPropertyRuleDto `json:"property_rule,omitempty" `

    /*
        默认值     */
    DefaultValue  *string `json:"default_value,omitempty" `

    /*
        属性等级     */
    PropertyGrade  *int64 `json:"property_grade,omitempty" `

    /*
        商品属性单位     */
    AttrUnit  *string `json:"attr_unit,omitempty" `

    /*
        属性名称     */
    PropertyName  *string `json:"property_name,omitempty" `

    /*
        属性可选项     */
    Options  *[]AlibabaJymItemPropertyDefQueryGoodsPropertyOptionDto `json:"options,omitempty" `

    /*
        占位符     */
    Placeholder  *string `json:"placeholder,omitempty" `

    /*
        属性ID     */
    PropertyId  *int64 `json:"property_id,omitempty" `

    /*
        内容类型,1:"常规",2:"游戏帐号",3:"游戏密码",4:"QQ号",5:"单价数量",6:"买家服务器",7:"开始发货时间",8:"终止发货时间",9:"原价",10:"联系手机",11:"账户类型",12:"发货服务器",13:"角色名称",14:"备用名称",15:"第三名称"  ,16:"确认帐号",17:"解锁密码",18:"面额",19:"充值类型",20:"回收折扣",21:"帐号绑定",22:"多选"     */
    ContentType  *int64 `json:"content_type,omitempty" `

}

func (s *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto) SetPropertyContentType(v int64) *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto {
    s.PropertyContentType = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto) SetVisible(v bool) *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto {
    s.Visible = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto) SetCategoryAttrType(v int64) *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto {
    s.CategoryAttrType = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto) SetPropertyRule(v AlibabaJymItemPropertyDefQueryPropertyRuleDto) *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto {
    s.PropertyRule = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto) SetDefaultValue(v string) *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto {
    s.DefaultValue = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto) SetPropertyGrade(v int64) *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto {
    s.PropertyGrade = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto) SetAttrUnit(v string) *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto {
    s.AttrUnit = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto) SetPropertyName(v string) *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto {
    s.PropertyName = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto) SetOptions(v []AlibabaJymItemPropertyDefQueryGoodsPropertyOptionDto) *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto {
    s.Options = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto) SetPlaceholder(v string) *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto {
    s.Placeholder = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto) SetPropertyId(v int64) *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto {
    s.PropertyId = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto) SetContentType(v int64) *AlibabaJymItemPropertyDefQueryGoodsPublishPropertyDefDetailDto {
    s.ContentType = &v
    return s
}
