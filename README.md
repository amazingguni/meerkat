# Meerkat

![Go](https://github.com/amazingguni/meerkat/workflows/Go/badge.svg)

Github에 호스팅되는 repository의 commit별 코드 품질을 저장하는 서버입니다.

## Getting Started

### Install Golang (go 1.15.6)

Please check the official golang installation guide before you start. [Official Documentation](https://golang.org/doc/install) Also make sure you have installed a go 1.15.6 version.

### Install Golang (go 1.15.6)

You need to make sure you defined the environment variable below.

- https://github.com/golang/go/wiki/SettingGOPATH

```sh
$ echo $GOPATH
/Users/xesina/go
$ echo $GOROOT
/usr/local/go/
$ echo $PATH
...:/usr/local/go/bin:/Users/xesina/test//bin:/usr/local/go/bin
```

### Install Dependency

```sh
$ go mod download
```

### Run

```sh
$ go run main.go
```

### Build

```sh
$ go build
```

### Run with fresh

Using [Fresh](https://github.com/gravityblast/fresh), you can build and (re)start your web application everytime you save a Go.

````sh
$ go get github.com/pilu/fresh
$ fresh
```

### Test

```sh
$ go test ./...
````
