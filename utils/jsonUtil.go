package utils

import (
	"encoding/json"
	"os"
	"zhanhuo/entity"
)

func ReadJons(fileName string) *entity.Account {
	fileName = "./" + fileName + ".json"
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	ret := &entity.Account{}
	if err = json.NewDecoder(f).Decode(ret); err != nil {
		panic(err)
	}
	return ret
}

func WriteJons(account *entity.Account) {
	fileName := "./" + account.Name + ".json"
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data, err := json.MarshalIndent(account, "", "  ")
	if err != nil {
		panic(err)
	}
	f.Write(data)
}

func ReadDevices() [][]string {
	f, err := os.Open("./devices.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var ret [][]string
	if err = json.NewDecoder(f).Decode(&ret); err != nil {
		panic(err)
	}
	return ret
}
