[![GitHub Release Version](https://img.shields.io/github/v/release/smashedr/bup?logo=github)](https://github.com/smashedr/bup/releases)
[![GitHub Downloads](https://img.shields.io/github/downloads/smashedr/bup/total?logo=rolldown&logoColor=white)](https://github.com/smashedr/bup/releases/latest)
[![Go Version](https://img.shields.io/github/go-mod/go-version/smashedr/bup?logo=go&logoColor=white&label=go)](https://github.com/smashedr/bup/blob/master/go.mod)
[![Deployment Docs](https://img.shields.io/github/deployments/smashedr/bup/docs?logo=materialformkdocs&logoColor=white&label=docs)](https://github.com/smashedr/bup/deployments/docs)
[![Deployment Preview](https://img.shields.io/github/deployments/smashedr/bup/preview?logo=materialformkdocs&logoColor=white&label=preview)](https://github.com/smashedr/bup/deployments/preview)
[![Workflow Release](https://img.shields.io/github/actions/workflow/status/smashedr/bup/release.yaml?logo=testcafe&logoColor=white&label=release)](https://github.com/smashedr/bup/actions/workflows/release.yaml)
[![Workflow Lint](https://img.shields.io/github/actions/workflow/status/smashedr/bup/lint.yaml?logo=testcafe&logoColor=white&label=lint)](https://github.com/smashedr/bup/actions/workflows/lint.yaml)
[![GitHub Last Commit](https://img.shields.io/github/last-commit/smashedr/bup?logo=listenhub&label=updated)](https://github.com/smashedr/bup/pulse)
[![GitHub Repo Size](https://img.shields.io/github/repo-size/smashedr/bup?logo=buffer&label=repo%20size)](https://github.com/smashedr/bup?tab=readme-ov-file#readme)
[![GitHub Top Language](https://img.shields.io/github/languages/top/smashedr/bup?logo=devbox)](https://github.com/smashedr/bup?tab=readme-ov-file#readme)
[![GitHub Contributors](https://img.shields.io/github/contributors-anon/smashedr/bup?logo=southwestairlines)](https://github.com/smashedr/bup/graphs/contributors)
[![GitHub Issues](https://img.shields.io/github/issues/smashedr/bup?logo=codeforces&logoColor=white)](https://github.com/smashedr/bup/issues)
[![GitHub Discussions](https://img.shields.io/github/discussions/smashedr/bup?logo=theconversation)](https://github.com/smashedr/bup/discussions)
[![GitHub Forks](https://img.shields.io/github/forks/smashedr/bup?style=flat&logo=forgejo&logoColor=white)](https://github.com/smashedr/bup/forks)
[![GitHub Repo Stars](https://img.shields.io/github/stars/smashedr/bup?style=flat&logo=gleam&logoColor=white)](https://github.com/smashedr/bup/stargazers)
[![GitHub Org Stars](https://img.shields.io/github/stars/cssnr?style=flat&logo=apachespark&logoColor=white&label=org%20stars)](https://cssnr.github.io/)
[![Discord](https://img.shields.io/discord/899171661457293343?logo=discord&logoColor=white&label=discord&color=7289da)](https://discord.gg/wXy6m2X8wY)
[![Ko-fi](https://img.shields.io/badge/Ko--fi-72a5f2?logo=kofi&label=support)](https://ko-fi.com/cssnr)

# Back UP

[![Homebrew](https://img.shields.io/badge/homebrew-gray?style=flat-square&logo=homebrew)](#homebrew)
[![Bash](https://img.shields.io/badge/bash-gray?style=flat-square&logo=stackedit&logoColor=white)](#bash)
[![Powershell](https://img.shields.io/badge/powershell-gray?style=flat-square&logo=cashapp&logoColor=lightblue)](#powershell)
[![Go](https://img.shields.io/badge/source-gray?style=flat-square&logo=go)](#source)
[![Docker Installer](https://img.shields.io/badge/docker_installer-gray?style=flat-square&logo=docker)](#docker)
[![Windows Installer](https://img.shields.io/badge/windows_installer-gray?style=flat-square&logo=data:image/svg%2bxml;base64,PHN2ZyB3aWR0aD0iMjQ5MCIgaGVpZ2h0PSIyNTAwIiB2aWV3Qm94PSIwIDAgMjU2IDI1NyIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiBwcmVzZXJ2ZUFzcGVjdFJhdGlvPSJ4TWlkWU1pZCI+PHBhdGggZD0iTTAgMzYuMzU3TDEwNC42MiAyMi4xMWwuMDQ1IDEwMC45MTQtMTA0LjU3LjU5NUwwIDM2LjM1OHptMTA0LjU3IDk4LjI5M2wuMDggMTAxLjAwMkwuMDgxIDIyMS4yNzVsLS4wMDYtODcuMzAyIDEwNC40OTQuNjc3em0xMi42ODItMTE0LjQwNUwyNTUuOTY4IDB2MTIxLjc0bC0xMzguNzE2IDEuMVYyMC4yNDZ6TTI1NiAxMzUuNmwtLjAzMyAxMjEuMTkxLTEzOC43MTYtMTkuNTc4LS4xOTQtMTAxLjg0TDI1NiAxMzUuNnoiIGZpbGw9IiMwMEFERUYiLz48L3N2Zz4=)](https://github.com/smashedr/bup/releases/latest/download/bup_Windows_Installer.exe)

<a title="BackUP" href="https://smashedr.github.io/bup/" target="_blank">
<img alt="BackUP" align="right" width="128" height="auto" src="https://raw.githubusercontent.com/smashedr/bup/refs/heads/master/docs/assets/images/logo.png"></a>

- [Install](#install)
- [Usage](#usage)
- [Development](#development)
- [Support](#Support)
- [Contributing](#contributing)

Back UP `bup` CLI written in Go.

Creates an archive of the `source` directory and puts it in the `destination` directory
in a sub-folder with the `name` of the `source` directory and a timestamped filename.

Remembers your `destination` directory and uses the current directory as `source` by default.

Supports directory excludes stored in the config file with the saved destination.

[![VHS Tape](https://cssnr.s3.amazonaws.com/bup/demo.gif)](https://smashedr.github.io/bup/)

## Install

[![Latest Release](https://img.shields.io/github/v/release/smashedr/bup?logo=github&label=latest%20release)](https://github.com/smashedr/bup/releases)
[![Windows Installer](https://img.shields.io/badge/download-windows--installer.exe-blue?logo=data:image/svg%2bxml;base64,PHN2ZyB3aWR0aD0iMjQ5MCIgaGVpZ2h0PSIyNTAwIiB2aWV3Qm94PSIwIDAgMjU2IDI1NyIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiBwcmVzZXJ2ZUFzcGVjdFJhdGlvPSJ4TWlkWU1pZCI+PHBhdGggZD0iTTAgMzYuMzU3TDEwNC42MiAyMi4xMWwuMDQ1IDEwMC45MTQtMTA0LjU3LjU5NUwwIDM2LjM1OHptMTA0LjU3IDk4LjI5M2wuMDggMTAxLjAwMkwuMDgxIDIyMS4yNzVsLS4wMDYtODcuMzAyIDEwNC40OTQuNjc3em0xMi42ODItMTE0LjQwNUwyNTUuOTY4IDB2MTIxLjc0bC0xMzguNzE2IDEuMVYyMC4yNDZ6TTI1NiAxMzUuNmwtLjAzMyAxMjEuMTkxLTEzOC43MTYtMTkuNTc4LS4xOTQtMTAxLjg0TDI1NiAxMzUuNnoiIGZpbGw9IiMwMEFERUYiLz48L3N2Zz4=)](https://github.com/smashedr/bup/releases/latest/download/bup_Windows_Installer.exe)

#### Homebrew

```shell
brew install cssnr/tap/bup
```

#### Bash

```shell
curl https://i.jpillora.com/smashedr/bup! | bash
```

💾 Alternatively, you can manually [download a release](https://github.com/smashedr/bup/releases).

#### PowerShell

```powershell
iex (iwr -useb 'https://raw.githubusercontent.com/smashedr/bup/refs/heads/master/scripts/install.ps1').Content
```

🪟 Windows users can download the [Windows Installer.exe](https://github.com/smashedr/bup/releases/latest/download/bup_Windows_Installer.exe).

#### Source

```shell
go install github.com/smashedr/bup@latest
```

#### Docker

```shell
docker run --rm -itv ~/bin:/out ghcr.io/smashedr/ir:latest -b /out smashedr/bup
```

_Note: Docker requires you to mount the target bin directory._

[![View Documentation](https://img.shields.io/badge/view_documentation-blue?style=for-the-badge&logo=googledocs&logoColor=white)](https://smashedr.github.io/bup/)

## Usage

Specify `source` and `destination`.

```shell
bup backup [source] [destination]
```

Use the `b` alias, current directory, and saved `destination`.

```shell
bup b
```

The `list` command list backups by name.

```shell
bup l
```

The `info` command prints the configuration.

```shell
bup i
```

[![View Documentation](https://img.shields.io/badge/view_documentation-blue?style=for-the-badge&logo=googledocs&logoColor=white)](https://smashedr.github.io/bup/)

# Development

Go: <https://go.dev/doc/install>

```shell
go run main.go
```

Task: <https://taskfile.dev/docs/installation>

```shell
task build
task lint
```

Docs: <https://zensical.org/docs/get-started>

```shell
task docs
```

Inno Setup: <https://jrsoftware.org/isdl.php>

```shell
task pathmgr
task inno
```

# Support

If you run into any issues or need help getting started, please do one of the following:

- Report an Issue: <https://github.com/smashedr/bup/issues>
- Q&A Discussion: <https://github.com/smashedr/bup/discussions/categories/q-a>
- Request a Feature: <https://github.com/smashedr/bup/issues/new?template=1-feature.yaml>
- Chat with us on Discord: <https://discord.gg/wXy6m2X8wY>

[![Features](https://img.shields.io/badge/features-brightgreen?style=for-the-badge&logo=rocket&logoColor=white)](https://github.com/smashedr/bup/issues/new?template=1-feature.yaml)
[![Issues](https://img.shields.io/badge/issues-red?style=for-the-badge&logo=southwestairlines&logoColor=white)](https://github.com/smashedr/bup/issues)
[![Discussions](https://img.shields.io/badge/discussions-blue?style=for-the-badge&logo=theconversation&logoColor=white)](https://github.com/smashedr/bup/discussions)
[![Discord](https://img.shields.io/badge/discord-5865F2?style=for-the-badge&logo=discord&logoColor=white)](https://discord.gg/wXy6m2X8wY)

# Contributing

If you would like to submit a PR, please review the [CONTRIBUTING.md](#contributing-ov-file).

Please consider making a donation to support the development of this project
and [additional](https://cssnr.com/) open source projects.

[![Ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/cssnr)

For a full list of current projects visit: [https://cssnr.github.io/](https://cssnr.github.io/)
