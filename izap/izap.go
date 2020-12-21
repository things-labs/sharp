package izap

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config 日志配置
type Config struct {
	Level       string `yaml:"level" json:"level"`             // 日志等级, debug,info,warn,error,dpanic,panic,fatal 默认info
	Format      string `yaml:"format" json:"format"`           // 编码格式: json or console 默认json
	EncodeLevel string `yaml:"encodeLevel" json:"encodeLevel"` // 编码器类型, 默认 LowercaseLevelEncoder
	Writer      string `yaml:"write" json:"write"`             // 输出: file,console,multi 默认 console
	Stack       bool   `yaml:"stack" json:"stack"`             // 使能栈调试输出 , 默认false
	Path        string `yaml:"path" json:"path"`               // 日志存放路径, 默认 empty
	// see lumberjack.Logger
	FileName   string `yaml:"fileName" json:"fileName"`     // 文件名,空字符使用默认    默认<processname>-lumberjack.log
	MaxSize    int    `yaml:"maxSize" json:"maxSize"`       // 每个日志文件最大尺寸(MB) 默认100MB,
	MaxAge     int    `yaml:"maxAge" json:"maxAge"`         // 日志文件保存天数, 默认0不删除
	MaxBackups int    `yaml:"maxBackups" json:"maxBackups"` // 日志文件保存备份数, 默认0都保存
	LocalTime  bool   `yaml:"localTime" json:"localTime"`   // 是否格式化时间戳, 默认UTC时间
	Compress   bool   `yaml:"compress" json:"compress"`     // 压缩文件,采用gzip, 默认不压缩
}

var Logger = zap.NewNop()
var Sugar = Logger.Sugar()

func New(c Config) *zap.Logger {
	var options []zap.Option
	var encoder zapcore.Encoder

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    toEncodeLevel(c.EncodeLevel),
		EncodeTime:     zapcore.ISO8601TimeEncoder, // 修改输出时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if c.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 设置日志输出等级
	level := zap.NewAtomicLevelAt(toLevel(c.Level))
	// 初始化core
	core := zapcore.NewCore(encoder, toWriter(&c), level)

	// 添加显示文件名和行号,跳过封装调用层,栈调用,及使能等级
	if c.Stack {
		stackLevel := zap.NewAtomicLevel()
		stackLevel.SetLevel(zap.WarnLevel) // 只显示栈的错误等级
		options = append(options,
			zap.AddCaller(),
			zap.AddCallerSkip(1),
			zap.AddStacktrace(stackLevel),
		)
	}
	return zap.New(core, options...)
}

func ReplaceGlobals(l *zap.Logger) {
	Logger = l
	Sugar = l.Sugar()
	zap.ReplaceGlobals(Logger)
}

// SetLevel 设置zap默认目志等级,线程安全
func SetLevel(l zapcore.Level) {
	Logger.Core().Enabled(l)
}

func toLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func toEncodeLevel(l string) zapcore.LevelEncoder {
	switch l {
	case "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	case "LowercaseLevelEncoder": // 小写编码器(默认)
		fallthrough
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

func toWriter(c *Config) zapcore.WriteSyncer {
	switch strings.ToLower(c.Writer) {
	case "file":
		return zapcore.AddSync(&lumberjack.Logger{ // 文件切割
			Filename:   filepath.Join(c.Path, c.FileName),
			MaxSize:    c.MaxSize,
			MaxAge:     c.MaxAge,
			MaxBackups: c.MaxBackups,
			LocalTime:  c.LocalTime,
			Compress:   c.Compress,
		})
	case "multi":
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&lumberjack.Logger{ // 文件切割
			Filename:   filepath.Join(c.Path, c.FileName),
			MaxSize:    c.MaxSize,
			MaxAge:     c.MaxAge,
			MaxBackups: c.MaxBackups,
			LocalTime:  c.LocalTime,
			Compress:   c.Compress,
		}))
	case "console":
		fallthrough
	default:
		return zapcore.AddSync(os.Stdout)
	}
}
