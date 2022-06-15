package fbrip

import (
	"encoding/json"
	"io/ioutil"
)

type RipPayload struct {
	User    jsonUser    `json:"user"`
	Actions jsonActions `json:"actions"`
}

type jsonUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type jsonActions struct {
	GetBasicInfo bool           `json:"getBasicInfo"`
	Reactions    []jsonReaction `json:"reactions"`
	Scrap        jsonScrap      `json:"scrap"`
}

type jsonReaction struct {
	ReactionId string `json:"reactionId"`
	PostUrl    string `json:"postUrl"`
}

type jsonScrap struct {
	Urls             []string `json:"urls"`
	OutputFolderPath string   `json:"outputFolderPath"`
	NamePrefix       string   `json:"namePrefix"`
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
