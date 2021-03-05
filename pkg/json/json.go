package json

import (
	"awesomeProject/nix/entity"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//ParseJSON implement parsing functions from format json to Post struct
func ParseJSON(res *http.Response) []entity.Post {

	p := []entity.Post{}

	err := json.NewDecoder(res.Body).Decode(&p)

	fmt.Println(p)

	if err != nil {
		log.Println("failed to decode json to struct post : ", err)
	}
	return p
}

