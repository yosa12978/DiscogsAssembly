package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yosa12978/DiscogsAssembly/helpers"
	"github.com/yosa12978/DiscogsAssembly/models"
	"github.com/yosa12978/DiscogsAssembly/repos"
)

type ImageService interface {
	Download(uri, output string) error
}

type imageService struct {
	discogsRepo repos.DiscogsRepo
}

func NewImageService() ImageService {
	service := new(imageService)
	service.discogsRepo = repos.NewDiscogsRepo()
	return new(imageService)
}

func (service *imageService) Download(url, output string) error {
	release, err := service.discogsRepo.GetRelease(url)
	if err != nil {
		return err
	}
	path, err := service.createPath(output, release)
	if err != nil {
		return err
	}
	if err := service.downloadImages(path, release); err != nil {
		return err
	}
	service.saveMetadata(path, release)
	return nil
}

func (service *imageService) createPath(output string, release *models.Release) (string, error) {
	path := fmt.Sprintf("%s/%s %d", output, release.Title, release.Id)
	fmt.Printf("%d path: %v\n", release.Id, path)
	err := os.MkdirAll(path, 0700)
	return path, err
}

func (service *imageService) downloadImages(path string, release *models.Release) error {
	for i := 0; i < len(release.Images); i++ {
		imagePath := fmt.Sprintf("%s/img%d.jpeg", path, i)
		err := helpers.DownloadFile(release.Images[i].Uri, imagePath)
		if err != nil {
			log.Printf(err.Error())
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}

func (service *imageService) saveMetadata(path string, release *models.Release) {
	metadata := models.ToMetadata(release)
	metadataPath := fmt.Sprintf("%s/discogs.json", path)
	metadataFile, err := os.Create(metadataPath)
	if err != nil {
		log.Fatalf("error creating metadata file")
	}
	defer metadataFile.Close()
	enc := json.NewEncoder(metadataFile)
	enc.SetIndent("", "    ")
	enc.Encode(metadata)
}
