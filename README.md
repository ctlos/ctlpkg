<p align="center">
<img style="align: center; padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="238px" height="238px" src="./assets/images/logo.png" />
</p>

<h2 align="center">CtlOS package repository - ctlpkg</h2>

[![Generic badge](https://img.shields.io/badge/LICENSE-GPLv3-orange.svg)](https://dancheg97.ru/dancheg97/ctlpkg/src/branch/main/LICENSE)
[![Generic badge](https://img.shields.io/badge/GITEA-REPO-red.svg)](https://dancheg97.ru/dancheg97/ctlpkg)
[![Generic badge](https://img.shields.io/badge/GITHUB-REPO-white.svg)](https://github.com/ctlos/ctlpkg)
[![Generic badge](https://img.shields.io/badge/DOCKER-REGISTRY-blue.svg)](https://dancheg97.ru/dancheg97/-/packages/container/ctlpkg/latest)
[![Build Status](https://drone.dancheg97.ru/api/badges/dancheg97/ctlpkg/status.svg)](https://drone.dancheg97.ru/dancheg97/ctlpkg)

Dockerized pacman repository with friendly user interface and public API. Project goal is to quickly set up personal pacman repostitory without pain and hustle.

---

## Configurations

Environment variables/flags:

- 📄 - `CTLPKG_REPO` - `repo` - repository name on the web page
- 😀 - `CTLPKG_USER` - `user` - user name in system, will be used to eject `yay` packages
- 🌐 - `CTLPKG_PORT` - `port` - publically exposed port, `8080` default
- 📫 - `CTLPKG_API_ADRESS` - `api-adress` - adress for backend api calls via `grpc-web`
- 📦 - `CTLPKG_INIT_PKGS` - `init-pkgs` - initial packages to download on start
- 📒 - `CTLPKG_LOGS_FORMAT` - `logs-fmt` - format for logs (can be text/json/pretty)
- 📂 - `CTLPKG_WEB_DIR` - `web-dir` - directory with flutter web app
- 🔐 - `CTLPKG_LOGINS` - `logins` - list of logins and passwords separated by '|' symbol

---

## Deploy

- with `docker`:

```sh
docker run -p 8080:8080 -e CTLPKG_LOGS_FMT=text dancheg97.ru/dancheg97/ctlpkg:latest
```

- with `docker-compose`:

```yml
services:
  pacman:
    image: dancheg97.ru/dancheg97/ctlpkg:latest
    command: run
    environment:
      CTLPKG_INIT_PKGS: yay
      CTLPKG_API_ADRESS: http://localhost:8080/
      CTLPKG_LOGS_FMT: text
      CTLPKG_LOGINS: user1|pass1|user2|pass2
    ports:
      - 8080:8080
```

## Add to pacman conf

Add those lines to your `/etc/pacman.conf`, to get things to work:

```conf
[localhost]
SigLevel = Optional TrustAll
Server = http://localhost:8080/pkg
```

You can test it with this commands:

```sh
sudo pacman -R yay
sudo pacman -Sy yay
```

## Contribute

For applicaiton development you need to install following software:

- `go`
- `gofumpt`
- `golangci-lint`
- `buf`
- `flutter`
- `flutter webkit`

All frontend dart code is located in `lib` folder, all backend go code is
located in `cmd` folder.
