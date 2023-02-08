package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int32    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    Relations
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func InfArtist() []Artists {
	tmp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer tmp.Body.Close()
	body, err := ioutil.ReadAll(tmp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var artist []Artists
	json.Unmarshal(body, &artist)
	return artist
}

func NameArtist(s int) Artists {
	tmp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer tmp.Body.Close()
	body, err := ioutil.ReadAll(tmp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var artist []Artists
	json.Unmarshal(body, &artist)
	return artist[s-1]
}

func LocationDateArtist(s int) Relations {
	tmp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + strconv.Itoa(s))
	if err != nil {
		log.Fatal(err)
	}
	defer tmp.Body.Close()
	body, err := ioutil.ReadAll(tmp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var loc Relations
	json.Unmarshal(body, &loc)
	return loc
}
