package v2ray

import (
	"log"
	"os"

	v2commlog "github.com/v2fly/v2ray-core/v5/common/log"
)

// This struct creates our own log writer without datatime stamp
// As Android adds time stamps on each line
type consoleLogWriter struct {
	logger *log.Logger
}

func (w *consoleLogWriter) Write(s string) error {
	w.logger.Print(s)
	return nil
}

func (w *consoleLogWriter) Close() error {
	return nil
}

// This logger won't print data/time stamps
func createStdoutLogWriter() v2commlog.WriterCreator {
	return func() v2commlog.Writer {
		return &consoleLogWriter{
			logger: log.New(os.Stdout, "PROXY_GEN:", 0)}
	}
}
