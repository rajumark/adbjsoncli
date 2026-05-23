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
      "CPU usage from 21088ms to 711ms ago (2026-05-23 16:04:15.323 to 2026-05-23 16:04:35.700):",
      "4.1% 644/system_server: 1.9% user + 2.2% kernel / faults: 1771 minor 2 major",
      "0.1% 451/android.hardware.graphics.composer3-service.ranchu: 0% user + 0.1% kernel / faults: 22 minor",
      "0.1% 447/android.hardware.sensors-service.multihal: 0% user + 0% kernel",
      "0.1% 489/surfaceflinger: 0% user + 0% kernel / faults: 120 minor 2 major",
      "0.1% 3376/com.google.android.youtube: 0% user + 0% kernel / faults: 163 minor 354 major",
      "0% 491/adbd: 0% user + 0% kernel / faults: 3329 minor 14 major",
      "0% 889/com.android.systemui: 0% user + 0% kernel / faults: 431 minor",
      "0% 2710/com.google.android.gms.unstable: 0% user + 0% kernel / faults: 458 minor 726 major",
      "0% 3737/artd: 0% user + 0% kernel / faults: 2300 minor",
      "0% 65/kswapd0: 0% user + 0% kernel",
      "0% 67/kworker/u17:0-blk_crypto_wq: 0% user + 0% kernel",
      "0% 286/logd: 0% user + 0% kernel / faults: 79 minor 3 major",
      "0% 337/jbd2/dm-5-8: 0% user + 0% kernel",
      "0% 1033/com.android.phone: 0% user + 0% kernel / faults: 488 minor 1 major",
      "0% 2212/com.google.android.googlequicksearchbox:search: 0% user + 0% kernel / faults: 279 minor 469 major",
      "0% 2249/com.google.android.apps.messaging: 0% user + 0% kernel / faults: 455 minor 1178 major",
      "0% 2565/com.android.vending: 0% user + 0% kernel / faults: 586 minor 1286 major",
      "0% 3180/com.google.android.rkpdapp: 0% user + 0% kernel / faults: 169 minor 34 major",
      "0% 5832/com.google.android.photopicker: 0% user + 0% kernel / faults: 359 minor 262 major",
      "0% 5910/kworker/0:1-events: 0% user + 0% kernel",
      "+0% 6206/dumpsys: 0% user + 0% kernel",
      "39% TOTAL: 0.6% user + 0.9% kernel + 37% iowait + 0.5% irq + 0% softirq"
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
Load: 0.3 / 0.51 / 0.32
CPU usage from 21105ms to 728ms ago (2026-05-23 16:04:15.323 to 2026-05-23 16:04:35.700):
  4.1% 644/system_server: 1.9% user + 2.2% kernel / faults: 1771 minor 2 major
  0.1% 451/android.hardware.graphics.composer3-service.ranchu: 0% user + 0.1% kernel / faults: 22 minor
  0.1% 447/android.hardware.sensors-service.multihal: 0% user + 0% kernel
  0.1% 489/surfaceflinger: 0% user + 0% kernel / faults: 120 minor 2 major
  0.1% 3376/com.google.android.youtube: 0% user + 0% kernel / faults: 163 minor 354 major
  0% 491/adbd: 0% user + 0% kernel / faults: 3329 minor 14 major
  0% 889/com.android.systemui: 0% user + 0% kernel / faults: 431 minor
  0% 2710/com.google.android.gms.unstable: 0% user + 0% kernel / faults: 458 minor 726 major
  0% 3737/artd: 0% user + 0% kernel / faults: 2300 minor
  0% 65/kswapd0: 0% user + 0% kernel
  0% 67/kworker/u17:0-blk_crypto_wq: 0% user + 0% kernel
  0% 286/logd: 0% user + 0% kernel / faults: 79 minor 3 major
  0% 337/jbd2/dm-5-8: 0% user + 0% kernel
  0% 1033/com.android.phone: 0% user + 0% kernel / faults: 488 minor 1 major
  0% 2212/com.google.android.googlequicksearchbox:search: 0% user + 0% kernel / faults: 279 minor 469 major
  0% 2249/com.google.android.apps.messaging: 0% user + 0% kernel / faults: 455 minor 1178 major
  0% 2565/com.android.vending: 0% user + 0% kernel / faults: 586 minor 1286 major
  0% 3180/com.google.android.rkpdapp: 0% user + 0% kernel / faults: 169 minor 34 major
  0% 5832/com.google.android.photopicker: 0% user + 0% kernel / faults: 359 minor 262 major
  0% 5910/kworker/0:1-events: 0% user + 0% kernel
 +0% 6206/dumpsys: 0% user + 0% kernel
39% TOTAL: 0.6% user + 0.9% kernel + 37% iowait + 0.5% irq + 0% softirq
```
