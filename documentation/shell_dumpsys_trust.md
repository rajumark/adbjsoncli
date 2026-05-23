# `adbjson shell dumpsys trust`

## adbjson

**Command:**
```bash
adbjson shell dumpsys trust
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "User \"Owner\" (id=0, flags=0x4c13) (current)": "trustState=UNTRUSTED, trustManaged=0, deviceLocked=0, isActiveUnlockRunning=0, strongAuthRequired=0x0"
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys trust
```

**Output:**
```
Trust manager state:
 User "Owner" (id=0, flags=0x4c13) (current): trustState=UNTRUSTED, trustManaged=0, deviceLocked=0, isActiveUnlockRunning=0, strongAuthRequired=0x0
   Enabled agents:
   Events:

target service=10219/com.google.android.gms/.chimera.PersistentInternalBoundBrokerService@-2147483648
connected=true
```
