package output

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hablullah/go-hijri"
	"log"
	"strings"
	"time"
)

func ShowDateGregorian() {
	today := time.Now().Format("Monday, January 2 2006")
	datePrint := color.New(color.FgMagenta, color.Bold).PrintFunc()

	datePrint("\n " + today + "\n")
}

func ShowDateHijri() {

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
	curPrayerPrint := color.New(color.FgCyan, color.Bold).PrintfFunc()
	nextPrayerPrint := color.New(color.FgGreen, color.Bold).PrintfFunc()

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

func ShowPrevPrayer(prevPrayer string, prayers []Prayer) {
	for _, prayer := range prayers {
		if prevPrayer == prayer.Name {
			formatter := "   %-7s : %s\n"
			nextPrayerPrint := color.New(color.FgCyan, color.Bold).PrintfFunc()
			nextPrayerTime := prayer.Time
			nextPrayerPrint(formatter, prevPrayer, nextPrayerTime)
			return
		}
	}
	fmt.Printf("\nError showing next prayer. Try again with --list flag")
}

func ShowNextPrayer(nextPrayer string, prayers []Prayer) {
	for _, prayer := range prayers {
		if nextPrayer == prayer.Name {
			formatter := "   %-7s : %s\n\n"
			nextPrayerPrint := color.New(color.FgGreen, color.Bold).PrintfFunc()
			nextPrayerTime := prayer.Time
			nextPrayerPrint(formatter, nextPrayer, nextPrayerTime)
			return
		}
	}
	fmt.Printf("\nError showing next prayer. Try again with --list flag")
}

func ShowTimeRemaining(prevPrayer string, nextPrayer string, prayers []Prayer) {
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
	nextTime, err := time.Parse(timeFormat, nextPrayerTime)
	if err != nil {
		log.Fatalln(err)
	}

	timeRemaining := nextTime.Sub(nowTime).String()
	formatter := "\n   %s%s%s\n\n"
	timeRemainingPrint := color.New(color.FgYellow, color.Bold).PrintfFunc()

	if prevPrayer == "Isha" && nextTime.Before(nowTime) {
		// Next prayer is tomorrow, add 24 hours
		tomorrow := nextTime.Add(24 * time.Hour)
		timeRemaining = tomorrow.Sub(nowTime).String()
	} else {
		timeRemaining = nextTime.Sub(nowTime).String()
	}
	timeRemainingStr := strings.Replace(timeRemaining, "0s", "", 1)
	timeRemainingPrint(formatter, timeRemainingStr, " until ", nextPrayer)
}
