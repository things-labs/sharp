package builder

var (
	// BuildDate 编译日期 由外部ldflags指定
	BuildTime = ""
	// GitCommit git提交版本 由外部ldflags指定
	GitCommit = ""
	// Version 版本 由外部ldflags指定
	Version = "v0.0.1 Beta"
	// APIVersion api版本 由外部ldflags指定
	APIVersion = ""
	// Model 型号 由外部ldflags指定
	Model = ""
)
