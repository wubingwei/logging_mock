package log2file

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	Frequency *logrus.Logger
	Forward   *logrus.Logger
)

const (
	frequencyPath = "./log/frequency.log"
	forwardPath   = "./log/forward.log"
)

func init() {
	Frequency, _ = newLog(frequencyPath, &EmptyFormatter{})
	Forward, _ = newLog(forwardPath, &EmptyFormatter{})
}

type EmptyFormatter struct{}

func (f *EmptyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	sb := strings.Builder{}
	sb.WriteString(entry.Message)
	sb.WriteString("\n")
	return []byte(sb.String()), nil
}

func newLog(fileName string, formatter logrus.Formatter) (*logrus.Logger, error) {
	absPath, _ := filepath.Abs(fileName)
	err := mkLogDir(absPath)
	if err != nil {
		return nil, errors.Wrapf(err, "mkdir from log path error: %s\n", err.Error())
	}
	log := logrus.New()
	writer, err := rotatelogs.New(
		absPath+".%Y-%m-%d-%H",
		rotatelogs.WithLinkName(absPath),            // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(2*time.Hour),          // 文件最大保存时间
		rotatelogs.WithRotationTime(60*time.Minute), // 日志切割时间间隔
	)
	if err != nil {
		return nil, errors.Wrap(err, "newLogrus "+absPath)
	}
	log.SetFormatter(formatter)
	log.SetOutput(writer)
	log.SetLevel(logrus.DebugLevel)
	return log, nil
}

func mkLogDir(logPath string) error {
	dir, _ := filepath.Split(logPath)
	if len(dir) > 0 {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
