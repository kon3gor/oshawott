package engine

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kon3gor/oshawott/internal/core"
	"github.com/kon3gor/oshawott/pkg/keys"
)

type Engine struct {
	ctx         core.OshawottContext
	keyProvider keys.KeyProvider
}

func (e Engine) SaveUrl(url string) (keys.Key, error) {
	saved, found := core.GetKey(e.ctx, url)
	if found {
		return keys.Key(saved), nil
	}

	key, err := e.keyProvider.GetKey(e.ctx, url)
	if err != nil {
		fmt.Println(err)
		return key, err
	}

	go core.SaveUrl(e.ctx, string(key), url)

	return key, nil
}

func (e Engine) Start() {
	port := fmt.Sprintf(":%d", e.ctx.Port)
	http.HandleFunc("/save", wrap(e, save))
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/", wrap(e, resolve))
	fmt.Printf(fmt.Sprintf("Running on port %d", e.ctx.Port))
	log.Fatalln(http.ListenAndServe(port, nil))
}
