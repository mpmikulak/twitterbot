package main

import (
	"encoding/csv"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"log"
	"os"
)

const (
	consumerKey    = "uek7mBKUGMQRWdIlWiK8sM0gi"
	consumerSecret = "I762buScjVTzv3OknlrUZL6ag4oV7h944ZjDs9GFDoQrg4kXat"
	accessToken    = "794621011450806272-TZqLoOib2YiHyfIz6JPCEVyWwGzYK0H"
	tokenSecret    = "VwOSDh8eHPdQFsfEmyEzibyTkVist702K5CrinLYpz04w"
)

func main() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, tokenSecret)
	defer api.Close()

	urlData := make(map[string][]string)
	populateUrlData(urlData)	

	/*stream := api.PublicStreamFilter(urlData)

	for tweet := range stream.C {
		v := tweet.(anaconda.Tweet)
		fmt.Printf("%s: %s\n", v.User.Name, v.Text)
	}*/
}

func populateUrlData(m map[string][]string) {
	file, err := os.Open("searchTerms.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("yes")

	for _, record := range data {
		
	fmt.Println(record)
	}
	
}
