package output

import (
	"fmt"
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

	prevPrayer, nextPrayer := GetCurrentPrayers(prayers)

	fmt.Printf("\n")
	ShowPrayersList(prevPrayer, nextPrayer, prayers)
	ShowPrayerLoader(prevPrayer, nextPrayer)
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
	ShowPrayerLoader(prevPrayer, nextPrayer)
}

func AladhanHandler(aladhanTimes times.AladhanAPIResponse, nextOnly bool) {
	if nextOnly {
		SingleAladhan(aladhanTimes)
	} else {
		ListAladhan(aladhanTimes)
	}
}
