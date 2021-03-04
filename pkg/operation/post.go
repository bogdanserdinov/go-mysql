package operation

import (
	"awesomeProject/nix/pkg/db"
	"awesomeProject/nix/pkg/json"
	"fmt"
	"log"
	"net/http"
)

func getPosts(id int){
	request := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId=%d",id)
	response, err := http.Get(request)
	if err != nil {
		log.Fatal("failed in get response from url : ",err)
	}

	p := json.ParseJSON(response)	//parse post info

	go db.WriteToDBPost(p)
}

