package core

import (
	"log"
	"regexp"
)

// FilterTheQuestion - функция, предназначенная для очистки входящего запроса от лишних символов
func FilterTheQuestion(question string) string {
	reg, error := regexp.Compile("[^ a-z A-Z а-я А-Я 0-9 , ! ? .]")
	if error != nil {
		log.Fatal(error)
	}
	question = reg.ReplaceAllString(question, "")

	return question
}
