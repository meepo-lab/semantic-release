package provider

import (
	"github.com/ted-vo/semantic-release/v3/pkg/semrel"
)

type Provider interface {
	Init(map[string]string) error
	Name() string
	Version() string
	GetInfo() (*RepositoryInfo, error)
	GetCommits(fromSha, toSha string) ([]*semrel.RawCommit, error)
	GetReleases(re string) ([]*semrel.Release, error)
	CreateRelease(*CreateReleaseConfig) error
	CommitFilesChanged(filePaths []string, message string) (string, error)
}
