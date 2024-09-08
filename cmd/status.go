package cmd

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/ngc7293/plink/internal/api"
	"github.com/ngc7293/plink/internal/tui"
)

func ShowStatus() error {
	config, err := ReadConfig()

	if err != nil {
		slog.Error("could not read config", "error", err)
		return err
	}

	job, err := api.GetJob(config.Host, config.Username, config.Password)

	if err != nil {
		slog.Error("could not get status", "error", err)
		return err
	}

	if job == nil {
		fmt.Println("No job currently running")
		return err
	}

	fmt.Print("\n\n")
	for {
		job, err = api.GetJob(config.Host, config.Username, config.Password)
		if err != nil {
			slog.Error("could not get status", "error", err)
			return err
		}

		elapsed := time.Duration(job.TimePrinting * int(time.Second))
		remaining := time.Duration(job.TimeRemaining * int(time.Second))

		fmt.Print(tui.DrawStatusWidget(job.State, job.File.DisplayName, job.Progress/100, elapsed, remaining))
		time.Sleep(100 * time.Millisecond)
	}
}
