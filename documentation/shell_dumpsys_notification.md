# `adbjson shell dumpsys notification`

## adbjson

**Command:**
```bash
adbjson shell dumpsys notification
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "listener_count": 21,
    "notification_count": 0
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys notification
```

**Output:**
```
Current Notification Manager state:
  Notification List:
    NotificationRecord(0x06538df3: pkg=com.google.android.deskclock user=UserHandle{0} id=2147483642 tag=null importance=3 key=0|com.google.android.deskclock|2147483642|null|10173: Notification(channel=Stopwatch v2 shortcut=null contentView=com.google.android.deskclock/0x7f0e004e vibrate=null sound=null defaults=0 flags=AUTO_CANCEL|LOCAL_ONLY color=0xffb8c6ee groupKey=5 actions=2 vis=PRIVATE))
      uid=10173 userId=0
      opPkg=com.google.android.deskclock
      icon=Icon(typ=RESOURCE pkg=com.google.android.deskclock id=0x7f0803e5)
      flags=AUTO_CANCEL|LOCAL_ONLY
      originalFlags=AUTO_CANCEL|LOCAL_ONLY
      pri=2
      key=0|com.google.android.deskclock|2147483642|null|10173
      seen=false
      groupKey=0|com.google.android.deskclock|g:5
      notification=
            fullscreenIntent=null
            contentIntent=PendingIntent{3c5aab0: PendingIntentRecord{dc989e9 com.google.android.deskclock startActivity (allowlist: 6355546:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            deleteIntent=PendingIntent{94b7d29: PendingIntentRecord{3981cb3 com.google.android.deskclock broadcastIntent (allowlist: 6355546:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            number=0
            groupAlertBehavior=0
            when=1779532067910/1779532067910
            tickerText=null
            vis=0
            contentView=com.google.android.deskclock/0x7f0e004e (0 bytes): android.widget.RemoteViews@bfcf8ae
            bigContentView=null
            headsUpContentView=null
            color=0xffb8c6ee
            timeout=PT72H
            actions={
                [0] "Start" -> PendingIntent{30f164f: PendingIntentRecord{12e4a22 com.google.android.deskclock broadcastIntent (allowlist: 6355546:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
                [1] "Reset" -> PendingIntent{652c1dc: PendingIntentRecord{5f4f270 com.google.android.deskclock broadcastIntent (allowlist: 6355546:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
              }
            extras={
                android.title=null
                android.reduced.images=Boolean (true)
                android.subText=null
                android.template=String (android.app.Notification$DecoratedCustomViewStyle)
                android.showChronometer=Boolean (false)
                android.text=null
                android.progress=Integer (0)
                androidx.core.app.extra.COMPAT_TEMPLATE=String (androidx.core.app.NotificationCompat$DecoratedCustomViewStyle)
                android.progressMax=Integer (0)
                android.appInfo=ApplicationInfo (ApplicationInfo{475c2e5 com.google.android.deskclock})
                android.contains.customView=Boolean (true)
                android.showWhen=Boolean (true)
                android.largeIcon=null
                android.infoText=null
                android.progressIndeterminate=Boolean (false)
                android.remoteInputHistory=null
                android.shortCriticalText=null
            }
      publicNotification=
            None
      stats=SingleNotificationStats{posttimeElapsedMs=6931, posttimeToFirstClickMs=-1, posttimeToDismissMs=-1, airtimeCount=0, airtimeMs=0, currentAirtimeStartElapsedMs=-1, airtimeExpandedMs=0, posttimeToFirstVisibleExpansionMs=-1, currentAirtimeExpandedStartElapsedMs=-1, requestedImportance=4, naturalImportance=3, isNoisy=false}
      mContactAffinity=0.0
      mPackagePriority=0
      mPackageVisibility=-1000
      mSystemImportance=UNSPECIFIED
      mAsstImportance=UNSPECIFIED
      mRuleImportance=UNSPECIFIED
      mImportance=DEFAULT
      mImportanceExplanation=app
      mProposedImportance=UNSPECIFIED
      mIsAppImportanceLocked=false
      mSensitiveContent=false
      mCanceledAfterLifetimeExtension=false
      mIntercept=false
      mHidden==false
      mGlobalSortKey=crtcl=0x0002:intrsv=2:grnk=0x0000:gsmry=1:nsk:rnk=0x0000
      mRankingTimeMs=1779532067942
      mCreationTimeMs=1779532067942
      mVisibleSinceMs=0
      mUpdateTimeMs=1779532067942
      mInterruptionTimeMs=1779532067944
      mSuppressedVisualEffects= 0
      mSound= null
      mVibration= null
      mAttributes= AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null
      mLight= null
      mShowBadge=false
      mColorized=false
      mAllowBubble=false
      isBubble=false
      mIsInterruptive=true
      effectiveNotificationChannel=NotificationChannel{mId='Stopwatch v2', mName=Stopwatch, mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      mAdjustments=[]
      mRuleChannelReason=0
      mRuleSoundReason=0
      mRuleLightReason=0
      shortcut=null found valid? false
      mUserVisOverride=-1000
      hasSummarization=false
      bundleType=0
    NotificationRecord(0x09a3f1ba: pkg=android user=UserHandle{0} id=2345 tag=ChFBbmRyb2lkTG9ja1NjcmVlbhIRTm9TY3JlZW5Mb2NrSXNzdWUYAA== importance=3 key=0|android|2345|ChFBbmRyb2lkTG9ja1NjcmVlbhIRTm9TY3JlZW5Mb2NrSXNzdWUYAA==|1000: Notification(channel=safety_center_recommendation shortcut=null contentView=null vibrate=null sound=null defaults=0 flags=CAN_COLORIZE color=0xff1a73e8 vis=PRIVATE))
      uid=1000 userId=0
      opPkg=android
      icon=Icon(typ=RESOURCE pkg=com.google.android.safetycenter.resources id=0x7f020001)
      flags=CAN_COLORIZE
      originalFlags=CAN_COLORIZE
      pri=0
      key=0|android|2345|ChFBbmRyb2lkTG9ja1NjcmVlbhIRTm9TY3JlZW5Mb2NrSXNzdWUYAA==|1000
      seen=false\n\n... (truncated,     1938 lines total)
```
