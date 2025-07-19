# Prayer Times CLI

A simple command-line tool that automatically detects your location using your IP address and displays Islamic prayer times for your area.

## Features

- 🌍 **Automatic location detection** - Uses your public IP to determine location
- 🕌 **Complete prayer schedule** - Shows all 5 daily prayers plus Duha
- ⚡ **Fast and lightweight** - Single binary with no dependencies
- 🔒 **Privacy-focused** - No API keys or personal data required
- 🌐 **Global support** - Works worldwide with accurate prayer times

## Installation

TBD

## Usage

Simply run the command:

```bash
salah
```

**Example output:**
```
📍 Location: New York, United States (40.7128, -74.0060)

🕌 Prayer Times for Today:
   Fajr    : 05:30 AM
   Duha    : 07:15 AM
   Dhuhr   : 12:45 PM
   Asr     : 03:30 PM
   Maghrib : 06:15 PM
   Isha    : 07:30 PM

Timezone: America/New_York
```

### Command Options

```bash
prayer-times --help        # Show help
prayer-times --version     # Show version
prayer-times --json        # Output in JSON format
```

## How It Works

1. **IP Detection**: Gets your public IP address
2. **Geolocation**: Converts IP to geographic coordinates using ip-api.com
3. **Prayer Calculation**: Fetches accurate prayer times from IslamicFinder API
4. **Display**: Shows formatted prayer times for your location

## APIs Used

- **IP Geolocation**: [ip-api.com](http://ip-api.com) (free, no key required)
- **Prayer Times**: [IslamicFinder API](https://islamicfinder.us) (free, no key required)

## Configuration

No configuration needed! The tool works out of the box with sensible defaults:

- **Calculation Method**: Islamic Society of North America (ISNA)
- **Juristic Method**: Standard (Shafi, Maliki, Hanbali)
- **Time Format**: 12-hour with AM/PM

## Contributing

Contributions welcome! Please feel free to:

1. Report bugs via [GitHub Issues](https://github.com/yourusername/prayer-times/issues)
2. Submit feature requests
3. Create pull requests

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Acknowledgments

- [ip-api.com](https://httpbin.org/ip) for free IP geolocation
- [IslamicFinder](https://islamicfinder.us) for accurate prayer time calculations
- The Go community for excellent tooling

---

**Made with ❤️ for the Muslim community**

*"And establish prayer at the two ends of the day and at the approach of the night. Indeed, good deeds do away with misdeeds. That is a reminder for those who remember."* - Quran 11:114
