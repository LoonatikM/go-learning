package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Movie struct {
	Title      string  `json:"title"`
	Year       int     `json:"year"`
	Rating     float64 `json:"rating"`
	secretNote string
}

func SaveMovies(filename string, movies []Movie) error {
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		return fmt.Errorf("save movies: %w", err)
	}

	err = os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("save movies: %w", err)
	}
	return nil
}

func LoadMovies(filename string) ([]Movie, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("load movies: %w", err)
	}
	var movies []Movie

	err = json.Unmarshal(data, &movies) // ⚠️ передаём УКАЗАТЕЛЬ
	if err != nil {
		return nil, fmt.Errorf("load movies: %w", err)
	}
	return movies, nil

}
func main() {
	movies := []Movie{
		{Title: "Inception", Year: 2010, Rating: 8.8},
		{Title: "Elsbet", Year: 2024, Rating: 7.9},
		{Title: "friends", Year: 1996, Rating: 9.0},
	}
	err := SaveMovies("movies.json", movies)
	if err != nil {
		fmt.Println(err)
		return
	}

	loadedMovies, err := LoadMovies("movies.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, m := range loadedMovies {
		fmt.Printf("%s (%d) - %.1f \n", m.Title, m.Year, m.Rating)

	}

}
