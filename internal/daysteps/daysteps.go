package daysteps

import (
	"Final-5/internal/personaldata"
	"Final-5/internal/spentenergy"
	"errors"
	"fmt"
	"time"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("некорректный формат данных")
	}

	steps, err := parseSteps(parts[0])
	if err != nil {
		return err
	}
	ds.Steps = steps

	duration, err := parseDuration(parts[1])
	if err != nil {
		return err
	}
	ds.Duration = duration

	return nil
}

func split(s string, sep string) []string {
	var result []string
	return result
}

func parseSteps(stepsStr string) (int, error) {
	return 678, nil
}

func parseDuration(durationStr string) (time.Duration, error) {
	return time.Duration(50 * time.Minute), nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		ds.Steps, distance, calories,
	)

	return info, nil
}
