package prefdefaults

import (
	"github.com/mikeschinkel/prefsctl/macprefs"
)

type DomainDefaults = macprefs.DomainPrefDefaults
type DomainPrefs = macprefs.PrefDefaultsMap
type Pref = macprefs.PrefDefault

//goland:noinspection SpellCheckingInspection
func montereyPrefDefaults() DomainDefaults {
	return DomainDefaults{
		"GlobalPreferences": DomainPrefs{
			"AppleActionOnDoubleClick": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"AppleAntiAliasingThreshold": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"AppleAquaColorVariant": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"AppleEnableSwipeNavigateWithScrolls": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"AppleKeyboardUIMode": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"AppleMeasurementUnits": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"AppleMetricUnits": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"AppleMiniaturizeOnDoubleClick": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"AppleShowAllExtensions": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"AppleShowScrollBars": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"AppleTemperatureUnit": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"AppleWindowTabbingMode": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.mouse.scaling": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.sound.beep.flash": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.springing.delay": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.springing.enabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.swipescrolldirection": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.trackpad.forceClick": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.trackpad.scaling": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"InitialKeyRepeat": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"KeyRepeat": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NSAutomaticCapitalizationEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NSAutomaticDashSubstitutionEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NSAutomaticPeriodSubstitutionEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NSAutomaticQuoteSubstitutionEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NSAutomaticSpellingCorrectionEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NSAutomaticTextCompletionEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NSAutomaticWindowAnimationsEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"WebAutomaticSpellingCorrectionEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"WebKitDeveloperExtras": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"AppleLanguage": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
					&macprefs.SetupSets,
					&macprefs.StringType,
				},
			},
			"AppleLocale": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SetupSets,
					&macprefs.StringType,
				},
			},
			"646F6E7A_00000000_00000001_6E7A6361_696D6963:": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"AKLastIDMSEnvironment:": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"AudioQuest inc.AudioQuest DragonFly:": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"Blue Microphones Yeti Stereo Microphone:": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"Canon MF5900 Series:": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"ContextMenuGesture:": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"Fujitsu ScanSnap iX500:": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"Yubico YubiKey OTP+FIDO+CCID:": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"AppleLanguagesDidMigrate:": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"NavPanelFileListModeForOpenMode:": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"NavPanelFileListModeForSaveMode:": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NSLinguisticDataAssetsRequestLastInterval:": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"QLPanelAnimationDuration:": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.security.sosaccount": DomainPrefs{
			"SOSEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.SafariCloudHistoryPushAgent": DomainPrefs{
			"AcknowledgedPushNotifications": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"UnacknowledgedPushNotifications": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.print.custompresets.forprinter.Canon_MF5900_Series": DomainPrefs{
			"com.apple.print.lastPresetPref": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"com.apple.print.lastPresetPrefType": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.finder": DomainPrefs{
			"SidebarTagsSctionDisclosedState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"QLEnableTextSelection": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"PreviewPaneGalleryWidth": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"FK_ArrangeBy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"PreferencesWindow.LastSelection": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"FXEnableExtensionChangeWarning": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NSWindowTabbingShoudShowTabBarKey-com.apple.finder.TBrowserWindow": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"FXICloudDriveDesktop": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ShowMountedServersOnDesktop": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"SidebarShowingiCloudDesktop": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"FXToolbarUpgradedToTenNine": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"ShowRemovableMediaOnDesktop": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ShowStatusBar": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"FXSidebarUpgradedSharedSearchToTwelve": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"FXInfoWindowWidth": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"DownloadsFolderListViewSettingsVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"SidebariCloudDriveSectionDisclosedState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AppleShowAllFiles": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"FXICloudDriveDeclinedUpgrade": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"FXArrangeGroupViewBy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"RecentsArrangeGroupViewBy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"FXPreferredSearchViewStyleVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"PreviewPaneWidth": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"FXAtLeastOneScreenHasBeenInHIDPI": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"CopyProgressWindowLocation": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"SidebarShowingSignedIntoiCloud": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ViewOptionsWindow.Location": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ShowHardDrivesOnDesktop": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"FXQuickActionsConfigUpgradeLevel": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"FXPreferredViewStyle": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"QuitMenuItem": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"FXDefaultSearchScope": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"FXPreferredSearchViewStyle": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"DisableAllAnimations": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"SidebarWidth": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ShowSidebar": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TagsCloudSerialNumber": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NewWindowTarget": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ShowPathbar": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"FXSidebarUpgradedToTenFourteen": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastTrashState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ShowExternalHardDrivesOnDesktop": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"GoToField": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"_FXSortFoldersFirstOnDesktop": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"FK_AppCentricShowSidebar": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"FXSidebarUpgradedToTenTen": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"FXConnectToLastURL": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"FXICloudDriveEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"FXICloudDriveDocuments": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"_FXSortFoldersFirst": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"_FXShowPosixPathInTitle": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"EmptyTrashProgressWindowLocation": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"SearchRecentsSavedViewStyle": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"FXDesktopTouchBarUpgradedToTenTwelveOne": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"FXICloudDriveFirstSyncDownComplete": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"SidebarDevicesSectionDisclosedState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"MountProgressWindowLocation": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"FK_SidebarWidth": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"SearchRecentsSavedViewStyleVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"FXPreferredGroupBy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"PreviewPaneInfoExpanded": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"FXLastSearchScope": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"InspectorWindow.Location": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"SidebarPlacesSectionDisclosedState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"FXToolbarUpgradedToTenEight": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"FXICloudLoggedIn": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"FXToolbarUpgradedToTenSeven": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.messages.pinning": DomainPrefs{
			"IMPinningPinConfigMigrationKey-v2-r2": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.identityservicesd": DomainPrefs{
			"ReRegisteredForDevicesv56": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"hasRegIdentityContainer": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"ReRegisteredForDevicesv55": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"persister-migration-personal-session-token-cache-v4": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"ReRegisteredForDevices": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"DidCleanLegacyAccountPrefs": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"ImportedLegacyIMAccounts": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"hasUnregIdentityContainer": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"MigratedToNewDisabledState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"triggeredRemoteSessionVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"ReRegisteredForDevicesHash": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"ImportedLegacyIDSAccounts2": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"GDRRequestMadeForRelayRepair": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"persister-migration-com.apple.identityservices.userStore": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"ReRegisteredForDevicesv1350": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"performed-user-intent-migrate": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.Spotlight": DomainPrefs{
			"ModelName": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.AppManaged,
				},
			},
			"engagementCount-com.apple.Spotlight": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"engagementCount-com.apple.mail": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"SSMessageTracingWindowShowCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"windowHeight": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"version": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"useCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"__NSEnableTSMDocumentWindowLevel": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"SSMessageTracingWindowHideCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"engagementCount-com.apple.Spotlight.suggestions": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"SPMessageTracingWindowHideCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"showedFTE": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"SPMessageTracingWindowShowCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"userHasMovedWindow": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.DictionaryServices": DomainPrefs{
			"DCSPreferenceVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.PTPCamera": DomainPrefs{
			"FilenamePrefix": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NextFilenameCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.imtranscoding.IMTranscoderAgent": DomainPrefs{
			"kCKMediaObjectManagerDefaultsOSVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.scriptmenu": DomainPrefs{
			"ScriptMenuEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.TTY": DomainPrefs{
			"RTTSettingsVersionPreference": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"RTTContinuityRTTIsSupportedPreference": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"TUSupportsRelayCallingPreference": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"TUIsRelayCallingEnabledPreference": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.voicetrigger": DomainPrefs{
			"VoiceTrigger Enabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.DiskUtility": DomainPrefs{
			"OperationProgress ExpandedHeight": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"DUDebugMenuEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"advanced-image-options": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"SidebarShowAllDevices": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"OperationProgress DetailsVisible": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.photos.shareddefaults": DomainPrefs{
			"CPLDefaultDownload": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"iCloudPhotoLibraryLastResetWelcomePromptDBVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"downloadAndKeepOriginals": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"clearPurgeableResources": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"DiskSpaceWasLow": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"repushAssetsWithImportedByBundleIdentifier": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.AppManaged,
				},
			},
			"repushVideoAssetsMetadata": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.AppManaged,
				},
			},
		},
		"com.apple.parsecd": DomainPrefs{
			"PARDefaultsVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"PARTestSeed": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.universalaccess": DomainPrefs{
			"login": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"switchCoalescePressesDuration": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"lastNightShiftEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"switchAutoScanPanelInterval": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"closeViewZoomedIn": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"grayscale": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"switchSweepingCursorSpeed": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"sessionChange": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"switchHoldBeforeRepeatDuration": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"closeViewZoomFactorBeforeTermination": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"switchHideUITimeout": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"switchAutoScanElementInterval": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"lastNightShiftActive": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"virtualKeyboardOnOff": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"headMouseEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"increaseContrast": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"mouseDriverCursorSize": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"selectedTab": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"lastNightShiftDate": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"closeViewZoomDisplayID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"hoverTextIsAlwaysOn": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"closeViewSplitScreenRatio": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"hoverTextIsHoveringAndVisible": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AssistiveControlType": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"closeViewDesiredZoomFactor": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"slowKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"closeViewScrollWheelToggle": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"mouseDriver": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"differentiateWithoutColor": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"slowKeyDelay": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"switchMinimumPressDuration": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"closeViewZoomFactor": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"switchOnOffKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"keyboardAccessFocusRingTimeout": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"stickyKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"reduceTransparency": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"dwellHideUITimeout": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"dwellEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"contrast": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"useStickyKeysShortcutKeys": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"lastNightShiftMode": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"hoverTextEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"whiteOnBlack": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"closeViewHotkeysEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"switchFirstElementDelay": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"voiceOverOnOffKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.Terminal": DomainPrefs{
			"HasMigratedDefaults": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"SecureKeyboardEntry": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Default Window Settings": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"DefaultProfilesVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"ProfileCurrentVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"TTAppPreferences Selected Tab": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"Startup Window Settings": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.coreservices.useractivityd": DomainPrefs{
			"kLocalPasteboardBlobName": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"kRemotePasteboardBlobName": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.ColorSyncUtility": DomainPrefs{
			"CSUtilitiesContentSize": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"CSProfsUtilsGroupBy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"CSUtilitiesSelectedIdentifier": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.weather": DomainPrefs{
			"userId": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.Music": DomainPrefs{
			"dontAskForAllPlaylistItemRemoval": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"storeSupportsPasswordSettings": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"notifications-warming-last-displayed-time": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"miniplayerVPos": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"com.apple.Music.WebKit2ShouldPrintBackgrounds": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"storeSupportsPreviousPurchases": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"com.apple.Music.WebKit2AsynchronousSpellCheckingEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"miniPlayerQueueVisible": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"eqWindowHPos": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"userWantsPlaybackNotifications": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"miniplayerUserSetHeight": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"com.apple.Music.WebKit2ShowsToolTipOverTruncatedText": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"soundEnhancerAmount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"preloadFilesIntoMemory": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"updateLevel": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"videoWindowVSize": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"com.apple.Music.WebKit2EnableInheritURIQueryComponent": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"automaticallyDownloadArtwork": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"refreshedHLSKeysTime": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"miniplayerWidthInDesignCoords": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"userMaxConcurrentDownloads": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.Music.WebKit2SuppressesIncrementalRendering": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"miniplayerHPos": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"videoWindowVPos": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"log-PlayQueue-LastSelectedTab": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"dontWarnDownloadCloudPurchases": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"storeSupportsUPP": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"showWelcomeScreenState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"checkedHLSKeysTime": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"storeSupportsCloudPurchases": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"gridViewSearchDoesNotSwitchView": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"eqName": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ITUserPrefsMigrated": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"cddbPrefsOK": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"eqPrefsVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"playbackIsFullscreen": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"debugAssertCategoriesVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"debugAssertCategoriesEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"firstLaunchShowWelcomeScreenState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"encoderName": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.Music.WebKit2UserInterfaceDirectionPolicy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"haveRadioState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"videoWindowHSize": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"eqWindowVPos": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"haveAskedToCheckForAvailableDownloads": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"musicVideoPlaybackLocation": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"volumeWSG": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"dontShowRestrictionsPrefs": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"persistentUserID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"NSApplicationCrashOnExceptions": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"miniPlayerLargeArtVisible": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"searchSavedTab": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"artworkDownloadDSID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"videoWindowHPos": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"controllableInterfaceGUID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"userHasActivatedFullScreenVisualizer": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"storeMediaTypeFlags": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"imported-eq-presets": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"soundEnhancerEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"miniplayerSnapMode": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"doesStoreSupportCloudMusicLibrary": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"eqEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.airport.airportutility": DomainPrefs{
			"dontPerformBaseRestartWarning": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"version": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.studentd": DomainPrefs{
			"LastDateProviderSessionToken": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"DeviceIdentifier": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.systemuiserver": DomainPrefs{
			"__NSEnableTSMDocumentWindowLevel": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.IMCoreSpotlight": DomainPrefs{
			"IMCSIdxVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"IMCSLastIndexedRowID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"IMCSNeedsIndexing": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"IMCSIndexTotalRecords": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"IMCSIdxProductVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"IMCSBypassIndexVersionCheckV2": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.bird": DomainPrefs{
			"optimize-storage": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"icloud-drive.account-migration-status.294735135": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.xpc.activity2": DomainPrefs{
			"ProductBuildVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.iPod": DomainPrefs{
			"com.apple.PreferenceSync.ExcludeAllSyncKeys": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.newscore2": DomainPrefs{
			"report_concern_user_id": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"provider_user_id": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"instance_identifier": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.cloudd": DomainPrefs{
			"com.apple.private.cloudkit.shouldUseGeneratedDeviceID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.driver.AppleBluetoothMultitouch.trackpad": DomainPrefs{
			"TrackpadFourFingerHorizSwipeGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadPinch": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"USBMouseStopsTrackpad": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadFourFingerVertSwipeGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadRotate": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadHorizScroll": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadTwoFingerDoubleTapGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"version": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"TrackpadMomentumScroll": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadFourFingerPinchGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"UserPreferences": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Clicking": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadThreeFingerTapGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"HIDScrollZoomModifierMask": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadScroll": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"DragLock": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadFiveFingerPinchGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadThreeFingerHorizSwipeGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadThreeFingerVertSwipeGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadTwoFingerFromRightEdgeSwipeGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ForceSuppressed": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"TrackpadCornerSecondaryClick": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Dragging": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadHandResting": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadRightClick": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadThreeFingerDrag": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.stocks.account": DomainPrefs{
			"deleteOnNextLaunch": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.preference.trackpad": DomainPrefs{
			"ForceClickSavedState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.SetupAssistant": DomainPrefs{
			"PreviousSystemVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeSyncSetup2": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"SkipExpressSettingsUpdating": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"LastSeenDiagnosticsProductVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"SkipFirstLoginOptimization": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"LastSeenCloudProductVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeCloudSetup": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"PreviousBuildVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeApplePaySetup": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeTrueTonePrivacy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeScreenTime": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"MiniBuddyShouldLaunchToResumeSetup": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"DidSeeTouchIDSetup": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"NSAddServicesToContextMenus": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"DidSeeActivationLock": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeSyncSetup": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeAvatarSetup": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"MiniBuddyLaunchReason": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"DidSeeSiriSetup": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeTrueTone": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeePrivacy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastPrivacyBundleVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastSeenSiriProductVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeAppStore": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastPreLoginTasksPerformedVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeAppearanceSetup": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastSeenSyncProductVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastSeeniCloudStorageServicesProductVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastPreLoginTasksPerformedBuild": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"MiniBuddyLaunchedPostMigration": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"DidSeeAccessibility": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastSeenBuddyBuildVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeiCloudLoginForStorageServices": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.ScriptEditor2": DomainPrefs{
			"OSAStandardAdditions ChooseApplication Bounds": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.preferences.extensions.ServicesWithUI": DomainPrefs{
			"userHasOrdered": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.MobileSMS": DomainPrefs{
			"TextSize": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"SyndicationOnboardingVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"LegacyAppSidebarPersistedWidth": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"QuickSaveHasBeenUsedBefore": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"CatalystPreferenceMigrationVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"PlaySoundsKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"BusinessChatPrivacyPageDisplayed": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"sForceUnknownFilteringCompleted": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"CKLastSelectedItemIdentifier": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"kCKMediaObjectManagerDefaultsOSVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"CatalystCustomFontMigrationVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"NotifyAboutMessagesFromUnknownContacts": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"CKMigratedAutoSpamReports26375262": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"SidebarPersistedWidth": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ReadReceiptSettingsConfirmed": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"sForceSMSSpamFilteringCompleted": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"KeepMessageForDays": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"PendingCleared": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"TextFontSize": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.preferences.extensions.ShareMenu": DomainPrefs{
			"userHasOrdered": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.mmcs": DomainPrefs{
			"report.LastFailedCheckDate": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"report.sha256": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"report.LastSuccessfulCheckDate": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"report.TTL": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.appstored": DomainPrefs{
			"BadgeCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ArcadeSubscriptionState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ArcadePayoutDeviceID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"WelcomeNotificationExcludedFromSample": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ArcadeDeviceID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"ArcadeDeviceGUID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"LastOSBuildVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"WelcomeNotificationLastAppStoreConnectionProductVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.PubSubAgent": DomainPrefs{
			"TigerMigrationCompleted": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.commcenter.callservices": DomainPrefs{
			"last.known.icloud.id": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"associated.account": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.print.PrinterProxy": DomainPrefs{
			"IK_Scanner_downloadURL": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"IK_Scanner_selectedTag": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"IK_Scanner_PostProcessApplication": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.amsengagementd": DomainPrefs{
			"AMSMetricsIdentifierUserRecordName": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"AMSMetricsIdentifierZoneCreated": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AMSMetricsIdentifierZoneSubscriptionCreated": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.iChat": DomainPrefs{
			"SaveConversationsOnClose": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"PMPrintingExpandedStateForPrint2": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"UnifiedChatWindowControllerSelectionGUIDSet": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"DidCheckForDuplicateChats": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"HasPromptedSMSRelay": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"HasPromptediMessageFTU": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"DidRegenerateGroupID63841559": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"messageTracerSMSSent": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"messageTracerSubmitDate": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"BuddyPictureSetToGenericByUser": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"messageTracerIMessagesSent": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"PlaySoundsKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"messageTracerSMSUsed": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"messageTracerNumUpgradeOffers": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ChatWindowControllerUnifiedIsVisible": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"messageTracerMessagesSent": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"WebIconDatabaseEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"DidMarkGroupPhotosAsUnpurgeable": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"messageTracerUpgradesAccepted": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"messageTracerIMessagesReceived": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"DaemonConnectionWaitTime": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"NotifyAboutMessagesFromUnknownContacts": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"KeepMessagesVersionID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"messageTracerCharactersSent": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"messageTracerMessagesReceived": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"messageTracerUpgradesDeclined": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"messageTracerSMSReceived": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"KeepMessageForDays": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"UnifiedChatListViewControllerAutomaticallySortsChats": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"messageTracerIMessageUsed": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"DidMigratePersonCentricIDs": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastIMDNotificationPostedDate": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ChatWindowControllerUnifiedChatListWidth": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"LastFailedMessageIMDNotificationPostedDate": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"SODefaultTranscriptName": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"imageBrowser.disableOpenGL": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.TextEdit": DomainPrefs{
			"PlainTextEncoding": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"RichText": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"PlainTextEncodingForWrite": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.HIToolbox": DomainPrefs{
			"AppleFnUsageType": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"AppleCurrentKeyboardLayoutInputSourceID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"AppleDictationAutoEnable": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.preferences.extensions.FinderSync": DomainPrefs{
			"userHasOrdered": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.coreservices.uiagent": DomainPrefs{
			"CSUIHaveCenteredGatekeeperProgressWindow": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"CSUIHasSafariBeenLaunched": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"CSUIRecommendSafariBackOffInterval": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"CSUILastOSVersionWhereSafariRecommendationWasMade": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.CharacterPicker": DomainPrefs{
			"version": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.CharacterPaletteIM": DomainPrefs{
			"CVFontSize": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"CVStartAsLargeWindow": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"version": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.assistant.backedup": DomainPrefs{
			"Hands Free Mode": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Session Language": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Cloud Sync Enabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Cloud Sync User ID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"MultiUser VoiceIdentification Enabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.icloud.fmfd.notbackedup": DomainPrefs{
			"kKeychainUtilMigrationVersionKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.security.cloudkeychainproxy3.keysToRegister": DomainPrefs{
			"EnsurePeerRegistration": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"DSID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"KeyAccountUUID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.mediaaccessibility.public": DomainPrefs{
			"MACaptionDisplayType": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.speech.recognition.AppleSpeechRecognition.prefs": DomainPrefs{
			"DictationIMPresentedOfflineUpgradeSuggestion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"DictationIMCommandsWindowIsOpen": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"DictationIMCorrectionCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"DictationIMUpgradedTo10_16": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"DictationIMUpgradedTo10_15": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"DictationIMPlaySoundUponRecognition": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"DictationIMLocaleIdentifier": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"DictationIMMessageTracesSinceLastReport": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"DictationIMUseOnlyOfflineDictation": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.loginwindow": DomainPrefs{
			"TALLogoutSavesState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"MiniBuddyLaunch": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"oneTimeSSMigrationComplete": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.preferences.softwareupdate": DomainPrefs{
			"LatestMajorOSSeenByUserBundleIdentifier": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.accessibility.universalAccessAuthWarn": DomainPrefs{
			"ThirdPartyCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.internetconnect": DomainPrefs{
			"ServiceID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.TextInputMenu": DomainPrefs{
			"visible": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.messages.nicknames": DomainPrefs{
			"NicknameAppleIDAndiCloudAccountMatchAndExist": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"MeCardPendingNicknamesVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"MeCardSharingVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"MeCardSharingEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"MeCardSharingImageForkedFromMeCard": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"MeCardWhitelistBlacklistNicknamesVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"NicknameInfoRequestedFromPeers": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"MeCardSharingAudience": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NicknameScrutinyDoNotHandle": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"ReuploadLocalNicknamesVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.ProblemReporter": DomainPrefs{
			"ShowComments": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ShowDetails": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.gamed": DomainPrefs{
			"natType": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"GKStoreFrontIdentifier": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"GKActiveEnvironmentKeyV2": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"GKLastPushTokenPlayerID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"GKLastPushTokenEnvironment": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"GKPushEnvironmentKeyV2": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"GKLastProtocolVersionUsedKeyV2": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"GKLoginCancelled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.Safari.SandboxBroker": DomainPrefs{
			"DidMigrateDownloadFolderToSandbox": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"ShowDevelopMenu": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"DidMigrateResourcesToSandbox": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidMigrateDownloadMetadataToSandboxGroupContainer": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.preferences.extensions.XcodeSourceEditor": DomainPrefs{
			"userHasOrdered": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.AppleMediaServices.notbackedup": DomainPrefs{
			"AMSDidRetrieveDeviceOffers": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AMSDidRetrieveDeviceOffersEligibility": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.imdsmsrecordstore": DomainPrefs{
			"DeleteSequenceNumber": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.iTunes": DomainPrefs{
			"com.apple.iTunes.WebKit2ShowsToolTipOverTruncatedText": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"doesAccountArtistListHaveSharePermission": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"dontAutomaticallySyncIPods": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"MacBuddyStoreID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"com.apple.iTunes.WebKit2ShouldPrintBackgrounds": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"com.apple.iTunes.WebKit2SuppressesIncrementalRendering": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"com.apple.iTunes.WebKit2AsynchronousSpellCheckingEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AutomaticDeviceBackupsDisabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"disableRadio": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Store Apple ID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"debugAssertCategoriesVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"com.apple.iTunes.WebKit2UserInterfaceDirectionPolicy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"com.apple.iTunes.WebKit2EnableInheritURIQueryComponent": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"isAccountEnrolledInITunesMatch": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"log-PlayQueue-LastSelectedTab": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"isAccountAdmin": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"NSApplicationCrashOnExceptions": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"com.apple.iTunes.WebKit2DeveloperExtrasEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"DeviceBackupsDisabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"disableShareLibraryInfo": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"WirelessBuddyID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"Store DSID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"storefront": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"isAccountEnrolledInAppleMusic": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"log-push": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"debugAssertCategoriesEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.textInput.keyboardServices.textReplacement": DomainPrefs{
			"osxMajorVersForCloudKitSync": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"TRCKSyncMaxCountThreshold": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"_KSTRCloudKitMigratable": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"KSDidPushAllLocalRecordsOnce-2": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"osxMinorVersForCloudKitSync": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"internalMigrationPercent": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"KSCKSubscriptionProd-TextReplacements": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"seedMigrationPercent": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"seedMigrationPercent1": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"KSDidPushMigrationStatusOnce-2": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"gmMigrationPercent1": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"KSDidMigrateToCloudKitOnCloud": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"KSCloudKitMigrationDidComplete": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"kTRCKSyncCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"KSCKDidSetupRecordZoneProd-TextReplacements": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"iOSMajorVersForCloudKitSync": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"kDidMigrateToUUIDRecordName": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"internalMigrationPercent1": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"kDidInsertSampleShortcutForPeer": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"iOSMinorVersForCloudKitSync": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"gmMigrationPercent": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"userDidFallInMigrationAllowBatch": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"KSSampleShortcutWasImported_CK": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"KSDidPullLegacyEntriesFromPeers": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"TRCKSyncCountHalflifeInHours": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"KSLastKnownUserID-TextReplacements": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"osxMinorSubversionForCloudKitSync": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"KSLSShouldUpdateCache": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.diagnosticextensionsd": DomainPrefs{
			"directoriesCleanupDone": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.Messages": DomainPrefs{
			"hasPerformedNewDeviceBringUpSync": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"hasTriedToEnableCKAndSyncAfterFirstImagentConnection": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.NewDeviceOutreach": DomainPrefs{
			"InitialOutreachActivityScheduled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"CachedWarrantyValidityDuration": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"PostUpgradeActivityCompleted": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"PostUpgradeOSVersionKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"CachedWarrantyLocale": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"CachedWarrantyVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.icloud.searchpartyuseragent": DomainPrefs{
			"PersistenceLayerVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastOSLaunchVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"lastLaunchBootSessionUUID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"lastLaunchLocale": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"FMIPStateManager.fmipState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.cloudpaird": DomainPrefs{
			"MagicCloudPairingAccessorySubscriptionManateeID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"MagicCloudPairingManateeUpgradedAccount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"MagicCloudPairingManateeUpgraded": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"MagicCloudPairingProtectedAccessorySubscription": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"UploadedHSA2KeysForLocalDevice": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"MagicCloudPairingAccessorySubscriptionID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"MagicCloudPairingMasterEncryptionPreferences": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"SecuredAccessoryZoneSubscription": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"MagicCloudPairingMasterSubscriptionManateeID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"MagicCloudPairingMasterSubscriptionID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"SecuredZoneSubscription": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"MagicCloudPairingAccessoryEncryptionPreferences": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.syncserver": DomainPrefs{
			"SyncServicesResetWorldRunOnce": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"RunCompletelyDisabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.screencaptureui": DomainPrefs{
			"com.apple.AnnotationKit.userDefaultsVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"com.apple.AnnotationKit.arrowHeadStyle": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.AnnotationKit.strokeIsDashed": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.AnnotationKit.hasShadow": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.AnnotationKit.strokeWidth": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.AnnotationKit.highlightStyle": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"toolbarOrigin": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"com.apple.AnnotationKit.brushStyle": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.desktopservices": DomainPrefs{
			"DSDontWriteNetworkStores": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.security.pboxd": DomainPrefs{
			"ILMediaBrowserMediaType": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ILMediaBrowserSelectedGroupIndex1": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ILMediaBrowserSelectedGroup1": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.keychainaccess": DomainPrefs{
			"Keychain Item List Sorting": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Keychain Item List Sort Descending": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Last Selected Category": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"VPNSSItemsChecked": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"Item Preview Closed": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"Last Selected Keychain": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"Keychains List Closed": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.imessage": DomainPrefs{
			"PreviewTranscodingWasInitializedToDefaultValue": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"PreviewTranscodingQualityOnWiFiWasInitializedToDefaultValue": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.Safari.SafeBrowsing": DomainPrefs{
			"HasMigratedSafeBrowsingEnabledDefaults": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.AppleMediaServices": DomainPrefs{
			"AMSIncludeFullResponseInHARLogging": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"AMSShowSandboxAccountUI": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"AMSDeviceBiometricsState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"AMSUserDefaultsincludeFullResponseInHARLogging": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"AMSMigratedToNewAccountStore": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"AMSUserDefaultsIncludeFullRequestInHARLogging": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"AMSLogHARData": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"AMSLastMigratedBuildVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"AMSDidSetUpServerDataCache": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AMSIncludeFullRequestInHARLogging": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"AMSMigratedStorageToDefaultsForNonAMSInternal": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.PhotoBooth": DomainPrefs{
			"AVDesktopPlaybackControlsControllerShowsDurationInsteadOfTimeRemainingDefaultsKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"PBLaunchedAtLeastOnceOnLion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"UserCameraUniqueIDPref": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.MobileSMSPreview": DomainPrefs{
			"IMPreviewGenerationMaxPxWidth": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"IMPreviewGenerationMinHeight": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"IMPreviewGenerationScreenScale": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"IMPreviewGenerationMinWidth": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.ActivityMonitor": DomainPrefs{
			"OpenMainWindow": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ShowCategoryAppsinLast12Hours": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"SelectedTab": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ShowCategory": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"HasSavedState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.airplay": DomainPrefs{
			"showInMenuBarIfPresent": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.messages.facetime": DomainPrefs{
			"FaceTimeTwoTimeCallthroughEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.remindd": DomainPrefs{
			"CloudKitAccountStatus": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"preferredDefaultListObjectIDUrl": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ThrottlingPolicyCurrentBatchCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"isDatabaseMigrated": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"preferredDefaultListID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"cloudKitSchemaCatchUpSyncLastSuccessBuildVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"ThrottlingPolicyCurrentLevelIndex": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"spotlightIndexVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"useExtraneousAlarmBackOffThrottleInterval": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"CloudConfigurationPath": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"lastDatabaseMigrationSystemBuildVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.ncprefs": DomainPrefs{
			"content_visibility": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"sort_order": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"version": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.itunesstored": DomainPrefs{
			"AccountsNotificationPlugin.ActiveStorefrontIdentifier": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.driver.AppleHIDMouse": DomainPrefs{
			"Button4Force": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ScrollH": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ButtonDominance": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Button4": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Button4Click": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Button3": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Button1": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Button2": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ScrollSSize": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ScrollV": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ScrollS": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.preferences.extensions.CredentialProvider": DomainPrefs{
			"userHasOrdered": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.itunescloud": DomainPrefs{
			"ICDefaultsKeyLastActiveAccountDSID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.stockholm": DomainPrefs{
			"RemoteAdminV2": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"PKWalletShouldAutomaticallyRegisterKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.Siri": DomainPrefs{
			"VoiceTriggerUserEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"StatusMenuVisible": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.sharingd": DomainPrefs{
			"AirDropRandomHashUUIDKey4": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"OneTimeAirDropReset2": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"OneTimeAirDropReset": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"HashManager-StoredDatabaseVersionKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"AirDropRandomHashUUIDKey2": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"DiscoverableMode": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"AutoUnlockPresentedWiFiError": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AutoUnlockPresentedBluetoothError": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AirDropRandomHashUUIDKey3": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"AutoUnlockWatchCurrentlyInList": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AutoUnlockWatchExistedInUnlockList": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AirDropRandomHashUUIDKey1": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.AppStoreComponents": DomainPrefs{
			"ASCLocaleID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.DFSLoginPlugin": DomainPrefs{
			"WindowLeft": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"WindowTop": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.iCal": DomainPrefs{
			"BirthdayEventsGenerationLastLocale": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"last calendar view description": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"BirthdayEventsGenerationVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"icaluuid": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"CalendarSidebarShown": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"AvailabilityShowTwentyFourHours": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"CalForceTruthFileRestoreHashKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"last selected calendar list item": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"CalDefaultCalendar": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"CalAgentNS_Preference_DefaultReminderCalendar": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"display birthdays calendar": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NSEventConcurrentProcessingEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"BirthdayEventsGenerationAttemptsToResetKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"TimeZone support enabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Show Week Numbers": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"iCal version": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"first shown minute of day": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"CalendarSidebarView": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"RemindersLastMigratedSystemVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastReminderMigrationCleanupVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AllDayAreaHeight": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"lastViewsTimeZone": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"privacyPaneHasBeenAcknowledgedVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"CalendarSidebarWidth": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"last minute of day time range": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"first minute of day time range": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"CalendarListMiniMonthVisibleMonths": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"IncludeDebugMenu": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"WR_DONT_ASK_FOR_DEFAULT": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"kLastABCReportTimeKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NSDontMakeMainWindowKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.assistant": DomainPrefs{
			"Experiment Identifierinvocation_feedback_experiment": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"Has Set Up MultiUser Shared Record Subscription": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"Initial Interstitial Delay for VoiceTrigger One-Shot without Audio Ducking": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"Initial Interstitial Delay - stark": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"MultiUser Shared Data Needs Sync": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"Initial Interstitial Delay": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"Myriad Device Delay": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"Myriad Device Class": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"Myriad Device Adjust": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"Has Set Up Account Status Subscription": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"Voice Trigger Needs Sync": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"Myriad Device Trump Delay": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"Experiment Identifier": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"Flush Session Tickets Cache": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"Server Has Provisioned Myriad": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"Myriad VTEndtimeDistanceThreshold": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"Has Set Up Key Value Subscription": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"Borealis Education Header Display Count": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"Experiment Identifiersiri_vox_tap_to_siri_behavior_experiment": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"Has Set Up Voice Trigger Subscription": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"Last Known Analytics Store State": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"Experiment Identifiersiri_vox_sounds_experiment": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"Manual Endpointing Threshold": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.dt.Xcode": DomainPrefs{
			"IBPreferencesMigrated": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"IBAppliesAutoResizingRulesWhileResizing": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"XCFontAndColorCurrentTheme": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"IDEKeyBindingCurrentPreferenceSet": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"IDESourceControlPreferencesVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"DVTFontAndColorLastUpdatedToolsVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"DVTDownloadableAutomaticUpdate": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"IDEUserWantsToEnableDeveloperSystemPolicyMode": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.AMPLibraryAgent": DomainPrefs{
			"iTunes-media-folder-url": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"debugAssertCategoriesVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"have-shown-cloud-UI": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"storeSupportsUPP": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"updateLevel": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"ITUserPrefsMigrated": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"books-migrated": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"storeSupportsPreviousPurchases": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"persistentUserID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"storeSupportsCloudPurchases": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"imported-media-domains": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"books-persistent-id": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"storeMediaTypeFlags": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"podcasts-migrated": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"storeSupportsPasswordSettings": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"imported-media-domains-modification-date": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"show-home-sharing-turned-on-in-iTunes-warning": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"devices-persistent-id": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"debugAssertCategoriesEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"podcasts-persistent-id": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.wifi.keychain-format": DomainPrefs{
			"Version": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.facetime.bag": DomainPrefs{
			"URL": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"CacheTime": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.TelephonyUtilities": DomainPrefs{
			"GondolaLastAccountsChangedState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"GondolaGeneratedIdentifier": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"registeredProvidersVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"CachedVCCaps": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"relayCallingDisabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"GondolaLatestVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"GondolaSyncedVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"FaceTimePhotosEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.findmy.fmfcore.notbackedup": DomainPrefs{
			"frontMostWindow": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"publicAPSToken": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"FMFGarbageCollectorIdentityKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"FMFLimitedPrecisionPrefKey.limitedPrecision": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"windowVisible": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.AppStore": DomainPrefs{
			"__uniquePageGroupID-9.WebKit2UserInterfaceDirectionPolicy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"__uniquePageGroupID-12.WebKit2DeveloperExtrasEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"mostRecentTabIdentifier": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AutoPlayVideoSetting": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.AppStore.WebKit2UserInterfaceDirectionPolicy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"appStoreBadgeCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"__uniquePageGroupID-1.WebKit2ShowsToolTipOverTruncatedText": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"UserSetAutoPlayVideoSetting": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"lastBootstrapTimeZone": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"__uniquePageGroupID-9.WebKit2SuppressesIncrementalRendering": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ASAcknowledgedOnboardingVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"__uniquePageGroupID-12.WebKit2ShowsToolTipOverTruncatedText": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"__uniquePageGroupID-9.WebKit2ShowsToolTipOverTruncatedText": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"com.apple.AppStore.WebKit2SuppressesIncrementalRendering": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"__uniquePageGroupID-1.WebKit2DeveloperExtrasEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"__uniquePageGroupID-9.WebKit2DeveloperExtrasEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AVDesktopPlaybackControlsControllerShowsDurationInsteadOfTimeRemainingDefaultsKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"__uniquePageGroupID-12.WebKit2UserInterfaceDirectionPolicy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"__uniquePageGroupID-12.WebKit2SuppressesIncrementalRendering": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"com.apple.AppStore.WebKit2ShowsToolTipOverTruncatedText": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"com.apple.AppStore.WebKit2DeveloperExtrasEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"__uniquePageGroupID-1.WebKit2UserInterfaceDirectionPolicy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"__uniquePageGroupID-1.WebKit2SuppressesIncrementalRendering": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.AdPlatforms": DomainPrefs{
			"AppStorePAAvailable": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"LatestPAVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"acknowledgedPersonalizedAdsVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.keyboard": DomainPrefs{
			"KeyboardWordOrSentenceTrackingForPFL": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.assistant.support": DomainPrefs{
			"Siri Data Sharing Opt-In Status": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Assistant Enabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Dictation Enabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.preferences.extensions.QuickLook": DomainPrefs{
			"userHasOrdered": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.internetconfigspec": DomainPrefs{
			"WWWHomePage": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.icloud.fmfd": DomainPrefs{
			"numberOfFriendsFollowersKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"kFMFDStoredDataVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"storedDSID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"lazyInitTimeSecsStoredKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"kFMFDlastLoggedInPrsId": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.newscore": DomainPrefs{
			"notificationEnableAssetPrefetching": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.news.notification_receipt_event_endpoint": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"news_url_resolution_subscription_status": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"com.apple.news.default_event_endpoint": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"notificationAssetPrefetchingRequiresWatch": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"notificationArticleDiversificationSimilarityExpectationStart": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"force_refresh_user_segmentation": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"notificationArticleDiversificationUniquePublisherExpectationYIntercept": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"notificationArticleDiversificationUniquePublisherExpectationSlope": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"news_carplay_is_enabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.newsd.download.hasUnfinishedWork": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"notificationArticleDiversificationSimilarityExpectationEnd": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"use_parsec_results_for_widget": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.AddressBook": DomainPrefs{
			"ABBirthDayVisible": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ABMetaDataChangeCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ABTextSizeIncrement": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ABVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"ABLastImportShown": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ABDefaultSourceID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NSPreferencesSelectedIndex": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AB21vCardEncoding": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NSPreferencesContentSize": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ABPrivateVCardFieldsEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.SpeechRecognitionCore": DomainPrefs{
			"AllowAudioDucking": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.commerce": DomainPrefs{
			"LastUpdateNotificationOSMajorVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"PurchasesInflight": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.networkresolutiond": DomainPrefs{
			"_networkDevices": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.TV": DomainPrefs{
			"automaticallyDownloadArtwork": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"userWantsPlaybackNotifications": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"gridViewSearchDoesNotSwitchView": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"persistentUserID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"miniplayerUserSetHeight": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"storeMediaTypeFlags": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"dontAskForAllPlaylistItemRemoval": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"checkedHLSKeysTime": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"storeSupportsCloudPurchases": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"debugAssertCategoriesVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"com.apple.TV.WebKit2EnableInheritURIQueryComponent": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"cddbPrefsOK": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"com.apple.TV.WebKit2SuppressesIncrementalRendering": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"miniplayerVPos": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"controllableInterfaceGUID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"automaticallyDeleteVideoAssetsAfterWatching": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"videoWindowVPos": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"miniplayerHPos": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"playbackIsFullscreen": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"storeSupportsPasswordSettings": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"com.apple.TV.WebKit2UserInterfaceDirectionPolicy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"haveAskedToCheckForAvailableDownloads": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"videoWindowHPos": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"miniplayerSnapMode": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"storeSupportsUPP": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"showWelcomeScreenState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"videoWindowVSize": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"updateLevel": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"ITUserPrefsMigrated": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"miniplayerWidthInDesignCoords": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"storeSupportsPreviousPurchases": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"tabViewMode": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"preloadFilesIntoMemory": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"userMaxConcurrentDownloads": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NSApplicationCrashOnExceptions": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"videoWindowHSize": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"checkForAvailableDownloads": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"miniPlayerQueueVisible": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"debugAssertCategoriesEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"firstLaunchShowWelcomeScreenState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"com.apple.TV.WebKit2ShowsToolTipOverTruncatedText": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"musicVideoPlaybackLocation": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"miniPlayerLargeArtVisible": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"userHasActivatedFullScreenVisualizer": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"com.apple.TV.WebKit2AsynchronousSpellCheckingEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"com.apple.TV.WebKit2ShouldPrintBackgrounds": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.preference.general": DomainPrefs{
			"version": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.menuextra.battery": DomainPrefs{
			"ShowPercent": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.imdpersistence.IMDPersistenceAgent": DomainPrefs{
			"LogAllIOErrors": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.SafariServices": DomainPrefs{
			"SearchProviderIdentifierMigratedToSystemPreference": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.siri.shortcuts": DomainPrefs{
			"WFServicesShortcutsMigrated": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"WFDidUnconflictShortcuts": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"LegacyShortcutsZoneSubscriptionUnsubscribed": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"WFDefaultShortcutsVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"WFLastSyncedFlagsHash": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.SystemProfiler": DomainPrefs{
			"PrefsVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"SPLastDocumentsSize": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.remindd.babysitter": DomainPrefs{
			"LastSystemVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.internal.ck": DomainPrefs{
			"DictationOnDeviceSamplingDisabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"DictationSamplingRates": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"DefaultCacheKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"DefaultWarmupScriptsExtension": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"WarmupScriptIdentifiers": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"HasSetUpRecordZoneSubscription": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ByteCodeCacheEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"DefaultBootTimeUpdateScripts": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"WarmupModularScriptIdentifiers": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"DisableFBFForUEI": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.internetconfigpriv": DomainPrefs{
			"WWWHomePage": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.dock": DomainPrefs{
			"wvous-br-corner": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"showAppExposeGestureEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"autohide": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"tilesize": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"no-bouncing": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"region": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"showhidden": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"launchanim": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"showMissionControlGestureEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"version": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"wvous-tr-modifier": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"magnification": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"largesize": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"mod-count": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"trash-full": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"minimize-to-application": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"loc": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"expose-animation-duration": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"wvous-tr-corner": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"autohide-delay": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.AdLib": DomainPrefs{
			"adprivacydMaxSegmentSendInterval": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"personalizedAdsMigrated": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"partiality-segment": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"CKDPIDSyncState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"forceLimitAdTracking": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"adprivacydSegmentInterval": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"allowIdentifierForAdvertising": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"allowApplePersonalizedAdvertising": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.amp.mediasharingd": DomainPrefs{
			"public-sharing-enabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"home-sharing-sequence-id": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"home-sharing-group-id": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"home-sharing-computer-id": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"photo-sharing-enabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"shared-library-id": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"shared-library-name": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"public-sharing-is-protected": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"shared-library-machine-id": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"home-sharing-enabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.systempreferences": DomainPrefs{
			"com.apple.preferences.sharing.selectedservice": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ThirdPartyCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"SoundTab": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"com.apple.dtandsspref.lastselectedtab": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NSFullScreenMenuItemEverywhere": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"trackpad.lastselectedtab": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"ShowAllMode": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ICQ_iCloudFirstRunNotificationShown": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"DisableAutoLoginButtonIsHidden": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"com.apple.preferences.accounts.outline.usersparent": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"mouse.lastselectedtab": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NSQuitAlwaysKeepsWindows": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.SecurityPref.Privacy.LastSourceSelected": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"SecurityPrefTab": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"com.apple.preference.keyboard.lastselectedtab": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"com.apple.datetimepref.lastselectedtab": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.UIKit": DomainPrefs{
			"hasAccessibilityBeenMigrated": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.SafariBookmarksSyncAgent": DomainPrefs{
			"CloudBookmarksSupplementalClientIdentifier": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"NewestLaunchedSafariBookmarksSyncAgentVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.VoiceOverTraining": DomainPrefs{
			"doNotShowSplashScreen": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.newsd": DomainPrefs{
			"FCAppConfigurationBundleShortVersionKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.itunescloud.daemon": DomainPrefs{
			"ICDDefaultsKeyKnownActiveAccountDSID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.Image_Capture": DomainPrefs{
			"IK_lastUsedDeviceUUID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"IK_Camera_selectedTag": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"IK_lastUsedDeviceIsRemote": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"IK_Scanner_downloadURL": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"IKC_sort_ascending": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.imagekit.cameraviewmode": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"IK_prefsVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"IK_Camera_selectedPathType": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"IKC_sort_key": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"IK_Accessory_selectedTag": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"IK_Scanner_PostProcessApplication": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"IK_Camera_preferPostPocessingApp": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"IK_lastUsedDeviceName": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.preferences.extensions.SharedLinks": DomainPrefs{
			"userHasOrdered": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.driver.AppleBluetoothMultitouch.mouse": DomainPrefs{
			"MouseOneFingerDoubleTapGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"MouseTwoFingerHorizSwipeGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"MouseVerticalScroll": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"UserPreferences": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"MouseHorizontalScroll": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"MouseButtonDivision": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"MouseMomentumScroll": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"MouseButtonMode": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"MouseTwoFingerDoubleTapGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.dock.external.extra.x86_64": DomainPrefs{
			"ShowSolidarity": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.commcenter.data": DomainPrefs{
			"pw_ver": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.imagecapture": DomainPrefs{
			"loggingLevel": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.notificationcenterui": DomainPrefs{
			"bannerTime": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.pipagent": DomainPrefs{
			"Size": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.AMPDevicesAgent": DomainPrefs{
			"debugAssertCategoriesEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"persistentUserID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"dontAutomaticallySyncIPods": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"updateLevel": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"storeMediaTypeFlags": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"AutomaticDeviceBackupsDisabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"storeSupportsCloudPurchases": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"debugAssertCategoriesVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"ITUserPrefsMigrated": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"userMaxConcurrentDownloads": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.MobileSMS.CKDNDList": DomainPrefs{
			"CatalystDNDMigrationVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"CKDNDMigrationKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.protectedcloudstorage.protectedcloudkeysyncing": DomainPrefs{
			"registrySyncIdentifier": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.VoiceOver4.local": DomainPrefs{
			"SCRCUserDefaultsAllowAirPlay": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"SCRCUserDefaultsUnplannedShutdownCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"SCRCUserDefaultsPlannedShutdownSuccessful": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.suggestions": DomainPrefs{
			"SuggestionsAllowGeocode": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"spToLearnMigrationPerformed": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"findToShowMigrationPerformed": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.NetworkUtility": DomainPrefs{
			"NUPingCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NUTracerouteAddress": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NUPortScanEnd": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NUSelectedTabItem": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NULookupAddress": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NUFingerPerson": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NUPortScanAddress": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NUPingChoice": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NUPingAddress": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NUNetstatChoice": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NUPortScanStart": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NUWhoisAddress": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NUWhoisSelectedServer": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"NUPortScanRange": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.siri.VoiceShortcuts": DomainPrefs{
			"VCLSDatabaseUUIDKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"VCLSDataSequenceKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.Safari": DomainPrefs{
			"ResetCloudHistory": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"WebKitHistoryItemLimit": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"WebKitInitialTimedLayoutDelay": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ShowFullURLInSmartSearchField": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"com.apple.Safari.ContentPageGroupIdentifier.WebKit2DeveloperExtrasEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"cloudBookmarksMigrationEligibilityDataInvalidated": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"UniversalSearchEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"WebKitDeveloperExtrasEnabledPreferenceKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"SuppressSearchSuggestions": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"WebKitHistoryAgeInDaysLimit": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"IncludeDebugMenu": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"IncludeDevelopMenu": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"IncludeInternalDebugMenu": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.accounts": DomainPrefs{
			"ACOneTimeSMTPFixAccountSync": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"ACOneTimeLDAPFixAccountSync": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.visualintelligence": DomainPrefs{
			"bypassAvailability": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"enableBooksDomain": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"enableLandmarkDomain": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"enableArtDomain": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"firstTimeExperience": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"enableCoarseClassification": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"initialized": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"enablePetsDomain": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"enableScreenshots": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"debugUI": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"enableSafariApp": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"sendLocationInfo": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"enableAlbumsDomain": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"enableQuickLook": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"enableNatureDomain": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"enablePhotosApp": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"sendOCRText": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.accountsd": DomainPrefs{
			"com.apple.mail.searchableIndex.lastProcessedAttachmentIDKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"LastSystemVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.fileproviderd": DomainPrefs{
			"com.apple.fileproviderd.did-drop-daemon-corespotlight-index": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.screencapture": DomainPrefs{
			"location": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"style": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"captureDelay": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"last-selection-display": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"video": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.SoftwareUpdate": DomainPrefs{
			"AutoUpdateMajorOSVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.Preferences": DomainPrefs{
			"UserDictionaryImportedSinceMaintenance": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"UserDictionaryLocalPeerIdentityCurrent": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.sociallayerd.CloudKit.ckwriter": DomainPrefs{
			"accountLastKnownUserRecordID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"accountAvailable": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"version": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"writerDone": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.classroom": DomainPrefs{
			"allowClassroomLockDevice-00000000-0000-0000-0000-000000000000:0": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"allowClassroomOpenApp-00000000-0000-0000-0000-000000000000:0-ask": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"allowClassroomScreenObservation-00000000-0000-0000-0000-000000000000:0": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"allowClassroomOpenURL-ask": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"allowClassroomScreenObservation-00000000-0000-0000-0000-000000000000:0-ask": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"allowClassroomOpenURL": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"allowClassroomLockDevice-00000000-0000-0000-0000-000000000000:0-ask": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"allowClassroomOpenURL-00000000-0000-0000-0000-000000000000:0-ask": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"allowClassroomOpenApp-00000000-0000-0000-0000-000000000000:0": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"forceUnpromptedRemoteScreenObservation-00000000-0000-0000-0000-000000000000:0-ask": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"allowClassroomOpenURL-00000000-0000-0000-0000-000000000000:0": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"forceUnpromptedRemoteScreenObservation-00000000-0000-0000-0000-000000000000:0": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"allowClassroomOpenApp": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"allowClassroomOpenApp-ask": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.CloudKit": DomainPrefs{
			"AccountInfoValidationCounter": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.findmy.fmipcore.notbackedup": DomainPrefs{
			"publicAPSToken": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"itemLearnMoreURL": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"frontMostWindow": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"windowVisible": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"deviceImageVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"FMIPGarbageCollectorIdentityKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"FMIPLimitedPrecisionPrefKey.limitedPrecision": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.tipsd": DomainPrefs{
			"TPSWelcomeNotificationViewedVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"TPSTipsAppInstalled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"TPSOfflineLastProcessedRemoteContentIdentifier": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"DeliveryInfoVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"TPSLastMajorVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"ConsecutiveNotificationsCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"TPSWelcomeNotificationReminderState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.CoreGraphics": DomainPrefs{
			"DisplayUseForcedGray": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"DisplayUseInvertedPolarity": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.DirectoryUtility": DomainPrefs{
			"FirstLaunch": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.AvatarUI.Staryu": DomainPrefs{
			"AVTAvatarHasDisplayedAnimojiSplashScreen": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AVTAvatarHasDisplayedCameraEffectsSplashScreen": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AVTAvatarHasDisplayedPaddleView": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AVTAvatarHasDisplayedSplashScreen": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AVTAvatarUILastCacheVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.networkConnect": DomainPrefs{
			"VPNShowTime": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.UserAccountUpdater": DomainPrefs{
			"com.apple.HidLibraryFolderAlready": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.madrid": DomainPrefs{
			"IMCloudKitStartingInitialSync": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"CloudKitIsSyncing": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"CloudKitIsRemovedFromBackup": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"CloudKitIsEligibleForTruthZone": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"CloudKitEligibleToToggleMiCSwitch": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"IMCloudKitSyncControllerSyncTypeKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"CloudKitCheckedMiCSwitchEligibilityOnImagentLaunch": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"createdChatZone": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"hasCompletedInitialCKSync": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"CloudKitSyncingEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"IMCloudKitSyncControllerSyncRecordTypeKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"initialSyncRecordHasBeenWritten": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"IMCloudKitSyncControllerSyncStateKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"IMCloudKitStartingPeriodicSync": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"IMCloudKitSyncPaused": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"RequestPriorityRamp": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"IMCloudKitAccountStatusKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"IMCloudKitStartingEnabledSettingChange": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"IMCKFinishedFetchingAttachmentsFromCloudKit": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"CKMOCAccountsMatch": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"IMCloudKitStartingDisableDevices": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"oneTimeMOCAccountCheckV2": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"enableCKSyncingV2": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.rapport": DomainPrefs{
			"familySyncedName": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.menuextra.clock": DomainPrefs{
			"DateFormat": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.imessage.bag": DomainPrefs{
			"URL": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"CacheTime": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.sms": DomainPrefs{
			"hasBeenApprovedForSMSRelay": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.Automator": DomainPrefs{
			"NSSplitView AMDocumentMinor Expanded Position": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"NSSplitView AMLibraryActionsMajor Expanded Position": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AMLogTabViewSelectedIndex": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.weather.internal": DomainPrefs{
			"userId": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.FaceTime": DomainPrefs{
			"DidMigrateToSixteenByNineAspectRatio": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"kPromptEnableReadRecipts": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"CustomRingtone": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"PhoneNumberUpgradeShown": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"PreferredVideoDeviceUID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.AssistiveControl.virtualKeyboard": DomainPrefs{
			"SCLaunchedAsSlave": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"PanelPosition": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"DesiredPanelWindowPosition": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.Accessibility": DomainPrefs{
			"InvertColorsEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"GrayscaleDisplay": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"AXSClassicInvertColorsPreference": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"KeyRepeatInterval": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"CommandAndControlEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ApplicationAccessibilityEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"DarkenSystemColors": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"FullKeyboardAccessFocusRingEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"GenericAccessibilityClientEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"AccessibilityEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"KeyRepeatEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"KeyRepeatDelay": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ZoomTouchEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"DifferentiateWithoutColor": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"FullKeyboardAccessEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"BrailleInputDeviceConnected": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"AutomationEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ReduceMotionEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"SpeakThisEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"EnhancedBackgroundContrastEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"VoiceOverTouchEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.ipTelephony": DomainPrefs{
			"ImsLoggingEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.ColorSyncCalibrator": DomainPrefs{
			"722524374X-shareMode": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"722524374X-targetTemp": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"722524374X-name": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"722524374X-expertMode": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.StatusKitAgent": DomainPrefs{
			"reauthCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.AppleMultitouchTrackpad": DomainPrefs{
			"TrackpadThreeFingerDrag": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadPinch": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"USBMouseStopsTrackpad": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadFourFingerVertSwipeGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ActuationStrength": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ActuateDetents": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"FirstClickThreshold": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"SecondClickThreshold": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadTwoFingerDoubleTapGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadHorizScroll": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadMomentumScroll": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadFourFingerPinchGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadRotate": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Clicking": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadThreeFingerTapGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadThreeFingerHorizSwipeGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadScroll": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"DragLock": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadFiveFingerPinchGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"ForceSuppressed": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"TrackpadThreeFingerVertSwipeGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadTwoFingerFromRightEdgeSwipeGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"UserPreferences": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"version": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"HIDScrollZoomModifierMask": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadHandResting": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadRightClick": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Dragging": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadFourFingerHorizSwipeGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"TrackpadCornerSecondaryClick": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.ImageCaptureService": DomainPrefs{
			"IK_Accessory_selectedTag": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.screensaver": DomainPrefs{
			"tokenRemovalAction": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.AppleMultitouchMouse": DomainPrefs{
			"MouseOneFingerDoubleTapGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"MouseTwoFingerHorizSwipeGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"MouseVerticalScroll": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"UserPreferences": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"MouseHorizontalScroll": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"MouseButtonDivision": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"MouseMomentumScroll": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"version": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"MouseButtonMode": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"MouseTwoFingerDoubleTapGesture": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.universalaccessAuthWarning": DomainPrefs{
			"2::/Applications/Logi Options.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::com.logitech.manager.daemon": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Microsoft Teams (work or school).app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::com.lightpillar.Mosaic-setapp": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"1::/Applications/CloudBerry Backup.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::com.getcleanshot.app-setapp": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::/Applications/Setapp/Mosaic.app/Contents/MacOS/Mosaic": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"/Applications/Clipy.app/Contents/MacOS/Clipy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/xScope.app/Contents/MacOS/xScope": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Microsoft Teams.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"1::/Users/mikeschinkel/Library/Application Support/JetBrains/Toolbox/apps/Goland/ch-0/232.8660.185/GoLand.app/Contents/MacOS/goland": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"1::/Applications": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::com.microsoft.teams2": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Discord.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Microsoft Edge.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::/Applications/Logi Options.app/Contents/Support/LogiMgrDaemon.app/Contents/MacOS/LogiMgrDaemon": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Skitch.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/ScreenFlow.app/Contents/Library/LoginItems/ScreenFlowHelper.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"/Applications/Alfred 4.app/Contents/MacOS/Alfred": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Toast 19 Pro.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"1::com.vmware.fusion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::com.lastpass.lastpassmacdesktop": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"com.clipy-app.Clipy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::com.google.Chrome": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"1::/Users/mikeschinkel/Library/Application Support/JetBrains/Toolbox/apps/Goland/ch-0/232.8660.185/GoLand.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"1::/Applications/Microsoft Teams.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/ScreenFlow.app/Contents/Library/LoginItems/ScreenFlowHelper.app/Contents/MacOS/ScreenFlowHelper": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Microsoft Teams (work or school).app/Contents/MacOS/MSTeams": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/RingCentral.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::com.amazon.Amazon-Chime": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"com.mutedeck.mac": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Parallels Toolbox.app/Contents/Applications/Hide Menu Icons.app/Contents/MacOS/Hide Menu Icons": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::/Applications/Snappy.app/Contents/MacOS/Snappy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::me.waydabber.BetterDummy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/GoToMeeting.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Google Chrome.app/Contents/MacOS/Google Chrome": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::WSG985FR47.net.telestream.screenflowhelper": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Slack.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"/Applications/Alfred 4.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"1::/Applications/BetterDisplay.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::com.logi.cp-dev-mgr": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"com.runningwithcrayons.Alfred": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Xnapper.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/xScope.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::com.microsoft.edgemac": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"/Applications/zoom.us.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"/System/Applications/Automator.app/Contents/MacOS/Automator": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"/Applications/Clipy.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/BetterDisplay.app/Contents/MacOS/BetterDummy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Snip.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::ro.nextwave.Snappy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::com.parallels.toolbox.HideMenuIcons": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Google Chrome.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"1::/Applications/VMware Fusion-8.5.10.app/Contents/MacOS/VMware Fusion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::com.ringcentral.glip": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Acorn.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"/Applications/MuteDeck/MuteDeck.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::/Library/Application Support/Logitech.localized/LogiOptionsPlus/logioptionsplus_agent.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Timing.app/Contents/Library/LoginItems/TimingHelper.app/Contents/MacOS/TimingHelper": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::/Applications/Logi Options.app/Contents/Support/LogiMgrDaemon.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::/Applications/Setapp/Mosaic.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"1::me.waydabber.BetterDummy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::/Applications/LastPass.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::info.eurocomp.TimingHelper": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::com.amazon.Amazon-Chime": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"1::/Applications/BetterDisplay.app/Contents/MacOS/BetterDummy": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"1::/Applications/iTerm.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"1::/Applications/OBS.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"1::/Applications/HP RGS Receiver.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::J8RPQ294UB.com.skitch.SkitchHelper": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::com.logitech.Logi-Options": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::/Applications/Logi Options.app/Contents/MacOS/Logi Options": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"us.zoom.xos": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"1::/Applications/Slack.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::/Applications/Snappy.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"1::com.jetbrains.goland": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::info.eurocomp.Timing2": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Xnapper.app/Contents/MacOS/Xnapper": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/RingCentral.app/Contents/MacOS/RingCentral": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"/Applications/zoom.us.app/Contents/MacOS/zoom.us": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/BetterDisplay.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/zoom.us.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"com.apple.Automator": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Skitch.app/Contents/Library/LoginItems/J8RPQ294UB.com.skitch.SkitchHelper.app/Contents/MacOS/J8RPQ294UB.com.skitch.SkitchHelper": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Amazon Chime.app/Contents/MacOS/Amazon Chime": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Setapp/CleanShot X.app/Contents/MacOS/CleanShot X Setapp": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/ScreenFlow.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::/Applications/Amazon Chime.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"1::/Applications/VirtualBox.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"/Applications/MuteDeck/MuteDeck.app/Contents/MacOS/MuteDeck": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Timing.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::/Applications/Amazon Chime.app/Contents/MacOS/Amazon Chime": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Timing.app/Contents/MacOS/Timing": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Amazon Chime.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Skitch.app/Contents/Library/LoginItems/J8RPQ294UB.com.skitch.SkitchHelper.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::com.iconfactory.xScope": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::/Applications/LastPass.app/Contents/MacOS/LastPass": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"/System/Applications/Automator.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"1::/Applications/VMware Fusion-8.5.10.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Loom.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/var/tmp/LogMeIn Rescue - 4C61/Support-LogMeInRescue.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Setapp/CleanShot X.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::com.devuap.beautyshotapp": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Timing.app/Contents/Library/LoginItems/TimingHelper.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/var/tmp/LogMeIn Rescue - 4C61": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"3::/Applications/Parallels Toolbox.app/Contents/Applications/Hide Menu Icons.app": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"2::/Library/Application Support/Logitech.localized/LogiOptionsPlus/logioptionsplus_agent.app/Contents/MacOS/logioptionsplus_agent": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.findmy": DomainPrefs{
			"restoreState": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"coarseServiceAcknowledged_v1.0": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"onboarding_v2.0": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.FolderActionsDispatcher": DomainPrefs{
			"folderActionsEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.SocialLayer": DomainPrefs{
			"SharedWithYouEnabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.security.KCN": DomainPrefs{
			"absentCircleWithNoReason": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
			"lastCircleStatus": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"pendingApplicationReminderInterval": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.GEO": DomainPrefs{
			"DefaultsSanitizedVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.SharedWebCredentials": DomainPrefs{
			"recheckVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.suggestd": DomainPrefs{
			"DeviceIdentifier": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.CallHistorySyncHelper": DomainPrefs{
			"/Users/mikeschinkel/Library/Application Support/CallHistoryTransactions/transaction.log": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"CallHistoryDeviceCount": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"/Users/mikeschinkel/Library/Application Support/CallHistoryTransactions/transactions.log": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.dashboard": DomainPrefs{
			"devmode": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"mcx-disabled": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.SystemManaged,
				},
			},
		},
		"com.apple.FontRegistry.user": DomainPrefs{
			"Version": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.imagent": DomainPrefs{
			"Setting.EnableReadReceipts": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
			"Setting.GlobalReadReceiptsVersionID": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"changeToken-syncStoreVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"IMChatVocabularyUpdaterDidPerformInitialUpdateKey": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
			"attachmentZoneChangeToken-syncStoreVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"messagesChangeToken-syncStoreVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
			"archivedMessagesChangeToken-syncStoreVersion": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.preferences.extensions.PhotoProjects": DomainPrefs{
			"userHasOrdered": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
		"com.apple.messaging.watchdog": DomainPrefs{
			"watchdogWatermark": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.RuntimeState,
				},
			},
		},
		"com.apple.preferences.extensions.PhotoEditing": DomainPrefs{
			"userHasOrdered": &Pref{
				NoDefault: true,
				Labels: macprefs.Labels{
					&macprefs.UserManaged,
				},
			},
		},
	}
}
