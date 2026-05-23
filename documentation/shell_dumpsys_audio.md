# `adbjson shell dumpsys audio`

## adbjson

**Command:**
```bash
adbjson shell dumpsys audio
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "(Total messages": "0, polling=true, quitting=false)",
    "(not logged)  AudioPlaybackConfiguration piid:15 deviceIds:[] type:android.media.SoundPool u/pid:1000/644 state:idle attr:AudioAttributes": "usage=USAGE_ASSISTANCE_SONIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null sessionId:0 mutedState:none  FormatInfo{isSpatialized=false, channelMask=0x0, sampleRate=0}",
    "- STREAM_ACCESSIBILITY (aliased to": "STREAM_MUSIC):",
    "- STREAM_DTMF (aliased to": "STREAM_RING):",
    "- STREAM_SYSTEM (aliased to": "STREAM_RING):",
    "- STREAM_SYSTEM_ENFORCED (aliased to": "STREAM_RING):",
    "- STREAM_TTS (aliased to": "STREAM_MUSIC):",
    "-AudioAttributes": "usage=USAGE_CALL_ASSISTANT content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null",
    "05-23 15:57:46:888 init()": "setting state to STATE_NOT_SUPPORTED due to effect not expected",
    "05-23 15:57:46:953 registerAudioPolicy for android.media.audiopolicy.AudioPolicy$1@b9afad9 u/pid:1000/644 with config:reg:32:ap:0 for virtual deviceId": "0 isFocusPolicy:false isTestFocusPolicy:false",
    "05-23 15:57:47:024 setModeOwnerToken": "null -> []",
    "05-23 15:57:47:027 stopBluetoothSco": "resetBluetoothSco",
    "05-23 15:57:47:027 updateCommunicationRoute, preferredCommunicationDevice": "null eventSource: resetBluetoothSco",
    "05-23 15:57:47:030 new player piid:15 uid/pid:1000/644 package:com.android.server.telecom type:android.media.SoundPool attr:AudioAttributes": "usage=USAGE_ASSISTANCE_SONIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null session:0",
    "05-23 15:57:47:266 getAvailableCommunicationDevices": "no EARPIECE!",
    "05-23 15:57:47:432 new player piid:23 uid/pid:10195/889 package:com.android.systemui type:android.media.SoundPool attr:AudioAttributes": "usage=USAGE_ASSISTANCE_SONIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null session:0",
    "05-23 15:57:47:559 BT profile service": "connecting HEADSET profile",
    "05-23 15:59:06:095 BT profile service": "disconnecting A2DP profile",
    "05-23 15:59:06:096 BT profile service": "disconnecting HEADSET profile",
    "05-23 15:59:18:449 BT profile service": "connecting HEADSET profile",
    "05-23 15:59:18:450 BT profile service": "connecting HEARING_AID profile",
    "05-23 15:59:41:793 BT profile service": "disconnecting HEADSET profile",
    "Active communication device": "AudioDeviceAttributes: role:output type:speaker addr: name:sdk_gphone16k_arm64 profiles:[] descriptors:[]",
    "Applied Preferred communication device": "null",
    "AudioAttributes": "usage=USAGE_ASSISTANT content=CONTENT_TYPE_UNKNOWN flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null forVolume: true stream: STREAM_ASSISTANT(11)",
    "AudioDeviceAttributes": "role:output type:speaker addr: name: profiles:[] descriptors:[]",
    "AudioPlaybackConfiguration piid:23 deviceIds:[] type:android.media.SoundPool u/pid:10195/889 state:idle attr:AudioAttributes": "usage=USAGE_ASSISTANCE_SONIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null sessionId:0 mutedState:none  FormatInfo{isSpatialized=false, channelMask=0x0, sampleRate=0}",
    "Bluetooth A2DP suspended, requested ext": "false, requested int: false, applied false",
    "Bluetooth LE Audio suspended, requested ext": "false, requested int: false, applied false",
    "Bluetooth SCO on, requested": "false, applied: false",
    "Computed Preferred communication device": "null",
    "Current": "40000000 (default): 5",
    "Current time": "Sat May 23 16:04:37 GMT+05:30 2026",
    "Default attenuation (dB)": "-6",
    "Device type": "0x20000004, driving stream 3",
    "Devices": "speaker",
    "Legacy Stream Type": "11 Volume Group Id: 3",
    "Max": "15",
    "Min": "0",
    "Mode owner token": "null",
    "Muted": "false",
    "Muted Internally": "false",
    "Name": "AUDIO_STREAM_CALL_ASSISTANT Id: 14",
    "Notify on duck": "true",
    "Stream volumes (device": "index)",
    "Streams": "UNKNOWN_STREAM_14",
    "Supported Legacy Stream Types": "{ 14 }",
    "Volume Group": "AUDIO_STREAM_ASSISTANT",
    "Volume Groups (device": "index)",
    "last cache clear time": "05-23 15:57:47:322",
    "mA2dp": "null",
    "mAccessibilityStrategyId": "7",
    "mAudioModeOwner": "AudioDeviceBroker$AudioModeInfo[mMode=0, mPid=0, mUid=0, mToken=null]",
    "mAvrcpAbsVolSupported": "false",
    "mBluetoothHeadset": "null",
    "mBluetoothHeadsetDevice": "null",
    "mCommunicationStrategyId": "1",
    "mHearingAid": "null",
    "mLeAudio": "null",
    "mLeAudioPeripheral": "null",
    "mScoAudioState": "SCO_STATE_INACTIVE",
    "mScoManagedByAudio": "false",
    "mSupportsBleHearingAids": "false",
    "spatialized speaker masks": "none"
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys audio
```

**Output:**
```
AudioService Dumpsys
Current time: Sat May 23 16:04:37 GMT+05:30 2026

# IPC
## Native audioserver lifecycle events
05-23 15:57:46:871 AudioService()
05-23 15:57:46:876 Controller start task complete

## AudioHandler
Message handler (watch for unhandled messages):
  Handler (com.android.server.audio.AudioService$AudioHandler) {44c275b} @ 416981
    Looper (AudioService, tid 139) {370881}
      (MessageQueue is using DeliQueue implementation)
      (Total messages: 0, polling=true, quitting=false)


# Stream activity
## PlaybackActivityMonitor
  playback listeners:
 PlayMonitorClient:S uid:1000 pid:644
 PlayMonitorClient:S uid:1000 pid:644
 PlayMonitorClient:S uid:1000 pid:644



  players:
(not logged)  AudioPlaybackConfiguration piid:15 deviceIds:[] type:android.media.SoundPool u/pid:1000/644 state:idle attr:AudioAttributes: usage=USAGE_ASSISTANCE_SONIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null sessionId:0 mutedState:none  FormatInfo{isSpatialized=false, channelMask=0x0, sampleRate=0}
  AudioPlaybackConfiguration piid:23 deviceIds:[] type:android.media.SoundPool u/pid:10195/889 state:idle attr:AudioAttributes: usage=USAGE_ASSISTANCE_SONIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null sessionId:0 mutedState:none  FormatInfo{isSpatialized=false, channelMask=0x0, sampleRate=0}

  ducked players piids:

  faded out players piids:

  muted player piids due to call/ring:

  banned uids:


  muted players (piids) awaiting device connection:


  current piid to portId map:


### Playback activity
05-23 15:57:47:030 new player piid:15 uid/pid:1000/644 package:com.android.server.telecom type:android.media.SoundPool attr:AudioAttributes: usage=USAGE_ASSISTANCE_SONIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null session:0
05-23 15:57:47:432 new player piid:23 uid/pid:10195/889 package:com.android.systemui type:android.media.SoundPool attr:AudioAttributes: usage=USAGE_ASSISTANCE_SONIFICATION content=CONTENT_TYPE_SONIFICATION flags=0x800(FLAG_MUTE_HAPTIC)  tags= bundle=null session:0


  allowed capture policies:
## RecordActivityMonitor


### Recording Activity

  mHardeningOverride=0
## Hardening enforcement

## MediaFocusControl
  has focus policy:false

Audio Focus stack entries (last is top of stack):


No external focus policy



 Notify on duck:  true


focus commands as seen by MediaFocusControl
Multi Audio Focus enabled :false
------------------------------
AudioFocus isolation - Uid to active focus request:
<Flag disabled>


No additional audio focus environments.

# Routing
## AudioDeviceBroker
  Message handler (watch for unhandled messages):
    Handler (com.android.server.audio.AudioDeviceBroker$BrokerHandler) {3ccabf8} @ 416981
      Looper (AudioDeviceBroker, tid 136) {540dd3b}
        (MessageQueue is using DeliQueue implementation)
        (Total messages: 0, polling=true, quitting=false)

    BECOMING_NOISY_INTENT_DEVICES_SET=
 0x400 0x800 0x8000000 0x80 0x200 0x100 0x2000 0x4000 0x4000000 0x20000000 0x20000001 0x20000 0x20000002 0x4 0x20000004 0x8
    Preferred devices for strategy:

    Non-default devices for strategy:

    Connected devices:

    APM Connected device:

    Preferred devices for capture preset:\n\n... (truncated,      758 lines total)
```
