package json

import (
	"encoding/json"
	"log"
	"net/http"
	"../../entity"
)

//ParseJSON implement parsing functions from format json to Post struct
func ParseJSON(res *http.Response) []Post{

	p := []Post{}

	err := json.NewDecoder(res.Body).Decode(&p)

	if err != nil{
		log.Println("failed to decode json to struct post : ",err)
	}
	return p
}

