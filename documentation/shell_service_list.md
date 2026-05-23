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
    "service_count": 326,
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
        "interface": "android.security.apc.IProtectedConfirmation",
        "name": "android.security.apc"
      },
      {
        "id": "69",
        "interface": "android.security.authorization.IKeystoreAuthorization",
        "name": "android.security.authorization"
      },
      {
        "id": "70",
        "interface": "android.security.compat.IKeystoreCompatService",
        "name": "android.security.compat"
      },
      {
        "id": "71",
        "interface": "android.security.identity.ICredentialStoreFactory",
        "name": "android.security.identity"
      },
      {
        "id": "72",
        "interface": "android.security.legacykeystore.ILegacyKeystore",
        "name": "android.security.legacykeystore"
      },
      {
        "id": "73",
        "interface": "android.security.maintenance.IKeystoreMaintenance",
        "name": "android.security.maintenance"
      },
      {
        "id": "74",
        "interface": "android.security.metrics.IKeystoreMetrics",
        "name": "android.security.metrics"
      },
      {
        "id": "75",
        "interface": "",
        "name": "android.service.gatekeeper.IGateKeeperService"
      },
      {
        "id": "76",
        "interface": "android.system.keystore2.IKeystoreService",
        "name": "android.system.keystore2.IKeystoreService/default"
      },
      {
        "id": "77",
        "interface": "",
        "name": "android.system.net.netd.INetd/default"
      },
      {
        "id": "78",
        "interface": "",
        "name": "android.system.suspend.ISystemSuspend/default"
      },
      {
        "id": "79",
        "interface": "",
        "name": "android.system.vold.IVold/default"
      },
      {
        "id": "80",
        "interface": "android.os.profiling.anomaly.IAnomalyDetectorService",
        "name": "anomaly_detector"
      },
      {
        "id": "81",
        "interface": "",
        "name": "app_binding"
      },
      {
        "id": "82",
        "interface": "android.app.appfunctions.IAppFunctionManager",
        "name": "app_function"
      },
      {
        "id": "83",
        "interface": "android.apphibernation.IAppHibernationService",
        "name": "app_hibernation"
      },
      {
        "id": "84",
        "interface": "android.content.integrity.IAppIntegrityManager",
        "name": "app_integrity"
      },
      {
        "id": "85",
        "interface": "android.app.prediction.IPredictionManager",
        "name": "app_prediction"
      },
      {
        "id": "86",
        "interface": "android.app.appsearch.aidl.IAppSearchManager",
        "name": "app_search"
      },
      {
        "id": "87",
        "interface": "com.android.internal.app.IAppOpsService",
        "name": "appops"
      },
      {
        "id": "88",
        "interface": "com.android.internal.appwidget.IAppWidgetService",
        "name": "appwidget"
      },
      {
        "id": "89",
        "interface": "",
        "name": "artd"
      },
      {
        "id": "90",
        "interface": "android.security.attestationverification.IAttestationVerificationManagerService",
        "name": "attestation_verification"
      },
      {
        "id": "91",
        "interface": "android.media.IAudioService",
        "name": "audio"
      },
      {
        "id": "92",
        "interface": "android.hardware.biometrics.IAuthService",
        "name": "auth"
      },
      {
        "id": "93",
        "interface": "android.security.authenticationpolicy.IAuthenticationPolicyService",
        "name": "authentication_policy"
      },
      {
        "id": "94",
        "interface": "android.view.autofill.IAutoFillManager",
        "name": "autofill"
      },
      {
        "id": "95",
        "interface": "android.content.pm.IBackgroundInstallControlService",
        "name": "background_install_control"
      },
      {
        "id": "96",
        "interface": "android.app.backup.IBackupManager",
        "name": "backup"
      },
      {
        "id": "97",
        "interface": "",
        "name": "battery"
      },
      {
        "id": "98",
        "interface": "android.os.IBatteryPropertiesRegistrar",
        "name": "batteryproperties"
      },
      {
        "id": "99",
        "interface": "com.android.internal.app.IBatteryStats",
        "name": "batterystats"
      },
      {
        "id": "100",
        "interface": "",
        "name": "binder_calls_stats"
      },
      {
        "id": "101",
        "interface": "android.os.binder.IBinderStatsConsumerService",
        "name": "binder_stats_consumer"
      },
      {
        "id": "102",
        "interface": "android.hardware.biometrics.IBiometricService",
        "name": "biometric"
      },
      {
        "id": "103",
        "interface": "android.app.blob.IBlobStoreManager",
        "name": "blob_store"
      },
      {
        "id": "104",
        "interface": "android.bluetooth.IBluetoothManager",
        "name": "bluetooth_manager"
      },
      {
        "id": "105",
        "interface": "android.os.IDumpstate",
        "name": "bugreport"
      },
      {
        "id": "106",
        "interface": "",
        "name": "cacheinfo"
      },
      {
        "id": "107",
        "interface": "com.android.internal.telephony.ICarrierConfigLoader",
        "name": "carrier_config"
      },
      {
        "id": "108",
        "interface": "android.content.IClipboard",
        "name": "clipboard"
      },
      {
        "id": "109",
        "interface": "android.hardware.display.IColorDisplayManager",
        "name": "color_display"
      },
      {
        "id": "110",
        "interface": "android.companion.ICompanionDeviceManager",
        "name": "companiondevice"
      },
      {
        "id": "111",
        "interface": "android.net.IConnectivityManager",
        "name": "connectivity"
      },
      {
        "id": "112",
        "interface": "android.net.connectivity.aidl.ConnectivityNative",
        "name": "connectivity_native"
      },
      {
        "id": "113",
        "interface": "android.net.IIpConnectivityMetrics",
        "name": "connmetrics"
      },
      {
        "id": "114",
        "interface": "android.content.IContentService",
        "name": "content"
      },
      {
        "id": "115",
        "interface": "android.view.contentcapture.IContentCaptureManager",
        "name": "content_capture"
      },
      {
        "id": "116",
        "interface": "android.app.contentsafety.IContentSafetyManager",
        "name": "content_safety"
      },
      {
        "id": "117",
        "interface": "android.app.contentsuggestions.IContentSuggestionsManager",
        "name": "content_suggestions"
      },
      {
        "id": "118",
        "interface": "android.hardware.location.IContextHubService",
        "name": "contexthub"
      },
      {
        "id": "119",
        "interface": "android.app.modes.IContextualModeManager",
        "name": "contextual_mode"
      },
      {
        "id": "120",
        "interface": "android.location.ICountryDetector",
        "name": "country_detector"
      },
      {
        "id": "121",
        "interface": "",
        "name": "cpuinfo"
      },
      {
        "id": "122",
        "interface": "android.credentials.ICredentialManager",
        "name": "credential"
      },
      {
        "id": "123",
        "interface": "android.content.pm.ICrossProfileApps",
        "name": "crossprofileapps"
      },
      {
        "id": "124",
        "interface": "android.content.pm.IDataLoaderManager",
        "name": "dataloader_manager"
      },
      {
        "id": "125",
        "interface": "",
        "name": "dbinfo"
      },
      {
        "id": "126",
        "interface": "",
        "name": "device_config"
      },
      {
        "id": "127",
        "interface": "android.os.IDeviceIdentifiersPolicyService",
        "name": "device_identifiers"
      },
      {
        "id": "128",
        "interface": "android.devicelock.IDeviceLockService",
        "name": "device_lock"
      },
      {
        "id": "129",
        "interface": "android.app.admin.IDevicePolicyManager",
        "name": "device_policy"
      },
      {
        "id": "130",
        "interface": "android.hardware.devicestate.IDeviceStateManager",
        "name": "device_state"
      },
      {
        "id": "131",
        "interface": "android.os.IDeviceIdleController",
        "name": "deviceidle"
      },
      {
        "id": "132",
        "interface": "",
        "name": "devicestoragemonitor"
      },
      {
        "id": "133",
        "interface": "",
        "name": "diskstats"
      },
      {
        "id": "134",
        "interface": "android.hardware.display.IDisplayManager",
        "name": "display"
      },
      {
        "id": "135",
        "interface": "",
        "name": "dnsresolver"
      },
      {
        "id": "136",
        "interface": "android.content.pm.verify.domain.IDomainVerificationManager",
        "name": "domain_verification"
      },
      {
        "id": "137",
        "interface": "android.service.dreams.IDreamManager",
        "name": "dreams"
      },
      {
        "id": "138",
        "interface": "drm.IDrmManagerService",
        "name": "drm.drmManager"
      },
      {
        "id": "139",
        "interface": "com.android.internal.os.IDropBoxManagerService",
        "name": "dropbox"
      },
      {
        "id": "140",
        "interface": "android.os.instrumentation.IDynamicInstrumentationManager",
        "name": "dynamic_instrumentation"
      },
      {
        "id": "141",
        "interface": "android.os.image.IDynamicSystemService",
        "name": "dynamic_system"
      },
      {
        "id": "142",
        "interface": "android.app.ecm.IEnhancedConfirmationManager",
        "name": "ecm_enhanced_confirmation"
      },
      {
        "id": "143",
        "interface": "",
        "name": "emergency_affordance"
      },
      {
        "id": "144",
        "interface": "android.os.IExternalVibratorService",
        "name": "external_vibrator_service"
      },
      {
        "id": "145",
        "interface": "android.flags.IFeatureFlags",
        "name": "feature_flags"
      },
      {
        "id": "146",
        "interface": "android.os.storage.IFileService",
        "name": "file"
      },
      {
        "id": "147",
        "interface": "android.security.IFileIntegrityService",
        "name": "file_integrity"
      },
      {
        "id": "148",
        "interface": "android.hardware.fingerprint.IFingerprintService",
        "name": "fingerprint"
      },
      {
        "id": "149",
        "interface": "com.android.internal.graphics.fonts.IFontManager",
        "name": "font"
      },
      {
        "id": "150",
        "interface": "android.app.IGameManagerService",
        "name": "game"
      },
      {
        "id": "151",
        "interface": "",
        "name": "gfxinfo"
      },
      {
        "id": "152",
        "interface": "android.graphicsenv.IGpuService",
        "name": "gpu"
      },
      {
        "id": "153",
        "interface": "android.app.IGrammaticalInflectionManager",
        "name": "grammatical_inflection"
      },
      {
        "id": "154",
        "interface": "android.view.IGraphicsStats",
        "name": "graphicsstats"
      },
      {
        "id": "155",
        "interface": "android.os.IHardwarePropertiesManager",
        "name": "hardware_properties"
      },
      {
        "id": "156",
        "interface": "android.health.connect.aidl.IHealthConnectService",
        "name": "healthconnect"
      },
      {
        "id": "157",
        "interface": "com.android.internal.telephony.IMms",
        "name": "imms"
      },
      {
        "id": "158",
        "interface": "",
        "name": "incident"
      },
      {
        "id": "159",
        "interface": "android.os.IIncidentCompanion",
        "name": "incidentcompanion"
      },
      {
        "id": "160",
        "interface": "android.os.incremental.IIncrementalService",
        "name": "incremental"
      },
      {
        "id": "161",
        "interface": "android.hardware.input.IInputManager",
        "name": "input"
      },
      {
        "id": "162",
        "interface": "com.android.internal.view.IInputMethodManager",
        "name": "input_method"
      },
      {
        "id": "163",
        "interface": "android.os.IInputFlinger",
        "name": "inputflinger"
      },
      {
        "id": "164",
        "interface": "",
        "name": "installd"
      },
      {
        "id": "165",
        "interface": "android.security.intrusiondetection.IIntrusionDetectionService",
        "name": "intrusion_detection"
      },
      {
        "id": "166",
        "interface": "com.android.internal.telephony.IOns",
        "name": "ions"
      },
      {
        "id": "167",
        "interface": "com.android.internal.telephony.IPhoneSubInfo",
        "name": "iphonesubinfo"
      },
      {
        "id": "168",
        "interface": "android.net.IIpSecService",
        "name": "ipsec"
      },
      {
        "id": "169",
        "interface": "com.android.internal.telephony.ISms",
        "name": "isms"
      },
      {
        "id": "170",
        "interface": "com.android.internal.telephony.ISub",
        "name": "isub"
      },
      {
        "id": "171",
        "interface": "android.app.job.IJobScheduler",
        "name": "jobscheduler"
      },
      {
        "id": "172",
        "interface": "android.content.pm.ILauncherApps",
        "name": "launcherapps"
      },
      {
        "id": "173",
        "interface": "android.permission.ILegacyPermissionManager",
        "name": "legacy_permission"
      },
      {
        "id": "174",
        "interface": "android.hardware.lights.ILightsManager",
        "name": "lights"
      },
      {
        "id": "175",
        "interface": "android.app.ILocaleManager",
        "name": "locale"
      },
      {
        "id": "176",
        "interface": "android.location.ILocationManager",
        "name": "location"
      },
      {
        "id": "177",
        "interface": "",
        "name": "location_time_zone_manager"
      },
      {
        "id": "178",
        "interface": "com.android.internal.widget.ILockSettings",
        "name": "lock_settings"
      },
      {
        "id": "179",
        "interface": "android.os.logcat.ILogcatManagerService",
        "name": "logcat"
      },
      {
        "id": "180",
        "interface": "",
        "name": "looper_stats"
      },
      {
        "id": "181",
        "interface": "android.os.IServiceManager",
        "name": "manager"
      },
      {
        "id": "182",
        "interface": "",
        "name": "mdns"
      },
      {
        "id": "183",
        "interface": "android.media.IAudioFlingerService",
        "name": "media.audio_flinger"
      },
      {
        "id": "184",
        "interface": "android.media.IAudioPolicyService",
        "name": "media.audio_policy"
      },
      {
        "id": "185",
        "interface": "android.hardware.ICameraService",
        "name": "media.camera"
      },
      {
        "id": "186",
        "interface": "android.hardware.ICameraServiceProxy",
        "name": "media.camera.proxy"
      },
      {
        "id": "187",
        "interface": "android.IMediaExtractorService",
        "name": "media.extractor"
      },
      {
        "id": "188",
        "interface": "android.media.IMediaMetricsService",
        "name": "media.metrics"
      },
      {
        "id": "189",
        "interface": "android.media.IMediaPlayerService",
        "name": "media.player"
      },
      {
        "id": "190",
        "interface": "android.media.IResourceManagerService",
        "name": "media.resource_manager"
      },
      {
        "id": "191",
        "interface": "android.media.IResourceObserverService",
        "name": "media.resource_observer"
      },
      {
        "id": "192",
        "interface": "android.media.IMediaCommunicationService",
        "name": "media_communication"
      },
      {
        "id": "193",
        "interface": "android.media.metrics.IMediaMetricsManager",
        "name": "media_metrics"
      },
      {
        "id": "194",
        "interface": "android.media.projection.IMediaProjectionManager",
        "name": "media_projection"
      },
      {
        "id": "195",
        "interface": "android.media.IMediaResourceMonitor",
        "name": "media_resource_monitor"
      },
      {
        "id": "196",
        "interface": "android.media.IMediaRouterService",
        "name": "media_router"
      },
      {
        "id": "197",
        "interface": "android.media.session.ISessionManager",
        "name": "media_session"
      },
      {
        "id": "198",
        "interface": "",
        "name": "meminfo"
      },
      {
        "id": "199",
        "interface": "android.hardware.memtrack.IMemtrack",
        "name": "memtrack.proxy"
      },
      {
        "id": "200",
        "interface": "android.media.midi.IMidiManager",
        "name": "midi"
      },
      {
        "id": "201",
        "interface": "android.os.storage.IStorageManager",
        "name": "mount"
      },
      {
        "id": "202",
        "interface": "android.media.musicrecognition.IMusicRecognitionManager",
        "name": "music_recognition"
      },
      {
        "id": "203",
        "interface": "android.nearby.INearbyManager",
        "name": "nearby"
      },
      {
        "id": "204",
        "interface": "",
        "name": "netd"
      },
      {
        "id": "205",
        "interface": "android.net.metrics.INetdEventListener",
        "name": "netd_listener"
      },
      {
        "id": "206",
        "interface": "android.net.INetworkPolicyManager",
        "name": "netpolicy"
      },
      {
        "id": "207",
        "interface": "android.net.INetworkStatsService",
        "name": "netstats"
      },
      {
        "id": "208",
        "interface": "android.os.INetworkManagementService",
        "name": "network_management"
      },
      {
        "id": "209",
        "interface": "android.net.INetworkScoreService",
        "name": "network_score"
      },
      {
        "id": "210",
        "interface": "android.net.INetworkStackConnector",
        "name": "network_stack"
      },
      {
        "id": "211",
        "interface": "",
        "name": "network_time_update_service"
      },
      {
        "id": "212",
        "interface": "com.android.internal.net.INetworkWatchlistManager",
        "name": "network_watchlist"
      },
      {
        "id": "213",
        "interface": "android.app.INotificationManager",
        "name": "notification"
      },
      {
        "id": "214",
        "interface": "android.app.ondeviceintelligence.IOnDeviceIntelligenceManager",
        "name": "on_device_intelligence"
      },
      {
        "id": "215",
        "interface": "android.ondevicepersonalization.IOnDevicePersonalizationSystemService",
        "name": "ondevicepersonalization_system_service"
      },
      {
        "id": "216",
        "interface": "",
        "name": "ot_daemon"
      },
      {
        "id": "217",
        "interface": "android.content.om.IOverlayManager",
        "name": "overlay"
      },
      {
        "id": "218",
        "interface": "android.net.IPacProxyManager",
        "name": "pac_proxy"
      },
      {
        "id": "219",
        "interface": "android.content.pm.IPackageManager",
        "name": "package"
      },
      {
        "id": "220",
        "interface": "android.content.pm.IPackageManagerNative",
        "name": "package_native"
      },
      {
        "id": "221",
        "interface": "android.app.privatecompute.IPccSandboxManager",
        "name": "pcc_sandbox"
      },
      {
        "id": "222",
        "interface": "android.app.privatecompute.IPccSandboxManagerNative",
        "name": "pcc_sandbox_native"
      },
      {
        "id": "223",
        "interface": "android.app.people.IPeopleManager",
        "name": "people"
      },
      {
        "id": "224",
        "interface": "android.os.IHintManager",
        "name": "performance_hint"
      },
      {
        "id": "225",
        "interface": "android.os.IPermissionController",
        "name": "permission"
      },
      {
        "id": "226",
        "interface": "android.permission.IPermissionChecker",
        "name": "permission_checker"
      },
      {
        "id": "227",
        "interface": "android.permission.IPermissionManager",
        "name": "permissionmgr"
      },
      {
        "id": "228",
        "interface": "com.android.internal.telephony.ITelephony",
        "name": "phone"
      },
      {
        "id": "229",
        "interface": "android.app.pinner.IPinnerService",
        "name": "pinner"
      },
      {
        "id": "230",
        "interface": "com.android.internal.compat.IPlatformCompat",
        "name": "platform_compat"
      },
      {
        "id": "231",
        "interface": "com.android.internal.compat.IPlatformCompatNative",
        "name": "platform_compat_native"
      },
      {
        "id": "232",
        "interface": "android.os.IPowerManager",
        "name": "power"
      },
      {
        "id": "233",
        "interface": "android.os.IPowerStatsService",
        "name": "powerstats"
      },
      {
        "id": "234",
        "interface": "android.print.IPrintManager",
        "name": "print"
      },
      {
        "id": "235",
        "interface": "android.os.IProcessInfoService",
        "name": "processinfo"
      },
      {
        "id": "236",
        "interface": "com.android.internal.app.procstats.IProcessStats",
        "name": "procstats"
      },
      {
        "id": "237",
        "interface": "android.os.IProfilingService",
        "name": "profiling_service"
      },
      {
        "id": "238",
        "interface": "com.android.internal.protolog.IProtoLogConfigurationService",
        "name": "protolog_configuration"
      },
      {
        "id": "239",
        "interface": "android.ranging.IRangingAdapter",
        "name": "ranging"
      },
      {
        "id": "240",
        "interface": "android.scheduling.IRebootReadinessManager",
        "name": "reboot_readiness"
      },
      {
        "id": "241",
        "interface": "android.os.IRecoverySystem",
        "name": "recovery"
      },
      {
        "id": "242",
        "interface": "android.security.rkp.IRemoteProvisioning",
        "name": "remote_provisioning"
      },
      {
        "id": "243",
        "interface": "android.content.res.IResourcesManager",
        "name": "resources"
      },
      {
        "id": "244",
        "interface": "android.content.IRestrictionsManager",
        "name": "restrictions"
      },
      {
        "id": "245",
        "interface": "android.app.role.IRoleManager",
        "name": "role"
      },
      {
        "id": "246",
        "interface": "android.content.rollback.IRollbackManager",
        "name": "rollback"
      },
      {
        "id": "247",
        "interface": "",
        "name": "runtime"
      },
      {
        "id": "248",
        "interface": "android.safetycenter.ISafetyCenterManager",
        "name": "safety_center"
      },
      {
        "id": "249",
        "interface": "android.os.ISchedulingPolicyService",
        "name": "scheduling_policy"
      },
      {
        "id": "250",
        "interface": "android.app.sdksandbox.ISdkSandboxManager",
        "name": "sdk_sandbox"
      },
      {
        "id": "251",
        "interface": "android.app.ISearchManager",
        "name": "search"
      },
      {
        "id": "252",
        "interface": "android.app.search.ISearchUiManager",
        "name": "search_ui"
      },
      {
        "id": "253",
        "interface": "android.security.keystore.IKeyAttestationApplicationIdProvider",
        "name": "sec_key_att_app_id_provider"
      },
      {
        "id": "254",
        "interface": "android.se.omapi.ISecureElementService",
        "name": "secure_element"
      },
      {
        "id": "255",
        "interface": "android.os.ISecurityStateManager",
        "name": "security_state"
      },
      {
        "id": "256",
        "interface": "android.view.selectiontoolbar.ISelectionToolbarManager",
        "name": "selection_toolbar"
      },
      {
        "id": "257",
        "interface": "android.view.ISensitiveContentProtectionManager",
        "name": "sensitive_content_protection_service"
      },
      {
        "id": "258",
        "interface": "android.hardware.ISensorPrivacyManager",
        "name": "sensor_privacy"
      },
      {
        "id": "259",
        "interface": "android.gui.SensorServer",
        "name": "sensorservice"
      },
      {
        "id": "260",
        "interface": "android.hardware.serial.ISerialManager",
        "name": "serial"
      },
      {
        "id": "261",
        "interface": "android.net.nsd.INsdManager",
        "name": "servicediscovery"
      },
      {
        "id": "262",
        "interface": "",
        "name": "settings"
      },
      {
        "id": "263",
        "interface": "android.content.pm.IShortcutService",
        "name": "shortcut"
      },
      {
        "id": "264",
        "interface": "com.android.internal.telephony.IIccPhoneBook",
        "name": "simphonebook"
      },
      {
        "id": "265",
        "interface": "android.app.slice.ISliceManager",
        "name": "slice"
      },
      {
        "id": "266",
        "interface": "android.app.smartspace.ISmartspaceManager",
        "name": "smartspace"
      },
      {
        "id": "267",
        "interface": "com.android.internal.app.ISoundTriggerService",
        "name": "soundtrigger"
      },
      {
        "id": "268",
        "interface": "android.media.soundtrigger_middleware.ISoundTriggerMiddlewareService",
        "name": "soundtrigger_middleware"
      },
      {
        "id": "269",
        "interface": "android.speech.IRecognitionServiceManager",
        "name": "speech_recognition"
      },
      {
        "id": "270",
        "interface": "android.os.IStatsd",
        "name": "stats"
      },
      {
        "id": "271",
        "interface": "android.os.IStatsBootstrapAtomService",
        "name": "statsbootstrap"
      },
      {
        "id": "272",
        "interface": "android.os.IStatsCompanionService",
        "name": "statscompanion"
      },
      {
        "id": "273",
        "interface": "android.os.IStatsManagerService",
        "name": "statsmanager"
      },
      {
        "id": "274",
        "interface": "com.android.internal.statusbar.IStatusBarService",
        "name": "statusbar"
      },
      {
        "id": "275",
        "interface": "android.os.IStoraged",
        "name": "storaged"
      },
      {
        "id": "276",
        "interface": "android.os.storaged.IStoragedPrivate",
        "name": "storaged_pri"
      },
      {
        "id": "277",
        "interface": "android.app.usage.IStorageStatsManager",
        "name": "storagestats"
      },
      {
        "id": "278",
        "interface": "android.app.supervision.ISupervisionManager",
        "name": "supervision"
      },
      {
        "id": "279",
        "interface": "",
        "name": "suspend_control"
      },
      {
        "id": "280",
        "interface": "",
        "name": "suspend_control_internal"
      },
      {
        "id": "281",
        "interface": "android.os.ISystemConfig",
        "name": "system_config"
      },
      {
        "id": "282",
        "interface": "",
        "name": "system_server_dumper"
      },
      {
        "id": "283",
        "interface": "android.os.ISystemUpdateManager",
        "name": "system_update"
      },
      {
        "id": "284",
        "interface": "android.companion.datatransfer.continuity.ITaskContinuityManager",
        "name": "task_continuity"
      },
      {
        "id": "285",
        "interface": "com.android.internal.telecom.ITelecomService",
        "name": "telecom"
      },
      {
        "id": "286",
        "interface": "com.android.internal.telephony.ITelephonyRegistry",
        "name": "telephony.registry"
      },
      {
        "id": "287",
        "interface": "android.telephony.ims.aidl.IImsRcsController",
        "name": "telephony_ims"
      },
      {
        "id": "288",
        "interface": "com.android.internal.telephony.IPhoneNumber",
        "name": "telephony_phone_number"
      },
      {
        "id": "289",
        "interface": "",
        "name": "testharness"
      },
      {
        "id": "290",
        "interface": "android.net.ITetheringConnector",
        "name": "tethering"
      },
      {
        "id": "291",
        "interface": "android.service.textclassifier.ITextClassifierService",
        "name": "textclassification"
      },
      {
        "id": "292",
        "interface": "com.android.internal.textservice.ITextServicesManager",
        "name": "textservices"
      },
      {
        "id": "293",
        "interface": "android.speech.tts.ITextToSpeechManager",
        "name": "texttospeech"
      },
      {
        "id": "294",
        "interface": "android.os.IThermalService",
        "name": "thermalservice"
      },
      {
        "id": "295",
        "interface": "android.net.connectivity.android.net.thread.IThreadNetworkManager",
        "name": "thread_network"
      },
      {
        "id": "296",
        "interface": "android.app.timedetector.ITimeDetectorService",
        "name": "time_detector"
      },
      {
        "id": "297",
        "interface": "android.app.timezonedetector.ITimeZoneDetectorService",
        "name": "time_zone_detector"
      },
      {
        "id": "298",
        "interface": "android.tracing.ITracingServiceProxy",
        "name": "tracing.proxy"
      },
      {
        "id": "299",
        "interface": "android.os.ITradeInMode",
        "name": "tradeinmode"
      },
      {
        "id": "300",
        "interface": "android.view.translation.ITranslationManager",
        "name": "translation"
      },
      {
        "id": "301",
        "interface": "com.android.internal.os.IBinaryTransparencyService",
        "name": "transparency"
      },
      {
        "id": "302",
        "interface": "android.app.trust.ITrustManager",
        "name": "trust"
      },
      {
        "id": "303",
        "interface": "android.app.IUiModeManager",
        "name": "uimode"
      },
      {
        "id": "304",
        "interface": "android.os.IUpdateLock",
        "name": "updatelock"
      },
      {
        "id": "305",
        "interface": "com.android.uprobestats.IUprobeStatsBridgeService",
        "name": "uprobestats_bridge"
      },
      {
        "id": "306",
        "interface": "android.app.IUriGrantsManager",
        "name": "uri_grants"
      },
      {
        "id": "307",
        "interface": "android.app.usage.IUsageStatsManager",
        "name": "usagestats"
      },
      {
        "id": "308",
        "interface": "android.hardware.usb.IUsbManager",
        "name": "usb"
      },
      {
        "id": "309",
        "interface": "android.os.IUserManager",
        "name": "user"
      },
      {
        "id": "310",
        "interface": "android.uwb.IUwbAdapter",
        "name": "uwb"
      },
      {
        "id": "311",
        "interface": "android.net.vcn.IVcnManagementService",
        "name": "vcn_management"
      },
      {
        "id": "312",
        "interface": "android.os.IVibratorManagerService",
        "name": "vibrator_manager"
      },
      {
        "id": "313",
        "interface": "android.companion.virtual.IVirtualDeviceManager",
        "name": "virtualdevice"
      },
      {
        "id": "314",
        "interface": "android.companion.virtualnative.IVirtualDeviceManagerNative",
        "name": "virtualdevice_native"
      },
      {
        "id": "315",
        "interface": "com.android.internal.app.IVoiceInteractionManagerService",
        "name": "voiceinteraction"
      },
      {
        "id": "316",
        "interface": "",
        "name": "vold"
      },
      {
        "id": "317",
        "interface": "android.net.IVpnManager",
        "name": "vpn_management"
      },
      {
        "id": "318",
        "interface": "android.app.IWallpaperManager",
        "name": "wallpaper"
      },
      {
        "id": "319",
        "interface": "android.app.wallpapereffectsgeneration.IWallpaperEffectsGenerationManager",
        "name": "wallpaper_effects_generation"
      },
      {
        "id": "320",
        "interface": "android.app.wearable.IWearableSensingManager",
        "name": "wearable_sensing"
      },
      {
        "id": "321",
        "interface": "android.webkit.IWebViewUpdateService",
        "name": "webviewupdate"
      },
      {
        "id": "322",
        "interface": "android.net.wifi.IWifiManager",
        "name": "wifi"
      },
      {
        "id": "323",
        "interface": "android.net.wifi.p2p.IWifiP2pManager",
        "name": "wifip2p"
      },
      {
        "id": "324",
        "interface": "android.net.wifi.IWifiScanner",
        "name": "wifiscanner"
      },
      {
        "id": "325",
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
Found 326 services:
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
68	android.security.apc: [android.security.apc.IProtectedConfirmation]
69	android.security.authorization: [android.security.authorization.IKeystoreAuthorization]
70	android.security.compat: [android.security.compat.IKeystoreCompatService]
71	android.security.identity: [android.security.identity.ICredentialStoreFactory]
72	android.security.legacykeystore: [android.security.legacykeystore.ILegacyKeystore]
73	android.security.maintenance: [android.security.maintenance.IKeystoreMaintenance]
74	android.security.metrics: [android.security.metrics.IKeystoreMetrics]
75	android.service.gatekeeper.IGateKeeperService: []
76	android.system.keystore2.IKeystoreService/default: [android.system.keystore2.IKeystoreService]
77	android.system.net.netd.INetd/default: []
78	android.system.suspend.ISystemSuspend/default: []
79	android.system.vold.IVold/default: []
80	anomaly_detector: [android.os.profiling.anomaly.IAnomalyDetectorService]
81	app_binding: []
82	app_function: [android.app.appfunctions.IAppFunctionManager]
83	app_hibernation: [android.apphibernation.IAppHibernationService]
84	app_integrity: [android.content.integrity.IAppIntegrityManager]
85	app_prediction: [android.app.prediction.IPredictionManager]
86	app_search: [android.app.appsearch.aidl.IAppSearchManager]
87	appops: [com.android.internal.app.IAppOpsService]
88	appwidget: [com.android.internal.appwidget.IAppWidgetService]
89	artd: []
90	attestation_verification: [android.security.attestationverification.IAttestationVerificationManagerService]
91	audio: [android.media.IAudioService]
92	auth: [android.hardware.biometrics.IAuthService]
93	authentication_policy: [android.security.authenticationpolicy.IAuthenticationPolicyService]
94	autofill: [android.view.autofill.IAutoFillManager]
95	background_install_control: [android.content.pm.IBackgroundInstallControlService]
96	backup: [android.app.backup.IBackupManager]
97	battery: []
98	batteryproperties: [android.os.IBatteryPropertiesRegistrar]\n\n... (truncated,      327 lines total)
```
