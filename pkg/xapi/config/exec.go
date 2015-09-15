package config

import (
	"sync"

	execdriver "github.com/docker/docker/pkg/xapi/types/exec"
)

type execConfig struct {
	sync.Mutex
	ID            string
	Running       bool
	ExitCode      int
	ProcessConfig *execdriver.ProcessConfig
	StreamConfig
	OpenStdin  bool
	OpenStderr bool
	OpenStdout bool
	Container  *Container
	canRemove  bool

	// waitStart will be closed immediately after the exec is really started.
	waitStart chan struct{}
}

type execStore struct {
	s map[string]*execConfig
	sync.RWMutex
}
