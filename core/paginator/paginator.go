package paginator

// DefaultPageSize 默认页大小
var DefaultPageSize = 10

// Param 分页查询参数
type Param struct {
	PageIndex int `form:"pageIndex" json:"pageIndex"`
	PageSize  int `form:"pageSize" json:"pageSize"`
}

// Info 分页一些信息
type Info struct {
	PageIndex int `json:"pageIndex,omitempty"`
	PageSize  int `json:"pageSize,omitempty"`
	Total     int `json:"total"`
}

// DefaultParam 默认分页参数值
func DefaultParam() Param {
	return Param{1, DefaultPageSize}
}

// Inspect 校验分页查询参数有效性,非法值将设置为默认值,
// 默认分页索引: 第1页
// 默认分页大小: DefaultPageSize
// 可修改DefaultPageSize改变全局默认分页大小或传入参数
func (sf *Param) Inspect(pageSize ...int) *Param {
	if sf.PageIndex <= 0 {
		sf.PageIndex = 1
	}
	if sf.PageSize <= 0 {
		if len(pageSize) > 0 && pageSize[0] > 0 {
			sf.PageSize = pageSize[0]
		} else {
			sf.PageSize = DefaultPageSize
		}
	}
	return sf
}
