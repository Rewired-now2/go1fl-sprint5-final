package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// Функция для расчёта дистанции в километрах
func Distance(steps int, height float64) float64 {
	// Расчёт длины шага
	stepLength := height * stepLengthCoefficient
	// Общее количество метров
	distanceInMeters := float64(steps) * stepLength
	// Переводим в километры
	distanceInKm := distanceInMeters / mInKm
	return distanceInKm
}

// Функция для расчёта средней скорости
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Проверка корректности входных данных
	if steps < 0 || duration <= 0 {
		return 0
	}
	// Расчёт дистанции
	distance := Distance(steps, height)
	// Переводим продолжительность в часы
	durationInHours := duration.Seconds() / 3600
	// Средняя скорость
	return distance / durationInHours
}

// Функция для расчёта калорий, потраченных при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверка входных данных
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные параметры")
	}
	// Расчёт средней скорости
	meanSpeed := MeanSpeed(steps, height, duration)
	// Переводим продолжительность в минуты
	durationInMinutes := duration.Seconds() / 60
	// Калории, потраченные на бег
	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories, nil
}

// Функция для расчёта калорий, потраченных при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверка входных данных
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные параметры")
	}
	// Расчёт средней скорости
	meanSpeed := MeanSpeed(steps, height, duration)
	// Переводим продолжительность в минуты
	durationInMinutes := duration.Seconds() / 60
	// Калории, потраченные на ходьбу
	calories := (weight * meanSpeed * durationInMinutes) / minInH
	// Применяем коэффициент для ходьбы
	calories *= walkingCaloriesCoefficient
	return calories, nil
}
