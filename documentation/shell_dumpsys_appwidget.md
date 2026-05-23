# `adbjson shell dumpsys appwidget`

## adbjson

**Command:**
```bash
adbjson shell dumpsys appwidget
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
adb shell dumpsys appwidget
```

**Output:**
```
Providers:
  [0] provider ProviderId{uid:10162, app:10162, cmp:ComponentInfo{com.android.chrome/org.chromium.chrome.browser.quickactionsearchwidget.QuickActionSearchWidgetProvider$QuickActionSearchWidgetProviderSearch}, profile:UserHandle{0}}
    min=(66561x12289)   minResize=(66561x12289) updatePeriodMillis=0 resizeMode=3 widgetCategory=5 autoAdvanceViewId=-1 initialLayout=#7f0e0346 initialKeyguardLayout=#0   zombie=false
  [1] provider ProviderId{uid:10162, app:10162, cmp:ComponentInfo{com.android.chrome/org.chromium.chrome.browser.quickactionsearchwidget.QuickActionSearchWidgetProvider$QuickActionSearchWidgetProviderDino}, profile:UserHandle{0}}
    min=(28161x28161)   minResize=(28161x28161) updatePeriodMillis=0 resizeMode=0 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0e0343 initialKeyguardLayout=#0   zombie=false
  [2] provider ProviderId{uid:10162, app:10162, cmp:ComponentInfo{com.android.chrome/com.google.android.apps.chrome.appwidget.bookmarks.BookmarkThumbnailWidgetProvider}, profile:UserHandle{0}}
    min=(46081x38401)   minResize=(16385x12289) updatePeriodMillis=86400000 resizeMode=3 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0e00a2 initialKeyguardLayout=#0   zombie=false
  [3] provider ProviderId{uid:10162, app:10162, cmp:ComponentInfo{com.android.chrome/org.chromium.chrome.browser.searchwidget.SearchWidgetProvider}, profile:UserHandle{0}}
    min=(61441x12289)   minResize=(61441x12289) updatePeriodMillis=86400000 resizeMode=1 widgetCategory=5 autoAdvanceViewId=-1 initialLayout=#7f0e0382 initialKeyguardLayout=#0   zombie=false
  [4] provider ProviderId{uid:10195, app:10195, cmp:ComponentInfo{com.android.systemui/com.android.systemui.people.widget.PeopleSpaceWidgetProvider}, profile:UserHandle{0}}
    min=(34817x0)   minResize=(15361x12801) updatePeriodMillis=60000 resizeMode=3 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0d0219 initialKeyguardLayout=#0   zombie=false
  [5] provider ProviderId{uid:10163, app:10163, cmp:ComponentInfo{com.google.android.apps.docs/com.google.android.apps.docs.drive.widget.suggestion.SuggestionAppWidgetProvider}, profile:UserHandle{0}}
    min=(66561x28161)   minResize=(36353x25601) updatePeriodMillis=14400000 resizeMode=3 widgetCategory=3 autoAdvanceViewId=-1 initialLayout=#7f0d0044 initialKeyguardLayout=#0   zombie=false
  [6] provider ProviderId{uid:10163, app:10163, cmp:ComponentInfo{com.google.android.apps.docs/com.google.android.apps.docs.drive.widget.CakemixAppWidgetProvider}, profile:UserHandle{0}}
    min=(75777x18433)   minResize=(36353x12289) updatePeriodMillis=79200000 resizeMode=3 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0d0240 initialKeyguardLayout=#0   zombie=false
  [7] provider ProviderId{uid:10168, app:10168, cmp:ComponentInfo{com.google.android.apps.maps/com.google.android.apps.gmm.widget.SearchWidgetProvider}, profile:UserHandle{0}}
    min=(56321x25601)   minResize=(40961x25601) updatePeriodMillis=43200000 resizeMode=3 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0e022c initialKeyguardLayout=#0   zombie=false
  [8] provider ProviderId{uid:10168, app:10168, cmp:ComponentInfo{com.google.android.apps.maps/com.google.android.apps.gmm.widget.traffic.TrafficWidgetProvider}, profile:UserHandle{0}}
    min=(56321x25601)   minResize=(35841x20481) updatePeriodMillis=0 resizeMode=3 widgetCategory=3 autoAdvanceViewId=-1 initialLayout=#7f0e0364 initialKeyguardLayout=#0   zombie=false
  [9] provider ProviderId{uid:10171, app:10171, cmp:ComponentInfo{com.google.android.apps.photos/com.google.android.apps.photos.widget.people.WidgetProviderPeoplePets}, profile:UserHandle{0}}
    min=(0x0)   minResize=(0x0) updatePeriodMillis=0 resizeMode=3 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0e080f initialKeyguardLayout=#0   zombie=false
  [10] provider ProviderId{uid:10171, app:10171, cmp:ComponentInfo{com.google.android.apps.photos/com.google.android.apps.photos.widget.WidgetProvider}, profile:UserHandle{0}}
    min=(15361x25601)   minResize=(15361x25601) updatePeriodMillis=0 resizeMode=3 widgetCategory=3 autoAdvanceViewId=-1 initialLayout=#7f0e080f initialKeyguardLayout=#0   zombie=false
  [11] provider ProviderId{uid:10159, app:10159, cmp:ComponentInfo{com.google.android.apps.wellbeing/com.google.android.apps.wellbeing.widget.screentime.ScreenTimeWidgetProviderReceiver_Receiver}, profile:UserHandle{0}}
    min=(25601x23041)   minResize=(25601x12801) updatePeriodMillis=0 resizeMode=3 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0d00a3 initialKeyguardLayout=#0   zombie=false
  [12] provider ProviderId{uid:10183, app:10183, cmp:ComponentInfo{com.google.android.apps.youtube.music/com.google.android.apps.youtube.music.player.widget.MusicWidgetProvider}, profile:UserHandle{0}}
    min=(89089x14337)   minResize=(49153x14337) updatePeriodMillis=1800000 resizeMode=1 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0e0049 initialKeyguardLayout=#0   zombie=false
  [13] provider ProviderId{uid:10183, app:10183, cmp:ComponentInfo{com.google.android.apps.youtube.music/com.google.android.apps.youtube.music.player.widget.gm3.FreeformMusicWidgetProvider}, profile:UserHandle{0}}
    min=(38401x25601)   minResize=(38401x25601) updatePeriodMillis=1800000 resizeMode=3 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0e004a initialKeyguardLayout=#0   zombie=false
  [14] provider ProviderId{uid:10160, app:10160, cmp:ComponentInfo{com.google.android.calendar/com.android.calendar.widget.CalendarAppWidgetProvider}, profile:UserHandle{0}}
    min=(28161x28161)   minResize=(28161x28161) updatePeriodMillis=0 resizeMode=3 widgetCategory=3 autoAdvanceViewId=-1 initialLayout=#7f0e0257 initialKeyguardLayout=#0   zombie=false
  [15] provider ProviderId{uid:10160, app:10160, cmp:ComponentInfo{com.google.android.calendar/com.google.android.calendar.widgetmonth.MonthViewWidgetProvider}, profile:UserHandle{0}}
    min=(64001x64001)   minResize=(64001x64001) updatePeriodMillis=0 resizeMode=3 widgetCategory=3 autoAdvanceViewId=-1 initialLayout=#7f0e0257 initialKeyguardLayout=#0   zombie=false
  [16] provider ProviderId{uid:10165, app:10165, cmp:ComponentInfo{com.google.android.contacts/com.google.android.apps.contacts.widget.besties.BestiesWidgetProvider}, profile:UserHandle{0}}
    min=(97281x26625)   minResize=(30721x10241) updatePeriodMillis=86400000 resizeMode=3 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0e01b3 initialKeyguardLayout=#0   zombie=false
  [17] provider ProviderId{uid:10165, app:10165, cmp:ComponentInfo{com.google.android.contacts/com.google.android.apps.contacts.widget.bestiessinglecontact.BestiesSingleContactWidgetProvider}, profile:UserHandle{0}}
    min=(34817x30721)   minResize=(5121x5121) updatePeriodMillis=10800000 resizeMode=3 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0e0662 initialKeyguardLayout=#0   zombie=false
  [18] provider ProviderId{uid:10165, app:10165, cmp:ComponentInfo{com.google.android.contacts/com.google.android.apps.contacts.widget.favoritecontactsgrid.FavoriteContactsWidgetProviderUpgradeWrapper}, profile:UserHandle{0}}
    min=(97281x26625)   minResize=(30721x10241) updatePeriodMillis=86400000 resizeMode=3 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0e01b3 initialKeyguardLayout=#0   zombie=false
  [19] provider ProviderId{uid:10165, app:10165, cmp:ComponentInfo{com.google.android.contacts/com.google.android.apps.contacts.widget.singlecontact.SingleContactWidgetProvider}, profile:UserHandle{0}}
    min=(34817x30721)   minResize=(5121x5121) updatePeriodMillis=10800000 resizeMode=3 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0e0662 initialKeyguardLayout=#0   zombie=false
  [20] provider ProviderId{uid:10173, app:10173, cmp:ComponentInfo{com.google.android.deskclock/com.android.alarmclock.AnalogAppWidgetProvider}, profile:UserHandle{0}}
    min=(28161x28161)   minResize=(1x1) updatePeriodMillis=0 resizeMode=3 widgetCategory=3 autoAdvanceViewId=-1 initialLayout=#7f0e00b2 initialKeyguardLayout=#0   zombie=false
  [21] provider ProviderId{uid:10173, app:10173, cmp:ComponentInfo{com.google.android.deskclock/com.android.alarmclock.StopwatchAppWidgetProvider}, profile:UserHandle{0}}
    min=(38401x38401)   minResize=(28161x28161) updatePeriodMillis=0 resizeMode=3 widgetCategory=3 autoAdvanceViewId=-1 initialLayout=#7f0e00b2 initialKeyguardLayout=#0   zombie=false
  [22] provider ProviderId{uid:10173, app:10173, cmp:ComponentInfo{com.google.android.deskclock/com.android.alarmclock.DigitalAppWidgetProvider}, profile:UserHandle{0}}
    min=(38401x25601)   minResize=(21761x1) updatePeriodMillis=0 resizeMode=3 widgetCategory=3 autoAdvanceViewId=-1 initialLayout=#7f0e00b2 initialKeyguardLayout=#0   zombie=false
  [23] provider ProviderId{uid:10173, app:10173, cmp:ComponentInfo{com.google.android.deskclock/com.android.alarmclock.DigitalStackedAppWidgetProvider}, profile:UserHandle{0}}
    min=(46081x35841)   minResize=(21761x1) updatePeriodMillis=0 resizeMode=3 widgetCategory=3 autoAdvanceViewId=-1 initialLayout=#7f0e00b2 initialKeyguardLayout=#0   zombie=false
  [24] provider ProviderId{uid:10173, app:10173, cmp:ComponentInfo{com.google.android.deskclock/com.android.alarmclock.DigitalCitiesAppWidgetProvider}, profile:UserHandle{0}}
    min=(76801x30721)   minResize=(21761x1) updatePeriodMillis=0 resizeMode=3 widgetCategory=3 autoAdvanceViewId=-1 initialLayout=#7f0e00b2 initialKeyguardLayout=#0   zombie=false
  [25] provider ProviderId{uid:10174, app:10174, cmp:ComponentInfo{com.google.android.gm/com.google.android.gm.widget.GmailWidgetProvider}, profile:UserHandle{0}}
    min=(64001x28161)   minResize=(38401x28161) updatePeriodMillis=79200000 resizeMode=3 widgetCategory=3 autoAdvanceViewId=-1 initialLayout=#7f0e0e4c initialKeyguardLayout=#0   zombie=false
  [26] provider ProviderId{uid:10174, app:10174, cmp:ComponentInfo{com.google.android.gm/com.google.android.gm.widget.GoogleMailWidgetProvider}, profile:UserHandle{0}}
    min=(64001x28161)   minResize=(38401x28161) updatePeriodMillis=79200000 resizeMode=3 widgetCategory=3 autoAdvanceViewId=-1 initialLayout=#7f0e0e4c initialKeyguardLayout=#0   zombie=false
  [27] provider ProviderId{uid:10158, app:10158, cmp:ComponentInfo{com.google.android.googlequicksearchbox/com.google.android.apps.gsa.staticplugins.smartspace.widget.SmartspaceWidgetProvider}, profile:UserHandle{0}}
    min=(84737x12801)   minResize=(84737x12801) updatePeriodMillis=7200000 resizeMode=3 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0e051f initialKeyguardLayout=#0   zombie=false
  [28] provider ProviderId{uid:10158, app:10158, cmp:ComponentInfo{com.google.android.googlequicksearchbox/com.google.android.googlequicksearchbox.SearchWidgetProvider}, profile:UserHandle{0}}
    min=(71681x10241)   minResize=(20481x10241) updatePeriodMillis=0 resizeMode=1 widgetCategory=5 autoAdvanceViewId=-1 initialLayout=#7f0e0a90 initialKeyguardLayout=#0   zombie=false
  [29] provider ProviderId{uid:10158, app:10158, cmp:ComponentInfo{com.google.android.googlequicksearchbox/com.google.android.apps.gsa.staticplugins.searchwidget.GoogleSearchWidgetProvider}, profile:UserHandle{0}}
    min=(71681x10241)   minResize=(20481x10241) updatePeriodMillis=0 resizeMode=1 widgetCategory=5 autoAdvanceViewId=-1 initialLayout=#7f0e0a46 initialKeyguardLayout=#0   zombie=false
  [30] provider ProviderId{uid:10158, app:10158, cmp:ComponentInfo{com.google.android.googlequicksearchbox/com.google.android.apps.search.widgets.stocks.StocksWidgetReceiver}, profile:UserHandle{0}}
    min=(35841x23041)   minResize=(35841x23041) updatePeriodMillis=86400000 resizeMode=3 widgetCategory=3 autoAdvanceViewId=-1 initialLayout=#7f0e113b initialKeyguardLayout=#7f0e113b   zombie=false
  [31] provider ProviderId{uid:10157, app:10157, cmp:ComponentInfo{com.google.android.settings.intelligence/com.google.android.settings.intelligence.modules.batterywidget.impl.BatteryAppWidgetProvider}, profile:UserHandle{0}}
    min=(46081x12289)   minResize=(28161x12289) updatePeriodMillis=3600000 resizeMode=3 widgetCategory=3 autoAdvanceViewId=-1 initialLayout=#7f0e004f initialKeyguardLayout=#0   zombie=false
  [32] provider ProviderId{uid:10182, app:10182, cmp:ComponentInfo{com.google.android.youtube/com.google.android.apps.youtube.app.widget.YtQuickActionsWidgetProvider}, profile:UserHandle{0}}
    min=(48641x30721)   minResize=(48641x12801) updatePeriodMillis=72000000 resizeMode=3 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0e0075 initialKeyguardLayout=#0   zombie=false
  [33] provider ProviderId{uid:10182, app:10182, cmp:ComponentInfo{com.google.android.youtube/com.google.android.apps.youtube.app.widget.YtSearchWidgetProvider}, profile:UserHandle{0}}
    min=(57345x12801)   minResize=(48641x12801) updatePeriodMillis=72000000 resizeMode=3 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0e0075 initialKeyguardLayout=#0   zombie=false
  [34] provider ProviderId{uid:10228, app:10228, cmp:ComponentInfo{com.raju.shingadiya.debug/com.raju.shingadiya.widget.NessoUnreadSmallWidgetReceiver}, profile:UserHandle{0}}
    min=(42497x42497)   minResize=(42497x42497) updatePeriodMillis=1800000 resizeMode=0 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0d017e initialKeyguardLayout=#0   zombie=false
  [35] provider ProviderId{uid:10228, app:10228, cmp:ComponentInfo{com.raju.shingadiya.debug/com.raju.shingadiya.widget.NessoUnreadMediumWidgetReceiver}, profile:UserHandle{0}}
    min=(91137x43009)   minResize=(91137x43009) updatePeriodMillis=1800000 resizeMode=0 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0d017e initialKeyguardLayout=#0   zombie=false
  [36] provider ProviderId{uid:10228, app:10228, cmp:ComponentInfo{com.raju.shingadiya.debug/com.raju.shingadiya.widget.NessoUnreadLargeWidgetReceiver}, profile:UserHandle{0}}
    min=(91137x43009)   minResize=(91137x43009) updatePeriodMillis=1800000 resizeMode=2 widgetCategory=1 autoAdvanceViewId=-1 initialLayout=#7f0d017e initialKeyguardLayout=#0   zombie=false
 
Widgets:
  [0] id=6
    host=HostId{user:0, app:10192, hostId:1025, pkg:com.google.android.apps.nexuslauncher}
    provider=ProviderId{uid:10158, app:10158, cmp:ComponentInfo{com.google.android.googlequicksearchbox/com.google.android.apps.gsa.staticplugins.searchwidget.GoogleSearchWidgetProvider}, profile:UserHandle{0}}
    host.callbacks=com.android.internal.appwidget.IAppWidgetHost$Stub$Proxy@cf3ef99
    views=android.widget.RemoteViews@33d25e
    views_bitmap_memory=0
 
Hosts:
  [0] hostId=HostId{user:0, app:10192, hostId:1025, pkg:com.google.android.apps.nexuslauncher}
    callbacks=com.android.internal.appwidget.IAppWidgetHost$Stub$Proxy@cf3ef99
    widgets.size=1 zombie=false
  [1] hostId=HostId{user:0, app:10192, hostId:1024, pkg:com.google.android.apps.nexuslauncher}
    callbacks=com.android.internal.appwidget.IAppWidgetHost$Stub$Proxy@2dd663f
    widgets.size=0 zombie=false
 
Grants:
```
