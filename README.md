<p align="center">
<img style="align: center; padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="238px" height="238px" src="https://gitea.dancheg97.ru/repo-avatars/65-c2b763bfe9d206d7f362412b1e59e301" />
</p>

<h2 align="center">Go pacman</h2>

[![Generic badge](https://img.shields.io/badge/LICENSE-GPLv3-red.svg)](https://gitea.dancheg97.ru/dancheg97/go-pacman/src/branch/main/LICENSE)
[![Generic badge](https://img.shields.io/badge/GITEA-REPO-blue.svg)](https://gitea.dancheg97.ru/dancheg97/go-pacman)
[![Build Status](https://drone.dancheg97.ru/api/badges/dancheg97/go-pacman/status.svg)](https://drone.dancheg97.ru/dancheg97/go-pacman)

Dockerized pacman repository with gRPC API for packages.

Project aims to quickly set up and running personal pacman repostitory without pain and hussle.

Currently supports single gRPC method `Add()`, which will add packages from AUR repository to registry.

Environment variables:

- `REPO` - 📄 repository name on the web page
- `USER` - 😀 user name in system
- `GRPC_PORT` - 🌐 gRPC API port for repository packages
- `FILE_PORT` - 🌐 port for static file server to access packages
- `INIT_PKGS` - 📦 initial packages for download

You can run repository via `docker-compose`:

```yml
services:
  gitea.dancheg97.ru/dancheg97/go-pacman:latest:
    command: run
    environment:
      REPO: localhost
    volumes:
      - ./pacman:/var/cache/pacman/pkg
    ports:
      - 9080:9080
      - 8080:8080
```

Or, alternatively, if you have virtual machine without password for sudo user,
you can install package using go:

```sh
go install gitea.dancheg97.ru/dancheg97/go-pacman@latest
```

And run it as a CLI programm:

```sh
go-pacman -h
go-pacman run -h
```
