package cmd

import (
	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
)

func SpecRegistryToPlugin(registry specs.Registry) managedplugin.Registry {
	switch registry {
	case specs.RegistryGithub:
		return managedplugin.RegistryGithub
	case specs.RegistryLocal:
		return managedplugin.RegistryLocal
	case specs.RegistryGrpc:
		return managedplugin.RegistryGrpc
	default:
		panic("unknown registry " + registry.String())
	}
}
