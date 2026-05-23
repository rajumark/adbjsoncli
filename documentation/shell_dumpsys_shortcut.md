# `adbjson shell dumpsys shortcut`

## adbjson

**Command:**
```bash
adbjson shell dumpsys shortcut
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "#Failures": "0",
    "Backup Allowed": "true",
    "Cached launcher": "com.google.android.apps.nexuslauncher",
    "Calls": "0",
    "Config": "Max icon dim: 252",
    "Icon format": "PNG",
    "Icon quality": "100",
    "IsShadow": "false (installed)",
    "Last app scan FP": "google/sdk_gphone16k_arm64/emu64a16k:17/CP21.260330.005/15181570:user/dev-keys",
    "Last known FG": "418672",
    "Last package update time": "1777904375358",
    "Last reset": "[1779529036263] 2026-05-23 15:07:16",
    "Launcher": "com.android.systemui  Package user: 0  Owner user: 0",
    "Launcher permission check": "count=280, total=2.8ms, avg=0.010ms, max calls/s=256 max dur/s=2.6ms max time=0.7ms",
    "Non-persistent": "user ID:0",
    "Now": "[1779532479683] 2026-05-23 16:04:39  Raw last reset: [1779529036263] 2026-05-23 15:07:16  Last reset: [1779529036263] 2026-05-23 15:07:16  Next reset: [1779615436263] 2026-05-24 15:07:16",
    "Package": "com.android.settings  UID: 1000",
    "Path": "bitmaps/ has 0 files, size=0 (0 B)",
    "Pending saves": "Num=0 Executor=java.util.concurrent.ThreadPoolExecutor@8114f1[Running, pool size = 0, active threads = 0, queued tasks = 0, completed tasks = 3]",
    "Restore from FP": "null",
    "Total bitmap size": "0 (0 B)",
    "User": "0  Known locales: en-US  Last app scan: [1779532068234] 2026-05-23 15:57:48",
    "Version": "37",
    "VoiceInteractionManager": "com.google.android.googlequicksearchbox",
    "asyncPreloadUserDelay": "count=1, total=0.4ms, avg=0.396ms, max calls/s=0 max dur/s=0.0ms max time=0.4ms",
    "checkLauncherActivity": "count=284, total=4.3ms, avg=0.015ms, max calls/s=0 max dur/s=0.0ms max time=0.6ms",
    "checkPackageChanges": "count=1, total=207.6ms, avg=207.582ms, max calls/s=0 max dur/s=0.0ms max time=207.6ms",
    "cleanupDanglingBitmaps": "count=1, total=3.0ms, avg=3.044ms, max calls/s=0 max dur/s=0.0ms max time=3.0ms",
    "getActivity+metadata": "count=39, total=1.2ms, avg=0.031ms, max calls/s=0 max dur/s=0.0ms max time=0.4ms",
    "getApplicationInfo": "count=13, total=24.4ms, avg=1.878ms, max calls/s=0 max dur/s=0.0ms max time=24.0ms",
    "getApplicationResources": "count=48, total=21.5ms, avg=0.448ms, max calls/s=0 max dur/s=0.0ms max time=10.8ms",
    "getDefaultLauncher()": "count=17, total=0.1ms, avg=0.005ms, max calls/s=13 max dur/s=0.1ms max time=0.0ms",
    "getHomeActivities()": "count=1, total=0.0ms, avg=0.020ms, max calls/s=0 max dur/s=0.0ms max time=0.0ms",
    "getInstalledPackages": "count=1, total=3.8ms, avg=3.831ms, max calls/s=0 max dur/s=0.0ms max time=3.8ms",
    "getLauncherActivity": "count=3, total=0.0ms, avg=0.016ms, max calls/s=0 max dur/s=0.0ms max time=0.0ms",
    "getPackageInfo()": "count=0, total=0.0ms, avg=0.000ms, max calls/s=0 max dur/s=0.0ms max time=0.0ms",
    "getPackageInfo(SIG)": "count=51, total=55.7ms, avg=1.092ms, max calls/s=15 max dur/s=55.2ms max time=55.1ms",
    "isActivityEnabled": "count=12, total=0.3ms, avg=0.022ms, max calls/s=0 max dur/s=0.0ms max time=0.1ms",
    "maxShortcutsPerActivity": "15",
    "maxUpdatesPerInterval": "10",
    "packageUpdateCheck": "count=280, total=3.7ms, avg=0.013ms, max calls/s=0 max dur/s=0.0ms max time=0.7ms",
    "resetInterval": "86400000",
    "resourceNameLookup": "count=32, total=22.1ms, avg=0.690ms, max calls/s=0 max dur/s=0.0ms max time=8.0ms",
    "saveDelayMillis": "3000"
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys shortcut
```

**Output:**
```
Now: [1779532479701] 2026-05-23 16:04:39  Raw last reset: [1779529036263] 2026-05-23 15:07:16  Last reset: [1779529036263] 2026-05-23 15:07:16  Next reset: [1779615436263] 2026-05-24 15:07:16

  Config:    Max icon dim: 252
    Icon format: PNG
    Icon quality: 100
    saveDelayMillis: 3000
    resetInterval: 86400000
    maxUpdatesPerInterval: 10
    maxShortcutsPerActivity: 15

  Stats:
    getHomeActivities(): count=1, total=0.0ms, avg=0.020ms, max calls/s=0 max dur/s=0.0ms max time=0.0ms
    Launcher permission check: count=280, total=2.8ms, avg=0.010ms, max calls/s=256 max dur/s=2.6ms max time=0.7ms
    getPackageInfo(): count=0, total=0.0ms, avg=0.000ms, max calls/s=0 max dur/s=0.0ms max time=0.0ms
    getPackageInfo(SIG): count=51, total=55.7ms, avg=1.092ms, max calls/s=15 max dur/s=55.2ms max time=55.1ms
    getApplicationInfo: count=13, total=24.4ms, avg=1.878ms, max calls/s=0 max dur/s=0.0ms max time=24.0ms
    cleanupDanglingBitmaps: count=1, total=3.0ms, avg=3.044ms, max calls/s=0 max dur/s=0.0ms max time=3.0ms
    getActivity+metadata: count=39, total=1.2ms, avg=0.031ms, max calls/s=0 max dur/s=0.0ms max time=0.4ms
    getInstalledPackages: count=1, total=3.8ms, avg=3.831ms, max calls/s=0 max dur/s=0.0ms max time=3.8ms
    checkPackageChanges: count=1, total=207.6ms, avg=207.582ms, max calls/s=0 max dur/s=0.0ms max time=207.6ms
    getApplicationResources: count=48, total=21.5ms, avg=0.448ms, max calls/s=0 max dur/s=0.0ms max time=10.8ms
    resourceNameLookup: count=32, total=22.1ms, avg=0.690ms, max calls/s=0 max dur/s=0.0ms max time=8.0ms
    getLauncherActivity: count=3, total=0.0ms, avg=0.016ms, max calls/s=0 max dur/s=0.0ms max time=0.0ms
    checkLauncherActivity: count=284, total=4.3ms, avg=0.015ms, max calls/s=0 max dur/s=0.0ms max time=0.6ms
    isActivityEnabled: count=12, total=0.3ms, avg=0.022ms, max calls/s=0 max dur/s=0.0ms max time=0.1ms
    packageUpdateCheck: count=280, total=3.7ms, avg=0.013ms, max calls/s=0 max dur/s=0.0ms max time=0.7ms
    asyncPreloadUserDelay: count=1, total=0.4ms, avg=0.396ms, max calls/s=0 max dur/s=0.0ms max time=0.4ms
    getDefaultLauncher(): count=17, total=0.1ms, avg=0.005ms, max calls/s=13 max dur/s=0.1ms max time=0.0ms

  #Failures: 0

  User: 0  Known locales: en-US  Last app scan: [1779532068234] 2026-05-23 15:57:48
      Last app scan FP: google/sdk_gphone16k_arm64/emu64a16k:17/CP21.260330.005/15181570:user/dev-keys
      Restore from FP: null
      Cached launcher: com.google.android.apps.nexuslauncher

      Launcher: android  Package user: 0  Owner user: 0

        PackageInfo:
          IsShadow: false (installed)
          Version: -1
          Last package update time: 0


      Launcher: com.google.android.as  Package user: 0  Owner user: 0

        PackageInfo:
          IsShadow: false (installed)
          Version: -1
          Last package update time: 0


      Launcher: com.google.android.apps.nexuslauncher  Package user: 0  Owner user: 0

        PackageInfo:
          IsShadow: false (installed)
          Version: -1
          Last package update time: 0


      Launcher: com.android.systemui  Package user: 0  Owner user: 0

        PackageInfo:
          IsShadow: false (installed)
          Version: -1
          Last package update time: 0


      Package: com.google.android.youtube  UID: 10182
        Calls: 0
        Last known FG: 0
        Last reset: [1779529036263] 2026-05-23 15:07:16

        PackageInfo:
          IsShadow: false (installed)
          Version: 1561171562
          Backup Allowed: true
          Last package update time: 1778994484925

        Shortcuts:
          ShortcutInfo {id=subscriptions-shortcut, flags=0x1a4 [ImManIc-rStr]
            packageName=com.google.android.youtube
            activity=ComponentInfo{com.google.android.youtube/com.google.android.youtube.app.honeycomb.Shell$HomeActivity}
            shortLabel=Subscriptions, resId=2132020803[subscriptions]
            longLabel=null, resId=0[null]
            disabledMessage=null, resId=0[null]
            disabledReason=[Not disabled]
            categories=null
            persons=null
            icon=null
            rank=2, timestamp=1779532068440
            intents=[Intent { act=com.google.android.youtube.action.open.subscriptions flg=0x1000c000 }/PersistableBundle[{source=shortcut}]]
            extras=null
            iconRes=2131232291[drawable/ic_shortcut_subscriptions], bitmapPath=null, iconUri=null}
          ShortcutInfo {id=shorts-shortcut, flags=0x1a4 [ImManIc-rStr]
            packageName=com.google.android.youtube
            activity=ComponentInfo{com.google.android.youtube/com.google.android.youtube.app.honeycomb.Shell$HomeActivity}
            shortLabel=Shorts, resId=2132020450[shorts]
            longLabel=null, resId=0[null]
            disabledMessage=null, resId=0[null]\n\n... (truncated,      695 lines total)
```
