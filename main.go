package main

import (
	"fmt"
	lessWatcher "github.com/radovskyb/watcher"
	"os"
	"ssh/agent"
	"ssh/config"
	"ssh/logger"
	"ssh/watcher"
	"strings"
)

type Project struct {
	Watcher *watcher.EventFSLess
}

func Constructor(options agent.Options) *agent.Agent {
	return &agent.Agent{
		SSHOptions: options,
	}
}

func main() {
	p := Project{}

	err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	//a := Constructor(agent.Options{
	//	Ip:       "192.133.1.105",
	//	Password: "kLkeu9is9N",
	//	Login:    "root",
	//})

	a := Constructor(agent.Options{
		Ip:       os.Getenv("IpAddress"),
		Password: os.Getenv("Password"),
		Login:    os.Getenv("Login"),
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
