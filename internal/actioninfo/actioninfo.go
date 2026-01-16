package actioninfo

import (
	"fmt"
	"log"
)

// DataParser интерфейс определяет методы для парсинга и получения информации
type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error // метод для парсинга строки
	ActionInfo() (string, error)   // метод для получения информации
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию

	// Перебираем все строки в dataset
	for i, data := range dataset {
		// Парсим строку
		err := dp.Parse(data)
		if err != nil {
			// Логируем ошибку и переходим к следующей строке
			log.Printf("Ошибка парсинга строки %d: %v", i+1, err)
			continue
		}

		// Получаем информацию об активности
		info, err := dp.ActionInfo()
		if err != nil {
			// Логируем ошибку
			log.Printf("Ошибка получения информации для строки %d: %v", i+1, err)
			continue
		}

		// Выводим разделитель и информацию
		fmt.Printf("=== Результат обработки ===\n%s\n", info)
	}
}
