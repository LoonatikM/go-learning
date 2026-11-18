package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title      string  `json:"title"`
	Year       int     `json:"year"`
	Rating     float64 `json:"rating"`
	secretNote string
}

func main() {
	m := Movie{Title: "Elsbet", Year: 2020, Rating: 7.9, secretNote: "good"}

	data, err := json.Marshal(m) // возвращает ([]byte, error)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	inMov := []byte(`{"title": "Inception", "year": 2010, "rating": 8.8}`)
	var m2 Movie

	err = json.Unmarshal(inMov, &m2) // ⚠️ передаём УКАЗАТЕЛЬ
	if err != nil {
		fmt.Println("плохой JSON:", err)
		return
	}
	fmt.Printf("%s (%d) - %.1f \n", m2.Title, m2.Year, m2.Rating)
}
