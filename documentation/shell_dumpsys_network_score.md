# `adbjson shell dumpsys network_score`

## adbjson

**Command:**
```bash
adbjson shell dumpsys network_score
```

**Output:**
```json
{
  "status": 0,
  "output": "Current scorer: NetworkScorerAppData{packageUid=10219, mRecommendationService=ComponentInfo{com.google.android.gms/com.google.android.gms.chimera.PersistentBoundBrokerService}, mRecommendationServiceLabel=Google Play services, mEnableUseOpenWifiActivity=null, mNetworkAvailableNotificationChannelId=null}\nScoringServiceConnection: ComponentInfo{com.google.android.gms/com.google.android.gms.chimera.PersistentBoundBrokerService}, bound: true, connected: true"
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys network_score
```

**Output:**
```
Current scorer: NetworkScorerAppData{packageUid=10219, mRecommendationService=ComponentInfo{com.google.android.gms/com.google.android.gms.chimera.PersistentBoundBrokerService}, mRecommendationServiceLabel=Google Play services, mEnableUseOpenWifiActivity=null, mNetworkAvailableNotificationChannelId=null}
ScoringServiceConnection: ComponentInfo{com.google.android.gms/com.google.android.gms.chimera.PersistentBoundBrokerService}, bound: true, connected: true
```
