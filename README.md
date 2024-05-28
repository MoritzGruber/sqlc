# sqlc: A SQL Compiler

![go](https://github.com/MoritzGruber/sqlc/workflows/go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/MoritzGruber/sqlc)](https://goreportcard.com/report/github.com/MoritzGruber/sqlc)

sqlc generates **type-safe code** from SQL. Here's how it works:

1. You write queries in SQL.
1. You run sqlc to generate code with type-safe interfaces to those queries.
1. You write application code that calls the generated code.

Check out [an interactive example](https://play.sqlc.dev/) to see it in action, and the [introductory blog post](https://conroy.org/introducing-sqlc) for the motivation behind sqlc.

## Overview

- [Documentation](https://docs.sqlc.dev)
- [Installation](https://docs.sqlc.dev/en/latest/overview/install.html)
- [Playground](https://play.sqlc.dev)
- [Website](https://sqlc.dev)
- [Downloads](https://downloads.sqlc.dev/)
- [Community](https://discord.gg/EcXzGe5SEs)

## Supported languages

- [sqlc-gen-go](https://github.com/MoritzGruber/sqlc-gen-go)
- [sqlc-gen-kotlin](https://github.com/MoritzGruber/sqlc-gen-kotlin)
- [sqlc-gen-python](https://github.com/MoritzGruber/sqlc-gen-python)
- [sqlc-gen-typescript](https://github.com/MoritzGruber/sqlc-gen-typescript)

Additional languages can be added via [plugins](https://docs.sqlc.dev/en/latest/reference/language-support.html#community-language-support).

## Acknowledgments

sqlc was inspired by [PugSQL](https://pugsql.org/) and
[HugSQL](https://www.hugsql.org/).
