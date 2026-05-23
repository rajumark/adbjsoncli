# `adbjson shell dumpsys webviewupdate`

## adbjson

**Command:**
```bash
adbjson shell dumpsys webviewupdate
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "Any WebView package installed": "true",
    "Current WebView package (name, version)": "(com.google.android.webview, 148.0.7778.120)",
    "Minimum WebView version code": "763221804",
    "Minimum targetSdkVersion": "33",
    "Number of relros finished": "1",
    "Number of relros started": "1",
    "Preferred WebView package (name, version)": "(com.google.android.webview, 148.0.7778.120)",
    "Valid package com.google.android.webview (versionName": "148.0.7778.120, versionCode: 777812003, targetSdkVersion: 36) is  installed/enabled for all users",
    "WebView package dirty": "false"
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys webviewupdate
```

**Output:**
```
Current WebView Update Service state
  Current WebView package (name, version): (com.google.android.webview, 148.0.7778.120)
  Minimum targetSdkVersion: 33
  Minimum WebView version code: 763221804
  Number of relros started: 1
  Number of relros finished: 1
  WebView package dirty: false
  Any WebView package installed: true
  Preferred WebView package (name, version): (com.google.android.webview, 148.0.7778.120)
  WebView packages:
    Valid package com.google.android.webview (versionName: 148.0.7778.120, versionCode: 777812003, targetSdkVersion: 36) is  installed/enabled for all users
    com.google.android.webview.beta is NOT installed.
    com.google.android.webview.dev is NOT installed.
    com.google.android.webview.canary is NOT installed.
    com.google.android.webview.debug is NOT installed.
    com.android.webview is NOT installed.
```
