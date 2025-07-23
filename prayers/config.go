package prayers

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	City          string `toml:"city"`
	Country       string `toml:"country"`
	Compact       bool   `toml:"Compact"`
	UseColors     bool   `toml:"useColors"`
	UseArabic     bool   `toml:"useArabic"`
	GregorianDate bool   `toml:"gregorianDate"`
	HijriDate     bool   `toml:"hijriDate"`
	LocateByIp    bool   `toml:"locateByIp"`
}

type Flags struct {
	City    string
	Country string
	Compact bool
	Help    bool
}

func GetDefaultConfig() Config {
	return Config{
		City:          "",
		Country:       "",
		Compact:       false,
		UseColors:     true,
		UseArabic:     false,
		GregorianDate: true,
		HijriDate:     false,
		LocateByIp:    true,
	}
}

func ValidateConfig(config *Config) {

	if config.GregorianDate && config.HijriDate {
		fmt.Fprintln(os.Stderr, "Warning: cannot set both 'GregorianDate' and 'HijriDate' values to 'true'.")
		fmt.Fprintln(os.Stderr, "Using Gregorian Date as default.")
		config.HijriDate = false
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
