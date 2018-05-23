package models

import "time"

type User struct {
	Id         int       `orm:"column(Id)"`
	Username   string    `orm:"column(Username)"`
	Password   string    `orm:"column(Password)"`
	Email      string    `orm:"column(Email)"`
	LoginCount int       `orm:"column(LoginCount)"`
	LastTime   time.Time `orm:"column(LastTime)"`
	LastIp     string    `orm:"column(LastIp)"`
	State      int8      `orm:"column(State)"`
	Created    time.Time `orm:"column(Created)"`
	Updated    time.Time `orm:"column(Updated)"`
}

func (m *User) TableName() string {
	return TableName("user")
}
