package entity

//Post represent post tab on json
type Post struct{
	UserID int		`json:"userId" gorm:"column:UserID"`
	ID int			`json:"id" gorm:"primary_key"`
	Title string	`json:"title" gorm:"type:text"`
	Body string		`json:"body" gorm:"type:text"`
}

//Comment represent comment tab on json
type Comment struct{
	PostID int		`json:"postId" gorm:"column:PostID"`
	ID int			`json:"id" gorm:"primary_key"`
	Name string		`json:"name" gorm:"type:text"`
	Email string	`json:"email" gorm:"type:varchar(60)"`
	Body string		`json:"body" gorm:"type:text"`
}