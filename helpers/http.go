package helpers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func HttpGet(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return b, CheckStatus(res.StatusCode)
}

func DownloadFile(URL, fileName string) error {
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return fmt.Errorf("received %s", response.Status)
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}

var (
	ErrUnauthorized        error = errors.New("discogs: unauthorized")
	ErrNotFound                  = errors.New("discogs: not found")
	ErrForbidden                 = errors.New("discogs: forbidden")
	ErrMethodNotAllowed          = errors.New("discogs: method not allowed")
	ErrInternalServer            = errors.New("discogs: internal server error")
	ErrUnprocessableEntity       = errors.New("discogs: unprocessable entity")
)

func CheckStatus(code int) error {
	switch code {
	case 401:
		return ErrUnauthorized
	case 404:
		return ErrNotFound
	case 403:
		return ErrForbidden
	case 405:
		return ErrMethodNotAllowed
	case 422:
		return ErrUnprocessableEntity
	case 500:
		return ErrInternalServer
	default:
		return nil
	}
}
