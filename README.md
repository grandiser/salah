<div align="center">

# Salah

*And establish prayer at the two ends of the day and at the approach of the night. Indeed, good deeds do away with misdeeds. That is a reminder for those who remember.* - Quran 11:114

Minimal, customizable, and neofetch-like Prayer Times CLI, written in Go

---
<img width="623" height="385" alt="image" src="https://github.com/user-attachments/assets/513d5624-93c0-4d82-80a2-f53a31623602" />

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
- `calculationMethod` = Specify a calculation method for prayers using the ID or Shorthand notation. See the API reference for more information: https://aladhan.com/prayer-times-api#get-/methods.
- `compact`: Compact output
- `useArabic`: Use Arabic for Dates and Prayer Names
- `hijriDate`: Show Hijri date
- `locateByIp`: Use IP-based location

---

---

## Display Examples
|   |   |
|---|---|
| <img width="486" height="286" alt="image" src="https://github.com/user-attachments/assets/f69660ae-053d-41f4-9f39-25d39c410bc5" /> | <img width="486" height="286" alt="image" src="https://github.com/user-attachments/assets/2e6828ed-6045-4ece-bb1b-637a3153feb9" /> |
| <img width="486" height="286" alt="image" src="https://github.com/user-attachments/assets/1e5c663f-1ef0-4d50-b39f-4cb8fe01185b" /> |   |


---

## How It Works

1. **IP Detection amd Geolocation:** Gets your public IP address using [ip-api.com](http://ip-api.com)
2. **Prayer Times:** Fetches prayer times from [AlAdhan API](https://aladhan.com/prayer-times-api)
3. **Display:** Shows formatted prayer times in your terminal

---

## License

MIT License — see [LICENSE](LICENSE) for details.

---

## Acknowledgements

- [AlAdhan API](https://aladhan.com/prayer-times-api) for providing accurate and reliable prayer times data.
- [IP-API](http://ip-api.com) for providing free and accurate geolocation data.
- Inspired by [stormy](https://github.com/ashish0kumar/stormy), a minimal, neofetch-like weather CLI.

