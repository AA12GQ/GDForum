package v1

import (
	"GDForum/app/requests"
	"GDForum/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

type Check404Controller struct {
	BaseAPIController
}

func (ctrl Check404Controller) Check404(c *gin.Context){
	 body := make([]string, 10,10)
	str := make([]string, 0,10)
	check := make([]requests.ResponseCheck,0,50)
	var request  *requests.Check404Input
	if err := c.ShouldBindJSON(&request); err!= nil {
		return
	}
	 resCheck := requests.ResponseCheck{}
	for key,value := range request.U{
		url := value
		body[key] = fetch(url)
		for _,v := range request.K{
			value := v
			rune := strings.Contains(body[key],value)
			if rune {
				str = append(str,value)
			}
		}
		resCheck.U = []string{url}
		resCheck.K = str
		check = append(check,resCheck)
	}

	response.JSON(c,gin.H{
		"data" : check,
	})

}

func fetch (url string) string {
	//fmt.Println("Fetch Url", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err:", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http status code:", resp.StatusCode)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error", err)
		return ""
	}
	return string(body)
}