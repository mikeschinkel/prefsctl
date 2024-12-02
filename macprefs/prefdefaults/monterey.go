package prefdefaults

import (
	"github.com/mikeschinkel/prefsctl/macprefs/filters"
)

//goland:noinspection SpellCheckingInspection
func montereyPrefDefaults() DomainDefaults {
	return DomainDefaults{
		"GlobalPreferences": DomainPrefs{
			"AppleActionOnDoubleClick": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"AppleAntiAliasingThreshold": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"AppleAquaColorVariant": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"AppleEnableSwipeNavigateWithScrolls": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"AppleKeyboardUIMode": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"AppleMeasurementUnits": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"AppleMetricUnits": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"AppleMiniaturizeOnDoubleClick": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"AppleShowAllExtensions": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"AppleShowScrollBars": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"AppleTemperatureUnit": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"AppleWindowTabbingMode": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.mouse.scaling": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.sound.beep.flash": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.springing.delay": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.springing.enabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.swipescrolldirection": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.trackpad.forceClick": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.trackpad.scaling": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"InitialKeyRepeat": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"KeyRepeat": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NSAutomaticCapitalizationEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NSAutomaticDashSubstitutionEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NSAutomaticPeriodSubstitutionEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NSAutomaticQuoteSubstitutionEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NSAutomaticSpellingCorrectionEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NSAutomaticTextCompletionEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NSAutomaticWindowAnimationsEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"WebAutomaticSpellingCorrectionEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"WebKitDeveloperExtras": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"AppleLanguage": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
					&filters.SetupSets,
					&filters.StringType,
				},
			},
			"AppleLocale": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SetupSets,
					&filters.StringType,
				},
			},
			"646F6E7A_00000000_00000001_6E7A6361_696D6963:": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"AKLastIDMSEnvironment:": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"AudioQuest inc.AudioQuest DragonFly:": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"Blue Microphones Yeti Stereo Microphone:": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"Canon MF5900 Series:": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"ContextMenuGesture:": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"Fujitsu ScanSnap iX500:": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"Yubico YubiKey OTP+FIDO+CCID:": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"AppleLanguagesDidMigrate:": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"NavPanelFileListModeForOpenMode:": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"NavPanelFileListModeForSaveMode:": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NSLinguisticDataAssetsRequestLastInterval:": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"QLPanelAnimationDuration:": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.security.sosaccount": DomainPrefs{
			"SOSEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.SafariCloudHistoryPushAgent": DomainPrefs{
			"AcknowledgedPushNotifications": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"UnacknowledgedPushNotifications": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.print.custompresets.forprinter.Canon_MF5900_Series": DomainPrefs{
			"com.apple.print.lastPresetPref": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"com.apple.print.lastPresetPrefType": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.finder": DomainPrefs{
			"SidebarTagsSctionDisclosedState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"QLEnableTextSelection": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"PreviewPaneGalleryWidth": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"FK_ArrangeBy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"PreferencesWindow.LastSelection": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"FXEnableExtensionChangeWarning": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NSWindowTabbingShoudShowTabBarKey-com.apple.finder.TBrowserWindow": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"FXICloudDriveDesktop": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ShowMountedServersOnDesktop": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"SidebarShowingiCloudDesktop": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"FXToolbarUpgradedToTenNine": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"ShowRemovableMediaOnDesktop": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ShowStatusBar": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"FXSidebarUpgradedSharedSearchToTwelve": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"FXInfoWindowWidth": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"DownloadsFolderListViewSettingsVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"SidebariCloudDriveSectionDisclosedState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AppleShowAllFiles": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"FXICloudDriveDeclinedUpgrade": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"FXArrangeGroupViewBy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"RecentsArrangeGroupViewBy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"FXPreferredSearchViewStyleVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"PreviewPaneWidth": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"FXAtLeastOneScreenHasBeenInHIDPI": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"CopyProgressWindowLocation": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"SidebarShowingSignedIntoiCloud": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ViewOptionsWindow.Location": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ShowHardDrivesOnDesktop": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"FXQuickActionsConfigUpgradeLevel": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"FXPreferredViewStyle": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"QuitMenuItem": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"FXDefaultSearchScope": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"FXPreferredSearchViewStyle": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"DisableAllAnimations": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"SidebarWidth": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ShowSidebar": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TagsCloudSerialNumber": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NewWindowTarget": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ShowPathbar": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"FXSidebarUpgradedToTenFourteen": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"LastTrashState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ShowExternalHardDrivesOnDesktop": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"GoToField": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"_FXSortFoldersFirstOnDesktop": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"FK_AppCentricShowSidebar": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"FXSidebarUpgradedToTenTen": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"FXConnectToLastURL": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"FXICloudDriveEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"FXICloudDriveDocuments": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"_FXSortFoldersFirst": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"_FXShowPosixPathInTitle": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"EmptyTrashProgressWindowLocation": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"SearchRecentsSavedViewStyle": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"FXDesktopTouchBarUpgradedToTenTwelveOne": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"FXICloudDriveFirstSyncDownComplete": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"SidebarDevicesSectionDisclosedState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"MountProgressWindowLocation": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"FK_SidebarWidth": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"SearchRecentsSavedViewStyleVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"FXPreferredGroupBy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"PreviewPaneInfoExpanded": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"FXLastSearchScope": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"InspectorWindow.Location": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"SidebarPlacesSectionDisclosedState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"FXToolbarUpgradedToTenEight": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"FXICloudLoggedIn": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"FXToolbarUpgradedToTenSeven": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.messages.pinning": DomainPrefs{
			"IMPinningPinConfigMigrationKey-v2-r2": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.identityservicesd": DomainPrefs{
			"ReRegisteredForDevicesv56": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"hasRegIdentityContainer": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"ReRegisteredForDevicesv55": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"persister-migration-personal-session-token-cache-v4": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"ReRegisteredForDevices": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"DidCleanLegacyAccountPrefs": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"ImportedLegacyIMAccounts": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"hasUnregIdentityContainer": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"MigratedToNewDisabledState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"triggeredRemoteSessionVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"ReRegisteredForDevicesHash": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"ImportedLegacyIDSAccounts2": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"GDRRequestMadeForRelayRepair": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"persister-migration-com.apple.identityservices.userStore": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"ReRegisteredForDevicesv1350": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"performed-user-intent-migrate": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.Spotlight": DomainPrefs{
			"ModelName": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.AppManaged,
				},
			},
			"engagementCount-com.apple.Spotlight": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"engagementCount-com.apple.mail": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"SSMessageTracingWindowShowCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"windowHeight": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"useCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"__NSEnableTSMDocumentWindowLevel": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"SSMessageTracingWindowHideCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"engagementCount-com.apple.Spotlight.suggestions": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"SPMessageTracingWindowHideCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"showedFTE": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"SPMessageTracingWindowShowCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"userHasMovedWindow": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.DictionaryServices": DomainPrefs{
			"DCSPreferenceVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.PTPCamera": DomainPrefs{
			"FilenamePrefix": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NextFilenameCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.imtranscoding.IMTranscoderAgent": DomainPrefs{
			"kCKMediaObjectManagerDefaultsOSVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.scriptmenu": DomainPrefs{
			"ScriptMenuEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.TTY": DomainPrefs{
			"RTTSettingsVersionPreference": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"RTTContinuityRTTIsSupportedPreference": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"TUSupportsRelayCallingPreference": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"TUIsRelayCallingEnabledPreference": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.voicetrigger": DomainPrefs{
			"VoiceTrigger Enabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.DiskUtility": DomainPrefs{
			"OperationProgress ExpandedHeight": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"DUDebugMenuEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"advanced-image-options": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"SidebarShowAllDevices": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"OperationProgress DetailsVisible": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.photos.shareddefaults": DomainPrefs{
			"CPLDefaultDownload": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"iCloudPhotoLibraryLastResetWelcomePromptDBVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"downloadAndKeepOriginals": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"clearPurgeableResources": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"DiskSpaceWasLow": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"repushAssetsWithImportedByBundleIdentifier": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.AppManaged,
				},
			},
			"repushVideoAssetsMetadata": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.AppManaged,
				},
			},
		},
		"com.apple.parsecd": DomainPrefs{
			"PARDefaultsVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"PARTestSeed": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.universalaccess": DomainPrefs{
			"login": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"switchCoalescePressesDuration": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"lastNightShiftEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"switchAutoScanPanelInterval": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"closeViewZoomedIn": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"grayscale": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"switchSweepingCursorSpeed": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"sessionChange": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"switchHoldBeforeRepeatDuration": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"closeViewZoomFactorBeforeTermination": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"switchHideUITimeout": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"switchAutoScanElementInterval": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"lastNightShiftActive": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"virtualKeyboardOnOff": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"headMouseEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"increaseContrast": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"mouseDriverCursorSize": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"selectedTab": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"lastNightShiftDate": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"closeViewZoomDisplayID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"hoverTextIsAlwaysOn": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"closeViewSplitScreenRatio": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"hoverTextIsHoveringAndVisible": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AssistiveControlType": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"closeViewDesiredZoomFactor": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"slowKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"closeViewScrollWheelToggle": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"mouseDriver": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"differentiateWithoutColor": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"slowKeyDelay": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"switchMinimumPressDuration": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"closeViewZoomFactor": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"switchOnOffKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"keyboardAccessFocusRingTimeout": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"stickyKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"reduceTransparency": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"dwellHideUITimeout": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"dwellEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"contrast": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"useStickyKeysShortcutKeys": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"lastNightShiftMode": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"hoverTextEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"whiteOnBlack": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"closeViewHotkeysEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"switchFirstElementDelay": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"voiceOverOnOffKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.Terminal": DomainPrefs{
			"HasMigratedDefaults": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"SecureKeyboardEntry": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Default Window Settings": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"DefaultProfilesVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"ProfileCurrentVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"TTAppPreferences Selected Tab": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"Startup Window Settings": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.coreservices.useractivityd": DomainPrefs{
			"kLocalPasteboardBlobName": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"kRemotePasteboardBlobName": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.ColorSyncUtility": DomainPrefs{
			"CSUtilitiesContentSize": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"CSProfsUtilsGroupBy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"CSUtilitiesSelectedIdentifier": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.weather": DomainPrefs{
			"userId": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.Music": DomainPrefs{
			"dontAskForAllPlaylistItemRemoval": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"storeSupportsPasswordSettings": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"notifications-warming-last-displayed-time": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"miniplayerVPos": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"com.apple.Music.WebKit2ShouldPrintBackgrounds": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"storeSupportsPreviousPurchases": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"com.apple.Music.WebKit2AsynchronousSpellCheckingEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"miniPlayerQueueVisible": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"eqWindowHPos": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"userWantsPlaybackNotifications": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"miniplayerUserSetHeight": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"com.apple.Music.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"soundEnhancerAmount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"preloadFilesIntoMemory": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"updateLevel": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"videoWindowVSize": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"com.apple.Music.WebKit2EnableInheritURIQueryComponent": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"automaticallyDownloadArtwork": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"refreshedHLSKeysTime": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"miniplayerWidthInDesignCoords": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"userMaxConcurrentDownloads": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.Music.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"miniplayerHPos": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"videoWindowVPos": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"log-PlayQueue-LastSelectedTab": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"dontWarnDownloadCloudPurchases": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"storeSupportsUPP": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"showWelcomeScreenState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"checkedHLSKeysTime": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"storeSupportsCloudPurchases": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"gridViewSearchDoesNotSwitchView": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"eqName": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ITUserPrefsMigrated": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"cddbPrefsOK": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"eqPrefsVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"playbackIsFullscreen": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"debugAssertCategoriesVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"debugAssertCategoriesEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"firstLaunchShowWelcomeScreenState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"encoderName": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.Music.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"haveRadioState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"videoWindowHSize": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"eqWindowVPos": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"haveAskedToCheckForAvailableDownloads": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"musicVideoPlaybackLocation": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"volumeWSG": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"dontShowRestrictionsPrefs": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"persistentUserID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"NSApplicationCrashOnExceptions": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"miniPlayerLargeArtVisible": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"searchSavedTab": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"artworkDownloadDSID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"videoWindowHPos": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"controllableInterfaceGUID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"userHasActivatedFullScreenVisualizer": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"storeMediaTypeFlags": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"imported-eq-presets": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"soundEnhancerEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"miniplayerSnapMode": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"doesStoreSupportCloudMusicLibrary": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"eqEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.airport.airportutility": DomainPrefs{
			"dontPerformBaseRestartWarning": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.studentd": DomainPrefs{
			"LastDateProviderSessionToken": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"DeviceIdentifier": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.systemuiserver": DomainPrefs{
			"__NSEnableTSMDocumentWindowLevel": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.IMCoreSpotlight": DomainPrefs{
			"IMCSIdxVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"IMCSLastIndexedRowID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"IMCSNeedsIndexing": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"IMCSIndexTotalRecords": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"IMCSIdxProductVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"IMCSBypassIndexVersionCheckV2": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.bird": DomainPrefs{
			"optimize-storage": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"icloud-drive.account-migration-status.294735135": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.xpc.activity2": DomainPrefs{
			"ProductBuildVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.iPod": DomainPrefs{
			"com.apple.PreferenceSync.ExcludeAllSyncKeys": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.newscore2": DomainPrefs{
			"report_concern_user_id": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"provider_user_id": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"instance_identifier": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.cloudd": DomainPrefs{
			"com.apple.private.cloudkit.shouldUseGeneratedDeviceID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.driver.AppleBluetoothMultitouch.trackpad": DomainPrefs{
			"TrackpadFourFingerHorizSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadPinch": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"USBMouseStopsTrackpad": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadFourFingerVertSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadRotate": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadHorizScroll": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadTwoFingerDoubleTapGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"TrackpadMomentumScroll": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadFourFingerPinchGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"UserPreferences": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Clicking": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadThreeFingerTapGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"HIDScrollZoomModifierMask": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadScroll": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"DragLock": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadFiveFingerPinchGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadThreeFingerHorizSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadThreeFingerVertSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadTwoFingerFromRightEdgeSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ForceSuppressed": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"TrackpadCornerSecondaryClick": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Dragging": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadHandResting": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadRightClick": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadThreeFingerDrag": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.stocks.account": DomainPrefs{
			"deleteOnNextLaunch": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.preference.trackpad": DomainPrefs{
			"ForceClickSavedState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.SetupAssistant": DomainPrefs{
			"PreviousSystemVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"DidSeeSyncSetup2": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"SkipExpressSettingsUpdating": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"LastSeenDiagnosticsProductVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"SkipFirstLoginOptimization": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"LastSeenCloudProductVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"DidSeeCloudSetup": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"PreviousBuildVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"DidSeeApplePaySetup": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"DidSeeTrueTonePrivacy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"DidSeeScreenTime": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"MiniBuddyShouldLaunchToResumeSetup": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"DidSeeTouchIDSetup": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"NSAddServicesToContextMenus": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"DidSeeActivationLock": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"DidSeeSyncSetup": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"DidSeeAvatarSetup": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"MiniBuddyLaunchReason": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"DidSeeSiriSetup": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"DidSeeTrueTone": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"DidSeePrivacy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"LastPrivacyBundleVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"LastSeenSiriProductVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"DidSeeAppStore": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"LastPreLoginTasksPerformedVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"DidSeeAppearanceSetup": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"LastSeenSyncProductVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"LastSeeniCloudStorageServicesProductVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"LastPreLoginTasksPerformedBuild": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"MiniBuddyLaunchedPostMigration": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"DidSeeAccessibility": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"LastSeenBuddyBuildVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"DidSeeiCloudLoginForStorageServices": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.ScriptEditor2": DomainPrefs{
			"OSAStandardAdditions ChooseApplication Bounds": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.preferences.extensions.ServicesWithUI": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.MobileSMS": DomainPrefs{
			"TextSize": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"SyndicationOnboardingVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"LegacyAppSidebarPersistedWidth": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"QuickSaveHasBeenUsedBefore": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"CatalystPreferenceMigrationVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"PlaySoundsKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"BusinessChatPrivacyPageDisplayed": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"sForceUnknownFilteringCompleted": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"CKLastSelectedItemIdentifier": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"kCKMediaObjectManagerDefaultsOSVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"CatalystCustomFontMigrationVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"NotifyAboutMessagesFromUnknownContacts": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"CKMigratedAutoSpamReports26375262": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"SidebarPersistedWidth": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ReadReceiptSettingsConfirmed": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"sForceSMSSpamFilteringCompleted": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"KeepMessageForDays": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"PendingCleared": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"TextFontSize": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.preferences.extensions.ShareMenu": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.mmcs": DomainPrefs{
			"report.LastFailedCheckDate": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"report.sha256": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"report.LastSuccessfulCheckDate": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"report.TTL": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.appstored": DomainPrefs{
			"BadgeCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ArcadeSubscriptionState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ArcadePayoutDeviceID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"WelcomeNotificationExcludedFromSample": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ArcadeDeviceID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"ArcadeDeviceGUID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"LastOSBuildVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"WelcomeNotificationLastAppStoreConnectionProductVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.PubSubAgent": DomainPrefs{
			"TigerMigrationCompleted": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.commcenter.callservices": DomainPrefs{
			"last.known.icloud.id": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"associated.account": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.print.PrinterProxy": DomainPrefs{
			"IK_Scanner_downloadURL": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"IK_Scanner_selectedTag": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"IK_Scanner_PostProcessApplication": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.amsengagementd": DomainPrefs{
			"AMSMetricsIdentifierUserRecordName": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"AMSMetricsIdentifierZoneCreated": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AMSMetricsIdentifierZoneSubscriptionCreated": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.iChat": DomainPrefs{
			"SaveConversationsOnClose": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"PMPrintingExpandedStateForPrint2": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"UnifiedChatWindowControllerSelectionGUIDSet": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"DidCheckForDuplicateChats": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"HasPromptedSMSRelay": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"HasPromptediMessageFTU": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"DidRegenerateGroupID63841559": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"messageTracerSMSSent": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"messageTracerSubmitDate": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"BuddyPictureSetToGenericByUser": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"messageTracerIMessagesSent": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"PlaySoundsKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"messageTracerSMSUsed": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"messageTracerNumUpgradeOffers": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ChatWindowControllerUnifiedIsVisible": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"messageTracerMessagesSent": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"WebIconDatabaseEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"DidMarkGroupPhotosAsUnpurgeable": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"messageTracerUpgradesAccepted": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"messageTracerIMessagesReceived": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"DaemonConnectionWaitTime": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"NotifyAboutMessagesFromUnknownContacts": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"KeepMessagesVersionID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"messageTracerCharactersSent": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"messageTracerMessagesReceived": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"messageTracerUpgradesDeclined": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"messageTracerSMSReceived": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"KeepMessageForDays": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"UnifiedChatListViewControllerAutomaticallySortsChats": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"messageTracerIMessageUsed": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"DidMigratePersonCentricIDs": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"LastIMDNotificationPostedDate": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ChatWindowControllerUnifiedChatListWidth": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"LastFailedMessageIMDNotificationPostedDate": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"SODefaultTranscriptName": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"imageBrowser.disableOpenGL": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.TextEdit": DomainPrefs{
			"PlainTextEncoding": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"RichText": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"PlainTextEncodingForWrite": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.HIToolbox": DomainPrefs{
			"AppleFnUsageType": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"AppleCurrentKeyboardLayoutInputSourceID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"AppleDictationAutoEnable": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.preferences.extensions.FinderSync": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.coreservices.uiagent": DomainPrefs{
			"CSUIHaveCenteredGatekeeperProgressWindow": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"CSUIHasSafariBeenLaunched": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"CSUIRecommendSafariBackOffInterval": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"CSUILastOSVersionWhereSafariRecommendationWasMade": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.CharacterPicker": DomainPrefs{
			"version": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.CharacterPaletteIM": DomainPrefs{
			"CVFontSize": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"CVStartAsLargeWindow": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.assistant.backedup": DomainPrefs{
			"Hands Free Mode": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Session Language": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Cloud Sync Enabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Cloud Sync User ID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"MultiUser VoiceIdentification Enabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.icloud.fmfd.notbackedup": DomainPrefs{
			"kKeychainUtilMigrationVersionKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.security.cloudkeychainproxy3.keysToRegister": DomainPrefs{
			"EnsurePeerRegistration": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"DSID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"KeyAccountUUID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.mediaaccessibility.public": DomainPrefs{
			"MACaptionDisplayType": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.speech.recognition.AppleSpeechRecognition.prefs": DomainPrefs{
			"DictationIMPresentedOfflineUpgradeSuggestion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"DictationIMCommandsWindowIsOpen": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"DictationIMCorrectionCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"DictationIMUpgradedTo10_16": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"DictationIMUpgradedTo10_15": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"DictationIMPlaySoundUponRecognition": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"DictationIMLocaleIdentifier": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"DictationIMMessageTracesSinceLastReport": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"DictationIMUseOnlyOfflineDictation": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.loginwindow": DomainPrefs{
			"TALLogoutSavesState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"MiniBuddyLaunch": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"oneTimeSSMigrationComplete": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.preferences.softwareupdate": DomainPrefs{
			"LatestMajorOSSeenByUserBundleIdentifier": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.accessibility.universalAccessAuthWarn": DomainPrefs{
			"ThirdPartyCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.internetconnect": DomainPrefs{
			"ServiceID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.TextInputMenu": DomainPrefs{
			"visible": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.messages.nicknames": DomainPrefs{
			"NicknameAppleIDAndiCloudAccountMatchAndExist": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"MeCardPendingNicknamesVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"MeCardSharingVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"MeCardSharingEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"MeCardSharingImageForkedFromMeCard": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"MeCardWhitelistBlacklistNicknamesVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"NicknameInfoRequestedFromPeers": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"MeCardSharingAudience": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NicknameScrutinyDoNotHandle": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"ReuploadLocalNicknamesVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.ProblemReporter": DomainPrefs{
			"ShowComments": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ShowDetails": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.gamed": DomainPrefs{
			"natType": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"GKStoreFrontIdentifier": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"GKActiveEnvironmentKeyV2": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"GKLastPushTokenPlayerID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"GKLastPushTokenEnvironment": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"GKPushEnvironmentKeyV2": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"GKLastProtocolVersionUsedKeyV2": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"GKLoginCancelled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.Safari.SandboxBroker": DomainPrefs{
			"DidMigrateDownloadFolderToSandbox": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"ShowDevelopMenu": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"DidMigrateResourcesToSandbox": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"DidMigrateDownloadMetadataToSandboxGroupContainer": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.preferences.extensions.XcodeSourceEditor": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.AppleMediaServices.notbackedup": DomainPrefs{
			"AMSDidRetrieveDeviceOffers": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AMSDidRetrieveDeviceOffersEligibility": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.imdsmsrecordstore": DomainPrefs{
			"DeleteSequenceNumber": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.iTunes": DomainPrefs{
			"com.apple.iTunes.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"doesAccountArtistListHaveSharePermission": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"dontAutomaticallySyncIPods": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"MacBuddyStoreID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"com.apple.iTunes.WebKit2ShouldPrintBackgrounds": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"com.apple.iTunes.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"com.apple.iTunes.WebKit2AsynchronousSpellCheckingEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AutomaticDeviceBackupsDisabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"disableRadio": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Store Apple ID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"debugAssertCategoriesVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"com.apple.iTunes.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"com.apple.iTunes.WebKit2EnableInheritURIQueryComponent": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"isAccountEnrolledInITunesMatch": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"log-PlayQueue-LastSelectedTab": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"isAccountAdmin": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"NSApplicationCrashOnExceptions": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"com.apple.iTunes.WebKit2DeveloperExtrasEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"DeviceBackupsDisabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"disableShareLibraryInfo": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"WirelessBuddyID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"Store DSID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"storefront": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"isAccountEnrolledInAppleMusic": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"log-push": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"debugAssertCategoriesEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.textInput.keyboardServices.textReplacement": DomainPrefs{
			"osxMajorVersForCloudKitSync": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"TRCKSyncMaxCountThreshold": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"_KSTRCloudKitMigratable": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"KSDidPushAllLocalRecordsOnce-2": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"osxMinorVersForCloudKitSync": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"internalMigrationPercent": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"KSCKSubscriptionProd-TextReplacements": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"seedMigrationPercent": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"seedMigrationPercent1": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"KSDidPushMigrationStatusOnce-2": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"gmMigrationPercent1": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"KSDidMigrateToCloudKitOnCloud": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"KSCloudKitMigrationDidComplete": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"kTRCKSyncCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"KSCKDidSetupRecordZoneProd-TextReplacements": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"iOSMajorVersForCloudKitSync": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"kDidMigrateToUUIDRecordName": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"internalMigrationPercent1": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"kDidInsertSampleShortcutForPeer": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"iOSMinorVersForCloudKitSync": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"gmMigrationPercent": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"userDidFallInMigrationAllowBatch": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"KSSampleShortcutWasImported_CK": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"KSDidPullLegacyEntriesFromPeers": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"TRCKSyncCountHalflifeInHours": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"KSLastKnownUserID-TextReplacements": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"osxMinorSubversionForCloudKitSync": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"KSLSShouldUpdateCache": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.diagnosticextensionsd": DomainPrefs{
			"directoriesCleanupDone": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.Messages": DomainPrefs{
			"hasPerformedNewDeviceBringUpSync": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"hasTriedToEnableCKAndSyncAfterFirstImagentConnection": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.NewDeviceOutreach": DomainPrefs{
			"InitialOutreachActivityScheduled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"CachedWarrantyValidityDuration": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"PostUpgradeActivityCompleted": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"PostUpgradeOSVersionKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"CachedWarrantyLocale": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"CachedWarrantyVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.icloud.searchpartyuseragent": DomainPrefs{
			"PersistenceLayerVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"LastOSLaunchVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"lastLaunchBootSessionUUID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"lastLaunchLocale": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"FMIPStateManager.fmipState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.cloudpaird": DomainPrefs{
			"MagicCloudPairingAccessorySubscriptionManateeID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"MagicCloudPairingManateeUpgradedAccount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"MagicCloudPairingManateeUpgraded": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"MagicCloudPairingProtectedAccessorySubscription": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"UploadedHSA2KeysForLocalDevice": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"MagicCloudPairingAccessorySubscriptionID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"MagicCloudPairingMasterEncryptionPreferences": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"SecuredAccessoryZoneSubscription": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"MagicCloudPairingMasterSubscriptionManateeID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"MagicCloudPairingMasterSubscriptionID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"SecuredZoneSubscription": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"MagicCloudPairingAccessoryEncryptionPreferences": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.syncserver": DomainPrefs{
			"SyncServicesResetWorldRunOnce": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"RunCompletelyDisabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.screencaptureui": DomainPrefs{
			"com.apple.AnnotationKit.userDefaultsVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"com.apple.AnnotationKit.arrowHeadStyle": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.AnnotationKit.strokeIsDashed": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.AnnotationKit.hasShadow": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.AnnotationKit.strokeWidth": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.AnnotationKit.highlightStyle": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"toolbarOrigin": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"com.apple.AnnotationKit.brushStyle": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.desktopservices": DomainPrefs{
			"DSDontWriteNetworkStores": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.security.pboxd": DomainPrefs{
			"ILMediaBrowserMediaType": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ILMediaBrowserSelectedGroupIndex1": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ILMediaBrowserSelectedGroup1": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.keychainaccess": DomainPrefs{
			"Keychain Item List Sorting": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Keychain Item List Sort Descending": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Last Selected Category": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"VPNSSItemsChecked": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"Item Preview Closed": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"Last Selected Keychain": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"Keychains List Closed": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.imessage": DomainPrefs{
			"PreviewTranscodingWasInitializedToDefaultValue": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"PreviewTranscodingQualityOnWiFiWasInitializedToDefaultValue": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.Safari.SafeBrowsing": DomainPrefs{
			"HasMigratedSafeBrowsingEnabledDefaults": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.AppleMediaServices": DomainPrefs{
			"AMSIncludeFullResponseInHARLogging": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"AMSShowSandboxAccountUI": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"AMSDeviceBiometricsState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"AMSUserDefaultsincludeFullResponseInHARLogging": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"AMSMigratedToNewAccountStore": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"AMSUserDefaultsIncludeFullRequestInHARLogging": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"AMSLogHARData": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"AMSLastMigratedBuildVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"AMSDidSetUpServerDataCache": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AMSIncludeFullRequestInHARLogging": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"AMSMigratedStorageToDefaultsForNonAMSInternal": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.PhotoBooth": DomainPrefs{
			"AVDesktopPlaybackControlsControllerShowsDurationInsteadOfTimeRemainingDefaultsKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"PBLaunchedAtLeastOnceOnLion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"UserCameraUniqueIDPref": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.MobileSMSPreview": DomainPrefs{
			"IMPreviewGenerationMaxPxWidth": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"IMPreviewGenerationMinHeight": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"IMPreviewGenerationScreenScale": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"IMPreviewGenerationMinWidth": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.ActivityMonitor": DomainPrefs{
			"OpenMainWindow": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ShowCategoryAppsinLast12Hours": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"SelectedTab": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ShowCategory": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"HasSavedState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.airplay": DomainPrefs{
			"showInMenuBarIfPresent": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.messages.facetime": DomainPrefs{
			"FaceTimeTwoTimeCallthroughEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.remindd": DomainPrefs{
			"CloudKitAccountStatus": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"preferredDefaultListObjectIDUrl": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ThrottlingPolicyCurrentBatchCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"isDatabaseMigrated": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"preferredDefaultListID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"cloudKitSchemaCatchUpSyncLastSuccessBuildVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"ThrottlingPolicyCurrentLevelIndex": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"spotlightIndexVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"useExtraneousAlarmBackOffThrottleInterval": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"CloudConfigurationPath": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"lastDatabaseMigrationSystemBuildVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.ncprefs": DomainPrefs{
			"content_visibility": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"sort_order": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.itunesstored": DomainPrefs{
			"AccountsNotificationPlugin.ActiveStorefrontIdentifier": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.driver.AppleHIDMouse": DomainPrefs{
			"Button4Force": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ScrollH": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ButtonDominance": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Button4": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Button4Click": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Button3": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Button1": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Button2": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ScrollSSize": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ScrollV": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ScrollS": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.preferences.extensions.CredentialProvider": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.itunescloud": DomainPrefs{
			"ICDefaultsKeyLastActiveAccountDSID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.stockholm": DomainPrefs{
			"RemoteAdminV2": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"PKWalletShouldAutomaticallyRegisterKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.Siri": DomainPrefs{
			"VoiceTriggerUserEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"StatusMenuVisible": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.sharingd": DomainPrefs{
			"AirDropRandomHashUUIDKey4": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"OneTimeAirDropReset2": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"OneTimeAirDropReset": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"HashManager-StoredDatabaseVersionKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"AirDropRandomHashUUIDKey2": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"DiscoverableMode": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"AutoUnlockPresentedWiFiError": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AutoUnlockPresentedBluetoothError": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AirDropRandomHashUUIDKey3": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"AutoUnlockWatchCurrentlyInList": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AutoUnlockWatchExistedInUnlockList": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AirDropRandomHashUUIDKey1": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.AppStoreComponents": DomainPrefs{
			"ASCLocaleID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.DFSLoginPlugin": DomainPrefs{
			"WindowLeft": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"WindowTop": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.iCal": DomainPrefs{
			"BirthdayEventsGenerationLastLocale": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"last calendar view description": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"BirthdayEventsGenerationVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"icaluuid": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"CalendarSidebarShown": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"AvailabilityShowTwentyFourHours": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"CalForceTruthFileRestoreHashKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"last selected calendar list item": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"CalDefaultCalendar": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"CalAgentNS_Preference_DefaultReminderCalendar": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"display birthdays calendar": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NSEventConcurrentProcessingEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"BirthdayEventsGenerationAttemptsToResetKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"TimeZone support enabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Show Week Numbers": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"iCal version": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"first shown minute of day": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"CalendarSidebarView": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"RemindersLastMigratedSystemVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"LastReminderMigrationCleanupVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AllDayAreaHeight": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"lastViewsTimeZone": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"privacyPaneHasBeenAcknowledgedVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"CalendarSidebarWidth": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"last minute of day time range": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"first minute of day time range": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"CalendarListMiniMonthVisibleMonths": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"IncludeDebugMenu": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"WR_DONT_ASK_FOR_DEFAULT": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"kLastABCReportTimeKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NSDontMakeMainWindowKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.assistant": DomainPrefs{
			"Experiment Identifierinvocation_feedback_experiment": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"Has Set Up MultiUser Shared Record Subscription": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"Initial Interstitial Delay for VoiceTrigger One-Shot without Audio Ducking": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"Initial Interstitial Delay - stark": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"MultiUser Shared Data Needs Sync": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"Initial Interstitial Delay": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"Myriad Device Delay": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"Myriad Device Class": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"Myriad Device Adjust": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"Has Set Up Account Status Subscription": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"Voice Trigger Needs Sync": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"Myriad Device Trump Delay": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"Experiment Identifier": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"Flush Session Tickets Cache": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"Server Has Provisioned Myriad": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"Myriad VTEndtimeDistanceThreshold": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"Has Set Up Key Value Subscription": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"Borealis Education Header Display Count": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"Experiment Identifiersiri_vox_tap_to_siri_behavior_experiment": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"Has Set Up Voice Trigger Subscription": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"Last Known Analytics Store State": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"Experiment Identifiersiri_vox_sounds_experiment": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"Manual Endpointing Threshold": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.dt.Xcode": DomainPrefs{
			"IBPreferencesMigrated": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"IBAppliesAutoResizingRulesWhileResizing": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"XCFontAndColorCurrentTheme": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"IDEKeyBindingCurrentPreferenceSet": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"IDESourceControlPreferencesVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"DVTFontAndColorLastUpdatedToolsVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"DVTDownloadableAutomaticUpdate": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"IDEUserWantsToEnableDeveloperSystemPolicyMode": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.AMPLibraryAgent": DomainPrefs{
			"iTunes-media-folder-url": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"debugAssertCategoriesVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"have-shown-cloud-UI": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"storeSupportsUPP": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"updateLevel": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"ITUserPrefsMigrated": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"books-migrated": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"storeSupportsPreviousPurchases": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"persistentUserID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"storeSupportsCloudPurchases": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"imported-media-domains": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"books-persistent-id": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"storeMediaTypeFlags": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"podcasts-migrated": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"storeSupportsPasswordSettings": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"imported-media-domains-modification-date": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"show-home-sharing-turned-on-in-iTunes-warning": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"devices-persistent-id": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"debugAssertCategoriesEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"podcasts-persistent-id": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.wifi.keychain-format": DomainPrefs{
			"Version": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.facetime.bag": DomainPrefs{
			"URL": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"CacheTime": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.TelephonyUtilities": DomainPrefs{
			"GondolaLastAccountsChangedState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"GondolaGeneratedIdentifier": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"registeredProvidersVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"CachedVCCaps": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"relayCallingDisabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"GondolaLatestVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"GondolaSyncedVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"FaceTimePhotosEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.findmy.fmfcore.notbackedup": DomainPrefs{
			"frontMostWindow": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"publicAPSToken": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"FMFGarbageCollectorIdentityKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"FMFLimitedPrecisionPrefKey.limitedPrecision": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"windowVisible": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.AppStore": DomainPrefs{
			"__uniquePageGroupID-9.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"__uniquePageGroupID-12.WebKit2DeveloperExtrasEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"mostRecentTabIdentifier": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AutoPlayVideoSetting": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.AppStore.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"appStoreBadgeCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"__uniquePageGroupID-1.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"UserSetAutoPlayVideoSetting": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"lastBootstrapTimeZone": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"__uniquePageGroupID-9.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ASAcknowledgedOnboardingVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"__uniquePageGroupID-12.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"__uniquePageGroupID-9.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"com.apple.AppStore.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"__uniquePageGroupID-1.WebKit2DeveloperExtrasEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"__uniquePageGroupID-9.WebKit2DeveloperExtrasEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AVDesktopPlaybackControlsControllerShowsDurationInsteadOfTimeRemainingDefaultsKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"__uniquePageGroupID-12.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"__uniquePageGroupID-12.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"com.apple.AppStore.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"com.apple.AppStore.WebKit2DeveloperExtrasEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"__uniquePageGroupID-1.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"__uniquePageGroupID-1.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.AdPlatforms": DomainPrefs{
			"AppStorePAAvailable": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"LatestPAVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"acknowledgedPersonalizedAdsVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.keyboard": DomainPrefs{
			"KeyboardWordOrSentenceTrackingForPFL": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.assistant.support": DomainPrefs{
			"Siri Data Sharing Opt-In Status": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Assistant Enabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Dictation Enabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.preferences.extensions.QuickLook": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.internetconfigspec": DomainPrefs{
			"WWWHomePage": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.icloud.fmfd": DomainPrefs{
			"numberOfFriendsFollowersKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"kFMFDStoredDataVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"storedDSID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"lazyInitTimeSecsStoredKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"kFMFDlastLoggedInPrsId": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.newscore": DomainPrefs{
			"notificationEnableAssetPrefetching": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.news.notification_receipt_event_endpoint": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"news_url_resolution_subscription_status": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"com.apple.news.default_event_endpoint": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"notificationAssetPrefetchingRequiresWatch": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"notificationArticleDiversificationSimilarityExpectationStart": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"force_refresh_user_segmentation": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"notificationArticleDiversificationUniquePublisherExpectationYIntercept": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"notificationArticleDiversificationUniquePublisherExpectationSlope": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"news_carplay_is_enabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.newsd.download.hasUnfinishedWork": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"notificationArticleDiversificationSimilarityExpectationEnd": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"use_parsec_results_for_widget": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.AddressBook": DomainPrefs{
			"ABBirthDayVisible": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ABMetaDataChangeCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ABTextSizeIncrement": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ABVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"ABLastImportShown": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ABDefaultSourceID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NSPreferencesSelectedIndex": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AB21vCardEncoding": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NSPreferencesContentSize": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ABPrivateVCardFieldsEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.SpeechRecognitionCore": DomainPrefs{
			"AllowAudioDucking": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.commerce": DomainPrefs{
			"LastUpdateNotificationOSMajorVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"PurchasesInflight": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.networkresolutiond": DomainPrefs{
			"_networkDevices": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.TV": DomainPrefs{
			"automaticallyDownloadArtwork": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"userWantsPlaybackNotifications": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"gridViewSearchDoesNotSwitchView": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"persistentUserID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"miniplayerUserSetHeight": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"storeMediaTypeFlags": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"dontAskForAllPlaylistItemRemoval": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"checkedHLSKeysTime": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"storeSupportsCloudPurchases": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"debugAssertCategoriesVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"com.apple.TV.WebKit2EnableInheritURIQueryComponent": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"cddbPrefsOK": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"com.apple.TV.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"miniplayerVPos": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"controllableInterfaceGUID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"automaticallyDeleteVideoAssetsAfterWatching": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"videoWindowVPos": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"miniplayerHPos": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"playbackIsFullscreen": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"storeSupportsPasswordSettings": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"com.apple.TV.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"haveAskedToCheckForAvailableDownloads": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"videoWindowHPos": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"miniplayerSnapMode": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"storeSupportsUPP": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"showWelcomeScreenState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"videoWindowVSize": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"updateLevel": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"ITUserPrefsMigrated": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"miniplayerWidthInDesignCoords": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"storeSupportsPreviousPurchases": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"tabViewMode": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"preloadFilesIntoMemory": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"userMaxConcurrentDownloads": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NSApplicationCrashOnExceptions": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"videoWindowHSize": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"checkForAvailableDownloads": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"miniPlayerQueueVisible": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"debugAssertCategoriesEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"firstLaunchShowWelcomeScreenState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"com.apple.TV.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"musicVideoPlaybackLocation": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"miniPlayerLargeArtVisible": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"userHasActivatedFullScreenVisualizer": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"com.apple.TV.WebKit2AsynchronousSpellCheckingEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"com.apple.TV.WebKit2ShouldPrintBackgrounds": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.preference.general": DomainPrefs{
			"version": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.menuextra.battery": DomainPrefs{
			"ShowPercent": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.imdpersistence.IMDPersistenceAgent": DomainPrefs{
			"LogAllIOErrors": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.SafariServices": DomainPrefs{
			"SearchProviderIdentifierMigratedToSystemPreference": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.siri.shortcuts": DomainPrefs{
			"WFServicesShortcutsMigrated": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"WFDidUnconflictShortcuts": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"LegacyShortcutsZoneSubscriptionUnsubscribed": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"WFDefaultShortcutsVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"WFLastSyncedFlagsHash": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.SystemProfiler": DomainPrefs{
			"PrefsVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"SPLastDocumentsSize": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.remindd.babysitter": DomainPrefs{
			"LastSystemVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.internal.ck": DomainPrefs{
			"DictationOnDeviceSamplingDisabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"DictationSamplingRates": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"DefaultCacheKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"DefaultWarmupScriptsExtension": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"WarmupScriptIdentifiers": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"HasSetUpRecordZoneSubscription": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ByteCodeCacheEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"DefaultBootTimeUpdateScripts": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"WarmupModularScriptIdentifiers": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"DisableFBFForUEI": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.internetconfigpriv": DomainPrefs{
			"WWWHomePage": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.dock": DomainPrefs{
			"wvous-br-corner": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"showAppExposeGestureEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"autohide": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"tilesize": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"no-bouncing": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"region": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"showhidden": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"launchanim": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"showMissionControlGestureEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"wvous-tr-modifier": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"magnification": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"largesize": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"mod-count": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"trash-full": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"minimize-to-application": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"loc": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"expose-animation-duration": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"wvous-tr-corner": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"autohide-delay": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.AdLib": DomainPrefs{
			"adprivacydMaxSegmentSendInterval": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"personalizedAdsMigrated": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"partiality-segment": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"CKDPIDSyncState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"forceLimitAdTracking": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"adprivacydSegmentInterval": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"allowIdentifierForAdvertising": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"allowApplePersonalizedAdvertising": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.amp.mediasharingd": DomainPrefs{
			"public-sharing-enabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"home-sharing-sequence-id": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"home-sharing-group-id": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"home-sharing-computer-id": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"photo-sharing-enabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"shared-library-id": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"shared-library-name": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"public-sharing-is-protected": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"shared-library-machine-id": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"home-sharing-enabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.systempreferences": DomainPrefs{
			"com.apple.preferences.sharing.selectedservice": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ThirdPartyCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"SoundTab": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"com.apple.dtandsspref.lastselectedtab": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NSFullScreenMenuItemEverywhere": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"trackpad.lastselectedtab": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"ShowAllMode": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ICQ_iCloudFirstRunNotificationShown": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"DisableAutoLoginButtonIsHidden": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"com.apple.preferences.accounts.outline.usersparent": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"mouse.lastselectedtab": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NSQuitAlwaysKeepsWindows": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.SecurityPref.Privacy.LastSourceSelected": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"SecurityPrefTab": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"com.apple.preference.keyboard.lastselectedtab": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"com.apple.datetimepref.lastselectedtab": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.UIKit": DomainPrefs{
			"hasAccessibilityBeenMigrated": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.SafariBookmarksSyncAgent": DomainPrefs{
			"CloudBookmarksSupplementalClientIdentifier": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"NewestLaunchedSafariBookmarksSyncAgentVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.VoiceOverTraining": DomainPrefs{
			"doNotShowSplashScreen": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.newsd": DomainPrefs{
			"FCAppConfigurationBundleShortVersionKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.itunescloud.daemon": DomainPrefs{
			"ICDDefaultsKeyKnownActiveAccountDSID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.Image_Capture": DomainPrefs{
			"IK_lastUsedDeviceUUID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"IK_Camera_selectedTag": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"IK_lastUsedDeviceIsRemote": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"IK_Scanner_downloadURL": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"IKC_sort_ascending": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.imagekit.cameraviewmode": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"IK_prefsVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"IK_Camera_selectedPathType": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"IKC_sort_key": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"IK_Accessory_selectedTag": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"IK_Scanner_PostProcessApplication": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"IK_Camera_preferPostPocessingApp": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"IK_lastUsedDeviceName": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.preferences.extensions.SharedLinks": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.driver.AppleBluetoothMultitouch.mouse": DomainPrefs{
			"MouseOneFingerDoubleTapGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"MouseTwoFingerHorizSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"MouseVerticalScroll": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"UserPreferences": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"MouseHorizontalScroll": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"MouseButtonDivision": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"MouseMomentumScroll": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"MouseButtonMode": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"MouseTwoFingerDoubleTapGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.dock.external.extra.x86_64": DomainPrefs{
			"ShowSolidarity": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.commcenter.data": DomainPrefs{
			"pw_ver": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.imagecapture": DomainPrefs{
			"loggingLevel": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.notificationcenterui": DomainPrefs{
			"bannerTime": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.pipagent": DomainPrefs{
			"Size": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.AMPDevicesAgent": DomainPrefs{
			"debugAssertCategoriesEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"persistentUserID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"dontAutomaticallySyncIPods": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"updateLevel": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"storeMediaTypeFlags": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"AutomaticDeviceBackupsDisabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"storeSupportsCloudPurchases": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"debugAssertCategoriesVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"ITUserPrefsMigrated": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"userMaxConcurrentDownloads": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.MobileSMS.CKDNDList": DomainPrefs{
			"CatalystDNDMigrationVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"CKDNDMigrationKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.protectedcloudstorage.protectedcloudkeysyncing": DomainPrefs{
			"registrySyncIdentifier": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.VoiceOver4.local": DomainPrefs{
			"SCRCUserDefaultsAllowAirPlay": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"SCRCUserDefaultsUnplannedShutdownCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"SCRCUserDefaultsPlannedShutdownSuccessful": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.suggestions": DomainPrefs{
			"SuggestionsAllowGeocode": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"spToLearnMigrationPerformed": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"findToShowMigrationPerformed": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.NetworkUtility": DomainPrefs{
			"NUPingCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NUTracerouteAddress": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NUPortScanEnd": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NUSelectedTabItem": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NULookupAddress": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NUFingerPerson": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NUPortScanAddress": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NUPingChoice": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NUPingAddress": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NUNetstatChoice": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NUPortScanStart": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NUWhoisAddress": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NUWhoisSelectedServer": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"NUPortScanRange": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.siri.VoiceShortcuts": DomainPrefs{
			"VCLSDatabaseUUIDKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"VCLSDataSequenceKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.Safari": DomainPrefs{
			"ResetCloudHistory": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"WebKitHistoryItemLimit": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"WebKitInitialTimedLayoutDelay": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ShowFullURLInSmartSearchField": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"com.apple.Safari.ContentPageGroupIdentifier.WebKit2DeveloperExtrasEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"cloudBookmarksMigrationEligibilityDataInvalidated": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"UniversalSearchEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"WebKitDeveloperExtrasEnabledPreferenceKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"SuppressSearchSuggestions": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"WebKitHistoryAgeInDaysLimit": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"IncludeDebugMenu": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"IncludeDevelopMenu": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"IncludeInternalDebugMenu": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.accounts": DomainPrefs{
			"ACOneTimeSMTPFixAccountSync": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"ACOneTimeLDAPFixAccountSync": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.visualintelligence": DomainPrefs{
			"bypassAvailability": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"enableBooksDomain": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"enableLandmarkDomain": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"enableArtDomain": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"firstTimeExperience": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"enableCoarseClassification": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"initialized": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"enablePetsDomain": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"enableScreenshots": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"debugUI": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"enableSafariApp": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"sendLocationInfo": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"enableAlbumsDomain": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"enableQuickLook": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"enableNatureDomain": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"enablePhotosApp": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"sendOCRText": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.accountsd": DomainPrefs{
			"com.apple.mail.searchableIndex.lastProcessedAttachmentIDKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"LastSystemVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.fileproviderd": DomainPrefs{
			"com.apple.fileproviderd.did-drop-daemon-corespotlight-index": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.screencapture": DomainPrefs{
			"location": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"style": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"captureDelay": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"last-selection-display": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"video": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.SoftwareUpdate": DomainPrefs{
			"AutoUpdateMajorOSVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.Preferences": DomainPrefs{
			"UserDictionaryImportedSinceMaintenance": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"UserDictionaryLocalPeerIdentityCurrent": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.sociallayerd.CloudKit.ckwriter": DomainPrefs{
			"accountLastKnownUserRecordID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"accountAvailable": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"writerDone": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.classroom": DomainPrefs{
			"allowClassroomLockDevice-00000000-0000-0000-0000-000000000000:0": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"allowClassroomOpenApp-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"allowClassroomScreenObservation-00000000-0000-0000-0000-000000000000:0": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"allowClassroomOpenURL-ask": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"allowClassroomScreenObservation-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"allowClassroomOpenURL": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"allowClassroomLockDevice-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"allowClassroomOpenURL-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"allowClassroomOpenApp-00000000-0000-0000-0000-000000000000:0": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"forceUnpromptedRemoteScreenObservation-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"allowClassroomOpenURL-00000000-0000-0000-0000-000000000000:0": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"forceUnpromptedRemoteScreenObservation-00000000-0000-0000-0000-000000000000:0": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"allowClassroomOpenApp": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"allowClassroomOpenApp-ask": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.CloudKit": DomainPrefs{
			"AccountInfoValidationCounter": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.findmy.fmipcore.notbackedup": DomainPrefs{
			"publicAPSToken": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"itemLearnMoreURL": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"frontMostWindow": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"windowVisible": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"deviceImageVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"FMIPGarbageCollectorIdentityKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"FMIPLimitedPrecisionPrefKey.limitedPrecision": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.tipsd": DomainPrefs{
			"TPSWelcomeNotificationViewedVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"TPSTipsAppInstalled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"TPSOfflineLastProcessedRemoteContentIdentifier": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"DeliveryInfoVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"TPSLastMajorVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"ConsecutiveNotificationsCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"TPSWelcomeNotificationReminderState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.CoreGraphics": DomainPrefs{
			"DisplayUseForcedGray": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"DisplayUseInvertedPolarity": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.DirectoryUtility": DomainPrefs{
			"FirstLaunch": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.AvatarUI.Staryu": DomainPrefs{
			"AVTAvatarHasDisplayedAnimojiSplashScreen": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AVTAvatarHasDisplayedCameraEffectsSplashScreen": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AVTAvatarHasDisplayedPaddleView": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AVTAvatarHasDisplayedSplashScreen": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AVTAvatarUILastCacheVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.networkConnect": DomainPrefs{
			"VPNShowTime": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.UserAccountUpdater": DomainPrefs{
			"com.apple.HidLibraryFolderAlready": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.madrid": DomainPrefs{
			"IMCloudKitStartingInitialSync": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"CloudKitIsSyncing": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"CloudKitIsRemovedFromBackup": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"CloudKitIsEligibleForTruthZone": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"CloudKitEligibleToToggleMiCSwitch": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"IMCloudKitSyncControllerSyncTypeKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"CloudKitCheckedMiCSwitchEligibilityOnImagentLaunch": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"createdChatZone": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"hasCompletedInitialCKSync": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"CloudKitSyncingEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"IMCloudKitSyncControllerSyncRecordTypeKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"initialSyncRecordHasBeenWritten": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"IMCloudKitSyncControllerSyncStateKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"IMCloudKitStartingPeriodicSync": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"IMCloudKitSyncPaused": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"RequestPriorityRamp": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"IMCloudKitAccountStatusKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"IMCloudKitStartingEnabledSettingChange": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"IMCKFinishedFetchingAttachmentsFromCloudKit": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"CKMOCAccountsMatch": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"IMCloudKitStartingDisableDevices": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"oneTimeMOCAccountCheckV2": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"enableCKSyncingV2": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.rapport": DomainPrefs{
			"familySyncedName": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.menuextra.clock": DomainPrefs{
			"DateFormat": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.imessage.bag": DomainPrefs{
			"URL": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"CacheTime": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.sms": DomainPrefs{
			"hasBeenApprovedForSMSRelay": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.Automator": DomainPrefs{
			"NSSplitView AMDocumentMinor Expanded Position": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"NSSplitView AMLibraryActionsMajor Expanded Position": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AMLogTabViewSelectedIndex": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.weather.internal": DomainPrefs{
			"userId": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.FaceTime": DomainPrefs{
			"DidMigrateToSixteenByNineAspectRatio": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"kPromptEnableReadRecipts": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"CustomRingtone": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"PhoneNumberUpgradeShown": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"PreferredVideoDeviceUID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.AssistiveControl.virtualKeyboard": DomainPrefs{
			"SCLaunchedAsSlave": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"PanelPosition": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"DesiredPanelWindowPosition": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.Accessibility": DomainPrefs{
			"InvertColorsEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"GrayscaleDisplay": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"AXSClassicInvertColorsPreference": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"KeyRepeatInterval": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"CommandAndControlEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ApplicationAccessibilityEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"DarkenSystemColors": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"FullKeyboardAccessFocusRingEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"GenericAccessibilityClientEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"AccessibilityEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"KeyRepeatEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"KeyRepeatDelay": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ZoomTouchEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"DifferentiateWithoutColor": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"FullKeyboardAccessEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"BrailleInputDeviceConnected": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"AutomationEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ReduceMotionEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"SpeakThisEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"EnhancedBackgroundContrastEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"VoiceOverTouchEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.ipTelephony": DomainPrefs{
			"ImsLoggingEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.ColorSyncCalibrator": DomainPrefs{
			"722524374X-shareMode": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"722524374X-targetTemp": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"722524374X-name": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"722524374X-expertMode": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.StatusKitAgent": DomainPrefs{
			"reauthCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.AppleMultitouchTrackpad": DomainPrefs{
			"TrackpadThreeFingerDrag": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadPinch": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"USBMouseStopsTrackpad": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadFourFingerVertSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ActuationStrength": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ActuateDetents": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"FirstClickThreshold": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"SecondClickThreshold": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadTwoFingerDoubleTapGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadHorizScroll": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadMomentumScroll": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadFourFingerPinchGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadRotate": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Clicking": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadThreeFingerTapGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadThreeFingerHorizSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadScroll": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"DragLock": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadFiveFingerPinchGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"ForceSuppressed": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"TrackpadThreeFingerVertSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadTwoFingerFromRightEdgeSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"UserPreferences": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"HIDScrollZoomModifierMask": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadHandResting": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadRightClick": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Dragging": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadFourFingerHorizSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"TrackpadCornerSecondaryClick": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.ImageCaptureService": DomainPrefs{
			"IK_Accessory_selectedTag": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.screensaver": DomainPrefs{
			"tokenRemovalAction": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.AppleMultitouchMouse": DomainPrefs{
			"MouseOneFingerDoubleTapGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"MouseTwoFingerHorizSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"MouseVerticalScroll": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"UserPreferences": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"MouseHorizontalScroll": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"MouseButtonDivision": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"MouseMomentumScroll": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"MouseButtonMode": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"MouseTwoFingerDoubleTapGesture": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.universalaccessAuthWarning": DomainPrefs{
			"2::/Applications/Logi Options.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::com.logitech.manager.daemon": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Microsoft Teams (work or school).app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::com.lightpillar.Mosaic-setapp": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"1::/Applications/CloudBerry Backup.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::com.getcleanshot.app-setapp": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::/Applications/Setapp/Mosaic.app/Contents/MacOS/Mosaic": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"/Applications/Clipy.app/Contents/MacOS/Clipy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/xScope.app/Contents/MacOS/xScope": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Microsoft Teams.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"1::/Users/mikeschinkel/Library/Application Support/JetBrains/Toolbox/apps/Goland/ch-0/232.8660.185/GoLand.app/Contents/MacOS/goland": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"1::/Applications": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::com.microsoft.teams2": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Discord.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Microsoft Edge.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::/Applications/Logi Options.app/Contents/Support/LogiMgrDaemon.app/Contents/MacOS/LogiMgrDaemon": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Skitch.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/ScreenFlow.app/Contents/Library/LoginItems/ScreenFlowHelper.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"/Applications/Alfred 4.app/Contents/MacOS/Alfred": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Toast 19 Pro.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"1::com.vmware.fusion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::com.lastpass.lastpassmacdesktop": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"com.clipy-app.Clipy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::com.google.Chrome": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"1::/Users/mikeschinkel/Library/Application Support/JetBrains/Toolbox/apps/Goland/ch-0/232.8660.185/GoLand.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"1::/Applications/Microsoft Teams.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/ScreenFlow.app/Contents/Library/LoginItems/ScreenFlowHelper.app/Contents/MacOS/ScreenFlowHelper": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Microsoft Teams (work or school).app/Contents/MacOS/MSTeams": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/RingCentral.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::com.amazon.Amazon-Chime": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"com.mutedeck.mac": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Parallels Toolbox.app/Contents/Applications/Hide Menu Icons.app/Contents/MacOS/Hide Menu Icons": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::/Applications/Snappy.app/Contents/MacOS/Snappy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::me.waydabber.BetterDummy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/GoToMeeting.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Google Chrome.app/Contents/MacOS/Google Chrome": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::WSG985FR47.net.telestream.screenflowhelper": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Slack.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"/Applications/Alfred 4.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"1::/Applications/BetterDisplay.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::com.logi.cp-dev-mgr": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"com.runningwithcrayons.Alfred": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Xnapper.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/xScope.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::com.microsoft.edgemac": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"/Applications/zoom.us.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"/System/Applications/Automator.app/Contents/MacOS/Automator": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"/Applications/Clipy.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/BetterDisplay.app/Contents/MacOS/BetterDummy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Snip.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::ro.nextwave.Snappy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::com.parallels.toolbox.HideMenuIcons": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Google Chrome.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"1::/Applications/VMware Fusion-8.5.10.app/Contents/MacOS/VMware Fusion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::com.ringcentral.glip": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Acorn.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"/Applications/MuteDeck/MuteDeck.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::/Library/Application Support/Logitech.localized/LogiOptionsPlus/logioptionsplus_agent.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Timing.app/Contents/Library/LoginItems/TimingHelper.app/Contents/MacOS/TimingHelper": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::/Applications/Logi Options.app/Contents/Support/LogiMgrDaemon.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::/Applications/Setapp/Mosaic.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"1::me.waydabber.BetterDummy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::/Applications/LastPass.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::info.eurocomp.TimingHelper": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::com.amazon.Amazon-Chime": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"1::/Applications/BetterDisplay.app/Contents/MacOS/BetterDummy": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"1::/Applications/iTerm.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"1::/Applications/OBS.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"1::/Applications/HP RGS Receiver.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::J8RPQ294UB.com.skitch.SkitchHelper": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::com.logitech.Logi-Options": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::/Applications/Logi Options.app/Contents/MacOS/Logi Options": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"us.zoom.xos": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"1::/Applications/Slack.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::/Applications/Snappy.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"1::com.jetbrains.goland": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::info.eurocomp.Timing2": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Xnapper.app/Contents/MacOS/Xnapper": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/RingCentral.app/Contents/MacOS/RingCentral": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"/Applications/zoom.us.app/Contents/MacOS/zoom.us": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/BetterDisplay.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/zoom.us.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"com.apple.Automator": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Skitch.app/Contents/Library/LoginItems/J8RPQ294UB.com.skitch.SkitchHelper.app/Contents/MacOS/J8RPQ294UB.com.skitch.SkitchHelper": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Amazon Chime.app/Contents/MacOS/Amazon Chime": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Setapp/CleanShot X.app/Contents/MacOS/CleanShot X Setapp": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/ScreenFlow.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::/Applications/Amazon Chime.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"1::/Applications/VirtualBox.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"/Applications/MuteDeck/MuteDeck.app/Contents/MacOS/MuteDeck": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Timing.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::/Applications/Amazon Chime.app/Contents/MacOS/Amazon Chime": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Timing.app/Contents/MacOS/Timing": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Amazon Chime.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Skitch.app/Contents/Library/LoginItems/J8RPQ294UB.com.skitch.SkitchHelper.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::com.iconfactory.xScope": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::/Applications/LastPass.app/Contents/MacOS/LastPass": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"/System/Applications/Automator.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"1::/Applications/VMware Fusion-8.5.10.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Loom.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/var/tmp/LogMeIn Rescue - 4C61/Support-LogMeInRescue.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Setapp/CleanShot X.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::com.devuap.beautyshotapp": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Timing.app/Contents/Library/LoginItems/TimingHelper.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/var/tmp/LogMeIn Rescue - 4C61": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"3::/Applications/Parallels Toolbox.app/Contents/Applications/Hide Menu Icons.app": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"2::/Library/Application Support/Logitech.localized/LogiOptionsPlus/logioptionsplus_agent.app/Contents/MacOS/logioptionsplus_agent": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.findmy": DomainPrefs{
			"restoreState": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"coarseServiceAcknowledged_v1.0": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"onboarding_v2.0": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.FolderActionsDispatcher": DomainPrefs{
			"folderActionsEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.SocialLayer": DomainPrefs{
			"SharedWithYouEnabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.security.KCN": DomainPrefs{
			"absentCircleWithNoReason": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
			"lastCircleStatus": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"pendingApplicationReminderInterval": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.GEO": DomainPrefs{
			"DefaultsSanitizedVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.SharedWebCredentials": DomainPrefs{
			"recheckVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.suggestd": DomainPrefs{
			"DeviceIdentifier": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.CallHistorySyncHelper": DomainPrefs{
			"/Users/mikeschinkel/Library/Application Support/CallHistoryTransactions/transaction.log": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"CallHistoryDeviceCount": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"/Users/mikeschinkel/Library/Application Support/CallHistoryTransactions/transactions.log": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.dashboard": DomainPrefs{
			"devmode": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"mcx-disabled": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.SystemManaged,
				},
			},
		},
		"com.apple.FontRegistry.user": DomainPrefs{
			"Version": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.imagent": DomainPrefs{
			"Setting.EnableReadReceipts": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
			"Setting.GlobalReadReceiptsVersionID": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"changeToken-syncStoreVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"IMChatVocabularyUpdaterDidPerformInitialUpdateKey": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
			"attachmentZoneChangeToken-syncStoreVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"messagesChangeToken-syncStoreVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
			"archivedMessagesChangeToken-syncStoreVersion": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.VersionMigration,
				},
			},
		},
		"com.apple.preferences.extensions.PhotoProjects": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
		"com.apple.messaging.watchdog": DomainPrefs{
			"watchdogWatermark": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.RuntimeState,
				},
			},
		},
		"com.apple.preferences.extensions.PhotoEditing": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: filters.Labels{
					&filters.UserManaged,
				},
			},
		},
	}
}
