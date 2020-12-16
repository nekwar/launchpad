package phase

import (
	"github.com/Mirantis/mcc/pkg/phase"
	"github.com/Mirantis/mcc/pkg/product/mke/api"
)

// CleanUp phase is used by reset for performing post-uninstall cleanups
type CleanUp struct {
	phase.BasicPhase
}

// Title for the phase
func (p *CleanUp) Title() string {
	return "Clean up"
}

// Run does all the prep work on the hosts in parallel
func (p *CleanUp) Run() error {
	err := phase.RunParallelOnHosts(p.Config.Spec.Hosts, p.Config, p.cleanupEnv)
	if err != nil {
		return err
	}

	return nil
}

func (p *CleanUp) cleanupEnv(h *api.Host, c *api.ClusterConfig) error {
	if len(h.Environment) > 0 {
		return h.Configurer.CleanupEnvironment(h.Environment)
	}
	return nil
}
