# Salah

Minimal, customizable, and neofetch-like Prayer Times CLI, written in Go 

---

## Features

- Automatic location detection
- Accurate daily prayer times
- Fast, single-binary CLI
- No API keys or personal data required
- Hijri and Gregorian date support
- Compact mode for minimal output

---

## Installation

### Linux/macOS

1. Option 1: Build from source:

```bash
git clone https://github.com/grandiser/salah.git
cd salah
go build -o salah

# Move to a directory in your PATH
sudo mv salah /usr/local/bin/
```

2. Option 2: Via Go Install

```bash
go install github.com/grandiser/salah@latest
```

### Windows

1. Option 1: Build from source (requires Go):

Open Command Prompt or PowerShell and run:

```powershell
git clone https://github.com/grandiser/salah.git
cd salah
go build -o salah.exe
```

You can then run `salah.exe` directly, or move it to a directory in your PATH for easier access.

2. Option 2: Via Go Install (requires Go in PATH):

```powershell
go install github.com/grandiser/salah@latest
```

The binary will be placed in your Go bin directory (usually `%USERPROFILE%\go\bin`). Add this directory to your PATH to run `salah` from anywhere.

---

## Usage

Run with default (auto-detect location):

```bash
salah
```

Specify city and country:

```bash
salah --city "Casablanca" --country "Morocco"
```

Compact output:

```bash
salah --compact
```

Show help:

```bash
salah --help
```

---

## Command-line Options

- `--city <name>`: Specify city for prayer times
- `--country <name>`: Specify country (must be used with --city)
- `--compact`: Compact display mode
- `--help`: Show help message

---

## Configuration

A config file is created at `~/.config/salah/salah.toml` (Linux/macOS) or `%AppData%/salah/salah.toml` (Windows) on first run. You can edit this file to set defaults for:

- `city`: Default city
- `country`: Default country
- `compact`: Compact output
- `useColors`: Enable/disable colored output
- `useArabic`: Show Hijri date in Arabic
- `gregorianDate`: Show Gregorian date
- `hijriDate`: Show Hijri date
- `locateByIp`: Use IP-based location

---

## How It Works

1. **IP Detection:** Gets your public IP address using [httpbin.org/ip](https://httpbin.org/ip)
2. **Geolocation:** Converts IP to coordinates using [ip-api.com](http://ip-api.com)
3. **City Geocoding:** Uses [Open-Meteo Geocoding API](https://open-meteo.com/) for city-to-coordinates
4. **Prayer Times:** Fetches prayer times from [AlAdhan API](https://aladhan.com/prayer-times-api)
5. **Display:** Shows formatted prayer times in your terminal

---

## Example Output

```
       ﷽   
   ╭────────۞────────╮
   │ Fajr    : 05:30 │
   │ Sunrise : 06:45 │
   │ Dhuhr   : 12:45 │
   │ Asr     : 16:15 │
   │ Maghrib : 19:10 │
   │ Isha    : 20:30 │
   ╰────────۞────────╯

 Al Arba'a, 28 Muḥarram 1447
 Next: Maghrib ▣▣▣▣▣▢▢▢▢▢ 01h 23m
```

---

## License

MIT License — see [LICENSE](LICENSE) for details.

---

## Acknowledgements

- [AlAdhan API](https://aladhan.com/prayer-times-api) for providing accurate and reliable prayer times data.
- Inspired by [stormy](https://github.com/ashish0kumar/stormy), a minimal, neofetch-like weather CLI.

**Made with ❤️ for the Muslim community** 
