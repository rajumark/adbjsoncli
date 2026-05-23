# `adbjson shell dumpsys statusbar`

## adbjson

**Command:**
```bash
adbjson shell dumpsys statusbar
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
adb shell dumpsys statusbar
```

**Output:**
```
  displayId=0
    mDisabled1=0x0
    mDisabled2=0x0
  mDisableRecords.size=2
    [0] userId=0 what1=0x00000000 what2=0x00000000 pkg=com.android.systemui token=android.os.BinderProxy@27ab73f
    [1] userId=0 what1=0x00000000 what2=0x00000000 pkg=com.android.systemui token=android.os.BinderProxy@bcf680c
  mCurrentUserId=0
  mIcons=
  mCurrentRequestAddTilePackages=[
  ]
  mShowPowerMenuCallbacks=[
  ]
  TileRequestTracker:
```
