package writeDisk

import (
	"bitbucket.org/biglog/common"
	"bitbucket.org/biglog/config"
	"github.com/golang/glog"
	"os"
	"time"
)

func NewWriteDisck() *WriteDisck {
	t := &WriteDisck{
		"",
		config.GetFilePath(),
		config.GetMaxFileLen(),
		0,
		nil,
	}

	t.openNewFile()

	return t
}

type WriteDisck struct {
	fileName   string
	filePath   string
	maxFileLen int
	fileLen    int

	file *os.File
}

func (w *WriteDisck) WriteLog(log []byte) {
	if w.file == nil {
		w.openNewFile()
		if w.file == nil {
			return
		}
	}

	n, err := w.file.Write(log)

	if err != nil {
		glog.Fatalln("Write log failed!!!")
	}

	w.fileLen = w.fileLen + n

	//文件块写满后,创建新文件块
	if w.fileLen >= w.maxFileLen {
		w.file.Close()
		w.file = nil
		w.openNewFile()
	}
}

func (w *WriteDisck) openNewFile() error {
	if w.file != nil {
		return nil
	}

	uuid := common.NewUUID()
	ntime := time.Now()
	nowTime := ntime.Format("20060102150405")
	fileName := w.filePath + uuid + "_" + nowTime + ".log"

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0)

	if err == nil {
		w.file = file
		w.fileName = fileName
	} else {
		w.file = nil
		glog.Fatalln("Error crate log file by time:%s", nowTime)
	}

	return err
}
