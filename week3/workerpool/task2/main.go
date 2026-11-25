package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

var ErrInvalidEmail = errors.New("invalid email")

type SendResult struct {
	Email string
	Err   error
}

func sendEmail(email string) error {
	time.Sleep(300 * time.Millisecond)
	if !strings.Contains(email, "@") {
		return ErrInvalidEmail
	}
	return nil
}

func emailWorker(id int, jobs <-chan string, results chan<- SendResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for email := range jobs {
		err := sendEmail(email)
		results <- SendResult{Email: email, Err: err}
	}
}

func main() {
	emails := []string{"@good1-adress", "@good2-adress", "@good3-adress", "bad1-adress", "@good4-adress", "@good5-adress", "bad2-adress", "@good6-adress"}
	jobs := make(chan string, len(emails))
	results := make(chan SendResult, len(emails))
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go emailWorker(w, jobs, results, &wg)
	}

	// 2. кладём 9 задач и закрываем очередь
	for _, e := range emails {
		jobs <- e
	}
	close(jobs)

	// 3. когда все worker'ы закончат — закрываем results
	go func() {
		wg.Wait()
		close(results)
	}()

	// 4. собираем результаты

	ok, failed := 0, 0
	for r := range results {
		if r.Err != nil {
			fmt.Printf("ошибка: %s (%v)\n", r.Email, r.Err)
			failed++
		} else {
			fmt.Println("отправлено:", r.Email)
			ok++
		}
	}
	fmt.Printf("успешно: %d, ошибок: %d", ok, failed)
}
