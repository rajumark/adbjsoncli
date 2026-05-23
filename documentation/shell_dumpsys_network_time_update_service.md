# `adbjson shell dumpsys network_time_update_service`

## adbjson

**Command:**
```bash
adbjson shell dumpsys network_time_update_service
```

**Output:**
```json
{
  "status": 0,
  "output": "mDefaultNetwork=null\nmNormalPollingIntervalMillis=64800000\nmShortPollingIntervalMillis=60000\nmTryAgainTimesMax=3\nmLastRefreshAttemptElapsedRealtimeMillis=PT4.434S (4434)\nmTryAgainCounter=0\n\nNtpTrustedTime:\n  getNtpConfig()=NtpConnectionInfo{mServerUris=[ntp://time.android.com], mTimeout=PT5S}\n  mNtpConfigForTests=null\n  mLastSuccessfulNtpServerUri=ntp://time.android.com\n  mTimeResult=TimeResult{unixEpochTime=2026-05-23T10:27:49.084Z, elapsedRealtime=PT4.535S, mUncertaintyMillis=43, mNtpServerSocketAddress=time.android.com/216.239.35.4:123}\n  mTimeResult.getAgeMillis()=PT6M54.235S\n\nDebug log:\n  PT4.536S / 2026-05-23T10:27:45.548445Z - refreshIfRequiredAndReschedule: network=100, reason=network available, initialTimeResult=null, shouldAttemptRefresh=true, refreshSuccessful=true, currentElapsedRealtimeMillis=PT4.535S (4535), latestTimeResult=TimeResult{unixEpochTime=2026-05-23T10:27:49.084Z, elapsedRealtime=PT4.535S, mUncertaintyMillis=43, mNtpServerSocketAddress=time.android.com/216.239.35.4:123}, mTryAgainCounter=0, refreshAttemptDelayMillis=64800000, nextRefreshElapsedRealtimeMillis=PT18H4.535S (64804535)\n  PT2M53.263S / 2026-05-23T10:30:34.275826Z - refreshIfRequiredAndReschedule: network=101, reason=network available, initialTimeResult=TimeResult{unixEpochTime=2026-05-23T10:27:49.084Z, elapsedRealtime=PT4.535S, mUncertaintyMillis=43, mNtpServerSocketAddress=time.android.com/216.239.35.4:123}, shouldAttemptRefresh=false, refreshSuccessful=false, currentElapsedRealtimeMillis=PT2M53.263S (173263), latestTimeResult=TimeResult{unixEpochTime=2026-05-23T10:27:49.084Z, elapsedRealtime=PT4.535S, mUncertaintyMillis=43, mNtpServerSocketAddress=time.android.com/216.239.35.4:123}, mTryAgainCounter=0, refreshAttemptDelayMillis=64800000, nextRefreshElapsedRealtimeMillis=PT18H4.535S (64804535)"
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys network_time_update_service
```

**Output:**
```
mDefaultNetwork=null
mNormalPollingIntervalMillis=64800000
mShortPollingIntervalMillis=60000
mTryAgainTimesMax=3
mLastRefreshAttemptElapsedRealtimeMillis=PT4.434S (4434)
mTryAgainCounter=0

NtpTrustedTime:
  getNtpConfig()=NtpConnectionInfo{mServerUris=[ntp://time.android.com], mTimeout=PT5S}
  mNtpConfigForTests=null
  mLastSuccessfulNtpServerUri=ntp://time.android.com
  mTimeResult=TimeResult{unixEpochTime=2026-05-23T10:27:49.084Z, elapsedRealtime=PT4.535S, mUncertaintyMillis=43, mNtpServerSocketAddress=time.android.com/216.239.35.4:123}
  mTimeResult.getAgeMillis()=PT6M54.253S

Debug log:
  PT4.536S / 2026-05-23T10:27:45.548445Z - refreshIfRequiredAndReschedule: network=100, reason=network available, initialTimeResult=null, shouldAttemptRefresh=true, refreshSuccessful=true, currentElapsedRealtimeMillis=PT4.535S (4535), latestTimeResult=TimeResult{unixEpochTime=2026-05-23T10:27:49.084Z, elapsedRealtime=PT4.535S, mUncertaintyMillis=43, mNtpServerSocketAddress=time.android.com/216.239.35.4:123}, mTryAgainCounter=0, refreshAttemptDelayMillis=64800000, nextRefreshElapsedRealtimeMillis=PT18H4.535S (64804535)
  PT2M53.263S / 2026-05-23T10:30:34.275826Z - refreshIfRequiredAndReschedule: network=101, reason=network available, initialTimeResult=TimeResult{unixEpochTime=2026-05-23T10:27:49.084Z, elapsedRealtime=PT4.535S, mUncertaintyMillis=43, mNtpServerSocketAddress=time.android.com/216.239.35.4:123}, shouldAttemptRefresh=false, refreshSuccessful=false, currentElapsedRealtimeMillis=PT2M53.263S (173263), latestTimeResult=TimeResult{unixEpochTime=2026-05-23T10:27:49.084Z, elapsedRealtime=PT4.535S, mUncertaintyMillis=43, mNtpServerSocketAddress=time.android.com/216.239.35.4:123}, mTryAgainCounter=0, refreshAttemptDelayMillis=64800000, nextRefreshElapsedRealtimeMillis=PT18H4.535S (64804535)
```
