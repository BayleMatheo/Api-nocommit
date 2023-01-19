package Groupie

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func getData(url string, data interface{}) {
	rawData := getRawData(url)
	err := json.Unmarshal(rawData, &data)
	if err != nil {
		log.Panic("Problème dans la fonction getData lors du déclassement des données :", err)
	}
}

func getRawData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Panic("Problème dans la fonction getRawData lors de l'obtention de la réponse :", err)
		return nil
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panic("Problème dans la fonction getRawData lors de la lecture de la réponse :", err)
		return nil
	}
	return responseData
}

func getArtist() []rawArtist {
	artists := []rawArtist{}
	getData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	return artists
}

func getArtists(Id string) rawArtist {
	artist := rawArtist{}
	getData("https://groupietrackers.herokuapp.com/api/artists/"+Id, &artist)
	return artist
}
