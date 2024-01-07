package services

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/yosa12978/DiscogsAssembly/helpers"
	"github.com/yosa12978/DiscogsAssembly/models"
	"github.com/yosa12978/DiscogsAssembly/repos"
)

type ImageService interface {
	Download(release *models.Release, output string) error
}

type imageService struct {
	discogsRepo repos.DiscogsRepo
}

func NewImageService() ImageService {
	service := new(imageService)
	service.discogsRepo = repos.NewDiscogsRepo()
	return service
}

func (service *imageService) Download(release *models.Release, output string) error {
	path, err := service.createPath(output, release)
	if err != nil {
		return err
	}
	if err := service.downloadImages(path, release); err != nil {
		return err
	}
	return nil
}

func (service *imageService) createPath(output string, release *models.Release) (string, error) {
	path := fmt.Sprintf("%s/%s %d", output, release.Title, release.Id)
	err := os.MkdirAll(path, 0700)
	return path, err
}

func (service *imageService) downloadImages(path string, release *models.Release) error {
	var artists []string
	for i := 0; i < len(release.Artists); i++ {
		artists = append(artists, release.Artists[i].Name)
	}
	fmt.Printf("Downloading %s - %s\n", strings.Join(artists, ", "), release.Title)

	for i := 0; i < len(release.Images); i++ {
		imagePath := fmt.Sprintf("%s/img%d.jpeg", path, i)
		err := helpers.DownloadFile(release.Images[i].Uri, imagePath)
		if err != nil {
			fmt.Println(err.Error())
		}
		time.Sleep(1500 * time.Millisecond)
	}
	fmt.Println("Download complete")
	return nil
}
