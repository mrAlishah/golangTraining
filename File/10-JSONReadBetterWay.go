package main

import (
	"encoding/json"
	"os"
	"testing"
)

// https://medium.com/kanoteknologi/better-way-to-read-and-write-json-file-in-golang-9d575b7254f2
//*The unmarshal way
/*
This is not efficient since the we need to unmarshal and hold the whole data in memory then doing
some iteration to filter the data. It might be ok with small JSON file. But will be a problem if we have big JSON file.
*/
func readJSON(fileName string, filter func(map[string]interface{}) bool) []map[string]interface{} {
	datas := []map[string]interface{}{}

	file, _ := os.ReadFile(fileName)
	json.Unmarshal(file, &datas)

	filteredData := []map[string]interface{}{}

	for _, data := range datas {
		// Do some filtering
		if filter(data) {
			filteredData = append(filteredData, data)
		}
	}

	return filteredData
}

//*The JSON Decoder way = The best way
/*
Better way to read JSON file is using json.Decoder. Because instead of unmarshal the whole content of a file
the decoder will decode one line/record at a time while we doing filtering in data. This is much more efficient
and less burden in memory.
*/
func readJSONToken(fileName string, filter func(map[string]interface{}) bool) []map[string]interface{} {
	file, _ := os.Open(fileName)
	defer file.Close()

	decoder := json.NewDecoder(file)

	filteredData := []map[string]interface{}{}

	// Read the array open bracket
	decoder.Token()

	data := map[string]interface{}{}
	for decoder.More() {
		decoder.Decode(&data)

		if filter(data) {
			filteredData = append(filteredData, data)
		}
	}

	return filteredData
}

func BenchmarkMovieRead(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readJSON("movie.json", func(data map[string]interface{}) bool {
			return data["year"].(float64) >= 2010
		})
	}
}

func BenchmarkMovieReadToken(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readJSONToken("movie.json", func(data map[string]interface{}) bool {
			return data["year"].(float64) >= 2010
		})
	}
}

func BenchmarkQRead(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readJSON("question.json", func(data map[string]interface{}) bool {
			return data["show_number"].(string) == "4680"
		})
	}
}

func BenchmarkQReadToken(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readJSONToken("question.json", func(data map[string]interface{}) bool {
			return data["show_number"].(string) == "4680"
		})
	}
}
