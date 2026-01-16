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
		return errors.New("ошибка в формате строки")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return err
	}

	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
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
