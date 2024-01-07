package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/yosa12978/DiscogsAssembly/models"
)

type MetadataService interface {
	SaveMetadata(path, name string, release *models.Release) error
}

type metadataService struct {
}

func NewMetadataService() MetadataService {
	return new(metadataService)
}

func (service *metadataService) SaveMetadata(path, name string, release *models.Release) error {
	metadata := models.ToMetadata(release)
	os.MkdirAll(path, 0700)
	metadataPath := fmt.Sprintf("%s/%s.json", path, name)
	if _, err := os.Stat(metadataPath); !errors.Is(err, os.ErrNotExist) {
		return err
	}
	metadataFile, err := os.Create(metadataPath)
	if err != nil {
		return err
	}
	defer metadataFile.Close()
	enc := json.NewEncoder(metadataFile)
	enc.SetIndent("", "    ")
	return enc.Encode(metadata)
}
