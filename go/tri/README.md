# Building an Awesome CLI App in Go – OSCON 2017

> <https://spf13.com/presentation/building-an-awesome-cli-app-in-go-oscon/>
> A workshop outlining the techniques, principles, and libraries you need to create user-friendly command-line interfaces and command suites.
> It cover everything from how to design and build commands to working with and parsing flags, config files and remote config systems, and how to work with environment variables and 12-factor apps.

This workshop covers:

- CLI application design
- Introduction to the Go programming language
- Introduction to the Cobra CLI framework (used by Kubernetes, Docker, Hugo, etc)
- Building a Todo application in Go using Cobra

## Dev

```sh
# Install Cobra-CLI generator
# @see <https://github.com/spf13/cobra-cli/blob/main/README.md>
go install github.com/spf13/cobra-cli@latest

go mod init github.com/younger-1/pkg/go/tri                                                                                       at  github.com/younger-1/pkg

cobra-cli init --author "Xavier Young" --license apache --viper
cobra-cli add add --author "Xavier Young" --license apache
cobra-cli add list --author "Xavier Young" --license apache
```

## Example

```sh
go run . add "eat fruits"
go run . add hello world -p3
go run . add "format list" -p1                                                                                                          at  github.com/younger-1/pkg
go run . list
go run . done 2
go run . list --done
go run . list --all
TRI_DATAFILE=~/.tri-data.json go run . list
```

## 12-Factor 要求代码和配置严格分离

> @see <https://12factor.net/zh_cn/config>

配置在不同部署间存在差异。判断一个应用是否正确地将配置排除在代码之外，一个简单的方法是看该应用的基准代码是否可以立刻开源，而不用担心会暴露任何敏感的信息。

- 在代码中使用常量保存配置
- 使用配置文件，但不把它们纳入版本控制系统
- 12-Factor推荐将应用的配置存储于 环境变量 中

Config options priority

- command-line options/flags > env variables > config file > default in code

<https://github.com/ilyakaznacheev/cleanenv>

- <https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64>
- <https://dev.to/ilyakaznacheev/clean-configuration-management-in-golang-1c89>

<https://github.com/peterbourgon/ff>

- Flags-first package for configuration


## TODOs

- [ ] edit
- [ ] search
