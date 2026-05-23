# `adbjson shell dumpsys uimode`

## adbjson

**Command:**
```bash
adbjson shell dumpsys uimode
```

**Output:**
```json
{
  "status": 0,
  "output": {}
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys uimode
```

**Output:**
```
Current UI Mode Service state:
  mDockState=0 mLastBroadcastState=0
 mStartDreamImmediatelyOnDock=true  mNightMode=1 (no)  mOverrideOn/Off=false/false  mAttentionModeThemeOverlay=1000 mNightModeLocked=true
  mCarModeEnabled=false (carModeApps=
 mWaitForDeviceInactive=false mComputedNightMode=false customStart=22:00 customEnd06:00 mCarModeEnableFlags=0 mEnableCarDockLaunch=false
  mCurUiMode=0x11 mUiModeLocked=false mSetUiMode=0x11
  mHoldingConfiguration=false mSystemReady=true
  mTwilightService.getLastTwilightState()=null
```
