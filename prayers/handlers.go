package prayers

import (
	"fmt"
	"github.com/grandiser/salah/apis"
)

type Prayer struct {
	Name string
	Time string
}

func AladhanHandler(aladhanTimes apis.AladhanAPIResponse, config Config) {
	if config.Compact {
		SingleAladhan(aladhanTimes, config)
	} else {
		ListAladhan(aladhanTimes, config)
	}
}

func ListAladhan(aladhanTimes apis.AladhanAPIResponse, config Config) {
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
	ShowBasmalah()
	ShowPrayersList(prevPrayer, nextPrayer, prayers)
	showDate(aladhanTimes, config)
	ShowPrayerLoader(prevPrayer, nextPrayer)
}

func SingleAladhan(aladhanTimes apis.AladhanAPIResponse, config Config) {
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
	showDate(aladhanTimes, config)
	ShowPrayerLoader(prevPrayer, nextPrayer)
}
