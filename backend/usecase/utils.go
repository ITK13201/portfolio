package usecase

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
)

func newCsvReader(r io.Reader) *csv.Reader {
	br := bufio.NewReader(r)
	bs, err := br.Peek(3)
	if err != nil {
		return csv.NewReader(br)
	}
	if bs[0] == 0xEF && bs[1] == 0xBB && bs[2] == 0xBF {
		br.Discard(3)
	}
	return csv.NewReader(br)
}

func LoadCsv(filePath string) ([]map[string]string, error) {
	var data []map[string]string

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := newCsvReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	headers := rows[0]
	rows = rows[1:]
	for i := 0; i < len(rows); i++ {
		entity := make(map[string]string)
		for j := 0; j < len(rows[0]); j++ {
			entity[headers[j]] = rows[i][j]
		}
		data = append(data, entity)
	}
	return data, nil
}

func LoadJson(filePath string) ([]map[string]string, error) {
	var data []map[string]string

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	d := json.NewDecoder(file)
	err = d.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func must(func()) {}
