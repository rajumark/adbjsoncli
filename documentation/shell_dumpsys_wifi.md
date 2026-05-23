# `adbjson shell dumpsys wifi`

## adbjson

**Command:**
```bash
adbjson shell dumpsys wifi
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
adb shell dumpsys wifi
```

**Output:**
```
Verbose logging is off
mVerboseLoggingLevel 0
Stay-awake conditions: 1
mInIdleMode false
mScanPending false
SupportedFeatures: [WIFI_FEATURE_INFRA, WIFI_FEATURE_P2P, WIFI_FEATURE_MOBILE_HOTSPOT, WIFI_FEATURE_WPA3_SAE, WIFI_FEATURE_WPA3_SUITE_B, WIFI_FEATURE_OWE, WIFI_FEATURE_DPP, WIFI_FEATURE_AP_RAND_MAC, WIFI_FEATURE_MBO, WIFI_FEATURE_FILS_SHA256, WIFI_FEATURE_FILS_SHA384, WIFI_FEATURE_SAE_PK, WIFI_FEATURE_DPP_ENROLLEE_RESPONDER, WIFI_FEATURE_PASSPOINT_TERMS_AND_CONDITIONS, WIFI_FEATURE_SAE_H2E, WIFI_FEATURE_WFD_R2, WIFI_FEATURE_DECORATED_IDENTITY, WIFI_FEATURE_TRUST_ON_FIRST_USE, WIFI_FEATURE_SET_TLS_MINIMUM_VERSION, WIFI_FEATURE_TLS_V1_3, WIFI_FEATURE_WEP, WIFI_FEATURE_WPA_PERSONAL]
SettingsStore:
WifiState 0
AirplaneModeOn false
ScanAlwaysAvailable false
WifiScoringState true
WifiPasspointState true
WifiMultiInternetMode 0
WifiStateApm false
WifiStateBt false
WifiStateUser 0
AirplaneModeEnhancementEnabled true
SatelliteModeOn false
Dump of WifiDeviceStateChangeManager
WifiDeviceStateChangeManager - Log Begin ----
mLastActiveScreenStateQuery: 
mLastScreenOn: true
WifiDeviceStateChangeManager - Log End ----
Dump of WifiActiveModeWarden
Current wifi mode: DisabledState
Wi-Fi is disabled
NumActiveModeManagers: 0
mIsMultiplePrimaryBugreportTaken: false
WifiController:
 total records=18
  rec[0]: time=05-23 15:57:44.304 processed=DefaultState org=DisabledState dest=<null> what=CMD_UPDATE_AP_CAPABILITY 1 0 num ClientModeManagers:0 num SoftApManagers:0 SupportedFeatures=42 MaximumSupportedClientNumber=16 SupportedChannelListIn24g[] SupportedChannelListIn5g[] SupportedChannelListIn6g[] SupportedChannelListIn60g[] mCountryCodeFromDrivernull
  rec[1]: time=05-23 15:57:45.120 processed=DefaultState org=DisabledState dest=<null> what=CMD_UPDATE_AP_CAPABILITY 1 0 num ClientModeManagers:0 num SoftApManagers:0 SupportedFeatures=42 MaximumSupportedClientNumber=16 SupportedChannelListIn24g[] SupportedChannelListIn5g[] SupportedChannelListIn6g[] SupportedChannelListIn60g[] mCountryCodeFromDrivernull
  rec[2]: time=05-23 15:57:48.830 processed=DefaultState org=DisabledState dest=<null> what=CMD_UPDATE_AP_CAPABILITY 1 0 num ClientModeManagers:0 num SoftApManagers:0 SupportedFeatures=42 MaximumSupportedClientNumber=16 SupportedChannelListIn24g[] SupportedChannelListIn5g[] SupportedChannelListIn6g[] SupportedChannelListIn60g[] mCountryCodeFromDrivernull
  rec[3]: time=05-23 15:59:55.354 processed=DisabledState org=DisabledState dest=EnabledState what=CMD_WIFI_TOGGLED 0 0 num ClientModeManagers:1 num SoftApManagers:0 WorkSource{2000 com.android.shell}
  rec[4]: time=05-23 15:59:55.459 processed=EnabledState org=EnabledState dest=<null> what=CMD_WIFI_TOGGLED 0 0 num ClientModeManagers:1 num SoftApManagers:0 WorkSource{2000 com.android.shell}
  rec[5]: time=05-23 15:59:55.571 processed=EnabledState org=EnabledState dest=DisabledState what=CMD_STA_STOPPED 0 0 num ClientModeManagers:0 num SoftApManagers:0
  rec[6]: time=05-23 15:59:55.571 processed=DefaultState org=DisabledState dest=<null> what=CMD_UPDATE_AP_CAPABILITY 1 0 num ClientModeManagers:0 num SoftApManagers:0 SupportedFeatures=42 MaximumSupportedClientNumber=16 SupportedChannelListIn24g[] SupportedChannelListIn5g[] SupportedChannelListIn6g[] SupportedChannelListIn60g[] mCountryCodeFromDrivernull
  rec[7]: time=05-23 15:59:55.571 processed=DefaultState org=DisabledState dest=<null> what=CMD_UPDATE_AP_CAPABILITY 2 0 num ClientModeManagers:0 num SoftApManagers:0 SupportedFeatures=42 MaximumSupportedClientNumber=16 SupportedChannelListIn24g[] SupportedChannelListIn5g[] SupportedChannelListIn6g[] SupportedChannelListIn60g[] mCountryCodeFromDrivernull
  rec[8]: time=05-23 16:00:27.833 processed=DisabledState org=DisabledState dest=EnabledState what=CMD_WIFI_TOGGLED 0 0 num ClientModeManagers:1 num SoftApManagers:0 WorkSource{2000 com.android.shell}
  rec[9]: time=05-23 16:00:27.904 processed=EnabledState org=EnabledState dest=<null> what=CMD_WIFI_TOGGLED 0 0 num ClientModeManagers:1 num SoftApManagers:0 WorkSource{2000 com.android.shell}
  rec[10]: time=05-23 16:00:28.051 processed=EnabledState org=EnabledState dest=DisabledState what=CMD_STA_STOPPED 0 0 num ClientModeManagers:0 num SoftApManagers:0
  rec[11]: time=05-23 16:00:28.051 processed=DefaultState org=DisabledState dest=<null> what=CMD_UPDATE_AP_CAPABILITY 1 0 num ClientModeManagers:0 num SoftApManagers:0 SupportedFeatures=42 MaximumSupportedClientNumber=16 SupportedChannelListIn24g[] SupportedChannelListIn5g[] SupportedChannelListIn6g[] SupportedChannelListIn60g[] mCountryCodeFromDrivernull
  rec[12]: time=05-23 16:00:28.051 processed=DefaultState org=DisabledState dest=<null> what=CMD_UPDATE_AP_CAPABILITY 2 0 num ClientModeManagers:0 num SoftApManagers:0 SupportedFeatures=42 MaximumSupportedClientNumber=16 SupportedChannelListIn24g[] SupportedChannelListIn5g[] SupportedChannelListIn6g[] SupportedChannelListIn60g[] mCountryCodeFromDrivernull
  rec[13]: time=05-23 16:00:33.987 processed=DisabledState org=DisabledState dest=EnabledState what=CMD_WIFI_TOGGLED 0 0 num ClientModeManagers:1 num SoftApManagers:0 WorkSource{2000 com.android.shell}
  rec[14]: time=05-23 16:00:34.089 processed=EnabledState org=EnabledState dest=<null> what=CMD_WIFI_TOGGLED 0 0 num ClientModeManagers:1 num SoftApManagers:0 WorkSource{2000 com.android.shell}
  rec[15]: time=05-23 16:00:34.299 processed=EnabledState org=EnabledState dest=DisabledState what=CMD_STA_STOPPED 0 0 num ClientModeManagers:0 num SoftApManagers:0
  rec[16]: time=05-23 16:00:34.299 processed=DefaultState org=DisabledState dest=<null> what=CMD_UPDATE_AP_CAPABILITY 1 0 num ClientModeManagers:0 num SoftApManagers:0 SupportedFeatures=42 MaximumSupportedClientNumber=16 SupportedChannelListIn24g[] SupportedChannelListIn5g[] SupportedChannelListIn6g[] SupportedChannelListIn60g[] mCountryCodeFromDrivernull
  rec[17]: time=05-23 16:00:34.299 processed=DefaultState org=DisabledState dest=<null> what=CMD_UPDATE_AP_CAPABILITY 2 0 num ClientModeManagers:0 num SoftApManagers:0 SupportedFeatures=42 MaximumSupportedClientNumber=16 SupportedChannelListIn24g[] SupportedChannelListIn5g[] SupportedChannelListIn6g[] SupportedChannelListIn60g[] mCountryCodeFromDrivernull
 curState=DisabledState
Dump of ActiveModeWarden.Graveyard
Stopped ClientModeManagers: 3 total
Dump of ClientModeManager id=134341
current StateMachine mode: StateMachine not active
mRole: null
mPreviousRole: ROLE_CLIENT_PRIMARY
mTargetRoleChangeInfo: null
mClientInterfaceName: null
mIfaceIsUp: false
mSecondaryInternet: false
mIsDbs: false
WifiClientModeManager:
 total records=2
 rec[0]: time=05-23 15:59:55.372 processed=IdleState org=IdleState dest=StartedState what=CMD_START 0 0 Role: ROLE_CLIENT_PRIMARY, RequestorWs: WorkSource{2000 com.android.shell}, ModeListener: com.android.server.wifi.ActiveModeWarden$ClientListener@7d83a50
 rec[1]: time=05-23 15:59:55.447 processed=StartedState org=StartedState dest=ConnectModeState what=CMD_SWITCH_TO_CONNECT_MODE 0 0 Role: ROLE_CLIENT_PRIMARY, RequestorWs: WorkSource{2000 com.android.shell}, ModeListener: com.android.server.wifi.ActiveModeWarden$ClientListener@7d83a50
curState=ConnectModeState

No active ClientModeImpl instance
Dump of ConcreteClientModeManager.Graveyard
Stopped ClientModeImpls: 1 total
Dump of ClientModeImpl id=134435
WifiClientModeImpl:
 total records=1
 rec[0]: time=05-23 15:59:55.457 processed=ConnectableState org=DisconnectedState dest=<null> what=CMD_IPCLIENT_CREATED screen=on 1 0
curState=DisconnectedState
SupplicantStateTracker:
 total records=0
 curState=<QUIT>
mAuthFailureInSupplicantBroadcast false
mAuthFailureReason 0

mIsRestrictedNetworkDebug false
mLinkProperties {LinkAddresses: [ ] DnsAddresses: [ ] Domains: null MTU: 0 Routes: [ ]}
mWifiInfo SSID: <unknown ssid>, BSSID: <none>, MAC: 02:15:b2:00:00:00, IP: null, Security type: -1, Supplicant state: DISCONNECTED, Wi-Fi standard: unknown, RSSI: -127, Link speed: -1Mbps, Tx Link speed: -1Mbps, Max Supported Tx Link speed: -1Mbps, Calculated Tx : 0Mbps, Rx Link speed: -1Mbps, Max Supported Rx Link speed: -1Mbps, Calculated Rx : 0Mbps, Frequency: -1MHz, Net ID: -1, Metered hint: false, score: 0, isUsable: true, CarrierMerged: false, SubscriptionId: -1, IsPrimary: 0, Trusted: false, Restricted: false, Ephemeral: false, OEM paid: false, OEM private: false, OSU AP: false, FQDN: <none>, Provider friendly name: <none>, Requesting package name: <none><none>MLO Information: , Is TID-To-Link negotiation supported by the AP: false, AP MLD Address: <none>, AP MLO Link Id: <none>, AP MLO Affiliated links: <none>, Vendor Data: <none>
mDhcpResultsParcelable baseConfiguration nullleaseDuration 0mtu 0serverAddress nullserverHostName nullvendorInfo null
mLastSignalLevel -1
mLastTxKbps -1
mLastRxKbps -1
mLastBssid null
mLastNetworkId -1
mLastSubId -1
mLastSimBasedConnectionCarrierName null
mSuspendOptimizationsEnabled true
mSuspendOptNeedsDisabled 0
mPowerSaveDisableRequests 0
WifiScoreReport:
time,session,netid,rssi,filtered_rssi,rssi_threshold,freq,txLinkSpeed,rxLinkSpeed,txTput,rxTput,calculatedTx,calculatedRx,bcnCnt,tx_good,tx_retry,tx_bad,rx_pps,nudrq,nuds,internalScorerType, internalScore, internalAdjustedScore, internalIsUsable, externalScore,{linkId,linkRssi,linkFreq,txLinkSpeed,rxLinkSpeed,linkBcnCnt,linkTxGood,linkTxRetry,linkTxBad,linkRxGood,linkMloState,linkUsageState}
externalScorerActive=false
mShouldReduceNetworkScore=false
QosPolicyRequestHandler:
mQosRequestDialogToken: 0\n\n... (truncated,     2146 lines total)
```
