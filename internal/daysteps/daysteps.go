package daysteps

import (
	"Final-5/internal/personaldata"
	"Final-5/internal/spentenergy"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) error {
	// Проверяем наличие пробела перед или после запятой
	if strings.Contains(datastring, " ,") || strings.Contains(datastring, ", ") {
		return errors.New("неверный формат данных: пробелы вокруг запятой")
	}

	// Теперь разделяем строку по запятой
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("неверный формат данных: должно быть два параметра (шаги и продолжительность)")
	}

	stepsStr := parts[0]
	if strings.Contains(stepsStr, " ") {
		return errors.New("неверное количество шагов: лишние пробелы")
	}

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return fmt.Errorf("неверное количество шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше 0")
	}
	ds.Steps = steps

	durationStr := parts[1]
	if strings.Contains(durationStr, " ") {
		return errors.New("неверная продолжительность: лишние пробелы")
	}
	if durationStr == "" {
		return errors.New("неверная продолжительность: продолжительность не может быть пустой")
	}
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("неверный формат продолжительности: %w", err)
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше 0")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", errors.New("количество шагов должно быть больше 0")
	}

	if ds.Duration <= 0 {
		return "", errors.New("продолжительность должна быть больше 0")
	}
	if ds.Personal.Weight <= 0 {
		return "", errors.New("вес должен быть больше 0")
	}
	if ds.Personal.Height <= 0 {
		return "", errors.New("рост должен быть больше 0")
	}
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories,
	)

	return info, nil
}

func parseSteps(stepsStr string) (int, error) {
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, fmt.Errorf("неверное количество шагов: %w", err)
	}
	if steps <= 0 {
		return 0, errors.New("шагов должно быть больше 0")
	}
	return steps, nil
}

func parseDuration(durationStr string) (time.Duration, error) {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, fmt.Errorf("неверный формат продолжительности: %w", err)
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше 0")
	}
	return duration, nil
}
