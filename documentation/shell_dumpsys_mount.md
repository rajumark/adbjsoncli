# `adbjson shell dumpsys mount`

## adbjson

**Command:**
```bash
adbjson shell dumpsys mount
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "CE unlocked users": "[0]",
    "Internal storage (null) total size": "8000000000 (7629 MiB)",
    "Last maintenance": "2026-05-20 17:31:43",
    "Media cloud providers": "{0=com.google.android.apps.photos.cloudpicker}",
    "Primary storage UUID": "null",
    "System unlocked users": "[0]"
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys mount
```

**Output:**
```
Disks:

Volumes:
  VolumeInfo{emulated;0}:
    type=EMULATED diskId=null partGuid= mountFlags=PRIMARY|VISIBLE_FOR_WRITE mountUserId=0 state=MOUNTED 
    fsType=null fsUuid=null fsLabel=null 
    path=/storage/emulated internalPath=/data/media 

Records:

Primary storage UUID: null

Internal storage (null) total size: 8000000000 (7629 MiB)

CE unlocked users: [0]
System unlocked users: [0]

mObbMounts:

mObbPathToStateMap:

Media cloud providers: {0=com.google.android.apps.photos.cloudpicker}

Last maintenance: 2026-05-20 17:31:43
```
