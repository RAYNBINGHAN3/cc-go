package src

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"sync"
)

//var ua *useragent

const AgentPath = "src/storage/"
const AgentJsonFile = "fake_agent.json"

type useragent struct {
	data []string
	lock sync.Once
}

//func init() {
//	ua = newUA()
//}

func newUA() *useragent {
	ua := new(useragent)
	err := ua.bufReadByFile()
	if err != nil {
		log.Fatal(err)
	}
	return ua
}

func (u *useragent) random() string {
	l := len(u.data)
	randomK := rand.Intn(l)
	return u.data[randomK]
}

func (u *useragent) bufReadByFile() error {
	docPath := AgentPath + AgentJsonFile
	content, err := ioutil.ReadFile(docPath)
	if err != nil {
		return err
	}

	var temp map[string][]string
	err = json.Unmarshal(content, &temp)
	if err != nil {
		return err
	}
	for _, v := range temp {
		u.data = append(u.data, v...)
	}

	return nil
}
