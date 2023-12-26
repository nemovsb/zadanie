package zaplogger

import (
	"errors"
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	Production  = "prod"
	Development = "dev"
	None        = "none"
)

const payloadKey = "payload"

var ErrUnsupportedMode = errors.New("unsupported zapLogger mode")

func Provider(mode string) (logger *zap.Logger, cleanup func(), err error) {
	var zapLogger *zap.Logger

	switch mode {
	case Development:
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zapLogger, err = config.Build()
	case Production:
		config := zap.NewProductionConfig()
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		zapLogger, err = config.Build()
	case None:
		zapLogger = zap.NewNop()
	default:
		err = fmt.Errorf("%w: %s", ErrUnsupportedMode, mode)
	}

	if err != nil {
		return nil, nil, err
	}

	undoRedirectStdLog := zap.RedirectStdLog(zapLogger)
	cleanup = func() {
		if errSync := zapLogger.Sync(); errSync != nil && errSync.Error() != "sync /dev/stderr: invalid argument" {
			log.Println("can't sync zap logger", errSync)
		}

		undoRedirectStdLog()
	}

	zapLogger = zapLogger.WithOptions(zap.AddCallerSkip(0)).With(zap.Namespace(payloadKey))

	return zapLogger, cleanup, nil
}
