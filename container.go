package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Configs struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

func ContainerLogFilePath(name string) (path string) {
	containerId := ContainerIdByName(name)
	if containerId == "" {
		return
	}

	filePath := fmt.Sprintf("%s-json.log", containerId)
	path = filepath.Join(DockerRoot, "containers", containerId, filePath)
	return
}

func ContainerIdByName(name string) (containerId string) {
	pattern := filepath.Join(DockerRoot, "containers/*/configs.json")
	paths, err := filepath.Glob(pattern)
	if err != nil {
		logger.Error("GetContainerIdByName.Glob", "err", err)
		return
	}

	for _, path := range paths {
		file, err := ioutil.ReadFile(path)
		if err != nil {
			logger.Error("GetContainerIdByName.ReadFile", "err", err)
		}

		configs := &Configs{}
		err = json.Unmarshal(file, &configs)
		if err != nil {
			logger.Error("GetContainerIdByName.Json", "err", err)

		}

		if configs.Name == "/"+name {
			return configs.ID
		}
	}

	return
}
