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
	allGenerations := db.getAllJSON()
	newGeneration := Generation{UnsortedNumbers: unsorted, SortedNumbers: sorted, Time: time, CreatedAt: created}
	if length := len(allGenerations.Generations); length < SIZE {
		allGenerations.Generations = append(allGenerations.Generations, newGeneration)
		rawDataOut, _ := json.MarshalIndent(&allGenerations, "", "  ")
		_ = ioutil.WriteFile(db.path, rawDataOut, 0)
	} else {
		copy(allGenerations.Generations[0:], allGenerations.Generations[1:])
		allGenerations.Generations[length-1] = Generation{}
		allGenerations.Generations = allGenerations.Generations[:length-1]
		allGenerations.Generations = append(allGenerations.Generations, newGeneration)
		rawDataOut, _ := json.MarshalIndent(&allGenerations, "", "  ")
		_ = ioutil.WriteFile(db.path, rawDataOut, 0)
	}
}

// Возвращает срез всех JSON файлов в db.json
func (db *Driver) getAllJSON() JSON {
	rawDataIn, _ := ioutil.ReadFile(db.path)
	var allGenerations JSON
	_ = json.Unmarshal(rawDataIn, &allGenerations)
	return allGenerations
}

// Возвращает JSON, готовый для передачи на фронт
// func (db *Driver) GetAllMarshaledJSON() []byte {
// 	generations := db.getAllJSON().Generations
// 	rawDataOut, _ := json.Marshal(&generations)
// 	return rawDataOut
// }
