package main

import (
	"testing"
)

func Test_ContainerLogPath(t *testing.T) {
	DockerRoot = "./test-fixtures/"
	path := ContainerLogFilePath("a")

	if path == "" {
		t.Error("path is empty")
	}
}
