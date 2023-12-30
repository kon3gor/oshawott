package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kon3gor/oshawott"
)

type server struct {
	engine oshawott.Engine
	ctx    oshawott.AppContext
}

func Start(ctx oshawott.AppContext, e oshawott.Engine) {
	port := fmt.Sprintf(":%d", ctx.Port)

	s := server{e, ctx}
	http.HandleFunc("/save", s.save)
	http.HandleFunc("/ping", s.ping)
	http.HandleFunc("/", s.resolve)

	fmt.Printf("Running on port %d\n", ctx.Port)
	log.Fatalln(http.ListenAndServe(port, nil))
}
