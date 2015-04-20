package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

func ContainerLogFilePaths(name string) (paths []string) {
	containerIds := ContainerIdsByName(name)
	if len(containerIds) <= 0 {
		return
	}

	for _, cid := range containerIds {

		filePath := fmt.Sprintf("%s-json.log", cid)
		path := filepath.Join(DockerRoot, "containers", cid, filePath)
		paths = append(paths, path)
	}
	return
}

func ContainerIdsByName(name string) (containerIds []string) {
	pattern := filepath.Join(DockerRoot, "containers/*/config.json")
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

		config := &Config{}
		err = json.Unmarshal(file, &config)
		if err != nil {
			logger.Error("GetContainerIdByName.Json", "err", err)

		}

		if config.Name == "/"+name {
			containerIds = append(containerIds, config.ID)
		}
	}

	return
}
