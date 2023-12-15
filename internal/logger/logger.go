package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"openapi/internal/config"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	log := zap.New(core, zap.AddCaller())
	Logger = log.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {

	//dir := filepath.Dir(config.CFG.Log.LogFile)
	//if _, err := os.Stat(dir); os.IsNotExist(err) {
	//	err := os.MkdirAll(dir, 0666)
	//	if err != nil {
	//		panic(fmt.Errorf("Fatal error mkdir: %s \n", err))
	//	}
	//}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.CFG.Log.LogFile,
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}

	//file, _ := os.OpenFile(config.CFG.Log.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	return zapcore.AddSync(lumberJackLogger)
}
