<p align="center">
<img style="align: center; padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="238px" height="238px" src="https://gitea.dancheg97.ru/repo-avatars/65-c2b763bfe9d206d7f362412b1e59e301" />
</p>

<h2 align="center">Go pacman</h2>

[![Generic badge](https://img.shields.io/badge/LICENSE-GPLv3-red.svg)](https://gitea.dancheg97.ru/dancheg97/go-pacman/src/branch/main/LICENSE)
[![Generic badge](https://img.shields.io/badge/GITEA-REPO-orgnge.svg)](https://gitea.dancheg97.ru/dancheg97/go-pacman)
[![Generic badge](https://img.shields.io/badge/DOCKER-CONTAINER-blue.svg)](https://gitea.dancheg97.ru/dancheg97/-/packages/container/go-pacman/latest)
[![Build Status](https://drone.dancheg97.ru/api/badges/dancheg97/go-pacman/status.svg)](https://drone.dancheg97.ru/dancheg97/go-pacman)

Dockerized pacman repository with gRPC API for package management. Project goal is to quickly set up personal pacman repostitory without pain and hustle.

gRPC API is available at port `9080` by default, and should not be exposed publically, because it executes `yay` shell command under the hood. Also, this API supports reflection, so you can explore it using `grpcurl`, `grpclicker-vscode` or `insomnia`.

---

### Configurations

Environment variables/flags:

- 📄 - `REPO` - `repo` - repository name on the web page
- 😀 - `USER` - `user` - user name in system, will be used to eject `yay` packages
- 🌐 - `GRPC_PORT` - `grpc-port` - gRPC API port for repository packages
- 🌐 - `FILE_PORT` - `file-port` - port for static file server to access packages
- 📦 - `INIT_PKGS` - `init-pkgs` - initial packages to download on start
- 📒 - `LOGS_FMT` - `logs-fmt` - format for logs (can be text/json/pretty)

You can run repository via `docker-compose`:

```yml
services:
  pacman:
    image: gitea.dancheg97.ru/dancheg97/go-pacman:latest
    command: run
    environment:
      INIT_PKGS: onlyoffice-bin yay
      LOGS_FMT: pretty
    ports:
      - 9080:9080
      - 8080:8080
```

---

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
