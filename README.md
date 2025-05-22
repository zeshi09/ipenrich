# IPEnrich 🔍

**IPEnrich** is a powerful terminal-based IP enrichment dashboard built in Go.\
It parses your server logs (e.g. `auth.log`), extracts IP addresses, and
enriches them with:

- 🌍 **GeoIP** info (country, city, organization)
- 🚨 **AbuseIPDB** threat scores
- 🧪 **VirusTotal** last analysis stats

All displayed in a scrollable, colored, interactive **TUI** interface.

---

## 📦 Installation

Download the latest release from
[Releases](https://github.com/zeshi09/ipenrich/releases), or:

```bash
git clone https://github.com/zeshi09/ipenrich
cd ipenrich
go build -o ipenrich main.go
```

---

## ⚙️ Usage

```bash
ipenrich /var/log/auth.log
```

- Displays a scrollable dashboard
- Press `j/k` to scroll, `q` to quit

Example:

```
ipenrich ./auth.log
```

---

## 🔐 API Keys

You can enrich IPs using:

| API        | Required? | Env Variable        |
| ---------- | --------- | ------------------- |
| GeoIP      | No        | –                   |
| AbuseIPDB  | Optional  | `ABUSEIPDB_API_KEY` |
| VirusTotal | Optional  | `VT_API_KEY`        |

```bash
export ABUSEIPDB_API_KEY=your_key
export VT_API_KEY=your_key
```

---

## 🖥 Features

- [x] Scrollable viewport
- [x] Colored threat levels
- [x] Horizontal alignment with Lipgloss
- [x] Live progress bar
- [x] Go 1.24+ compatible

---

## 🧪 Sample Screenshot

![IPEnrich TUI screenshot](https://raw.githubusercontent.com/zeshi09/ipenrich/main/assets/screenshot.png)

---

## 💡 Future plans

- Export to CSV/JSON
- Web UI version
- Threat Intel integrations (Shodan, GreyNoise, etc)

---

## 🧠 Author

Developed by [@zeshi09](https://github.com/zeshi09) — made for real-world TI
use.

---

## ⚖ License

MIT License. Use it freely, give credit 🙏
