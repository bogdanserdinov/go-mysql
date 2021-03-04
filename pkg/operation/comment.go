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

func getComment(p entity.Post) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%d",p.ID)

	response,err := http.Get(url)

	if err != nil{
		log.Fatal("failed to get json command : ",err)
	}

	c := []entity.Comment{}

	err = json.NewDecoder(response.Body).Decode(&c)

	if err != nil{
		log.Fatal("failed in decoding json to comment : ",err)
	}

	var mutex = &sync.Mutex{}

	for _,value := range c{
		go db.WriteToDBComment(value,mutex)
		fmt.Println(value)
	}


	err = response.Body.Close()
	if err != nil{
		log.Fatal("failed to close response.Body : ",err)
	}
}

