package operation

import (
	"awesomeProject/nix/entity"
	"awesomeProject/nix/pkg/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// GetComment required for getting comments from comment on JSON placeholder
func GetComment(p entity.Post) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%d",p.ID)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("failed to get json command : ", err)
	}

	var c []entity.Comment

	err = json.NewDecoder(response.Body).Decode(&c)
	if err != nil {
		log.Fatal("failed in decoding json to comment : ", err)
	}

	if err = response.Body.Close(); err != nil {
		log.Fatal("failed to close response.Body : ", err)
	}

	var mutex = &sync.Mutex{}

	gormDB := db.OpenDataBase()
	for _,value := range c{
		go db.WriteToDBComment(value,gormDB,mutex)
	}
}
