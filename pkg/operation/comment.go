package operation

import (
	"awesomeProject/nix/entity"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// GetComment required for getting comments from comment on JSON placeholder
func GetComment(id int) []entity.Comment {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%d", id)

	response, err := http.Get(url)

	if err != nil {
		log.Fatal("failed to get json command : ", err)
	}

	c := []entity.Comment{}

	err = json.NewDecoder(response.Body).Decode(&c)

	if err != nil {
		log.Fatal("failed in decoding json to comment : ", err)
	}

	if err = response.Body.Close(); err != nil {
		log.Fatal("failed to close response.Body : ", err)
	}
	return c
}
