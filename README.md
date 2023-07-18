# noiota

## Description

Linter to check if [iota](https://go.dev/ref/spec#Iota) is being used. `iota` may cause subtle bugs if declarations are re-ordered or `const` blocks merged changing constant values, specially if the values are serialized either to network or databases.

## Install

```shell
go install github.com/jacoelho/noiota
```
