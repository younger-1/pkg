## How to Write Go Code

> <https://go.dev/doc/code>
> This document demonstrates the development of a simple Go package inside a module
> and introduces the go tool, the standard way to fetch, build, and install Go modules, packages, and commands.


1. Go programs are organized into packages.

2. A package is a collection of source files in the same directory that are compiled together.
Package name should match directory name.
Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.

3. A module is a collection of related Go packages that are released together.
Module is indicated by `go.mod`, which declares the **module path**: the *import path prefix* for all packages within the module.
The module contains the packages in the directory containing its `go.mod` file as well as subdirectories of that directory, up to the next subdirectory containing another `go.mod` file (if any).

