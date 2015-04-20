package main

import (
	"github.com/mgutz/logxi/v1"
	"github.com/namsral/flag"
	"os"
	"strings"
	"time"
)

var logger log.Logger

var (
	DockerRoot    string
	ContainerName string
	Output        string
	GraphiteHost  string
	Parser        string
	Prefix        string
	Interval      string
	LogsterPath   string
)

func init() {

	logger = log.NewLogger(os.Stdout, "logster-docker-runner:"+VERSION)

	flag.StringVar(&DockerRoot, "docker_root", "/var/lib/docker", "docker root")
	flag.StringVar(&ContainerName, "container_name", "", "container name")
	flag.StringVar(&Output, "output", "stdout", "logster output")
	flag.StringVar(&GraphiteHost, "graphite_host", "", "graphite host")
	flag.StringVar(&Parser, "parser", "", "logster parser")
	flag.StringVar(&Prefix, "prefix", "", "metric prefix")
	flag.StringVar(&Interval, "interval", "1m", "logster running interval")
	flag.StringVar(&LogsterPath, "logster_path", "logster", "logster file path")

}

func main() {
	flag.Parse()

	if Prefix != "" {
		if strings.Contains(Prefix, "{{HOSTNAME}}") {
			hostname, _ := os.Hostname()
			Prefix = strings.Replace(Prefix, "{{HOSTNAME}}", hostname, 1)
		}
	}

	duration, err := time.ParseDuration(Interval)
	if err != nil {
		logger.Error("interval", "err", err)
		os.Exit(1)
	}

	if ContainerName == "" {
		logger.Error("container name not found")
		os.Exit(1)
	}

	ticker := time.NewTicker(duration)

	logger.Info("start logster runner")

	for _ = range ticker.C {

		//Sometimes,docker can't remove old removed container fs. Even through `docker rm` was successful. docker 1.4.1
		for _, logPath := range ContainerLogFilePaths(ContainerName) {
			RunLogster(logPath)
		}

	}

	logger.Info("bye")
}
