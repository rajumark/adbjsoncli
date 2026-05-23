# `adbjson shell dumpsys telecom`

## adbjson

**Command:**
```bash
adbjson shell dumpsys telecom
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "(Total messages": "0, polling=true, quitting=false)",
    "2026-05-23T15:57:47.559275 - Got BluetoothHearingAid": "android.bluetooth.BluetoothHearingAid@bc245b8",
    "2026-05-23T15:57:47.559295 - Got BluetoothHeadset": "android.bluetooth.BluetoothHeadset@815101b",
    "2026-05-23T15:59:18.450125 - Got BluetoothHeadset": "android.bluetooth.BluetoothHeadset@815101b",
    "2026-05-23T15:59:18.450184 - Got BluetoothHearingAid": "android.bluetooth.BluetoothHearingAid@bc245b8",
    "Bound": "N",
    "Current call": "<none>",
    "DefaultCallRedirectionApp": "null",
    "DefaultCallScreeningApp": "null",
    "Init Path": "On mainline",
    "System Dialer": "ComponentInfo{com.google.android.dialer/com.android.incallui.InCallServiceImpl}",
    "TelecomUI package": "com.google.android.telecomui",
    "User 0": "com.google.android.dialer",
    "[[X] PhoneAccount": "ComponentInfo{com.android.phone/com.android.services.telephony.TelephonyConnectionService}, 1, UserHandle{0} Capabilities: CallProvider MultiUser PlaceEmerg SimSub  Audio Routes: BESW Schemes: tel voicemail  Extras: Bundle[{android.telecom.extra.SUPPORTS_VIDEO_CALLING_FALLBACK=true}] GroupId: *** SC Restrictions: [ ]]",
    "[✅]": "disconnectOnTimeoutWhileDisconnecting disconnect_on_timeout_while_disconnecting",
    "[❌]": "disconnectVoipOnHoldFail              disconnect_voip_on_hold_fail",
    "defaultOutgoing": "ComponentInfo{com.android.phone/com.android.services.telephony.TelephonyConnectionService}, 1, UserHandle{0}",
    "defaultVoiceSubId": "1",
    "isConnected": "false",
    "mCurrentTtyMode": "0",
    "outgoingPhoneAccountForTelScheme": "ComponentInfo{com.android.phone/com.android.services.telephony.TelephonyConnectionService}, 1, UserHandle{0}",
    "simCallManager": "null",
    "test emergency PhoneAccount filter": "[]",
    "xmlVersion": "10"
  }
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys telecom
```

**Output:**
```
Init Path: On mainline
TelecomUI package: com.google.android.telecomui
CallsManager: 
  mCalls: 
  mCallAudioManager:
    All calls:
    Active dialing, or connecting calls:
    Ringing calls:
    Holding calls:
    Foreground call:
    null
    CallAudioModeStateMachine:
      History:
      2026-05-23T15:57:47.269958 - Enter UNFOCUSED
      Pending Msg:
      Looper (CallAudioModeStateMachine, tid 224) {5d8dc2}
        (MessageQueue is using DeliQueue implementation)
        (Total messages: 0, polling=true, quitting=false)
    mCallAudioRouteAdapter:
    BluetoothDeviceManager:
      2026-05-23T15:57:47.559275 - Got BluetoothHearingAid: android.bluetooth.BluetoothHearingAid@bc245b8
      2026-05-23T15:57:47.559295 - Got BluetoothHeadset: android.bluetooth.BluetoothHeadset@815101b
      2026-05-23T15:59:06.096475 - Lost BluetoothHearingAid service. Removing all tracked devices.
      2026-05-23T15:59:06.096491 - Lost BluetoothHeadset service. Removing all tracked devices
      2026-05-23T15:59:18.450125 - Got BluetoothHeadset: android.bluetooth.BluetoothHeadset@815101b
      2026-05-23T15:59:18.450184 - Got BluetoothHearingAid: android.bluetooth.BluetoothHearingAid@bc245b8
      2026-05-23T15:59:41.794530 - Lost BluetoothHearingAid service. Removing all tracked devices.
      2026-05-23T15:59:41.794571 - Lost BluetoothHeadset service. Removing all tracked devices
  mTtyManager:
    mCurrentTtyMode: 0
  mInCallController:
    combinedInCallServiceMap (InCalls registered):
    ServiceConnections (InCalls bound):
    CarModeTracker:
      Current car mode apps:
      Car mode history:
  mCallDiagnosticServiceController:
    activeCallDiagnosticService: 
    isConnected: false
  mAudioModeTracker:
    Audio Mode History:
  mLocalVoicemailController:
    Current call: <none>
    Service pkg: 
    Bound: N
    Local voicemail History:
  mCallAnomalyWatchdog:
    Anomaly log:
    Pending timeouts: 
    Pending destruction: 
  mEmergencyCallDiagnosticLogger:
      PERSISTED DIAGNOSTIC DATA FROM DROP BOX
      END OF PERSISTED DIAGNOSTIC DATA FROM DROP BOX
  mDefaultDialerCache:
    System Dialer: ComponentInfo{com.google.android.dialer/com.android.incallui.InCallServiceImpl}
    User 0: com.google.android.dialer
  mConnectionServiceRepository:
    mServiceCache:
  mRoleManager:
    DefaultCallRedirectionApp: null
    DefaultCallScreeningApp: null
    DefaultCallCompanionApps: 
  mConnectionSvrFocusMgr:
    Call Focus History:
  CallAudioWatchdog:
    Active Sessions:
    Audio sessions Sessions:
PhoneAccountRegistrar: 
  xmlVersion: 10
  defaultOutgoing: ComponentInfo{com.android.phone/com.android.services.telephony.TelephonyConnectionService}, 1, UserHandle{0}
  outgoingPhoneAccountForTelScheme: ComponentInfo{com.android.phone/com.android.services.telephony.TelephonyConnectionService}, 1, UserHandle{0}
  defaultVoiceSubId: 1
  simCallManager: null
  phoneAccounts:
    [[X] PhoneAccount: ComponentInfo{com.android.phone/com.android.services.telephony.TelephonyConnectionService}, 1, UserHandle{0} Capabilities: CallProvider MultiUser PlaceEmerg SimSub  Audio Routes: BESW Schemes: tel voicemail  Extras: Bundle[{android.telecom.extra.SUPPORTS_VIDEO_CALLING_FALLBACK=true}] GroupId: *** SC Restrictions: [ ]]
    test emergency PhoneAccount filter: []
  localVoicemailTimeouts:
Flag Configurations (framework - com.android.server.telecom): 
  	[✅]: addCallUriForMissedCalls                     add_call_uri_for_missed_calls
  	[✅]: bulkStateUpdateCall                          bulk_state_update_call
  	[✅]: businessCallComposer                         business_call_composer
  	[✅]: callDetailsIdChanges                         call_details_id_changes
  	[✅]: callSequencingCallResumeFailed               call_sequencing_call_resume_failed
  	[❌]: callSequencingMetrics                        call_sequencing_metrics
  	[✅]: cleanupVerifyCallState                       cleanup_verify_call_state
  	[✅]: conferenceModifyMergeFail                    conference_modify_merge_fail
  	[✅]: doNotSendCallToNullIcs                       do_not_send_call_to_null_ics
  	[✅]: ensureAudioModeUpdatesOnForegroundCallChange ensure_audio_mode_updates_on_foreground_call_change
  	[✅]: filterVoipCallLogs                           filter_voip_call_logs
  	[✅]: getRegisteredPhoneAccounts                   get_registered_phone_accounts
  	[✅]: hdPlusCall                                   hd_plus_call
  	[✅]: integratedCallLogs                           integrated_call_logs
  	[✅]: notifyAvailableEndpointsOnIcsConnected       notify_available_endpoints_on_ics_connected
  	[✅]: preventIllegalAudioProcessingExit            prevent_illegal_audio_processing_exit
  	[✅]: reuseOriginalConnRemoteConfApi               reuse_original_conn_remote_conf_api
  	[✅]: revertDisconnectingDuringMerge               revert_disconnecting_during_merge
  	[✅]: selectPhoneAccountBeforeMakingRoom           select_phone_account_before_making_room
  	[✅]: setMuteState                                 set_mute_state
  	[✅]: supportDisplayNameCallLog                    support_display_name_call_log
  	[✅]: synchronizeConnState                         synchronize_conn_state
  	[❌]: telecomAppLabelProxyHsumAware                telecom_app_label_proxy_hsum_aware
  	[✅]: telecomMainUserInBlockCheck                  telecom_main_user_in_block_check
  	[✅]: telecomMainUserInGetRespondMessageApp        telecom_main_user_in_get_respond_message_app
  	[✅]: telecomMainlineBlockedNumbersManager         telecom_mainline_blocked_numbers_manager
  	[✅]: telecomResolveHiddenDependencies             telecom_resolve_hidden_dependencies
  	[✅]: telephonyHasDefaultButTelecomDoesNot         telephony_has_default_but_telecom_does_not
  	[✅]: transactionalCsVerifier                      transactional_cs_verifier
  	[✅]: transactionalVideoState                      transactional_video_state
  	[✅]: unregisterUnresolvableAccounts               unregister_unresolvable_accounts
  	[✅]: useDeviceProvidedSerializedRingerVibration   use_device_provided_serialized_ringer_vibration
  	[❌]: vibrationAccountsForMainSetting              vibration_accounts_for_main_setting
  	[✅]: voipAppActionsSupport                        voip_app_actions_support
  	[❌]: voipDndFocus                                 voip_dnd_focus
Flag Configurations (module API - android.telecom): 
  	[✅]: callConnectedIndicatorPreference    call_connected_indicator_preference
  	[✅]: callDetailsGetAssociatedUserApi2    call_details_get_associated_user_api2
  	[✅]: callEndpointRequestedApi            call_endpoint_requested_api
  	[✅]: changeRttToAudio                    change_rtt_to_audio
  	[✅]: deprecateSelfManagedCs              deprecate_self_managed_cs
  	[✅]: enableAudioProcessingUseCase        enable_audio_processing_use_case
  	[✅]: explicitCallTransfer                explicit_call_transfer
  	[✅]: filterVoipCallLogs                  filter_voip_call_logs
  	[✅]: getLastKnownCellIdentity            get_last_known_cell_identity
  	[✅]: integratedCallLogsStage2            integrated_call_logs_stage2
  	[✅]: isInExternalCall                    is_in_external_call
  	[✅]: isUsingCrs                          is_using_crs
  	[✅]: isUsingUnidirectionalVideoService   is_using_unidirectional_video_service
  	[✅]: isUsingVideoRingback                is_using_video_ringback
  	[✅]: localVoicemail                      local_voicemail
  	[✅]: moveGetTtyModeToTelephonyManager    move_get_tty_mode_to_telephony_manager
  	[✅]: multiPartyAnchorConf                multi_party_anchor_conf
  	[❌]: optOutPremiumNetwork                opt_out_premium_network
  	[✅]: phoneAccountChanged                 phone_account_changed
  	[✅]: placeCallToAlternateNumber          place_call_to_alternate_number
  	[✅]: promoteExtraDoNotLogCallToSystemApi promote_extra_do_not_log_call_to_system_api
  	[✅]: releaseIconAsApi                    release_icon_as_api
  	[✅]: remotelyHostedProperty              remotely_hosted_property
  	[✅]: sendOriginalNumberOnPlaceCall       send_original_number_on_place_call
  	[✅]: telecomMainlineApi                  telecom_mainline_api
Flag Configurations (module bugfix - com.android.internal.telecom): 
  	[❌]: addEscapeHatchForStuckVoip            add_escape_hatch_for_stuck_voip
  	[❌]: callAudioRouteRf                      call_audio_route_rf
  	[❌]: cleanupCallsInSelectAccount           cleanup_calls_in_select_account
  	[✅]: connectionServiceBal                  connection_service_bal
  	[✅]: delayRequestedHandleSelection         delay_requested_handle_selection
  	[✅]: disconnectOnTimeoutWhileDisconnecting disconnect_on_timeout_while_disconnecting
  	[❌]: disconnectVoipOnHoldFail              disconnect_voip_on_hold_fail
TransactionManager: 
  Pending Transactions:
  Ongoing Transaction:
  Completed Transactions:
Historical Events:
```
