package logs

import (
	"anime-community/config"
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var errorLogger *zap.SugaredLogger
var initOnce sync.Once

func Init() {
	initOnce.Do(func() {
		encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			MessageKey:  "msg",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "ts",
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05"))
			},
			CallerKey:    "file",
			EncodeCaller: zapcore.ShortCallerEncoder,
			EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendInt64(int64(d) / 1000000)
			},
		})

		// 不同日志输入不同文件
		infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.InfoLevel
		})

		errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})

		infoWriter := getLogWriter(config.GetServerConfig().LogConfig.FilePath)
		errorWriter := getLogWriter(config.GetServerConfig().LogConfig.ErrFilePath)

		core := zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
		)

		log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
		errorLogger = log.Sugar()
	})
}

func getLogWriter(filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    config.GetServerConfig().LogConfig.MaxSize,    // mb
		MaxBackups: config.GetServerConfig().LogConfig.MaxBackups, // 保留旧文件的最大个数
		MaxAge:     config.GetServerConfig().LogConfig.MaxAge,     // 保留旧文件的最大天数
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Debugf(ctx context.Context, template string, args ...interface{}) {
	errorLogger.Debugf(addTraceId(ctx, template), args...)
}

func Infof(ctx context.Context, template string, args ...interface{}) {
	errorLogger.Infof(addTraceId(ctx, template), args...)
}

func Warnf(ctx context.Context, template string, args ...interface{}) {
	errorLogger.Warnf(addTraceId(ctx, template), args...)
}

func Errorf(ctx context.Context, template string, args ...interface{}) {
	errorLogger.Errorf(addTraceId(ctx, template), args...)
}

func DPanicf(ctx context.Context, template string, args ...interface{}) {
	errorLogger.DPanicf(addTraceId(ctx, template), args...)
}

func Panicf(ctx context.Context, template string, args ...interface{}) {
	errorLogger.Panicf(addTraceId(ctx, template), args...)
}

func Fatalf(ctx context.Context, template string, args ...interface{}) {
	errorLogger.Fatalf(addTraceId(ctx, template), args...)
}

func Sync() {
	errorLogger.Sync()
}

func addTraceId(ctx context.Context, template string) string {
	if traceId, ok := ctx.Value(traceIDKey{}).(string); ok && traceId != "" {
		return traceId + " " + template
	}
	return template
}
