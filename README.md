[![GitHub Release Version](https://img.shields.io/github/v/release/smashedr/bup?logo=github)](https://github.com/smashedr/bup/releases)
[![GitHub Downloads](https://img.shields.io/github/downloads/smashedr/bup/total?logo=rolldown&logoColor=white)](https://github.com/smashedr/bup/releases/latest)
[![Image Size](https://badges.cssnr.com/ghcr/size/smashedr/bup)](https://github.com/smashedr/bup/pkgs/container/bup)
[![Go Version](https://img.shields.io/github/go-mod/go-version/smashedr/bup?logo=go&logoColor=white&label=go)](https://github.com/smashedr/bup/blob/master/go.mod)
[![Deployment Docs](https://img.shields.io/github/deployments/smashedr/bup/docs?logo=materialformkdocs&logoColor=white&label=docs)](https://github.com/smashedr/bup/deployments/docs)
[![Deployment Preview](https://img.shields.io/github/deployments/smashedr/bup/preview?logo=materialformkdocs&logoColor=white&label=preview)](https://github.com/smashedr/bup/deployments/preview)
[![Workflow Release](https://img.shields.io/github/actions/workflow/status/smashedr/bup/release.yaml?logo=testcafe&logoColor=white&label=release)](https://github.com/smashedr/bup/actions/workflows/release.yaml)
[![Workflow Lint](https://img.shields.io/github/actions/workflow/status/smashedr/bup/lint.yaml?logo=testcafe&logoColor=white&label=lint)](https://github.com/smashedr/bup/actions/workflows/lint.yaml)
[![GitHub Last Commit](https://img.shields.io/github/last-commit/smashedr/bup?logo=speedtest&label=updated)](https://github.com/smashedr/bup?tab=readme-ov-file#readme)
[![GitHub Repo Size](https://img.shields.io/github/repo-size/smashedr/bup?logo=buffer&label=repo%20size)](https://github.com/smashedr/bup?tab=readme-ov-file#readme)
[![GitHub Top Language](https://img.shields.io/github/languages/top/smashedr/bup?logo=devbox)](https://github.com/smashedr/bup?tab=readme-ov-file#readme)
[![GitHub Contributors](https://img.shields.io/github/contributors-anon/smashedr/bup?logo=southwestairlines)](https://github.com/smashedr/bup/graphs/contributors)
[![GitHub Issues](https://img.shields.io/github/issues/smashedr/bup?logo=codeforces&logoColor=white)](https://github.com/smashedr/bup/issues)
[![GitHub Discussions](https://img.shields.io/github/discussions/smashedr/bup?logo=rocketdotchat&logoColor=white)](https://github.com/smashedr/bup/discussions)
[![GitHub Forks](https://img.shields.io/github/forks/smashedr/bup?style=flat&logo=forgejo&logoColor=white)](https://github.com/smashedr/bup/forks)
[![GitHub Repo Stars](https://img.shields.io/github/stars/smashedr/bup?style=flat&logo=gleam&logoColor=white)](https://github.com/smashedr/bup/stargazers)
[![GitHub Org Stars](https://img.shields.io/github/stars/cssnr?style=flat&logo=apachespark&logoColor=white&label=org%20stars)](https://cssnr.github.io/)
[![Discord](https://img.shields.io/discord/899171661457293343?logo=discord&logoColor=white&label=discord&color=7289da)](https://discord.gg/wXy6m2X8wY)
[![Ko-fi](https://img.shields.io/badge/Ko--fi-72a5f2?logo=kofi&label=support)](https://ko-fi.com/cssnr)

# Back UP

[![Homebrew](https://img.shields.io/badge/brew_install-smashedr%2Ftest%2Fbup-blue?style=flat-square&logo=homebrew)](#homebrew)
[![Docker](https://img.shields.io/badge/docker_run-ghcr.io%2Fsmashedr%2Fbup-blue?style=flat-square&logo=docker)](#docker)
[![GitHub](https://img.shields.io/badge/curl_--L-i.jpillora.com%2Fsmashedr%2Fbup!|bash-blue?style=flat-square&logo=github)](#github)

- [Install](#install)
- [Usage](#usage)
- [Development](#development)
- [Contributing](#contributing)

Back UP `bup` CLI written in Go.

Creates an archive of the `source` directory and puts it in the `destination` directory
in a sub-folder with the `name` of the `source` directory and a timestamped filename.

Remembers your `destination` directory and uses the current directory as `source` by default.

Supports directory excludes stored in the config file with the saved destination.

[![View Documentation](https://img.shields.io/badge/view_documentation-blue?style=for-the-badge&logo=googledocs&logoColor=white)](https://smashedr.github.io/bup/)

## Install

#### GitHub

```shell
curl https://i.jpillora.com/smashedr/bup! | bash
```

Alternatively, you can manually [download a release](https://github.com/smashedr/bup/releases).

#### Homebrew

```shell
brew install smashedr/test/bup
```

#### Docker

```shell
docker run --rm ghcr.io/smashedr/bup:latest
```

#### Source

```shell
go install github.com/smashedr/bup@latest
```

## Usage

- Specify `source` and `destination`

```shell
bup backup [source] [destination]
```

- Use the `b` alias, `$(cwd)` source, and saved `destination`

```shell
bup b
```

<details><summary>📟 View Command Output</summary>

```text
PS C:\Users\Shane\IdeaProjects\django-vue> bup b
Config File: C:\Users\Shane\.config\bup.yaml
Enter Destination Path: C:\Users\Shane\Backup
Saved Default Destination: C:\Users\Shane\Backup
Source: C:\Users\Shane\IdeaProjects\django-vue
Destination: C:\Users\Shane\Backup
Name: django-vue
Proceed? (y/N): y
Directory: C:\Users\Shane\Backup\django-vue
Excludes: [.cache .venv build dist node_modules venv]
Archive: C:\Users\Shane\Backup\django-vue\26-01-19-20-48-20.zip
Success!
```

</details>

- The `list` command is INOP but prints the configuration

```shell
bup l
```

<details><summary>📟 View Command Output</summary>

```text
PS C:\Users\Shane\IdeaProjects\django-vue> bup l
Config File: C:\Users\Shane\.config\bup.yaml
args: []
cfgFile:
destination: C:\Users\Shane\Backup
excludes: [.cache .venv build dist node_modules venv]
 - List Command INOP
```

</details>

[![View Documentation](https://img.shields.io/badge/view_documentation-blue?style=for-the-badge&logo=googledocs&logoColor=white)](https://smashedr.github.io/bup/)

# Development

Go: <https://go.dev/doc/install>

```shell
go run main.go
```

Task: <https://taskfile.dev/docs/installation>

```shell
task build
```

Docs: <https://zensical.org/docs/get-started>

```shell
zensical serve --open
```

# Contributing

Please consider making a donation to support the development of this project
and [additional](https://cssnr.com/) open source projects.

[![Ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/cssnr)

For a full list of current projects visit: [https://cssnr.github.io/](https://cssnr.github.io/)
