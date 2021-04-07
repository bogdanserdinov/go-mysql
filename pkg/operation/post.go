package operation

import (
	"fmt"
	"github.com/bogdanserdinov/go-mysql/entity"
	"github.com/bogdanserdinov/go-mysql/pkg/json"
	"log"
	"net/http"
)

// GetComment required for getting posts from post on JSON placeholder
func GetPosts(id int) []entity.Post {
	request := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId=%d", id)
	response, err := http.Get(request)
	if err != nil {
		log.Fatal("failed in get response from url : ", err)
	}
	p := json.ParseJSON(response) //parse post info
	return p
}
