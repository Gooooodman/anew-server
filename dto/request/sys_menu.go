package request

// 获取菜单列表结构体
type MenuListReq struct {
	Name       string `json:"name" form:"name"`
	Title      string `json:"title" form:"title"`
	Path       string `json:"path" form:"path"`
	Component  string `json:"component" form:"component"`
	Redirect   string `json:"redirect"`
	Status     *bool  `json:"status" form:"status"`
	Visible    *bool  `json:"visible" form:"visible"`
	Breadcrumb *bool  `json:"breadcrumb" form:"breadcrumb"`
	Creator    string `json:"creator" form:"creator"`
	PageInfo          // 分页参数
}

// 创建菜单结构体
type CreateMenuReq struct {
	Name       string `json:"name" validate:"required"`
	Title      string `json:"title"`
	Icon       string `json:"icon"`
	Path       string `json:"path"`
	Redirect   string `json:"redirect"`
	Component  string `json:"component"`
	Permission string `json:"permission"`
	Sort       int    `json:"sort"`
	Status     *bool  `json:"status"`
	Visible    *bool  `json:"visible"`
	Breadcrumb *bool  `json:"breadcrumb"`
	ParentId   uint   `json:"parentId"`
	Creator    string `json:"creator"`
}

// 翻译需要校验的字段名称
func (s CreateMenuReq) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["Name"] = "菜单名称"
	return m
}
