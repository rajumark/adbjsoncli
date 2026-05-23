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
  "output": {
    "ssid": "\"AndroidWifi\""
  }
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
SupportedFeatures: [WIFI_FEATURE_INFRA, WIFI_FEATURE_P2P, WIFI_FEATURE_MOBILE_HOTSPOT, WIFI_FEATURE_WPA3_SAE, WIFI_FEATURE_WPA3_SUITE_B, WIFI_FEATURE_OWE, WIFI_FEATURE_DPP, WIFI_FEATURE_AP_RAND_MAC, WIFI_FEATURE_MBO, WIFI_FEATURE_FILS_SHA256, WIFI_FEATURE_FILS_SHA384, WIFI_FEATURE_SAE_PK, WIFI_FEATURE_DPP_ENROLLEE_RESPONDER, WIFI_FEATURE_PASSPOINT_TERMS_AND_CONDITIONS, WIFI_FEATURE_WFD_R2, WIFI_FEATURE_DECORATED_IDENTITY, WIFI_FEATURE_TRUST_ON_FIRST_USE, WIFI_FEATURE_SET_TLS_MINIMUM_VERSION, WIFI_FEATURE_TLS_V1_3, WIFI_FEATURE_WEP, WIFI_FEATURE_WPA_PERSONAL]
SettingsStore:
WifiState 1
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
Current wifi mode: EnabledState
Wi-Fi is enabled
NumActiveModeManagers: 1
mIsMultiplePrimaryBugreportTaken: false
WifiController:
 total records=4
  rec[0]: time=05-23 15:37:35.173 processed=DefaultState org=EnabledState dest=<null> what=CMD_UPDATE_AP_CAPABILITY 1 0 num ClientModeManagers:1 num SoftApManagers:0 SupportedFeatures=42 MaximumSupportedClientNumber=16 SupportedChannelListIn24g[] SupportedChannelListIn5g[] SupportedChannelListIn6g[] SupportedChannelListIn60g[] mCountryCodeFromDrivernull
  rec[1]: time=05-23 15:37:40.116 processed=DefaultState org=EnabledState dest=<null> what=CMD_UPDATE_AP_CAPABILITY 1 0 num ClientModeManagers:1 num SoftApManagers:0 SupportedFeatures=42 MaximumSupportedClientNumber=16 SupportedChannelListIn24g[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11] SupportedChannelListIn5g[] SupportedChannelListIn6g[] SupportedChannelListIn60g[] mCountryCodeFromDriverUS
  rec[2]: time=05-23 15:37:40.116 processed=DefaultState org=EnabledState dest=<null> what=CMD_UPDATE_AP_CAPABILITY 2 0 num ClientModeManagers:1 num SoftApManagers:0 SupportedFeatures=42 MaximumSupportedClientNumber=16 SupportedChannelListIn24g[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11] SupportedChannelListIn5g[] SupportedChannelListIn6g[] SupportedChannelListIn60g[] mCountryCodeFromDriverUS
  rec[3]: time=05-23 15:37:43.641 processed=DefaultState org=EnabledState dest=<null> what=CMD_UPDATE_AP_CAPABILITY 1 0 num ClientModeManagers:1 num SoftApManagers:0 SupportedFeatures=42 MaximumSupportedClientNumber=16 SupportedChannelListIn24g[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11] SupportedChannelListIn5g[] SupportedChannelListIn6g[] SupportedChannelListIn60g[] mCountryCodeFromDriverUS
 curState=EnabledState
Dump of ClientModeManager id=4074
current StateMachine mode: ConnectModeState
mRole: ROLE_CLIENT_PRIMARY
mPreviousRole: null
mTargetRoleChangeInfo: Role: ROLE_CLIENT_PRIMARY, RequestorWs: WorkSource{1000 com.android.settings}, ModeListener: com.android.server.wifi.ActiveModeWarden$ClientListener@40d2d30
mClientInterfaceName: wlan0
mIfaceIsUp: true
mSecondaryInternet: false
mIsDbs: false
WifiClientModeManager:
 total records=2
 rec[0]: time=05-23 15:37:32.519 processed=IdleState org=IdleState dest=StartedState what=CMD_START 0 0 Role: ROLE_CLIENT_PRIMARY, RequestorWs: WorkSource{1000 com.android.settings}, ModeListener: com.android.server.wifi.ActiveModeWarden$ClientListener@40d2d30
 rec[1]: time=05-23 15:37:32.585 processed=StartedState org=StartedState dest=ConnectModeState what=CMD_SWITCH_TO_CONNECT_MODE 0 0 Role: ROLE_CLIENT_PRIMARY, RequestorWs: WorkSource{1000 com.android.settings}, ModeListener: com.android.server.wifi.ActiveModeWarden$ClientListener@40d2d30
curState=ConnectModeState

Dump of ClientModeImpl id=4187
WifiClientModeImpl:
 total records=41
 rec[0]: time=05-23 15:37:33.333 processed=ConnectableState org=DisconnectedState dest=<null> what=CMD_IPCLIENT_CREATED screen=on 1 0
 rec[1]: time=05-23 15:37:33.346 processed=ConnectableState org=DisconnectedState dest=<null> what=CMD_ENABLE_RSSI_POLL screen=on 1 0
 rec[2]: time=05-23 15:37:33.346 processed=ConnectableState org=DisconnectedState dest=<null> what=CMD_SET_SUSPEND_OPT_ENABLED screen=on 0 0
 rec[3]: time=05-23 15:37:33.361 processed=ConnectableState org=DisconnectedState dest=<null> what=CMD_RESET_SIM_NETWORKS screen=on 2 0
 rec[4]: time=05-23 15:37:35.133 processed=ConnectableState org=DisconnectedState dest=<null> what=CMD_RESET_SIM_NETWORKS screen=on 1 0
 rec[5]: time=05-23 15:37:40.116 processed=ConnectableState org=DisconnectedState dest=<null> what=CMD_SCREEN_STATE_CHANGED screen=on 1 0
 rec[6]: time=05-23 15:37:40.116 processed=ConnectableState org=DisconnectedState dest=<null> what=CMD_ENABLE_RSSI_POLL screen=on 1 0
 rec[7]: time=05-23 15:37:40.118 processed=ConnectableState org=DisconnectedState dest=<null> what=CMD_SET_SUSPEND_OPT_ENABLED screen=on 0 0
 rec[8]: time=05-23 15:37:45.102 processed=ConnectableState org=DisconnectedState dest=L2ConnectingState what=CMD_START_CONNECT screen=on 0 1010 targetConfigKey="AndroidWifi"NONE BSSID=null targetBssid=00:13:10:85:fe:01 roam=false
 rec[9]: time=05-23 15:37:45.394 processed=ConnectingOrConnectedState org=L2ConnectingState dest=<null> what=SUPPLICANT_STATE_CHANGE_EVENT screen=on 0 0 ssid: "AndroidWifi" bssid: 00:13:10:85:fe:01 nid: 0 frequencyMhz: 0 state: AUTHENTICATING
 rec[10]: time=05-23 15:37:45.754 processed=ConnectingOrConnectedState org=L2ConnectingState dest=<null> what=SUPPLICANT_STATE_CHANGE_EVENT screen=on 0 0 ssid: "AndroidWifi" bssid: 00:13:10:85:fe:01 nid: 0 frequencyMhz: 0 state: ASSOCIATING
 rec[11]: time=05-23 15:37:46.357 processed=ConnectingOrConnectedState org=L2ConnectingState dest=<null> what=SUPPLICANT_STATE_CHANGE_EVENT screen=on 0 0 ssid: "AndroidWifi" bssid: 00:13:10:85:fe:01 nid: 0 frequencyMhz: 2447 state: ASSOCIATED
 rec[12]: time=05-23 15:37:46.385 processed=ConnectableState org=L2ConnectingState dest=<null> what=ASSOCIATED_BSSID_EVENT screen=on 0 0 BSSID=00:13:10:85:fe:01 Target Bssid=00:13:10:85:fe:01 Last Bssid=00:13:10:85:fe:01 roam=false
 rec[13]: time=05-23 15:37:46.524 processed=ConnectingOrConnectedState org=L2ConnectingState dest=L3ProvisioningState what=NETWORK_CONNECTION_EVENT screen=on 0 false 00:13:10:85:fe:01 nid=0 "AndroidWifi"NONE last=
 rec[14]: time=05-23 15:37:46.977 processed=ConnectingOrConnectedState org=L3ProvisioningState dest=<null> what=SUPPLICANT_STATE_CHANGE_EVENT screen=on 0 0 ssid: "AndroidWifi" bssid: 00:13:10:85:fe:01 nid: 0 frequencyMhz: 2447 state: COMPLETED
 rec[15]: time=05-23 15:37:47.404 processed=ConnectableState org=L3ProvisioningState dest=<null> what=CMD_CONFIG_ND_OFFLOAD screen=on 1 1
 rec[16]: time=05-23 15:37:47.404 processed=ConnectableState org=L3ProvisioningState dest=<null> what=CMD_INSTALL_PACKET_FILTER screen=on len=4096
 rec[17]: time=05-23 15:37:47.404 processed=ConnectableState org=L3ProvisioningState dest=<null> what=CMD_INSTALL_PACKET_FILTER screen=on len=228
 rec[18]: time=05-23 15:37:47.404 processed=ConnectableState org=L3ProvisioningState dest=<null> what=CMD_CONFIG_ND_OFFLOAD screen=on 1 0
 rec[19]: time=05-23 15:37:47.617 processed=ConnectableState org=L3ProvisioningState dest=<null> what=CMD_INSTALL_PACKET_FILTER screen=on len=234
 rec[20]: time=05-23 15:37:47.617 processed=ConnectableState org=L3ProvisioningState dest=<null> what=CMD_INSTALL_PACKET_FILTER screen=on len=314
 rec[21]: time=05-23 15:37:47.718 processed=ConnectableState org=L3ProvisioningState dest=<null> what=CMD_UPDATE_LINKPROPERTIES screen=on 1 0 
 rec[22]: time=05-23 15:37:47.722 processed=ConnectableState org=L3ProvisioningState dest=<null> what=CMD_SET_MAX_DTIM_MULTIPLIER screen=on maximum multiplier=1
 rec[23]: time=05-23 15:37:47.753 processed=L2ConnectedState org=L3ProvisioningState dest=<null> what=CMD_PRE_DHCP_ACTION screen=on 1 0 txpkts=5,0,0
 rec[24]: time=05-23 15:37:47.753 processed=L2ConnectedState org=L3ProvisioningState dest=<null> what=CMD_PRE_DHCP_ACTION_COMPLETE screen=on 0 0
 rec[25]: time=05-23 15:37:47.882 processed=L2ConnectedState org=L3ProvisioningState dest=<null> what=CMD_POST_DHCP_ACTION screen=on 
 rec[26]: time=05-23 15:37:47.988 processed=L2ConnectedState org=L3ProvisioningState dest=<null> what=CMD_IPV4_PROVISIONING_SUCCESS screen=on DhcpResultsParcelable{baseConfiguration: IP address 10.0.2.16/24 Gateway 10.0.2.2  DNS servers: [ 10.0.2.3 ] Domains , leaseDuration: 86400, mtu: 0, serverAddress: 10.0.2.2, vendorInfo: null, serverHostName: , captivePortalApiUrl: null}
 rec[27]: time=05-23 15:37:47.988 processed=ConnectableState org=L3ProvisioningState dest=<null> what=CMD_UPDATE_LINKPROPERTIES screen=on 1 0 v4r
 rec[28]: time=05-23 15:37:47.991 processed=ConnectableState org=L3ProvisioningState dest=<null> what=CMD_INSTALL_PACKET_FILTER screen=on len=328
 rec[29]: time=05-23 15:37:48.003 processed=ConnectableState org=L3ProvisioningState dest=<null> what=CMD_INSTALL_PACKET_FILTER screen=on len=651
 rec[30]: time=05-23 15:37:48.006 processed=ConnectableState org=L3ProvisioningState dest=<null> what=CMD_UPDATE_LINKPROPERTIES screen=on 1 0 v4 v4r v4dns
 rec[31]: time=05-23 15:37:48.006 processed=L2ConnectedState org=L3ProvisioningState dest=L3ConnectedState what=CMD_IP_CONFIGURATION_SUCCESSFUL screen=on 1 0
 rec[32]: time=05-23 15:37:48.686 processed=ConnectableState org=L3ConnectedState dest=<null> what=CMD_INSTALL_PACKET_FILTER screen=on len=1080
 rec[33]: time=05-23 15:37:48.686 processed=ConnectableState org=L3ConnectedState dest=<null> what=CMD_INSTALL_PACKET_FILTER screen=on len=1334
 rec[34]: time=05-23 15:37:48.686 processed=ConnectableState org=L3ConnectedState dest=<null> what=CMD_INSTALL_PACKET_FILTER screen=on len=1464
 rec[35]: time=05-23 15:37:48.686 processed=ConnectableState org=L3ConnectedState dest=<null> what=CMD_INSTALL_PACKET_FILTER screen=on len=1516
 rec[36]: time=05-23 15:37:48.776 processed=ConnectableState org=L3ConnectedState dest=<null> what=CMD_UPDATE_LINKPROPERTIES screen=on 1 0 v4 v4r v4dns v6r v6dns
 rec[37]: time=05-23 15:37:49.066 processed=ConnectableState org=L3ConnectedState dest=<null> what=CMD_INSTALL_PACKET_FILTER screen=on len=1513
 rec[38]: time=05-23 15:37:55.284 processed=L3ConnectedState org=L3ConnectedState dest=<null> what=CMD_NETWORK_STATUS screen=on 1 0
 rec[39]: time=05-23 15:38:04.421 processed=L2ConnectedState org=L3ConnectedState dest=<null> what=CMD_ONESHOT_RSSI_POLL screen=on 0 0 "AndroidWifi" 00:13:10:85:fe:01 rssi=-50 f=2447 sc=100 link=2 tx=63.1, 0.0, 0.0 rx=103.2 bcn=0 [on:0 tx:0 rx:0 period:1414423877] from screen [on:0 period:1414423877] score=100
 rec[40]: time=05-23 15:40:23.741 processed=L2ConnectedState org=L3ConnectedState dest=<null> what=CMD_ONESHOT_RSSI_POLL screen=on 0 0 "AndroidWifi" 00:13:10:85:fe:01 rssi=-50 f=2447 sc=100 link=1 tx=0.0, 0.0, 0.0 rx=0.0 bcn=0 [on:0 tx:0 rx:0 period:139320] from screen [on:0 period:1414563197] score=100
curState=L3ConnectedState
SupplicantStateTracker:
 total records=2
  rec[0]: time=05-23 15:37:45.379 processed=DefaultState org=UninitializedState dest=HandshakeState what=147462(0x24006)
  rec[1]: time=05-23 15:37:46.935 processed=DefaultState org=HandshakeState dest=CompletedState what=147462(0x24006)
 curState=CompletedState
mAuthFailureInSupplicantBroadcast false
mAuthFailureReason 0

mIsRestrictedNetworkDebug false
mLinkProperties {InterfaceName: wlan0 LinkAddresses: [ fe80::eecc:225f:e844:5677/64,10.0.2.16/24,fec0::e3c9:906d:cb80:5fe/64,fec0::a37b:76f7:57b1:2420/64 ] DnsAddresses: [ /fec0::3,/10.0.2.3 ] Domains: null MTU: 0 ServerAddress: /10.0.2.2 TcpBufferSizes: 524288,1048576,2097152,262144,524288,1048576 Routes: [ fe80::/64 -> :: wlan0 mtu 0,::/0 -> fe80::2 wlan0 mtu 0,fec0::/64 -> :: wlan0 mtu 0,10.0.2.0/24 -> 0.0.0.0 wlan0 mtu 0,0.0.0.0/0 -> 10.0.2.2 wlan0 mtu 0 ]}
mWifiInfo SSID: "AndroidWifi", BSSID: 00:13:10:85:fe:01, MAC: 02:15:b2:00:00:00, IP: /10.0.2.16, Security type: 0, Supplicant state: COMPLETED, Wi-Fi standard: legacy, RSSI: -50, Link speed: 1Mbps, Tx Link speed: 1Mbps, Max Supported Tx Link speed: 11Mbps, Calculated Tx : 0Mbps, Rx Link speed: 2Mbps, Max Supported Rx Link speed: 11Mbps, Calculated Rx : 0Mbps, Frequency: 2447MHz, Net ID: 0, Metered hint: false, score: 100, isUsable: true, CarrierMerged: false, SubscriptionId: -1, IsPrimary: 1, Trusted: true, Restricted: false, Ephemeral: false, OEM paid: false, OEM private: false, OSU AP: false, FQDN: <none>, Provider friendly name: <none>, Requesting package name: <none>"AndroidWifi"openMLO Information: , Is TID-To-Link negotiation supported by the AP: false, AP MLD Address: <none>, AP MLO Link Id: <none>, AP MLO Affiliated links: <none>, Vendor Data: <none>
mDhcpResultsParcelable baseConfiguration IP address 10.0.2.16/24 Gateway 10.0.2.2  DNS servers: [ 10.0.2.3 ] Domains leaseDuration 86400mtu 0serverAddress 10.0.2.2serverHostName vendorInfo null
mLastSignalLevel 4
mLastTxKbps 12000
mLastRxKbps 30000
mLastBssid 00:13:10:85:fe:01
mLastNetworkId 0
mLastSubId -1
mLastSimBasedConnectionCarrierName null
mSuspendOptimizationsEnabled true
mSuspendOptNeedsDisabled 4
mPowerSaveDisableRequests 0
IpClient logs have moved to dumpsys network_stack
WifiScoreReport:
time,session,netid,rssi,filtered_rssi,rssi_threshold,freq,txLinkSpeed,rxLinkSpeed,txTput,rxTput,calculatedTx,calculatedRx,bcnCnt,tx_good,tx_retry,tx_bad,rx_pps,nudrq,nuds,internalScorerType, internalScore, internalAdjustedScore, internalIsUsable, externalScore,{linkId,linkRssi,linkFreq,txLinkSpeed,rxLinkSpeed,linkBcnCnt,linkTxGood,linkTxRetry,linkTxBad,linkRxGood,linkMloState,linkUsageState}
5-23 15:37:47.156,0,101,-50,0.0,-80.0,2447,2,-1,12,30,0Mbps,0Mbps,0,0.00,0.00,0.00,0.00,1,1,ML,45,100,true,-1
5-23 15:37:50.457,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,10.79,0.00,0.00,7.10,2,2,ML,31,100,true,-1
5-23 15:37:53.637,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,7.14,0.00,0.00,4.25,3,2,ML,31,100,true,-1
5-23 15:37:56.679,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,19.76,0.00,0.00,20.39,4,3,ML,31,100,true,-1
5-23 15:37:59.699,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,28.58,0.00,0.00,32.37,5,3,ML,33,100,true,-1
5-23 15:38:2.710,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,63.17,0.00,0.00,103.29,6,4,ML,30,100,true,-1
5-23 15:38:5.725,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,77.62,0.00,0.00,167.63,7,4,ML,30,100,true,-1
5-23 15:38:8.726,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,30.35,0.00,0.00,62.93,8,5,ML,30,100,true,-1
5-23 15:38:11.732,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,13.25,0.00,0.00,25.22,9,5,ML,28,100,true,-1
5-23 15:38:14.736,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,7.18,0.00,0.00,10.52,10,6,ML,28,100,true,-1
5-23 15:38:17.741,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,3.90,0.00,0.00,4.91,11,6,ML,28,100,true,-1
5-23 15:38:20.745,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,1.43,0.00,0.00,1.80,12,7,ML,31,100,true,-1
5-23 15:38:23.749,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,2.84,0.00,0.00,1.92,13,7,ML,26,100,true,-1
5-23 15:38:26.751,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,1.46,0.00,0.00,1.12,14,8,ML,26,100,true,-1
5-23 15:38:29.755,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,0.96,0.00,0.00,0.83,15,8,ML,26,100,true,-1
5-23 15:38:32.763,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,0.35,0.00,0.00,0.30,16,9,ML,30,100,true,-1
5-23 15:38:35.768,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.55,0.00,0.00,0.53,17,9,ML,26,100,true,-1
5-23 15:38:38.772,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,0.20,0.00,0.00,0.19,17,9,ML,82,100,true,-1
5-23 15:38:41.778,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,0.07,0.00,0.00,0.07,17,9,ML,82,100,true,-1
5-23 15:38:44.782,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.02,0.00,0.00,0.02,17,9,ML,82,100,true,-1
5-23 15:38:47.787,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.01,0.00,0.00,0.00,17,9,ML,82,100,true,-1
5-23 15:38:50.791,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.00,0.00,0.00,0.00,17,9,ML,77,100,true,-1
5-23 15:38:53.796,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,1.05,0.00,0.00,0.84,18,10,ML,26,100,true,-1
5-23 15:38:56.800,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.80,0.00,0.00,0.73,19,10,ML,26,100,true,-1
5-23 15:38:59.806,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.29,0.00,0.00,0.26,20,11,ML,30,100,true,-1
5-23 15:39:2.810,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,0.55,0.00,0.00,0.54,21,11,ML,26,100,true,-1
5-23 15:39:5.817,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.62,0.00,0.00,0.62,22,12,ML,27,100,true,-1
5-23 15:39:8.820,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,0.65,0.00,0.00,0.64,22,12,ML,82,100,true,-1
5-23 15:39:11.823,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,0.23,0.00,0.00,0.23,22,12,ML,82,100,true,-1
5-23 15:39:14.828,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,0.08,0.00,0.00,0.08,22,12,ML,82,100,true,-1
5-23 15:39:17.830,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,3.40,0.00,0.00,3.82,23,13,ML,26,100,true,-1
5-23 15:39:20.834,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,1.67,0.00,0.00,1.82,24,13,ML,26,100,true,-1
5-23 15:39:23.840,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,0.61,0.00,0.00,0.67,25,14,ML,30,100,true,-1
5-23 15:39:26.843,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.64,0.00,0.00,0.66,26,14,ML,26,100,true,-1
5-23 15:39:29.848,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.23,0.00,0.00,0.24,27,15,ML,30,100,true,-1
5-23 15:39:32.853,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,0.50,0.00,0.00,0.51,27,15,ML,82,100,true,-1
5-23 15:39:35.858,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,0.18,0.00,0.00,0.18,27,15,ML,82,100,true,-1
5-23 15:39:38.866,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,0.06,0.00,0.00,0.06,27,15,ML,82,100,true,-1
5-23 15:39:41.876,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.02,0.00,0.00,0.02,27,15,ML,82,100,true,-1
5-23 15:39:44.887,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.00,0.00,0.00,0.00,27,15,ML,82,100,true,-1
5-23 15:39:47.900,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.00,0.00,0.00,0.00,27,15,ML,77,100,true,-1
5-23 15:39:50.911,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.00,0.00,0.00,0.00,27,15,ML,77,100,true,-1
5-23 15:39:53.919,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.00,0.00,0.00,0.00,27,15,ML,77,100,true,-1
5-23 15:39:56.922,1,101,-50,0.0,-80.0,2447,11,2,12,30,0Mbps,0Mbps,0,0.21,0.00,0.00,0.21,27,15,ML,81,100,true,-1
5-23 15:39:59.928,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.91,0.00,0.00,0.70,28,16,ML,27,100,true,-1
5-23 15:40:2.935,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.75,0.00,0.00,0.68,29,16,ML,27,100,true,-1
5-23 15:40:5.942,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,0.27,0.00,0.00,0.25,30,17,ML,32,100,true,-1
5-23 15:40:8.951,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.52,0.00,0.00,0.51,31,17,ML,27,100,true,-1
5-23 15:40:11.962,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.19,0.00,0.00,0.18,32,18,ML,32,100,true,-1
5-23 15:40:14.971,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,0.49,0.00,0.00,0.48,32,18,ML,82,100,true,-1
5-23 15:40:17.976,1,101,-50,0.0,-80.0,2447,2,2,12,30,0Mbps,0Mbps,0,0.18,0.00,0.00,0.17,32,18,ML,82,100,true,-1
5-23 15:40:20.982,1,101,-50,0.0,-80.0,2447,1,2,12,30,0Mbps,0Mbps,0,0.06,0.00,0.00,0.06,32,18,ML,82,100,true,-1
externalScorerActive=false
mShouldReduceNetworkScore=false
QosPolicyRequestHandler:
mQosRequestDialogToken: 0
mNumQosPoliciesInRequest: 0
mQosResourcesAvailable: false
mQosPolicyStatusList size: 0
mQosPolicyRequestQueue size: 0

Dump of ConcreteClientModeManager.Graveyard
Stopped ClientModeImpls: 0 total


Dump of ActiveModeWarden.Graveyard
Stopped ClientModeManagers: 0 total
Stopped SoftApManagers: 0 total

STA + STA Concurrency Supported: false
STA + AP Concurrency Supported: false
Dump of HalDeviceManager:
  mManagerStatusListeners: [com.android.server.wifi.HalDeviceManager$ManagerStatusListenerProxy@a3f7bfa]
  mInterfaceInfoCache: {Pair{wlan0 0}={name=wlan0, type=0, destroyedListeners.size()=1, RequestorWs=WorkSource{1000 com.android.settings}, creationTime=4100}}
  mDebugChipsInfo: [{chipId=0, availableModes=[{id=0, availableCombinations=[{limits=[{maxIfaces=1, types=[0]}, {maxIfaces=1, types=[3]}]}]}, {id=1, availableCombinations=[{limits=[{maxIfaces=1, types=[1]}]}]}], currentModeIdValid=true, currentModeId=0, chipCapabilities={}, radioCombinations=null, bandCombinations=null, ifaces[1].length=0, ifaces[0].length=1, ifaces[2].length=0, ifaces[3].length=0}]
Wifi handler thread overruns
2026-05-23T15:37:33.303906 - ClientModeImpl.ClientModeImpl$ConnectableState.Enter was running for 669 ms
2026-05-23T15:37:40.097902 - WifiService#handleBootCompleted was running for 227
2026-05-23T15:37:42.372734 - WifiService#handleUserUnlock was running for 484
2026-05-23T15:37:43.519863 - TelephonyRegistryManager$1#onSubscriptionsChanged was running for 265
2026-05-23T15:37:43.896941 - WifiCarrierInfoManager$2#onReceive was running for 142
2026-05-23T15:37:44.597424 - WifiScannerInternal#onResults was running for 111
2026-05-23T15:37:44.999622 - WifiScannerInternal#onResults was running for 397
2026-05-23T15:37:45.102762 - ClientModeImpl.ClientModeImpl$ConnectableState.CMD_START_CONNECT was running for 102 ms
2026-05-23T15:37:45.379442 - ClientModeImpl.ClientModeImpl$L2ConnectingState.Enter was running for 269 ms
2026-05-23T15:37:46.355004 - ClientModeImpl.ClientModeImpl$ConnectingOrConnectedState.SUPPLICANT_STATE_CHANGE_EVENT was running for 562 ms
2026-05-23T15:37:46.524515 - ClientModeImpl.ClientModeImpl$ConnectingOrConnectedState.NETWORK_CONNECTION_EVENT was running for 139 ms
2026-05-23T15:37:46.764799 - ClientModeImpl.ClientModeImpl$L2ConnectedState.Enter was running for 240 ms
2026-05-23T15:37:46.935425 - ClientModeImpl.ClientModeImpl$L3ProvisioningState.Enter was running for 171 ms
2026-05-23T15:37:47.381875 - ClientModeImpl.ClientModeImpl$L2ConnectedState.CMD_RSSI_POLL was running for 364 ms
2026-05-23T15:37:48.684228 - ClientModeImpl.ClientModeImpl$L3ConnectedState.Enter was running for 678 ms
2026-05-23T15:37:50.611648 - ClientModeImpl.ClientModeImpl$L2ConnectedState.CMD_RSSI_POLL was running for 228 ms
2026-05-23T15:38:04.418395 - WifiService#dump was running for 260
2026-05-23T15:40:23.741760 - WifiService#dump was running for 233
Dump of MakeBeforeBreakManager
mMakeBeforeBreakInfo=null
mInternalState MBB_STATE_NONE

dump of InterfaceConflictManager:
  mUserApprovalNeeded=false
  mUserApprovalNeededOverride=false
  mUserApprovalNeededOverrideValue=false
  mUserApprovalPending=false
  mUserApprovalPendingTag=null
  mUserJustApproved=false
  mUserApprovalNotRequireForDisconnectedP2p=false

mTxPkts 844
mRxPkts 1360
mLastActivity 0
mRegisteredCallbacks 1

Locks held:
Dump of WifiLockManager
Locks acquired: 0 full high perf, 0 full low latency
Locks released: 0 full high perf, 0 full low latency
Connection state: STA=true, P2P=false, Aware=false
Screen state: true
Current operation mode: 0

Locks held:

Low-latency uid watchlist:

mMulticastEnabled 3
mMulticastDisabled 2
Active lock owners: {1000=1}
Inactive lock owners: {}
Multicast Locks held:
    Multicaster{AdbMulticastLock-comm uid=1000}

WifiScoreCard:
GssGGvAECIDKtZYHIn0IAxCPExomCAwRAAAAAADAgsAZAAAAAABM3UAhAAAAAAAAScApAAAAAAAA
ScAiJggKEQAAAAAAADNAGQAAAAAAgEJAIQAAAAAAAPA/KQAAAAAAAABAKiYIDBEAAAAAgFfOQBkA
AACIoB2AQSEAAAAAAEBpQCkAAAAAALatQCKDAQgEEI8TGiwIDBEAAAAAAMCCwBkAAAAAAEzdQCEA
AAAAAABJwCkAAAAAAABJwEIECGMQDCImCAoRAAAAAAAAM0AZAAAAAACAQkAhAAAAAAAA8D8pAAAA
AAAAAEAqJggMEQAAAAAATtpAGQAAAFgCnZdBIQAAAAAAeIFAKQAAAAAAu7pAIn0IDRCPExomCAwR
AAAAAADAgsAZAAAAAABM3UAhAAAAAAAAScApAAAAAAAAScAiJggLEQAAAAAAADJAGQAAAAAAAEBA
IQAAAAAAAPA/KQAAAAAAAABAKiYIDBEAAAAAUI73QBkAAEDWoQnQQSEAAAAAAJCcQCkAAAAAgFnR
QCJbCAEQjxMaLAhREQAAAAAApK/AGQAAAAAguAhBIQAAAAAAAEnAKQAAAAAAAEnAQgQIYxBRIiYI
TxEAAAAAAIBfQBkAAAAAAGBzQCEAAAAAAADwPykAAAAAAAAmQDAAOoQBCkAKHgoECAAQAAoECAAQ
AAoECAAQAAoECAAQAAoECAAQABIeCgQIABAACgQIABAACgQIABAACgQIABAACgQIABAAEkAKHgoE
CAAQAAoECAAQAAoECAAQAAoECAAQAAoECAAQABIeCgQIABAACgQIABAACgQIABAACgQIABAACgQI
ABAAIAAoZTLRAQis7ctDEhQIDBAAGAAgACgAMAA4AEAASABQABoUCAAQABgAIAAoADAAOABAAEgA
UAAiFAgAEAAYACAAKAAwADgAQABIAFAAKI8TMoQBCkAKHgoECAAQAAoECAAQAAoECAAQAAoECAAQ
AAoECAAQABIeCgQIABAACgQIABAACgQIABAACgQIABAACgQIABAAEkAKHgoECAAQAAoECAAQAAoE
CAAQAAoECAAQAAoECAAQABIeCgQIABAACgQIABAACgQIABAACgQIABAACgQIABAA

WifiMetrics:
mConnectionEvents:
startTime=5-23 15:37:45.28, SSID="AndroidWifi", BSSID=00:13:10:85:fe:01, durationMillis=2983, roamType=ROAM_UNRELATED, connectionResult=1, level2FailureCode=NONE, connectivityLevelFailureCode=NONE, signalStrength=-50, wifiState=WIFI_DISCONNECTED, screenOn=true, mRouterFingerprint=mConnectionEvent.roamType=0, mChannelInfo=2447, mDtim=0, mAuthentication=1, mHidden=false, mRouterTechnology=4, mSupportsIpv6=false, mEapMethod=0, mAuthPhase2Method=0, mOcspType=0, mPmkCache=false, mMaxSupportedTxLinkSpeedMbps=11, mMaxSupportedRxLinkSpeedMbps=11, mIsFrameworkInitiatedRoaming=false, mIsIncorrectlyConfiguredAsHidden=false, mWifiStandard=4, mIs11bSupported=false, mIsMboSupported=false, mIsOceSupported=false, mIsFilsSupported=false, mIsIndividualTwtSupported=false, mIsBroadcastTwtSupported=false, mIsRestrictedTwtSupported=false, mIsTwtRequired=false, mIs11mcSupported=false, mIs11azSupported=false, mApType6Ghz=AP_TYPE_6GHZ_UNKNOWN, mIsEcpsPriorityAccessSupported=false, mHsRelease=null, mChannelWidth0falsefalsefalsefalse, useRandomizedMac=true, useAggressiveMac=false, connectionNominator=NOMINATOR_SAVED, networkSelectorExperimentId=42330058, numBssidInBlocklist=0, level2FailureReason=FAILURE_REASON_UNKNOWN, networkType=TYPE_OPEN, networkCreator=CREATOR_USER, numConsecutiveConnectionFailure=0, isOsuProvisioned=false interfaceName=wlan0 interfaceRole=ROLE_CLIENT_PRIMARY, isFirstConnectionAfterBoot=true, isCarrierWifi=false, isOobPseudonymEnabled=false, uid=1010
mWifiLogProto.numSavedNetworks=1
mWifiLogProto.numSavedNetworksWithMacRandomization=0
mWifiLogProto.numOpenNetworks=1
mWifiLogProto.numLegacyPersonalNetworks=0
mWifiLogProto.numLegacyEnterpriseNetworks=0
mWifiLogProto.numEnhancedOpenNetworks=0
mWifiLogProto.numWpa3PersonalNetworks=0
mWifiLogProto.numWpa3EnterpriseNetworks=0
mWifiLogProto.numWapiPersonalNetworks=0
mWifiLogProto.numWapiEnterpriseNetworks=0
mWifiLogProto.numHiddenNetworks=0
mWifiLogProto.numPasspointNetworks=0
mWifiLogProto.isLocationEnabled=false
mWifiLogProto.isScanningAlwaysEnabled=false
mWifiLogProto.isVerboseLoggingEnabled=false
mWifiLogProto.isEnhancedMacRandomizationForceEnabled=false
mWifiLogProto.isWifiWakeEnabled=true
mWifiLogProto.numNetworksAddedByUser=0
mWifiLogProto.numNetworksAddedByApps=1
mWifiLogProto.numNonEmptyScanResults=5
mWifiLogProto.numEmptyScanResults=0
mWifiLogProto.numConnecitvityOneshotScans=5
mWifiLogProto.numOneshotScans=5
mWifiLogProto.numOneshotHasDfsChannelScans=1
mWifiLogProto.numBackgroundScans=0
mWifiLogProto.numExternalAppOneshotScanRequests=0
mWifiLogProto.numExternalForegroundAppOneshotScanRequestsThrottled=0
mWifiLogProto.numExternalBackgroundAppOneshotScanRequestsThrottled=0
mWifiLogProto.meteredNetworkStatsSaved=
num_metered: 0
num_override_metered: 0
num_override_unmetered: 0
num_unmetered: 1

mWifiLogProto.meteredNetworkStatsSuggestion=
num_metered: 0
num_override_metered: 0
num_override_unmetered: 0
num_unmetered: 0

mScanReturnEntries:
  SCAN_UNKNOWN: 0
  SCAN_SUCCESS: 5
  SCAN_FAILURE_INTERRUPTED: 0
  SCAN_FAILURE_INVALID_CONFIGURATION: 0
  FAILURE_WIFI_DISABLED: 0
mSystemStateEntries: <state><screenOn> : <scansInitiated>
  WIFI_UNKNOWN       ON: 0
  WIFI_DISABLED      ON: 0
  WIFI_DISCONNECTED  ON: 1
  WIFI_ASSOCIATED    ON: 4
  WIFI_UNKNOWN      OFF: 0
  WIFI_DISABLED     OFF: 0
  WIFI_DISCONNECTED OFF: 0
  WIFI_ASSOCIATED   OFF: 0
mWifiLogProto.numConnectivityWatchdogPnoGood=0
mWifiLogProto.numConnectivityWatchdogPnoBad=0
mWifiLogProto.numConnectivityWatchdogBackgroundGood=0
mWifiLogProto.numConnectivityWatchdogBackgroundBad=0
mWifiLogProto.numLastResortWatchdogTriggers=0
mWifiLogProto.numLastResortWatchdogBadAssociationNetworksTotal=0
mWifiLogProto.numLastResortWatchdogBadAuthenticationNetworksTotal=0
mWifiLogProto.numLastResortWatchdogBadDhcpNetworksTotal=0
mWifiLogProto.numLastResortWatchdogBadOtherNetworksTotal=0
mWifiLogProto.numLastResortWatchdogAvailableNetworksTotal=0
mWifiLogProto.numLastResortWatchdogTriggersWithBadAssociation=0
mWifiLogProto.numLastResortWatchdogTriggersWithBadAuthentication=0
mWifiLogProto.numLastResortWatchdogTriggersWithBadDhcp=0
mWifiLogProto.numLastResortWatchdogTriggersWithBadOther=0
mWifiLogProto.numLastResortWatchdogSuccesses=0
mWifiLogProto.watchdogTotalConnectionFailureCountAfterTrigger=0
mWifiLogProto.watchdogTriggerToConnectionSuccessDurationMs=-1
mWifiLogProto.recordDurationSec=172
mWifiLogProto.rssiPollCount: {"2447":[{"-50":52}]}
mWifiLogProto.rssiPollDeltaCount: Printing counts for [-127, 127]
  0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 
mWifiLogProto.linkSpeedCounts: 
1:{29, 1450, 72500} 2:{22, 1100, 55000} 11:{1, 50, 2500} 
mWifiLogProto.alertReasonCounts=()
mWifiLogProto.numTotalScanResults=5
mWifiLogProto.numOpenNetworkScanResults=5
mWifiLogProto.numLegacyPersonalNetworkScanResults=0
mWifiLogProto.numLegacyEnterpriseNetworkScanResults=0
mWifiLogProto.numEnhancedOpenNetworkScanResults=0
mWifiLogProto.numWpa3PersonalNetworkScanResults=0
mWifiLogProto.numWpa3EnterpriseNetworkScanResults=0
mWifiLogProto.numWapiPersonalNetworkScanResults=0
mWifiLogProto.numWapiEnterpriseNetworkScanResults=0
mWifiLogProto.numHiddenNetworkScanResults=0
mWifiLogProto.numHotspot2R1NetworkScanResults=0
mWifiLogProto.numHotspot2R2NetworkScanResults=0
mWifiLogProto.numHotspot2R3NetworkScanResults=0
mWifiLogProto.numMboSupportedNetworkScanResults=0
mWifiLogProto.numMboCellularDataAwareNetworkScanResults=0
mWifiLogProto.numOceSupportedNetworkScanResults=0
mWifiLogProto.numFilsSupportedNetworkScanResults=0
mWifiLogProto.num11AxNetworkScanResults=0
mWifiLogProto.num6GNetworkScanResults0
mWifiLogProto.num6GPscNetworkScanResults0
mWifiLogProto.numBssidFilteredDueToMboAssocDisallowInd=0
mWifiLogProto.numConnectToNetworkSupportingMbo=0
mWifiLogProto.numConnectToNetworkSupportingOce=0
mWifiLogProto.numSteeringRequest=0
mWifiLogProto.numForceScanDueToSteeringRequest=0
mWifiLogProto.numMboCellularSwitchRequest=0
mWifiLogProto.numSteeringRequestIncludingMboAssocRetryDelay=0
mWifiLogProto.numConnectRequestWithFilsAkm=0
mWifiLogProto.numL2ConnectionThroughFilsAuthentication=0
mWifiLogProto.recentFailureAssociationStatus={}
mWifiLogProto.numScans=5
mWifiLogProto.WifiScoreCount: [0, 60]
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 
mWifiLogProto.WifiUsabilityScoreCount: [0, 100]
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 
mWifiLogProto.SoftApManagerReturnCodeCounts:
  SUCCESS: 0
  FAILED_GENERAL_ERROR: 0
  FAILED_NO_CHANNEL: 0
  FAILED_UNSUPPORTED_CONFIGURATION: 0

mWifiLogProto.numHalCrashes=0
mWifiLogProto.numWificondCrashes=0
mWifiLogProto.numSupplicantCrashes=0
mWifiLogProto.numHostapdCrashes=0
mWifiLogProto.numSetupClientInterfaceFailureDueToHal=0
mWifiLogProto.numSetupClientInterfaceFailureDueToWificond=0
mWifiLogProto.numSetupClientInterfaceFailureDueToSupplicant=0
mWifiLogProto.numSetupSoftApInterfaceFailureDueToHal=0
mWifiLogProto.numSetupSoftApInterfaceFailureDueToWificond=0
mWifiLogProto.numSetupSoftApInterfaceFailureDueToHostapd=0
StaEventList:
05-23 15:37:33.333 WIFI_ENABLED screenOn=false cellularData=false adaptiveConnectivity=true totalTxBytes=136 totalRxBytes=80 interfaceName=wlan0 interfaceRole=ROLE_CLIENT_PRIMARY
05-23 15:37:45.102 CMD_START_CONNECT screenOn=true cellularData=true adaptiveConnectivity=true, ConfigInfo: allowed_key_management=1 allowed_protocols=3 allowed_auth_algorithms=0 allowed_pairwise_ciphers=0 allowed_group_ciphers=0 hidden_ssid=false is_passpoint=false is_ephemeral=false has_ever_connected=true scan_rssi=-50 scan_freq=2447 mobileTxBytes=8515 mobileRxBytes=25392 totalTxBytes=11059 totalRxBytes=26620 interfaceName=wlan0 interfaceRole=ROLE_CLIENT_PRIMARY
05-23 15:37:46.369 CMD_ASSOCIATED_BSSID screenOn=true cellularData=true adaptiveConnectivity=true, supplicantStateChangeEvents: { AUTHENTICATING ASSOCIATING ASSOCIATED } mobileTxBytes=9328 mobileRxBytes=25973 totalTxBytes=12818 totalRxBytes=27674 interfaceName=wlan0 interfaceRole=ROLE_CLIENT_PRIMARY
05-23 15:37:46.385 NETWORK_CONNECTION_EVENT screenOn=true cellularData=true adaptiveConnectivity=true mobileTxBytes=9328 mobileRxBytes=25973 totalTxBytes=12818 totalRxBytes=27674 interfaceName=wlan0 interfaceRole=ROLE_CLIENT_PRIMARY
05-23 15:37:48.005 CMD_IP_CONFIGURATION_SUCCESSFUL lastRssi=-50 lastFreq=2447 lastLinkSpeed=2 screenOn=true cellularData=true adaptiveConnectivity=true, supplicantStateChangeEvents: { COMPLETED } mobileTxBytes=10355 mobileRxBytes=26884 totalTxBytes=15465 totalRxBytes=30491 interfaceName=wlan0 interfaceRole=ROLE_CLIENT_PRIMARY
05-23 15:37:55.259 NETWORK_AGENT_VALID_NETWORK lastRssi=-50 lastFreq=2447 lastLinkSpeed=1 screenOn=true cellularData=true adaptiveConnectivity=true mobileTxBytes=248855 mobileRxBytes=917988 totalTxBytes=272384 totalRxBytes=936412 interfaceName=wlan0 interfaceRole=ROLE_CLIENT_PRIMARY
UserActionEvents:
mWifiLogProto.numPasspointProviders=0
mWifiLogProto.numPasspointProviderInstallation=0
mWifiLogProto.numPasspointProviderInstallSuccess=0
mWifiLogProto.numPasspointProviderUninstallation=0
mWifiLogProto.numPasspointProviderUninstallSuccess=0
mWifiLogProto.numPasspointProvidersSuccessfullyConnected=0
mWifiLogProto.installedPasspointProfileTypeForR1:{}
mWifiLogProto.installedPasspointProfileTypeForR2:{}
mWifiLogProto.passpointProvisionStats.numProvisionSuccess=0
mWifiLogProto.passpointProvisionStats.provisionFailureCount:{}
mWifiLogProto.totalNumberOfPasspointConnectionsWithVenueUrl=0
mWifiLogProto.totalNumberOfPasspointConnectionsWithTermsAndConditionsUrl=0
mWifiLogProto.totalNumberOfPasspointAcceptanceOfTermsAndConditions=0
mWifiLogProto.totalNumberOfPasspointProfilesWithDecoratedIdentity=0
mWifiLogProto.passpointDeauthImminentScope={}
mWifiLogProto.numRadioModeChangeToMcc=0
mWifiLogProto.numRadioModeChangeToScc=0
mWifiLogProto.numRadioModeChangeToSbs=0
mWifiLogProto.numRadioModeChangeToDbs=0
mWifiLogProto.numSoftApUserBandPreferenceUnsatisfied=0
mTotalSsidsInScanHistogram:{1=1}
mTotalBssidsInScanHistogram:{1=1}
mAvailableOpenSsidsInScanHistogram:{1=1}
mAvailableOpenBssidsInScanHistogram:{1=1}
mAvailableSavedSsidsInScanHistogram:{1=1}
mAvailableSavedBssidsInScanHistogram:{1=1}
mAvailableOpenOrSavedSsidsInScanHistogram:{1=1}
mAvailableOpenOrSavedBssidsInScanHistogram:{1=1}
mAvailableSavedPasspointProviderProfilesInScanHistogram:{0=1}
mAvailableSavedPasspointProviderBssidsInScanHistogram:{0=1}
mWifiLogProto.partialAllSingleScanListenerResults=4
mWifiLogProto.fullBandAllSingleScanListenerResults=1
mWifiAwareMetrics:
mLastEnableUsageMs:0
mLastEnableUsageInThisSampleWindowMs:0
mAvailableTimeMs:0
mHistogramAwareAvailableDurationMs:
mLastEnableAwareMs:0
mLastEnableAwareInThisSampleWindowMs:0
mEnabledTimeMs:0
mHistogramAwareEnabledDurationMs:
mAttachDataByUid:
mAttachStatusData:
mHistogramAttachDuration:
mMaxPublishInApp:0
mMaxSubscribeInApp:0
mMaxDiscoveryInApp:0
mMaxPublishInSystem:0
mMaxSubscribeInSystem:0
mMaxDiscoveryInSystem:0
mPublishStatusData:
mSubscribeStatusData:
mHistogramPublishDuration:
mHistogramSubscribeDuration:
mAppsWithDiscoverySessionResourceFailure:
mMaxPublishWithRangingInApp:0
mMaxSubscribeWithRangingInApp:0
mMaxPublishWithRangingInSystem:0
mMaxSubscribeWithRangingInSystem:0
mHistogramSubscribeGeofenceMin:
mHistogramSubscribeGeofenceMax:
mNumSubscribesWithRanging:0
mNumMatchesWithRanging:0
mNumMatchesWithoutRangingForRangingEnabledSubscribes:0
mMaxNdiInApp:0
mMaxNdpInApp:0
mMaxSecureNdpInApp:0
mMaxNdiInSystem:0
mMaxNdpInSystem:0
mMaxSecureNdpInSystem:0
mMaxNdpPerNdi:0
mInBandNdpStatusData:
mOutOfBandNdpStatusData:
mNdpCreationTimeDuration:
mNdpCreationTimeMin:-1
mNdpCreationTimeMax:0
mNdpCreationTimeSum:0
mNdpCreationTimeSumSq:0
mNdpCreationTimeNumSamples:0
mHistogramNdpDuration:
mNdpRequestType:
mRttMetrics:
RTT Metrics:
mNumStartRangingCalls:0
mOverallStatusHistogram:{}
mMeasurementDurationApOnlyHistogram{}
mMeasurementDurationWithAwareHistogram{}
AP:numCalls=0, numIndividualCalls=0, perUidInfo={}, numRequestsHistogram={}, requestGapHistogram={}, statusHistogram={}, measuredDistanceHistogram={}
AWARE:numCalls=0, numIndividualCalls=0, perUidInfo={}, numRequestsHistogram={}, requestGapHistogram={}, statusHistogram={}, measuredDistanceHistogram={}
mNumStartContinuousRangingCalls:0
mContinuousRangingStartStatusHistogram:{}
mContinuousRangingTerminationReasonHistogram:{}
mContinuousRangingResultStatusHistogram:{}
mContinuousRangingIntervalHistogram:{}
mContinuousRangingSessionDurationHistogram:{}
mContinuousRangingDistancePerSession:{}
mPnoScanMetrics.numPnoScanAttempts=0
mPnoScanMetrics.numPnoScanFailed=0
mPnoScanMetrics.numPnoScanStartedOverOffload=0
mPnoScanMetrics.numPnoScanFailedOverOffload=0
mPnoScanMetrics.numPnoFoundNetworkEvents=0
mWifiLinkLayerUsageStats.loggingDurationMs=0
mWifiLinkLayerUsageStats.radioOnTimeMs=0
mWifiLinkLayerUsageStats.radioTxTimeMs=0
mWifiLinkLayerUsageStats.radioRxTimeMs=0
mWifiLinkLayerUsageStats.radioScanTimeMs=0
mWifiLinkLayerUsageStats.radioNanScanTimeMs=0
mWifiLinkLayerUsageStats.radioBackgroundScanTimeMs=0
mWifiLinkLayerUsageStats.radioRoamScanTimeMs=0
mWifiLinkLayerUsageStats.radioPnoScanTimeMs=0
mWifiLinkLayerUsageStats.radioHs20ScanTimeMs=0
mWifiLinkLayerUsageStats per Radio Stats: 
mWifiLogProto.connectToNetworkNotificationCount={}
mWifiLogProto.connectToNetworkNotificationActionCount={}
mWifiLogProto.openNetworkRecommenderBlocklistSize=0
mWifiLogProto.isWifiNetworksAvailableNotificationOn=true
mWifiLogProto.numOpenNetworkRecommendationUpdates=0
mWifiLogProto.numOpenNetworkConnectMessageFailedToSend=0
mWifiLogProto.observedHotspotR1ApInScanHistogram={0=1}
mWifiLogProto.observedHotspotR2ApInScanHistogram={0=1}
mWifiLogProto.observedHotspotR3ApInScanHistogram={0=1}
mWifiLogProto.observedHotspotR1EssInScanHistogram={0=1}
mWifiLogProto.observedHotspotR2EssInScanHistogram={0=1}
mWifiLogProto.observedHotspotR3EssInScanHistogram={0=1}
mWifiLogProto.observedHotspotR1ApsPerEssInScanHistogram={}
mWifiLogProto.observedHotspotR2ApsPerEssInScanHistogram={}
mWifiLogProto.observedHotspotR3ApsPerEssInScanHistogram={}
mWifiLogProto.observed80211mcSupportingApsInScanHistogram{0=1}
mWifiLogProto.bssidBlocklistStats:
networkSelectionFilteredBssidCount={0=5}
mBlockedBssidPerReasonCount={}
mBlockedConfigurationPerReasonCount={}, highMovementMultipleScansFeatureEnabled=true, numHighMovementConnectionSkipped=0, numHighMovementConnectionStarted=0, mBlockedBssidPerReasonCount={}, mBlockedConfigurationPerReasonCount={}
mSoftApTetheredEvents:
mSoftApLocalOnlyEvents:
-------WifiWake metrics-------
mTotalSessions: 0
mTotalWakeups: 0
mIgnoredStarts: 0
mIsInSession: false
Stored Sessions: 0
----end of WifiWake metrics----
mWifiLogProto.isMacRandomizationOn=false
mWifiLogProto.scoreExperimentId=
mExperimentValues.wifiDataStallMinTxBad=1
mExperimentValues.wifiDataStallMinTxSuccessWithoutRx=50
mExperimentValues.linkSpeedCountsLoggingEnabled=true
mExperimentValues.dataStallDurationMs=1500
mExperimentValues.dataStallTxTputThrKbps=2000
mExperimentValues.dataStallRxTputThrKbps=2000
mExperimentValues.dataStallTxPerThr=90
mExperimentValues.dataStallCcaLevelThr=256
WifiIsUnusableEventList: 
Hardware Version: 
mWifiUsabilityStatsEntriesRingBuffer:
timestamp_ms=18264,rssi=0,link_speed_mbps=0,total_tx_success=0,total_tx_retries=0,total_tx_bad=0,total_rx_success=0,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=0,wifi_usability_score=0,seq_num_to_framework=0,prediction_horizon_sec=0,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=0,probe_elapsed_time_ms_since_last_update=0,probe_mcs_rate_since_last_update=0,rx_link_speed_mbps=0,seq_num_inside_framework=0,is_same_bssid_and_freq=false,device_mobility_state=0,time_slice_duty_cycle_in_percent=0,channel_utilization_ratio=0,is_throughput_sufficient=false,is_wifi_scoring_enabled=false,is_cellular_data_available=false,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=0,is_network_capabilities_downstream_sufficient=0,is_network_capabilities_upstream_sufficient=0,is_throughput_predictor_downstream_sufficient=0,is_throughput_predictor_upstream_sufficient=0,is_bluetooth_connected=false,uwb_adapter_state=0,is_low_latency_activated=false,max_supported_tx_linkspeed=0,max_supported_rx_linkspeed=0,voip_mode=0,thread_device_role=0,capture_event_type=2,capture_event_type_subcode=-1,status_data_stall=0
timestamp_ms=18905,rssi=-50,link_speed_mbps=2,total_tx_success=0,total_tx_retries=0,total_tx_bad=0,total_rx_success=0,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=-1,seq_num_inside_framework=0,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=10,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=1,is_network_capabilities_downstream_sufficient=0,is_network_capabilities_upstream_sufficient=0,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=22206,rssi=-50,link_speed_mbps=1,total_tx_success=51,total_tx_retries=0,total_tx_bad=0,total_rx_success=33,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=1,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=0,is_network_capabilities_upstream_sufficient=0,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=25252,rssi=-50,link_speed_mbps=1,total_tx_success=68,total_tx_retries=0,total_tx_bad=0,total_rx_success=42,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=2,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=0,is_network_capabilities_upstream_sufficient=0,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=28289,rssi=-50,link_speed_mbps=2,total_tx_success=150,total_tx_retries=0,total_tx_bad=0,total_rx_success=132,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=3,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=31303,rssi=-50,link_speed_mbps=2,total_tx_success=252,total_tx_retries=0,total_tx_bad=0,total_rx_success=251,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=4,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=34316,rssi=-50,link_speed_mbps=2,total_tx_success=486,total_tx_retries=0,total_tx_bad=0,total_rx_success=657,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=5,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=37326,rssi=-50,link_speed_mbps=1,total_tx_success=745,total_tx_retries=0,total_tx_bad=0,total_rx_success=1274,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=6,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=40329,rssi=-50,link_speed_mbps=1,total_tx_success=754,total_tx_retries=0,total_tx_bad=0,total_rx_success=1281,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=7,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=43335,rssi=-50,link_speed_mbps=2,total_tx_success=764,total_tx_retries=0,total_tx_bad=0,total_rx_success=1291,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=8,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=46337,rssi=-50,link_speed_mbps=1,total_tx_success=775,total_tx_retries=0,total_tx_bad=0,total_rx_success=1297,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=9,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=49344,rssi=-50,link_speed_mbps=1,total_tx_success=781,total_tx_retries=0,total_tx_bad=0,total_rx_success=1302,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=10,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=52348,rssi=-50,link_speed_mbps=1,total_tx_success=781,total_tx_retries=0,total_tx_bad=0,total_rx_success=1302,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=11,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=55351,rssi=-50,link_speed_mbps=1,total_tx_success=792,total_tx_retries=0,total_tx_bad=0,total_rx_success=1308,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=12,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=58353,rssi=-50,link_speed_mbps=2,total_tx_success=794,total_tx_retries=0,total_tx_bad=0,total_rx_success=1310,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=13,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=61358,rssi=-50,link_speed_mbps=2,total_tx_success=796,total_tx_retries=0,total_tx_bad=0,total_rx_success=1312,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=14,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=64364,rssi=-50,link_speed_mbps=2,total_tx_success=796,total_tx_retries=0,total_tx_bad=0,total_rx_success=1312,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=15,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=67370,rssi=-50,link_speed_mbps=1,total_tx_success=798,total_tx_retries=0,total_tx_bad=0,total_rx_success=1314,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=16,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=70374,rssi=-50,link_speed_mbps=2,total_tx_success=798,total_tx_retries=0,total_tx_bad=0,total_rx_success=1314,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=17,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=73381,rssi=-50,link_speed_mbps=2,total_tx_success=798,total_tx_retries=0,total_tx_bad=0,total_rx_success=1314,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=18,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=76383,rssi=-50,link_speed_mbps=1,total_tx_success=798,total_tx_retries=0,total_tx_bad=0,total_rx_success=1314,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=19,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=79390,rssi=-50,link_speed_mbps=1,total_tx_success=798,total_tx_retries=0,total_tx_bad=0,total_rx_success=1314,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=20,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=82394,rssi=-50,link_speed_mbps=1,total_tx_success=798,total_tx_retries=0,total_tx_bad=0,total_rx_success=1314,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=21,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=85399,rssi=-50,link_speed_mbps=1,total_tx_success=803,total_tx_retries=0,total_tx_bad=0,total_rx_success=1318,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=22,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=88403,rssi=-50,link_speed_mbps=1,total_tx_success=805,total_tx_retries=0,total_tx_bad=0,total_rx_success=1320,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=23,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=91408,rssi=-50,link_speed_mbps=1,total_tx_success=805,total_tx_retries=0,total_tx_bad=0,total_rx_success=1320,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=24,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=94412,rssi=-50,link_speed_mbps=2,total_tx_success=807,total_tx_retries=0,total_tx_bad=0,total_rx_success=1322,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=25,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=97419,rssi=-50,link_speed_mbps=1,total_tx_success=809,total_tx_retries=0,total_tx_bad=0,total_rx_success=1324,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=26,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=100423,rssi=-50,link_speed_mbps=2,total_tx_success=811,total_tx_retries=0,total_tx_bad=0,total_rx_success=1326,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=27,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=103425,rssi=-50,link_speed_mbps=2,total_tx_success=811,total_tx_retries=0,total_tx_bad=0,total_rx_success=1326,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=28,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=106430,rssi=-50,link_speed_mbps=2,total_tx_success=811,total_tx_retries=0,total_tx_bad=0,total_rx_success=1326,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=29,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=109431,rssi=-50,link_speed_mbps=1,total_tx_success=827,total_tx_retries=0,total_tx_bad=0,total_rx_success=1344,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=30,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=112437,rssi=-50,link_speed_mbps=2,total_tx_success=829,total_tx_retries=0,total_tx_bad=0,total_rx_success=1346,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=31,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=115442,rssi=-50,link_speed_mbps=2,total_tx_success=829,total_tx_retries=0,total_tx_bad=0,total_rx_success=1346,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=32,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=118444,rssi=-50,link_speed_mbps=1,total_tx_success=831,total_tx_retries=0,total_tx_bad=0,total_rx_success=1348,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=33,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=121451,rssi=-50,link_speed_mbps=1,total_tx_success=831,total_tx_retries=0,total_tx_bad=0,total_rx_success=1348,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=34,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=124455,rssi=-50,link_speed_mbps=2,total_tx_success=833,total_tx_retries=0,total_tx_bad=0,total_rx_success=1350,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=35,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=127462,rssi=-50,link_speed_mbps=2,total_tx_success=833,total_tx_retries=0,total_tx_bad=0,total_rx_success=1350,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=36,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=130469,rssi=-50,link_speed_mbps=2,total_tx_success=833,total_tx_retries=0,total_tx_bad=0,total_rx_success=1350,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=37,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=133478,rssi=-50,link_speed_mbps=1,total_tx_success=833,total_tx_retries=0,total_tx_bad=0,total_rx_success=1350,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=38,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=136491,rssi=-50,link_speed_mbps=1,total_tx_success=833,total_tx_retries=0,total_tx_bad=0,total_rx_success=1350,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=39,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=139503,rssi=-50,link_speed_mbps=1,total_tx_success=833,total_tx_retries=0,total_tx_bad=0,total_rx_success=1350,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=40,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=142515,rssi=-50,link_speed_mbps=1,total_tx_success=833,total_tx_retries=0,total_tx_bad=0,total_rx_success=1350,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=41,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=145522,rssi=-50,link_speed_mbps=1,total_tx_success=833,total_tx_retries=0,total_tx_bad=0,total_rx_success=1350,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=42,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=148523,rssi=-50,link_speed_mbps=11,total_tx_success=834,total_tx_retries=0,total_tx_bad=0,total_rx_success=1351,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=43,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=151531,rssi=-50,link_speed_mbps=1,total_tx_success=838,total_tx_retries=0,total_tx_bad=0,total_rx_success=1354,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=44,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=154539,rssi=-50,link_speed_mbps=1,total_tx_success=840,total_tx_retries=0,total_tx_bad=0,total_rx_success=1356,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=45,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=157545,rssi=-50,link_speed_mbps=2,total_tx_success=840,total_tx_retries=0,total_tx_bad=0,total_rx_success=1356,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=46,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=160555,rssi=-50,link_speed_mbps=1,total_tx_success=842,total_tx_retries=0,total_tx_bad=0,total_rx_success=1358,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=47,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=163566,rssi=-50,link_speed_mbps=1,total_tx_success=842,total_tx_retries=0,total_tx_bad=0,total_rx_success=1358,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=48,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=166573,rssi=-50,link_speed_mbps=2,total_tx_success=844,total_tx_retries=0,total_tx_bad=0,total_rx_success=1360,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=49,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=169579,rssi=-50,link_speed_mbps=2,total_tx_success=844,total_tx_retries=0,total_tx_bad=0,total_rx_success=1360,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=50,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
timestamp_ms=172584,rssi=-50,link_speed_mbps=1,total_tx_success=844,total_tx_retries=0,total_tx_bad=0,total_rx_success=1360,total_radio_on_time_ms=0,total_radio_tx_time_ms=0,total_radio_rx_time_ms=0,total_scan_time_ms=0,total_nan_scan_time_ms=0,total_background_scan_time_ms=0,total_roam_scan_time_ms=0,total_pno_scan_time_ms=0,total_hotspot_2_scan_time_ms=0,wifi_score=-1,wifi_usability_score=-1,seq_num_to_framework=-1,prediction_horizon_sec=-1,total_cca_busy_freq_time_ms=0,total_radio_on_freq_time_ms=0,total_beacon_rx=0,probe_status_since_last_update=1,probe_elapsed_time_ms_since_last_update=-1,probe_mcs_rate_since_last_update=-1,rx_link_speed_mbps=2,seq_num_inside_framework=51,is_same_bssid_and_freq=true,device_mobility_state=0,time_slice_duty_cycle_in_percent=-1,access_category=0,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=1,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=2,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,access_category=3,contention_time_min_micros=0,contention_time_max_micros=0,contention_time_avg_micros=0,contention_num_samples=0,channel_utilization_ratio=80,is_throughput_sufficient=true,is_wifi_scoring_enabled=true,is_cellular_data_available=true,sta_count=0,channel_utilization=0,wifi_link_count=0,mlo_mode=0,tx_transmitted_bytes=0,rx_transmitted_bytes=0,label_bad_event_count=0,wifi_framework_state=2,is_network_capabilities_downstream_sufficient=1,is_network_capabilities_upstream_sufficient=1,is_throughput_predictor_downstream_sufficient=1,is_throughput_predictor_upstream_sufficient=1,is_bluetooth_connected=false,uwb_adapter_state=1,is_low_latency_activated=false,max_supported_tx_linkspeed=11,max_supported_rx_linkspeed=11,voip_mode=-1,thread_device_role=0,capture_event_type=1,capture_event_type_subcode=0,status_data_stall=0
mWifiUsabilityStatsTrainingExamples:
mMobilityStatePnoStatsMap:
device_mobility_state=0,num_times_entered_state=1,total_duration_ms=0,pno_duration_ms=0
WifiP2pMetrics:
mConnectionEvents:
mGroupEvents:
mWifiP2pStatsProto.numPersistentGroup=0
mWifiP2pStatsProto.numTotalPeerScans=0
mWifiP2pStatsProto.numTotalServiceScans=0
mDppMetrics:
---Easy Connect/DPP metrics---
mWifiDppLogProto.numDppConfiguratorInitiatorRequests=0
mWifiDppLogProto.numDppEnrolleeInitiatorRequests=0
mWifiDppLogProto.numDppEnrolleeResponderRequests=0
mWifiDppLogProto.numDppEnrolleeResponderSuccess=0
mWifiDppLogProto.numDppEnrolleeSuccess=0
mWifiDppLogProto.numDppR1CapableEnrolleeResponderDevices=0
mWifiDppLogProto.numDppR2CapableEnrolleeResponderDevices=0
mWifiDppLogProto.numDppR2EnrolleeResponderIncompatibleConfiguration=0
---End of Easy Connect/DPP metrics---
mWifiConfigStoreReadDurationHistogram:{0=1, 3=1}
mWifiConfigStoreWriteDurationHistogram:{0=1, 1=1}
mLinkProbeSuccessRssiCounts:{}
mLinkProbeFailureRssiCounts:{}
mLinkProbeSuccessLinkSpeedCounts:{}
mLinkProbeFailureLinkSpeedCounts:{}
mLinkProbeSuccessSecondsSinceLastTxSuccessHistogram:{}
mLinkProbeFailureSecondsSinceLastTxSuccessHistogram:{}
mLinkProbeSuccessElapsedTimeMsHistogram:{}
mLinkProbeFailureReasonCounts:{}
mLinkProbeExperimentProbeCounts:{}
mNetworkSelectionExperimentPairNumChoicesCounts:{Pair{42902385 42330058}=NetworkSelectionExperimentResults{sameSelectionNumChoicesCounter={1=1}, differentSelectionNumChoicesCounter={}}, Pair{42598152 42330058}=NetworkSelectionExperimentResults{sameSelectionNumChoicesCounter={1=1}, differentSelectionNumChoicesCounter={}}, Pair{42504592 42330058}=NetworkSelectionExperimentResults{sameSelectionNumChoicesCounter={1=1}, differentSelectionNumChoicesCounter={}}}
mLinkProbeStaEventCount:0
mWifiNetworkRequestApiLog:
num_apps: 0
num_concurrent_connection: 0
num_connect_on_primary_iface: 0
num_connect_on_secondary_iface: 0
num_connect_success_on_primary_iface: 0
num_connect_success_on_secondary_iface: 0
num_request: 0
num_user_approval_bypass: 0
num_user_reject: 0

mWifiNetworkRequestApiMatchSizeHistogram:
{}
mWifiNetworkRequestApiConnectionDurationSecOnPrimaryIfaceHistogram:
{}
mWifiNetworkRequestApiConnectionDurationSecOnSecondaryIfaceHistogram:
{}
mWifiNetworkRequestApiConcurrentConnectionDurationSecHistogram:
{}
mWifiNetworkSuggestionApiLog:
num_connect_failure: 0
num_connect_success: 0
num_modification: 0
num_multiple_suggestions: 0
num_priority_groups: 0
num_saved_networks_with_configured_suggestion: 0
user_revoke_app_suggestion_permission: 0

mWifiNetworkSuggestionApiMatchSizeHistogram:
{}
mWifiNetworkSuggestionApiAppTypeCounter:
{}
mWifiNetworkSuggestionPriorityGroups:
{}
mWifiNetworkSuggestionCoexistSavedNetworks:
{}
mUserApprovalSuggestionAppUiUserReaction:
mUserApprovalCarrierUiUserReaction:
mNetworkIdToNominatorId:
{0=2}
mWifiLockStats:
high_perf_active_time_ms: 0
low_latency_active_time_ms: 0

mWifiLockHighPerfAcqDurationSecHistogram:
{}
mWifiLockLowLatencyAcqDurationSecHistogram:
{}
mWifiLockHighPerfActiveSessionDurationSecHistogram:
{}
mWifiLockLowLatencyActiveSessionDurationSecHistogram:
{}
mWifiToggleStats:
num_toggle_off_normal: 0
num_toggle_off_privileged: 0
num_toggle_on_normal: 0
num_toggle_on_privileged: 0

mWifiLogProto.numAddOrUpdateNetworkCalls=0
mWifiLogProto.numEnableNetworkCalls=0
mWifiLogProto.txLinkSpeedCount2g={1=29, 2=22, 11=1}
mWifiLogProto.txLinkSpeedCount5gLow={}
mWifiLogProto.txLinkSpeedCount5gMid={}
mWifiLogProto.txLinkSpeedCount5gHigh={}
mWifiLogProto.txLinkSpeedCount6gLow={}
mWifiLogProto.txLinkSpeedCount6gMid={}
mWifiLogProto.txLinkSpeedCount6gHigh={}
mWifiLogProto.rxLinkSpeedCount2g={2=51}
mWifiLogProto.rxLinkSpeedCount5gLow={}
mWifiLogProto.rxLinkSpeedCount5gMid={}
mWifiLogProto.rxLinkSpeedCount5gHigh={}
mWifiLogProto.rxLinkSpeedCount6gLow={}
mWifiLogProto.rxLinkSpeedCount6gMid={}
mWifiLogProto.rxLinkSpeedCount6gHigh={}
mWifiLogProto.numIpRenewalFailure=0
mWifiLogProto.connectionDurationStats=connectionDurationSufficientThroughputMs=0, connectionDurationInSufficientThroughputMs=0, connectionDurationInSufficientThroughputDefaultWifiMs=0, connectionDurationCellularDataOffMs=0
mWifiLogProto.isExternalWifiScorerOn=false
mWifiLogProto.wifiOffMetrics=numWifiOff=0, numWifiOffDeferring=0, numWifiOffDeferringTimeout=0, wifiOffDeferringTimeHistogram={}
mWifiLogProto.softApConfigLimitationMetrics=numSecurityTypeResetToDefault=0, numMaxClientSettingResetToDefault=0, numClientControlByUserResetToDefault=0, maxClientSettingWhenReachHistogram={}
mChannelUtilizationHistogram2G:
{[75,100)=52}
mChannelUtilizationHistogramAbove2G:
{}
mTxThroughputMbpsHistogram2G:
{[10,15)=52}
mRxThroughputMbpsHistogram2G:
{[10,15)=52}
mTxThroughputMbpsHistogramAbove2G:
{}
mRxThroughputMbpsHistogramAbove2G:
{}
mCarrierWifiMetrics:
numConnectionSuccess=0, numConnectionAuthFailure=0, numConnectionNonAuthFailure0
FirstConnectAfterBootStats{wifiEnabledAtBoot=Attempt{timestampSinceBootMillis=4074,isSuccess=true},firstNetworkSelectionAttempt{timestampSinceBootMillis=16600,isSuccess=true},firstL2ConnectionAttempt{timestampSinceBootMillis=18365,isSuccess=true},firstL3ConnectionAttempt{timestampSinceBootMillis=20114,isSuccess=true}}
WifiToWifiSwitchStats{isMakeBeforeBreakSupported=false,wifiToWifiSwitchTriggerCount=0,makeBeforeBreakTriggerCount=0,makeBeforeBreakNoInternetCount=0,makeBeforeBreakRecoverPrimaryCount=0,makeBeforeBreakInternetValidatedCount=0,makeBeforeBreakSuccessCount=0,makeBeforeBreakLingerCompletedCount=0,makeBeforeBreakLingeringDurationSeconds={}}
mInitPartialScanTotalCount:
0
mInitPartialScanSuccessCount:
0
mInitPartialScanFailureCount:
0
mInitPartialScanSuccessHistogram:
{}
mInitPartialScanFailureHistogram:
{}

Dump of WifiNetworkSuggestionsManager
WifiNetworkSuggestionsManager - Networks Begin ----
WifiNetworkSuggestionsManager - Networks End ----

Dump of WifiBackupRestore

Dump of BackupRestoreController
mLastBackupDataRetrievedTimestamp: N/A
mLastBackupDataRestoredTimestamp: N/A

ScoringParams: rssi2=-83:-80:-73:-60,rssi5=-80:-77:-70:-57,rssi6=-80:-77:-70:-57,pps=0:16:100,horizon=15,nud=8,expid=0


Dump of WifiSettingsConfigStore
Settings:
wifi_native_supported_sta_bands=3
wifi_p2p_pending_factory_reset=false
d2d_allowed_when_infra_sta_disabled=false
wifi_p2p_device_name=null
wifi_scan_always_enabled=false
default_wifi_networks_available_notification_on=true
wifi_sta_factory_mac_address=02:15:b2:00:00:00
wifi_static_chip_info=[{"chipId":0,"availableModes":[{"id":0,"availableCombinations":[{"limits":[{"maxIfaces":1,"types":[0]},{"maxIfaces":1,"types":[3]}]}]},{"id":1,"availableCombinations":[{"limits":[{"maxIfaces":1,"types":[1]}]}]}]}]
wifi_scan_throttle_enabled=true
wifi_verbose_logging_enabled=false
default_wifi_wakeup_enabled=true
wifi_native_extended_supported_features=[J@a81e3d2
wifi_available_soft_ap_freqs_mhz=[2412,2417,2422,2427,2432,2437,2442,2447,2452,2457,2462]
wifi_networks_available_notification_on=true
default_wifi_scan_always_enabled=false
wifi_last_country_code=US
wifi_wakeup_enabled=true
wep_allowed=true
supplicant_hal_aidl_service_version=5
Migration data for shared to private settings migration:
wifi_scan_always_enabled=false
wifi_wakeup_enabled=true
wep_allowed=true
d2d_allowed_when_infra_sta_disabled=false
wifi_networks_available_notification_on=true

mRevertCountryCodeOnCellularLoss: false
DefaultCountryCode(system property): null
DefaultCountryCode(config store): null
mTelephonyCountryCode: US
mTelephonyCountryTimestamp: 05-23 15:37:34.210
mOverrideCountryCode: null
mAllCmmReadyTimestamp: 05-23 15:37:32.632
isAllCmmReady: false
mAmmToReadyForChangeMap: {ConcreteClientModeManager{id=4074 iface=wlan0 role=ROLE_CLIENT_PRIMARY}=false}
mDisconnectWifiToForceUpdateCount: 0
mDriverCountryCode: US
mDriverCountryCodeUpdatedTimestamp: 05-23 15:37:34.257
mFrameworkCountryCode: null
mFrameworkCountryCodeUpdatedTimestamp: N/A
isDriverSupportedRegChangedEvent: false
providerId=6, ScoreFilter=Score(Policies : 0), Filter=[ Transports: WIFI Capabilities: NOT_METERED&INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VPN&NOT_ROAMING&NOT_CONGESTED&NOT_SUSPENDED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED LinkUpBandwidth>=1048576Kbps LinkDnBandwidth>=1048576Kbps Specifier: <android.net.MatchAllNetworkSpecifier@0> SubscriptionIds: {1} UnderlyingNetworks: Null], requests=42
  {NetworkRequest [ REQUEST id=40, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1001 RequestorUid: 1001 RequestorPkg: com.android.phone UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=64, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10195 RequestorUid: 10195 RequestorPkg: com.android.systemui UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=33, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10146 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=31, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=35, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1000 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=1, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VPN&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=48, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10195 RequestorUid: 10195 RequestorPkg: com.android.systemui UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=29, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10165 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=25, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10171 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=27, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10167 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=62, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10195 RequestorUid: 10195 RequestorPkg: com.android.systemui UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=37, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1073 RequestorUid: 1073 RequestorPkg: com.google.android.networkstack UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=58, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10195 RequestorUid: 10195 RequestorPkg: com.android.systemui UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=23, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10205 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=53, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1001 RequestorUid: 1001 RequestorPkg: com.android.phone UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=51, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1001 RequestorUid: 1001 RequestorPkg: com.android.phone UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=21, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10225 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=76, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10167 RequestorUid: 10167 RequestorPkg: com.google.android.inputmethod.latin UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=79, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1000 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=81, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10147 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=86, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10159 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=94, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10171 RequestorUid: 10171 RequestorPkg: com.google.android.apps.photos UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=109, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10149 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=113, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10158 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=125, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10166 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=151, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=164, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10226 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=170, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10163 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=178, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10155 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=182, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10144 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=188, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=190, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=193, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=195, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=202, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10157 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=204, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=206, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=210, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=220, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10226 RequestorUid: 10226 RequestorPkg: com.google.android.rkpdapp UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=224, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=226, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=229, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=true}
WifiNetworkFactory: mGenericConnectionReqCount 42
WifiNetworkFactory: mActiveSpecificNetworkRequest null
WifiNetworkFactory: mUserApprovedAccessPointMap {}
WifiNetworkFactory: mLocalOnlyDisconnectionStatusListenerPerApp {}
providerId=7, ScoreFilter=Score(Policies : 0), Filter=[ Transports: WIFI Capabilities: NOT_METERED&INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VPN&NOT_ROAMING&NOT_CONGESTED&NOT_SUSPENDED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED LinkUpBandwidth>=1048576Kbps LinkDnBandwidth>=1048576Kbps Specifier: <android.net.MatchAllNetworkSpecifier@0> SubscriptionIds: {1} UnderlyingNetworks: Null], requests=42
  {NetworkRequest [ REQUEST id=40, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1001 RequestorUid: 1001 RequestorPkg: com.android.phone UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=64, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10195 RequestorUid: 10195 RequestorPkg: com.android.systemui UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=33, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10146 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=31, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=35, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1000 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=1, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VPN&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=48, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10195 RequestorUid: 10195 RequestorPkg: com.android.systemui UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=29, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10165 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=25, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10171 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=27, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10167 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=62, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10195 RequestorUid: 10195 RequestorPkg: com.android.systemui UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=37, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1073 RequestorUid: 1073 RequestorPkg: com.google.android.networkstack UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=58, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10195 RequestorUid: 10195 RequestorPkg: com.android.systemui UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=23, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10205 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=53, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1001 RequestorUid: 1001 RequestorPkg: com.android.phone UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=51, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1001 RequestorUid: 1001 RequestorPkg: com.android.phone UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=21, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10225 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=76, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10167 RequestorUid: 10167 RequestorPkg: com.google.android.inputmethod.latin UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=79, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1000 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=81, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10147 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=86, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10159 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=94, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10171 RequestorUid: 10171 RequestorPkg: com.google.android.apps.photos UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=109, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10149 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=113, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10158 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=125, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10166 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=151, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=164, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10226 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=170, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10163 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=178, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10155 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=182, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10144 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=188, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=190, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=193, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=195, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=202, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10157 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=204, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=206, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=210, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=220, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10226 RequestorUid: 10226 RequestorPkg: com.google.android.rkpdapp UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=224, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=226, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=229, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=true}
UntrustedWifiNetworkFactory: mConnectionReqCount 0
providerId=9, ScoreFilter=Score(Policies : 0), Filter=[ Transports: WIFI Capabilities: NOT_METERED&INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VPN&NOT_ROAMING&NOT_CONGESTED&NOT_SUSPENDED&OEM_PAID&OEM_PRIVATE&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED LinkUpBandwidth>=1048576Kbps LinkDnBandwidth>=1048576Kbps Specifier: <android.net.MatchAllNetworkSpecifier@0> UnderlyingNetworks: Null], requests=42
  {NetworkRequest [ REQUEST id=40, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1001 RequestorUid: 1001 RequestorPkg: com.android.phone UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=64, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10195 RequestorUid: 10195 RequestorPkg: com.android.systemui UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=33, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10146 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=31, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=35, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1000 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=1, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VPN&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=48, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10195 RequestorUid: 10195 RequestorPkg: com.android.systemui UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=29, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10165 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=25, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10171 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=27, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10167 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=62, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10195 RequestorUid: 10195 RequestorPkg: com.android.systemui UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=37, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1073 RequestorUid: 1073 RequestorPkg: com.google.android.networkstack UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=58, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10195 RequestorUid: 10195 RequestorPkg: com.android.systemui UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=23, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10205 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=53, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1001 RequestorUid: 1001 RequestorPkg: com.android.phone UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=51, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1001 RequestorUid: 1001 RequestorPkg: com.android.phone UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=21, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10225 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=76, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10167 RequestorUid: 10167 RequestorPkg: com.google.android.inputmethod.latin UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=79, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1000 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=81, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10147 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=86, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10159 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=94, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10171 RequestorUid: 10171 RequestorPkg: com.google.android.apps.photos UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=109, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10149 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=113, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10158 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=125, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10166 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=151, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=164, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10226 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=170, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10163 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=178, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10155 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=182, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10144 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=188, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=190, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=193, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=195, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=202, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10157 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=204, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=206, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=210, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=220, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10226 RequestorUid: 10226 RequestorPkg: com.google.android.rkpdapp UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=224, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=226, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=true}
  {NetworkRequest [ REQUEST id=229, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=true}
OemPaidWifiNetworkFactory: mOemPaidConnectionReqCount 0
OemPaidWifiNetworkFactory: mOemPrivateConnectionReqCount 0
providerId=8, ScoreFilter=Score(Policies : 0), Filter=[ Transports: WIFI Capabilities: NOT_METERED&INTERNET&TRUSTED&NOT_VPN&NOT_ROAMING&NOT_CONGESTED&NOT_SUSPENDED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED LinkUpBandwidth>=1048576Kbps LinkDnBandwidth>=1048576Kbps Specifier: <android.net.MatchAllNetworkSpecifier@0> SubscriptionIds: {1} UnderlyingNetworks: Null], requests=0
RestrictedWifiNetworkFactory: mConnectionReqCount 0
providerId=10, ScoreFilter=Score(Policies : 0), Filter=[ Transports: WIFI Capabilities: NOT_METERED&INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VPN&NOT_ROAMING&NOT_CONGESTED&NOT_SUSPENDED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED LinkUpBandwidth>=1048576Kbps LinkDnBandwidth>=1048576Kbps Specifier: <android.net.MatchAllNetworkSpecifier@0> UnderlyingNetworks: Null], requests=42
  {NetworkRequest [ REQUEST id=40, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1001 RequestorUid: 1001 RequestorPkg: com.android.phone UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=64, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10195 RequestorUid: 10195 RequestorPkg: com.android.systemui UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=33, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10146 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=31, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=35, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1000 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=1, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VPN&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=48, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10195 RequestorUid: 10195 RequestorPkg: com.android.systemui UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=29, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10165 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=25, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10171 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=27, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10167 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=62, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10195 RequestorUid: 10195 RequestorPkg: com.android.systemui UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=37, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1073 RequestorUid: 1073 RequestorPkg: com.google.android.networkstack UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=58, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10195 RequestorUid: 10195 RequestorPkg: com.android.systemui UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=23, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10205 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=53, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1001 RequestorUid: 1001 RequestorPkg: com.android.phone UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=51, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1001 RequestorUid: 1001 RequestorPkg: com.android.phone UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=21, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10225 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=76, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10167 RequestorUid: 10167 RequestorPkg: com.google.android.inputmethod.latin UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=79, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 1000 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=81, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10147 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=86, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10159 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=94, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10171 RequestorUid: 10171 RequestorPkg: com.google.android.apps.photos UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=109, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10149 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=113, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10158 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=125, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10166 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=151, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=164, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10226 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=170, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10163 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=178, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10155 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=182, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10144 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=188, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=190, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=193, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=195, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=202, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10157 RequestorUid: 1000 RequestorPkg: android UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=204, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=206, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10182 RequestorUid: 10182 RequestorPkg: com.google.android.youtube UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=210, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=220, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10226 RequestorUid: 10226 RequestorPkg: com.google.android.rkpdapp UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=224, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=226, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=false}
  {NetworkRequest [ REQUEST id=229, [ Capabilities: INTERNET&NOT_RESTRICTED&TRUSTED&NOT_VCN_MANAGED&NOT_BANDWIDTH_CONSTRAINED Uid: 10153 RequestorUid: 10153 RequestorPkg: com.android.vending UnderlyingNetworks: Null] ], requested=false}
Dump of MultiInternetWifiNetworkFactory
Dump of MultiInternetManager
WifiMultiInternet: mStaConcurrencyMultiInternetMode 0
Dump of SsidTranslator
mCurrentLocaleCharset: null
mCharsetsPerLocaleLanguage Begin ---
mCharsetsPerLocaleLanguage End ---
mTranslatedBssids Begin ---
mTranslatedBssids End ---
mUntranslatedBssids Begin ---
mUntranslatedBssids End ---
Wlan Wake Reasons: totalCmdEventWake 0 totalDriverFwLocalWake 0 totalRxDataWake 0 rxUnicast 0 rxMulticast 0 rxBroadcast 0 icmp 0 icmp6 0 icmp6Ra 0 icmp6Na 0 icmp6Ns 0 ipv4RxMulticast 0 ipv6Multicast 0 otherRxMulticast 0

Dump of WifiConfigManager
WifiConfigManager - Log Begin ----
2026-05-23T15:37:32.469502 - clearInternalData: Clearing all internal data
2026-05-23T15:37:33.360308 - onCellularConnectivityChanged:0
2026-05-23T15:37:35.488090 - onCellularConnectivityChanged:1
WifiConfigManager - Log End ----
WifiConfigManager - mIsCurrentUserAdmin:true
WifiConfigManager - Configured networks Begin ----
* ID: 0 SSID: "AndroidWifi" PROVIDER-NAME: null BSSID: null FQDN: null HOME-PROVIDER-NETWORK: false PRIO: 0 HIDDEN: false PMF: false CarrierId: -1 SubscriptionId: -1 SubscriptionGroup: null Currently Connected: true User Selected: false
 NetworkSelectionStatus NETWORK_SELECTION_ENABLED
 hasEverConnected: true
 hasNeverDetectedCaptivePortal: true
 hasEverValidatedInternetAccess: true
 mCandidateSecurityParams: Security Parameters:
 Type: 0
 Enabled: true
 KeyMgmt: NONE
 Protocols: WPA RSN
 AuthAlgorithms:
 PairwiseCiphers:
 GroupCiphers:
 GroupMgmtCiphers:
 SuiteBCiphers:
 RequirePmf: false
 IsAddedByAutoUpgrade: false
 IsSaeH2eOnlyMode: false
 IsSaePkOnlyMode: false
 mLastUsedSecurityParams: Security Parameters:
 Type: 0
 Enabled: true
 KeyMgmt: NONE
 Protocols: WPA RSN
 AuthAlgorithms:
 PairwiseCiphers:
 GroupCiphers:
 GroupMgmtCiphers:
 SuiteBCiphers:
 RequirePmf: false
 IsAddedByAutoUpgrade: false
 IsSaeH2eOnlyMode: false
 IsSaePkOnlyMode: false
 numAssociation 1
 validatedInternetAccess shared trusted
 macRandomizationSetting: 3
 mRandomizedMacAddress: 46:a0:89:b8:78:55
 randomizedMacExpirationTimeMs: 2026-05-24T15:37:47.928
 randomizedMacLastModifiedTimeMs: <none>
 persistentMacRandomizationSeed: 0
 mIsSendDhcpHostnameEnabled: true
 deletionPriority: 0
 KeyMgmt: NONE Protocols: WPA RSN
 AuthAlgorithms:
 PairwiseCiphers:
 GroupCiphers:
 GroupMgmtCiphers:
 SuiteBCiphers:
 PSK/SAE: 
SecurityParams List:
Security Parameters:
 Type: 0
 Enabled: true
 KeyMgmt: NONE
 Protocols: WPA RSN
 AuthAlgorithms:
 PairwiseCiphers:
 GroupCiphers:
 GroupMgmtCiphers:
 SuiteBCiphers:
 RequirePmf: false
 IsAddedByAutoUpgrade: false
 IsSaeH2eOnlyMode: false
 IsSaePkOnlyMode: false
Security Parameters:
 Type: 6
 Enabled: true
 KeyMgmt: OWE
 Protocols: RSN
 AuthAlgorithms:
 PairwiseCiphers: CCMP GCMP_256 GCMP_128
 GroupCiphers: CCMP GCMP_256 GCMP_128
 GroupMgmtCiphers:
 SuiteBCiphers:
 RequirePmf: true
 IsAddedByAutoUpgrade: true
 IsSaeH2eOnlyMode: false
 IsSaePkOnlyMode: false

Enterprise config:
 ocsp: 0
 trust_on_first_use: false
 user_approve_no_ca_cert: false
 selected_rcoi: 0
 minimum_tls_version: 0
 enable_conservative_peer_mode: false
 tofu_dialog_state: 0
 tofu_connection_state: 0
IP config:
IP assignment: DHCP
Proxy settings: NONE
 cuid=10190 cname=com.google.android.googlesdksetup luid=10190 lname=com.google.android.googlesdksetup lcuid=10190 allowAutojoin=true mAllowedAutoJoinInAdvancedProtection=true noInternetAccessExpected=false mostRecentlyConnected=true 
lastConnected: 2026-05-23T15:37:48.117 

numRebootsSinceLastUse: 0
recentFailure: Association Rejection code: 0, last update time: 0
bssidAllowlist unset
vendorData unset
IsDppConfigurator: true
HasEncryptedPreSharedKey: false
 setWifi7Enabled=true
 mIsAllowedToUpdateByOtherUsers=true
 mCreatorUserId=0

WifiConfigManager - Configured networks End ----
WifiConfigManager - ConfigurationMap Begin ----
mPerId={0=* ID: 0 SSID: "AndroidWifi" PROVIDER-NAME: null BSSID: null FQDN: null HOME-PROVIDER-NETWORK: false PRIO: 0 HIDDEN: false PMF: false CarrierId: -1 SubscriptionId: -1 SubscriptionGroup: null Currently Connected: true User Selected: false
 NetworkSelectionStatus NETWORK_SELECTION_ENABLED
 hasEverConnected: true
 hasNeverDetectedCaptivePortal: true
 hasEverValidatedInternetAccess: true
 mCandidateSecurityParams: Security Parameters:
 Type: 0
 Enabled: true
 KeyMgmt: NONE
 Protocols: WPA RSN
 AuthAlgorithms:
 PairwiseCiphers:
 GroupCiphers:
 GroupMgmtCiphers:
 SuiteBCiphers:
 RequirePmf: false
 IsAddedByAutoUpgrade: false
 IsSaeH2eOnlyMode: false
 IsSaePkOnlyMode: false
 mLastUsedSecurityParams: Security Parameters:
 Type: 0
 Enabled: true
 KeyMgmt: NONE
 Protocols: WPA RSN
 AuthAlgorithms:
 PairwiseCiphers:
 GroupCiphers:
 GroupMgmtCiphers:
 SuiteBCiphers:
 RequirePmf: false
 IsAddedByAutoUpgrade: false
 IsSaeH2eOnlyMode: false
 IsSaePkOnlyMode: false
 numAssociation 1
 validatedInternetAccess shared trusted
 macRandomizationSetting: 3
 mRandomizedMacAddress: 46:a0:89:b8:78:55
 randomizedMacExpirationTimeMs: 2026-05-24T15:37:47.928
 randomizedMacLastModifiedTimeMs: <none>
 persistentMacRandomizationSeed: 0
 mIsSendDhcpHostnameEnabled: true
 deletionPriority: 0
 KeyMgmt: NONE Protocols: WPA RSN
 AuthAlgorithms:
 PairwiseCiphers:
 GroupCiphers:
 GroupMgmtCiphers:
 SuiteBCiphers:
 PSK/SAE: 
SecurityParams List:
Security Parameters:
 Type: 0
 Enabled: true
 KeyMgmt: NONE
 Protocols: WPA RSN
 AuthAlgorithms:
 PairwiseCiphers:
 GroupCiphers:
 GroupMgmtCiphers:
 SuiteBCiphers:
 RequirePmf: false
 IsAddedByAutoUpgrade: false
 IsSaeH2eOnlyMode: false
 IsSaePkOnlyMode: false
Security Parameters:
 Type: 6
 Enabled: true
 KeyMgmt: OWE
 Protocols: RSN
 AuthAlgorithms:
 PairwiseCiphers: CCMP GCMP_256 GCMP_128
 GroupCiphers: CCMP GCMP_256 GCMP_128
 GroupMgmtCiphers:
 SuiteBCiphers:
 RequirePmf: true
 IsAddedByAutoUpgrade: true
 IsSaeH2eOnlyMode: false
 IsSaePkOnlyMode: false

Enterprise config:
 ocsp: 0
 trust_on_first_use: false
 user_approve_no_ca_cert: false
 selected_rcoi: 0
 minimum_tls_version: 0
 enable_conservative_peer_mode: false
 tofu_dialog_state: 0
 tofu_connection_state: 0
IP config:
IP assignment: DHCP
Proxy settings: NONE
 cuid=10190 cname=com.google.android.googlesdksetup luid=10190 lname=com.google.android.googlesdksetup lcuid=10190 allowAutojoin=true mAllowedAutoJoinInAdvancedProtection=true noInternetAccessExpected=false mostRecentlyConnected=true 
lastConnected: 2026-05-23T15:37:48.117 

numRebootsSinceLastUse: 0
recentFailure: Association Rejection code: 0, last update time: 0
bssidAllowlist unset
vendorData unset
IsDppConfigurator: true
HasEncryptedPreSharedKey: false
 setWifi7Enabled=true
 mIsAllowedToUpdateByOtherUsers=true
 mCreatorUserId=0
}
mPerIDForCurrentUser={0=* ID: 0 SSID: "AndroidWifi" PROVIDER-NAME: null BSSID: null FQDN: null HOME-PROVIDER-NETWORK: false PRIO: 0 HIDDEN: false PMF: false CarrierId: -1 SubscriptionId: -1 SubscriptionGroup: null Currently Connected: true User Selected: false
 NetworkSelectionStatus NETWORK_SELECTION_ENABLED
 hasEverConnected: true
 hasNeverDetectedCaptivePortal: true
 hasEverValidatedInternetAccess: true
 mCandidateSecurityParams: Security Parameters:
 Type: 0
 Enabled: true
 KeyMgmt: NONE
 Protocols: WPA RSN
 AuthAlgorithms:
 PairwiseCiphers:
 GroupCiphers:
 GroupMgmtCiphers:
 SuiteBCiphers:
 RequirePmf: false
 IsAddedByAutoUpgrade: false
 IsSaeH2eOnlyMode: false
 IsSaePkOnlyMode: false
 mLastUsedSecurityParams: Security Parameters:
 Type: 0
 Enabled: true
 KeyMgmt: NONE
 Protocols: WPA RSN
 AuthAlgorithms:
 PairwiseCiphers:
 GroupCiphers:
 GroupMgmtCiphers:
 SuiteBCiphers:
 RequirePmf: false
 IsAddedByAutoUpgrade: false
 IsSaeH2eOnlyMode: false
 IsSaePkOnlyMode: false
 numAssociation 1
 validatedInternetAccess shared trusted
 macRandomizationSetting: 3
 mRandomizedMacAddress: 46:a0:89:b8:78:55
 randomizedMacExpirationTimeMs: 2026-05-24T15:37:47.928
 randomizedMacLastModifiedTimeMs: <none>
 persistentMacRandomizationSeed: 0
 mIsSendDhcpHostnameEnabled: true
 deletionPriority: 0
 KeyMgmt: NONE Protocols: WPA RSN
 AuthAlgorithms:
 PairwiseCiphers:
 GroupCiphers:
 GroupMgmtCiphers:
 SuiteBCiphers:
 PSK/SAE: 
SecurityParams List:
Security Parameters:
 Type: 0
 Enabled: true
 KeyMgmt: NONE
 Protocols: WPA RSN
 AuthAlgorithms:
 PairwiseCiphers:
 GroupCiphers:
 GroupMgmtCiphers:
 SuiteBCiphers:
 RequirePmf: false
 IsAddedByAutoUpgrade: false
 IsSaeH2eOnlyMode: false
 IsSaePkOnlyMode: false
Security Parameters:
 Type: 6
 Enabled: true
 KeyMgmt: OWE
 Protocols: RSN
 AuthAlgorithms:
 PairwiseCiphers: CCMP GCMP_256 GCMP_128
 GroupCiphers: CCMP GCMP_256 GCMP_128
 GroupMgmtCiphers:
 SuiteBCiphers:
 RequirePmf: true
 IsAddedByAutoUpgrade: true
 IsSaeH2eOnlyMode: false
 IsSaePkOnlyMode: false

Enterprise config:
 ocsp: 0
 trust_on_first_use: false
 user_approve_no_ca_cert: false
 selected_rcoi: 0
 minimum_tls_version: 0
 enable_conservative_peer_mode: false
 tofu_dialog_state: 0
 tofu_connection_state: 0
IP config:
IP assignment: DHCP
Proxy settings: NONE
 cuid=10190 cname=com.google.android.googlesdksetup luid=10190 lname=com.google.android.googlesdksetup lcuid=10190 allowAutojoin=true mAllowedAutoJoinInAdvancedProtection=true noInternetAccessExpected=false mostRecentlyConnected=true 
lastConnected: 2026-05-23T15:37:48.117 

numRebootsSinceLastUse: 0
recentFailure: Association Rejection code: 0, last update time: 0
bssidAllowlist unset
vendorData unset
IsDppConfigurator: true
HasEncryptedPreSharedKey: false
 setWifi7Enabled=true
 mIsAllowedToUpdateByOtherUsers=true
 mCreatorUserId=0
}
mScanResultMatchInfoMapForCurrentUser={ScanResultMatchInfo: SSID: "AndroidWifi", from scan result: false, SecurityParams List:Security Parameters:
 Type: 0
 Enabled: true
 KeyMgmt: NONE
 Protocols: WPA RSN
 AuthAlgorithms:
 PairwiseCiphers:
 GroupCiphers:
 GroupMgmtCiphers:
 SuiteBCiphers:
 RequirePmf: false
 IsAddedByAutoUpgrade: false
 IsSaeH2eOnlyMode: false
 IsSaePkOnlyMode: false
Security Parameters:
 Type: 6
 Enabled: true
 KeyMgmt: OWE
 Protocols: RSN
 AuthAlgorithms:
 PairwiseCiphers: CCMP GCMP_256 GCMP_128
 GroupCiphers: CCMP GCMP_256 GCMP_128
 GroupMgmtCiphers:
 SuiteBCiphers:
 RequirePmf: true
 IsAddedByAutoUpgrade: true
 IsSaeH2eOnlyMode: false
 IsSaePkOnlyMode: false
=* ID: 0 SSID: "AndroidWifi" PROVIDER-NAME: null BSSID: null FQDN: null HOME-PROVIDER-NETWORK: false PRIO: 0 HIDDEN: false PMF: false CarrierId: -1 SubscriptionId: -1 SubscriptionGroup: null Currently Connected: true User Selected: false
 NetworkSelectionStatus NETWORK_SELECTION_ENABLED
 hasEverConnected: true
 hasNeverDetectedCaptivePortal: true
 hasEverValidatedInternetAccess: true
 mCandidateSecurityParams: Security Parameters:
 Type: 0
 Enabled: true
 KeyMgmt: NONE
 Protocols: WPA RSN
 AuthAlgorithms:
 PairwiseCiphers:
 GroupCiphers:
 GroupMgmtCiphers:
 SuiteBCiphers:
 RequirePmf: false
 IsAddedByAutoUpgrade: false
 IsSaeH2eOnlyMode: false
 IsSaePkOnlyMode: false
 mLastUsedSecurityParams: Security Parameters:
 Type: 0
 Enabled: true
 KeyMgmt: NONE
 Protocols: WPA RSN
 AuthAlgorithms:
 PairwiseCiphers:
 GroupCiphers:
 GroupMgmtCiphers:
 SuiteBCiphers:
 RequirePmf: false
 IsAddedByAutoUpgrade: false
 IsSaeH2eOnlyMode: false
 IsSaePkOnlyMode: false
 numAssociation 1
 validatedInternetAccess shared trusted
 macRandomizationSetting: 3
 mRandomizedMacAddress: 46:a0:89:b8:78:55
 randomizedMacExpirationTimeMs: 2026-05-24T15:37:47.928
 randomizedMacLastModifiedTimeMs: <none>
 persistentMacRandomizationSeed: 0
 mIsSendDhcpHostnameEnabled: true
 deletionPriority: 0
 KeyMgmt: NONE Protocols: WPA RSN
 AuthAlgorithms:
 PairwiseCiphers:
 GroupCiphers:
 GroupMgmtCiphers:
 SuiteBCiphers:
 PSK/SAE: 
SecurityParams List:
Security Parameters:
 Type: 0
 Enabled: true
 KeyMgmt: NONE
 Protocols: WPA RSN
 AuthAlgorithms:
 PairwiseCiphers:
 GroupCiphers:
 GroupMgmtCiphers:
 SuiteBCiphers:
 RequirePmf: false
 IsAddedByAutoUpgrade: false
 IsSaeH2eOnlyMode: false
 IsSaePkOnlyMode: false
Security Parameters:
 Type: 6
 Enabled: true
 KeyMgmt: OWE
 Protocols: RSN
 AuthAlgorithms:
 PairwiseCiphers: CCMP GCMP_256 GCMP_128
 GroupCiphers: CCMP GCMP_256 GCMP_128
 GroupMgmtCiphers:
 SuiteBCiphers:
 RequirePmf: true
 IsAddedByAutoUpgrade: true
 IsSaeH2eOnlyMode: false
 IsSaePkOnlyMode: false

Enterprise config:
 ocsp: 0
 trust_on_first_use: false
 user_approve_no_ca_cert: false
 selected_rcoi: 0
 minimum_tls_version: 0
 enable_conservative_peer_mode: false
 tofu_dialog_state: 0
 tofu_connection_state: 0
IP config:
IP assignment: DHCP
Proxy settings: NONE
 cuid=10190 cname=com.google.android.googlesdksetup luid=10190 lname=com.google.android.googlesdksetup lcuid=10190 allowAutojoin=true mAllowedAutoJoinInAdvancedProtection=true noInternetAccessExpected=false mostRecentlyConnected=true 
lastConnected: 2026-05-23T15:37:48.117 

numRebootsSinceLastUse: 0
recentFailure: Association Rejection code: 0, last update time: 0
bssidAllowlist unset
vendorData unset
IsDppConfigurator: true
HasEncryptedPreSharedKey: false
 setWifi7Enabled=true
 mIsAllowedToUpdateByOtherUsers=true
 mCreatorUserId=0
}
mCurrentUserId=0
WifiConfigManager - ConfigurationMap End ----
WifiConfigManager - Next network ID to be allocated 1
WifiConfigManager - Last selected network ID -1
WifiConfigManager - PNO scan frequency culling enabled = true
WifiConfigManager - PNO scan recency sorting enabled = true
Dump of WifiConfigStore
WifiConfigStore - Store File Begin ----
Name: /data/misc/apexdata/com.android.wifi/WifiConfigStore.xml, File Id: 0, Credentials encrypted: false
Name: /data/misc/apexdata/com.android.wifi/WifiConfigStoreSoftAp.xml, File Id: 1, Credentials encrypted: false
Name: /data/misc_ce/0/apexdata/com.android.wifi/WifiConfigStore.xml, File Id: 2, Credentials encrypted: false
Name: /data/misc_ce/0/apexdata/com.android.wifi/WifiConfigStoreNetworkSuggestions.xml, File Id: 3, Credentials encrypted: false
Name: /data/misc_ce/0/apexdata/com.android.wifi/WifiConfigStoreSoftAp.xml, File Id: 4, Credentials encrypted: false
Name: /data/misc_ce/0/apexdata/com.android.wifi/WifiConfigStoreAware.xml, File Id: 5, Credentials encrypted: false
WifiConfigStore - Store Data Begin ----
StoreData => Name: WifiCarrierInfoStoreManagerDataStores, File Id: 0, File Name: WifiConfigStore.xml
StoreData => Name: ImsiPrivacyProtectionExemptionMap, File Id: 2, File Name: WifiConfigStore.xml
StoreData => Name: PairingConfigManagerData, File Id: 5, File Name: WifiConfigStoreAware.xml
StoreData => Name: NetworkList, File Id: 0, File Name: WifiConfigStore.xml
StoreData => Name: NetworkList, File Id: 2, File Name: WifiConfigStore.xml
StoreData => Name: MacAddressMap, File Id: 0, File Name: WifiConfigStore.xml
StoreData => Name: Settings, File Id: 0, File Name: WifiConfigStore.xml
StoreData => Name: Settings, File Id: 2, File Name: WifiConfigStore.xml
StoreData => Name: NetworkSuggestionMap, File Id: 3, File Name: WifiConfigStoreNetworkSuggestions.xml
StoreData => Name: PasspointConfigData, File Id: 2, File Name: WifiConfigStore.xml
StoreData => Name: PasspointConfigData, File Id: 0, File Name: WifiConfigStore.xml
StoreData => Name: OpenNetworkNotifierBlacklistConfigData, File Id: 2, File Name: WifiConfigStore.xml
StoreData => Name: NetworkRequestMap, File Id: 2, File Name: WifiConfigStore.xml
StoreData => Name: SoftAp, File Id: 1, File Name: WifiConfigStoreSoftAp.xml
StoreData => Name: SoftAp, File Id: 4, File Name: WifiConfigStoreSoftAp.xml
StoreData => Name: WakeupConfigStoreData, File Id: 2, File Name: WifiConfigStore.xml
StoreData => Name: RoamingPolicies, File Id: 0, File Name: WifiConfigStore.xml
WifiConfigStore - Store Data End ----
WifiCarrierInfoManager: 
mImsiEncryptionInfoAvailable={}
mImsiPrivacyProtectionExemptionMap={}
mMergedCarrierNetworkOffloadMap<subId, enabled>={}
mUnmergedCarrierNetworkOffloadMap<subId, enabled>={}
mSubIdToSimInfoSparseArray={}
mActiveSubInfos=[[SubscriptionInfo: id=1 iccId=898603186[****] simSlotIndex=0 portIndex=0 isEmbedded=false carrierId=1 displayName=T-Mobile carrierName=T-Mobile isOpportunistic=false groupUuid=null groupOwner= isGroupDisabled=false displayNameSource=CARRIER iconTint=-9033797 number=[****] dataRoaming=0 mcc=310 mnc=260 ehplmns=[] hplmns=[] cardString=898603186[****] cardId=0 nativeAccessRules=null carrierConfigAccessRules=[cert: 92B5F8117FBD9BD5738FF168A4FA12CBE284BE834EDE1A7BB44DD8455BA15920 pkg: com.tmobile.rsuapp access: 0, cert: 92B5F8117FBD9BD5738FF168A4FA12CBE284BE834EDE1A7BB44DD8455BA15920 pkg: com.tmobile.echolocate access: 0, cert: 92B5F8117FBD9BD5738FF168A4FA12CBE284BE834EDE1A7BB44DD8455BA15920 pkg: com.tmobile.services.nameid access: 0, cert: 92B5F8117FBD9BD5738FF168A4FA12CBE284BE834EDE1A7BB44DD8455BA15920 pkg: com.tmobile.pr.mytmobile access: 0, cert: 92B5F8117FBD9BD5738FF168A4FA12CBE284BE834EDE1A7BB44DD8455BA15920 pkg: com.tmobile.vvm.application access: 0, cert: 3D1A4BEF6EE7AF7D34D120E7B1AAC0DD245585DE6237CF100F68333AFACFF562 pkg: com.tmobile.rsuapp access: 0, cert: 3D1A4BEF6EE7AF7D34D120E7B1AAC0DD245585DE6237CF100F68333AFACFF562 pkg: com.tmobile.echolocate access: 0, cert: 3D1A4BEF6EE7AF7D34D120E7B1AAC0DD245585DE6237CF100F68333AFACFF562 pkg: com.tmobile.services.nameid access: 0, cert: 3D1A4BEF6EE7AF7D34D120E7B1AAC0DD245585DE6237CF100F68333AFACFF562 pkg: com.tmobile.pr.mytmobile access: 0, cert: 3D1A4BEF6EE7AF7D34D120E7B1AAC0DD245585DE6237CF100F68333AFACFF562 pkg: com.tmobile.vvm.application access: 0, cert: 6892793FC413019D2DF609DFED7AF622D0F2D8FCF96EFA7E3FB87EEA34E10B93 pkg: com.tmobile.rsuapp access: 0, cert: 6892793FC413019D2DF609DFED7AF622D0F2D8FCF96EFA7E3FB87EEA34E10B93 pkg: com.tmobile.echolocate access: 0, cert: 6892793FC413019D2DF609DFED7AF622D0F2D8FCF96EFA7E3FB87EEA34E10B93 pkg: com.tmobile.services.nameid access: 0, cert: 6892793FC413019D2DF609DFED7AF622D0F2D8FCF96EFA7E3FB87EEA34E10B93 pkg: com.tmobile.pr.mytmobile access: 0, cert: 6892793FC413019D2DF609DFED7AF622D0F2D8FCF96EFA7E3FB87EEA34E10B93 pkg: com.tmobile.vvm.application access: 0, cert: 7B68FD9D4E7610C9CB35FC0C6CC06EA04C6906E3DFA9F48F9A05460AF36BFFFC pkg: com.tmobile.rsuapp access: 0, cert: 7B68FD9D4E7610C9CB35FC0C6CC06EA04C6906E3DFA9F48F9A05460AF36BFFFC pkg: com.tmobile.echolocate access: 0, cert: 7B68FD9D4E7610C9CB35FC0C6CC06EA04C6906E3DFA9F48F9A05460AF36BFFFC pkg: com.tmobile.services.nameid access: 0, cert: 7B68FD9D4E7610C9CB35FC0C6CC06EA04C6906E3DFA9F48F9A05460AF36BFFFC pkg: com.tmobile.pr.mytmobile access: 0, cert: 7B68FD9D4E7610C9CB35FC0C6CC06EA04C6906E3DFA9F48F9A05460AF36BFFFC pkg: com.tmobile.vvm.application access: 0] countryIso=us profileClass=-1 mType=LOCAL_SIM areUiccApplicationsEnabled=true usageSetting=DEFAULT isOnlyNonTerrestrialNetwork=false serviceCapabilities=[1, 2, 3] transferStatus=0 isSatelliteESOSSupported=false isPrivateNetwork=false streamingAppMaxDownlinkKbps=-1 streamingAppMaxUplinkKbps=-1]]
mCachedCarrierConfigPerSubId={1=PersistableBundle[{allow_hold_in_rtt_call_bool=true, imsvoice.voice_on_default_bearer_supported_bool=false, imsvoice.rtp_inactivity_time_threshold_millis_long=5000, imsss.ut_transport_type_int=1, imsemergency.emergency_call_setup_timer_on_current_network_sec_int=0, carrier_volte_provisioned_bool=false, support_cdma_1x_voice_calls_bool=true, support_ss_over_cdma_bool=false, disable_charge_indication_bool=false, ims.sip_timer_t4_millis_int=17000, support_tdscdma_bool=false, low_battery_alert_interval_int=-1, httpSocketTimeout=60000, allowed_initial_attach_apn_types_string_array=[ia, default], disable_supplementary_services_in_airplane_mode_bool=false, call_barring_default_service_class_int=1, call_redirection_service_component_name_string=null, carrier_wfc_supports_wifi_only_bool=true, carrier_default_actions_on_default_network_available_string_array=[false: 7, true: 8], imsvoice.minimum_session_expires_timer_sec_int=90, vonr_on_by_default_bool=true, iwlan_handover_policy_string_array=[source=GERAN|UTRAN|EUTRAN|NGRAN|IWLAN, target=GERAN|UTRAN|EUTRAN|NGRAN|IWLAN, type=allowed], operator_selection_expand_bool=true, disable_cdma_activation_code_bool=false, simplified_network_settings_bool=false, imsemergency.emergency_registration_timer_millis_int=10000, call_forwarding_when_unanswered_supported_bool=true, imsvt.h264_payload_description_bundle=PersistableBundle[mParcelledData.dataSize=172], support_tdscdma_roaming_networks_string_array=null, opportunistic.5g_data_switch_hysteresis_time_long=2000, ims.sip_timer_d_millis_int=130000, gps.lpp_profile=2, call_barring_visibility_bool=false, gba_mode_int=1, ims.registration_subscribe_expiry_timer_sec_int=600000, low_battery_alert_threshold_int=-1, gps.es_extension_sec=0, enable_apps_string_array=null, 5g_icon_display_secondary_grace_period_string=, carrier_default_wfc_ims_roaming_enabled_bool=true, cdma_nonroaming_networks_string_array=null, imsss.ut_as_server_fqdn_string=, imsemergency.emergency_lte_preferred_after_nr_failed_bool=false, signal_strength_nr_nsa_use_lte_as_primary_bool=true, aliasMinChars=2, fdn_number_length_limit_int=20, support_ims_call_forwarding_while_roaming_bool=true, wfc_data_spn_format_idx_int=0, hide_sim_lock_settings_bool=false, premium_capability_network_setup_time_millis_long=300000, satellite_esos_supported_bool=false, drop_video_call_when_answering_audio_call_bool=false, data_stall_recovery_timers_long_array=[180000, 180000, 180000, 180000], satellite_connection_hysteresis_sec_int=180, opportunistic_esim_download_via_wifi_only_bool=false, imsss.xcap_over_ut_supported_rats_int_array=[3, 5, 6], ci_action_on_sys_update_bool=false, carrier_certificate_string_array=[92B5F8117FBD9BD5738FF168A4FA12CBE284BE834EDE1A7BB44DD8455BA15920:com.tmobile.rsuapp,com.tmobile.echolocate,com.tmobile.services.nameid,com.tmobile.pr.mytmobile,com.tmobile.vvm.application, 3D1A4BEF6EE7AF7D34D120E7B1AAC0DD245585DE6237CF100F68333AFACFF562:com.tmobile.rsuapp,com.tmobile.echolocate,com.tmobile.services.nameid,com.tmobile.pr.mytmobile,com.tmobile.vvm.application, 6892793FC413019D2DF609DFED7AF622D0F2D8FCF96EFA7E3FB87EEA34E10B93:com.tmobile.rsuapp,com.tmobile.echolocate,com.tmobile.services.nameid,com.tmobile.pr.mytmobile,com.tmobile.vvm.application, 7B68FD9D4E7610C9CB35FC0C6CC06EA04C6906E3DFA9F48F9A05460AF36BFFFC:com.tmobile.rsuapp,com.tmobile.echolocate,com.tmobile.services.nameid,com.tmobile.pr.mytmobile,com.tmobile.vvm.application], convert_cdma_caller_id_mmi_codes_while_roaming_on_3gpp_bool=false, supports_customized_ringing_signal_bool=false, support_enhanced_call_blocking_bool=true, iwlan.ike_rekey_hard_timer_in_sec=14400, vvm_port_number_int=1808, premium_capability_notification_display_timeout_millis_long=1800000, wifi_connectivity_extend_cell_delay=-1, imsemergency.emergency_requires_ims_registration_bool=false, carrier_default_wfc_ims_enabled_bool=true, iwlan.ike_session_encryption_aes_ctr_key_size_int_array=[128, 192, 256], dial_string_replace_string_array=null, opportunistic_network_entry_threshold_bandwidth_int=1024, imsemergency.emergency_requires_volte_enabled_bool=false, opportunistic_network_data_switch_exit_hysteresis_time_long=3000, imsss.use_csfb_on_xcap_over_ut_failure_bool=true, gsm_nonroaming_networks_string_array=null, hide_radio_info_on_user_build_bool=false, imsserviceentitlement.ims_provisioning_bool=false, wifi.avoid_5ghz_softap_for_laa_bool=false, hide_ims_apn_bool=false, carrier_network_service_wlan_class_override_string=, vvm_destination_number_string=122, ims.ipsec_encryption_algorithms_int_array=[0, 1, 2], opportunistic.5g_backoff_time_long=10000, enable_cross_sim_calling_on_opportunistic_data_bool=false, vonr_setting_visibility_bool=true, imssms.sms_over_ims_supported_bool=true, dtmf_type_enabled_bool=false, ims.sip_timer_f_millis_int=128000, imsemergency.emergency_requires_vonr_enabled_bool=false, allow_video_calling_fallback_bool=true, nr_advanced_threshold_bandwidth_khz_int=0, imssms.sms_rp_cause_values_to_fallback_int_array=[1, 8, 10, 11, 21, 27, 28, 29, 30, 38, 42, 47, 50, 69, 81, 95, 96, 97, 98, 99, 111, 127], ngran_ssrsrq_hysteresis_db_int=2, ntn_lte_rssnr_thresholds_int_array=[-3, 1, 5, 13], eutran_rsrp_hysteresis_db_int=2, opportunistic.5g_ping_pong_time_long=60000, config_ims_package_override_string=com.google.android.ims, data_stall_recovery_should_skip_bool_array=[false, false, true, false, false], only_auto_select_in_home_network=false, iwlan.supported_prf_algorithms_int_array=[2, 4, 5, 6, 7], carrier_roaming_satellite_upsell_supported_bool=false, include_lte_for_nr_advanced_threshold_bandwidth_bool=false, show_data_connected_roaming_notification=false, gps.es_supl_data_plane_only_roaming_plmn_string_array=null, imsserviceentitlement.fcm_sender_id_string=, smsToMmsTextThreshold=-1, imsvt.video_as_bandwidth_kbps_int=960, opportunistic.exit_threshold_ss_rsrp_int_bundle=PersistableBundle[EMPTY_PARCEL], inflate_signal_strength_bool=false, iwlan.epdg_address_ip_type_preference_int=0, force_home_network_bool=false, spdi_override_string_array=null, show_wfc_location_privacy_policy_bool=false, imsemergency.refresh_geolocation_timeout_millis_int=5000, imsvt.video_rtp_dscp_int=40, gps.supl_es=1, enableSMSDeliveryReports=true, iwlan.natt_keep_alive_timer_sec_int=20, cellular_usage_setting_int=-1, missed_incoming_call_sms_originator_string_array=[], missed_incoming_call_sms_pattern_string_array=[], opportunistic.entry_threshold_ss_rsrp_int_bundle=PersistableBundle[EMPTY_PARCEL], auto_retry_enabled_bool=false, ims.enable_presence_capability_exchange_bool=true, opportunistic_network_entry_threshold_rssnr_int=5, play_call_recording_tone_bool=false, imsss.ut_terminal_based_services_int_array=[], carrier_force_disable_etws_cmas_test_bool=false, satellite_supported_disaster_plmn_string_array=[], call_waiting_service_class_int=1, allow_emergency_video_calls_bool=false, parameters_used_for_ntn_lte_signal_bar_int=1, support_add_conference_participants_bool=false, csp_enabled_bool=false, nr_timers_reset_on_endc_to_sa_transit_bool=false, default_sim_call_manager_string=, config_plans_package_override_string=, ims.ipv4_sip_mtu_size_cellular_int=1500, imsss.terminal_based_call_waiting_default_enabled_bool=true, ims.ims_pdn_enabled_in_no_vops_support_int_array=[], carrier_app_required_during_setup_bool=false, ims.sip_timer_h_millis_int=128000, call_forwarding_blocks_while_roaming_string_array=null, emergency_notification_delay_int=-1, support_video_conference_call_bool=false, ims_dtmf_tone_delay_int=0, wifi.hotspot_maximum_client_count=0, iwlan.supported_child_session_encryption_algorithms_int_array=[12], ims.registration_expiry_timer_sec_int=600000, ims.request_uri_type_int=0, ims.supported_rats_int_array=[6, 3, 5], imsvoice.amrwb_payload_description_bundle=PersistableBundle[mParcelledData.dataSize=164], imsemergency.emergency_scan_timer_sec_int=10, show_signal_strength_in_sim_status_bool=true, imsvt.video_rs_bandwidth_bps_int=8000, cellular_service_capabilities_int_array=[1, 2, 3], imsvt.video_qos_precondition_supported_bool=true, gba_ua_security_protocol_int=65536, data_rapid_notification_bool=true, display_no_data_notification_on_permanent_failure_bool=false, carrier_instant_lettering_escaped_chars_string=, carrier_default_wfc_ims_roaming_mode_int=2, bsf.bsf_transport_type_int=1, hide_enable_2g_bool=false, config_ims_mmtel_package_override_string=null, editable_wfc_mode_bool=true, imsvoice.audio_rr_bandwidth_bps_int=2000, imsvoice.audio_rtp_inactivity_timer_millis_int=20000, support_direct_fdn_dialing_bool=false, telephony_unsupported_network_capability_string_array=null, carrier_ut_provisioning_required_bool=false, disallow_adding_apn_string_array=null, imsvoice.conference_subscribe_type_int=1, imsemergency.cross_stack_redial_timer_sec_int=120, iwlan.diffie_hellman_groups_int_array=[2, 5, 14], opp_auto_data_switch_availability_switchback_millis_long=150000, treat_downgraded_video_calls_as_video_calls_bool=false, wifi_calls_can_be_hd_audio=true, support_downgrade_vt_to_audio_bool=true, wifi.suggestion_ssid_list_with_mac_randomization_disabled=[], use_default_ims_apn_when_absent_bool=true, allow_merging_rtt_calls_bool=false, satellite_information_redirect_url_string=, cdma_enhanced_roaming_indicator_for_home_network_int_array=[1], notify_handover_video_from_lte_to_wifi_bool=false, ims.sip_timer_j_millis_int=128000, use_otasp_for_provisioning_bool=false, always_play_remote_hold_tone_bool=false, opportunistic_network_data_switch_hysteresis_time_long=10000, telephony_data_handover_retry_rules_string_array=[retry_interval=1000|2000|4000|8000|16000, maximum_retries=5], ping_test_before_data_switch_bool=true, use_modem_display_network_type_bool=false, esim_download_retry_backoff_timer_sec_int=60, maxMessageTextSize=-1, duration_blocking_disabled_after_emergency_int=7200, default_vm_number_roaming_and_ims_unregistered_string=, imsvoice.srvcc_type_int_array=[0, 1, 2, 3], carrier_metered_roaming_apn_types_strings=[default, mms, dun, supl, enterprise], imssms.sms_over_ims_send_retry_delay_millis_int=2000, rtt_supported_while_roaming_bool=false, additional_settings_call_waiting_visibility_bool=true, carrier_ussd_method_int=0, show_video_call_charges_alert_dialog_bool=false, utran_rscp_hysteresis_db_int=2, carrier_network_service_wlan_package_override_string=, ims.enable_presence_publish_bool=false, supports_sdp_negotiation_of_d2d_rtp_header_extensions_bool=false, regional_satellite_earfcn_bundle=PersistableBundle[EMPTY_PARCEL], show_apn_setting_cdma_bool=false, use_usim_bool=false, carrier_instant_lettering_length_limit_int=64, premium_capability_notification_backoff_hysteresis_time_millis_long=1800000, imssms.sms_max_retry_over_ims_count_int=3, operator_name_filter_pattern_string=, carrier_service_name_array=[], opportunistic_carrier_ids_int_array=[0], iwlan.ike_local_id_type_int=3, data_switch_validation_min_gap_long=86400000, carrier_instant_lettering_invalid_chars_string=, ntn_5g_nr_ssrsrp_thresholds_int_array=[-110, -90, -80, -65], carrier_setup_app_string=, editable_wfc_roaming_mode_bool=false, supportMmsContentDisposition=true, support_manage_ims_conference_call_bool=true, supports_call_composer_bool=false, 5g_nr_ssrsrp_thresholds_int_array=[-110, -90, -80, -65], opportunistic_network_exit_threshold_rssnr_int=1, imsemergency.ims_reasoninfo_code_to_retry_emergency_int_array=[], allow_non_emergency_calls_in_ecm_bool=true, ims.sa_disable_policy_int=0, allow_cdma_eri_bool=false, satellite_entitlement_status_refresh_days_int=1, opportunistic.5g_data_switch_exit_hysteresis_time_long_bundle=PersistableBundle[EMPTY_PARCEL], carrier_data_service_wlan_package_override_string=, show_carrier_data_icon_pattern_string=, allow_hold_call_during_emergency_bool=true, use_rcs_presence_bool=false, opportunistic.entry_threshold_ss_rsrp_int=-111, call_forwarding_when_busy_supported_bool=true, config_telephony_use_own_number_for_voicemail_bool=false, cdma_dtmf_tone_delay_int=100, editable_voicemail_number_bool=false, carrier_data_call_apn_retry_after_disconnect_long=3000, filtered_cnap_names_string_array=null, enabledTransID=false, aliasEnabled=false, imsemergency.emergency_qos_precondition_supported_bool=true, parameters_use_for_ntn_5g_nr_signal_bar_int=1, satellite_entitlement_supported_bool=false, imsrtt.text_qos_precondition_supported_bool=true, imsss.ut_supported_when_roaming_bool=true, satellite_supported_emergency_plmn_string_array=[], carrier_supports_opp_data_auto_provisioning_bool=false, imsvoice.session_expires_timer_sec_int=1800, bandwidth_string_array=[GPRS:24,24, EDGE:70,18, UMTS:115,115, CDMA:14,14, 1xRTT:30,30, EvDo_0:750,48, EvDo_A:950,550, HSDPA:4300,620, HSUPA:4300,1800, HSPA:4300,1800, EvDo_B:1500,550, eHRPD:750,48, iDEN:14,14, LTE:30000,15000, HSPA+:13000,3400, GSM:24,24, TD_SCDMA:115,115, LTE_CA:30000,15000, NR_NSA:47000,18000, NR_NSA_MMWAVE:145000,60000, NR_SA:145000,60000, NR_SA_MMWAVE:145000,60000], imssms.sms_rp_cause_values_to_retry_over_ims_int_array=[41], iwlan.supported_ike_session_aead_algorithms_int_array=[], vvm_disabled_capabilities_string_array=null, ngran_ssrsrp_hysteresis_db_int=2, satellite_configs_per_plmn_bundle=PersistableBundle[EMPTY_PARCEL], show_4g_for_lte_data_icon_bool=false, check_pricing_with_carrier_data_roaming_bool=false, carrier_roaming_satellite_upsell_notification_maximum_daily_count_int=5, ims.rcs_feature_tag_allowed_string_array=[+g.3gpp.icsi-ref="urn%3Aurn-7%3A3gpp-service.ims.icsi.oma.cpm.msg", +g.3gpp.icsi-ref="urn%3Aurn-7%3A3gpp-service.ims.icsi.oma.cpm.largemsg", +g.3gpp.icsi-ref="urn%3Aurn-7%3A3gpp-service.ims.icsi.oma.cpm.deferred", +g.gsma.rcs.cpm.pager-large, +g.3gpp.icsi-ref="urn%3Aurn-7%3A3gpp-service.ims.icsi.oma.cpm.session", +g.3gpp.icsi-ref="urn%3Aurn-7%3A3gpp-service.ims.icsi.oma.cpm.filetransfer", +g.3gpp.iari-ref="urn%3Aurn-7%3A3gpp-application.ims.iari.rcs.fthttp", +g.3gpp.iari-ref="urn%3Aurn-7%3A3gpp-application.ims.iari.rcs.ftsms", +g.3gpp.iari-ref="urn%3Aurn-7%3A3gpp-service.ims.icsi.gsma.callcomposer", +g.gsma.callcomposer, +g.3gpp.icsi-ref="urn%3Aurn-7%3A3gpp-service.ims.icsi.gsma.callunanswered", +g.3gpp.icsi-ref="urn%3Aurn-7%3A3gpp-service.ims.icsi.gsma.sharedmap", +g.3gpp.icsi-ref="urn%3Aurn-7%3A3gpp-service.ims.icsi.gsma.sharedsketch", +g.3gpp.iari-ref="urn%3Aurn-7%3A3gpp-application.ims.iari.rcs.geopush", +g.3gpp.iari-ref="urn%3Aurn-7%3A3gpp-application.ims.iari.rcs.geosms", +g.3gpp.iari-ref="urn%3Aurn-7%3A3gpp-application.ims.iari.rcs.chatbot", +g.3gpp.iari-ref="urn%3Aurn-7%3A3gpp-application.ims.iari.rcs.chatbot.sa", +g.gsma.rcs.botversion="#=1,#=2", +g.gsma.rcs.cpimext], call_barring_supports_deactivate_all_bool=true, show_avoid_bad_wifi_bool=false, ims_conference_size_limit_int=5, satellite_sos_max_datagram_size_bytes_int=255, imsss.ut_as_server_port_int=80, apn.settings_default_protocol_string=, carrier_data_service_wlan_class_override_string=, identify_high_definition_calls_in_call_log_bool=false, ims.enable_presence_group_subscribe_bool=false, call_barring_supports_password_change_bool=true, imsvoice.multiendpoint_supported_bool=false, supportHttpCharsetHeader=false, ngran_sssinr_hysteresis_db_int=2, allow_hold_video_call_bool=true, ims.sip_server_port_number_int=5060, ims.wifi_off_deferring_time_millis_int=4000, iwlan.supported_integrity_algorithms_int_array=[5, 2, 12, 13, 14], vvm_prefetch_bool=true, show_5g_slice_icon_bool=true, hide_enhanced_4g_lte_bool=false, default_rtt_mode_int=0, ims.allow_non_global_phone_number_format_bool=false, imsemergency.emergency_over_cs_supported_access_network_types_int_array=[2, 1], opp_auto_data_switch_performance_stability_millis_long=120000, satellite_supported_msg_apps_string_array=[com.google.android.apps.messaging], volte_5g_limited_alert_dialog_bool=false, data_connected_roaming_notification_excluded_mccs_string_array=[310, 311, 312, 313, 314, 315, 316], carrier_qualified_networks_service_class_override_string=, display_hd_audio_property_bool=true, ims.non_rcs_capabilities_cache_expiration_sec_int=2592000, has_in_call_noise_suppression_bool=false, enhanced_4g_lte_on_by_default_bool=true, data_limit_notification_bool=true, imssms.sms_over_ims_format_int=0, config_show_orig_dial_string_for_cdma=false, voicemail_notification_persistent_bool=false, carrier_default_actions_on_reset_string_array=[6, 8], opportunistic_network_max_backoff_time_long=60000, world_phone_bool=false, ims_reasoninfo_mapping_string_array=null, imsemergency.emergency_domain_preference_int_array=[2, 1, 3], call_screening_app=, disconnect_cause_play_busytone_int_array=[], gps.a_glonass_pos_protocol_select=0, default_vm_number_string=, proxy_connectivity_delay_cell=-1, data_limit_threshold_bytes_long=-1, imsemergency.maximum_number_of_emergency_tries_over_vowifi_int=1, imsvoice.session_refresh_method_int=1, premium_capability_maximum_monthly_notification_count_int=10, imsvt.video_rr_bandwidth_bps_int=6000, use_acs_for_rcs_bool=false, esim_max_download_retry_attempts_int=5, support_ims_conference_event_package_bool=true, allowAttachAudio=true, use_caller_id_ussd_bool=false, show_call_blocking_disabled_notification_always_bool=false, carrier_service_number_array=[], ci_action_on_sys_update_extra_string=, imssms.emergency_sms_over_emergency_pdn_int_array=[], rtt_upgrade_supported_for_downgraded_vt_call=true, ignore_sim_network_locked_events_bool=false, cdma_home_registered_plmn_name_override_bool=false, iwlan.supported_child_session_aead_algorithms_int_array=[], notify_international_call_on_wfc_bool=false, use_ip_for_calling_indicator_bool=false, iwlan.add_ke_to_child_session_rekey_bool=false, call_forwarding_map_non_number_to_voicemail_bool=false, imsss.ut_iptype_roaming_int=2, carrier_instant_lettering_available_bool=false, satellite_connected_notification_throttle_millis_int=604800000, premium_capability_purchase_condition_backoff_hysteresis_time_millis_long=1800000, nr_advanced_capable_pco_id_int=0, opportunistic.exit_threshold_ss_rsrq_double_bundle=PersistableBundle[EMPTY_PARCEL], data_switch_validation_timeout_long=5000, imsvoice.session_privacy_type_int=0, show_spn_for_home_in_choose_network_setting_bool=false, lte_rsrq_thresholds_int_array=[-20, -17, -14, -11], iwlan.dpd_timer_sec_int=120, stk_disable_launch_browser_bool=false, show_ims_registration_status_bool=false, carrier_use_ims_first_for_emergency_bool=true, carrier_vt_available_bool=true, imsvoice.rtp_packet_loss_rate_threshold_int=40, nr_advanced_pci_change_secondary_timer_seconds_int=0, smsToMmsTextLengthThreshold=-1, display_voicemail_number_as_default_call_forwarding_number=false, config_wifi_disable_in_ecbm=false, eutran_rssnr_hysteresis_db_int=2, ci_action_on_sys_update_intent_string=, ims.subscribe_retry_duration_millis_long=-1, bsf.bsf_server_fqdn_string=, always_show_data_rat_icon_bool=false, ascii_7_bit_support_for_long_message_bool=false, additional_call_setting_bool=true, enabled_4g_opportunistic_network_scan_bool=true, ntn_lte_rsrq_thresholds_int_array=[-20, -17, -14, -11], editable_voicemail_number_setting_bool=true, carrier_app_no_wake_signal_config=null, sms_requires_destination_number_conversion_bool=false, imsss.ut_server_based_services_int_array=[0, 1, 2, 6, 3, 4, 5, 7, 8, 9, 10, 11, 12, 13, 17, 14, 15, 16, 18, 19, 20, 21], carrier_default_actions_on_redirection_string_array=[9, 4, 1], support_swap_after_merge_bool=true, recipientLimit=2147483647, support_phone_number_source_ts43_bool=false, smdp_server_address_string=, imsvt.video_codec_capability_payload_types_bundle=PersistableBundle[mParcelledData.dataSize=92], opp_auto_data_switch_validation_max_retries_int=7, format_incoming_number_to_national_for_jp_bool=false, imsserviceentitlement.show_vowifi_webview_bool=false, ims.gruu_enabled_bool=false, store_sim_pin_for_unattended_reboot_bool=true, call_waiting_over_ut_warning_bool=false, hide_digits_helper_text_on_stk_input_screen_bool=true, notify_vt_handover_to_wifi_failure_bool=false, carrier_supports_caller_id_vertical_service_codes_bool=false, 5g_watchdog_time_ms_long=3600000, gsm_dtmf_tone_delay_int=0, imsvoice.ringback_timer_millis_int=90000, show_onscreen_dial_button_bool=true, support_multi_anchor_conf_bool=false, imssms.sms_csfb_retry_on_failure_bool=true, uaProfTagName=x-wap-profile, sim_network_unlock_allow_dismiss_bool=true, carrier_roaming_satellite_purchase_mode_timeout_sec_int=300, satellite_display_name_string=, show_carrier_id_in_sim_status_bool=false, hide_preset_apn_details_bool=false, imsemergency.start_quick_cross_stack_redial_timer_when_registered_bool=true, cdma_3waycall_flash_delay_int=0, carrier_settings_activity_component_name_string=, carrier_default_wfc_ims_mode_int=1, iwlan.child_session_aes_cbc_key_size_int_array=[128, 192, 256], call_forwarding_visibility_bool=true, enable_eap_method_prefix_bool=false, imsserviceentitlement.skip_wfc_activation_bool=false, force_imei_bool=false, carrier_wfc_ims_available_bool=true, enable_nr_advanced_for_roaming_bool=false, display_call_strength_indicator_bool=true, imsss.ut_iptype_home_int=2, aliasMaxChars=48, imsvoice.audio_codec_capability_payload_types_bundle=PersistableBundle[mParcelledData.dataSize=388], satellite_roaming_p2p_sms_inactivity_timeout_sec_int=180, opportunistic.5g_data_switch_hysteresis_time_long_bundle=PersistableBundle[EMPTY_PARCEL], hide_preferred_network_type_bool=false, support_emergency_sms_over_ims_bool=false, cdma_roaming_networks_string_array=null, imsemergency.emergency_over_ims_roaming_supported_3gpp_network_types_int_array=[3], supports_device_to_device_communication_using_dtmf_bool=false, imsss.network_initiated_ussd_over_ims_supported_bool=true, maxImageHeight=1944, iwlan.supports_eap_aka_fast_reauth_bool=false, imsemergency.retry_emergency_on_ims_pdn_bool=false, gsm_roaming_networks_string_array=null, satellite_roaming_turn_off_session_for_emergency_call_bool=true, imsrtt.text_rs_bandwidth_bps_int=100, gps.enable_ni_supl_message_injection_bool=false, carrier_roaming_satellite_upsell_notification_hysteresis_sec_int=900, opportunistic_network_ping_pong_time_long=60000, ims.publish_service_desc_feature_tag_map_override_string_array=[], imsemergency.emergency_over_ims_supported_3gpp_network_types_int_array=[3], gps.persist_lpp_mode_bool=true, data_warning_threshold_bytes_long=-1, supports_device_to_device_communication_using_rtp_bool=false, iwlan.epdg_plmn_priority_int_array=[0, 1, 2], roaming_operator_string_array=null, imsemergency.scan_limited_service_after_volte_failure_bool=false, emergency_messaging_supported_bool=false, carrier_roaming_satellite_emergency_messaging_provider_per_country_bundle=PersistableBundle[EMPTY_PARCEL], unmetered_network_types_string_array=[NR_NSA, NR_NSA_MMWAVE, NR_SA, NR_SA_MMWAVE], supports_video_back_tone_bool=false, lte_endc_using_user_data_for_rrc_detection_bool=false, ims.ims_single_registration_required_bool=false, gps.es_supl_control_plane_support_int=0, imsemergency.maximum_cellular_search_timer_sec_int=0, wfc_flight_mode_spn_format_idx_int=-1, imsrtt.text_on_default_bearer_supported_bool=false, satellite_data_support_mode_int=0, iwlan.epdg_pco_id_ipv4_int=0, iwlan.epdg_pco_id_ipv6_int=0, imssms.sms_max_retry_count_int=3, limited_sim_function_notification_for_dsds_bool=false, ims.rcs_requires_provisioning_bundle=PersistableBundle[EMPTY_PARCEL], enableMMSDeliveryReports=false, carrier_vvm_package_name_string=com.tmobile.vvm.application, satellite_ignore_data_roaming_setting_bool=true, disable_voice_barring_notification_bool=false, world_mode_enabled_bool=false, parameters_use_for_5g_nr_signal_bar_int=1, gsm_rssi_thresholds_int_array=[-107, -103, -97, -89], premium_capability_supported_on_lte_bool=false, carrier_default_data_roaming_enabled_bool=false, imsvoice.audio_as_bandwidth_kbps_int=41, ratchet_rat_families=[1,2, 7,8,12, 3,11,9,10,15, 14,19], opportunistic.exit_threshold_ss_rsrq_double=-18.5, ims.keep_pdn_up_in_no_vops_bool=false, ims.rcs_bulk_capability_exchange_bool=false, supported_premium_capabilities_int_array=[], carrier_volte_available_bool=true, lte_rssnr_thresholds_int_array=[-3, 1, 5, 13], show_gid1_in_sim_status_bool=false, wfc_operator_error_codes_string_array=[REG09|0], iwlan.supported_ike_session_encryption_algorithms_int_array=[12], undelivered_sms_message_expiration_time=604800000, allow_local_dtmf_tones_bool=true, apn_matched_required=true, ims.phone_context_domain_name_string=, imsi_carrier_public_key_wlan_string=null, ims.ipsec_authentication_algorithms_int_array=[0, 1], support_conference_call_bool=true, show_forwarded_number_bool=false, iwlan.supports_ike_session_multiple_sa_proposals_bool=false, supports_business_call_composer_bool=false, log_calls_answered_elsewhere_bool=true, opportunistic.5g_max_backoff_time_long=60000, geran_rssi_hysteresis_db_int=2, gba_ua_security_organization_int=1, ims.geolocation_pidf_in_sip_register_support_int_array=[2], opportunistic.entry_threshold_ss_rsrq_double_bundle=PersistableBundle[EMPTY_PARCEL], ignore_rtt_mode_setting_bool=true, carrier_roaming_ntn_connect_type_int=0, wifi.avoid_5ghz_wifi_direct_for_laa_bool=false, iwlan.epdg_authentication_method_int=0, carrier_supports_tethering_bool=true, support_pause_ims_video_calls_bool=false, vvm_legacy_mode_enabled_bool=false, parameters_used_for_lte_signal_bar_int=1, imssms.sms_tr2_timer_millis_int=15000, is_private_network_bool=false, iwlan.ike_remote_id_type_int=2, userAgent=, mms_network_release_timeout_millis_int=5000, imsrtt.text_as_bandwidth_kbps_int=4, satellite_attach_supported_bool=false, enhanced_4g_lte_title_variant_bool=false, mmsCloseConnection=true, gps.supl_host=supl.google.com, gps.supl_mode=1, gps.supl_port=7275, show_wifi_calling_icon_in_status_bar_bool=false, pnn_override_string_array=null, cross_sim_spn_format_int=1, support_ims_conference_call_bool=true, config_cellBroadcastAppLinks=true, allow_add_call_during_video_call=true, skip_cf_fail_to_disable_dialog_bool=false, iwlan.retransmit_timer_sec_int_array=[500, 1000, 2000, 4000, 8000], wcdma_default_signal_strength_measurement_string=rssi, preferred_ike_protocol_int=-1, carrier_volte_provisioning_required_bool=false, use_wfc_home_network_mode_in_roaming_network_bool=false, imsemergency.quick_cross_stack_redial_timer_sec_int=0, monthly_data_cycle_day_int=-1, ims.ipv6_sip_mtu_size_cellular_int=1500, ci_action_on_sys_update_extra_val_string=, no_reply_timer_for_cfnry_sec_int=20, rtt_upgrade_supported_bool=false, carrier_roaming_satellite_emergency_messaging_redirection_destination_string=, apn_expand_bool=true, imsrtt.text_codec_capability_payload_types_bundle=PersistableBundle[mParcelledData.dataSize=144], satellite_upsell_notification_throttle_hours_int=24, imsi_key_availability_int=0, carrier_vvm_package_name_string_array=null, show_4g_for_3g_data_icon_bool=false, carrier_allow_turnoff_ims_bool=true, opportunistic_network_entry_threshold_rsrp_int=-108, nrarfcns_rsrp_boost_int_array=null, iwlan.ike_session_encryption_aes_cbc_key_size_int_array=[128, 192, 256], capabilities_exempt_from_single_dc_check_int_array=[4], wfc_spn_use_root_locale=false, imsi_key_download_url_string=null, carrier_provisioning_app_string=, allow_adding_apns_bool=true, switch_data_to_primary_if_primary_is_oos_bool=true, maxSubjectLength=40, support_adhoc_conference_calls_bool=false, ims.rcs_request_forbidden_by_sip_489_bool=false, satellite_nidd_apn_name_string=, ehplmn_override_string_array=null, carrier_supported_satellite_services_per_provider_bundle=PersistableBundle[mParcelledData.dataSize=44], tty_supported_bool=true, imsvoice.rtp_jitter_threshold_millis_int=120, support_no_reply_timer_for_cfnry_bool=true, vvm_type_string=vvm_type_cvvm, boosted_lte_earfcns_string_array=null, carrier_default_redirection_url_string_array=null, iwlan.child_session_aes_gcm_key_size_int_array=[], use_call_waiting_ussd_bool=false, is_ims_conference_size_enforced_bool=false, override_wfc_roaming_mode_while_using_ntn_bool=true, mms_max_ntn_payload_size_bytes_int=-1, enable_dialer_key_vibration_bool=true, opportunistic_network_entry_or_exit_hysteresis_time_long=10000, data_connected_roaming_notification_included_mcc_mncs_string_array=[310032, 310033, 310110, 310140, 310370, 310400, 310470, 310500, 310970, 311170, 311250], enhanced_4g_lte_title_variant_int=0, unthrottle_data_retry_when_tac_changes_bool=false, time_to_switch_back_to_primary_if_opportunistic_oos_long=60000, satellite_roaming_esos_inactivity_timeout_sec_int=600, network_temp_not_metered_supported_bool=true, imsvoice.session_refresher_type_int=1, gps.use_emergency_pdn_for_emergency_supl=1, imsrtt.text_rr_bandwidth_bps_int=300, read_only_apn_fields_string_array=null, mmi_two_digit_number_pattern_string_array=[], broadcast_emergency_call_state_changes_bool=false, imsemergency.emergency_cdma_preferred_numbers_string_array=[], nr_timers_reset_if_non_endc_and_rrc_idle_bool=false, carrier_supported_satellite_notification_hysteresis_sec_int=180, network_notification_delay_int=-1, prefer_in_service_sim_for_normal_routed_emergency_calls_bool=false, vt_upgrade_supported_for_downgraded_rtt_call=true, cdma_roaming_mode_int=-1, iwlan.mcc_mncs_string_array=[], carrier_volte_tty_supported_bool=false, show_precise_failed_cause_bool=false, gsm_cdma_calls_can_be_hd_audio=false, wfc_carrier_name_override_by_pnn_bool=true, carrier_network_service_wwan_package_override_string=, 4g_only_bool=false, video_calls_can_be_hd_audio=true, httpParams=, prefer_3g_visibility_bool=true, imsvoice.audio_inactivity_call_end_reasons_int_array=[1, 2, 3, 0], rtt_supported_bool=true, enableGroupMms=true, imsvoice.carrier_volte_roaming_available_bool=true, hide_lte_plus_data_icon_bool=true, support_emergency_dialer_shortcut_bool=true, subscription_group_uuid_string=, auto_unhold_on_remote_disconnect_bool=false, carrier_cross_sim_ims_available_bool=false, carrier_default_2g_protection_enabled_bool=false, call_barring_over_ut_warning_bool=false, restart_radio_on_pdp_fail_regular_deactivation_bool=false, rtt_auto_upgrade_bool=false, imsemergency.emergency_domain_preference_roaming_int_array=[2, 1, 3], boosted_nrarfcns_string_array=null, carrier_config_applied_bool=true, imsvoice.amrnb_payload_description_bundle=PersistableBundle[mParcelledData.dataSize=164], additional_nr_advanced_bands_int_array=[41], carrier_ims_gba_required_bool=true, lte_enabled_bool=true, show_cdma_choices_bool=false, show_roaming_indicator_bool=false, apn.settings_default_roaming_protocol_string=, ims.registration_event_package_supported_bool=true, carrier_provisions_wifi_merged_networks_bool=false, allow_video_call_in_low_battery_bool=true, show_blocking_pay_phone_option_bool=false, use_rcs_sip_options_bool=false, wcdma_rscp_thresholds_int_array=[-115, -105, -95, -85], gps.supl_ver=0x20000, carrier_data_service_wwan_package_override_string=, wfc_spn_format_idx_int=1, voice_privacy_disable_ui_bool=false, lte_plus_threshold_bandwidth_khz_int=20000, carrier_settings_enable_bool=false, imsserviceentitlement.default_service_entitlement_status_bool=false, opl_override_opl_string_array=null, imsemergency.prefer_ims_emergency_when_voice_calls_on_cs_bool=false, carrier_roaming_ntn_emergency_call_to_satellite_handover_type_int=2, opp_auto_data_switch_availability_stability_millis_long=10000, data_warning_notification_bool=true, local_disconnect_empty_ims_conference_bool=false, ims.registration_retry_max_timer_millis_int=1800000, imswfc.emergency_call_over_emergency_pdn_bool=false, cdma_home_registered_plmn_name_string=, premium_capability_purchase_url_string=null, emailGatewayNumber=, gba_ua_tls_cipher_suite_int=47, utran_ecno_hysteresis_db_int=2, opportunistic.exit_threshold_ss_rsrp_int=-120, mdn_is_additional_voicemail_number_bool=false, carrier_network_service_wwan_class_override_string=, opportunistic.entry_threshold_ss_rsrq_double=-18.5, carrier_roaming_satellite_default_services_int_array=[2, 3, 6], imsvoice.ringing_timer_millis_int=90000, allow_metered_network_for_cert_download_bool=false, telephony_network_capability_priorities_string_array=[eims:90, supl:80, mms:70, xcap:70, cbs:50, mcx:50, fota:50, ims:40, rcs:40, dun:30, enterprise:20, internet:20, prioritize_bandwidth:20, prioritize_latency:20, prioritize_unified_communications:20], imsserviceentitlement.entitlement_server_url_string=, ims.mmtel_requires_provisioning_bundle=PersistableBundle[EMPTY_PARCEL], call_forwarding_when_unreachable_supported_bool=true, ntn_5g_nr_ssrsrq_thresholds_int_array=[-31, -19, -7, 6], iwlan.child_sa_rekey_soft_timer_sec_int=3600, imsemergency.emergency_over_ims_supported_rats_int_array=[3, 5], unloggable_numbers_string_array=null, 5g_nr_ssrsrq_thresholds_int_array=[-31, -19, -7, 6], only_single_dc_allowed_int_array=[4, 7, 5, 6, 12], opp_auto_data_switch_ping_before_switch_bool=true, imsi_carrier_public_key_epdg_string=null, ims.ims_user_agent_string=#MANUFACTURER#_#MODEL#_Android#AV#_#BUILD#, default_preferred_apn_name_string=, ims.geolocation_pidf_in_sip_invite_support_int_array=[2], imsemergency.emergency_network_scan_type_int=0, support_clir_network_default_bool=true, carrier_supports_ss_over_ut_bool=true, always_show_emergency_alert_onoff_bool=false, iwlan.max_retries_int=3, imssms.sms_tr1_timer_millis_int=130000, imsemergency.emergency_vowifi_requires_condition_int=0, imsss.ut_requires_ims_registration_bool=false, imsserviceentitlement.entitlement_version_int=2, emergency_sms_mode_timer_ms_int=0, nr_timers_reset_on_plmn_change_bool=false, bsf.bsf_server_port_int=80, avoid_bad_wifi_bool=true, auto_data_switch_rat_signal_score_string_bundle=PersistableBundle[mParcelledData.dataSize=1024], emergency_number_prefix_string_array=[], iwlan.ike_session_encryption_aes_gcm_key_size_int_array=[], carrier_volte_override_wfc_provisioning_bool=false, ntn_5g_nr_sssinr_thresholds_int_array=[-5, 5, 15, 30], 5g_icon_configuration_string=connected_mmwave:5G_Plus,connected:5G,not_restricted_rrc_idle:5G,not_restricted_rrc_con:5G, imsemergency.emergency_callback_mode_supported_bool=false, default_vm_number_roaming_string=, 5g_nr_sssinr_thresholds_int_array=[-5, 5, 15, 30], rtt_downgrade_supported_bool=false, default_mtu_int=1440, caller_id_over_ut_warning_bool=false, satellite_entitlement_app_name_string=androidSatmode, imsvoice.session_timer_supported_bool=true, use_hfa_for_provisioning_bool=false, carrier_eri_file_name_string=eri.xml, imssms.sms_over_ims_supported_rats_int_array=[3, 5], opportunistic_time_to_scan_after_capability_switch_to_primary_long=120000, delay_ims_tear_down_until_call_end_bool=false, international_roaming_dial_string_replace_string_array=null, carrier_allow_deflect_ims_call_bool=false, maxMessageSize=1048576, gps.nfw_proxy_apps=, vvm_client_prefix_string=//VVM, disable_dun_apn_while_roaming_with_preset_apn_bool=false, config_ims_rcs_package_override_string=null, hide_carrier_network_settings_bool=false, gps.gps_lock=3, ims.sip_preferred_transport_int=2, imsvoice.include_caller_id_service_codes_in_sip_invite_bool=false, naiSuffix=, data_stall_recovery_timers_randomization_millis_long_array=[0, 0, 0, 0], carrier_rcs_provisioning_required_bool=true, support_ims_conference_event_package_on_peer_bool=true, editable_enhanced_4g_lte_bool=true, carrier_metered_apn_types_strings=[default, mms, dun, supl, enterprise], ims.sa_disable_policy_for_emergency_int=0, non_roaming_operator_string_array=null, spn_display_condition_override_int=-1, emergency_call_to_satellite_t911_handover_timeout_millis_int=30000, opportunistic_network_backoff_time_long=10000, ims.rcs_request_retry_interval_millis_long=1200000, iwlan.epdg_address_priority_int_array=[1, 0], premium_capability_maximum_daily_notification_count_int=2, allow_merge_wifi_calls_when_vowifi_off_bool=true, vilte_data_is_metered_bool=true, spn_display_rule_use_roaming_from_service_state_bool=false, satellite_roaming_p2p_sms_supported_bool=false, use_only_rsrp_for_lte_signal_bar_bool=false, show_4glte_for_lte_data_icon_bool=false, lte_rsrp_thresholds_int_array=[-128, -118, -108, -98], telephony_data_setup_retry_rules_string_array=[capabilities=eims, retry_interval=1000, maximum_retries=20, permanent_fail_causes=8|27|28|29|30|32|33|35|50|51|111|-5|-6|65537|65538|-3|65543|65547|2252|2253|2254, retry_interval=2500, capabilities=mms|supl|cbs|rcs, retry_interval=2000, capabilities=internet|enterprise|dun|ims|fota|xcap|mcx|prioritize_bandwidth|prioritize_latency|prioritize_unified_communications, retry_interval=2500|3000|5000|10000|15000|20000|40000|60000|120000|240000|600000|1200000|1800000, maximum_retries=20], supports_unidirectional_video_service_bool=false, rtt_supported_for_vt_bool=false, use_only_dialed_sim_ecc_list_bool=false, carrier_nr_availabilities_int_array=[1, 2], enableMMSReadReports=false, iwlan.child_sa_rekey_hard_timer_sec_int=7200, allow_hold_in_ims_call=true, bandwidth_nr_nsa_use_lte_value_for_uplink_bool=false, imswfc.pidf_short_code_string_array=[], nr_timers_reset_on_voice_qos_bool=false, prefer_2g_bool=false, show_single_operator_row_in_choose_network_setting_bool=true, ntn_lte_rsrp_thresholds_int_array=[-128, -118, -108, -98], is_opportunistic_subscription_bool=false, iwlan.supports_child_session_multiple_sa_proposals_bool=false, call_forwarding_over_ut_warning_bool=false, carrier_vowifi_tty_supported_bool=false, maxImageWidth=2592, iwlan.ike_rekey_soft_timer_sec_int=7200, smart_forwarding_config_component_name_string=, ratchet_nr_advanced_bandwidth_if_rrc_idle_bool=true, remove_satellite_plmn_in_manual_network_scan_bool=true, eutran_rsrq_hysteresis_db_int=2, show_iccid_in_sim_status_bool=false, nr_advanced_bands_secondary_timer_seconds_int=0, lte_earfcns_rsrp_boost_int=0, call_composer_picture_server_url_string=https://ue.fcs.mstore.msg.t-mobile.com/restclient/V1/FCS/Upload, vonr_enabled_bool=false, carrier_promote_wfc_on_call_fail_bool=false, opportunistic_network_exit_threshold_rsrp_int=-118, vvm_ssl_enabled_bool=false, iwlan.child_session_aes_ctr_key_size_int_array=[128, 192, 256], carrier_roaming_satellite_t911_to_esos_handover_supported_bool=false, radio_restart_failure_causes_int_array=[], hide_voicemail_number_setting_bool=false, show_vowifi_drop_dialog_on_dsds_bool=false, wfc_emergency_address_carrier_app_string=, ims.use_sip_uri_for_presence_subscribe_bool=false, imsss.ut_supported_when_ps_data_off_bool=true, support_wps_over_ims_bool=true, ims.sip_timer_t1_millis_int=2000, volte_replacement_rat_int=3, carrier_config_version_string=, ims.registration_retry_base_timer_millis_int=30000, imsvt.video_rtcp_inactivity_timer_millis_int=0, satellite_technology_type_int_array=[], imsvoice.conference_factory_uri_string=, 5g_icon_display_grace_period_string=connected_mmwave,legacy,30;connected,legacy,30;not_restricted_rrc_idle,legacy,30;not_restricted_rrc_con,legacy,30, wcdma_ecno_thresholds_int_array=[-24, -14, -6, 1], enable_carrier_display_name_resolver_bool=false, roaming_unmetered_network_types_string_array=[], ignore_data_enabled_changed_for_video_calls=true, iwlan.epdg_static_address_roaming_string=, allow_emergency_numbers_in_call_log_bool=false, ims.sip_over_ipsec_enabled_bool=true, uaProfUrl=, enabledMMS=false, imsvoice.audio_rtcp_inactivity_timer_millis_int=20000, opportunistic.5g_data_switch_exit_hysteresis_time_long=2000, imsvoice.mo_call_request_timeout_millis_int=5000, ims.sip_timer_t2_millis_int=16000, imsvoice.oip_source_from_header_bool=false, apn_settings_default_apn_types_string_array=null, imsvoice.voice_qos_precondition_supported_bool=true, imsvoice.dedicated_bearer_wait_timer_millis_int=8000, auto_retry_failed_wifi_emergency_call=false, read_only_apn_types_string_array=[dun], support_3gpp_call_forwarding_while_roaming_bool=true, allow_ussd_requests_via_telephony_manager_bool=true, imsss.terminal_based_call_waiting_sync_type_int=3, enabledNotifyWapMMSC=false, vvm_cellular_data_required_bool=false, imsvt.video_rtp_inactivity_timer_millis_int=0, imsvoice.prack_supported_for_18x_bool=false, carrier_app_wake_signal_config=[com.android.carrierdefaultapp/.CarrierDefaultBroadcastReceiver:com.android.internal.telephony.CARRIER_SIGNAL_RESET], notify_handover_video_from_wifi_to_lte_bool=false, prevent_clir_activation_and_deactivation_code_bool=false, ims.sip_timer_b_millis_int=128000, hide_tty_hco_vco_with_rtt=true, show_operator_name_in_statusbar_bool=false, sendMultipartSmsAsSeparateMessages=false, require_entitlement_checks_bool=true, iwlan.epdg_static_address_string=, imsemergency.emergency_over_cs_roaming_supported_access_network_types_int_array=[2, 1], use_call_forwarding_ussd_bool=false, carrier_roaming_satellite_upsell_notification_maximum_monthly_count_int=20, satellite_roaming_screen_off_inactivity_timeout_sec_int=30, carrier_data_service_wwan_class_override_string=, feature_access_codes_string_array=null, nr_advanced_requires_single_cc_above_bandwidth_threshold=false, min_udp_port_4500_nat_timeout_sec_int=300, imsvt.video_on_default_bearer_supported_bool=false, imsvoice.audio_rs_bandwidth_bps_int=600, sim_country_iso_override_string=, always_show_primary_signal_bar_in_opportunistic_network_boolean=false, rcs_config_server_url_string=, additional_settings_caller_id_visibility_bool=true, carrier_qualified_networks_service_package_override_string=, carrier_instant_lettering_encoding_string=, carrier_allow_transfer_ims_call_bool=false, opp_auto_data_switch_policy_int=0, carrier_name_string=T-Mobile, enableMultipartSMS=true, carrier_auto_cancel_cs_notification=true, ims.sip_timer_c_millis_int=210000, carrier_name_override_bool=true}]}
mCarrierAutoJoinResetCheckedForOobPseudonym=true
mCarrierPrivilegedPackagesBySimSlot=[ 
[]
]
NonCarrierMergedNetworksStatusTracker - Log Begin ----
mSubscriptionId=-1
dumpTimeMs=175369
mDisableStartTimeMs=0
mMinDisableDurationMs=0
mMaxDisableDurationMs=0
mListener=com.android.server.wifi.WifiConfigManager$3@d3c9894
mTemporarilyDisabledNonCarrierMergedListAtStart=
NonCarrierMergedNetworksStatusTracker - Log End ----

WifiApConfigStore config: ssid = "AndroidAP_9979" 
 Passphrase = <non-empty> 
 HiddenSsid = false 
 Channels = {1=0} 
 SecurityType = 1 
 MaxClient = 0 
 AutoShutdownEnabled = true 
 ShutdownTimeoutMillis = -1 
 ClientControlByUser = false 
 BlockedClientList = [] 
 AllowedClientList= [] 
 MacRandomizationSetting = 2 
 BridgedModeInstanceOpportunisticEnabled = true 
 BridgedModeOpportunisticShutdownTimeoutMillis = -1 
 Ieee80211axEnabled = true 
 Ieee80211beEnabled = true 
 isUserConfiguration = false 
 vendorElements = [] 
 mPersistentRandomizedMacAddress = 52:ce:28:07:23:7f 
 mAllowedAcsChannels2g = [] 
 mAllowedAcsChannels5g = [] 
 mAllowedAcsChannels6g = [] 
 mMaxChannelBandwidth = -1 
 mVendorData = [] 
 mIsClientIsolationEnabled = false 
 mIsBandOptimizationEnabled = true

Dump of PasspointManager
mEnabled: true
PasspointManager - Providers Begin ---
PasspointManager - Providers End ---
PasspointManager - Next provider ID to be assigned 0
Last sweep 0:00:03.572 ago.
ANQPRequestManager - Begin ---
ANQPRequestManager - End ---
Dump of PasspointNetworkNominateHelper
PasspointNetworkNominateHelper --- end ---

Chipset information :-----------------------------------------------
FW Version is: 1.0
Driver Version is: 1.0
Supported Feature set: -1
--------------------------------------------------------------------
Bug dump 0
system time = 5-23 15:38:4.198
kernel time = 35.799
reason = 7
kernel log: 

system log: 
--------- beginning of main
05-23 15:38:04.077  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: restrictions
05-23 15:38:04.077  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: role
05-23 15:38:04.077  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: rollback
05-23 15:38:04.077  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: runtime
05-23 15:38:04.077  1563  4441 I FSAPI   : (REDACTED) %s completed request (took %d ms): %s
05-23 15:38:04.077  4473  4473 E libbinder.Parcel: Reading a NULL string not supported here.
05-23 15:38:04.077  1563  1676 I FSAPI   : (REDACTED) Starting sending%srequest for %s (timeout=%dms)...
05-23 15:38:04.077  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: safety_center
05-23 15:38:04.077  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: scheduling_policy
05-23 15:38:04.078  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: sdk_sandbox
05-23 15:38:04.078  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: search
05-23 15:38:04.078  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: search_ui
05-23 15:38:04.078  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: sec_key_att_app_id_provider
05-23 15:38:04.078  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: secure_element
05-23 15:38:04.079  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: security_state
05-23 15:38:04.079  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: selection_toolbar
05-23 15:38:04.079  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: sensitive_content_protection_service
05-23 15:38:04.079  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: sensor_privacy
05-23 15:38:04.079  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: sensorservice
05-23 15:38:04.079  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: serial
05-23 15:38:04.079  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: servicediscovery
05-23 15:38:04.080  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: settings
05-23 15:38:04.080  4473  4473 E libbinder.Parcel: Reading a NULL string not supported here.
05-23 15:38:04.080  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: shortcut
05-23 15:38:04.080  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: simphonebook
05-23 15:38:04.081  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: slice
05-23 15:38:04.081  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: smartspace
05-23 15:38:04.081  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: soundtrigger
05-23 15:38:04.081  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: soundtrigger_middleware
05-23 15:38:04.081  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: speech_recognition
05-23 15:38:04.082  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: stats
05-23 15:38:04.082  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: statsbootstrap
05-23 15:38:04.082  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: statscompanion
05-23 15:38:04.082  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: statsmanager
05-23 15:38:04.082  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: statusbar
05-23 15:38:04.082  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: storaged
05-23 15:38:04.082  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: storaged_pri
05-23 15:38:04.083  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: storagestats
05-23 15:38:04.083  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: supervision
05-23 15:38:04.083  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: system_config
05-23 15:38:04.083  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: system_server_dumper
05-23 15:38:04.083  4473  4473 E libbinder.Parcel: Reading a NULL string not supported here.
05-23 15:38:04.083  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: system_update
05-23 15:38:04.084  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: task_continuity
05-23 15:38:04.084  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: telecom
05-23 15:38:04.086  1563  1676 E jtri    : *~*~*~ Previous channel {0} was garbage collected without being shut down! ~*~*~*
05-23 15:38:04.086  1563  1676 E jtri    :     Make sure to call shutdown()/shutdownNow()
05-23 15:38:04.086  1563  1676 E jtri    : java.lang.RuntimeException: ManagedChannel allocation site
05-23 15:38:04.086  1563  1676 E jtri    : 	at jtrh.<init>(:com.google.android.gms@261733035@26.17.33 (260400-911611531):21)
05-23 15:38:04.086  1563  1676 E jtri    : 	at jtri.<init>(:com.google.android.gms@261733035@26.17.33 (260400-911611531):10)
05-23 15:38:04.086  1563  1676 E jtri    : 	at jtrg.a(:com.google.android.gms@261733035@26.17.33 (260400-911611531):300)
05-23 15:38:04.086  1563  1676 E jtri    : 	at jsyy.a(:com.google.android.gms@261733035@26.17.33 (260400-911611531):5)
05-23 15:38:04.086  1563  1676 E jtri    : 	at jtey.a(:com.google.android.gms@261733035@26.17.33 (260400-911611531):18)
05-23 15:38:04.086  1563  1676 E jtri    : 	at avvu.c(:com.google.android.gms@261733035@26.17.33 (260400-911611531):63)
05-23 15:38:04.086  1563  1676 E jtri    : 	at avvu.<init>(:com.google.android.gms@261733035@26.17.33 (260400-911611531):41)
05-23 15:38:04.086  1563  1676 E jtri    : 	at avug.<init>(:com.google.android.gms@261733035@26.17.33 (260400-911611531):199)
05-23 15:38:04.086  1563  1676 E jtri    : 	at azba.mv(:com.google.android.gms@261733035@26.17.33 (260400-911611531):9)
05-23 15:38:04.086  1563  1676 E jtri    : 	at hetz.mv(:com.google.android.gms@261733035@26.17.33 (260400-911611531):14)
05-23 15:38:04.086  1563  1676 E jtri    : 	at azbe.g(:com.google.android.gms@261733035@26.17.33 (260400-911611531):3)
05-23 15:38:04.086  1563  1676 E jtri    : 	at azbe.e(:com.google.android.gms@261733035@26.17.33 (260400-911611531):2)
05-23 15:38:04.086  1563  1676 E jtri    : 	at aypz.call(:com.google.android.gms@261733035@26.17.33 (260400-911611531):30)
05-23 15:38:04.086  1563  1676 E jtri    : 	at java.util.concurrent.FutureTask.run(FutureTask.java:328)
05-23 15:38:04.086  1563  1676 E jtri    : 	at bkcg.c(:com.google.android.gms@261733035@26.17.33 (260400-911611531):50)
05-23 15:38:04.086  1563  1676 E jtri    : 	at bkcg.run(:com.google.android.gms@261733035@26.17.33 (260400-911611531):85)
05-23 15:38:04.086  1563  1676 E jtri    : 	at java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1100)
05-23 15:38:04.086  1563  1676 E jtri    : 	at java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:624)
05-23 15:38:04.086  1563  1676 E jtri    : 	at bkhx.run(:com.google.android.gms@261733035@26.17.33 (260400-911611531):8)
05-23 15:38:04.086  1563  1676 E jtri    : 	at java.lang.Thread.run(Thread.java:1572)
05-23 15:38:04.086  1563  1676 E jtri    : *~*~*~ Previous channel {0} was garbage collected without being shut down! ~*~*~*
05-23 15:38:04.086  1563  1676 E jtri    :     Make sure to call shutdown()/shutdownNow()
05-23 15:38:04.086  1563  1676 E jtri    : java.lang.RuntimeException: ManagedChannel allocation site
05-23 15:38:04.086  1563  1676 E jtri    : 	at jtrh.<init>(:com.google.android.gms@261733035@26.17.33 (260400-911611531):21)
05-23 15:38:04.086  1563  1676 E jtri    : 	at jtri.<init>(:com.google.android.gms@261733035@26.17.33 (260400-911611531):10)
05-23 15:38:04.086  1563  1676 E jtri    : 	at jtrg.a(:com.google.android.gms@261733035@26.17.33 (260400-911611531):300)
05-23 15:38:04.086  1563  1676 E jtri    : 	at jsyy.a(:com.google.android.gms@261733035@26.17.33 (260400-911611531):5)
05-23 15:38:04.086  1563  1676 E jtri    : 	at jtey.a(:com.google.android.gms@261733035@26.17.33 (260400-911611531):18)
05-23 15:38:04.086  1563  1676 E jtri    : 	at avvu.c(:com.google.android.gms@261733035@26.17.33 (260400-911611531):63)
05-23 15:38:04.086  1563  1676 E jtri    : 	at avvu.<init>(:com.google.android.gms@261733035@26.17.33 (260400-911611531):60)
05-23 15:38:04.086  1563  1676 E jtri    : 	at avug.<init>(:com.google.android.gms@261733035@26.17.33 (260400-911611531):199)
05-23 15:38:04.086  1563  1676 E jtri    : 	at azba.mv(:com.google.android.gms@261733035@26.17.33 (260400-911611531):9)
05-23 15:38:04.086  1563  1676 E jtri    : 	at hetz.mv(:com.google.android.gms@261733035@26.17.33 (260400-911611531):14)
05-23 15:38:04.086  1563  1676 E jtri    : 	at azbe.g(:com.google.android.gms@261733035@26.17.33 (260400-911611531):3)
05-23 15:38:04.086  1563  1676 E jtri    : 	at azbe.e(:com.google.android.gms@261733035@26.17.33 (260400-911611531):2)
05-23 15:38:04.086  1563  1676 E jtri    : 	at aypz.call(:com.google.android.gms@261733035@26.17.33 (260400-911611531):30)
05-23 15:38:04.086  1563  1676 E jtri    : 	at java.util.concurrent.FutureTask.run(FutureTask.java:328)
05-23 15:38:04.086  1563  1676 E jtri    : 	at bkcg.c(:com.google.android.gms@261733035@26.17.33 (260400-911611531):50)
05-23 15:38:04.086  1563  1676 E jtri    : 	at bkcg.run(:com.google.android.gms@261733035@26.17.33 (260400-911611531):85)
05-23 15:38:04.086  1563  1676 E jtri    : 	at java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1100)
05-23 15:38:04.086  1563  1676 E jtri    : 	at java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:624)
05-23 15:38:04.086  1563  1676 E jtri    : 	at bkhx.run(:com.google.android.gms@261733035@26.17.33 (260400-911611531):8)
05-23 15:38:04.086  1563  1676 E jtri    : 	at java.lang.Thread.run(Thread.java:1572)
05-23 15:38:04.086  1563  1676 E jtri    : *~*~*~ Previous channel {0} was garbage collected without being shut down! ~*~*~*
05-23 15:38:04.086  1563  1676 E jtri    :     Make sure to call shutdown()/shutdownNow()
05-23 15:38:04.086  1563  1676 E jtri    : java.lang.RuntimeException: ManagedChannel allocation site
05-23 15:38:04.086  1563  1676 E jtri    : 	at jtrh.<init>(:com.google.android.gms@261733035@26.17.33 (260400-911611531):21)
05-23 15:38:04.086  1563  1676 E jtri    : 	at jtri.<init>(:com.google.android.gms@261733035@26.17.33 (260400-911611531):10)
05-23 15:38:04.086  1563  1676 E jtri    : 	at jtrg.a(:com.google.android.gms@261733035@26.17.33 (260400-911611531):300)
05-23 15:38:04.086  1563  1676 E jtri    : 	at jsyy.a(:com.google.android.gms@261733035@26.17.33 (260400-911611531):5)
05-23 15:38:04.086  1563  1676 E jtri    : 	at jtey.a(:com.google.android.gms@261733035@26.17.33 (260400-911611531):18)
05-23 15:38:04.086  1563  1676 E jtri    : 	at avvu.c(:com.google.android.gms@261733035@26.17.33 (260400-911611531):63)
05-23 15:38:04.086  1563  1676 E jtri    : 	at avvu.<init>(:com.google.android.gms@261733035@26.17.33 (260400-911611531):20)
05-23 15:38:04.086  1563  1676 E jtri    : 	at avug.<init>(:com.google.android.gms@261733035@26.17.33 (260400-911611531):199)
05-23 15:38:04.086  1563  1676 E jtri    : 	at azba.mv(:com.google.android.gms@261733035@26.17.33 (260400-911611531):9)
05-23 15:38:04.086  1563  1676 E jtri    : 	at hetz.mv(:com.google.android.gms@261733035@26.17.33 (260400-911611531):14)
05-23 15:38:04.086  1563  1676 E jtri    : 	at azbe.g(:com.google.android.gms@261733035@26.17.33 (260400-911611531):3)
05-23 15:38:04.086  1563  1676 E jtri    : 	at azbe.e(:com.google.android.gms@261733035@26.17.33 (260400-911611531):2)
05-23 15:38:04.086  1563  1676 E jtri    : 	at aypz.call(:com.google.android.gms@261733035@26.17.33 (260400-911611531):30)
05-23 15:38:04.086  1563  1676 E jtri    : 	at java.util.concurrent.FutureTask.run(FutureTask.java:328)
05-23 15:38:04.086  1563  1676 E jtri    : 	at bkcg.c(:com.google.android.gms@261733035@26.17.33 (260400-911611531):50)
05-23 15:38:04.086  1563  1676 E jtri    : 	at bkcg.run(:com.google.android.gms@261733035@26.17.33 (260400-911611531):85)
05-23 15:38:04.086  1563  1676 E jtri    : 	at java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1100)
05-23 15:38:04.086  1563  1676 E jtri    : 	at java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:624)
05-23 15:38:04.086  1563  1676 E jtri    : 	at bkhx.run(:com.google.android.gms@261733035@26.17.33 (260400-911611531):8)
05-23 15:38:04.086  1563  1676 E jtri    : 	at java.lang.Thread.run(Thread.java:1572)
05-23 15:38:04.086  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: telephony.registry
05-23 15:38:04.086  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: telephony_ims
05-23 15:38:04.087  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: telephony_phone_number
05-23 15:38:04.088  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: testharness
05-23 15:38:04.088  4473  4473 E libbinder.Parcel: Reading a NULL string not supported here.
05-23 15:38:04.088  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: tethering
05-23 15:38:04.089  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: textclassification
05-23 15:38:04.089  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: textservices
05-23 15:38:04.089  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: texttospeech
05-23 15:38:04.089  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: thermalservice
05-23 15:38:04.089  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: thread_network
05-23 15:38:04.089  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: time_detector
05-23 15:38:04.090  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: time_zone_detector
05-23 15:38:04.090  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: tracing.proxy
05-23 15:38:04.090  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: tradeinmode
05-23 15:38:04.090  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: translation
05-23 15:38:04.090  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: transparency
05-23 15:38:04.090  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: trust
05-23 15:38:04.090  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: uimode
05-23 15:38:04.091  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: updatelock
05-23 15:38:04.091  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: uprobestats_bridge
05-23 15:38:04.092  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: uri_grants
05-23 15:38:04.092  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: usagestats
05-23 15:38:04.092  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: usb
05-23 15:38:04.093  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: user
05-23 15:38:04.093  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: uwb
05-23 15:38:04.094  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: vcn_management
05-23 15:38:04.094  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: vibrator_manager
05-23 15:38:04.095  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: virtualdevice
05-23 15:38:04.096  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: virtualdevice_native
05-23 15:38:04.097  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: voiceinteraction
05-23 15:38:04.098  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: vpn_management
05-23 15:38:04.098  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: wallpaper
05-23 15:38:04.099  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: wallpaper_effects_generation
05-23 15:38:04.099  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: wearable_sensing
05-23 15:38:04.099  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: webviewupdate
05-23 15:38:04.100  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: wifi
05-23 15:38:04.101  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: wifip2p
05-23 15:38:04.102  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: wifiscanner
05-23 15:38:04.102  4473  4473 W libbinder.BackendUnifiedServiceManager: Thread Pool max thread count is 0. Cannot cache binder as linkToDeath cannot be implemented. serviceName: window
05-23 15:38:04.108  3939  4007 I GoogleRestorePhotosBackupApiServiceImpl: called getPhotosBackupSettings
05-23 15:38:04.118   496   496 I adbd    : adbd service requested 'shell,v2,TERM=xterm-256color,raw:dumpsys diskstats'
05-23 15:38:04.118  3939  4007 I GoogleRestorePhotosBackupApiServiceImpl: called getPhotosBackupSettings
05-23 15:38:04.126  1563  1676 W AndroidBackupEventsLogg: Flag disabled, logging without sampling [CONTEXT service_id=229 ]
05-23 15:38:04.126  1563  1676 I AndroidBackupEventsLogg: (REDACTED) Logging without sampling %s
05-23 15:38:04.127  1563  4441 I FSAPI   : (REDACTED) %s completed request (took %d ms): %s
05-23 15:38:04.127  1563  2609 I FSAPI   : (REDACTED) Starting sending%srequest for %s (timeout=%dms)...
05-23 15:38:04.127   640  2055 W DiskStatsService: exception reading diskstats cache file
05-23 15:38:04.127   640  2055 W DiskStatsService: java.io.FileNotFoundException: /data/system/diskstats_cache.json: open failed: ENOENT (No such file or directory)
05-23 15:38:04.127   640  2055 W DiskStatsService: 	at libcore.io.IoBridge.open(IoBridge.java:574)
05-23 15:38:04.127   640  2055 W DiskStatsService: 	at libcore.io.IoUtils$FileReader.<init>(IoUtils.java:378)
05-23 15:38:04.127   640  2055 W DiskStatsService: 	at libcore.io.IoUtils.readFileAsString(IoUtils.java:291)
05-23 15:38:04.127   640  2055 W DiskStatsService: 	at com.android.server.DiskStatsService.reportCachedValues(DiskStatsService.java:209)
05-23 15:38:04.127   640  2055 W DiskStatsService: 	at com.android.server.DiskStatsService.dump(DiskStatsService.java:147)
05-23 15:38:04.127   640  2055 W DiskStatsService: 	at android.os.Binder.doDump(Binder.java:1009)
05-23 15:38:04.127   640  2055 W DiskStatsService: 	at android.os.Binder.dump(Binder.java:999)
05-23 15:38:04.127   640  2055 W DiskStatsService: 	at android.os.Binder.onTransact(Binder.java:925)
05-23 15:38:04.127   640  2055 W DiskStatsService: 	at android.os.Binder.execTransactInternal(Binder.java:1369)
05-23 15:38:04.127   640  2055 W DiskStatsService: 	at android.os.Binder.execTransact(Binder.java:1323)
05-23 15:38:04.127   640  2055 W DiskStatsService: Caused by: android.system.ErrnoException: open failed: ENOENT (No such file or directory)
05-23 15:38:04.127   640  2055 W DiskStatsService: 	at libcore.io.Linux.open(Native Method)
05-23 15:38:04.127   640  2055 W DiskStatsService: 	at libcore.io.ForwardingOs.open(ForwardingOs.java:579)
05-23 15:38:04.127   640  2055 W DiskStatsService: 	at libcore.io.BlockGuardOs.open(BlockGuardOs.java:274)
05-23 15:38:04.127   640  2055 W DiskStatsService: 	at libcore.io.IoBridge.open(IoBridge.java:560)
05-23 15:38:04.127   640  2055 W DiskStatsService: 	... 9 more
05-23 15:38:04.130  1563  4441 I FSAPI   : (REDACTED) %s completed request (took %d ms): %s
05-23 15:38:04.131  1563  2609 I FSAPI   : (REDACTED) Starting sending%srequest for %s (timeout=%dms)...
05-23 15:38:04.135  1210  1325 I NearbySharing: There is no account metadata, return default instance.
05-23 15:38:04.136  1210  1325 I NearbySharing: There is no account metadata, return default instance.
05-23 15:38:04.136  1210  1325 I NearbySharing: There is no account metadata, return default instance.
05-23 15:38:04.137  1563  2609 W FSAPI   : Failed to get feature status for NEARBY_SHARING. [CONTEXT service_id=329 ]
05-23 15:38:04.137  1563  2609 W FSAPI   : bihk: 35505: 
05-23 15:38:04.137  1563  2609 W FSAPI   : 	at dsyp.a(:com.google.android.gms@261733035@26.17.33 (260400-911611531):8)
05-23 15:38:04.137  1563  2609 W FSAPI   : 	at dsyp.dispatchTransaction(:com.google.android.gms@261733035@26.17.33 (260400-911611531):16)
05-23 15:38:04.137  1563  2609 W FSAPI   : 	at cbns.a(:com.google.android.gms@261733035@26.17.33 (260400-911611531):42)
05-23 15:38:04.137  1563  2609 W FSAPI   : 	at zwl.onTransact(:com.google.android.gms@261733035@26.17.33 (260400-911611531):23)
05-23 15:38:04.137  1563  2609 W FSAPI   : 	at android.os.Binder.execTransactInternal(Binder.java:1369)
05-23 15:38:04.137  1563  2609 W FSAPI   : 	at android.os.Binder.execTransact(Binder.java:1323)
05-23 15:38:04.137  1563  4441 I FSAPI   : (REDACTED) %s completed request (took %d ms): %s
05-23 15:38:04.137  1563  1617 W GmsCoreDailyFeatureLogg: Unable to get Feature State: Caught for NEARBY_SHARING class bihk: 35505:  [CONTEXT service_id=231 ]
05-23 15:38:04.137  1563  2609 I FSAPI   : (REDACTED) Starting sending%srequest for %s (timeout=%dms)...
05-23 15:38:04.139  1563  2609 I Pay     : Unable to create Wallet contextual card: missing account [CONTEXT service_id=198 ]
05-23 15:38:04.139  1563  4441 I FSAPI   : (REDACTED) %s completed request (took %d ms): %s
05-23 15:38:04.139  1563  1625 I FSAPI   : (REDACTED) Starting sending%srequest for %s (timeout=%dms)...
05-23 15:38:04.140  1563  1625 I Multidevice: [CrossDeviceServicesFeatureStatusIntentOperation] providing fetcher [CONTEXT service_id=348 ]
05-23 15:38:04.142  1210  2294 I Multidevice: (REDACTED) [FeatureSettingsApiService] getXdFeatureStatus( %s ): %s
05-23 15:38:04.142   496   496 I adbd    : adbd service requested 'shell,v2,TERM=xterm-256color,raw:dumpsys wifi'
05-23 15:38:04.142  1210  2294 I Multidevice: (REDACTED) [FeatureSettingsApiService] getXdFeatureStatus( %s ): %s
05-23 15:38:04.143  1563  4441 I FSAPI   : (REDACTED) %s completed request (took %d ms): %s
05-23 15:38:04.143  1563  1625 I FSAPI   : (REDACTED) Starting sending%srequest for %s (timeout=%dms)...
05-23 15:38:04.145  1210  1261 I PhoneskyFopFeatureStatu: null account [CONTEXT service_id=349 ]
05-23 15:38:04.145  1210  2137 I PhoneskyFopFeatureStatu: null account [CONTEXT service_id=349 ]
05-23 15:38:04.145  1210  4455 I FSAPI   : (REDACTED) %s completed request (took %d ms): %s
05-23 15:38:04.148  1210  2137 I NetworkScheduler.Stats: (REDACTED) HousekeepingTask %s/%s finished executing. cause:%s result: %s elapsed_millis: %s uptime_millis: %s exec_start_elapsed_seconds: %s
05-23 15:38:04.150  1563  1563 D BoundBrokerSvc: onUnbind: Intent { act=com.google.android.gms.common.stats.FeatureLoggingTask.ACTION_TASK_READY dat=chimera-action:/... xflg=0x4 cmp=com.google.android.gms/.chimera.GmsInternalBoundBrokerService }
05-23 15:38:04.152  1563  1563 D BoundBrokerSvc: onRebind: Intent { act=com.google.android.gms.common.stats.NotificationLoggingTask.ACTION_TASK_READY dat=chimera-action:/... xflg=0x4 cmp=com.google.android.gms/.chimera.GmsInternalBoundBrokerService }
05-23 15:38:04.154  1210  1261 I NetworkScheduler.Stats: (REDACTED) Task %s/%s started execution. cause:%s exec_start_elapsed_seconds: %s
05-23 15:38:04.155  1563  1625 W ChimeraUtils: Module com.google.android.gms.core missing resource null(0)
05-23 15:38:04.157  1210  2137 I NetworkScheduler.Stats: (REDACTED) HousekeepingTask %s/%s finished executing. cause:%s result: %s elapsed_millis: %s uptime_millis: %s exec_start_elapsed_seconds: %s
05-23 15:38:04.157  1563  1563 D BoundBrokerSvc: onUnbind: Intent { act=com.google.android.gms.common.stats.NotificationLoggingTask.ACTION_TASK_READY dat=chimera-action:/... xflg=0x4 cmp=com.google.android.gms/.chimera.GmsInternalBoundBrokerService }
05-23 15:38:04.161  1563  1563 D BoundBrokerSvc: onRebind: Intent { act=com.google.android.gms.adsidentity.service.AdservicesStatusService.ACTION_TASK_READY dat=chimera-action:/... xflg=0x4 cmp=com.google.android.gms/.chimera.GmsInternalBoundBrokerService }
05-23 15:38:04.162  1210  2137 I NetworkScheduler.Stats: (REDACTED) Task %s/%s started execution. cause:%s exec_start_elapsed_seconds: %s
05-23 15:38:04.166  1563  1617 I AdservicesStatusTask: Ad ID Response time: 3 milliseconds

ring-buffer = ring0

--------------------------------------------------------------------
--------------------------------------------------------------------
Bug dump 1
system time = 5-23 15:40:23.518
kernel time = 175.119
reason = 7
kernel log: 

system log: 
--------- beginning of main
05-23 15:40:11.232   640   798 D InetDiagMessage: Destroyed 0 sockets, proto=IPPROTO_TCP, family=AF_INET, states=14
05-23 15:40:11.232   640   798 D InetDiagMessage: Destroyed 0 sockets, proto=IPPROTO_TCP, family=AF_INET6, states=14
05-23 15:40:11.232   640   798 D InetDiagMessage: Destroyed live tcp sockets for uids={10205} in 1ms
05-23 15:40:11.232   640   798 D InetDiagMessage: Destroyed 0 sockets, proto=IPPROTO_TCP, family=AF_INET, states=14
05-23 15:40:11.233   640   798 D InetDiagMessage: Destroyed 0 sockets, proto=IPPROTO_TCP, family=AF_INET6, states=14
05-23 15:40:11.233   640   798 D InetDiagMessage: Destroyed live tcp sockets for uids={20205} in 1ms
05-23 15:40:11.965   640   792 I HalDevMgr: bestIfaceCreationProposal is null, requestIface=STA, existingIface=[name=wlan0 type=STA]
05-23 15:40:11.969   983  1118 D IpClient/wlan0: interfaceLinkStateChanged: ifindex 16 up
05-23 15:40:11.973   983  1118 D IpClient/wlan0: interfaceLinkStateChanged: ifindex 16 up
--------- beginning of system
05-23 15:40:11.975   640   666 D OomAdjuster: Not killing cached processes
05-23 15:40:12.208   884   884 I wpa_supplicant: wlan0: CTRL-EVENT-BEACON-LOSS 
05-23 15:40:14.971   640   823 I AdbWifiNetworkMonitor: Wi-Fi network available
05-23 15:40:14.971   640   823 I AdbWifiNetworkMonitor: Received the same Wi-Fi SSID. Ignoring.
05-23 15:40:14.972   640   792 I HalDevMgr: bestIfaceCreationProposal is null, requestIface=STA, existingIface=[name=wlan0 type=STA]
05-23 15:40:14.972  2485  2917 D WM-WorkConstraintsTrack: NetworkRequestConstraintController onCapabilitiesChanged callback
05-23 15:40:14.973  1765  2233 I BugleRcsEngine: Connected state: [1], networkType: [WIFI] [CONTEXT thread_id=57 ]
05-23 15:40:14.974   640   798 W libbinder.IPCThreadState: Sending oneway calls to frozen process.
05-23 15:40:14.974   640   798 W libbinder.IPCThreadState: Sending oneway calls to frozen process.
05-23 15:40:14.974   640   798 W libbinder.IPCThreadState: Sending oneway calls to frozen process.
05-23 15:40:14.974  1765  2479 I BugleRcsEngine: handleMessage processing message:[NOTIFY_UPTIME_IGNORE_STATE_CHANGED] with [non-null]:RcsEngineImpl reference [CONTEXT log_prefix="RcsEngineImpl[DUAL_REG]:[76fa9490-ffc6]>Handler" thread_id=69 ]
05-23 15:40:14.975  1210  1804 I NullBinder: NullBinder for android.net.action.RECOMMEND_NETWORKS triggering remote TransactionTooLargeException due to Service without Chimera impl, calling uid: 1000, calling pid: 0
05-23 15:40:14.975  1210  1804 W libbinder.Binder: Large reply transaction of 1056768 bytes, interface descriptor , function: UNKNOWN_FUNCTION_NAME, code: 1, flags: 17
05-23 15:40:14.976  1765  2479 I BugleRcsEngine: No RCS Configuration was found in Bugle for simID: redacted-pii:sim_id[chars:20,last3:897] [CONTEXT log_prefix="ProvisioningEngineDataRetriever" thread_id=69 ]
05-23 15:40:16.219   640   733 D ActivityManager: freezing 1870 com.google.android.apps.photos
05-23 15:40:16.221   640   798 D InetDiagMessage: Destroyed 0 sockets, proto=IPPROTO_TCP, family=AF_INET, states=14
05-23 15:40:16.221   640   798 D InetDiagMessage: Destroyed 0 sockets, proto=IPPROTO_TCP, family=AF_INET6, states=14
05-23 15:40:16.222   640   798 D InetDiagMessage: Destroyed live tcp sockets for uids={10171} in 0ms
05-23 15:40:16.236   640   733 D ActivityManager: freezing 1563 com.google.android.gms
05-23 15:40:16.237   640   733 D ActivityManager: freezing 4103 com.google.android.adservices.api
05-23 15:40:16.267   640   798 D InetDiagMessage: Destroyed 0 sockets, proto=IPPROTO_TCP, family=AF_INET, states=14
05-23 15:40:16.267   640   798 D InetDiagMessage: Destroyed 0 sockets, proto=IPPROTO_TCP, family=AF_INET6, states=14
05-23 15:40:16.267   640   798 D InetDiagMessage: Destroyed live tcp sockets for uids={10205} in 0ms
05-23 15:40:16.823   496   496 I adbd    : adbd service requested 'shell,v2,TERM=xterm-256color,raw:dumpsys package com.google.android.gms'
05-23 15:40:16.865  1447  1448 I artd    : GetBestInfo: odex next to the dex file (/data/app/~~3Ntp6qzFTHAsJLZ0AfssTw==/com.google.android.gms-DeZv5XqXlHqhcJB6HJ-mrw==/oat/arm64/base.odex) is kOatUpToDate with filter 'speed-profile' executable 'false'
05-23 15:40:16.865  1447  1448 I artd    : GetBestInfo: dm (/data/app/~~3Ntp6qzFTHAsJLZ0AfssTw==/com.google.android.gms-DeZv5XqXlHqhcJB6HJ-mrw==/split_AdsDynamite_installtime.dm) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.866  1447  1448 I artd    : GetBestInfo: odex next to the dex file (/data/app/~~3Ntp6qzFTHAsJLZ0AfssTw==/com.google.android.gms-DeZv5XqXlHqhcJB6HJ-mrw==/oat/arm64/split_CronetDynamite_installtime.odex) is kOatUpToDate with filter 'speed-profile' executable 'false'
05-23 15:40:16.866  1447  1448 I artd    : GetBestInfo: dm (/data/app/~~3Ntp6qzFTHAsJLZ0AfssTw==/com.google.android.gms-DeZv5XqXlHqhcJB6HJ-mrw==/split_DynamiteLoader_installtime.dm) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.867  1447  1448 I artd    : GetBestInfo: dm (/data/app/~~3Ntp6qzFTHAsJLZ0AfssTw==/com.google.android.gms-DeZv5XqXlHqhcJB6HJ-mrw==/split_DynamiteModulesA_installtime.dm) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.867  1447  1448 I artd    : GetBestInfo: dm (/data/app/~~3Ntp6qzFTHAsJLZ0AfssTw==/com.google.android.gms-DeZv5XqXlHqhcJB6HJ-mrw==/split_DynamiteModulesC_installtime.dm) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.868  1447  1448 I artd    : GetBestInfo: dm (/data/app/~~3Ntp6qzFTHAsJLZ0AfssTw==/com.google.android.gms-DeZv5XqXlHqhcJB6HJ-mrw==/split_GoogleCertificates_installtime.dm) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.868  1447  1448 I artd    : GetBestInfo: dm (/data/app/~~3Ntp6qzFTHAsJLZ0AfssTw==/com.google.android.gms-DeZv5XqXlHqhcJB6HJ-mrw==/split_MapsDynamite_installtime.dm) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.868  1447  1448 I artd    : GetBestInfo: dm (/data/app/~~3Ntp6qzFTHAsJLZ0AfssTw==/com.google.android.gms-DeZv5XqXlHqhcJB6HJ-mrw==/split_MeasurementDynamite_installtime.dm) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.869  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000001/CronetDynamite.apk has no usable artifacts
05-23 15:40:16.869  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000002/DynamiteLoader.apk has no usable artifacts
05-23 15:40:16.869  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/0000001c/dl-VisionOcr.optional_261733100000.apk has no usable artifacts
05-23 15:40:16.869  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000005/GoogleCertificates.apk has no usable artifacts
05-23 15:40:16.870  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/0000001d/dl-PlayCloudSearch.optional_261733100000.apk has no usable artifacts
05-23 15:40:16.870  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/0000001f/dl-Appsearch.optional_261733100400.apk has no usable artifacts
05-23 15:40:16.870  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/0000000e/dl-Appsearch.optional_261136100400.apk has no usable artifacts
05-23 15:40:16.870  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000018/dl-IdentityCredentialsPlatform.optional_261631100400.apk has no usable artifacts
05-23 15:40:16.871  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000016/dl-PlayCloudSearch.optional_261631100000.apk has no usable artifacts
05-23 15:40:16.871  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/0000000a/dl-MlkitOcrCommon.optional_261136100400.apk has no usable artifacts
05-23 15:40:16.871  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000008/dl-TfliteDynamiteDynamite.integ_252130102100400.apk has no usable artifacts
05-23 15:40:16.872  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000009/dl-VisionOcr.optional_261136100000.apk has no usable artifacts
05-23 15:40:16.872  1447  1448 I artd    : GetBestInfo: /data/user/0/com.google.android.gms/app_dg_cache/87E7746227FF4E457A2EB56043B95E41006DE49F/the.apk has no usable artifacts
05-23 15:40:16.872  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/0000001e/dl-IdentityCredentialsPlatform.optional_261733100400.apk has no usable artifacts
05-23 15:40:16.873  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/0000000d/dl-MlkitBarcodeUi.optional_261136100400.apk has no usable artifacts
05-23 15:40:16.873  1447  1448 I artd    : GetBestInfo: vdex next to the dex file (/data/user/0/com.google.android.gms/app_dg_cache/86153AA318A3F208A8C511836585279D8C7E8D94/oat/arm64/the.vdex) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.873  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/0000001b/dl-MlkitBarcodeUi.optional_261733100400.apk has no usable artifacts
05-23 15:40:16.874  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000017/dl-Appsearch.optional_261631100400.apk has no usable artifacts
05-23 15:40:16.874  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000012/dl-MlkitOcrCommon.optional_261631100400.apk has no usable artifacts
05-23 15:40:16.874  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000010/dl-PlayCloudSearch.optional_261136100000.apk has no usable artifacts
05-23 15:40:16.874  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000011/dl-VisionOcr.optional_261631100000.apk has no usable artifacts
05-23 15:40:16.875  1447  1448 I artd    : GetBestInfo: vdex next to the dex file (/data/user/0/com.google.android.gms/app_dg_cache/21A90BF1C1C388089206DCE215FB60D7882BFA18/oat/arm64/the.vdex) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.875  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000019/dl-MlkitOcrCommon.optional_261733100400.apk has no usable artifacts
05-23 15:40:16.876  1447  1448 I artd    : GetBestInfo: vdex next to the dex file (/data/user/0/com.google.android.gms/app_dg_cache/754BEC187BC4F26BBC31300F93DBF9B7506C1861/oat/arm64/the.vdex) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.876  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000007/MeasurementDynamite.apk has no usable artifacts
05-23 15:40:16.876  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000014/dl-MlkitBarcodeUi.optional_261631100400.apk has no usable artifacts
05-23 15:40:17.345  1023  1149 D SatelliteController: isInCarrierRoamingNbIotNtn: satellite is disabled
05-23 15:40:17.345  1023  1149 D SatelliteController: getCarrierRoamingNtnConnectType: Checking connect type from PLMN config for subId: 1
05-23 15:40:17.345  1023  1149 D TelephonyConfigUpdateInstallReceiver: getConfigParser: domain=satellite
05-23 15:40:17.345  1023  1149 V SatelliteController: satelliteConfigParser is not ready
05-23 15:40:17.345  1023  1149 D SatelliteController: isSatelliteAttachSupportedViaConfigupdater: return null (satelliteConfig is null)
05-23 15:40:17.345  1023  1149 D SatelliteController: getSatellitePerPlmnConfiguration: invalid subId or not supported via carrier.
05-23 15:40:17.345  1023  1149 D SatelliteController: getCarrierRoamingNtnConnectType: config: null
05-23 15:40:17.345  1023  1149 D TelephonyConfigUpdateInstallReceiver: getConfigParser: domain=satellite
05-23 15:40:17.345  1023  1149 V SatelliteController: satelliteConfigParser is not ready
05-23 15:40:17.345  1023  1149 D SatelliteController: getCarrierRoamingNtnConnectTypeViaConfigUpdater: return null (satelliteConfig is null)
05-23 15:40:17.346  1023  1149 D SatelliteController: getCarrierRoamingNtnConnectType: Falling back to global carrier config connect type: 0
05-23 15:40:17.346  1023  1149 D TelephonyConfigUpdateInstallReceiver: getConfigParser: domain=satellite
05-23 15:40:17.346  1023  1149 V SatelliteController: satelliteConfigParser is not ready
05-23 15:40:17.346  1023  1149 D SatelliteController: isSatelliteAttachSupportedViaConfigupdater: return null (satelliteConfig is null)
05-23 15:40:17.346  1023  1149 D SatelliteController: getCarrierRoamingNtnConnectType: Checking connect type from PLMN config for subId: 1
05-23 15:40:17.346  1023  1149 D TelephonyConfigUpdateInstallReceiver: getConfigParser: domain=satellite
05-23 15:40:17.346  1023  1149 V SatelliteController: satelliteConfigParser is not ready
05-23 15:40:17.346  1023  1149 D SatelliteController: isSatelliteAttachSupportedViaConfigupdater: return null (satelliteConfig is null)
05-23 15:40:17.346  1023  1149 D SatelliteController: getSatellitePerPlmnConfiguration: invalid subId or not supported via carrier.
05-23 15:40:17.346  1023  1149 D SatelliteController: getCarrierRoamingNtnConnectType: config: null
05-23 15:40:17.346  1023  1149 D TelephonyConfigUpdateInstallReceiver: getConfigParser: domain=satellite
05-23 15:40:17.346  1023  1149 V SatelliteController: satelliteConfigParser is not ready
05-23 15:40:17.346  1023  1149 D SatelliteController: getCarrierRoamingNtnConnectTypeViaConfigUpdater: return null (satelliteConfig is null)
05-23 15:40:17.346  1023  1149 D SatelliteController: getCarrierRoamingNtnConnectType: Falling back to global carrier config connect type: 0
05-23 15:40:17.346  1023  1149 D TelephonyConfigUpdateInstallReceiver: getConfigParser: domain=satellite
05-23 15:40:17.346  1023  1149 V SatelliteController: satelliteConfigParser is not ready
05-23 15:40:17.346  1023  1149 D SatelliteController: isSatelliteAttachSupportedViaConfigupdater: return null (satelliteConfig is null)
05-23 15:40:17.978   640   792 I HalDevMgr: bestIfaceCreationProposal is null, requestIface=STA, existingIface=[name=wlan0 type=STA]
05-23 15:40:20.141   640   796 E Nl80211Native: getChannelsMhzForBand: Wiphy index not recorded for band 8
05-23 15:40:20.141   640   796 E Nl80211Native: getChannelsMhzForBand: Wiphy index not recorded for band 16
05-23 15:40:20.141   640   796 D Nl80211Native: Ignoring unsupported scan type 2
05-23 15:40:20.141   640   796 I Nl80211Proxy: Sending Nl80211 message: GenericNetlinkMsg{ nlHeader{StructNlMsgHdr{ nlmsg_len{56}, nlmsg_type{29()}, nlmsg_flags{5(NLM_F_REQUEST|NLM_F_ACK)}, nlmsg_seq{21}, nlmsg_pid{0} }}, genNlHeader{StructGenNlMsgHdr{ command{33}, version{1}, reserved{0} }}, attributes{[StructNlAttr{ nla_len{8}, nla_type{-32723}, nla_value{04000000}, }, StructNlAttr{ nla_len{8}, nla_type{3}, nla_value{10000000}, }, StructNlAttr{ nla_len{12}, nla_type{-32724}, nla_value{080000008F090000}, }, StructNlAttr{ nla_len{8}, nla_type{158}, nla_value{00000000}, }]} }
05-23 15:40:20.141   640   796 I Nl80211Proxy: Received NLMSG_ERROR with error 0 for message GenericNetlinkMsg{ nlHeader{StructNlMsgHdr{ nlmsg_len{56}, nlmsg_type{29()}, nlmsg_flags{5(NLM_F_REQUEST|NLM_F_ACK)}, nlmsg_seq{21}, nlmsg_pid{0} }}, genNlHeader{StructGenNlMsgHdr{ command{33}, version{1}, reserved{0} }}, attributes{[StructNlAttr{ nla_len{8}, nla_type{-32723}, nla_value{04000000}, }, StructNlAttr{ nla_len{8}, nla_type{3}, nla_value{10000000}, }, StructNlAttr{ nla_len{12}, nla_type{-32724}, nla_value{080000008F090000}, }, StructNlAttr{ nla_len{8}, nla_type{158}, nla_value{00000000}, }]} }
05-23 15:40:20.192   640   792 D WifiNative: Scan result ready event
05-23 15:40:20.192   640   796 I Nl80211Proxy: Sending Nl80211 message: GenericNetlinkMsg{ nlHeader{StructNlMsgHdr{ nlmsg_len{20}, nlmsg_type{29()}, nlmsg_flags{769(NLM_F_REQUEST|NLM_F_DUMP)}, nlmsg_seq{22}, nlmsg_pid{0} }}, genNlHeader{StructGenNlMsgHdr{ command{5}, version{1}, reserved{0} }}, attributes{[]} }
05-23 15:40:20.192   640   796 I Nl80211Proxy: Received NLMSG_DONE for message GenericNetlinkMsg{ nlHeader{StructNlMsgHdr{ nlmsg_len{20}, nlmsg_type{29()}, nlmsg_flags{769(NLM_F_REQUEST|NLM_F_DUMP)}, nlmsg_seq{22}, nlmsg_pid{0} }}, genNlHeader{StructGenNlMsgHdr{ command{5}, version{1}, reserved{0} }}, attributes{[]} }
05-23 15:40:20.192   640   796 W Nl80211Utils: Malformed NEW_INTERFACE response: missing attributes
05-23 15:40:20.193   640   796 I Nl80211Proxy: Sending Nl80211 message: GenericNetlinkMsg{ nlHeader{StructNlMsgHdr{ nlmsg_len{28}, nlmsg_type{29()}, nlmsg_flags{769(NLM_F_REQUEST|NLM_F_DUMP)}, nlmsg_seq{23}, nlmsg_pid{0} }}, genNlHeader{StructGenNlMsgHdr{ command{32}, version{1}, reserved{0} }}, attributes{[StructNlAttr{ nla_len{8}, nla_type{3}, nla_value{10000000}, }]} }
05-23 15:40:20.193   640   796 I Nl80211Proxy: Received NLMSG_DONE for message GenericNetlinkMsg{ nlHeader{StructNlMsgHdr{ nlmsg_len{28}, nlmsg_type{29()}, nlmsg_flags{769(NLM_F_REQUEST|NLM_F_DUMP)}, nlmsg_seq{23}, nlmsg_pid{0} }}, genNlHeader{StructGenNlMsgHdr{ command{32}, version{1}, reserved{0} }}, attributes{[StructNlAttr{ nla_len{8}, nla_type{3}, nla_value{10000000}, }]} }
05-23 15:40:20.193  1210  1804 I WifiScanner: onFullResults
05-23 15:40:20.195   640   792 I HalDevMgr: bestIfaceCreationProposal is null, requestIface=STA, existingIface=[name=wlan0 type=STA]
05-23 15:40:20.983   640   823 I AdbWifiNetworkMonitor: Wi-Fi network available
05-23 15:40:20.983   640   823 I AdbWifiNetworkMonitor: Received the same Wi-Fi SSID. Ignoring.
05-23 15:40:20.983   640   792 I HalDevMgr: bestIfaceCreationProposal is null, requestIface=STA, existingIface=[name=wlan0 type=STA]
05-23 15:40:20.985   640   798 W libbinder.IPCThreadState: Sending oneway calls to frozen process.
05-23 15:40:20.985   640   798 W libbinder.IPCThreadState: Sending oneway calls to frozen process.
05-23 15:40:20.985   640   798 W libbinder.IPCThreadState: Sending oneway calls to frozen process.
05-23 15:40:20.985   640   798 W libbinder.IPCThreadState: Sending oneway calls to frozen process.
05-23 15:40:20.985   640   798 W libbinder.IPCThreadState: Sending oneway calls to frozen process.
05-23 15:40:20.985   640   798 W libbinder.IPCThreadState: Sending oneway calls to frozen process.
05-23 15:40:20.985  2485  2917 D WM-WorkConstraintsTrack: NetworkRequestConstraintController onCapabilitiesChanged callback
05-23 15:40:20.986  1765  2233 I BugleRcsEngine: Connected state: [1], networkType: [WIFI] [CONTEXT thread_id=57 ]
05-23 15:40:20.986  1765  2479 I BugleRcsEngine: handleMessage processing message:[NOTIFY_UPTIME_IGNORE_STATE_CHANGED] with [non-null]:RcsEngineImpl reference [CONTEXT log_prefix="RcsEngineImpl[DUAL_REG]:[76fa9490-ffc6]>Handler" thread_id=69 ]
05-23 15:40:20.987  1210  2128 I NullBinder: NullBinder for android.net.action.RECOMMEND_NETWORKS triggering remote TransactionTooLargeException due to Service without Chimera impl, calling uid: 1000, calling pid: 0
05-23 15:40:20.987  1210  2128 W libbinder.Binder: Large reply transaction of 1056768 bytes, interface descriptor , function: UNKNOWN_FUNCTION_NAME, code: 1, flags: 17
05-23 15:40:20.988  1765  2479 I BugleRcsEngine: No RCS Configuration was found in Bugle for simID: redacted-pii:sim_id[chars:20,last3:897] [CONTEXT log_prefix="ProvisioningEngineDataRetriever" thread_id=69 ]
05-23 15:40:23.444   496   496 I adbd    : adbd service requested 'shell,v2,TERM=xterm-256color,raw:dumpsys power'
05-23 15:40:23.472   496   496 I adbd    : adbd service requested 'shell,v2,TERM=xterm-256color,raw:dumpsys power'
05-23 15:40:23.504   496   496 I adbd    : adbd service requested 'shell,v2,TERM=xterm-256color,raw:dumpsys wifi'

ring-buffer = ring0

--------------------------------------------------------------------
--------------------------------------------------------------------
Bug dump 2
system time = 5-23 15:40:23.768
kernel time = 175.369
reason = 7
kernel log: 

system log: 
--------- beginning of main
05-23 15:40:11.233   640   798 D InetDiagMessage: Destroyed 0 sockets, proto=IPPROTO_TCP, family=AF_INET6, states=14
05-23 15:40:11.233   640   798 D InetDiagMessage: Destroyed live tcp sockets for uids={20205} in 1ms
05-23 15:40:11.965   640   792 I HalDevMgr: bestIfaceCreationProposal is null, requestIface=STA, existingIface=[name=wlan0 type=STA]
05-23 15:40:11.969   983  1118 D IpClient/wlan0: interfaceLinkStateChanged: ifindex 16 up
05-23 15:40:11.973   983  1118 D IpClient/wlan0: interfaceLinkStateChanged: ifindex 16 up
--------- beginning of system
05-23 15:40:11.975   640   666 D OomAdjuster: Not killing cached processes
05-23 15:40:12.208   884   884 I wpa_supplicant: wlan0: CTRL-EVENT-BEACON-LOSS 
05-23 15:40:14.971   640   823 I AdbWifiNetworkMonitor: Wi-Fi network available
05-23 15:40:14.971   640   823 I AdbWifiNetworkMonitor: Received the same Wi-Fi SSID. Ignoring.
05-23 15:40:14.972   640   792 I HalDevMgr: bestIfaceCreationProposal is null, requestIface=STA, existingIface=[name=wlan0 type=STA]
05-23 15:40:14.972  2485  2917 D WM-WorkConstraintsTrack: NetworkRequestConstraintController onCapabilitiesChanged callback
05-23 15:40:14.973  1765  2233 I BugleRcsEngine: Connected state: [1], networkType: [WIFI] [CONTEXT thread_id=57 ]
05-23 15:40:14.974   640   798 W libbinder.IPCThreadState: Sending oneway calls to frozen process.
05-23 15:40:14.974   640   798 W libbinder.IPCThreadState: Sending oneway calls to frozen process.
05-23 15:40:14.974   640   798 W libbinder.IPCThreadState: Sending oneway calls to frozen process.
05-23 15:40:14.974  1765  2479 I BugleRcsEngine: handleMessage processing message:[NOTIFY_UPTIME_IGNORE_STATE_CHANGED] with [non-null]:RcsEngineImpl reference [CONTEXT log_prefix="RcsEngineImpl[DUAL_REG]:[76fa9490-ffc6]>Handler" thread_id=69 ]
05-23 15:40:14.975  1210  1804 I NullBinder: NullBinder for android.net.action.RECOMMEND_NETWORKS triggering remote TransactionTooLargeException due to Service without Chimera impl, calling uid: 1000, calling pid: 0
05-23 15:40:14.975  1210  1804 W libbinder.Binder: Large reply transaction of 1056768 bytes, interface descriptor , function: UNKNOWN_FUNCTION_NAME, code: 1, flags: 17
05-23 15:40:14.976  1765  2479 I BugleRcsEngine: No RCS Configuration was found in Bugle for simID: redacted-pii:sim_id[chars:20,last3:897] [CONTEXT log_prefix="ProvisioningEngineDataRetriever" thread_id=69 ]
05-23 15:40:16.219   640   733 D ActivityManager: freezing 1870 com.google.android.apps.photos
05-23 15:40:16.221   640   798 D InetDiagMessage: Destroyed 0 sockets, proto=IPPROTO_TCP, family=AF_INET, states=14
05-23 15:40:16.221   640   798 D InetDiagMessage: Destroyed 0 sockets, proto=IPPROTO_TCP, family=AF_INET6, states=14
05-23 15:40:16.222   640   798 D InetDiagMessage: Destroyed live tcp sockets for uids={10171} in 0ms
05-23 15:40:16.236   640   733 D ActivityManager: freezing 1563 com.google.android.gms
05-23 15:40:16.237   640   733 D ActivityManager: freezing 4103 com.google.android.adservices.api
05-23 15:40:16.267   640   798 D InetDiagMessage: Destroyed 0 sockets, proto=IPPROTO_TCP, family=AF_INET, states=14
05-23 15:40:16.267   640   798 D InetDiagMessage: Destroyed 0 sockets, proto=IPPROTO_TCP, family=AF_INET6, states=14
05-23 15:40:16.267   640   798 D InetDiagMessage: Destroyed live tcp sockets for uids={10205} in 0ms
05-23 15:40:16.823   496   496 I adbd    : adbd service requested 'shell,v2,TERM=xterm-256color,raw:dumpsys package com.google.android.gms'
05-23 15:40:16.865  1447  1448 I artd    : GetBestInfo: odex next to the dex file (/data/app/~~3Ntp6qzFTHAsJLZ0AfssTw==/com.google.android.gms-DeZv5XqXlHqhcJB6HJ-mrw==/oat/arm64/base.odex) is kOatUpToDate with filter 'speed-profile' executable 'false'
05-23 15:40:16.865  1447  1448 I artd    : GetBestInfo: dm (/data/app/~~3Ntp6qzFTHAsJLZ0AfssTw==/com.google.android.gms-DeZv5XqXlHqhcJB6HJ-mrw==/split_AdsDynamite_installtime.dm) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.866  1447  1448 I artd    : GetBestInfo: odex next to the dex file (/data/app/~~3Ntp6qzFTHAsJLZ0AfssTw==/com.google.android.gms-DeZv5XqXlHqhcJB6HJ-mrw==/oat/arm64/split_CronetDynamite_installtime.odex) is kOatUpToDate with filter 'speed-profile' executable 'false'
05-23 15:40:16.866  1447  1448 I artd    : GetBestInfo: dm (/data/app/~~3Ntp6qzFTHAsJLZ0AfssTw==/com.google.android.gms-DeZv5XqXlHqhcJB6HJ-mrw==/split_DynamiteLoader_installtime.dm) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.867  1447  1448 I artd    : GetBestInfo: dm (/data/app/~~3Ntp6qzFTHAsJLZ0AfssTw==/com.google.android.gms-DeZv5XqXlHqhcJB6HJ-mrw==/split_DynamiteModulesA_installtime.dm) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.867  1447  1448 I artd    : GetBestInfo: dm (/data/app/~~3Ntp6qzFTHAsJLZ0AfssTw==/com.google.android.gms-DeZv5XqXlHqhcJB6HJ-mrw==/split_DynamiteModulesC_installtime.dm) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.868  1447  1448 I artd    : GetBestInfo: dm (/data/app/~~3Ntp6qzFTHAsJLZ0AfssTw==/com.google.android.gms-DeZv5XqXlHqhcJB6HJ-mrw==/split_GoogleCertificates_installtime.dm) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.868  1447  1448 I artd    : GetBestInfo: dm (/data/app/~~3Ntp6qzFTHAsJLZ0AfssTw==/com.google.android.gms-DeZv5XqXlHqhcJB6HJ-mrw==/split_MapsDynamite_installtime.dm) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.868  1447  1448 I artd    : GetBestInfo: dm (/data/app/~~3Ntp6qzFTHAsJLZ0AfssTw==/com.google.android.gms-DeZv5XqXlHqhcJB6HJ-mrw==/split_MeasurementDynamite_installtime.dm) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.869  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000001/CronetDynamite.apk has no usable artifacts
05-23 15:40:16.869  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000002/DynamiteLoader.apk has no usable artifacts
05-23 15:40:16.869  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/0000001c/dl-VisionOcr.optional_261733100000.apk has no usable artifacts
05-23 15:40:16.869  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000005/GoogleCertificates.apk has no usable artifacts
05-23 15:40:16.870  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/0000001d/dl-PlayCloudSearch.optional_261733100000.apk has no usable artifacts
05-23 15:40:16.870  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/0000001f/dl-Appsearch.optional_261733100400.apk has no usable artifacts
05-23 15:40:16.870  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/0000000e/dl-Appsearch.optional_261136100400.apk has no usable artifacts
05-23 15:40:16.870  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000018/dl-IdentityCredentialsPlatform.optional_261631100400.apk has no usable artifacts
05-23 15:40:16.871  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000016/dl-PlayCloudSearch.optional_261631100000.apk has no usable artifacts
05-23 15:40:16.871  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/0000000a/dl-MlkitOcrCommon.optional_261136100400.apk has no usable artifacts
05-23 15:40:16.871  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000008/dl-TfliteDynamiteDynamite.integ_252130102100400.apk has no usable artifacts
05-23 15:40:16.872  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000009/dl-VisionOcr.optional_261136100000.apk has no usable artifacts
05-23 15:40:16.872  1447  1448 I artd    : GetBestInfo: /data/user/0/com.google.android.gms/app_dg_cache/87E7746227FF4E457A2EB56043B95E41006DE49F/the.apk has no usable artifacts
05-23 15:40:16.872  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/0000001e/dl-IdentityCredentialsPlatform.optional_261733100400.apk has no usable artifacts
05-23 15:40:16.873  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/0000000d/dl-MlkitBarcodeUi.optional_261136100400.apk has no usable artifacts
05-23 15:40:16.873  1447  1448 I artd    : GetBestInfo: vdex next to the dex file (/data/user/0/com.google.android.gms/app_dg_cache/86153AA318A3F208A8C511836585279D8C7E8D94/oat/arm64/the.vdex) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.873  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/0000001b/dl-MlkitBarcodeUi.optional_261733100400.apk has no usable artifacts
05-23 15:40:16.874  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000017/dl-Appsearch.optional_261631100400.apk has no usable artifacts
05-23 15:40:16.874  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000012/dl-MlkitOcrCommon.optional_261631100400.apk has no usable artifacts
05-23 15:40:16.874  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000010/dl-PlayCloudSearch.optional_261136100000.apk has no usable artifacts
05-23 15:40:16.874  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000011/dl-VisionOcr.optional_261631100000.apk has no usable artifacts
05-23 15:40:16.875  1447  1448 I artd    : GetBestInfo: vdex next to the dex file (/data/user/0/com.google.android.gms/app_dg_cache/21A90BF1C1C388089206DCE215FB60D7882BFA18/oat/arm64/the.vdex) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.875  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000019/dl-MlkitOcrCommon.optional_261733100400.apk has no usable artifacts
05-23 15:40:16.876  1447  1448 I artd    : GetBestInfo: vdex next to the dex file (/data/user/0/com.google.android.gms/app_dg_cache/754BEC187BC4F26BBC31300F93DBF9B7506C1861/oat/arm64/the.vdex) is kOatUpToDate with filter 'verify' executable 'false'
05-23 15:40:16.876  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000007/MeasurementDynamite.apk has no usable artifacts
05-23 15:40:16.876  1447  1448 I artd    : GetBestInfo: /data/user_de/0/com.google.android.gms/app_chimera/m/00000014/dl-MlkitBarcodeUi.optional_261631100400.apk has no usable artifacts
05-23 15:40:17.345  1023  1149 D SatelliteController: isInCarrierRoamingNbIotNtn: satellite is disabled
05-23 15:40:17.345  1023  1149 D SatelliteController: getCarrierRoamingNtnConnectType: Checking connect type from PLMN config for subId: 1
05-23 15:40:17.345  1023  1149 D TelephonyConfigUpdateInstallReceiver: getConfigParser: domain=satellite
05-23 15:40:17.345  1023  1149 V SatelliteController: satelliteConfigParser is not ready
05-23 15:40:17.345  1023  1149 D SatelliteController: isSatelliteAttachSupportedViaConfigupdater: return null (satelliteConfig is null)
05-23 15:40:17.345  1023  1149 D SatelliteController: getSatellitePerPlmnConfiguration: invalid subId or not supported via carrier.
05-23 15:40:17.345  1023  1149 D SatelliteController: getCarrierRoamingNtnConnectType: config: null
05-23 15:40:17.345  1023  1149 D TelephonyConfigUpdateInstallReceiver: getConfigParser: domain=satellite
05-23 15:40:17.345  1023  1149 V SatelliteController: satelliteConfigParser is not ready
05-23 15:40:17.345  1023  1149 D SatelliteController: getCarrierRoamingNtnConnectTypeViaConfigUpdater: return null (satelliteConfig is null)
05-23 15:40:17.346  1023  1149 D SatelliteController: getCarrierRoamingNtnConnectType: Falling back to global carrier config connect type: 0
05-23 15:40:17.346  1023  1149 D TelephonyConfigUpdateInstallReceiver: getConfigParser: domain=satellite
05-23 15:40:17.346  1023  1149 V SatelliteController: satelliteConfigParser is not ready
05-23 15:40:17.346  1023  1149 D SatelliteController: isSatelliteAttachSupportedViaConfigupdater: return null (satelliteConfig is null)
05-23 15:40:17.346  1023  1149 D SatelliteController: getCarrierRoamingNtnConnectType: Checking connect type from PLMN config for subId: 1
05-23 15:40:17.346  1023  1149 D TelephonyConfigUpdateInstallReceiver: getConfigParser: domain=satellite
05-23 15:40:17.346  1023  1149 V SatelliteController: satelliteConfigParser is not ready
05-23 15:40:17.346  1023  1149 D SatelliteController: isSatelliteAttachSupportedViaConfigupdater: return null (satelliteConfig is null)
05-23 15:40:17.346  1023  1149 D SatelliteController: getSatellitePerPlmnConfiguration: invalid subId or not supported via carrier.
05-23 15:40:17.346  1023  1149 D SatelliteController: getCarrierRoamingNtnConnectType: config: null
05-23 15:40:17.346  1023  1149 D TelephonyConfigUpdateInstallReceiver: getConfigParser: domain=satellite
05-23 15:40:17.346  1023  1149 V SatelliteController: satelliteConfigParser is not ready
05-23 15:40:17.346  1023  1149 D SatelliteController: getCarrierRoamingNtnConnectTypeViaConfigUpdater: return null (satelliteConfig is null)
05-23 15:40:17.346  1023  1149 D SatelliteController: getCarrierRoamingNtnConnectType: Falling back to global carrier config connect type: 0
05-23 15:40:17.346  1023  1149 D TelephonyConfigUpdateInstallReceiver: getConfigParser: domain=satellite
05-23 15:40:17.346  1023  1149 V SatelliteController: satelliteConfigParser is not ready
05-23 15:40:17.346  1023  1149 D SatelliteController: isSatelliteAttachSupportedViaConfigupdater: return null (satelliteConfig is null)
05-23 15:40:17.978   640   792 I HalDevMgr: bestIfaceCreationProposal is null, requestIface=STA, existingIface=[name=wlan0 type=STA]
05-23 15:40:20.141   640   796 E Nl80211Native: getChannelsMhzForBand: Wiphy index not recorded for band 8
05-23 15:40:20.141   640   796 E Nl80211Native: getChannelsMhzForBand: Wiphy index not recorded for band 16
05-23 15:40:20.141   640   796 D Nl80211Native: Ignoring unsupported scan type 2
05-23 15:40:20.141   640   796 I Nl80211Proxy: Sending Nl80211 message: GenericNetlinkMsg{ nlHeader{StructNlMsgHdr{ nlmsg_len{56}, nlmsg_type{29()}, nlmsg_flags{5(NLM_F_REQUEST|NLM_F_ACK)}, nlmsg_seq{21}, nlmsg_pid{0} }}, genNlHeader{StructGenNlMsgHdr{ command{33}, version{1}, reserved{0} }}, attributes{[StructNlAttr{ nla_len{8}, nla_type{-32723}, nla_value{04000000}, }, StructNlAttr{ nla_len{8}, nla_type{3}, nla_value{10000000}, }, StructNlAttr{ nla_len{12}, nla_type{-32724}, nla_value{080000008F090000}, }, StructNlAttr{ nla_len{8}, nla_type{158}, nla_value{00000000}, }]} }
05-23 15:40:20.141   640   796 I Nl80211Proxy: Received NLMSG_ERROR with error 0 for message GenericNetlinkMsg{ nlHeader{StructNlMsgHdr{ nlmsg_len{56}, nlmsg_type{29()}, nlmsg_flags{5(NLM_F_REQUEST|NLM_F_ACK)}, nlmsg_seq{21}, nlmsg_pid{0} }}, genNlHeader{StructGenNlMsgHdr{ command{33}, version{1}, reserved{0} }}, attributes{[StructNlAttr{ nla_len{8}, nla_type{-32723}, nla_value{04000000}, }, StructNlAttr{ nla_len{8}, nla_type{3}, nla_value{10000000}, }, StructNlAttr{ nla_len{12}, nla_type{-32724}, nla_value{080000008F090000}, }, StructNlAttr{ nla_len{8}, nla_type{158}, nla_value{00000000}, }]} }
05-23 15:40:20.192   640   792 D WifiNative: Scan result ready event
05-23 15:40:20.192   640   796 I Nl80211Proxy: Sending Nl80211 message: GenericNetlinkMsg{ nlHeader{StructNlMsgHdr{ nlmsg_len{20}, nlmsg_type{29()}, nlmsg_flags{769(NLM_F_REQUEST|NLM_F_DUMP)}, nlmsg_seq{22}, nlmsg_pid{0} }}, genNlHeader{StructGenNlMsgHdr{ command{5}, version{1}, reserved{0} }}, attributes{[]} }
05-23 15:40:20.192   640   796 I Nl80211Proxy: Received NLMSG_DONE for message GenericNetlinkMsg{ nlHeader{StructNlMsgHdr{ nlmsg_len{20}, nlmsg_type{29()}, nlmsg_flags{769(NLM_F_REQUEST|NLM_F_DUMP)}, nlmsg_seq{22}, nlmsg_pid{0} }}, genNlHeader{StructGenNlMsgHdr{ command{5}, version{1}, reserved{0} }}, attributes{[]} }
05-23 15:40:20.192   640   796 W Nl80211Utils: Malformed NEW_INTERFACE response: missing attributes
05-23 15:40:20.193   640   796 I Nl80211Proxy: Sending Nl80211 message: GenericNetlinkMsg{ nlHeader{StructNlMsgHdr{ nlmsg_len{28}, nlmsg_type{29()}, nlmsg_flags{769(NLM_F_REQUEST|NLM_F_DUMP)}, nlmsg_seq{23}, nlmsg_pid{0} }}, genNlHeader{StructGenNlMsgHdr{ command{32}, version{1}, reserved{0} }}, attributes{[StructNlAttr{ nla_len{8}, nla_type{3}, nla_value{10000000}, }]} }
05-23 15:40:20.193   640   796 I Nl80211Proxy: Received NLMSG_DONE for message GenericNetlinkMsg{ nlHeader{StructNlMsgHdr{ nlmsg_len{28}, nlmsg_type{29()}, nlmsg_flags{769(NLM_F_REQUEST|NLM_F_DUMP)}, nlmsg_seq{23}, nlmsg_pid{0} }}, genNlHeader{StructGenNlMsgHdr{ command{32}, version{1}, reserved{0} }}, attributes{[StructNlAttr{ nla_len{8}, nla_type{3}, nla_value{10000000}, }]} }
05-23 15:40:20.193  1210  1804 I WifiScanner: onFullResults
05-23 15:40:20.195   640   792 I HalDevMgr: bestIfaceCreationProposal is null, requestIface=STA, existingIface=[name=wlan0 type=STA]
05-23 15:40:20.983   640   823 I AdbWifiNetworkMonitor: Wi-Fi network available
05-23 15:40:20.983   640   823 I AdbWifiNetworkMonitor: Received the same Wi-Fi SSID. Ignoring.
05-23 15:40:20.983   640   792 I HalDevMgr: bestIfaceCreationProposal is null, requestIface=STA, existingIface=[name=wlan0 type=STA]
05-23 15:40:20.985   640   798 W libbinder.IPCThreadState: Sending oneway calls to frozen process.
05-23 15:40:20.985   640   798 W libbinder.IPCThreadState: Sending oneway calls to frozen process.
05-23 15:40:20.985   640   798 W libbinder.IPCThreadState: Sending oneway calls to frozen process.
05-23 15:40:20.985   640   798 W libbinder.IPCThreadState: Sending oneway calls to frozen process.
05-23 15:40:20.985   640   798 W libbinder.IPCThreadState: Sending oneway calls to frozen process.
05-23 15:40:20.985   640   798 W libbinder.IPCThreadState: Sending oneway calls to frozen process.
05-23 15:40:20.985  2485  2917 D WM-WorkConstraintsTrack: NetworkRequestConstraintController onCapabilitiesChanged callback
05-23 15:40:20.986  1765  2233 I BugleRcsEngine: Connected state: [1], networkType: [WIFI] [CONTEXT thread_id=57 ]
05-23 15:40:20.986  1765  2479 I BugleRcsEngine: handleMessage processing message:[NOTIFY_UPTIME_IGNORE_STATE_CHANGED] with [non-null]:RcsEngineImpl reference [CONTEXT log_prefix="RcsEngineImpl[DUAL_REG]:[76fa9490-ffc6]>Handler" thread_id=69 ]
05-23 15:40:20.987  1210  2128 I NullBinder: NullBinder for android.net.action.RECOMMEND_NETWORKS triggering remote TransactionTooLargeException due to Service without Chimera impl, calling uid: 1000, calling pid: 0
05-23 15:40:20.987  1210  2128 W libbinder.Binder: Large reply transaction of 1056768 bytes, interface descriptor , function: UNKNOWN_FUNCTION_NAME, code: 1, flags: 17
05-23 15:40:20.988  1765  2479 I BugleRcsEngine: No RCS Configuration was found in Bugle for simID: redacted-pii:sim_id[chars:20,last3:897] [CONTEXT log_prefix="ProvisioningEngineDataRetriever" thread_id=69 ]
05-23 15:40:23.444   496   496 I adbd    : adbd service requested 'shell,v2,TERM=xterm-256color,raw:dumpsys power'
05-23 15:40:23.472   496   496 I adbd    : adbd service requested 'shell,v2,TERM=xterm-256color,raw:dumpsys power'
05-23 15:40:23.504   496   496 I adbd    : adbd service requested 'shell,v2,TERM=xterm-256color,raw:dumpsys wifi'
05-23 15:40:23.736   640   792 W LastMileLogger: Failed to read event trace: /sys/kernel/debug/tracing/instances/wifi/trace
05-23 15:40:23.740   640   792 I android_os_HwBinder: HwBinder: Starting thread pool for getting: android.hidl.manager@1.0::IServiceManager/default
05-23 15:40:23.741   289   289 I hwservicemanager: getTransport: Cannot find entry android.hardware.wifi.hostapd@1.0::IHostapd/default in either framework or device VINTF manifest.
05-23 15:40:23.759   496   496 I adbd    : adbd service requested 'shell,v2,TERM=xterm-256color,raw:dumpsys wifi'

ring-buffer = ring0

--------------------------------------------------------------------
Last Flush Time: {}
--------------------------------------------------------------------
No fates fetched for "Last failed connection fates"
HAL provided zero fates for "Latest fates"
No last mile log for "Last failed last-mile log"
No last mile log for "Latest last-mile log"
--------------------------------------------------------------------
Dump of WifiConnectivityManager
WifiConnectivityManager - Log Begin ----
mIsLocationModeEnabled: true
mPnoScanEnabledByFramework: true
mEnablePnoScanAfterWifiToggle: true
mMultiInternetConnectionState 0
2026-05-23T15:37:32.631069 - Set WiFi enabled
2026-05-23T15:37:32.631073 - Stopping WifiConnectivityManager
2026-05-23T15:37:33.313073 - handleConnectionStateChanged: state=disconnected
2026-05-23T15:37:33.313611 - startConnectivityScan: screenOn=false wifiState=disconnected scanImmediately=true wifiEnabled=true mAutoJoinEnabled=false mAutoJoinEnabledExternal=true mAutoJoinEnabledExternalSetByDeviceAdmin=false mPnoScanEnabledByFramework=true mEnablePnoScanAfterWifiToggle=true mSpecificNetworkRequestInProgress=false mTrustedConnectionAllowed=false isSufficiencyCheckEnabled=true isAssociatedNetworkSelectionEnabled=true noPotentialNetworkAvailable=false
2026-05-23T15:37:39.959539 - handleScreenStateChanged: screenOn=true
2026-05-23T15:37:39.959582 - startConnectivityScan: screenOn=true wifiState=disconnected scanImmediately=false wifiEnabled=true mAutoJoinEnabled=false mAutoJoinEnabledExternal=true mAutoJoinEnabledExternalSetByDeviceAdmin=false mPnoScanEnabledByFramework=true mEnablePnoScanAfterWifiToggle=true mSpecificNetworkRequestInProgress=false mTrustedConnectionAllowed=false isSufficiencyCheckEnabled=true isAssociatedNetworkSelectionEnabled=true noPotentialNetworkAvailable=false
2026-05-23T15:37:40.121117 - setTrustedConnectionAllowed: allowed=true
2026-05-23T15:37:40.121125 - Starting up WifiConnectivityManager
2026-05-23T15:37:40.121212 - startConnectivityScan: screenOn=true wifiState=disconnected scanImmediately=true wifiEnabled=true mAutoJoinEnabled=true mAutoJoinEnabledExternal=true mAutoJoinEnabledExternalSetByDeviceAdmin=false mPnoScanEnabledByFramework=true mEnablePnoScanAfterWifiToggle=true mSpecificNetworkRequestInProgress=false mTrustedConnectionAllowed=true isSufficiencyCheckEnabled=true isAssociatedNetworkSelectionEnabled=true noPotentialNetworkAvailable=false
2026-05-23T15:37:40.121397 - schedulePeriodicScanTimer intervalMs 20000
2026-05-23T15:37:44.861524 - AllSingleScanListener onResults: start network selection
2026-05-23T15:37:44.876391 - About to run SavedNetworkNominator :
2026-05-23T15:37:44.913219 - About to run NetworkSuggestionNominator :
2026-05-23T15:37:44.913308 - did not see any matching auto-join enabled network suggestions.
2026-05-23T15:37:44.914824 - removeAutoUpgradeSecurityParamsIfNecessary: SSID: "AndroidWifi" baseSecurityType: 2 upgradableSecurityType: 4 isLegacyNetworkInRange: false isUpgradableTypeOnlyInRange: false isAutoUpgradeEnabled: true
2026-05-23T15:37:44.914844 - removeAutoUpgradeSecurityParamsIfNecessary: SSID: "AndroidWifi" baseSecurityType: 0 upgradableSecurityType: 6 isLegacyNetworkInRange: true isUpgradableTypeOnlyInRange: false isAutoUpgradeEnabled: true
2026-05-23T15:37:44.914851 - Remove upgradable security type 6 for the network.
2026-05-23T15:37:44.914868 - removeAutoUpgradeSecurityParamsIfNecessary: SSID: "AndroidWifi" baseSecurityType: 3 upgradableSecurityType: 9 isLegacyNetworkInRange: false isUpgradableTypeOnlyInRange: false isAutoUpgradeEnabled: true
2026-05-23T15:37:44.914928 - Candidate { config = 0, bssid = 00:13:10:85:fe:01, freq = 2447, channelWidth = 0, rssi = -50, Mbps = 49, nominator = 0, pInternet = 50, numRebootsSinceLastUse = 1, saved, trusted, open }
2026-05-23T15:37:44.916300 - BubbleFunScorer_v2 would choose 0 score 14.411264143668568+/-0.30216115261975496 expid 42598152
2026-05-23T15:37:44.916313 - CompatibilityScorer would choose 0 score 99.95+/-10.0 expid 42504592
2026-05-23T15:37:44.998555 - ScoreCardBasedScorer would choose 0 score 100.0+/-10.0 expid 42902385
2026-05-23T15:37:44.998598 - ThroughputScorer chooses 0 score 3560.95+/-10.0 expid 42330058
2026-05-23T15:37:44.998643 - AllSingleScanListener:  WNS candidate-"AndroidWifi"
2026-05-23T15:37:44.999603 - connectToNetwork(ConcreteClientModeManager{id=4074 iface=wlan0 role=ROLE_CLIENT_PRIMARY}): Connect to "AndroidWifi" : 00:13:10:85:fe:01 from Disconnected
2026-05-23T15:37:44.999611 - noteConnectionAttempt: timeMillis=16600
2026-05-23T15:37:45.102848 - handleConnectionStateChanged: state=transitioning
2026-05-23T15:37:45.102867 - startConnectivityScan: screenOn=true wifiState=transitioning scanImmediately=false wifiEnabled=true mAutoJoinEnabled=true mAutoJoinEnabledExternal=true mAutoJoinEnabledExternalSetByDeviceAdmin=false mPnoScanEnabledByFramework=true mEnablePnoScanAfterWifiToggle=true mSpecificNetworkRequestInProgress=false mTrustedConnectionAllowed=true isSufficiencyCheckEnabled=true isAssociatedNetworkSelectionEnabled=true noPotentialNetworkAvailable=false
2026-05-23T15:37:45.102873 - cancelPeriodicScanTimer
2026-05-23T15:37:48.114314 - handleConnectionStateChanged: state=connected
2026-05-23T15:37:48.117035 - startConnectivityScan: screenOn=true wifiState=connected scanImmediately=false wifiEnabled=true mAutoJoinEnabled=true mAutoJoinEnabledExternal=true mAutoJoinEnabledExternalSetByDeviceAdmin=false mPnoScanEnabledByFramework=true mEnablePnoScanAfterWifiToggle=true mSpecificNetworkRequestInProgress=false mTrustedConnectionAllowed=true isSufficiencyCheckEnabled=true isAssociatedNetworkSelectionEnabled=true noPotentialNetworkAvailable=false
2026-05-23T15:37:48.117049 - Last periodic single scan started 7996ms ago, defer this new scan request.
2026-05-23T15:37:48.117051 - schedulePeriodicScanTimer intervalMs 12004
2026-05-23T15:38:00.123685 - Current connected network: 0
2026-05-23T15:38:00.124295 - No full band scan because current network is sufficient
2026-05-23T15:38:00.126225 - schedulePeriodicScanTimer intervalMs 20000
2026-05-23T15:38:00.177458 - AllSingleScanListener onResults: start network selection
2026-05-23T15:38:00.177710 - Current connected network: 0
2026-05-23T15:38:00.177737 - wlan0: Current connected network already sufficient. Skip network selection.
2026-05-23T15:38:00.178095 - AllSingleScanListener:  No candidates
2026-05-23T15:38:20.132335 - Current connected network: 0
2026-05-23T15:38:20.132382 - No full band scan because current network is sufficient
2026-05-23T15:38:20.133318 - schedulePeriodicScanTimer intervalMs 40000
2026-05-23T15:38:20.181479 - AllSingleScanListener onResults: start network selection
2026-05-23T15:38:20.181588 - Current connected network: 0
2026-05-23T15:38:20.181618 - wlan0: Current connected network already sufficient. Skip network selection.
2026-05-23T15:38:20.181630 - AllSingleScanListener:  No candidates
2026-05-23T15:39:00.134079 - Current connected network: 0
2026-05-23T15:39:00.134107 - No full band scan because current network is sufficient
2026-05-23T15:39:00.134184 - schedulePeriodicScanTimer intervalMs 80000
2026-05-23T15:39:00.177088 - AllSingleScanListener onResults: start network selection
2026-05-23T15:39:00.177155 - Current connected network: 0
2026-05-23T15:39:00.177174 - wlan0: Current connected network already sufficient. Skip network selection.
2026-05-23T15:39:00.177182 - AllSingleScanListener:  No candidates
2026-05-23T15:40:20.140074 - Current connected network: 0
2026-05-23T15:40:20.140202 - No full band scan because current network is sufficient
2026-05-23T15:40:20.140602 - schedulePeriodicScanTimer intervalMs 160000
2026-05-23T15:40:20.195876 - AllSingleScanListener onResults: start network selection
2026-05-23T15:40:20.196032 - Current connected network: 0
2026-05-23T15:40:20.196083 - wlan0: Current connected network already sufficient. Skip network selection.
2026-05-23T15:40:20.196102 - AllSingleScanListener:  No candidates
WifiConnectivityManager - Log End ----
WifiOpenNetworkNotifier: 
mSettingEnabled true
currentTime: 1779531023988
mNotificationRepeatTime: 0
mState: 0
mBlocklistedSsids: {}
Dump of WifiBlocklistMonitor
WifiBlocklistMonitor - Bssid blocklist begin ----
WifiBlocklistMonitor - Bssid blocklist end ----
Dump of BSSID to Affiliated BSSID mapping
WifiBlocklistMonitor - Bssid blocklist logs begin ----
List of SSIDs to never block:
WifiBlocklistMonitor - Bssid blocklist logs end ----
Dump of ExternalPnoScanRequestManager
ExternalPnoScanRequestManager - Log Begin ----
No external PNO scan request set.
mCurrentRequestOnPnoNetworkFoundCount: 0
ExternalPnoScanRequestManager - Log End ----
Dump of WifiConnectivityHelper
WifiConnectivityHelper - Log Begin ----
mFirmwareRoamingSupported: false
mMaxNumBlocklistBssid: -1
mMaxNumAllowlistSsid: -1
WifiConnectivityHelper - Log End ----
Dump of WifiHealthMonitor
WifiHealthMonitor - Log Begin ----
System Info Stats
current SW build: OS build version: CP21.260330.005 dev-keys Wifi stack version: 370399999 Wifi driver version: 1.0 Wifi firmware version: 1.0
currScanStats: last scan time: 1779530864602 APs found at 2G: 1 APs found above 2g: 0
prevScanStats: last scan time: 1779298310796 APs found at 2G: 1 APs found above 2g: 0
configured network connection stats
SSID: "AndroidWifi"
 LastRssiPollTime: 172584 LastRssiPoll: -50 LastTxSpeedPoll: -1
 StatsRecent:  ConnectAttempt: 12 ConnectFailure: 0 ConnectDurSec: 0 AssocRej: 0 AssocTimeout: 0 AuthFailure: 0 ShortDiscNonlocal: 0 DisconnectNonlocal: 0 Disconnect: 0 ConsecutiveConnectFailure: 0 ConnectFailureDiscon: 0 ConsecutiveWrongPassword: 0
 StatsCurr:  ConnectAttempt: 0 ConnectFailure: 0 ConnectDurSec: 0 AssocRej: 0 AssocTimeout: 0 AuthFailure: 0 ShortDiscNonlocal: 0 DisconnectNonlocal: 0 Disconnect: 0 ConsecutiveConnectFailure: 0 ConnectFailureDiscon: 0 ConsecutiveWrongPassword: 0
 StatsPrev:  ConnectAttempt: 0 ConnectFailure: 0 ConnectDurSec: 0 AssocRej: 0 AssocTimeout: 0 AuthFailure: 0 ShortDiscNonlocal: 0 DisconnectNonlocal: 0 Disconnect: 0 ConsecutiveConnectFailure: 0 ConnectFailureDiscon: 0 ConsecutiveWrongPassword: 0 BandwidthStats:
 avgKbps:  0 0 0 0 0
 count:  0 0 0 0 0
 avgKbps:  0 0 0 0 0
 count:  0 0 0 0 0

 avgKbps:  0 0 0 0 0
 count:  0 0 0 0 0
 avgKbps:  0 0 0 0 0
 count:  0 0 0 0 0


networks with failure increase: 

networks with failure drop: 

networks with high failure without previous stats: 

WifiHealthMonitor - Log End ----
Dump of WifiScoreCard
current SSID(s):{iface=wlan0,ssid="AndroidWifi"}
 BW Estimation Stats
2G
 Tx
 Count
 0 0 0 0 0
 AvgKbps
 0 0 0 0 0
 BwEst error
 0 0 0 0 0
 L2 error
 0 0 0 0 0
 Rx
 Count
 0 0 0 0 0
 AvgKbps
 0 0 0 0 0
 BwEst error
 0 0 0 0 0
 L2 error
 0 0 0 0 0
5G
 Tx
 Count
 0 0 0 0 0
 AvgKbps
 0 0 0 0 0
 BwEst error
 0 0 0 0 0
 L2 error
 0 0 0 0 0
 Rx
 Count
 0 0 0 0 0
 AvgKbps
 0 0 0 0 0
 BwEst error
 0 0 0 0 0
 L2 error
 0 0 0 0 0

Dump of WakeupController
USE_PLATFORM_WIFI_WAKE: true
mWifiWakeupEnabled: true
isOnboarded: false
configStore hasBeenRead: true
mIsActive: false
mNumScansHandled: 0
WakeupLock: 
mNumScans: 0
mIsInitialized: true
Locked networks: 0
Dump of WifiLastResortWatchdog
WifiLastResortWatchdog - Log Begin ----
2026-05-23T15:37:48.128090 - connectedStateTransition: isEntering = true
2026-05-23T15:37:48.128104 - connectedStateTransition: setWatchdogTriggerEnabled to true
WifiLastResortWatchdog - Log End ----
Dump of AdaptiveConnectivityEnabledSettingObserver
mAdaptiveConnectivityEnabled=true
Dump of WifiGlobals
mPollRssiIntervalMillis=3000
mIsPollRssiIntervalOverridden=false
mPollRssiShortIntervalMillis=3000
mPollRssiLongIntervalMillis=6000
mIpReachabilityDisconnectEnabled=true
mIsBluetoothConnected=false
mIsWpa3SaeUpgradeOffloadEnabled=false
mIsUsingExternalScorer=false
mIsWepAllowed=true
IsD2dSupportedWhenInfraStaDisabled=false
mIsWpa3SaeH2eSupported=true
mIsMultiInternetSameBandConnectionAllowed=false
mIsMultiInternetSameBssidConnectionAllowed=false
carrierId=1839, eapFailureCode=0, displayNotification=false, threshold=1, durationMs=14400000
carrierId=1839, eapFailureCode=1026, displayNotification=false, threshold=1, durationMs=14400000
carrierId=1839, eapFailureCode=1031, displayNotification=true, threshold=1, durationMs=86400000
carrierId=1839, eapFailureCode=16384, displayNotification=false, threshold=1, durationMs=300000
carrierId=1839, eapFailureCode=16385, displayNotification=false, threshold=3, durationMs=7200000
mSendDhcpHostnameRestriction=0
Dump of SarManager
isSarSupported: false
isSarVoiceCallSupported: false
isSarSoftApSupported: false

Dump of SarInfo
Current values:
    Voice Call state is: false
    Wifi Client state is: true
    Wifi Soft AP state is: false
    Wifi ScanOnly state is: false
    Earpiece state is : false
Last reported values:
    Soft AP state is: false
    Voice Call state is: false
    Earpiece state is: false
Last reported scenario: -2
Reported 1779531023 seconds ago

Dump of LastCallerInfoManager
API key=1 API name=ScanningEnabled: tid=792 uid=1000 pid=640 packageName=android toggleState=true
API key=33 API name=API_WIFI_SCANNER_START_SCAN: tid=792 uid=1000 pid=640 packageName=null toggleState=true

Dump of WifiNative
mIsLocationModeEnabled: true
mLastLocationModeEnabledTimeMs: 11561
Dump of HostapdHal
AIDL service declared: true
HIDL service declared: false
Initialized: false
Dump of SupplicantStaIfaceHal
Implemented: true
Dump of SupplicantStaIfaceHalAidlVendorImpl
Local Version: 5
Service Version: 5
mISupplicant: true
ifaces: [wlan0]


Dump of WifiRoamingConfigStore
DEVICE_ADMIN_POLICIES

NON_ADMIN_POLICIES

Dump of WifiResourceCache
WifiResourceCache - resource value Begin ----
Resource Name: config_wifi_revert_country_code_on_cellular_loss, value: false
Resource Name: config_wifi_connected_mac_randomization_supported, value: false
Resource Name: config_wifi6ghzSupport, value: false
Resource Name: config_wifiAfcSupported, value: false
Resource Name: config_wifiHardwareSoftapMaxClientCount, value: 16
Resource Name: config_wifiSaeUpgradeEnabled, value: true
Resource Name: config_wifi60ghzSupport, value: false
Resource Name: config_wifi24ghzSupport, value: true
Resource Name: config_wifiWepDeprecated, value: false
Resource Name: config_wifiOweUpgradeEnabled, value: true
Resource Name: config_wifiLowConnectedScoreThresholdToTriggerScanForMbb, value: 55
Resource Name: config_wifiWpaPersonalDeprecated, value: false
Resource Name: config_wifiSoftap24ghzSupported, value: true
Resource Name: config_wifiSoftapAcsIncludeDfs, value: false
Resource Name: config_wifiVerboseLoggingAlwaysOnLevel, value: 1
Resource Name: config_internalScorerType, value: 1
Resource Name: config_wifi5ghzSupport, value: false
Resource Name: config_wifiDriverSupportedNl80211RegChangedEvent, value: false
Resource Name: config_wifi_p2p_mac_randomization_supported, value: false
Resource Name: config_wifiSoftap2gChannelList, value: 1-11
Resource Name: config_wifiAdjustPollRssiIntervalEnabled, value: false
Resource Name: config_wifiWepAllowedControlSupported, value: true
Resource Name: config_wifiAllowInsecureEnterpriseConfigurationsForSettingsAndSUW, value: false
Resource Name: config_wifiD2dAllowedControlSupportedWhenInfraStaDisabled, value: false
WifiResourceCache - resource value End ----
boundToExternalScorer=failure, lastScorerBindingState=-1
Dump of WifiPowerStatsManager:
--- WifiPowerStats History ---
--- End of WifiPowerStats History ---
```
