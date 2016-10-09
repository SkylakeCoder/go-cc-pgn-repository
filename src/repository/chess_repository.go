package repository

import (
	"math/rand"
	"encoding/json"
	"os"
	"log"
	"io/ioutil"
)

var db map[string] []string = nil

func Init() {
	if db == nil {
		db = make(map[string] []string, 0)
	} else {
		log.Fatalln("db != nil...")
	}
}

func Record(key string, value string) {
	_, ok := db[key]
	if !ok {
		db[key] = []string {}
	}
	list := db[key]
	exist := false
	for _, v := range list {
		if v == value {
			exist = true
			break
		}
	}
	if !exist {
		db[key] = append(list, value)
	}
}

func Search(key string) string {
	v, ok := db[key]
	if !ok {
		return ""
	}
	random := rand.Intn(len(v))
	return v[random];
}

func Save() {
	bytes, err := json.Marshal(db)
	if err != nil {
		log.Fatalln("error when save...")
		os.Exit(1)
	}
	ioutil.WriteFile("db.json", bytes, 0644)
	log.Println("data saved success.")
}
