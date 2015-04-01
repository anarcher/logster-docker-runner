package main

import (
	"fmt"
)

func LogsterArgs(logPath string) (args []string) {

	args = append(args, fmt.Sprintf("--output=%s", Output))

	if Prefix != "" {
		args = append(args, fmt.Sprintf("--metric-prefix=%s", Prefix))
	}

	if GraphiteHost != "" {
		args = append(args, fmt.Sprintf("--graphite-host=", GraphiteHost))
	}

	args = append(args, Parser)
	args = append(args, logPath)

	return
}
