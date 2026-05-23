# `adbjson shell dumpsys activity activities`

## adbjson

**Command:**
```bash
adbjson shell dumpsys activity activities
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "current_focus": "Window{4879fe9 u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}",
    "focused_app": "ActivityRecord{165322630 u0 com.google.android.apps.nexuslauncher/.NexusLauncherActivity t2}",
    "stack_count": 0,
    "task_count": 0
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys activity activities
```

**Output:**
```
ACTIVITY MANAGER ACTIVITIES (dumpsys activity activities)
Display #0 (activities from top to bottom):
  * Task{14a2074 #1 type=home U=0 visible=true visibleRequested=true mode=fullscreen translucent=false sz=1}
    fullscreenRequestAllowMode=NONE
    * Task{654f047 #2 type=home I=com.google.android.apps.nexuslauncher/.NexusLauncherActivity U=0 rootTaskId=1 visible=true visibleRequested=true mode=fullscreen translucent=false sz=2}
      fullscreenRequestAllowMode=NONE
      isSleeping=false
      topResumedActivity=ActivityRecord{165322630 u0 com.google.android.apps.nexuslauncher/.NexusLauncherActivity t2}
      * Hist  #1: ActivityRecord{165322630 u0 com.google.android.apps.nexuslauncher/.NexusLauncherActivity t2}
        packageName=com.google.android.apps.nexuslauncher processName=com.google.android.apps.nexuslauncher
        launchedFromUid=0 launchedFromPackage=null launchedFromFeature=null userId=0
        app=ProcessRecord{693582f 1043:com.google.android.apps.nexuslauncher/u0a192}
        Intent { act=android.intent.action.MAIN cat=[android.intent.category.HOME] flg=0x10000100 cmp=com.google.android.apps.nexuslauncher/.NexusLauncherActivity (has extras) }
        rootOfTask=true task=Task{654f047 #2 type=home I=com.google.android.apps.nexuslauncher/.NexusLauncherActivity}
        taskAffinity=null
        mActivityComponent=com.google.android.apps.nexuslauncher/.NexusLauncherActivity
        baseDir=/system_ext/priv-app/NexusLauncherRelease/NexusLauncherRelease.apk
        dataDir=/data/user/0/com.google.android.apps.nexuslauncher
        stateNotNeeded=true componentSpecified=true mActivityType=home
        compat={420dpi} theme=0x7f140024
        mLastReportedConfigurations:
          mGlobalConfig={1.0 310mcc260mnc [en_US] ldltr sw411dp w411dp h923dp 420dpi nrml long compactNeeded port finger qwerty/v/v dpad/v winConfig={ mBounds=Rect(0, 0 - 1080, 2424) mAppBounds=Rect(0, 0 - 1080, 2424) mMaxBounds=Rect(0, 0 - 1080, 2424) mDisplayRotation=ROTATION_0 mWindowingMode=fullscreen mActivityType=undefined mAlwaysOnTop=undefined mRotation=ROTATION_0} as.3 s.62 fontWeightAdjustment=0}
          mOverrideConfig={1.0 310mcc260mnc [en_US] ldltr sw411dp w411dp h923dp 420dpi nrml long compactNeeded port finger qwerty/v/v dpad/v winConfig={ mBounds=Rect(0, 0 - 1080, 2424) mAppBounds=Rect(0, 0 - 1080, 2424) mMaxBounds=Rect(0, 0 - 1080, 2424) mDisplayRotation=ROTATION_0 mWindowingMode=fullscreen mActivityType=home mAlwaysOnTop=undefined mRotation=ROTATION_0} as.3 s.7 fontWeightAdjustment=0}
        mLastReportedActivityWindowInfo=ActivityWindowInfo{isEmbedded=false, taskBounds=Rect(0, 0 - 1080, 2424), taskFragmentBounds=Rect(0, 0 - 1080, 2424)}
        CurrentConfiguration={1.0 310mcc260mnc [en_US] ldltr sw411dp w411dp h923dp 420dpi nrml long compactNeeded port finger qwerty/v/v dpad/v winConfig={ mBounds=Rect(0, 0 - 1080, 2424) mAppBounds=Rect(0, 0 - 1080, 2424) mMaxBounds=Rect(0, 0 - 1080, 2424) mDisplayRotation=ROTATION_0 mWindowingMode=fullscreen mActivityType=home mAlwaysOnTop=undefined mRotation=ROTATION_0} as.3 s.8 fontWeightAdjustment=0}
        RequestedOverrideConfiguration={0.0 ?mcc0mnc ?localeList ?layoutDir ?swdp ?wdp ?hdp ?density ?lsize ?long ?round ?ldr ?wideColorGamut ?orien ?uimode ?night ?touch ?keyb/?/? ?nav/? winConfig={ mBounds=Rect(0, 0 - 0, 0) mAppBounds=null mMaxBounds=Rect(0, 0 - 0, 0) mDisplayRotation=undefined mWindowingMode=undefined mActivityType=home mAlwaysOnTop=undefined mRotation=undefined} as.3 ?fontWeightAdjustment}
        taskDescription: label="null" icon=null iconResource=/0 iconFilename=null primaryColor=fff1f0f6
          backgroundColor=ffededf6 statusBarColor=0 navigationBarColor=0
         backgroundColorFloating=fff1f0f6
        launchFailed=false launchCount=1 lastLaunchTime=-6m53s315ms
        mHaveState=false mIcicle=null
        state=RESUMED finishing=false
        keysPaused=false inHistory=true idle=true
        occludesParent=true mNoDisplay=false immersive=false launchMode=2
        mActivityType=home
        windows=[Window{4879fe9 u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}]
        windowType=2
        mOccludesParent=true
        overrideOrientation=SCREEN_ORIENTATION_NOSENSOR
        requestedOrientation=SCREEN_ORIENTATION_NOSENSOR
        mVisibleRequested=true mVisible=true mClientVisible=true reportedDrawn=true reportedVisible=true
        mNumInterestingWindows=2 mNumDrawnWindows=2 allDrawn=true)
        startingData=null firstWindowDrawn=true
        nowVisible=true lastVisibleTime=-6m45s269ms
        resizeMode=RESIZE_MODE_RESIZEABLE
        mLastReportedMultiWindowMode=false mLastReportedPictureInPictureMode=false
        supportsSizeChanges=SIZE_CHANGES_UNSUPPORTED_METADATA
        configChanges=0x4ffb
        neverSandboxDisplayApis=false
        alwaysSandboxDisplayApis=false
        isTransparentPolicyRunning=false
        areBoundsLetterboxed=false
        isLetterboxRunning=false
        mRecreateConfigMask=0
        AppCompatCameraDisplayRotationPolicy:
          isTreatmentEnabledForDisplay=false
        AppCompatCameraSimReqOrientationPolicy:
        CameraStateMonitor:
          activeCameraConnections= mCameraAppInfoSet={ mCameraAppInfoSet={} }
          mAvailableRotateAndCropModesForCamera={}
      * TaskFragment{5f9b6f5 mode=multi-window organizerUid=10192 organizerProc=com.google.android.apps.nexuslauncher}
        mForceHiddenFlags=4
        fullscreenRequestAllowMode=NONE

  * Task{85cf62 #3 name=Bubbles type=undefined U=0 visible=false visibleRequested=false mode=multi-window translucent=true sz=1}
    fullscreenRequestAllowMode=NONE
    mCreatedByOrganizer=true
    mDisablePip=true
    * Task{8e47946 #0 name=Bubbles-visibility-barrier type=undefined U=0 rootTaskId=3 visible=false visibleRequested=false mode=multi-window translucent=true sz=0}
      fullscreenRequestAllowMode=NONE
      mCreatedByOrganizer=true
      isSleeping=false

  * Task{5f558f3 #4 name=SplitRoot type=undefined U=0 visible=false visibleRequested=false mode=fullscreen translucent=true sz=2}
    fullscreenRequestAllowMode=ENTER
    mCreatedByOrganizer=true
    * Task{cc5c029 #6 name=side type=undefined U=0 rootTaskId=4 visible=false visibleRequested=false mode=multi-window translucent=true sz=0}
      mBounds=Rect(0, 2424 - 1080, 3636)
      fullscreenRequestAllowMode=ENTER
      mCreatedByOrganizer=true
      mDisallowOverrideBoundsForChildren=true
      isSleeping=false
    * Task{98319b0 #5 name=main type=undefined U=0 rootTaskId=4 visible=false visibleRequested=false mode=multi-window translucent=true sz=0}
      fullscreenRequestAllowMode=ENTER
      mCreatedByOrganizer=true
      mDisallowOverrideBoundsForChildren=true
      isSleeping=false

  Resumed activities in task display areas (from top to bottom):
    Resumed: ActivityRecord{165322630 u0 com.google.android.apps.nexuslauncher/.NexusLauncherActivity t2}

  ResumedActivity: ActivityRecord{165322630 u0 com.google.android.apps.nexuslauncher/.NexusLauncherActivity t2}

ActivityTaskSupervisor state:
  topDisplayFocusedRootTask=Task{14a2074 #1 type=home}
  Display: mDisplayId=0 (organized)
    init=1080x2424 420dpi mMinSizeOfResizeableTaskDp=220 cur=1080x2424 app=1080x2424 rng=1080x1080-2424x2424
    deferred=false mLayoutNeeded=false
  mLastOrientationSource=WindowedMagnification:0:31@226782821
  deepestLastOrientationSource=ActivityRecord{165322630 u0 com.google.android.apps.nexuslauncher/.NexusLauncherActivity t2}\n\n... (truncated,      399 lines total)
```
