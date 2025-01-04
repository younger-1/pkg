## Building an Awesome CLI App in Go – OSCON 2017

> <https://spf13.com/presentation/building-an-awesome-cli-app-in-go-oscon/>
> A workshop outlining the techniques, principles, and libraries you need to create user-friendly command-line interfaces and command suites.
> It cover everything from how to design and build commands to working with and parsing flags, config files and remote config systems, and how to work with environment variables and 12-factor apps.

This workshop covers:

- CLI application design
- Introduction to the Go programming language
- Introduction to the Cobra CLI framework (used by Kubernetes, Docker, Hugo, etc)
- Building a Todo application in Go using Cobra

```sh
# Install Cobra-CLI generator
# @see <https://github.com/spf13/cobra-cli/blob/main/README.md>
go install github.com/spf13/cobra-cli@latest

go mod init github.com/younger-1/code-playground/go/tri                                                                                       at  github.com/younger-1/code-playground

cobra-cli init --author "Xavier Young" --license apache --viper
cobra-cli add add --author "Xavier Young" --license apache
cobra-cli add list --author "Xavier Young" --license apache
```
