package engine

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/kon3gor/oshawott/internal/core"
)

type oshawottHandler func(http.ResponseWriter, *http.Request, Engine)

type httpHandler func(http.ResponseWriter, *http.Request)

type saveReq struct {
	Url string `json:"url"`
}

func wrap(engine Engine, handler oshawottHandler) httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, engine)
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func save(w http.ResponseWriter, r *http.Request, engine Engine) {
	if r.Method == http.MethodGet {
		index(w, r)
		return
	}

	defer r.Body.Close()
	var body saveReq
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		//todo: handle this shit
	}

	url := body.Url
	if url == "" {
		//todo: handle this shit
	}

	key, err := engine.SaveUrl(url)
	fmt.Println(err)
	fmt.Fprint(w, fmt.Sprintf("http://localhost:8080/%s", key))
}

func resolve(w http.ResponseWriter, r *http.Request, engine Engine) {
	path := r.URL.Path
	_, hash, _ := strings.Cut(path, "/")
	url, found := core.GetValue(e.ctx, hash)
	if !found {
		//todo: handle this shit
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusSeeOther)

}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, core.IndexPage)
}
