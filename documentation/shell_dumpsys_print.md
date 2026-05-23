# `adbjson shell dumpsys print`

## adbjson

**Command:**
```bash
adbjson shell dumpsys print
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
adb shell dumpsys print
```

**Output:**
```
PRINT MANAGER STATE (dumpsys print)
{
  user_states={
    user_id=0
    installed_services={
      component_name={
        package_name=com.android.bips
        class_name=com.android.bips.BuiltInPrintService
      }
      add_printers_activity=com.android.bips.ui.AddPrintersActivity
      advanced_options_activity=com.android.bips.ui.MoreOptionsActivity
    }
    actives_services={
      component_name={
        package_name=com.android.bips
        class_name=com.android.bips.BuiltInPrintService
      }
      is_destroyed=false
      is_bound=false
      has_discovery_session=false
      has_active_print_jobs=false
      is_discovering_printers=false
    }
    print_spooler_state={
      is_destroyed=false
      is_bound=true
      internal_state={
      }
    }
  }
}
```
