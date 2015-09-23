package xapi

import (
	"io"
	"time"

	// everything from here needs to move to types
	// consists mainly of XYZConfig structs to pass information
	"github.com/docker/docker/daemon"
	"github.com/docker/docker/runconfig"

	"github.com/docker/docker/pkg/xapi/types"
)

// Backend is all the methods that need to be implemented to provide
// all the stuff a server wants
type Backend interface {
	SystemInfo() (*types.Info, error)
	//ContainerStart(name string, hostConfig *runconfig.HostConfig) error
	// NetworkApiRouter()

	// Get(string) // get container NO
	Exists(id string) bool
	// ContainerCopy(string, string)
	// ContainerStatPath(string, string)
	// ContainerArchivePath(string, string)
	// ContainerExtractToDir(string, string, bool, io.Reader)
	// ContainerInspect(string)
	ContainerInspect(name string) (*types.ContainerJSON, error)
	ContainerInspect120(name string) (*types.ContainerJSON120, error)
	ContainerInspectPre120(name string) (*types.ContainerJSONPre120, error)

	// Containers(config)
	ContainerStats(prefixOrName string, config *daemon.ContainerStatsConfig) error
	// ContainerLogs(c, logsConfig)
	ContainerExport(name string, out io.Writer) error
	// ContainerStart(string, hostConfig)
	// ContainerStop(string, seconds)
	// ContainerKill(name, sig)
	// ContainerRestart(string,
	ContainerPause(name string) error
	ContainerUnpause(name string) error
	ContainerWait(name string, timeout time.Duration) (int, error)
	// ContainerChanges(string)
	ContainerTop(name string, psArgs string) (*types.ContainerProcessList, error)
	// ContainerRename(name, newName)
	// ContainerCreate(name, config, hostConfig)
	// ContainerRm(name, config)
	ContainerResize(name string, height, width int) error
	ContainerExecResize(name string, height, width int) error

	// ContainerAttachWithLogs(cont, attachWithLogsConfig)
	// ContainerWsAttachWithLogs(cont, wsAttachWithLogsConfig)

	ContainerExecStart(execName string, stdin io.ReadCloser, stdout io.Writer, stderr io.Writer) error
	// two different versions of ExecConfig, oi vey!
	ContainerExecCreate(config *runconfig.ExecConfig) (string, error)
	ContainerExecInspect(id string) (*daemon.ExecConfig, error)
	// Repositories()

	// ContainerAttachWithLogs(cont, attachWithLogsConfig)
	// ContainerWsAttachWithLogs(cont, wsAttachWithLogsConfig)
	// ContainerExecInspect(string)
	// ContainerExecCreate(execConfig)
	// ContainerExecStart(execName, stdin, stdout, stderr)
	// ContainerExecResize(string, height, width)

	// NetworkApiRouter()
	// ImageDelete(name, force, noprune)
	// EventsService
	// RegistryService
	// ContainerInspectPre120(namevar)
	VolumeInspect(name string) (*types.Volume, error)
}
