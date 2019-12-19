package model

import "github.com/jikeSen/godemos/library"

type User struct {
    Id           int `gorm:"primary_key:id"`
    Cellphone    string
    Nickname     string
    Avatar       string
    Token        string
    Status       int
    ShareCodeUrl string
    IsAuth       int
    ParentUid    int
    InitTime     int
}

//修改默认表名
func (User) TableName() string {
    return "gp_user"
}

func (user *User) AddUser() {
    library.Db.Table("gp_user").Create(user)
}

func (user *User) UpdateUserInfo() {
    library.Db.Model(&user).Update(user)
}

// 查询列表demo
func (user *User) GetUserList() (u []User) {
    // in 查询
    /*library.Db.Table("qy_user").Where("id in (?)",[]int{14607944,14607934}).Find(&u)*/

    // where 查询
    library.Db.Table("gp_user").Where("id > ?", 10044).Find(&u)

    return u
}
