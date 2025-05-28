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

func getCategoriesData() ([]dataTypes.CategoryData, error) {
	fmt.Println("========= Getting Category Data =========")
	Uri := "https://tcgcsv.com/tcgplayer/categories"
	buffer, err := GetCSVData(Uri)
	if err != nil {
		fmt.Println("GetCategory CSVData Error:", err)
		return nil, err
	}

	var csvData dataTypes.CategoriesDataResponse
	err = json.Unmarshal(buffer, &csvData)
	if err != nil {
		fmt.Println("GetCategory Unmarshal Error:", err)
		return nil, err
	}

	fmt.Printf("========= Found %d Category Data =========\n", csvData.TotalItems)

	firebase.UpdateCategoriesDataToArray(&csvData.Results, "Categories", "Games")

	fmt.Println("========= Uploaded Category Data to Firebase =========")

	return csvData.Results, nil
}

func GetGroupsData(categoryId string) ([]dataTypes.GroupData, error) {
	fmt.Printf("========= Getting Groups Data for %s =========\n", categoryId)

	Uri := "https://tcgcsv.com/tcgplayer/" + categoryId + "/groups"
	buffer, err := GetCSVData(Uri)
	if err != nil {
		fmt.Println("GetGroups CSVData Error:", err)
		return nil, err
	}

	var csvData dataTypes.GroupsDataResponse
	err = json.Unmarshal(buffer, &csvData)
	if err != nil {
		fmt.Println("GetGroups Unmarshal Error:", err)
		return nil, err
	}
	fmt.Printf("========= Found %d Groups Data =========\n", csvData.TotalItems)	
	
	firebase.UpdateGroupsDataToArray(&csvData.Results, "Groups", categoryId)

	fmt.Printf("========= Uploaded Groups Data to Firebase for %s =========\n", categoryId)

	return csvData.Results, nil
}

func GetProductsData(categoryId string, groupId string) ([]dataTypes.ProductData, error) {
	fmt.Printf("========= Getting Product Data for %s - %s =========\n", categoryId, groupId)

	Uri := "https://tcgcsv.com/tcgplayer/" + categoryId + "/" + groupId + "/products"
	buffer, err := GetCSVData(Uri)
	if err != nil {
		fmt.Println("GetProducts CSVData Error:", err)
		return nil, err
	}

	// Convert []byte to struct
	var csvData dataTypes.ProductsDataResponse
	err = json.Unmarshal(buffer, &csvData)
	if err != nil {
		fmt.Println("GetProducts Unmarshal Error:", err)
		return nil, err
	}

	fmt.Printf("========= Found %d Products Data =========\n", csvData.TotalItems)
	
	firebase.UpdateProductsDataToArray(&csvData.Results, categoryId, groupId)

	fmt.Printf("========= Uploaded Products Data to Firebase for %s - %s =========\n", categoryId, groupId)


	return csvData.Results, nil
}