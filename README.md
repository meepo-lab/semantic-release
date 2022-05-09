# :package::rocket: semantic-release

## This is hard fork from [go-semantic-release](https://github.com/go-semantic-release/semantic-release) for personal customize update.

[![CI](https://github.com/ted-vo/semantic-release/workflows/CI/badge.svg?branch=master)](https://github.com/ted-vo/semantic-release/actions?query=workflow%3ACI+branch%3Amaster)
[![pipeline status](https://gitlab.com/ted-vo/semantic-release/badges/master/pipeline.svg)](https://gitlab.com/ted-vo/semantic-release/pipelines)
[![Go Report Card](https://goreportcard.com/badge/github.com/ted-vo/semantic-release)](https://goreportcard.com/report/github.com/ted-vo/semantic-release)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/ted-vo/semantic-release/v3)](https://pkg.go.dev/github.com/ted-vo/semantic-release/v3)

> fully automated package/module/image publishing

A more lightweight and standalone version of [semantic-release](https://github.com/semantic-release/semantic-release).

## How does it work?

Instead of writing [meaningless commit messages](http://whatthecommit.com/), we can take our time to think about the changes in the codebase and write them down. Following the [AngularJS Commit Message Conventions](https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y/edit) it is then possible to generate a helpful changelog and to derive the next semantic version number from them.

When `semantic-release` is setup it will do that after every successful continuous integration build of your default branch and publish the new version for you. This way no human is directly involved in the release process and your releases are guaranteed to be [unromantic and unsentimental](http://sentimentalversioning.org/).

_Source: [semantic-release/semantic-release#how-does-it-work](https://github.com/semantic-release/semantic-release#how-does-it-work)_

You can enforce semantic commit messages using [a git hook](https://github.com/hazcod/semantic-commit-hook).

## Installation

### Option 1: Use the ted-vo GitHub Action ([ted-vo/action](https://github.com/ted-vo/action))

### Option 2: Install `semantic-release` manually

```bash
curl -SL https://get-release.xyz/semantic-release/linux/amd64 -o ./semantic-release && chmod +x ./semantic-release
```

## Plugin System

Since v2, semantic-release is equipped with a plugin system. The plugins are standalone binaries that use [hashicorp/go-plugin](https://github.com/hashicorp/go-plugin) as a plugin library. `semantic-release` automatically downloads the necessary plugins if they don't exist locally. The plugins are stored in the `.semrel` directory of the current working directory in the following format: `.semrel/<os>_<arch>/<plugin name>/<version>/`. The ted-vo plugins API (`https://plugins.ted-vo.xyz/api/v1/`) is used to resolve plugins to the correct binary. The served content of the API can be found [here](https://github.com/ted-vo/ted-vo.github.io/tree/plugin-index), and a list of all existing plugins can be found [here](https://plugins.ted-vo.xyz/api/v1/plugins.json).

### Plugin Types

* Commit Analyzer ([Docs](https://pkg.go.dev/github.com/ted-vo/semantic-release/v3/pkg/analyzer?tab=doc#CommitAnalyzer), [Example](https://github.com/ted-vo/commit-analyzer-cz))
* CI Condition ([Docs](https://pkg.go.dev/github.com/ted-vo/semantic-release/v3/pkg/condition?tab=doc#CICondition), [Example](https://github.com/ted-vo/condition-github))
* Changelog Generator ([Docs](https://pkg.go.dev/github.com/ted-vo/semantic-release/v3/pkg/generator?tab=doc#ChangelogGenerator), [Example](https://github.com/ted-vo/changelog-generator-default))
* Provider ([Docs](https://pkg.go.dev/github.com/ted-vo/semantic-release/v3/pkg/provider?tab=doc#Provider), [Example](https://github.com/ted-vo/provider-github))
* Files Updater ([Docs](https://pkg.go.dev/github.com/ted-vo/semantic-release/v3/pkg/updater?tab=doc#FilesUpdater), [Example](https://github.com/ted-vo/files-updater-npm))
* Hooks ([Docs](https://pkg.go.dev/github.com/ted-vo/semantic-release/v3/pkg/hooks?tab=doc#Hooks), [Example](https://github.com/ted-vo/hooks-goreleaser))

### Configuration

Plugins can be configured using CLI flags or the `.semrelrc` config file. By using a `@` sign after the plugin name, the required version of the plugin can be specified. Otherwise, any locally installed version will be used. If the plugin does not exist locally, the latest version will be downloaded. This is an example of the `.semrelrc` config file:

```json
{
  "plugins": {
    "commit-analyzer": {
      "name": "default@^1.0.0"
    },
    "ci-condition": {
      "name": "default"
    },
    "changelog-generator": {
      "name": "default",
      "options": {
        "emojis": "true"
      }
    },
    "provider": {
      "name": "gitlab",
      "options": {
        "gitlab_projectid": "123456"
      }
    },
    "files-updater": {
      "names": ["npm"]
    }
  }
}
```

## Example GitHub Actions

## Example GitLab CI Config

### GitLab token

It is necessary to create a new Gitlab personal access token with the `api` scope [here](https://gitlab.com/profile/personal_access_tokens).
Ensure the CI variable is protected and masked as the `GITLAB_TOKEN` has a lot of rights. There is an open issue for project specific [tokens](https://gitlab.com/gitlab-org/gitlab/issues/756)
You can set the GitLab token via the `GITLAB_TOKEN` environment variable or the `-token` flag.

.gitlab-ci.yml

```yml
 stages:
  # other stages
  - release

release:
  image: registry.gitlab.com/ted-vo/semantic-release:latest # Replace this with the current release
  stage: release
  # Remove this if you want a release created for each push to master
  when: manual
  only:
    - master
  script:
    - release
```

## Beta release support

Beta release support empowers you to release beta, rc, etc. versions with `semantic-release` (e.g. v2.0.0-beta.1). To enable this feature you need to create a new branch (e.g. beta/v2) and check in a `.semrelrc` file with the following content:

``` json
{
  "maintainedVersion": "2-beta"
}
```

If you commit to this branch a new incremental pre-release is created everytime you push. (2.0.0-beta.1, 2.0.0-beta.2, ...)

## Licence

The [MIT License (MIT)](http://opensource.org/licenses/MIT)

Copyright © 2020 [Christoph Witzko](https://twitter.com/christophwitzko)
