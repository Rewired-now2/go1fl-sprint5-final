package personaldata

import "fmt"

// Структура для хранения персональных данных
type Personal struct {
	Name   string
	Weight float64 // Вес в килограммах
	Height float64 // Рост в метрах
}

// Метод для вывода персональных данных
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n", p.Name, p.Weight, p.Height)
}
