package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/kon3gor/oshawott"
	"github.com/kon3gor/oshawott/internal/core"
)

type saveReq struct {
	Url string `json:"url"`
}

func (s server) ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func (s server) save(w http.ResponseWriter, r *http.Request) {
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

	key, err := s.engine.SaveUrl(url)
	fmt.Println(err)
	fmt.Fprint(w, key)
}

func (s server) resolve(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	_, hash, _ := strings.Cut(path, "/")
	url, err := s.engine.GetUrl(oshawott.Key(hash))
	if err != nil {
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
