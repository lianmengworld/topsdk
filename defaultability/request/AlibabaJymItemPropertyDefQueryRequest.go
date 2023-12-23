package request


type AlibabaJymItemPropertyDefQueryRequest struct {
    /*
        二级类目ID     */
    CategoryId  *int64 `json:"category_id" required:"true" `
}

func (s *AlibabaJymItemPropertyDefQueryRequest) SetCategoryId(v int64) *AlibabaJymItemPropertyDefQueryRequest {
    s.CategoryId = &v
    return s
}

func (req *AlibabaJymItemPropertyDefQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CategoryId != nil) {
        paramMap["category_id"] = *req.CategoryId
    }
    return paramMap
}

func (req *AlibabaJymItemPropertyDefQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}