<p align="center">
<img style="align: center; padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="238px" height="238px" src="./assets/images/logo.png" />
</p>

<h2 align="center">Ctlos package repository - ctlpkg</h2>

[![Generic badge](https://img.shields.io/badge/LICENSE-GPLv3-orange.svg)](https://dancheg97.ru/dancheg97/ctlpkg/src/branch/main/LICENSE)
[![Generic badge](https://img.shields.io/badge/GITEA-REPO-red.svg)](https://dancheg97.ru/dancheg97/ctlpkg)
[![Generic badge](https://img.shields.io/badge/GITHUB-REPO-white.svg)](https://github.com/ctlos/ctlpkg)
[![Generic badge](https://img.shields.io/badge/DOCKER-REGISTRY-blue.svg)](https://dancheg97.ru/dancheg97/-/packages/container/ctlpkg/latest)
[![Build Status](https://drone.dancheg97.ru/api/badges/dancheg97/ctlpkg/status.svg)](https://drone.dancheg97.ru/dancheg97/ctlpkg)

Dockerized pacman repository with friendly user interface and public API. Project goal is to quickly set up personal pacman repostitory without pain and hustle.

---

## Environment configurations

Environment variables/flags:

- 📄 - `CTLPKG_REPO` - `repo` - repository name on the web page
- 😀 - `CTLPKG_USER` - `user` - user name in system, will be used to eject `yay` packages
- 🌐 - `CTLPKG_GRPC_PORT` - `grpc-port` - port for gRPC API
- 🌐 - `CTLPKG_HTTP_PORT` - `http-port` - port for static file servers and rest API
- 📦 - `CTLPKG_INIT_PKGS` - `init-pkgs` - initial packages to download on start
- 📒 - `CTLPKG_LOGS_FMT` - `logs-fmt` - format for logs (can be text/json/pretty)
- 📂 - `CTLPKG_WEB_DIR` - `web-dir` - directory with flutter web app

## Start docker service

You can run repository via `docker-compose`:

```yml
services:
  pacman:
    image: dancheg97.ru/dancheg97/ctlpkg:latest
    command: run
    environment:
      INIT_PKGS: onlyoffice-bin yay
      LOGS_FMT: pretty
    volumes:
      - ./go-pacman:/var/cache/pacman/pkg
    ports:
      - 9080:9080
      - 8080:8080
```

## Roadmap

- add list with added packages to flutter app
- add auth to flutter app
- add ability to update and add packages from UI
- fix problem with static file servers
