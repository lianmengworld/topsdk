package domain


type AlibabaJymItemExternalGoodsDetailQueryGoodsCategoryDto struct {
    /*
        二级类目名称     */
    SecondCategoryName  *string `json:"second_category_name,omitempty" `

    /*
        二级类目ID     */
    SecondCategoryId  *int64 `json:"second_category_id,omitempty" `

    /*
        一级类目名称     */
    FirstCategoryName  *string `json:"first_category_name,omitempty" `

    /*
        一级类目ID     */
    FirstCategoryId  *int64 `json:"first_category_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsDetailQueryGoodsCategoryDto) SetSecondCategoryName(v string) *AlibabaJymItemExternalGoodsDetailQueryGoodsCategoryDto {
    s.SecondCategoryName = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryGoodsCategoryDto) SetSecondCategoryId(v int64) *AlibabaJymItemExternalGoodsDetailQueryGoodsCategoryDto {
    s.SecondCategoryId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryGoodsCategoryDto) SetFirstCategoryName(v string) *AlibabaJymItemExternalGoodsDetailQueryGoodsCategoryDto {
    s.FirstCategoryName = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryGoodsCategoryDto) SetFirstCategoryId(v int64) *AlibabaJymItemExternalGoodsDetailQueryGoodsCategoryDto {
    s.FirstCategoryId = &v
    return s
}
