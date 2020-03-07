package pagination

// 分页默认每页大小
const (
	DefaultPageSize = 100
)

// Param 分页查询参数
type Param struct {
	PageIndex int `form:"pagenum" json:"pageIndex"`
	PageSize  int `form:"pagesize" json:"pageSize"`
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
	return Param{1, DefaultPageSize}
}

// Inspect 校验分页查询参数有效性,非法值将设置为默认值,
// 默认第1页,
// 默认分大小DefaultPageSize,可修改相应默认值
func (sf *Param) Inspect(pageSize ...int) *Param {
	if sf.PageIndex <= 0 {
		sf.PageIndex = 1
	}
	if sf.PageSize <= 0 {
		sf.PageSize = append(pageSize, DefaultPageSize)[0]
	}
	return sf
}
