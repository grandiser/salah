<div align="center">

# Salah

*And establish prayer at the two ends of the day and at the approach of the night. Indeed, good deeds do away with misdeeds. That is a reminder for those who remember.* - Quran 11:114

Minimal, customizable, and neofetch-like Prayer Times CLI, written in Go

---
<img width="597" height="365" alt="image" src="https://github.com/user-attachments/assets/77590643-46c2-4a78-b2bf-83d625d09c61" />

---

</div>

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
git clone github.com/grandiser/salah.git
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

---

## Display Examples

|  |  |
|---|---|
| <img width="486" height="286" alt="image" src="https://github.com/user-attachments/assets/081e86aa-da34-4cc3-b19e-cb2b8eb35ac6" /> | <img width="486" height="286" alt="image" src="https://github.com/user-attachments/assets/87a70c8e-3f98-4a96-b696-84b44f3c72e1" /> |
| <img width="486" height="286" alt="image" src="https://github.com/user-attachments/assets/b4ddfadf-f04f-41dd-b3da-486fa5930353" /> | <img width="486" height="286" alt="image" src="https://github.com/user-attachments/assets/cc04cbe6-6245-4933-bfad-8f2322a5f298" /> |

---

## How It Works

1. **IP Detection:** Gets your public IP address using [httpbin.org/ip](https://httpbin.org/ip)
2. **Geolocation:** Converts IP to coordinates using [ip-api.com](http://ip-api.com)
3. **City Geocoding:** Uses [Open-Meteo Geocoding API](https://open-meteo.com/) for city-to-coordinates
4. **Prayer Times:** Fetches prayer times from [AlAdhan API](https://aladhan.com/prayer-times-api)
5. **Display:** Shows formatted prayer times in your terminal

---

## License

MIT License — see [LICENSE](LICENSE) for details.

---

## Acknowledgements

- [AlAdhan API](https://aladhan.com/prayer-times-api) for providing accurate and reliable prayer times data.
- Inspired by [stormy](https://github.com/ashish0kumar/stormy), a minimal, neofetch-like weather CLI.

