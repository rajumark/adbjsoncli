# `adbjson shell dumpsys battery`

## adbjson

**Command:**
```bash
adbjson shell dumpsys battery
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "AC powered": "false",
    "Capacity level": "-1",
    "Charge counter": "10000",
    "Charging policy": "0",
    "Charging state": "0",
    "Dock powered": "false",
    "Max charging current": "0",
    "Max charging voltage": "0",
    "The last voltage value sent via the battery changed broadcast": "5000",
    "Time when the latest updated value of the Max charging current was sent via battery changed broadcast": "+3s391ms",
    "Time when the latest updated value of the voltage was sent via battery changed broadcast": "+3s391ms",
    "USB powered": "false",
    "Wireless powered": "false",
    "health": "2",
    "level": "100",
    "present": "true",
    "scale": "100",
    "status": "4",
    "technology": "Li-ion",
    "temperature": "250",
    "voltage": "5000"
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys battery
```

**Output:**
```
Current Battery Service state:
  AC powered: false
  USB powered: false
  Wireless powered: false
  Dock powered: false
  Max charging current: 0
 Time when the latest updated value of the Max charging current was sent via battery changed broadcast: +3s391ms
  Max charging voltage: 0
  Charge counter: 10000
  status: 4
  health: 2
  present: true
  level: 100
  scale: 100
  voltage: 5000
 Time when the latest updated value of the voltage was sent via battery changed broadcast: +3s391ms
 The last voltage value sent via the battery changed broadcast: 5000
  temperature: 250
  technology: Li-ion
  Charging state: 0
  Charging policy: 0
  Capacity level: -1
```
