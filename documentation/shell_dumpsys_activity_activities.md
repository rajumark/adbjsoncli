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
    "current_focus": "Window{86fb04f u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}",
    "focused_app": "ActivityRecord{101347192 u0 com.google.android.apps.nexuslauncher/.NexusLauncherActivity t2}",
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
  * Task{f21feb6 #1 type=home U=0 visible=true visibleRequested=true mode=fullscreen translucent=false sz=1}
    fullscreenRequestAllowMode=NONE
    * Task{f383651 #2 type=home I=com.google.android.apps.nexuslauncher/.NexusLauncherActivity U=0 rootTaskId=1 visible=true visibleRequested=true mode=fullscreen translucent=false sz=2}
      fullscreenRequestAllowMode=NONE
      isSleeping=false
      topResumedActivity=ActivityRecord{101347192 u0 com.google.android.apps.nexuslauncher/.NexusLauncherActivity t2}
      * Hist  #1: ActivityRecord{101347192 u0 com.google.android.apps.nexuslauncher/.NexusLauncherActivity t2}
        packageName=com.google.android.apps.nexuslauncher processName=com.google.android.apps.nexuslauncher
        launchedFromUid=0 launchedFromPackage=null launchedFromFeature=null userId=0
        app=ProcessRecord{ad84df2 1033:com.google.android.apps.nexuslauncher/u0a192}
        Intent { act=android.intent.action.MAIN cat=[android.intent.category.HOME] flg=0x10000100 cmp=com.google.android.apps.nexuslauncher/.NexusLauncherActivity (has extras) }
        rootOfTask=true task=Task{f383651 #2 type=home I=com.google.android.apps.nexuslauncher/.NexusLauncherActivity}
        taskAffinity=null
        mActivityComponent=com.google.android.apps.nexuslauncher/.NexusLauncherActivity
        baseDir=/system_ext/priv-app/NexusLauncherRelease/NexusLauncherRelease.apk
        dataDir=/data/user/0/com.google.android.apps.nexuslauncher
        stateNotNeeded=true componentSpecified=true mActivityType=home
        compat={420dpi} theme=0x7f140024
        mLastReportedConfigurations:
          mGlobalConfig={1.0 310mcc260mnc [en_US] ldltr sw411dp w411dp h923dp 420dpi nrml long compactNeeded port finger qwerty/v/v dpad/v winConfig={ mBounds=Rect(0, 0 - 1080, 2424) mAppBounds=Rect(0, 0 - 1080, 2424) mMaxBounds=Rect(0, 0 - 1080, 2424) mDisplayRotation=ROTATION_0 mWindowingMode=fullscreen mActivityType=undefined mAlwaysOnTop=undefined mRotation=ROTATION_0} as.3 s.56 fontWeightAdjustment=0}
          mOverrideConfig={1.0 310mcc260mnc [en_US] ldltr sw411dp w411dp h923dp 420dpi nrml long compactNeeded port finger qwerty/v/v dpad/v winConfig={ mBounds=Rect(0, 0 - 1080, 2424) mAppBounds=Rect(0, 0 - 1080, 2424) mMaxBounds=Rect(0, 0 - 1080, 2424) mDisplayRotation=ROTATION_0 mWindowingMode=fullscreen mActivityType=home mAlwaysOnTop=undefined mRotation=ROTATION_0} as.3 s.7 fontWeightAdjustment=0}
        mLastReportedActivityWindowInfo=ActivityWindowInfo{isEmbedded=false, taskBounds=Rect(0, 0 - 1080, 2424), taskFragmentBounds=Rect(0, 0 - 1080, 2424)}
        CurrentConfiguration={1.0 310mcc260mnc [en_US] ldltr sw411dp w411dp h923dp 420dpi nrml long compactNeeded port finger qwerty/v/v dpad/v winConfig={ mBounds=Rect(0, 0 - 1080, 2424) mAppBounds=Rect(0, 0 - 1080, 2424) mMaxBounds=Rect(0, 0 - 1080, 2424) mDisplayRotation=ROTATION_0 mWindowingMode=fullscreen mActivityType=home mAlwaysOnTop=undefined mRotation=ROTATION_0} as.3 s.8 fontWeightAdjustment=0}
        RequestedOverrideConfiguration={0.0 ?mcc0mnc ?localeList ?layoutDir ?swdp ?wdp ?hdp ?density ?lsize ?long ?round ?ldr ?wideColorGamut ?orien ?uimode ?night ?touch ?keyb/?/? ?nav/? winConfig={ mBounds=Rect(0, 0 - 0, 0) mAppBounds=null mMaxBounds=Rect(0, 0 - 0, 0) mDisplayRotation=undefined mWindowingMode=undefined mActivityType=home mAlwaysOnTop=undefined mRotation=undefined} as.3 ?fontWeightAdjustment}
        taskDescription: label="null" icon=null iconResource=/0 iconFilename=null primaryColor=fff1f0f6
          backgroundColor=ffededf6 statusBarColor=0 navigationBarColor=0
         backgroundColorFloating=fff1f0f6
        launchFailed=false launchCount=1 lastLaunchTime=-2m50s934ms
        mHaveState=false mIcicle=null
        state=RESUMED finishing=false
        keysPaused=false inHistory=true idle=true
        occludesParent=true mNoDisplay=false immersive=false launchMode=2
        mActivityType=home
        windows=[Window{86fb04f u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}]
        windowType=2
        mOccludesParent=true
        overrideOrientation=SCREEN_ORIENTATION_NOSENSOR
        requestedOrientation=SCREEN_ORIENTATION_NOSENSOR
        mVisibleRequested=true mVisible=true mClientVisible=true reportedDrawn=true reportedVisible=true
        mNumInterestingWindows=2 mNumDrawnWindows=2 allDrawn=true)
        startingData=null firstWindowDrawn=true
        nowVisible=true lastVisibleTime=-2m35s835ms
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
      * TaskFragment{e1d7385 mode=multi-window organizerUid=10192 organizerProc=com.google.android.apps.nexuslauncher}
        mForceHiddenFlags=4
        fullscreenRequestAllowMode=NONE

  * Task{519e95 #3 name=Bubbles type=undefined U=0 visible=false visibleRequested=false mode=multi-window translucent=true sz=1}
    fullscreenRequestAllowMode=NONE
    mCreatedByOrganizer=true
    mDisablePip=true
    * Task{3760b75 #0 name=Bubbles-visibility-barrier type=undefined U=0 rootTaskId=3 visible=false visibleRequested=false mode=multi-window translucent=true sz=0}
      fullscreenRequestAllowMode=NONE
      mCreatedByOrganizer=true
      isSleeping=false

  * Task{d501baa #4 name=SplitRoot type=undefined U=0 visible=false visibleRequested=false mode=fullscreen translucent=true sz=2}
    fullscreenRequestAllowMode=ENTER
    mCreatedByOrganizer=true
    * Task{fefb338 #6 name=side type=undefined U=0 rootTaskId=4 visible=false visibleRequested=false mode=multi-window translucent=true sz=0}
      mBounds=Rect(0, 2424 - 1080, 3636)
      fullscreenRequestAllowMode=ENTER
      mCreatedByOrganizer=true
      mDisallowOverrideBoundsForChildren=true
      isSleeping=false
    * Task{305c39b #5 name=main type=undefined U=0 rootTaskId=4 visible=false visibleRequested=false mode=multi-window translucent=true sz=0}
      fullscreenRequestAllowMode=ENTER
      mCreatedByOrganizer=true
      mDisallowOverrideBoundsForChildren=true
      isSleeping=false

  Resumed activities in task display areas (from top to bottom):
    Resumed: ActivityRecord{101347192 u0 com.google.android.apps.nexuslauncher/.NexusLauncherActivity t2}

  ResumedActivity: ActivityRecord{101347192 u0 com.google.android.apps.nexuslauncher/.NexusLauncherActivity t2}

ActivityTaskSupervisor state:
  topDisplayFocusedRootTask=Task{f21feb6 #1 type=home}
  Display: mDisplayId=0 (organized)
    init=1080x2424 420dpi mMinSizeOfResizeableTaskDp=220 cur=1080x2424 app=1080x2424 rng=1080x1080-2424x2424
    deferred=false mLayoutNeeded=false
  mLastOrientationSource=WindowedMagnification:0:31@88611326
  deepestLastOrientationSource=ActivityRecord{101347192 u0 com.google.android.apps.nexuslauncher/.NexusLauncherActivity t2}
  overrideConfig={1.0 310mcc260mnc [en_US] ldltr sw411dp w411dp h923dp 420dpi nrml long compactNeeded port finger qwerty/v/v dpad/v winConfig={ mBounds=Rect(0, 0 - 1080, 2424) mAppBounds=Rect(0, 0 - 1080, 2424) mMaxBounds=Rect(0, 0 - 1080, 2424) mDisplayRotation=ROTATION_0 mWindowingMode=fullscreen mActivityType=undefined mAlwaysOnTop=undefined mRotation=ROTATION_0} as.3 s.38 fontWeightAdjustment=0}
  mHasSetIgnoreOrientationRequest=false ignoreOrientationRequest=false
  mSleeping=false mAllSleepTokens=[]
  mLayoutSeq=103
  mImeWindow=Window{4506a47 u0 InputMethod}
  mImeLayeringTarget=Window{86fb04f u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}
  mImeInputTarget=Window{86fb04f u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}
  mImeControlTarget=Window{86fb04f u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}
  ImeContainer
    mNeedsLayer=false
    mImeWindowToken=ImeWindowToken{4934df4 u0 android.os.Binder@7863c7}
  mImeParent=ActivityRecord{101347192 u0 com.google.android.apps.nexuslauncher/.NexusLauncherActivity t2}
  mRemoteInsetsControlTarget=RemoteInsetsControlTarget{11a235a displayId=0 requestedVisibleTypes=503 animatingTypes=0}
  mCurrentFocus=Window{86fb04f u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}
  mFocusedApp=ActivityRecord{101347192 u0 com.google.android.apps.nexuslauncher/.NexusLauncherActivity t2}
  mFixedRotationLaunchingApp=null

  mHoldScreenWindow=null
  mObscuringWindow=Window{1f83155 u0 com.android.systemui.wallpapers.ImageWallpaper}
  mLastWakeLockHoldingWindow=null
  mLastWakeLockObscuringWindow=null

  mIsHardwareRendererOutputDisabled=false
  mAreClientRenderingLimitationsEnabled=false

  displayId=0
  mWallpaperTarget=Window{86fb04f u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}
  mLastWallpaperZoomOut=0.0
  token WallpaperWindowToken{c116005 showWhenLocked=true}:
    mWallpaperX=0.0
    mWallpaperY=0.5
    mWallpaperXStep=0.33333334
    mWallpaperYStep=1.0
    mWallpaperDisplayOffsetX=NA
    mWallpaperDisplayOffsetY=NA

  mSystemGestureExclusion=SkRegion((0,142,1080,2424))

  Display areas in top down Z order:
    * Leaf:36:36
    * OneHanded:34:35
      * Leaf:34:35
    * Leaf:33:33
    * OneHanded:32:32
      * Leaf:32:32
    * WindowedMagnification:0:31 (organized)
      * OneHanded:26:31
        * Leaf:26:31
      * Leaf:24:25
      * OneHanded:0:23
        * Leaf:23:23
        * AppZoomOut:20:22 (organized)
          * Leaf:20:22
        * Leaf:17:19
        * AppZoomOut:16:16 (organized)
          * Leaf:16:16
        * Leaf:15:15
        * AppZoomOut:2:14 (organized)
          * ImePlaceholder:13:14
            * ImeContainer
          * Leaf:3:12
          * DefaultTaskDisplayArea (organized)
        * Leaf:1:1
        * AppZoomOut:0:0 (organized)
          * Leaf:0:0

  Task display areas in top down Z order:
    TaskDisplayArea DefaultTaskDisplayArea
      overrideConfig={0.0 ?mcc0mnc ?localeList ?layoutDir ?swdp ?wdp ?hdp ?density ?lsize ?long ?round ?ldr ?wideColorGamut ?orien ?uimode ?night ?touch ?keyb/?/? ?nav/? winConfig={ mBounds=Rect(0, 0 - 0, 0) mAppBounds=null mMaxBounds=Rect(0, 0 - 0, 0) mDisplayRotation=undefined mWindowingMode=fullscreen mActivityType=undefined mAlwaysOnTop=undefined mRotation=undefined} ?fontWeightAdjustment}
      mPreferredTopFocusableRootTask=Task{f21feb6 #1 type=home}
      mLastFocusedRootTask=Task{f21feb6 #1 type=home}
      Application tokens in top down Z order:
      * Task{f21feb6 #1 type=home U=0 visible=true visibleRequested=true mode=fullscreen translucent=false sz=1}
        bounds=[0,0][1080,2424]
        * Task{f383651 #2 type=home I=com.google.android.apps.nexuslauncher/.NexusLauncherActivity U=0 rootTaskId=1 visible=true visibleRequested=true mode=fullscreen translucent=false sz=2}
          bounds=[0,0][1080,2424]
          * ActivityRecord{101347192 u0 com.google.android.apps.nexuslauncher/.NexusLauncherActivity t2}
          * TaskFragment{e1d7385 mode=multi-window organizerUid=10192 organizerProc=com.google.android.apps.nexuslauncher}
            bounds=[0,0][1080,2424]
      * Task{519e95 #3 name=Bubbles type=undefined U=0 visible=false visibleRequested=false mode=multi-window translucent=true sz=1}
        bounds=[0,0][1080,2424]
        * Task{3760b75 #0 name=Bubbles-visibility-barrier type=undefined U=0 rootTaskId=3 visible=false visibleRequested=false mode=multi-window translucent=true sz=0}
          bounds=[0,0][1080,2424]
      * Task{d501baa #4 name=SplitRoot type=undefined U=0 visible=false visibleRequested=false mode=fullscreen translucent=true sz=2}
        bounds=[0,0][1080,2424]
        * Task{fefb338 #6 name=side type=undefined U=0 rootTaskId=4 visible=false visibleRequested=false mode=multi-window translucent=true sz=0}
          bounds=[0,2424][1080,3636]
        * Task{305c39b #5 name=main type=undefined U=0 rootTaskId=4 visible=false visibleRequested=false mode=multi-window translucent=true sz=0}
          bounds=[0,0][1080,2424]

  rootHomeTask=Task=1

  PinnedTaskController
    mIsImeShowing=false
    mImeHeight=0
    mMinAspectRatio=0.41841003
    mMaxAspectRatio=2.39

  DisplayFrames w=1080 h=2424 r=0

  DisplayPolicy
    mCarDockEnablesAccelerometer=true mDeskDockEnablesAccelerometer=true
    mDockMode=EXTRA_DOCK_STATE_UNDOCKED mLidState=LID_OPEN
    mAwake=true mScreenOnEarly=true mScreenOnFully=true
    mKeyguardDrawComplete=true mWindowManagerDrawComplete=true
    mHdmiPlugged=false
    mLastBehavior=DEFAULT
    mShowingDream=false mDreamingLockscreen=false
    mStatusBar=Window{836a84b u0 StatusBar}
    mExpandedPanel=Window{8491cd2 u0 NotificationShade}
    isKeyguardShowing=false
    mNavigationBar=Window{96172ec u0 Taskbar}
    mNavBarOpacityMode=2
    mNavigationBarCanMove=false
    mHasBottomNavigationBar=true
    mLeftGestureHost=Window{96172ec u0 Taskbar}
    mTopGestureHost=Window{836a84b u0 StatusBar}
    mRightGestureHost=Window{96172ec u0 Taskbar}
    mBottomGestureHost=Window{96172ec u0 Taskbar}
    mFocusedWindow=Window{86fb04f u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}
    mTopFullscreenOpaqueWindowState=Window{86fb04f u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}
    mSystemUiControllingWindow=Window{86fb04f u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}
    mSystemBarColorApps={ActivityRecord{101347192 u0 com.google.android.apps.nexuslauncher/.NexusLauncherActivity t2}}
    mNavBarColorWindowCandidate=Window{86fb04f u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}
    mNavBarBackgroundWindowCandidate=Window{86fb04f u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}
    mLastStatusBarAppearanceRegions=
      AppearanceRegion{ bounds=[0,0][1080,2424]}
    mLastLetterboxDetails=
    mStatusBarBackgroundWindows=
      Window{86fb04f u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}
    mImeInsetsConsumed=false
    mForceShowNavigationBarEnabled=false mAllowLockscreenWhenOn=false
    mAllowsSystemBarRemoteInsetsController=false
    mDecorInsetsInfo:
      ROTATION_0={overrideNonDecorInsets=[0,142][0,63], overrideConfigInsets=[0,142][0,63], overrideNonDecorFrame=[0,142][1080,2361], overrideConfigFrame=[0,142][1080,2361]}
      ROTATION_90={overrideNonDecorInsets=[142,0][0,63], overrideConfigInsets=[142,137][0,63], overrideNonDecorFrame=[142,0][2424,1017], overrideConfigFrame=[142,137][2424,1017]}
      ROTATION_180={overrideNonDecorInsets=[0,0][0,0], overrideConfigInsets=[0,0][0,0], overrideNonDecorFrame=[0,0][0,0], overrideConfigFrame=[0,0][0,0]}
      ROTATION_270={overrideNonDecorInsets=[0,0][0,0], overrideConfigInsets=[0,0][0,0], overrideNonDecorFrame=[0,0][0,0], overrideConfigFrame=[0,0][0,0]}
    SystemGestures:
      mDisplayCutoutTouchableRegionSize=32
      mSwipeStartThreshold=Rect(63, 174 - 63, 63)
      mSwipeDistanceThreshold=63

  DisplayRotation
    mCurrentAppOrientation=SCREEN_ORIENTATION_NOSENSOR
    mLastOrientation=5
    mRotation=0 mDeferredRotationPauseCount=0
    mLandscapeRotation=ROTATION_90 mSeascapeRotation=ROTATION_270
    mPortraitRotation=ROTATION_0 mUpsideDownRotation=ROTATION_180
    mSupportAutoRotation=true
    WindowOrientationListener
      mEnabled=true
      mCurrentRotation=ROTATION_0
      mSensorType=null
      mSensor={Sensor name="Goldfish 3-axis Accelerometer", vendor="The Android Open Source Project", version=1, type=1, maxRange=39.300102, resolution=2.480159E-4, power=3.0, minDelay=10000}
      mRate=2
      AccelSensorJudge
        mProposedRotation=0
        mPredictedRotation=0
        mLastFilteredX=0.0
        mLastFilteredY=9.776321
        mLastFilteredZ=0.812345
        mLastFilteredTimestampNanos=175726405501 (24.146915ms ago)
        mTiltHistory={last: 5.0}
        mFlat=false
        mSwinging=false
        mAccelerating=false
        mOverhead=false
        mTouched=false
        mTiltToleranceConfig=[[-25, 70], [-25, 65], [-25, 60], [-25, 65]]

    mCarDockRotation=-1 mDeskDockRotation=-1
    mUserRotationMode=USER_ROTATION_FREE mUserRotation=ROTATION_0 mCameraRotationMode=0 mAllowAllRotations=unknown
    mDemoHdmiRotation=ROTATION_90 mDemoHdmiRotationLock=false mUndockedHdmiRotation=-1
    mLidOpenRotation=-1
    mFixedToUserRotation=false

  InputConsumers:
    name=recents_animation_input_consumer pid=1033 user=UserHandle{0}

  WindowInsetsStateController
    InsetsState
      mDisplayFrame=Rect(0, 0 - 1080, 2424)
      mDisplayCutout=DisplayCutout{insets=Rect(0, 142 - 0, 0) waterfall=Insets{left=0, top=0, right=0, bottom=0} boundingRect={Bounds=[Rect(0, 0 - 0, 0), Rect(485, 0 - 595, 142), Rect(0, 0 - 0, 0), Rect(0, 0 - 0, 0)]} cutoutPathParserInfo={CutoutPathParserInfo{displayWidth=1080 displayHeight=2424 physicalDisplayWidth=1080 physicalDisplayHeight=2424 density={2.625} cutoutSpec={m 581.5,86.5 a 42,42 0 0 0 -84,0 42,42 0 0 0 84,0 z @left} rotation={0} scale={1.0} physicalPixelDisplaySizeRatio={1.0}}} sideOverrides={}}
      mRoundedCorners=RoundedCorners{[RoundedCorner{position=TopLeft, radius=132, center=Point(132, 132)}, RoundedCorner{position=TopRight, radius=132, center=Point(948, 132)}, RoundedCorner{position=BottomRight, radius=132, center=Point(948, 2292)}, RoundedCorner{position=BottomLeft, radius=132, center=Point(132, 2292)}]}
      mRoundedCornerFrame=Rect(0, 0 - 0, 0)
      mPrivacyIndicatorBounds=PrivacyIndicatorBounds {static bounds=Rect(827, 0 - 1043, 142) rotation=0}
      mDisplayShape=DisplayShape{type=2, displayWidth=1080, displayHeight=2424, physicalPixelDisplaySizeRatio=1.0, rotation=0, offsetX=0, offsetY=0, scale=1.0}
        InsetsSource id=f1880000 type=statusBars frame=[0,0][1080,142] visible=true flags= sideHint=TOP
        InsetsSource id=f1880005 type=mandatorySystemGestures frame=[0,0][1080,174] visible=true flags= sideHint=TOP
        InsetsSource id=f1880006 type=tappableElement frame=[0,0][1080,142] visible=true flags= sideHint=TOP
        InsetsSource id=3 type=ime frame=[0,0][0,0] visible=false flags=INVALID sideHint=NONE
        InsetsSource id=27 type=displayCutout frame=[0,0][1080,142] visible=true flags= sideHint=TOP
        InsetsSource id=70790001 type=navigationBars frame=[0,2361][1080,2424] visible=true flags=SUPPRESS_SCRIM|ANIMATE_RESIZING sideHint=BOTTOM
        InsetsSource id=70790004 type=systemGestures frame=[0,0][78,2424] visible=true flags= sideHint=LEFT
        InsetsSource id=70790005 type=mandatorySystemGestures frame=[0,2340][1080,2424] visible=true flags= sideHint=BOTTOM
        InsetsSource id=70790006 type=tappableElement frame=[0,0][0,0] visible=true flags= sideHint=NONE
        InsetsSource id=70790024 type=systemGestures frame=[1002,0][1080,2424] visible=true flags= sideHint=RIGHT
    Control map:
      Window{86fb04f u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}:
        InsetsSourceControl: {3 mType=ime mSurfacePosition=Point(336, 2298) mInsetsHint=Insets{left=0, top=0, right=0, bottom=0} mLeash=null}
        InsetsSourceControl: {f1880000 mType=statusBars initiallyVisible mSurfacePosition=Point(0, 0) mInsetsHint=Insets{left=0, top=142, right=0, bottom=0} mLeash=Surface(name=Surface(name=836a84b StatusBar#70)/@0x8fa3b41 - animation-leash of insets_animation#105)/@0x2734523}
        InsetsSourceControl: {70790001 mType=navigationBars initiallyVisible mSurfacePosition=Point(0, 2361) mInsetsHint=Insets{left=0, top=0, right=0, bottom=63} mLeash=Surface(name=Surface(name=96172ec Taskbar#78)/@0x992b2c6 - animation-leash of insets_animation#104)/@0x6374a20}
    InsetsSourceProviders:
      InsetsSourceProvider
        mSource=InsetsSource id=70790024 type=systemGestures frame=[1002,0][1080,2424] visible=true flags= sideHint=RIGHT
        mSourceFrame=Rect(1002, 0 - 1080, 2424)
        mIsLeashInitialized=false mHasPendingPosition=false
        mWin=Window{96172ec u0 Taskbar}
      InsetsSourceProvider
        mSource=InsetsSource id=70790006 type=tappableElement frame=[0,0][0,0] visible=true flags= sideHint=NONE
        mSourceFrame=Rect(0, 0 - 0, 0)
        mOverrideFrames={2031=Rect(0, 0 - 0, 0)}
        mIsLeashInitialized=false mHasPendingPosition=false
        mWin=Window{96172ec u0 Taskbar}
      InsetsSourceProvider
        mSource=InsetsSource id=70790005 type=mandatorySystemGestures frame=[0,2340][1080,2424] visible=true flags= sideHint=BOTTOM
        mSourceFrame=Rect(0, 2340 - 1080, 2424)
        mIsLeashInitialized=false mHasPendingPosition=false
        mWin=Window{96172ec u0 Taskbar}
      InsetsSourceProvider
        mSource=InsetsSource id=70790004 type=systemGestures frame=[0,0][78,2424] visible=true flags= sideHint=LEFT
        mSourceFrame=Rect(0, 0 - 78, 2424)
        mIsLeashInitialized=false mHasPendingPosition=false
        mWin=Window{96172ec u0 Taskbar}
      InsetsSourceProvider
        mSource=InsetsSource id=70790001 type=navigationBars frame=[0,2361][1080,2424] visible=true flags=SUPPRESS_SCRIM|ANIMATE_RESIZING sideHint=BOTTOM
        mSourceFrame=Rect(0, 2361 - 1080, 2424)
        mControl=InsetsSourceControl mId=70790001 mType=navigationBars mLeash=Surface(name=Surface(name=96172ec Taskbar#78)/@0x992b2c6 - animation-leash of insets_animation#104)/@0x6374a20 mInitiallyVisible=true mSurfacePosition=Point(0, 2361) mInsetsHint=Insets{left=0, top=0, right=0, bottom=63} mSkipAnimationOnce=false mImeStatsToken=null
        mInsetsHint=Insets{left=0, top=0, right=0, bottom=63}
        mIsLeashInitialized=true mHasPendingPosition=false
        mWin=Window{96172ec u0 Taskbar}
        mAdapter=ControlAdapter mCapturedLeash=Surface(name=Surface(name=96172ec Taskbar#78)/@0x992b2c6 - animation-leash of insets_animation#104)/@0x6374a20
        mControlTarget=Window{86fb04f u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}
      ImeInsetsSourceProvider
        mSource=InsetsSource id=3 type=ime frame=[0,0][0,0] visible=false flags=INVALID sideHint=NONE
        mSourceFrame=Rect(336, 2298 - 744, 2424)
        mControl=InsetsSourceControl mId=3 mType=ime mLeash=Surface(name=Surface(name=4506a47 InputMethod#89)/@0x7b4639d - animation-leash of insets_animation#106)/@0x2e07fd9 mInitiallyVisible=false mSurfacePosition=Point(336, 2298) mInsetsHint=Insets{left=0, top=0, right=0, bottom=0} mSkipAnimationOnce=false mImeStatsToken=null
        mInsetsHint=Insets{left=0, top=0, right=0, bottom=0} stale
        mIsLeashInitialized=true mHasPendingPosition=false
        mWin=Window{4506a47 u0 InputMethod}
        mAdapter=ControlAdapter mCapturedLeash=Surface(name=Surface(name=4506a47 InputMethod#89)/@0x7b4639d - animation-leash of insets_animation#106)/@0x2e07fd9
        mControlTarget=Window{86fb04f u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}
        mImeShowing=false mLastDrawn=false
      InsetsSourceProvider
        mSource=InsetsSource id=f1880006 type=tappableElement frame=[0,0][1080,142] visible=true flags= sideHint=TOP
        mSourceFrame=Rect(0, 0 - 1080, 142)
        mIsLeashInitialized=false mHasPendingPosition=false
        mWin=Window{836a84b u0 StatusBar}
      InsetsSourceProvider
        mSource=InsetsSource id=f1880005 type=mandatorySystemGestures frame=[0,0][1080,174] visible=true flags= sideHint=TOP
        mSourceFrame=Rect(0, 0 - 1080, 174)
        mIsLeashInitialized=false mHasPendingPosition=false
        mWin=Window{836a84b u0 StatusBar}
      InsetsSourceProvider
        mSource=InsetsSource id=f1880000 type=statusBars frame=[0,0][1080,142] visible=true flags= sideHint=TOP
        mSourceFrame=Rect(0, 0 - 1080, 142)
        mControl=InsetsSourceControl mId=f1880000 mType=statusBars mLeash=Surface(name=Surface(name=836a84b StatusBar#70)/@0x8fa3b41 - animation-leash of insets_animation#105)/@0x2734523 mInitiallyVisible=true mSurfacePosition=Point(0, 0) mInsetsHint=Insets{left=0, top=142, right=0, bottom=0} mSkipAnimationOnce=false mImeStatsToken=null
        mInsetsHint=Insets{left=0, top=142, right=0, bottom=0}
        mIsLeashInitialized=true mHasPendingPosition=false
        mWin=Window{836a84b u0 StatusBar}
        mAdapter=ControlAdapter mCapturedLeash=Surface(name=Surface(name=836a84b StatusBar#70)/@0x8fa3b41 - animation-leash of insets_animation#105)/@0x2734523
        mControlTarget=Window{86fb04f u0 com.google.android.apps.nexuslauncher/com.google.android.apps.nexuslauncher.NexusLauncherActivity}
  InsetsPolicy
    status: WINDOW_STATE_SHOWING
    nav: WINDOW_STATE_SHOWING

  DisplayWindowSettingsProvider
    SettingsEntry{mWindowingMode=0, mUserRotationMode=null, mUserRotation=null, mForcedWidth=0, mForcedHeight=0, mForcedDensity=0, mForcedDensityRatio=0.0, mForcedScalingMode=null, mRemoveContentMode=0, mShouldShowWithInsecureKeyguard=null, mShouldShowSystemDecors=null, mIsHomeSupported=null, mShouldShowIme=null, mFixedToUserRotation=null, mIgnoreOrientationRequest=null, mIgnoreDisplayCutout=null, mDontMoveToTop=null, mForceAppsUniversalResizable=null, mCanStealTopFocus=null}

  KeyguardController:
    mKeyguardShowing=false
    mAodShowing=false
    mKeyguardGoingAway=false
   KeyguardShowing=false AodShowing=false KeyguardGoingAway=false DismissalRequested=false  Occluded=false DismissingKeyguardActivity=null TurnScreenOnActivity=null at display=0
    mDismissalRequested=false

  LockTaskController:
    mLockTaskModeState=NONE
    mLockTaskModeTasks=
    mLockTaskPackages (userId:packages)=

  mCurTaskIdForUser={0=39}
  mUserRootTaskInFront={}
  mVisibilityTransactionDepth=0
  isHomeRecentsComponent=true
  mNoHistoryActivities=[]

  TaskOrganizerController:
      android.window.ITaskOrganizer$Stub$Proxy@39ee457 uid=10195 (active):
        (fullscreen) Task{f21feb6 #1 type=home}
        (multi-window) Task{519e95 #3 name=Bubbles type=undefined}
        (fullscreen) Task{d501baa #4 name=SplitRoot type=undefined}
        (multi-window) Task{305c39b #5 name=main type=undefined}
        (multi-window) Task{fefb338 #6 name=side type=undefined}
        (multi-window) Task{3760b75 #0 name=Bubbles-visibility-barrier type=undefined}

  VisibleActivityProcess:[ ProcessRecord{ad84df2 1033:com.google.android.apps.nexuslauncher/u0a192}]
  NumNonAppVisibleWindowUidMap:[ 10195:4 10192:1]
```
