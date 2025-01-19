package prefdefaults

// Corner modifiers - stable across macOS versions
const (
	ModifierKeyNone    string = "0"
	ModifierKeyShift   string = "1048576"
	ModifierKeyControl string = "524288"
	ModifierKeyOption  string = "262144"
	ModifierKeyCommand string = "131072"
)

// Corner actions - may vary by macOS version
const (
	CornerActionNone               string = "0"
	CornerActionMissionControl     string = "1"
	CornerActionAppWindows         string = "2"
	CornerActionDesktop            string = "3"
	CornerActionDashboard          string = "4" // Deprecated
	CornerActionScreenSaver        string = "5"
	CornerActionDisableScreenSaver string = "6"
	CornerActionNotificationCenter string = "7"
	CornerActionSleep              string = "10"
	CornerActionLaunchpad          string = "11"
	CornerActionQuickNote          string = "12" // Added in newer versions
)
