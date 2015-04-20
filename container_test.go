package main

import (
	"testing"
)

func Test_ContainerLogPath(t *testing.T) {
	DockerRoot = "./test-fixtures/"
	paths := ContainerLogFilePaths("a")

	if len(paths) <= 0 {
		t.Error("path is empty")
	}
}
