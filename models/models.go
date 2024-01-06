package models

type FolderDetail struct {
	Releases []CollectionItem
}

type CollectionItem struct {
	Id int `json:"id"`
}

type Release struct {
	Id      int      `json:"id"`
	Title   string   `json:"title"`
	Country string   `json:"country"`
	Uri     string   `json:"uri"`
	Artists []Artist `json:"artists"`
	Styles  []string `json:"styles"`
	Images  []Image  `json:"images"`
	Genres  []string `json:"genres"`
	Year    int      `json:"year"`
}

type Image struct {
	Uri string `json:"uri"`
}

type Artist struct {
	Name string `json:"name"`
}

type ReleaseMetadata struct {
	Id      int      `json:"discogs_id"`
	Uri     string   `json:"discogs_url"`
	Artists []string `json:"artists"`
	Title   string   `json:"title"`
	Genres  []string `json:"genres"`
	Styles  []string `json:"styles"`
	Country string   `json:"country"`
	Year    int      `json:"year"`
	Images  []string `json:"images"`
}

func ToMetadata(release *Release) *ReleaseMetadata {
	var images []string
	for i := 0; i < len(release.Images); i++ {
		images = append(images, release.Images[i].Uri)
	}
	var artists []string
	for i := 0; i < len(release.Artists); i++ {
		artists = append(artists, release.Artists[i].Name)
	}

	return &ReleaseMetadata{
		Id:      release.Id,
		Uri:     release.Uri,
		Title:   release.Title,
		Country: release.Country,
		Genres:  release.Genres,
		Styles:  release.Styles,
		Artists: artists,
		Year:    release.Year,
		Images:  images,
	}
}

type User struct {
	Id            int    `json:"id"`
	Username      string `json:"username"`
	Resource_url  string `json:"resource_url"`
	Consumer_name string `json:"consumer_name"`
}

type Collection struct {
	Folders []Folder `json:"folders"`
}

type Folder struct {
	Id    int    `json:"id"`
	Count int    `json:"count"`
	Name  string `json:"name"`
	URL   string `json:"resource_url"`
}

/*
   {
     "id": 0,
     "count": 23,
     "name": "All",
     "resource_url": "https://api.discogs.com/users/example/collection/folders/0"
   },
*/

type Config struct {
	Discogs Token `json:"discogs"`
}

type Token struct {
	Token string `json:"token"`
}
