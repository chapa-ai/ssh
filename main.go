package main

import (
	"fmt"
	lessWatcher "github.com/radovskyb/watcher"
	"ssh/agent"
	"ssh/logger"
	"ssh/watcher"
	"strings"
)

type Project struct {
	Watcher *watcher.EventFSLess
}

//type EventFSLess struct {
//	Watcher
//	Callback func(log logger.LoggerInterface, event *lessWatcher.Event)
//	Timeout  int32
//}
//
//type Watcher struct {
//	Singleton
//	Path         string
//	ExcludeMatch []string
//}

//type Singleton struct {
//	instance sync.Once
//	exitOnce sync.Once
//}

func Constructor(options agent.Options) *agent.Agent {
	return &agent.Agent{
		SSHOptions: options,
	}
}

func main() {
	p := Project{}

	a := Constructor(agent.Options{
		Ip:       "162.55.165.223",
		Password: "thil=ee8aL",
		Login:    "root",
	})

	p.Watcher = &watcher.EventFSLess{

		Watcher: watcher.Watcher{Path: "/Users/anton/Desktop/modem", ExcludeMatch: []string{".idea", ".git"}},
		Callback: func(log logger.LoggerInterface, event *lessWatcher.Event) {
			path := strings.Split(event.Path, "/Users/anton/Desktop/modem/")

			log.Info(fmt.Sprintf("Event: %v, %v", event.Path, event.FileInfo))
			log.Info(fmt.Sprintf("Path 1: %v", strings.Split(event.Path, "/Users/anton/Desktop/modem/")))

			log.Info(fmt.Sprintf("Path: %v, length: %v", path, len(path)))
			log.Info(fmt.Sprintf("P: %v", path[1]))

			destinationPath := "/srv/modem/" + path[1]

			log.Info(fmt.Sprintf("DestinationPath: %v", destinationPath))

			a.CopyFileFromHost(agent.Mirror{
				HostPath:        event.Path,
				DestinationPath: destinationPath,
			})

		},
		Timeout: 100,
	}

	p.Watcher.Watch()

}
