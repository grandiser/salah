package output

import (
	"log"
	"math"
	"strings"
	"time"
)

func ConvertStringToTime(timeStr string) time.Time {
	timeFormat := "15:04"

	timeTime, err := time.Parse(timeFormat, timeStr)
	if err != nil {
		log.Fatalln(err)
	}
	return timeTime
}

func GetCurrentPrayers(prayers []Prayer) (prevPrayer Prayer, nextPrayer Prayer) {
	now := time.Now()
	curTime := now.Format("15:04")

	for idx, prayer := range prayers {
		prevName, prevTime := prayer.Name, prayer.Time

		if idx+1 == len(prayers) {
			prevPrayer = Prayer{prevName, prevTime}
			nextPrayer = Prayer{"Fajr", prayers[0].Time}
			return prevPrayer, nextPrayer
		}

		nextName, nextTime := prayers[idx+1].Name, prayers[idx+1].Time

		if curTime > prevPrayer.Time && curTime < nextPrayer.Time {
			prevPrayer = Prayer{prevName, prevTime}
			nextPrayer = Prayer{nextName, nextTime}
			return prevPrayer, nextPrayer
		}
	}
	// fallback
	return prayers[0], prayers[1]
}

func GetLoadingSquares(prevPrayer Prayer, nextPrayer Prayer) string {
	loaded := "▣"
	unloaded := "▢"

	timeBetween := CalculateTimeDiff(prevPrayer, prevPrayer.Time, nextPrayer.Time)
	timeRemaining := GetTimeRemaining(prevPrayer, nextPrayer)
	nSquares := int(math.Ceil(float64((timeRemaining / timeBetween) * 10)))

	loadingSquares := strings.Repeat(loaded, nSquares) + strings.Repeat(unloaded, 10-nSquares)
	return loadingSquares
}

func CalculateTimeDiff(prevPrayer Prayer, prevTimeStr string, nextTimeStr string) time.Duration {
	prevTime := ConvertStringToTime(prevTimeStr)
	nextTime := ConvertStringToTime(nextTimeStr)

	timeDiff := nextTime.Sub(prevTime)

	if prevPrayer.Name == "Isha" && nextTime.Before(prevTime) {
		// Next prayer is tomorrow, add 24 hours
		tomorrow := nextTime.Add(24 * time.Hour)
		timeDiff = tomorrow.Sub(prevTime)
	} else {
		timeDiff = nextTime.Sub(prevTime)
	}

	return timeDiff
}

func GetTimeRemaining(prevPrayer Prayer, nextPrayer Prayer) time.Duration {
	timeFormat := "15:04"

	nowTimeStr := time.Now().Format(timeFormat)
	timeRemaining := CalculateTimeDiff(prevPrayer, nowTimeStr, nextPrayer.Time)

	return timeRemaining
}
