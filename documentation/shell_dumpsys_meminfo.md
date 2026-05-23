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
    "free_ram": "1,120,674K (  733,906K cached pss +   203,536K cached kernel +   183,232K free)",
    "lost_ram": "78,613K",
    "total_ram": "2,042,240K (status normal)",
    "used_ram": "1,542,874K (1,162,890K used pss +   379,984K kernel)",
    "zram": "233,184K physical used for   956,064K in swap (1,531,664K total swap)"
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
Uptime: 415039 Realtime: 415039


Total RSS by process:
    292,432K: system (pid 644)
    231,072K: com.google.android.gms.persistent (pid 1246)
    180,224K: com.android.systemui (pid 889)
    140,288K: com.google.android.apps.nexuslauncher (pid 1043 / activities)
    136,720K: com.google.android.gms (pid 1589)
    131,584K: com.google.android.as (pid 1695)
    128,928K: com.google.android.inputmethod.latin (pid 1530)
    116,480K: com.google.android.gms.unstable (pid 2710)
    114,432K: com.google.android.googlequicksearchbox:interactor (pid 1677)
    103,904K: com.android.phone (pid 1033)
    100,480K: com.google.android.apps.messaging (pid 2249)
     91,760K: com.google.android.providers.media.module (pid 1664)
     91,312K: com.google.android.youtube (pid 3376)
     90,144K: com.google.android.googlequicksearchbox:search (pid 2212)
     89,360K: com.android.vending (pid 2565)
     88,528K: com.google.android.as.oss (pid 2188)
     85,056K: com.android.vending:background (pid 3452)
     84,288K: com.android.settings (pid 4567)
     83,280K: com.google.android.settings.intelligence (pid 5247)
     82,944K: com.android.networkstack.process (pid 996)
     82,832K: com.google.android.apps.messaging:rcs (pid 1743)
     81,760K: com.google.android.ext.services (pid 1069)
     81,440K: android.process.acore (pid 1449)
     80,624K: com.google.android.apps.photos (pid 1848)
     80,096K: com.google.android.apps.wellbeing (pid 2126)
     79,088K: com.google.android.contacts (pid 3205)
     76,256K: com.google.android.apps.wallpaper (pid 2059)
     76,016K: com.google.android.deskclock (pid 1383)
     75,776K: com.google.android.partnersetup (pid 3345)
     75,760K: com.google.android.permissioncontroller (pid 1332)
     74,064K: com.android.carrierdefaultapp (pid 5404)
     74,000K: com.google.android.photopicker (pid 5832)
     73,872K: com.android.chrome (pid 3580)
     73,424K: com.google.android.rkpdapp (pid 3180)
     73,344K: android.process.media (pid 1960)
     71,824K: com.android.providers.calendar (pid 3690)
     71,808K: com.google.android.configupdater (pid 3163)
     71,760K: com.android.printspooler (pid 4371)
     71,392K: zygote64 (pid 431)
     70,336K: com.google.android.projection.gearhead:car (pid 3361)
     69,040K: com.android.imsserviceentitlement (pid 3395)
     69,040K: com.android.emulator.multidisplay (pid 1765)
     68,928K: .adservices (pid 3276)
     68,848K: com.android.externalstorage (pid 3441)
     68,032K: com.google.android.healthconnect.controller (pid 3302)
     67,984K: com.android.se (pid 1026)
     67,840K: com.google.android.ondevicepersonalization.services (pid 3317)
     67,600K: com.google.android.packageinstaller (pid 3331)
     67,392K: com.google.android.federatedcompute (pid 3290)
     51,360K: webview_zygote (pid 964)
     19,360K: surfaceflinger (pid 489)
     11,456K: netd (pid 429)
     11,408K: artd (pid 3737)
     10,464K: audioserver (pid 471)
      8,224K: adbd (pid 491)
      8,176K: keystore2 (pid 313)
      7,312K: logd (pid 286)
      7,152K: mediaserver (pid 524)
      7,056K: android.hardware.wifi-service (pid 449)
      6,944K: android.hardware.graphics.composer3-service.ranchu (pid 451)
      6,944K: cameraserver (pid 509)
      6,624K: android.hardware.camera.provider@2.7-service-google (pid 443)
      6,560K: init (pid 1)
      6,544K: media.extractor (pid 522)
      6,160K: media.swcodec (pid 529)
      6,080K: lmkd (pid 287)
      5,920K: android.hardware.media.c2-service-goldfish (pid 445)
      5,792K: statsd (pid 430)
      5,744K: android.hardware.radio-service.ranchu (pid 446)
      5,648K: android.hardware.gnss-service.ranchu (pid 950)
      5,584K: android.hardware.audio.service (pid 442)
      5,472K: storaged (pid 528)
      5,392K: installd (pid 518)
      5,360K: media.metrics (pid 523)
      5,344K: android.hardware.vibrator-service.example (pid 469)
      5,344K: gpuservice (pid 487)
      5,312K: android.system.suspend-service (pid 312)
      5,264K: android.hardware.sensors-service.multihal (pid 447)
      5,200K: android.hardware.bluetooth-service.default (pid 450)
      5,152K: vold (pid 303)
      5,120K: android.hardware.health-service.example (pid 444)
      5,040K: android.hardware.graphics.allocator-service.ranchu (pid 492)
      4,896K: android.hardware.security.keymint-service (pid 316)
      4,896K: android.hardware.camera.provider.ranchu (pid 493)
      4,864K: servicemanager (pid 288)
      4,800K: ip6tables-restore (pid 5581)
      4,784K: android.hardware.usb-service.example (pid 448)
      4,784K: android.hardware.uwb-service (pid 494)
      4,736K: ot-daemon (pid 867)
      4,672K: iptables-restore (pid 5579)
      4,576K: android.hardware.power.stats-service.example (pid 465)
      4,560K: android.hardware.power-service.example (pid 463)
      4,512K: android.hardware.thermal-service.example (pid 467)
      4,432K: gatekeeperd (pid 530)
      4,416K: ueventd (pid 99)\n\n... (truncated,      580 lines total)
```
