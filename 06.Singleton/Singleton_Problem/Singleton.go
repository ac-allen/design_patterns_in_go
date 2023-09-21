package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

// ! Breaks dependency inversion principle
// ! Unit test is going to integration test
// ! as reading a live database

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

func GetSingletonDatabase() *singletonDatabase {
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

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city)
	}
	return result
}

func main() {
	cities := []string{"Seoul", "Mexico City"}
	tp := GetTotalPopulation(cities)
	ok := tp == (17500000 + 17400000)
	fmt.Println(ok)
}
