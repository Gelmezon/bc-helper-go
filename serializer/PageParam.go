package serializer

type PageParam struct {
	Page     int `form:"page" json:"page,omitempty"`
	PageSize int `form:"pageSize" json:"pageSize,omitempty"`
}

func CreatePage(page, pageSize int) *PageParam {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	return &PageParam{
		Page:     page,
		PageSize: pageSize,
	}
}

func (p PageParam) Begin() (res int) {
	res = (p.Page - 1) * p.PageSize
	return
}

func (p PageParam) End() (res int) {
	res = p.Page * p.PageSize
	return
}
