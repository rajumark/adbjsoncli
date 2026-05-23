# `adbjson shell dumpsys diskstats`

## adbjson

**Command:**
```bash
adbjson shell dumpsys diskstats
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "cache-free": "1208200K / 6082144K total = 19% free",
    "data-free": "1208200K / 6082144K total = 19% free",
    "disk_write_speed_kbps": "14998",
    "file_based_encryption": ": true",
    "latency": "0ms [512B Data Write]",
    "metadata-free": "9588K / 11248K total = 85% free",
    "recent_disk_write_speed_(kb/s)_=": "14998",
    "system-free": "0K / 580644K total = 0% free"
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys diskstats
```

**Output:**
```
Latency: 0ms [512B Data Write]
Recent Disk Write Speed (kB/s) = 14998
Data-Free: 1208200K / 6082144K total = 19% free
Cache-Free: 1208200K / 6082144K total = 19% free
System-Free: 0K / 580644K total = 0% free
Metadata-Free: 9588K / 11248K total = 85% free
File-based Encryption: true
```
