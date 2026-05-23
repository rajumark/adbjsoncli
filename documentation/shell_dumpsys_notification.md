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
    NotificationRecord(0x092c03d3: pkg=com.google.android.deskclock user=UserHandle{0} id=2147483642 tag=null importance=3 key=0|com.google.android.deskclock|2147483642|null|10173: Notification(channel=Stopwatch v2 shortcut=null contentView=com.google.android.deskclock/0x7f0e004e vibrate=null sound=null defaults=0 flags=AUTO_CANCEL|LOCAL_ONLY color=0xffbac5ee groupKey=5 actions=2 vis=PRIVATE))
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
            contentIntent=PendingIntent{91a1b10: PendingIntentRecord{d46d52d com.google.android.deskclock startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            deleteIntent=PendingIntent{2bc9609: PendingIntentRecord{24b7d57 com.google.android.deskclock broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            number=0
            groupAlertBehavior=0
            when=1779530861063/1779530861063
            tickerText=null
            vis=0
            contentView=com.google.android.deskclock/0x7f0e004e (0 bytes): android.widget.RemoteViews@3f8d00e
            bigContentView=null
            headsUpContentView=null
            color=0xffbac5ee
            timeout=PT72H
            actions={
                [0] "Start" -> PendingIntent{9290a2f: PendingIntentRecord{29f1044 com.google.android.deskclock broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
                [1] "Reset" -> PendingIntent{6edd83c: PendingIntentRecord{9eaaa62 com.google.android.deskclock broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
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
                android.appInfo=ApplicationInfo (ApplicationInfo{d75df1a com.google.android.deskclock})
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
      stats=SingleNotificationStats{posttimeElapsedMs=12670, posttimeToFirstClickMs=-1, posttimeToDismissMs=-1, airtimeCount=0, airtimeMs=0, currentAirtimeStartElapsedMs=-1, airtimeExpandedMs=0, posttimeToFirstVisibleExpansionMs=-1, currentAirtimeExpandedStartElapsedMs=-1, requestedImportance=4, naturalImportance=3, isNoisy=false}
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
      mRankingTimeMs=1779530861067
      mCreationTimeMs=1779530861067
      mVisibleSinceMs=0
      mUpdateTimeMs=1779530861067
      mInterruptionTimeMs=1779530861069
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
    NotificationRecord(0x0d461a4b: pkg=android user=UserHandle{0} id=2345 tag=ChFBbmRyb2lkTG9ja1NjcmVlbhIRTm9TY3JlZW5Mb2NrSXNzdWUYAA== importance=3 key=0|android|2345|ChFBbmRyb2lkTG9ja1NjcmVlbhIRTm9TY3JlZW5Mb2NrSXNzdWUYAA==|1000: Notification(channel=safety_center_recommendation shortcut=null contentView=null vibrate=null sound=null defaults=0 flags=CAN_COLORIZE color=0xff1a73e8 vis=PRIVATE))
      uid=1000 userId=0
      opPkg=android
      icon=Icon(typ=RESOURCE pkg=com.google.android.safetycenter.resources id=0x7f020001)
      flags=CAN_COLORIZE
      originalFlags=CAN_COLORIZE
      pri=0
      key=0|android|2345|ChFBbmRyb2lkTG9ja1NjcmVlbhIRTm9TY3JlZW5Mb2NrSXNzdWUYAA==|1000
      seen=false
      groupKey=0|android|2345|ChFBbmRyb2lkTG9ja1NjcmVlbhIRTm9TY3JlZW5Mb2NrSXNzdWUYAA==|1000
      notification=
            fullscreenIntent=null
            contentIntent=PendingIntent{3de3c28: PendingIntentRecord{c9748cc android startActivity}}
            deleteIntent=PendingIntent{ae3fd41: PendingIntentRecord{e605715 android broadcastIntent}}
            number=0
            groupAlertBehavior=0
            when=1779530853648/1779530853648
            tickerText=null
            vis=0
            contentView=null
            bigContentView=null
            headsUpContentView=null
            color=0xff1a73e8
            timeout=PT72H
            extras={
                android.title=String [length=17]
                android.reduced.images=Boolean (true)
                android.template=String (android.app.Notification$BigTextStyle)
                android.text=String [length=68]
                android.appInfo=ApplicationInfo (ApplicationInfo{b82d2e6 android})
                android.showWhen=Boolean (true)
                android.substName=String (Security & privacy)
                android.bigText=String [length=68]
            }
      publicNotification=
            None
      stats=SingleNotificationStats{posttimeElapsedMs=5253, posttimeToFirstClickMs=-1, posttimeToDismissMs=-1, airtimeCount=0, airtimeMs=0, currentAirtimeStartElapsedMs=-1, airtimeExpandedMs=0, posttimeToFirstVisibleExpansionMs=-1, currentAirtimeExpandedStartElapsedMs=-1, requestedImportance=3, naturalImportance=3, isNoisy=true}
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
      mGlobalSortKey=crtcl=0x0002:intrsv=2:grnk=0x0001:gsmry=1:nsk:rnk=0x0001
      mRankingTimeMs=1779530853650
      mCreationTimeMs=1779530853650
      mVisibleSinceMs=0
      mUpdateTimeMs=1779530853650
      mInterruptionTimeMs=1779530853652
      mSuppressedVisualEffects= 0
      mSound= content://settings/system/notification_sound
      mVibration= null
      mAttributes= AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null
      mLight= null
      mShowBadge=true
      mColorized=false
      mAllowBubble=false
      isBubble=false
      mIsInterruptive=true
      effectiveNotificationChannel=NotificationChannel{mId='safety_center_recommendation', mName=Warnings, mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='safety_center_channels', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      mAdjustments=[]
      mRuleChannelReason=0
      mRuleSoundReason=0
      mRuleLightReason=0
      shortcut=null found valid? false
      mUserVisOverride=-1000
      hasSummarization=false
      bundleType=0
    NotificationRecord(0x00cc11cf: pkg=com.google.android.apps.messaging user=UserHandle{0} id=7 tag=incoming_message_group_key importance=4 key=0|com.google.android.apps.messaging|7|incoming_message_group_key|10154: Notification(channel=bugle_default_channel shortcut=null contentView=null vibrate=null sound=null defaults=0 flags=ONLY_ALERT_ONCE|AUTO_CANCEL|GROUP_SUMMARY color=0xff1a73e8 category=msg groupKey=incoming_message_group_key vis=PRIVATE))
      uid=10154 userId=0
      opPkg=com.google.android.apps.messaging
      icon=Icon(typ=RESOURCE pkg=com.google.android.apps.messaging id=0x7f0806b2)
      flags=ONLY_ALERT_ONCE|AUTO_CANCEL|GROUP_SUMMARY
      originalFlags=ONLY_ALERT_ONCE|AUTO_CANCEL|GROUP_SUMMARY
      pri=2
      key=0|com.google.android.apps.messaging|7|incoming_message_group_key|10154
      seen=false
      groupKey=0|com.google.android.apps.messaging|g:incoming_message_group_key
      notification=
            fullscreenIntent=null
            contentIntent=PendingIntent{96ad75c: PendingIntentRecord{64f6036 com.google.android.apps.messaging startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            deleteIntent=PendingIntent{d43ea65: PendingIntentRecord{2202dd1 com.google.android.apps.messaging broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            number=0
            groupAlertBehavior=2
            when=1779198318544/1779198318544
            tickerText=null
            vis=0
            contentView=null
            bigContentView=null
            headsUpContentView=null
            color=0xff1a73e8
            timeout=PT72H
            extras={
                android.title=null
                android.reduced.images=Boolean (true)
                android.subText=null
                android.showChronometer=Boolean (false)
                android.text=null
                android.progress=Integer (0)
                android.progressMax=Integer (0)
                android.appInfo=ApplicationInfo (ApplicationInfo{673c33a com.google.android.apps.messaging})
                android.showWhen=Boolean (true)
                android.largeIcon=null
                android.infoText=null
                android.progressIndeterminate=Boolean (false)
                android.remoteInputHistory=null
                android.shortCriticalText=null
            }
      publicNotification=
            None
      stats=SingleNotificationStats{posttimeElapsedMs=33069, posttimeToFirstClickMs=-1, posttimeToDismissMs=-1, airtimeCount=0, airtimeMs=0, currentAirtimeStartElapsedMs=-1, airtimeExpandedMs=0, posttimeToFirstVisibleExpansionMs=-1, currentAirtimeExpandedStartElapsedMs=-1, requestedImportance=4, naturalImportance=4, isNoisy=true}
      mContactAffinity=0.0
      mPackagePriority=0
      mPackageVisibility=-1000
      mSystemImportance=UNSPECIFIED
      mAsstImportance=UNSPECIFIED
      mRuleImportance=UNSPECIFIED
      mImportance=HIGH
      mImportanceExplanation=app
      mProposedImportance=UNSPECIFIED
      mIsAppImportanceLocked=false
      mSensitiveContent=false
      mCanceledAfterLifetimeExtension=false
      mIntercept=false
      mHidden==false
      mGlobalSortKey=crtcl=0x0002:intrsv=2:grnk=0x0002:gsmry=0:nsk:rnk=0x0003
      mRankingTimeMs=1779198318544
      mCreationTimeMs=1779530881264
      mVisibleSinceMs=0
      mUpdateTimeMs=1779530881264
      mInterruptionTimeMs=1779530881264
      mSuppressedVisualEffects= 0
      mSound= content://settings/system/notification_sound
      mVibration= Composed{segments=[Step{amplitude=0.0, duration=0}, Step{amplitude=-1.0, duration=350}, Step{amplitude=0.0, duration=250}, Step{amplitude=-1.0, duration=350}], repeat=-1}
      mAttributes= AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null
      mLight= Light{color=-1, onMs=500, offMs=2000}
      mShowBadge=true
      mColorized=false
      mAllowBubble=false
      isBubble=false
      mIsInterruptive=false
      effectiveNotificationChannel=NotificationChannel{mId='bugle_default_channel', mName=Incoming messages, mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='_bugle_default_settings_group', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      mAdjustments=[]
      mRuleChannelReason=0
      mRuleSoundReason=0
      mRuleLightReason=0
      shortcut=null found valid? false
      mUserVisOverride=-1000
      hasSummarization=false
      bundleType=0
    NotificationRecord(0x0da1fbeb: pkg=com.google.android.apps.messaging user=UserHandle{0} id=2 tag=com.google.android.apps.messaging:incoming_message:3 importance=4 key=0|com.google.android.apps.messaging|2|com.google.android.apps.messaging:incoming_message:3|10154: Notification(channel=bugle_default_channel shortcut=3 contentView=null vibrate=null sound=null tick defaults=0 flags=AUTO_CANCEL color=0xff1a73e8 category=msg groupKey=incoming_message_group_key actions=2 vis=PRIVATE))
      uid=10154 userId=0
      opPkg=com.google.android.apps.messaging
      icon=Icon(typ=RESOURCE pkg=com.google.android.apps.messaging id=0x7f0806b2)
      flags=AUTO_CANCEL
      originalFlags=AUTO_CANCEL
      pri=0
      key=0|com.google.android.apps.messaging|2|com.google.android.apps.messaging:incoming_message:3|10154
      seen=false
      groupKey=0|com.google.android.apps.messaging|g:incoming_message_group_key
      notification=
            fullscreenIntent=null
            contentIntent=PendingIntent{b7c0d48: PendingIntentRecord{dfb8499 com.google.android.apps.messaging startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            deleteIntent=PendingIntent{d8dc7e1: PendingIntentRecord{958cb9d com.google.android.apps.messaging broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            number=1
            groupAlertBehavior=1
            when=1779198318544/1779198318544
            tickerText=(650) 55...
            vis=0
            contentView=null
            bigContentView=null
            headsUpContentView=null
            color=0xff1a73e8
            timeout=PT72H
            actions={
                [0] "Mark as read" -> PendingIntent{14ad906: PendingIntentRecord{3fc69e0 com.google.android.apps.messaging broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
                [1] "Reply" -> PendingIntent{ba26bc7: PendingIntentRecord{71e0be3 com.google.android.apps.messaging broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
              }
            extras={
                android.title=String [length=14]
                android.hiddenConversationTitle=null
                android.reduced.images=Boolean (true)
                extra_im_notification_message_ids=ArrayList ([6])
                android.subText=null
                android.template=String (android.app.Notification$MessagingStyle)
                android.showChronometer=Boolean (false)
                extra_im_notification_earliest_timestamp=Long (1779198318544)
                android.text=String [length=5]
                android.progress=Integer (0)
                androidx.core.app.extra.COMPAT_TEMPLATE=String (androidx.core.app.NotificationCompat$MessagingStyle)
                android.progressMax=Integer (0)
                android.selfDisplayName=String [length=3]
                android.conversationUnreadMessageCount=Integer (0)
                android.appInfo=ApplicationInfo (ApplicationInfo{3e8b5f4 com.google.android.apps.messaging})
                android.messages=Bundle[] (1)
                android.showWhen=Boolean (true)
                extra_im_notification_channel_importance=Integer (4)
                android.largeIcon=null
                android.messagingStyleUser=Bundle (Bundle[mParcelledData.dataSize=492])
                android.messagingUser=Person (android.app.Person@ca984bdf)
                android.infoText=null
                android.wearable.EXTENSIONS=Bundle (Bundle[{actions=[android.app.Notification$Action@82b7f92, android.app.Notification$Action@84ffd63, android.app.Notification$Action@6d53d60, android.app.Notification$Action@9be9219]}])
                extra_im_notification_latest_timestamp=Long (1779198318544)
                android.progressIndeterminate=Boolean (false)
                extra_im_notification_participant_normalized_destination=String [length=12]
                android.remoteInputHistory=null
                android.shortCriticalText=null
                extra_im_notification_conversation_id=String [length=1]
                extra_im_is_notification_from_rbm=Boolean (false)
                extra_im_notification_latest_analytics_id=Long (2536288054746613231)
                extra_im_notification_latest_rcs_message_id=String [length=0]
                android.isGroupConversation=Boolean (false)
            }
      publicNotification=
            None
      stats=SingleNotificationStats{posttimeElapsedMs=33069, posttimeToFirstClickMs=-1, posttimeToDismissMs=-1, airtimeCount=0, airtimeMs=0, currentAirtimeStartElapsedMs=-1, airtimeExpandedMs=0, posttimeToFirstVisibleExpansionMs=-1, currentAirtimeExpandedStartElapsedMs=-1, requestedImportance=3, naturalImportance=4, isNoisy=true}
      mContactAffinity=0.0
      mPackagePriority=0
      mPackageVisibility=-1000
      mSystemImportance=UNSPECIFIED
      mAsstImportance=UNSPECIFIED
      mRuleImportance=UNSPECIFIED
      mImportance=HIGH
      mImportanceExplanation=app
      mProposedImportance=UNSPECIFIED
      mIsAppImportanceLocked=false
      mSensitiveContent=false
      mCanceledAfterLifetimeExtension=false
      mIntercept=false
      mHidden==false
      mGlobalSortKey=crtcl=0x0002:intrsv=2:grnk=0x0002:gsmry=1:nsk:rnk=0x0002
      mRankingTimeMs=1779198318544
      mCreationTimeMs=1779530881264
      mVisibleSinceMs=0
      mUpdateTimeMs=1779530881264
      mInterruptionTimeMs=1779530881468
      mSuppressedVisualEffects= 0
      mSound= content://settings/system/notification_sound
      mVibration= Composed{segments=[Step{amplitude=0.0, duration=0}, Step{amplitude=-1.0, duration=350}, Step{amplitude=0.0, duration=250}, Step{amplitude=-1.0, duration=350}], repeat=-1}
      mAttributes= AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null
      mLight= Light{color=-1, onMs=500, offMs=2000}
      mShowBadge=true
      mColorized=false
      mAllowBubble=false
      isBubble=false
      mIsInterruptive=true
      effectiveNotificationChannel=NotificationChannel{mId='bugle_default_channel', mName=Incoming messages, mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='_bugle_default_settings_group', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      mAdjustments=[]
      mRuleChannelReason=0
      mRuleSoundReason=0
      mRuleLightReason=0
      shortcut=3 found valid? true
      mUserVisOverride=-1000
      hasSummarization=false
      bundleType=0
    NotificationRecord(0x0efe02de: pkg=com.google.android.apps.messaging user=UserHandle{0} id=2 tag=com.google.android.apps.messaging:incoming_message:2 importance=4 key=0|com.google.android.apps.messaging|2|com.google.android.apps.messaging:incoming_message:2|10154: Notification(channel=bugle_default_channel shortcut=2 contentView=null vibrate=null sound=null tick defaults=0 flags=AUTO_CANCEL color=0xff1a73e8 category=msg groupKey=incoming_message_group_key actions=2 vis=PRIVATE))
      uid=10154 userId=0
      opPkg=com.google.android.apps.messaging
      icon=Icon(typ=RESOURCE pkg=com.google.android.apps.messaging id=0x7f0806b2)
      flags=AUTO_CANCEL
      originalFlags=AUTO_CANCEL
      pri=0
      key=0|com.google.android.apps.messaging|2|com.google.android.apps.messaging:incoming_message:2|10154
      seen=false
      groupKey=0|com.google.android.apps.messaging|g:incoming_message_group_key
      notification=
            fullscreenIntent=null
            contentIntent=PendingIntent{cbb0cbf: PendingIntentRecord{7813d2d com.google.android.apps.messaging startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            deleteIntent=PendingIntent{440cf8c: PendingIntentRecord{be9e4e5 com.google.android.apps.messaging broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            number=1
            groupAlertBehavior=1
            when=1779198313253/1779198313253
            tickerText=65055541...
            vis=0
            contentView=null
            bigContentView=null
            headsUpContentView=null
            color=0xff1a73e8
            timeout=PT72H
            actions={
                [0] "Mark as read" -> PendingIntent{11d36d5: PendingIntentRecord{449e84f com.google.android.apps.messaging broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
                [1] "Reply" -> PendingIntent{de36eea: PendingIntentRecord{78f262 com.google.android.apps.messaging broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
              }
            extras={
                android.title=String [length=11]
                android.hiddenConversationTitle=null
                android.reduced.images=Boolean (true)
                extra_im_notification_message_ids=ArrayList ([4])
                android.subText=null
                android.template=String (android.app.Notification$MessagingStyle)
                android.showChronometer=Boolean (false)
                extra_im_notification_earliest_timestamp=Long (1779198313253)
                android.text=String [length=32]
                android.progress=Integer (0)
                androidx.core.app.extra.COMPAT_TEMPLATE=String (androidx.core.app.NotificationCompat$MessagingStyle)
                android.progressMax=Integer (0)
                android.selfDisplayName=String [length=3]
                android.conversationUnreadMessageCount=Integer (0)
                android.appInfo=ApplicationInfo (ApplicationInfo{89db5db com.google.android.apps.messaging})
                android.messages=Bundle[] (1)
                android.showWhen=Boolean (true)
                extra_im_notification_channel_importance=Integer (4)
                android.largeIcon=null
                android.messagingStyleUser=Bundle (Bundle[mParcelledData.dataSize=492])
                android.messagingUser=Person (android.app.Person@12ca7dba)
                android.infoText=null
                android.wearable.EXTENSIONS=Bundle (Bundle[{actions=[android.app.Notification$Action@6876b51, android.app.Notification$Action@c018fb6, android.app.Notification$Action@337d4b7, android.app.Notification$Action@e1f8424]}])
                extra_im_notification_latest_timestamp=Long (1779198313253)
                android.progressIndeterminate=Boolean (false)
                extra_im_notification_participant_normalized_destination=String [length=11]
                android.remoteInputHistory=null
                android.shortCriticalText=null
                extra_im_notification_conversation_id=String [length=1]
                extra_im_is_notification_from_rbm=Boolean (false)
                extra_im_notification_latest_analytics_id=Long (-2249342395652203734)
                extra_im_notification_latest_rcs_message_id=String [length=0]
                android.isGroupConversation=Boolean (false)
            }
      publicNotification=
            None
      stats=SingleNotificationStats{posttimeElapsedMs=33068, posttimeToFirstClickMs=-1, posttimeToDismissMs=-1, airtimeCount=0, airtimeMs=0, currentAirtimeStartElapsedMs=-1, airtimeExpandedMs=0, posttimeToFirstVisibleExpansionMs=-1, currentAirtimeExpandedStartElapsedMs=-1, requestedImportance=3, naturalImportance=4, isNoisy=true}
      mContactAffinity=0.0
      mPackagePriority=0
      mPackageVisibility=-1000
      mSystemImportance=UNSPECIFIED
      mAsstImportance=UNSPECIFIED
      mRuleImportance=UNSPECIFIED
      mImportance=HIGH
      mImportanceExplanation=app
      mProposedImportance=UNSPECIFIED
      mIsAppImportanceLocked=false
      mSensitiveContent=false
      mCanceledAfterLifetimeExtension=false
      mIntercept=false
      mHidden==false
      mGlobalSortKey=crtcl=0x0002:intrsv=2:grnk=0x0002:gsmry=1:nsk:rnk=0x0004
      mRankingTimeMs=1779198313253
      mCreationTimeMs=1779530881263
      mVisibleSinceMs=0
      mUpdateTimeMs=1779530881263
      mInterruptionTimeMs=1779530881467
      mSuppressedVisualEffects= 0
      mSound= content://settings/system/notification_sound
      mVibration= Composed{segments=[Step{amplitude=0.0, duration=0}, Step{amplitude=-1.0, duration=350}, Step{amplitude=0.0, duration=250}, Step{amplitude=-1.0, duration=350}], repeat=-1}
      mAttributes= AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null
      mLight= Light{color=-1, onMs=500, offMs=2000}
      mShowBadge=true
      mColorized=false
      mAllowBubble=false
      isBubble=false
      mIsInterruptive=true
      effectiveNotificationChannel=NotificationChannel{mId='bugle_default_channel', mName=Incoming messages, mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='_bugle_default_settings_group', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      mAdjustments=[]
      mRuleChannelReason=0
      mRuleSoundReason=0
      mRuleLightReason=0
      shortcut=2 found valid? true
      mUserVisOverride=-1000
      hasSummarization=false
      bundleType=0
    NotificationRecord(0x05866b8d: pkg=com.google.android.apps.messaging user=UserHandle{0} id=2 tag=com.google.android.apps.messaging:incoming_message:1 importance=4 key=0|com.google.android.apps.messaging|2|com.google.android.apps.messaging:incoming_message:1|10154: Notification(channel=bugle_default_channel shortcut=1 contentView=null vibrate=null sound=null tick defaults=0 flags=AUTO_CANCEL color=0xff1a73e8 category=msg groupKey=incoming_message_group_key actions=2 vis=PRIVATE))
      uid=10154 userId=0
      opPkg=com.google.android.apps.messaging
      icon=Icon(typ=RESOURCE pkg=com.google.android.apps.messaging id=0x7f0806b2)
      flags=AUTO_CANCEL
      originalFlags=AUTO_CANCEL
      pri=0
      key=0|com.google.android.apps.messaging|2|com.google.android.apps.messaging:incoming_message:1|10154
      seen=false
      groupKey=0|com.google.android.apps.messaging|g:incoming_message_group_key
      notification=
            fullscreenIntent=null
            contentIntent=PendingIntent{b0ef142: PendingIntentRecord{cd52b80 com.google.android.apps.messaging startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            deleteIntent=PendingIntent{deb0553: PendingIntentRecord{34b2abd com.google.android.apps.messaging broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            number=1
            groupAlertBehavior=1
            when=1779198310201/1779198310201
            tickerText=65055541...
            vis=0
            contentView=null
            bigContentView=null
            headsUpContentView=null
            color=0xff1a73e8
            timeout=PT72H
            actions={
                [0] "Mark as read" -> PendingIntent{9e2be90: PendingIntentRecord{934f8b2 com.google.android.apps.messaging broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
                [1] "Reply" -> PendingIntent{efb3389: PendingIntentRecord{a84d467 com.google.android.apps.messaging broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
              }
            extras={
                android.title=String [length=12]
                android.hiddenConversationTitle=null
                android.reduced.images=Boolean (true)
                extra_im_notification_message_ids=ArrayList ([2])
                android.subText=null
                android.template=String (android.app.Notification$MessagingStyle)
                android.showChronometer=Boolean (false)
                extra_im_notification_earliest_timestamp=Long (1779198310201)
                android.text=String [length=32]
                android.progress=Integer (0)
                androidx.core.app.extra.COMPAT_TEMPLATE=String (androidx.core.app.NotificationCompat$MessagingStyle)
                android.progressMax=Integer (0)
                android.selfDisplayName=String [length=3]
                android.conversationUnreadMessageCount=Integer (0)
                android.appInfo=ApplicationInfo (ApplicationInfo{f9edf8e com.google.android.apps.messaging})
                android.messages=Bundle[] (1)
                android.showWhen=Boolean (true)
                extra_im_notification_channel_importance=Integer (4)
                android.largeIcon=null
                android.messagingStyleUser=Bundle (Bundle[mParcelledData.dataSize=492])
                android.messagingUser=Person (android.app.Person@d5d95571)
                android.infoText=null
                android.wearable.EXTENSIONS=Bundle (Bundle[{actions=[android.app.Notification$Action@25d33bc, android.app.Notification$Action@6da7f45, android.app.Notification$Action@22d669a, android.app.Notification$Action@7b3cbcb]}])
                extra_im_notification_latest_timestamp=Long (1779198310201)
                android.progressIndeterminate=Boolean (false)
                extra_im_notification_participant_normalized_destination=String [length=12]
                android.remoteInputHistory=null
                android.shortCriticalText=null
                extra_im_notification_conversation_id=String [length=1]
                extra_im_is_notification_from_rbm=Boolean (false)
                extra_im_notification_latest_analytics_id=Long (-7877412699053206930)
                extra_im_notification_latest_rcs_message_id=String [length=0]
                android.isGroupConversation=Boolean (false)
            }
      publicNotification=
            None
      stats=SingleNotificationStats{posttimeElapsedMs=33066, posttimeToFirstClickMs=-1, posttimeToDismissMs=-1, airtimeCount=0, airtimeMs=0, currentAirtimeStartElapsedMs=-1, airtimeExpandedMs=0, posttimeToFirstVisibleExpansionMs=-1, currentAirtimeExpandedStartElapsedMs=-1, requestedImportance=3, naturalImportance=4, isNoisy=true}
      mContactAffinity=0.0
      mPackagePriority=0
      mPackageVisibility=-1000
      mSystemImportance=UNSPECIFIED
      mAsstImportance=UNSPECIFIED
      mRuleImportance=UNSPECIFIED
      mImportance=HIGH
      mImportanceExplanation=app
      mProposedImportance=UNSPECIFIED
      mIsAppImportanceLocked=false
      mSensitiveContent=false
      mCanceledAfterLifetimeExtension=false
      mIntercept=false
      mHidden==false
      mGlobalSortKey=crtcl=0x0002:intrsv=2:grnk=0x0002:gsmry=1:nsk:rnk=0x0005
      mRankingTimeMs=1779198310201
      mCreationTimeMs=1779530881261
      mVisibleSinceMs=0
      mUpdateTimeMs=1779530881261
      mInterruptionTimeMs=1779530881465
      mSuppressedVisualEffects= 0
      mSound= content://settings/system/notification_sound
      mVibration= Composed{segments=[Step{amplitude=0.0, duration=0}, Step{amplitude=-1.0, duration=350}, Step{amplitude=0.0, duration=250}, Step{amplitude=-1.0, duration=350}], repeat=-1}
      mAttributes= AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null
      mLight= Light{color=-1, onMs=500, offMs=2000}
      mShowBadge=true
      mColorized=false
      mAllowBubble=false
      isBubble=false
      mIsInterruptive=true
      effectiveNotificationChannel=NotificationChannel{mId='bugle_default_channel', mName=Incoming messages, mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='_bugle_default_settings_group', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      mAdjustments=[]
      mRuleChannelReason=0
      mRuleSoundReason=0
      mRuleLightReason=0
      shortcut=1 found valid? true
      mUserVisOverride=-1000
      hasSummarization=false
      bundleType=0
    NotificationRecord(0x09a7cfa8: pkg=com.google.android.dialer user=UserHandle{0} id=1 tag=GroupSummary_MissedCall importance=3 key=0|com.google.android.dialer|1|GroupSummary_MissedCall|10147: Notification(channel=phone_missed_call shortcut=null contentView=null vibrate=null sound=null defaults=LIGHTS flags=SHOW_LIGHTS|ONLY_ALERT_ONCE|AUTO_CANCEL|GROUP_SUMMARY color=0xffc0c6dc groupKey=MissedCallGroup vis=PRIVATE publicVersion=Notification(channel=phone_missed_call shortcut=null contentView=null vibrate=null sound=null defaults=0 flags=ONLY_ALERT_ONCE|AUTO_CANCEL color=0xffc0c6dc groupKey=MissedCallGroup vis=PRIVATE)))
      uid=10147 userId=0
      opPkg=com.google.android.dialer
      icon=Icon(typ=RESOURCE pkg=com.google.android.dialer id=0x7f08049b)
      flags=SHOW_LIGHTS|ONLY_ALERT_ONCE|AUTO_CANCEL|GROUP_SUMMARY
      originalFlags=SHOW_LIGHTS|ONLY_ALERT_ONCE|AUTO_CANCEL|GROUP_SUMMARY
      pri=0
      key=0|com.google.android.dialer|1|GroupSummary_MissedCall|10147
      seen=false
      groupKey=0|com.google.android.dialer|g:MissedCallGroup
      notification=
            fullscreenIntent=null
            contentIntent=PendingIntent{1a8cac1: PendingIntentRecord{4e0e125 com.google.android.dialer startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            deleteIntent=PendingIntent{4eb5266: PendingIntentRecord{19a711c com.google.android.dialer broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            number=0
            groupAlertBehavior=0
            when=1779198292329/1779198292329
            tickerText=null
            vis=0
            contentView=null
            bigContentView=null
            headsUpContentView=null
            color=0xffc0c6dc
            timeout=PT72H
            extras={
                android.title=String [length=12]
                android.reduced.images=Boolean (true)
                android.subText=null
                android.showChronometer=Boolean (false)
                android.text=String [length=14]
                android.progress=Integer (0)
                android.progressMax=Integer (0)
                android.appInfo=ApplicationInfo (ApplicationInfo{c3159a7 com.google.android.dialer})
                android.showWhen=Boolean (true)
                android.largeIcon=null
                android.infoText=null
                android.wearable.EXTENSIONS=Bundle (Bundle[{dismissalId=7d20f624-c178-43af-9390-91d15b5ea79d}])
                android.progressIndeterminate=Boolean (false)
                android.remoteInputHistory=null
                android.shortCriticalText=null
            }
      publicNotification=
            fullscreenIntent=null
            contentIntent=PendingIntent{87e3e54: PendingIntentRecord{4e0e125 com.google.android.dialer startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            deleteIntent=PendingIntent{db651fd: PendingIntentRecord{19a711c com.google.android.dialer broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            number=0
            groupAlertBehavior=0
            when=1779198292329/1779198292329
            tickerText=null
            vis=0
            contentView=null
            bigContentView=null
            headsUpContentView=null
            color=0xffc0c6dc
            timeout=PT0S
            extras={
                android.title=String [length=12]
                android.reduced.images=Boolean (true)
                android.subText=null
                android.showChronometer=Boolean (false)
                android.text=null
                android.progress=Integer (0)
                android.progressMax=Integer (0)
                android.appInfo=ApplicationInfo (ApplicationInfo{a4a2ef2 com.google.android.dialer})
                android.showWhen=Boolean (true)
                android.largeIcon=null
                android.infoText=null
                android.progressIndeterminate=Boolean (false)
                android.remoteInputHistory=null
                android.shortCriticalText=null
            }
      stats=SingleNotificationStats{posttimeElapsedMs=19510, posttimeToFirstClickMs=-1, posttimeToDismissMs=-1, airtimeCount=0, airtimeMs=0, currentAirtimeStartElapsedMs=-1, airtimeExpandedMs=0, posttimeToFirstVisibleExpansionMs=-1, currentAirtimeExpandedStartElapsedMs=-1, requestedImportance=3, naturalImportance=3, isNoisy=true}
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
      mGlobalSortKey=crtcl=0x0002:intrsv=2:grnk=0x0007:gsmry=0:nsk:rnk=0x0006
      mRankingTimeMs=1779198292329
      mCreationTimeMs=1779530867709
      mVisibleSinceMs=0
      mUpdateTimeMs=1779530867709
      mInterruptionTimeMs=1779530867709
      mSuppressedVisualEffects= 0
      mSound= null
      mVibration= Composed{segments=[Step{amplitude=0.0, duration=0}, Step{amplitude=-1.0, duration=350}, Step{amplitude=0.0, duration=250}, Step{amplitude=-1.0, duration=350}], repeat=-1}
      mAttributes= AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null
      mLight= Light{color=-1, onMs=500, offMs=2000}
      mShowBadge=true
      mColorized=false
      mAllowBubble=false
      isBubble=false
      mIsInterruptive=false
      effectiveNotificationChannel=NotificationChannel{mId='phone_missed_call', mName=Missed calls, mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      mAdjustments=[]
      mRuleChannelReason=0
      mRuleSoundReason=0
      mRuleLightReason=0
      shortcut=null found valid? false
      mUserVisOverride=-1000
      hasSummarization=false
      bundleType=0
    NotificationRecord(0x0d0fe943: pkg=com.google.android.dialer user=UserHandle{0} id=1 tag=MissedCall_content://call_log/calls/3 importance=3 key=0|com.google.android.dialer|1|MissedCall_content://call_log/calls/3|10147: Notification(channel=phone_missed_call shortcut=null contentView=null vibrate=null sound=null defaults=LIGHTS flags=SHOW_LIGHTS|ONLY_ALERT_ONCE|AUTO_CANCEL color=0xffc0c6dc category=missed_call groupKey=MissedCallGroup actions=2 vis=PRIVATE publicVersion=Notification(channel=phone_missed_call shortcut=null contentView=null vibrate=null sound=null defaults=0 flags=ONLY_ALERT_ONCE|AUTO_CANCEL color=0xffc0c6dc groupKey=MissedCallGroup vis=PRIVATE)))
      uid=10147 userId=0
      opPkg=com.google.android.dialer
      icon=Icon(typ=RESOURCE pkg=com.google.android.dialer id=0x7f08049b)
      flags=SHOW_LIGHTS|ONLY_ALERT_ONCE|AUTO_CANCEL
      originalFlags=SHOW_LIGHTS|ONLY_ALERT_ONCE|AUTO_CANCEL
      pri=0
      key=0|com.google.android.dialer|1|MissedCall_content://call_log/calls/3|10147
      seen=false
      groupKey=0|com.google.android.dialer|g:MissedCallGroup
      notification=
            fullscreenIntent=null
            contentIntent=PendingIntent{5aaebc0: PendingIntentRecord{1464d4d com.google.android.dialer startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            deleteIntent=PendingIntent{41b10f9: PendingIntentRecord{19a711c com.google.android.dialer broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            number=0
            groupAlertBehavior=0
            when=1779198292329/1779198292329
            tickerText=null
            vis=0
            contentView=null
            bigContentView=null
            headsUpContentView=null
            color=0xffc0c6dc
            timeout=PT72H
            actions={
                [0] "Call back" -> PendingIntent{848483e: PendingIntentRecord{d0a9002 com.google.android.dialer startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
                [1] "Message" -> PendingIntent{abdd69f: PendingIntentRecord{8ec1d13 com.google.android.dialer startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
              }
            extras={
                android.title=String [length=12]
                android.hiddenConversationTitle=null
                android.reduced.images=Boolean (true)
                android.subText=null
                android.template=String (android.app.Notification$MessagingStyle)
                android.showChronometer=Boolean (false)
                android.text=String [length=11]
                android.progress=Integer (0)
                androidx.core.app.extra.COMPAT_TEMPLATE=String (androidx.core.app.NotificationCompat$MessagingStyle)
                android.progressMax=Integer (0)
                android.selfDisplayName=String [length=12]
                android.conversationUnreadMessageCount=Integer (0)
                android.appInfo=ApplicationInfo (ApplicationInfo{83303ec com.google.android.dialer})
                android.messages=Bundle[] (1)
                android.showWhen=Boolean (true)
                android.largeIcon=Icon (Icon(typ=BITMAP size=126x126))
                android.messagingStyleUser=Bundle (Bundle[mParcelledData.dataSize=556])
                android.messagingUser=Person (android.app.Person@bd6c1395)
                android.infoText=null
                android.progressIndeterminate=Boolean (false)
                android.remoteInputHistory=null
                android.shortCriticalText=null
                android.isGroupConversation=Boolean (false)
            }
      publicNotification=
            fullscreenIntent=null
            contentIntent=PendingIntent{47caa4a: PendingIntentRecord{1464d4d com.google.android.dialer startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            deleteIntent=PendingIntent{4333dbb: PendingIntentRecord{19a711c com.google.android.dialer broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            number=0
            groupAlertBehavior=0
            when=1779198292329/1779198292329
            tickerText=null
            vis=0
            contentView=null
            bigContentView=null
            headsUpContentView=null
            color=0xffc0c6dc
            timeout=PT0S
            extras={
                android.title=String [length=11]
                android.reduced.images=Boolean (true)
                android.subText=null
                android.showChronometer=Boolean (false)
                android.text=null
                android.progress=Integer (0)
                android.progressMax=Integer (0)
                android.appInfo=ApplicationInfo (ApplicationInfo{4272d8 com.google.android.dialer})
                android.showWhen=Boolean (true)
                android.largeIcon=null
                android.infoText=null
                android.progressIndeterminate=Boolean (false)
                android.remoteInputHistory=null
                android.shortCriticalText=null
            }
      stats=SingleNotificationStats{posttimeElapsedMs=20031, posttimeToFirstClickMs=-1, posttimeToDismissMs=-1, airtimeCount=0, airtimeMs=0, currentAirtimeStartElapsedMs=-1, airtimeExpandedMs=0, posttimeToFirstVisibleExpansionMs=-1, currentAirtimeExpandedStartElapsedMs=-1, requestedImportance=3, naturalImportance=3, isNoisy=true}
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
      mGlobalSortKey=crtcl=0x0002:intrsv=2:grnk=0x0007:gsmry=1:nsk:rnk=0x0007
      mRankingTimeMs=1779198292329
      mCreationTimeMs=1779530868111
      mVisibleSinceMs=0
      mUpdateTimeMs=1779530868423
      mInterruptionTimeMs=1779530868423
      mSuppressedVisualEffects= 0
      mSound= null
      mVibration= Composed{segments=[Step{amplitude=0.0, duration=0}, Step{amplitude=-1.0, duration=350}, Step{amplitude=0.0, duration=250}, Step{amplitude=-1.0, duration=350}], repeat=-1}
      mAttributes= AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null
      mLight= Light{color=-1, onMs=500, offMs=2000}
      mShowBadge=true
      mColorized=false
      mAllowBubble=false
      isBubble=false
      mIsInterruptive=false
      effectiveNotificationChannel=NotificationChannel{mId='phone_missed_call', mName=Missed calls, mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      mAdjustments=[]
      mRuleChannelReason=0
      mRuleSoundReason=0
      mRuleLightReason=0
      shortcut=null found valid? false
      mUserVisOverride=-1000
      hasSummarization=false
      bundleType=0
    NotificationRecord(0x0658e631: pkg=com.google.android.dialer user=UserHandle{0} id=1 tag=MissedCall_content://call_log/calls/2 importance=3 key=0|com.google.android.dialer|1|MissedCall_content://call_log/calls/2|10147: Notification(channel=phone_missed_call shortcut=null contentView=null vibrate=null sound=null defaults=LIGHTS flags=SHOW_LIGHTS|ONLY_ALERT_ONCE|AUTO_CANCEL color=0xffc0c6dc category=missed_call groupKey=MissedCallGroup actions=2 vis=PRIVATE publicVersion=Notification(channel=phone_missed_call shortcut=null contentView=null vibrate=null sound=null defaults=0 flags=ONLY_ALERT_ONCE|AUTO_CANCEL color=0xffc0c6dc groupKey=MissedCallGroup vis=PRIVATE)))
      uid=10147 userId=0
      opPkg=com.google.android.dialer
      icon=Icon(typ=RESOURCE pkg=com.google.android.dialer id=0x7f08049b)
      flags=SHOW_LIGHTS|ONLY_ALERT_ONCE|AUTO_CANCEL
      originalFlags=SHOW_LIGHTS|ONLY_ALERT_ONCE|AUTO_CANCEL
      pri=0
      key=0|com.google.android.dialer|1|MissedCall_content://call_log/calls/2|10147
      seen=false
      groupKey=0|com.google.android.dialer|g:MissedCallGroup
      notification=
            fullscreenIntent=null
            contentIntent=PendingIntent{4e32116: PendingIntentRecord{f444830 com.google.android.dialer startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            deleteIntent=PendingIntent{5cdfa97: PendingIntentRecord{19a711c com.google.android.dialer broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            number=0
            groupAlertBehavior=0
            when=1779198289427/1779198289427
            tickerText=null
            vis=0
            contentView=null
            bigContentView=null
            headsUpContentView=null
            color=0xffc0c6dc
            timeout=PT72H
            actions={
                [0] "Call back" -> PendingIntent{2d7e484: PendingIntentRecord{3318ca9 com.google.android.dialer startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
                [1] "Message" -> PendingIntent{fa4b46d: PendingIntentRecord{8aa922e com.google.android.dialer startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
              }
            extras={
                android.title=String [length=11]
                android.hiddenConversationTitle=null
                android.reduced.images=Boolean (true)
                android.subText=null
                android.template=String (android.app.Notification$MessagingStyle)
                android.showChronometer=Boolean (false)
                android.text=String [length=11]
                android.progress=Integer (0)
                androidx.core.app.extra.COMPAT_TEMPLATE=String (androidx.core.app.NotificationCompat$MessagingStyle)
                android.progressMax=Integer (0)
                android.selfDisplayName=String [length=11]
                android.conversationUnreadMessageCount=Integer (0)
                android.appInfo=ApplicationInfo (ApplicationInfo{16838a2 com.google.android.dialer})
                android.messages=Bundle[] (1)
                android.showWhen=Boolean (true)
                android.largeIcon=Icon (Icon(typ=BITMAP size=126x126))
                android.messagingStyleUser=Bundle (Bundle[mParcelledData.dataSize=548])
                android.messagingUser=Person (android.app.Person@f5f97919)
                android.infoText=null
                android.progressIndeterminate=Boolean (false)
                android.remoteInputHistory=null
                android.shortCriticalText=null
                android.isGroupConversation=Boolean (false)
            }
      publicNotification=
            fullscreenIntent=null
            contentIntent=PendingIntent{430c4f0: PendingIntentRecord{f444830 com.google.android.dialer startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            deleteIntent=PendingIntent{6652a69: PendingIntentRecord{19a711c com.google.android.dialer broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            number=0
            groupAlertBehavior=0
            when=1779198289427/1779198289427
            tickerText=null
            vis=0
            contentView=null
            bigContentView=null
            headsUpContentView=null
            color=0xffc0c6dc
            timeout=PT0S
            extras={
                android.title=String [length=11]
                android.reduced.images=Boolean (true)
                android.subText=null
                android.showChronometer=Boolean (false)
                android.text=null
                android.progress=Integer (0)
                android.progressMax=Integer (0)
                android.appInfo=ApplicationInfo (ApplicationInfo{2353cee com.google.android.dialer})
                android.showWhen=Boolean (true)
                android.largeIcon=null
                android.infoText=null
                android.progressIndeterminate=Boolean (false)
                android.remoteInputHistory=null
                android.shortCriticalText=null
            }
      stats=SingleNotificationStats{posttimeElapsedMs=19936, posttimeToFirstClickMs=-1, posttimeToDismissMs=-1, airtimeCount=0, airtimeMs=0, currentAirtimeStartElapsedMs=-1, airtimeExpandedMs=0, posttimeToFirstVisibleExpansionMs=-1, currentAirtimeExpandedStartElapsedMs=-1, requestedImportance=3, naturalImportance=3, isNoisy=true}
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
      mGlobalSortKey=crtcl=0x0002:intrsv=2:grnk=0x0007:gsmry=1:nsk:rnk=0x0008
      mRankingTimeMs=1779198289427
      mCreationTimeMs=1779530868125
      mVisibleSinceMs=0
      mUpdateTimeMs=1779530868447
      mInterruptionTimeMs=1779530868447
      mSuppressedVisualEffects= 0
      mSound= null
      mVibration= Composed{segments=[Step{amplitude=0.0, duration=0}, Step{amplitude=-1.0, duration=350}, Step{amplitude=0.0, duration=250}, Step{amplitude=-1.0, duration=350}], repeat=-1}
      mAttributes= AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null
      mLight= Light{color=-1, onMs=500, offMs=2000}
      mShowBadge=true
      mColorized=false
      mAllowBubble=false
      isBubble=false
      mIsInterruptive=false
      effectiveNotificationChannel=NotificationChannel{mId='phone_missed_call', mName=Missed calls, mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      mAdjustments=[]
      mRuleChannelReason=0
      mRuleSoundReason=0
      mRuleLightReason=0
      shortcut=null found valid? false
      mUserVisOverride=-1000
      hasSummarization=false
      bundleType=0
    NotificationRecord(0x074fa58f: pkg=com.google.android.dialer user=UserHandle{0} id=1 tag=MissedCall_content://call_log/calls/1 importance=3 key=0|com.google.android.dialer|1|MissedCall_content://call_log/calls/1|10147: Notification(channel=phone_missed_call shortcut=null contentView=null vibrate=null sound=null defaults=LIGHTS flags=SHOW_LIGHTS|ONLY_ALERT_ONCE|AUTO_CANCEL color=0xffc0c6dc category=missed_call groupKey=MissedCallGroup actions=2 vis=PRIVATE publicVersion=Notification(channel=phone_missed_call shortcut=null contentView=null vibrate=null sound=null defaults=0 flags=ONLY_ALERT_ONCE|AUTO_CANCEL color=0xffc0c6dc groupKey=MissedCallGroup vis=PRIVATE)))
      uid=10147 userId=0
      opPkg=com.google.android.dialer
      icon=Icon(typ=RESOURCE pkg=com.google.android.dialer id=0x7f08049b)
      flags=SHOW_LIGHTS|ONLY_ALERT_ONCE|AUTO_CANCEL
      originalFlags=SHOW_LIGHTS|ONLY_ALERT_ONCE|AUTO_CANCEL
      pri=0
      key=0|com.google.android.dialer|1|MissedCall_content://call_log/calls/1|10147
      seen=false
      groupKey=0|com.google.android.dialer|g:MissedCallGroup
      notification=
            fullscreenIntent=null
            contentIntent=PendingIntent{cf5401c: PendingIntentRecord{cb89d9e com.google.android.dialer startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            deleteIntent=PendingIntent{f9d0425: PendingIntentRecord{19a711c com.google.android.dialer broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            number=0
            groupAlertBehavior=0
            when=1779198285415/1779198285415
            tickerText=null
            vis=0
            contentView=null
            bigContentView=null
            headsUpContentView=null
            color=0xffc0c6dc
            timeout=PT72H
            actions={
                [0] "Call back" -> PendingIntent{6bc39fa: PendingIntentRecord{d5707f com.google.android.dialer startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
                [1] "Message" -> PendingIntent{92b0bab: PendingIntentRecord{461c84c com.google.android.dialer startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
              }
            extras={
                android.title=String [length=14]
                android.hiddenConversationTitle=null
                android.reduced.images=Boolean (true)
                android.subText=null
                android.template=String (android.app.Notification$MessagingStyle)
                android.showChronometer=Boolean (false)
                android.text=String [length=11]
                android.progress=Integer (0)
                androidx.core.app.extra.COMPAT_TEMPLATE=String (androidx.core.app.NotificationCompat$MessagingStyle)
                android.progressMax=Integer (0)
                android.selfDisplayName=String [length=14]
                android.conversationUnreadMessageCount=Integer (0)
                android.appInfo=ApplicationInfo (ApplicationInfo{6244208 com.google.android.dialer})
                android.messages=Bundle[] (1)
                android.showWhen=Boolean (true)
                android.largeIcon=Icon (Icon(typ=BITMAP size=126x126))
                android.messagingStyleUser=Bundle (Bundle[mParcelledData.dataSize=564])
                android.messagingUser=Person (android.app.Person@24bf98cf)
                android.infoText=null
                android.progressIndeterminate=Boolean (false)
                android.remoteInputHistory=null
                android.shortCriticalText=null
                android.isGroupConversation=Boolean (false)
            }
      publicNotification=
            fullscreenIntent=null
            contentIntent=PendingIntent{183fbc6: PendingIntentRecord{cb89d9e com.google.android.dialer startActivity (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            deleteIntent=PendingIntent{d0cb787: PendingIntentRecord{19a711c com.google.android.dialer broadcastIntent (allowlist: 3cd0aae:+30s0ms/0/NOTIFICATION_SERVICE/NotificationManagerService)}}
            number=0
            groupAlertBehavior=0
            when=1779198285415/1779198285415
            tickerText=null
            vis=0
            contentView=null
            bigContentView=null
            headsUpContentView=null
            color=0xffc0c6dc
            timeout=PT0S
            extras={
                android.title=String [length=11]
                android.reduced.images=Boolean (true)
                android.subText=null
                android.showChronometer=Boolean (false)
                android.text=null
                android.progress=Integer (0)
                android.progressMax=Integer (0)
                android.appInfo=ApplicationInfo (ApplicationInfo{7bf76b4 com.google.android.dialer})
                android.showWhen=Boolean (true)
                android.largeIcon=null
                android.infoText=null
                android.progressIndeterminate=Boolean (false)
                android.remoteInputHistory=null
                android.shortCriticalText=null
            }
      stats=SingleNotificationStats{posttimeElapsedMs=20021, posttimeToFirstClickMs=-1, posttimeToDismissMs=-1, airtimeCount=0, airtimeMs=0, currentAirtimeStartElapsedMs=-1, airtimeExpandedMs=0, posttimeToFirstVisibleExpansionMs=-1, currentAirtimeExpandedStartElapsedMs=-1, requestedImportance=3, naturalImportance=3, isNoisy=true}
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
      mGlobalSortKey=crtcl=0x0002:intrsv=2:grnk=0x0007:gsmry=1:nsk:rnk=0x0009
      mRankingTimeMs=1779198285415
      mCreationTimeMs=1779530868136
      mVisibleSinceMs=0
      mUpdateTimeMs=1779530868776
      mInterruptionTimeMs=1779530868776
      mSuppressedVisualEffects= 0
      mSound= null
      mVibration= Composed{segments=[Step{amplitude=0.0, duration=0}, Step{amplitude=-1.0, duration=350}, Step{amplitude=0.0, duration=250}, Step{amplitude=-1.0, duration=350}], repeat=-1}
      mAttributes= AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null
      mLight= Light{color=-1, onMs=500, offMs=2000}
      mShowBadge=true
      mColorized=false
      mAllowBubble=false
      isBubble=false
      mIsInterruptive=false
      effectiveNotificationChannel=NotificationChannel{mId='phone_missed_call', mName=Missed calls, mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      mAdjustments=[]
      mRuleChannelReason=0
      mRuleSoundReason=0
      mRuleLightReason=0
      shortcut=null found valid? false
      mUserVisOverride=-1000
      hasSummarization=false
      bundleType=0
  
  mMaxPackageEnqueueRate=5.0
  hideSilentStatusBar=true

  Notification attention state:
      mSoundNotificationKey=null
      mVibrateNotificationKey=null
      mDisableNotificationEffects=false
      mCallState=CALL_STATE_IDLE
      mSystemReady=true
      mNotificationPulseEnabled=true
  mArchive=Archive (0 notifications)

  Snoozed notifications:

 Pending snoozed notifications

  Ranking Config:
    mSignalExtractors.length = 10
      NotificationChannelExtractor
      NotificationAdjustmentExtractor
      BubbleExtractor
      ValidateNotificationPeople
      PriorityExtractor
      ZenModeExtractor
      ImportanceExtractor
      VisibilityExtractor
      BadgeExtractor
      CriticalNotificationExtractor

 Notification Preferences:
    per-package config version: 4
PackagePreferences:
      AppSettings: com.android.ons (1001) fixedImportance=true
      AppSettings: com.google.android.safetycenter.resources (10207)
      AppSettings: com.google.android.providers.media.module (10220) importance=DEFAULT userSet=false fixedImportance=true
        NotificationChannel{mId='PhotoPickerSyncChannel', mName=Med..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='native_transcode_alert_channel', mName=Tra..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='native_transcode_progress_channel', mName=Tra..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.android.internal.display.cutout.emulation.corner (10006)
      AppSettings: com.android.emergency (10200)
      AppSettings: com.android.managedprovisioning.auto_generated_rro_product__ (10047)
      AppSettings: com.android.settings.auto_generated_rro_product__ (10054)
      AppSettings: com.android.localtransport (1000) fixedImportance=true
      AppSettings: com.android.pacprocessor (10136)
      AppSettings: com.android.wallpaper (10196) importance=NONE userSet=false
      AppSettings: com.android.systemui.emulation.pixel_3 (10062)
      AppSettings: com.google.android.printservice.recommendation (10134)
      AppSettings: com.android.credentialmanager (10107)
      AppSettings: com.google.android.wifi.dialog (10213)
      AppSettings: com.google.mainline.adservices (10175)
      AppSettings: com.android.htmlviewer (10135)
      AppSettings: com.google.android.apps.docs (10163) importance=NONE userSet=false
        NotificationChannel{mId='SCAN_DOWNLOADS', mName=Sca..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='CONTENT_SYNC_OTHER', mName=Syn..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='ACCOUNT_UPDATES', mName=Acc..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.android.systemui.auto_generated_rro_product__ (10085)
      AppSettings: com.google.android.configupdater (10144)
      AppSettings: com.android.systemui.emulation.pixel_8 (10076)
      AppSettings: com.google.android.nfc (1027) importance=DEFAULT userSet=false fixedImportance=true
      AppSettings: com.android.internal.emulation.pixel_7 (10028)
      AppSettings: com.android.internal.emulation.pixel_3a_xl (10020)
      AppSettings: com.android.storagemanager (10194) importance=DEFAULT userSet=false fixedImportance=true
      AppSettings: com.google.android.apps.wallpaper (10198)
      AppSettings: com.google.android.appsearch.apk (10218)
      AppSettings: com.android.systemui.plugin.globalactions.wallet (10193)
      AppSettings: com.android.carrierconfig.auto_generated_rro_product__ (10002)
      AppSettings: com.google.android.uwb.resources.goldfish.overlay (10092)
      AppSettings: com.android.dynsystem (1000) importance=DEFAULT userSet=false fixedImportance=true
        NotificationChannel{mId='com.android.dynsystem', mName=Dyn..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.google.android.dialer (10147) importance=DEFAULT userSet=false defaultAppLocked=true
        Delegate: com.google.android.gms (10219) enabled=true
        NotificationChannel{mId='phone_ongoing_call', mName=Ong..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=true, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='phone_incoming_call', mName=Inc..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=true, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='phone_emergency_call', mName=Eme..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='phone_low_priority', mName=Bac..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='phone_voicemail', mName=Voi..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='phone_missed_call', mName=Mis..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='phone_feedback', mName=Fee..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='phone_default', mName=Def..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.android.internal.display.cutout.emulation.double (10007)
      AppSettings: com.android.internal.emulation.pixel_8_pro (10032)
      AppSettings: com.android.providers.partnerbookmarks (10137)
      AppSettings: com.google.android.apps.customization.pixel (10172)
      AppSettings: com.android.systemui.emulation.pixel_6a (10072)
      AppSettings: com.android.internal.display.cutout.emulation.emu01 (10008)
      AppSettings: com.android.internal.systemui.navbar.transparent (10091)
      AppSettings: com.google.android.captiveportallogin (10127) importance=DEFAULT userSet=false
      AppSettings: com.google.android.overlay.largescreensettingsprovider (10046)
      AppSettings: com.android.systemui.accessibility.accessibilitymenu.auto_generated_rro_product__ (10000)
      AppSettings: com.android.certinstaller (10129)
      AppSettings: com.google.android.partnersetup (10152)
      AppSettings: com.raju.shingadiya.debug (10228) importance=DEFAULT userSet=false
      AppSettings: com.android.providers.contacts (10102)
      AppSettings: com.android.proxyhandler (10119)
      AppSettings: com.android.providers.contacts.auto_generated_rro_product__ (10004)
      AppSettings: com.android.wallpaper.livepicker (10115)
      AppSettings: com.google.android.apps.messaging (10154) importance=DEFAULT userSet=false
        Delegate: com.google.android.gms (10219) enabled=true
        NotificationChannel{mId='bugle_misc_channel', mName=Oth..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='bugle_default_channel', mName=Inc..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='_bugle_default_settings_group', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='bugle_broadcast_receiver_channel', mName=Bac..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=true, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannelGroup{mId='bugle_web_group', mName=Messages for web, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
        NotificationChannelGroup{mId='_bugle_default_settings_group', mName=Default settings, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
        NotificationChannelGroup{mId='bugle_conversations_group', mName=Conversations, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
      AppSettings: com.google.android.googlequicksearchbox (10158) importance=DEFAULT userSet=false
      AppSettings: com.android.theme.font.notoserifsource (10041)
      AppSettings: com.android.internal.emulation.pixel_4a (10023)
      AppSettings: com.google.android.apps.youtube.music (10183) importance=NONE userSet=false
      AppSettings: com.android.bluetoothmidiservice (10124)
      AppSettings: com.android.internal.emulation.pixel_10_pro_xl (10015)
      AppSettings: com.google.android.gsf (10191)
      AppSettings: com.android.providers.telephony.auto_generated_characteristics_rro (10088)
      AppSettings: com.android.providers.telephony (1001) fixedImportance=true
      AppSettings: com.android.systemui.emulation.pixel_7_pro (10074)
      AppSettings: com.google.android.modulemetadata (10170)
      AppSettings: com.android.internal.emulation.pixel_9a (10038)
      AppSettings: android.auto_generated_rro_vendor__ (10099)
      AppSettings: com.google.android.packageinstaller (10113) importance=DEFAULT userSet=false fixedImportance=true
      AppSettings: com.android.internal.emulation.pixel_fold (10039)
      AppSettings: com.android.systemui.emulation.pixel_4 (10066)
      AppSettings: com.android.wallpapercropper (10197)
      AppSettings: com.android.systemui.emulation.pixel_4_xl (10067)
      AppSettings: com.android.internal.emulation.pixel_3 (10017)
      AppSettings: com.android.internal.systemui.navbar.gestural (10050)
      AppSettings: com.android.cellbroadcastreceiver (10105)
      AppSettings: com.android.systemui.emulation.pixel_9 (10079)
      AppSettings: com.android.calllogbackup (10102)
      AppSettings: android (1000) importance=DEFAULT userSet=false fixedImportance=true
        NotificationChannel{mId='ABUSIVE_BACKGROUND_APPS', mName=Bac..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='NETWORK_ALERTS', mName=Net..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='DEVELOPER', mName=Dev..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='APM_ALERTS', mName=APM..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='SYSTEM_CHANGES_ALERTS', mName=Sys..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='safety_center_critical_warning', mName=Cri..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='safety_center_channels', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='PARENTAL_CONTROLS', mName=Par..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='NETWORK_STATUS', mName=Net..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='PHYSICAL_KEYBOARD', mName=Phy..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='HEAVY_WEIGHT_APP', mName=App..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION_EVENT content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='ACCOUNT', mName=Acc..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='DEVICE_ADMIN_ALERTS', mName=Ale..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='ACCESSIBILITY_SECURITY_POLICY', mName=Acc..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='FOREGROUND_SERVICE', mName=App..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='USB', mName=USB..., mDescription=, mImportance=1, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=1, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='VPN', mName=VPN..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='TIME', mName=Tim..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='safety_center_recommendation', mName=War..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='safety_center_channels', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='UPDATES', mName=Upd..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='safety_center_information', mName=Rec..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='safety_center_channels', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='RETAIL_MODE', mName=Ret..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='CAR_MODE', mName=Car..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='ACCESSIBILITY_HEARING_DEVICE', mName=Hea..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='ACCESSIBILITY_MAGNIFICATION', mName=Mag..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='bg_user_sound_channel', mName=Bac..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='SECURITY', mName=Sec..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='NETWORK_AVAILABLE', mName=Net..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='DEVELOPER_IMPORTANT', mName=Imp..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='ALERTS', mName=Ale..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannelGroup{mId='safety_center_channels', mName=Security & privacy, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
      AppSettings: com.google.android.telecomui (10223)
      AppSettings: com.android.systemui.emulation.pixel_9_pro_xl (10082)
      AppSettings: com.google.android.overlay.googleconfig (10042)
      AppSettings: com.android.internal.emulation.pixel_10_pro_fold (10014)
      AppSettings: com.android.internal.emulation.pixel_8 (10031)
      AppSettings: com.google.android.inputmethod.latin (10167)
        NotificationChannel{mId='com.google.android.apps.inputmethod.libs.dataservice.superpacks', mName=Dow..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='com.google.android.inputmethod.latin', mName=Dow..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.android.internal.ranchu.commonoverlay (10053)
      AppSettings: com.google.android.photopicker (10221)
      AppSettings: com.android.internal.systemui.navbar.threebutton (10049)
      AppSettings: com.android.systemui.emulation.pixel_3_xl (10063)
      AppSettings: com.google.android.marvin.talkback (10184) importance=NONE userSet=false
      AppSettings: com.google.android.networkstack (1073) fixedImportance=true
        NotificationChannel{mId='connected_note', mName=Net..., mDescription=hasDescription , mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='connected_note_loud', mName=Cap..., mDescription=hasDescription , mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.android.intentresolver (10114)
      AppSettings: com.android.providers.userdictionary (10102)
      AppSettings: com.google.android.settings.intelligence (10157) importance=DEFAULT userSet=false
      AppSettings: com.android.phone.auto_generated_rro_vendor__ (10098)
      AppSettings: com.android.systemui.emulation.pixel_4a (10068)
      AppSettings: com.google.android.sdksandbox (10206)
      AppSettings: com.android.mms.service (1001) fixedImportance=true
      AppSettings: com.google.android.cellbroadcastservice (1073) fixedImportance=true
      AppSettings: com.google.android.cellbroadcastreceiver (10201) importance=DEFAULT userSet=false fixedImportance=true
      AppSettings: android.auto_generated_characteristics_rro (10093)
      AppSettings: com.android.systemui.emulation.pixel_9a (10083)
      AppSettings: com.android.devicediagnostics.auto_generated_rro_product__ (10005)
      AppSettings: com.android.devicelockcontroller (10202) importance=DEFAULT userSet=false fixedImportance=true
      AppSettings: com.android.systemui.emulation.pixel_9_pro (10080)
      AppSettings: com.google.android.adservices.api (10205) importance=DEFAULT userSet=false
        NotificationChannel{mId='PRIVACY_SANDBOX_CHANNEL', mName=Pri..., mDescription=hasDescription , mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.android.server.telecom (1000) importance=DEFAULT userSet=false fixedImportance=true
        NotificationChannel{mId='TelecomMissedCalls', mName=Mis..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='TelecomDisconnectedCalls', mName=Dis..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='TelecomCallStreaming', mName=Cal..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='TelecomInCallServiceCrash', mName=Cra..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='LocalVoicemail', mName=Voi..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='TelecomCallBlocking', mName=Cal..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='TelecomIncomingCalls', mName=Inc..., mDescription=, mImportance=5, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=5, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='TelecomBackgroundAudioProcessing', mName=Bac..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.android.shell (2000) importance=DEFAULT userSet=false fixedImportance=true
      AppSettings: com.google.android.telecomresources (10224)
      AppSettings: com.android.wallpaperbackup (10141)
      AppSettings: com.google.android.hotspot2.osulogin (10212)
      AppSettings: com.google.android.rkpdapp (10226)
      AppSettings: com.android.emulator.multidisplay (1000) fixedImportance=true
      AppSettings: com.android.internal.emulation.pixel_7a (10030)
      AppSettings: com.google.android.accessibility.switchaccess (10178)
      AppSettings: com.google.android.glasses.core (10143) importance=NONE userSet=false
      AppSettings: com.android.providers.calendar (10104)
      AppSettings: com.google.android.webapp.service (10203) importance=DEFAULT userSet=false
      AppSettings: com.android.emulator.radio.config (10188)
      AppSettings: com.google.android.ext.services (10225)
      AppSettings: com.android.companiondevicemanager (10130)
      AppSettings: com.android.simappdialog (10139)
      AppSettings: com.google.android.healthconnect.controller (10209)
      AppSettings: com.google.android.apps.accessibility.voiceaccess (10180) importance=NONE userSet=false
      AppSettings: com.google.android.apps.pulse (10155)
        Delegate: com.google.android.gms (10219) enabled=true
      AppSettings: com.android.vpndialogs (10122)
      AppSettings: com.google.android.apps.restore (10148) importance=DEFAULT userSet=false
      AppSettings: com.google.mainline.telemetry (10176)
      AppSettings: com.google.android.marvin.talkbackoverlay (10096) importance=NONE userSet=false
      AppSettings: com.android.systemui.emulation.pixel_5 (10069)
      AppSettings: com.google.android.markup (10169)
      AppSettings: com.google.android.gms (10219) importance=DEFAULT userSet=false
        NotificationChannel{mId='gms.pay_default', mName=All..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='gms.gpay_wallet', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='nearby_sharing_alert', mName=Ale..., mDescription=hasDescription , mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='NEARBY_SHARING_CHANNEL_GROUP_ID', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='eew_notification_v2', mName=Ear..., mDescription=hasDescription , mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=android.resource://com.google.android.gms/raw/be_aware_alert_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='Personal_Safety_Id', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION_EVENT content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='eew_update', mName=Ear..., mDescription=hasDescription , mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='Personal_Safety_Id', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='DEVICES_WITH_YOUR_ACCOUNT', mName=Dev..., mDescription=, mImportance=1, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='DEVICES_CHANNEL_GROUP_ID', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=1, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='com.google.android.gms.backup.g1.storagealerts.notification.channel.id', mName=Sto..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='com.google.android.gms.backup.g1.notification.group.id', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='mt-notification-instant-hotspot-hun-channel-id', mName=Ins..., mDescription=, mImportance=4, mBypassDnd=true, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='com.google.android.gms.multidevice.NOTIFICATION_GROUP_ID', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='eew_alert_v2', mName=Ear..., mDescription=hasDescription , mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=android.resource://com.google.android.gms/raw/be_aware_alert_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='Personal_Safety_Id', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION_EVENT content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='com.google.android.gms.mobiledataplan.NOTIFICATION.LOW_BALANCE', mName=Dat..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='com.google.android.gms.mobiledataplan.NOTIFICATION', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='com.google.android.gms.mobiledataplan.NOTIFICATION.ACCOUNT_ALERT', mName=Acc..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='com.google.android.gms.mobiledataplan.NOTIFICATION', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='nearby_sharing_app', mName=App..., mDescription=hasDescription , mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='NEARBY_SHARING_CHANNEL_GROUP_ID', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='DEVICES_WITHIN_REACH_REBRANDED', mName=Dev..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='DEVICES_CHANNEL_GROUP_ID', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='DEVICES_REBRANDED', mName=Dev..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='DEVICES_CHANNEL_GROUP_ID', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='mt-notification-channel-id', mName=Pro..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='mt-notification-channel-group-id', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='crisis_alerts_1_1', mName=Cri..., mDescription=hasDescription , mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='Personal_Safety_Id', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION_EVENT content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='crisis_alerts_2_1', mName=Cri..., mDescription=hasDescription , mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='Personal_Safety_Id', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION_EVENT content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='crisis_alerts_3_1', mName=Cri..., mDescription=hasDescription , mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='Personal_Safety_Id', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='mt-notification-connected-channel-id', mName=Sha..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='mt-notification-channel-group-id', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='nearby_sharing_file', mName=Fil..., mDescription=hasDescription , mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='NEARBY_SHARING_CHANNEL_GROUP_ID', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='com.google.android.gms.backup.g1.featureupdates.notification.channel.id', mName=Fea..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='com.google.android.gms.backup.g1.notification.group.id', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='com.google.android.gms.backup.g1.statusalerts.notification.channel.id', mName=Sta..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='com.google.android.gms.backup.g1.notification.group.id', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='com.google.android.gms.mobiledataplan.NOTIFICATION.UPSELL_OFFER', mName=Dat..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='com.google.android.gms.mobiledataplan.NOTIFICATION', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='nearby_sharing_privacy', mName=Pri..., mDescription=hasDescription , mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='NEARBY_SHARING_CHANNEL_GROUP_ID', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='dck.alerts', mName=Imp..., mDescription=hasDescription , mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='dck', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannelGroup{mId='mt-notification-channel-group-id', mName=Instant Tethering, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
        NotificationChannelGroup{mId='DEVICES_CHANNEL_GROUP_ID', mName=Devices, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
        NotificationChannelGroup{mId='com.google.android.gms.backup.g1.notification.group.id', mName=Back up and restore, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
        NotificationChannelGroup{mId='com.google.android.gms.mobiledataplan.NOTIFICATION', mName=Device Plans, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
        NotificationChannelGroup{mId='NEARBY_CHANNEL_GROUP_ID', mName=Nearby, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
        NotificationChannelGroup{mId='com.google.android.gms.multidevice.NOTIFICATION_GROUP_ID', mName=Cross-device services, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
        NotificationChannelGroup{mId='gms.gpay_wallet', mName=Google Pay & Wallet, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
        NotificationChannelGroup{mId='NEARBY_SHARING_CHANNEL_GROUP_ID', mName=Quick Share, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
        NotificationChannelGroup{mId='dck', mName=Digital Car Key, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
        NotificationChannelGroup{mId='Personal_Safety_Id', mName=Personal Safety, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
      AppSettings: com.google.android.health.connect.backuprestore (10210)
      AppSettings: com.android.systemui.accessibility.accessibilitymenu (10199)
      AppSettings: com.android.systemui.emulation.pixel_10_pro (10059)
      AppSettings: com.android.internal.emulation.pixel_4 (10021)
      AppSettings: com.android.DeviceAsWebcam (1000) importance=DEFAULT userSet=false fixedImportance=true
      AppSettings: com.android.eyedropper (10112)
      AppSettings: com.android.internal.display.cutout.emulation.tall (10010)
      AppSettings: com.google.android.bluetooth (1002) importance=DEFAULT userSet=false fixedImportance=true
        NotificationChannel{mId='opp_notification_channel', mName=Blu..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.android.camera2 (10161) importance=NONE userSet=false
      AppSettings: com.android.internal.emulation.pixel_9 (10034)
      AppSettings: com.android.providers.telephony.auto_generated_rro_product__ (10089)
      AppSettings: com.google.android.networkstack.tethering (1073) importance=DEFAULT userSet=false fixedImportance=true
        NotificationChannel{mId='TETHERING_STATUS', mName=Hot..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.google.android.uwb.resources (10204)
      AppSettings: com.google.android.webview (10181) importance=NONE userSet=false
      AppSettings: com.android.musicfx (10117)
      AppSettings: com.google.android.odad (10151) importance=DEFAULT userSet=false
      AppSettings: com.google.android.feedback (10189)
      AppSettings: com.google.android.contacts (10165) importance=DEFAULT userSet=true
        NotificationChannel{mId='PEOPLE_PROMPTS_CHANNEL', mName=Rem..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='SUGGESTIONS_CHANNEL', mName=Sug..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='DEFAULT_CHANNEL', mName=Pro..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=true, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.google.android.apps.nexuslauncher (10192) importance=DEFAULT userSet=false
      AppSettings: com.android.internal.emulation.pixel_7_pro (10029)
      AppSettings: com.android.simappdialog.auto_generated_rro_product__ (10056)
      AppSettings: com.android.systemui.emulation.pixel_7a (10075)
      AppSettings: com.android.role.notes.enabled (10051)
      AppSettings: com.android.externalstorage (10111)
      AppSettings: com.google.android.documentsui (10109) importance=DEFAULT userSet=false
      AppSettings: com.android.imsserviceentitlement (10149)
      AppSettings: com.android.cts.priv.ctsshim (10214) importance=NONE userSet=false
      AppSettings: com.android.providers.blockednumber (10102)
      AppSettings: com.google.android.glasses.companion (10164) importance=DEFAULT userSet=false
      AppSettings: com.android.soundpicker (10110) fixedImportance=true
      AppSettings: com.google.android.permissioncontroller (10208) importance=DEFAULT userSet=false fixedImportance=true
      AppSettings: com.android.carrierconfig.auto_generated_rro_vendor__ (10095)
      AppSettings: com.google.android.ondevicepersonalization.services (10216)
      AppSettings: com.google.android.apps.wellbeing (10159) importance=DEFAULT userSet=false
        NotificationChannel{mId='discovery_notifications', mName=Sug..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='app_limit_updates', mName=Upd..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=true, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.android.systemui.emulation.pixel_6_pro (10071)
      AppSettings: com.google.android.tag (10121)
      AppSettings: com.android.keychain (1000) fixedImportance=true
      AppSettings: com.android.inputdevices (1000) fixedImportance=true
      AppSettings: com.google.android.as (10146) importance=DEFAULT userSet=false
        NotificationChannel{mId='Captions', mName=Liv..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=true, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.android.se (1068)
      AppSettings: com.android.bips.auto_generated_rro_product__ (10001)
      AppSettings: com.android.storagemanager.auto_generated_rro_product__ (10057)
      AppSettings: com.android.contactspicker (10106)
      AppSettings: com.android.systemui.emulation.pixel_6 (10070)
      AppSettings: com.android.internal.emulation.pixel_5 (10024)
      AppSettings: com.google.android.overlay.permissioncontroller (10043)
      AppSettings: com.google.android.apps.safetyhub (10156) importance=DEFAULT userSet=false defaultAppLocked=true
      AppSettings: com.android.mtp (10110) importance=DEFAULT userSet=false fixedImportance=true
        NotificationChannel{mId='device_notification_channel', mName=Unc..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.android.internal.display.cutout.emulation.hole (10009)
      AppSettings: com.android.internal.emulation.pixel_9_pro (10035)
      AppSettings: com.android.systemui.emulation.pixel_fold (10084)
      AppSettings: com.android.phone.auto_generated_characteristics_rro (10086)
      AppSettings: com.android.traceur.auto_generated_rro_product__ (10090)
      AppSettings: com.google.android.overlay.pixelconfigcommon (10052)
      AppSettings: com.android.chrome (10162) importance=NONE userSet=false
        NotificationChannel{mId='incognito', mName=Inc..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='general', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='tips', mName=Chr..., mDescription=, mImportance=0, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='general', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=0, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='media', mName=Pla..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='general', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='browser', mName=Bro..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='general', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='downloads', mName=Act..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='general', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannelGroup{mId='general', mName=General, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
      AppSettings: com.android.providers.contactkeys (10102)
      AppSettings: com.android.bips (10103) importance=DEFAULT userSet=false fixedImportance=true
      AppSettings: com.android.printspooler (10138) importance=DEFAULT userSet=false fixedImportance=true
        NotificationChannel{mId='PRINT_PROGRESS', mName=Run..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='PRINT_FAILURES', mName=Fai..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.android.internal.emulation.pixel_9_pro_xl (10037)
      AppSettings: com.google.android.wifi.resources (10211)
      AppSettings: com.android.sharedstoragebackup (10120)
      AppSettings: com.android.systemui.emulation.pixel_3a_xl (10065)
      AppSettings: com.google.android.connectivity.resources.goldfish.overlay (10100)
      AppSettings: com.android.carrierdefaultapp (10128) importance=DEFAULT userSet=false fixedImportance=true
      AppSettings: com.android.traceur (10140) importance=DEFAULT userSet=false fixedImportance=true
        NotificationChannel{mId='trace-is-being-recorded', mName=Tra..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='system-tracing', mName=Sav..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.google.android.federatedcompute (10217)
      AppSettings: com.google.android.deskclock (10173) importance=DEFAULT userSet=false
        NotificationChannel{mId='Timers', mName=Tim..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=true, mDeletedTimeMs=1778993995320, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_UNKNOWN content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='Missed Alarms', mName=Mis..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=true, mDeletedTimeMs=1778993995335, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_UNKNOWN content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='Snoozed Alarms v2', mName=Sno..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='Workflows', mName=Goo..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=true, mDeletedTimeMs=1778993995339, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_UNKNOWN content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='Upcoming Alarms v2', mName=Upc..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='Stopwatch', mName=Sto..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=true, mDeletedTimeMs=1778993995339, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_UNKNOWN content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='Stopwatch v2', mName=Sto..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='Snoozed Alarms', mName=Sno..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=true, mDeletedTimeMs=1778993995341, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_UNKNOWN content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='Bedtime', mName=Bed..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='Missed Alarms v2', mName=Mis..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='Upcoming Alarms', mName=Upc..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=true, mDeletedTimeMs=1778993995343, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_UNKNOWN content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='Workflows v2', mName=Rou..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='Timers v2', mName=Tim..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='Firing', mName=Fir..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_ALARM content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.android.providers.settings (1000) importance=DEFAULT userSet=false fixedImportance=true
      AppSettings: com.android.systemui.auto_generated_rro_vendor__ (10097)
      AppSettings: com.google.android.calendar (10160) importance=DEFAULT userSet=false
      AppSettings: com.android.internal.emulation.pixel_3a (10019)
      AppSettings: com.google.android.networkstack.tethering.emulator (10040)
      AppSettings: com.android.internal.emulation.pixel_4_xl (10022)
      AppSettings: com.android.systemui.emulation.pixel_8_pro (10077)
      AppSettings: com.google.android.as.oss (10145)
      AppSettings: com.android.internal.emulation.pixel_8a (10033)
      AppSettings: com.android.stk (1001) importance=DEFAULT userSet=false fixedImportance=true
      AppSettings: com.google.android.apps.maps (10168) importance=NONE userSet=false
      AppSettings: com.android.internal.emulation.pixel_10_pro (10013)
      AppSettings: com.android.location.fused (1000) fixedImportance=true
      AppSettings: com.google.android.tts (10166) importance=NONE userSet=false
      AppSettings: com.google.android.telephony.satellite (10186)
      AppSettings: com.google.android.projection.gearhead (10142) importance=DEFAULT userSet=false
      AppSettings: com.android.internal.emulation.pixel_3_xl (10018)
      AppSettings: com.android.providers.downloads.ui (10110) fixedImportance=true
      AppSettings: com.android.systemui (10195) importance=DEFAULT userSet=false fixedImportance=true
        NotificationChannel{mId='ALR', mName=Ale..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='BAT', mName=Bat..., mDescription=, mImportance=5, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=file:///product/media/audio/ui/LowBattery.ogg, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION_EVENT content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=5, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='DSK', mName=Sto..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='HNT', mName=Hin..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='INS', mName=Ins..., mDescription=, mImportance=1, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=1, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='STP', mName=Set..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='SCN_HEADSUP', mName=Scr..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='screen_record', mName=Scr..., mDescription=hasDescription , mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='ContextualEduNotificationChannel', mName=And..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='TutorialSchedulerNotificationChannel', mName=And..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.google.android.gm (10174) importance=NONE userSet=false
      AppSettings: com.android.nfc (1027) fixedImportance=true
      AppSettings: com.google.android.youtube (10182) importance=NONE userSet=false
        Delegate: com.google.android.gms (10219) enabled=true
        NotificationChannel{mId='generic_notifications', mName=Gen..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='1', mName=Sub..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='2', mName=Liv..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='3', mName=Com..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='5', mName=Rec..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='6', mName=Pro..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='7', mName=Loc..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='10', mName=Sen..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='90', mName=Fee..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='OfflineNotifications', mName=Dow..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.google.android.googlesdksetup (10190)
      AppSettings: com.android.devicediagnostics (10108)
      AppSettings: com.android.internal.emulation.pixel_10 (10012)
      AppSettings: com.google.android.avatarpicker (10185)
      AppSettings: com.android.systemui.emulation.pixel_7 (10073)
      AppSettings: com.android.internal.systemui.navbar.twobutton (10048)
      AppSettings: com.android.internal.emulation.pixel_9_pro_fold (10036)
      AppSettings: com.android.carrierconfig (10187)
      AppSettings: com.android.appsearch.aiseal.config (10132)
      AppSettings: android.auto_generated_rro_product__ (10094)
      AppSettings: com.google.android.apps.photos (10171) importance=NONE userSet=false
        Delegate: com.google.android.gms (10219) enabled=true
        NotificationChannel{mId='backup_2_suggestions', mName=Sug..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='backup', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='backup_1_progress', mName=Pro..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='backup', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='feature_drops', mName=Fea..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='movies', mName=Mov..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='free_up_space', mName=Fre..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='other', mName=Oth..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='promotions', mName=Pro..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='restore', mName=Res..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='video_boost', mName=Vid..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='download', mName=Dow..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=false, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='assistant', mName=Cre..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='memories', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='backup_3_alerts', mName=Ale..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='backup', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='new_memories', mName=New..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='memories', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='sharing', mName=Sha..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=true, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=true, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannelGroup{mId='backup', mName=Backup, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
        NotificationChannelGroup{mId='memories', mName=Memories, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
      AppSettings: com.android.settings (1000) importance=DEFAULT userSet=false fixedImportance=true
      AppSettings: com.android.companiondevicemanager.auto_generated_characteristics_rro (10003)
      AppSettings: com.android.internal.emulation.pixel_2_xl (10016)
      AppSettings: com.android.internal.emulation.pixel_6 (10025)
      AppSettings: com.android.systemui.emulation.pixel_9_pro_fold (10081)
      AppSettings: com.android.backupconfirm (10101)
      AppSettings: com.google.android.connectivity.resources (10222)
      AppSettings: com.android.egg (10131) importance=DEFAULT userSet=false fixedImportance=true
      AppSettings: com.android.systemui.emulation.pixel_10_pro_fold (10060)
      AppSettings: com.android.systemui.emulation.pixel_3a (10064)
      AppSettings: com.android.cts.ctsshim (10215) importance=NONE userSet=false
      AppSettings: com.android.phone.auto_generated_rro_product__ (10087)
      AppSettings: com.google.android.ext.shared (10133)
      AppSettings: com.android.dreams.basic (10123)
      AppSettings: com.google.android.overlay.googlewebview (10044)
      AppSettings: com.android.phone (1001) importance=DEFAULT userSet=false fixedImportance=true
        NotificationChannel{mId='satellite', mName=sat..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='voiceMail', mName=Voi..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='mobileDataAlertNew', mName=Mob..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='sim', mName=SIM..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=null, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='sms', mName=SMS..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='wfc', mName=Wi-..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='alert', mName=Ale..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='simHighPriority', mName=Hig..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='callForwardNew', mName=Cal..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.android.internal.display.cutout.emulation.waterfall (10011)
      AppSettings: com.google.android.overlay.largescreenconfig (10045)
      AppSettings: com.android.bookmarkprovider (10125)
      AppSettings: com.android.systemui.emulation.pixel_8a (10078)
      AppSettings: com.android.managedprovisioning (10116) importance=DEFAULT userSet=false fixedImportance=true
      AppSettings: com.android.vending (10153) importance=DEFAULT userSet=false
        Delegate: com.google.android.gms (10219) enabled=true
        NotificationChannel{mId='your-community', mName=You..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='alerts', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='security-and-errors', mName=Sec..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='essentials', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='account', mName=Acc..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='alerts', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='leagues-updates', mName=Lea..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='alerts', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='community-posts', mName=Com..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=true, mDeletedTimeMs=1778088999650, mGroup='alerts', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='follow-genres', mName=Gen..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='alerts', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='updates-completed', mName=Upd..., mDescription=, mImportance=0, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='alerts', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=0, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='open-app-reminders', mName=Ope..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='alerts', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='required', mName=Req..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='essentials', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='updates-available', mName=Upd..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='alerts', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='play-protect', mName=Pla..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='essentials', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='setup', mName=Set..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='essentials', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='maintenance-v2', mName=Mai..., mDescription=, mImportance=1, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='essentials', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=1, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='payments-deals-and-recommendations', mName=Pay..., mDescription=, mImportance=2, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='alerts', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=2, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='auto-open', mName=App..., mDescription=, mImportance=4, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='alerts', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=false, mAllowBubbles=-1, mImportanceLockedDefaultApp=false, mOriginalImp=4, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannelGroup{mId='alerts', mName=Alerts, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
        NotificationChannelGroup{mId='essentials', mName=Essentials, mDescription=, mBlocked=false, mChannels=[], mUserLockedFields=0}
      AppSettings: com.google.android.gms.supervision (10150) importance=NONE userSet=false
      AppSettings: com.android.internal.emulation.pixel_6_pro (10026)
      AppSettings: com.android.virtualmachine.res (10227)
      AppSettings: com.android.providers.settings.auto_generated_rro_product__ (10055)
      AppSettings: com.google.android.soundpicker (10177)
      AppSettings: com.android.cameraextensions (10126)
      AppSettings: com.android.systemui.emulation.pixel_10 (10058)
      AppSettings: com.android.providers.downloads (10110) importance=DEFAULT userSet=false fixedImportance=true
        NotificationChannel{mId='active', mName=In ..., mDescription=, mImportance=1, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=1, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='complete', mName=Don..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
        NotificationChannel{mId='waiting', mName=Que..., mDescription=, mImportance=3, mBypassDnd=false, mLockscreenVisibility=-1000, mSound=content://settings/system/notification_sound, mLights=false, mLightColor=0, mVibrationPattern=null, mVibrationEffect=null, mUserLockedFields=0, mUserVisibleTaskShown=false, mVibrationEnabled=false, mShowBadge=true, mDeleted=false, mDeletedTimeMs=-1, mGroup='null', mAudioAttributes=AudioAttributes: usage=USAGE_NOTIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null, mBlockableSystem=true, mAllowBubbles=-1, mImportanceLockedDefaultApp=true, mOriginalImp=3, mParent=null, mConversationId=null, mDemoted=false, mImportantConvo=false, mLastNotificationUpdateTimeMs=0, isBundle=false, emoji=null}
      AppSettings: com.android.internal.emulation.pixel_6a (10027)
      AppSettings: com.android.systemui.emulation.pixel_10_pro_xl (10061)
Restored without uid:

  Notification listeners:
    Allowed notification listeners:
      com.google.android.as/com.google.android.apps.miphone.aiai.common.notification.service.AiAiNotificationListenerService:com.google.android.apps.nexuslauncher/com.android.launcher3.notification.NotificationListener (user: 0 isPrimary: true)
    Approved uids for user 0: [10192, 10146]
    Has user set:
      userId=0 value={com.google.android.as/com.google.android.apps.miphone.aiai.common.notification.service.AiAiNotificationListenerService, com.google.android.apps.nexuslauncher/com.android.launcher3.notification.NotificationListener}
    All notification listeners (2) enabled for current profiles:
      ComponentInfo{com.google.android.apps.nexuslauncher/com.android.launcher3.notification.NotificationListener}
      ComponentInfo{com.google.android.as/com.google.android.apps.miphone.aiai.common.notification.service.AiAiNotificationListenerService}
    Live notification listeners (9):
      ComponentInfo{android/com.android.server.biometrics.BiometricNotificationLogger} (user -1): android.service.notification.NotificationListenerService$NotificationListenerWrapper@9dbd92e (sc=null) SYSTEM
      ComponentInfo{android/com.android.server.am.AppFGSTracker$NotificationListener} (user -1): android.service.notification.NotificationListenerService$NotificationListenerWrapper@65baccf (sc=null) SYSTEM
      ComponentInfo{com.android.systemui/com.android.systemui.statusbar.NotificationListener} (user -1): android.service.notification.INotificationListener$Stub$Proxy@822c04f (sc=null) SYSTEM
      ComponentInfo{android/com.android.server.media.MediaSessionService$NotificationListener} (user -1): android.service.notification.NotificationListenerService$NotificationListenerWrapper@97e23dc (sc=null) SYSTEM
      ComponentInfo{android/com.android.server.SensitiveContentProtectionManagerService$NotificationListener} (user -1): android.service.notification.NotificationListenerService$NotificationListenerWrapper@1ca7ce5 (sc=null) SYSTEM
      ComponentInfo{android/com.android.server.people.data.DataManager} (user 0): android.service.notification.NotificationListenerService$NotificationListenerWrapper@6193772 (sc=null) SYSTEM
      ComponentInfo{com.google.android.apps.nexuslauncher/com.android.launcher3.notification.NotificationListener} (user 0): android.service.notification.INotificationListener$Stub$Proxy@5a5beb (sc=6b2f908)
      ComponentInfo{com.google.android.as/com.google.android.apps.miphone.aiai.common.notification.service.AiAiNotificationAssistantService} (user 0): android.service.notification.INotificationListener$Stub$Proxy@cdc0b5f (sc=6f17c33) GUEST
      ComponentInfo{com.google.android.as/com.google.android.apps.miphone.aiai.common.notification.service.AiAiNotificationListenerService} (user 0): android.service.notification.INotificationListener$Stub$Proxy@13ced48 (sc=d91d823)
    Snoozed notification listeners (0):
    mListenerHints: 0
    mListenersDisablingEffects: ()

  NotificationListenerStats:

  Notification assistant services:
    Allowed notification assistants:
      com.google.android.as/com.google.android.apps.miphone.aiai.common.notification.service.AiAiNotificationAssistantService (user: 0 isPrimary: true isUserChanged: false)
    Approved uids for user 0: [10146]
    Has user set:
    All notification assistants (1) enabled for current profiles:
      ComponentInfo{com.google.android.as/com.google.android.apps.miphone.aiai.common.notification.service.AiAiNotificationAssistantService}
    Live notification assistants (1):
      ComponentInfo{com.google.android.as/com.google.android.apps.miphone.aiai.common.notification.service.AiAiNotificationAssistantService} (user 0): android.service.notification.INotificationListener$Stub$Proxy@cdc0b5f (sc=6f17c33)
    Snoozed notification assistants (0):
    Unsupported Adjustment keys: 
      0: [key_type, key_summarization]
    (user) Denied Adjustment keys (see rules for KEY_TYPE): 
      user 0: [key_summarization]
    Disallowed adjustment pkg count: 

  Zen Mode:
    mInterruptionFilter=1
    mZenMode=ZEN_MODE_OFF
    mConsolidatedPolicy=NotificationManager.Policy[priorityCategories=PRIORITY_CATEGORY_ALARMS,PRIORITY_CATEGORY_MEDIA,PRIORITY_CATEGORY_MESSAGES,PRIORITY_CATEGORY_CALLS,PRIORITY_CATEGORY_REPEAT_CALLERS,PRIORITY_CATEGORY_CONVERSATIONS,priorityCallSenders=PRIORITY_SENDERS_STARRED,priorityMessageSenders=PRIORITY_SENDERS_STARRED,priorityConvSenders=important,suppressedVisualEffects=SUPPRESSED_EFFECT_SCREEN_OFF,SUPPRESSED_EFFECT_SCREEN_ON,SUPPRESSED_EFFECT_FULL_SCREEN_INTENT,SUPPRESSED_EFFECT_LIGHTS,SUPPRESSED_EFFECT_PEEK,SUPPRESSED_EFFECT_AMBIENT,hasPriorityChannels=true,allowPriorityChannels=true,allowSoundFor=PRIORITY_CATEGORY_ALARMS,PRIORITY_CATEGORY_MEDIA,PRIORITY_CATEGORY_MESSAGES,PRIORITY_CATEGORY_CALLS,PRIORITY_CATEGORY_REPEAT_CALLERS,PRIORITY_CATEGORY_CONVERSATIONS,allowVibrationFor=PRIORITY_CATEGORY_ALARMS,PRIORITY_CATEGORY_MEDIA,PRIORITY_CATEGORY_MESSAGES,PRIORITY_CATEGORY_CALLS,PRIORITY_CATEGORY_REPEAT_CALLERS,PRIORITY_CATEGORY_CONVERSATIONS]
    mConsolidatedDeviceEffects=[]
    mConfigs[u=0]=ZenModeConfig[user=0
hasPriorityChannels=true,
automaticRules={
ZenRule[id=e8f0824a48424476bc499cad45c4a8d6,state=STATE_FALSE,enabled=FALSE,conditionOverride=OVERRIDE_NONE,name=Bedtime,zenMode=ZEN_MODE_IMPORTANT_INTERRUPTIONS,conditionId=condition://com.google.android.apps.wellbeing.azr/winddown,pkg=com.google.android.apps.wellbeing,component=null,configActivity=ComponentInfo{com.google.android.apps.wellbeing/com.google.android.apps.wellbeing.winddown.ui.WindDownAutomaticZenRuleEntryActivity},creationTime=1777904374607,enabler=null,zenPolicy=ZenPolicy{priorityCategories=[reminders=disallow events=disallow messages=disallow calls=allow repeatCallers=allow alarms=allow media=allow system=disallow convs=disallow ], interruptionType=[alarms=all ], visualEffects=[fullScreenIntent=disallow lights=disallow peek=disallow statusBar=disallow badge=disallow ambient=disallow notificationList=disallow ], priorityCallsSenders=starred_contacts, priorityMessagesSenders=none, priorityConversationSenders=none, allowChannels=priority},condition=Condition[state=STATE_FALSE,id=condition://com.google.android.apps.wellbeing.azr/winddown,summary=Bedtime,line1=,line2=,icon=-1,source=SOURCE_UNKNOWN,flags=2],deviceEffects=[grayscale, dimWallpaper, nightMode],allowManualInvocation=true,iconResName=com.google.android.apps.wellbeing:drawable/ic_new_bedtime_mode,triggerDescription=From 11:00 PM – 7:00 AM,type=3,disabledOrigin=4,legacySuppressedEffects=0,lastActivation=null,lastManualActivation=null,lastDeactivation=null,lastManualDeactivation=null],
ZenRule[id=f6eedd0dd63142d086a28fed88d06658,state=STATE_FALSE,enabled=TRUE,conditionOverride=OVERRIDE_NONE,name=Transit,zenMode=ZEN_MODE_IMPORTANT_INTERRUPTIONS,conditionId=condition://com.google.android.settings.intelligence/Transit,pkg=com.google.android.settings.intelligence,component=null,configActivity=ComponentInfo{com.google.android.settings.intelligence/com.google.android.settings.intelligence.modules.transit.impl.TransitSettingsActivity},creationTime=1778088972488,enabler=null,zenPolicy=ZenPolicy{priorityCategories=[reminders=disallow events=disallow messages=allow calls=allow repeatCallers=allow alarms=allow media=allow system=disallow convs=allow ], interruptionType=[alarms=all ], visualEffects=[fullScreenIntent=disallow lights=disallow peek=disallow statusBar=allow badge=allow ambient=disallow notificationList=allow ], priorityCallsSenders=starred_contacts, priorityMessagesSenders=starred_contacts, priorityConversationSenders=important, allowChannels=priority},condition=null,deviceEffects=null,allowManualInvocation=true,iconResName=com.google.android.settings.intelligence:drawable/transit_ic_transit,triggerDescription=Set volume, Bluetooth, and more,type=-1,disabledOrigin=0,legacySuppressedEffects=0,lastActivation=null,lastManualActivation=null,lastDeactivation=null,lastManualDeactivation=null],
ZenRule[id=cf96a56c6bc64b4a8eaa09dcaeb4d912,state=STATE_FALSE,enabled=FALSE,conditionOverride=OVERRIDE_NONE,name=Driving,zenMode=ZEN_MODE_IMPORTANT_INTERRUPTIONS,conditionId=condition://com.google.android.gms/false/2?provider=Driving%20Mode&rule_source=driving_mode&driving_mode_enable_dnd_rule=false,pkg=com.google.android.gms,component=ComponentInfo{com.google.android.gms/com.google.android.location.settings.DrivingConditionProvider},configActivity=ComponentInfo{com.google.android.gms/com.google.android.location.settings.DrivingBehaviorSettingV31Activity},creationTime=1778088967225,enabler=null,zenPolicy=ZenPolicy{priorityCategories=[reminders=disallow events=disallow messages=allow calls=allow repeatCallers=allow alarms=allow media=allow system=disallow convs=allow ], interruptionType=[alarms=all ], visualEffects=[fullScreenIntent=disallow lights=disallow peek=disallow statusBar=allow badge=allow ambient=disallow notificationList=allow ], priorityCallsSenders=starred_contacts, priorityMessagesSenders=starred_contacts, priorityConversationSenders=important, allowChannels=priority},condition=null,deviceEffects=null,allowManualInvocation=true,iconResName=null,triggerDescription=Using device's Bluetooth connection,type=4,disabledOrigin=4,legacySuppressedEffects=0,lastActivation=null,lastManualActivation=null,lastDeactivation=null,lastManualDeactivation=null]},
manualRule=ZenRule[id=MANUAL_RULE,state=STATE_FALSE,enabled=TRUE,conditionOverride=OVERRIDE_NONE,name=null,zenMode=ZEN_MODE_OFF,conditionId=,pkg=android,component=null,configActivity=null,creationTime=0,enabler=null,zenPolicy=ZenPolicy{priorityCategories=[reminders=disallow events=disallow messages=allow calls=allow repeatCallers=allow alarms=allow media=allow system=disallow convs=allow ], interruptionType=[alarms=all ], visualEffects=[fullScreenIntent=disallow lights=disallow peek=disallow statusBar=allow badge=allow ambient=disallow notificationList=allow ], priorityCallsSenders=starred_contacts, priorityMessagesSenders=starred_contacts, priorityConversationSenders=important, allowChannels=priority},condition=null,deviceEffects=null,allowManualInvocation=true,iconResName=null,triggerDescription=null,type=0,disabledOrigin=0,legacySuppressedEffects=0,lastActivation=null,lastManualActivation=null,lastDeactivation=null,lastManualDeactivation=null],
deletedRules={}]
    mUser=0
    mConfig=ZenModeConfig[user=0
hasPriorityChannels=true,
automaticRules={
ZenRule[id=e8f0824a48424476bc499cad45c4a8d6,state=STATE_FALSE,enabled=FALSE,conditionOverride=OVERRIDE_NONE,name=Bedtime,zenMode=ZEN_MODE_IMPORTANT_INTERRUPTIONS,conditionId=condition://com.google.android.apps.wellbeing.azr/winddown,pkg=com.google.android.apps.wellbeing,component=null,configActivity=ComponentInfo{com.google.android.apps.wellbeing/com.google.android.apps.wellbeing.winddown.ui.WindDownAutomaticZenRuleEntryActivity},creationTime=1777904374607,enabler=null,zenPolicy=ZenPolicy{priorityCategories=[reminders=disallow events=disallow messages=disallow calls=allow repeatCallers=allow alarms=allow media=allow system=disallow convs=disallow ], interruptionType=[alarms=all ], visualEffects=[fullScreenIntent=disallow lights=disallow peek=disallow statusBar=disallow badge=disallow ambient=disallow notificationList=disallow ], priorityCallsSenders=starred_contacts, priorityMessagesSenders=none, priorityConversationSenders=none, allowChannels=priority},condition=Condition[state=STATE_FALSE,id=condition://com.google.android.apps.wellbeing.azr/winddown,summary=Bedtime,line1=,line2=,icon=-1,source=SOURCE_UNKNOWN,flags=2],deviceEffects=[grayscale, dimWallpaper, nightMode],allowManualInvocation=true,iconResName=com.google.android.apps.wellbeing:drawable/ic_new_bedtime_mode,triggerDescription=From 11:00 PM – 7:00 AM,type=3,disabledOrigin=4,legacySuppressedEffects=0,lastActivation=null,lastManualActivation=null,lastDeactivation=null,lastManualDeactivation=null],
ZenRule[id=f6eedd0dd63142d086a28fed88d06658,state=STATE_FALSE,enabled=TRUE,conditionOverride=OVERRIDE_NONE,name=Transit,zenMode=ZEN_MODE_IMPORTANT_INTERRUPTIONS,conditionId=condition://com.google.android.settings.intelligence/Transit,pkg=com.google.android.settings.intelligence,component=null,configActivity=ComponentInfo{com.google.android.settings.intelligence/com.google.android.settings.intelligence.modules.transit.impl.TransitSettingsActivity},creationTime=1778088972488,enabler=null,zenPolicy=ZenPolicy{priorityCategories=[reminders=disallow events=disallow messages=allow calls=allow repeatCallers=allow alarms=allow media=allow system=disallow convs=allow ], interruptionType=[alarms=all ], visualEffects=[fullScreenIntent=disallow lights=disallow peek=disallow statusBar=allow badge=allow ambient=disallow notificationList=allow ], priorityCallsSenders=starred_contacts, priorityMessagesSenders=starred_contacts, priorityConversationSenders=important, allowChannels=priority},condition=null,deviceEffects=null,allowManualInvocation=true,iconResName=com.google.android.settings.intelligence:drawable/transit_ic_transit,triggerDescription=Set volume, Bluetooth, and more,type=-1,disabledOrigin=0,legacySuppressedEffects=0,lastActivation=null,lastManualActivation=null,lastDeactivation=null,lastManualDeactivation=null],
ZenRule[id=cf96a56c6bc64b4a8eaa09dcaeb4d912,state=STATE_FALSE,enabled=FALSE,conditionOverride=OVERRIDE_NONE,name=Driving,zenMode=ZEN_MODE_IMPORTANT_INTERRUPTIONS,conditionId=condition://com.google.android.gms/false/2?provider=Driving%20Mode&rule_source=driving_mode&driving_mode_enable_dnd_rule=false,pkg=com.google.android.gms,component=ComponentInfo{com.google.android.gms/com.google.android.location.settings.DrivingConditionProvider},configActivity=ComponentInfo{com.google.android.gms/com.google.android.location.settings.DrivingBehaviorSettingV31Activity},creationTime=1778088967225,enabler=null,zenPolicy=ZenPolicy{priorityCategories=[reminders=disallow events=disallow messages=allow calls=allow repeatCallers=allow alarms=allow media=allow system=disallow convs=allow ], interruptionType=[alarms=all ], visualEffects=[fullScreenIntent=disallow lights=disallow peek=disallow statusBar=allow badge=allow ambient=disallow notificationList=allow ], priorityCallsSenders=starred_contacts, priorityMessagesSenders=starred_contacts, priorityConversationSenders=important, allowChannels=priority},condition=null,deviceEffects=null,allowManualInvocation=true,iconResName=null,triggerDescription=Using device's Bluetooth connection,type=4,disabledOrigin=4,legacySuppressedEffects=0,lastActivation=null,lastManualActivation=null,lastDeactivation=null,lastManualDeactivation=null]},
manualRule=ZenRule[id=MANUAL_RULE,state=STATE_FALSE,enabled=TRUE,conditionOverride=OVERRIDE_NONE,name=null,zenMode=ZEN_MODE_OFF,conditionId=,pkg=android,component=null,configActivity=null,creationTime=0,enabler=null,zenPolicy=ZenPolicy{priorityCategories=[reminders=disallow events=disallow messages=allow calls=allow repeatCallers=allow alarms=allow media=allow system=disallow convs=allow ], interruptionType=[alarms=all ], visualEffects=[fullScreenIntent=disallow lights=disallow peek=disallow statusBar=allow badge=allow ambient=disallow notificationList=allow ], priorityCallsSenders=starred_contacts, priorityMessagesSenders=starred_contacts, priorityConversationSenders=important, allowChannels=priority},condition=null,deviceEffects=null,allowManualInvocation=true,iconResName=null,triggerDescription=null,type=0,disabledOrigin=0,legacySuppressedEffects=0,lastActivation=null,lastManualActivation=null,lastDeactivation=null,lastManualDeactivation=null],
deletedRules={}]
    mSuppressedEffects=0
    mDefaultPhoneApp=ComponentInfo{com.google.android.dialer/com.android.dialer.DialtactsActivity}
    RepeatCallers.mThresholdMinutes=0
    mSubscriptions={}

  Zen Log:
    Interception Events:
    State Changes:
    2026-05-23T15:37:32.120505 - config: readXml (ORIGIN_INIT) from uid 1000,
ZenModeConfig[user=0
hasPriorityChannels=true,
automaticRules={
ZenRule[id=e8f0824a48424476bc499cad45c4a8d6,state=STATE_FALSE,enabled=FALSE,conditionOverride=OVERRIDE_NONE,name=Bedtime,zenMode=ZEN_MODE_IMPORTANT_INTERRUPTIONS,conditionId=condition://com.google.android.apps.wellbeing.azr/winddown,pkg=com.google.android.apps.wellbeing,component=null,configActivity=ComponentInfo{com.google.android.apps.wellbeing/com.google.android.apps.wellbeing.winddown.ui.WindDownAutomaticZenRuleEntryActivity},creationTime=1777904374607,enabler=null,zenPolicy=ZenPolicy{priorityCategories=[reminders=disallow events=disallow messages=disallow calls=allow repeatCallers=allow alarms=allow media=allow system=disallow convs=disallow ], interruptionType=[alarms=all ], visualEffects=[fullScreenIntent=disallow lights=disallow peek=disallow statusBar=disallow badge=disallow ambient=disallow notificationList=disallow ], priorityCallsSenders=starred_contacts, priorityMessagesSenders=none, priorityConversationSenders=none, allowChannels=priority},condition=Condition[state=STATE_FALSE,id=condition://com.google.android.apps.wellbeing.azr/winddown,summary=Bedtime,line1=,line2=,icon=-1,source=SOURCE_UNKNOWN,flags=2],deviceEffects=[grayscale, dimWallpaper, nightMode],allowManualInvocation=true,iconResName=com.google.android.apps.wellbeing:drawable/ic_new_bedtime_mode,triggerDescription=From 11:00 PM – 7:00 AM,type=3,disabledOrigin=4,legacySuppressedEffects=0,lastActivation=null,lastManualActivation=null,lastDeactivation=null,lastManualDeactivation=null],
ZenRule[id=f6eedd0dd63142d086a28fed88d06658,state=STATE_FALSE,enabled=TRUE,conditionOverride=OVERRIDE_NONE,name=Transit,zenMode=ZEN_MODE_IMPORTANT_INTERRUPTIONS,conditionId=condition://com.google.android.settings.intelligence/Transit,pkg=com.google.android.settings.intelligence,component=null,configActivity=ComponentInfo{com.google.android.settings.intelligence/com.google.android.settings.intelligence.modules.transit.impl.TransitSettingsActivity},creationTime=1778088972488,enabler=null,zenPolicy=ZenPolicy{priorityCategories=[reminders=disallow events=disallow messages=allow calls=allow repeatCallers=allow alarms=allow media=allow system=disallow convs=allow ], interruptionType=[alarms=all ], visualEffects=[fullScreenIntent=disallow lights=disallow peek=disallow statusBar=allow badge=allow ambient=disallow notificationList=allow ], priorityCallsSenders=starred_contacts, priorityMessagesSenders=starred_contacts, priorityConversationSenders=important, allowChannels=priority},condition=null,deviceEffects=null,allowManualInvocation=true,iconResName=com.google.android.settings.intelligence:drawable/transit_ic_transit,triggerDescription=Set volume, Bluetooth, and more,type=-1,disabledOrigin=0,legacySuppressedEffects=0,lastActivation=null,lastManualActivation=null,lastDeactivation=null,lastManualDeactivation=null],
ZenRule[id=cf96a56c6bc64b4a8eaa09dcaeb4d912,state=STATE_FALSE,enabled=FALSE,conditionOverride=OVERRIDE_NONE,name=Driving,zenMode=ZEN_MODE_IMPORTANT_INTERRUPTIONS,conditionId=condition://com.google.android.gms/false/2?provider=Driving%20Mode&rule_source=driving_mode&driving_mode_enable_dnd_rule=false,pkg=com.google.android.gms,component=ComponentInfo{com.google.android.gms/com.google.android.location.settings.DrivingConditionProvider},configActivity=ComponentInfo{com.google.android.gms/com.google.android.location.settings.DrivingBehaviorSettingV31Activity},creationTime=1778088967225,enabler=null,zenPolicy=ZenPolicy{priorityCategories=[reminders=disallow events=disallow messages=allow calls=allow repeatCallers=allow alarms=allow media=allow system=disallow convs=allow ], interruptionType=[alarms=all ], visualEffects=[fullScreenIntent=disallow lights=disallow peek=disallow statusBar=allow badge=allow ambient=disallow notificationList=allow ], priorityCallsSenders=starred_contacts, priorityMessagesSenders=starred_contacts, priorityConversationSenders=important, allowChannels=priority},condition=null,deviceEffects=null,allowManualInvocation=true,iconResName=null,triggerDescription=Using device's Bluetooth connection,type=4,disabledOrigin=4,legacySuppressedEffects=0,lastActivation=null,lastManualActivation=null,lastDeactivation=null,lastManualDeactivation=null]},
manualRule=ZenRule[id=MANUAL_RULE,state=STATE_FALSE,enabled=TRUE,conditionOverride=OVERRIDE_NONE,name=null,zenMode=ZEN_MODE_OFF,conditionId=,pkg=android,component=null,configActivity=null,creationTime=0,enabler=null,zenPolicy=ZenPolicy{priorityCategories=[reminders=disallow events=disallow messages=allow calls=allow repeatCallers=allow alarms=allow media=allow system=disallow convs=allow ], interruptionType=[alarms=all ], visualEffects=[fullScreenIntent=disallow lights=disallow peek=disallow statusBar=allow badge=allow ambient=disallow notificationList=allow ], priorityCallsSenders=starred_contacts, priorityMessagesSenders=starred_contacts, priorityConversationSenders=important, allowChannels=priority},condition=null,deviceEffects=null,allowManualInvocation=true,iconResName=null,triggerDescription=null,type=0,disabledOrigin=0,legacySuppressedEffects=0,lastActivation=null,lastManualActivation=null,lastDeactivation=null,lastManualDeactivation=null],
deletedRules={}],
Diff[hasPriorityChannels:false->true,
automaticRule[e8f0824a48424476bc499cad45c4a8d6]:ZenRuleDiff{added},
automaticRule[EVERY_NIGHT_DEFAULT_RULE]:ZenRuleDiff{removed},
automaticRule[f6eedd0dd63142d086a28fed88d06658]:ZenRuleDiff{added},
automaticRule[cf96a56c6bc64b4a8eaa09dcaeb4d912]:ZenRuleDiff{added}]
    2026-05-23T15:37:32.124415 - set_consolidated_policy: NotificationManager.Policy[priorityCategories=PRIORITY_CATEGORY_ALARMS,PRIORITY_CATEGORY_MEDIA,PRIORITY_CATEGORY_MESSAGES,PRIORITY_CATEGORY_CALLS,PRIORITY_CATEGORY_REPEAT_CALLERS,PRIORITY_CATEGORY_CONVERSATIONS,priorityCallSenders=PRIORITY_SENDERS_STARRED,priorityMessageSenders=PRIORITY_SENDERS_STARRED,priorityConvSenders=important,suppressedVisualEffects=SUPPRESSED_EFFECT_SCREEN_OFF,SUPPRESSED_EFFECT_SCREEN_ON,SUPPRESSED_EFFECT_FULL_SCREEN_INTENT,SUPPRESSED_EFFECT_LIGHTS,SUPPRESSED_EFFECT_PEEK,SUPPRESSED_EFFECT_AMBIENT,hasPriorityChannels=true,allowPriorityChannels=true,allowSoundFor=PRIORITY_CATEGORY_ALARMS,PRIORITY_CATEGORY_MEDIA,PRIORITY_CATEGORY_MESSAGES,PRIORITY_CATEGORY_CALLS,PRIORITY_CATEGORY_REPEAT_CALLERS,PRIORITY_CATEGORY_CONVERSATIONS,allowVibrationFor=PRIORITY_CATEGORY_ALARMS,PRIORITY_CATEGORY_MEDIA,PRIORITY_CATEGORY_MESSAGES,PRIORITY_CATEGORY_CALLS,PRIORITY_CATEGORY_REPEAT_CALLERS,PRIORITY_CATEGORY_CONVERSATIONS],readXml
    2026-05-23T15:37:32.124528 - set_zen_mode: off,readXml
    2026-05-23T15:37:32.124551 - set_zen_mode: off,updated setting
    2026-05-23T15:37:32.130899 - set_zen_mode: off,init
    2026-05-23T15:37:32.130908 - set_zen_mode: off,updated setting
    2026-05-23T15:37:40.857306 - config: setAzrState: e8f0824a48424476bc499cad45c4a8d6 (ORIGIN_APP) no changes
    2026-05-23T15:37:40.857396 - set_zen_mode: off,setAzrState: e8f0824a48424476bc499cad45c4a8d6
    2026-05-23T15:37:40.857404 - set_zen_mode: off,updated setting
    2026-05-23T15:37:43.387476 - config: zmc.onServiceAdded:ComponentInfo{com.google.android.gms/com.google.android.location.settings.DrivingConditionProvider} (ORIGIN_SYSTEM) no changes
    2026-05-23T15:37:43.387586 - set_zen_mode: off,zmc.onServiceAdded:ComponentInfo{com.google.android.gms/com.google.android.location.settings.DrivingConditionProvider}
    2026-05-23T15:37:43.387596 - set_zen_mode: off,updated setting
    2026-05-23T15:38:15.169490 - config: setAzrState: e8f0824a48424476bc499cad45c4a8d6 (ORIGIN_APP) no changes
    2026-05-23T15:38:15.169614 - set_zen_mode: off,setAzrState: e8f0824a48424476bc499cad45c4a8d6
    2026-05-23T15:38:15.169623 - set_zen_mode: off,updated setting
    2026-05-23T15:38:15.176904 - config: setAzrState: e8f0824a48424476bc499cad45c4a8d6 (ORIGIN_APP) no changes
    2026-05-23T15:38:15.176968 - set_zen_mode: off,setAzrState: e8f0824a48424476bc499cad45c4a8d6
    2026-05-23T15:38:15.176975 - set_zen_mode: off,updated setting
    Other Events:

  Condition providers:
    Allowed condition providers:
      com.google.android.apps.diagnosticstool:com.google.android.apps.safetyhub:com.google.intelligence.sense:com.google.android.apps.wellbeing:com.google.android.dialer:com.google.android.gms:com.google.android.settings.intelligence:com.google.android.GoogleCamera (user: 0 isPrimary: true)
      com.google.android.as:com.google.android.apps.nexuslauncher (user: 0 isPrimary: false)
    Approved uids for user 0: [10192, 10146, 10147, 10219, 10156, 10157, 10159]
    Has user set:
      userId=0 value={com.google.android.apps.diagnosticstool, com.google.android.apps.safetyhub, com.google.android.as, com.google.intelligence.sense, com.google.android.apps.wellbeing, com.google.android.dialer, com.google.android.gms, com.google.android.apps.nexuslauncher, com.google.android.settings.intelligence, com.google.android.GoogleCamera}
    All condition providers (1) enabled for current profiles:
      ComponentInfo{com.google.android.gms/com.google.android.location.settings.DrivingConditionProvider}
    Live condition providers (4):
      ComponentInfo{android/com.android.server.notification.CountdownConditionProvider} (user 0): android.service.notification.ConditionProviderService$Provider@f3a68c1 (sc=null) SYSTEM
      ComponentInfo{android/com.android.server.notification.ScheduleConditionProvider} (user 0): android.service.notification.ConditionProviderService$Provider@e0f9866 (sc=null) SYSTEM
      ComponentInfo{android/com.android.server.notification.EventConditionProvider} (user 0): android.service.notification.ConditionProviderService$Provider@1f727a7 (sc=null) SYSTEM
      ComponentInfo{android/com.android.server.notification.CustomManualConditionProvider} (user 0): android.service.notification.ConditionProviderService$Provider@67af454 (sc=null) SYSTEM
    Snoozed condition providers (1):
      User: 0
        com.google.android.gms/com.google.android.location.settings.DrivingConditionProvider
    mRecords(0):
    mSystemConditionProviders: {schedule, event, custom_manual, countdown}
    EventConditionProvider:
      mConnected=true
      mRegistered=false
      mBootComplete=true
      mNextAlarmTime=0
      mSubscriptions=
      mTrackers=
        user=0
          mCallback=null
          mRegistered=false
          u=0
    ScheduleConditionProvider:
      mConnected=true
      mRegistered=false
      mSubscriptions=
      snoozed due to alarm: 
      mNextAlarmTime=0
    CountdownConditionProvider:
      mConnected=true
      mTime=0
    CustomManualConditionProvider: ENABLED

  Group summaries:
    0|com.google.android.apps.messaging|g:incoming_message_group_key -> 0|com.google.android.apps.messaging|7|incoming_message_group_key|10154
    0|com.google.android.dialer|g:MissedCallGroup -> 0|com.google.android.dialer|1|GroupSummary_MissedCall|10147

  Usage Stats:
    AggregatedStats{
      key='com.google.android.as',
      numEnqueuedByApp=1,
      numPostedByApp=0,
      numUpdatedByApp=0,
      numRemovedByApp=1,
      numPeopleCacheHit=0,
      numWithStaredPeople=0,
      numWithValidPeople=0,
      numPeopleCacheMiss=0,
      numBlocked=0,
      numSuspendedByAdmin=0,
      numWithActions=0,
      numPrivate=0,
      numSecret=0,
      numInterrupt=0,
      numWithBigText=0,
      numWithBigPicture=0
      numForegroundService=0
      numUserInitiatedJob=0
      numOngoing=0
      numAutoCancel=0
      numWithLargeIcon=0
      numWithInbox=0
      numWithMediaSession=0
      numWithTitle=0
      numWithText=0
      numWithSubText=0
      numWithInfoText=0
      numRateViolations=0
      numAlertViolations=0
      numQuotaViolations=0
      numImagesRemoved=0
      numTooOld=0
      note_imp_noisy_: [0, 0, 0, 0, 0, 0]
      note_imp_quiet_: [0, 0, 0, 0, 0, 0]
      note_importance_: [0, 0, 0, 0, 0, 0]
      numUndecorateRVs=0
    }
    AggregatedStats{
      key='com.google.android.dialer',
      numEnqueuedByApp=12,
      numPostedByApp=4,
      numUpdatedByApp=8,
      numRemovedByApp=0,
      numPeopleCacheHit=20,
      numWithStaredPeople=0,
      numWithValidPeople=0,
      numPeopleCacheMiss=0,
      numBlocked=0,
      numSuspendedByAdmin=0,
      numWithActions=9,
      numPrivate=12,
      numSecret=0,
      numInterrupt=0,
      numWithBigText=0,
      numWithBigPicture=0
      numForegroundService=0
      numUserInitiatedJob=0
      numOngoing=0
      numAutoCancel=12
      numWithLargeIcon=12
      numWithInbox=0
      numWithMediaSession=0
      numWithTitle=12
      numWithText=12
      numWithSubText=0
      numWithInfoText=0
      numRateViolations=0
      numAlertViolations=3
      numQuotaViolations=0
      numImagesRemoved=0
      numTooOld=0
      note_imp_noisy_: [0, 0, 0, 12, 0, 0]
      note_imp_quiet_: [0, 0, 0, 0, 0, 0]
      note_importance_: [0, 0, 0, 12, 0, 0]
      numUndecorateRVs=0
    }
    AggregatedStats{
      key='android',
      numEnqueuedByApp=1,
      numPostedByApp=1,
      numUpdatedByApp=0,
      numRemovedByApp=0,
      numPeopleCacheHit=3,
      numWithStaredPeople=0,
      numWithValidPeople=0,
      numPeopleCacheMiss=0,
      numBlocked=0,
      numSuspendedByAdmin=0,
      numWithActions=0,
      numPrivate=1,
      numSecret=0,
      numInterrupt=0,
      numWithBigText=1,
      numWithBigPicture=0
      numForegroundService=0
      numUserInitiatedJob=0
      numOngoing=0
      numAutoCancel=0
      numWithLargeIcon=0
      numWithInbox=0
      numWithMediaSession=0
      numWithTitle=1
      numWithText=1
      numWithSubText=0
      numWithInfoText=0
      numRateViolations=0
      numAlertViolations=0
      numQuotaViolations=0
      numImagesRemoved=0
      numTooOld=0
      note_imp_noisy_: [0, 0, 0, 1, 0, 0]
      note_imp_quiet_: [0, 0, 0, 0, 0, 0]
      note_importance_: [0, 0, 0, 1, 0, 0]
      numUndecorateRVs=0
    }
    AggregatedStats{
      key='com.google.android.deskclock',
      numEnqueuedByApp=1,
      numPostedByApp=1,
      numUpdatedByApp=0,
      numRemovedByApp=0,
      numPeopleCacheHit=3,
      numWithStaredPeople=0,
      numWithValidPeople=0,
      numPeopleCacheMiss=0,
      numBlocked=0,
      numSuspendedByAdmin=0,
      numWithActions=1,
      numPrivate=1,
      numSecret=0,
      numInterrupt=0,
      numWithBigText=0,
      numWithBigPicture=0
      numForegroundService=0
      numUserInitiatedJob=0
      numOngoing=0
      numAutoCancel=1
      numWithLargeIcon=1
      numWithInbox=0
      numWithMediaSession=0
      numWithTitle=0
      numWithText=0
      numWithSubText=0
      numWithInfoText=0
      numRateViolations=0
      numAlertViolations=0
      numQuotaViolations=0
      numImagesRemoved=0
      numTooOld=0
      note_imp_noisy_: [0, 0, 0, 0, 0, 0]
      note_imp_quiet_: [0, 0, 0, 0, 1, 0]
      note_importance_: [0, 0, 0, 1, 0, 0]
      numUndecorateRVs=0
    }
    AggregatedStats{
      key='__global',
      numEnqueuedByApp=19,
      numPostedByApp=10,
      numUpdatedByApp=8,
      numRemovedByApp=1,
      numPeopleCacheHit=30,
      numWithStaredPeople=0,
      numWithValidPeople=0,
      numPeopleCacheMiss=0,
      numBlocked=0,
      numSuspendedByAdmin=0,
      numWithActions=13,
      numPrivate=18,
      numSecret=0,
      numInterrupt=0,
      numWithBigText=1,
      numWithBigPicture=0
      numForegroundService=0
      numUserInitiatedJob=0
      numOngoing=0
      numAutoCancel=17
      numWithLargeIcon=17
      numWithInbox=0
      numWithMediaSession=0
      numWithTitle=16
      numWithText=16
      numWithSubText=0
      numWithInfoText=0
      numRateViolations=0
      numAlertViolations=0
      numQuotaViolations=0
      numImagesRemoved=0
      numTooOld=0
      note_imp_noisy_: [0, 0, 0, 16, 1, 0]
      note_imp_quiet_: [0, 0, 0, 0, 1, 0]
      note_importance_: [0, 0, 0, 14, 4, 0]
      numUndecorateRVs=0
    }
    AggregatedStats{
      key='com.google.android.apps.messaging',
      numEnqueuedByApp=4,
      numPostedByApp=4,
      numUpdatedByApp=0,
      numRemovedByApp=0,
      numPeopleCacheHit=4,
      numWithStaredPeople=0,
      numWithValidPeople=0,
      numPeopleCacheMiss=0,
      numBlocked=0,
      numSuspendedByAdmin=0,
      numWithActions=3,
      numPrivate=4,
      numSecret=0,
      numInterrupt=0,
      numWithBigText=0,
      numWithBigPicture=0
      numForegroundService=0
      numUserInitiatedJob=0
      numOngoing=0
      numAutoCancel=4
      numWithLargeIcon=4
      numWithInbox=0
      numWithMediaSession=0
      numWithTitle=3
      numWithText=3
      numWithSubText=0
      numWithInfoText=0
      numRateViolations=0
      numAlertViolations=0
      numQuotaViolations=0
      numImagesRemoved=0
      numTooOld=0
      note_imp_noisy_: [0, 0, 0, 3, 1, 0]
      note_imp_quiet_: [0, 0, 0, 0, 0, 0]
      note_importance_: [0, 0, 0, 0, 4, 0]
      numUndecorateRVs=0
    }
    mStatsArrays.size(): 1
    mStats.size(): 6

  TimeToLive alarms:
    mKeys [Pair{259205251 0|android|2345|ChFBbmRyb2lkTG9ja1NjcmVlbhIRTm9TY3JlZW5Mb2NrSXNzdWUYAA==|1000}, Pair{259212670 0|com.google.android.deskclock|2147483642|null|10173}, Pair{259219322 0|com.google.android.dialer|1|GroupSummary_MissedCall|10147}, Pair{259220113 0|com.google.android.dialer|1|MissedCall_content://call_log/calls/3|10147}, Pair{259220223 0|com.google.android.dialer|1|MissedCall_content://call_log/calls/2|10147}, Pair{259220378 0|com.google.android.dialer|1|MissedCall_content://call_log/calls/1|10147}, Pair{259232862 0|com.google.android.apps.messaging|2|com.google.android.apps.messaging:incoming_message:1|10154}, Pair{259232864 0|com.google.android.apps.messaging|2|com.google.android.apps.messaging:incoming_message:2|10154}, Pair{259232865 0|com.google.android.apps.messaging|2|com.google.android.apps.messaging:incoming_message:3|10154}]

  GroupHelper:
    Ungrouped notifications:
        0|com.google.android.deskclock|g:Aggregate_AlertingSection
            0|com.google.android.deskclock|2147483642|null|10173
        0|android|g:Aggregate_AlertingSection
            0|android|2345|ChFBbmRyb2lkTG9ja1NjcmVlbhIRTm9TY3JlZW5Mb2NrSXNzdWUYAA==|1000

  Configurable parameters:
    nls_completion_duration_ms=+10s0ms 
```
