package runner

import (
	"bytes"
	"context"
	"encoding/json"
	event "faas/internal/events"
	"faas/internal/responses"
	"fmt"
	"os/exec"
	"time"
)

const timeout = 30 * time.Second

type Runner struct {
}

func New() *Runner {
	return &Runner{}
}

func (r *Runner) Run(ev event.HTTPEndpoint) responses.Output {

	st, err := json.Marshal(ev)

	if err != nil {
		fmt.Printf("Error marshalling event: %v\n", err)
		return responses.FromError(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	cmd := exec.CommandContext(ctx, "/opt/faas_dev/to_run", string(st))

	var res bytes.Buffer

	cmd.Stdout = &res

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error starting subprocess: %v\n", err)
		return responses.FromError(err)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Printf("Error waiting for subprocess: %v\n", err)
		return responses.FromError(err)
	}

	fmt.Printf("Subprocess output: %s\n", res.String())

	return responses.Output{
		Message: "Success",
		Body:    res.String(),
	}
}
