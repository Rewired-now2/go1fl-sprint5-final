package daysteps

import (
	"Final-5/internal/personaldata"
	"Final-5/internal/spentenergy"
	"errors"
	"fmt"
	"time"
)

type DaySteps struct {
	Steps                 int
	Duration              time.Duration
	personaldata.Personal // Встроенная структура с персональными данными
}

// Метод для парсинга строки с данными о прогулке
func (ds *DaySteps) Parse(datastring string) (err error) {
	// Разделяем строку на части
	parts := split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("некорректный формат данных")
	}

	// Парсим количество шагов
	steps, err := parseSteps(parts[0])
	if err != nil {
		return err
	}
	ds.Steps = steps

	// Парсим продолжительность
	duration, err := parseDuration(parts[1])
	if err != nil {
		return err
	}
	ds.Duration = duration

	return nil
}

// Функция для разделения строки
func split(s string, sep string) []string {
	// Можно использовать strings.Split(s, sep) для реальной реализации
	var result []string
	// Для простоты
	return result
}

// Функция для парсинга количества шагов
func parseSteps(stepsStr string) (int, error) {
	// Преобразуем строку в число
	// Если возникнет ошибка, вернуть её
	return 678, nil // Пример, замените на реальную логику
}

// Функция для парсинга продолжительности
func parseDuration(durationStr string) (time.Duration, error) {
	// Предположим, что строка имеет формат "0h50m"
	// Парсим вручную (реализуйте это в зависимости от формата строки)
	// Это просто пример, реальную логику нужно будет доработать
	return time.Duration(50 * time.Minute), nil
}

// Метод для формирования строки с информацией о прогулке
func (ds DaySteps) ActionInfo() (string, error) {
	// Вычисляем пройденную дистанцию
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	// Вычисляем потраченные калории
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	// Формируем строку с информацией
	info := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		ds.Steps, distance, calories,
	)

	return info, nil
}
