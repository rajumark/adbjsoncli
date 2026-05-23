# `adbjson shell dumpsys sensorservice`

## adbjson

**Command:**
```bash
adbjson shell dumpsys sensorservice
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "0000000000) Goldfish 3-axis Accelerometer | The Android Open Source Project | ver": "1 | type: android.sensor.accelerometer(1) | perm: n/a | flags: 0x00000050",
    "0x00000001) Goldfish 3-axis Gyroscope | The Android Open Source Project | ver": "1 | type: android.sensor.gyroscope(4) | perm: n/a | flags: 0x00000050",
    "0x00000002) Goldfish 3-axis Magnetic field sensor | The Android Open Source Project | ver": "1 | type: android.sensor.magnetic_field(2) | perm: n/a | flags: 0x00000010",
    "0x00000003) Goldfish Orientation sensor | The Android Open Source Project | ver": "1 | type: android.sensor.orientation(3) | perm: n/a | flags: 0x00000010",
    "0x00000004) Goldfish Ambient Temperature sensor | The Android Open Source Project | ver": "1 | type: android.sensor.ambient_temperature(13) | perm: n/a | flags: 0x00000012",
    "0x00000005) Goldfish Proximity sensor | The Android Open Source Project | ver": "1 | type: android.sensor.proximity(8) | perm: n/a | flags: 0x00000013",
    "0x00000006) Goldfish Light sensor     | The Android Open Source Project | ver": "1 | type: android.sensor.light(5) | perm: n/a | flags: 0x00000012",
    "0x00000007) Goldfish Pressure sensor  | The Android Open Source Project | ver": "1 | type: android.sensor.pressure(6) | perm: n/a | flags: 0x00000010",
    "0x00000008) Goldfish Humidity sensor  | The Android Open Source Project | ver": "1 | type: android.sensor.relative_humidity(12) | perm: n/a | flags: 0x00000012",
    "0x00000009) Goldfish 3-axis Magnetic field sensor (uncalibrated) | The Android Open Source Project | ver": "1 | type: android.sensor.magnetic_field_uncalibrated(14) | perm: n/a | flags: 0x00000010",
    "0x0000000a) Goldfish 3-axis Gyroscope (uncalibrated) | The Android Open Source Project | ver": "1 | type: android.sensor.gyroscope_uncalibrated(16) | perm: n/a | flags: 0x00000010",
    "0x00000011) Goldfish 3-axis Accelerometer Uncalibrated | The Android Open Source Project | ver": "1 | type: android.sensor.accelerometer_uncalibrated(35) | perm: n/a | flags: 0x00000050",
    "0x5f636779) Corrected Gyroscope Sensor | AOSP            | ver": "1 | type: android.sensor.gyroscope(4) | perm: n/a | flags: 0x00000000",
    "0x5f676172) Game Rotation Vector Sensor | AOSP            | ver": "3 | type: android.sensor.game_rotation_vector(15) | perm: n/a | flags: 0x00000000",
    "0x5f676273) Gyroscope Bias (debug)    | AOSP            | ver": "1 | type: android.sensor.accelerometer(1) | perm: n/a | flags: 0x00000000",
    "0x5f67656f) GeoMag Rotation Vector Sensor | AOSP            | ver": "3 | type: android.sensor.geomagnetic_rotation_vector(20) | perm: n/a | flags: 0x00000000",
    "0x5f677276) Gravity Sensor            | AOSP            | ver": "3 | type: android.sensor.gravity(9) | perm: n/a | flags: 0x00000000",
    "0x5f6c696e) Linear Acceleration Sensor | AOSP            | ver": "3 | type: android.sensor.linear_acceleration(10) | perm: n/a | flags: 0x00000000",
    "0x5f726f76) Rotation Vector Sensor    | AOSP            | ver": "3 | type: android.sensor.rotation_vector(11) | perm: n/a | flags: 0x00000000",
    "0x5f797072) Orientation Sensor        | AOSP            | ver": "1 | type: android.sensor.orientation(3) | perm: n/a | flags: 0x00000000",
    "15:57:44 + 0x00000000 pid=  644 uid= 1000 samplingPeriod=   66667us batchingPeriod=  100000us result=OK (sensor, package)": "(Goldfish 3-axis Accelerometer, com.android.server.wm.WindowOrientationListener$AccelSensorJudge)",
    "15:57:46 + 0x00000000 pid=  644 uid= 1000 samplingPeriod=  200000us batchingPeriod= 2000000us result=OK (sensor, package)": "(Goldfish 3-axis Accelerometer, com.android.server.power.FaceDownDetector)",
    "15:57:52 + 0x5f6c696e pid= 2710 uid=10219 samplingPeriod=   20000us batchingPeriod=       0us result=OK (sensor, package)": "(Linear Acceleration Sensor , com.google.ccc.abuse.droidguard.events.b)",
    "15:57:52 + 0x5f726f76 pid= 2710 uid=10219 samplingPeriod=   20000us batchingPeriod=       0us result=OK (sensor, package)": "(Rotation Vector Sensor     , com.google.ccc.abuse.droidguard.events.b)",
    "15:57:53 + 0x5f6c696e pid= 2710 uid=10219 samplingPeriod=   20000us batchingPeriod=       0us result=OK (sensor, package)": "(Linear Acceleration Sensor , com.google.ccc.abuse.droidguard.events.b)",
    "15:57:53 + 0x5f726f76 pid= 2710 uid=10219 samplingPeriod=   20000us batchingPeriod=       0us result=OK (sensor, package)": "(Rotation Vector Sensor     , com.google.ccc.abuse.droidguard.events.b)",
    "15:57:53 - 0x5f6c696e pid= 2710 uid=10219 samplingPeriod=   N/A  us batchingPeriod=   N/A  us result=OK (sensor, package)": "(Linear Acceleration Sensor , com.google.ccc.abuse.droidguard.events.b)",
    "15:57:53 - 0x5f726f76 pid= 2710 uid=10219 samplingPeriod=   N/A  us batchingPeriod=   N/A  us result=OK (sensor, package)": "(Rotation Vector Sensor     , com.google.ccc.abuse.droidguard.events.b)",
    "15:57:54 - 0x5f6c696e pid= 2710 uid=10219 samplingPeriod=   N/A  us batchingPeriod=   N/A  us result=OK (sensor, package)": "(Linear Acceleration Sensor , com.google.ccc.abuse.droidguard.events.b)",
    "15:57:54 - 0x5f726f76 pid= 2710 uid=10219 samplingPeriod=   N/A  us batchingPeriod=   N/A  us result=OK (sensor, package)": "(Rotation Vector Sensor     , com.google.ccc.abuse.droidguard.events.b)",
    "15:58:24 + 0x5f6c696e pid= 2710 uid=10219 samplingPeriod=   20000us batchingPeriod=       0us result=OK (sensor, package)": "(Linear Acceleration Sensor , com.google.ccc.abuse.droidguard.events.b)",
    "15:58:24 + 0x5f726f76 pid= 2710 uid=10219 samplingPeriod=   20000us batchingPeriod=       0us result=OK (sensor, package)": "(Rotation Vector Sensor     , com.google.ccc.abuse.droidguard.events.b)",
    "15:58:24 - 0x5f6c696e pid= 2710 uid=10219 samplingPeriod=   N/A  us batchingPeriod=   N/A  us result=OK (sensor, package)": "(Linear Acceleration Sensor , com.google.ccc.abuse.droidguard.events.b)",
    "15:58:24 - 0x5f726f76 pid= 2710 uid=10219 samplingPeriod=   N/A  us batchingPeriod=   N/A  us result=OK (sensor, package)": "(Rotation Vector Sensor     , com.google.ccc.abuse.droidguard.events.b)",
    "16:02:52 + 0x00000007 pid= 1246 uid=10219 samplingPeriod=  100000us batchingPeriod=       0us result=OK (sensor, package)": "(Goldfish Pressure sensor   , com.google.android.location.fused.PressureProvider)",
    "16:02:56 + 0x00000000 pid= 1246 uid=10219 samplingPeriod=   20000us batchingPeriod=       0us result=OK (sensor, package)": "(Goldfish 3-axis Accelerometer, com.google.android.location.fused.StepDetector)",
    "16:02:56 + 0x00000000 pid= 1246 uid=10219 samplingPeriod=  200000us batchingPeriod=       0us result=OK (sensor, package)": "(Goldfish 3-axis Accelerometer, unknown_package_pid_1246)",
    "16:02:56 + 0x00000002 pid= 1246 uid=10219 samplingPeriod=  200000us batchingPeriod=       0us result=OK (sensor, package)": "(Goldfish 3-axis Magnetic field sensor, unknown_package_pid_1246)",
    "16:02:56 + 0x00000009 pid= 1246 uid=10219 samplingPeriod=  200000us batchingPeriod=       0us result=OK (sensor, package)": "(Goldfish 3-axis Magnetic field sensor (uncalibrated), unknown_package_pid_1246)",
    "16:02:56 + 0x0000000a pid= 1246 uid=10219 samplingPeriod=  200000us batchingPeriod=       0us result=OK (sensor, package)": "(Goldfish 3-axis Gyroscope (uncalibrated), unknown_package_pid_1246)",
    "16:02:56 + 0x5f6c696e pid= 2710 uid=10219 samplingPeriod=   20000us batchingPeriod=       0us result=OK (sensor, package)": "(Linear Acceleration Sensor , com.google.ccc.abuse.droidguard.events.b)",
    "16:02:56 + 0x5f726f76 pid= 2710 uid=10219 samplingPeriod=   20000us batchingPeriod=       0us result=OK (sensor, package)": "(Rotation Vector Sensor     , com.google.ccc.abuse.droidguard.events.b)",
    "16:02:56 - 0x00000000 pid= 1246 uid=10219 samplingPeriod=   N/A  us batchingPeriod=   N/A  us result=OK (sensor, package)": "(Goldfish 3-axis Accelerometer, com.google.android.location.fused.StepDetector)",
    "16:02:56 - 0x00000002 pid= 1246 uid=10219 samplingPeriod=   N/A  us batchingPeriod=   N/A  us result=OK (sensor, package)": "(Goldfish 3-axis Magnetic field sensor, unknown_package_pid_1246)",
    "16:02:56 - 0x00000007 pid= 1246 uid=10219 samplingPeriod=   N/A  us batchingPeriod=   N/A  us result=OK (sensor, package)": "(Goldfish Pressure sensor   , com.google.android.location.fused.PressureProvider)",
    "16:02:56 - 0x00000009 pid= 1246 uid=10219 samplingPeriod=   N/A  us batchingPeriod=   N/A  us result=OK (sensor, package)": "(Goldfish 3-axis Magnetic field sensor (uncalibrated), unknown_package_pid_1246)",
    "16:02:56 - 0x0000000a pid= 1246 uid=10219 samplingPeriod=   N/A  us batchingPeriod=   N/A  us result=OK (sensor, package)": "(Goldfish 3-axis Gyroscope (uncalibrated), unknown_package_pid_1246)",
    "16:02:56 - 0x5f6c696e pid= 2710 uid=10219 samplingPeriod=   N/A  us batchingPeriod=   N/A  us result=OK (sensor, package)": "(Linear Acceleration Sensor , com.google.ccc.abuse.droidguard.events.b)",
    "16:02:56 - 0x5f726f76 pid= 2710 uid=10219 samplingPeriod=   N/A  us batchingPeriod=   N/A  us result=OK (sensor, package)": "(Rotation Vector Sensor     , com.google.ccc.abuse.droidguard.events.b)",
    "16:03:10 + 0x00000000 pid= 1246 uid=10219 samplingPeriod=   20000us batchingPeriod=       0us result=OK (sensor, package)": "(Goldfish 3-axis Accelerometer, com.google.android.location.collectionlib.SensorScanner$SensorEventListenerAdapter)",
    "16:03:11 - 0x00000000 pid= 1246 uid=10219 samplingPeriod=   N/A  us batchingPeriod=   N/A  us result=OK (sensor, package)": "(Goldfish 3-axis Accelerometer, com.google.android.location.collectionlib.SensorScanner$SensorEventListenerAdapter)",
    "Captured at": "16:04:38.807",
    "Client 0 [PID": "1246, UID: 10219, IsFrozen: false]",
    "Client 1 [PID": "2710, UID: 10219, IsFrozen: false]",
    "Client 10 [PID": "1677, UID: 10158, IsFrozen: false]",
    "Client 11 [PID": "1589, UID: 10219, IsFrozen: false]",
    "Client 12 [PID": "1043, UID: 10192, IsFrozen: false]",
    "Client 13 [PID": "889, UID: 10195, IsFrozen: false]",
    "Client 14 [PID": "1246, UID: 10219, IsFrozen: false]",
    "Client 15 [PID": "1043, UID: 10192, IsFrozen: false]",
    "Client 2 [PID": "1246, UID: 10219, IsFrozen: false]",
    "Client 3 [PID": "1246, UID: 10219, IsFrozen: false]",
    "Client 4 [PID": "1246, UID: 10219, IsFrozen: false]",
    "Client 5 [PID": "2212, UID: 10158, IsFrozen: false]",
    "Client 6 [PID": "1695, UID: 10146, IsFrozen: false]",
    "Client 7 [PID": "1246, UID: 10219, IsFrozen: false]",
    "Client 8 [PID": "2126, UID: 10159, IsFrozen: false]",
    "Client 9 [PID": "1589, UID: 10219, IsFrozen: false]",
    "Connection Number": "1",
    "Final Duration": "00:00:00 | Active Time: 100% (Suspended for 00:00:00 due to sensor access restriction)",
    "Goldfish 3-axis Accelerometer": "last 50 events",
    "Goldfish 3-axis Accelerometer 0x00000000 | first flush pending": "false | pending flush events 0",
    "Goldfish 3-axis Gyroscope (uncalibrated)": "last 10 events",
    "Goldfish 3-axis Magnetic field sensor": "last 10 events",
    "Goldfish Pressure sensor": "last 10 events",
    "Linear Acceleration Sensor": "last 10 events",
    "Mode": "NORMAL",
    "Operating Mode": "NORMAL",
    "Rotation Vector Sensor": "last 10 events",
    "Sensor Privacy": "disabled",
    "Session Start": "15:57:44 | Active Duration: 00:06:54 | Suspended Duration: 00:00:00",
    "WakeLock Status": "not held",
    "com.android.server.power.FaceDownDetector | WakeLockRefCount 0 | uid 1000 | cache size 0 | max cache size 0 | has sensor access": "true",
    "com.android.server.wm.WindowOrientationListener$AccelSensorJudge | WakeLockRefCount 0 | uid 1000 | cache size 0 | max cache size 0 | has sensor access": "true"
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys sensorservice
```

**Output:**
```
Captured at: 16:04:38.823
Sensor Device:
Total 12 h/w sensors, 12 running 0 disabled clients:
0x00000000) active-count = 2; sampling_period(ms) = {200.0, 66.7}, selected = 66.67 ms; batching_period(ms) = {2000.0, 100.0}, selected = 100.00 ms
Sensor List:
0000000000) Goldfish 3-axis Accelerometer | The Android Open Source Project | ver: 1 | type: android.sensor.accelerometer(1) | perm: n/a | flags: 0x00000050
	continuous | minRate=2.00Hz | maxRate=100.00Hz | no batching | non-wakeUp | data-injection, has-additional-info, 
0x00000001) Goldfish 3-axis Gyroscope | The Android Open Source Project | ver: 1 | type: android.sensor.gyroscope(4) | perm: n/a | flags: 0x00000050
	continuous | minRate=2.00Hz | maxRate=100.00Hz | no batching | non-wakeUp | data-injection, has-additional-info, 
0x00000002) Goldfish 3-axis Magnetic field sensor | The Android Open Source Project | ver: 1 | type: android.sensor.magnetic_field(2) | perm: n/a | flags: 0x00000010
	continuous | minRate=2.00Hz | maxRate=100.00Hz | no batching | non-wakeUp | data-injection, 
0x00000003) Goldfish Orientation sensor | The Android Open Source Project | ver: 1 | type: android.sensor.orientation(3) | perm: n/a | flags: 0x00000010
	continuous | minRate=2.00Hz | maxRate=100.00Hz | no batching | non-wakeUp | data-injection, 
0x00000004) Goldfish Ambient Temperature sensor | The Android Open Source Project | ver: 1 | type: android.sensor.ambient_temperature(13) | perm: n/a | flags: 0x00000012
	on-change | maxDelay=0us | minDelay=0us | no batching | non-wakeUp | data-injection, 
0x00000005) Goldfish Proximity sensor | The Android Open Source Project | ver: 1 | type: android.sensor.proximity(8) | perm: n/a | flags: 0x00000013
	on-change | maxDelay=0us | minDelay=0us | no batching | wakeUp | data-injection, 
0x00000006) Goldfish Light sensor     | The Android Open Source Project | ver: 1 | type: android.sensor.light(5) | perm: n/a | flags: 0x00000012
	on-change | maxDelay=0us | minDelay=0us | no batching | non-wakeUp | data-injection, 
0x00000007) Goldfish Pressure sensor  | The Android Open Source Project | ver: 1 | type: android.sensor.pressure(6) | perm: n/a | flags: 0x00000010
	continuous | minRate=2.00Hz | maxRate=100.00Hz | no batching | non-wakeUp | data-injection, 
0x00000008) Goldfish Humidity sensor  | The Android Open Source Project | ver: 1 | type: android.sensor.relative_humidity(12) | perm: n/a | flags: 0x00000012
	on-change | maxDelay=0us | minDelay=0us | no batching | non-wakeUp | data-injection, 
0x00000009) Goldfish 3-axis Magnetic field sensor (uncalibrated) | The Android Open Source Project | ver: 1 | type: android.sensor.magnetic_field_uncalibrated(14) | perm: n/a | flags: 0x00000010
	continuous | minRate=2.00Hz | maxRate=100.00Hz | no batching | non-wakeUp | data-injection, 
0x0000000a) Goldfish 3-axis Gyroscope (uncalibrated) | The Android Open Source Project | ver: 1 | type: android.sensor.gyroscope_uncalibrated(16) | perm: n/a | flags: 0x00000010
	continuous | minRate=2.00Hz | maxRate=100.00Hz | no batching | non-wakeUp | data-injection, 
0x00000011) Goldfish 3-axis Accelerometer Uncalibrated | The Android Open Source Project | ver: 1 | type: android.sensor.accelerometer_uncalibrated(35) | perm: n/a | flags: 0x00000050
	continuous | minRate=2.00Hz | maxRate=100.00Hz | no batching | non-wakeUp | data-injection, has-additional-info, 
0x5f636779) Corrected Gyroscope Sensor | AOSP            | ver: 1 | type: android.sensor.gyroscope(4) | perm: n/a | flags: 0x00000000
	continuous | maxDelay=0us | maxRate=100.00Hz | no batching | non-wakeUp | 
0x5f676172) Game Rotation Vector Sensor | AOSP            | ver: 3 | type: android.sensor.game_rotation_vector(15) | perm: n/a | flags: 0x00000000
	continuous | maxDelay=0us | maxRate=100.00Hz | no batching | non-wakeUp | 
0x5f676273) Gyroscope Bias (debug)    | AOSP            | ver: 1 | type: android.sensor.accelerometer(1) | perm: n/a | flags: 0x00000000
	continuous | maxDelay=0us | maxRate=100.00Hz | no batching | non-wakeUp | 
0x5f67656f) GeoMag Rotation Vector Sensor | AOSP            | ver: 3 | type: android.sensor.geomagnetic_rotation_vector(20) | perm: n/a | flags: 0x00000000
	continuous | maxDelay=0us | maxRate=100.00Hz | no batching | non-wakeUp | 
0x5f677276) Gravity Sensor            | AOSP            | ver: 3 | type: android.sensor.gravity(9) | perm: n/a | flags: 0x00000000
	continuous | maxDelay=0us | maxRate=100.00Hz | no batching | non-wakeUp | 
0x5f6c696e) Linear Acceleration Sensor | AOSP            | ver: 3 | type: android.sensor.linear_acceleration(10) | perm: n/a | flags: 0x00000000
	continuous | maxDelay=0us | maxRate=100.00Hz | no batching | non-wakeUp | 
0x5f726f76) Rotation Vector Sensor    | AOSP            | ver: 3 | type: android.sensor.rotation_vector(11) | perm: n/a | flags: 0x00000000
	continuous | maxDelay=0us | maxRate=100.00Hz | no batching | non-wakeUp | 
0x5f797072) Orientation Sensor        | AOSP            | ver: 1 | type: android.sensor.orientation(3) | perm: n/a | flags: 0x00000000
	continuous | maxDelay=0us | maxRate=100.00Hz | no batching | non-wakeUp | 
Fusion States:
9-axis fusion disabled (0 clients), gyro-rate= 200.00Hz, q=< 0.6773, -1.51434e-07, 0.000178644, 0.735707 > (1), b=< 1.90774e-10, 2.36082e-10, 2.00988e-10 >
game fusion(no mag) disabled (0 clients), gyro-rate= 200.00Hz, q=< 0.573339, -0.360574, -0.391521, 0.62288 > (1), b=< 0, 0, 0 >
geomag fusion (no gyro) disabled (0 clients), gyro-rate= 200.00Hz, q=< 0, 0, 0, 0 > (0), b=< 0, 0, 0 >
Recent Sensor events:
Rotation Vector Sensor: last 10 events
	 1 (ts=13.173068000, wall=15:57:54.185) 0.68, 0.00, 0.00, 0.74, 0.00, 
	 2 (ts=13.193068000, wall=15:57:54.205) 0.68, 0.00, 0.00, 0.74, 0.00, 
	 3 (ts=13.213068000, wall=15:57:54.225) 0.68, 0.00, 0.00, 0.74, 0.00, 
	 4 (ts=13.233068000, wall=15:57:54.245) 0.68, 0.00, 0.00, 0.74, 0.00, 
	 5 (ts=13.253068000, wall=15:57:54.265) 0.68, 0.00, 0.00, 0.74, 0.00, 
	 6 (ts=13.273068000, wall=15:57:54.285) 0.68, 0.00, 0.00, 0.74, 0.00, 
	 7 (ts=13.293068000, wall=15:57:54.305) 0.68, 0.00, 0.00, 0.74, 0.00, 
	 8 (ts=13.313068000, wall=15:57:54.325) 0.68, 0.00, 0.00, 0.74, 0.00, 
	 9 (ts=13.333068000, wall=15:57:54.345) 0.68, 0.00, 0.00, 0.74, 0.00, 
	10 (ts=13.353068000, wall=15:57:54.365) 0.68, 0.00, 0.00, 0.74, 0.00, 
Goldfish 3-axis Gyroscope (uncalibrated): last 10 events
	 1 (ts=43.695350519, wall=15:58:24.708) 0.00, 0.00, 0.00, -0.00, 0.00, -0.00, 
	 2 (ts=43.700350519, wall=15:58:24.713) 0.00, 0.00, 0.00, -0.00, 0.00, -0.00, 
	 3 (ts=43.705350519, wall=15:58:24.717) 0.00, 0.00, 0.00, -0.00, 0.00, -0.00, 
	 4 (ts=43.710350519, wall=15:58:24.722) 0.00, 0.00, 0.00, -0.00, 0.00, -0.00, 
	 5 (ts=43.715350519, wall=15:58:24.728) 0.00, 0.00, 0.00, -0.00, 0.00, -0.00, 
	 6 (ts=43.720350519, wall=15:58:24.733) 0.00, 0.00, 0.00, -0.00, 0.00, -0.00, 
	 7 (ts=315.188202732, wall=16:02:56.200) 0.00, 0.00, 0.00, -0.00, -0.00, -0.00, 
	 8 (ts=315.193202732, wall=16:02:56.205) 0.00, 0.00, 0.00, -0.00, -0.00, -0.00, 
	 9 (ts=315.198202732, wall=16:02:56.210) 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 
	10 (ts=315.203202732, wall=16:02:56.215) 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 
Goldfish Pressure sensor: last 10 events
	 1 (ts=314.171581772, wall=16:02:55.184) 1013.25, 0.00, 0.00, 
	 2 (ts=314.271581772, wall=16:02:55.284) 1013.25, 0.00, 0.00, 
	 3 (ts=314.371581772, wall=16:02:55.384) 1013.25, 0.00, 0.00, 
	 4 (ts=314.471581772, wall=16:02:55.484) 1013.25, 0.00, 0.00, 
	 5 (ts=314.571581772, wall=16:02:55.584) 1013.25, 0.00, 0.00, 
	 6 (ts=314.671581772, wall=16:02:55.684) 1013.25, 0.00, 0.00, 
	 7 (ts=314.771581772, wall=16:02:55.784) 1013.25, 0.00, 0.00, 
	 8 (ts=314.871581772, wall=16:02:55.884) 1013.25, 0.00, 0.00, 
	 9 (ts=314.971581772, wall=16:02:55.984) 1013.25, 0.00, 0.00, 
	10 (ts=315.071581772, wall=16:02:56.084) 1013.25, 0.00, 0.00, 
Linear Acceleration Sensor: last 10 events
	 1 (ts=13.173068000, wall=15:57:54.185) -0.00, 0.00, 0.00, 
	 2 (ts=13.193068000, wall=15:57:54.205) -0.00, 0.00, 0.00, 
	 3 (ts=13.213068000, wall=15:57:54.225) -0.00, 0.00, 0.00, 
	 4 (ts=13.233068000, wall=15:57:54.245) -0.00, 0.00, 0.00, 
	 5 (ts=13.253068000, wall=15:57:54.265) -0.00, 0.00, 0.00, 
	 6 (ts=13.273068000, wall=15:57:54.285) -0.00, 0.00, 0.00, 
	 7 (ts=13.293068000, wall=15:57:54.305) -0.00, 0.00, 0.00, 
	 8 (ts=13.313068000, wall=15:57:54.325) -0.00, 0.00, 0.00, 
	 9 (ts=13.333068000, wall=15:57:54.345) -0.00, 0.00, 0.00, 
	10 (ts=13.353068000, wall=15:57:54.365) -0.00, 0.00, 0.00, 
Goldfish 3-axis Magnetic field sensor: last 10 events
	 1 (ts=13.319572171, wall=15:57:54.332) 0.00, 9.88, -47.75, 
	 2 (ts=13.329572171, wall=15:57:54.342) 0.00, 9.88, -47.75, 
	 3 (ts=13.339572171, wall=15:57:54.352) 0.00, 9.88, -47.75, 
	 4 (ts=13.349572171, wall=15:57:54.362) 0.00, 9.88, -47.75, 
	 5 (ts=13.359572171, wall=15:57:54.372) 0.00, 9.88, -47.75, \n\n... (truncated,      246 lines total)
```
