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

	prevPrayer, nextPrayer := GetCurrentPrayers(prayers)

	ShowDateGregorian()
	ShowTimeRemaining(prevPrayer, nextPrayer, prayers)
	ShowPrayersList(prevPrayer, nextPrayer, prayers)
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
	ShowDateGregorian()
	ShowTimeRemaining(prevPrayer, nextPrayer, prayers)
	ShowPrevPrayer(prevPrayer, prayers)
	ShowNextPrayer(nextPrayer, prayers)
}

func AladhanHandler(aladhanTimes times.AladhanAPIResponse, nextOnly bool) {
	if nextOnly {
		SingleAladhan(aladhanTimes)
	} else {
		ListAladhan(aladhanTimes)
	}
}
