# `adbjson shell dumpsys cpuinfo`

## adbjson

**Command:**
```bash
adbjson shell dumpsys cpuinfo
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "processes": [
      "CPU usage from 6587ms to 1487ms ago (2026-05-23 15:39:17.132 to 2026-05-23 15:39:22.232):",
      "71% 640/system_server: 36% user + 34% kernel / faults: 4184 minor",
      "1.3% 496/adbd: 0.1% user + 1.1% kernel / faults: 17040 minor",
      "0.3% 448/android.hardware.sensors-service.multihal: 0.1% user + 0.1% kernel",
      "0% 9/kworker/0:0-events: 0% user + 0% kernel",
      "0% 38/kworker/2:0-rcu_gp: 0% user + 0% kernel",
      "0% 41/rcuop/2: 0% user + 0% kernel",
      "0.1% 454/android.hardware.graphics.composer3-service.ranchu: 0% user + 0.1% kernel",
      "0% 494/surfaceflinger: 0% user + 0% kernel",
      "0.1% 1470/kworker/0:3-virtio_vsock: 0% user + 0.1% kernel",
      "0% 1563/com.google.android.gms: 0% user + 0% kernel / faults: 136 minor 233 major",
      "0% 3513/com.google.android.youtube: 0% user + 0% kernel / faults: 94 minor 207 major",
      "+0% 5350/kworker/3:2-events: 0% user + 0% kernel",
      "+0% 5424/dumpsys: 0% user + 0% kernel",
      "63% TOTAL: 10% user + 10% kernel + 42% iowait + 1.1% irq"
    ]
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys cpuinfo
```

**Output:**
```
Load: 3.84 / 2.49 / 0.98
CPU usage from 6603ms to 1503ms ago (2026-05-23 15:39:17.132 to 2026-05-23 15:39:22.232):
  71% 640/system_server: 36% user + 34% kernel / faults: 4184 minor
  1.3% 496/adbd: 0.1% user + 1.1% kernel / faults: 17040 minor
  0.3% 448/android.hardware.sensors-service.multihal: 0.1% user + 0.1% kernel
  0% 9/kworker/0:0-events: 0% user + 0% kernel
  0% 38/kworker/2:0-rcu_gp: 0% user + 0% kernel
  0% 41/rcuop/2: 0% user + 0% kernel
  0.1% 454/android.hardware.graphics.composer3-service.ranchu: 0% user + 0.1% kernel
  0% 494/surfaceflinger: 0% user + 0% kernel
  0.1% 1470/kworker/0:3-virtio_vsock: 0% user + 0.1% kernel
  0% 1563/com.google.android.gms: 0% user + 0% kernel / faults: 136 minor 233 major
  0% 3513/com.google.android.youtube: 0% user + 0% kernel / faults: 94 minor 207 major
 +0% 5350/kworker/3:2-events: 0% user + 0% kernel
 +0% 5424/dumpsys: 0% user + 0% kernel
63% TOTAL: 10% user + 10% kernel + 42% iowait + 1.1% irq
```
