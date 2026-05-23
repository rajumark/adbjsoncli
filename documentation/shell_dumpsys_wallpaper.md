# `adbjson shell dumpsys wallpaper`

## adbjson

**Command:**
```bash
adbjson shell dumpsys wallpaper
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "User 0": "id=1: mWhich=3: mSystemWasBoth=false: mBindSource=INITIALIZE_FALLBACK"
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys wallpaper
```

**Output:**
```
mDefaultWallpaperComponent=ComponentInfo{com.android.systemui/com.android.systemui.wallpapers.ImageWallpaper}
mImageWallpaper=ComponentInfo{com.android.systemui/com.android.systemui.wallpapers.ImageWallpaper}
mCurrentUserIsUnlocked=true
System wallpaper state:
 User 0: id=2: mWhich=3: mSystemWasBoth=false: mBindSource=UNKNOWN
 Display state:
  displayId=0
  mWidth=4848  mHeight=2424
  mPadding=Rect(0, 0 - 0, 0)
  mCropHint=Rect(0, 0 - 0, 0)
  mCropHints={}
  mSampleSize=1.0
  mName=
  mAllowBackup=true
  isColorExtracted=false
  mWallpaperComponent=ComponentInfo{com.android.systemui/com.android.systemui.wallpapers.ImageWallpaper}
  isColorExtracted=false
  mUidToDimAmount:
  Wallpaper connection com.android.server.wallpaper.WallpaperManagerService$WallpaperConnection@40e454:
     mDisplayId=0
     mToken=android.os.Binder@6fc5fbe
     mEngine=android.service.wallpaper.IWallpaperEngine$Stub$Proxy@542777d
    mService=android.service.wallpaper.IWallpaperService$Stub$Proxy@5632672
    mLastDiedTime=-418088
Lock wallpaper state:
 (null entry)
Fallback wallpaper state:
 User 0: id=1: mWhich=3: mSystemWasBoth=false: mBindSource=INITIALIZE_FALLBACK
 Display state:
  displayId=0
  mWidth=4848  mHeight=2424
  mPadding=Rect(0, 0 - 0, 0)
  mCropHint=Rect(0, 0 - 0, 0)
  mCropHints={}
  mSampleSize=1.0
  mName=
  mAllowBackup=false
  isColorExtracted=false
  mWallpaperComponent=ComponentInfo{com.android.systemui/com.android.systemui.wallpapers.GradientColorWallpaper}
  isColorExtracted=false
  mUidToDimAmount:
  Wallpaper connection com.android.server.wallpaper.WallpaperManagerService$WallpaperConnection@7d381b0:
    mInfo.component=ComponentInfo{com.android.systemui/com.android.systemui.wallpapers.GradientColorWallpaper}
    mService=android.service.wallpaper.IWallpaperService$Stub$Proxy@f1a4ac3
    mLastDiedTime=-418088
```
