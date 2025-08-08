package prayers

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mehran-prs/gopersian"
	"runtime"
)

// Color codes list : https://www.ditig.com/256-colors-cheat-sheet

var userOS string

func init() {
	userOS = runtime.GOOS
}

func BasmalahFormatter(basmalah string) string {
	switch userOS {

	case "darwin":
		return fmt.Sprintf("%-7s%s%-3s\n", "", basmalah, "")
	case "linux":
		return fmt.Sprintf("%-6s%s%-3s\n", "", basmalah, "")
	case "windows":
		basmalah, err := gopersian.Bidi("بِسْمِ ٱللّٰهِ ٱلرَّحْمَٰنِ ٱلرَّحِيم")
		if err != nil {
			fmt.Println("Error converting arabic text bidirectionally")
		}
		return fmt.Sprintf("%-2s%s\n\n", "", basmalah)
	default:
		return ""
	}
}

func BasmalahPrinter(basmalahSprint string) {
	// Choose colors here
	switch userOS {

	case "darwin":
		fmt.Printf(basmalahSprint)
	case "linux":
		whitePrint := color.New(color.FgHiWhite, color.Bold).PrintFunc()
		whitePrint(basmalahSprint)
	case "windows":
		whitePrint := color.New(color.FgHiWhite, color.Bold).PrintFunc()
		whitePrint(basmalahSprint)
		return

	default:
		return
	}
}

func DateFormatter(todayDate string) string {
	switch userOS {

	case "darwin":
		return fmt.Sprintf("%-1s%s\n", "", todayDate)
	case "linux":
		return fmt.Sprintf("%-1s%s\n", "", todayDate)
	case "windows":
		return fmt.Sprintf("%-1s%s\n", "", todayDate)

	default:
		return ""
	}
}

func DatePrinter(dateSprint string) {
	// Choose colors here
	switch userOS {

	case "darwin":
		fmt.Printf(boldColor256(157, dateSprint))
	case "linux":
		fmt.Printf(boldColor256(157, dateSprint))
	case "windows":
		fmt.Printf(boldColor256(157, dateSprint))
	default:
		return
	}
}

func LoaderFormatter(nextPrayerName string, loadingSquares string, timeRemaining string) string {
	switch userOS {

	case "darwin":
		return fmt.Sprintf("%-1sNext: %s %s %s\n", "", nextPrayerName, loadingSquares, timeRemaining)
	case "linux":
		return fmt.Sprintf("%-1sNext: %s %s %s\n", "", nextPrayerName, loadingSquares, timeRemaining)
	case "windows":
		return fmt.Sprintf("%-1sNext: %s %s %s\n", "", nextPrayerName, loadingSquares, timeRemaining)
	default:
		return ""
	}
}

func LoaderPrinter(loaderSprint string) {
	// Choose colors here
	switch userOS {

	case "darwin":
		fmt.Printf(boldColor256(153, loaderSprint))
	case "linux":
		fmt.Printf(boldColor256(153, loaderSprint))
	case "windows":
		fmt.Printf(boldColor256(153, loaderSprint))
	default:
		return
	}
}

func PrayerFormatter(prayerName string, prayerTime string) string {
	switch userOS := runtime.GOOS; userOS {

	case "darwin":
		return fmt.Sprintf("%-7s : %s", prayerName, prayerTime)
	case "linux":
		return fmt.Sprintf("%-7s : %s", prayerName, prayerTime)
	case "windows":
		return fmt.Sprintf("%-7s : %s", prayerName, prayerTime)
	default:
		return ""
	}

}

func PrayerColorer(prayerSprint string, isPrev bool, isNext bool) string {
	// Change Colors here
	switch userOS {
	case "darwin":
		if isPrev {
			return fmt.Sprintf(boldColor256(153, prayerSprint))
		} else if isNext {
			return fmt.Sprintf(boldColor256(102, prayerSprint))
		} else {
			return fmt.Sprintf(boldColor256(102, prayerSprint))
		}
	case "windows":
		if isPrev {
			return fmt.Sprintf(boldColor256(153, prayerSprint))
		} else if isNext {
			return fmt.Sprintf(boldColor256(102, prayerSprint))
		} else {
			return fmt.Sprintf(boldColor256(102, prayerSprint))
		}
	case "linux":
		if isPrev {
			return fmt.Sprintf(boldColor256(153, prayerSprint))
		} else if isNext {
			//return fmt.Sprintf(boldColor256(15, prayerSprint))
			return prayerSprint
		} else {
			return prayerSprint
		}

	default:
		return ""
	}
}

func TableFormatter(coloredPrayer string) string {
	switch userOS {

	case "darwin":
		return fmt.Sprintf("   │ %s │\n", coloredPrayer)
	case "linux":
		return fmt.Sprintf("   │ %s │\n", coloredPrayer)
	case "windows":
		return fmt.Sprintf("   │ %s │\n", coloredPrayer)
	default:
		return ""
	}
}
