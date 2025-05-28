package main

import (
	"fmt"
	"time"
	os "os"
	"strconv"

	"github.com/potofgreedtcg/TCGProject-Data/dataParser"
)

func main() {

	start := time.Now()
    fmt.Println("========= Starting Data Updating =========")

	// categoriesCsvData, err := dataParser.GetCategoriesData()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }


	groupsCsvData, err := dataParser.GetGroupsData(strconv.Itoa(3))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	
	fmt.Printf("Game Data: %d\n Set Data: %d\n", len(groupsCsvData), len(groupsCsvData))

	// for _, category := range gameCsvData {
	// 	fmt.Printf("Game: %d\n", category.GroupId)
	// }


	// for _, set := range setCsvData {
	// 	fmt.Printf("Product Data: %d - %d \n", set.CategoryId, set.GroupId)
	// 	productCsvData, err := dataParser.GetProductsData(strconv.Itoa(set.CategoryId), strconv.Itoa(set.GroupId))
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		os.Exit(1)
	// 	}
	// 	fmt.Printf("Product Items: %d \n", len(productCsvData))
	// }
	
	
	fmt.Println("========= Finished Data Updating =========")
	duration := time.Since(start)
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	fmt.Printf("Execution time: %02d:%02d:%02d\n", hours, minutes, seconds)
}

