package agent

import (
	"github.com/sirupsen/logrus"
	Errors "ssh/errors"
	"ssh/logger"
	"ssh/singleton"
	s "ssh/ssh"
)

type Mirror struct {
	HostPath        string
	DestinationPath string
	ExcludeMatch    []string
}

type Agent struct {
	s.Object
	*s.MetaInfo
	singleton.Singleton

	ssh        *s.SSH
	SSHOptions Options
	LOG        logger.LoggerInterface
}

type Options struct {
	Ip       string
	Login    string
	Password string
}

func (a *Agent) init() {
	a.ssh.Construct(func() {
		a.ssh = &s.SSH{
			Options: s.Options{
				Ip:       a.SSHOptions.Ip,
				Password: a.SSHOptions.Password,
				Login:    a.SSHOptions.Login,
			},
		}
		//a.LOG = log.New(log.D{"agent": "SFTP-AGENT"})
		logrus.Printf("agent: SFTP-AGENT")

		a.LOG.Debug("init agent")
	})

}

func (a *Agent) PingSSH() bool {
	a.init()
	return a.ssh.TestConnection()
}

func (a *Agent) Watch(options Mirror) {
	a.init()
	if !a.PingSSH() {
		a.LOG.Error(Errors.ErrorSshConnection.Error())
		return
	}
}

func (a *Agent) CopyFileFromHost(options Mirror) error {
	a.init()
	return a.ssh.CopyFileFromHost(options.HostPath, options.DestinationPath)
}
