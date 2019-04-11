package main

import (
	pkg "local/proxy_server/pipe"
	"log"
	"sync"
)

func main() {
	pipe := pkg.NewPipe("0.0.0.0:9006", "http://127.0.0.1:9001", false, false)
	var pipeWG sync.WaitGroup
	pipeWG.Add(1)
	go func(p *pkg.Pipe) {
		defer pipeWG.Done()
		defer log.Printf("pipe gg")
		p.Start()
	}(pipe)
	pipeWG.Wait()
	return
}
