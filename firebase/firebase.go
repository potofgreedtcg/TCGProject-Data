package firebase

import (
    "context"
    "log"
    "fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"unicode"
	os "os"
		
    firebase "firebase.google.com/go/v4"
    "google.golang.org/api/option"
    "cloud.google.com/go/firestore"
    dataTypes "github.com/potofgreedtcg/TCGProject-Data/dataTypes"
)

func InitializeFirebase() (*firebase.App, error) {

	// Create the configuration
	config := &firebase.Config{
		ProjectID: "potofgreedtcg-cc276", // Add your Firebase project ID here
	}

	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
	  return nil, fmt.Errorf("error initializing app: %v", err)
	}

	return app, nil
}

func GetFirestoreClient(app *firebase.App) (*firestore.Client, error) {
    client, err := app.Firestore(context.Background())
    if err != nil {
        log.Printf("Error getting Firestore client: %v\n", err)
        return nil, err
    }
    return client, nil
}

// Update Game Data to Array in Firestore document
func UpdateCategoriesDataToArray(data *[]dataTypes.CategoryData, Collection string, document string) {
    app, err := InitializeFirebase()
    if err != nil {
		log.Printf("Error setting document: %v\n", err)
    }

    client, err := GetFirestoreClient(app)
    if err != nil {
        log.Printf("Error getting Firestore client: %v\n", err)
    }

	defer client.Close()

    // Set the array in the document
    _, err = client.Collection(Collection).Doc(document).Set(context.Background(), map[string]interface{}{
        "results": data,
    })
    if err != nil {
        log.Println("Error setting array: %v", err)
    }
}

// Update Set Data to Array in Firestore document
func UpdateGroupsDataToArray(data *[]dataTypes.GroupData, Collection string, document string) {

    app, err := InitializeFirebase()
    if err != nil {
		log.Printf("Error setting document: %v\n", err)
    }

    client, err := GetFirestoreClient(app)
    if err != nil {
        log.Printf("Error getting Firestore client: %v\n", err)
    }

	defer client.Close()

    for _, group := range *data {
        cleaned := cleanString(group.Name)

		filePath := fmt.Sprintf("images/groups/")
		filename := fmt.Sprintf("%s%d.jpg", filePath, group.GroupId)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			os.MkdirAll(filePath, 0755)
		}

		groupUrl := fmt.Sprintf("https://tcgplayer-cdn.tcgplayer.com/set_icon/%s.png", cleaned)
		err := SaveImage(groupUrl, filename)
		if err != nil {
			log.Printf("Error downloading image for Group %d: %v", group.GroupId, err)
		}
		fmt.Printf("Downloaded image for Group %d\n", group.GroupId)
	}

    // Set the array in the document
    _, err = client.Collection(Collection).Doc(document).Set(context.Background(), map[string]interface{}{
        "results": data,
    })
    if err != nil {
        log.Println("error setting array: %v", err)
    }
	
}

func UpdateProductsDataToArray(data *[]dataTypes.ProductData, Collection string, document string) {
    app, err := InitializeFirebase()
    if err != nil {
		log.Printf("Error setting document: %v\n", err)
    }

    client, err := GetFirestoreClient(app)
    if err != nil {
        log.Printf("Error getting Firestore client: %v\n", err)
    }

	defer client.Close()

    for _, product := range *data {
		filePath := fmt.Sprintf("images/%s/%s/", Collection, document)
		filename := fmt.Sprintf("%s%d.jpg", filePath, product.ProductId)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			os.MkdirAll(filePath, 0755)
		}
		

		err := SaveImage(product.ImageUrl, filename)
		if err != nil {
			log.Printf("Error downloading image for product %d: %v", product.ProductId, err)
		}
		fmt.Printf("Downloaded image for product %d\n", product.ProductId)
	}

    // Set the array in the document
    _, err = client.Collection(Collection).Doc(document).Set(context.Background(), map[string]interface{}{
        "results": data,
    })
    if err != nil {
        log.Println("error setting array: %v", err)
    }
}

func SaveImage(url string, filename string) error {
    // Download the image
    imageBytes, err := DownloadImage(url)
    if err != nil {
        return fmt.Errorf("error downloading image: %v", err)
    }

    // Save to file
    err = ioutil.WriteFile(filename, imageBytes, 0644)
    if err != nil {
        return fmt.Errorf("error saving image: %v", err)
    }

    return nil
}

func DownloadImage(url string) ([]byte, error) {
    // Make HTTP GET request
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error making request: %v", err)
    }
    defer resp.Body.Close()

    // Check if the response status is OK
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("bad status code: %d", resp.StatusCode)
    }

    // Read the response body
    imageBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response: %v", err)
    }

    return imageBytes, nil
}

func cleanString(input string) string {
	var builder strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}