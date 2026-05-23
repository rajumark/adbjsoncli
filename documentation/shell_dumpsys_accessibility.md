# `adbjson shell dumpsys accessibility`

## adbjson

**Command:**
```bash
adbjson shell dumpsys accessibility
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "1 valid display": "0",
    "Client list broadcasts count": "-1",
    "Client list callbacks": "5",
    "Client list killed": "false",
    "MagnificationConfig[mode": "1, activated: false, scale: 1.0, centerX: 540.0, centerY: 1180.5]",
    "Number of proxy connections": "0",
    "SystemUI uid": "10195",
    "button mode": "1"
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys accessibility
```

**Output:**
```
ACCESSIBILITY MANAGER (dumpsys accessibility)

currentUserId=0
hasMagnificationConnection=false
Magnifier on display#0
    MagnificationConfig[mode: 1, activated: false, scale: 1.0, centerX: 540.0, centerY: 1180.5] 
    Magnification region=SkRegion((0,0,1080,2361))
    IdOfLastServiceToMagnify=-1
    SupportWindowMagnification=true
    WindowMagnificationConnectionState=DISCONNECTED
User state[
     attributes:{id=0, touchExplorationEnabled=false, serviceHandlesDoubleTap=false, requestMultiFingerGestures=false, requestTwoFingerPassthrough=false, sendMotionEventsEnabledfalse, displayMagnificationEnabled=false, autoclickEnabled=false, nonInteractiveUiTimeout=0, interactiveUiTimeout=0, installedServiceCount=5, magnificationModes={0=1}, magnificationCapabilities=3, audioDescriptionByDefaultEnabled=false, magnificationCursorFollowingMode=0, magnificationFollowTypingEnabled=true, magnificationFollowKeyboardEnabled=true, alwaysOnMagnificationEnabled=true}
     button mode: 1
     shortcut key:{}
     button:{}
     button target:{null}
     gesture:{}
     qs shortcut targets:{}
     a11y tiles in QS panel:{}
     top row key:{}
     quick access dialog:{}
     keyboard shortcuts targets:{}
     Bound services:{}
     Enabled services:{}
     Binding services:{}
     Crashed services:{}
     Client list info:{
          Client list callbacks: 7
          Client list killed: false
          Client list broadcasts count: -1
          Registered clients:{
[com.google.android.apps.wellbeing][com.google.android.googlequicksearchbox][com.google.android.contacts][com.google.android.inputmethod.latin][com.google.android.youtube][com.google.android.apps.photos][com.google.android.apps.nexuslauncher]}]
Global Info [ Top focused display Id = 0
     Active Window Id = -1
     Top Focused Window Id = -1
     Accessibility Focused Window Id = -1 ]


Window attributes:[{}]
Global client list info:{
    Client list callbacks: 5
    Client list killed: false
    Client list broadcasts count: -1
    Registered clients:{
[com.google.android.gms][com.android.server.telecom, com.android.emulator.multidisplay, com.android.inputdevices, com.android.providers.settings, com.android.keychain, com.android.localtransport, android, com.android.DeviceAsWebcam, com.android.location.fused, com.android.settings, com.android.dynsystem][com.google.android.permissioncontroller][com.google.android.as][com.android.systemui]

Proxy manager state:
    Number of proxy connections: 0
    Registered proxy connections:
Accessibility Display Listener:
    SystemUI uid: 10195
    1 valid display: 0
```
