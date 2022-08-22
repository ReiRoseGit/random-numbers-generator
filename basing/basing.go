package basing

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

// Структура, хранящая путь к файлу db.json
type Driver struct {
	path string
}

// Структура хранит срез сгенерированных JSON файлов
type JSON struct {
	Generations []Generation `json:"generations"`
}

// СТруктура конкретной генерации
type Generation struct {
	UnsortedNumbers []int         `json:"unsorted_numbers"`
	SortedNumbers   []int         `json:"sorted_numbers"`
	Time            time.Duration `json:"time"`
	CreatedAt       string        `json:"created_at"`
}

// Максимальное кол-во элементов
const SIZE = 10

func New(path string) *Driver {
	return &Driver{path: path}
}

// Читает данные из файла и добавляет новое значение, при необходимости удаляя первый элемент
func (db *Driver) AddJSON(unsorted, sorted []int, time time.Duration, created string) {
	generations := db.getAllJSON()
	newGeneration := Generation{UnsortedNumbers: unsorted, SortedNumbers: sorted, Time: time, CreatedAt: created}
	if length := len(generations); length < SIZE {
		generations = append(generations, newGeneration)
		rawDataOut, _ := json.MarshalIndent(&generations, "", "  ")
		_ = ioutil.WriteFile(db.path, rawDataOut, 0)
	} else {
		copy(generations[0:], generations[1:])
		generations[length-1] = Generation{}
		generations = generations[:length-1]
		generations = append(generations, newGeneration)
		rawDataOut, _ := json.MarshalIndent(&generations, "", "  ")
		_ = ioutil.WriteFile(db.path, rawDataOut, 0)
	}
}

// Возвращает срез всех JSON файлов в db.json
func (db *Driver) getAllJSON() []Generation {
	rawDataIn, _ := ioutil.ReadFile(db.path)
	var allGenerations JSON
	_ = json.Unmarshal(rawDataIn, &allGenerations)
	return allGenerations.Generations
}

// Возвращает JSON, готовый для передачи на фронт
func (db *Driver) GetAllMarshaledJSON() []byte {
	generations := db.getAllJSON()
	rawDataOut, _ := json.MarshalIndent(&generations, "", "  ")
	return rawDataOut
}
