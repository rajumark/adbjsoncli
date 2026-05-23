# `adbjson shell dumpsys role`

## adbjson

**Command:**
```bash
adbjson shell dumpsys role
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
adb shell dumpsys role
```

**Output:**
```
ROLE STATE (dumpsys role):
{
  user_states={
    user_id=0
    version=1
    packages_hash=CD19BFA8ED561FCCCEF1EDD3A0175AA48F056B0E55E6201E9171B954E840173A
    roles=[
      {
        name=android.app.role.SYSTEM_CONTACTS
        fallback_enabled=true
        holders=com.google.android.contacts
      }
      {
        name=android.app.role.FINANCED_DEVICE_KIOSK
        fallback_enabled=true
      }
      {
        name=android.app.role.SYSTEM_SETTINGS_INTELLIGENCE
        fallback_enabled=true
        holders=com.google.android.settings.intelligence
      }
      {
        name=android.app.role.RETAIL_DEMO
        fallback_enabled=true
      }
      {
        name=android.app.role.SYSTEM_SUPERVISION
        fallback_enabled=true
      }
      {
        name=android.app.role.SYSTEM_CALL_STREAMING
        fallback_enabled=true
        holders=com.google.android.gms
      }
      {
        name=android.app.role.DEVICE_POLICY_MANAGEMENT
        fallback_enabled=true
      }
      {
        name=android.app.role.SUPERVISION
        fallback_enabled=true
      }
      {
        name=android.app.role.COMPANION_DEVICE_WATCH
        fallback_enabled=true
      }
      {
        name=android.app.role.COMPANION_DEVICE_GLASSES
        fallback_enabled=true
      }
      {
        name=android.app.role.CALL_SCREENING
        fallback_enabled=true
      }
      {
        name=android.app.role.SYSTEM_NOTIFICATION_INTELLIGENCE
        fallback_enabled=true
        holders=com.google.android.as
      }
      {
        name=android.app.role.SYSTEM_BLUETOOTH_STACK
        fallback_enabled=true
        holders=com.google.android.bluetooth
      }
      {
        name=android.app.role.COMPANION_DEVICE_APP_STREAMING
        fallback_enabled=true
      }
      {
        name=android.app.role.SYSTEM_DOCUMENT_MANAGER
        fallback_enabled=true
        holders=com.google.android.documentsui
      }
      {
        name=android.app.role.SYSTEM_GALLERY
        fallback_enabled=true
        holders=com.google.android.apps.photos
      }
      {
        name=android.app.role.COMPANION_DEVICE_MEDICAL
        fallback_enabled=true
      }
      {
        name=android.app.role.SYSTEM_APP_PROTECTION_SERVICE
        fallback_enabled=true
        holders=com.google.android.odad
      }
      {
        name=android.app.role.SYSTEM_SHELL
        fallback_enabled=true
        holders=com.android.shell
      }
      {
        name=android.app.role.SYSTEM_DEPENDENCY_INSTALLER
        fallback_enabled=true
      }
      {
        name=android.app.role.SYSTEM_WELLBEING
        fallback_enabled=true
        holders=com.google.android.apps.wellbeing\n\n... (truncated,      218 lines total)
```
