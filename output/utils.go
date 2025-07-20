package output

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"time"
)

func ShowDate() {
	today := time.Now().Format("Monday, January 2 2006")
	datePrint := color.New(color.FgMagenta, color.Bold).PrintFunc()

	datePrint("\n" + today + "\n")
}

func GetCurrentPrayers(prayers []Prayer) (string, string) {
	now := time.Now()
	curTime := now.Format("15:04")

	for idx, prayer := range prayers {
		curPrayerName, curPrayerTime := prayer.Name, prayer.Time

		if idx+1 == len(prayers) {
			return curPrayerName, "Fajr"
		}

		nextPrayerName, nextPrayerTime := prayers[idx+1].Name, prayers[idx+1].Time

		if curTime > curPrayerTime && curTime < nextPrayerTime {
			return curPrayerName, nextPrayerName
		}
	}
	return "Error", "Error"
}

func ShowPrayersList(curPrayer string, nextPrayer string, prayers []Prayer) {
	formatter := "   %-7s : %s\n"
	curPrayerPrint := color.New(color.FgGreen, color.Bold).PrintfFunc()
	nextPrayerPrint := color.New(color.FgYellow, color.Bold).PrintfFunc()

	for _, prayer := range prayers {
		if curPrayer == prayer.Name {
			curPrayerPrint(formatter, curPrayer, prayer.Time)

		} else if nextPrayer == prayer.Name {
			nextPrayerPrint(formatter, nextPrayer, prayer.Time)

		} else {
			fmt.Printf(formatter, prayer.Name, prayer.Time)
		}
	}
	fmt.Printf("\n")
}

func ShowNextPrayer(nextPrayer string, prayers []Prayer) {
	for _, prayer := range prayers {
		if nextPrayer == prayer.Name {
			formatter := "   %-7s : %s\n"
			nextPrayerPrint := color.New(color.FgCyan, color.Bold).PrintfFunc()
			nextPrayerTime := prayer.Time
			nextPrayerPrint(formatter, nextPrayer, nextPrayerTime)
			return
		}
	}
	fmt.Printf("\nError showing next prayer. Try again with --list flag")
}

func ShowTimesBetween(curPrayer string, nextPrayer string, prayers []Prayer, showElapsed bool, showRemaining bool) {
	var prevPrayerTime string
	var nextPrayerTime string

	for _, prayer := range prayers {
		if curPrayer == prayer.Name {
			prevPrayerTime = prayer.Time
			fmt.Printf(prevPrayerTime)
		}
		if nextPrayer == prayer.Name {
			nextPrayerTime = prayer.Time
			fmt.Printf(nextPrayerTime)
		}
	}

	nowTimeStr := time.Now().Format("15:03")

	nowTime, err := time.Parse("15:03", nowTimeStr)
	if err != nil {
		log.Fatalln(err)
	}

	if showElapsed && showRemaining {
		prevTime, err := time.Parse("15:03", prevPrayerTime)
		if err != nil {
			log.Fatalln(err)
		}

		nextTime, err := time.Parse("15:03", nextPrayerTime)
		if err != nil {
			log.Fatalln(err)
		}

		timeElapsed := nowTime.Sub(prevTime).String()
		timeRemaining := nextTime.Sub(nowTime).String()

		timeElapsedPrint := color.New(color.FgGreen, color.Bold).PrintfFunc()
		timeRemainingPrint := color.New(color.FgYellow, color.Bold).PrintfFunc()

		timeElapsedPrint(timeElapsed)
		timeRemainingPrint(timeRemaining)
	}

}
