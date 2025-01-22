package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

type JsonRecord struct {
	Delete struct {
		Index string `json:"_index"`
		Id    string `json:"_id"`
	} `json:"delete"`
}

const ndjsonContents = `
{ "delete": { "_index": "my-index", "_id": "1" } }
{ "delete": { "_index": "my-index", "_id": "2" } }
{ "delete": { "_index": "my-index", "_id": "3" } }
`

func main() {
	f := strings.NewReader(ndjsonContents)

	var records []JsonRecord
	d := json.NewDecoder(f)
	for {
		var record JsonRecord
		err := d.Decode(&record)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}
		records = append(records, record)
	}

	fmt.Printf("%+v\n", records)
}
