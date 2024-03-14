package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"
)

func GetFileLogger() (*zap.Logger, error) {
	eConf := zapcore.EncoderConfig{
		MessageKey: "MESSAGE:",
		LevelKey:   "LEVEL",
		TimeKey:    "TIME",
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format("02 Jan 2006 15:04:05"))
		},
	}
	zap.NewDevelopmentEncoderConfig()
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:         "json",
		EncoderConfig:    eConf,
		Development:      true,
		OutputPaths:      []string{"app-logs.log"},
		ErrorOutputPaths: []string{"errors.json"},
	}
	logger, err := cfg.Build()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer logger.Sync()
	return logger, nil
}

func GetSmartLogger(errorFileName, logFileName string) (*zap.Logger, error) {
	//eConf := zapcore.EncoderConfig{
	//	TimeKey:       "TIME",
	//	MessageKey:    "MESSAGE",
	//	StacktraceKey: "F",
	//	CallerKey:     "someKey",
	//
	//	EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	//		enc.AppendString(t.Format("15:04:05 2 Jan 2006"))
	//	},
	//}
	newConf := zapcore.EncoderConfig{
		MessageKey:     "m",
		LevelKey:       "l",
		TimeKey:        "t",
		NameKey:        "n",
		CallerKey:      "c",
		FunctionKey:    "f",
		StacktraceKey:  "s",
		SkipLineEnding: true,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("15:04:05 2 Jan 2006"))
		},
	}
	testConf := zap.NewDevelopmentEncoderConfig()
	testConf.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("15:04:05 2 Jan 2006"))
	}
	testConf.CallerKey = "F"
	testConf.EncodeCaller = zapcore.ShortCallerEncoder
	errFile, err := os.OpenFile("errors.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return nil, err
	}
	errPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zap.ErrorLevel
	})
	//logFile, err := os.OpenFile("app-logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	//if err != nil {
	//	return nil, err
	//}

	errEncoder := zapcore.NewJSONEncoder(newConf)
	writer := zapcore.AddSync(errFile)
	//logEncoder := zapcore.NewConsoleEncoder(eConf)
	errCore := zapcore.NewCore(errEncoder, writer, errPriority)
	//logCore := zapcore.NewCore(logEncoder, logFile, zap.InfoLevel)
	tee := zapcore.NewTee(errCore)
	logger := zap.New(tee, zap.AddCaller())
	return logger, nil
}

func FileLogger(filename string) *zap.Logger {
	config := zap.NewDevelopmentEncoderConfig()
	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("15:04:05 2 Jan 2006"))
	}
	fileEncoder := zapcore.NewJSONEncoder(config)
	logFile, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
	)

	logger := zap.New(core, zap.AddCaller())

	return logger
}

func GetSmarterLogger() (*zap.Logger, error) {
	errConf := zap.NewDevelopmentEncoderConfig()
	errConf.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("15:04:05 2 Jan 2006"))
	}
	errEncoder := zapcore.NewJSONEncoder(errConf)
	errFile, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return nil, err
	}
	errSync := zapcore.AddSync(errFile)
	errCore := zapcore.NewCore(errEncoder, errSync, zap.ErrorLevel)

	logsConf := zap.NewProductionEncoderConfig()
	logsConf.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("15:04:05 2 Jan 2006"))
	}

	logsEncoder := zapcore.NewConsoleEncoder(logsConf)

	logsFile, err := os.OpenFile("app-logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return nil, err
	}
	logsSync := zapcore.AddSync(logsFile)
	logCore := zapcore.NewCore(logsEncoder, logsSync, zap.InfoLevel)
	tee := zapcore.NewTee(errCore, logCore)
	logger := zap.New(tee, zap.AddCaller())
	return logger, nil

}
