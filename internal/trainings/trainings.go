package trainings

import (
	"Final-5/internal/personaldata"
	"Final-5/internal/spentenergy"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("неверный формат данных: должно быть три параметра (шаги, тип тренировки, продолжительность)")
	}
	stepsStr := strings.TrimSpace(parts[0])
	if stepsStr == "" {
		return errors.New("неверное количество шагов: не может быть пустым")
	}
	steps, err := strconv.Atoi(stepsStr)
	if err != nil || steps <= 0 {
		return errors.New("неверное количество шагов: должно быть числовым и больше 0")
	}

	t.TrainingType = strings.TrimSpace(parts[1])
	if t.TrainingType == "" {
		return errors.New("неверный тип тренировки: не может быть пустым")
	}

	durationStr := strings.TrimSpace(parts[2])
	if durationStr == "" {
		return errors.New("неверная продолжительность: не может быть пустой")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("неверный формат продолжительности: %w", err)
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше 0")
	}
	t.Steps = steps
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.Personal.Weight <= 0 {
		return "", errors.New("вес должен быть больше 0")
	}
	if t.Personal.Height <= 0 {
		return "", errors.New("рост должен быть больше 0")
	}
	if t.Steps <= 0 {
		return "", errors.New("количество шагов должно быть больше 0")
	}

	if t.Duration <= 0 {
		return "", errors.New("продолжительность должна быть больше 0")
	}

	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	var calories float64
	var err error
	if t.TrainingType == "Бег" {
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	} else if t.TrainingType == "Ходьба" {
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	} else {
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}
	durationInHours := t.Duration.Seconds() / 3600
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, durationInHours, distance, meanSpeed, calories), nil
}
