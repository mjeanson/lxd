package network

import (
	"github.com/lxc/lxd/lxd/cluster"
	"github.com/lxc/lxd/lxd/state"
	"github.com/lxc/lxd/shared/api"
)

// Network represents a LXD network.
type Network interface {
	// Load.
	init(state *state.State, id int64, name string, netType string, description string, config map[string]string, status string)
	fillConfig(config map[string]string) error

	// Config.
	Validate(config map[string]string) error
	Name() string
	Type() string
	Status() string
	Config() map[string]string
	IsUsed() (bool, error)
	HasDHCPv4() bool
	HasDHCPv6() bool
	DHCPv4Ranges() []DHCPRange
	DHCPv6Ranges() []DHCPRange

	// Actions.
	Create(clusterNotification bool) error
	Start() error
	Stop() error
	Rename(name string) error
	Update(newNetwork api.NetworkPut, targetNode string, clusterNotification bool) error
	HandleHeartbeat(heartbeatData *cluster.APIHeartbeat) error
	Delete(clusterNotification bool) error
}
