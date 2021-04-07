package json

import (
	"encoding/json"
	"github.com/bogdanserdinov/go-mysql/entity"
	"log"
	"net/http"
)

//ParseJSON implement parsing functions from format json to Post struct
func ParseJSON(res *http.Response) []entity.Post {

	var p []entity.Post

	err := json.NewDecoder(res.Body).Decode(&p)

	if err != nil {
		log.Println("failed to decode json to struct post : ", err)
	}
	return p
}

