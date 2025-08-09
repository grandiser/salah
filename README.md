<div align="center">

# Salah

*And establish prayer at the two ends of the day and at the approach of the night. Indeed, good deeds do away with misdeeds. That is a reminder for those who remember.* - Quran 11:114

Minimal, customizable, and neofetch-like Prayer Times CLI, written in Go

---
<img width="542" height="372" alt="image" src="https://github.com/user-attachments/assets/73f7ead1-828d-4ef3-8c91-4c5a17dbc10e" />

---

</div>

## Features

- Automatic location detection
- Accurate daily prayer times
- English/Arabic Support
- No API keys or personal data required
- Hijri and Gregorian date support
- Compact mode for minimal output

## Installation

### Build from Source

**Linux/macOS:**
```bash
git clone github.com/grandiser/salah.git
cd salah
go build -o salah
sudo mv salah /usr/local/bin/
```

**Windows:**
```powershell
git clone https://github.com/grandiser/salah.git
cd salah
go build -o salah.exe
```

### Via Go Install

```bash
go install github.com/grandiser/salah@latest
```

### Pre-built Binaries

1. Download the binary for your platform from the [Releases page](https://github.com/grandiser/salah/releases)

2. **Linux/macOS:** Make executable and install:
```bash
chmod +x salah-linux-amd64
sudo mv salah-linux-amd64 /usr/local/bin/salah
```

3. **Windows:** Move to a directory in your PATH or run directly:
```powershell
# Run directly
./salah-windows-amd64.exe

# Or move to PATH directory
move salah-windows-amd64.exe C:\Windows\System32\salah.exe
```

## Usage

```bash
# Default (auto-detect location)
salah

# Specify city and country
salah --city "Casablanca" --country "Morocco"

# Compact output
salah --compact

# Show help
salah --help
```

## Command-line Options

* `--city <name>`: Specify city for prayer times
* `--country <name>`: Specify country (must be used with --city)
* `--compact`: Compact display mode
* `--help`: Show help message

## Configuration

A config file is created at `~/.config/salah/salah.toml` (Linux/macOS) or `%AppData%/salah/salah.toml` (Windows) on first run:

* `city`: Default city
* `country`: Default country
* `calculationMethod`: Calculation method ID or shorthand ([API reference](https://aladhan.com/prayer-times-api#get-/methods))
* `compact`: Compact output
* `useArabic`: Use Arabic for Dates and Prayer Names
* `hijriDate`: Show Hijri date
* `locateByIp`: Use IP-based location

## Display Examples

| macOS | Linux Mint |
| ----- | ---------- |
| <img width="542" height="372" alt="image" src="https://github.com/user-attachments/assets/acb700ca-2360-48a3-93f2-cd62243f88fe" /> | <img width="502" height="289" alt="image" src="https://github.com/user-attachments/assets/4fff0131-51f2-40a7-8144-6ebf6838f255" /> |
| <img width="542" height="372" alt="image" src="https://github.com/user-attachments/assets/202e7626-550b-4adc-b3ed-b4f5b3d2b2cb" /> | <img width="502" height="289" alt="image" src="https://github.com/user-attachments/assets/15d2afb8-0647-45e3-aa95-53c091496f6e" /> |

## How It Works

1. **IP Detection and Geolocation:** Gets your public IP address using [ip-api.com](http://ip-api.com)
2. **Prayer Times:** Fetches prayer times from [AlAdhan API](https://aladhan.com/prayer-times-api)
3. **Display:** Shows formatted prayer times in your terminal

## License

MIT License — see [LICENSE](LICENSE) for details.

## Acknowledgements

* [AlAdhan API](https://aladhan.com/prayer-times-api) for providing accurate and reliable prayer times data
* [IP-API](http://ip-api.com) for providing free and accurate geolocation data
* Inspired by [stormy](https://github.com/ashish0kumar/stormy), a minimal, neofetch-like weather CLI
