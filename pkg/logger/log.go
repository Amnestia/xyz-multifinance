package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

type leveledWriter struct {
	io.Writer
	level map[zerolog.Level]bool
}

func (lw *leveledWriter) WriteLevel(lv zerolog.Level, p []byte) (n int, err error) {
	if lw.level[lv] {
		return lw.Writer.Write(p)
	}
	return len(p), nil
}

func ensureDir(path string) (err error) {
	d := strings.Split(path, "/")
	d = d[:len(d)-1]
	dir := strings.Join(d, "/")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModeDir)
		if err != nil {
			return fmt.Errorf("failed to create dir %s : %s", dir, err)
		}
	}
	return nil
}

func InitLogger(serviceName, infoLog, errorLog string) (err error) {
	infoPath := filepath.Clean(infoLog)
	if err = ensureDir(infoPath); err != nil {
		return err
	}
	errorPath := filepath.Clean(errorLog)
	if err = ensureDir(errorPath); err != nil {
		return err
	}
	fileInfo, err := os.OpenFile(infoPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open %s : %s", infoLog, err)
	}
	fileError, err := os.OpenFile(errorPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open %s : %s", errorLog, err)
	}
	w := zerolog.MultiLevelWriter(
		zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		},
		leveledWriter{
			zerolog.MultiLevelWriter(fileInfo),
			map[zerolog.Level]bool{
				zerolog.InfoLevel:  true,
				zerolog.DebugLevel: true,
			},
		},
		leveledWriter{
			zerolog.MultiLevelWriter(fileError),
			map[zerolog.Level]bool{
				zerolog.WarnLevel:  true,
				zerolog.ErrorLevel: true,
				zerolog.PanicLevel: true,
				zerolog.FatalLevel: true,
			},
		},
	)
	zerolog.TimeFieldFormat = time.RFC3339
	Logger = zerolog.New(w).With().Timestamp().Logger()
	return
}

// Error log error helper to show context
// ex: [context] msg : error
func ErrorWrap(err error, context string, msg ...string) error {
	context = strings.TrimSpace(context)
	if len(context) > 0 {
		context = fmt.Sprintf("[%s]", context)
	}
	message := fmt.Sprintf("%v %v : %v", context, strings.Join(msg, ", "), err)
	return errors.Wrap(err, message)
}
