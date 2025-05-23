package dataParser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/potofgreedtcg/TCGProject-Data/firebase"
	"github.com/potofgreedtcg/TCGProject-Data/dataTypes"
)

func GetCSVData(uri string) ([]byte, error) {
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println("HTTP Get Error:", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read All Error:", err)
		return nil, err
	}

	return body, nil
}

func GetGameData() ([]dataTypes.GameData, error) {
	fmt.Println("========= Getting Game Data =========")
	Uri := "https://tcgcsv.com/tcgplayer/categories"
	buffer, err := GetCSVData(Uri)
	if err != nil {
		fmt.Println("GetGame CSVData Error:", err)
		return nil, err
	}

	var csvData dataTypes.GameDataResponse
	err = json.Unmarshal(buffer, &csvData)
	if err != nil {
		fmt.Println("GetGame Unmarshal Error:", err)
		return nil, err
	}

	fmt.Printf("========= Found %d Game Data =========\n", csvData.TotalItems)

	firebase.UpdateGameDataToArray(&csvData.Results, "Categories", "Games")

	fmt.Println("========= Uploaded Game Data to Firebase =========")

	return csvData.Results, nil
}

func GetSetData(gameId string) ([]dataTypes.SetData, error) {
	fmt.Printf("========= Getting Set Data for %s =========\n", gameId)

	Uri := "https://tcgcsv.com/tcgplayer/" + gameId + "/groups"
	buffer, err := GetCSVData(Uri)
	if err != nil {
		fmt.Println("GetSet CSVData Error:", err)
		return nil, err
	}

	var csvData dataTypes.SetDataResponse
	err = json.Unmarshal(buffer, &csvData)
	if err != nil {
		fmt.Println("GetSet Unmarshal Error:", err)
		return nil, err
	}
	fmt.Printf("========= Found %d Set Data =========\n", csvData.TotalItems)
	
	firebase.UpdateSetDataToArray(&csvData.Results, "Sets", gameId)

	fmt.Printf("========= Uploaded Set Data to Firebase for %s =========\n", gameId)

	return csvData.Results, nil
}

func GetProductData(categoryId string, groupId string) ([]dataTypes.ProductData, error) {
	fmt.Printf("========= Getting Product Data for %s - %s =========\n", categoryId, groupId)

	Uri := "https://tcgcsv.com/tcgplayer/" + categoryId + "/" + groupId + "/products"
	buffer, err := GetCSVData(Uri)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	// Convert []byte to struct
	var csvData dataTypes.ProductDataResponse
	err = json.Unmarshal(buffer, &csvData)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	fmt.Printf("========= Found %d Product Data =========\n", csvData.TotalItems)
	
	firebase.UpdateProductDataToArray(&csvData.Results, categoryId, groupId)

	fmt.Printf("========= Uploaded Product Data to Firebase for %s - %s =========\n", categoryId, groupId)


	return csvData.Results, nil
}