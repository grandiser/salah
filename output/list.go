package output

import (
	"github.com/grandiser/salah/times"
)

type Prayer struct {
	Name string
	Time string
}

func OutputListAladhan(aladhanTimes times.AladhanAPIResponse) {
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
	ShowPrayerTimes(curPrayer, nextPrayer, prayers)
}

func OutputListIslamicFinder(islamicFinderTimes times.IslamicFinder) {
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
	ShowPrayerTimes(curPrayer, nextPrayer, prayers)
}
