package response

import (
	"anew-server/models"
)

// 接口信息响应, 字段含义见models
type ApiListResp struct {
	Id        uint             `json:"id"`
	Method    string           `json:"method"`
	Path      string           `json:"path"`
	Category  string           `json:"category"`
	Creator   string           `json:"creator"`
	Desc      string           `json:"desc"`
	Title     string           `json:"title"`
	CreatedAt models.LocalTime `json:"createdAt"`
}

// 权限接口信息响应, 字段含义见models
type ApiGroupByCategoryResp struct {
	Title    string        `json:"title"`    // 标题
	Category string        `json:"category"` // 分组名称
	Children []ApiListResp `json:"children"` // 前端以树形图结构展示, 这里用children表示
}

// 接口树信息响应, 包含有权限访问的id列表
type ApiTreeWithAccessResponseStruct struct {
	List      []ApiGroupByCategoryResp `json:"list"`
	AccessIds []uint                   `json:"accessIds"`
}
