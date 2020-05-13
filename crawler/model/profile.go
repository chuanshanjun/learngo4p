package model

import "encoding/json"

type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hukou      string
	Xinzuo     string
	House      string
	Car        string
}

func FromJsonObj(o interface{}) (Profile, error) {
	// 传进来一个interface{}
	var profile Profile
	// 将interface{}转成string
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	// 再将string专程profile
	err = json.Unmarshal(s, &profile)
	return profile, err
}
