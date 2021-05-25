package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*这个结构会保存解析后的返回数据。
他们会形成有层级的XML，可以忽略一些无用的数据*/
type Status struct {
	Text string
}

type User struct {
	XMLName xml.Name
	Status  Status
}

func main() {
	// perform an HTTP request for the twitter status of user: Googland
	resp, _ := http.Get("http://twitter.com/users/Googland.xml")
	// initialize the structure of the XML response
	user := User{xml.Name{"", "user"}, Status{""}}
	// unmarshal the XML into our structures
	defer resp.Body.Close()
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		fmt.Printf("error: %s", err.Error())
	} else {
		fmt.Printf("%s ---", body)
		xml.Unmarshal(body, &user)
	}
	fmt.Printf("name: %s ", user.XMLName)
	fmt.Printf("status: %s", user.Status.Text)
}