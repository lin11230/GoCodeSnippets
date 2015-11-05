package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jmcvetta/restclient"
)

type MapNode map[string][]string

type SSDB struct {
	Dns map[string]MapNode `json:"dns",ommitifempty`
}

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Get("https://gcssdb.nexusguard.com/ssdbapi.php?cmd=listjson")
	if err != nil {
		log.Fatal(err)
	}

	robots, err := ioutil.ReadAll(res.body)
	res.body.close()
	if err != nil {
		log.Fatal(err)
	}
	var row SSDB
	row = SSDB{}
	err = json.Unmarshal(robots, &row.Dns)

	fmt.Printf("%s", robots)
	fmt.Println("")
	fmt.Println("")
	for _, itemName := range row.Dns {
		for k, _ := range itemName {
			fmt.Println(k)
		}
	}
	fmt.Println("")
	fmt.Println("")

	type Foo struct {
		Email    string
		Password string
	}
	type Result struct {
		Token string
	}
	f := Foo{
		Email:    "admin@nxg.com",
		Password: "z5M9ceO6oA",
	}
	s := Result{}
	r := restclient.RequestResponse{
		Url:    "https://login.nxg.me/apiv1/auth?email=admin@nxg.com&password=z5M9ceO6oA",
		Method: restclient.POST,
		Data:   &f,
		Result: &s,
	}
	status, err := restclient.Do(&r)
	if err != nil {
		panic(err)
	}
	if status == 200 {
		println(s.Token)
	}
}
