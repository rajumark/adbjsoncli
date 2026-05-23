# `adbjson shell service list`

## adbjson

**Command:**
```bash
adbjson shell service list
```

**Output:**
```json
{
  "status": 0,
  "output": {
    "service_count": 327,
    "services": [
      {
        "id": "0",
        "interface": "",
        "name": "DockObserver"
      },
      {
        "id": "1",
        "interface": "android.ui.ISurfaceComposer",
        "name": "SurfaceFlinger"
      },
      {
        "id": "2",
        "interface": "android.gui.ISurfaceComposer",
        "name": "SurfaceFlingerAIDL"
      },
      {
        "id": "3",
        "interface": "android.view.accessibility.IAccessibilityManager",
        "name": "accessibility"
      },
      {
        "id": "4",
        "interface": "android.accounts.IAccountManager",
        "name": "account"
      },
      {
        "id": "5",
        "interface": "android.app.IActivityManager",
        "name": "activity"
      },
      {
        "id": "6",
        "interface": "android.app.IActivityManagerStructured",
        "name": "activity_structured"
      },
      {
        "id": "7",
        "interface": "android.app.IActivityTaskManager",
        "name": "activity_task"
      },
      {
        "id": "8",
        "interface": "android.debug.IAdbManager",
        "name": "adb"
      },
      {
        "id": "9",
        "interface": "android.app.adservices.IAdServicesManager",
        "name": "adservices_manager"
      },
      {
        "id": "10",
        "interface": "android.security.advancedprotection.IAdvancedProtectionService",
        "name": "advanced_protection"
      },
      {
        "id": "11",
        "interface": "android.app.IAlarmManager",
        "name": "alarm"
      },
      {
        "id": "12",
        "interface": "android.os.allowlist.IAllowlistService",
        "name": "allowlist"
      },
      {
        "id": "13",
        "interface": "android.app.ambientcontext.IAmbientContextManager",
        "name": "ambient_context"
      },
      {
        "id": "14",
        "interface": "android.frameworks.cameraservice.service.ICameraService",
        "name": "android.frameworks.cameraservice.service.ICameraService/default"
      },
      {
        "id": "15",
        "interface": "android.frameworks.devicestate.IDeviceStateService",
        "name": "android.frameworks.devicestate.IDeviceStateService/default"
      },
      {
        "id": "16",
        "interface": "android.frameworks.location.altitude.IAltitudeService",
        "name": "android.frameworks.location.altitude.IAltitudeService/default"
      },
      {
        "id": "17",
        "interface": "android.frameworks.sensorservice.ISensorManager",
        "name": "android.frameworks.sensorservice.ISensorManager/default"
      },
      {
        "id": "18",
        "interface": "android.frameworks.stats.IStats",
        "name": "android.frameworks.stats.IStats/default"
      },
      {
        "id": "19",
        "interface": "android.frameworks.vibrator.IVibratorControlService",
        "name": "android.frameworks.vibrator.IVibratorControlService/default"
      },
      {
        "id": "20",
        "interface": "",
        "name": "android.hardware.authsecret.IAuthSecret/default"
      },
      {
        "id": "21",
        "interface": "",
        "name": "android.hardware.biometrics.fingerprint.IFingerprint/default"
      },
      {
        "id": "22",
        "interface": "",
        "name": "android.hardware.bluetooth.IBluetoothHci/default"
      },
      {
        "id": "23",
        "interface": "",
        "name": "android.hardware.bluetooth.audio.IBluetoothAudioProviderFactory/default"
      },
      {
        "id": "24",
        "interface": "",
        "name": "android.hardware.bluetooth.offload.leaudio.IHciProxy/default"
      },
      {
        "id": "25",
        "interface": "",
        "name": "android.hardware.camera.provider.ICameraProvider/internal/0"
      },
      {
        "id": "26",
        "interface": "",
        "name": "android.hardware.camera.provider.ICameraProvider/internal/1"
      },
      {
        "id": "27",
        "interface": "android.hardware.cas.IMediaCasService",
        "name": "android.hardware.cas.IMediaCasService/default"
      },
      {
        "id": "28",
        "interface": "",
        "name": "android.hardware.contexthub.IContextHub/default"
      },
      {
        "id": "29",
        "interface": "android.hardware.drm.IDrmFactory",
        "name": "android.hardware.drm.IDrmFactory/widevine"
      },
      {
        "id": "30",
        "interface": "",
        "name": "android.hardware.gatekeeper.IGatekeeper/default"
      },
      {
        "id": "31",
        "interface": "",
        "name": "android.hardware.gnss.IGnss/default"
      },
      {
        "id": "32",
        "interface": "android.hardware.graphics.allocator.IAllocator",
        "name": "android.hardware.graphics.allocator.IAllocator/default"
      },
      {
        "id": "33",
        "interface": "",
        "name": "android.hardware.graphics.composer3.IComposer/default"
      },
      {
        "id": "34",
        "interface": "",
        "name": "android.hardware.health.IHealth/default"
      },
      {
        "id": "35",
        "interface": "",
        "name": "android.hardware.identity.IIdentityCredentialStore/default"
      },
      {
        "id": "36",
        "interface": "",
        "name": "android.hardware.light.ILights/default"
      },
      {
        "id": "37",
        "interface": "android.hardware.media.c2.IComponentStore",
        "name": "android.hardware.media.c2.IComponentStore/default"
      },
      {
        "id": "38",
        "interface": "android.hardware.media.c2.IComponentStore",
        "name": "android.hardware.media.c2.IComponentStore/software"
      },
      {
        "id": "39",
        "interface": "android.hardware.neuralnetworks.IDevice",
        "name": "android.hardware.neuralnetworks.IDevice/nnapi-sample_all"
      },
      {
        "id": "40",
        "interface": "android.hardware.neuralnetworks.IDevice",
        "name": "android.hardware.neuralnetworks.IDevice/nnapi-sample_quant"
      },
      {
        "id": "41",
        "interface": "android.hardware.neuralnetworks.IDevice",
        "name": "android.hardware.neuralnetworks.IDevice/nnapi-sample_sl_shim"
      },
      {
        "id": "42",
        "interface": "",
        "name": "android.hardware.power.IPower/default"
      },
      {
        "id": "43",
        "interface": "",
        "name": "android.hardware.power.stats.IPowerStats/default"
      },
      {
        "id": "44",
        "interface": "",
        "name": "android.hardware.radio.config.IRadioConfig/default"
      },
      {
        "id": "45",
        "interface": "",
        "name": "android.hardware.radio.data.IRadioData/slot1"
      },
      {
        "id": "46",
        "interface": "",
        "name": "android.hardware.radio.ims.IRadioIms/slot1"
      },
      {
        "id": "47",
        "interface": "",
        "name": "android.hardware.radio.ims.media.IImsMedia/default"
      },
      {
        "id": "48",
        "interface": "",
        "name": "android.hardware.radio.messaging.IRadioMessaging/slot1"
      },
      {
        "id": "49",
        "interface": "",
        "name": "android.hardware.radio.modem.IRadioModem/slot1"
      },
      {
        "id": "50",
        "interface": "",
        "name": "android.hardware.radio.network.IRadioNetwork/slot1"
      },
      {
        "id": "51",
        "interface": "",
        "name": "android.hardware.radio.sap.ISap/slot1"
      },
      {
        "id": "52",
        "interface": "",
        "name": "android.hardware.radio.sim.IRadioSim/slot1"
      },
      {
        "id": "53",
        "interface": "",
        "name": "android.hardware.radio.voice.IRadioVoice/slot1"
      },
      {
        "id": "54",
        "interface": "",
        "name": "android.hardware.rebootescrow.IRebootEscrow/default"
      },
      {
        "id": "55",
        "interface": "",
        "name": "android.hardware.security.keymint.IKeyMintDevice/default"
      },
      {
        "id": "56",
        "interface": "android.hardware.security.keymint.IRemotelyProvisionedComponent",
        "name": "android.hardware.security.keymint.IRemotelyProvisionedComponent/default"
      },
      {
        "id": "57",
        "interface": "",
        "name": "android.hardware.security.secureclock.ISecureClock/default"
      },
      {
        "id": "58",
        "interface": "",
        "name": "android.hardware.security.sharedsecret.ISharedSecret/default"
      },
      {
        "id": "59",
        "interface": "",
        "name": "android.hardware.security.sharedsecret.ISharedSecret/gatekeeper"
      },
      {
        "id": "60",
        "interface": "",
        "name": "android.hardware.sensors.ISensors/default"
      },
      {
        "id": "61",
        "interface": "",
        "name": "android.hardware.thermal.IThermal/default"
      },
      {
        "id": "62",
        "interface": "",
        "name": "android.hardware.threadnetwork.IThreadChip/chip0"
      },
      {
        "id": "63",
        "interface": "",
        "name": "android.hardware.usb.IUsb/default"
      },
      {
        "id": "64",
        "interface": "",
        "name": "android.hardware.uwb.IUwb/default"
      },
      {
        "id": "65",
        "interface": "",
        "name": "android.hardware.vibrator.IVibrator/default"
      },
      {
        "id": "66",
        "interface": "",
        "name": "android.hardware.vibrator.IVibratorManager/default"
      },
      {
        "id": "67",
        "interface": "",
        "name": "android.hardware.wifi.IWifi/default"
      },
      {
        "id": "68",
        "interface": "",
        "name": "android.hardware.wifi.supplicant.ISupplicant/default"
      },
      {
        "id": "69",
        "interface": "android.security.apc.IProtectedConfirmation",
        "name": "android.security.apc"
      },
      {
        "id": "70",
        "interface": "android.security.authorization.IKeystoreAuthorization",
        "name": "android.security.authorization"
      },
      {
        "id": "71",
        "interface": "android.security.compat.IKeystoreCompatService",
        "name": "android.security.compat"
      },
      {
        "id": "72",
        "interface": "android.security.identity.ICredentialStoreFactory",
        "name": "android.security.identity"
      },
      {
        "id": "73",
        "interface": "android.security.legacykeystore.ILegacyKeystore",
        "name": "android.security.legacykeystore"
      },
      {
        "id": "74",
        "interface": "android.security.maintenance.IKeystoreMaintenance",
        "name": "android.security.maintenance"
      },
      {
        "id": "75",
        "interface": "android.security.metrics.IKeystoreMetrics",
        "name": "android.security.metrics"
      },
      {
        "id": "76",
        "interface": "",
        "name": "android.service.gatekeeper.IGateKeeperService"
      },
      {
        "id": "77",
        "interface": "android.system.keystore2.IKeystoreService",
        "name": "android.system.keystore2.IKeystoreService/default"
      },
      {
        "id": "78",
        "interface": "",
        "name": "android.system.net.netd.INetd/default"
      },
      {
        "id": "79",
        "interface": "",
        "name": "android.system.suspend.ISystemSuspend/default"
      },
      {
        "id": "80",
        "interface": "",
        "name": "android.system.vold.IVold/default"
      },
      {
        "id": "81",
        "interface": "android.os.profiling.anomaly.IAnomalyDetectorService",
        "name": "anomaly_detector"
      },
      {
        "id": "82",
        "interface": "",
        "name": "app_binding"
      },
      {
        "id": "83",
        "interface": "android.app.appfunctions.IAppFunctionManager",
        "name": "app_function"
      },
      {
        "id": "84",
        "interface": "android.apphibernation.IAppHibernationService",
        "name": "app_hibernation"
      },
      {
        "id": "85",
        "interface": "android.content.integrity.IAppIntegrityManager",
        "name": "app_integrity"
      },
      {
        "id": "86",
        "interface": "android.app.prediction.IPredictionManager",
        "name": "app_prediction"
      },
      {
        "id": "87",
        "interface": "android.app.appsearch.aidl.IAppSearchManager",
        "name": "app_search"
      },
      {
        "id": "88",
        "interface": "com.android.internal.app.IAppOpsService",
        "name": "appops"
      },
      {
        "id": "89",
        "interface": "com.android.internal.appwidget.IAppWidgetService",
        "name": "appwidget"
      },
      {
        "id": "90",
        "interface": "",
        "name": "artd"
      },
      {
        "id": "91",
        "interface": "android.security.attestationverification.IAttestationVerificationManagerService",
        "name": "attestation_verification"
      },
      {
        "id": "92",
        "interface": "android.media.IAudioService",
        "name": "audio"
      },
      {
        "id": "93",
        "interface": "android.hardware.biometrics.IAuthService",
        "name": "auth"
      },
      {
        "id": "94",
        "interface": "android.security.authenticationpolicy.IAuthenticationPolicyService",
        "name": "authentication_policy"
      },
      {
        "id": "95",
        "interface": "android.view.autofill.IAutoFillManager",
        "name": "autofill"
      },
      {
        "id": "96",
        "interface": "android.content.pm.IBackgroundInstallControlService",
        "name": "background_install_control"
      },
      {
        "id": "97",
        "interface": "android.app.backup.IBackupManager",
        "name": "backup"
      },
      {
        "id": "98",
        "interface": "",
        "name": "battery"
      },
      {
        "id": "99",
        "interface": "android.os.IBatteryPropertiesRegistrar",
        "name": "batteryproperties"
      },
      {
        "id": "100",
        "interface": "com.android.internal.app.IBatteryStats",
        "name": "batterystats"
      },
      {
        "id": "101",
        "interface": "",
        "name": "binder_calls_stats"
      },
      {
        "id": "102",
        "interface": "android.os.binder.IBinderStatsConsumerService",
        "name": "binder_stats_consumer"
      },
      {
        "id": "103",
        "interface": "android.hardware.biometrics.IBiometricService",
        "name": "biometric"
      },
      {
        "id": "104",
        "interface": "android.app.blob.IBlobStoreManager",
        "name": "blob_store"
      },
      {
        "id": "105",
        "interface": "android.bluetooth.IBluetoothManager",
        "name": "bluetooth_manager"
      },
      {
        "id": "106",
        "interface": "android.os.IDumpstate",
        "name": "bugreport"
      },
      {
        "id": "107",
        "interface": "",
        "name": "cacheinfo"
      },
      {
        "id": "108",
        "interface": "com.android.internal.telephony.ICarrierConfigLoader",
        "name": "carrier_config"
      },
      {
        "id": "109",
        "interface": "android.content.IClipboard",
        "name": "clipboard"
      },
      {
        "id": "110",
        "interface": "android.hardware.display.IColorDisplayManager",
        "name": "color_display"
      },
      {
        "id": "111",
        "interface": "android.companion.ICompanionDeviceManager",
        "name": "companiondevice"
      },
      {
        "id": "112",
        "interface": "android.net.IConnectivityManager",
        "name": "connectivity"
      },
      {
        "id": "113",
        "interface": "android.net.connectivity.aidl.ConnectivityNative",
        "name": "connectivity_native"
      },
      {
        "id": "114",
        "interface": "android.net.IIpConnectivityMetrics",
        "name": "connmetrics"
      },
      {
        "id": "115",
        "interface": "android.content.IContentService",
        "name": "content"
      },
      {
        "id": "116",
        "interface": "android.view.contentcapture.IContentCaptureManager",
        "name": "content_capture"
      },
      {
        "id": "117",
        "interface": "android.app.contentsafety.IContentSafetyManager",
        "name": "content_safety"
      },
      {
        "id": "118",
        "interface": "android.app.contentsuggestions.IContentSuggestionsManager",
        "name": "content_suggestions"
      },
      {
        "id": "119",
        "interface": "android.hardware.location.IContextHubService",
        "name": "contexthub"
      },
      {
        "id": "120",
        "interface": "android.app.modes.IContextualModeManager",
        "name": "contextual_mode"
      },
      {
        "id": "121",
        "interface": "android.location.ICountryDetector",
        "name": "country_detector"
      },
      {
        "id": "122",
        "interface": "",
        "name": "cpuinfo"
      },
      {
        "id": "123",
        "interface": "android.credentials.ICredentialManager",
        "name": "credential"
      },
      {
        "id": "124",
        "interface": "android.content.pm.ICrossProfileApps",
        "name": "crossprofileapps"
      },
      {
        "id": "125",
        "interface": "android.content.pm.IDataLoaderManager",
        "name": "dataloader_manager"
      },
      {
        "id": "126",
        "interface": "",
        "name": "dbinfo"
      },
      {
        "id": "127",
        "interface": "",
        "name": "device_config"
      },
      {
        "id": "128",
        "interface": "android.os.IDeviceIdentifiersPolicyService",
        "name": "device_identifiers"
      },
      {
        "id": "129",
        "interface": "android.devicelock.IDeviceLockService",
        "name": "device_lock"
      },
      {
        "id": "130",
        "interface": "android.app.admin.IDevicePolicyManager",
        "name": "device_policy"
      },
      {
        "id": "131",
        "interface": "android.hardware.devicestate.IDeviceStateManager",
        "name": "device_state"
      },
      {
        "id": "132",
        "interface": "android.os.IDeviceIdleController",
        "name": "deviceidle"
      },
      {
        "id": "133",
        "interface": "",
        "name": "devicestoragemonitor"
      },
      {
        "id": "134",
        "interface": "",
        "name": "diskstats"
      },
      {
        "id": "135",
        "interface": "android.hardware.display.IDisplayManager",
        "name": "display"
      },
      {
        "id": "136",
        "interface": "",
        "name": "dnsresolver"
      },
      {
        "id": "137",
        "interface": "android.content.pm.verify.domain.IDomainVerificationManager",
        "name": "domain_verification"
      },
      {
        "id": "138",
        "interface": "android.service.dreams.IDreamManager",
        "name": "dreams"
      },
      {
        "id": "139",
        "interface": "drm.IDrmManagerService",
        "name": "drm.drmManager"
      },
      {
        "id": "140",
        "interface": "com.android.internal.os.IDropBoxManagerService",
        "name": "dropbox"
      },
      {
        "id": "141",
        "interface": "android.os.instrumentation.IDynamicInstrumentationManager",
        "name": "dynamic_instrumentation"
      },
      {
        "id": "142",
        "interface": "android.os.image.IDynamicSystemService",
        "name": "dynamic_system"
      },
      {
        "id": "143",
        "interface": "android.app.ecm.IEnhancedConfirmationManager",
        "name": "ecm_enhanced_confirmation"
      },
      {
        "id": "144",
        "interface": "",
        "name": "emergency_affordance"
      },
      {
        "id": "145",
        "interface": "android.os.IExternalVibratorService",
        "name": "external_vibrator_service"
      },
      {
        "id": "146",
        "interface": "android.flags.IFeatureFlags",
        "name": "feature_flags"
      },
      {
        "id": "147",
        "interface": "android.os.storage.IFileService",
        "name": "file"
      },
      {
        "id": "148",
        "interface": "android.security.IFileIntegrityService",
        "name": "file_integrity"
      },
      {
        "id": "149",
        "interface": "android.hardware.fingerprint.IFingerprintService",
        "name": "fingerprint"
      },
      {
        "id": "150",
        "interface": "com.android.internal.graphics.fonts.IFontManager",
        "name": "font"
      },
      {
        "id": "151",
        "interface": "android.app.IGameManagerService",
        "name": "game"
      },
      {
        "id": "152",
        "interface": "",
        "name": "gfxinfo"
      },
      {
        "id": "153",
        "interface": "android.graphicsenv.IGpuService",
        "name": "gpu"
      },
      {
        "id": "154",
        "interface": "android.app.IGrammaticalInflectionManager",
        "name": "grammatical_inflection"
      },
      {
        "id": "155",
        "interface": "android.view.IGraphicsStats",
        "name": "graphicsstats"
      },
      {
        "id": "156",
        "interface": "android.os.IHardwarePropertiesManager",
        "name": "hardware_properties"
      },
      {
        "id": "157",
        "interface": "android.health.connect.aidl.IHealthConnectService",
        "name": "healthconnect"
      },
      {
        "id": "158",
        "interface": "com.android.internal.telephony.IMms",
        "name": "imms"
      },
      {
        "id": "159",
        "interface": "",
        "name": "incident"
      },
      {
        "id": "160",
        "interface": "android.os.IIncidentCompanion",
        "name": "incidentcompanion"
      },
      {
        "id": "161",
        "interface": "android.os.incremental.IIncrementalService",
        "name": "incremental"
      },
      {
        "id": "162",
        "interface": "android.hardware.input.IInputManager",
        "name": "input"
      },
      {
        "id": "163",
        "interface": "com.android.internal.view.IInputMethodManager",
        "name": "input_method"
      },
      {
        "id": "164",
        "interface": "android.os.IInputFlinger",
        "name": "inputflinger"
      },
      {
        "id": "165",
        "interface": "",
        "name": "installd"
      },
      {
        "id": "166",
        "interface": "android.security.intrusiondetection.IIntrusionDetectionService",
        "name": "intrusion_detection"
      },
      {
        "id": "167",
        "interface": "com.android.internal.telephony.IOns",
        "name": "ions"
      },
      {
        "id": "168",
        "interface": "com.android.internal.telephony.IPhoneSubInfo",
        "name": "iphonesubinfo"
      },
      {
        "id": "169",
        "interface": "android.net.IIpSecService",
        "name": "ipsec"
      },
      {
        "id": "170",
        "interface": "com.android.internal.telephony.ISms",
        "name": "isms"
      },
      {
        "id": "171",
        "interface": "com.android.internal.telephony.ISub",
        "name": "isub"
      },
      {
        "id": "172",
        "interface": "android.app.job.IJobScheduler",
        "name": "jobscheduler"
      },
      {
        "id": "173",
        "interface": "android.content.pm.ILauncherApps",
        "name": "launcherapps"
      },
      {
        "id": "174",
        "interface": "android.permission.ILegacyPermissionManager",
        "name": "legacy_permission"
      },
      {
        "id": "175",
        "interface": "android.hardware.lights.ILightsManager",
        "name": "lights"
      },
      {
        "id": "176",
        "interface": "android.app.ILocaleManager",
        "name": "locale"
      },
      {
        "id": "177",
        "interface": "android.location.ILocationManager",
        "name": "location"
      },
      {
        "id": "178",
        "interface": "",
        "name": "location_time_zone_manager"
      },
      {
        "id": "179",
        "interface": "com.android.internal.widget.ILockSettings",
        "name": "lock_settings"
      },
      {
        "id": "180",
        "interface": "android.os.logcat.ILogcatManagerService",
        "name": "logcat"
      },
      {
        "id": "181",
        "interface": "",
        "name": "looper_stats"
      },
      {
        "id": "182",
        "interface": "android.os.IServiceManager",
        "name": "manager"
      },
      {
        "id": "183",
        "interface": "",
        "name": "mdns"
      },
      {
        "id": "184",
        "interface": "android.media.IAudioFlingerService",
        "name": "media.audio_flinger"
      },
      {
        "id": "185",
        "interface": "android.media.IAudioPolicyService",
        "name": "media.audio_policy"
      },
      {
        "id": "186",
        "interface": "android.hardware.ICameraService",
        "name": "media.camera"
      },
      {
        "id": "187",
        "interface": "android.hardware.ICameraServiceProxy",
        "name": "media.camera.proxy"
      },
      {
        "id": "188",
        "interface": "android.IMediaExtractorService",
        "name": "media.extractor"
      },
      {
        "id": "189",
        "interface": "android.media.IMediaMetricsService",
        "name": "media.metrics"
      },
      {
        "id": "190",
        "interface": "android.media.IMediaPlayerService",
        "name": "media.player"
      },
      {
        "id": "191",
        "interface": "android.media.IResourceManagerService",
        "name": "media.resource_manager"
      },
      {
        "id": "192",
        "interface": "android.media.IResourceObserverService",
        "name": "media.resource_observer"
      },
      {
        "id": "193",
        "interface": "android.media.IMediaCommunicationService",
        "name": "media_communication"
      },
      {
        "id": "194",
        "interface": "android.media.metrics.IMediaMetricsManager",
        "name": "media_metrics"
      },
      {
        "id": "195",
        "interface": "android.media.projection.IMediaProjectionManager",
        "name": "media_projection"
      },
      {
        "id": "196",
        "interface": "android.media.IMediaResourceMonitor",
        "name": "media_resource_monitor"
      },
      {
        "id": "197",
        "interface": "android.media.IMediaRouterService",
        "name": "media_router"
      },
      {
        "id": "198",
        "interface": "android.media.session.ISessionManager",
        "name": "media_session"
      },
      {
        "id": "199",
        "interface": "",
        "name": "meminfo"
      },
      {
        "id": "200",
        "interface": "android.hardware.memtrack.IMemtrack",
        "name": "memtrack.proxy"
      },
      {
        "id": "201",
        "interface": "android.media.midi.IMidiManager",
        "name": "midi"
      },
      {
        "id": "202",
        "interface": "android.os.storage.IStorageManager",
        "name": "mount"
      },
      {
        "id": "203",
        "interface": "android.media.musicrecognition.IMusicRecognitionManager",
        "name": "music_recognition"
      },
      {
        "id": "204",
        "interface": "android.nearby.INearbyManager",
        "name": "nearby"
      },
      {
        "id": "205",
        "interface": "",
        "name": "netd"
      },
      {
        "id": "206",
        "interface": "android.net.metrics.INetdEventListener",
        "name": "netd_listener"
      },
      {
        "id": "207",
        "interface": "android.net.INetworkPolicyManager",
        "name": "netpolicy"
      },
      {
        "id": "208",
        "interface": "android.net.INetworkStatsService",
        "name": "netstats"
      },
      {
        "id": "209",
        "interface": "android.os.INetworkManagementService",
        "name": "network_management"
      },
      {
        "id": "210",
        "interface": "android.net.INetworkScoreService",
        "name": "network_score"
      },
      {
        "id": "211",
        "interface": "android.net.INetworkStackConnector",
        "name": "network_stack"
      },
      {
        "id": "212",
        "interface": "",
        "name": "network_time_update_service"
      },
      {
        "id": "213",
        "interface": "com.android.internal.net.INetworkWatchlistManager",
        "name": "network_watchlist"
      },
      {
        "id": "214",
        "interface": "android.app.INotificationManager",
        "name": "notification"
      },
      {
        "id": "215",
        "interface": "android.app.ondeviceintelligence.IOnDeviceIntelligenceManager",
        "name": "on_device_intelligence"
      },
      {
        "id": "216",
        "interface": "android.ondevicepersonalization.IOnDevicePersonalizationSystemService",
        "name": "ondevicepersonalization_system_service"
      },
      {
        "id": "217",
        "interface": "",
        "name": "ot_daemon"
      },
      {
        "id": "218",
        "interface": "android.content.om.IOverlayManager",
        "name": "overlay"
      },
      {
        "id": "219",
        "interface": "android.net.IPacProxyManager",
        "name": "pac_proxy"
      },
      {
        "id": "220",
        "interface": "android.content.pm.IPackageManager",
        "name": "package"
      },
      {
        "id": "221",
        "interface": "android.content.pm.IPackageManagerNative",
        "name": "package_native"
      },
      {
        "id": "222",
        "interface": "android.app.privatecompute.IPccSandboxManager",
        "name": "pcc_sandbox"
      },
      {
        "id": "223",
        "interface": "android.app.privatecompute.IPccSandboxManagerNative",
        "name": "pcc_sandbox_native"
      },
      {
        "id": "224",
        "interface": "android.app.people.IPeopleManager",
        "name": "people"
      },
      {
        "id": "225",
        "interface": "android.os.IHintManager",
        "name": "performance_hint"
      },
      {
        "id": "226",
        "interface": "android.os.IPermissionController",
        "name": "permission"
      },
      {
        "id": "227",
        "interface": "android.permission.IPermissionChecker",
        "name": "permission_checker"
      },
      {
        "id": "228",
        "interface": "android.permission.IPermissionManager",
        "name": "permissionmgr"
      },
      {
        "id": "229",
        "interface": "com.android.internal.telephony.ITelephony",
        "name": "phone"
      },
      {
        "id": "230",
        "interface": "android.app.pinner.IPinnerService",
        "name": "pinner"
      },
      {
        "id": "231",
        "interface": "com.android.internal.compat.IPlatformCompat",
        "name": "platform_compat"
      },
      {
        "id": "232",
        "interface": "com.android.internal.compat.IPlatformCompatNative",
        "name": "platform_compat_native"
      },
      {
        "id": "233",
        "interface": "android.os.IPowerManager",
        "name": "power"
      },
      {
        "id": "234",
        "interface": "android.os.IPowerStatsService",
        "name": "powerstats"
      },
      {
        "id": "235",
        "interface": "android.print.IPrintManager",
        "name": "print"
      },
      {
        "id": "236",
        "interface": "android.os.IProcessInfoService",
        "name": "processinfo"
      },
      {
        "id": "237",
        "interface": "com.android.internal.app.procstats.IProcessStats",
        "name": "procstats"
      },
      {
        "id": "238",
        "interface": "android.os.IProfilingService",
        "name": "profiling_service"
      },
      {
        "id": "239",
        "interface": "com.android.internal.protolog.IProtoLogConfigurationService",
        "name": "protolog_configuration"
      },
      {
        "id": "240",
        "interface": "android.ranging.IRangingAdapter",
        "name": "ranging"
      },
      {
        "id": "241",
        "interface": "android.scheduling.IRebootReadinessManager",
        "name": "reboot_readiness"
      },
      {
        "id": "242",
        "interface": "android.os.IRecoverySystem",
        "name": "recovery"
      },
      {
        "id": "243",
        "interface": "android.security.rkp.IRemoteProvisioning",
        "name": "remote_provisioning"
      },
      {
        "id": "244",
        "interface": "android.content.res.IResourcesManager",
        "name": "resources"
      },
      {
        "id": "245",
        "interface": "android.content.IRestrictionsManager",
        "name": "restrictions"
      },
      {
        "id": "246",
        "interface": "android.app.role.IRoleManager",
        "name": "role"
      },
      {
        "id": "247",
        "interface": "android.content.rollback.IRollbackManager",
        "name": "rollback"
      },
      {
        "id": "248",
        "interface": "",
        "name": "runtime"
      },
      {
        "id": "249",
        "interface": "android.safetycenter.ISafetyCenterManager",
        "name": "safety_center"
      },
      {
        "id": "250",
        "interface": "android.os.ISchedulingPolicyService",
        "name": "scheduling_policy"
      },
      {
        "id": "251",
        "interface": "android.app.sdksandbox.ISdkSandboxManager",
        "name": "sdk_sandbox"
      },
      {
        "id": "252",
        "interface": "android.app.ISearchManager",
        "name": "search"
      },
      {
        "id": "253",
        "interface": "android.app.search.ISearchUiManager",
        "name": "search_ui"
      },
      {
        "id": "254",
        "interface": "android.security.keystore.IKeyAttestationApplicationIdProvider",
        "name": "sec_key_att_app_id_provider"
      },
      {
        "id": "255",
        "interface": "android.se.omapi.ISecureElementService",
        "name": "secure_element"
      },
      {
        "id": "256",
        "interface": "android.os.ISecurityStateManager",
        "name": "security_state"
      },
      {
        "id": "257",
        "interface": "android.view.selectiontoolbar.ISelectionToolbarManager",
        "name": "selection_toolbar"
      },
      {
        "id": "258",
        "interface": "android.view.ISensitiveContentProtectionManager",
        "name": "sensitive_content_protection_service"
      },
      {
        "id": "259",
        "interface": "android.hardware.ISensorPrivacyManager",
        "name": "sensor_privacy"
      },
      {
        "id": "260",
        "interface": "android.gui.SensorServer",
        "name": "sensorservice"
      },
      {
        "id": "261",
        "interface": "android.hardware.serial.ISerialManager",
        "name": "serial"
      },
      {
        "id": "262",
        "interface": "android.net.nsd.INsdManager",
        "name": "servicediscovery"
      },
      {
        "id": "263",
        "interface": "",
        "name": "settings"
      },
      {
        "id": "264",
        "interface": "android.content.pm.IShortcutService",
        "name": "shortcut"
      },
      {
        "id": "265",
        "interface": "com.android.internal.telephony.IIccPhoneBook",
        "name": "simphonebook"
      },
      {
        "id": "266",
        "interface": "android.app.slice.ISliceManager",
        "name": "slice"
      },
      {
        "id": "267",
        "interface": "android.app.smartspace.ISmartspaceManager",
        "name": "smartspace"
      },
      {
        "id": "268",
        "interface": "com.android.internal.app.ISoundTriggerService",
        "name": "soundtrigger"
      },
      {
        "id": "269",
        "interface": "android.media.soundtrigger_middleware.ISoundTriggerMiddlewareService",
        "name": "soundtrigger_middleware"
      },
      {
        "id": "270",
        "interface": "android.speech.IRecognitionServiceManager",
        "name": "speech_recognition"
      },
      {
        "id": "271",
        "interface": "android.os.IStatsd",
        "name": "stats"
      },
      {
        "id": "272",
        "interface": "android.os.IStatsBootstrapAtomService",
        "name": "statsbootstrap"
      },
      {
        "id": "273",
        "interface": "android.os.IStatsCompanionService",
        "name": "statscompanion"
      },
      {
        "id": "274",
        "interface": "android.os.IStatsManagerService",
        "name": "statsmanager"
      },
      {
        "id": "275",
        "interface": "com.android.internal.statusbar.IStatusBarService",
        "name": "statusbar"
      },
      {
        "id": "276",
        "interface": "android.os.IStoraged",
        "name": "storaged"
      },
      {
        "id": "277",
        "interface": "android.os.storaged.IStoragedPrivate",
        "name": "storaged_pri"
      },
      {
        "id": "278",
        "interface": "android.app.usage.IStorageStatsManager",
        "name": "storagestats"
      },
      {
        "id": "279",
        "interface": "android.app.supervision.ISupervisionManager",
        "name": "supervision"
      },
      {
        "id": "280",
        "interface": "",
        "name": "suspend_control"
      },
      {
        "id": "281",
        "interface": "",
        "name": "suspend_control_internal"
      },
      {
        "id": "282",
        "interface": "android.os.ISystemConfig",
        "name": "system_config"
      },
      {
        "id": "283",
        "interface": "",
        "name": "system_server_dumper"
      },
      {
        "id": "284",
        "interface": "android.os.ISystemUpdateManager",
        "name": "system_update"
      },
      {
        "id": "285",
        "interface": "android.companion.datatransfer.continuity.ITaskContinuityManager",
        "name": "task_continuity"
      },
      {
        "id": "286",
        "interface": "com.android.internal.telecom.ITelecomService",
        "name": "telecom"
      },
      {
        "id": "287",
        "interface": "com.android.internal.telephony.ITelephonyRegistry",
        "name": "telephony.registry"
      },
      {
        "id": "288",
        "interface": "android.telephony.ims.aidl.IImsRcsController",
        "name": "telephony_ims"
      },
      {
        "id": "289",
        "interface": "com.android.internal.telephony.IPhoneNumber",
        "name": "telephony_phone_number"
      },
      {
        "id": "290",
        "interface": "",
        "name": "testharness"
      },
      {
        "id": "291",
        "interface": "android.net.ITetheringConnector",
        "name": "tethering"
      },
      {
        "id": "292",
        "interface": "android.service.textclassifier.ITextClassifierService",
        "name": "textclassification"
      },
      {
        "id": "293",
        "interface": "com.android.internal.textservice.ITextServicesManager",
        "name": "textservices"
      },
      {
        "id": "294",
        "interface": "android.speech.tts.ITextToSpeechManager",
        "name": "texttospeech"
      },
      {
        "id": "295",
        "interface": "android.os.IThermalService",
        "name": "thermalservice"
      },
      {
        "id": "296",
        "interface": "android.net.connectivity.android.net.thread.IThreadNetworkManager",
        "name": "thread_network"
      },
      {
        "id": "297",
        "interface": "android.app.timedetector.ITimeDetectorService",
        "name": "time_detector"
      },
      {
        "id": "298",
        "interface": "android.app.timezonedetector.ITimeZoneDetectorService",
        "name": "time_zone_detector"
      },
      {
        "id": "299",
        "interface": "android.tracing.ITracingServiceProxy",
        "name": "tracing.proxy"
      },
      {
        "id": "300",
        "interface": "android.os.ITradeInMode",
        "name": "tradeinmode"
      },
      {
        "id": "301",
        "interface": "android.view.translation.ITranslationManager",
        "name": "translation"
      },
      {
        "id": "302",
        "interface": "com.android.internal.os.IBinaryTransparencyService",
        "name": "transparency"
      },
      {
        "id": "303",
        "interface": "android.app.trust.ITrustManager",
        "name": "trust"
      },
      {
        "id": "304",
        "interface": "android.app.IUiModeManager",
        "name": "uimode"
      },
      {
        "id": "305",
        "interface": "android.os.IUpdateLock",
        "name": "updatelock"
      },
      {
        "id": "306",
        "interface": "com.android.uprobestats.IUprobeStatsBridgeService",
        "name": "uprobestats_bridge"
      },
      {
        "id": "307",
        "interface": "android.app.IUriGrantsManager",
        "name": "uri_grants"
      },
      {
        "id": "308",
        "interface": "android.app.usage.IUsageStatsManager",
        "name": "usagestats"
      },
      {
        "id": "309",
        "interface": "android.hardware.usb.IUsbManager",
        "name": "usb"
      },
      {
        "id": "310",
        "interface": "android.os.IUserManager",
        "name": "user"
      },
      {
        "id": "311",
        "interface": "android.uwb.IUwbAdapter",
        "name": "uwb"
      },
      {
        "id": "312",
        "interface": "android.net.vcn.IVcnManagementService",
        "name": "vcn_management"
      },
      {
        "id": "313",
        "interface": "android.os.IVibratorManagerService",
        "name": "vibrator_manager"
      },
      {
        "id": "314",
        "interface": "android.companion.virtual.IVirtualDeviceManager",
        "name": "virtualdevice"
      },
      {
        "id": "315",
        "interface": "android.companion.virtualnative.IVirtualDeviceManagerNative",
        "name": "virtualdevice_native"
      },
      {
        "id": "316",
        "interface": "com.android.internal.app.IVoiceInteractionManagerService",
        "name": "voiceinteraction"
      },
      {
        "id": "317",
        "interface": "",
        "name": "vold"
      },
      {
        "id": "318",
        "interface": "android.net.IVpnManager",
        "name": "vpn_management"
      },
      {
        "id": "319",
        "interface": "android.app.IWallpaperManager",
        "name": "wallpaper"
      },
      {
        "id": "320",
        "interface": "android.app.wallpapereffectsgeneration.IWallpaperEffectsGenerationManager",
        "name": "wallpaper_effects_generation"
      },
      {
        "id": "321",
        "interface": "android.app.wearable.IWearableSensingManager",
        "name": "wearable_sensing"
      },
      {
        "id": "322",
        "interface": "android.webkit.IWebViewUpdateService",
        "name": "webviewupdate"
      },
      {
        "id": "323",
        "interface": "android.net.wifi.IWifiManager",
        "name": "wifi"
      },
      {
        "id": "324",
        "interface": "android.net.wifi.p2p.IWifiP2pManager",
        "name": "wifip2p"
      },
      {
        "id": "325",
        "interface": "android.net.wifi.IWifiScanner",
        "name": "wifiscanner"
      },
      {
        "id": "326",
        "interface": "android.view.IWindowManager",
        "name": "window"
      }
    ]
  }
}
```

---

## adb

**Command:**
```bash
adb shell service list
```

**Output:**
```
Found 327 services:
0	DockObserver: []
1	SurfaceFlinger: [android.ui.ISurfaceComposer]
2	SurfaceFlingerAIDL: [android.gui.ISurfaceComposer]
3	accessibility: [android.view.accessibility.IAccessibilityManager]
4	account: [android.accounts.IAccountManager]
5	activity: [android.app.IActivityManager]
6	activity_structured: [android.app.IActivityManagerStructured]
7	activity_task: [android.app.IActivityTaskManager]
8	adb: [android.debug.IAdbManager]
9	adservices_manager: [android.app.adservices.IAdServicesManager]
10	advanced_protection: [android.security.advancedprotection.IAdvancedProtectionService]
11	alarm: [android.app.IAlarmManager]
12	allowlist: [android.os.allowlist.IAllowlistService]
13	ambient_context: [android.app.ambientcontext.IAmbientContextManager]
14	android.frameworks.cameraservice.service.ICameraService/default: [android.frameworks.cameraservice.service.ICameraService]
15	android.frameworks.devicestate.IDeviceStateService/default: [android.frameworks.devicestate.IDeviceStateService]
16	android.frameworks.location.altitude.IAltitudeService/default: [android.frameworks.location.altitude.IAltitudeService]
17	android.frameworks.sensorservice.ISensorManager/default: [android.frameworks.sensorservice.ISensorManager]
18	android.frameworks.stats.IStats/default: [android.frameworks.stats.IStats]
19	android.frameworks.vibrator.IVibratorControlService/default: [android.frameworks.vibrator.IVibratorControlService]
20	android.hardware.authsecret.IAuthSecret/default: []
21	android.hardware.biometrics.fingerprint.IFingerprint/default: []
22	android.hardware.bluetooth.IBluetoothHci/default: []
23	android.hardware.bluetooth.audio.IBluetoothAudioProviderFactory/default: []
24	android.hardware.bluetooth.offload.leaudio.IHciProxy/default: []
25	android.hardware.camera.provider.ICameraProvider/internal/0: []
26	android.hardware.camera.provider.ICameraProvider/internal/1: []
27	android.hardware.cas.IMediaCasService/default: [android.hardware.cas.IMediaCasService]
28	android.hardware.contexthub.IContextHub/default: []
29	android.hardware.drm.IDrmFactory/widevine: [android.hardware.drm.IDrmFactory]
30	android.hardware.gatekeeper.IGatekeeper/default: []
31	android.hardware.gnss.IGnss/default: []
32	android.hardware.graphics.allocator.IAllocator/default: [android.hardware.graphics.allocator.IAllocator]
33	android.hardware.graphics.composer3.IComposer/default: []
34	android.hardware.health.IHealth/default: []
35	android.hardware.identity.IIdentityCredentialStore/default: []
36	android.hardware.light.ILights/default: []
37	android.hardware.media.c2.IComponentStore/default: [android.hardware.media.c2.IComponentStore]
38	android.hardware.media.c2.IComponentStore/software: [android.hardware.media.c2.IComponentStore]
39	android.hardware.neuralnetworks.IDevice/nnapi-sample_all: [android.hardware.neuralnetworks.IDevice]
40	android.hardware.neuralnetworks.IDevice/nnapi-sample_quant: [android.hardware.neuralnetworks.IDevice]
41	android.hardware.neuralnetworks.IDevice/nnapi-sample_sl_shim: [android.hardware.neuralnetworks.IDevice]
42	android.hardware.power.IPower/default: []
43	android.hardware.power.stats.IPowerStats/default: []
44	android.hardware.radio.config.IRadioConfig/default: []
45	android.hardware.radio.data.IRadioData/slot1: []
46	android.hardware.radio.ims.IRadioIms/slot1: []
47	android.hardware.radio.ims.media.IImsMedia/default: []
48	android.hardware.radio.messaging.IRadioMessaging/slot1: []
49	android.hardware.radio.modem.IRadioModem/slot1: []
50	android.hardware.radio.network.IRadioNetwork/slot1: []
51	android.hardware.radio.sap.ISap/slot1: []
52	android.hardware.radio.sim.IRadioSim/slot1: []
53	android.hardware.radio.voice.IRadioVoice/slot1: []
54	android.hardware.rebootescrow.IRebootEscrow/default: []
55	android.hardware.security.keymint.IKeyMintDevice/default: []
56	android.hardware.security.keymint.IRemotelyProvisionedComponent/default: [android.hardware.security.keymint.IRemotelyProvisionedComponent]
57	android.hardware.security.secureclock.ISecureClock/default: []
58	android.hardware.security.sharedsecret.ISharedSecret/default: []
59	android.hardware.security.sharedsecret.ISharedSecret/gatekeeper: []
60	android.hardware.sensors.ISensors/default: []
61	android.hardware.thermal.IThermal/default: []
62	android.hardware.threadnetwork.IThreadChip/chip0: []
63	android.hardware.usb.IUsb/default: []
64	android.hardware.uwb.IUwb/default: []
65	android.hardware.vibrator.IVibrator/default: []
66	android.hardware.vibrator.IVibratorManager/default: []
67	android.hardware.wifi.IWifi/default: []
68	android.hardware.wifi.supplicant.ISupplicant/default: []
69	android.security.apc: [android.security.apc.IProtectedConfirmation]
70	android.security.authorization: [android.security.authorization.IKeystoreAuthorization]
71	android.security.compat: [android.security.compat.IKeystoreCompatService]
72	android.security.identity: [android.security.identity.ICredentialStoreFactory]
73	android.security.legacykeystore: [android.security.legacykeystore.ILegacyKeystore]
74	android.security.maintenance: [android.security.maintenance.IKeystoreMaintenance]
75	android.security.metrics: [android.security.metrics.IKeystoreMetrics]
76	android.service.gatekeeper.IGateKeeperService: []
77	android.system.keystore2.IKeystoreService/default: [android.system.keystore2.IKeystoreService]
78	android.system.net.netd.INetd/default: []
79	android.system.suspend.ISystemSuspend/default: []
80	android.system.vold.IVold/default: []
81	anomaly_detector: [android.os.profiling.anomaly.IAnomalyDetectorService]
82	app_binding: []
83	app_function: [android.app.appfunctions.IAppFunctionManager]
84	app_hibernation: [android.apphibernation.IAppHibernationService]
85	app_integrity: [android.content.integrity.IAppIntegrityManager]
86	app_prediction: [android.app.prediction.IPredictionManager]
87	app_search: [android.app.appsearch.aidl.IAppSearchManager]
88	appops: [com.android.internal.app.IAppOpsService]
89	appwidget: [com.android.internal.appwidget.IAppWidgetService]
90	artd: []
91	attestation_verification: [android.security.attestationverification.IAttestationVerificationManagerService]
92	audio: [android.media.IAudioService]
93	auth: [android.hardware.biometrics.IAuthService]
94	authentication_policy: [android.security.authenticationpolicy.IAuthenticationPolicyService]
95	autofill: [android.view.autofill.IAutoFillManager]
96	background_install_control: [android.content.pm.IBackgroundInstallControlService]
97	backup: [android.app.backup.IBackupManager]
98	battery: []
99	batteryproperties: [android.os.IBatteryPropertiesRegistrar]
100	batterystats: [com.android.internal.app.IBatteryStats]
101	binder_calls_stats: []
102	binder_stats_consumer: [android.os.binder.IBinderStatsConsumerService]
103	biometric: [android.hardware.biometrics.IBiometricService]
104	blob_store: [android.app.blob.IBlobStoreManager]
105	bluetooth_manager: [android.bluetooth.IBluetoothManager]
106	bugreport: [android.os.IDumpstate]
107	cacheinfo: []
108	carrier_config: [com.android.internal.telephony.ICarrierConfigLoader]
109	clipboard: [android.content.IClipboard]
110	color_display: [android.hardware.display.IColorDisplayManager]
111	companiondevice: [android.companion.ICompanionDeviceManager]
112	connectivity: [android.net.IConnectivityManager]
113	connectivity_native: [android.net.connectivity.aidl.ConnectivityNative]
114	connmetrics: [android.net.IIpConnectivityMetrics]
115	content: [android.content.IContentService]
116	content_capture: [android.view.contentcapture.IContentCaptureManager]
117	content_safety: [android.app.contentsafety.IContentSafetyManager]
118	content_suggestions: [android.app.contentsuggestions.IContentSuggestionsManager]
119	contexthub: [android.hardware.location.IContextHubService]
120	contextual_mode: [android.app.modes.IContextualModeManager]
121	country_detector: [android.location.ICountryDetector]
122	cpuinfo: []
123	credential: [android.credentials.ICredentialManager]
124	crossprofileapps: [android.content.pm.ICrossProfileApps]
125	dataloader_manager: [android.content.pm.IDataLoaderManager]
126	dbinfo: []
127	device_config: []
128	device_identifiers: [android.os.IDeviceIdentifiersPolicyService]
129	device_lock: [android.devicelock.IDeviceLockService]
130	device_policy: [android.app.admin.IDevicePolicyManager]
131	device_state: [android.hardware.devicestate.IDeviceStateManager]
132	deviceidle: [android.os.IDeviceIdleController]
133	devicestoragemonitor: []
134	diskstats: []
135	display: [android.hardware.display.IDisplayManager]
136	dnsresolver: []
137	domain_verification: [android.content.pm.verify.domain.IDomainVerificationManager]
138	dreams: [android.service.dreams.IDreamManager]
139	drm.drmManager: [drm.IDrmManagerService]
140	dropbox: [com.android.internal.os.IDropBoxManagerService]
141	dynamic_instrumentation: [android.os.instrumentation.IDynamicInstrumentationManager]
142	dynamic_system: [android.os.image.IDynamicSystemService]
143	ecm_enhanced_confirmation: [android.app.ecm.IEnhancedConfirmationManager]
144	emergency_affordance: []
145	external_vibrator_service: [android.os.IExternalVibratorService]
146	feature_flags: [android.flags.IFeatureFlags]
147	file: [android.os.storage.IFileService]
148	file_integrity: [android.security.IFileIntegrityService]
149	fingerprint: [android.hardware.fingerprint.IFingerprintService]
150	font: [com.android.internal.graphics.fonts.IFontManager]
151	game: [android.app.IGameManagerService]
152	gfxinfo: []
153	gpu: [android.graphicsenv.IGpuService]
154	grammatical_inflection: [android.app.IGrammaticalInflectionManager]
155	graphicsstats: [android.view.IGraphicsStats]
156	hardware_properties: [android.os.IHardwarePropertiesManager]
157	healthconnect: [android.health.connect.aidl.IHealthConnectService]
158	imms: [com.android.internal.telephony.IMms]
159	incident: []
160	incidentcompanion: [android.os.IIncidentCompanion]
161	incremental: [android.os.incremental.IIncrementalService]
162	input: [android.hardware.input.IInputManager]
163	input_method: [com.android.internal.view.IInputMethodManager]
164	inputflinger: [android.os.IInputFlinger]
165	installd: []
166	intrusion_detection: [android.security.intrusiondetection.IIntrusionDetectionService]
167	ions: [com.android.internal.telephony.IOns]
168	iphonesubinfo: [com.android.internal.telephony.IPhoneSubInfo]
169	ipsec: [android.net.IIpSecService]
170	isms: [com.android.internal.telephony.ISms]
171	isub: [com.android.internal.telephony.ISub]
172	jobscheduler: [android.app.job.IJobScheduler]
173	launcherapps: [android.content.pm.ILauncherApps]
174	legacy_permission: [android.permission.ILegacyPermissionManager]
175	lights: [android.hardware.lights.ILightsManager]
176	locale: [android.app.ILocaleManager]
177	location: [android.location.ILocationManager]
178	location_time_zone_manager: []
179	lock_settings: [com.android.internal.widget.ILockSettings]
180	logcat: [android.os.logcat.ILogcatManagerService]
181	looper_stats: []
182	manager: [android.os.IServiceManager]
183	mdns: []
184	media.audio_flinger: [android.media.IAudioFlingerService]
185	media.audio_policy: [android.media.IAudioPolicyService]
186	media.camera: [android.hardware.ICameraService]
187	media.camera.proxy: [android.hardware.ICameraServiceProxy]
188	media.extractor: [android.IMediaExtractorService]
189	media.metrics: [android.media.IMediaMetricsService]
190	media.player: [android.media.IMediaPlayerService]
191	media.resource_manager: [android.media.IResourceManagerService]
192	media.resource_observer: [android.media.IResourceObserverService]
193	media_communication: [android.media.IMediaCommunicationService]
194	media_metrics: [android.media.metrics.IMediaMetricsManager]
195	media_projection: [android.media.projection.IMediaProjectionManager]
196	media_resource_monitor: [android.media.IMediaResourceMonitor]
197	media_router: [android.media.IMediaRouterService]
198	media_session: [android.media.session.ISessionManager]
199	meminfo: []
200	memtrack.proxy: [android.hardware.memtrack.IMemtrack]
201	midi: [android.media.midi.IMidiManager]
202	mount: [android.os.storage.IStorageManager]
203	music_recognition: [android.media.musicrecognition.IMusicRecognitionManager]
204	nearby: [android.nearby.INearbyManager]
205	netd: []
206	netd_listener: [android.net.metrics.INetdEventListener]
207	netpolicy: [android.net.INetworkPolicyManager]
208	netstats: [android.net.INetworkStatsService]
209	network_management: [android.os.INetworkManagementService]
210	network_score: [android.net.INetworkScoreService]
211	network_stack: [android.net.INetworkStackConnector]
212	network_time_update_service: []
213	network_watchlist: [com.android.internal.net.INetworkWatchlistManager]
214	notification: [android.app.INotificationManager]
215	on_device_intelligence: [android.app.ondeviceintelligence.IOnDeviceIntelligenceManager]
216	ondevicepersonalization_system_service: [android.ondevicepersonalization.IOnDevicePersonalizationSystemService]
217	ot_daemon: []
218	overlay: [android.content.om.IOverlayManager]
219	pac_proxy: [android.net.IPacProxyManager]
220	package: [android.content.pm.IPackageManager]
221	package_native: [android.content.pm.IPackageManagerNative]
222	pcc_sandbox: [android.app.privatecompute.IPccSandboxManager]
223	pcc_sandbox_native: [android.app.privatecompute.IPccSandboxManagerNative]
224	people: [android.app.people.IPeopleManager]
225	performance_hint: [android.os.IHintManager]
226	permission: [android.os.IPermissionController]
227	permission_checker: [android.permission.IPermissionChecker]
228	permissionmgr: [android.permission.IPermissionManager]
229	phone: [com.android.internal.telephony.ITelephony]
230	pinner: [android.app.pinner.IPinnerService]
231	platform_compat: [com.android.internal.compat.IPlatformCompat]
232	platform_compat_native: [com.android.internal.compat.IPlatformCompatNative]
233	power: [android.os.IPowerManager]
234	powerstats: [android.os.IPowerStatsService]
235	print: [android.print.IPrintManager]
236	processinfo: [android.os.IProcessInfoService]
237	procstats: [com.android.internal.app.procstats.IProcessStats]
238	profiling_service: [android.os.IProfilingService]
239	protolog_configuration: [com.android.internal.protolog.IProtoLogConfigurationService]
240	ranging: [android.ranging.IRangingAdapter]
241	reboot_readiness: [android.scheduling.IRebootReadinessManager]
242	recovery: [android.os.IRecoverySystem]
243	remote_provisioning: [android.security.rkp.IRemoteProvisioning]
244	resources: [android.content.res.IResourcesManager]
245	restrictions: [android.content.IRestrictionsManager]
246	role: [android.app.role.IRoleManager]
247	rollback: [android.content.rollback.IRollbackManager]
248	runtime: []
249	safety_center: [android.safetycenter.ISafetyCenterManager]
250	scheduling_policy: [android.os.ISchedulingPolicyService]
251	sdk_sandbox: [android.app.sdksandbox.ISdkSandboxManager]
252	search: [android.app.ISearchManager]
253	search_ui: [android.app.search.ISearchUiManager]
254	sec_key_att_app_id_provider: [android.security.keystore.IKeyAttestationApplicationIdProvider]
255	secure_element: [android.se.omapi.ISecureElementService]
256	security_state: [android.os.ISecurityStateManager]
257	selection_toolbar: [android.view.selectiontoolbar.ISelectionToolbarManager]
258	sensitive_content_protection_service: [android.view.ISensitiveContentProtectionManager]
259	sensor_privacy: [android.hardware.ISensorPrivacyManager]
260	sensorservice: [android.gui.SensorServer]
261	serial: [android.hardware.serial.ISerialManager]
262	servicediscovery: [android.net.nsd.INsdManager]
263	settings: []
264	shortcut: [android.content.pm.IShortcutService]
265	simphonebook: [com.android.internal.telephony.IIccPhoneBook]
266	slice: [android.app.slice.ISliceManager]
267	smartspace: [android.app.smartspace.ISmartspaceManager]
268	soundtrigger: [com.android.internal.app.ISoundTriggerService]
269	soundtrigger_middleware: [android.media.soundtrigger_middleware.ISoundTriggerMiddlewareService]
270	speech_recognition: [android.speech.IRecognitionServiceManager]
271	stats: [android.os.IStatsd]
272	statsbootstrap: [android.os.IStatsBootstrapAtomService]
273	statscompanion: [android.os.IStatsCompanionService]
274	statsmanager: [android.os.IStatsManagerService]
275	statusbar: [com.android.internal.statusbar.IStatusBarService]
276	storaged: [android.os.IStoraged]
277	storaged_pri: [android.os.storaged.IStoragedPrivate]
278	storagestats: [android.app.usage.IStorageStatsManager]
279	supervision: [android.app.supervision.ISupervisionManager]
280	suspend_control: []
281	suspend_control_internal: []
282	system_config: [android.os.ISystemConfig]
283	system_server_dumper: []
284	system_update: [android.os.ISystemUpdateManager]
285	task_continuity: [android.companion.datatransfer.continuity.ITaskContinuityManager]
286	telecom: [com.android.internal.telecom.ITelecomService]
287	telephony.registry: [com.android.internal.telephony.ITelephonyRegistry]
288	telephony_ims: [android.telephony.ims.aidl.IImsRcsController]
289	telephony_phone_number: [com.android.internal.telephony.IPhoneNumber]
290	testharness: []
291	tethering: [android.net.ITetheringConnector]
292	textclassification: [android.service.textclassifier.ITextClassifierService]
293	textservices: [com.android.internal.textservice.ITextServicesManager]
294	texttospeech: [android.speech.tts.ITextToSpeechManager]
295	thermalservice: [android.os.IThermalService]
296	thread_network: [android.net.connectivity.android.net.thread.IThreadNetworkManager]
297	time_detector: [android.app.timedetector.ITimeDetectorService]
298	time_zone_detector: [android.app.timezonedetector.ITimeZoneDetectorService]
299	tracing.proxy: [android.tracing.ITracingServiceProxy]
300	tradeinmode: [android.os.ITradeInMode]
301	translation: [android.view.translation.ITranslationManager]
302	transparency: [com.android.internal.os.IBinaryTransparencyService]
303	trust: [android.app.trust.ITrustManager]
304	uimode: [android.app.IUiModeManager]
305	updatelock: [android.os.IUpdateLock]
306	uprobestats_bridge: [com.android.uprobestats.IUprobeStatsBridgeService]
307	uri_grants: [android.app.IUriGrantsManager]
308	usagestats: [android.app.usage.IUsageStatsManager]
309	usb: [android.hardware.usb.IUsbManager]
310	user: [android.os.IUserManager]
311	uwb: [android.uwb.IUwbAdapter]
312	vcn_management: [android.net.vcn.IVcnManagementService]
313	vibrator_manager: [android.os.IVibratorManagerService]
314	virtualdevice: [android.companion.virtual.IVirtualDeviceManager]
315	virtualdevice_native: [android.companion.virtualnative.IVirtualDeviceManagerNative]
316	voiceinteraction: [com.android.internal.app.IVoiceInteractionManagerService]
317	vold: []
318	vpn_management: [android.net.IVpnManager]
319	wallpaper: [android.app.IWallpaperManager]
320	wallpaper_effects_generation: [android.app.wallpapereffectsgeneration.IWallpaperEffectsGenerationManager]
321	wearable_sensing: [android.app.wearable.IWearableSensingManager]
322	webviewupdate: [android.webkit.IWebViewUpdateService]
323	wifi: [android.net.wifi.IWifiManager]
324	wifip2p: [android.net.wifi.p2p.IWifiP2pManager]
325	wifiscanner: [android.net.wifi.IWifiScanner]
326	window: [android.view.IWindowManager]
```
