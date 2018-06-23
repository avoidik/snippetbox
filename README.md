# Demo Go Application

This is the demo application which was developed throughout Alex Edwards [Let's Go](https://lets-go.alexedwards.net/) book with minor changes 

- [sqlite3](https://github.com/mattn/go-sqlite3) instead of [mysql](https://github.com/golang/go/wiki/SQLDrivers)
- database initialization using transactions
- console interrupts and graceful shutdown support
- [alice](https://github.com/justinas/alice) for middleware chains
- [glide](https://github.com/Masterminds/glide) dependecies manager
- additional checks for command-line parameters

## How to run

```
git clone https://github.com/avoidik/snippetbox
cd snippetbox
go run $GOROOT/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
go run cmd/web/*
```

## Useful links

Alex Edwards blog
https://www.alexedwards.net/blog/

Go Basics
http://openmymind.net/The-Little-Go-Book/

Pat - muxer
https://github.com/bmizerany/pat

Bone - muxer
https://github.com/go-zoo/bone

Gorilla - muxer
https://github.com/gorilla/mux

Gorilla - sessions
https://github.com/gorilla/sessions

SCS - sessions
https://github.com/alexedwards/scs

Alice - middleware chaining
https://github.com/justinas/alice

Nosurf - CSRF protection middleware
https://github.com/justinas/nosurf

Gorilla - CSRF protection middleware
https://github.com/gorilla/csrf
