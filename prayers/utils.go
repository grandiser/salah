package prayers

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/grandiser/salah/apis"
)

func boldColor256(colorCode int, text string) string {
	// Color codes list : https://www.ditig.com/256-colors-cheat-sheet
	return fmt.Sprintf("\033[1;38;5;%dm%s\033[0m", colorCode, text)
}

func GetHijriDate(aladhanTimes apis.AladhanAPIResponse, config Config) string {
	var hijri string
	dayNum := aladhanTimes.Data.Date.Hijri.Day
	format := "%s, %s %s %s"

	if config.UseArabic {
		arMonth := aladhanTimes.Data.Date.Hijri.Month.Ar
		year := aladhanTimes.Data.Date.Hijri.Year
		arWeekday := aladhanTimes.Data.Date.Hijri.Weekday.Ar
		hijri = fmt.Sprintf(format, arWeekday, dayNum, arMonth, year)
	} else {
		enMonth := aladhanTimes.Data.Date.Hijri.Month.En
		year := aladhanTimes.Data.Date.Hijri.Year
		enWeekday := aladhanTimes.Data.Date.Hijri.Weekday.En
		hijri = fmt.Sprintf(format, enWeekday, dayNum, enMonth, year)
	}
	return hijri
}

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

		if curTime > prevTime && curTime < nextTime {
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

	var nSquares int

	timeBetween := CalculateTimeDiff(prevPrayer, prevPrayer.Time, nextPrayer.Time)
	timeRemaining := GetTimeRemaining(prevPrayer, nextPrayer)

	nSquares = int(math.Ceil((float64(timeRemaining) / float64(timeBetween)) * 10))
	// avoids error when you run 'salah' during the first minute of the new prayer
	if nSquares > 10 {
		nSquares = 10
	}

	loadingSquares := strings.Repeat(loaded, 10-nSquares) + strings.Repeat(unloaded, nSquares)
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
