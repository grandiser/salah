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

1. **Build from source:**

```bash
git clone https://github.com/grandiser/salah.git
cd salah
go build -o salah
```

2. **Move to your PATH:**

```bash
sudo mv salah /usr/local/bin/
```

---

## Usage

Run with default (auto-detect location):

```bash
./salah
```

Specify city and country:

```bash
./salah --city "Casablanca" --country "Morocco"
```

Compact output:

```bash
./salah --compact
```

Show help:

```bash
./salah --help
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
- `compact`: Compact output (true/false)
- `useColors`: Enable/disable colored output
- `useArabic`: Show Hijri date in Arabic
- `gregorianDate`: Show Gregorian date
- `hijriDate`: Show Hijri date
- `locateByIp`: Use IP-based location (true/false)

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

 Monday, January 2 2025
 Next: Maghrib ▣▣▣▣▣▢▢▢▢▢ 01h 23m
```

---

## License

MIT License — see [LICENSE](LICENSE) for details.

---

**Made with ❤️ for the Muslim community** 
