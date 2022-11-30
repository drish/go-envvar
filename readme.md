# go-envvar

Opnionated utility package for handling environment variables, based on [brendanstennett/envvar](https://github.com/brendanstennett/envvar/)

---

![Build](https://github.com/drish/go-envvar/actions/workflows/ci.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/drish/go-envvar?status.svg)](https://godoc.org/github.com/drish/go-envvar)
![](https://img.shields.io/badge/license-MIT-blue.svg)


# Example config.yml

```yml

```

# Usage

```go
func init() {
    envvar.Load("../config.yaml", os.Getenv("APP_ENV"))
    // if no environment argument is passed, it defaults to `local`
}
```