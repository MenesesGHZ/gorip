package fbrip

import (
	"encoding/json"
	"io/ioutil"
)

type RipPayload struct {
	User User `json:"user"`
	Actions `json:"actions"`
}

type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Actions struct {
	GetBasicInfo bool `json:"getBasicInfo"`
	Reactions []Reaction `json:"reactions"`
	Scrap Scrap `json:"scrap"`
}

type Reaction struct {
	ReactionId string `json:"reactionId"`
	PostUrl string `json:"postUrl"`
}

type Scrap struct {
	Urls []string `json:"urls"`
	OutputFolderPath string `json:"outputFolderPath"`
	NamePrefix string `json:"namePrefix"`
}

func ReadRip(path string) (*RipPayload, error) {
	var payload RipPayload
	jsonByteSlice, err := ioutil.ReadFile(path)
	if err != nil {
		return &payload, err
	}
	json.Unmarshal(jsonByteSlice, &payload)
	return &payload, nil 
}