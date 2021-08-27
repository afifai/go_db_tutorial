package main

import (
	"encoding/json"
)

const data = `[{"nama": "Afif A. Iskandar", "usia": 20, "lokasi": "Bandung"},
				{"nama": "Budi Nugroho", "usia": 26, "lokasi": "Bekasi"},
				{"nama": "Dadang Subur", "usia": 40, "lokasi": "Cimahi"},
				{"nama": "Rony Alfarisi", "usia": 15, "lokasi": "Bogor"},
				{"nama": "Zufri", "usia": 22, "lokasi": "Tangerang"}]`

func GetData() (people []Person) {
	err := json.Unmarshal([]byte(data), &people)
	if err != nil {
		panic(err.Error)
	}
	return people
}
