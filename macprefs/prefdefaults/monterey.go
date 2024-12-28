package prefdefaults

//goland:noinspection SpellCheckingInspection
func montereyPrefDefaults() DomainDefaults {
	return DomainDefaults{
		"com.apple.Accessibility": DomainPrefs{
			"AccessibilityEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ApplicationAccessibilityEnabled": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AutomationEnabled": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AXSClassicInvertColorsPreference": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"BrailleInputDeviceConnected": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CommandAndControlEnabled": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DarkenSystemColors": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DifferentiateWithoutColor": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"EnhancedBackgroundContrastEnabled": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FullKeyboardAccessEnabled": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FullKeyboardAccessFocusRingEnabled": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"GenericAccessibilityClientEnabled": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"GrayscaleDisplay": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"InvertColorsEnabled": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KeyRepeatDelay": DomainPref{
				Type:    "float",
				Default: "0.25",
				Labels: NewLabels(
					DefaultsSet,
					TypeVerified,
					UserManaged,
				),
			},
			"KeyRepeatEnabled": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ReduceMotionEnabled": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SpeakThisEnabled": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"VoiceOverTouchEnabled": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ZoomTouchEnabled": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.accounts": DomainPrefs{
			"ACOneTimeLDAPFixAccountSync": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ACOneTimeSMTPFixAccountSync": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ActivityMonitor": DomainPrefs{
			"cacheTableSortDescriptors": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Column Width": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"OpenMainWindow": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SelectedTab": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowCategory": DomainPref{
				Type:    "int",
				Default: "100",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowCategoryAppsinLast12Hours": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserColumnSortPerTab": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserColumnsPerTab v5.0": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserColumnsPerTab v6.0": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.AddressBook": DomainPrefs{
			"AB21vCardEncoding": DomainPref{
				Type:    "string",
				Default: "macintosh",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ABBirthDayVisible": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ABCleanWindowController-MainCleanWindow-groupList": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ABCleanWindowController-MainCleanWindow-personListController": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ABLastImportShown": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ABMetadataLastOilChange": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ABPrivateVCardFieldsEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ABTextSizeIncrement": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSPreferencesSelectedIndex": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.airplay": DomainPrefs{
			"showInMenuBarIfPresent": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.airport.airportutility": DomainPrefs{
			"configuredAirPortIDs": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dontPerformBaseRestartWarning": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"RecentNetworks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.amp.mediasharingd": DomainPrefs{
			"home-sharing-enabled": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"home-sharing-sequence-id": DomainPref{
				Type:    "int",
				Default: "1183",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"home-sharing-settings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"photo-sharing-enabled": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"photo-sharing-settings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"public-sharing-enabled": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"public-sharing-is-protected": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"public-sharing-settings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"shared-library-name": DomainPref{
				Type:    "string",
				Default: "Mike’s MacBook Pro Library",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.AMPDevicesAgent": DomainPrefs{
			"AutomaticDeviceBackupsDisabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"debugAssertCategoriesEnabled": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dontAutomaticallySyncIPods": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ITUserPrefsMigrated": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MobileDeviceSoftwareUpdateQueue": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeMediaTypeFlags": DomainPref{
				Type:    "int",
				Default: "3422",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsCloudPurchases": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"updateLevel": DomainPref{
				Type:    "int",
				Default: "43",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userMaxConcurrentDownloads": DomainPref{
				Type:    "int",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.AMPLibraryAgent": DomainPrefs{
			"books-migrated": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"debugAssertCategoriesEnabled": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"have-shown-cloud-UI": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"imported-media-domains": DomainPref{
				Type:    "int",
				Default: "15",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"iTunes-media-folder-bookmark": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"iTunes-media-folder-url": DomainPref{
				Type:    "string",
				Default: "file:///Users/mikeschinkel/Music/iTunes/iTunes%20Media/",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ITUserPrefsMigrated": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"podcasts-migrated": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"show-home-sharing-turned-on-in-iTunes-warning": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeMediaTypeFlags": DomainPref{
				Type:    "int",
				Default: "3422",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsCloudPurchases": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsPasswordSettings": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsPreviousPurchases": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsUPP": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"updateLevel": DomainPref{
				Type:    "int",
				Default: "43",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.amsaccountsd": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.animoji": DomainPrefs{
			"AVTRecentStickers": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ap.adprivacyd": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.appkit.xpc.openAndSavePanelService": DomainPrefs{
			"NSNavRecentPlaces": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.appleaccountd": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastCloudSyncTimestampKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.AppleMultitouchMouse": DomainPrefs{
			"MouseButtonDivision": DomainPref{
				Type:    "int",
				Default: "55",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseButtonMode": DomainPref{
				Type:    "string",
				Default: "TwoButton",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseHorizontalScroll": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseMomentumScroll": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseOneFingerDoubleTapGesture": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseTwoFingerDoubleTapGesture": DomainPref{
				Type:    "int",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseTwoFingerHorizSwipeGesture": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseVerticalScroll": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserPreferences": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.AppleMultitouchTrackpad": DomainPrefs{
			"ActuateDetents": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ActuationStrength": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Clicking": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Dragging": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DragLock": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FirstClickThreshold": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ForceSuppressed": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"HIDScrollZoomModifierMask": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SecondClickThreshold": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadCornerSecondaryClick": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFiveFingerPinchGesture": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFourFingerHorizSwipeGesture": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFourFingerPinchGesture": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFourFingerVertSwipeGesture": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadHandResting": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadHorizScroll": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadMomentumScroll": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadPinch": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadRightClick": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadRotate": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadScroll": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerDrag": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerHorizSwipeGesture": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerTapGesture": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerVertSwipeGesture": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadTwoFingerDoubleTapGesture": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadTwoFingerFromRightEdgeSwipeGesture": DomainPref{
				Type:    "int",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"USBMouseStopsTrackpad": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserPreferences": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.AppStore": DomainPrefs{
			"AutoPlayVideoSetting": DomainPref{
				Type:    "string",
				Default: "on",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AVDesktopPlaybackControlsControllerShowsDurationInsteadOfTimeRemainingDefaultsKey": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AppStore.WebKit2DeveloperExtrasEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AppStore.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AppStore.WebKit2SuppressesIncrementalRendering": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AppStore.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"JE.MediaAPIToken": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastBootstrapTimeZone": DomainPref{
				Type:    "string",
				Default: "America/New_York",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserSetAutoPlayVideoSetting": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.assistant": DomainPrefs{
			"Account Status Server Change Token": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Accounts": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Has Set Up Account Status Subscription": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Has Set Up Key Value Subscription": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Has Set Up MultiUser Shared Record Subscription": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Has Set Up Voice Trigger Subscription": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Initial Interstitial Delay": DomainPref{
				Type:    "float",
				Default: "3.5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Initial Interstitial Delay - stark": DomainPref{
				Type:    "int",
				Default: "5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Initial Interstitial Delay for VoiceTrigger One-Shot without Audio Ducking": DomainPref{
				Type:    "float",
				Default: "1.75",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Manual Endpointing Threshold": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MultiUser Shared Data Needs Sync": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Myriad Device Adjust": DomainPref{
				Type:    "int",
				Default: "-10",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Myriad Device Class": DomainPref{
				Type:    "int",
				Default: "9",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Myriad Device Delay": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Myriad Device Trump Delay": DomainPref{
				Type:    "float",
				Default: "1.75",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Myriad Last Win": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Myriad VTEndtimeDistanceThreshold": DomainPref{
				Type:    "float",
				Default: "0.5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Server Has Provisioned Myriad": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Voice Trigger Needs Sync": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.assistant.backedup": DomainPrefs{
			"Cloud Sync Enabled": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Cloud Sync User ID": DomainPref{
				Type:    "string",
				Default: "_44c58d2cbdc7b03bb2cdefc57dcb364d",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Hands Free Mode": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Modification Dates": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MultiUser VoiceIdentification Enabled": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Output Voice": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Output Voice v3": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Session Language": DomainPref{
				Type:    "string",
				Default: "en-US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.assistant.support": DomainPrefs{
			"Assistant Enabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Dictation Enabled": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Offline Dictation Status": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Siri Data Sharing Opt-In Status": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.assistantd": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.AssistiveControl.virtualKeyboard": DomainPrefs{
			"DesiredPanelWindowPosition": DomainPref{
				Type:    "string",
				Default: "Absolute",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PanelPosition": DomainPref{
				Type:    "string",
				Default: "Absolute",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SCLaunchedAsSlave": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ATS": DomainPrefs{
			"ATSAutoActivationAppSpecific": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Automator": DomainPrefs{
			"AMLogTabViewSelectedIndex": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.AvatarUI.Staryu": DomainPrefs{
			"AVTAvatarHasDisplayedAnimojiSplashScreen": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AVTAvatarHasDisplayedCameraEffectsSplashScreen": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AVTAvatarHasDisplayedPaddleView": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AVTAvatarHasDisplayedSplashScreen": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.BezelServices": DomainPrefs{
			"afActionHistory": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"BatteryHistory": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ChargingEvents": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TransportElapsedSeconds": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.bookdatastored": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"kJaliscoDAAPClientLastSuccessfulFullUpdateDateDictionary": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.CacheDelete": DomainPrefs{
			"com.apple.cache_delete.purge_markers": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.CharacterPaletteIM": DomainPrefs{
			"CVFontSize": DomainPref{
				Type:    "int",
				Default: "26",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CVStartAsLargeWindow": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.CharacterPicker": DomainPrefs{
			"State:com.apple.CharacterPaletteIM": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"State:com.apple.iChat": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"State:com.apple.MobileSMS": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.classroom": DomainPrefs{
			"allowClassroomLockDevice-00000000-0000-0000-0000-000000000000:0": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomLockDevice-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenApp": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenApp-00000000-0000-0000-0000-000000000000:0": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenApp-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenApp-ask": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenURL": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenURL-00000000-0000-0000-0000-000000000000:0": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenURL-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenURL-ask": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomScreenObservation-00000000-0000-0000-0000-000000000000:0": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomScreenObservation-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"forceUnpromptedRemoteScreenObservation-00000000-0000-0000-0000-000000000000:0": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"forceUnpromptedRemoteScreenObservation-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.CloudKit": DomainPrefs{
			"AccountInfoValidationCounter": DomainPref{
				Type:    "int",
				Default: "236",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.cloudphotod": DomainPrefs{
			"_CPLUpgradeHistory-SystemLibrary": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CPLEngineParameters-SystemLibrary": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.CloudSubscriptionFeatures.cache": DomainPrefs{
			"mail.custom-domains.transfer": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"mail.hide-my-email.create": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"networking.privacy.attestation": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"networking.privacy.subscriber": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.CloudSubscriptionFeatures.geoCache": DomainPrefs{
			"001794-05-7e3c72e0-68de-4a6b-a51d-5ef3201ca25b-com.apple.systempreferences-mail.custom-domains.purchase": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"001794-05-7e3c72e0-68de-4a6b-a51d-5ef3201ca25b-com.apple.systempreferences-mail.hide-my-email.create": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"001794-05-7e3c72e0-68de-4a6b-a51d-5ef3201ca25b-com.apple.systempreferences-networking.privacy.subscriber": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.cmfsyncagent": DomainPrefs{
			"__kCMFBlockListStoreTopLevelKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ColorSyncCalibrator": DomainPrefs{
			"722524374X-expertMode": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"722524374X-name": DomainPref{
				Type:    "string",
				Default: "DELL U3223QE Calibrated",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"722524374X-shareMode": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"722524374X-targetTemp": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ColorSyncUtility": DomainPrefs{
			"CSProfsUtilsGroupBy": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.commcenter": DomainPrefs{
			"PhoneServices": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.CommCenter.counts": DomainPrefs{
			"CellularBytesRecved": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CellularBytesSent": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.commerce.knownclients": DomainPrefs{
			"com.apple.appstore:708": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.appstore:709": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.appstore:712": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.configurator:709": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.configurator:712": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.ibooks:708": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.ibooks:709": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.spotlight:708": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.spotlight:709": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.spotlight:712": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.fiplab.batteryhealth:712": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.multimarkdown.composer4.mac:712": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.coreauthd": DomainPrefs{
			"refKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.CoreDuet": DomainPrefs{
			"CloudSync2Policies": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudSyncPoliciesLastDownload": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudSyncPoliciesLastDownloadAttempt": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.CoreGraphics": DomainPrefs{
			"DisplayUseForcedGray": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DisplayUseInvertedPolarity": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.coreservices.uiagent": DomainPrefs{
			"CSUIHasSafariBeenLaunched": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CSUIHaveCenteredGatekeeperProgressWindow": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CSUILast32BitPromptDate_v2 for capicli": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CSUILast32BitPromptDate_v2 for com.charlessoft.pacifist": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CSUILast32BitPromptDate_v2 for com.tuffcode.HTTPScoop": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CSUILast32BitPromptDate_v2 for com.vmware.proxyApp.564d2494dc9a7951-62392db59de0d65c.564471774": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CSUILast32BitPromptDate_v2 for com.vmware.proxyApp.564d89fdf692a7b6-17f20d885122b12a.429132423": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CSUILast32BitPromptDate_v2 for com.yourcompany.Notify": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CSUILast32BitPromptDate_v2 for gen_snapshot": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CSUILast32BitPromptDate_v2 for net.nightproductions.prefSetter": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CSUILast32BitPromptDate_v2 for net.openvpn.uninstaller": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CSUILastOSVersionWhereSafariRecommendationWasMade": DomainPref{
				Type:    "string",
				Default: "10.15.6",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CSUIRecommendSafariBackOffInterval": DomainPref{
				Type:    "int",
				Default: "259200",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SGTRecentFileSearches": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.coreservices.useractivityd.dynamicuseractivites": DomainPrefs{
			"kUADynamicUserActivitiesKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.corespotlightui": DomainPrefs{
			"lastSyncDisabledBundleIds": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.dashboard": DomainPrefs{
			"devmode": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"mcx-disabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.desktopservices": DomainPrefs{
			"DSDontWriteNetworkStores": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.DFSLoginPlugin": DomainPrefs{
			"WindowLeft": DomainPref{
				Type:    "int",
				Default: "313",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WindowTop": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.DiagnosticExtensions.extensionTracker": DomainPrefs{
			"com.apple.DiagnosticExtensions.extensionTracker": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.digihub": DomainPrefs{
			"com.apple.digihub.blank.cd.appeared": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.DirectoryUtility": DomainPrefs{
			"FirstLaunch": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"perHost": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.DiskUtility": DomainPrefs{
			"advanced-image-options": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DUDebugMenuEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"OperationProgress DetailsVisible": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"OperationProgress ExpandedHeight": DomainPref{
				Type:    "int",
				Default: "488",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SidebarShowAllDevices": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.dock": DomainPrefs{
			"autohide": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"autohide-delay": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"largesize": DomainPref{
				Type:    "int",
				Default: "75",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"last-analytics-stamp": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"launchanim": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"loc": DomainPref{
				Type:    "string",
				Default: "en_US:US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"magnification": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"minimize-to-application": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"no-bouncing": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"persistent-apps": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"persistent-others": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"recent-apps": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"region": DomainPref{
				Type:    "string",
				Default: "US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"showAppExposeGestureEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"showhidden": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"showMissionControlGestureEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"tilesize": DomainPref{
				Type:    "int",
				Default: "70",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"trash-full": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"wvous-br-corner": DomainPref{
				Type:    "int",
				Default: "14",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"wvous-tr-corner": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"wvous-tr-modifier": DomainPref{
				Type:    "int",
				Default: "1048576",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.dock.external.extra.x86_64": DomainPrefs{
			"ShowSolidarity": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.donotdisturbd": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DNDSModeConfigurationsContactHistoryToken": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.driver.AppleBluetoothMultitouch.mouse": DomainPrefs{
			"MouseButtonDivision": DomainPref{
				Type:    "int",
				Default: "55",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseButtonMode": DomainPref{
				Type:    "string",
				Default: "TwoButton",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseHorizontalScroll": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseMomentumScroll": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseOneFingerDoubleTapGesture": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseTwoFingerDoubleTapGesture": DomainPref{
				Type:    "int",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseTwoFingerHorizSwipeGesture": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseVerticalScroll": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserPreferences": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.driver.AppleBluetoothMultitouch.trackpad": DomainPrefs{
			"Clicking": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Dragging": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DragLock": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ForceSuppressed": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"HIDScrollZoomModifierMask": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadCornerSecondaryClick": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFiveFingerPinchGesture": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFourFingerHorizSwipeGesture": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFourFingerPinchGesture": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFourFingerVertSwipeGesture": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadHandResting": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadHorizScroll": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadMomentumScroll": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadPinch": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadRightClick": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadRotate": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadScroll": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerDrag": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerHorizSwipeGesture": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerTapGesture": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerVertSwipeGesture": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadTwoFingerDoubleTapGesture": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadTwoFingerFromRightEdgeSwipeGesture": DomainPref{
				Type:    "int",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"USBMouseStopsTrackpad": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserPreferences": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.driver.AppleHIDMouse": DomainPrefs{
			"Button1": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Button2": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Button3": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Button4": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Button4Click": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Button4Force": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ButtonDominance": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ScrollH": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ScrollS": DomainPref{
				Type:    "int",
				Default: "4",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ScrollSSize": DomainPref{
				Type:    "int",
				Default: "30",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ScrollV": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.dt.Xcode": DomainPrefs{
			"DVTIgnoredDevices": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DVTSourceControlAccountDefaultsKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IBAppliesAutoResizingRulesWhileResizing": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IBPreferencesMigrated": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IDEDefaultPrimaryEditorFrameSizeForPaths": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IDEKeyBindingCurrentPreferenceSet": DomainPref{
				Type:    "string",
				Default: "Default.idekeybindings",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IDEKeyBindingExcludedPreferenceSet": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IDELastLSRegisterHashes": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IDERecentEditorDocuments": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IDESourceControlHostAccounts_10": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IDESourceControlKnownSSHHostsDefaultsKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IDEUserWantsToEnableDeveloperSystemPolicyMode": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"XCFontAndColorCurrentTheme": DomainPref{
				Type:    "string",
				Default: "Default (Light).xccolortheme",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.EmojiCache": DomainPrefs{
			"18F132_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"18G1012_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"18G103_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"18G2022_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"18G95_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"19H1030_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"19H114_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"19H1217_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"19H1323_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"19H1419_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"19H1519_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"19H1615_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"19H1715_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"19H1824_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"19H1922_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"19H2026_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"19H2_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"19H524_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"21G115_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"21G1974_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"21G320_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"21G419_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"21G725_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"21H1123_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"21H1320_en-US": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.EmojiPreferences": DomainPrefs{
			"EMFDefaultsKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.FaceTime": DomainPrefs{
			"CustomRingtone": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DidMigrateToSixteenByNineAspectRatio": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"kPromptEnableReadRecipts": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PhoneNumberUpgradeShown": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PNRltc": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PreferredVideoDeviceUID": DomainPref{
				Type:    "string",
				Default: "CC26373HVV3G1HNBN",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.facetime.bag": DomainPrefs{
			"CacheCertificate": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CachedBag": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CachedSignature": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CacheTime": DomainPref{
				Type:    "int",
				Default: "3600",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"URL": DomainPref{
				Type:    "string",
				Default: "https://init.ess.apple.com/WebObjects/VCInit.woa/wa/getBag?ix=3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.finder": DomainPrefs{
			"_FXShowPosixPathInTitle": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"_FXSortFoldersFirst": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"_FXSortFoldersFirstOnDesktop": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleShowAllFiles": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ComputerViewSettings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DesktopViewSettings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DisableAllAnimations": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FavoriteTagNames": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FK_AppCentricShowSidebar": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FK_ArrangeBy": DomainPref{
				Type:    "string",
				Default: "None",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FK_DefaultIconViewSettings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FK_iCloudIconViewSettings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FK_SearchIconViewSettings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FK_SidebarWidth": DomainPref{
				Type:    "int",
				Default: "204",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FK_StandardViewOptions2": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FK_StandardViewSettings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXArrangeGroupViewBy": DomainPref{
				Type:    "string",
				Default: "Name",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXAtLeastOneScreenHasBeenInHIDPI": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXConnectToLastURL": DomainPref{
				Type:    "string",
				Default: "smb://synologynas.local",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXDefaultSearchScope": DomainPref{
				Type:    "string",
				Default: "SCcf",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXDesktopTouchBarUpgradedToTenTwelveOne": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXDesktopVolumePositions": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXEnableExtensionChangeWarning": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudDriveDeclinedUpgrade": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudDriveDesktop": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudDriveDocuments": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudDriveEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudDriveFirstSyncDownComplete": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudLoggedIn": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXInfoPanesExpanded": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXInfoWindowWidth": DomainPref{
				Type:    "int",
				Default: "400",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXLastSearchScope": DomainPref{
				Type:    "string",
				Default: "SCcf",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXPreferredGroupBy": DomainPref{
				Type:    "string",
				Default: "Kind",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXPreferredSearchViewStyle": DomainPref{
				Type:    "string",
				Default: "Nlsv",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXPreferredViewStyle": DomainPref{
				Type:    "string",
				Default: "clmv",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXQuickActionsConfigUpgradeLevel": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXRecentFolders": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXSidebarUpgradedSharedSearchToTwelve": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXSidebarUpgradedToTenFourteen": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXSidebarUpgradedToTenTen": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXSyncExtensionToolbarItemsAutomaticallyAdded": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXToolbarUpgradedToTenEight": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXToolbarUpgradedToTenNine": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXToolbarUpgradedToTenSeven": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"GoToField": DomainPref{
				Type:    "string",
				Default: "/Volumes/Backups/2022-",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"GoToFieldHistory": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ICloudViewSettings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MeetingRoomViewSetting": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NetworkViewSettings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NewWindowTarget": DomainPref{
				Type:    "string",
				Default: "PfHm",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PackageViewSettings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PreviewPaneGalleryWidth": DomainPref{
				Type:    "int",
				Default: "267",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PreviewPaneInfoExpanded": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PreviewPaneWidth": DomainPref{
				Type:    "int",
				Default: "267",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"QLEnableTextSelection": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"QuitMenuItem": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"RecentMoveAndCopyDestinations": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"RecentsArrangeGroupViewBy": DomainPref{
				Type:    "string",
				Default: "Date Last Opened",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SearchRecentsSavedViewStyle": DomainPref{
				Type:    "string",
				Default: "Nlsv",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SearchRecentsViewSettings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SearchViewSettings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SGTRecentFileSearches": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowExternalHardDrivesOnDesktop": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowHardDrivesOnDesktop": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowMountedServersOnDesktop": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowPathbar": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowRemovableMediaOnDesktop": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowSidebar": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowStatusBar": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SidebarShowingiCloudDesktop": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SidebarShowingSignedIntoiCloud": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SidebarWidth": DomainPref{
				Type:    "int",
				Default: "256",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SlicesRootAttributes": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SmartSharedSearchViewSettings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"StandardViewOptions": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"StandardViewSettings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TagsCloudSerialNumber": DomainPref{
				Type:    "int",
				Default: "23",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrashViewSettings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ViewSettingsDictionary": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.findmy": DomainPrefs{
			"coarseServiceAcknowledged_v1.0": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"onboarding_v2.0": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.FolderActionsDispatcher": DomainPrefs{
			"folderActions": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"folderActionsEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.gamed": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"GKLastPushTokenDateKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"GKLastPushTokenEnvironment": DomainPref{
				Type:    "int",
				Default: "5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"GKLastPushTokenKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"GKLastPushTokenPlayerID": DomainPref{
				Type:    "string",
				Default: "U:097e62590e16c6eb3396c6958826eb60",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"GKLoginCancelled": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"GKURLPatternCacheBagURLUserDefaultsKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"GKURLPatternCacheExpirationDateUserDefaultsKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"natType": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.help": DomainPrefs{
			"RegisteredBooks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"RegistrationOrder": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.HIToolbox": DomainPrefs{
			"AppleCurrentKeyboardLayoutInputSourceID": DomainPref{
				Type:    "string",
				Default: "com.apple.keylayout.US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleDictationAutoEnable": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleEnabledInputSources": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleFnUsageType": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleSelectedInputSources": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.iApps": DomainPrefs{
			"PhotoBoothLibraries": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.iCal": DomainPrefs{
			"AccountDisplayOrder": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AllDayAreaHeight": DomainPref{
				Type:    "int",
				Default: "60",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AvailabilityShowTwentyFourHours": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"BirthdayEventsGenerationAttemptsToResetKey": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"BirthdayEventsGenerationLastLocale": DomainPref{
				Type:    "string",
				Default: "en_US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CalDefaultCalendar": DomainPref{
				Type:    "string",
				Default: "UseLastSelectedAsDefaultCalendar",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CalendarListMiniMonthVisibleMonths": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CalendarSidebarShown": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CalendarSidebarView": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CalendarSidebarWidth": DomainPref{
				Type:    "int",
				Default: "205",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CalForceTruthFileRestoreHashKey": DomainPref{
				Type:    "string",
				Default: "16A1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CollapsedTopLevelNodes": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DelegatesInSeparateWindows": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DisabledCalendars": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"display birthdays calendar": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"first minute of day time range": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IncludeDebugMenu": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"last calendar view description": DomainPref{
				Type:    "string",
				Default: "Monthly",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"last minute of day time range": DomainPref{
				Type:    "int",
				Default: "1440",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastViewsTimeZone": DomainPref{
				Type:    "string",
				Default: "America/New_York",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSDontMakeMainWindowKey": DomainPref{
				Type:    "string",
				Default: "NO",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSEventConcurrentProcessingEnabled": DomainPref{
				Type:    "string",
				Default: "YES",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"OwnerEmailAddress": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SelectedEvents": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Show Week Numbers": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TabbedWindowGroups": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TimeZone support enabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"view rects": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WR_DONT_ASK_FOR_DEFAULT": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.iChat": DomainPrefs{
			"AccountSortOrder": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AccountsToLogInAtLaunch": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"BuddyPictureSetToGenericByUser": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ChatWindowControllerUnifiedChatListWidth": DomainPref{
				Type:    "int",
				Default: "284",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ChatWindowControllerUnifiedContents": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ChatWindowControllerUnifiedIsVisible": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DaemonConnectionWaitTime": DomainPref{
				Type:    "int",
				Default: "15",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DidCheckForDuplicateChats": DomainPref{
				Type:    "int",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DidMarkGroupPhotosAsUnpurgeable": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DidMigratePersonCentricIDs": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DidRegenerateGroupID63841559": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DownloadsBookmark": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"HasPromptediMessageFTU": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"HasPromptedSMSRelay": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"imageBrowser.disableOpenGL": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KeepMessageForDays": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KeepMessagesVersionID": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NotifyAboutMessagesFromUnknownContacts": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PlaySoundsKey": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PMPrintingExpandedStateForPrint2": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SaveConversationsOnClose": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SODefaultTranscriptName": DomainPref{
				Type:    "string",
				Default: "ChatTranscriptViewController",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UnifiedChatListViewControllerAutomaticallySortsChats": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UnifiedChatWindowControllerSelectionGUIDSet": DomainPref{
				Type:    "string",
				Default: "SMS;-;22395",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WebIconDatabaseEnabled": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.iChat.AIM": DomainPrefs{
			"Accounts": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ActiveAccounts": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"OnlineAccounts": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.iChat.inputLine": DomainPrefs{
			"inputLineSettingsKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.iChat.Jabber": DomainPrefs{
			"Accounts": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ActiveAccounts": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"OnlineAccounts": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.icloud.fmip.clientconfiguration": DomainPrefs{
			"object": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.icloud.fmip.voiceassistantsync": DomainPrefs{
			"object": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.icloud.searchpartyuseragent": DomainPrefs{
			"baseDate.LocalBeaconingManager": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastLaunchLocale": DomainPref{
				Type:    "string",
				Default: "en",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastPruneDate.BeaconManagerService": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ids.deviceproperties": DomainPrefs{
			"6B23BF90-1A72-4FF8-B110-34B616AA2005": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CC2F9C5B-CFEA-4446-932F-1FEA359FC21D": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"D7F530C5-9B62-407E-86E8-06897F41B8BE": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"E8311575-FFDD-4240-92B3-64379B26E092": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Image_Capture": DomainPrefs{
			"com.apple.imagekit.cameraviewmode": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_Accessory_selectedTag": DomainPref{
				Type:    "int",
				Default: "1000",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_Camera_preferPostPocessingApp": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_Camera_selectedPathType": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_Camera_selectedTag": DomainPref{
				Type:    "int",
				Default: "1100",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_Camera_user-directories-0": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_lastUsedDeviceIsRemote": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_lastUsedDeviceName": DomainPref{
				Type:    "string",
				Default: "Mike's iPhone8",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_Scanner_downloadURL": DomainPref{
				Type:    "string",
				Default: "~/Pictures",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_Scanner_PostProcessApplication": DomainPref{
				Type:    "string",
				Default: "/Applications/Preview.app",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IKC_sort_ascending": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IKC_sort_key": DomainPref{
				Type:    "string",
				Default: "creationDate",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.imagecapture": DomainPrefs{
			"loggingLevel": DomainPref{
				Type:    "int",
				Default: "3073",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ImageCaptureService": DomainPrefs{
			"IK_Accessory_selectedTag": DomainPref{
				Type:    "int",
				Default: "1000",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.imagent": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMChatVocabularyUpdaterDidPerformInitialUpdateKey": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Setting.EnableReadReceipts": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Setting.GlobalReadReceiptsVersionID": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.imdpersistence.IMDPersistenceAgent": DomainPrefs{
			"LogAllIOErrors": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.imessage": DomainPrefs{
			"PreviewTranscodingQualityOnWiFiWasInitializedToDefaultValue": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PreviewTranscodingWasInitializedToDefaultValue": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.imessage.bag": DomainPrefs{
			"CacheCertificate": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CachedBag": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CachedSignature": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CacheTime": DomainPref{
				Type:    "int",
				Default: "6006",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"URL": DomainPref{
				Type:    "string",
				Default: "https://init-p01md.apple.com/bag?ix=2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.imservice.ids.FaceTime": DomainPrefs{
			"ActiveAccounts": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"OnlineAccounts": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.imservice.ids.iMessage": DomainPrefs{
			"ActiveAccounts": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"OnlineAccounts": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.imtranscoding.IMTranscoderAgent": DomainPrefs{
			"kCKMediaObjectManagerDefaultsClasses": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"kCKMediaObjectManagerDefaultsDynTypes": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"kCKMediaObjectManagerDefaultsUTITypes": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.internetconfigpriv": DomainPrefs{
			"DownloadFolder": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WWWHomePage": DomainPref{
				Type:    "string",
				Default: "http://livepage.apple.com/",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.internetconfigspec": DomainPrefs{
			"UserInformation": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WWWHomePage": DomainPref{
				Type:    "string",
				Default: "http://livepage.apple.com/",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ipTelephony": DomainPrefs{
			"ImsLoggingEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.iTunes": DomainPrefs{
			"AutomaticDeviceBackupsDisabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"bwui": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.iTunes.WebKit2AsynchronousSpellCheckingEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.iTunes.WebKit2DeveloperExtrasEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.iTunes.WebKit2EnableInheritURIQueryComponent": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.iTunes.WebKit2ShouldPrintBackgrounds": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.iTunes.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.iTunes.WebKit2SuppressesIncrementalRendering": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.iTunes.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Database Location": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"debugAssertCategoriesEnabled": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DeviceBackupsDisabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"disableRadio": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"disableShareLibraryInfo": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"doesAccountArtistListHaveSharePermission": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dontAutomaticallySyncIPods": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"gnot:1:Gracenote Match Registered User ID": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"gnot:2:Gracenote CDDB Lookup Registered User ID": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Home Sharing Settings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"isAccountAdmin": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"isAccountEnrolledInAppleMusic": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"isAccountEnrolledInITunesMatch": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"log-PlayQueue-LastSelectedTab": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"log-push": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MacBuddyStoreID": DomainPref{
				Type:    "string",
				Default: "mikeschinkel@gmail.com",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MobileDeviceSoftwareUpdateQueue": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PLGD:HIToolbox.rsrc•AAC Encoder": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"pref:130:Preferences": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"pref:400:Touch Remote Preferences": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"RDoc:132:Documents": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"recentSearches": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"recentSearchHints": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"rprf:0000000000000000": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"rspl:1:AirTunes Speaker List": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Store Apple ID": DomainPref{
				Type:    "string",
				Default: "mikeschinkel@gmail.com",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storefront": DomainPref{
				Type:    "string",
				Default: "143441-1,32",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"StoreUserInfo": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"stpr:294735135:Store Preferences": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.iTunes.eq": DomainPrefs{
			"eqps:129:EQPresets": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.itunescloud": DomainPrefs{
			"ICDefaultsCachedSubscriptionStatus": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ICDefaultsKeyLastCookieHeadersForRevokingMusicUserTokens": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ICDefaultsKeyLastKnownSubscriptionStatusBaseCacheKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ICDefaultsKeyMediaLibraryAccessApplicationIdentifiersWithTCCAcceptanceDates": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.java.util.prefs": DomainPrefs{
			"/": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.keyboard": DomainPrefs{
			"KeyboardWordOrSentenceTrackingForPFL": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.keyboardservicesd": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserDictionary": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.keychainaccess": DomainPrefs{
			"Item Preview Closed": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Keychain Item List Sort Descending": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Keychain Item List Sorting": DomainPref{
				Type:    "string",
				Default: "modified-date",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Keychains List Closed": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Last Selected Category": DomainPref{
				Type:    "string",
				Default: "All Items",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Last Selected Keychain": DomainPref{
				Type:    "string",
				Default: "/System/Library/Keychains/SystemRootCertificates.keychain",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"VPNSSItemsChecked": DomainPref{
				Type:    "string",
				Default: "Version 10.15.7 (Build 19H1715)",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.loginwindow": DomainPrefs{
			"MiniBuddyLaunch": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"oneTimeSSMigrationComplete": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TALLogoutSavesState": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.madrid": DomainPrefs{
			"CKMOCAccountsMatch": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitCheckedMiCSwitchEligibilityOnImagentLaunch": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitEligibleToToggleMiCSwitch": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitIsEligibleForTruthZone": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitIsRemovedFromBackup": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitIsSyncing": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitSyncingEnabled": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"createdChatZone": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableCKSyncingV2": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"hasCompletedInitialCKSync": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCKFinishedFetchingAttachmentsFromCloudKit": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitAccountStatusKey": DomainPref{
				Type:    "int",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitAnalyticSyncDatesDictionary": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitStartingDisableDevices": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitStartingEnabledSettingChange": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitStartingInitialSync": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitStartingPeriodicSync": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitSyncControllerSyncRecordTypeKey": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitSyncControllerSyncStateKey": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitSyncControllerSyncTypeKey": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitSyncPaused": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"initialSyncRecordHasBeenWritten": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"LastChatSyncTime": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"RequestPriorityRamp": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.mail": DomainPrefs{
			"NSUserKeyEquivalents": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.mediaaccessibility.public": DomainPrefs{
			"MACaptionDisplayType": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.mediaanalysisd": DomainPrefs{
			"PhotoLibraries": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PreheatLastRun": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.menuextra.battery": DomainPrefs{
			"ShowPercent": DomainPref{
				Type:    "string",
				Default: "YES",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.menuextra.clock": DomainPrefs{
			"DateFormat": DomainPref{
				Type:    "string",
				Default: "EEE MMM d  j:mm a",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Messages": DomainPrefs{
			"hasPerformedNewDeviceBringUpSync": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"hasTriedToEnableCKAndSyncAfterFirstImagentConnection": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.messages.facetime": DomainPrefs{
			"FaceTimeTwoTimeCallthroughEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.messages.nicknames": DomainPrefs{
			"MeCardSharingAudience": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MeCardSharingEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MeCardSharingImageForkedFromMeCard": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NicknameAppleIDAndiCloudAccountMatchAndExist": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NicknameInfoRequestedFromPeers": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NicknameScrutinyDoNotHandle": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.messages.pinning": DomainPrefs{
			"IMPinningPinConfigMigrationKey-v2-r2": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.messageshelper.MessageController": DomainPrefs{
			"SOInputLineSettings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.metrickitd": DomainPrefs{
			"MXClientDataRetrieved": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MXClientTeamID": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.mlruntimed": DomainPrefs{
			"LastCompletionDate-com.apple.NewsDES": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.MobileDeviceCrashCopy": DomainPrefs{
			"LastDeleted": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShouldPrompt": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShouldSubmit": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.MobileSMS": DomainPrefs{
			"BusinessChatPrivacyPageDisplayed": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CKCallPreferredServiceToIDMap": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CKEntryViewLayoutMetricsInfo": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CKMigratedAutoSpamReports26375262": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"kCKBrowserSelectionControllerIndexPathDictionaryV3Key": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"kCKBrowserSelectionControllerLaunchTimeDictionaryKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"kCKBrowserSelectionControllerSeenDictionaryKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"kCKBrowserSelectionControllerVersionDictionaryKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"kCKMediaObjectManagerDefaultsClasses": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"kCKMediaObjectManagerDefaultsDynTypes": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"kCKMediaObjectManagerDefaultsUTITypes": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KeepMessageForDays": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"kPluginSummaryTextDefaultsKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"LegacyAppSidebarPersistedWidth": DomainPref{
				Type:    "int",
				Default: "284",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NotifyAboutMessagesFromUnknownContacts": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"OnboardingFinishTime": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PendingCleared": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PlaySoundsKey": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"QuickSaveHasBeenUsedBefore": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ReadReceiptSettingsConfirmed": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Scrutiny": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sForceSMSSpamFilteringCompleted": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sForceUnknownFilteringCompleted": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SidebarPersistedWidth": DomainPref{
				Type:    "int",
				Default: "320",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TextFontSize": DomainPref{
				Type:    "int",
				Default: "13",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TextSize": DomainPref{
				Type:    "int",
				Default: "4",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.MobileSMSPreview": DomainPrefs{
			"IMPreviewGenerationMaxPxWidth": DomainPref{
				Type:    "int",
				Default: "300",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMPreviewGenerationMinHeight": DomainPref{
				Type:    "int",
				Default: "50",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMPreviewGenerationMinWidth": DomainPref{
				Type:    "int",
				Default: "50",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMPreviewGenerationScreenScale": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Music": DomainPrefs{
			"automaticallyDownloadArtwork": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"bwui": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"cddbPrefsOK": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.Music.WebKit2AsynchronousSpellCheckingEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.Music.WebKit2EnableInheritURIQueryComponent": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.Music.WebKit2ShouldPrintBackgrounds": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.Music.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.Music.WebKit2SuppressesIncrementalRendering": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.Music.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"debugAssertCategoriesEnabled": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"devp:2:Device Information": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"doesStoreSupportCloudMusicLibrary": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dontAskForAllPlaylistItemRemoval": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dontShowRestrictionsPrefs": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dontWarnDownloadCloudPurchases": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"encoderName": DomainPref{
				Type:    "string",
				Default: "AAC Encoder",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"eqEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"eqName": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"eqWindowHPos": DomainPref{
				Type:    "int",
				Default: "1059",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"eqWindowVPos": DomainPref{
				Type:    "int",
				Default: "617",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"gnot:1:Gracenote Match Registered User ID": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"gnot:2:Gracenote CDDB Lookup Registered User ID": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"gridViewSearchDoesNotSwitchView": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"haveAskedToCheckForAvailableDownloads": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Home Sharing Settings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"imported-eq-presets": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ITUserPrefsMigrated": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastMusicTabs": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"library-bookmark": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"log-PlayQueue-LastSelectedTab": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"media-folder-bookmark": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerHPos": DomainPref{
				Type:    "int",
				Default: "1722",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniPlayerLargeArtVisible": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniPlayerQueueVisible": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerSnapMode": DomainPref{
				Type:    "int",
				Default: "5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerUserSetHeight": DomainPref{
				Type:    "int",
				Default: "826",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerVPos": DomainPref{
				Type:    "int",
				Default: "-958",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerWidthInDesignCoords": DomainPref{
				Type:    "int",
				Default: "534",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"musicVideoPlaybackLocation": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"playbackIsFullscreen": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PLGD:Music•AAC Encoder": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:3d7cff22c30ce399:108cd281017e134d": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:3d7cff22c30ce399:23c80390c3230d0b": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:3d7cff22c30ce399:33cdb311c1d3000": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:3d7cff22c30ce399:39fb487bb789c556": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:3d7cff22c30ce399:5": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:3d7cff22c30ce399:5f95ba817db924bd": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:3d7cff22c30ce399:9b5cec37a177fa3e": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:3d7cff22c30ce399:ad9854d9acf570b3": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:3d7cff22c30ce399:ca10e523ad6d719c": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:3d7cff22c30ce399:cb975be86b377bd1": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:3d7cff22c30ce399:d7029e8cc09035fa": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:3d7cff22c30ce399:de054297cbb94475": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:3d7cff22c30ce399:e637bd9dae791bfe": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:3d7cff22c30ce399:edd0b99296e638db": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:3d7cff22c30ce399:ef3a85e10d48207": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:3d7cff22c30ce399:f2e02089ee363beb": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:3d7cff22c30ce399:f5a16fda0eb89f8e": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:3d7cff22c30ce399:fdf5a5350cf284": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"preloadFilesIntoMemory": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"RDoc:132:Documents": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"recentSearches": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"recentSearchHints": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"rprf:0000000000000000": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"rprf:00002074CF733A84": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"searchSavedTab": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sidebarInfo": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sidebarItemInfo": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"smartShuffleOptions": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"soundEnhancerAmount": DomainPref{
				Type:    "int",
				Default: "191",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"soundEnhancerEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeBrowseColumns": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeMediaTypeFlags": DomainPref{
				Type:    "int",
				Default: "3422",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storePublishedPlaylistColumns": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsCloudPurchases": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsPasswordSettings": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsPreviousPurchases": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsUPP": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"StoreUserInfo": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"updateLevel": DomainPref{
				Type:    "int",
				Default: "43",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userHasActivatedFullScreenVisualizer": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userMaxConcurrentDownloads": DomainPref{
				Type:    "int",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userWantsPlaybackNotifications": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"videoWindowHPos": DomainPref{
				Type:    "int",
				Default: "4025",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"videoWindowHSize": DomainPref{
				Type:    "int",
				Default: "3492",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"videoWindowVPos": DomainPref{
				Type:    "int",
				Default: "739",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"videoWindowVSize": DomainPref{
				Type:    "int",
				Default: "1964",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"viewOptionsWindowSectionConfig": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"volumeWSG": DomainPref{
				Type:    "float",
				Default: "0.1204547",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Music.eq": DomainPrefs{
			"eqps:129:EQPresets": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ncprefs": DomainPrefs{
			"apps": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"content_visibility": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dnd_prefs": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sort_order": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.NetAuthAgent": DomainPrefs{
			"PreviouslySelectedShares": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.networkConnect": DomainPrefs{
			"VPNShowTime": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.networkserviceproxy": DomainPrefs{
			"NSPConfiguration": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSPProxyAgentManagerPreferences": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSPServiceStatusManagerInfo": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSPSignatureInfo": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.NetworkUtility": DomainPrefs{
			"NUFingerPerson": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NULookupAddress": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NUNetstatChoice": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NUPingAddress": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NUPingChoice": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NUPortScanAddress": DomainPref{
				Type:    "string",
				Default: "10.10.10.51",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NUPortScanEnd": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NUPortScanRange": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NUPortScanStart": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NUSelectedTabItem": DomainPref{
				Type:    "string",
				Default: "NUTraceroute",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NUTracerouteAddress": DomainPref{
				Type:    "string",
				Default: "104.130.122.30",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NUWhoisAddress": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NUWhoisSelectedServer": DomainPref{
				Type:    "string",
				Default: "whois.internic.net",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NUWhoisServers": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.NewDeviceOutreach": DomainPrefs{
			"CachedWarrantyLocale": DomainPref{
				Type:    "string",
				Default: "en_US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CachedWarrantyValidityDuration": DomainPref{
				Type:    "int",
				Default: "28800",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"InitialOutreachActivityScheduled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PostUpgradeActivityCompleted": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PostUpgradeOSVersionKey": DomainPref{
				Type:    "string",
				Default: "12.7.6",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.newscore": DomainPrefs{
			"com.apple.news.default_event_endpoint": DomainPref{
				Type:    "string",
				Default: "https://news-events.apple.com/analyticseventsv2/async",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.news.notification_receipt_channel_data_dictionary": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.news.notification_receipt_event_app_data": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.news.notification_receipt_event_device_data": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.news.notification_receipt_event_endpoint": DomainPref{
				Type:    "string",
				Default: "https://news-notification-events.apple.com/analyticseventsv2/async",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.news.notification_receipt_event_session_data": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.news.notification_receipt_event_subscription_data": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.news.notification_receipt_event_user_data": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.news.notification_receipt_user_channel_context_data_dictionary": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.newsd.download.hasUnfinishedWork": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"force_refresh_user_segmentation": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"news_carplay_is_enabled": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"news_url_resolution_subscription_status": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"notificationArticleDiversificationSimilarityExpectationEnd": DomainPref{
				Type:    "float",
				Default: "0.1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"notificationArticleDiversificationSimilarityExpectationStart": DomainPref{
				Type:    "float",
				Default: "0.2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"notificationArticleDiversificationUniquePublisherExpectationSlope": DomainPref{
				Type:    "float",
				Default: "0.75",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"notificationArticleDiversificationUniquePublisherExpectationYIntercept": DomainPref{
				Type:    "float",
				Default: "0.25",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"notificationAssetPrefetchingRequiresWatch": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"notificationEnableAssetPrefetching": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"notificationPersonalizationTreatmentDictionary": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"use_parsec_results_for_widget": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.newsd": DomainPrefs{
			"FCAppConfigurationBundleShortVersionKey": DomainPref{
				Type:    "float",
				Default: "6.0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"last-known-requests": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.notificationcenterui": DomainPrefs{
			"bannerTime": DomainPref{
				Type:    "int",
				Default: "100",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TodayView": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.passd": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.PhotoBooth": DomainPrefs{
			"AVDesktopPlaybackControlsControllerShowsDurationInsteadOfTimeRemainingDefaultsKey": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"LibraryBookmark": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PBLaunchedAtLeastOnceOnLion": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserCameraUniqueIDPref": DomainPref{
				Type:    "string",
				Default: "CC26373HVV3G1HNBN",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.photos.shareddefaults": DomainPrefs{
			"clearPurgeableResources": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CPLDefaultDownload": DomainPref{
				Type:    "string",
				Default: "Default",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DiskSpaceWasLow": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"downloadAndKeepOriginals": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"iCloudPhotoLibraryHasBeenEnabledAccountIdentifiers": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"repushVideoAssetsMetadata": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Preferences": DomainPrefs{
			"UserDictionaryImportedSinceMaintenance": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserDictionaryUbiquityIdentityTokenCurrent": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserDictionaryUbiquityIdentityTokenLastKnown": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.extensions.CredentialProvider": DomainPrefs{
			"displayOrder": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userHasOrdered": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.extensions.FinderSync": DomainPrefs{
			"displayOrder": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userHasOrdered": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.extensions.PhotoEditing": DomainPrefs{
			"displayOrder": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userHasOrdered": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.extensions.PhotoProjects": DomainPrefs{
			"displayOrder": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userHasOrdered": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.extensions.QuickLook": DomainPrefs{
			"displayOrder": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userHasOrdered": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.extensions.ServicesWithUI": DomainPrefs{
			"displayOrder": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userHasOrdered": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.extensions.SharedLinks": DomainPrefs{
			"displayOrder": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userHasOrdered": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.extensions.ShareMenu": DomainPrefs{
			"displayOrder": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SHKRecentServices": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userHasOrdered": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.extensions.XcodeSourceEditor": DomainPrefs{
			"displayOrder": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userHasOrdered": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.softwareupdate": DomainPrefs{
			"ProductKeysLastSeenByUser": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.print.add": DomainPrefs{
			"defaultTableView": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.print.custompapers": DomainPrefs{
			"33 Ponce Mailbox Label": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Narrow": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Thin Borders": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.print.custompresets.forprinter.Canon_MF5900_Series": DomainPrefs{
			"com.apple.print.lastPresetPref": DomainPref{
				Type:    "string",
				Default: "Default Settings",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.print.lastPresetPrefType": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.print.lastUsedSettingsPref": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.print.PrinterProxy": DomainPrefs{
			"IK_Scanner_downloadURL": DomainPref{
				Type:    "string",
				Default: "~/Pictures",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_Scanner_PostProcessApplication": DomainPref{
				Type:    "string",
				Default: "/System/Applications/Preview.app",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_Scanner_selectedTag": DomainPref{
				Type:    "int",
				Default: "1000",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.proactive.PersonalizationPortrait": DomainPrefs{
			"entityBackfillTimestamp": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"entityMappingTrieSha256": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ProblemReporter": DomainPrefs{
			"ShowComments": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowDetails": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.PTPCamera": DomainPrefs{
			"FilenamePrefix": DomainPref{
				Type:    "string",
				Default: "IMG_",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ModuleSpecificPrefs": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.PubSubAgent": DomainPrefs{
			"TigerMigrationCompleted": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.rapport": DomainPrefs{
			"familySyncedName": DomainPref{
				Type:    "string",
				Default: "MacBookPro",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.remindd": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudConfigurationPath": DomainPref{
				Type:    "string",
				Default: "/System/Library/PrivateFrameworks/ReminderKit.framework/Versions/A/Resources/CloudConfigurations/Normal.plist",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitAccountStatus": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitZonesNeedingFetchChanges": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"isDatabaseMigrated": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastCloudConfigurationDownload": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastSuggestedAttributesAutoTrainingToken": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastSyncPoll": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SubscriptionIDs": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ThrottlingPolicyCurrentLevelIndex": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"useExtraneousAlarmBackOffThrottleInterval": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.remindd.babysitter": DomainPrefs{
			"FailedWaiters": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"RefreshingWaiters": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"RestrictedWaiters": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Safari": DomainPrefs{
			"cloudBookmarksMigrationEligibilityDataInvalidated": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.Safari.ContentPageGroupIdentifier.WebKit2DeveloperExtrasEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IncludeDebugMenu": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IncludeDevelopMenu": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IncludeInternalDebugMenu": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ResetCloudHistory": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowFullURLInSmartSearchField": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SuppressSearchSuggestions": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UniversalSearchEnabled": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WebKitDeveloperExtrasEnabledPreferenceKey": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WebKitHistoryAgeInDaysLimit": DomainPref{
				Type:    "int",
				Default: "360",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WebKitHistoryItemLimit": DomainPref{
				Type:    "int",
				Default: "9999",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WebKitInitialTimedLayoutDelay": DomainPref{
				Type:    "float",
				Default: "0.25",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Safari.SafeBrowsing": DomainPrefs{
			"HasMigratedSafeBrowsingEnabledDefaults": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Safari.SandboxBroker": DomainPrefs{
			"DidMigrateDownloadFolderToSandbox": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DidMigrateDownloadMetadataToSandboxGroupContainer": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DidMigrateResourcesToSandbox": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Homepage": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowDevelopMenu": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.SafariBookmarksSyncAgent": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"LastDatabaseServerChangeToken.com.apple.SafariShared.CloudTabs": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"LastSavedAccountHash": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MigrationStateEncodedRecordData": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.SafariCloudHistoryPushAgent": DomainPrefs{
			"AcknowledgedPushNotifications": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UnacknowledgedPushNotifications": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.SafariServices": DomainPrefs{
			"SearchProviderIdentifierMigratedToSystemPreference": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.scheduler": DomainPrefs{
			"AbsoluteSchedule": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.screencapture": DomainPrefs{
			"captureDelay": DomainPref{
				Type:    "int",
				Default: "5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"last-selection": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"last-selection-display": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"location": DomainPref{
				Type:    "string",
				Default: "/Users/mikeschinkel/Desktop/Screenshots",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"style": DomainPref{
				Type:    "string",
				Default: "display",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"video": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.screencaptureui": DomainPrefs{
			"com.apple.AnnotationKit.arrowHeadStyle": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AnnotationKit.brushStyle": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AnnotationKit.fillColor": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AnnotationKit.font": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AnnotationKit.hasShadow": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AnnotationKit.highlightStyle": DomainPref{
				Type:    "int",
				Default: "765200",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AnnotationKit.strokeColor": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AnnotationKit.strokeIsDashed": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AnnotationKit.strokeWidth": DomainPref{
				Type:    "int",
				Default: "5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AnnotationKit.textAttributes": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.screensaver": DomainPrefs{
			"tokenRemovalAction": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.scriptmenu": DomainPrefs{
			"ScriptMenuEnabled": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.security.ctkd-db": DomainPrefs{
			"classes": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"tokens": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ServicesMenu.Services": DomainPrefs{
			"NSServices": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.sharingd": DomainPrefs{
			"AutoUnlockPresentedBluetoothError": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AutoUnlockPresentedWiFiError": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AutoUnlockWatchCurrentlyInList": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AutoUnlockWatchExistedInUnlockList": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DiscoverableMode": DomainPref{
				Type:    "string",
				Default: "Contacts Only",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"HashManager-LastConsumedHistoryTokenKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"HashManager-LastRebuiltDateKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"HashManager-LastUpdatedDateKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"HashManager-StoredDatabaseVersionKey": DomainPref{
				Type:    "int",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"OneTimeAirDropReset": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"OneTimeAirDropReset2": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.shortcuts": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MainViewLayoutMode": DomainPref{
				Type:    "unknown",
				Default: "grid",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSUserKeyEquivalents": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Siri": DomainPrefs{
			"StatusMenuVisible": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"VoiceTriggerUserEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.siri-distributed-evaluation": DomainPrefs{
			"LastCompletionDate-com.apple.dodml.TestPlugin": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"LastCompletionDate-com.apple.fides.borealis": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"LastCompletionDate-com.apple.NewsDES": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"LastCompletionDate-com.apple.proactive.PMLDESPlugin": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"LastCompletionDate-com.apple.reminders.des": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"LastCompletionDate-com.apple.siri.SiriAudioDESPlugin": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"LastCompletionDate-com.apple.stocks.des": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.siri.DialogEngine": DomainPrefs{
			"dateTime": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dateTime#hoursMinutesSeconds": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.siri.embeddedspeech": DomainPrefs{
			"InstalledLanguages": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.siri.morphun": DomainPrefs{
			"subscriptions": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.siri.shortcuts": DomainPrefs{
			"LegacyShortcutsZoneSubscriptionUnsubscribed": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WFDidUnconflictShortcuts": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WFKnownHomeSceneNames": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WFRecentDocuments": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WFServicesShortcutsMigrated": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WFSidebarAdditionalDestinations": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WFSidebarHiddenWorkflowTypes": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.siri.VoiceShortcuts": DomainPrefs{
			"VCLSDataSequenceKey": DomainPref{
				Type:    "int",
				Default: "63392",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.siriactionsd": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.sms": DomainPrefs{
			"hasBeenApprovedForSMSRelay": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.SocialLayer": DomainPrefs{
			"SharedWithYouApps": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SharedWithYouEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.sociallayerd": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.sociallayerd.CloudKit.ckwriter": DomainPrefs{
			"accountAvailable": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"accountLastKnownUserRecordID": DomainPref{
				Type:    "string",
				Default: "_4d72b079e6dd351d59dcdf0b9afd41d6",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastKnownApps": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"syncEngineMetadata": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"writerDone": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.SoftwareUpdate": DomainPrefs{
			"AutoUpdateMajorOSVersion": DomainPref{
				Type:    "float",
				Default: "10.15",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.spaces": DomainPrefs{
			"app-bindings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SpacesDisplayConfiguration": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.speech.recognition.AppleSpeechRecognition.CustomCommands": DomainPrefs{
			"Accessibility.DisableCommandAndControl": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.CursorMoveDownPixelsWithCardinalNumber": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.CursorMoveLeftPixelsWithCardinalNumber": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.CursorMoveRightPixelsWithCardinalNumber": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.CursorMoveUpPixelsWithCardinalNumber": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.CursorTriplePress": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.DecrementItem": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.DecrementItemUsingCardinalNumber": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.DragFromItemToItem": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.EndDragAtItem": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.Find": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.IncrementItem": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.IncrementItemUsingCardinalNumber": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.OverlayDecrementLabel": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.OverlayDecrementLabelUsingCardinalNumber": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.OverlayIncrementLabel": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.OverlayIncrementLabelUsingCardinalNumber": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.OverlayMoveCursor": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.OverlayTriplePress": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.PressAndHoldAtItem": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.PressItemWithoutPrefix": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.SetValueByPercentAtItem": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.StartCommandSearch": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"System.StartDragAtItem": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.speech.recognition.AppleSpeechRecognition.prefs": DomainPrefs{
			"DictationIMCommandCounts": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DictationIMCommandsWindowIsOpen": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DictationIMLocaleIdentifier": DomainPref{
				Type:    "string",
				Default: "en_US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DictationIMMessageTracesSinceLastReport": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DictationIMPlaySoundUponRecognition": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DictationIMPresentedOfflineUpgradeSuggestion": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DictationIMTargetApplications": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DictationIMUpgradedTo10_15": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DictationIMUpgradedTo10_16": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DictationIMUseOnlyOfflineDictation": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DictationIMUserHintsHistory": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SpeechDataInstallationLog": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.speech.voice.prefs": DomainPrefs{
			"SpeechDataInstallationLog": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SpeechDataPurgeList": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"VoiceStatistics": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.SpeechRecognitionCore": DomainPrefs{
			"AllowAudioDucking": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.SpeechRecognitionCore.speechrecognitiond": DomainPrefs{
			"InstalledLanguages": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Spotlight": DomainPrefs{
			"__NSEnableTSMDocumentWindowLevel": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"engagementCount-com.apple.mail": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"engagementCount-com.apple.Spotlight": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"engagementCount-com.apple.Spotlight.suggestions": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ModelName": DomainPref{
				Type:    "string",
				Default: "Canned",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"orderedItems": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"reloadShortcuts": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"showedFTE": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userHasMovedWindow": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"windowHeight": DomainPref{
				Type:    "int",
				Default: "430",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.StatusKitAgent": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastPublishTime": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastReauthTime": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.stockholm": DomainPrefs{
			"PKWalletShouldAutomaticallyRegisterKey": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.suggestd": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.suggestions": DomainPrefs{
			"AppCanShowSiriSuggestionsBlacklist": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"findToShowMigrationPerformed": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"spToLearnMigrationPerformed": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SuggestionsAllowGeocode": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.symbolichotkeys": DomainPrefs{
			"AppleSymbolicHotKeys": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.syncdefaultsd": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.systempreferences": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.preferences.accounts.outline.usersparent": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.preferences.energysaver": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.preferences.keyboard.text.sortDescriptors": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.preferences.sharing.selectedservice": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.PreferenceSync.ExcludeSyncKeys": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.SecurityPref.Privacy.LastSourceSelected": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DidShowPrefBundleIDs": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DisableAutoLoginButtonIsHidden": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DSKDesktopPrefPane": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSFullScreenMenuItemEverywhere": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSOutlineView Items DSKSourceOfPicturesOutlineView": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSQuitAlwaysKeepsWindows": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SecurityPrefTab": DomainPref{
				Type:    "string",
				Default: "General",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SGTRecentFileSearches": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowAllMode": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SoundTab": DomainPref{
				Type:    "string",
				Default: "output",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.SystemProfiler": DomainPrefs{
			"CPU Names": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.talagent": DomainPrefs{
			"LastKeyChange": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.TelephonyUtilities": DomainPrefs{
			"activeConversationsListKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FaceTimePhotosEnabled": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"registeredProviders": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"relayCallingDisabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"thumperCallingCapabilitiesStates": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.TelephonyUtilities.sharePlayAppPolicies": DomainPrefs{
			"CPAppPolicy.ContactsChangeHistoryToken": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Terminal": DomainPrefs{
			"Default Window Settings": DomainPref{
				Type:    "string",
				Default: "Basic",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"HasMigratedDefaults": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SecureKeyboardEntry": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Startup Window Settings": DomainPref{
				Type:    "string",
				Default: "Basic",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TTAppPreferences Selected Tab": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Window Settings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.TextEdit": DomainPrefs{
			"PlainTextEncoding": DomainPref{
				Type:    "int",
				Default: "4",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PlainTextEncodingForWrite": DomainPref{
				Type:    "int",
				Default: "4",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"RichText": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.textInput.keyboardServices.textReplacement": DomainPrefs{
			"_KSTRCloudKitMigratable": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"configPullTime": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"gmMigrationPercent": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"gmMigrationPercent1": DomainPref{
				Type:    "int",
				Default: "999",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"internalMigrationPercent": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"internalMigrationPercent1": DomainPref{
				Type:    "int",
				Default: "999",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"iOSMajorVersForCloudKitSync": DomainPref{
				Type:    "int",
				Default: "11",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"iOSMinorVersForCloudKitSync": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"kDidMigrateToUUIDRecordName": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSCKDidSetupRecordZoneProd-TextReplacements": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSCKSubscriptionProd-TextReplacements": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSCloudKitMigrationDidComplete": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSDidMigrateToCloudKitOnCloud": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSDidPullLegacyEntriesFromPeers": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSDidPushAllLocalRecordsOnce-2": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSDidPushMigrationStatusOnce-2": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSLastKnownUserID-TextReplacements": DomainPref{
				Type:    "string",
				Default: "_694927be690dcbcc4d7ebaccd1c628c5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSSampleShortcutWasImported_CK": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"kTRCKLastSyncTime": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"osxMajorVersForCloudKitSync": DomainPref{
				Type:    "int",
				Default: "10",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"osxMinorSubversionForCloudKitSync": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"osxMinorVersForCloudKitSync": DomainPref{
				Type:    "int",
				Default: "13",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"seedMigrationPercent": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"seedMigrationPercent1": DomainPref{
				Type:    "int",
				Default: "999",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TRCKSyncCountHalflifeInHours": DomainPref{
				Type:    "int",
				Default: "10",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TRCKSyncMaxCountThreshold": DomainPref{
				Type:    "int",
				Default: "200",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userDidFallInMigrationAllowBatch": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"whitelistedMigrationAccounts": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.TextInputMenu": DomainPrefs{
			"visible": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.tipsd": DomainPrefs{
			"DeliveryEventIdentifierToDeliveryEventMap": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DiscoverabilityDeliveryInfoMap": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Experiment": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TPSAnalyticsHistoricTipsDataProcessor_DateLastRun": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TPSAnalyticsStore": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TPSAnalyticsUsageEventsProcessor_DateLastRun": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TPSDCloudHintDisplayed": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TPSTipsAppInstalled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.TMHelperAgent": DomainPrefs{
			"LastNotificationDates": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.touristd": DomainPrefs{
			"firstOSLogin": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.translationd": DomainPrefs{
			"InstalledLocales": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.triald": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.TTY": DomainPrefs{
			"IncomingCallsTTYPreference": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"RTTContinuityRTTIsSupportedPreference": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"RTTSettingsVersionPreference": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TTYHardwareEnabledPreference": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TTYShouldBeRealtimePreference": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TTYSoftwareEnabledPreference": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TUIsRelayCallingEnabledPreference": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TUSupportsRelayCallingPreference": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.TV": DomainPrefs{
			"automaticallyDeleteVideoAssetsAfterWatching": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"automaticallyDownloadArtwork": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"bwui": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"cddbPrefsOK": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"checkForAvailableDownloads": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.TV.WebKit2AsynchronousSpellCheckingEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.TV.WebKit2EnableInheritURIQueryComponent": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.TV.WebKit2ShouldPrintBackgrounds": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.TV.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.TV.WebKit2SuppressesIncrementalRendering": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.TV.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"debugAssertCategoriesEnabled": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dontAskForAllPlaylistItemRemoval": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"gridViewSearchDoesNotSwitchView": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"haveAskedToCheckForAvailableDownloads": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Home Sharing Settings": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ITUserPrefsMigrated": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastViewdLibraryViewMode": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"library-bookmark": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"media-folder-bookmark": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerHPos": DomainPref{
				Type:    "int",
				Default: "2723",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniPlayerLargeArtVisible": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniPlayerQueueVisible": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerSnapMode": DomainPref{
				Type:    "int",
				Default: "5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerUserSetHeight": DomainPref{
				Type:    "int",
				Default: "771",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerVPos": DomainPref{
				Type:    "int",
				Default: "821",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerWidthInDesignCoords": DomainPref{
				Type:    "int",
				Default: "288",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"musicVideoPlaybackLocation": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"playbackIsFullscreen": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:a2de93531b7e1bda:3884c713d2a16431": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:a2de93531b7e1bda:4192905156c353e3": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:a2de93531b7e1bda:5": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:a2de93531b7e1bda:66852efdd705ebec": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:a2de93531b7e1bda:97d81438db99e819": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:a2de93531b7e1bda:ceb4675ad60ea843": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:a2de93531b7e1bda:ed5ef2132826af84": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PPr4:LIB:a2de93531b7e1bda:f8b16edff78b2283": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"preloadFilesIntoMemory": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"RDoc:132:Documents": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"recentSearches": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"recentSearchHints": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sidebarInfo": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"smartShuffleOptions": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeBrowseColumns": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeMediaTypeFlags": DomainPref{
				Type:    "int",
				Default: "3422",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storePublishedPlaylistColumns": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsCloudPurchases": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsPasswordSettings": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsPreviousPurchases": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsUPP": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"tabBarTabs": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"tabViewMode": DomainPref{
				Type:    "int",
				Default: "22",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"updateLevel": DomainPref{
				Type:    "int",
				Default: "43",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userHasActivatedFullScreenVisualizer": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userMaxConcurrentDownloads": DomainPref{
				Type:    "int",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userWantsPlaybackNotifications": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"videoWindowHPos": DomainPref{
				Type:    "int",
				Default: "461",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"videoWindowHSize": DomainPref{
				Type:    "int",
				Default: "1906",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"videoWindowVPos": DomainPref{
				Type:    "int",
				Default: "139",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"videoWindowVSize": DomainPref{
				Type:    "int",
				Default: "1072",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"viewOptionsWindowSectionConfig": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.typographypanel": DomainPrefs{
			"NSTypographyPanel_Monaco": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.UIKit": DomainPrefs{
			"hasAccessibilityBeenMigrated": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.universalaccess": DomainPrefs{
			"AssistiveControlType": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"axShortcutExposedFeatures": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"closeViewDesiredZoomFactor": DomainPref{
				Type:    "float",
				Default: "1.1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"closeViewHotkeysEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"closeViewScrollWheelToggle": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"closeViewSplitScreenRatio": DomainPref{
				Type:    "float",
				Default: "0.2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"closeViewZoomDisplayID": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"closeViewZoomedIn": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"closeViewZoomFactor": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"closeViewZoomFactorBeforeTermination": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.custommenu.apps": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"contrast": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"differentiateWithoutColor": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dwellEnabled": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dwellHideUITimeout": DomainPref{
				Type:    "int",
				Default: "15",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"grayscale": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"headMouseEnabled": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"hoverTextEnabled": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"hoverTextIsAlwaysOn": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"hoverTextIsHoveringAndVisible": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"increaseContrast": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"keyboardAccessFocusRingTimeout": DomainPref{
				Type:    "int",
				Default: "15",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastNightShiftActive": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastNightShiftEnabled": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastNightShiftMode": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"login": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"mouseDriver": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"mouseDriverCursorSize": DomainPref{
				Type:    "float",
				Default: "1.44109",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"reduceTransparency": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"selectedTab": DomainPref{
				Type:    "int",
				Default: "14",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sessionChange": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"slowKey": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"slowKeyDelay": DomainPref{
				Type:    "int",
				Default: "250",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"stickyKey": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"switchAutoScanElementInterval": DomainPref{
				Type:    "float",
				Default: "0.5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"switchAutoScanPanelInterval": DomainPref{
				Type:    "float",
				Default: "0.5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"switchCoalescePressesDuration": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"switchFirstElementDelay": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"switchHideUITimeout": DomainPref{
				Type:    "int",
				Default: "15",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"switchHoldBeforeRepeatDuration": DomainPref{
				Type:    "int",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"switchMinimumPressDuration": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"switchOnOffKey": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"switchSweepingCursorSpeed": DomainPref{
				Type:    "int",
				Default: "5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserAssignableHotKeys": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"useStickyKeysShortcutKeys": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"virtualKeyboardOnOff": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"voiceOverOnOffKey": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"whiteOnBlack": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.UserAccountUpdater": DomainPrefs{
			"com.apple.HidLibraryFolderAlready": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.visualintelligence": DomainPrefs{
			"bypassAvailability": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"debugUI": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableAlbumsDomain": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableArtDomain": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableBooksDomain": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableCoarseClassification": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableLandmarkDomain": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableNatureDomain": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enablePetsDomain": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enablePhotosApp": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableQuickLook": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableSafariApp": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableScreenshots": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"firstTimeExperience": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"initialized": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sendLocationInfo": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sendOCRText": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.voicememod": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.VoiceOver": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.VoiceOver4.local": DomainPrefs{
			"SCRCUserDefaultsAllowAirPlay": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SCRCUserDefaultsPlannedShutdownSuccessful": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.VoiceOverTraining": DomainPrefs{
			"doNotShowSplashScreen": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.voicetrigger": DomainPrefs{
			"VoiceTrigger Enabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.weather": DomainPrefs{
			"cachedAppConfig": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.weather.internal": DomainPrefs{
			"cachedAppConfig": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.WorkflowKit.BackgroundShortcutRunner": DomainPrefs{
			"CKPerBootTasks": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"GlobalPreferences": DomainPrefs{
			"646F6E7A_00000000_00000001_6E7A6361_696D6963": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AKLastEmailListRequestDateKey": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AKLastIDMSEnvironment": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleActionOnDoubleClick": DomainPref{
				Type:    "string",
				Default: "Maximize",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleAntiAliasingThreshold": DomainPref{
				Type:    "int",
				Default: "4",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleAquaColorVariant": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleEnableSwipeNavigateWithScrolls": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleKeyboardUIMode": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleLanguages": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleLanguagesDidMigrate": DomainPref{
				Type:    "string",
				Default: "12.7.6",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleLocale": DomainPref{
				Type:    "string",
				Default: "en_US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleMeasurementUnits": DomainPref{
				Type:    "string",
				Default: "Inches",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleMetricUnits": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleMiniaturizeOnDoubleClick": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleShowAllExtensions": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleShowScrollBars": DomainPref{
				Type:    "string",
				Default: "Always",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleTemperatureUnit": DomainPref{
				Type:    "string",
				Default: "Fahrenheit",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleWindowTabbingMode": DomainPref{
				Type:    "string",
				Default: "always",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AudioQuest inc. AudioQuest DragonFly": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Blue Microphones Yeti Stereo Microphone": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Canon MF5900 Series": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.finder.SyncExtensions": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.mouse.scaling": DomainPref{
				Type:    "int",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.sound.beep.flash": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.springing.delay": DomainPref{
				Type:    "float",
				Default: "0.5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.springing.enabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.swipescrolldirection": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.trackpad.forceClick": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.trackpad.scaling": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ContextMenuGesture": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Fujitsu ScanSnap iX500": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"InitialKeyRepeat": DomainPref{
				Type:    "int",
				Default: "15",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KeyRepeat": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NavPanelFileListModeForOpenMode": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NavPanelFileListModeForSaveMode": DomainPref{
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSAutomaticCapitalizationEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSAutomaticDashSubstitutionEnabled": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSAutomaticPeriodSubstitutionEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSAutomaticQuoteSubstitutionEnabled": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSAutomaticSpellingCorrectionEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSAutomaticTextCompletionEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSAutomaticWindowAnimationsEnabled": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSLinguisticDataAssetsRequested": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSLinguisticDataAssetsRequestedByChecker": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSLinguisticDataAssetsRequestLastInterval": DomainPref{
				Type:    "int",
				Default: "86400",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSLinguisticDataAssetsRequestTime": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSPreferredWebServices": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSUserDictionaryReplacementItems": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"QLPanelAnimationDuration": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WebAutomaticSpellingCorrectionEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WebKitDeveloperExtras": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Yubico YubiKey OTP+FIDO+CCID": DomainPref{
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
	}
}
