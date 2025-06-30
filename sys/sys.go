package sys

import (
	"io"
	"os"

	"github.com/ADM87/ggame/sys/logger"
)

// =======================================================================
// System Variables
// =======================================================================

var (
	sysLogger = logger.NewLoggerWith(
		logger.LevelWarn|logger.LevelError,
		map[logger.LogLevel]string{
			logger.LevelDebug: logger.Gray("[DEBUG] "),
			logger.LevelInfo:  logger.Blue("[INFO] "),
			logger.LevelWarn:  logger.Yellow("[WARN] "),
			logger.LevelError: logger.Red("[ERROR] "),
		},
		map[logger.LogLevel]string{
			logger.LevelDebug: logger.StyleGray,
			logger.LevelInfo:  logger.StyleBlue,
			logger.LevelWarn:  logger.StyleYellow,
			logger.LevelError: logger.StyleRed,
		},
		map[logger.LogLevel]io.Writer{
			logger.LevelDebug: os.Stdout,
			logger.LevelInfo:  os.Stdout,
			logger.LevelWarn:  os.Stdout,
			logger.LevelError: os.Stderr,
		},
	)

	sysVersion = "0.0.0-unreleased" // This should be set by the build system, e.g., using ldflags
)

// =======================================================================
// Logger
// =======================================================================

func Logger() logger.Logger {
	return sysLogger
}

func SetLogger(l logger.Logger) {
	sysLogger = l
}

// =======================================================================
// Version
// =======================================================================

func Version() string {
	return sysVersion
}

func SetVersion(v string) {
	sysVersion = v
}

// =======================================================================
// System Methods
// =======================================================================

func Shutdown() {
	Logger().Info("Shutting down...")
	os.Exit(0)
}

func ShutdownWith(code int) {
	Logger().Infof("Shutting down with exit code %d...", code)
	os.Exit(code)
}
