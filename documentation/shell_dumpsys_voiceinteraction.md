# `adbjson shell dumpsys voiceinteraction`

## adbjson

**Command:**
```bash
adbjson shell dumpsys voiceinteraction
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "1 supported user": "[0]",
    "mCurUser": "0",
    "mCurUserSupported": "true",
    "mEnableService": "true",
    "mIsHdsRequired": "false",
    "mTemporarilyDisabled": "false"
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys voiceinteraction
```

**Output:**
```
VOICE INTERACTION MANAGER (dumpsys voiceinteraction)
  mEnableService: true
  mTemporarilyDisabled: false
  mCurUser: 0
  mCurUserSupported: true
  mIsHdsRequired: false
  1 supported user: [0]
  Enrolled KeyphraseSoundModels:
  mUser=0
  mComponent=com.google.android.googlequicksearchbox/com.google.android.voiceinteraction.GsaVoiceInteractionService
  Session service=com.google.android.apps.gsa.nga.engine.viss.NgaVoiceInteractionSessionService
  Service info:
    name=com.google.android.voiceinteraction.GsaVoiceInteractionService
    packageName=com.google.android.googlequicksearchbox
    labelRes=0x7f151f42 nonLocalizedLabel=null icon=0x0 banner=0x0
    processName=com.google.android.googlequicksearchbox:interactor
    enabled=true exported=true directBootAware=false
    permission=android.permission.BIND_VOICE_INTERACTION
    flags=0x0
    ApplicationInfo:
      name=com.google.android.apps.gsa.binaries.velvet.app.VelvetMultiprocessRoot_Application
      packageName=com.google.android.googlequicksearchbox
      labelRes=0x7f150063 nonLocalizedLabel=null icon=0x7f110010 banner=0x0
      className=com.google.android.apps.gsa.binaries.velvet.app.VelvetMultiprocessRoot_Application
      processName=com.google.android.googlequicksearchbox
      taskAffinity=null
      uid=10158 flags=0xa0dbbec5 privateFlags=0x8c080518 theme=0x7f160e72
      requiresSmallestWidthDp=0 compatibleWidthLimitDp=0 largestWidthLimitDp=0
      sourceDir=/data/app/~~hlb79vT4o4DcBas6YlfMYg==/com.google.android.googlequicksearchbox-STno5SGMiFkQpBeVXpiIqg==/base.apk
      splitSourceDirs=[/data/app/~~hlb79vT4o4DcBas6YlfMYg==/com.google.android.googlequicksearchbox-STno5SGMiFkQpBeVXpiIqg==/split_config.xxhdpi.apk, /data/app/~~hlb79vT4o4DcBas6YlfMYg==/com.google.android.googlequicksearchbox-STno5SGMiFkQpBeVXpiIqg==/split_lens_ondevice_engine_play_ml_module.apk]
      resourceDirs=[/product/overlay/EmulationPixel10/EmulationPixel10Overlay.apk, /product/overlay/NavigationBarModeGestural/NavigationBarModeGesturalOverlay.apk]
      overlayPaths=[/product/overlay/EmulationPixel10/EmulationPixel10Overlay.apk, /product/overlay/NavigationBarModeGestural/NavigationBarModeGesturalOverlay.apk, /data/resource-cache/com.android.systemui-neutral-weVW.frro, /data/resource-cache/com.android.systemui-accent-cdQ7.frro, /data/resource-cache/com.android.systemui-dynamic-85XO.frro]
      seinfo=default:privapp:targetSdkVersion=37:partition=product
      seinfoUser=:complete
      dataDir=/data/user/0/com.google.android.googlequicksearchbox
      deviceProtectedDataDir=/data/user_de/0/com.google.android.googlequicksearchbox
      credentialProtectedDataDir=/data/user/0/com.google.android.googlequicksearchbox
      sharedLibraryFiles=[/system/framework/org.apache.http.legacy.jar, /system_ext/framework/androidx.window.extensions.jar, /system_ext/framework/androidx.window.sidecar.jar]
      splitClassLoaderNames=[null, null]
      enabled=true minSdkVersion=32 targetSdkVersion=37 versionCode=301740578 targetSandboxVersion=1
      supportsRtl=true
      fullBackupContent=true
      dataExtractionRules=@xml/2132344866
      crossProfile=true
      networkSecurityConfigRes=0x7f1901b6
      category=7
      HiddenApiEnforcementPolicy=2
      usesNonSdkApi=false
      allowsPlaybackCapture=true
      nativeHeapZeroInitialized=0
      localeConfigRes=0x7f1901af
      enableOnBackInvokedCallback=true
      allowCrossUidActivitySwitchFromBelow=true
      mPageSizeAppCompatFlags=0
      isAppLockSupported=false
      isAppLockEnabled=false
      createTimestamp=7082
  Recognition service=com.google.android.voicesearch.serviceapi.GoogleRecognitionService
  Hotword detection service=com.google.android.apps.gsa.hotword.hotworddetectionservice.GsaHotwordDetectionService
  Settings activity=null
  Supports assist=true
  Supports launch from keyguard=true
  mBound=true mService=android.service.voice.IVoiceInteractionService$Stub$Proxy@1759841
  No Hotword detection connection
```
