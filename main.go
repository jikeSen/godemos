package main

import (
	"fmt"
	"github.com/jikeSen/godemos/bfPost"
	"sync"
)

/*func init() {
	library.InitRedis()
}*/

func main() {

	/*listUser := make([]int, 10, 10)
	listUser[0] = 10042
	listUser[1] = 10043
	listUser[2] = 10044
	listUser[3] = 10044
	listUser[4] = 10045
	listUser[5] = 10046
	listUser[6] = 10047
	listUser[7] = 10048
	listUser[8] = 10049
	listUser[9] = 10050
	library.RedisConn.Select(9)
	err := library.RedisConn.SAdd("Demo:listUser", ).Err()

	if err != nil {
		panic(err)
	}*/

	sw := sync.WaitGroup{}
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

	fmt.Println("主程序结束")
}

func f(wg *sync.WaitGroup) {
	wg.Done()
}
