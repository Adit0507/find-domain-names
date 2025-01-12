package thesaurus

import (
	"encoding/json"
	"errors"
	"net/http"
)

// defines neccesary API key
// is provides the Synonym method thats responsible for doing work of accessing the endpoint
type BigHuge struct {
	APIKey string
}

type words struct {
	Syn []string `json:"syn"`
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string
	resp, err := http.Get("https://words.bighugelabs.com/api/2/" + b.APIKey + "/word" + "/json")
	if err != nil {
		return syns, errors.New("bighuge: Failed when looking for synonyms for " + term + "" + err.Error())
	}

	var data synonyms
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return syns, err
	}
	if data.Noun != nil {
		syns = append(syns, data.Noun.Syn...)
	}
	if data.Verb != nil {
		syns = append(syns, data.Verb.Syn...)
	}

	return syns, nil
}
