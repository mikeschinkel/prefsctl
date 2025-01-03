package prefdefaults

// TODO: Add these
//mouse-over-hilite-stack: Controls highlight effect on stacks
//persistent-apps: List of permanently docked applications
//persistent-others: List of permanently docked folders/stacks
//mru-spaces: Controls if Spaces reorder based on usage
//show-process-indicators: Controls running app indicators (dots)
//minimize-to-application: Controls if windows minimize to app icon
//contents-immutable: Prevents dock contents modification
//size-immutable: Prevents dock size modification
//position-immutable: Prevents dock position modification

//mouse-over-hilite-stack: Default: true
//show-process-indicators: Default: true (the little dots under running apps)
//mru-spaces: Default: true (spaces reorder based on use)
//For these persistence-related ones, they're trickier because they contain arrays/lists:
//persistent-apps: Default: Contains default apps like Finder, System Settings, etc.
//persistent-others: Default: Empty array (or contains Downloads folder depending on macOS version)
//contents-immutable: false
//size-immutable: false
//position-immutable: false

func comAppleDock() DomainPrefs {
	return DomainPrefs{
		"autohide": DomainPref{
			Descr:   "Controls if Dock automatically hides",
			Type:    "bool",
			Default: "false",
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"autohide-delay": DomainPref{
			Descr:   "Controls hide/show animation delay",
			Type:    "float",
			Default: "0.2",
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"show-recents": DomainPref{
			Descr:   "Controls if recently used apps are shown in a separate section of the dock",
			Type:    "bool",
			Default: "true",
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"static-only": DomainPref{
			Descr:   "Controls if only opened apps are shown in the dock (empties dock when no apps running)",
			Type:    "bool",
			Default: "false",
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"scroll-to-open": DomainPref{
			Descr:   "Controls if scrolling up on a dock icon shows all Space's opened windows for an app, or open stack",
			Type:    "bool",
			Default: "false",
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"largesize": DomainPref{
			Descr:   "Maximum size of dock icons when magnification is enabled",
			Type:    "int",
			Default: "128",
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"mineffect": DomainPref{
			Descr:   "Controls the minimization animation of the dock",
			Type:    "string",
			Default: "genie",
			Options: []string{
				"genie",
				"scale",
				"suck",
			},
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
				OptionsVerified,
			),
		},
		"launchanim": DomainPref{
			Descr:   "Dock launching animation",
			Type:    "bool",
			Default: "true",
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"magnification": DomainPref{
			Descr:   "Icon magnification on hover",
			Type:    "bool",
			Default: "false",
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"minimize-to-application": DomainPref{
			Descr:   "Controls if minimizing windows minimize to the app's icon",
			Type:    "bool",
			Default: "false",
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"no-bouncing": DomainPref{
			Descr:   "Control whether app icons can 'bounce' for notifications",
			Type:    "bool",
			Default: "false",
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"showAppExposeGestureEnabled": DomainPref{
			Descr:   "Enables a downward swipe gesture to launch App Expose",
			Type:    "bool",
			Default: "true",
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"showhidden": DomainPref{
			Descr:   "Controls whether hidden applications appear translucent in the Dock or are completely hidden",
			Type:    "bool",
			Default: "false",
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"showMissionControlGestureEnabled": DomainPref{
			Descr:   "Enables an upward swipe gesture to launch Mission Control",
			Type:    "bool",
			Default: "true",
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"tilesize": DomainPref{
			Descr:   "Size of icons when not hovered and thus not maximized",
			Type:    "int",
			Default: "48",
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"trash-full": DomainPref{
			Descr:   "",
			Type:    "bool",
			Default: "false",
			Labels: NewLabels(
				DefaultsSet,
				RuntimeState,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"wvous-tr-corner": DomainPref{
			Descr:   "The current action for the top-right hot corner",
			Type:    "int",
			Default: CornerActionNone,
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"wvous-tr-modifier": DomainPref{
			Descr:   "The modifier key for the top-right hot corner",
			Type:    "int",
			Default: ModifierKeyNone,
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"wvous-tl-corner": DomainPref{
			Descr:   "The current action for the top-left hot corner",
			Type:    "int",
			Default: CornerActionNone,
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"wvous-tl-modifier": DomainPref{
			Descr:   "The modifier key for the top-left hot corner",
			Type:    "int",
			Default: ModifierKeyNone,
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"wvous-br-corner": DomainPref{
			Descr:   "The current action for the bottom-right hot corner",
			Type:    "int",
			Default: CornerActionNone,
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"wvous-br-modifier": DomainPref{
			Descr:   "The modifier key for the bottom-right hot corner",
			Type:    "int",
			Default: ModifierKeyNone,
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"wvous-bl-corner": DomainPref{
			Descr:   "The current action for the bottom-left hot corner",
			Type:    "int",
			Default: CornerActionNone,
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"wvous-bl-modifier": DomainPref{
			Descr:   "The modifier key for the bottom-left hot corner",
			Type:    "int",
			Default: ModifierKeyNone,
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
			),
		},
		"orientation": DomainPref{
			Descr:   "Location for the dock on screen",
			Type:    "string",
			Default: "bottom",
			Options: []string{
				"left",
				"bottom",
				"right",
			},
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				DefaultVerified,
			),
		},
		"autohide-time-modifier": DomainPref{
			Descr:   "Control the opening and closing animation times for the dock",
			Type:    "float",
			Default: "0.5",
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				DefaultVerified,
			),
		},
		//"loc": DomainPref{
		//	Descr:   "",
		//	Type:    "string",
		//	Default: "en_US:US",
		//	Labels: NewLabels(
		//		SetupSets,
		//		UserManaged,
		//		TypeVerified,
		//		ClassVerified,
		//	),
		//},
		//"region": DomainPref{
		//	Descr:   "",
		//	Type:    "string",
		//	Default: "US",
		//	Labels: NewLabels(
		//		SetupSets,
		//		UserManaged,
		//		TypeVerified,
		//		ClassVerified,
		//	),
		//},
	}
}
