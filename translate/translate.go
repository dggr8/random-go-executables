package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func GetPrivateKey() (private_key string) {
	jsonFile, err := os.Open("../../translate.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened the key.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	private_key = result["private_key"].(string)
	return
}

func createClientWithKey() {
	// import "cloud.google.com/go/translate"
	// import "google.golang.org/api/option"
	// import "golang.org/x/text/language"
	ctx := context.Background()

	var apiKey = GetPrivateKey()
	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Translate(ctx, []string{"Hello, world!"}, language.Russian, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", resp)
}

func main() {
	createClientWithKey()
}
