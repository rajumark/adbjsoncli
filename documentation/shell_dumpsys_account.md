# `adbjson shell dumpsys account`

## adbjson

**Command:**
```bash
adbjson shell dumpsys account
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "Accounts": "0",
    "Active Sessions": "0",
    "RegisteredServicesCache": "4 services",
    "ServiceInfo": "AuthenticatorDescription {type=com.google.android.gm.legacyimap}, ComponentInfo{com.google.android.gm/com.android.email.service.LegacyImapAuthenticatorService}, uid 10174, lastUpdateTime 1778994578779"
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys account
```

**Output:**
```
User UserInfo{0:Owner:4c13}:
  Accounts: 0
  
  AccountId, Action_Type, timestamp, UID, TableName, Key
  Accounts History
  
  Active Sessions: 0
  
  RegisteredServicesCache: 4 services
    ServiceInfo: AuthenticatorDescription {type=com.google.android.gm.pop3}, ComponentInfo{com.google.android.gm/com.android.email.service.Pop3AuthenticatorService}, uid 10174, lastUpdateTime 1778994578779
    ServiceInfo: AuthenticatorDescription {type=com.google}, ComponentInfo{com.google.android.gms/com.google.android.gms.auth.account.authenticator.GoogleAccountAuthenticatorService}, uid 10219, lastUpdateTime 1778392535660
    ServiceInfo: AuthenticatorDescription {type=com.google.android.gm.exchange}, ComponentInfo{com.google.android.gm/com.android.email.service.EasAuthenticatorService}, uid 10174, lastUpdateTime 1778994578779
    ServiceInfo: AuthenticatorDescription {type=com.google.android.gm.legacyimap}, ComponentInfo{com.google.android.gm/com.android.email.service.LegacyImapAuthenticatorService}, uid 10174, lastUpdateTime 1778994578779
  
  Account visibility:
  
```
