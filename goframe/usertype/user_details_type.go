package usertype

type UserDetails struct {
	Id          int64  `orm:"id" json:"id"`
	UserId      int64  `orm:"user_id" json:"user_id"`
	UserName    string `orm:"user_name" json:"user_name"`
	UserPhone   string `orm:"user_phone" json:"user_phone"`
	UserAddress string `orm:"user_address" json:"user_address"`
}
