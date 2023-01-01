package progresFile

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"ssh/logger"
	"ssh/singleton"
	"ssh/ssh/files/bar"
)

type MetaInfo struct {
	Path        string
	Name        string
	Label       string
	Required    bool
	Description string
}

type File1 struct {
	io.Reader
	MetaInfo

	Link *os.File
	Path string

	//ssh.Object
	singleton.Singleton
	LOG logger.LoggerInterface

	OnFinish func(f *os.File)
}

type File struct {
	File1

	bar   *bar.Bar
	total int64
}

func (f *File) finish() {
	f.Destruct(func() {
		f.bar.Finish()
		if f.OnFinish != nil {
			f.OnFinish(f.Link)
		}
	})
}

func (f *File) Read(p []byte) (int, error) {
	if f.Reader == nil && len(f.Path) > 0 {
		f.Init()
	}
	n, err := f.Reader.Read(p)

	if f.total == 0 {
		if f.LOG == nil {
			//f.LOG = log.New(log.D{"from": f.Link.Name()})
			logrus.Printf("from: %v", f.Link.Name())
		}
		f.LOG.Debug("start read")
		f.bar = &bar.Bar{
			//Description: fmt.Sprintf("%s ...", f.Description),
		}
		f.bar.NewOption(int64(n), f.GetFileSize())
	}

	if err == nil {
		f.total += int64(n)
		f.bar.IncCur(f.total)
	}

	switch {
	case f.total == f.GetFileSize():
		f.finish()
	}

	return n, err
}

func (f *File) Init() {
	if len(f.Path) > 0 {
		file, err := os.Open(f.Path)
		if err != nil {
			return
		}
		f.Link = file
		f.Reader = file
	}
}

func (f *File1) GetFileSize() int64 {
	stat, _ := f.Link.Stat()
	return stat.Size()
}
