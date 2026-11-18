package main

import (
	"errors"
	"fmt"
	"os"
)

func SaveNote(filename, text string) error {
	err := os.WriteFile(filename, []byte(text), 0644)
	if err != nil {
		return fmt.Errorf("save note: %w", err)
	}
	return nil

}

func LoadNote(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("load note: %w", err)
	}
	return string(data), nil

}

func main() {
	// 1. Сохранить
	err := SaveNote("note.txt", "Сегодня учу файлы в Go")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 2. Прочитать и напечатать
	text, err := LoadNote("note.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(text)

	// 3. Прочитать несуществующий файл
	_, err = LoadNote("missing.txt")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("файл не найден")
	} else if err != nil {
		fmt.Println(err)
	}

}
