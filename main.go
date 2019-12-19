package main

import (
    "fmt"
    "github.com/jikeSen/godemos/library"
    "github.com/jikeSen/godemos/model"
    jsoniter "github.com/json-iterator/go"
    "sync"
    "time"
    "unsafe"
)

func init() {
    library.MysqlConnect()
}

func main() {
    UpdateUser()
    /*sw := sync.WaitGroup{}
      sw.Add(10)
      uid := 10042
      for i := 0; i < 10; i++ {
          uid = uid + 1
          go func(uid int) {
              bfPost.GetChook(uid)
              f(&sw)
          }(uid)

      }

      sw.Wait()

      fmt.Println("主程序结束")*/
    //library.InitRabbit()

}

func f(wg *sync.WaitGroup) {
    wg.Done()
}

func Adder() {
    u := model.User{}
    u.Cellphone = "15565704122"
    u.Nickname = "1556570412332"
    u.IsAuth = 0
    u.Avatar = "/default/avatar.jpg"
    u.ParentUid = 10010
    u.Status = 1
    u.ShareCodeUrl = "/user/1316e1068c2a509e3732de03e7b7561f.png"
    init_time := time.Now().Unix()
    int_num := *(*int)(unsafe.Pointer(&init_time))
    u.InitTime = int_num

    u.AddUser()
}

func GetUserList() {
    u := model.User{}

    r := u.GetUserList()

    var json = jsoniter.ConfigCompatibleWithStandardLibrary

    resByte := make([]byte, 0, 2014)
    resByte, err := json.Marshal(&r)

    if err != nil {
        panic(err)
    }

    fmt.Println(string(resByte))

}

func UpdateUser() {
    u := model.User{}
    u.Id = 10045
    u.Nickname = "jksen"
    u.UpdateUserInfo()
}
