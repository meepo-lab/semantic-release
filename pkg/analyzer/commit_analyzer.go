package analyzer

import (
	"github.com/ted-vo/semantic-release/v3/pkg/semrel"
)

type CommitAnalyzer interface {
	Init(map[string]string) error
	Name() string
	Version() string
	Analyze([]*semrel.RawCommit) []*semrel.Commit
}
