package data

import (
	"fiber-demo/app/demo/internal/conf"
	"fiber-demo/pkgs/utils"
	"fmt"
	kzap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"go.uber.org/zap"
	"os"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"go.uber.org/zap/zapcore"

	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDiscovery,
	NewRegistrar,
	NewLogger,
	NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewLogger(conf *conf.Logger) log.Logger {
	if ok, _ := utils.PathExists(conf.Directory); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", conf.Directory)
		_ = os.Mkdir(conf.Directory, os.ModePerm)
	}
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	cores := [...]zapcore.Core{
		getEncoderCore(conf, fmt.Sprintf("./%s/server_debug.log", conf.Directory), debugPriority),
		getEncoderCore(conf, fmt.Sprintf("./%s/server_info.log", conf.Directory), infoPriority),
		getEncoderCore(conf, fmt.Sprintf("./%s/server_warn.log", conf.Directory), warnPriority),
		getEncoderCore(conf, fmt.Sprintf("./%s/server_error.log", conf.Directory), errorPriority),
	}
	zLog := zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	if conf.ShowLine {
		zLog = zLog.WithOptions(zap.AddCaller())
	}
	id, _ := os.Hostname()
	return log.With(kzap.NewLogger(zLog),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", conf.Prefix,
	)
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig(conf *conf.Logger) (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  conf.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case conf.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case conf.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case conf.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case conf.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder(conf *conf.Logger) zapcore.Encoder {
	if conf.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig(conf))
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig(conf))
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(conf *conf.Logger, fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := utils.GetWriteSyncer(fileName, conf.LogInConsole) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(getEncoder(conf), writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}
