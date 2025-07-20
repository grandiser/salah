package output

import (
	"github.com/grandiser/salah/times"
)

type Prayer struct {
	Name string
	Time string
}

func ListAladhan(aladhanTimes times.AladhanAPIResponse) {
	prayers := []Prayer{
		{"Fajr", aladhanTimes.Data.Timings.Fajr},
		{"Sunrise", aladhanTimes.Data.Timings.Sunrise},
		{"Dhuhr", aladhanTimes.Data.Timings.Dhuhr},
		{"Asr", aladhanTimes.Data.Timings.Asr},
		{"Maghrib", aladhanTimes.Data.Timings.Maghrib},
		{"Isha", aladhanTimes.Data.Timings.Isha},
	}

	curPrayer, nextPrayer := GetCurrentPrayers(prayers)

	ShowDate()
	ShowPrayersList(curPrayer, nextPrayer, prayers)
}

func ListIslamicFinder(islamicFinderTimes times.IslamicFinder) {
	prayers := []Prayer{
		{"Fajr", islamicFinderTimes.Results.Fajr},
		{"Sunrise", islamicFinderTimes.Results.Duha},
		{"Dhuhr", islamicFinderTimes.Results.Dhuhr},
		{"Asr", islamicFinderTimes.Results.Asr},
		{"Maghrib", islamicFinderTimes.Results.Maghrib},
		{"Isha", islamicFinderTimes.Results.Isha},
	}

	curPrayer, nextPrayer := GetCurrentPrayers(prayers)
	ShowDate()
	ShowPrayersList(curPrayer, nextPrayer, prayers)
}

func SingleAladhan(aladhanTimes times.AladhanAPIResponse) {
	prayers := []Prayer{
		{"Fajr", aladhanTimes.Data.Timings.Fajr},
		{"Sunrise", aladhanTimes.Data.Timings.Sunrise},
		{"Dhuhr", aladhanTimes.Data.Timings.Dhuhr},
		{"Asr", aladhanTimes.Data.Timings.Asr},
		{"Maghrib", aladhanTimes.Data.Timings.Maghrib},
		{"Isha", aladhanTimes.Data.Timings.Isha},
	}
	prevPrayer, nextPrayer := GetCurrentPrayers(prayers)
	ShowDate()
	ShowTimesBetween(prevPrayer, nextPrayer, prayers, true, true)
	ShowNextPrayer(nextPrayer, prayers)
}

func SingleIslamicFinder(islamicFinderTimes times.IslamicFinder) {
	prayers := []Prayer{
		{"Fajr", islamicFinderTimes.Results.Fajr},
		{"Sunrise", islamicFinderTimes.Results.Duha},
		{"Dhuhr", islamicFinderTimes.Results.Dhuhr},
		{"Asr", islamicFinderTimes.Results.Asr},
		{"Maghrib", islamicFinderTimes.Results.Maghrib},
		{"Isha", islamicFinderTimes.Results.Isha},
	}

	prevPrayer, nextPrayer := GetCurrentPrayers(prayers)
	ShowDate()
	ShowTimesBetween(prevPrayer, nextPrayer, prayers, true, true)
	ShowNextPrayer(nextPrayer, prayers)
}

func AladhanHandler(aladhanTimes times.AladhanAPIResponse, listAll bool) {
	if listAll {
		ListAladhan(aladhanTimes)
	} else {
		SingleAladhan(aladhanTimes)
	}
}

func IslamicFinderHandler(islamicFinderTimes times.IslamicFinder, listAll bool) {
	if listAll {
		ListIslamicFinder(islamicFinderTimes)
	} else {
		SingleIslamicFinder(islamicFinderTimes)
	}
}
