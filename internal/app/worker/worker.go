package worker

import (
	"fmt"
	"io"
	"time"

	"webhooq/internal/config"
)

func Run(cfg config.Config, out io.Writer) error {
	ticker := time.NewTicker(time.Duration(cfg.WorkerPollMs) * time.Millisecond)
	defer ticker.Stop()

	fmt.Fprintln(out, "worker started (skeleton mode)")
	for range ticker.C {
		fmt.Fprintln(out, "worker tick: claim + deliver pipeline not implemented yet")
	}

	return nil
}
