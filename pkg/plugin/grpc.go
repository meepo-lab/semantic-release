package plugin

import (
	"context"
	"errors"

	"github.com/hashicorp/go-plugin"
	"github.com/ted-vo/semantic-release/v3/pkg/analyzer"
	"github.com/ted-vo/semantic-release/v3/pkg/condition"
	"github.com/ted-vo/semantic-release/v3/pkg/generator"
	"github.com/ted-vo/semantic-release/v3/pkg/hooks"
	"github.com/ted-vo/semantic-release/v3/pkg/provider"
	"github.com/ted-vo/semantic-release/v3/pkg/updater"
	"google.golang.org/grpc"
)

type GRPCWrapper struct {
	Type string
	Impl interface{}
	plugin.NetRPCUnsupportedPlugin
}

func (p *GRPCWrapper) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	switch p.Type {
	case analyzer.CommitAnalyzerPluginName:
		analyzer.RegisterCommitAnalyzerPluginServer(s, &analyzer.CommitAnalyzerServer{
			Impl: p.Impl.(analyzer.CommitAnalyzer),
		})
	case condition.CIConditionPluginName:
		condition.RegisterCIConditionPluginServer(s, &condition.CIConditionServer{
			Impl: p.Impl.(condition.CICondition),
		})
	case generator.ChangelogGeneratorPluginName:
		generator.RegisterChangelogGeneratorPluginServer(s, &generator.ChangelogGeneratorServer{
			Impl: p.Impl.(generator.ChangelogGenerator),
		})
	case provider.PluginName:
		provider.RegisterProviderPluginServer(s, &provider.Server{
			Impl: p.Impl.(provider.Provider),
		})
	case updater.FilesUpdaterPluginName:
		updater.RegisterFilesUpdaterPluginServer(s, &updater.FilesUpdaterServer{
			Impl: p.Impl.(updater.FilesUpdater),
		})
	case hooks.PluginName:
		hooks.RegisterHooksPluginServer(s, &hooks.Server{
			Impl: p.Impl.(hooks.Hooks),
		})
	default:
		return errors.New("unknown plugin type")
	}
	return nil
}

func (p *GRPCWrapper) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	switch p.Type {
	case analyzer.CommitAnalyzerPluginName:
		return &analyzer.CommitAnalyzerClient{
			Impl: analyzer.NewCommitAnalyzerPluginClient(c),
		}, nil
	case condition.CIConditionPluginName:
		return &condition.CIConditionClient{
			Impl: condition.NewCIConditionPluginClient(c),
		}, nil
	case generator.ChangelogGeneratorPluginName:
		return &generator.ChangelogGeneratorClient{
			Impl: generator.NewChangelogGeneratorPluginClient(c),
		}, nil
	case provider.PluginName:
		return &provider.Client{
			Impl: provider.NewProviderPluginClient(c),
		}, nil
	case updater.FilesUpdaterPluginName:
		return &updater.FilesUpdaterClient{
			Impl: updater.NewFilesUpdaterPluginClient(c),
		}, nil
	case hooks.PluginName:
		return &hooks.Client{
			Impl: hooks.NewHooksPluginClient(c),
		}, nil
	}
	return nil, errors.New("unknown plugin type")
}
