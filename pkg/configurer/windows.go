package configurer

import (
	"fmt"
	"strings"

	"github.com/Mirantis/mcc/pkg/config"
)

// WindowsConfigurer is a generic windows host configurer
type WindowsConfigurer struct {
	Host *config.Host
}

// InstallEngine install Docker EE engine on Windows
func (c *WindowsConfigurer) InstallEngine(engineConfig *config.EngineConfig) error {
	if c.Host.Metadata.EngineVersion == engineConfig.Version {
		return nil
	}
	scriptURL := fmt.Sprintf("%sinstall.ps1", config.EngineInstallURL)
	dlCommand := fmt.Sprintf("$ProgressPreference = 'SilentlyContinue'; [Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12; Invoke-WebRequest %s -UseBasicParsing -OutFile install.ps1", scriptURL)
	err := c.Host.ExecCmd("powershell", dlCommand, false)
	if err != nil {
		return err
	}
	installCommand := fmt.Sprintf("set DOWNLOAD_URL=%s && set DOCKER_VERSION=%s && set CHANNEL=%s && powershell -ExecutionPolicy Bypass -File install.ps1 -Verbose", engineConfig.RepoURL, engineConfig.Version, engineConfig.Channel)
	err = c.Host.Exec(installCommand)
	if err != nil {
		return err
	}

	return nil
}

// RestartEngine restarts Docker EE engine
func (c *WindowsConfigurer) RestartEngine() error {
	// TODO: handle restart
	return nil
}

// ResolveHostname resolves hostname
func (c *WindowsConfigurer) ResolveHostname() string {
	output, err := c.Host.ExecWithOutput("powershell $env:COMPUTERNAME")
	if err != nil {
		return "localhost"
	}
	return strings.TrimSpace(output)
}

// ResolveInternalIP resolves internal ip from private interface
func (c *WindowsConfigurer) ResolveInternalIP() string {
	output, err := c.Host.ExecWithOutput(fmt.Sprintf(`powershell -Command "(Get-NetIPAddress -AddressFamily IPv4 -InterfaceAlias \"%s\").IPAddress"`, c.Host.PrivateInterface))
	if err != nil {
		return c.Host.Address
	}
	return strings.TrimSpace(output)
}

// IsContainerized checks if host is actually a container
func (c *WindowsConfigurer) IsContainerized() bool {
	return false
}

// DockerCommandf accepts a printf-like template string and arguments
// and builds a command string for running the docker cli on the host
func (c *WindowsConfigurer) DockerCommandf(template string, args ...interface{}) string {
	return fmt.Sprintf("docker.exe %s", fmt.Sprintf(template, args...))
}
