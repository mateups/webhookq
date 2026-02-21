package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"webhooq/internal/config"
	"webhooq/internal/targets"
)

type listTargetsResponse struct {
	Items []targets.Target `json:"items"`
}

func Run(cfg config.Config, out io.Writer) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	mux.HandleFunc("/v1/targets", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, http.StatusOK, listTargetsResponse{Items: []targets.Target{}})
		case http.MethodPost:
			writeJSON(w, http.StatusNotImplemented, map[string]string{"error": "not implemented yet"})
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	srv := &http.Server{
		Addr:              cfg.APIListenAddr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	fmt.Fprintf(out, "api listening on %s\n", cfg.APIListenAddr)
	return srv.ListenAndServe()
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
