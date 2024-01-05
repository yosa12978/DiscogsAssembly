package repos

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
	"github.com/yosa12978/DiscogsAssembly/helpers"
	"github.com/yosa12978/DiscogsAssembly/models"
)

var UnauthorizedError error = errors.New("401 Unauthorized")

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

	release_url := fmt.Sprintf("https://api.discogs.com/releases/%s?token=%s", id, token)
	s, err := helpers.HttpGet(release_url)
	if err != nil {
		return nil, err
	}

	var release models.Release
	err = json.Unmarshal([]byte(s), &release)
	return &release, err
}

func (repo *discogsRepo) GetCurrentUser() (*models.User, error) {
	token := viper.Get("discogs.token")
	addr := fmt.Sprintf("https://api.discogs.com/oauth/identity?token=%s", token)

	res, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == 401 {
		return nil, UnauthorizedError
	}
	b, err := io.ReadAll(res.Body)
	var user models.User
	err = json.Unmarshal(b, &user)
	return &user, err
}

func (repo *discogsRepo) GetFolders(username string) ([]models.Folder, error) {
	addr := fmt.Sprintf("https://api.discogs.com/users/%s/collection/folders", username)
	s, err := helpers.HttpGet(addr)
	if err != nil {
		return nil, err
	}
	var collection models.Collection
	err = json.Unmarshal([]byte(s), &collection)
	return collection.Folders, err
}

func (repo *discogsRepo) GetFolder(username, id string) (*models.FolderDetail, error) {
	addr := fmt.Sprintf("https://api.discogs.com/users/%s/collection/%s", username, id)
	s, err := helpers.HttpGet(addr)
	if err != nil {
		return nil, err
	}
	var folder models.FolderDetail
	err = json.Unmarshal([]byte(s), &folder)
	return &folder, err
}
