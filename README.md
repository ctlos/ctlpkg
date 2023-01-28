<p align="center">
<img style="align: center; padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="238px" height="238px" src="https://gitea.dancheg97.ru/repo-avatars/14-4297f15da3e76c29478ec89973007622" />
</p>

<h2 align="center">Go gRPC template</h2>

[![Generic badge](https://img.shields.io/badge/LICENSE-MIT-red.svg)](https://gitea.dancheg97.ru/templates/go-sqlc/src/branch/main/LICENSE)
[![Generic badge](https://img.shields.io/badge/GITEA-REPO-blue.svg)](https://gitea.dancheg97.ru/templates/go-sqlc)
[![Build Status](https://drone.dancheg97.ru/api/badges/templates/go-sqlc/status.svg)](https://drone.dancheg97.ru/templates/go-sqlc)

Example description.

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

---

Repository stack:

- [goose](https://github.com/pressly/goose) - write your migrations in sql format and simply add them to `/migrations` folder in goose compatible format. You can test migrations by running. `docker-compose` or when generating `sqlc` code for queries.
- [sqlc](https://github.com/kyleconroy/sqlc) - you can simply write your queries for database in `sqlc.sql`, and they will be automatically generated. `sqlc` is a compiler, so queries are automatically checked for compatability with provided via migrations database schema.
- [grpc](https://grpc.io/) - you can describe your service and messages in `proto` file code for provided schema will be generated, type safe and fast API. Multiple middleware solutions are included with service.
- [docker](https://www.docker.com/) - 2 stage `dockerfile` and clever `docker-compose` with migrations separated to different container (for optional horizontal scaling of service) are already provided with service.
- [golangci-lint](https://golangci-lint.run/) - linter config stored in root directory, already includes a lot of usefull linters, mostly commented with helpful descriptions of different linters. You can tune it for your personal needs.
- [drone](https://www.drone.io/) - already written pipeline for `drone-ci`, which includes `lint`, `test`, `build` and pushing service to registry.
- [logrus](https://github.com/sirupsen/logrus) - one of the most popular libraries for logging in go, has perfect interfaces and adapters for both grpc server and postgres modules.
- [postgres](https://www.postgresql.org/) - simply best relational database of all time.
