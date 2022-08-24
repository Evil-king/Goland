package usertype

type User struct {
	Id   int64  `orm:"id" json:"id"`
	Name string `orm:"name" json:"name"`
}
