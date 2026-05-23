# `adbjson shell dumpsys meminfo`

## adbjson

**Command:**
```bash
adbjson shell dumpsys meminfo
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "free_ram": "936,275K (  598,307K cached pss +   241,552K cached kernel +    96,416K free)",
    "lost_ram": "216,382K",
    "total_ram": "2,042,240K (status normal)",
    "used_ram": "1,566,594K (1,206,578K used pss +   360,016K kernel)",
    "zram": "233,760K physical used for   932,592K in swap (1,531,664K total swap)"
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys meminfo
```

**Output:**
```
Applications Memory Usage (in Kilobytes):
Uptime: 115002 Realtime: 115002


Total RSS by process:
    251,760K: system (pid 640)
    172,944K: com.google.android.gms.persistent (pid 1210)
    150,080K: com.android.systemui (pid 899)
    145,696K: com.google.android.apps.messaging (pid 2485)
    116,064K: com.google.android.googlequicksearchbox:interactor (pid 1728)
    115,424K: com.google.android.gms.unstable (pid 2924)
    113,440K: com.google.android.gms (pid 1563)
     98,480K: com.google.android.as (pid 1738)
     94,704K: com.google.android.apps.nexuslauncher (pid 1033 / activities)
     93,200K: com.android.vending:background (pid 4527)
     90,512K: com.android.phone (pid 1023)
     89,424K: com.android.vending (pid 2775)
     89,168K: com.google.android.youtube (pid 3513)
     88,432K: com.google.android.providers.media.module (pid 1710)
     83,264K: com.google.android.apps.messaging:rcs (pid 1765)
     80,608K: com.google.android.apps.photos (pid 1870)
     80,016K: com.google.android.apps.wellbeing (pid 2546)
     79,392K: com.google.android.apps.restore (pid 3939)
     79,328K: com.google.android.inputmethod.latin (pid 1530)
     77,904K: com.android.networkstack.process (pid 983)
     75,904K: com.google.android.contacts (pid 4034)
     74,784K: com.android.settings (pid 4632)
     74,128K: com.google.android.as.oss (pid 2210)
     70,832K: zygote64 (pid 432)
     70,112K: android.process.acore (pid 1374)
     69,808K: com.google.android.bluetooth (pid 990)
     69,072K: com.google.android.partnersetup (pid 3724)
     68,880K: com.google.android.permissioncontroller (pid 1117)
     68,640K: com.google.android.rkpdapp (pid 4386)
     68,544K: com.android.vending:instant_app_installer (pid 4180)
     68,512K: com.google.android.apps.wallpaper (pid 2430)
     67,632K: com.android.chrome (pid 3994)
     67,520K: com.google.android.adservices.api (pid 4103)
     66,864K: com.google.android.ext.services (pid 1039)
     65,872K: android.process.media (pid 2135)
     65,504K: com.android.providers.calendar (pid 4133)
     62,896K: com.android.imsserviceentitlement (pid 4144)
     61,936K: com.android.emulator.multidisplay (pid 1816)
     59,760K: com.android.se (pid 1012)
     51,184K: webview_zygote (pid 976)
     17,728K: surfaceflinger (pid 494)
     10,288K: artd (pid 1447)
     10,112K: netd (pid 430)
      9,424K: keystore2 (pid 313)
      9,376K: media.extractor (pid 536)
      9,184K: mediaserver (pid 538)
      8,416K: audioserver (pid 490)
      8,304K: adbd (pid 496)
      7,104K: wpa_supplicant (pid 884)
      6,848K: cameraserver (pid 525)
      6,816K: logd (pid 284)
      6,320K: android.hardware.graphics.composer3-service.ranchu (pid 454)
      5,984K: android.hardware.media.c2-service-goldfish (pid 444)
      5,952K: media.swcodec (pid 541)
      5,936K: media.metrics (pid 537)
      5,808K: lmkd (pid 287)
      5,808K: statsd (pid 431)
      5,776K: installd (pid 533)
      5,520K: vold (pid 302)
      5,408K: android.hardware.camera.provider@2.7-service-google (pid 439)
      5,296K: android.hardware.radio-service.ranchu (pid 447)
      5,296K: android.hardware.wifi-service (pid 450)
      5,280K: drmserver (pid 510)
      5,024K: gpuservice (pid 492)
      5,008K: storaged (pid 540)
      4,960K: android.hardware.sensors-service.multihal (pid 448)
      4,928K: servicemanager (pid 288)
      4,912K: android.hardware.bluetooth-service.default (pid 451)
      4,832K: init (pid 1)
      4,784K: android.hardware.health-service.example (pid 440)
      4,768K: ueventd (pid 100)
      4,768K: credstore (pid 491)
      4,688K: android.system.suspend-service (pid 312)
      4,624K: ot-daemon (pid 872)
      4,624K: android.hardware.gnss-service.ranchu (pid 952)
      4,608K: android.hardware.usb-service.example (pid 449)
      4,592K: android.hardware.graphics.allocator-service.ranchu (pid 497)
      4,560K: android.hardware.power-service.example (pid 481)
      4,544K: android.hardware.security.keymint-service (pid 314)
      4,432K: prng_seeder (pid 282)
      4,384K: android.hardware.neuralnetworks-service-sample-limited (pid 478)
      4,368K: android.hardware.camera.provider.ranchu (pid 498)
      4,304K: android.hardware.neuralnetworks-service-sample-all (pid 477)
      4,304K: android.hardware.drm-service.widevine-rikers (pid 489)
      4,144K: android.hardware.cas-service.example (pid 467)
      4,144K: android.hardware.biometrics.fingerprint-service.ranchu (pid 844)
      4,128K: gatekeeperd (pid 542)
      4,112K: traced (pid 523)
      4,096K: android.hardware.neuralnetworks-shim-service-sample (pid 479)
      3,984K: pmgd (pid 1586)
      3,952K: android.hardware.audio.service (pid 438)
      3,920K: android.hardware.threadnetwork-service (pid 485)
      3,904K: android.hardware.authsecret-service.example (pid 457)
      3,888K: android.hardware.uwb-service (pid 499)
      3,856K: android.hardware.identity-service.example (pid 455)\n\n... (truncated,      544 lines total)
```
