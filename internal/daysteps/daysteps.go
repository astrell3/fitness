package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	cal "github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("неверный формат данных")
	}

	stepsStr := strings.TrimSpace(parts[0])
	durationStr := strings.TrimSpace(parts[1])

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, 0, err
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть положительным")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, 0, err
	}

	if duration <= 0 {
		return 0, 0, fmt.Errorf("продолжительность должна быть положительной")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distance := (float64(steps) * stepLength) / mInKm

	calories, err := cal.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)
}
