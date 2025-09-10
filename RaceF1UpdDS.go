package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type RaceResult struct {
	Pilot     string
	TotalTime time.Duration
	LapTimes  []time.Duration
}

func main() {
	rand.Seed(time.Now().UnixNano())

	players := map[string]string{
		"Ferrari":  "Leclerc",
		"Red Bull": "Max Verstappen",
		"Mclaren":  "Lando Noris",
		"Apx":      "Sonny Hayes",
	}

	var wg sync.WaitGroup
	results := make(chan RaceResult, len(players))
	var raceResults []RaceResult

	// Запускаем всех гонщиков
	for _, pilot := range players {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			results <- Racing(p)
		}(pilot)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		raceResults = append(raceResults, result)
	}

	sort.Slice(raceResults, func(i, j int) bool {
		return raceResults[i].TotalTime < raceResults[j].TotalTime
	})

	// Выводим результаты
	fmt.Println("\nРезультаты гонки:")
	for i, result := range raceResults {
		fmt.Printf("%d место: %s\n", i+1, result.Pilot)
		fmt.Printf("Общее время: %v\n", result.TotalTime)
		fmt.Println("Время по кругам:")
		for lap, lapTime := range result.LapTimes {
			fmt.Printf("Круг %d: %v\n", lap+1, lapTime)
		}
		fmt.Println()
	}
}

func Racing(pilot string) RaceResult {
	var totalTime time.Duration
	lapTimes := make([]time.Duration, 3)

	for lap := 0; lap < 3; lap++ {
		lapTime := time.Duration(rand.Intn(1000)+500) * time.Millisecond
		lapTimes[lap] = lapTime
		totalTime += lapTime
		time.Sleep(lapTime)
		fmt.Printf("%s завершил круг %d за %v\n", pilot, lap+1, lapTime)
	}

	return RaceResult{
		Pilot:     pilot,
		TotalTime: totalTime,
		LapTimes:  lapTimes,
	}
}
