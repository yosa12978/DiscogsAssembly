package repos

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
	"github.com/yosa12978/DiscogsAssembly/helpers"
	"github.com/yosa12978/DiscogsAssembly/models"
)

type DiscogsRepo interface {
	GetRelease(id string) (*models.Release, error)
	GetFolder(username, id string) (*models.FolderDetail, error)
	GetFolders(username string) ([]models.Folder, error)
	GetCurrentUser() (*models.User, error)
}

type discogsRepo struct {
}

func NewDiscogsRepo() DiscogsRepo {
	return new(discogsRepo)
}

func (repo *discogsRepo) GetRelease(id string) (*models.Release, error) {
	token := viper.Get("discogs.token")

	addr := fmt.Sprintf("https://api.discogs.com/releases/%s?token=%s", id, token)

	resp, err := helpers.HttpGet(addr)
	if err != nil {
		return nil, err
	}

	var release models.Release
	err = json.Unmarshal(resp, &release)
	return &release, err
}

func (repo *discogsRepo) GetCurrentUser() (*models.User, error) {
	token := viper.Get("discogs.token")
	addr := fmt.Sprintf("https://api.discogs.com/oauth/identity?token=%s", token)

	resp, err := helpers.HttpGet(addr)
	if err != nil {
		return nil, err
	}

	var user models.User
	err = json.Unmarshal(resp, &user)
	return &user, err
}

func (repo *discogsRepo) GetFolders(username string) ([]models.Folder, error) {
	addr := fmt.Sprintf("https://api.discogs.com/users/%s/collection/folders", username)
	resp, err := helpers.HttpGet(addr)
	if err != nil {
		return nil, err
	}
	var collection models.Collection
	err = json.Unmarshal(resp, &collection)
	return collection.Folders, err
}

func (repo *discogsRepo) GetFolder(username, id string) (*models.FolderDetail, error) {
	addr := fmt.Sprintf("https://api.discogs.com/users/%s/collection/%s", username, id)
	resp, err := helpers.HttpGet(addr)
	if err != nil {
		return nil, err
	}
	var folder models.FolderDetail
	err = json.Unmarshal(resp, &folder)
	return &folder, err
}
