# `adbjson shell dumpsys procstats`

## adbjson

**Command:**
```bash
adbjson shell dumpsys procstats
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "(Last Act)": "0.74%",
    "BFgs": "143MB (1115 samples)",
    "BTop": "2.2MB (660 samples)",
    "Bnd Fgs": "0.09%",
    "Bnd Top": "0.01%",
    "Cached": "247MB (376 samples)",
    "Crit": "+11s898ms",
    "External": "0x over 0",
    "External Slow": "1840x over +8s24ms",
    "Fgs": "49KB (566 samples)",
    "Free": "104MB (376 samples)",
    "Frozen": "148MB (1097 samples)",
    "Imp Bg": "0.02%",
    "Imp Fg": "0.03%",
    "ImpBg": "45MB (1898 samples)",
    "ImpFg": "593KB (358 samples)",
    "Internal All Procs (Memory Change)": "0x over 0",
    "Internal All Procs (Polling)": "36x over +3ms",
    "Internal Single": "70x over +39ms",
    "Kernel": "352MB (376 samples)",
    "LastAct": "33MB (517 samples)",
    "Low": "+17s883ms",
    "Mod": "+49s346ms",
    "Native": "310MB (376 samples)",
    "Persist": "181MB (313 samples)",
    "Persistent": "96% (0.00-16MB-37MB/0.00-5.0MB-29MB/86MB-74MB-150MB over 34)",
    "Receiver": "0.00% (0.00-0.00-0.00/0.00-0.00-0.00/87MB-87MB-87MB over 1)",
    "Receivr": "69KB (1965 samples)",
    "SOff/Norm": "+366ms",
    "SOn/Norm": "+2h16m11s210ms",
    "ServRst": "0.89MB (145 samples)",
    "Service": "3.6MB (2173 samples)",
    "Service Rs": "0.01%",
    "TOTAL": "1.8GB",
    "Top": "61MB (51 samples)",
    "Z-Ram": "214MB (376 samples)"
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys procstats
```

**Output:**
```
CURRENT STATS:
  * com.android.systemui / u0a195 / v37:
         TOTAL: 100% (98MB-110MB-119MB/38MB-74MB-81MB/120MB-168MB-177MB over 14)
    Persistent: 100% (98MB-110MB-119MB/38MB-74MB-81MB/120MB-168MB-177MB over 14)
  * com.android.networkstack.process / 1073 / v370399999:
         TOTAL: 100% (14MB-14MB-15MB/7.1MB-9.2MB-10MB/78MB-81MB-83MB over 14)
    Persistent: 100% (14MB-14MB-15MB/7.1MB-9.2MB-10MB/78MB-81MB-83MB over 14)
  * com.android.phone / 1001 / v37:
         TOTAL: 100% (26MB-26MB-27MB/15MB-19MB-20MB/92MB-101MB-103MB over 14)
    Persistent: 100% (26MB-26MB-27MB/15MB-19MB-20MB/92MB-101MB-103MB over 14)
  * com.google.android.apps.nexuslauncher / u0a192 / v907:
         TOTAL: 100% (58MB-64MB-69MB/10MB-24MB-33MB/127MB-124MB-138MB over 14)
           Top: 100% (58MB-64MB-69MB/10MB-24MB-33MB/127MB-124MB-138MB over 14)
  * com.android.se / 1068 / v37:
         TOTAL: 100% (8.7MB-8.8MB-9.8MB/1.7MB-2.1MB-5.2MB/69MB-67MB-69MB over 14)
    Persistent: 100% (8.7MB-8.8MB-9.8MB/1.7MB-2.1MB-5.2MB/69MB-67MB-69MB over 14)
  * com.google.android.ext.services / u0a225 / v309999900:
         TOTAL: 100% (14MB-15MB-16MB/6.0MB-7.6MB-8.8MB/81MB-80MB-81MB over 14)
       Bnd Fgs: 100% (14MB-15MB-16MB/6.0MB-7.6MB-8.8MB/81MB-80MB-81MB over 14)
  * com.google.android.gms.persistent / u0a219 / v261733035:
         TOTAL: 100% (82MB-115MB-129MB/55MB-83MB-96MB/152MB-212MB-231MB over 14)
       Bnd Top: 19% (123MB-126MB-129MB/90MB-93MB-96MB/231MB-228MB-231MB over 7)
       Bnd Fgs: 80% (82MB-104MB-117MB/55MB-74MB-84MB/152MB-195MB-217MB over 7)
           Fgs: 0.05%
        Imp Fg: 0.06%
        Imp Bg: 0.05%
       Service: 0.91%
  * com.google.android.inputmethod.latin / u0a167 / v175754038:
         TOTAL: 99% (60MB-62MB-72MB/40MB-42MB-53MB/135MB-127MB-135MB over 14)
        Imp Bg: 99% (60MB-62MB-72MB/40MB-42MB-53MB/135MB-127MB-135MB over 14)
  * com.google.android.providers.media.module / u0a220 / v37:
         TOTAL: 99% (22MB-23MB-23MB/14MB-15MB-15MB/91MB-90MB-91MB over 14)
    Persistent: 99% (22MB-23MB-23MB/14MB-15MB-15MB/91MB-90MB-91MB over 14)
  * com.android.emulator.multidisplay / 1000 / v37:
         TOTAL: 99% (8.7MB-8.8MB-9.7MB/2.1MB-3.3MB-5.2MB/70MB-69MB-70MB over 14)
    Persistent: 99% (8.7MB-8.8MB-9.7MB/2.1MB-3.3MB-5.2MB/70MB-69MB-70MB over 14)
  * com.google.android.googlequicksearchbox:interactor / u0a158 / v301740578:
         TOTAL: 99% (31MB-43MB-56MB/9.3MB-22MB-26MB/115MB-110MB-115MB over 15)
       Bnd Fgs: 99% (31MB-43MB-56MB/9.3MB-22MB-26MB/115MB-110MB-115MB over 15)
  * com.google.android.apps.messaging:rcs / u0a154 / v308183063:
         TOTAL: 98% (37MB-37MB-38MB/8.1MB-8.6MB-9.1MB/85MB-82MB-85MB over 14)
       Bnd Fgs: 98% (37MB-37MB-38MB/8.1MB-8.6MB-9.1MB/85MB-82MB-85MB over 14)
  * com.google.android.as / u0a146 / v16818899:
         TOTAL: 98% (54MB-71MB-74MB/31MB-46MB-50MB/109MB-125MB-129MB over 14)
       Bnd Fgs: 98% (54MB-71MB-74MB/31MB-46MB-50MB/109MB-125MB-129MB over 14)
           Fgs: 0.02%
  * com.google.android.as.oss / u0a145 / v133583:
         TOTAL: 98% (17MB-21MB-22MB/8.0MB-14MB-16MB/80MB-84MB-87MB over 14)
       Bnd Fgs: 98% (17MB-21MB-22MB/8.0MB-14MB-16MB/80MB-84MB-87MB over 14)
           Fgs: 0.02%
  * com.google.android.googlequicksearchbox:search / u0a158 / v301740578:
         TOTAL: 31% (46MB-46MB-46MB/37MB-37MB-37MB/112MB-112MB-112MB over 1)
       Bnd Fgs: 31% (46MB-46MB-46MB/37MB-37MB-37MB/112MB-112MB-112MB over 1)
  * com.google.android.bluetooth / 1002 / v37:
         TOTAL: 26% (19MB-19MB-19MB/13MB-13MB-13MB/86MB-86MB-86MB over 1)
    Persistent: 26% (19MB-19MB-19MB/13MB-13MB-13MB/86MB-86MB-86MB over 1)
       Service: 0.01%
  * com.google.android.gms / u0a219 / v261733035:
         TOTAL: 15%
       Bnd Top: 2.2%
       Bnd Fgs: 2.3%
           Fgs: 0.05%
        Imp Bg: 0.02%
       Service: 10%
    Service Rs: 0.01%
    (Last Act): 1.1%
  * com.google.android.permissioncontroller / u0a208 / v330000000:
         TOTAL: 9.6%
        Imp Fg: 9.6%
       Service: 0.01%
  * com.google.android.apps.messaging / u0a154 / v308183063:
         TOTAL: 4.6%
       Bnd Fgs: 4.5%
           Fgs: 0.05%
        Imp Bg: 0.01%
       Service: 0.01%
      Receiver: 0.02%
    (Last Act): 9.8%
  * android.process.acore / u0a102 / v37:
         TOTAL: 3.9%
       Bnd Fgs: 3.9%
        Imp Bg: 0.01%
    (Last Act): 4.8%
  * com.android.printspooler / u0a138 / v37:
         TOTAL: 3.2%
       Bnd Fgs: 2.4%
        Imp Fg: 0.78%
       Service: 0.01%
  * com.google.android.apps.photos / u0a171 / v51684210:
         TOTAL: 3.1%
       Bnd Top: 1.5%
        Imp Bg: 1.7%
       Service: 0.02%
    (Last Act): 13% (40MB-40MB-40MB/20MB-20MB-20MB/93MB-93MB-93MB over 2)
  * com.google.android.gms.unstable / u0a219 / v261733035:
         TOTAL: 3.1%
       Bnd Top: 1.8%
       Bnd Fgs: 0.11%
        Imp Bg: 0.01%
       Service: 1.2%\n\n... (truncated,      825 lines total)
```
