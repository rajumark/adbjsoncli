# `adbjson shell dumpsys media_session`

## adbjson

**Command:**
```bash
adbjson shell dumpsys media_session
```

**Output:**
```json
{
  "status": 0,
  "output": "MEDIA SESSION SERVICE (dumpsys media_session)\n\n5 sessions listeners.\nGlobal priority session is com.android.server.telecom/HeadsetMediaButton/1 (userId=0)\n  HeadsetMediaButton com.android.server.telecom/HeadsetMediaButton/1 (userId=0)\n    ownerPid=644, ownerUid=1000, userId=0\n    package=com.android.server.telecom\n    launchIntent=null\n    mediaButtonReceiver=null\n    active=false\n    flags=65537\n    rating type=0\n    controllers: 0\n    state=null\n    audioAttrs=AudioAttributes: usage=USAGE_VOICE_COMMUNICATION content=CONTENT_TYPE_SPEECH flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null\n    volumeType=LOCAL, controlType=ABSOLUTE, max=0, current=0, volumeControlId=null\n    metadata: null\n    queueTitle=null, size=0\nUser Records:\nRecord for full_user=0\n  Volume key long-press listener: null\n  Volume key long-press listener package: \n  Media key listener: null\n  Media key listener package: \n  OnMediaKeyEventDispatchedListener: added 0 listener(s)\n  OnMediaKeyEventSessionChangedListener: added 0 listener(s)\n  Last MediaButtonReceiver: null\n  Media button session is null\n  Sessions Stack - have 0 sessions:\nAudio playback (lastly played comes first)\nMedia session config:\n  media_session_calback_fgs_allowlist_duration_ms: [cur: 10000, def: 10000]\n  media_session_callback_fgs_while_in_use_temp_allow_duration_ms: [cur: 10000, def: 10000]\n  media_session_temp_user_engaged_duration_ms: [cur: 600000, def: 600000]"
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys media_session
```

**Output:**
```
MEDIA SESSION SERVICE (dumpsys media_session)

5 sessions listeners.
Global priority session is com.android.server.telecom/HeadsetMediaButton/1 (userId=0)
  HeadsetMediaButton com.android.server.telecom/HeadsetMediaButton/1 (userId=0)
    ownerPid=644, ownerUid=1000, userId=0
    package=com.android.server.telecom
    launchIntent=null
    mediaButtonReceiver=null
    active=false
    flags=65537
    rating type=0
    controllers: 0
    state=null
    audioAttrs=AudioAttributes: usage=USAGE_VOICE_COMMUNICATION content=CONTENT_TYPE_SPEECH flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null
    volumeType=LOCAL, controlType=ABSOLUTE, max=0, current=0, volumeControlId=null
    metadata: null
    queueTitle=null, size=0
User Records:
Record for full_user=0
  Volume key long-press listener: null
  Volume key long-press listener package: 
  Media key listener: null
  Media key listener package: 
  OnMediaKeyEventDispatchedListener: added 0 listener(s)
  OnMediaKeyEventSessionChangedListener: added 0 listener(s)
  Last MediaButtonReceiver: null
  Media button session is null
  Sessions Stack - have 0 sessions:
Audio playback (lastly played comes first)
Media session config:
  media_session_calback_fgs_allowlist_duration_ms: [cur: 10000, def: 10000]
  media_session_callback_fgs_while_in_use_temp_allow_duration_ms: [cur: 10000, def: 10000]
  media_session_temp_user_engaged_duration_ms: [cur: 600000, def: 600000]
```
