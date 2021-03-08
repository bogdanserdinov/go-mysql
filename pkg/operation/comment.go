package operation

import (
	"awesomeProject/nix/entity"
	"awesomeProject/nix/pkg/db"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	c := []entity.Comment{}

	err = json.NewDecoder(response.Body).Decode(&c)
	if err != nil {
		log.Fatal("failed in decoding json to comment : ", err)
	}

	if err = response.Body.Close(); err != nil {
		log.Fatal("failed to close response.Body : ", err)
	}

	var mutex1 = &sync.Mutex{}

	dsn := "root:blablabla29032002@tcp(127.0.0.1:3306)/public?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed in opening db : ",err)
	}

	for _,value := range c{
		go db.WriteToDBComment(value,gormDB,mutex1)
	}
}
