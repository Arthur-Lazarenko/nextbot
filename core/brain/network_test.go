package brain

import (
	"fmt"
	"math/rand"
)

func ExampleFeedForward() {

	// установка случайности в нулевое значение
	rand.Seed(0)

	// создание шаблона XOR для обучения сети
	patterns := [][][]float64{
		{{0, 0}, {0}},
		{{0, 1}, {1}},
		{{1, 0}, {1}},
		{{1, 1}, {0}},
	}

	// создание экземпляра передачи
	ff := &FeedForward{}

	// инициализация нейронной сети, структура сети будет содержать 2 входа, 2 скрытых узла и 1 выход
	ff.Init(2, 2, 1)

	/*
		Обучение сети с использованием шаблона XOR:

		Тренировка будет проходить 1000 итераций,
		скорость обучения установлена равной 0.6, а коэффициент импульса - 0.4,
		последний параметр отвечает за получение отчётов об ошибках при обучении.
	*/
	ff.Train(patterns, 1000, 0.6, 0.4, false)

	// тестирование обученной сети
	ff.Test(patterns)

	// ручное тестирование
	inputs := []float64{1, 1}
	fmt.Println(ff.Update(inputs))

	// Output:
	// [0 0] -> [0.057503945708445206]  :  [0]
	// [0 1] -> [0.9301006350712101]  :  [1]
	// [1 0] -> [0.9278099662272838]  :  [1]
	// [1 1] -> [0.09740879532462123]  :  [0]
	// [0.09740879532462123]
}
