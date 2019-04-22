package engine

const Version = "1.0.0"

const configuration = `
ip: 0.0.0.0
port: 8080
cookie: tigocookie
log:
  trace: stdout
  info: stdout
  warning: stdout
  error: stdout
`

const handler = `
package handlers

import "github.com/karldoenitz/Tigo/TigoWeb"

type PingHandler struct {
	TigoWeb.BaseHandler
}

func (p *PingHandler) Get() {
	p.ResponseAsText("PONG")
}
`

const main = `
package main

import (
	"github.com/karldoenitz/Tigo/TigoWeb"
	"%s/handlers"
)

var routers = []TigoWeb.Router{
	{"/ping", &handlers.PingHandler{}, nil},
}

func main() {
	application := TigoWeb.Application{ConfigPath: "./server.yaml", UrlRouters: routers}
	application.Run()
}
`
