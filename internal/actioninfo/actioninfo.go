package actioninfo

import (
	"fmt"
)

// Интерфейс для парсинга и формирования информации
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Функция для обработки данных
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			fmt.Println("Ошибка при парсинге:", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Println("Ошибка при получении информации:", err)
			continue
		}
		fmt.Println(info)
	}
}
