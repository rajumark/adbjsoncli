# `adbjson shell dumpsys translation`

## adbjson

**Command:**
```bash
adbjson shell dumpsys translation
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "0": "defaultName=com.google.android.as/com.google.android.apps.miphone.aiai.app.AiAiTranslationService (enabled)",
    "1 supported user": "[0]",
    "Allow instant service": "false",
    "Cached services": "1",
    "Debug": "false Verbose: false",
    "Disabled by UserManager": "false",
    "Name resolver": "defaultName=com.google.android.as/com.google.android.apps.miphone.aiai.app.AiAiTranslationService (enabled)",
    "Package policy flags": "36",
    "Service Label": "Android System Intelligence",
    "Service UID": "10146",
    "Setup complete": "true",
    "Target SDK": "36",
    "User": "0",
    "Users disabled by restriction": "null"
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys translation
```

**Output:**
```
Debug: false Verbose: false
Package policy flags: 36
1 supported user: [0]
Name resolver: FrameworkResourcesServiceNamer: resId=17040021, numberTemps=0, enabledDefaults=0
    0: defaultName=com.google.android.as/com.google.android.apps.miphone.aiai.app.AiAiTranslationService (enabled)

Users disabled by restriction: null
Allow instant service: false
Cached services: 1
Service at 0: 
    User: 0
    Service Label: Android System Intelligence
    Target SDK: 36
    Name resolver: defaultName=com.google.android.as/com.google.android.apps.miphone.aiai.app.AiAiTranslationService (enabled)

    Disabled by UserManager: false
    Setup complete: true
    Service UID: 10146


  No requested UiTranslation Activity.
```
