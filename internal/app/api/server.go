package api

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"webhooq/internal/config"
)

func Run(cfg config.Config, out io.Writer, targetsHandler *TargetsHandler) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		newJSONResponse(w, nil)
	})

	mux.HandleFunc("GET /v1/targets", targetsHandler.ListTargets)
	mux.HandleFunc("POST /v1/targets", targetsHandler.CreateTarget)
	mux.HandleFunc("GET /v1/targets/{id}", targetsHandler.GetTarget)

	srv := &http.Server{
		Addr:              cfg.APIListenAddr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	fmt.Fprintf(out, "api listening on %s\n", cfg.APIListenAddr)
	return srv.ListenAndServe()
}
