package plugin

import (
	"github.com/hashicorp/go-plugin"
	"github.com/ted-vo/semantic-release/v3/pkg/analyzer"
	"github.com/ted-vo/semantic-release/v3/pkg/condition"
	"github.com/ted-vo/semantic-release/v3/pkg/generator"
	"github.com/ted-vo/semantic-release/v3/pkg/hooks"
	"github.com/ted-vo/semantic-release/v3/pkg/provider"
	"github.com/ted-vo/semantic-release/v3/pkg/updater"
)

var Handshake = plugin.HandshakeConfig{
	MagicCookieKey:   "GO_SEMANTIC_RELEASE_MAGIC_COOKIE",
	MagicCookieValue: "beepboop",
}

type CommitAnalyzerFunc func() analyzer.CommitAnalyzer
type CIConditionFunc func() condition.CICondition
type ChangelogGeneratorFunc func() generator.ChangelogGenerator
type ProviderFunc func() provider.Provider
type FilesUpdaterFunc func() updater.FilesUpdater
type HooksFunc func() hooks.Hooks

type ServeOpts struct {
	CommitAnalyzer     CommitAnalyzerFunc
	CICondition        CIConditionFunc
	ChangelogGenerator ChangelogGeneratorFunc
	Provider           ProviderFunc
	FilesUpdater       FilesUpdaterFunc
	Hooks              HooksFunc
}

func Serve(opts *ServeOpts) {
	pluginSet := make(plugin.PluginSet)

	if opts.CommitAnalyzer != nil {
		pluginSet[analyzer.CommitAnalyzerPluginName] = &GRPCWrapper{
			Type: analyzer.CommitAnalyzerPluginName,
			Impl: opts.CommitAnalyzer(),
		}
	}

	if opts.CICondition != nil {
		pluginSet[condition.CIConditionPluginName] = &GRPCWrapper{
			Type: condition.CIConditionPluginName,
			Impl: opts.CICondition(),
		}
	}

	if opts.ChangelogGenerator != nil {
		pluginSet[generator.ChangelogGeneratorPluginName] = &GRPCWrapper{
			Type: generator.ChangelogGeneratorPluginName,
			Impl: opts.ChangelogGenerator(),
		}
	}

	if opts.Provider != nil {
		pluginSet[provider.PluginName] = &GRPCWrapper{
			Type: provider.PluginName,
			Impl: opts.Provider(),
		}
	}

	if opts.FilesUpdater != nil {
		pluginSet[updater.FilesUpdaterPluginName] = &GRPCWrapper{
			Type: updater.FilesUpdaterPluginName,
			Impl: opts.FilesUpdater(),
		}
	}

	if opts.Hooks != nil {
		pluginSet[hooks.PluginName] = &GRPCWrapper{
			Type: hooks.PluginName,
			Impl: opts.Hooks(),
		}
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: Handshake,
		VersionedPlugins: map[int]plugin.PluginSet{
			1: pluginSet,
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
