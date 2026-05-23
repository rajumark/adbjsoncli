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
    "mHoldingWakeLockSuspendBlocker": "false",
    "mInterceptedPowerKeyForProximity": "false",
    "mIsFaceDown": "false",
    "mIsPowered": "false",
    "mLastEnhancedDischargeTimeUpdatedElapsed": "0",
    "mLastFlipTime": "0",
    "mLastGlobalSleepTimeRealtime": "0 (415479 ms ago)",
    "mLastGlobalWakeTimeRealtime": "0 (415479 ms ago)",
    "mLastInteractivePowerHintTime": "349977 (65502 ms ago)",
    "mLastScreenBrightnessBoostTime": "0 (415479 ms ago)",
    "mLastSleepReason": "application",
    "mLastSleepTime": "0 (415479 ms ago)",
    "mLastWakeTime": "0 (415479 ms ago)",
    "mLightDeviceIdleMode": "false",
    "mLowPowerStandbyActive": "false",
    "mNotifyLongDispatched": "-36s487ms",
    "mNotifyLongNextCheck": "(none)",
    "mNotifyLongScheduled": "(none)",
    "mPlugType": "0",
    "mProximityPositive": "false",
    "mRequestWaitForNegativeProximity": "false",
    "mSandmanScheduled": "false",
    "mScreenBrightnessBoostInProgress": "false",
    "mStayOn": "false",
    "mSystemReady": "true",
    "mUseAutoSuspend": "true",
    "mWakeLockSummary": "0x0",
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
  mWakeLockSummary=0x0
  mNotifyLongScheduled=(none)
  mNotifyLongDispatched=-36s508ms
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
  mLastWakeTime=0 (415500 ms ago)
  mLastSleepTime=0 (415500 ms ago)
  mLastSleepReason=application
  mLastGlobalWakeTimeRealtime=0 (415500 ms ago)
  mLastGlobalSleepTimeRealtime=0 (415500 ms ago)
  mLastInteractivePowerHintTime=349977 (65523 ms ago)
  mLastScreenBrightnessBoostTime=0 (415500 ms ago)
  mScreenBrightnessBoostInProgress=false
  mHoldingWakeLockSuspendBlocker=false
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
  UID 1001:   ACTIVE  count=0 state=0\n\n... (truncated,     1497 lines total)
```
