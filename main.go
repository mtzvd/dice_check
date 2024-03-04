package main

import (
	"fmt"
)

type config struct {
	diceNumb    int
	sides       int
	threshold   int
	successNumb int
	modifier    int
}

// Abs возвращает модуль числа
func Abs(number int) int {
	if number < 0 {
		return number * -1
	}

	return number
}

// isSuccess проверяет является ли комбинация успешной
func isSuccess(threshold, successNumb int, dices []int) bool {
	aboveThreshold := 0
	for _, dice := range dices {
		if dice >= threshold {
			aboveThreshold++
		}
	}

	return aboveThreshold >= successNumb
}

// generateModifiers возвращает слайс со всеми комбинациями распределения модификатора по коробкам
func generateModifiers(modifier, diceNumb int) [][]int {
	absMod := Abs(modifier)
	if diceNumb == 1 {
		// Если осталась одна коробка, возвращаем все оставшиеся единицы в ней
		return [][]int{{absMod}}
	}

	// Инициализация начального состояния для первой коробки
	initialCombinations := make([][]int, absMod+1)
	for i := range initialCombinations {
		initialCombinations[i] = []int{i}
	}

	// Построение комбинаций для каждой коробки
	for box := 2; box <= diceNumb; box++ {
		newCombinations := [][]int{}
		for _, combo := range initialCombinations {
			sum := 0
			for _, value := range combo {
				sum += value
			}
			for i := 0; i <= absMod-sum; i++ {
				if box == diceNumb && sum+i != absMod {
					continue // Для последней коробки убедимся, что используем все единицы
				}

				// Добавляем новое значение к существующим комбинациям
				newCombo := append([]int(nil), combo...)
				newCombo = append(newCombo, i)
				newCombinations = append(newCombinations, newCombo)
			}
		}
		initialCombinations = newCombinations
	}

	// Применяем модификацию для отрицательных модификаторов
	if modifier < 0 {
		for i, combo := range initialCombinations {
			for j, value := range combo {
				initialCombinations[i][j] = -value
			}
		}
	}

	return initialCombinations
}

// generateCombinations создает все возможные комбинации
func generateCombinations(diceNumb, sides int, combinationsChan chan<- []int) {
	defer close(combinationsChan) // Закрываем канал по завершению функции

	// Инициализируем начальную комбинацию с пустым списком
	result := [][]int{{}}

	// Для каждой кости в игре
	for d := 0; d < diceNumb; d++ {
		newResult := [][]int{}

		// Проходим по всем текущим комбинациям
		for _, combo := range result {

			// Добавляем к каждой комбинации все возможные значения текущей кости
			for s := 1; s <= sides; s++ {
				newCombo := append([]int(nil), combo...)
				newCombo = append(newCombo, s)

				// Добавляем новую комбинацию во временный результат для следующего уровня
				newResult = append(newResult, newCombo)
			}
		}

		// Обновляем основной результат новыми комбинациями
		result = newResult
	}

	// Выводим комбинации
	for _, combo := range result {
		fmt.Printf("Сгенерировано: %v\n", combo)
	}

	// После генерации всех комбинаций, отправляем их в канал
	fmt.Printf("Начинаю проверку\n")
	for _, combo := range result {
		combinationsChan <- combo
	}
}

// summSlices поэлементно складывает слайсы, длина результирующего слайса равна длине первого слайса
func summSlices(a, b []int) []int {
	result := make([]int, len(a))
	for i := range a {
		result[i] = a[i] + b[i]
	}

	return result
}

// processCombination возвращает может ли комбинация быть успешной с применением модификатора
func processCombination(combination []int, modifierTable [][]int, threshold, successNumb int) bool {

	// Проходим по всем комбинациям модификаторов
	for _, modCombo := range modifierTable {

		// Складываем текущую комбинацию с модификатором
		modifiedCombo := summSlices(combination, modCombo)

		// Проверяем, является ли модифицированная комбинация успешной
		if isSuccess(threshold, successNumb, modifiedCombo) {
			return true // Возвращаем true, если комбинация успешна
		}
	}

	return false // Возвращаем false, если ни одна комбинация не привела к успеху
}

func main() {

	//Задаем параметры
	cfg := config{
		threshold:   4, //число, на кости считается успех, если выпало не меньше этого числа.
		successNumb: 3, //число костей, на которых должен выпасть успех (с учетом модификатора), чтобы результат броска считался успешным.
		diceNumb:    3, //число костей в броске.
		sides:       6, //число граней каждой кости.
		modifier:    1, //модификатор броска (может быть отрицательным).
	}

	//Обнуляем счётчики
	allCases := 0
	succesCases := 0

	//Генерируем таблицу модификатора
	fmt.Printf("Генерирую таблицу всех вариантов распределения модификатора %v на количество костей=%v  \n", cfg.modifier, cfg.diceNumb)
	modifierTable := generateModifiers(cfg.modifier, cfg.diceNumb)
	fmt.Printf("Таблица распределения модификатора %v на количество костей = %v  \n %v\n\n", cfg.modifier, cfg.diceNumb, modifierTable)

	//Создаем канал для обработки комбинаций
	combinationsChan := make(chan []int)

	//Запускаем генератор
	fmt.Printf("Генерирую все комбинации %vd%v\n", cfg.diceNumb, cfg.sides)
	go generateCombinations(cfg.diceNumb, cfg.sides, combinationsChan)

	// Читаем комбинаций из канала и обрабатываем
	for combo := range combinationsChan {
		result := processCombination(combo, modifierTable, cfg.threshold, cfg.successNumb)
		fmt.Println(combo, ": ", result)
		if result {
			succesCases++
		}
		allCases++
	}

	//Выводим результат
	fmt.Printf("Успешных бросков %v из %v\n", succesCases, allCases)
	fmt.Printf("Вероятность успеха при \n%+v \n=%v", cfg, float32(succesCases)/float32(allCases))
}
