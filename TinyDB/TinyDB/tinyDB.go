package TinyDB

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

type TinyDB struct {
	data map[string]string
	mu   sync.Mutex
}

func NewTinyDB(filePath string) (*TinyDB, error) {
	db := &TinyDB{data: make(map[string]string)}

	// Load existing data from file
	if _, err := os.Stat(filePath); err == nil {
		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(file, &db.data)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

// Add a new entry
func (db *TinyDB) Put(key, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[key] = value
}

// Get an entry
func (db *TinyDB) Get(key string) (string, bool) {

	value, exists := db.data[key]
	return value, exists
}

// Delete an entry
func (db *TinyDB) Delete(key string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.data, key)
}

func (db *TinyDB) Save(filePath string) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	data, err := json.Marshal(db.data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

func (db *TinyDB) Search(value string) []string {
	var keys []string
	for k, v := range db.data {
		if v == value {
			keys = append(keys, k)
		}
	}
	return keys
}
