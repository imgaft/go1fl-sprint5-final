package spentenergy

import (
	"fmt"
	"time"
)

const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

// Distance вычисляет пройденную дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	return float64(steps) * height * stepLengthCoefficient / mInKm
}

// MeanSpeed вычисляет среднюю скорость в километрах в час.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps < 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

// RunningSpentCalories рассчитывает расход калорий при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("некорректные параметры активности")
	}
	return weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH, nil
}

// WalkingSpentCalories рассчитывает расход калорий при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return calories * walkingCaloriesCoefficient, nil
}
