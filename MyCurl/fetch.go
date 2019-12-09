package MyCurl

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	Host = "http://api.park.com"
)

// post 请求
func HttpPost(HttpUrl string, PostForm interface{}) string {

	client := &http.Client{}

	urlStr := Host + HttpUrl

	jsons, _ := json.Marshal(PostForm)
	result := string(jsons)
	jsoninfo := strings.NewReader(result)

	request, err := http.NewRequest("POST", urlStr, jsoninfo)

	request.Header.Set("admin", createAdmin())
	request.Header.Set("Content-Type", "application/json:JSON")

	ress, err := client.Do(request)

	if err != nil {
		log.Printf("调用接口异常%v", err.Error())
	}

	defer ress.Body.Close()

	responseByte := make([]byte, 1024)
	reader := bufio.NewReader(ress.Body)
	n, err := reader.Read(responseByte)

	if err != nil {
		log.Printf("读取数据错误%v", err.Error())
	}

	body := responseByte[:n]
	fmt.Println(string(body))
	return string(body)
}

func createAdmin() string {
	y := time.Now().Year()
	m := int(time.Now().Month())
	d := time.Now().Day()

	dd := "01"
	if d < 10 {
		dd = "0" + strconv.Itoa(d)
	} else {
		dd = strconv.Itoa(d)
	}

	secStr := strconv.Itoa(y) + strconv.Itoa(m) + dd + "admin"
	return md5Str(secStr)
}

func md5Str(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
