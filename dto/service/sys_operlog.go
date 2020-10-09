package service

import (
	"anew-server/dto/request"
	"anew-server/models"
	"anew-server/pkg/common"
	"fmt"
	"strings"
)

// 获取操作日志
func (s *MysqlService) GetOperLogs(req *request.OperLogListReq) ([]models.SysOperLog, error) {
	var err error
	list := make([]models.SysOperLog, 0)
	query := common.Mysql
	method := strings.TrimSpace(req.Method)
	if method != "" {
		query = query.Where("method LIKE ?", fmt.Sprintf("%%%s%%", method))
	}
	path := strings.TrimSpace(req.Path)
	if path != "" {
		query = query.Where("path LIKE ?", fmt.Sprintf("%%%s%%", path))
	}
	ip := strings.TrimSpace(req.Ip)
	if ip != "" {
		query = query.Where("ip LIKE ?", fmt.Sprintf("%%%s%%", ip))
	}
	status := strings.TrimSpace(req.Status)
	if status != "" {
		query = query.Where("status LIKE ?", fmt.Sprintf("%%%s%%", status))
	}
	query = query.Order("id DESC")
	// 查询条数
	err = query.Find(&list).Count(&req.PageInfo.Total).Error
	if err == nil {
		if req.PageInfo.All {
			// 不使用分页
			err = query.Find(&list).Error
		} else {
			// 获取分页参数
			limit, offset := req.GetLimit()
			err = query.Limit(limit).Offset(offset).Find(&list).Error
		}
	}
	return list, err
}

// 批量删除操作日志
func (s *MysqlService) DeleteOperationLogByIds(ids []uint) (err error) {
	return s.tx.Where("id IN (?)", ids).Delete(models.SysOperLog{}).Error
}
