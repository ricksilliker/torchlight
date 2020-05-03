package torchlight

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

type MayaScene struct {
	ScenePath string `json:"scene"`
	Extension string `json:"ext"`
}

type MayaController interface {
	GetFileReferences(string) ([]FileReference, error)
}

func GetMayaAsciiSceneData(scenePath string) (string, error) {
	fileBytes, err := ioutil.ReadFile(scenePath)
	if err != nil {
		logrus.Error("failed to parse maya scene")
		return "", err
	}

	fileStr := string(fileBytes)

	return fileStr, nil
}