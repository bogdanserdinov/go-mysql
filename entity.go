package entity

//Post represent post tab on json
type Post struct{
	UserID int		`json:"userId"`
	ID int			`json:"id"`
	Title string	`json:"title"`
	Body string		`json:"body"`
}

//Comment represent comment tab on json
type Comment struct{
	PostID int		`json:"postId"`
	ID int			`json:"id"`
	Name string		`json:"name"`
	Email string	`json:"email"`
	Body string		`json:"body"`
}
