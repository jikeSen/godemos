package MyCurl

import (
    "bufio"
    "crypto/md5"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"
    "os"
    "path/filepath"
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
func Download(url string, filename string) {
    downloadDestFolder := "/Users/jksen/Downloads/test"

    res, err := http.Get(url)
    if err != nil {
        log.Printf("http.Get -> %v", err)
        return
    }
    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Printf("ioutil.ReadAll -> %s", err.Error())
        return
    }
    defer res.Body.Close()
    _ = os.MkdirAll(downloadDestFolder, 0777)
    //filename = time.Now().U
    filenames := CreateCaptcha() + ".jpg"
    if err = ioutil.WriteFile(downloadDestFolder+string(filepath.Separator)+filenames, data, 0777); err != nil {
        log.Println("Error Saving:", filenames, err)
    } else {
        log.Println("Saved:", filenames)
    }
}

func CreateCaptcha() string {
    return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
