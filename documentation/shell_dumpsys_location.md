# `adbjson shell dumpsys location`

## adbjson

**Command:**
```bash
adbjson shell dumpsys location
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "05-23 15:57:43.473": "passive provider [u0] enabled",
    "05-23 15:57:43.736": "gps provider [u0] enabled",
    "05-23 15:57:43.737": "network provider +registration 1000/android[GnssService]/B7F6F231 -> Request[PASSIVE, minUpdateInterval=0, hiddenFromAppOps, WorkSource{1000 android}]",
    "05-23 15:57:43.771": "gps provider [u0] enabled",
    "05-23 15:57:43.877": "fused provider [u0] enabled",
    "05-23 15:57:43.878": "[u0] visible",
    "05-23 15:57:47.856": "passive provider request = ProviderRequest[@0]",
    "05-23 15:57:47.870": "passive provider +registration 1000/android[UwbCountryCode]/AAEC14E4 -> Request[PASSIVE, minUpdateInterval=+1h0m0s0ms, minUpdateDistance=5000.0, WorkSource{1000 android}]",
    "05-23 15:57:48.614": "network provider [u0] enabled",
    "05-23 15:57:48.889": "network provider [u0] disabled",
    "05-23 15:57:48.905": "network provider request = ProviderRequest[OFF]",
    "05-23 15:57:49.244": "passive provider +registration 10219/com.google.android.gms[fused_location_provider]/B0B82D52 -> Request[@+24855d3h14m7s806ms BALANCED, minUpdateInterval=0, hiddenFromAppOps, WorkSource{10219 com.google.android.gms}]",
    "05-23 15:57:49.622": "passive provider +registration 10219/com.google.android.gms[network_location_provider]/A397556B -> Request[PASSIVE, minUpdateInterval=0, hiddenFromAppOps, WorkSource{10219 com.google.android.gms}]",
    "05-23 15:57:50.422": "network provider +registration 10146/com.google.android.as/B1E02796 -> Request[@+6h0m0s0ms BALANCED, minUpdateDistance=100000.0, WorkSource{10146 com.google.android.as}]",
    "05-23 16:02:52.075": "gps provider request = ProviderRequest[@0, HIGH_ACCURACY, WorkSource{10219 com.google.android.gms}]",
    "05-23 16:02:56.115": "gps provider received location[1]",
    "05-23 16:02:56.116": "passive provider delivered location[1] to 1000/android[SensorNotificationService]/3C81C501",
    "05-23 16:02:56.147": "fused provider received location[1]",
    "05-23 16:02:56.148": "passive provider delivered location[1] to 10219/com.google.android.gms[network_location_provider]/A397556B",
    "05-23 16:02:56.170": "gps provider request = ProviderRequest[OFF]",
    "1000/android[GnssService]": "min/max interval = passive/passive, total/active/foreground duration = +6m54s607ms/+259ms/+6m54s607ms, locations = 0",
    "1000/android[SensorNotificationService]": "min/max interval = passive/passive, total/active/foreground duration = +6m50s488ms/+6m50s488ms/+6m50s488ms, locations = 1",
    "1000/android[UwbCountryCode]": "min/max interval = passive/passive, total/active/foreground duration = +6m50s474ms/+6m50s474ms/+6m50s474ms, locations = 1",
    "10146/com.google.android.as": "min/max interval = 21600s/21600s, total/active/foreground duration = +6m47s922ms/0/+6m47s921ms, locations = 0",
    "10219/com.google.android.gms[fused_location_provider]": "min/max interval = 0s/0s, total/active/foreground duration = +4s95ms/+4s95ms/+4s95ms, locations = 1",
    "10219/com.google.android.gms[network_location_provider]": "min/max interval = passive/passive, total/active/foreground duration = +6m48s722ms/+6m48s721ms/+6m48s722ms, locations = 2",
    "Amount of time (while on battery) Top 4 Avg CN0 <= 20.0 dB-Hz (min)": "0.0",
    "Amount of time (while on battery) Top 4 Avg CN0 > 20.0 dB-Hz (min)": "0.06806666666666666",
    "Antenna Infos": "null",
    "Capabilities": "[SCHEDULING ACCUMULATED_DELTA_RANGE(unknown) TOTAL_POWER SINGLEBAND_TRACKING_POWER MULTIBAND_TRACKING_POWER SINGLEBAND_ACQUISITION_POWER MULTIBAND_ACQUISITION_POWER OTHER_MODES_POWER]",
    "Energy consumed while on battery (mAh)": "0.0",
    "GNSS Hardware Model Name": "Android Studio Emulator GPS",
    "KPI logging end time": "+6m57s332ms",
    "KPI logging start time": "+2s715ms",
    "Location Setting": "true",
    "Number of CN0 reports": "296",
    "Number of L5 CN0 reports": "0",
    "Number of TTFF reports": "1",
    "Number of location reports": "1",
    "Number of position accuracy reports": "1",
    "Percentage location failure": "0.0",
    "Position accuracy mean (m)": "5.0",
    "Position accuracy standard deviation (m)": "0.0",
    "TTFF mean (sec)": "4.032",
    "TTFF standard deviation (sec)": "0.0",
    "Time on battery (min)": "7.138916666666667",
    "Top 4 Avg CN0 mean (dB-Hz)": "30.0",
    "Top 4 Avg CN0 standard deviation (dB-Hz)": "0.0",
    "Total number of L5 sv status messages processed": "0",
    "Total number of L5 sv status messages processed, where sv is used in fix": "0",
    "Total number of sv status messages processed": "1776",
    "Total number of sv status messages processed, where sv is used in fix": "0",
    "current users": "u[0]",
    "service": "unregistered"
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys location
```

**Output:**
```
Location Manager State:
  User Info:
    current users: u[0]
  Location Settings:
    Location Setting: true
    Location Allow/Deny Packages:
    Throttling Allow Packages:
      com.google.android.gms
    Emergency Bypass Allow Packages:
      com.google.android.dialer[*]
      com.google.android.gms[.thunderbird]
      com.google.android.apps.scone[satellite_esos]
      com.android.phone[telephony]
  Location Providers:
    passive provider:
      service: registered
      listeners:
        10219/com.google.android.gms[fused_location_provider]/B0B82D52 Request[@+24855d3h14m7s806ms BALANCED, minUpdateInterval=0, hiddenFromAppOps, WorkSource{10219 com.google.android.gms}]
        10219/com.google.android.gms[network_location_provider]/A397556B Request[PASSIVE, minUpdateInterval=0, hiddenFromAppOps, WorkSource{10219 com.google.android.gms}]
        1000/android[UwbCountryCode]/AAEC14E4 Request[PASSIVE, minUpdateInterval=+1h0m0s0ms, minUpdateDistance=5000.0, WorkSource{1000 android}]
        1000/android[SensorNotificationService]/3C81C501 Request[PASSIVE, minUpdateInterval=+30m0s0ms, minUpdateDistance=100000.0, WorkSource{1000 android}]
      last location=Location[gps 37.421998,-122.084000 hAcc=5.0 et=+5m15s101ms alt=5.0 vAcc=0.5 vel=0.0 sAcc=0.5 bear=0.0 bAcc=30.0 {Bundle[{satellites=0, maxCn0=0, meanCn0=0}]}]
      enabled=true
      allowed=true
      identity=1000/android[LocationService]
      properties=ProviderProperties[powerUsage=Low, accuracy=Fine]
    network provider:
      service: ProviderRequest[OFF]
      listeners:
        1000/android[GnssService]/B7F6F231 Request[PASSIVE, minUpdateInterval=0, hiddenFromAppOps, WorkSource{1000 android}] (inactive)
        10146/com.google.android.as/B1E02796 Request[@+6h0m0s0ms BALANCED, minUpdateDistance=100000.0, WorkSource{10146 com.google.android.as}] (inactive)
      last location=null
      enabled=false
      allowed=false
      identity=10219/com.google.android.gms[network_location_provider]
      extra attribution tags={awareness_provider, activity_recognition_provider, network_location_provider, gnss_location_provider, network_location_calibration, current_semantic_location, fused_location_provider, wearable_flp_shim, geofencer_provider}
      properties=ProviderProperties[powerUsage=Medium, accuracy=Fine, requires=network, supports=[altitude]]
      target service=10219/com.google.android.gms/com.google.android.location.network.NetworkLocationService@2
      connected=true
    fused provider:
      service: ProviderRequest[OFF]
      last location=Location[fused 37.421998,-122.084000 hAcc=5.0 et=+5m15s101ms alt=5.0 vAcc=0.5 vel=0.0 sAcc=0.5]
      enabled=true
      allowed=true
      identity=10219/com.google.android.gms[fused_location_provider]
      extra attribution tags={awareness_provider, activity_recognition_provider, network_location_provider, gnss_location_provider, network_location_calibration, current_semantic_location, fused_location_provider, wearable_flp_shim, geofencer_provider}
      properties=ProviderProperties[powerUsage=Low, accuracy=Fine, supports=[bearing,speed,altitude]]
      target service=10219/com.google.android.gms/com.google.android.location.fused.FusedLocationService@1
      connected=true
    gps provider:
      service: ProviderRequest[OFF]
      last location=Location[gps 37.421998,-122.084000 hAcc=5.0 et=+5m15s101ms alt=5.0 vAcc=0.5 vel=0.0 sAcc=0.5 bear=0.0 bAcc=30.0 {Bundle[{satellites=0, maxCn0=0, meanCn0=0}]}]
      enabled=true
      allowed=true
      identity=1000/android[GnssService]
      properties=ProviderProperties[powerUsage=High, accuracy=Fine, requires=satellite, supports=[bearing,speed,altitude]]
      mStarted=false   (changed +1m42s191ms ago)
      mBatchingEnabled=true
      mBatchingStarted=false
      mBatchSize=4
      mFixInterval=0
      GNSS_KPI_START
        KPI logging start time: +2s715ms
        KPI logging end time: +6m57s350ms
        Number of location reports: 1
        Percentage location failure: 0.0
        Number of TTFF reports: 1
        TTFF mean (sec): 4.032
        TTFF standard deviation (sec): 0.0
        Number of position accuracy reports: 1
        Position accuracy mean (m): 5.0
        Position accuracy standard deviation (m): 0.0
        Number of CN0 reports: 296
        Top 4 Avg CN0 mean (dB-Hz): 30.0
        Top 4 Avg CN0 standard deviation (dB-Hz): 0.0
        Total number of sv status messages processed: 1776
        Total number of L5 sv status messages processed: 0
        Total number of sv status messages processed, where sv is used in fix: 0
        Total number of L5 sv status messages processed, where sv is used in fix: 0
        Number of L5 CN0 reports: 0
        Used-in-fix constellation types: 
      GNSS_KPI_END
      Power Metrics
        Time on battery (min): 7.139216666666667
        Amount of time (while on battery) Top 4 Avg CN0 > 20.0 dB-Hz (min): 0.06806666666666666
        Amount of time (while on battery) Top 4 Avg CN0 <= 20.0 dB-Hz (min): 0.0
        Energy consumed while on battery (mAh): 0.0
      Hardware Version: 
  Historical Aggregate Location Provider Data:
    passive:
      10219/com.google.android.gms[fused_location_provider]: min/max interval = 9223372036854775s/9223372036854775s, total/active/foreground duration = +6m49s118ms/+6m49s117ms/+6m49s118ms, locations = 2
      1000/android[UwbCountryCode]: min/max interval = passive/passive, total/active/foreground duration = +6m50s492ms/+6m50s492ms/+6m50s492ms, locations = 1
      1000/android[SensorNotificationService]: min/max interval = passive/passive, total/active/foreground duration = +6m50s506ms/+6m50s506ms/+6m50s506ms, locations = 1
      10219/com.google.android.gms[network_location_provider]: min/max interval = passive/passive, total/active/foreground duration = +6m48s740ms/+6m48s739ms/+6m48s740ms, locations = 2
    gps:
      10219/com.google.android.gms[fused_location_provider]: min/max interval = 0s/0s, total/active/foreground duration = +4s95ms/+4s95ms/+4s95ms, locations = 1
    network:
      10146/com.google.android.as: min/max interval = 21600s/21600s, total/active/foreground duration = +6m47s940ms/0/+6m47s939ms, locations = 0
      1000/android[GnssService]: min/max interval = passive/passive, total/active/foreground duration = +6m54s625ms/+259ms/+6m54s625ms, locations = 0
  Historical Aggregate Gnss Measurement Provider Data:
  GNSS Manager:
    Capabilities: [SCHEDULING ACCUMULATED_DELTA_RANGE(unknown) TOTAL_POWER SINGLEBAND_TRACKING_POWER MULTIBAND_TRACKING_POWER SINGLEBAND_ACQUISITION_POWER MULTIBAND_ACQUISITION_POWER OTHER_MODES_POWER]
    GNSS Hardware Model Name: Android Studio Emulator GPS
    Status Provider:
      service: registered
      listeners:
        10219/com.google.android.gms[network_location_provider]/0B24D37C
    Measurements Provider:
      service: unregistered
      last measurements=null
    Navigation Message Provider:
      service: unregistered
    Antenna Info Provider:
      Antenna Infos: null
      service: unregistered
  Geofence Manager:
    service: unregistered
  Event Log:
    05-23 15:57:43.473: passive provider [u0] enabled
    05-23 15:57:43.736: gps provider [u0] enabled
    05-23 15:57:43.737: network provider +registration 1000/android[GnssService]/B7F6F231 -> Request[PASSIVE, minUpdateInterval=0, hiddenFromAppOps, WorkSource{1000 android}]
    05-23 15:57:43.771: passive provider [u0] enabled
    05-23 15:57:43.771: gps provider [u0] enabled
    05-23 15:57:43.877: fused provider [u0] enabled
    05-23 15:57:43.878: [u0] visible
    05-23 15:57:47.856: passive provider +registration 1000/android[SensorNotificationService]/3C81C501 -> Request[PASSIVE, minUpdateInterval=+30m0s0ms, minUpdateDistance=100000.0, WorkSource{1000 android}]
    05-23 15:57:47.856: passive provider request = ProviderRequest[@0]
    05-23 15:57:47.870: passive provider +registration 1000/android[UwbCountryCode]/AAEC14E4 -> Request[PASSIVE, minUpdateInterval=+1h0m0s0ms, minUpdateDistance=5000.0, WorkSource{1000 android}]
    05-23 15:57:48.614: network provider [u0] enabled
    05-23 15:57:48.889: network provider [u0] disabled
    05-23 15:57:48.905: network provider request = ProviderRequest[OFF]
    05-23 15:57:49.244: passive provider +registration 10219/com.google.android.gms[fused_location_provider]/B0B82D52 -> Request[@+24855d3h14m7s806ms BALANCED, minUpdateInterval=0, hiddenFromAppOps, WorkSource{10219 com.google.android.gms}]
    05-23 15:57:49.622: passive provider +registration 10219/com.google.android.gms[network_location_provider]/A397556B -> Request[PASSIVE, minUpdateInterval=0, hiddenFromAppOps, WorkSource{10219 com.google.android.gms}]
    05-23 15:57:50.422: network provider +registration 10146/com.google.android.as/B1E02796 -> Request[@+6h0m0s0ms BALANCED, minUpdateDistance=100000.0, WorkSource{10146 com.google.android.as}]
    05-23 16:02:52.075: gps provider +registration 10219/com.google.android.gms[fused_location_provider]/AA011815 -> Request[@0 HIGH_ACCURACY, hiddenFromAppOps, WorkSource{10219 com.google.android.gms}]
    05-23 16:02:52.075: gps provider request = ProviderRequest[@0, HIGH_ACCURACY, WorkSource{10219 com.google.android.gms}]
    05-23 16:02:56.115: gps provider received location[1]
    05-23 16:02:56.116: gps provider delivered location[1] to 10219/com.google.android.gms[fused_location_provider]/AA011815
    05-23 16:02:56.116: passive provider delivered location[1] to 10219/com.google.android.gms[fused_location_provider]/B0B82D52
    05-23 16:02:56.116: passive provider delivered location[1] to 10219/com.google.android.gms[network_location_provider]/A397556B
    05-23 16:02:56.116: passive provider delivered location[1] to 1000/android[UwbCountryCode]/AAEC14E4
    05-23 16:02:56.116: passive provider delivered location[1] to 1000/android[SensorNotificationService]/3C81C501
    05-23 16:02:56.147: fused provider received location[1]
    05-23 16:02:56.148: passive provider delivered location[1] to 10219/com.google.android.gms[fused_location_provider]/B0B82D52
    05-23 16:02:56.148: passive provider delivered location[1] to 10219/com.google.android.gms[network_location_provider]/A397556B
    05-23 16:02:56.170: gps provider -registration 10219/com.google.android.gms[fused_location_provider]/AA011815
    05-23 16:02:56.170: gps provider request = ProviderRequest[OFF]
```
