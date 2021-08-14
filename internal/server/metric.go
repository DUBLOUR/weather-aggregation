package server

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Metric struct{
	dbFile string
}

func (m Metric) load() (map[string]int, error) {
	result := make(map[string]int)
	if _, err := os.Stat(m.dbFile); os.IsNotExist(err) {
		return result, nil
	}
	data, err := ioutil.ReadFile(m.dbFile)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

func (m Metric) save(data map[string]int) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(m.dbFile, []byte(jsonData), 0644)
}


func (m Metric) Inc(key string) error {
	a, err := m.load()
	if err != nil {
		return err
	}
	a[key] += 1
	return m.save(a)
}

func (m Metric) Get(key string) (int, error) {
	a, err := m.load()
	return a[key], err
}
