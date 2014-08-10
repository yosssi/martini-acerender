# Martini Acerender

[![wercker status](https://app.wercker.com/status/1cdc361a2374b8f28022c7fd2f8792b5/m "wercker status")](https://app.wercker.com/project/bykey/1cdc361a2374b8f28022c7fd2f8792b5)
[![GoDoc](https://godoc.org/github.com/yosssi/martini-acerender?status.svg)](https://godoc.org/github.com/yosssi/martini-acerender)
[![Coverage Status](https://img.shields.io/coveralls/yosssi/martini-acerender.svg)](https://coveralls.io/r/yosssi/martini-acerender?branch=master)

## Overview

Martini Acerender is a Martini middleware/handler for parsing [Ace](https://github.com/yosssi/ace) templates and rendering HTML.

## Example

main.go

```go
package main

import (
	"github.com/go-martini/martini"
	"github.com/yosssi/martini-acerender"
)

func main() {
	m := martini.Classic()
	m.Use(acerender.Renderer(acerender.Options{BaseDir: "assets"}))
	m.Get("/", func(r acerender.Render) {
		r.HTML(200, "base:inner", map[string]string{"Msg": "Hello Acerender"}, nil)
	})
	m.Run()
}
```

base.ace

```ace
= doctype html
html lang=en
  head
    meta charset=utf-8
    meta http-equiv=Content-Type content="text/html;charset=UTF-8"
    meta name=viewport content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0"
    title Ace example
    = css
      h1 { color: blue; }
  body
    h1 Base Template : {{.Msg}}
    #container.wrapper
      = yield main
      = yield sub
    = javascript
      alert('{{.Msg}}');
```

inner.ace

```ace
= content main
  h2 Inner Template - Main : {{.Msg}}

= content sub
  h3 Inner Template - Sub : {{.Msg}}
```
