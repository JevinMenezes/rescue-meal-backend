package rest

import "net/http"

func (s login.service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		if err := json.NewDecoder(r.Body).Decode(&req);
	}
}