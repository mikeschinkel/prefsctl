package prefdefaults

import (
	"github.com/mikeschinkel/prefsctl/macosutil"
)

//goland:noinspection SpellCheckingInspection
func comAppleDock() Domain {
	return Domain{
		Domain: "com.apple.dock",
		Prefs: []DomainPref{
			{
				Name:    "autohide",
				Descr:   "Controls if Dock automatically hides",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
				),
			},
			{
				Name:    "autohide-delay",
				Descr:   "Controls hide/show animation delay",
				Type:    "float",
				Default: "0.2",
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "autohide-time-modifier",
				Descr:   "Control the opening and closing animation times for the dock",
				Type:    "float",
				Default: "0.5",
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "largesize",
				Descr:   "Maximum size of dock icons when magnification is enabled",
				Type:    "int",
				Default: "128",
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "launchanim",
				Descr:   "Dock launching animation",
				Type:    "bool",
				Default: "true",
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "magnification",
				Descr:   "Icon magnification on hover",
				Type:    "bool",
				Default: "false",
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "mineffect",
				Descr:   "Controls the minimization animation of the dock",
				Type:    "string",
				Default: "genie",
				Options: []string{
					"genie",
					"scale",
					"suck",
				},
				Labels: FinalizedUserManaged,
			},
			{
				Name:    "minimize-to-application",
				Descr:   "Controls whether windows minimize into their parent application's Dock icon (true) or to a separate thumbnail on the right side of the Dock (false) where the thumbnail is a screenshot of the window.",
				Type:    "bool",  // Accepts both bool and 0/1 values
				Default: "false", // Not present in pristine install
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "no-bouncing",
				Descr:   "Control whether app icons can 'bounce' for notifications",
				Type:    "bool",
				Default: "false",
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "orientation",
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
				Labels: FinalizedUserManaged,
			},
			{
				Name:    "scroll-to-open",
				Descr:   "Controls if scrolling up on a dock icon shows all Space's opened windows for an app, or open stack",
				Type:    "bool",
				Default: "false",
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "show-recents",
				Descr:   "Controls if recently used apps are shown in a separate section of the dock",
				Type:    "bool",
				Default: "true",
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "showAppExposeGestureEnabled",
				Descr:   "Enables a downward swipe gesture to launch App Expose",
				Type:    "bool",
				Default: "true",
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "showhidden",
				Descr:   "Controls whether hidden applications appear translucent in the Dock or are completely hidden",
				Type:    "bool",
				Default: "false",
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "showMissionControlGestureEnabled",
				Descr:   "Enables an upward swipe gesture to launch Mission Control",
				Type:    "bool",
				Default: "true",
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "static-only",
				Descr:   "Controls if only opened apps are shown in the dock (empties dock when no apps running)",
				Type:    "bool",
				Default: "false",
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "tilesize",
				Descr:   "Size of icons when not hovered and thus not maximized",
				Type:    "int",
				Default: "48",
				Range:   []string{"16", "128"},
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "trash-full",
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels:  FinalizedRuntimeState,
			},
			{
				Name:    "wvous-bl-corner",
				Descr:   "The current action for the bottom-left hot corner",
				Type:    "int",
				Default: CornerActionNone,
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "wvous-bl-modifier",
				Descr:   "The modifier key for the bottom-left hot corner",
				Type:    "int",
				Default: ModifierKeyNone,
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "wvous-br-corner",
				Descr:   "The current action for the bottom-right hot corner",
				Type:    "int",
				Default: CornerActionNone,
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "wvous-br-modifier",
				Descr:   "The modifier key for the bottom-right hot corner",
				Type:    "int",
				Default: ModifierKeyNone,
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "wvous-tl-corner",
				Descr:   "The current action for the top-left hot corner",
				Type:    "int",
				Default: CornerActionNone,
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "wvous-tl-modifier",
				Descr:   "The modifier key for the top-left hot corner",
				Type:    "int",
				Default: ModifierKeyNone,
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "wvous-tr-corner",
				Descr:   "The current action for the top-right hot corner",
				Type:    "int",
				Default: CornerActionNone,
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "wvous-tr-modifier",
				Descr:   "The modifier key for the top-right hot corner",
				Type:    "int",
				Default: ModifierKeyNone,
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "mouse-over-hilite-stack",
				Descr:   `Controls if a squarish background "hilite" is displayed behind each file on mouse hover when a "stack" is expanded and extends over the desktop (true) or not background (false). A stack is a folder on the dock — typically shown on the right side — and the ~/Downloads folder is displayed in the dock by default.`,
				Default: "false",
				Type:    "bool",
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "show-process-indicators",
				Descr:   "Controls if a small black dot appears below each app icon on the dock when the app is currently running (true) or not (false).",
				Default: "true",
				Type:    "bool",
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "mru-spaces",
				Descr:   `Controls whether Mission Control's virtual desktops (Spaces) automatically reorders based on most recent use (true) or maintain their manual arrangement (false). The default of true can be disorienting as it changes the spatial relationship between spaces. (But reorder how? At top of screen or in swipe order? Cannot reproduce and difference in behavior between true or false. Also, see: https://apple.stackexchange.com/questions/214348/how-to-prevent-mac-from-changing-the-order-of-desktops-spaces"`,
				Default: "true",
				Type:    "bool",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			{
				Name:    "contents-immutable",
				Descr:   "Prevents adding or removing dock icons via drag-and-drop when true, and removes those options from the right-click options menu",
				Default: "false",
				Type:    "bool",
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "size-immutable",
				Descr:   "Prevents modifying dock size via the right-click Dock Preferences menu",
				Default: "0",
				Type:    "intBool",
				Labels:  FinalizedUserManaged,
			},
			{
				Name:    "position-immutable",
				Descr:   "Prevents modifying the dock 'orientation' via the right-click Dock Preferences menu from bottom, left, or right. HOWEVER, it appears that at times it is only immutable AFTER orientation has been set via the command line, but if it is then set via the Dock Preferences menu it will become immutable, which is probably not what is expected.",
				Default: "false",
				Type:    "bool",
				Labels:  FinalizedUserManaged,
			},
			{
				Name:  "persistent-apps",
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
				),
			},
			{
				Name:  "persistent-others",
				Descr: "List of 'Keep in Dock' folders/stacks",
				Type:  "unknown",
				Default: `[]string{
					"file://${HOME}/Downloads/",
				}`,
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		AfterApplyFunc: func() error {
			return macosutil.KillDock()
		},
	}
}
