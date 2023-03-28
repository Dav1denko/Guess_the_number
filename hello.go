package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	second := time.Now().Unix()
	rand.Seed(second)
	target := rand.Intn(100) + 1
	fmt.Println("Я выбираю число от 1 до 100")
	fmt.Println("Число выбрано")

	reader := bufio.NewReader(os.Stdin)

	success := false
	for guesses := 0; guesses < 10; guesses++ {

		fmt.Println("у вас осталось  ", 10-guesses, "попыток")
		fmt.Print("Я загадал число: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess > target {
			fmt.Println("Ваше значение больше чем загоданое")
		} else if guess < target {
			fmt.Println("Ваше значение меньше чем загоданое")

		} else {
			success = true
			fmt.Println("Поздравляю")
			break
		}
	}
	if !success {
		fmt.Println("Извините, ваши попытки закончились. Это число было ", target)
	}
}
