package main

import (
	"fmt"
	"os/exec"
)

func LogsterArgs(logPath string) (args []string) {

	args = append(args, fmt.Sprintf("--output=%s", Output))

	if Prefix != "" {
		args = append(args, fmt.Sprintf("--metric-prefix=%s", Prefix))
	}

	if GraphiteHost != "" {
		args = append(args, fmt.Sprintf("--graphite-host=%s", GraphiteHost))
	}

	args = append(args, Parser)
	args = append(args, logPath)

	return
}

func RunLogster(logFile string) {

	if logFile == "" {
		logger.Error("logfile not found", "file", logFile)
		return
	}

	logsterArgs := LogsterArgs(logFile)

	logger.Info("logster", "args", logsterArgs)

	cmd := exec.Command(LogsterPath, logsterArgs...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error("logster", "err", err.Error())
	}

	logger.Info("logster", "output", string(out))

}
