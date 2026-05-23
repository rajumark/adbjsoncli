# `adbjson shell dumpsys network_management`

## adbjson

**Command:**
```bash
adbjson shell dumpsys network_management
```

**Output:**
```json
{
  "status": 0,
  "output": "Flags:\ncom.android.server.net.use_metered_firewall_chains: true\n\nActive quota ifaces: {}\nActive alert ifaces: {}\nData saver mode: false\nUID bandwith control denied UIDs: []\nUID bandwith control allowed UIDs: []\nUID firewall  rule: []\nUID firewall standby chain enabled: true\nUID firewall standby rule: [10142:2,10161:2,10162:2,10163:2,10164:2,10167:2,10168:2,10174:2,10180:2,10183:2,20142:2,20161:2,20162:2,20163:2,20164:2,20167:2,20168:2,20174:2,20180:2,20183:2]\nUID firewall dozable chain enabled: false\nUID firewall dozable rule: []\nUID firewall powersave chain enabled: false\nUID firewall powersave rule: []\nUID firewall restricted mode chain enabled: false\nUID firewall restricted rule: []\nUID firewall low power standby chain enabled: false\nUID firewall low_power_standby rule: []\nUID firewall background chain enabled: true\nUID firewall background rule: [10104:1,10110:1,10119:1,10145:1,10149:1,10153:1,10154:1,10156:1,10158:1,10167:1,10202:1,10219:1,10225:1,10226:1,20104:1,20110:1,20119:1,20145:1,20149:1,20153:1,20154:1,20156:1,20158:1,20167:1,20202:1,20219:1,20225:1,20226:1]\nUID firewall metered allow chain enabled (Data saver mode): false\nUID firewall metered_allow rule: [10110:1,10145:1,10154:1,10158:1,10202:1,10219:1,10225:1,10226:1,20110:1,20145:1,20154:1,20158:1,20202:1,20219:1,20225:1,20226:1]\nUID firewall metered deny_user chain enabled (always-on): true\nUID firewall metered_deny_user rule: []\nUID firewall metered deny_admin chain enabled (always-on): true\nUID firewall metered_deny_admin rule: []\nFirewall enabled: false\nNetd service status: alive"
}
```

---

## adb

**Command:**
```bash
adb shell dumpsys network_management
```

**Output:**
```
Flags:
com.android.server.net.use_metered_firewall_chains: true

Active quota ifaces: {}
Active alert ifaces: {}
Data saver mode: false
UID bandwith control denied UIDs: []
UID bandwith control allowed UIDs: []
UID firewall  rule: []
UID firewall standby chain enabled: true
UID firewall standby rule: [10142:2,10161:2,10162:2,10163:2,10164:2,10167:2,10168:2,10174:2,10180:2,10183:2,20142:2,20161:2,20162:2,20163:2,20164:2,20167:2,20168:2,20174:2,20180:2,20183:2]
UID firewall dozable chain enabled: false
UID firewall dozable rule: []
UID firewall powersave chain enabled: false
UID firewall powersave rule: []
UID firewall restricted mode chain enabled: false
UID firewall restricted rule: []
UID firewall low power standby chain enabled: false
UID firewall low_power_standby rule: []
UID firewall background chain enabled: true
UID firewall background rule: [10104:1,10110:1,10119:1,10145:1,10149:1,10153:1,10154:1,10156:1,10158:1,10167:1,10202:1,10219:1,10225:1,10226:1,20104:1,20110:1,20119:1,20145:1,20149:1,20153:1,20154:1,20156:1,20158:1,20167:1,20202:1,20219:1,20225:1,20226:1]
UID firewall metered allow chain enabled (Data saver mode): false
UID firewall metered_allow rule: [10110:1,10145:1,10154:1,10158:1,10202:1,10219:1,10225:1,10226:1,20110:1,20145:1,20154:1,20158:1,20202:1,20219:1,20225:1,20226:1]
UID firewall metered deny_user chain enabled (always-on): true
UID firewall metered_deny_user rule: []
UID firewall metered deny_admin chain enabled (always-on): true
UID firewall metered_deny_admin rule: []
Firewall enabled: false
Netd service status: alive
```
