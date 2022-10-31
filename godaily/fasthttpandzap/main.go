package main

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

/*
当访问 http://localhost:8081 时
打印日志 {"level":"info","ts":1665220173.801674,"caller":"fasthttpandzap/main.go:15","msg":"hello, go module","uri":"/"}

当访问 http://localhost:8081/foo/bar 时
{"level":"info","ts":1626614126.9899719,"caller":"hellomodule/main.go:15","msg":"hello, go module","uri":"/foo/bar"}
*/
var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	logger.Info("hello, go module", zap.ByteString("uri", ctx.RequestURI()))
}

func main() {
	fasthttp.ListenAndServe(":8081", fastHTTPHandler)
}
