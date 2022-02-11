package main

import (
	//"encoding/json"
	"log"

	"github.com/gopherjs/gopherjs/js"
	"github.com/icza/screp/rep"
	"github.com/icza/screp/repparser"
)

type WebLogger struct{}

func (w *WebLogger) Write(p []byte) (n int, err error) {
	js.Global.Get("console").Call("log", string(p))
	return len(p), nil
}
type Rep struct {
	buf []byte
}

func New() *js.Object {
	return js.MakeWrapper(&Rep{})
}

func (p *Rep) Buffer() []byte {
	return p.buf
}

func (p *Rep) ParseBuffer(buf []byte) *rep.Replay {
	p.buf = buf
	// idk do some stuff here??
	r, _ := repparser.ParseConfig(buf, repparser.Config{Commands: true, MapData: true})
	// maybe serialize it to json before returning? but I feel gopherjs should be able to just return a pure JS object.
	return r
}

func main() {
	// logger for convenience, not really used
	var logger = log.New(&WebLogger{}, "", log.LstdFlags)

	// put Rep on the global object so we can access it from JS.
	js.Global.Set("Rep", map[string]interface{}{
		"New": New,
	})
	
	logger.Println("test from repinterface.go")
	
}
