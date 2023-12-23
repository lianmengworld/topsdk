package domain


type AlibabaJymItemPropertyDefQueryGoodsPropertyDto struct {
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

func (s *AlibabaJymItemPropertyDefQueryGoodsPropertyDto) SetPropertyContentType(v int64) *AlibabaJymItemPropertyDefQueryGoodsPropertyDto {
    s.PropertyContentType = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPropertyDto) SetVisible(v bool) *AlibabaJymItemPropertyDefQueryGoodsPropertyDto {
    s.Visible = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPropertyDto) SetCategoryAttrType(v int64) *AlibabaJymItemPropertyDefQueryGoodsPropertyDto {
    s.CategoryAttrType = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPropertyDto) SetPropertyRule(v AlibabaJymItemPropertyDefQueryPropertyRuleDto) *AlibabaJymItemPropertyDefQueryGoodsPropertyDto {
    s.PropertyRule = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPropertyDto) SetDefaultValue(v string) *AlibabaJymItemPropertyDefQueryGoodsPropertyDto {
    s.DefaultValue = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPropertyDto) SetPropertyGrade(v int64) *AlibabaJymItemPropertyDefQueryGoodsPropertyDto {
    s.PropertyGrade = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPropertyDto) SetAttrUnit(v string) *AlibabaJymItemPropertyDefQueryGoodsPropertyDto {
    s.AttrUnit = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPropertyDto) SetPropertyName(v string) *AlibabaJymItemPropertyDefQueryGoodsPropertyDto {
    s.PropertyName = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPropertyDto) SetOptions(v []AlibabaJymItemPropertyDefQueryGoodsPropertyOptionDto) *AlibabaJymItemPropertyDefQueryGoodsPropertyDto {
    s.Options = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPropertyDto) SetPlaceholder(v string) *AlibabaJymItemPropertyDefQueryGoodsPropertyDto {
    s.Placeholder = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPropertyDto) SetPropertyId(v int64) *AlibabaJymItemPropertyDefQueryGoodsPropertyDto {
    s.PropertyId = &v
    return s
}
func (s *AlibabaJymItemPropertyDefQueryGoodsPropertyDto) SetContentType(v int64) *AlibabaJymItemPropertyDefQueryGoodsPropertyDto {
    s.ContentType = &v
    return s
}
