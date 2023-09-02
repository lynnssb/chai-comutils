/**
 * @author:       wangxuebing
 * @fileName:     pagehandler.go
 * @date:         2023/1/14 16:37
 * @description:
 */

package core

import (
	"chai-comutils/common/utils/characterutil"
	"strings"
)

type PageResult struct {
	Page     int     //页码
	PageSize int     //每页数量
	Filter   *string //关键字
}

func PageHandle(page, pageSize int, filter *string) PageResult {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	} else if pageSize >= 1000 {
		pageSize = 1000
	}
	if filter != nil {
		_filter := strings.TrimSpace(*filter)
		if len(_filter) > 0 {
			_f := characterutil.StitchingBuilderStr("%", _filter, "%")
			filter = &_f
		}
	}
	return PageResult{
		Page:     (page - 1) * pageSize,
		PageSize: pageSize,
		Filter:   filter,
	}
}

func Filter(filter string) string {
	_filter := strings.TrimSpace(filter)
	if len(_filter) > 0 {
		return characterutil.StitchingBuilderStr("%", _filter, "%")
	}
	return ""
}
