package prefdefaults

import (
	"github.com/mikeschinkel/prefsctl/macprefs"
)

//goland:noinspection SpellCheckingInspection
func montereyPrefDefaults() DomainDefaults {
	return DomainDefaults{
		"GlobalPreferences": DomainPrefs{
			"AppleActionOnDoubleClick": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"AppleAntiAliasingThreshold": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"AppleAquaColorVariant": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"AppleEnableSwipeNavigateWithScrolls": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"AppleKeyboardUIMode": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"AppleMeasurementUnits": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"AppleMetricUnits": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"AppleMiniaturizeOnDoubleClick": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"AppleShowAllExtensions": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"AppleShowScrollBars": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"AppleTemperatureUnit": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"AppleWindowTabbingMode": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.mouse.scaling": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.sound.beep.flash": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.springing.delay": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.springing.enabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.swipescrolldirection": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.trackpad.forceClick": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.trackpad.scaling": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"InitialKeyRepeat": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"KeyRepeat": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NSAutomaticCapitalizationEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NSAutomaticDashSubstitutionEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NSAutomaticPeriodSubstitutionEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NSAutomaticQuoteSubstitutionEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NSAutomaticSpellingCorrectionEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NSAutomaticTextCompletionEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NSAutomaticWindowAnimationsEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"WebAutomaticSpellingCorrectionEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"WebKitDeveloperExtras": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"AppleLanguage": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
					&SetupSets,
					&StringType,
				},
			},
			"AppleLocale": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SetupSets,
					&StringType,
				},
			},
			"646F6E7A_00000000_00000001_6E7A6361_696D6963:": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"AKLastIDMSEnvironment:": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"AudioQuest inc.AudioQuest DragonFly:": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"Blue Microphones Yeti Stereo Microphone:": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"Canon MF5900 Series:": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"ContextMenuGesture:": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"Fujitsu ScanSnap iX500:": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"Yubico YubiKey OTP+FIDO+CCID:": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"AppleLanguagesDidMigrate:": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"NavPanelFileListModeForOpenMode:": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"NavPanelFileListModeForSaveMode:": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NSLinguisticDataAssetsRequestLastInterval:": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"QLPanelAnimationDuration:": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.security.sosaccount": DomainPrefs{
			"SOSEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.SafariCloudHistoryPushAgent": DomainPrefs{
			"AcknowledgedPushNotifications": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"UnacknowledgedPushNotifications": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.print.custompresets.forprinter.Canon_MF5900_Series": DomainPrefs{
			"com.apple.print.lastPresetPref": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"com.apple.print.lastPresetPrefType": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.finder": DomainPrefs{
			"SidebarTagsSctionDisclosedState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"QLEnableTextSelection": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"PreviewPaneGalleryWidth": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"FK_ArrangeBy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"PreferencesWindow.LastSelection": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"FXEnableExtensionChangeWarning": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NSWindowTabbingShoudShowTabBarKey-com.apple.finder.TBrowserWindow": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"FXICloudDriveDesktop": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ShowMountedServersOnDesktop": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"SidebarShowingiCloudDesktop": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"FXToolbarUpgradedToTenNine": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"ShowRemovableMediaOnDesktop": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ShowStatusBar": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"FXSidebarUpgradedSharedSearchToTwelve": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"FXInfoWindowWidth": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"DownloadsFolderListViewSettingsVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"SidebariCloudDriveSectionDisclosedState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AppleShowAllFiles": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"FXICloudDriveDeclinedUpgrade": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"FXArrangeGroupViewBy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"RecentsArrangeGroupViewBy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"FXPreferredSearchViewStyleVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"PreviewPaneWidth": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"FXAtLeastOneScreenHasBeenInHIDPI": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"CopyProgressWindowLocation": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"SidebarShowingSignedIntoiCloud": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ViewOptionsWindow.Location": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ShowHardDrivesOnDesktop": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"FXQuickActionsConfigUpgradeLevel": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"FXPreferredViewStyle": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"QuitMenuItem": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"FXDefaultSearchScope": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"FXPreferredSearchViewStyle": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"DisableAllAnimations": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"SidebarWidth": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ShowSidebar": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TagsCloudSerialNumber": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NewWindowTarget": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ShowPathbar": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"FXSidebarUpgradedToTenFourteen": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastTrashState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ShowExternalHardDrivesOnDesktop": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"GoToField": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"_FXSortFoldersFirstOnDesktop": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"FK_AppCentricShowSidebar": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"FXSidebarUpgradedToTenTen": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"FXConnectToLastURL": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"FXICloudDriveEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"FXICloudDriveDocuments": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"_FXSortFoldersFirst": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"_FXShowPosixPathInTitle": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"EmptyTrashProgressWindowLocation": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"SearchRecentsSavedViewStyle": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"FXDesktopTouchBarUpgradedToTenTwelveOne": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"FXICloudDriveFirstSyncDownComplete": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"SidebarDevicesSectionDisclosedState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"MountProgressWindowLocation": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"FK_SidebarWidth": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"SearchRecentsSavedViewStyleVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"FXPreferredGroupBy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"PreviewPaneInfoExpanded": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"FXLastSearchScope": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"InspectorWindow.Location": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"SidebarPlacesSectionDisclosedState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"FXToolbarUpgradedToTenEight": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"FXICloudLoggedIn": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"FXToolbarUpgradedToTenSeven": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.messages.pinning": DomainPrefs{
			"IMPinningPinConfigMigrationKey-v2-r2": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.identityservicesd": DomainPrefs{
			"ReRegisteredForDevicesv56": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"hasRegIdentityContainer": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"ReRegisteredForDevicesv55": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"persister-migration-personal-session-token-cache-v4": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"ReRegisteredForDevices": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"DidCleanLegacyAccountPrefs": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"ImportedLegacyIMAccounts": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"hasUnregIdentityContainer": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"MigratedToNewDisabledState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"triggeredRemoteSessionVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"ReRegisteredForDevicesHash": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"ImportedLegacyIDSAccounts2": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"GDRRequestMadeForRelayRepair": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"persister-migration-com.apple.identityservices.userStore": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"ReRegisteredForDevicesv1350": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"performed-user-intent-migrate": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.Spotlight": DomainPrefs{
			"ModelName": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.AppManaged,
				},
			},
			"engagementCount-com.apple.Spotlight": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"engagementCount-com.apple.mail": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"SSMessageTracingWindowShowCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"windowHeight": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"useCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"__NSEnableTSMDocumentWindowLevel": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"SSMessageTracingWindowHideCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"engagementCount-com.apple.Spotlight.suggestions": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"SPMessageTracingWindowHideCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"showedFTE": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"SPMessageTracingWindowShowCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"userHasMovedWindow": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.DictionaryServices": DomainPrefs{
			"DCSPreferenceVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.PTPCamera": DomainPrefs{
			"FilenamePrefix": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NextFilenameCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.imtranscoding.IMTranscoderAgent": DomainPrefs{
			"kCKMediaObjectManagerDefaultsOSVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.scriptmenu": DomainPrefs{
			"ScriptMenuEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.TTY": DomainPrefs{
			"RTTSettingsVersionPreference": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"RTTContinuityRTTIsSupportedPreference": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"TUSupportsRelayCallingPreference": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"TUIsRelayCallingEnabledPreference": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.voicetrigger": DomainPrefs{
			"VoiceTrigger Enabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.DiskUtility": DomainPrefs{
			"OperationProgress ExpandedHeight": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"DUDebugMenuEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"advanced-image-options": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"SidebarShowAllDevices": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"OperationProgress DetailsVisible": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.photos.shareddefaults": DomainPrefs{
			"CPLDefaultDownload": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"iCloudPhotoLibraryLastResetWelcomePromptDBVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"downloadAndKeepOriginals": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"clearPurgeableResources": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"DiskSpaceWasLow": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"repushAssetsWithImportedByBundleIdentifier": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.AppManaged,
				},
			},
			"repushVideoAssetsMetadata": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.AppManaged,
				},
			},
		},
		"com.apple.parsecd": DomainPrefs{
			"PARDefaultsVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"PARTestSeed": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.universalaccess": DomainPrefs{
			"login": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"switchCoalescePressesDuration": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"lastNightShiftEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"switchAutoScanPanelInterval": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"closeViewZoomedIn": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"grayscale": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"switchSweepingCursorSpeed": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"sessionChange": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"switchHoldBeforeRepeatDuration": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"closeViewZoomFactorBeforeTermination": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"switchHideUITimeout": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"switchAutoScanElementInterval": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"lastNightShiftActive": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"virtualKeyboardOnOff": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"headMouseEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"increaseContrast": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"mouseDriverCursorSize": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"selectedTab": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"lastNightShiftDate": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"closeViewZoomDisplayID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"hoverTextIsAlwaysOn": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"closeViewSplitScreenRatio": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"hoverTextIsHoveringAndVisible": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AssistiveControlType": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"closeViewDesiredZoomFactor": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"slowKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"closeViewScrollWheelToggle": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"mouseDriver": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"differentiateWithoutColor": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"slowKeyDelay": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"switchMinimumPressDuration": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"closeViewZoomFactor": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"switchOnOffKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"keyboardAccessFocusRingTimeout": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"stickyKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"reduceTransparency": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"dwellHideUITimeout": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"dwellEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"contrast": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"useStickyKeysShortcutKeys": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"lastNightShiftMode": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"hoverTextEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"whiteOnBlack": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"closeViewHotkeysEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"switchFirstElementDelay": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"voiceOverOnOffKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.Terminal": DomainPrefs{
			"HasMigratedDefaults": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"SecureKeyboardEntry": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Default Window Settings": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"DefaultProfilesVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"ProfileCurrentVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"TTAppPreferences Selected Tab": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"Startup Window Settings": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.coreservices.useractivityd": DomainPrefs{
			"kLocalPasteboardBlobName": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"kRemotePasteboardBlobName": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.ColorSyncUtility": DomainPrefs{
			"CSUtilitiesContentSize": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"CSProfsUtilsGroupBy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"CSUtilitiesSelectedIdentifier": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.weather": DomainPrefs{
			"userId": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.Music": DomainPrefs{
			"dontAskForAllPlaylistItemRemoval": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"storeSupportsPasswordSettings": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"notifications-warming-last-displayed-time": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"miniplayerVPos": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"com.apple.Music.WebKit2ShouldPrintBackgrounds": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"storeSupportsPreviousPurchases": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"com.apple.Music.WebKit2AsynchronousSpellCheckingEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"miniPlayerQueueVisible": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"eqWindowHPos": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"userWantsPlaybackNotifications": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"miniplayerUserSetHeight": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"com.apple.Music.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"soundEnhancerAmount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"preloadFilesIntoMemory": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"updateLevel": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"videoWindowVSize": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"com.apple.Music.WebKit2EnableInheritURIQueryComponent": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"automaticallyDownloadArtwork": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"refreshedHLSKeysTime": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"miniplayerWidthInDesignCoords": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"userMaxConcurrentDownloads": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.Music.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"miniplayerHPos": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"videoWindowVPos": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"log-PlayQueue-LastSelectedTab": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"dontWarnDownloadCloudPurchases": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"storeSupportsUPP": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"showWelcomeScreenState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"checkedHLSKeysTime": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"storeSupportsCloudPurchases": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"gridViewSearchDoesNotSwitchView": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"eqName": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ITUserPrefsMigrated": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"cddbPrefsOK": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"eqPrefsVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"playbackIsFullscreen": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"debugAssertCategoriesVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"debugAssertCategoriesEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"firstLaunchShowWelcomeScreenState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"encoderName": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.Music.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"haveRadioState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"videoWindowHSize": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"eqWindowVPos": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"haveAskedToCheckForAvailableDownloads": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"musicVideoPlaybackLocation": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"volumeWSG": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"dontShowRestrictionsPrefs": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"persistentUserID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"NSApplicationCrashOnExceptions": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"miniPlayerLargeArtVisible": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"searchSavedTab": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"artworkDownloadDSID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"videoWindowHPos": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"controllableInterfaceGUID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"userHasActivatedFullScreenVisualizer": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"storeMediaTypeFlags": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"imported-eq-presets": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"soundEnhancerEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"miniplayerSnapMode": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"doesStoreSupportCloudMusicLibrary": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"eqEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.airport.airportutility": DomainPrefs{
			"dontPerformBaseRestartWarning": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.studentd": DomainPrefs{
			"LastDateProviderSessionToken": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"DeviceIdentifier": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.systemuiserver": DomainPrefs{
			"__NSEnableTSMDocumentWindowLevel": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.IMCoreSpotlight": DomainPrefs{
			"IMCSIdxVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"IMCSLastIndexedRowID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"IMCSNeedsIndexing": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"IMCSIndexTotalRecords": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"IMCSIdxProductVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"IMCSBypassIndexVersionCheckV2": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.bird": DomainPrefs{
			"optimize-storage": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"icloud-drive.account-migration-status.294735135": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.xpc.activity2": DomainPrefs{
			"ProductBuildVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.iPod": DomainPrefs{
			"com.apple.PreferenceSync.ExcludeAllSyncKeys": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.newscore2": DomainPrefs{
			"report_concern_user_id": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"provider_user_id": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"instance_identifier": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.cloudd": DomainPrefs{
			"com.apple.private.cloudkit.shouldUseGeneratedDeviceID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.driver.AppleBluetoothMultitouch.trackpad": DomainPrefs{
			"TrackpadFourFingerHorizSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadPinch": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"USBMouseStopsTrackpad": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadFourFingerVertSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadRotate": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadHorizScroll": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadTwoFingerDoubleTapGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"TrackpadMomentumScroll": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadFourFingerPinchGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"UserPreferences": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Clicking": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadThreeFingerTapGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"HIDScrollZoomModifierMask": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadScroll": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"DragLock": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadFiveFingerPinchGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadThreeFingerHorizSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadThreeFingerVertSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadTwoFingerFromRightEdgeSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ForceSuppressed": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"TrackpadCornerSecondaryClick": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Dragging": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadHandResting": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadRightClick": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadThreeFingerDrag": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.stocks.account": DomainPrefs{
			"deleteOnNextLaunch": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.preference.trackpad": DomainPrefs{
			"ForceClickSavedState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.SetupAssistant": DomainPrefs{
			"PreviousSystemVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeSyncSetup2": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"SkipExpressSettingsUpdating": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"LastSeenDiagnosticsProductVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"SkipFirstLoginOptimization": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"LastSeenCloudProductVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeCloudSetup": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"PreviousBuildVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeApplePaySetup": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeTrueTonePrivacy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeScreenTime": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"MiniBuddyShouldLaunchToResumeSetup": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"DidSeeTouchIDSetup": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"NSAddServicesToContextMenus": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"DidSeeActivationLock": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeSyncSetup": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeAvatarSetup": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"MiniBuddyLaunchReason": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"DidSeeSiriSetup": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeTrueTone": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeePrivacy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastPrivacyBundleVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastSeenSiriProductVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeAppStore": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastPreLoginTasksPerformedVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeAppearanceSetup": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastSeenSyncProductVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastSeeniCloudStorageServicesProductVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastPreLoginTasksPerformedBuild": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"MiniBuddyLaunchedPostMigration": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"DidSeeAccessibility": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastSeenBuddyBuildVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidSeeiCloudLoginForStorageServices": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.ScriptEditor2": DomainPrefs{
			"OSAStandardAdditions ChooseApplication Bounds": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.preferences.extensions.ServicesWithUI": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.MobileSMS": DomainPrefs{
			"TextSize": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"SyndicationOnboardingVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"LegacyAppSidebarPersistedWidth": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"QuickSaveHasBeenUsedBefore": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"CatalystPreferenceMigrationVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"PlaySoundsKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"BusinessChatPrivacyPageDisplayed": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"sForceUnknownFilteringCompleted": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"CKLastSelectedItemIdentifier": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"kCKMediaObjectManagerDefaultsOSVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"CatalystCustomFontMigrationVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"NotifyAboutMessagesFromUnknownContacts": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"CKMigratedAutoSpamReports26375262": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"SidebarPersistedWidth": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ReadReceiptSettingsConfirmed": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"sForceSMSSpamFilteringCompleted": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"KeepMessageForDays": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"PendingCleared": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"TextFontSize": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.preferences.extensions.ShareMenu": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.mmcs": DomainPrefs{
			"report.LastFailedCheckDate": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"report.sha256": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"report.LastSuccessfulCheckDate": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"report.TTL": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.appstored": DomainPrefs{
			"BadgeCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ArcadeSubscriptionState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ArcadePayoutDeviceID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"WelcomeNotificationExcludedFromSample": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ArcadeDeviceID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"ArcadeDeviceGUID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"LastOSBuildVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"WelcomeNotificationLastAppStoreConnectionProductVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.PubSubAgent": DomainPrefs{
			"TigerMigrationCompleted": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.commcenter.callservices": DomainPrefs{
			"last.known.icloud.id": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"associated.account": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.print.PrinterProxy": DomainPrefs{
			"IK_Scanner_downloadURL": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"IK_Scanner_selectedTag": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"IK_Scanner_PostProcessApplication": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.amsengagementd": DomainPrefs{
			"AMSMetricsIdentifierUserRecordName": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"AMSMetricsIdentifierZoneCreated": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AMSMetricsIdentifierZoneSubscriptionCreated": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.iChat": DomainPrefs{
			"SaveConversationsOnClose": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"PMPrintingExpandedStateForPrint2": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"UnifiedChatWindowControllerSelectionGUIDSet": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"DidCheckForDuplicateChats": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"HasPromptedSMSRelay": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"HasPromptediMessageFTU": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"DidRegenerateGroupID63841559": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"messageTracerSMSSent": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"messageTracerSubmitDate": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"BuddyPictureSetToGenericByUser": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"messageTracerIMessagesSent": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"PlaySoundsKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"messageTracerSMSUsed": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"messageTracerNumUpgradeOffers": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ChatWindowControllerUnifiedIsVisible": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"messageTracerMessagesSent": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"WebIconDatabaseEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"DidMarkGroupPhotosAsUnpurgeable": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"messageTracerUpgradesAccepted": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"messageTracerIMessagesReceived": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"DaemonConnectionWaitTime": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"NotifyAboutMessagesFromUnknownContacts": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"KeepMessagesVersionID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"messageTracerCharactersSent": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"messageTracerMessagesReceived": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"messageTracerUpgradesDeclined": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"messageTracerSMSReceived": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"KeepMessageForDays": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"UnifiedChatListViewControllerAutomaticallySortsChats": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"messageTracerIMessageUsed": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"DidMigratePersonCentricIDs": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastIMDNotificationPostedDate": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ChatWindowControllerUnifiedChatListWidth": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"LastFailedMessageIMDNotificationPostedDate": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"SODefaultTranscriptName": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"imageBrowser.disableOpenGL": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.TextEdit": DomainPrefs{
			"PlainTextEncoding": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"RichText": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"PlainTextEncodingForWrite": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.HIToolbox": DomainPrefs{
			"AppleFnUsageType": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"AppleCurrentKeyboardLayoutInputSourceID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"AppleDictationAutoEnable": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.preferences.extensions.FinderSync": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.coreservices.uiagent": DomainPrefs{
			"CSUIHaveCenteredGatekeeperProgressWindow": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"CSUIHasSafariBeenLaunched": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"CSUIRecommendSafariBackOffInterval": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"CSUILastOSVersionWhereSafariRecommendationWasMade": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.CharacterPicker": DomainPrefs{
			"version": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.CharacterPaletteIM": DomainPrefs{
			"CVFontSize": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"CVStartAsLargeWindow": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.assistant.backedup": DomainPrefs{
			"Hands Free Mode": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Session Language": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Cloud Sync Enabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Cloud Sync User ID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"MultiUser VoiceIdentification Enabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.icloud.fmfd.notbackedup": DomainPrefs{
			"kKeychainUtilMigrationVersionKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.security.cloudkeychainproxy3.keysToRegister": DomainPrefs{
			"EnsurePeerRegistration": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"DSID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"KeyAccountUUID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.mediaaccessibility.public": DomainPrefs{
			"MACaptionDisplayType": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.speech.recognition.AppleSpeechRecognition.prefs": DomainPrefs{
			"DictationIMPresentedOfflineUpgradeSuggestion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"DictationIMCommandsWindowIsOpen": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"DictationIMCorrectionCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"DictationIMUpgradedTo10_16": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"DictationIMUpgradedTo10_15": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"DictationIMPlaySoundUponRecognition": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"DictationIMLocaleIdentifier": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"DictationIMMessageTracesSinceLastReport": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"DictationIMUseOnlyOfflineDictation": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.loginwindow": DomainPrefs{
			"TALLogoutSavesState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"MiniBuddyLaunch": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"oneTimeSSMigrationComplete": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.preferences.softwareupdate": DomainPrefs{
			"LatestMajorOSSeenByUserBundleIdentifier": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.accessibility.universalAccessAuthWarn": DomainPrefs{
			"ThirdPartyCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.internetconnect": DomainPrefs{
			"ServiceID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.TextInputMenu": DomainPrefs{
			"visible": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.messages.nicknames": DomainPrefs{
			"NicknameAppleIDAndiCloudAccountMatchAndExist": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"MeCardPendingNicknamesVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"MeCardSharingVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"MeCardSharingEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"MeCardSharingImageForkedFromMeCard": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"MeCardWhitelistBlacklistNicknamesVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"NicknameInfoRequestedFromPeers": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"MeCardSharingAudience": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NicknameScrutinyDoNotHandle": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"ReuploadLocalNicknamesVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.ProblemReporter": DomainPrefs{
			"ShowComments": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ShowDetails": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.gamed": DomainPrefs{
			"natType": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"GKStoreFrontIdentifier": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"GKActiveEnvironmentKeyV2": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"GKLastPushTokenPlayerID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"GKLastPushTokenEnvironment": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"GKPushEnvironmentKeyV2": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"GKLastProtocolVersionUsedKeyV2": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"GKLoginCancelled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.Safari.SandboxBroker": DomainPrefs{
			"DidMigrateDownloadFolderToSandbox": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"ShowDevelopMenu": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"DidMigrateResourcesToSandbox": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"DidMigrateDownloadMetadataToSandboxGroupContainer": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.preferences.extensions.XcodeSourceEditor": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.AppleMediaServices.notbackedup": DomainPrefs{
			"AMSDidRetrieveDeviceOffers": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AMSDidRetrieveDeviceOffersEligibility": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.imdsmsrecordstore": DomainPrefs{
			"DeleteSequenceNumber": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.iTunes": DomainPrefs{
			"com.apple.iTunes.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"doesAccountArtistListHaveSharePermission": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"dontAutomaticallySyncIPods": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"MacBuddyStoreID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"com.apple.iTunes.WebKit2ShouldPrintBackgrounds": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"com.apple.iTunes.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"com.apple.iTunes.WebKit2AsynchronousSpellCheckingEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AutomaticDeviceBackupsDisabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"disableRadio": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Store Apple ID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"debugAssertCategoriesVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"com.apple.iTunes.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"com.apple.iTunes.WebKit2EnableInheritURIQueryComponent": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"isAccountEnrolledInITunesMatch": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"log-PlayQueue-LastSelectedTab": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"isAccountAdmin": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"NSApplicationCrashOnExceptions": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"com.apple.iTunes.WebKit2DeveloperExtrasEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"DeviceBackupsDisabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"disableShareLibraryInfo": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"WirelessBuddyID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"Store DSID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"storefront": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"isAccountEnrolledInAppleMusic": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"log-push": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"debugAssertCategoriesEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.textInput.keyboardServices.textReplacement": DomainPrefs{
			"osxMajorVersForCloudKitSync": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"TRCKSyncMaxCountThreshold": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"_KSTRCloudKitMigratable": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"KSDidPushAllLocalRecordsOnce-2": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"osxMinorVersForCloudKitSync": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"internalMigrationPercent": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"KSCKSubscriptionProd-TextReplacements": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"seedMigrationPercent": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"seedMigrationPercent1": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"KSDidPushMigrationStatusOnce-2": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"gmMigrationPercent1": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"KSDidMigrateToCloudKitOnCloud": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"KSCloudKitMigrationDidComplete": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"kTRCKSyncCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"KSCKDidSetupRecordZoneProd-TextReplacements": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"iOSMajorVersForCloudKitSync": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"kDidMigrateToUUIDRecordName": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"internalMigrationPercent1": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"kDidInsertSampleShortcutForPeer": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"iOSMinorVersForCloudKitSync": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"gmMigrationPercent": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"userDidFallInMigrationAllowBatch": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"KSSampleShortcutWasImported_CK": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"KSDidPullLegacyEntriesFromPeers": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"TRCKSyncCountHalflifeInHours": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"KSLastKnownUserID-TextReplacements": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"osxMinorSubversionForCloudKitSync": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"KSLSShouldUpdateCache": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.diagnosticextensionsd": DomainPrefs{
			"directoriesCleanupDone": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.Messages": DomainPrefs{
			"hasPerformedNewDeviceBringUpSync": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"hasTriedToEnableCKAndSyncAfterFirstImagentConnection": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.NewDeviceOutreach": DomainPrefs{
			"InitialOutreachActivityScheduled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"CachedWarrantyValidityDuration": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"PostUpgradeActivityCompleted": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"PostUpgradeOSVersionKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"CachedWarrantyLocale": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"CachedWarrantyVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.icloud.searchpartyuseragent": DomainPrefs{
			"PersistenceLayerVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastOSLaunchVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"lastLaunchBootSessionUUID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"lastLaunchLocale": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"FMIPStateManager.fmipState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.cloudpaird": DomainPrefs{
			"MagicCloudPairingAccessorySubscriptionManateeID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"MagicCloudPairingManateeUpgradedAccount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"MagicCloudPairingManateeUpgraded": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"MagicCloudPairingProtectedAccessorySubscription": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"UploadedHSA2KeysForLocalDevice": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"MagicCloudPairingAccessorySubscriptionID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"MagicCloudPairingMasterEncryptionPreferences": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"SecuredAccessoryZoneSubscription": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"MagicCloudPairingMasterSubscriptionManateeID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"MagicCloudPairingMasterSubscriptionID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"SecuredZoneSubscription": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"MagicCloudPairingAccessoryEncryptionPreferences": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.syncserver": DomainPrefs{
			"SyncServicesResetWorldRunOnce": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"RunCompletelyDisabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.screencaptureui": DomainPrefs{
			"com.apple.AnnotationKit.userDefaultsVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"com.apple.AnnotationKit.arrowHeadStyle": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.AnnotationKit.strokeIsDashed": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.AnnotationKit.hasShadow": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.AnnotationKit.strokeWidth": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.AnnotationKit.highlightStyle": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"toolbarOrigin": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"com.apple.AnnotationKit.brushStyle": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.desktopservices": DomainPrefs{
			"DSDontWriteNetworkStores": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.security.pboxd": DomainPrefs{
			"ILMediaBrowserMediaType": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ILMediaBrowserSelectedGroupIndex1": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ILMediaBrowserSelectedGroup1": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.keychainaccess": DomainPrefs{
			"Keychain Item List Sorting": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Keychain Item List Sort Descending": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Last Selected Category": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"VPNSSItemsChecked": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"Item Preview Closed": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"Last Selected Keychain": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"Keychains List Closed": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.imessage": DomainPrefs{
			"PreviewTranscodingWasInitializedToDefaultValue": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"PreviewTranscodingQualityOnWiFiWasInitializedToDefaultValue": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.Safari.SafeBrowsing": DomainPrefs{
			"HasMigratedSafeBrowsingEnabledDefaults": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.AppleMediaServices": DomainPrefs{
			"AMSIncludeFullResponseInHARLogging": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"AMSShowSandboxAccountUI": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"AMSDeviceBiometricsState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"AMSUserDefaultsincludeFullResponseInHARLogging": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"AMSMigratedToNewAccountStore": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"AMSUserDefaultsIncludeFullRequestInHARLogging": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"AMSLogHARData": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"AMSLastMigratedBuildVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"AMSDidSetUpServerDataCache": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AMSIncludeFullRequestInHARLogging": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"AMSMigratedStorageToDefaultsForNonAMSInternal": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.PhotoBooth": DomainPrefs{
			"AVDesktopPlaybackControlsControllerShowsDurationInsteadOfTimeRemainingDefaultsKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"PBLaunchedAtLeastOnceOnLion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"UserCameraUniqueIDPref": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.MobileSMSPreview": DomainPrefs{
			"IMPreviewGenerationMaxPxWidth": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"IMPreviewGenerationMinHeight": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"IMPreviewGenerationScreenScale": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"IMPreviewGenerationMinWidth": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.ActivityMonitor": DomainPrefs{
			"OpenMainWindow": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ShowCategoryAppsinLast12Hours": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"SelectedTab": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ShowCategory": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"HasSavedState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.airplay": DomainPrefs{
			"showInMenuBarIfPresent": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.messages.facetime": DomainPrefs{
			"FaceTimeTwoTimeCallthroughEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.remindd": DomainPrefs{
			"CloudKitAccountStatus": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"preferredDefaultListObjectIDUrl": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ThrottlingPolicyCurrentBatchCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"isDatabaseMigrated": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"preferredDefaultListID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"cloudKitSchemaCatchUpSyncLastSuccessBuildVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"ThrottlingPolicyCurrentLevelIndex": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"spotlightIndexVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"useExtraneousAlarmBackOffThrottleInterval": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"CloudConfigurationPath": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"lastDatabaseMigrationSystemBuildVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.ncprefs": DomainPrefs{
			"content_visibility": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"sort_order": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.itunesstored": DomainPrefs{
			"AccountsNotificationPlugin.ActiveStorefrontIdentifier": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.driver.AppleHIDMouse": DomainPrefs{
			"Button4Force": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ScrollH": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ButtonDominance": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Button4": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Button4Click": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Button3": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Button1": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Button2": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ScrollSSize": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ScrollV": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ScrollS": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.preferences.extensions.CredentialProvider": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.itunescloud": DomainPrefs{
			"ICDefaultsKeyLastActiveAccountDSID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.stockholm": DomainPrefs{
			"RemoteAdminV2": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"PKWalletShouldAutomaticallyRegisterKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.Siri": DomainPrefs{
			"VoiceTriggerUserEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"StatusMenuVisible": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.sharingd": DomainPrefs{
			"AirDropRandomHashUUIDKey4": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"OneTimeAirDropReset2": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"OneTimeAirDropReset": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"HashManager-StoredDatabaseVersionKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"AirDropRandomHashUUIDKey2": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"DiscoverableMode": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"AutoUnlockPresentedWiFiError": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AutoUnlockPresentedBluetoothError": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AirDropRandomHashUUIDKey3": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"AutoUnlockWatchCurrentlyInList": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AutoUnlockWatchExistedInUnlockList": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AirDropRandomHashUUIDKey1": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.AppStoreComponents": DomainPrefs{
			"ASCLocaleID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.DFSLoginPlugin": DomainPrefs{
			"WindowLeft": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"WindowTop": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.iCal": DomainPrefs{
			"BirthdayEventsGenerationLastLocale": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"last calendar view description": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"BirthdayEventsGenerationVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"icaluuid": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"CalendarSidebarShown": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"AvailabilityShowTwentyFourHours": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"CalForceTruthFileRestoreHashKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"last selected calendar list item": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"CalDefaultCalendar": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"CalAgentNS_Preference_DefaultReminderCalendar": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"display birthdays calendar": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NSEventConcurrentProcessingEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"BirthdayEventsGenerationAttemptsToResetKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"TimeZone support enabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Show Week Numbers": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"iCal version": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"first shown minute of day": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"CalendarSidebarView": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"RemindersLastMigratedSystemVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"LastReminderMigrationCleanupVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AllDayAreaHeight": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"lastViewsTimeZone": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"privacyPaneHasBeenAcknowledgedVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"CalendarSidebarWidth": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"last minute of day time range": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"first minute of day time range": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"CalendarListMiniMonthVisibleMonths": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"IncludeDebugMenu": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"WR_DONT_ASK_FOR_DEFAULT": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"kLastABCReportTimeKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NSDontMakeMainWindowKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.assistant": DomainPrefs{
			"Experiment Identifierinvocation_feedback_experiment": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"Has Set Up MultiUser Shared Record Subscription": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"Initial Interstitial Delay for VoiceTrigger One-Shot without Audio Ducking": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"Initial Interstitial Delay - stark": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"MultiUser Shared Data Needs Sync": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"Initial Interstitial Delay": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"Myriad Device Delay": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"Myriad Device Class": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"Myriad Device Adjust": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"Has Set Up Account Status Subscription": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"Voice Trigger Needs Sync": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"Myriad Device Trump Delay": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"Experiment Identifier": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"Flush Session Tickets Cache": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"Server Has Provisioned Myriad": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"Myriad VTEndtimeDistanceThreshold": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"Has Set Up Key Value Subscription": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"Borealis Education Header Display Count": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"Experiment Identifiersiri_vox_tap_to_siri_behavior_experiment": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"Has Set Up Voice Trigger Subscription": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"Last Known Analytics Store State": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"Experiment Identifiersiri_vox_sounds_experiment": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"Manual Endpointing Threshold": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.dt.Xcode": DomainPrefs{
			"IBPreferencesMigrated": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"IBAppliesAutoResizingRulesWhileResizing": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"XCFontAndColorCurrentTheme": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"IDEKeyBindingCurrentPreferenceSet": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"IDESourceControlPreferencesVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"DVTFontAndColorLastUpdatedToolsVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"DVTDownloadableAutomaticUpdate": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"IDEUserWantsToEnableDeveloperSystemPolicyMode": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.AMPLibraryAgent": DomainPrefs{
			"iTunes-media-folder-url": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"debugAssertCategoriesVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"have-shown-cloud-UI": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"storeSupportsUPP": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"updateLevel": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"ITUserPrefsMigrated": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"books-migrated": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"storeSupportsPreviousPurchases": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"persistentUserID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"storeSupportsCloudPurchases": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"imported-media-domains": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"books-persistent-id": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"storeMediaTypeFlags": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"podcasts-migrated": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"storeSupportsPasswordSettings": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"imported-media-domains-modification-date": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"show-home-sharing-turned-on-in-iTunes-warning": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"devices-persistent-id": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"debugAssertCategoriesEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"podcasts-persistent-id": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.wifi.keychain-format": DomainPrefs{
			"Version": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.facetime.bag": DomainPrefs{
			"URL": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"CacheTime": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.TelephonyUtilities": DomainPrefs{
			"GondolaLastAccountsChangedState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"GondolaGeneratedIdentifier": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"registeredProvidersVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"CachedVCCaps": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"relayCallingDisabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"GondolaLatestVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"GondolaSyncedVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"FaceTimePhotosEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.findmy.fmfcore.notbackedup": DomainPrefs{
			"frontMostWindow": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"publicAPSToken": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"FMFGarbageCollectorIdentityKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"FMFLimitedPrecisionPrefKey.limitedPrecision": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"windowVisible": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.AppStore": DomainPrefs{
			"__uniquePageGroupID-9.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"__uniquePageGroupID-12.WebKit2DeveloperExtrasEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"mostRecentTabIdentifier": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AutoPlayVideoSetting": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.AppStore.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"appStoreBadgeCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"__uniquePageGroupID-1.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"UserSetAutoPlayVideoSetting": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"lastBootstrapTimeZone": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"__uniquePageGroupID-9.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ASAcknowledgedOnboardingVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"__uniquePageGroupID-12.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"__uniquePageGroupID-9.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"com.apple.AppStore.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"__uniquePageGroupID-1.WebKit2DeveloperExtrasEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"__uniquePageGroupID-9.WebKit2DeveloperExtrasEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AVDesktopPlaybackControlsControllerShowsDurationInsteadOfTimeRemainingDefaultsKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"__uniquePageGroupID-12.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"__uniquePageGroupID-12.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"com.apple.AppStore.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"com.apple.AppStore.WebKit2DeveloperExtrasEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"__uniquePageGroupID-1.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"__uniquePageGroupID-1.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.AdPlatforms": DomainPrefs{
			"AppStorePAAvailable": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"LatestPAVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"acknowledgedPersonalizedAdsVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.keyboard": DomainPrefs{
			"KeyboardWordOrSentenceTrackingForPFL": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.assistant.support": DomainPrefs{
			"Siri Data Sharing Opt-In Status": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Assistant Enabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Dictation Enabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.preferences.extensions.QuickLook": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.internetconfigspec": DomainPrefs{
			"WWWHomePage": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.icloud.fmfd": DomainPrefs{
			"numberOfFriendsFollowersKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"kFMFDStoredDataVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"storedDSID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"lazyInitTimeSecsStoredKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"kFMFDlastLoggedInPrsId": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.newscore": DomainPrefs{
			"notificationEnableAssetPrefetching": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.news.notification_receipt_event_endpoint": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"news_url_resolution_subscription_status": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"com.apple.news.default_event_endpoint": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"notificationAssetPrefetchingRequiresWatch": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"notificationArticleDiversificationSimilarityExpectationStart": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"force_refresh_user_segmentation": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"notificationArticleDiversificationUniquePublisherExpectationYIntercept": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"notificationArticleDiversificationUniquePublisherExpectationSlope": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"news_carplay_is_enabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.newsd.download.hasUnfinishedWork": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"notificationArticleDiversificationSimilarityExpectationEnd": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"use_parsec_results_for_widget": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.AddressBook": DomainPrefs{
			"ABBirthDayVisible": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ABMetaDataChangeCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ABTextSizeIncrement": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ABVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"ABLastImportShown": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ABDefaultSourceID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NSPreferencesSelectedIndex": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AB21vCardEncoding": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NSPreferencesContentSize": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ABPrivateVCardFieldsEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.SpeechRecognitionCore": DomainPrefs{
			"AllowAudioDucking": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.commerce": DomainPrefs{
			"LastUpdateNotificationOSMajorVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"PurchasesInflight": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.networkresolutiond": DomainPrefs{
			"_networkDevices": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.TV": DomainPrefs{
			"automaticallyDownloadArtwork": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"userWantsPlaybackNotifications": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"gridViewSearchDoesNotSwitchView": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"persistentUserID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"miniplayerUserSetHeight": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"storeMediaTypeFlags": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"dontAskForAllPlaylistItemRemoval": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"checkedHLSKeysTime": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"storeSupportsCloudPurchases": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"debugAssertCategoriesVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"com.apple.TV.WebKit2EnableInheritURIQueryComponent": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"cddbPrefsOK": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"com.apple.TV.WebKit2SuppressesIncrementalRendering": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"miniplayerVPos": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"controllableInterfaceGUID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"automaticallyDeleteVideoAssetsAfterWatching": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"videoWindowVPos": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"miniplayerHPos": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"playbackIsFullscreen": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"storeSupportsPasswordSettings": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"com.apple.TV.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"haveAskedToCheckForAvailableDownloads": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"videoWindowHPos": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"miniplayerSnapMode": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"storeSupportsUPP": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"showWelcomeScreenState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"videoWindowVSize": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"updateLevel": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"ITUserPrefsMigrated": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"miniplayerWidthInDesignCoords": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"storeSupportsPreviousPurchases": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"tabViewMode": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"preloadFilesIntoMemory": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"userMaxConcurrentDownloads": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NSApplicationCrashOnExceptions": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"videoWindowHSize": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"checkForAvailableDownloads": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"miniPlayerQueueVisible": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"debugAssertCategoriesEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"firstLaunchShowWelcomeScreenState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"com.apple.TV.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"musicVideoPlaybackLocation": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"miniPlayerLargeArtVisible": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"userHasActivatedFullScreenVisualizer": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"com.apple.TV.WebKit2AsynchronousSpellCheckingEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"com.apple.TV.WebKit2ShouldPrintBackgrounds": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.preference.general": DomainPrefs{
			"version": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.menuextra.battery": DomainPrefs{
			"ShowPercent": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.imdpersistence.IMDPersistenceAgent": DomainPrefs{
			"LogAllIOErrors": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.SafariServices": DomainPrefs{
			"SearchProviderIdentifierMigratedToSystemPreference": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.siri.shortcuts": DomainPrefs{
			"WFServicesShortcutsMigrated": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"WFDidUnconflictShortcuts": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"LegacyShortcutsZoneSubscriptionUnsubscribed": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"WFDefaultShortcutsVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"WFLastSyncedFlagsHash": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.SystemProfiler": DomainPrefs{
			"PrefsVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"SPLastDocumentsSize": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.remindd.babysitter": DomainPrefs{
			"LastSystemVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.internal.ck": DomainPrefs{
			"DictationOnDeviceSamplingDisabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"DictationSamplingRates": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"DefaultCacheKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"DefaultWarmupScriptsExtension": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"WarmupScriptIdentifiers": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"HasSetUpRecordZoneSubscription": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ByteCodeCacheEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"DefaultBootTimeUpdateScripts": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"WarmupModularScriptIdentifiers": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"DisableFBFForUEI": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.internetconfigpriv": DomainPrefs{
			"WWWHomePage": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.dock": DomainPrefs{
			"wvous-br-corner": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"showAppExposeGestureEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"autohide": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"tilesize": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"no-bouncing": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"region": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"showhidden": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"launchanim": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"showMissionControlGestureEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"wvous-tr-modifier": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"magnification": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"largesize": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"mod-count": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"trash-full": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"minimize-to-application": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"loc": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"expose-animation-duration": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"wvous-tr-corner": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"autohide-delay": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.AdLib": DomainPrefs{
			"adprivacydMaxSegmentSendInterval": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"personalizedAdsMigrated": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"partiality-segment": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"CKDPIDSyncState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"forceLimitAdTracking": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"adprivacydSegmentInterval": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"allowIdentifierForAdvertising": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"allowApplePersonalizedAdvertising": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.amp.mediasharingd": DomainPrefs{
			"public-sharing-enabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"home-sharing-sequence-id": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"home-sharing-group-id": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"home-sharing-computer-id": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"photo-sharing-enabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"shared-library-id": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"shared-library-name": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"public-sharing-is-protected": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"shared-library-machine-id": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"home-sharing-enabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.systempreferences": DomainPrefs{
			"com.apple.preferences.sharing.selectedservice": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ThirdPartyCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"SoundTab": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"com.apple.dtandsspref.lastselectedtab": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NSFullScreenMenuItemEverywhere": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"trackpad.lastselectedtab": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"ShowAllMode": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ICQ_iCloudFirstRunNotificationShown": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"DisableAutoLoginButtonIsHidden": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"com.apple.preferences.accounts.outline.usersparent": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"mouse.lastselectedtab": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NSQuitAlwaysKeepsWindows": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.SecurityPref.Privacy.LastSourceSelected": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"SecurityPrefTab": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"com.apple.preference.keyboard.lastselectedtab": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"com.apple.datetimepref.lastselectedtab": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.UIKit": DomainPrefs{
			"hasAccessibilityBeenMigrated": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.SafariBookmarksSyncAgent": DomainPrefs{
			"CloudBookmarksSupplementalClientIdentifier": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"NewestLaunchedSafariBookmarksSyncAgentVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.VoiceOverTraining": DomainPrefs{
			"doNotShowSplashScreen": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.newsd": DomainPrefs{
			"FCAppConfigurationBundleShortVersionKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.itunescloud.daemon": DomainPrefs{
			"ICDDefaultsKeyKnownActiveAccountDSID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.Image_Capture": DomainPrefs{
			"IK_lastUsedDeviceUUID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"IK_Camera_selectedTag": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"IK_lastUsedDeviceIsRemote": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"IK_Scanner_downloadURL": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"IKC_sort_ascending": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.imagekit.cameraviewmode": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"IK_prefsVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"IK_Camera_selectedPathType": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"IKC_sort_key": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"IK_Accessory_selectedTag": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"IK_Scanner_PostProcessApplication": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"IK_Camera_preferPostPocessingApp": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"IK_lastUsedDeviceName": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.preferences.extensions.SharedLinks": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.driver.AppleBluetoothMultitouch.mouse": DomainPrefs{
			"MouseOneFingerDoubleTapGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"MouseTwoFingerHorizSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"MouseVerticalScroll": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"UserPreferences": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"MouseHorizontalScroll": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"MouseButtonDivision": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"MouseMomentumScroll": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"MouseButtonMode": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"MouseTwoFingerDoubleTapGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.dock.external.extra.x86_64": DomainPrefs{
			"ShowSolidarity": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.commcenter.data": DomainPrefs{
			"pw_ver": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.imagecapture": DomainPrefs{
			"loggingLevel": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.notificationcenterui": DomainPrefs{
			"bannerTime": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.pipagent": DomainPrefs{
			"Size": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.AMPDevicesAgent": DomainPrefs{
			"debugAssertCategoriesEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"persistentUserID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"dontAutomaticallySyncIPods": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"updateLevel": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"storeMediaTypeFlags": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"AutomaticDeviceBackupsDisabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"storeSupportsCloudPurchases": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"debugAssertCategoriesVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"ITUserPrefsMigrated": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"userMaxConcurrentDownloads": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.MobileSMS.CKDNDList": DomainPrefs{
			"CatalystDNDMigrationVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"CKDNDMigrationKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.protectedcloudstorage.protectedcloudkeysyncing": DomainPrefs{
			"registrySyncIdentifier": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.VoiceOver4.local": DomainPrefs{
			"SCRCUserDefaultsAllowAirPlay": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"SCRCUserDefaultsUnplannedShutdownCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"SCRCUserDefaultsPlannedShutdownSuccessful": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.suggestions": DomainPrefs{
			"SuggestionsAllowGeocode": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"spToLearnMigrationPerformed": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"findToShowMigrationPerformed": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.NetworkUtility": DomainPrefs{
			"NUPingCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NUTracerouteAddress": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NUPortScanEnd": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NUSelectedTabItem": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NULookupAddress": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NUFingerPerson": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NUPortScanAddress": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NUPingChoice": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NUPingAddress": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NUNetstatChoice": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NUPortScanStart": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NUWhoisAddress": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NUWhoisSelectedServer": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"NUPortScanRange": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.siri.VoiceShortcuts": DomainPrefs{
			"VCLSDatabaseUUIDKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"VCLSDataSequenceKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.Safari": DomainPrefs{
			"ResetCloudHistory": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"WebKitHistoryItemLimit": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"WebKitInitialTimedLayoutDelay": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ShowFullURLInSmartSearchField": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"com.apple.Safari.ContentPageGroupIdentifier.WebKit2DeveloperExtrasEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"cloudBookmarksMigrationEligibilityDataInvalidated": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"UniversalSearchEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"WebKitDeveloperExtrasEnabledPreferenceKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"SuppressSearchSuggestions": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"WebKitHistoryAgeInDaysLimit": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"IncludeDebugMenu": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"IncludeDevelopMenu": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"IncludeInternalDebugMenu": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.accounts": DomainPrefs{
			"ACOneTimeSMTPFixAccountSync": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"ACOneTimeLDAPFixAccountSync": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.visualintelligence": DomainPrefs{
			"bypassAvailability": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"enableBooksDomain": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"enableLandmarkDomain": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"enableArtDomain": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"firstTimeExperience": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"enableCoarseClassification": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"initialized": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"enablePetsDomain": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"enableScreenshots": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"debugUI": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"enableSafariApp": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"sendLocationInfo": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"enableAlbumsDomain": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"enableQuickLook": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"enableNatureDomain": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"enablePhotosApp": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"sendOCRText": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.accountsd": DomainPrefs{
			"com.apple.mail.searchableIndex.lastProcessedAttachmentIDKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"LastSystemVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.fileproviderd": DomainPrefs{
			"com.apple.fileproviderd.did-drop-daemon-corespotlight-index": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.screencapture": DomainPrefs{
			"location": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"style": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"captureDelay": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"last-selection-display": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"video": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.SoftwareUpdate": DomainPrefs{
			"AutoUpdateMajorOSVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.Preferences": DomainPrefs{
			"UserDictionaryImportedSinceMaintenance": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"UserDictionaryLocalPeerIdentityCurrent": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.sociallayerd.CloudKit.ckwriter": DomainPrefs{
			"accountLastKnownUserRecordID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"accountAvailable": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"writerDone": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.classroom": DomainPrefs{
			"allowClassroomLockDevice-00000000-0000-0000-0000-000000000000:0": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"allowClassroomOpenApp-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"allowClassroomScreenObservation-00000000-0000-0000-0000-000000000000:0": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"allowClassroomOpenURL-ask": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"allowClassroomScreenObservation-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"allowClassroomOpenURL": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"allowClassroomLockDevice-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"allowClassroomOpenURL-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"allowClassroomOpenApp-00000000-0000-0000-0000-000000000000:0": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"forceUnpromptedRemoteScreenObservation-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"allowClassroomOpenURL-00000000-0000-0000-0000-000000000000:0": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"forceUnpromptedRemoteScreenObservation-00000000-0000-0000-0000-000000000000:0": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"allowClassroomOpenApp": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"allowClassroomOpenApp-ask": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.CloudKit": DomainPrefs{
			"AccountInfoValidationCounter": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.findmy.fmipcore.notbackedup": DomainPrefs{
			"publicAPSToken": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"itemLearnMoreURL": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"frontMostWindow": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"windowVisible": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"deviceImageVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"FMIPGarbageCollectorIdentityKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"FMIPLimitedPrecisionPrefKey.limitedPrecision": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.tipsd": DomainPrefs{
			"TPSWelcomeNotificationViewedVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"TPSTipsAppInstalled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"TPSOfflineLastProcessedRemoteContentIdentifier": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"DeliveryInfoVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"TPSLastMajorVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"ConsecutiveNotificationsCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"TPSWelcomeNotificationReminderState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.CoreGraphics": DomainPrefs{
			"DisplayUseForcedGray": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"DisplayUseInvertedPolarity": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.DirectoryUtility": DomainPrefs{
			"FirstLaunch": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.AvatarUI.Staryu": DomainPrefs{
			"AVTAvatarHasDisplayedAnimojiSplashScreen": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AVTAvatarHasDisplayedCameraEffectsSplashScreen": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AVTAvatarHasDisplayedPaddleView": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AVTAvatarHasDisplayedSplashScreen": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AVTAvatarUILastCacheVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.networkConnect": DomainPrefs{
			"VPNShowTime": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.UserAccountUpdater": DomainPrefs{
			"com.apple.HidLibraryFolderAlready": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.madrid": DomainPrefs{
			"IMCloudKitStartingInitialSync": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"CloudKitIsSyncing": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"CloudKitIsRemovedFromBackup": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"CloudKitIsEligibleForTruthZone": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"CloudKitEligibleToToggleMiCSwitch": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"IMCloudKitSyncControllerSyncTypeKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"CloudKitCheckedMiCSwitchEligibilityOnImagentLaunch": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"createdChatZone": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"hasCompletedInitialCKSync": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"CloudKitSyncingEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"IMCloudKitSyncControllerSyncRecordTypeKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"initialSyncRecordHasBeenWritten": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"IMCloudKitSyncControllerSyncStateKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"IMCloudKitStartingPeriodicSync": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"IMCloudKitSyncPaused": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"RequestPriorityRamp": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"IMCloudKitAccountStatusKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"IMCloudKitStartingEnabledSettingChange": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"IMCKFinishedFetchingAttachmentsFromCloudKit": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"CKMOCAccountsMatch": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"IMCloudKitStartingDisableDevices": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"oneTimeMOCAccountCheckV2": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"enableCKSyncingV2": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.rapport": DomainPrefs{
			"familySyncedName": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.menuextra.clock": DomainPrefs{
			"DateFormat": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.imessage.bag": DomainPrefs{
			"URL": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"CacheTime": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.sms": DomainPrefs{
			"hasBeenApprovedForSMSRelay": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.Automator": DomainPrefs{
			"NSSplitView AMDocumentMinor Expanded Position": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NSNavLastUserSetHideExtensionButtonState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"NSSplitView AMLibraryActionsMajor Expanded Position": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AMLogTabViewSelectedIndex": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.weather.internal": DomainPrefs{
			"userId": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.FaceTime": DomainPrefs{
			"DidMigrateToSixteenByNineAspectRatio": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"kPromptEnableReadRecipts": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"CustomRingtone": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"PhoneNumberUpgradeShown": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"PreferredVideoDeviceUID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.AssistiveControl.virtualKeyboard": DomainPrefs{
			"SCLaunchedAsSlave": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"PanelPosition": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"DesiredPanelWindowPosition": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.Accessibility": DomainPrefs{
			"InvertColorsEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"GrayscaleDisplay": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"AXSClassicInvertColorsPreference": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"KeyRepeatInterval": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"CommandAndControlEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ApplicationAccessibilityEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"DarkenSystemColors": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"FullKeyboardAccessFocusRingEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"GenericAccessibilityClientEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"AccessibilityEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"KeyRepeatEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"KeyRepeatDelay": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ZoomTouchEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"DifferentiateWithoutColor": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"FullKeyboardAccessEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"BrailleInputDeviceConnected": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"AutomationEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ReduceMotionEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"SpeakThisEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"EnhancedBackgroundContrastEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"VoiceOverTouchEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.ipTelephony": DomainPrefs{
			"ImsLoggingEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.ColorSyncCalibrator": DomainPrefs{
			"722524374X-shareMode": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"722524374X-targetTemp": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"722524374X-name": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"722524374X-expertMode": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.StatusKitAgent": DomainPrefs{
			"reauthCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.AppleMultitouchTrackpad": DomainPrefs{
			"TrackpadThreeFingerDrag": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadPinch": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"USBMouseStopsTrackpad": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadFourFingerVertSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ActuationStrength": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ActuateDetents": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"FirstClickThreshold": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"SecondClickThreshold": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadTwoFingerDoubleTapGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadHorizScroll": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadMomentumScroll": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadFourFingerPinchGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadRotate": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Clicking": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadThreeFingerTapGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadThreeFingerHorizSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadScroll": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"DragLock": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadFiveFingerPinchGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"ForceSuppressed": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"TrackpadThreeFingerVertSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadTwoFingerFromRightEdgeSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"UserPreferences": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"HIDScrollZoomModifierMask": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadHandResting": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadRightClick": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Dragging": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadFourFingerHorizSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"TrackpadCornerSecondaryClick": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.ImageCaptureService": DomainPrefs{
			"IK_Accessory_selectedTag": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.screensaver": DomainPrefs{
			"tokenRemovalAction": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.AppleMultitouchMouse": DomainPrefs{
			"MouseOneFingerDoubleTapGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"MouseTwoFingerHorizSwipeGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"MouseVerticalScroll": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"UserPreferences": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"MouseHorizontalScroll": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"MouseButtonDivision": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"MouseMomentumScroll": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"version": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"MouseButtonMode": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"MouseTwoFingerDoubleTapGesture": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.universalaccessAuthWarning": DomainPrefs{
			"2::/Applications/Logi Options.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::com.logitech.manager.daemon": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Microsoft Teams (work or school).app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::com.lightpillar.Mosaic-setapp": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"1::/Applications/CloudBerry Backup.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::com.getcleanshot.app-setapp": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::/Applications/Setapp/Mosaic.app/Contents/MacOS/Mosaic": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"/Applications/Clipy.app/Contents/MacOS/Clipy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/xScope.app/Contents/MacOS/xScope": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Microsoft Teams.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"1::/Users/mikeschinkel/Library/Application Support/JetBrains/Toolbox/apps/Goland/ch-0/232.8660.185/GoLand.app/Contents/MacOS/goland": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"1::/Applications": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::com.microsoft.teams2": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Discord.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Microsoft Edge.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::/Applications/Logi Options.app/Contents/Support/LogiMgrDaemon.app/Contents/MacOS/LogiMgrDaemon": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Skitch.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/ScreenFlow.app/Contents/Library/LoginItems/ScreenFlowHelper.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"/Applications/Alfred 4.app/Contents/MacOS/Alfred": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Toast 19 Pro.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"1::com.vmware.fusion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::com.lastpass.lastpassmacdesktop": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"com.clipy-app.Clipy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::com.google.Chrome": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"1::/Users/mikeschinkel/Library/Application Support/JetBrains/Toolbox/apps/Goland/ch-0/232.8660.185/GoLand.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"1::/Applications/Microsoft Teams.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/ScreenFlow.app/Contents/Library/LoginItems/ScreenFlowHelper.app/Contents/MacOS/ScreenFlowHelper": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Microsoft Teams (work or school).app/Contents/MacOS/MSTeams": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/RingCentral.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::com.amazon.Amazon-Chime": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"com.mutedeck.mac": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Parallels Toolbox.app/Contents/Applications/Hide Menu Icons.app/Contents/MacOS/Hide Menu Icons": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::/Applications/Snappy.app/Contents/MacOS/Snappy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::me.waydabber.BetterDummy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/GoToMeeting.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Google Chrome.app/Contents/MacOS/Google Chrome": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::WSG985FR47.net.telestream.screenflowhelper": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Slack.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"/Applications/Alfred 4.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"1::/Applications/BetterDisplay.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::com.logi.cp-dev-mgr": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"com.runningwithcrayons.Alfred": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Xnapper.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/xScope.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::com.microsoft.edgemac": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"/Applications/zoom.us.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"/System/Applications/Automator.app/Contents/MacOS/Automator": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"/Applications/Clipy.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/BetterDisplay.app/Contents/MacOS/BetterDummy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Snip.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::ro.nextwave.Snappy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::com.parallels.toolbox.HideMenuIcons": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Google Chrome.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"1::/Applications/VMware Fusion-8.5.10.app/Contents/MacOS/VMware Fusion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::com.ringcentral.glip": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Acorn.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"/Applications/MuteDeck/MuteDeck.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::/Library/Application Support/Logitech.localized/LogiOptionsPlus/logioptionsplus_agent.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Timing.app/Contents/Library/LoginItems/TimingHelper.app/Contents/MacOS/TimingHelper": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::/Applications/Logi Options.app/Contents/Support/LogiMgrDaemon.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::/Applications/Setapp/Mosaic.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"1::me.waydabber.BetterDummy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::/Applications/LastPass.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::info.eurocomp.TimingHelper": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::com.amazon.Amazon-Chime": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"1::/Applications/BetterDisplay.app/Contents/MacOS/BetterDummy": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"1::/Applications/iTerm.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"1::/Applications/OBS.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"1::/Applications/HP RGS Receiver.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::J8RPQ294UB.com.skitch.SkitchHelper": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::com.logitech.Logi-Options": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::/Applications/Logi Options.app/Contents/MacOS/Logi Options": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"us.zoom.xos": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"1::/Applications/Slack.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::/Applications/Snappy.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"1::com.jetbrains.goland": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::info.eurocomp.Timing2": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Xnapper.app/Contents/MacOS/Xnapper": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/RingCentral.app/Contents/MacOS/RingCentral": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"/Applications/zoom.us.app/Contents/MacOS/zoom.us": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/BetterDisplay.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/zoom.us.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"com.apple.Automator": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Skitch.app/Contents/Library/LoginItems/J8RPQ294UB.com.skitch.SkitchHelper.app/Contents/MacOS/J8RPQ294UB.com.skitch.SkitchHelper": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Amazon Chime.app/Contents/MacOS/Amazon Chime": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Setapp/CleanShot X.app/Contents/MacOS/CleanShot X Setapp": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/ScreenFlow.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::/Applications/Amazon Chime.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"1::/Applications/VirtualBox.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"/Applications/MuteDeck/MuteDeck.app/Contents/MacOS/MuteDeck": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Timing.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::/Applications/Amazon Chime.app/Contents/MacOS/Amazon Chime": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Timing.app/Contents/MacOS/Timing": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Amazon Chime.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Skitch.app/Contents/Library/LoginItems/J8RPQ294UB.com.skitch.SkitchHelper.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::com.iconfactory.xScope": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::/Applications/LastPass.app/Contents/MacOS/LastPass": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"/System/Applications/Automator.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"1::/Applications/VMware Fusion-8.5.10.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Loom.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/var/tmp/LogMeIn Rescue - 4C61/Support-LogMeInRescue.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Setapp/CleanShot X.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::com.devuap.beautyshotapp": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Timing.app/Contents/Library/LoginItems/TimingHelper.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/var/tmp/LogMeIn Rescue - 4C61": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"3::/Applications/Parallels Toolbox.app/Contents/Applications/Hide Menu Icons.app": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"2::/Library/Application Support/Logitech.localized/LogiOptionsPlus/logioptionsplus_agent.app/Contents/MacOS/logioptionsplus_agent": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.findmy": DomainPrefs{
			"restoreState": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"coarseServiceAcknowledged_v1.0": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"onboarding_v2.0": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.FolderActionsDispatcher": DomainPrefs{
			"folderActionsEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.SocialLayer": DomainPrefs{
			"SharedWithYouEnabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.security.KCN": DomainPrefs{
			"absentCircleWithNoReason": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
			"lastCircleStatus": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"pendingApplicationReminderInterval": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.GEO": DomainPrefs{
			"DefaultsSanitizedVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.SharedWebCredentials": DomainPrefs{
			"recheckVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.suggestd": DomainPrefs{
			"DeviceIdentifier": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.CallHistorySyncHelper": DomainPrefs{
			"/Users/mikeschinkel/Library/Application Support/CallHistoryTransactions/transaction.log": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"CallHistoryDeviceCount": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"/Users/mikeschinkel/Library/Application Support/CallHistoryTransactions/transactions.log": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.dashboard": DomainPrefs{
			"devmode": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"mcx-disabled": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&SystemManaged,
				},
			},
		},
		"com.apple.FontRegistry.user": DomainPrefs{
			"Version": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.imagent": DomainPrefs{
			"Setting.EnableReadReceipts": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
			"Setting.GlobalReadReceiptsVersionID": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"changeToken-syncStoreVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"IMChatVocabularyUpdaterDidPerformInitialUpdateKey": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
			"attachmentZoneChangeToken-syncStoreVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"messagesChangeToken-syncStoreVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
			"archivedMessagesChangeToken-syncStoreVersion": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&macprefs.VersionMigration,
				},
			},
		},
		"com.apple.preferences.extensions.PhotoProjects": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
		"com.apple.messaging.watchdog": DomainPrefs{
			"watchdogWatermark": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&RuntimeState,
				},
			},
		},
		"com.apple.preferences.extensions.PhotoEditing": DomainPrefs{
			"userHasOrdered": DomainPref{
				NoDefault: true,
				Labels: Labels{
					&UserManaged,
				},
			},
		},
	}
}
