package main

import (
	"bufio"
	"domain/src/thesaurus"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalln("Failed when looking for " + word)
		}
		if len(syns) == 0 {
			log.Fatalln("Couldnt find any synonnyms for " + word)
		}

		for _, syn := range syns {
			fmt.Println(syn)
		}
	}

}
