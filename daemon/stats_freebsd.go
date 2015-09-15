package daemon

import (
	"github.com/docker/docker/pkg/xapi/types"
	"github.com/opencontainers/runc/libcontainer"
)

// convertStatsToAPITypes converts the libcontainer.Stats to the api specific
// structs.  This is done to preserve API compatibility and versioning.
func convertStatsToAPITypes(ls *libcontainer.Stats) *types.Stats {
	// TODO FreeBSD. Refactor accordingly to fill in stats.
	s := &types.Stats{}
	return s
}
