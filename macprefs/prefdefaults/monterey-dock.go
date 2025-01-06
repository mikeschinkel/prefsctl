package prefdefaults

func comAppleDock() DomainPrefs {
	return DomainPrefs{
		"autohide": DomainPref{
			Descr:   "Controls if Dock automatically hides",
			Type:    "bool",
			Default: "false",
			Labels:  FinalizedUserManaged,
		},
		"autohide-delay": DomainPref{
			Descr:   "Controls hide/show animation delay",
			Type:    "float",
			Default: "0.2",
			Labels:  FinalizedUserManaged,
		},
		"autohide-time-modifier": DomainPref{
			Descr:   "Control the opening and closing animation times for the dock",
			Type:    "float",
			Default: "0.5",
			Labels:  FinalizedUserManaged,
		},
		"largesize": DomainPref{
			Descr:   "Maximum size of dock icons when magnification is enabled",
			Type:    "int",
			Default: "128",
			Labels:  FinalizedUserManaged,
		},
		"launchanim": DomainPref{
			Descr:   "Dock launching animation",
			Type:    "bool",
			Default: "true",
			Labels:  FinalizedUserManaged,
		},
		"magnification": DomainPref{
			Descr:   "Icon magnification on hover",
			Type:    "bool",
			Default: "false",
			Labels:  FinalizedUserManaged,
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
			Labels: FinalizedUserManagedWithOptions,
		},
		"minimize-to-application": DomainPref{
			Descr:   "Controls whether windows minimize into their parent application's Dock icon (true) or to a separate thumbnail on the right side of the Dock (false) where the thumbnail is a screenshot of the window.",
			Type:    "bool",  // Accepts both bool and 0/1 values
			Default: "false", // Not present in pristine install
			Labels:  FinalizedUserManaged,
		},
		"no-bouncing": DomainPref{
			Descr:   "Control whether app icons can 'bounce' for notifications",
			Type:    "bool",
			Default: "false",
			Labels:  FinalizedUserManaged,
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
			Requires: []string{
				"com.apple.dock/position-immutable=false",
			},
			Labels: FinalizedUserManagedWithOptions,
		},
		"scroll-to-open": DomainPref{
			Descr:   "Controls if scrolling up on a dock icon shows all Space's opened windows for an app, or open stack",
			Type:    "bool",
			Default: "false",
			Labels:  FinalizedUserManaged,
		},
		"show-recents": DomainPref{
			Descr:   "Controls if recently used apps are shown in a separate section of the dock",
			Type:    "bool",
			Default: "true",
			Labels:  FinalizedUserManaged,
		},
		"showAppExposeGestureEnabled": DomainPref{
			Descr:   "Enables a downward swipe gesture to launch App Expose",
			Type:    "bool",
			Default: "true",
			Labels:  FinalizedUserManaged,
		},
		"showhidden": DomainPref{
			Descr:   "Controls whether hidden applications appear translucent in the Dock or are completely hidden",
			Type:    "bool",
			Default: "false",
			Labels:  FinalizedUserManaged,
		},
		"showMissionControlGestureEnabled": DomainPref{
			Descr:   "Enables an upward swipe gesture to launch Mission Control",
			Type:    "bool",
			Default: "true",
			Labels:  FinalizedUserManaged,
		},
		"static-only": DomainPref{
			Descr:   "Controls if only opened apps are shown in the dock (empties dock when no apps running)",
			Type:    "bool",
			Default: "false",
			Labels:  FinalizedUserManaged,
		},
		"tilesize": DomainPref{
			Descr:   "Size of icons when not hovered and thus not maximized",
			Type:    "int",
			Default: "48",
			Range:   []string{"16", "128"},
			Labels:  FinalizedUserManaged,
		},
		"trash-full": DomainPref{
			Descr:   "",
			Type:    "bool",
			Default: "false",
			Labels:  FinalizedRuntimeState,
		},
		"wvous-bl-corner": DomainPref{
			Descr:   "The current action for the bottom-left hot corner",
			Type:    "int",
			Default: CornerActionNone,
			Labels:  FinalizedUserManaged,
		},
		"wvous-bl-modifier": DomainPref{
			Descr:   "The modifier key for the bottom-left hot corner",
			Type:    "int",
			Default: ModifierKeyNone,
			Labels:  FinalizedUserManaged,
		},
		"wvous-br-corner": DomainPref{
			Descr:   "The current action for the bottom-right hot corner",
			Type:    "int",
			Default: CornerActionNone,
			Labels:  FinalizedUserManaged,
		},
		"wvous-br-modifier": DomainPref{
			Descr:   "The modifier key for the bottom-right hot corner",
			Type:    "int",
			Default: ModifierKeyNone,
			Labels:  FinalizedUserManaged,
		},
		"wvous-tl-corner": DomainPref{
			Descr:   "The current action for the top-left hot corner",
			Type:    "int",
			Default: CornerActionNone,
			Labels:  FinalizedUserManaged,
		},
		"wvous-tl-modifier": DomainPref{
			Descr:   "The modifier key for the top-left hot corner",
			Type:    "int",
			Default: ModifierKeyNone,
			Labels:  FinalizedUserManaged,
		},
		"wvous-tr-corner": DomainPref{
			Descr:   "The current action for the top-right hot corner",
			Type:    "int",
			Default: CornerActionNone,
			Labels:  FinalizedUserManaged,
		},
		"wvous-tr-modifier": DomainPref{
			Descr:   "The modifier key for the top-right hot corner",
			Type:    "int",
			Default: ModifierKeyNone,
			Labels:  FinalizedUserManaged,
		},

		"mouse-over-hilite-stack": DomainPref{
			Descr:   `Controls if a squarish background "hilite" is displayed behind each file on mouse hover when a "stack" is expanded and extends over the desktop (true) or not background (false). A stack is a folder on the dock — typically shown on the right side — and the ~/Downloads folder is displayed in the dock by default.`,
			Default: "false",
			Type:    "bool",
			Labels:  FinalizedUserManaged,
		},
		"show-process-indicators": DomainPref{
			Descr:   "Controls if a small black dot appears below each app icon on the dock when the app is currently running (true) or not (false).",
			Default: "true",
			Type:    "bool",
			Labels:  FinalizedUserManaged,
		},
		"mru-spaces": DomainPref{
			Descr:   `Controls whether Mission Control's virtual desktops (Spaces) automatically reorders based on most recent use (true) or maintain their manual arrangement (false). The default of true can be disorienting as it changes the spatial relationship between spaces. (But reorder how? At top of screen or in swipe order? Cannot reproduce and difference in behavior between true or false. Also, see: https://apple.stackexchange.com/questions/214348/how-to-prevent-mac-from-changing-the-order-of-desktops-spaces"`,
			Default: "true",
			Type:    "bool",
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
				DescrNotVerified,
			),
		},
		"contents-immutable": DomainPref{
			Descr:   "Prevents adding or removing dock icons via drag-and-drop when true, and removes those options from the right-click options menu",
			Default: "false",
			Type:    "bool",
			Labels:  FinalizedUserManaged,
		},

		"size-immutable": DomainPref{
			Descr:   "Prevents modifying dock size via the right-click Dock Preferences menu",
			Default: "0",
			Type:    "intBool",
			Labels:  FinalizedUserManaged,
		},

		"position-immutable": DomainPref{
			Descr:   "Prevents modifying the dock 'orientation' via the right-click Dock Preferences menu from bottom, left, or right. HOWEVER, it appears that at times it is only immutable AFTER orientation has been set via the command line, but if it is then set via the Dock Preferences menu it will become immutable, which is probably not what is expected.",
			Default: "false",
			Type:    "bool",
			Labels:  FinalizedUserManaged,
		},

		"persistent-apps": DomainPref{
			Descr: "List of 'Keep in Dock' applications",
			Type:  "unknown",
			Default: `[]string{
				"file:///System/Applications/Launchpad.app/",
				"file:///Applications/Safari.app/",
				"file:///System/Applications/Messages.app/",
				"file:///System/Applications/Mail.app/",
				"file:///System/Applications/Maps.app/",
				"file:///System/Applications/Photos.app/",
				"file:///System/Applications/FaceTime.app/",
				"file:///System/Applications/Calendar.app/",
				"file:///System/Applications/Contacts.app/",
				"file:///System/Applications/Reminders.app/",
				"file:///System/Applications/Notes.app/",
				"file:///System/Applications/TV.app/",
				"file:///System/Applications/Music.app/",
				"file:///System/Applications/Podcasts.app/",
				"file:///System/Applications/News.app/",
				"file:///System/Applications/App%20Store.app/",
				"file:///System/Applications/System%20Preferences.app/",
			}`,
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeNotVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
				DescrVerified,
			),
		},
		"persistent-others": DomainPref{
			Descr: "List of 'Keep in Dock' folders/stacks",
			Type:  "unknown",
			Default: `[]string{
				"file://${HOME}/Downloads/",
			}`,
			Labels: NewLabels(
				DefaultsSet,
				UserManaged,
				TypeNotVerified,
				ClassVerified,
				SetsVerified,
				DefaultVerified,
				DescrVerified,
			),
		},
	}
}
