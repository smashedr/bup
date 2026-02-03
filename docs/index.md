---
icon: lucide/rocket
---

# :lucide-rocket: Get Started

[![Back UP](assets/images/logo.png){ align=right width=96 }](https://github.com/smashedr/bup?tab=readme-ov-file#readme)

[![GitHub Release Version](https://img.shields.io/github/v/release/smashedr/bup?logo=github)](https://github.com/smashedr/bup/releases)
[![GitHub Downloads](https://img.shields.io/github/downloads/smashedr/bup/total?logo=rolldown&logoColor=white)](https://github.com/smashedr/bup/releases/latest)
[![Image Size](https://badges.cssnr.com/ghcr/size/smashedr/bup)](https://github.com/smashedr/bup/pkgs/container/bup)
[![Go Version](https://img.shields.io/github/go-mod/go-version/smashedr/bup?logo=go&logoColor=white&label=go)](https://github.com/smashedr/bup/blob/master/go.mod)
[![GitHub Last Commit](https://img.shields.io/github/last-commit/smashedr/bup?logo=listenhub&label=updated)](https://github.com/smashedr/bup/pulse)
[![GitHub Repo Size](https://img.shields.io/github/repo-size/smashedr/bup?logo=buffer&label=repo%20size)](https://github.com/smashedr/bup?tab=readme-ov-file#readme)
[![GitHub Top Language](https://img.shields.io/github/languages/top/smashedr/bup?logo=devbox)](https://github.com/smashedr/bup?tab=readme-ov-file#readme)
[![GitHub Contributors](https://img.shields.io/github/contributors-anon/smashedr/bup?logo=southwestairlines)](https://github.com/smashedr/bup/graphs/contributors)
[![GitHub Issues](https://img.shields.io/github/issues/smashedr/bup?logo=codeforces&logoColor=white)](https://github.com/smashedr/bup/issues)
[![GitHub Discussions](https://img.shields.io/github/discussions/smashedr/bup?logo=theconversation&logoColor=white)](https://github.com/smashedr/bup/discussions)
[![GitHub Forks](https://img.shields.io/github/forks/smashedr/bup?style=flat&logo=forgejo&logoColor=white)](https://github.com/smashedr/bup/forks)
[![GitHub Repo Stars](https://img.shields.io/github/stars/smashedr/bup?style=flat&logo=gleam&logoColor=white)](https://github.com/smashedr/bup/stargazers)
[![GitHub Org Stars](https://img.shields.io/github/stars/cssnr?style=flat&logo=apachespark&logoColor=white&label=org%20stars)](https://cssnr.github.io/)
[![Discord](https://img.shields.io/discord/899171661457293343?logo=discord&logoColor=white&label=discord&color=7289da)](https://discord.gg/wXy6m2X8wY)
[![Ko-fi](https://img.shields.io/badge/Ko--fi-72a5f2?logo=kofi&label=support)](https://ko-fi.com/cssnr)

Back UP `bup` CLI written in Go.

Creates an archive of the `source` directory and puts it in the `destination` directory
in a sub-folder with the `name` of the `source` directory and a timestamped filename.

--8<-- "docs/snippets/install.md"

Remembers your `destination` directory and uses the current directory as `source` by default.

Supports directory excludes stored in the config file with the saved destination.

To get started see the [Quick Start](#quick-start) section or check out the [Features](#features).

If you run into any issues or have any questions, [support](support.md) is available.

[![VHS Tape](https://raw.githubusercontent.com/smashedr/repo-images/refs/heads/master/bup/demo.gif)](#quick-start)

## :lucide-sparkles: Features

- Backup a `source` directory to a `destination`
- Will save and reuse the `destination`
- Use the current directory for `source`
- Uses saved directory excludes
- Creates a timestamped archive
- Puts archives in named directory

## :lucide-plane-takeoff: Quick Start

Install.

--8<-- "docs/snippets/install.md"

Usage.

- Specify `source` and `destination`

```shell
bup backup [source] [destination]
```

- Use the `b` alias, `$pwd` source, and saved `destination`

```shell
bup b
```

- The `list` command list backups by name

```shell
bup l
```

- The `info` command prints the configuration

```shell
bup i
```

&nbsp;

!!! question

    If you need **help** getting started or run into any issues, [support](support.md) is available!
