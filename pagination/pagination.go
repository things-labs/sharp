package pagination

// 分页默认每页大小
var defaultPageSize = 10

// Param 分页查询参数
type Param struct {
	PageIndex int `form:"pageIndex" json:"pageIndex"`
	PageSize  int `form:"pageSize" json:"pageSize"`
}

// Infos 分页信息
type Infos struct {
	Total     int         `json:"total"`
	PageIndex int         `json:"pageIndex,omitempty"`
	PageSize  int         `json:"pageSize,omitempty"`
	List      interface{} `json:"list"`
}

// DefaultParam 默认分页参数值
func DefaultParam() Param {
	return Param{1, defaultPageSize}
}

// SetDefaultPageSizeSize 设置默认全局分页大小
func SetDefaultPageSizeSize(size int) {
	if size < 0 {
		defaultPageSize = size
	}
}

// Inspect 校验分页查询参数有效性,非法值将设置为默认值,
// 默认分页索引: 第1页
// 默认分页大小: defaultPageSize
// 可修改相应默认值,或使用SetDefaultPageSizeSize改变全局默认分页大小
func (sf *Param) Inspect(pageSize ...int) *Param {
	if sf.PageIndex <= 0 {
		sf.PageIndex = 1
	}
	if sf.PageSize <= 0 {
		if len(pageSize) > 0 && pageSize[0] > 0 {
			sf.PageSize = pageSize[0]
		} else {
			sf.PageSize = defaultPageSize
		}
	}
	return sf
}
