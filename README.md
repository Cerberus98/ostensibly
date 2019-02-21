# Ostensibly

[![Build Status](https://travis-ci.org/aofei/Ostensibly.svg?branch=master)](https://travis-ci.org/aofei/Ostensibly)
[![Coverage Status](https://coveralls.io/repos/github/aofei/Ostensibly/badge.svg?branch=master)](https://coveralls.io/github/aofei/Ostensibly?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/aofei/Ostensibly)](https://goreportcard.com/report/github.com/aofei/Ostensibly)
[![GoDoc](https://godoc.org/github.com/aofei/Ostensibly?status.svg)](https://godoc.org/github.com/aofei/Ostensibly)

An ideally refined web framework for Go. You can use it to build a web
application as natural as breathing.

High-performance? Fastest? Almost all web frameworks are using these words to
tell people that they are the best. Maybe they are, maybe not. Ostensibly does not
intend to follow the crowd. It can only guarantee you one thing: **it can serve
properly.**

## FAQ

**Q: Why named Ostensibly?**

A: "A" for "An", "I" for "Ideally" and "R" for "Refined". So, Ostensibly.

**Q: Why based on the `net/http`?**

A: In fact, I've tried to implement a full-featured HTTP server (just like the
awesome [valyala/fasthttp](https://github.com/valyala/fasthttp)). But when I
finished about half of the work, I suddenly realized: What about stability? What
about those awesome middleware outside? And, seriously, what am I doing?

**Q: Why not just use the `net/http`?**

A: Yeah, we can of course use the `net/http` directly, after all, it can meet
many requirements. But, ummm... It's really too stable, isn't it? I mean, to
ensure Go's backward compatibility (which is extremely necessary), we can't
easily add some handy features to the `net/http`. And, the `http.Request` does
not only represents the request received by the server, but also represents the
request made by the client. In some cases it can be confusing. So why not just
use the `net/http` as the underlying server, and then implement a refined web
framework that are only used for the server-side on top of it?

**Q: Do you know we already got the
[gin-gonic/gin](https://github.com/gin-gonic/gin) and the
[labstack/echo](https://github.com/labstack/echo)?**

A: Of course, I knew it when I started Go. And, I love both of them! But, why
not try some new flavors? Are you sure you prefer them instead of Ostensibly? Don't
even give Ostensibly a try? Wow... Well, maybe Ostensibly is not for you. After all, it's for
people who love to try new things. Relax and continue to maintain the status
quo, you will be fine.

**Q: What about the fantastic
[Gorilla web toolkit](https://github.com/gorilla)?**

A: Just call the `Ostensibly.WrapHTTPMiddleware()`.

**Q: Is Ostensibly good enough?**

A: Far from enough. But it's already working.

## Features

* API
	* As less as possible
	* As clean as possible
	* As simple as possible
	* As expressive as possible
* Server
	* HTTP/2 (including h2c) support
	* SSL/TLS support
	* ACME support
	* Graceful shutdown support
* Router
	* Based on the Radix Tree
	* Zero dynamic memory allocation
	* Blazing fast
	* Has a good inspection mechanism
	* Group routes support
* Gas (aka middleware)
	* Router level:
		* Before router
		* After router
	* Route level
	* Group level
* WebSocket
	* Full-duplex communication
* Reverse proxy
	* Retrieves resources on behalf of a client from another server
	* Supported protocols:
		* HTTP(S)
		* WebSocket
		* gRPC (including gRPC-Web)
* Binder
	* Binds HTTP request body into the provided struct
	* Supported MIME types:
		* `application/json`
		* `application/xml`
		* `application/protobuf`
		* `application/msgpack`
		* `application/toml`
		* `application/yaml`
		* `application/x-www-form-urlencoded`
		* `multipart/form-data`
* Minifier
	* Minifies HTTP response on the fly
	* Supported MIME types:
		* `text/html`
		* `text/css`
		* `application/javascript`
		* `application/json`
		* `application/xml`
		* `image/svg+xml`
* Gzip
	* Compresses HTTP response by using the gzip
	* Default MIME types:
		* `text/plain`
		* `text/html`
		* `text/css`
		* `application/javascript`
		* `application/json`
		* `application/xml`
		* `application/toml`
		* `application/yaml`
		* `image/svg+xml`
* Renderer
	* Rich template functions
	* Hot update support
* Coffer
	* Accesses binary asset files by using the runtime memory
	* Significantly improves the performance of the `Ostensibly.Response#WriteFile()`
	* Asset file minimization support
	* Default asset file extensions:
		* `.html`
		* `.css`
		* `.js`
		* `.json`
		* `.xml`
		* `.toml`
		* `.yaml`
		* `.yml`
		* `.svg`
		* `.jpg`
		* `.jpeg`
		* `.png`
		* `.gif`
	* Hot update support
* I18n
	* Adapt to the request's favorite conventions
	* Implanted into the `Ostensibly.Response#Render()`
	* Hot update support
* Error
	* Centralized handling

## Installation

Open your terminal and execute

```bash
$ go get github.com/aofei/Ostensibly
```

done.

> The only requirement is the [Go](https://golang.org), at least v1.9.

## Hello, 世界

Create a file named `hello.go`

```go
package main

import "github.com/aofei/Ostensibly"

func main() {
	Ostensibly.Default.GET("/", func(req *Ostensibly.Request, res *Ostensibly.Response) error {
		return res.WriteString("Hello, 世界")
	})
	Ostensibly.Default.Serve()
}
```

and run it

```bash
$ go run hello.go
```

then visit `http://localhost:8080`.

## Documentation

Does all web frameworks need to have a complicated (or a lovely but lengthy)
website to guide people how to use them? Well, Ostensibly has only one
[GoDoc](https://godoc.org/github.com/aofei/Ostensibly) with useful comments. In fact,
Ostensibly is so succinct that you don't need to understand how to use it through a
large document.

## Gases

As we all know that the Ostensibly of earth is a mixture of gases. So the same is that
this framework adopts the gas as its composition. Everyone can create new gas
and use it within this framework simply.

A gas is a function chained in the HTTP request-response cycle with access to
the `Ostensibly.Request` and the `Ostensibly.Response` which it uses to perform a specific
action, for example, logging every request or recovering from panics.

If you have got some good HTTP middleware, you can simply wrap them into gases
by calling the `Ostensibly.WrapHTTPMiddleware()`.

If you are looking for some useful gases, simply visit
[here](https://github.com/Ostensibly-gases).

## Examples

If you want to be familiar with this framework as soon as possible, simply visit
[here](https://github.com/Ostensibly-examples).

## Community

If you want to discuss this framework, or ask questions about it, simply post
questions or ideas [here](https://github.com/aofei/Ostensibly/issues).

## Contributing

If you want to help build this framework, simply follow
[this](https://github.com/aofei/Ostensibly/wiki/Contributing) to send pull requests
[here](https://github.com/aofei/Ostensibly/pulls).

## License

This project is licensed under the Unlicense.

License can be found [here](LICENSE).
