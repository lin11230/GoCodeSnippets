package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Card struct {
	Id           string   `json:"id,omitempty"`
	Desc         string   `json:"desc,omitempty"`
	IdBoard      string   `json:"idBoard,omitempty"`
	IdList       string   `json:"idList,omitempty"`
	Name         string   `json:"name,omitempty"`
	IdChecklists []string `json:"idCheckLists,omitempty"`
	IdMembers    []string `json:"idMembers,omitempty"`
}
type Member struct {
	Id       string `json:"id,omitempty"`
	FullName string `json:"fullName,omitempty"`
}

type CheckItem struct {
	Name  string `json:"name,omitempty"`
	State string `json:"state,omitempty"`
}
type Checklist struct {
	Name       string      `json:"name,omitempty"`
	CheckItems []CheckItem `json:"checkItems,omitempty"`
}

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	//Get members by board
	resMembers, err := client.Get("https://api.trello.com/1/boards/561f57695757ed7e676fbea2/members/?key=1d2c92ba1627a81e1e52651bf3251895&token=5eca8ed32c3bc14d50b02e2bfc80d52d152d999c83ddbcc99b4c8e370093f8fc")
	if err != nil {
		log.Fatal(err)
	}

	members, err := ioutil.ReadAll(resMembers.Body)
	resMembers.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var rowMembers []Member
	err = json.Unmarshal(members, &rowMembers)

	fmt.Printf("%s", rowMembers)
	fmt.Println("")
	fmt.Println("")
	for _, member := range rowMembers {
		fmt.Println("")
		fmt.Println("member id: ", member.Id)
		fmt.Println("member FullName: ", member.FullName)
		//Get cards by member's id
		resCards, err := client.Get("https://api.trello.com/1/boards/561f57695757ed7e676fbea2/members/" + member.Id + "/cards?key=1d2c92ba1627a81e1e52651bf3251895&token=5eca8ed32c3bc14d50b02e2bfc80d52d152d999c83ddbcc99b4c8e370093f8fc&cards=open")
		if err != nil {
			log.Fatal(err)
		}

		cards, err := ioutil.ReadAll(resCards.Body)
		resCards.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		var rowCards []Card
		err = json.Unmarshal(cards, &rowCards)

		//fmt.Printf("%s", rowCards)
		fmt.Println("")
		fmt.Println("Cards")
		for _, card := range rowCards {
			if card.IdList == "561f582446067acced353814" {
				fmt.Println("card name: ", card.Name)
				fmt.Println("card checklists id: ", card.IdChecklists)
				if len(card.IdChecklists) > 0 {
					for _, id := range card.IdChecklists {
						fmt.Println("idchecklists is: ", id)
						var query = fmt.Sprintf("https://api.trello.com/1/checklists/%s?key=1d2c92ba1627a81e1e52651bf3251895&token=5eca8ed32c3bc14d50b02e2bfc80d52d152d999c83ddbcc99b4c8e370093f8fc", id)
						resChecklists, err := client.Get(query)
						if err != nil {
							log.Fatal(err)
						}

						checklists, err := ioutil.ReadAll(resChecklists.Body)
						resChecklists.Body.Close()
						if err != nil {
							log.Fatal(err)
						}
						var rowChecklists Checklist
						rowChecklists = Checklist{}
						err = json.Unmarshal(checklists, &rowChecklists)
						if err != nil {
							log.Fatal(err)
						}

						fmt.Println("====== CheckItems ======")
						for _, item := range rowChecklists.CheckItems {
							fmt.Println("Check Item: ", item.Name)
							fmt.Println("Item state: ", item.State)

						}

					}

				}
				fmt.Println("")

			}
		}
		fmt.Println("")
		fmt.Println("")

	}

	//Get all cards by bord's id
	//res, err := client.Get("https://api.trello.com/1/boards/561f57695757ed7e676fbea2/cards?key=1d2c92ba1627a81e1e52651bf3251895&token=5eca8ed32c3bc14d50b02e2bfc80d52d152d999c83ddbcc99b4c8e370093f8fc")
	//if err != nil {
	//log.Fatal(err)
	//}

	//robots, err := ioutil.ReadAll(res.Body)
	//res.Body.Close()
	//if err != nil {
	//log.Fatal(err)
	//}
	//var row []Card
	//err = json.Unmarshal(robots, &row)

	//fmt.Printf("%s", row)
	//fmt.Println("")
	//fmt.Println("")
	//for _, card := range row {
	//fmt.Println(card.Name)
	//fmt.Println(card.IdBoard)
	//fmt.Println(card.IdMembers)
	//fmt.Println(card.IdChecklists)
	//}
}
