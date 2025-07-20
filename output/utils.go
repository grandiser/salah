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

func ShowTimesBetween(prevPrayer string, nextPrayer string, prayers []Prayer, showRemaining bool) {
	var nextPrayerTime string
	timeFormat := "15:04"

	for _, prayer := range prayers {
		if nextPrayer == prayer.Name {
			nextPrayerTime = prayer.Time
		}
	}

	nowTimeStr := time.Now().Format(timeFormat)
	nowTime, err := time.Parse(timeFormat, nowTimeStr)
	if err != nil {
		log.Fatalln(err)
	}

	nowTime.Format(timeFormat)

	if showRemaining {
		nextTime, err := time.Parse(timeFormat, nextPrayerTime)
		if err != nil {
			log.Fatalln(err)
		}

		timeRemaining := nextTime.Sub(nowTime).String()
		if prevPrayer == "Isha" {
			//TODO: FIX CONDTION SO THAT WHEN ITS BETWEEN ISHA AND FAJR THE TIME REMAINING ISNT 18HRS
			fajrTime := nextTime.Add(10)
			timeRemaining := nowTime.Sub(fajrTime).String()
			fmt.Printf(timeRemaining)
		}

		formatter := "   %-7s : %s\n"
		timeRemainingPrint := color.New(color.FgYellow, color.Bold).PrintfFunc()

		fmt.Printf("\n")
		timeRemainingPrint(formatter, "Time Remaining: ", timeRemaining)
		fmt.Printf("\n")
	}
}
