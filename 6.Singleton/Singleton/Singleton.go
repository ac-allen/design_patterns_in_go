package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// ? sync.Once or package level init -- thread safety
// ? laziness, only when someone need it
// ? sync is following laziness while init isn't

func readData(path string) (map[string]int, error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	file, err := os.Open(exPath + path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

var once sync.Once
var instance *singletonDatabase

func GetSignletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps, e := readData("/capitals.txt")
		db := singletonDatabase{}
		if e == nil {
			db.capitals = caps
		} else {
			fmt.Printf("error : %v", e)
		}
		instance = &db
	})
	return instance
}

func main() {
	db := GetSignletonDatabase()
	pop := db.GetPopulation("Seoul")
	fmt.Println("Population of Seoul = ", pop)
}
