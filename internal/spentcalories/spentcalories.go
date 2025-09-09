package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("неверный формат данных")
	}

	stepsStr := strings.TrimSpace(parts[0])
	trainingType := strings.TrimSpace(parts[1])
	durationStr := strings.TrimSpace(parts[2])

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, "", 0, err
	}

	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов должно быть положительным")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, "", 0, err
	}

	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("продолжительность должна быть положительной")
	}

	return steps, trainingType, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLenght := height * stepLengthCoefficient
	distanceM := float64(steps) * stepLenght
	return distanceM / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	dist := distance(steps, height)
	hours := duration.Hours()

	if hours == 0 {
		return 0
	}

	return dist / hours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, trainingType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var calories float64
	var dist float64
	var speed float64

	switch trainingType {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)

	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)

	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", trainingType)
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration.Hours(), dist, speed, calories), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть положительным")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть положительным")
	}

	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть положительным")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть положительной")
	}

	avgSpeed := meanSpeed(steps, height, duration)
	if avgSpeed <= 0 {
		return 0, fmt.Errorf("средняя скорость должна быть положительной")
	}

	durationMin := duration.Minutes()
	calories := (weight * avgSpeed * durationMin) / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть положительным")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть положительным")
	}

	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть положительным")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть положительной")
	}

	avgSpeed := meanSpeed(steps, height, duration)
	if avgSpeed <= 0 {
		return 0, fmt.Errorf("средняя скорость должна быть положительной")
	}

	durationMin := duration.Minutes()
	calories := (weight * avgSpeed * durationMin) / minInH
	calories *= walkingCaloriesCoefficient

	return calories, nil
}
