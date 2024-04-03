package boot

import (
	"bytes"
	"goscaffold/model/config"
	"os"
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

func init() {
	cfg := config.GetConfig()
	initLog(cfg)

}

func initLog(cfg *config.Config) {
	// TODO: 需要增加日志文件输出
	// 设置日志
	logrus.SetOutput(os.Stdout)
	if cfg.Server.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			b := bytes.Buffer{}
			b.WriteString(filename)
			b.WriteString(":")
			b.WriteString(strconv.Itoa(f.Line))
			return "", b.String()
		},
	})
}
