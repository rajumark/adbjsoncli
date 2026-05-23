# `adbjson shell dumpsys power`

## adbjson

**Command:**
```bash
adbjson shell dumpsys power
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "mBatteryLevel": "100",
    "mBatteryLevelLow": "false",
    "mBootCompleted": "true",
    "mDeviceIdleMode": "false",
    "mDeviceIdleTempWhitelist": "[]",
    "mDeviceIdleWhitelist": "[2000, 10104, 10105, 10110, 10156, 10200, 10201, 10202, 10219, 10226]",
    "mDirty": "0x0",
    "mDockState": "0",
    "mDreamsBatteryLevelDrain": "0",
    "mEnhancedDischargePredictionIsPersonalized": "false",
    "mEnhancedDischargeTimeElapsed": "0",
    "mHalAutoSuspendModeEnabled": "false",
    "mHalInteractiveModeEnabled": "true",
    "mHoldingDisplaySuspendBlocker": "true",
    "mHoldingWakeLockSuspendBlocker": "true",
    "mInterceptedPowerKeyForProximity": "false",
    "mIsFaceDown": "false",
    "mIsPowered": "false",
    "mLastEnhancedDischargeTimeUpdatedElapsed": "0",
    "mLastFlipTime": "0",
    "mLastGlobalSleepTimeRealtime": "0 (175051 ms ago)",
    "mLastGlobalWakeTimeRealtime": "0 (175051 ms ago)",
    "mLastInteractivePowerHintTime": "11453 (163598 ms ago)",
    "mLastScreenBrightnessBoostTime": "0 (175051 ms ago)",
    "mLastSleepReason": "application",
    "mLastSleepTime": "0 (175051 ms ago)",
    "mLastWakeTime": "0 (175051 ms ago)",
    "mLightDeviceIdleMode": "false",
    "mLowPowerStandbyActive": "false",
    "mNotifyLongDispatched": "-40s843ms",
    "mNotifyLongNextCheck": "(none)",
    "mNotifyLongScheduled": "+36s481ms",
    "mPlugType": "0",
    "mProximityPositive": "false",
    "mRequestWaitForNegativeProximity": "false",
    "mSandmanScheduled": "false",
    "mScreenBrightnessBoostInProgress": "false",
    "mStayOn": "false",
    "mSystemReady": "true",
    "mUseAutoSuspend": "true",
    "mWakeLockSummary": "0x1",
    "mWakefulness": "Awake",
    "mWakefulnessChanging": "false",
    "no_cached_wake_locks": "true"
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys power
```

**Output:**
```
POWER MANAGER (dumpsys power)

Power Manager State:
  Settings power_manager_constants:
    no_cached_wake_locks=true
  mDirty=0x0
  mWakefulness=Awake
  mWakefulnessChanging=false
  mIsPowered=false
  mPlugType=0
  mBatteryLevel=100
  mDreamsBatteryLevelDrain=0
  mDockState=0
  mStayOn=false
  mProximityPositive=false
  mBootCompleted=true
  mSystemReady=true
  mEnhancedDischargeTimeElapsed=0
  mLastEnhancedDischargeTimeUpdatedElapsed=0
  mEnhancedDischargePredictionIsPersonalized=false
  mUseAutoSuspend=true
  mHalAutoSuspendModeEnabled=false
  mHalInteractiveModeEnabled=true
  mWakeLockSummary=0x1
  mNotifyLongScheduled=+36s455ms
  mNotifyLongDispatched=-40s869ms
  mNotifyLongNextCheck=(none)
  mRequestWaitForNegativeProximity=false
  mInterceptedPowerKeyForProximity=false
  mSandmanScheduled=false
  mBatteryLevelLow=false
  mLightDeviceIdleMode=false
  mDeviceIdleMode=false
  mDeviceIdleWhitelist=[2000, 10104, 10105, 10110, 10156, 10200, 10201, 10202, 10219, 10226]
  mDeviceIdleTempWhitelist=[]
  mLowPowerStandbyActive=false
  mLastWakeTime=0 (175077 ms ago)
  mLastSleepTime=0 (175077 ms ago)
  mLastSleepReason=application
  mLastGlobalWakeTimeRealtime=0 (175077 ms ago)
  mLastGlobalSleepTimeRealtime=0 (175077 ms ago)
  mLastInteractivePowerHintTime=11453 (163624 ms ago)
  mLastScreenBrightnessBoostTime=0 (175077 ms ago)
  mScreenBrightnessBoostInProgress=false
  mHoldingWakeLockSuspendBlocker=true
  mHoldingDisplaySuspendBlocker=true
  mLastFlipTime=0
  mIsFaceDown=false

Settings and Configuration:
  mDecoupleHalAutoSuspendModeFromDisplayConfig=false
  mDecoupleHalInteractiveModeFromDisplayConfig=false
  mWakeUpWhenPluggedOrUnpluggedConfig=false
  mWakeUpWhenPluggedOrUnpluggedInTheaterModeConfig=false
  mTheaterModeEnabled=false
  mKeepDreamingWhenUnplugging=false
  mGlobalForceDisableWakelocks=false
  mGroupsToForceDisableWakelocks={}
  mSuspendWhenScreenOffDueToProximityConfig=false
  mDreamsSupportedConfig=true
  mDreamsEnabledByDefaultConfig=true
  mDreamsActivatedOnSleepByDefaultConfig=false
  mDreamsActivatedOnDockByDefaultConfig=true
  mDreamsActivatedWhilePosturedByDefaultConfig=false
  mDreamsActivatedOnlyWhileWirelessChargingConfig=false
  mDreamsEnabledOnBatteryConfig=false
  mDreamsBatteryLevelMinimumWhenPoweredConfig=-1
  mDreamsBatteryLevelMinimumWhenNotPoweredConfig=15
  mDreamsBatteryLevelDrainCutoffConfig=5
  mDreamsEnabledSetting=true
  mDreamsActivateOnSleepSetting=false
  mDreamsActivateOnDockSetting=true
  mDreamsActivateWhilePosturedSetting=false
  mDreamsOnlyOnWirelessChargingSetting=false
  mDozeAfterScreenOff=true
  mBrightWhenDozingConfig=false
  mAttentiveWarningDurationConfig=30000
  mStayOnWhilePluggedInSetting=1
  mUserActivityTimeoutOverrideFromWindowManager=-1
  mUserInactiveOverrideFromWindowManager=false
  mDozeScreenStateOverrideFromDreamManager=0
  mDrawWakeLockOverrideFromSidekick=false
  mDozeScreenBrightnessOverrideFromDreamManager=NaN
  mUseNormalBrightnessForDoze=false
  mScreenBrightnessMinimum=0.035433073
  mScreenBrightnessMaximum=1.0
  mScreenBrightnessDefault=0.39763778
  mDoubleTapWakeEnabled=false
  mForegroundProfile=0
  mUserId=0
  mWakeUpDelegate=null

Attentive timeout: -1 ms
Sleep timeout: -1 ms
Screen off timeout: 2147483647 ms
Screen dim duration: 7000 ms

UID states (changing=false changed=false):
  UID 1000:   ACTIVE  count=0 state=0
  UID 1001:   ACTIVE  count=0 state=0
  UID 1002:   ACTIVE  count=0 state=0
  UID 1068:   ACTIVE  count=0 state=0
  UID 1073:   ACTIVE  count=0 state=0
  UID u0a102: INACTIVE  count=0 state=19
  UID u0a104: INACTIVE  count=0 state=19
  UID u0a110: INACTIVE  count=0 state=19
  UID u0a145:   ACTIVE  count=0 state=5
  UID u0a146:   ACTIVE  count=0 state=5
  UID u0a148: INACTIVE  count=0 state=19
  UID u0a149: INACTIVE  count=0 state=19
  UID u0a152: INACTIVE  count=0 state=19
  UID u0a153: INACTIVE  count=0 state=19
  UID u0a154:   ACTIVE  count=0 state=5
  UID u0a158:   ACTIVE  count=0 state=5
  UID u0a159: INACTIVE  count=0 state=19
  UID u0a162: INACTIVE  count=0 state=19
  UID u0a165: INACTIVE  count=0 state=19
  UID u0a167:   ACTIVE  count=0 state=7
  UID u0a171: INACTIVE  count=0 state=19
  UID u0a182: INACTIVE  count=0 state=19
  UID u0a192:   ACTIVE  count=0 state=2
  UID u0a195:   ACTIVE  count=0 state=0
  UID u0a198: INACTIVE  count=0 state=19
  UID u0a205: INACTIVE  count=0 state=19
  UID u0a208: INACTIVE  count=0 state=19
  UID u0a219:   ACTIVE  count=1 state=5
  UID u0a220:   ACTIVE  count=0 state=0
  UID u0a225:   ACTIVE  count=0 state=5
  UID u0a226: INACTIVE  count=0 state=19

Looper state:
  Looper (PowerManagerService, tid 52) {d274111}
    (MessageQueue is using DeliQueue implementation)
    Message 0: { when=+24d20h28m33s23ms what=1 target=android.os.Handler async=true heapIndex=1 }
    Message 1: { when=+36s455ms what=4 target=android.os.Handler async=true heapIndex=0 }
    (Total messages: 2, polling=true, quitting=false)

Wake Locks: size=1
  PARTIAL_WAKE_LOCK                 '*gms_scheduler*/com.google.android.gms/.tapandpay.gcmtask.TapAndPayGcmTaskService' ACQ=-1m40s873ms LONG (uid=10219 isFrozen=false isAttributedUidCached=false pid=1210 ws=WorkSource{10219 com.google.android.gms} powerGroupId=0)

Suspend Blockers: size=5
  PowerManagerService.Booting: ref count=0 []
  PowerManagerService.WakeLocks: ref count=1 [unknown: (05-23 15:38:42.592)]
  PowerManagerService.Display: ref count=1 [holding display: (05-23 15:37:30.646)]
  PowerManagerService.Broadcasts: ref count=0 []
  PowerManagerService.WirelessChargerDetector: ref count=0 []

Display Power: com.android.server.power.PowerManagerService$1@8c66905

ScreenTimeoutConstants: 
  mAttentiveTimeoutConfig=-1
  mAttentiveTimeoutSetting=-1
  mMinimumScreenOffTimeoutConfig=10000
  mMaximumScreenDimDurationConfig=7000
  mMaximumScreenDimRatioConfig=0.20000005
  mMaximumScreenOffTimeoutFromDeviceAdmin=9223372036854775807 (enforced=false)
  mScreenOffTimeoutSetting=2147483647
  mSleepTimeoutSetting=-1

Battery saving stats:
  Battery Saver is currently: OFF
    Times full enabled: 0
    Times adaptive enabled: 0
  
  Drain stats:
                     Battery saver OFF                          ON
  NonDoze NonIntr:      0m      0mAh(  0%)      0.0mAh/h          0m      0mAh(  0%)      0.0mAh/h
             Intr:      0m      0mAh(  0%)      0.0mAh/h          0m      0mAh(  0%)      0.0mAh/h
  Deep    NonIntr:      0m      0mAh(  0%)      0.0mAh/h          0m      0mAh(  0%)      0.0mAh/h
             Intr:      0m      0mAh(  0%)      0.0mAh/h          0m      0mAh(  0%)      0.0mAh/h
  Light   NonIntr:      0m      0mAh(  0%)      0.0mAh/h          0m      0mAh(  0%)      0.0mAh/h
             Intr:      0m      0mAh(  0%)      0.0mAh/h          0m      0mAh(  0%)      0.0mAh/h

Battery saver policy (*NOTE* they only apply when battery saver is ON):
  Settings: battery_saver_constants
    value: 
  Settings: (overlay)
    value: 
  DeviceConfig: battery_saver
    location_mode: 
  mAccessibilityEnabled=false
  mAutomotiveProjectionActive=false
  mPolicyLevel=0
  
  Policy 'default full'
    advertise_is_enabled=true
    disable_vibration=false
    disable_animation=false
    defer_full_backup=true
    defer_keyvalue_backup=true
    enable_firewall=true
    enable_datasaver=false
    disable_launch_boost=true
    enable_brightness_adjustment=false
    adjust_brightness_factor=0.5
    location_mode=3
    force_all_apps_standby=true
    force_background_check=true
    disable_optional_sensors=true
    disable_aod=true
    soundtrigger_mode=1
    enable_quick_doze=true
    enable_night_mode=true
  
  Policy 'current full'
    advertise_is_enabled=true
    disable_vibration=false
    disable_animation=false
    defer_full_backup=true
    defer_keyvalue_backup=true
    enable_firewall=true
    enable_datasaver=false
    disable_launch_boost=true
    enable_brightness_adjustment=false
    adjust_brightness_factor=0.5
    location_mode=3
    force_all_apps_standby=true
    force_background_check=true
    disable_optional_sensors=true
    disable_aod=true
    soundtrigger_mode=1
    enable_quick_doze=true
    enable_night_mode=true
  
  Policy 'default adaptive'
    advertise_is_enabled=false
    disable_vibration=false
    disable_animation=false
    defer_full_backup=false
    defer_keyvalue_backup=false
    enable_firewall=false
    enable_datasaver=false
    disable_launch_boost=false
    enable_brightness_adjustment=false
    adjust_brightness_factor=1.0
    location_mode=0
    force_all_apps_standby=false
    force_background_check=false
    disable_optional_sensors=false
    disable_aod=false
    soundtrigger_mode=0
    enable_quick_doze=false
    enable_night_mode=false
  
  Policy 'current adaptive'
    advertise_is_enabled=false
    disable_vibration=false
    disable_animation=false
    defer_full_backup=false
    defer_keyvalue_backup=false
    enable_firewall=false
    enable_datasaver=false
    disable_launch_boost=false
    enable_brightness_adjustment=false
    adjust_brightness_factor=1.0
    location_mode=0
    force_all_apps_standby=false
    force_background_check=false
    disable_optional_sensors=false
    disable_aod=false
    soundtrigger_mode=0
    enable_quick_doze=false
    enable_night_mode=false
  
  Policy 'effective'
    advertise_is_enabled=false
    disable_vibration=false
    disable_animation=false
    defer_full_backup=false
    defer_keyvalue_backup=false
    enable_firewall=false
    enable_datasaver=false
    disable_launch_boost=false
    enable_brightness_adjustment=false
    adjust_brightness_factor=1.0
    location_mode=0
    force_all_apps_standby=false
    force_background_check=false
    disable_optional_sensors=false
    disable_aod=false
    soundtrigger_mode=0
    enable_quick_doze=false
    enable_night_mode=false

Battery saver state machine:
  Enabled=false
    full=false
    adaptive=false
  mState=1
  mLastChangedIntReason=0
  mLastChangedStrReason=null
  mBootCompleted=true
  mSettingsLoaded=true
  mBatteryStatusSet=true
  mIsPowered=false
  mBatteryLevel=100
  mIsBatteryLevelLow=false
  mSettingAutomaticBatterySaver=0
  mSettingBatterySaverEnabled=false
  mSettingBatterySaverEnabledSticky=false
  mSettingBatterySaverStickyAutoDisableEnabled=true
  mSettingBatterySaverStickyAutoDisableThreshold=90
  mSettingBatterySaverTriggerThreshold=0
  mBatterySaverStickyBehaviourDisabled=false
  mBatterySaverTurnedOffNotificationEnabled=true
  mDynamicPowerSavingsDefaultDisableThreshold=80
  mDynamicPowerSavingsDisableThreshold=80
  mDynamicPowerSavingsEnableBatterySaver=false
  mLastAdaptiveBatterySaverChangedExternallyElapsed=0
AttentionDetector:
 mIsSettingEnabled=false
 mMaxExtensionMillis=900000
 mPreDimCheckDurationMillis=2000
 mEffectivePostDimTimeout=0
 mLastUserActivityTime(excludingAttention)=11447
 mAttentionServiceSupported=false
 mRequested=false

Profile power states: size=0
Power Group User Activity:
groupId: 0
userActivitySummary=0x1
mWakeLockSummary=0x1
lastUserActivityTime=11453 (163624 ms ago)
lastUserActivityTimeNoChangeLights=0 (175077 ms ago)
mLastWakeReason=0
mLastSleepReason=14
mDimDuration=-1
mWakefulness=1
mIsDefaultGroupAdjacent=false
mSupportsSandman=true
mDreamManagerAttemptedDozing=false
mScreenOffTimeout=-1

Wireless Charger Detector State:
  mGravitySensor={Sensor name="Gravity Sensor", vendor="AOSP", version=3, type=9, maxRange=19.6133, resolution=2.480159E-4, power=12.7, minDelay=10000}
  mPoweredWirelessly=false
  mAtRest=false
  mRestX=0.0, mRestY=0.0, mRestZ=0.0
  mDetectionInProgress=false
  mDetectionStartTime=0 (never)
  mMustUpdateRestPosition=false
  mTotalSamples=0
  mMovingSamples=0
  mFirstSampleX=0.0, mFirstSampleY=0.0, mFirstSampleZ=0.0
  mLastSampleX=0.0, mLastSampleY=0.0, mLastSampleZ=0.0
Notifier:
Partial Wakelock Log:
Wake Lock Log
  05-23 15:37:51.229 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:35.195 - 1001 (com.android.ons,...) - ACQ *telephony-radio* (partial)
  05-23 15:37:51.207 - 10167 (com.google.android.inputmethod.latin) - ACQ *job*r/#mdd_download_task_unmetered_battery#@androidx.work.systemjobscheduler@com.google.android.inputmethod.latin/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:37:52.728 - 10219 (com.google.android.gms) - ACQ CMWakeLock (partial)
  05-23 15:37:52.719 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:52.741 - 10219 (com.google.android.gms) - REL wake:com.google.android.gms/.chimera.GmsIntentOperationService - 913
  05-23 15:37:52.746 - 10219 (com.google.android.gms) - REL CMWakeLock
  05-23 15:37:52.755 - 10219 (com.google.android.gms) - ACQ UlrDispSvcFastWL (partial)
  05-23 15:37:52.755 - 10219 (com.google.android.gms) - REL UlrDispSvcFastWL
  05-23 15:37:52.755 - 10219 (com.google.android.gms) - ACQ UlrDispSvcFastWL (partial)
  05-23 15:37:52.757 - 10219 (com.google.android.gms) - REL UlrDispSvcFastWL
  05-23 15:37:52.761 - 10219 (com.google.android.gms) - ACQ CallbackRunner (partial)
  05-23 15:37:52.767 - 10219 (com.google.android.gms) - REL CallbackRunner
  05-23 15:37:52.794 - 10219 (com.google.android.gms) - ACQ Icing (partial)
  05-23 15:37:52.812 - 10219 (com.google.android.gms) - ACQ IntentOp:.u.e.InstallationIntentOperation (partial)
  05-23 15:37:52.817 - 10219 (com.google.android.gms) - REL IntentOp:.u.e.InstallationIntentOperation
  05-23 15:37:52.841 - 10219 (com.google.android.gms) - REL Icing
  05-23 15:37:52.854 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:52.858 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:52.860 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:52.863 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:52.888 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:52.890 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:52.914 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:52.923 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:52.946 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:52.946 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:52.949 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:52.961 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:52.962 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:52.963 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:52.963 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:52.967 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:52.967 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:52.970 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:52.974 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:52.974 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:52.976 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:52.980 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:52.985 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:52.985 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:52.987 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:52.991 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:52.991 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:52.992 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:52.999 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.000 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.005 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.009 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.011 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.011 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.023 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.025 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.025 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.038 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.040 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService (partial)
  05-23 15:37:53.046 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.048 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.049 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.073 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService
  05-23 15:37:53.084 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.085 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.098 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.106 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.110 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.241 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.265 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.319 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.367 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.371 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.382 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.382 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.389 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.389 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.392 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.399 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.482 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.485 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.495 - 10219 (com.google.android.gms) - REL Icing
  05-23 15:37:53.509 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.510 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.541 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.542 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.549 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.551 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.577 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.g.HeartbeatAlarm$ConnectionInfoPersistService (partial)
  05-23 15:37:53.585 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.588 - 1000 (System) - ACQ *alarm* (partial)
  05-23 15:37:53.600 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.g.HeartbeatAlarm$ConnectionInfoPersistService
  05-23 15:37:53.603 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.605 - 1000 (System) - REL *alarm*
  05-23 15:37:53.622 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService (partial)
  05-23 15:37:53.629 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.635 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService
  05-23 15:37:53.637 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.658 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.662 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.689 - 1000 (System) - ACQ *alarm* (partial)
  05-23 15:37:53.695 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.741 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.741 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.744 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.745 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.748 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.751 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.757 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.763 - 1000 (System) - REL *alarm*
  05-23 15:37:53.769 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.769 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.773 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.773 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.773 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.775 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.776 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.778 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.779 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.780 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.780 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.788 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.794 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.794 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.799 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.802 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.805 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.805 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.816 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.823 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.828 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.830 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.832 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.842 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.843 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.845 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService (partial)
  05-23 15:37:53.852 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.857 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService
  05-23 15:37:53.857 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.862 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.862 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.865 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService (partial)
  05-23 15:37:53.867 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.867 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService
  05-23 15:37:53.867 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.870 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.870 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.872 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService (partial)
  05-23 15:37:53.878 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.881 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.881 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService
  05-23 15:37:53.881 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.889 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.889 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.894 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.895 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.899 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.900 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.903 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.905 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.913 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.913 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:53.920 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.920 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:53.923 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:53.926 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.s.s.PhenotypeConfigurator (partial)
  05-23 15:37:53.928 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:53.949 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.148 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.149 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:54.160 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:54.160 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:54.163 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.172 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:54.284 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.285 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.294 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.295 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.295 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.310 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.388 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.400 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.402 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.406 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.412 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.412 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.412 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.413 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.426 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.435 - 10152 (com.google.android.partnersetup) - ACQ *job*r/com.google.android.partnersetup/.AppHidingJobService (partial)
  05-23 15:37:54.435 - 10152 (com.google.android.partnersetup) - ACQ *job*r/com.google.android.partnersetup/.AppHidingJobService (partial)
  05-23 15:37:54.436 - 10152 (com.google.android.partnersetup) - REL *job*r/com.google.android.partnersetup/.AppHidingJobService
  05-23 15:37:54.448 - 10152 (com.google.android.partnersetup) - REL *job*r/com.google.android.partnersetup/.AppHidingJobService
  05-23 15:37:54.448 - 10152 (com.google.android.partnersetup) - ACQ *job*r/com.google.android.partnersetup/.ClientIdJobService (partial)
  05-23 15:37:54.448 - 10152 (com.google.android.partnersetup) - ACQ *job*r/com.google.android.partnersetup/.ClientIdJobService (partial)
  05-23 15:37:54.448 - 10152 (com.google.android.partnersetup) - REL *job*r/com.google.android.partnersetup/.ClientIdJobService
  05-23 15:37:54.463 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.463 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.463 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.463 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.475 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.476 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.476 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.479 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.485 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.486 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.486 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.486 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.562 - 10152 (com.google.android.partnersetup) - ACQ *alarm* (partial)
  05-23 15:37:54.563 - 10152 (com.google.android.partnersetup) - REL *alarm*
  05-23 15:37:54.576 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.576 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.577 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.593 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.596 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.596 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.599 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.605 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.641 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.642 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.642 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.647 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.648 - 10152 (com.google.android.partnersetup) - REL *job*r/com.google.android.partnersetup/.ClientIdJobService
  05-23 15:37:54.743 - 10152 (com.google.android.partnersetup) - ACQ *alarm* (partial)
  05-23 15:37:54.743 - 10152 (com.google.android.partnersetup) - REL *alarm*
  05-23 15:37:54.743 - 10152 (com.google.android.partnersetup) - ACQ *job*r/com.google.android.partnersetup/.ClientIdJobService (partial)
  05-23 15:37:54.743 - 10152 (com.google.android.partnersetup) - ACQ *job*r/com.google.android.partnersetup/.ClientIdJobService (partial)
  05-23 15:37:54.745 - 10152 (com.google.android.partnersetup) - REL *job*r/com.google.android.partnersetup/.ClientIdJobService
  05-23 15:37:54.849 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.852 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.852 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:54.855 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:54.896 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:55.007 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:55.031 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:55.041 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:55.041 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:55.045 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService (partial)
  05-23 15:37:55.046 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:55.051 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.s.w.PeopleAndroidUriWipeoutTask (partial)
  05-23 15:37:55.055 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.s.w.PeopleAndroidUriWipeoutTask
  05-23 15:37:55.075 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.m.s.QrlLoggerService (partial)
  05-23 15:37:55.101 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.m.s.QrlLoggerService
  05-23 15:37:55.124 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.c.s.FeatureLoggingTask (partial)
  05-23 15:37:55.131 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:55.153 - 10152 (com.google.android.partnersetup) - REL *job*r/com.google.android.partnersetup/.ClientIdJobService
  05-23 15:37:55.154 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:55.167 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:55.174 - 10219 (com.google.android.gms) - ACQ GCM_WORK (partial)
  05-23 15:37:55.174 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:55.175 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:55.199 - 10219 (com.google.android.gms) - REL GCM_WORK
  05-23 15:37:55.203 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:55.206 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:55.207 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:55.213 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:55.213 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:55.214 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:55.216 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:55.252 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:55.264 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:55.264 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:55.264 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:55.305 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:55.541 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:55.734 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:55.742 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:55.930 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.c.s.FeatureLoggingTask
  05-23 15:37:55.941 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.c.s.NotificationLoggingTask (partial)
  05-23 15:37:55.950 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.c.s.NotificationLoggingTask
  05-23 15:37:55.956 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.a.s.AdservicesStatusService (partial)
  05-23 15:37:56.227 - 10205 (com.google.android.adservices.api) - ACQ *job*r/com.google.android.adservices.api/com.android.adservices.service.measurement.reporting.ImmediateAggregateReportingJobService (partial)
  05-23 15:37:56.227 - 10205 (com.google.android.adservices.api) - ACQ *job*r/com.google.android.adservices.api/com.android.adservices.service.measurement.reporting.ImmediateAggregateReportingJobService (partial)
  05-23 15:37:56.228 - 10205 (com.google.android.adservices.api) - REL *job*r/com.google.android.adservices.api/com.android.adservices.service.measurement.reporting.ImmediateAggregateReportingJobService
  05-23 15:37:56.235 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.a.s.AdservicesStatusService
  05-23 15:37:56.241 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.a.s.LoggingBoundService (partial)
  05-23 15:37:56.244 - 10205 (com.google.android.adservices.api) - REL *job*r/com.google.android.adservices.api/com.android.adservices.service.measurement.reporting.ImmediateAggregateReportingJobService
  05-23 15:37:56.253 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.a.s.LoggingBoundService
  05-23 15:37:56.260 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.s.c.ActionPreferenceCleanupTask (partial)
  05-23 15:37:56.286 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.s.c.ActionPreferenceCleanupTask
  05-23 15:37:56.293 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.d.OdlhDatabaseCleanupJob (partial)
  05-23 15:37:56.327 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.d.OdlhDatabaseCleanupJob
  05-23 15:37:56.331 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.s.PersonalSafetyLoggerService (partial)
  05-23 15:37:56.345 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.s.PersonalSafetyLoggerService
  05-23 15:37:56.349 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/c.g.a.p.s.a.AutoLockLoggerService (partial)
  05-23 15:37:56.381 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/c.g.a.p.s.a.AutoLockLoggerService
  05-23 15:37:56.384 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.c.s.p.PermissionStateLoggingTask (partial)
  05-23 15:37:56.391 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.c.s.p.PermissionStateLoggingTask
  05-23 15:37:56.394 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.SemanticLocationCleanupJob (partial)
  05-23 15:37:56.399 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.SemanticLocationCleanupJob
  05-23 15:37:56.404 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.OdlhTombstonesCleanupJob (partial)
  05-23 15:37:56.412 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.OdlhTombstonesCleanupJob
  05-23 15:37:56.415 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.f.FederatedScheduleService (partial)
  05-23 15:37:56.420 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.f.FederatedScheduleService
  05-23 15:37:56.423 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.c.d.d.s.DatabaseCleanGmsTaskService (partial)
  05-23 15:37:56.444 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.c.d.d.s.DatabaseCleanGmsTaskService
  05-23 15:37:56.448 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.a.o.GpmBiometricsValueMigrationTaskBoundService (partial)
  05-23 15:37:56.456 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.a.o.GpmBiometricsValueMigrationTaskBoundService
  05-23 15:37:56.461 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.c.m.t.PasswordChangesSubscriptionTaskBoundService (partial)
  05-23 15:37:56.463 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.c.m.t.PasswordChangesSubscriptionTaskBoundService
  05-23 15:37:56.470 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.c.m.t.PasswordSharingSubscriptionTaskBoundService (partial)
  05-23 15:37:56.472 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.c.m.t.PasswordSharingSubscriptionTaskBoundService
  05-23 15:37:56.475 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.b.s.n.PhotosBackupMissingPermissionNotificationTask (partial)
  05-23 15:37:56.477 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.b.s.n.PhotosBackupMissingPermissionNotificationTask
  05-23 15:37:56.490 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.f.r.RegionMddMaintenanceService (partial)
  05-23 15:37:56.533 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.f.r.RegionMddMaintenanceService
  05-23 15:37:56.541 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService
  05-23 15:37:56.542 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:56.570 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:56.570 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:56.574 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService (partial)
  05-23 15:37:56.576 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:56.580 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.i.b.IpaGcmTaskService (partial)
  05-23 15:37:56.588 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.i.b.IpaGcmTaskService
  05-23 15:37:56.592 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.l.s.d.DrivingModeLoggerService (partial)
  05-23 15:37:56.634 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.l.s.d.DrivingModeLoggerService
  05-23 15:37:56.635 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService
  05-23 15:37:56.635 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:56.644 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:56.647 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:56.647 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:56.651 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService (partial)
  05-23 15:37:56.652 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:56.654 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService
  05-23 15:37:56.664 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:56.676 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:56.676 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:56.691 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:56.691 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:56.692 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:56.695 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:56.720 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:56.720 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:56.721 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:56.725 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:56.730 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:56.731 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:56.731 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:37:56.732 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:56.733 - 10171 (com.google.android.apps.photos) - ACQ *job*r/com.google.android.apps.photos/com.google.android.libraries.notifications.entrypoints.scheduled.ScheduledTaskService (partial)
  05-23 15:37:56.734 - 10171 (com.google.android.apps.photos) - ACQ *job*r/com.google.android.apps.photos/com.google.android.libraries.notifications.entrypoints.scheduled.ScheduledTaskService (partial)
  05-23 15:37:56.735 - 10171 (com.google.android.apps.photos) - REL *job*r/com.google.android.apps.photos/com.google.android.libraries.notifications.entrypoints.scheduled.ScheduledTaskService
  05-23 15:37:56.740 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:56.757 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:56.762 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:56.765 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:56.769 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:37:56.776 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService (partial)
  05-23 15:37:56.779 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:56.780 - 10171 (com.google.android.apps.photos) - REL *job*r/com.google.android.apps.photos/com.google.android.libraries.notifications.entrypoints.scheduled.ScheduledTaskService
  05-23 15:37:56.783 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService
  05-23 15:37:56.789 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:56.797 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService (partial)
  05-23 15:37:56.798 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:56.799 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService
  05-23 15:37:56.802 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:56.805 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:56.806 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:56.810 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.g.s.s.StorageUpdateTaskBoundService (partial)
  05-23 15:37:56.810 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:56.845 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.g.s.s.StorageUpdateTaskBoundService
  05-23 15:37:56.845 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:56.851 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:56.852 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:56.855 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService (partial)
  05-23 15:37:56.856 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:56.858 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.m.MddDownloadScheduleService (partial)
  05-23 15:37:56.861 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.m.MddDownloadScheduleService
  05-23 15:37:56.862 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService
  05-23 15:37:56.862 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:56.867 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:56.868 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:56.871 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService (partial)
  05-23 15:37:56.872 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:56.877 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.m.s.PeriodicSitrepService (partial)
  05-23 15:37:56.890 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.m.s.PeriodicSitrepService
  05-23 15:37:56.891 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService
  05-23 15:37:56.891 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:56.897 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:56.897 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:37:56.900 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService (partial)
  05-23 15:37:56.900 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:37:56.905 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.d.l.VmPrefetcherTaskService (partial)
  05-23 15:37:56.973 - 10153 (com.android.vending) - ACQ *job*r/com.android.vending/com.google.android.finsky.instantapps.hint.AppPreloadHygieneService (partial)
  05-23 15:37:56.973 - 10153 (com.android.vending) - ACQ *job*r/com.android.vending/com.google.android.finsky.instantapps.hint.AppPreloadHygieneService (partial)
  05-23 15:37:56.973 - 10153 (com.android.vending) - REL *job*r/com.android.vending/com.google.android.finsky.instantapps.hint.AppPreloadHygieneService
  05-23 15:37:57.017 - 10219 (com.google.android.gms) - ACQ IntentOp:.c.b.BackgroundBroadcastReceiverSupport$PersistentReceiverIntentOperation (partial)
  05-23 15:37:57.021 - 10219 (com.google.android.gms) - REL IntentOp:.c.b.BackgroundBroadcastReceiverSupport$PersistentReceiverIntentOperation
  05-23 15:37:57.039 - 10219 (com.google.android.gms) - ACQ IntentOp:.c.b.BackgroundBroadcastReceiverSupport$GmsReceiverIntentOperation (partial)
  05-23 15:37:57.042 - 10153 (com.android.vending) - REL *job*r/com.android.vending/com.google.android.finsky.scheduler.process.mainimpl.PhoneskyJobServiceMain
  05-23 15:37:57.053 - 10219 (com.google.android.gms) - REL IntentOp:.c.b.BackgroundBroadcastReceiverSupport$GmsReceiverIntentOperation
  05-23 15:37:57.112 - 10219 (com.google.android.gms) - ACQ CMWakeLock (partial)
  05-23 15:37:57.114 - 10219 (com.google.android.gms) - ACQ Icing (partial)
  05-23 15:37:57.119 - 10219 (com.google.android.gms) - REL CMWakeLock
  05-23 15:37:57.119 - 10153 (com.android.vending) - REL *job*r/com.android.vending/com.google.android.finsky.instantapps.hint.AppPreloadHygieneService
  05-23 15:37:57.123 - 10219 (com.google.android.gms) - ACQ UlrDispSvcFastWL (partial)
  05-23 15:37:57.123 - 10219 (com.google.android.gms) - REL UlrDispSvcFastWL
  05-23 15:37:57.123 - 10219 (com.google.android.gms) - ACQ UlrDispSvcFastWL (partial)
  05-23 15:37:57.130 - 10219 (com.google.android.gms) - REL UlrDispSvcFastWL
  05-23 15:37:57.131 - 10219 (com.google.android.gms) - ACQ CallbackRunner (partial)
  05-23 15:37:57.134 - 10219 (com.google.android.gms) - REL CallbackRunner
  05-23 15:37:57.138 - 10219 (com.google.android.gms) - ACQ CallbackRunner (partial)
  05-23 15:37:57.139 - 10219 (com.google.android.gms) - REL CallbackRunner
  05-23 15:37:57.147 - 10219 (com.google.android.gms) - REL Icing
  05-23 15:37:57.166 - 10219 (com.google.android.gms) - ACQ Icing (partial)
  05-23 15:37:57.171 - 10219 (com.google.android.gms) - REL Icing
  05-23 15:37:57.180 - 10219 (com.google.android.gms) - ACQ IntentOp:.u.e.InstallationIntentOperation (partial)
  05-23 15:37:57.181 - 10219 (com.google.android.gms) - REL IntentOp:.u.e.InstallationIntentOperation
  05-23 15:37:57.197 - 10152 (com.google.android.partnersetup) - ACQ *alarm* (partial)
  05-23 15:37:57.197 - 10152 (com.google.android.partnersetup) - REL *alarm*
  05-23 15:37:57.197 - 10152 (com.google.android.partnersetup) - ACQ *job*r/com.google.android.partnersetup/.AppHidingJobService (partial)
  05-23 15:37:57.197 - 10152 (com.google.android.partnersetup) - ACQ *job*r/com.google.android.partnersetup/.AppHidingJobService (partial)
  05-23 15:37:57.201 - 10152 (com.google.android.partnersetup) - REL *job*r/com.google.android.partnersetup/.AppHidingJobService
  05-23 15:37:57.209 - 10152 (com.google.android.partnersetup) - ACQ *alarm* (partial)
  05-23 15:37:57.209 - 10152 (com.google.android.partnersetup) - REL *alarm*
  05-23 15:37:57.209 - 10152 (com.google.android.partnersetup) - ACQ *job*r/com.google.android.partnersetup/.ClientIdJobService (partial)
  05-23 15:37:57.209 - 10152 (com.google.android.partnersetup) - ACQ *job*r/com.google.android.partnersetup/.ClientIdJobService (partial)
  05-23 15:37:57.209 - 10152 (com.google.android.partnersetup) - REL *job*r/com.google.android.partnersetup/.ClientIdJobService
  05-23 15:37:57.219 - 10152 (com.google.android.partnersetup) - REL *job*r/com.google.android.partnersetup/.AppHidingJobService
  05-23 15:37:57.318 - 10152 (com.google.android.partnersetup) - REL *job*r/com.google.android.partnersetup/.ClientIdJobService
  05-23 15:37:59.291 - 10153 (com.android.vending) - REL *job*r/com.android.vending/com.google.android.finsky.scheduler.process.mainimpl.PhoneskyJobServiceMain
  05-23 15:37:59.292 - 10153 (com.android.vending) - REL *job*e/com.android.vending/com.google.android.finsky.scheduler.process.mainimpl.PhoneskyJobServiceMain
  05-23 15:38:00.000 - 1000 (System) - ACQ *alarm* (partial)
  05-23 15:38:00.003 - 1000 (System) - REL *alarm*
  05-23 15:38:00.715 - 10219 (com.google.android.gms) - ACQ *alarm* (partial)
  05-23 15:38:00.720 - 10219 (com.google.android.gms) - ACQ GCM_WORK (partial)
  05-23 15:38:00.721 - 10219 (com.google.android.gms) - REL *alarm*
  05-23 15:38:00.721 - 10219 (com.google.android.gms) - REL GCM_WORK
  05-23 15:38:00.884 - 10219 (com.google.android.gms) - ACQ GCM_WORK (partial)
  05-23 15:38:00.895 - 10219 (com.google.android.gms) - ACQ GCM_READ (partial)
  05-23 15:38:00.895 - 10219 (com.google.android.gms) - REL GCM_READ
  05-23 15:38:00.897 - 10219 (com.google.android.gms) - REL GCM_WORK
  05-23 15:38:00.898 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:00.899 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:01.023 - 10219 (com.google.android.gms) - ACQ GCM_READ (partial)
  05-23 15:38:01.023 - 10219 (com.google.android.gms) - ACQ GCM_WORK (partial)
  05-23 15:38:01.023 - 10219 (com.google.android.gms) - REL GCM_READ
  05-23 15:38:01.025 - 10219 (com.google.android.gms) - ACQ GCM_READ (partial)
  05-23 15:38:01.026 - 10219 (com.google.android.gms) - REL GCM_WORK
  05-23 15:38:01.026 - 10219 (com.google.android.gms) - ACQ GCM_WORK (partial)
  05-23 15:38:01.027 - 10219 (com.google.android.gms) - REL GCM_READ
  05-23 15:38:01.028 - 10219 (com.google.android.gms) - REL GCM_WORK
  05-23 15:38:01.028 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:01.029 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:01.031 - 10154 (com.google.android.apps.messaging) - ACQ *alarm* (partial)
  05-23 15:38:01.031 - 10154 (com.google.android.apps.messaging) - REL *alarm*
  05-23 15:38:01.032 - 10154 (com.google.android.apps.messaging) - ACQ *job*r/#TikTokWorker#com.google.android.apps.messaging.shared.receiver.bootcomplete.BootCompleteWorker#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:01.032 - 10154 (com.google.android.apps.messaging) - ACQ *job*r/#TikTokWorker#com.google.android.apps.messaging.shared.receiver.bootcomplete.BootCompleteWorker#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:01.037 - 10154 (com.google.android.apps.messaging) - REL *job*r/#TikTokWorker#com.google.android.apps.messaging.shared.receiver.bootcomplete.BootCompleteWorker#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:01.106 - 10154 (com.google.android.apps.messaging) - ACQ WorkManager:TikTokListenableWorker startWork -> com.google.android.apps.messaging.shared.receiver.bootcomplete.BootCompleteWorker startWork() -> com.google.android.apps.messaging.shared.receiver.bootcomplete.BootCompleteWorker startWork() -> NoAccountWorkerFactory startWork() -> WorkQueueBatchingImpl#queueWorkItems -> WorkQueueBatchingImpl#schedulingDeferred (partial)
  05-23 15:38:01.116 - 10154 (com.google.android.apps.messaging) - ACQ NotificationManagerService:post:com.google.android.apps.messaging (partial)
  05-23 15:38:01.118 - 10154 (com.google.android.apps.messaging) - REL NotificationManagerService:post:com.google.android.apps.messaging
  05-23 15:38:01.150 - 10154 (com.google.android.apps.messaging) - ACQ WorkManager:TikTokListenableWorker startWork -> com.google.android.apps.messaging.shared.receiver.bootcomplete.BootCompleteWorker startWork() -> com.google.android.apps.messaging.shared.receiver.bootcomplete.BootCompleteWorker startWork() -> NoAccountWorkerFactory startWork() -> WorkQueueBatchingImpl#queueWorkItems -> WorkQueueBatchingImpl#schedulingDeferred (partial)
  05-23 15:38:01.159 - 10154 (com.google.android.apps.messaging) - ACQ WorkManager:TikTokListenableWorker startWork -> com.google.android.apps.messaging.shared.receiver.bootcomplete.BootCompleteWorker startWork() -> com.google.android.apps.messaging.shared.receiver.bootcomplete.BootCompleteWorker startWork() -> NoAccountWorkerFactory startWork() -> WorkQueueBatchingImpl#queueWorkItem -> WorkQueueBatchingImpl#queueWorkItems -> WorkQueueBatchingImpl#schedulingDeferred (partial)
  05-23 15:38:01.161 - 10154 (com.google.android.apps.messaging) - REL WorkManager:TikTokListenableWorker startWork -> com.google.android.apps.messaging.shared.receiver.bootcomplete.BootCompleteWorker startWork() -> com.google.android.apps.messaging.shared.receiver.bootcomplete.BootCompleteWorker startWork() -> NoAccountWorkerFactory startWork() -> WorkQueueBatchingImpl#queueWorkItems -> WorkQueueBatchingImpl#schedulingDeferred
  05-23 15:38:01.164 - 10154 (com.google.android.apps.messaging) - REL WorkManager:TikTokListenableWorker startWork -> com.google.android.apps.messaging.shared.receiver.bootcomplete.BootCompleteWorker startWork() -> com.google.android.apps.messaging.shared.receiver.bootcomplete.BootCompleteWorker startWork() -> NoAccountWorkerFactory startWork() -> WorkQueueBatchingImpl#queueWorkItem -> WorkQueueBatchingImpl#queueWorkItems -> WorkQueueBatchingImpl#schedulingDeferred
  05-23 15:38:01.167 - 10154 (com.google.android.apps.messaging) - REL *job*r/#TikTokWorker#com.google.android.apps.messaging.shared.receiver.bootcomplete.BootCompleteWorker#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:01.258 - 10171 (com.google.android.apps.photos) - REL *job*r/com.google.android.apps.photos/com.google.android.libraries.social.async.BackgroundTaskJobService
  05-23 15:38:01.261 - 10154 (com.google.android.apps.messaging) - ACQ NotificationManagerService:post:com.google.android.apps.messaging (partial)
  05-23 15:38:01.263 - 10154 (com.google.android.apps.messaging) - ACQ NotificationManagerService:post:com.google.android.apps.messaging (partial)
  05-23 15:38:01.263 - 10154 (com.google.android.apps.messaging) - ACQ NotificationManagerService:post:com.google.android.apps.messaging (partial)
  05-23 15:38:01.264 - 10154 (com.google.android.apps.messaging) - ACQ NotificationManagerService:post:com.google.android.apps.messaging (partial)
  05-23 15:38:01.277 - 10154 (com.google.android.apps.messaging) - REL WorkManager:TikTokListenableWorker startWork -> com.google.android.apps.messaging.shared.receiver.bootcomplete.BootCompleteWorker startWork() -> com.google.android.apps.messaging.shared.receiver.bootcomplete.BootCompleteWorker startWork() -> NoAccountWorkerFactory startWork() -> WorkQueueBatchingImpl#queueWorkItems -> WorkQueueBatchingImpl#schedulingDeferred
  05-23 15:38:01.304 - 10225 (com.google.android.ext.services) - ACQ *job*r/#ModelDownloadWorker#@androidx.work.systemjobscheduler@com.google.android.ext.services/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:01.304 - 10225 (com.google.android.ext.services) - ACQ *job*r/#ModelDownloadWorker#@androidx.work.systemjobscheduler@com.google.android.ext.services/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:01.305 - 10225 (com.google.android.ext.services) - REL *job*r/#ModelDownloadWorker#@androidx.work.systemjobscheduler@com.google.android.ext.services/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:01.315 - 10225 (com.google.android.ext.services) - REL *job*r/#ModelDownloadWorker#@androidx.work.systemjobscheduler@com.google.android.ext.services/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:01.469 - 10154 (com.google.android.apps.messaging) - REL NotificationManagerService:post:com.google.android.apps.messaging
  05-23 15:38:01.469 - 10154 (com.google.android.apps.messaging) - REL NotificationManagerService:post:com.google.android.apps.messaging
  05-23 15:38:01.469 - 10154 (com.google.android.apps.messaging) - REL NotificationManagerService:post:com.google.android.apps.messaging
  05-23 15:38:01.469 - 10154 (com.google.android.apps.messaging) - REL NotificationManagerService:post:com.google.android.apps.messaging
  05-23 15:38:02.523 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.d.l.VmPrefetcherTaskService
  05-23 15:38:02.529 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.c.t.UpdateAffiliationsTaskBoundService (partial)
  05-23 15:38:02.537 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.c.t.UpdateAffiliationsTaskBoundService
  05-23 15:38:02.540 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.a.a.c.f.o.c.PasskeysCacheUpdateService (partial)
  05-23 15:38:02.546 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.a.a.c.f.o.c.PasskeysCacheUpdateService
  05-23 15:38:02.549 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService
  05-23 15:38:02.551 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:02.559 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:02.559 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:02.563 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.c.CheckinService (partial)
  05-23 15:38:02.566 - 10219 (com.google.android.gms) - ACQ IntentOp:.c.b.BackgroundBroadcastReceiverSupport$PersistentReceiverIntentOperation (partial)
  05-23 15:38:02.566 - 10219 (com.google.android.gms) - ACQ IntentOp:.c.b.BackgroundBroadcastReceiverSupport$PersistentReceiverIntentOperation (partial)
  05-23 15:38:02.567 - 10219 (com.google.android.gms) - REL IntentOp:.c.b.BackgroundBroadcastReceiverSupport$PersistentReceiverIntentOperation
  05-23 15:38:02.567 - 10219 (com.google.android.gms) - REL IntentOp:.c.b.BackgroundBroadcastReceiverSupport$PersistentReceiverIntentOperation
  05-23 15:38:02.574 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:02.580 - 10219 (com.google.android.gms) - ACQ wake:com.google.android.gms/.chimera.GmsIntentOperationService - 913 (partial)
  05-23 15:38:02.581 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.c.CheckinService
  05-23 15:38:02.581 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:02.582 - 10219 (com.google.android.gms) - ACQ Checkin Service (partial)
  05-23 15:38:02.583 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:02.584 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:02.587 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.c.CheckinService (partial)
  05-23 15:38:02.589 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:02.592 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.c.CheckinService
  05-23 15:38:02.592 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:02.596 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:02.597 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:02.608 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.d.DroidGuardFastRefreshGmsTaskBoundService (partial)
  05-23 15:38:02.608 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.056 - 10219 (com.google.android.gms) - ACQ CMWakeLock (partial)
  05-23 15:38:03.067 - 10219 (com.google.android.gms) - REL Checkin Service
  05-23 15:38:03.067 - 10219 (com.google.android.gms) - REL wake:com.google.android.gms/.chimera.GmsIntentOperationService - 913
  05-23 15:38:03.070 - 10219 (com.google.android.gms) - REL CMWakeLock
  05-23 15:38:03.071 - 10219 (com.google.android.gms) - ACQ UlrDispSvcFastWL (partial)
  05-23 15:38:03.071 - 10219 (com.google.android.gms) - REL UlrDispSvcFastWL
  05-23 15:38:03.071 - 10219 (com.google.android.gms) - ACQ UlrDispSvcFastWL (partial)
  05-23 15:38:03.073 - 10219 (com.google.android.gms) - REL UlrDispSvcFastWL
  05-23 15:38:03.076 - 10219 (com.google.android.gms) - ACQ CallbackRunner (partial)
  05-23 15:38:03.077 - 10219 (com.google.android.gms) - REL CallbackRunner
  05-23 15:38:03.080 - 10219 (com.google.android.gms) - ACQ CallbackRunner (partial)
  05-23 15:38:03.080 - 10219 (com.google.android.gms) - REL CallbackRunner
  05-23 15:38:03.093 - 10219 (com.google.android.gms) - ACQ Icing (partial)
  05-23 15:38:03.109 - 10219 (com.google.android.gms) - ACQ IntentOp:.u.e.InstallationIntentOperation (partial)
  05-23 15:38:03.112 - 10219 (com.google.android.gms) - REL Icing
  05-23 15:38:03.112 - 10219 (com.google.android.gms) - REL IntentOp:.u.e.InstallationIntentOperation
  05-23 15:38:03.148 - 1000 (System) - ACQ *alarm* (partial)
  05-23 15:38:03.151 - 10152 (com.google.android.partnersetup) - ACQ *job*r/com.google.android.partnersetup/.AppHidingJobService (partial)
  05-23 15:38:03.151 - 10152 (com.google.android.partnersetup) - ACQ *job*r/com.google.android.partnersetup/.AppHidingJobService (partial)
  05-23 15:38:03.153 - 10152 (com.google.android.partnersetup) - REL *job*r/com.google.android.partnersetup/.AppHidingJobService
  05-23 15:38:03.153 - 10152 (com.google.android.partnersetup) - ACQ *job*r/com.google.android.partnersetup/.ClientIdJobService (partial)
  05-23 15:38:03.153 - 10152 (com.google.android.partnersetup) - ACQ *job*r/com.google.android.partnersetup/.ClientIdJobService (partial)
  05-23 15:38:03.153 - 10152 (com.google.android.partnersetup) - REL *job*r/com.google.android.partnersetup/.ClientIdJobService
  05-23 15:38:03.160 - 1000 (System) - REL *alarm*
  05-23 15:38:03.161 - 10152 (com.google.android.partnersetup) - REL *job*r/com.google.android.partnersetup/.AppHidingJobService
  05-23 15:38:03.262 - 10152 (com.google.android.partnersetup) - REL *job*r/com.google.android.partnersetup/.ClientIdJobService
  05-23 15:38:03.277 - 10171 (com.google.android.apps.photos) - ACQ BackgroundTaskService-UpdateFlagsTask (partial)
  05-23 15:38:03.280 - 10171 (com.google.android.apps.photos) - ACQ *job*r/com.google.android.apps.photos/com.google.android.libraries.social.async.BackgroundTaskJobService (partial)
  05-23 15:38:03.280 - 10171 (com.google.android.apps.photos) - ACQ *job*r/com.google.android.apps.photos/com.google.android.libraries.social.async.BackgroundTaskJobService (partial)
  05-23 15:38:03.281 - 10171 (com.google.android.apps.photos) - REL *job*r/com.google.android.apps.photos/com.google.android.libraries.social.async.BackgroundTaskJobService
  05-23 15:38:03.288 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.d.DroidGuardFastRefreshGmsTaskBoundService
  05-23 15:38:03.288 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.295 - 10171 (com.google.android.apps.photos) - ACQ BackgroundTaskService-UpdateFlagsTask (partial)
  05-23 15:38:03.295 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.296 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.302 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.g.c.SchedulePeriodicTasksService (partial)
  05-23 15:38:03.308 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.325 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.g.c.SchedulePeriodicTasksService
  05-23 15:38:03.326 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.350 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.359 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.384 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService (partial)
  05-23 15:38:03.404 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.421 - 10219 (com.google.android.gms) - ACQ CMWakeLock (partial)
  05-23 15:38:03.432 - 10219 (com.google.android.gms) - REL CMWakeLock
  05-23 15:38:03.447 - 10171 (com.google.android.apps.photos) - REL BackgroundTaskService-UpdateFlagsTask
  05-23 15:38:03.480 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService
  05-23 15:38:03.481 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.496 - 10171 (com.google.android.apps.photos) - REL BackgroundTaskService-UpdateFlagsTask
  05-23 15:38:03.498 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.502 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.534 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService (partial)
  05-23 15:38:03.538 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.540 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.s.s.PhenotypeConfigurator
  05-23 15:38:03.544 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.552 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.552 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.562 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService (partial)
  05-23 15:38:03.572 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.572 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService
  05-23 15:38:03.572 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.573 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService
  05-23 15:38:03.581 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.581 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.607 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService (partial)
  05-23 15:38:03.628 - 10154 (com.google.android.apps.messaging) - ACQ BuglePhenotypeBroadcastReceiver Receive broadcast -> ReceiverDoInBackground -> BuglePhenotypeBroadcastReceiver#doInBackgroundAsync -> PhenotypeHelper#onBuglePhenotypeChanged -> WorkQueueBatchingImpl#queueWorkItem -> WorkQueueBatchingImpl#queueWorkItems -> WorkQueueBatchingImpl#schedulingDeferred (partial)
  05-23 15:38:03.629 - 10154 (com.google.android.apps.messaging) - REL BuglePhenotypeBroadcastReceiver Receive broadcast -> ReceiverDoInBackground -> BuglePhenotypeBroadcastReceiver#doInBackgroundAsync -> PhenotypeHelper#onBuglePhenotypeChanged -> WorkQueueBatchingImpl#queueWorkItem -> WorkQueueBatchingImpl#queueWorkItems -> WorkQueueBatchingImpl#schedulingDeferred
  05-23 15:38:03.629 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService (partial)
  05-23 15:38:03.629 - 10219 (com.google.android.gms) - ACQ IntentOp:.u.e.InstallationIntentOperation (partial)
  05-23 15:38:03.630 - 10219 (com.google.android.gms) - REL IntentOp:.u.e.InstallationIntentOperation
  05-23 15:38:03.661 - 10154 (com.google.android.apps.messaging) - ACQ *job*r/#WorkQueueWorkerShim#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:03.661 - 10154 (com.google.android.apps.messaging) - ACQ *job*r/#WorkQueueWorkerShim#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:03.663 - 10154 (com.google.android.apps.messaging) - REL *job*r/#WorkQueueWorkerShim#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:03.667 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.667 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.667 - 10154 (com.google.android.apps.messaging) - REL *job*r/#WorkQueueWorkerShim#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:03.668 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.670 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService
  05-23 15:38:03.678 - 10219 (com.google.android.gms) - ACQ CMWakeLock (partial)
  05-23 15:38:03.678 - 10219 (com.google.android.gms) - REL CMWakeLock
  05-23 15:38:03.689 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.689 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService
  05-23 15:38:03.712 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.712 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.715 - 10219 (com.google.android.gms) - ACQ CMWakeLock (partial)
  05-23 15:38:03.715 - 10219 (com.google.android.gms) - REL CMWakeLock
  05-23 15:38:03.729 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.u.s.UdcContextInitService (partial)
  05-23 15:38:03.735 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService (partial)
  05-23 15:38:03.738 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.739 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.739 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.755 - 10219 (com.google.android.gms) - ACQ IntentOp:.a.a.AuthInitIntentOperation (partial)
  05-23 15:38:03.756 - 10219 (com.google.android.gms) - REL IntentOp:.a.a.AuthInitIntentOperation
  05-23 15:38:03.758 - 10219 (com.google.android.gms) - ACQ IntentOp:.a.f.FrpUpdateIntentOperation (partial)
  05-23 15:38:03.763 - 10219 (com.google.android.gms) - REL IntentOp:.a.f.FrpUpdateIntentOperation
  05-23 15:38:03.768 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService
  05-23 15:38:03.768 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.773 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.773 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.778 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService (partial)
  05-23 15:38:03.781 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.786 - 10219 (com.google.android.gms) - ACQ CMWakeLock (partial)
  05-23 15:38:03.787 - 10219 (com.google.android.gms) - REL CMWakeLock
  05-23 15:38:03.787 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService
  05-23 15:38:03.788 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.u.s.UdcContextInitService
  05-23 15:38:03.789 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.792 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.793 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.794 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.795 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.796 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService (partial)
  05-23 15:38:03.806 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService (partial)
  05-23 15:38:03.813 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.814 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService
  05-23 15:38:03.815 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService
  05-23 15:38:03.815 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.823 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService (partial)
  05-23 15:38:03.842 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService (partial)
  05-23 15:38:03.844 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.846 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService
  05-23 15:38:03.846 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.852 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.852 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.854 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService (partial)
  05-23 15:38:03.858 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.858 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService
  05-23 15:38:03.858 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.860 - 10219 (com.google.android.gms) - ACQ CMWakeLock (partial)
  05-23 15:38:03.861 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService
  05-23 15:38:03.862 - 10219 (com.google.android.gms) - REL CMWakeLock
  05-23 15:38:03.865 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.865 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.871 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.s.s.PhenotypeConfigurator (partial)
  05-23 15:38:03.876 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.c.u.QosUploaderService (partial)
  05-23 15:38:03.879 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.879 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.880 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.905 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.c.u.QosUploaderService
  05-23 15:38:03.905 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.912 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.913 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:03.913 - 10219 (com.google.android.gms) - ACQ CMWakeLock (partial)
  05-23 15:38:03.917 - 10219 (com.google.android.gms) - REL CMWakeLock
  05-23 15:38:03.918 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService (partial)
  05-23 15:38:03.922 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:03.928 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.l.f.l.FlpSettingsLoggerService (partial)
  05-23 15:38:03.933 - 10219 (com.google.android.gms) - ACQ CMWakeLock (partial)
  05-23 15:38:03.935 - 10219 (com.google.android.gms) - REL CMWakeLock
  05-23 15:38:03.935 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.l.f.l.FlpSettingsLoggerService
  05-23 15:38:03.955 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.ThunderbirdSchedulerService (partial)
  05-23 15:38:03.973 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.ThunderbirdSchedulerService
  05-23 15:38:03.980 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.l.g.GmsCoreLoggerFilesCleanupTask (partial)
  05-23 15:38:03.995 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.l.g.GmsCoreLoggerFilesCleanupTask
  05-23 15:38:04.000 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.c.s.AvailabilityFilesCleanupTask (partial)
  05-23 15:38:04.007 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.c.s.AvailabilityFilesCleanupTask
  05-23 15:38:04.010 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.l.g.e.LatencyMeasurementTaskService (partial)
  05-23 15:38:04.018 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.l.g.e.LatencyMeasurementTaskService
  05-23 15:38:04.024 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.w.AppSpecificPropertyRegistrationTask (partial)
  05-23 15:38:04.029 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.w.AppSpecificPropertyRegistrationTask
  05-23 15:38:04.036 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.s.w.PeopleAndroidUriWipeoutTask (partial)
  05-23 15:38:04.041 - 10219 (com.google.android.gms) - ACQ CMWakeLock (partial)
  05-23 15:38:04.041 - 10219 (com.google.android.gms) - REL CMWakeLock
  05-23 15:38:04.042 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.s.w.PeopleAndroidUriWipeoutTask
  05-23 15:38:04.053 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.m.s.QrlLoggerService (partial)
  05-23 15:38:04.059 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.m.s.QrlLoggerService
  05-23 15:38:04.065 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.c.s.FeatureLoggingTask (partial)
  05-23 15:38:04.148 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.c.s.FeatureLoggingTask
  05-23 15:38:04.151 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.c.s.NotificationLoggingTask (partial)
  05-23 15:38:04.157 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.c.s.NotificationLoggingTask
  05-23 15:38:04.161 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.a.s.AdservicesStatusService (partial)
  05-23 15:38:04.234 - 10205 (com.google.android.adservices.api) - ACQ *job*r/com.google.android.adservices.api/com.android.adservices.service.measurement.reporting.ImmediateAggregateReportingJobService (partial)
  05-23 15:38:04.234 - 10205 (com.google.android.adservices.api) - ACQ *job*r/com.google.android.adservices.api/com.android.adservices.service.measurement.reporting.ImmediateAggregateReportingJobService (partial)
  05-23 15:38:04.234 - 10205 (com.google.android.adservices.api) - REL *job*r/com.google.android.adservices.api/com.android.adservices.service.measurement.reporting.ImmediateAggregateReportingJobService
  05-23 15:38:04.238 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.a.s.AdservicesStatusService
  05-23 15:38:04.242 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.a.s.LoggingBoundService (partial)
  05-23 15:38:04.243 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.a.s.LoggingBoundService
  05-23 15:38:04.245 - 10205 (com.google.android.adservices.api) - REL *job*r/com.google.android.adservices.api/com.android.adservices.service.measurement.reporting.ImmediateAggregateReportingJobService
  05-23 15:38:04.248 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.s.c.ActionPreferenceCleanupTask (partial)
  05-23 15:38:04.251 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.s.c.ActionPreferenceCleanupTask
  05-23 15:38:04.254 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.d.OdlhDatabaseCleanupJob (partial)
  05-23 15:38:04.274 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.d.OdlhDatabaseCleanupJob
  05-23 15:38:04.276 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.s.PersonalSafetyLoggerService (partial)
  05-23 15:38:04.279 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.s.PersonalSafetyLoggerService
  05-23 15:38:04.281 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/c.g.a.p.s.a.AutoLockLoggerService (partial)
  05-23 15:38:04.287 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/c.g.a.p.s.a.AutoLockLoggerService
  05-23 15:38:04.290 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.c.s.p.PermissionStateLoggingTask (partial)
  05-23 15:38:04.292 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.c.s.p.PermissionStateLoggingTask
  05-23 15:38:04.296 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.SemanticLocationCleanupJob (partial)
  05-23 15:38:04.299 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.SemanticLocationCleanupJob
  05-23 15:38:04.302 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.OdlhTombstonesCleanupJob (partial)
  05-23 15:38:04.304 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.OdlhTombstonesCleanupJob
  05-23 15:38:04.307 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.f.FederatedScheduleService (partial)
  05-23 15:38:04.309 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.f.FederatedScheduleService
  05-23 15:38:04.313 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.c.d.d.s.DatabaseCleanGmsTaskService (partial)
  05-23 15:38:04.318 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.c.d.d.s.DatabaseCleanGmsTaskService
  05-23 15:38:04.321 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.a.o.GpmBiometricsValueMigrationTaskBoundService (partial)
  05-23 15:38:04.324 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.a.o.GpmBiometricsValueMigrationTaskBoundService
  05-23 15:38:04.327 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.c.m.t.PasswordChangesSubscriptionTaskBoundService (partial)
  05-23 15:38:04.328 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.c.m.t.PasswordChangesSubscriptionTaskBoundService
  05-23 15:38:04.331 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.c.m.t.PasswordSharingSubscriptionTaskBoundService (partial)
  05-23 15:38:04.332 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.c.m.t.PasswordSharingSubscriptionTaskBoundService
  05-23 15:38:04.335 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.b.s.n.PhotosBackupMissingPermissionNotificationTask (partial)
  05-23 15:38:04.337 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.b.s.n.PhotosBackupMissingPermissionNotificationTask
  05-23 15:38:04.341 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.f.r.RegionMddMaintenanceService (partial)
  05-23 15:38:04.349 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.f.r.RegionMddMaintenanceService
  05-23 15:38:04.350 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService
  05-23 15:38:04.350 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:04.356 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:04.356 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:04.361 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService (partial)
  05-23 15:38:04.362 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:04.366 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.i.b.IpaGcmTaskService (partial)
  05-23 15:38:04.372 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.i.b.IpaGcmTaskService
  05-23 15:38:04.382 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.l.s.d.DrivingModeLoggerService (partial)
  05-23 15:38:04.394 - 10154 (com.google.android.apps.messaging) - ACQ *job*r/#WorkQueueWorkerShim#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:04.395 - 10154 (com.google.android.apps.messaging) - ACQ *job*r/#WorkQueueWorkerShim#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:04.396 - 10154 (com.google.android.apps.messaging) - REL *job*r/#WorkQueueWorkerShim#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:04.402 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.l.s.d.DrivingModeLoggerService
  05-23 15:38:04.404 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService
  05-23 15:38:04.404 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:04.425 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:04.426 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:04.430 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService (partial)
  05-23 15:38:04.431 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:04.434 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.l.CleanBufferedLogsService (partial)
  05-23 15:38:04.441 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.l.CleanBufferedLogsService
  05-23 15:38:04.447 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.p.PulpMddMaintenanceService (partial)
  05-23 15:38:04.475 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.p.PulpMddMaintenanceService
  05-23 15:38:04.485 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.p.PulpMddDownloadScheduleService (partial)
  05-23 15:38:04.489 - 10154 (com.google.android.apps.messaging) - ACQ StartupHandlerImpl App Interactive Delay Timer Fired -> StartupHandlerImpl#onAppInteractiveInternal -> PhenotypeRestoreStartupTask -> PhenotypeHelper#registerPhenotype -> PhenotypeHelper#onBuglePhenotypeChanged -> WorkQueueBatchingImpl#queueWorkItem -> WorkQueueBatchingImpl#queueWorkItems -> WorkQueueBatchingImpl#schedulingDeferred (partial)
  05-23 15:38:04.501 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.p.PulpMddDownloadScheduleService
  05-23 15:38:04.508 - 10154 (com.google.android.apps.messaging) - REL StartupHandlerImpl App Interactive Delay Timer Fired -> StartupHandlerImpl#onAppInteractiveInternal -> PhenotypeRestoreStartupTask -> PhenotypeHelper#registerPhenotype -> PhenotypeHelper#onBuglePhenotypeChanged -> WorkQueueBatchingImpl#queueWorkItem -> WorkQueueBatchingImpl#queueWorkItems -> WorkQueueBatchingImpl#schedulingDeferred
  05-23 15:38:04.523 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService
  05-23 15:38:04.533 - 10154 (com.google.android.apps.messaging) - REL *job*r/#WorkQueueWorkerShim#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:04.540 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:04.542 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:04.543 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:04.546 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService (partial)
  05-23 15:38:04.546 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:04.549 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.c.f.BorderRouterSyncBoundService (partial)
  05-23 15:38:04.651 - 10165 (com.google.android.contacts) - ACQ *job*r/com.google.android.contacts/com.google.android.apps.contacts.shortcut.ShortcutJobService (partial)
  05-23 15:38:04.651 - 10165 (com.google.android.contacts) - ACQ *job*r/com.google.android.contacts/com.google.android.apps.contacts.shortcut.ShortcutJobService (partial)
  05-23 15:38:04.652 - 10165 (com.google.android.contacts) - REL *job*r/com.google.android.contacts/com.google.android.apps.contacts.shortcut.ShortcutJobService
  05-23 15:38:04.665 - 10165 (com.google.android.contacts) - REL *job*r/com.google.android.contacts/com.google.android.apps.contacts.shortcut.ShortcutJobService
  05-23 15:38:05.651 - 10165 (com.google.android.contacts) - ACQ *job*r/#Cp2UpdateWorker#@androidx.work.systemjobscheduler@com.google.android.contacts/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:05.652 - 10165 (com.google.android.contacts) - ACQ *job*r/#Cp2UpdateWorker#@androidx.work.systemjobscheduler@com.google.android.contacts/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:05.653 - 10165 (com.google.android.contacts) - REL *job*r/#Cp2UpdateWorker#@androidx.work.systemjobscheduler@com.google.android.contacts/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:05.692 - 10165 (com.google.android.contacts) - REL *job*r/#Cp2UpdateWorker#@androidx.work.systemjobscheduler@com.google.android.contacts/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:05.853 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.s.s.PhenotypeConfigurator
  05-23 15:38:05.855 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:05.857 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:05.858 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:05.860 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService (partial)
  05-23 15:38:05.861 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:05.863 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService (partial)
  05-23 15:38:05.868 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService
  05-23 15:38:05.871 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService (partial)
  05-23 15:38:05.874 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService
  05-23 15:38:05.877 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.a.a.b.a.GcmTaskService (partial)
  05-23 15:38:05.885 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.a.a.b.a.GcmTaskService
  05-23 15:38:05.888 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.a.a.b.a.GcmTaskService (partial)
  05-23 15:38:05.892 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.a.a.b.a.GcmTaskService
  05-23 15:38:05.894 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.d.OdlhPppCleanupJob (partial)
  05-23 15:38:05.899 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.d.OdlhPppCleanupJob
  05-23 15:38:05.901 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.g.HeartbeatAlarm$ConnectionInfoPersistService (partial)
  05-23 15:38:05.907 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.g.HeartbeatAlarm$ConnectionInfoPersistService
  05-23 15:38:05.909 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.a.f.s.FolsomPublicKeyUpdateService (partial)
  05-23 15:38:05.913 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.a.f.s.FolsomPublicKeyUpdateService
  05-23 15:38:05.917 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.SettingsLoggingService (partial)
  05-23 15:38:05.926 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.SettingsLoggingService
  05-23 15:38:05.929 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.d.DroidGuardFastRefreshGmsTaskBoundService (partial)
  05-23 15:38:05.932 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.d.DroidGuardFastRefreshGmsTaskBoundService
  05-23 15:38:05.935 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.SafeBrowsingUpdateTaskService (partial)
  05-23 15:38:05.940 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.SafeBrowsingUpdateTaskService
  05-23 15:38:05.943 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.SettingsSyncService (partial)
  05-23 15:38:05.946 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.SettingsSyncService
  05-23 15:38:05.948 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.StalePlacesCleaningService (partial)
  05-23 15:38:05.950 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.StalePlacesCleaningService
  05-23 15:38:05.952 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.b.s.n.BackupNotificationsTask (partial)
  05-23 15:38:05.957 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.b.s.n.BackupNotificationsTask
  05-23 15:38:05.960 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.r.ReinferenceService (partial)
  05-23 15:38:05.963 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.r.ReinferenceService
  05-23 15:38:05.965 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.a.b.s.c.CloudSyncBackupTaskService (partial)
  05-23 15:38:05.971 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.a.b.s.c.CloudSyncBackupTaskService
  05-23 15:38:05.973 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.d.g.ParcelRefreshTokensJob (partial)
  05-23 15:38:05.975 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.d.g.ParcelRefreshTokensJob
  05-23 15:38:05.978 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.LocationHistoryActiveProcessingService (partial)
  05-23 15:38:05.979 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.LocationHistoryActiveProcessingService
  05-23 15:38:05.982 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.d.w.WifiPlaceVisitProcessingService (partial)
  05-23 15:38:05.984 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.d.w.WifiPlaceVisitProcessingService
  05-23 15:38:05.986 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.d.g.ParcelSynchronizePersonalizationJob (partial)
  05-23 15:38:05.988 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.d.g.ParcelSynchronizePersonalizationJob
  05-23 15:38:05.990 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.d.h.TransientHistoricalBusynessProcessingService (partial)
  05-23 15:38:05.992 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.d.h.TransientHistoricalBusynessProcessingService
  05-23 15:38:05.993 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService
  05-23 15:38:05.995 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:05.996 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:05.997 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:05.999 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.i.p.IcingInternalCorporaUpdateService (partial)
  05-23 15:38:06.001 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:06.002 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.i.p.IcingInternalCorporaUpdateService
  05-23 15:38:06.002 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:06.005 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:06.005 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:06.009 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService (partial)
  05-23 15:38:06.010 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:06.011 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService
  05-23 15:38:06.011 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:06.015 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:06.016 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:06.018 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService (partial)
  05-23 15:38:06.018 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:06.021 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.RegularRetryVerificationTaskBoundService (partial)
  05-23 15:38:06.023 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.RegularRetryVerificationTaskBoundService
  05-23 15:38:06.026 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.l.CleanBufferedLogsService (partial)
  05-23 15:38:06.026 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.l.CleanBufferedLogsService
  05-23 15:38:06.029 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.p.PulpMddMaintenanceService (partial)
  05-23 15:38:06.031 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.p.PulpMddMaintenanceService
  05-23 15:38:06.033 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.p.PulpMddDownloadScheduleService (partial)
  05-23 15:38:06.034 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.p.PulpMddDownloadScheduleService
  05-23 15:38:06.035 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService
  05-23 15:38:06.035 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:06.038 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:06.038 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:06.039 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:07.569 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.c.f.BorderRouterSyncBoundService
  05-23 15:38:07.570 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService
  05-23 15:38:07.575 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:07.580 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:07.580 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:07.584 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService (partial)
  05-23 15:38:07.589 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService (partial)
  05-23 15:38:07.589 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:07.590 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService
  05-23 15:38:07.590 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:07.593 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:07.593 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:07.593 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService (partial)
  05-23 15:38:07.594 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:07.598 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService
  05-23 15:38:07.602 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService (partial)
  05-23 15:38:07.606 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.p.g.PayGcmTaskService
  05-23 15:38:07.608 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.a.a.b.a.GcmTaskService (partial)
  05-23 15:38:07.612 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.a.a.b.a.GcmTaskService
  05-23 15:38:07.614 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.a.a.b.a.GcmTaskService (partial)
  05-23 15:38:07.617 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.a.a.b.a.GcmTaskService
  05-23 15:38:07.620 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.d.OdlhPppCleanupJob (partial)
  05-23 15:38:07.622 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.d.OdlhPppCleanupJob
  05-23 15:38:07.624 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.g.HeartbeatAlarm$ConnectionInfoPersistService (partial)
  05-23 15:38:07.627 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.g.HeartbeatAlarm$ConnectionInfoPersistService
  05-23 15:38:07.630 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.a.f.s.FolsomPublicKeyUpdateService (partial)
  05-23 15:38:07.634 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.a.f.s.FolsomPublicKeyUpdateService
  05-23 15:38:07.637 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.SettingsLoggingService (partial)
  05-23 15:38:07.638 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.SettingsLoggingService
  05-23 15:38:07.641 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.d.DroidGuardFastRefreshGmsTaskBoundService (partial)
  05-23 15:38:07.642 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.d.DroidGuardFastRefreshGmsTaskBoundService
  05-23 15:38:07.645 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.SafeBrowsingUpdateTaskService (partial)
  05-23 15:38:07.648 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.SafeBrowsingUpdateTaskService
  05-23 15:38:07.650 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.SettingsSyncService (partial)
  05-23 15:38:07.652 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.SettingsSyncService
  05-23 15:38:07.654 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.StalePlacesCleaningService (partial)
  05-23 15:38:07.655 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.StalePlacesCleaningService
  05-23 15:38:07.657 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.b.s.n.BackupNotificationsTask (partial)
  05-23 15:38:07.659 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.b.s.n.BackupNotificationsTask
  05-23 15:38:07.661 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.r.ReinferenceService (partial)
  05-23 15:38:07.663 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.r.ReinferenceService
  05-23 15:38:07.665 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.a.b.s.c.CloudSyncBackupTaskService (partial)
  05-23 15:38:07.668 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.a.b.s.c.CloudSyncBackupTaskService
  05-23 15:38:07.674 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.d.g.ParcelRefreshTokensJob (partial)
  05-23 15:38:07.675 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.d.g.ParcelRefreshTokensJob
  05-23 15:38:07.678 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.LocationHistoryActiveProcessingService (partial)
  05-23 15:38:07.679 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.LocationHistoryActiveProcessingService
  05-23 15:38:07.681 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.d.w.WifiPlaceVisitProcessingService (partial)
  05-23 15:38:07.682 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.d.w.WifiPlaceVisitProcessingService
  05-23 15:38:07.685 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.d.g.ParcelSynchronizePersonalizationJob (partial)
  05-23 15:38:07.686 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.d.g.ParcelSynchronizePersonalizationJob
  05-23 15:38:07.688 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.d.h.TransientHistoricalBusynessProcessingService (partial)
  05-23 15:38:07.690 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.d.h.TransientHistoricalBusynessProcessingService
  05-23 15:38:07.690 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService
  05-23 15:38:07.690 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:07.694 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:07.694 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:07.695 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:10.378 - 10154 (com.google.android.apps.messaging) - ACQ bugle_datamodel_executor_wakelock (partial)
  05-23 15:38:10.380 - 10154 (com.google.android.apps.messaging) - ACQ com.google.android.apps.messaging.shared.datamodel.action.FixupMessageStatusOnStartupAction (partial)
  05-23 15:38:10.402 - 10154 (com.google.android.apps.messaging) - ACQ com.google.android.apps.messaging.shared.datamodel.action.ProcessPendingMessagesAction (partial)
  05-23 15:38:10.403 - 10154 (com.google.android.apps.messaging) - REL com.google.android.apps.messaging.shared.datamodel.action.FixupMessageStatusOnStartupAction
  05-23 15:38:10.404 - 10154 (com.google.android.apps.messaging) - ACQ com.google.android.apps.messaging.shared.datamodel.action.SelfParticipantsRefreshAction (partial)
  05-23 15:38:10.404 - 10154 (com.google.android.apps.messaging) - REL com.google.android.apps.messaging.shared.datamodel.action.ProcessPendingMessagesAction
  05-23 15:38:10.405 - 10154 (com.google.android.apps.messaging) - ACQ com.google.android.apps.messaging.shared.datamodel.action.SelfParticipantsRefreshAction (partial)
  05-23 15:38:10.405 - 10154 (com.google.android.apps.messaging) - REL com.google.android.apps.messaging.shared.datamodel.action.SelfParticipantsRefreshAction
  05-23 15:38:10.405 - 10154 (com.google.android.apps.messaging) - ACQ com.google.android.apps.messaging.shared.datamodel.action.GenericWorkerQueueAction (partial)
  05-23 15:38:10.405 - 10154 (com.google.android.apps.messaging) - REL com.google.android.apps.messaging.shared.datamodel.action.GenericWorkerQueueAction
  05-23 15:38:10.433 - 10154 (com.google.android.apps.messaging) - ACQ com.google.android.apps.messaging.shared.datamodel.action.SelfParticipantsRefreshAction (partial)
  05-23 15:38:10.433 - 10154 (com.google.android.apps.messaging) - ACQ com.google.android.apps.messaging.shared.datamodel.action.GenericWorkerQueueAction (partial)
  05-23 15:38:10.433 - 10154 (com.google.android.apps.messaging) - REL com.google.android.apps.messaging.shared.datamodel.action.SelfParticipantsRefreshAction
  05-23 15:38:10.433 - 10154 (com.google.android.apps.messaging) - REL com.google.android.apps.messaging.shared.datamodel.action.SelfParticipantsRefreshAction
  05-23 15:38:10.434 - 10154 (com.google.android.apps.messaging) - REL bugle_datamodel_executor_wakelock
  05-23 15:38:10.434 - 10154 (com.google.android.apps.messaging) - REL com.google.android.apps.messaging.shared.datamodel.action.GenericWorkerQueueAction
  05-23 15:38:10.434 - 10154 (com.google.android.apps.messaging) - ACQ com.google.android.apps.messaging.shared.datamodel.action.GenericWorkerQueueAction (partial)
  05-23 15:38:10.435 - 10154 (com.google.android.apps.messaging) - REL com.google.android.apps.messaging.shared.datamodel.action.GenericWorkerQueueAction
  05-23 15:38:11.040 - 1000 (System) - ACQ *alarm* (partial)
  05-23 15:38:11.051 - 1000 (System) - REL *alarm*
  05-23 15:38:12.722 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:12.726 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.c.CheckinService (partial)
  05-23 15:38:12.729 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:12.733 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.c.CheckinService
  05-23 15:38:12.733 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:12.736 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:12.737 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:12.743 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService (partial)
  05-23 15:38:12.743 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:12.746 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.s.SpamListSyncTaskService (partial)
  05-23 15:38:12.750 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.s.SpamListSyncTaskService
  05-23 15:38:12.754 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.c.f.BorderRouterSyncBoundService (partial)
  05-23 15:38:13.507 - 10171 (com.google.android.apps.photos) - REL *job*r/com.google.android.apps.photos/com.google.android.libraries.social.async.BackgroundTaskJobService
  05-23 15:38:13.585 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:13.588 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.i.b.IpaGcmTaskService (partial)
  05-23 15:38:13.590 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:13.590 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:13.591 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:13.592 - 10153 (com.android.vending) - ACQ *job*r/com.android.vending/com.google.android.finsky.scheduler.process.mainimpl.PhoneskyJobServiceMain (partial)
  05-23 15:38:13.592 - 10153 (com.android.vending) - ACQ *job*r/com.android.vending/com.google.android.finsky.scheduler.process.mainimpl.PhoneskyJobServiceMain (partial)
  05-23 15:38:13.592 - 10153 (com.android.vending) - REL *job*r/com.android.vending/com.google.android.finsky.scheduler.process.mainimpl.PhoneskyJobServiceMain
  05-23 15:38:13.592 - 10154 (com.google.android.apps.messaging) - ACQ *job*r/com.google.android.apps.messaging/.shared.datamodel.action.execution.ActionJobService (partial)
  05-23 15:38:13.592 - 10154 (com.google.android.apps.messaging) - ACQ *job*r/com.google.android.apps.messaging/.shared.datamodel.action.execution.ActionJobService (partial)
  05-23 15:38:13.593 - 10154 (com.google.android.apps.messaging) - REL *job*r/com.google.android.apps.messaging/.shared.datamodel.action.execution.ActionJobService
  05-23 15:38:13.594 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.i.b.IpaGcmTaskService
  05-23 15:38:13.595 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:13.602 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:13.602 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:13.607 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:13.609 - 10154 (com.google.android.apps.messaging) - ACQ bugle_datamodel_executor_wakelock (partial)
  05-23 15:38:13.610 - 10154 (com.google.android.apps.messaging) - ACQ com.google.android.apps.messaging.shared.datamodel.action.CountryCodeDetectorAction (partial)
  05-23 15:38:13.610 - 10154 (com.google.android.apps.messaging) - ACQ com.google.android.apps.messaging.shared.datamodel.action.CountryCodeDetectorAction (partial)
  05-23 15:38:13.610 - 10154 (com.google.android.apps.messaging) - REL bugle_datamodel_executor_wakelock
  05-23 15:38:13.610 - 10154 (com.google.android.apps.messaging) - REL com.google.android.apps.messaging.shared.datamodel.action.CountryCodeDetectorAction
  05-23 15:38:13.610 - 10154 (com.google.android.apps.messaging) - REL *job*r/com.google.android.apps.messaging/.shared.datamodel.action.execution.ActionJobService
  05-23 15:38:13.610 - 10154 (com.google.android.apps.messaging) - ACQ com.google.android.apps.messaging.shared.datamodel.action.CountryCodeDetectorAction (partial)
  05-23 15:38:13.610 - 10154 (com.google.android.apps.messaging) - REL com.google.android.apps.messaging.shared.datamodel.action.CountryCodeDetectorAction
  05-23 15:38:13.611 - 10154 (com.google.android.apps.messaging) - REL com.google.android.apps.messaging.shared.datamodel.action.CountryCodeDetectorAction
  05-23 15:38:13.627 - 10153 (com.android.vending) - REL *job*r/com.android.vending/com.google.android.finsky.scheduler.process.mainimpl.PhoneskyJobServiceMain
  05-23 15:38:13.650 - 10154 (com.google.android.apps.messaging) - ACQ *job*r/#TikTokWorker#com.google.android.apps.messaging.shared.contacts.sync.incremental.ContactsChangeWorker#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:13.650 - 10154 (com.google.android.apps.messaging) - ACQ *job*r/#TikTokWorker#com.google.android.apps.messaging.shared.contacts.sync.incremental.ContactsChangeWorker#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:13.651 - 10154 (com.google.android.apps.messaging) - REL *job*r/#TikTokWorker#com.google.android.apps.messaging.shared.contacts.sync.incremental.ContactsChangeWorker#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:13.670 - 10154 (com.google.android.apps.messaging) - REL *job*r/#TikTokWorker#com.google.android.apps.messaging.shared.contacts.sync.incremental.ContactsChangeWorker#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:14.703 - 10165 (com.google.android.contacts) - ACQ *job*r/com.google.android.contacts/com.google.android.apps.contacts.shortcut.ShortcutJobService (partial)
  05-23 15:38:14.703 - 10165 (com.google.android.contacts) - ACQ *job*r/com.google.android.contacts/com.google.android.apps.contacts.shortcut.ShortcutJobService (partial)
  05-23 15:38:14.704 - 10165 (com.google.android.contacts) - REL *job*r/com.google.android.contacts/com.google.android.apps.contacts.shortcut.ShortcutJobService
  05-23 15:38:14.711 - 10165 (com.google.android.contacts) - REL *job*r/com.google.android.contacts/com.google.android.apps.contacts.shortcut.ShortcutJobService
  05-23 15:38:15.051 - 10219 (com.google.android.gms) - ACQ *alarm* (partial)
  05-23 15:38:15.054 - 10219 (com.google.android.gms) - ACQ GCM_WORK (partial)
  05-23 15:38:15.055 - 10219 (com.google.android.gms) - REL *alarm*
  05-23 15:38:15.055 - 10219 (com.google.android.gms) - REL GCM_WORK
  05-23 15:38:15.160 - 10219 (com.google.android.gms) - ACQ CMWakeLock (partial)
  05-23 15:38:15.162 - 10219 (com.google.android.gms) - REL CMWakeLock
  05-23 15:38:15.171 - 10219 (com.google.android.gms) - ACQ CMWakeLock (partial)
  05-23 15:38:15.172 - 10219 (com.google.android.gms) - REL CMWakeLock
  05-23 15:38:15.180 - 10219 (com.google.android.gms) - ACQ CMWakeLock (partial)
  05-23 15:38:15.181 - 10159 (com.google.android.apps.wellbeing) - ACQ *job*r/#TikTokWorker#com.google.android.apps.wellbeing.autodnd.SettingsInjectionWorker#@androidx.work.systemjobscheduler@com.google.android.apps.wellbeing/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:15.181 - 10159 (com.google.android.apps.wellbeing) - ACQ *job*r/#TikTokWorker#com.google.android.apps.wellbeing.autodnd.SettingsInjectionWorker#@androidx.work.systemjobscheduler@com.google.android.apps.wellbeing/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:15.181 - 10219 (com.google.android.gms) - REL CMWakeLock
  05-23 15:38:15.182 - 10159 (com.google.android.apps.wellbeing) - REL *job*r/#TikTokWorker#com.google.android.apps.wellbeing.autodnd.SettingsInjectionWorker#@androidx.work.systemjobscheduler@com.google.android.apps.wellbeing/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:15.183 - 10159 (com.google.android.apps.wellbeing) - ACQ *job*r/#TikTokWorker#com.google.android.apps.wellbeing.task.Task#@androidx.work.systemjobscheduler@com.google.android.apps.wellbeing/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:15.183 - 10159 (com.google.android.apps.wellbeing) - ACQ *job*r/#TikTokWorker#com.google.android.apps.wellbeing.task.Task#@androidx.work.systemjobscheduler@com.google.android.apps.wellbeing/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:15.184 - 10219 (com.google.android.gms) - ACQ CMWakeLock (partial)
  05-23 15:38:15.184 - 10159 (com.google.android.apps.wellbeing) - REL *job*r/#TikTokWorker#com.google.android.apps.wellbeing.task.Task#@androidx.work.systemjobscheduler@com.google.android.apps.wellbeing/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:15.184 - 10219 (com.google.android.gms) - REL CMWakeLock
  05-23 15:38:15.187 - 10159 (com.google.android.apps.wellbeing) - ACQ *job*r/#TikTokWorker#com.google.android.apps.wellbeing.widget.WidgetFlagCheckWorker#@androidx.work.systemjobscheduler@com.google.android.apps.wellbeing/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:15.187 - 10159 (com.google.android.apps.wellbeing) - ACQ *job*r/#TikTokWorker#com.google.android.apps.wellbeing.widget.WidgetFlagCheckWorker#@androidx.work.systemjobscheduler@com.google.android.apps.wellbeing/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:15.187 - 10159 (com.google.android.apps.wellbeing) - REL *job*r/#TikTokWorker#com.google.android.apps.wellbeing.widget.WidgetFlagCheckWorker#@androidx.work.systemjobscheduler@com.google.android.apps.wellbeing/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:15.190 - 10219 (com.google.android.gms) - ACQ CMWakeLock (partial)
  05-23 15:38:15.191 - 10219 (com.google.android.gms) - REL CMWakeLock
  05-23 15:38:15.193 - 10159 (com.google.android.apps.wellbeing) - REL *job*r/#TikTokWorker#com.google.android.apps.wellbeing.widget.WidgetFlagCheckWorker#@androidx.work.systemjobscheduler@com.google.android.apps.wellbeing/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:15.195 - 10159 (com.google.android.apps.wellbeing) - REL *job*r/#TikTokWorker#com.google.android.apps.wellbeing.autodnd.SettingsInjectionWorker#@androidx.work.systemjobscheduler@com.google.android.apps.wellbeing/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:15.241 - 10159 (com.google.android.apps.wellbeing) - REL *job*r/#TikTokWorker#com.google.android.apps.wellbeing.task.Task#@androidx.work.systemjobscheduler@com.google.android.apps.wellbeing/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:15.702 - 10165 (com.google.android.contacts) - ACQ *job*r/#Cp2UpdateWorker#@androidx.work.systemjobscheduler@com.google.android.contacts/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:15.702 - 10165 (com.google.android.contacts) - ACQ *job*r/#Cp2UpdateWorker#@androidx.work.systemjobscheduler@com.google.android.contacts/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:15.703 - 10165 (com.google.android.contacts) - REL *job*r/#Cp2UpdateWorker#@androidx.work.systemjobscheduler@com.google.android.contacts/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:15.735 - 10165 (com.google.android.contacts) - REL *job*r/#Cp2UpdateWorker#@androidx.work.systemjobscheduler@com.google.android.contacts/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:15.763 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.c.f.BorderRouterSyncBoundService
  05-23 15:38:15.765 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.s.s.h.HousekeepingTrainTaskService
  05-23 15:38:15.765 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:15.770 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:15.770 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:15.771 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:19.786 - 10158 (com.google.android.googlequicksearchbox) - ACQ *job*r/#TikTokWorker#com.google.android.libraries.search.assistant.launcher.OpaEnabledBroadcastWorker#@androidx.work.systemjobscheduler@com.google.android.googlequicksearchbox/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:19.786 - 10158 (com.google.android.googlequicksearchbox) - ACQ *job*r/#TikTokWorker#com.google.android.libraries.search.assistant.launcher.OpaEnabledBroadcastWorker#@androidx.work.systemjobscheduler@com.google.android.googlequicksearchbox/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:19.786 - 10158 (com.google.android.googlequicksearchbox) - REL *job*r/#TikTokWorker#com.google.android.libraries.search.assistant.launcher.OpaEnabledBroadcastWorker#@androidx.work.systemjobscheduler@com.google.android.googlequicksearchbox/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:19.808 - 10158 (com.google.android.googlequicksearchbox) - REL *job*r/#TikTokWorker#com.google.android.libraries.search.assistant.launcher.OpaEnabledBroadcastWorker#@androidx.work.systemjobscheduler@com.google.android.googlequicksearchbox/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:22.065 - 10219 (com.google.android.gms) - ACQ *alarm* (partial)
  05-23 15:38:22.081 - 10219 (com.google.android.gms) - REL *alarm*
  05-23 15:38:22.083 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:22.089 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.a.s.GcmSchedulerWakeupService (partial)
  05-23 15:38:22.093 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:22.105 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.a.s.GcmSchedulerWakeupService
  05-23 15:38:22.105 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:22.110 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:22.110 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:22.111 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:23.582 - 10171 (com.google.android.apps.photos) - ACQ *job*r/com.google.android.apps.photos/com.google.android.libraries.social.mediamonitor.MediaMonitorJobSchedulerService (partial)
  05-23 15:38:23.583 - 10171 (com.google.android.apps.photos) - ACQ *job*r/com.google.android.apps.photos/com.google.android.libraries.social.mediamonitor.MediaMonitorJobSchedulerService (partial)
  05-23 15:38:23.583 - 10171 (com.google.android.apps.photos) - REL *job*r/com.google.android.apps.photos/com.google.android.libraries.social.mediamonitor.MediaMonitorJobSchedulerService
  05-23 15:38:23.625 - 10171 (com.google.android.apps.photos) - REL *job*r/com.google.android.apps.photos/com.google.android.libraries.social.mediamonitor.MediaMonitorJobSchedulerService
  05-23 15:38:23.670 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:23.673 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.i.p.IcingInternalCorporaUpdateService (partial)
  05-23 15:38:23.676 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:23.677 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.i.p.IcingInternalCorporaUpdateService
  05-23 15:38:23.678 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:23.680 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:23.681 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:23.685 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:23.701 - 10154 (com.google.android.apps.messaging) - ACQ *job*r/#TikTokWorker#com.google.android.apps.messaging.shared.contacts.sync.incremental.ContactsChangeWorker#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:23.701 - 10154 (com.google.android.apps.messaging) - ACQ *job*r/#TikTokWorker#com.google.android.apps.messaging.shared.contacts.sync.incremental.ContactsChangeWorker#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:38:23.701 - 10154 (com.google.android.apps.messaging) - REL *job*r/#TikTokWorker#com.google.android.apps.messaging.shared.contacts.sync.incremental.ContactsChangeWorker#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:23.712 - 10154 (com.google.android.apps.messaging) - REL *job*r/#TikTokWorker#com.google.android.apps.messaging.shared.contacts.sync.incremental.ContactsChangeWorker#@androidx.work.systemjobscheduler@com.google.android.apps.messaging/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:38:25.130 - 1000 (System) - ACQ *alarm* (partial)
  05-23 15:38:25.138 - 1000 (System) - REL *alarm*
  05-23 15:38:39.273 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:38:42.584 - 10153 (com.android.vending) - ACQ *alarm* (partial)
  05-23 15:38:42.585 - 10171 (com.google.android.apps.photos) - ACQ *job*r/com.google.android.apps.photos/.dbprocessor.impl.DatabaseProcessorJobService (partial)
  05-23 15:38:42.585 - 10171 (com.google.android.apps.photos) - ACQ *job*r/com.google.android.apps.photos/.dbprocessor.impl.DatabaseProcessorJobService (partial)
  05-23 15:38:42.587 - 10171 (com.google.android.apps.photos) - REL *job*r/com.google.android.apps.photos/.dbprocessor.impl.DatabaseProcessorJobService
  05-23 15:38:42.589 - 10219 (com.google.android.gms) - REL *alarm*
  05-23 15:38:42.590 - 10171 (com.google.android.apps.photos) - REL *job*r/com.google.android.apps.photos/.dbprocessor.impl.DatabaseProcessorJobService
  05-23 15:38:42.592 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:42.597 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService (partial)
  05-23 15:38:42.603 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService (partial)
  05-23 15:38:42.606 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:42.609 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService
  05-23 15:38:42.609 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:42.611 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:42.612 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:42.613 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:51.982 - 10219 (com.google.android.gms) - ACQ *alarm* (partial)
  05-23 15:38:51.984 - 10219 (com.google.android.gms) - REL *alarm*
  05-23 15:38:51.986 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:51.991 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.a.a.i.d.DeviceAccountTaskService (partial)
  05-23 15:38:51.992 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:51.999 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.a.a.i.d.DeviceAccountTaskService
  05-23 15:38:51.999 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:52.002 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:52.003 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:38:52.003 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:38:53.801 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:39:00.000 - 1000 (System) - ACQ *alarm* (partial)
  05-23 15:39:00.000 - 1000 (System) - REL *alarm*
  05-23 15:39:11.792 - 10153 (com.android.vending) - ACQ *job*r/com.android.vending/com.google.android.finsky.scheduler.process.mainimpl.PhoneskyJobServiceMain (partial)
  05-23 15:39:11.792 - 10153 (com.android.vending) - ACQ *job*r/com.android.vending/com.google.android.finsky.scheduler.process.mainimpl.PhoneskyJobServiceMain (partial)
  05-23 15:39:11.815 - 10153 (com.android.vending) - REL *job*r/com.android.vending/com.google.android.finsky.scheduler.process.mainimpl.PhoneskyJobServiceMain
  05-23 15:39:12.080 - 10153 (com.android.vending) - REL *job*r/com.android.vending/com.google.android.finsky.scheduler.process.mainimpl.PhoneskyJobServiceMain
  05-23 15:39:12.325 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:39:13.632 - 10220 (com.google.android.providers.media.module) - ACQ *alarm* (partial)
  05-23 15:39:13.632 - 10220 (com.google.android.providers.media.module) - REL *alarm*
  05-23 15:39:13.632 - 10220 (com.google.android.providers.media.module) - ACQ *job*r/#MediaServiceV2#@androidx.work.systemjobscheduler@com.google.android.providers.media.module/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:39:13.632 - 10220 (com.google.android.providers.media.module) - ACQ *job*r/#MediaServiceV2#@androidx.work.systemjobscheduler@com.google.android.providers.media.module/androidx.work.impl.background.systemjob.SystemJobService (partial)
  05-23 15:39:13.632 - 10220 (com.google.android.providers.media.module) - REL *job*r/#MediaServiceV2#@androidx.work.systemjobscheduler@com.google.android.providers.media.module/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:39:13.668 - 10220 (com.google.android.providers.media.module) - REL *job*r/#MediaServiceV2#@androidx.work.systemjobscheduler@com.google.android.providers.media.module/androidx.work.impl.background.systemjob.SystemJobService
  05-23 15:39:17.831 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:39:33.398 - 10219 (com.google.android.gms) - ACQ *alarm* (partial)
  05-23 15:39:33.401 - 10219 (com.google.android.gms) - ACQ CallbackRunner (partial)
  05-23 15:39:33.402 - 10219 (com.google.android.gms) - REL *alarm*
  05-23 15:39:33.418 - 10219 (com.google.android.gms) - ACQ CollectionLib-SigCollector (partial)
  05-23 15:39:33.423 - 10219 (com.google.android.gms) - REL CallbackRunner
  05-23 15:39:33.977 - 10219 (com.google.android.gms) - REL CollectionLib-SigCollector
  05-23 15:39:33.985 - 10219 (com.google.android.gms) - ACQ CallbackRunner (partial)
  05-23 15:39:33.994 - 10219 (com.google.android.gms) - ACQ PendingIntentClient (partial)
  05-23 15:39:33.995 - 10219 (com.google.android.gms) - REL PendingIntentClient
  05-23 15:39:33.996 - 10219 (com.google.android.gms) - ACQ PendingIntentClient (partial)
  05-23 15:39:33.998 - 10219 (com.google.android.gms) - ACQ GCoreFlp (partial)
  05-23 15:39:34.002 - 10219 (com.google.android.gms) - REL CallbackRunner
  05-23 15:39:34.004 - 10219 (com.google.android.gms) - REL GCoreFlp
  05-23 15:39:34.007 - 10219 (com.google.android.gms) - REL PendingIntentClient
  05-23 15:39:36.358 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  05-23 15:39:39.848 - 1000 (System) - ACQ *alarm* (partial)
  05-23 15:39:39.849 - 1000 (System) - REL *alarm*
  05-23 15:39:59.931 - 1073 (com.google.android.cellbroadcastservice,...) - ACQ IpReachabilityMonitor.wlan0 (partial)
  05-23 15:40:00.002 - 1000 (System) - ACQ *alarm* (partial)
  05-23 15:40:00.003 - 1000 (System) - REL *alarm*
  05-23 15:40:06.133 - 10171 (com.google.android.apps.photos) - ACQ *alarm* (partial)
  05-23 15:40:06.136 - 10171 (com.google.android.apps.photos) - ACQ *job*r/com.google.android.apps.photos/.dbprocessor.impl.DatabaseProcessorJobService (partial)
  05-23 15:40:06.136 - 10171 (com.google.android.apps.photos) - ACQ *job*r/com.google.android.apps.photos/.dbprocessor.impl.DatabaseProcessorJobService (partial)
  05-23 15:40:06.149 - 10171 (com.google.android.apps.photos) - REL *job*r/com.google.android.apps.photos/.dbprocessor.impl.DatabaseProcessorJobService
  05-23 15:40:06.151 - 10219 (com.google.android.gms) - REL *alarm*
  05-23 15:40:06.154 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:40:06.161 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService (partial)
  05-23 15:40:06.168 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:40:06.170 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService
  05-23 15:40:06.171 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:40:06.174 - 10171 (com.google.android.apps.photos) - REL *job*r/com.google.android.apps.photos/.dbprocessor.impl.DatabaseProcessorJobService
  05-23 15:40:06.178 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:40:06.178 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:40:06.182 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.a.j.NegotiationService (partial)
  05-23 15:40:06.184 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:40:06.196 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.a.j.NegotiationService
  05-23 15:40:06.197 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:40:06.201 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:40:06.201 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:40:06.204 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService (partial)
  05-23 15:40:06.206 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:40:06.219 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.g.TapAndPayGcmTaskService
  05-23 15:40:06.219 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:40:06.221 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:40:06.221 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:40:06.224 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*/c.g.a.g/.t.s.SpamListSyncTaskService (partial)
  05-23 15:40:06.226 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:40:06.228 - 10219 (com.google.android.gms) - REL *gms_scheduler*/c.g.a.g/.t.s.SpamListSyncTaskService
  05-23 15:40:06.228 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:40:06.231 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:40:06.231 - 10219 (com.google.android.gms) - ACQ *gms_scheduler*:internal (partial)
  05-23 15:40:06.232 - 10219 (com.google.android.gms) - REL *gms_scheduler*:internal
  05-23 15:40:18.477 - 1073 (com.google.android.cellbroadcastservice,...) - REL IpReachabilityMonitor.wlan0
  -
  Events: 1109, Time-Resets: 35
  Buffer, Bytes used: 3070
  Tag Database: size(128), entries: 128, Bytes used: 23056

Full Wakelock Log:
Wake Lock Log
  -
  Events: 0, Time-Resets: 0
  Buffer, Bytes used: 0
  Tag Database: size(16), entries: 0, Bytes used: 128


Wakefulness Session Observer:
default timeout: 2147483647
override timeout: -1
  Wakefulness Session Power Group powerGroupId: 0
    current wakefulness: 0
    current user activity event: 6
    current user activity duration: 163628
    previous user activity event: 0
    previous user activity duration: 175082
    is in override timeout: false
    mIsInteractive: false
    current screen policy: 3
    current screen policy duration: 175081
    previous screen policy: 3
    past screen policy duration: 0

FaceDownDetector:
  mFaceDown=false
  mActive=true
  mLastFlipTime=0
  mSensorMaxLatencyMicros=2000000
  mUserInteractionBackoffMillis=60000
  mPreviousResultTime=0
  mPreviousResultType=1
  mMillisSaved=0
  mZAccelerationThreshold=-9.5
  mAccelerationThreshold=0.2
  mTimeThreshold=PT1S
  mEnabledOverride=true
AmbientDisplaySuppressionController:
 ambientDisplaySuppressed=false
 mSuppressionTokens={}
 mSuppressions={}

Low Power Standby Controller:
  mIsActive=false
  mIsEnabled=false
  mSupportedConfig=false
  mEnabledByDefaultConfig=false
  mStandbyTimeoutConfig=0
  mEnableCustomPolicy=false
  Allowed UIDs=[]
  
  UID allowed reasons:

ScreenTimeoutOverridePolicy:
  mScreenTimeoutOverrideConfig=-1
  mLastAutoReleaseReason=-1
PowerManagerFlags:
 enable_early_screen_timeout_detector:                 true (def:true)
 per_display_wake_by_touch:                            true (def:true)
 lock_on_unplug:                                       true (def:true)
 disable_frozen_process_wakelocks:                     true (def:true)
 partial_sleep_wakelocks:                              true (def:true)
 separate_timeouts_flicker:                            true (def:true)
 wait_for_user_boot_complete:                          false (def:false)
```
