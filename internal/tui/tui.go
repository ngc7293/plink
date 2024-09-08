package tui

import (
	"fmt"
	"time"
)

func DrawProgressBar(progress float64, elapsed time.Duration, remaining time.Duration, width int) string {

	text := fmt.Sprintf("%3.0f%% (%v/%v)", progress*100, elapsed, elapsed+remaining)

	spaces := max(width-len(text), 0)

	line := "["
	for i := 1; i < spaces/2; i++ {
		line += "-"
	}
	line += text
	for i := 1; i < spaces-(spaces/2); i++ {
		line += "-"
	}
	line += "]"

	count := int(progress * float64(width))

	return "\x1b[38;2;255;255;255m\x1b[48;2;250;104;49m" + line[:count] + "\x1b[0m\x1b[38;2;250;104;49m" + line[count:] + "\x1b[0m"
}

func DrawState(state string) string {
	text := "Unknown"

	switch state {
	case "PRINTING":
		text = "\x1b[32mPrinting\x1b[0m"
	}

	return text
}

func DrawStatusWidget(state string, file string, progress float64, elapsed time.Duration, remaining time.Duration) string {
	clearall := "\x1b[2K\x1b[1F\x1b[2K\x1b[1F\x1b2K\r"
	statusline := fmt.Sprintf("[ %s ] %s", DrawState(state), file)
	progressline := DrawProgressBar(progress, elapsed, remaining, 72)

	return fmt.Sprintf("%s%s\n\n%s", clearall, statusline, progressline)
}
