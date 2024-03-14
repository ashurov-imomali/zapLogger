package main

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"main/logger"
	"os"
	"strconv"
	"time"
)

func main() {

	smartLogger, err := logger.GetSmarterLogger()
	if err != nil {
		log.Println(err)
		return
	}
	defer smartLogger.Sync()
	str := `24f`
	_, err = strconv.Atoi(str)
	if err != nil {
		smartLogger.Error("Can't Parse 2 int", zap.Error(err))
		return
	}
	smartLogger.Error("Can't parse 2 json.", zap.Error(errors.New("text")))
	smartLogger.Error("Can't parse 2 json.", zap.Error(errors.New("text")))
	smartLogger.Error("Can't parse 2 json.", zap.Error(errors.New("text")))
	smartLogger.Info("Start Deleting....")
	return

	eConf := zapcore.EncoderConfig{
		MessageKey: "MESSAGE:",
		LevelKey:   "LEVEL",
		TimeKey:    "TIME",
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format("02 Jan 2006 15:04:05"))
		},
	}
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.ErrorLevel),
		Encoding:         "json",
		EncoderConfig:    eConf,
		Development:      true,
		OutputPaths:      []string{"app-logs.log"},
		ErrorOutputPaths: []string{"errors.txt"},
	}
	logger, err := cfg.Build()
	if err != nil {
		log.Println(err)
	}
	zapcore.NewCore(zapcore.NewJSONEncoder(eConf), os.Stdout, zap.InfoLevel)
	defer logger.Sync()
	logger.Error("Can't parse 2 json.", zap.Error(errors.New("text")))
	logger.Info("Start parsing ...")
	//warnLevel := zap.NewAtomicLevelAt(zap.WarnLevel)
	//cfg := zap.Config{
	//	Level:         warnLevel,
	//	Encoding:      "json",
	//	OutputPaths:   []string{"stdout"},
	//	EncoderConfig: zap.NewDevelopmentEncoderConfig(),
	//}
	//zap.NewDevelopmentEncoderConfig()
	//zap.Time
	//build, err := cfg.Build()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//build.Debug("newMessage", zap.String("s", "s"))
	//warnLevel.SetLevel(zap.DebugLevel)
	//build.Debug("oldMessage", zap.String("s", "s"))
	//logger, err := zap.NewProduction()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//logger.Error("msg", zap.Error(errors.New("text")))
}
