package prefdefaults

import (
	"github.com/mikeschinkel/prefsctl/macprefs/kvfilters"
)

//goland:noinspection SpellCheckingInspection
func montereyPrefDefaults() DomainDefaults {
	return DomainDefaults{
		"GlobalPreferences": DomainPrefs{
			"AppleActionOnDoubleClick": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"AppleAntiAliasingThreshold": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"AppleAquaColorVariant": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"AppleEnableSwipeNavigateWithScrolls": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"AppleKeyboardUIMode": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"AppleMeasurementUnits": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"AppleMetricUnits": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"AppleMiniaturizeOnDoubleClick": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"AppleShowAllExtensions": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"AppleShowScrollBars": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"AppleTemperatureUnit": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"AppleWindowTabbingMode": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.mouse.scaling": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.sound.beep.flash": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.springing.delay": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.springing.enabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.swipescrolldirection": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.trackpad.forceClick": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.trackpad.scaling": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"InitialKeyRepeat": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"KeyRepeat": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NSAutomaticCapitalizationEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NSAutomaticDashSubstitutionEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NSAutomaticPeriodSubstitutionEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NSAutomaticQuoteSubstitutionEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NSAutomaticSpellingCorrectionEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NSAutomaticTextCompletionEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NSAutomaticWindowAnimationsEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"WebAutomaticSpellingCorrectionEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"WebKitDeveloperExtras": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"AppleLanguage": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
					&kvfilters.SetupSets,
					&kvfilters.StringType,
				},
			},
			"AppleLocale": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SetupSets,
					&kvfilters.StringType,
				},
			},
			"646F6E7A_00000000_00000001_6E7A6361_696D6963:": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"AKLastIDMSEnvironment:": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"AudioQuest inc.AudioQuest DragonFly:": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"Blue Microphones Yeti Stereo Microphone:": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"Canon MF5900 Series:": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"ContextMenuGesture:": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"Fujitsu ScanSnap iX500:": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"Yubico YubiKey OTP+FIDO+CCID:": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"AppleLanguagesDidMigrate:": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"NavPanelFileListModeForOpenMode:": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"NavPanelFileListModeForSaveMode:": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NSLinguisticDataAssetsRequestLastInterval:": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"QLPanelAnimationDuration:": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.security.sosaccount": DomainPrefs{
			"SOSEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.SafariCloudHistoryPushAgent": DomainPrefs{
			"AcknowledgedPushNotifications": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"UnacknowledgedPushNotifications": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.print.custompresets.forprinter.Canon_MF5900_Series": DomainPrefs{
			"com.apple.print.lastPresetPref": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"com.apple.print.lastPresetPrefType": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.finder": DomainPrefs{
			"SidebarTagsSctionDisclosedState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"QLEnableTextSelection": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"PreviewPaneGalleryWidth": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"FK_ArrangeBy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"PreferencesWindow.LastSelection": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"FXEnableExtensionChangeWarning": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NSWindowTabbingShoudShowTabBarKey-com.apple.finder.TBrowserWindow": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"FXICloudDriveDesktop": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ShowMountedServersOnDesktop": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"SidebarShowingiCloudDesktop": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"FXToolbarUpgradedToTenNine": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"ShowRemovableMediaOnDesktop": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ShowStatusBar": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"FXSidebarUpgradedSharedSearchToTwelve": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"FXInfoWindowWidth": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"DownloadsFolderListViewSettingsVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"SidebariCloudDriveSectionDisclosedState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AppleShowAllFiles": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"FXICloudDriveDeclinedUpgrade": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"FXArrangeGroupViewBy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"RecentsArrangeGroupViewBy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"FXPreferredSearchViewStyleVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"PreviewPaneWidth": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"FXAtLeastOneScreenHasBeenInHIDPI": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"CopyProgressWindowLocation": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"SidebarShowingSignedIntoiCloud": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ViewOptionsWindow.Location": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ShowHardDrivesOnDesktop": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"FXQuickActionsConfigUpgradeLevel": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"FXPreferredViewStyle": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"QuitMenuItem": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"FXDefaultSearchScope": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"FXPreferredSearchViewStyle": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"DisableAllAnimations": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"SidebarWidth": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ShowSidebar": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TagsCloudSerialNumber": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NewWindowTarget": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ShowPathbar": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"FXSidebarUpgradedToTenFourteen": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"LastTrashState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ShowExternalHardDrivesOnDesktop": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"GoToField": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"_FXSortFoldersFirstOnDesktop": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"FK_AppCentricShowSidebar": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"FXSidebarUpgradedToTenTen": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"FXConnectToLastURL": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"FXICloudDriveEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"FXICloudDriveDocuments": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"_FXSortFoldersFirst": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"_FXShowPosixPathInTitle": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"EmptyTrashProgressWindowLocation": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"SearchRecentsSavedViewStyle": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"FXDesktopTouchBarUpgradedToTenTwelveOne": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"FXICloudDriveFirstSyncDownComplete": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"SidebarDevicesSectionDisclosedState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"MountProgressWindowLocation": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"FK_SidebarWidth": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"SearchRecentsSavedViewStyleVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"FXPreferredGroupBy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"PreviewPaneInfoExpanded": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"FXLastSearchScope": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"InspectorWindow.Location": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"SidebarPlacesSectionDisclosedState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"FXToolbarUpgradedToTenEight": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"FXICloudLoggedIn": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"FXToolbarUpgradedToTenSeven": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.messages.pinning": DomainPrefs{
			"IMPinningPinConfigMigrationKey-v2-r2": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.identityservicesd": DomainPrefs{
			"ReRegisteredForDevicesv56": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"hasRegIdentityContainer": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"ReRegisteredForDevicesv55": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"persister-migration-personal-session-token-cache-v4": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"ReRegisteredForDevices": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"DidCleanLegacyAccountPrefs": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"ImportedLegacyIMAccounts": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"hasUnregIdentityContainer": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"MigratedToNewDisabledState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"triggeredRemoteSessionVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"ReRegisteredForDevicesHash": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"ImportedLegacyIDSAccounts2": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"GDRRequestMadeForRelayRepair": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"persister-migration-com.apple.identityservices.userStore": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"ReRegisteredForDevicesv1350": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"performed-user-intent-migrate": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.Spotlight": DomainPrefs{
			"ModelName": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.AppManaged,
				},
			},
			"engagementCount-com.apple.Spotlight": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"engagementCount-com.apple.mail": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"SSMessageTracingWindowShowCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"windowHeight": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"useCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"__NSEnableTSMDocumentWindowLevel": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"SSMessageTracingWindowHideCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"engagementCount-com.apple.Spotlight.suggestions": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"SPMessageTracingWindowHideCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"showedFTE": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"SPMessageTracingWindowShowCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"userHasMovedWindow": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.DictionaryServices": DomainPrefs{
			"DCSPreferenceVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.PTPCamera": DomainPrefs{
			"FilenamePrefix": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NextFilenameCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.imtranscoding.IMTranscoderAgent": DomainPrefs{
			"kCKMediaObjectManagerDefaultsOSVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.scriptmenu": DomainPrefs{
			"ScriptMenuEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.TTY": DomainPrefs{
			"RTTSettingsVersionPreference": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"RTTContinuityRTTIsSupportedPreference": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"TUSupportsRelayCallingPreference": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"TUIsRelayCallingEnabledPreference": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.voicetrigger": DomainPrefs{
			"VoiceTrigger Enabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.DiskUtility": DomainPrefs{
			"OperationProgress ExpandedHeight": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"DUDebugMenuEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"advanced-image-options": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"SidebarShowAllDevices": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"OperationProgress DetailsVisible": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.photos.shareddefaults": DomainPrefs{
			"CPLDefaultDownload": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"iCloudPhotoLibraryLastResetWelcomePromptDBVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"downloadAndKeepOriginals": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"clearPurgeableResources": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"DiskSpaceWasLow": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"repushAssetsWithImportedByBundleIdentifier": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.AppManaged,
				},
			},
			"repushVideoAssetsMetadata": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.AppManaged,
				},
			},
		},
		"com.apple.parsecd": DomainPrefs{
			"PARDefaultsVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"PARTestSeed": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.universalaccess": DomainPrefs{
			"login": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"switchCoalescePressesDuration": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"lastNightShiftEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"switchAutoScanPanelInterval": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"closeViewZoomedIn": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"grayscale": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"switchSweepingCursorSpeed": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"sessionChange": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"switchHoldBeforeRepeatDuration": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"closeViewZoomFactorBeforeTermination": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"switchHideUITimeout": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"switchAutoScanElementInterval": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"lastNightShiftActive": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"virtualKeyboardOnOff": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"headMouseEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"increaseContrast": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"mouseDriverCursorSize": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"selectedTab": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"lastNightShiftDate": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"closeViewZoomDisplayID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"hoverTextIsAlwaysOn": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"closeViewSplitScreenRatio": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"hoverTextIsHoveringAndVisible": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AssistiveControlType": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"closeViewDesiredZoomFactor": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"slowKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"closeViewScrollWheelToggle": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"mouseDriver": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"differentiateWithoutColor": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"slowKeyDelay": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"switchMinimumPressDuration": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"closeViewZoomFactor": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"switchOnOffKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"keyboardAccessFocusRingTimeout": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"stickyKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"reduceTransparency": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"dwellHideUITimeout": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"dwellEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"contrast": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"useStickyKeysShortcutKeys": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"lastNightShiftMode": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"hoverTextEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"whiteOnBlack": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"closeViewHotkeysEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"switchFirstElementDelay": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"voiceOverOnOffKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.Terminal": DomainPrefs{
			"HasMigratedDefaults": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"SecureKeyboardEntry": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Default Window Settings": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"DefaultProfilesVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"ProfileCurrentVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"TTAppPreferences Selected Tab": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"Startup Window Settings": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.coreservices.useractivityd": DomainPrefs{
			"kLocalPasteboardBlobName": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"kRemotePasteboardBlobName": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.ColorSyncUtility": DomainPrefs{
			"CSUtilitiesContentSize": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"CSProfsUtilsGroupBy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"CSUtilitiesSelectedIdentifier": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.weather": DomainPrefs{
			"userId": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.Music": DomainPrefs{
			"dontAskForAllPlaylistItemRemoval": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"storeSupportsPasswordSettings": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"notifications-warming-last-displayed-time": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"miniplayerVPos": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"com.apple.Music.WebKit2ShouldPrintBackgrounds": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"storeSupportsPreviousPurchases": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"com.apple.Music.WebKit2AsynchronousSpellCheckingEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"miniPlayerQueueVisible": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"eqWindowHPos": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"userWantsPlaybackNotifications": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"miniplayerUserSetHeight": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"com.apple.Music.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"soundEnhancerAmount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"preloadFilesIntoMemory": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"updateLevel": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"videoWindowVSize": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"com.apple.Music.WebKit2EnableInheritURIQueryComponent": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"automaticallyDownloadArtwork": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"refreshedHLSKeysTime": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"miniplayerWidthInDesignCoords": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"userMaxConcurrentDownloads": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.Music.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"miniplayerHPos": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"videoWindowVPos": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"log-PlayQueue-LastSelectedTab": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"dontWarnDownloadCloudPurchases": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"storeSupportsUPP": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"showWelcomeScreenState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"checkedHLSKeysTime": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"storeSupportsCloudPurchases": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"gridViewSearchDoesNotSwitchView": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"eqName": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ITUserPrefsMigrated": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"cddbPrefsOK": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"eqPrefsVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"playbackIsFullscreen": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"debugAssertCategoriesVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"debugAssertCategoriesEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"firstLaunchShowWelcomeScreenState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"encoderName": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.Music.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"haveRadioState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"videoWindowHSize": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"eqWindowVPos": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"haveAskedToCheckForAvailableDownloads": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"musicVideoPlaybackLocation": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"volumeWSG": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"dontShowRestrictionsPrefs": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"persistentUserID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"NSApplicationCrashOnExceptions": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"miniPlayerLargeArtVisible": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"searchSavedTab": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"artworkDownloadDSID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"videoWindowHPos": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"controllableInterfaceGUID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"userHasActivatedFullScreenVisualizer": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"storeMediaTypeFlags": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"imported-eq-presets": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"soundEnhancerEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"miniplayerSnapMode": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"doesStoreSupportCloudMusicLibrary": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"eqEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.airport.airportutility": DomainPrefs{
			"dontPerformBaseRestartWarning": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.studentd": DomainPrefs{
			"LastDateProviderSessionToken": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"DeviceIdentifier": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.systemuiserver": DomainPrefs{
			"__NSEnableTSMDocumentWindowLevel": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.IMCoreSpotlight": DomainPrefs{
			"IMCSIdxVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"IMCSLastIndexedRowID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"IMCSNeedsIndexing": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"IMCSIndexTotalRecords": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"IMCSIdxProductVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"IMCSBypassIndexVersionCheckV2": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.bird": DomainPrefs{
			"optimize-storage": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"icloud-drive.account-migration-status.294735135": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.xpc.activity2": DomainPrefs{
			"ProductBuildVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.iPod": DomainPrefs{
			"com.apple.PreferenceSync.ExcludeAllSyncKeys": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.newscore2": DomainPrefs{
			"report_concern_user_id": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"provider_user_id": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"instance_identifier": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.cloudd": DomainPrefs{
			"com.apple.private.cloudkit.shouldUseGeneratedDeviceID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.driver.AppleBluetoothMultitouch.trackpad": DomainPrefs{
			"TrackpadFourFingerHorizSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadPinch": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"USBMouseStopsTrackpad": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadFourFingerVertSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadRotate": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadHorizScroll": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadTwoFingerDoubleTapGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"TrackpadMomentumScroll": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadFourFingerPinchGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"UserPreferences": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Clicking": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadThreeFingerTapGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"HIDScrollZoomModifierMask": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadScroll": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"DragLock": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadFiveFingerPinchGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadThreeFingerHorizSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadThreeFingerVertSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadTwoFingerFromRightEdgeSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ForceSuppressed": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"TrackpadCornerSecondaryClick": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Dragging": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadHandResting": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadRightClick": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadThreeFingerDrag": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.stocks.account": DomainPrefs{
			"deleteOnNextLaunch": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.preference.trackpad": DomainPrefs{
			"ForceClickSavedState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.SetupAssistant": DomainPrefs{
			"PreviousSystemVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"DidSeeSyncSetup2": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"SkipExpressSettingsUpdating": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"LastSeenDiagnosticsProductVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"SkipFirstLoginOptimization": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"LastSeenCloudProductVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"DidSeeCloudSetup": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"PreviousBuildVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"DidSeeApplePaySetup": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"DidSeeTrueTonePrivacy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"DidSeeScreenTime": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"MiniBuddyShouldLaunchToResumeSetup": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"DidSeeTouchIDSetup": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"NSAddServicesToContextMenus": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"DidSeeActivationLock": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"DidSeeSyncSetup": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"DidSeeAvatarSetup": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"MiniBuddyLaunchReason": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"DidSeeSiriSetup": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"DidSeeTrueTone": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"DidSeePrivacy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"LastPrivacyBundleVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"LastSeenSiriProductVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"DidSeeAppStore": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"LastPreLoginTasksPerformedVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"DidSeeAppearanceSetup": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"LastSeenSyncProductVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"LastSeeniCloudStorageServicesProductVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"LastPreLoginTasksPerformedBuild": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"MiniBuddyLaunchedPostMigration": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"DidSeeAccessibility": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"LastSeenBuddyBuildVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"DidSeeiCloudLoginForStorageServices": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.ScriptEditor2": DomainPrefs{
			"OSAStandardAdditions ChooseApplication Bounds": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.preferences.extensions.ServicesWithUI": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.MobileSMS": DomainPrefs{
			"TextSize": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"SyndicationOnboardingVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"LegacyAppSidebarPersistedWidth": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"QuickSaveHasBeenUsedBefore": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"CatalystPreferenceMigrationVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"PlaySoundsKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"BusinessChatPrivacyPageDisplayed": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"sForceUnknownFilteringCompleted": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"CKLastSelectedItemIdentifier": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"kCKMediaObjectManagerDefaultsOSVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"CatalystCustomFontMigrationVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"NotifyAboutMessagesFromUnknownContacts": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"CKMigratedAutoSpamReports26375262": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"SidebarPersistedWidth": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ReadReceiptSettingsConfirmed": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"sForceSMSSpamFilteringCompleted": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"KeepMessageForDays": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"PendingCleared": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"TextFontSize": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.preferences.extensions.ShareMenu": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.mmcs": DomainPrefs{
			"report.LastFailedCheckDate": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"report.sha256": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"report.LastSuccessfulCheckDate": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"report.TTL": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.appstored": DomainPrefs{
			"BadgeCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ArcadeSubscriptionState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ArcadePayoutDeviceID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"WelcomeNotificationExcludedFromSample": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ArcadeDeviceID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"ArcadeDeviceGUID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"LastOSBuildVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"WelcomeNotificationLastAppStoreConnectionProductVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.PubSubAgent": DomainPrefs{
			"TigerMigrationCompleted": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.commcenter.callservices": DomainPrefs{
			"last.known.icloud.id": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"associated.account": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.print.PrinterProxy": DomainPrefs{
			"IK_Scanner_downloadURL": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"IK_Scanner_selectedTag": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"IK_Scanner_PostProcessApplication": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.amsengagementd": DomainPrefs{
			"AMSMetricsIdentifierUserRecordName": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"AMSMetricsIdentifierZoneCreated": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AMSMetricsIdentifierZoneSubscriptionCreated": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.iChat": DomainPrefs{
			"SaveConversationsOnClose": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"PMPrintingExpandedStateForPrint2": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"UnifiedChatWindowControllerSelectionGUIDSet": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"DidCheckForDuplicateChats": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"HasPromptedSMSRelay": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"HasPromptediMessageFTU": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"DidRegenerateGroupID63841559": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"messageTracerSMSSent": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"messageTracerSubmitDate": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"BuddyPictureSetToGenericByUser": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"messageTracerIMessagesSent": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"PlaySoundsKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"messageTracerSMSUsed": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"messageTracerNumUpgradeOffers": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ChatWindowControllerUnifiedIsVisible": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"messageTracerMessagesSent": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"WebIconDatabaseEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"DidMarkGroupPhotosAsUnpurgeable": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"messageTracerUpgradesAccepted": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"messageTracerIMessagesReceived": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"DaemonConnectionWaitTime": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"NotifyAboutMessagesFromUnknownContacts": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"KeepMessagesVersionID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"messageTracerCharactersSent": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"messageTracerMessagesReceived": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"messageTracerUpgradesDeclined": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"messageTracerSMSReceived": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"KeepMessageForDays": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"UnifiedChatListViewControllerAutomaticallySortsChats": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"messageTracerIMessageUsed": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"DidMigratePersonCentricIDs": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"LastIMDNotificationPostedDate": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ChatWindowControllerUnifiedChatListWidth": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"LastFailedMessageIMDNotificationPostedDate": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"SODefaultTranscriptName": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"imageBrowser.disableOpenGL": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.TextEdit": DomainPrefs{
			"PlainTextEncoding": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"RichText": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"PlainTextEncodingForWrite": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.HIToolbox": DomainPrefs{
			"AppleFnUsageType": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"AppleCurrentKeyboardLayoutInputSourceID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"AppleDictationAutoEnable": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.preferences.extensions.FinderSync": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.coreservices.uiagent": DomainPrefs{
			"CSUIHaveCenteredGatekeeperProgressWindow": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"CSUIHasSafariBeenLaunched": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"CSUIRecommendSafariBackOffInterval": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"CSUILastOSVersionWhereSafariRecommendationWasMade": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.CharacterPicker": DomainPrefs{
			"version": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.CharacterPaletteIM": DomainPrefs{
			"CVFontSize": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"CVStartAsLargeWindow": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.assistant.backedup": DomainPrefs{
			"Hands Free Mode": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Session Language": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Cloud Sync Enabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Cloud Sync User ID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"MultiUser VoiceIdentification Enabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.icloud.fmfd.notbackedup": DomainPrefs{
			"kKeychainUtilMigrationVersionKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.security.cloudkeychainproxy3.keysToRegister": DomainPrefs{
			"EnsurePeerRegistration": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"DSID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"KeyAccountUUID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.mediaaccessibility.public": DomainPrefs{
			"MACaptionDisplayType": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.speech.recognition.AppleSpeechRecognition.prefs": DomainPrefs{
			"DictationIMPresentedOfflineUpgradeSuggestion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"DictationIMCommandsWindowIsOpen": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"DictationIMCorrectionCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"DictationIMUpgradedTo10_16": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"DictationIMUpgradedTo10_15": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"DictationIMPlaySoundUponRecognition": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"DictationIMLocaleIdentifier": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"DictationIMMessageTracesSinceLastReport": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"DictationIMUseOnlyOfflineDictation": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.loginwindow": DomainPrefs{
			"TALLogoutSavesState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"MiniBuddyLaunch": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"oneTimeSSMigrationComplete": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.preferences.softwareupdate": DomainPrefs{
			"LatestMajorOSSeenByUserBundleIdentifier": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.accessibility.universalAccessAuthWarn": DomainPrefs{
			"ThirdPartyCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.internetconnect": DomainPrefs{
			"ServiceID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.TextInputMenu": DomainPrefs{
			"visible": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.messages.nicknames": DomainPrefs{
			"NicknameAppleIDAndiCloudAccountMatchAndExist": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"MeCardPendingNicknamesVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"MeCardSharingVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"MeCardSharingEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"MeCardSharingImageForkedFromMeCard": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"MeCardWhitelistBlacklistNicknamesVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"NicknameInfoRequestedFromPeers": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"MeCardSharingAudience": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NicknameScrutinyDoNotHandle": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"ReuploadLocalNicknamesVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.ProblemReporter": DomainPrefs{
			"ShowComments": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ShowDetails": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.gamed": DomainPrefs{
			"natType": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"GKStoreFrontIdentifier": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"GKActiveEnvironmentKeyV2": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"GKLastPushTokenPlayerID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"GKLastPushTokenEnvironment": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"GKPushEnvironmentKeyV2": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"GKLastProtocolVersionUsedKeyV2": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"GKLoginCancelled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.Safari.SandboxBroker": DomainPrefs{
			"DidMigrateDownloadFolderToSandbox": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"ShowDevelopMenu": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"DidMigrateResourcesToSandbox": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"DidMigrateDownloadMetadataToSandboxGroupContainer": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.preferences.extensions.XcodeSourceEditor": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.AppleMediaServices.notbackedup": DomainPrefs{
			"AMSDidRetrieveDeviceOffers": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AMSDidRetrieveDeviceOffersEligibility": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.imdsmsrecordstore": DomainPrefs{
			"DeleteSequenceNumber": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.iTunes": DomainPrefs{
			"com.apple.iTunes.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"doesAccountArtistListHaveSharePermission": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"dontAutomaticallySyncIPods": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"MacBuddyStoreID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"com.apple.iTunes.WebKit2ShouldPrintBackgrounds": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"com.apple.iTunes.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"com.apple.iTunes.WebKit2AsynchronousSpellCheckingEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AutomaticDeviceBackupsDisabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"disableRadio": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Store Apple ID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"debugAssertCategoriesVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"com.apple.iTunes.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"com.apple.iTunes.WebKit2EnableInheritURIQueryComponent": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"isAccountEnrolledInITunesMatch": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"log-PlayQueue-LastSelectedTab": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"isAccountAdmin": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"NSApplicationCrashOnExceptions": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"com.apple.iTunes.WebKit2DeveloperExtrasEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"DeviceBackupsDisabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"disableShareLibraryInfo": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"WirelessBuddyID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"Store DSID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"storefront": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"isAccountEnrolledInAppleMusic": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"log-push": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"debugAssertCategoriesEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.textInput.keyboardServices.textReplacement": DomainPrefs{
			"osxMajorVersForCloudKitSync": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"TRCKSyncMaxCountThreshold": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"_KSTRCloudKitMigratable": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"KSDidPushAllLocalRecordsOnce-2": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"osxMinorVersForCloudKitSync": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"internalMigrationPercent": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"KSCKSubscriptionProd-TextReplacements": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"seedMigrationPercent": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"seedMigrationPercent1": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"KSDidPushMigrationStatusOnce-2": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"gmMigrationPercent1": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"KSDidMigrateToCloudKitOnCloud": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"KSCloudKitMigrationDidComplete": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"kTRCKSyncCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"KSCKDidSetupRecordZoneProd-TextReplacements": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"iOSMajorVersForCloudKitSync": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"kDidMigrateToUUIDRecordName": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"internalMigrationPercent1": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"kDidInsertSampleShortcutForPeer": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"iOSMinorVersForCloudKitSync": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"gmMigrationPercent": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"userDidFallInMigrationAllowBatch": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"KSSampleShortcutWasImported_CK": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"KSDidPullLegacyEntriesFromPeers": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"TRCKSyncCountHalflifeInHours": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"KSLastKnownUserID-TextReplacements": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"osxMinorSubversionForCloudKitSync": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"KSLSShouldUpdateCache": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.diagnosticextensionsd": DomainPrefs{
			"directoriesCleanupDone": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.Messages": DomainPrefs{
			"hasPerformedNewDeviceBringUpSync": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"hasTriedToEnableCKAndSyncAfterFirstImagentConnection": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.NewDeviceOutreach": DomainPrefs{
			"InitialOutreachActivityScheduled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"CachedWarrantyValidityDuration": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"PostUpgradeActivityCompleted": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"PostUpgradeOSVersionKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"CachedWarrantyLocale": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"CachedWarrantyVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.icloud.searchpartyuseragent": DomainPrefs{
			"PersistenceLayerVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"LastOSLaunchVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"lastLaunchBootSessionUUID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"lastLaunchLocale": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"FMIPStateManager.fmipState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.cloudpaird": DomainPrefs{
			"MagicCloudPairingAccessorySubscriptionManateeID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"MagicCloudPairingManateeUpgradedAccount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"MagicCloudPairingManateeUpgraded": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"MagicCloudPairingProtectedAccessorySubscription": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"UploadedHSA2KeysForLocalDevice": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"MagicCloudPairingAccessorySubscriptionID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"MagicCloudPairingMasterEncryptionPreferences": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"SecuredAccessoryZoneSubscription": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"MagicCloudPairingMasterSubscriptionManateeID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"MagicCloudPairingMasterSubscriptionID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"SecuredZoneSubscription": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"MagicCloudPairingAccessoryEncryptionPreferences": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.syncserver": DomainPrefs{
			"SyncServicesResetWorldRunOnce": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"RunCompletelyDisabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.screencaptureui": DomainPrefs{
			"com.apple.AnnotationKit.userDefaultsVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"com.apple.AnnotationKit.arrowHeadStyle": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.AnnotationKit.strokeIsDashed": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.AnnotationKit.hasShadow": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.AnnotationKit.strokeWidth": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.AnnotationKit.highlightStyle": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"toolbarOrigin": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"com.apple.AnnotationKit.brushStyle": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.desktopservices": DomainPrefs{
			"DSDontWriteNetworkStores": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.security.pboxd": DomainPrefs{
			"ILMediaBrowserMediaType": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ILMediaBrowserSelectedGroupIndex1": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ILMediaBrowserSelectedGroup1": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.keychainaccess": DomainPrefs{
			"Keychain Item List Sorting": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Keychain Item List Sort Descending": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Last Selected Category": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"VPNSSItemsChecked": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"Item Preview Closed": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"Last Selected Keychain": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"Keychains List Closed": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.imessage": DomainPrefs{
			"PreviewTranscodingWasInitializedToDefaultValue": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"PreviewTranscodingQualityOnWiFiWasInitializedToDefaultValue": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.Safari.SafeBrowsing": DomainPrefs{
			"HasMigratedSafeBrowsingEnabledDefaults": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.AppleMediaServices": DomainPrefs{
			"AMSIncludeFullResponseInHARLogging": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"AMSShowSandboxAccountUI": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"AMSDeviceBiometricsState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"AMSUserDefaultsincludeFullResponseInHARLogging": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"AMSMigratedToNewAccountStore": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"AMSUserDefaultsIncludeFullRequestInHARLogging": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"AMSLogHARData": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"AMSLastMigratedBuildVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"AMSDidSetUpServerDataCache": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AMSIncludeFullRequestInHARLogging": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"AMSMigratedStorageToDefaultsForNonAMSInternal": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.PhotoBooth": DomainPrefs{
			"AVDesktopPlaybackControlsControllerShowsDurationInsteadOfTimeRemainingDefaultsKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"PBLaunchedAtLeastOnceOnLion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"UserCameraUniqueIDPref": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.MobileSMSPreview": DomainPrefs{
			"IMPreviewGenerationMaxPxWidth": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"IMPreviewGenerationMinHeight": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"IMPreviewGenerationScreenScale": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"IMPreviewGenerationMinWidth": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.ActivityMonitor": DomainPrefs{
			"OpenMainWindow": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ShowCategoryAppsinLast12Hours": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"SelectedTab": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ShowCategory": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"HasSavedState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.airplay": DomainPrefs{
			"showInMenuBarIfPresent": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.messages.facetime": DomainPrefs{
			"FaceTimeTwoTimeCallthroughEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.remindd": DomainPrefs{
			"CloudKitAccountStatus": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"preferredDefaultListObjectIDUrl": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ThrottlingPolicyCurrentBatchCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"isDatabaseMigrated": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"preferredDefaultListID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"cloudKitSchemaCatchUpSyncLastSuccessBuildVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"ThrottlingPolicyCurrentLevelIndex": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"spotlightIndexVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"useExtraneousAlarmBackOffThrottleInterval": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"CloudConfigurationPath": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"lastDatabaseMigrationSystemBuildVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.ncprefs": DomainPrefs{
			"content_visibility": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"sort_order": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.itunesstored": DomainPrefs{
			"AccountsNotificationPlugin.ActiveStorefrontIdentifier": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.driver.AppleHIDMouse": DomainPrefs{
			"Button4Force": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ScrollH": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ButtonDominance": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Button4": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Button4Click": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Button3": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Button1": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Button2": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ScrollSSize": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ScrollV": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ScrollS": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.preferences.extensions.CredentialProvider": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.itunescloud": DomainPrefs{
			"ICDefaultsKeyLastActiveAccountDSID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.stockholm": DomainPrefs{
			"RemoteAdminV2": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"PKWalletShouldAutomaticallyRegisterKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.Siri": DomainPrefs{
			"VoiceTriggerUserEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"StatusMenuVisible": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.sharingd": DomainPrefs{
			"AirDropRandomHashUUIDKey4": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"OneTimeAirDropReset2": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"OneTimeAirDropReset": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"HashManager-StoredDatabaseVersionKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"AirDropRandomHashUUIDKey2": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"DiscoverableMode": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"AutoUnlockPresentedWiFiError": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AutoUnlockPresentedBluetoothError": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AirDropRandomHashUUIDKey3": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"AutoUnlockWatchCurrentlyInList": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AutoUnlockWatchExistedInUnlockList": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AirDropRandomHashUUIDKey1": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.AppStoreComponents": DomainPrefs{
			"ASCLocaleID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.DFSLoginPlugin": DomainPrefs{
			"WindowLeft": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"WindowTop": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.iCal": DomainPrefs{
			"BirthdayEventsGenerationLastLocale": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"last calendar view description": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"BirthdayEventsGenerationVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"icaluuid": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"CalendarSidebarShown": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"AvailabilityShowTwentyFourHours": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"CalForceTruthFileRestoreHashKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"last selected calendar list item": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"CalDefaultCalendar": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"CalAgentNS_Preference_DefaultReminderCalendar": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"display birthdays calendar": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NSEventConcurrentProcessingEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"BirthdayEventsGenerationAttemptsToResetKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"TimeZone support enabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Show Week Numbers": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"iCal version": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"first shown minute of day": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"CalendarSidebarView": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"RemindersLastMigratedSystemVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"LastReminderMigrationCleanupVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AllDayAreaHeight": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"lastViewsTimeZone": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"privacyPaneHasBeenAcknowledgedVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"CalendarSidebarWidth": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"last minute of day time range": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"first minute of day time range": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"CalendarListMiniMonthVisibleMonths": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"IncludeDebugMenu": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"WR_DONT_ASK_FOR_DEFAULT": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"kLastABCReportTimeKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NSDontMakeMainWindowKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.assistant": DomainPrefs{
			"Experiment Identifierinvocation_feedback_experiment": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"Has Set Up MultiUser Shared Record Subscription": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"Initial Interstitial Delay for VoiceTrigger One-Shot without Audio Ducking": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"Initial Interstitial Delay - stark": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"MultiUser Shared Data Needs Sync": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"Initial Interstitial Delay": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"Myriad Device Delay": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"Myriad Device Class": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"Myriad Device Adjust": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"Has Set Up Account Status Subscription": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"Voice Trigger Needs Sync": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"Myriad Device Trump Delay": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"Experiment Identifier": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"Flush Session Tickets Cache": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"Server Has Provisioned Myriad": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"Myriad VTEndtimeDistanceThreshold": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"Has Set Up Key Value Subscription": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"Borealis Education Header Display Count": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"Experiment Identifiersiri_vox_tap_to_siri_behavior_experiment": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"Has Set Up Voice Trigger Subscription": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"Last Known Analytics Store State": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"Experiment Identifiersiri_vox_sounds_experiment": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"Manual Endpointing Threshold": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.dt.Xcode": DomainPrefs{
			"IBPreferencesMigrated": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"IBAppliesAutoResizingRulesWhileResizing": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"XCFontAndColorCurrentTheme": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"IDEKeyBindingCurrentPreferenceSet": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"IDESourceControlPreferencesVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"DVTFontAndColorLastUpdatedToolsVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"DVTDownloadableAutomaticUpdate": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"IDEUserWantsToEnableDeveloperSystemPolicyMode": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.AMPLibraryAgent": DomainPrefs{
			"iTunes-media-folder-url": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"debugAssertCategoriesVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"have-shown-cloud-UI": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"storeSupportsUPP": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"updateLevel": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"ITUserPrefsMigrated": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"books-migrated": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"storeSupportsPreviousPurchases": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"persistentUserID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"storeSupportsCloudPurchases": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"imported-media-domains": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"books-persistent-id": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"storeMediaTypeFlags": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"podcasts-migrated": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"storeSupportsPasswordSettings": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"imported-media-domains-modification-date": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"show-home-sharing-turned-on-in-iTunes-warning": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"devices-persistent-id": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"debugAssertCategoriesEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"podcasts-persistent-id": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.wifi.keychain-format": DomainPrefs{
			"Version": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.facetime.bag": DomainPrefs{
			"URL": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"CacheTime": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.TelephonyUtilities": DomainPrefs{
			"GondolaLastAccountsChangedState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"GondolaGeneratedIdentifier": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"registeredProvidersVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"CachedVCCaps": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"relayCallingDisabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"GondolaLatestVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"GondolaSyncedVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"FaceTimePhotosEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.findmy.fmfcore.notbackedup": DomainPrefs{
			"frontMostWindow": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"publicAPSToken": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"FMFGarbageCollectorIdentityKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"FMFLimitedPrecisionPrefKey.limitedPrecision": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"windowVisible": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.AppStore": DomainPrefs{
			"__uniquePageGroupID-9.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"__uniquePageGroupID-12.WebKit2DeveloperExtrasEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"mostRecentTabIdentifier": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AutoPlayVideoSetting": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.AppStore.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"appStoreBadgeCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"__uniquePageGroupID-1.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"UserSetAutoPlayVideoSetting": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"lastBootstrapTimeZone": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"__uniquePageGroupID-9.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ASAcknowledgedOnboardingVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"__uniquePageGroupID-12.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"__uniquePageGroupID-9.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"com.apple.AppStore.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"__uniquePageGroupID-1.WebKit2DeveloperExtrasEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"__uniquePageGroupID-9.WebKit2DeveloperExtrasEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AVDesktopPlaybackControlsControllerShowsDurationInsteadOfTimeRemainingDefaultsKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"__uniquePageGroupID-12.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"__uniquePageGroupID-12.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"com.apple.AppStore.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"com.apple.AppStore.WebKit2DeveloperExtrasEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"__uniquePageGroupID-1.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"__uniquePageGroupID-1.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.AdPlatforms": DomainPrefs{
			"AppStorePAAvailable": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"LatestPAVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"acknowledgedPersonalizedAdsVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.keyboard": DomainPrefs{
			"KeyboardWordOrSentenceTrackingForPFL": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.assistant.support": DomainPrefs{
			"Siri Data Sharing Opt-In Status": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Assistant Enabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Dictation Enabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.preferences.extensions.QuickLook": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.internetconfigspec": DomainPrefs{
			"WWWHomePage": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.icloud.fmfd": DomainPrefs{
			"numberOfFriendsFollowersKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"kFMFDStoredDataVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"storedDSID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"lazyInitTimeSecsStoredKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"kFMFDlastLoggedInPrsId": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.newscore": DomainPrefs{
			"notificationEnableAssetPrefetching": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.news.notification_receipt_event_endpoint": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"news_url_resolution_subscription_status": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"com.apple.news.default_event_endpoint": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"notificationAssetPrefetchingRequiresWatch": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"notificationArticleDiversificationSimilarityExpectationStart": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"force_refresh_user_segmentation": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"notificationArticleDiversificationUniquePublisherExpectationYIntercept": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"notificationArticleDiversificationUniquePublisherExpectationSlope": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"news_carplay_is_enabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.newsd.download.hasUnfinishedWork": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"notificationArticleDiversificationSimilarityExpectationEnd": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"use_parsec_results_for_widget": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.AddressBook": DomainPrefs{
			"ABBirthDayVisible": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ABMetaDataChangeCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ABTextSizeIncrement": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ABVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"ABLastImportShown": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ABDefaultSourceID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NSPreferencesSelectedIndex": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AB21vCardEncoding": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NSPreferencesContentSize": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ABPrivateVCardFieldsEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.SpeechRecognitionCore": DomainPrefs{
			"AllowAudioDucking": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.commerce": DomainPrefs{
			"LastUpdateNotificationOSMajorVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"PurchasesInflight": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.networkresolutiond": DomainPrefs{
			"_networkDevices": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.TV": DomainPrefs{
			"automaticallyDownloadArtwork": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"userWantsPlaybackNotifications": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"gridViewSearchDoesNotSwitchView": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"persistentUserID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"miniplayerUserSetHeight": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"storeMediaTypeFlags": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"dontAskForAllPlaylistItemRemoval": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"checkedHLSKeysTime": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"storeSupportsCloudPurchases": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"debugAssertCategoriesVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"com.apple.TV.WebKit2EnableInheritURIQueryComponent": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"cddbPrefsOK": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"com.apple.TV.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"miniplayerVPos": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"controllableInterfaceGUID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"automaticallyDeleteVideoAssetsAfterWatching": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"videoWindowVPos": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"miniplayerHPos": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"playbackIsFullscreen": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"storeSupportsPasswordSettings": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"com.apple.TV.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"haveAskedToCheckForAvailableDownloads": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"videoWindowHPos": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"miniplayerSnapMode": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"storeSupportsUPP": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"showWelcomeScreenState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"videoWindowVSize": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"updateLevel": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"ITUserPrefsMigrated": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"miniplayerWidthInDesignCoords": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"storeSupportsPreviousPurchases": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"tabViewMode": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"preloadFilesIntoMemory": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"userMaxConcurrentDownloads": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NSApplicationCrashOnExceptions": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"videoWindowHSize": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"checkForAvailableDownloads": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"miniPlayerQueueVisible": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"debugAssertCategoriesEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"firstLaunchShowWelcomeScreenState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"com.apple.TV.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"musicVideoPlaybackLocation": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"miniPlayerLargeArtVisible": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"userHasActivatedFullScreenVisualizer": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"com.apple.TV.WebKit2AsynchronousSpellCheckingEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"com.apple.TV.WebKit2ShouldPrintBackgrounds": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.preference.general": DomainPrefs{
			"version": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.menuextra.battery": DomainPrefs{
			"ShowPercent": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.imdpersistence.IMDPersistenceAgent": DomainPrefs{
			"LogAllIOErrors": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.SafariServices": DomainPrefs{
			"SearchProviderIdentifierMigratedToSystemPreference": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.siri.shortcuts": DomainPrefs{
			"WFServicesShortcutsMigrated": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"WFDidUnconflictShortcuts": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"LegacyShortcutsZoneSubscriptionUnsubscribed": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"WFDefaultShortcutsVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"WFLastSyncedFlagsHash": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.SystemProfiler": DomainPrefs{
			"PrefsVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"SPLastDocumentsSize": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.remindd.babysitter": DomainPrefs{
			"LastSystemVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.internal.ck": DomainPrefs{
			"DictationOnDeviceSamplingDisabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"DictationSamplingRates": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"DefaultCacheKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"DefaultWarmupScriptsExtension": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"WarmupScriptIdentifiers": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"HasSetUpRecordZoneSubscription": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ByteCodeCacheEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"DefaultBootTimeUpdateScripts": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"WarmupModularScriptIdentifiers": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"DisableFBFForUEI": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.internetconfigpriv": DomainPrefs{
			"WWWHomePage": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.dock": DomainPrefs{
			"wvous-br-corner": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"showAppExposeGestureEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"autohide": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"tilesize": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"no-bouncing": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"region": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"showhidden": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"launchanim": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"showMissionControlGestureEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"wvous-tr-modifier": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"magnification": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"largesize": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"mod-count": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"trash-full": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"minimize-to-application": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"loc": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"expose-animation-duration": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"wvous-tr-corner": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"autohide-delay": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.AdLib": DomainPrefs{
			"adprivacydMaxSegmentSendInterval": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"personalizedAdsMigrated": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"partiality-segment": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"CKDPIDSyncState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"forceLimitAdTracking": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"adprivacydSegmentInterval": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"allowIdentifierForAdvertising": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"allowApplePersonalizedAdvertising": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.amp.mediasharingd": DomainPrefs{
			"public-sharing-enabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"home-sharing-sequence-id": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"home-sharing-group-id": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"home-sharing-computer-id": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"photo-sharing-enabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"shared-library-id": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"shared-library-name": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"public-sharing-is-protected": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"shared-library-machine-id": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"home-sharing-enabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.systempreferences": DomainPrefs{
			"com.apple.preferences.sharing.selectedservice": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ThirdPartyCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"SoundTab": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"com.apple.dtandsspref.lastselectedtab": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NSFullScreenMenuItemEverywhere": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"trackpad.lastselectedtab": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"ShowAllMode": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ICQ_iCloudFirstRunNotificationShown": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"DisableAutoLoginButtonIsHidden": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"com.apple.preferences.accounts.outline.usersparent": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"mouse.lastselectedtab": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NSQuitAlwaysKeepsWindows": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.SecurityPref.Privacy.LastSourceSelected": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"SecurityPrefTab": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"com.apple.preference.keyboard.lastselectedtab": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"com.apple.datetimepref.lastselectedtab": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.UIKit": DomainPrefs{
			"hasAccessibilityBeenMigrated": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.SafariBookmarksSyncAgent": DomainPrefs{
			"CloudBookmarksSupplementalClientIdentifier": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"NewestLaunchedSafariBookmarksSyncAgentVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.VoiceOverTraining": DomainPrefs{
			"doNotShowSplashScreen": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.newsd": DomainPrefs{
			"FCAppConfigurationBundleShortVersionKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.itunescloud.daemon": DomainPrefs{
			"ICDDefaultsKeyKnownActiveAccountDSID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.Image_Capture": DomainPrefs{
			"IK_lastUsedDeviceUUID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"IK_Camera_selectedTag": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"IK_lastUsedDeviceIsRemote": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"IK_Scanner_downloadURL": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"IKC_sort_ascending": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.imagekit.cameraviewmode": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"IK_prefsVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"IK_Camera_selectedPathType": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"IKC_sort_key": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"IK_Accessory_selectedTag": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"IK_Scanner_PostProcessApplication": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"IK_Camera_preferPostPocessingApp": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"IK_lastUsedDeviceName": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.preferences.extensions.SharedLinks": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.driver.AppleBluetoothMultitouch.mouse": DomainPrefs{
			"MouseOneFingerDoubleTapGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"MouseTwoFingerHorizSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"MouseVerticalScroll": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"UserPreferences": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"MouseHorizontalScroll": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"MouseButtonDivision": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"MouseMomentumScroll": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"MouseButtonMode": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"MouseTwoFingerDoubleTapGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.dock.external.extra.x86_64": DomainPrefs{
			"ShowSolidarity": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.commcenter.data": DomainPrefs{
			"pw_ver": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.imagecapture": DomainPrefs{
			"loggingLevel": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.notificationcenterui": DomainPrefs{
			"bannerTime": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.pipagent": DomainPrefs{
			"Size": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.AMPDevicesAgent": DomainPrefs{
			"debugAssertCategoriesEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"persistentUserID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"dontAutomaticallySyncIPods": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"updateLevel": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"storeMediaTypeFlags": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"AutomaticDeviceBackupsDisabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"storeSupportsCloudPurchases": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"debugAssertCategoriesVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"ITUserPrefsMigrated": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"userMaxConcurrentDownloads": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.MobileSMS.CKDNDList": DomainPrefs{
			"CatalystDNDMigrationVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"CKDNDMigrationKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.protectedcloudstorage.protectedcloudkeysyncing": DomainPrefs{
			"registrySyncIdentifier": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.VoiceOver4.local": DomainPrefs{
			"SCRCUserDefaultsAllowAirPlay": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"SCRCUserDefaultsUnplannedShutdownCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"SCRCUserDefaultsPlannedShutdownSuccessful": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.suggestions": DomainPrefs{
			"SuggestionsAllowGeocode": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"spToLearnMigrationPerformed": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"findToShowMigrationPerformed": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.NetworkUtility": DomainPrefs{
			"NUPingCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NUTracerouteAddress": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NUPortScanEnd": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NUSelectedTabItem": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NULookupAddress": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NUFingerPerson": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NUPortScanAddress": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NUPingChoice": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NUPingAddress": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NUNetstatChoice": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NUPortScanStart": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NUWhoisAddress": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NUWhoisSelectedServer": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"NUPortScanRange": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.siri.VoiceShortcuts": DomainPrefs{
			"VCLSDatabaseUUIDKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"VCLSDataSequenceKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.Safari": DomainPrefs{
			"ResetCloudHistory": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"WebKitHistoryItemLimit": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"WebKitInitialTimedLayoutDelay": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ShowFullURLInSmartSearchField": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"com.apple.Safari.ContentPageGroupIdentifier.WebKit2DeveloperExtrasEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"cloudBookmarksMigrationEligibilityDataInvalidated": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"UniversalSearchEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"WebKitDeveloperExtrasEnabledPreferenceKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"SuppressSearchSuggestions": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"WebKitHistoryAgeInDaysLimit": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"IncludeDebugMenu": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"IncludeDevelopMenu": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"IncludeInternalDebugMenu": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.accounts": DomainPrefs{
			"ACOneTimeSMTPFixAccountSync": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"ACOneTimeLDAPFixAccountSync": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.visualintelligence": DomainPrefs{
			"bypassAvailability": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"enableBooksDomain": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"enableLandmarkDomain": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"enableArtDomain": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"firstTimeExperience": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"enableCoarseClassification": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"initialized": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"enablePetsDomain": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"enableScreenshots": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"debugUI": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"enableSafariApp": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"sendLocationInfo": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"enableAlbumsDomain": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"enableQuickLook": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"enableNatureDomain": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"enablePhotosApp": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"sendOCRText": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.accountsd": DomainPrefs{
			"com.apple.mail.searchableIndex.lastProcessedAttachmentIDKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"LastSystemVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.fileproviderd": DomainPrefs{
			"com.apple.fileproviderd.did-drop-daemon-corespotlight-index": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.screencapture": DomainPrefs{
			"location": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"style": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"captureDelay": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"last-selection-display": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"video": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.SoftwareUpdate": DomainPrefs{
			"AutoUpdateMajorOSVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.Preferences": DomainPrefs{
			"UserDictionaryImportedSinceMaintenance": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"UserDictionaryLocalPeerIdentityCurrent": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.sociallayerd.CloudKit.ckwriter": DomainPrefs{
			"accountLastKnownUserRecordID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"accountAvailable": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"writerDone": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.classroom": DomainPrefs{
			"allowClassroomLockDevice-00000000-0000-0000-0000-000000000000:0": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"allowClassroomOpenApp-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"allowClassroomScreenObservation-00000000-0000-0000-0000-000000000000:0": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"allowClassroomOpenURL-ask": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"allowClassroomScreenObservation-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"allowClassroomOpenURL": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"allowClassroomLockDevice-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"allowClassroomOpenURL-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"allowClassroomOpenApp-00000000-0000-0000-0000-000000000000:0": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"forceUnpromptedRemoteScreenObservation-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"allowClassroomOpenURL-00000000-0000-0000-0000-000000000000:0": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"forceUnpromptedRemoteScreenObservation-00000000-0000-0000-0000-000000000000:0": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"allowClassroomOpenApp": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"allowClassroomOpenApp-ask": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.CloudKit": DomainPrefs{
			"AccountInfoValidationCounter": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.findmy.fmipcore.notbackedup": DomainPrefs{
			"publicAPSToken": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"itemLearnMoreURL": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"frontMostWindow": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"windowVisible": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"deviceImageVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"FMIPGarbageCollectorIdentityKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"FMIPLimitedPrecisionPrefKey.limitedPrecision": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.tipsd": DomainPrefs{
			"TPSWelcomeNotificationViewedVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"TPSTipsAppInstalled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"TPSOfflineLastProcessedRemoteContentIdentifier": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"DeliveryInfoVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"TPSLastMajorVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"ConsecutiveNotificationsCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"TPSWelcomeNotificationReminderState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.CoreGraphics": DomainPrefs{
			"DisplayUseForcedGray": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"DisplayUseInvertedPolarity": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.DirectoryUtility": DomainPrefs{
			"FirstLaunch": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.AvatarUI.Staryu": DomainPrefs{
			"AVTAvatarHasDisplayedAnimojiSplashScreen": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AVTAvatarHasDisplayedCameraEffectsSplashScreen": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AVTAvatarHasDisplayedPaddleView": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AVTAvatarHasDisplayedSplashScreen": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AVTAvatarUILastCacheVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.networkConnect": DomainPrefs{
			"VPNShowTime": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.UserAccountUpdater": DomainPrefs{
			"com.apple.HidLibraryFolderAlready": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.madrid": DomainPrefs{
			"IMCloudKitStartingInitialSync": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"CloudKitIsSyncing": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"CloudKitIsRemovedFromBackup": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"CloudKitIsEligibleForTruthZone": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"CloudKitEligibleToToggleMiCSwitch": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"IMCloudKitSyncControllerSyncTypeKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"CloudKitCheckedMiCSwitchEligibilityOnImagentLaunch": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"createdChatZone": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"hasCompletedInitialCKSync": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"CloudKitSyncingEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"IMCloudKitSyncControllerSyncRecordTypeKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"initialSyncRecordHasBeenWritten": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"IMCloudKitSyncControllerSyncStateKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"IMCloudKitStartingPeriodicSync": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"IMCloudKitSyncPaused": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"RequestPriorityRamp": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"IMCloudKitAccountStatusKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"IMCloudKitStartingEnabledSettingChange": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"IMCKFinishedFetchingAttachmentsFromCloudKit": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"CKMOCAccountsMatch": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"IMCloudKitStartingDisableDevices": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"oneTimeMOCAccountCheckV2": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"enableCKSyncingV2": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.rapport": DomainPrefs{
			"familySyncedName": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.menuextra.clock": DomainPrefs{
			"DateFormat": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.imessage.bag": DomainPrefs{
			"URL": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"CacheTime": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.sms": DomainPrefs{
			"hasBeenApprovedForSMSRelay": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.Automator": DomainPrefs{
			"NSSplitView AMDocumentMinor Expanded Position": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"NSSplitView AMLibraryActionsMajor Expanded Position": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AMLogTabViewSelectedIndex": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.weather.internal": DomainPrefs{
			"userId": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.FaceTime": DomainPrefs{
			"DidMigrateToSixteenByNineAspectRatio": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"kPromptEnableReadRecipts": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"CustomRingtone": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"PhoneNumberUpgradeShown": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"PreferredVideoDeviceUID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.AssistiveControl.virtualKeyboard": DomainPrefs{
			"SCLaunchedAsSlave": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"PanelPosition": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"DesiredPanelWindowPosition": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.Accessibility": DomainPrefs{
			"InvertColorsEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"GrayscaleDisplay": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"AXSClassicInvertColorsPreference": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"KeyRepeatInterval": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"CommandAndControlEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ApplicationAccessibilityEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"DarkenSystemColors": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"FullKeyboardAccessFocusRingEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"GenericAccessibilityClientEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"AccessibilityEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"KeyRepeatEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"KeyRepeatDelay": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ZoomTouchEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"DifferentiateWithoutColor": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"FullKeyboardAccessEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"BrailleInputDeviceConnected": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"AutomationEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ReduceMotionEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"SpeakThisEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"EnhancedBackgroundContrastEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"VoiceOverTouchEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.ipTelephony": DomainPrefs{
			"ImsLoggingEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.ColorSyncCalibrator": DomainPrefs{
			"722524374X-shareMode": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"722524374X-targetTemp": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"722524374X-name": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"722524374X-expertMode": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.StatusKitAgent": DomainPrefs{
			"reauthCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.AppleMultitouchTrackpad": DomainPrefs{
			"TrackpadThreeFingerDrag": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadPinch": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"USBMouseStopsTrackpad": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadFourFingerVertSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ActuationStrength": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ActuateDetents": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"FirstClickThreshold": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"SecondClickThreshold": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadTwoFingerDoubleTapGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadHorizScroll": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadMomentumScroll": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadFourFingerPinchGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadRotate": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Clicking": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadThreeFingerTapGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadThreeFingerHorizSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadScroll": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"DragLock": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadFiveFingerPinchGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"ForceSuppressed": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"TrackpadThreeFingerVertSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadTwoFingerFromRightEdgeSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"UserPreferences": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"HIDScrollZoomModifierMask": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadHandResting": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadRightClick": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Dragging": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadFourFingerHorizSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"TrackpadCornerSecondaryClick": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.ImageCaptureService": DomainPrefs{
			"IK_Accessory_selectedTag": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.screensaver": DomainPrefs{
			"tokenRemovalAction": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.AppleMultitouchMouse": DomainPrefs{
			"MouseOneFingerDoubleTapGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"MouseTwoFingerHorizSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"MouseVerticalScroll": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"UserPreferences": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"MouseHorizontalScroll": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"MouseButtonDivision": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"MouseMomentumScroll": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"MouseButtonMode": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"MouseTwoFingerDoubleTapGesture": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.universalaccessAuthWarning": DomainPrefs{
			"2::/Applications/Logi Options.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::com.logitech.manager.daemon": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Microsoft Teams (work or school).app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::com.lightpillar.Mosaic-setapp": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"1::/Applications/CloudBerry Backup.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::com.getcleanshot.app-setapp": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::/Applications/Setapp/Mosaic.app/Contents/MacOS/Mosaic": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"/Applications/Clipy.app/Contents/MacOS/Clipy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/xScope.app/Contents/MacOS/xScope": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Microsoft Teams.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"1::/Users/mikeschinkel/Library/Application Support/JetBrains/Toolbox/apps/Goland/ch-0/232.8660.185/GoLand.app/Contents/MacOS/goland": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"1::/Applications": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::com.microsoft.teams2": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Discord.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Microsoft Edge.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::/Applications/Logi Options.app/Contents/Support/LogiMgrDaemon.app/Contents/MacOS/LogiMgrDaemon": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Skitch.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/ScreenFlow.app/Contents/Library/LoginItems/ScreenFlowHelper.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"/Applications/Alfred 4.app/Contents/MacOS/Alfred": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Toast 19 Pro.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"1::com.vmware.fusion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::com.lastpass.lastpassmacdesktop": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"com.clipy-app.Clipy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::com.google.Chrome": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"1::/Users/mikeschinkel/Library/Application Support/JetBrains/Toolbox/apps/Goland/ch-0/232.8660.185/GoLand.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"1::/Applications/Microsoft Teams.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/ScreenFlow.app/Contents/Library/LoginItems/ScreenFlowHelper.app/Contents/MacOS/ScreenFlowHelper": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Microsoft Teams (work or school).app/Contents/MacOS/MSTeams": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/RingCentral.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::com.amazon.Amazon-Chime": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"com.mutedeck.mac": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Parallels Toolbox.app/Contents/Applications/Hide Menu Icons.app/Contents/MacOS/Hide Menu Icons": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::/Applications/Snappy.app/Contents/MacOS/Snappy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::me.waydabber.BetterDummy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/GoToMeeting.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Google Chrome.app/Contents/MacOS/Google Chrome": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::WSG985FR47.net.telestream.screenflowhelper": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Slack.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"/Applications/Alfred 4.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"1::/Applications/BetterDisplay.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::com.logi.cp-dev-mgr": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"com.runningwithcrayons.Alfred": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Xnapper.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/xScope.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::com.microsoft.edgemac": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"/Applications/zoom.us.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"/System/Applications/Automator.app/Contents/MacOS/Automator": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"/Applications/Clipy.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/BetterDisplay.app/Contents/MacOS/BetterDummy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Snip.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::ro.nextwave.Snappy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::com.parallels.toolbox.HideMenuIcons": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Google Chrome.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"1::/Applications/VMware Fusion-8.5.10.app/Contents/MacOS/VMware Fusion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::com.ringcentral.glip": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Acorn.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"/Applications/MuteDeck/MuteDeck.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::/Library/Application Support/Logitech.localized/LogiOptionsPlus/logioptionsplus_agent.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Timing.app/Contents/Library/LoginItems/TimingHelper.app/Contents/MacOS/TimingHelper": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::/Applications/Logi Options.app/Contents/Support/LogiMgrDaemon.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::/Applications/Setapp/Mosaic.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"1::me.waydabber.BetterDummy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::/Applications/LastPass.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::info.eurocomp.TimingHelper": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::com.amazon.Amazon-Chime": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"1::/Applications/BetterDisplay.app/Contents/MacOS/BetterDummy": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"1::/Applications/iTerm.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"1::/Applications/OBS.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"1::/Applications/HP RGS Receiver.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::J8RPQ294UB.com.skitch.SkitchHelper": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::com.logitech.Logi-Options": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::/Applications/Logi Options.app/Contents/MacOS/Logi Options": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"us.zoom.xos": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"1::/Applications/Slack.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::/Applications/Snappy.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"1::com.jetbrains.goland": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::info.eurocomp.Timing2": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Xnapper.app/Contents/MacOS/Xnapper": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/RingCentral.app/Contents/MacOS/RingCentral": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"/Applications/zoom.us.app/Contents/MacOS/zoom.us": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/BetterDisplay.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/zoom.us.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"com.apple.Automator": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Skitch.app/Contents/Library/LoginItems/J8RPQ294UB.com.skitch.SkitchHelper.app/Contents/MacOS/J8RPQ294UB.com.skitch.SkitchHelper": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Amazon Chime.app/Contents/MacOS/Amazon Chime": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Setapp/CleanShot X.app/Contents/MacOS/CleanShot X Setapp": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/ScreenFlow.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::/Applications/Amazon Chime.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"1::/Applications/VirtualBox.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"/Applications/MuteDeck/MuteDeck.app/Contents/MacOS/MuteDeck": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Timing.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::/Applications/Amazon Chime.app/Contents/MacOS/Amazon Chime": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Timing.app/Contents/MacOS/Timing": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Amazon Chime.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Skitch.app/Contents/Library/LoginItems/J8RPQ294UB.com.skitch.SkitchHelper.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::com.iconfactory.xScope": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::/Applications/LastPass.app/Contents/MacOS/LastPass": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"/System/Applications/Automator.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"1::/Applications/VMware Fusion-8.5.10.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Loom.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/var/tmp/LogMeIn Rescue - 4C61/Support-LogMeInRescue.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Setapp/CleanShot X.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::com.devuap.beautyshotapp": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Timing.app/Contents/Library/LoginItems/TimingHelper.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/var/tmp/LogMeIn Rescue - 4C61": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"3::/Applications/Parallels Toolbox.app/Contents/Applications/Hide Menu Icons.app": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"2::/Library/Application Support/Logitech.localized/LogiOptionsPlus/logioptionsplus_agent.app/Contents/MacOS/logioptionsplus_agent": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.findmy": DomainPrefs{
			"restoreState": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"coarseServiceAcknowledged_v1.0": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"onboarding_v2.0": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.FolderActionsDispatcher": DomainPrefs{
			"folderActionsEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.SocialLayer": DomainPrefs{
			"SharedWithYouEnabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.security.KCN": DomainPrefs{
			"absentCircleWithNoReason": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
			"lastCircleStatus": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"pendingApplicationReminderInterval": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.GEO": DomainPrefs{
			"DefaultsSanitizedVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.SharedWebCredentials": DomainPrefs{
			"recheckVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.suggestd": DomainPrefs{
			"DeviceIdentifier": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.CallHistorySyncHelper": DomainPrefs{
			"/Users/mikeschinkel/Library/Application Support/CallHistoryTransactions/transaction.log": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"CallHistoryDeviceCount": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"/Users/mikeschinkel/Library/Application Support/CallHistoryTransactions/transactions.log": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.dashboard": DomainPrefs{
			"devmode": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"mcx-disabled": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.SystemManaged,
				},
			},
		},
		"com.apple.FontRegistry.user": DomainPrefs{
			"Version": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.imagent": DomainPrefs{
			"Setting.EnableReadReceipts": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
			"Setting.GlobalReadReceiptsVersionID": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"changeToken-syncStoreVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"IMChatVocabularyUpdaterDidPerformInitialUpdateKey": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
			"attachmentZoneChangeToken-syncStoreVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"messagesChangeToken-syncStoreVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
			"archivedMessagesChangeToken-syncStoreVersion": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.VersionMigration,
				},
			},
		},
		"com.apple.preferences.extensions.PhotoProjects": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
		"com.apple.messaging.watchdog": DomainPrefs{
			"watchdogWatermark": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.RuntimeState,
				},
			},
		},
		"com.apple.preferences.extensions.PhotoEditing": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: kvfilters.Labels{
					&kvfilters.UserManaged,
				},
			},
		},
	}
}
