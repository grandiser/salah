package prayers

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Config struct {
	City              string `toml:"city"`
	Country           string `toml:"country"`
	CalculationMethod string `toml:"calculationMethod"`
	Compact           bool   `toml:"Compact"`
	UseColors         bool   `toml:"useColors"`
	UseArabic         bool   `toml:"useArabic"`
	HijriDate         bool   `toml:"hijriDate"`
	LocateByIp        bool   `toml:"locateByIp"`
}

type Flags struct {
	City    string
	Country string
	Compact bool
	Help    bool
}

type Method struct {
	ID   string
	Name string
}

var methods = []Method{
	{ID: "0", Name: "JAFARI"},
	{ID: "1", Name: "KARACHI"},
	{ID: "2", Name: "ISNA"},
	{ID: "3", Name: "MWL"},
	{ID: "4", Name: "MAKKAH"},
	{ID: "5", Name: "EGYPT"},
	{ID: "7", Name: "TEHRAN"},
	{ID: "8", Name: "GULF"},
	{ID: "9", Name: "KUWAIT"},
	{ID: "10", Name: "QATAR"},
	{ID: "11", Name: "SINGAPORE"},
	{ID: "12", Name: "FRANCE"},
	{ID: "13", Name: "TURKEY"},
	{ID: "14", Name: "RUSSIA"},
	{ID: "15", Name: "MOONSIGHTING"},
	{ID: "16", Name: "DUBAI"},
	{ID: "17", Name: "JAKIM"},
	{ID: "18", Name: "TUNISIA"},
	{ID: "19", Name: "ALGERIA"},
	{ID: "20", Name: "KEMENAG"},
	{ID: "21", Name: "MOROCCO"},
	{ID: "22", Name: "PORTUGAL"},
	{ID: "23", Name: "JORDAN"},
	{ID: "99", Name: "CUSTOM"},
}

func GetDefaultConfig() Config {
	return Config{
		City:              "",
		Country:           "",
		CalculationMethod: "",
		Compact:           false,
		UseColors:         true,
		UseArabic:         false,
		HijriDate:         false,
		LocateByIp:        true,
	}
}

func ValidateConfig(config *Config) {
	if runtime.GOOS == "windows" && config.UseArabic {
		fmt.Fprintln(os.Stderr, "Warning: Arabic is not well supported on Windows.\nPlease change your config values.\nReverting to English.")
		config.UseArabic = false
	}

	if config.CalculationMethod != "" {
		userInputMethod := config.CalculationMethod
		upperCaseInputMethod := strings.ToUpper(userInputMethod)
		exists := false
		for _, method := range methods {
			if method.ID == upperCaseInputMethod || method.Name == upperCaseInputMethod {
				exists = true
				config.CalculationMethod = upperCaseInputMethod
				break
			}
		}
		if !exists {
			log.Fatalln("\n\n", userInputMethod, "is not a valid calculation method.\nIt is recommended to leave the config field as \"\" for automatic detection.\nOr use 'id' value or method abbreviation.\nSee the API Documentation for more information: \nhttps://aladhan.com/prayer-times-api#get-/methods")

		}
	}
	if !config.LocateByIp {
		if config.City == "" {
			log.Fatalln(os.Stderr, "Warning: cannot leave 'City' value empty. Please enter your city in the config file.")
		}
		if config.City != "" && config.Country == "" {
			fmt.Fprintln(os.Stderr, "Warning: leaving the 'Country' value empty may cause inaccurate results.\nPlease enter your country in the config file or use the --country flag.")
		}
	}
}

func GetConfigPath() string {
	var configDir string

	if runtime.GOOS == "windows" {
		// Windows: Use AppData directory
		dir, err := os.UserConfigDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to get config directory:", err)
			dir, err = os.UserHomeDir()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to get home directory:", err)
				return ""
			}
			configDir = filepath.Join(dir, "salah")
		} else {
			configDir = filepath.Join(dir, "salah")
		}
	} else {
		// Linux/macOS: Follow XDG Base Directory Specification
		xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
		if xdgConfigHome != "" {
			// Use XDG_CONFIG_HOME if set
			configDir = filepath.Join(xdgConfigHome, "salah")
		} else {
			// Fall back to ~/.config/salah
			homeDir, err := os.UserHomeDir()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to get home directory:", err)
				return ""
			}
			configDir = filepath.Join(homeDir, ".config", "salah")
		}
	}

	return filepath.Join(configDir, "salah.toml")
}

func ReadConfig() Config {
	configPath := GetConfigPath()
	if configPath == "" {
		return GetDefaultConfig()
	}

	// Create directory if it doesn't exist
	configDir := filepath.Dir(configPath)
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, 0755); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to create config directory:", err)
			return GetDefaultConfig()
		}
	}

	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Create default config
		defaultConfig := GetDefaultConfig()
		file, err := os.Create(configPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to create config file:", err)
			return defaultConfig
		}
		defer file.Close()

		if err := toml.NewEncoder(file).Encode(defaultConfig); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to write default config:", err)
			return defaultConfig
		}

		fmt.Printf("No config detected, config created at %s.\n", configPath)
		fmt.Println("Please edit the configuration file to add your city and country.")
		return defaultConfig
	}

	// Read existing config
	var config Config
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read config file:", err)
		return GetDefaultConfig()
	}

	if err := toml.Unmarshal(data, &config); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse config file, using defaults:", err)
	}

	ValidateConfig(&config)

	return config
}

func ParseFlags() Flags {
	flags := Flags{}

	flag.StringVar(&flags.City, "city", "", "City to get prayer times for")
	flag.StringVar(&flags.Country, "country", "", "Country in which city is located")
	flag.BoolVar(&flags.Compact, "compact", false, "Compact display mode")
	flag.BoolVar(&flags.Help, "help", false, "Show help")

	// Add usage information
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nConfig file is located at:", GetConfigPath())
	}

	flag.Parse()

	if flags.Help {
		flag.Usage()
		os.Exit(0)
	}

	return flags
}

func ApplyFlags(config *Config, flags Flags) {

	if flags.City != "" {
		config.LocateByIp = false
		config.City = flags.City
		config.Country = flags.Country
	}

	if flags.Country != "" {
		if flags.City != "" {
			config.LocateByIp = false
			config.City = flags.City
			config.Country = flags.Country
		} else {
			log.Fatalln(os.Stderr, "Error: The --country flag is meant to be used alongside a --city flag, not on its own.\n Try again by specifying the --city value.")
		}
	}
	if flags.Compact {
		config.Compact = true
	}

	ValidateConfig(config)
}
