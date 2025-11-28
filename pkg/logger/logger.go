package logger

import (
	"log"
	"os"

	constants "github.com/nielchaudhary/compass/pkg/constants"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	globalLogger *zap.Logger
	sugarLogger  *zap.SugaredLogger
)

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var colorCode string
	switch level {
	case zapcore.DebugLevel:
		colorCode = "\033[36m" // Cyan
	case zapcore.InfoLevel:
		colorCode = "\033[32m" // Green
	case zapcore.WarnLevel:
		colorCode = "\033[33m" // Yellow
	case zapcore.ErrorLevel:
		colorCode = "\033[31m" // Red
	case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		colorCode = "\033[35m" // Magenta
	default:
		colorCode = "\033[0m" // Reset
	}
	enc.AppendString(colorCode + level.CapitalString() + "\033[0m")
}

func newColoredDevelopmentLogger() (*zap.Logger, error) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    customLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)), nil
}

func InitLogger(env string) error {
	var logger *zap.Logger
	var err error

	switch env {
	case string(constants.Development), string(constants.Testing):
		logger, err = newColoredDevelopmentLogger()
	case string(constants.Production):
		// Production typically goes to files/log aggregators, so no colors
		logger, err = zap.NewProduction()
	default:
		logger, err = newColoredDevelopmentLogger()
	}

	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
		return err
	}

	globalLogger = logger
	sugarLogger = logger.Sugar()

	return nil
}

func GetLogger() *zap.SugaredLogger {
	if sugarLogger == nil {
		_ = InitLogger(string(constants.Development))
	}
	return sugarLogger
}

func Sync() error {
	if globalLogger != nil {
		err := globalLogger.Sync()
		// Ignore sync errors for stdout/stderr on terminal devices
		// This is expected behavior - fsync doesn't work on TTY devices
		if err != nil {
			// Check for the specific "inappropriate ioctl for device" error
			if osErr, ok := err.(*os.PathError); ok {
				// Ignore sync errors for /dev/stdout and /dev/stderr
				if osErr.Path == "/dev/stdout" || osErr.Path == "/dev/stderr" {
					return nil
				}
			}
		}
		return err
	}
	return nil
}
