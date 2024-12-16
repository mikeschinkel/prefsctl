package prefdefaults

//goland:noinspection SpellCheckingInspection
func montereyPrefDefaults() DomainDefaults {
	return DomainDefaults{
		"com.apple.Accessibility": DomainPrefs{
			"AccessibilityEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ApplicationAccessibilityEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AutomationEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AXSClassicInvertColorsPreference": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"BrailleInputDeviceConnected": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CommandAndControlEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DarkenSystemColors": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DifferentiateWithoutColor": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"EnhancedBackgroundContrastEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FullKeyboardAccessEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FullKeyboardAccessFocusRingEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"GenericAccessibilityClientEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"GrayscaleDisplay": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"InvertColorsEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KeyRepeatDelay": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "0.25",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KeyRepeatEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ReduceMotionEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SpeakThisEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"VoiceOverTouchEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ZoomTouchEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.accounts": DomainPrefs{
			"ACOneTimeLDAPFixAccountSync": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ACOneTimeSMTPFixAccountSync": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ActivityMonitor": DomainPrefs{
			"OpenMainWindow": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SelectedTab": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowCategory": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "100",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowCategoryAppsinLast12Hours": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.AddressBook": DomainPrefs{
			"AB21vCardEncoding": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "macintosh",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ABBirthDayVisible": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ABLastImportShown": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ABPrivateVCardFieldsEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ABTextSizeIncrement": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSPreferencesSelectedIndex": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.airplay": DomainPrefs{
			"showInMenuBarIfPresent": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.airport.airportutility": DomainPrefs{
			"dontPerformBaseRestartWarning": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.amp.mediasharingd": DomainPrefs{
			"home-sharing-enabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"home-sharing-sequence-id": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "1183",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"photo-sharing-enabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"public-sharing-enabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"public-sharing-is-protected": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"shared-library-name": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Mike’s MacBook Pro Library",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.AMPDevicesAgent": DomainPrefs{
			"AutomaticDeviceBackupsDisabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"debugAssertCategoriesEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dontAutomaticallySyncIPods": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ITUserPrefsMigrated": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeMediaTypeFlags": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3422",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsCloudPurchases": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"updateLevel": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "43",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userMaxConcurrentDownloads": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.AMPLibraryAgent": DomainPrefs{
			"books-migrated": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"debugAssertCategoriesEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"have-shown-cloud-UI": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"imported-media-domains": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "15",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"iTunes-media-folder-url": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "file:///Users/mikeschinkel/Music/iTunes/iTunes%20Media/",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ITUserPrefsMigrated": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"podcasts-migrated": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"show-home-sharing-turned-on-in-iTunes-warning": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeMediaTypeFlags": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3422",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsCloudPurchases": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsPasswordSettings": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsPreviousPurchases": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsUPP": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"updateLevel": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "43",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.AppleMultitouchMouse": DomainPrefs{
			"MouseButtonDivision": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "55",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseButtonMode": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "TwoButton",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseHorizontalScroll": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseMomentumScroll": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseOneFingerDoubleTapGesture": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseTwoFingerDoubleTapGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseTwoFingerHorizSwipeGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseVerticalScroll": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserPreferences": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.AppleMultitouchTrackpad": DomainPrefs{
			"ActuateDetents": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ActuationStrength": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Clicking": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Dragging": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DragLock": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FirstClickThreshold": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ForceSuppressed": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"HIDScrollZoomModifierMask": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SecondClickThreshold": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadCornerSecondaryClick": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFiveFingerPinchGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFourFingerHorizSwipeGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFourFingerPinchGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFourFingerVertSwipeGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadHandResting": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadHorizScroll": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadMomentumScroll": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadPinch": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadRightClick": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadRotate": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadScroll": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerDrag": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerHorizSwipeGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerTapGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerVertSwipeGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadTwoFingerDoubleTapGesture": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadTwoFingerFromRightEdgeSwipeGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"USBMouseStopsTrackpad": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserPreferences": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.AppStore": DomainPrefs{
			"AutoPlayVideoSetting": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "on",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AVDesktopPlaybackControlsControllerShowsDurationInsteadOfTimeRemainingDefaultsKey": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AppStore.WebKit2DeveloperExtrasEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AppStore.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AppStore.WebKit2SuppressesIncrementalRendering": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AppStore.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastBootstrapTimeZone": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "America/New_York",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserSetAutoPlayVideoSetting": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.assistant": DomainPrefs{
			"Has Set Up Account Status Subscription": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Has Set Up Key Value Subscription": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Has Set Up MultiUser Shared Record Subscription": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Has Set Up Voice Trigger Subscription": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Initial Interstitial Delay": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "3.5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Initial Interstitial Delay - stark": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Initial Interstitial Delay for VoiceTrigger One-Shot without Audio Ducking": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "1.75",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Manual Endpointing Threshold": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MultiUser Shared Data Needs Sync": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Myriad Device Adjust": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "-10",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Myriad Device Class": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "9",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Myriad Device Delay": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Myriad Device Trump Delay": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "1.75",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Myriad VTEndtimeDistanceThreshold": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "0.5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Server Has Provisioned Myriad": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Voice Trigger Needs Sync": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.assistant.backedup": DomainPrefs{
			"Cloud Sync Enabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Cloud Sync User ID": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "_44c58d2cbdc7b03bb2cdefc57dcb364d",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Hands Free Mode": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MultiUser VoiceIdentification Enabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Session Language": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "en-US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.assistant.support": DomainPrefs{
			"Assistant Enabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Dictation Enabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Siri Data Sharing Opt-In Status": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.AssistiveControl.virtualKeyboard": DomainPrefs{
			"DesiredPanelWindowPosition": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Absolute",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PanelPosition": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Absolute",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SCLaunchedAsSlave": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Automator": DomainPrefs{
			"AMLogTabViewSelectedIndex": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.AvatarUI.Staryu": DomainPrefs{
			"AVTAvatarHasDisplayedAnimojiSplashScreen": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AVTAvatarHasDisplayedCameraEffectsSplashScreen": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AVTAvatarHasDisplayedPaddleView": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AVTAvatarHasDisplayedSplashScreen": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.CharacterPaletteIM": DomainPrefs{
			"CVFontSize": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "26",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CVStartAsLargeWindow": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.classroom": DomainPrefs{
			"allowClassroomLockDevice-00000000-0000-0000-0000-000000000000:0": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomLockDevice-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenApp": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenApp-00000000-0000-0000-0000-000000000000:0": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenApp-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenApp-ask": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenURL": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenURL-00000000-0000-0000-0000-000000000000:0": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenURL-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenURL-ask": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomScreenObservation-00000000-0000-0000-0000-000000000000:0": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomScreenObservation-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"forceUnpromptedRemoteScreenObservation-00000000-0000-0000-0000-000000000000:0": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"forceUnpromptedRemoteScreenObservation-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.CloudKit": DomainPrefs{
			"AccountInfoValidationCounter": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "236",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ColorSyncCalibrator": DomainPrefs{
			"722524374X-expertMode": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"722524374X-name": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "DELL U3223QE Calibrated",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"722524374X-shareMode": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"722524374X-targetTemp": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ColorSyncUtility": DomainPrefs{
			"CSProfsUtilsGroupBy": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.CoreGraphics": DomainPrefs{
			"DisplayUseForcedGray": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DisplayUseInvertedPolarity": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.coreservices.uiagent": DomainPrefs{
			"CSUIHasSafariBeenLaunched": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CSUIHaveCenteredGatekeeperProgressWindow": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CSUILastOSVersionWhereSafariRecommendationWasMade": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "10.15.6",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CSUIRecommendSafariBackOffInterval": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "259200",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.dashboard": DomainPrefs{
			"devmode": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"mcx-disabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.desktopservices": DomainPrefs{
			"DSDontWriteNetworkStores": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.DFSLoginPlugin": DomainPrefs{
			"WindowLeft": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "313",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WindowTop": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.DirectoryUtility": DomainPrefs{
			"FirstLaunch": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.DiskUtility": DomainPrefs{
			"advanced-image-options": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DUDebugMenuEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"OperationProgress DetailsVisible": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"OperationProgress ExpandedHeight": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "488",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SidebarShowAllDevices": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.dock": DomainPrefs{
			"autohide": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"autohide-delay": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"largesize": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "75",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"launchanim": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"loc": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "en_US:US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"magnification": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"minimize-to-application": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"no-bouncing": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"region": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"showAppExposeGestureEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"showhidden": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"showMissionControlGestureEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"tilesize": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "70",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"trash-full": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"wvous-br-corner": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "14",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"wvous-tr-corner": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"wvous-tr-modifier": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "1048576",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.dock.external.extra.x86_64": DomainPrefs{
			"ShowSolidarity": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.driver.AppleBluetoothMultitouch.mouse": DomainPrefs{
			"MouseButtonDivision": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "55",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseButtonMode": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "TwoButton",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseHorizontalScroll": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseMomentumScroll": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseOneFingerDoubleTapGesture": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseTwoFingerDoubleTapGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseTwoFingerHorizSwipeGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseVerticalScroll": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserPreferences": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.driver.AppleBluetoothMultitouch.trackpad": DomainPrefs{
			"Clicking": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Dragging": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DragLock": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ForceSuppressed": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"HIDScrollZoomModifierMask": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadCornerSecondaryClick": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFiveFingerPinchGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFourFingerHorizSwipeGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFourFingerPinchGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFourFingerVertSwipeGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadHandResting": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadHorizScroll": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadMomentumScroll": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadPinch": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadRightClick": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadRotate": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadScroll": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerDrag": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerHorizSwipeGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerTapGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerVertSwipeGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadTwoFingerDoubleTapGesture": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadTwoFingerFromRightEdgeSwipeGesture": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"USBMouseStopsTrackpad": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserPreferences": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.driver.AppleHIDMouse": DomainPrefs{
			"Button1": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Button2": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Button3": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Button4": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Button4Click": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Button4Force": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ButtonDominance": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ScrollH": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ScrollS": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "4",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ScrollSSize": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "30",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ScrollV": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.dt.Xcode": DomainPrefs{
			"IBAppliesAutoResizingRulesWhileResizing": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IBPreferencesMigrated": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IDEKeyBindingCurrentPreferenceSet": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Default.idekeybindings",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IDEUserWantsToEnableDeveloperSystemPolicyMode": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"XCFontAndColorCurrentTheme": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Default (Light).xccolortheme",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.FaceTime": DomainPrefs{
			"DidMigrateToSixteenByNineAspectRatio": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"kPromptEnableReadRecipts": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PhoneNumberUpgradeShown": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PreferredVideoDeviceUID": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "CC26373HVV3G1HNBN",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.facetime.bag": DomainPrefs{
			"CacheTime": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "4212",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"URL": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "https://init.ess.apple.com/WebObjects/VCInit.woa/wa/getBag?ix=3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.finder": DomainPrefs{
			"_FXShowPosixPathInTitle": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"_FXSortFoldersFirst": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"_FXSortFoldersFirstOnDesktop": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleShowAllFiles": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DisableAllAnimations": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FK_AppCentricShowSidebar": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FK_ArrangeBy": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "None",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FK_SidebarWidth": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "204",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXArrangeGroupViewBy": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Name",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXAtLeastOneScreenHasBeenInHIDPI": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXConnectToLastURL": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "smb://synologynas.local",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXDefaultSearchScope": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "SCcf",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXDesktopTouchBarUpgradedToTenTwelveOne": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXEnableExtensionChangeWarning": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudDriveDeclinedUpgrade": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudDriveDesktop": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudDriveDocuments": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudDriveEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudDriveFirstSyncDownComplete": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudLoggedIn": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXInfoWindowWidth": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "400",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXLastSearchScope": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "SCcf",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXPreferredGroupBy": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Kind",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXPreferredSearchViewStyle": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Nlsv",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXPreferredViewStyle": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "clmv",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXQuickActionsConfigUpgradeLevel": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXSidebarUpgradedSharedSearchToTwelve": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXSidebarUpgradedToTenFourteen": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXSidebarUpgradedToTenTen": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXToolbarUpgradedToTenEight": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXToolbarUpgradedToTenNine": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXToolbarUpgradedToTenSeven": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"GoToField": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "/",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NewWindowTarget": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "PfHm",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PreviewPaneGalleryWidth": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "267",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PreviewPaneInfoExpanded": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PreviewPaneWidth": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "267",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"QLEnableTextSelection": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"QuitMenuItem": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"RecentsArrangeGroupViewBy": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Date Last Opened",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SearchRecentsSavedViewStyle": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Nlsv",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowExternalHardDrivesOnDesktop": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowHardDrivesOnDesktop": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowMountedServersOnDesktop": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowPathbar": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowRemovableMediaOnDesktop": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowSidebar": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowStatusBar": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SidebarShowingiCloudDesktop": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SidebarShowingSignedIntoiCloud": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SidebarWidth": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "256",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TagsCloudSerialNumber": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "23",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.findmy": DomainPrefs{
			"coarseServiceAcknowledged_v1.0": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"onboarding_v2.0": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.FolderActionsDispatcher": DomainPrefs{
			"folderActionsEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.gamed": DomainPrefs{
			"GKLastPushTokenEnvironment": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"GKLastPushTokenPlayerID": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "U:097e62590e16c6eb3396c6958826eb60",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"GKLoginCancelled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"natType": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.HIToolbox": DomainPrefs{
			"AppleCurrentKeyboardLayoutInputSourceID": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "com.apple.keylayout.US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleDictationAutoEnable": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleFnUsageType": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.iCal": DomainPrefs{
			"AllDayAreaHeight": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "60",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AvailabilityShowTwentyFourHours": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"BirthdayEventsGenerationAttemptsToResetKey": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"BirthdayEventsGenerationLastLocale": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "en_US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CalDefaultCalendar": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "UseLastSelectedAsDefaultCalendar",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CalendarListMiniMonthVisibleMonths": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CalendarSidebarShown": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CalendarSidebarView": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CalendarSidebarWidth": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "205",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CalForceTruthFileRestoreHashKey": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "16A1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"display birthdays calendar": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"first minute of day time range": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IncludeDebugMenu": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"last calendar view description": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Monthly",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"last minute of day time range": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "1440",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastViewsTimeZone": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "America/New_York",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSDontMakeMainWindowKey": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "NO",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSEventConcurrentProcessingEnabled": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "YES",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Show Week Numbers": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TimeZone support enabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WR_DONT_ASK_FOR_DEFAULT": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.iChat": DomainPrefs{
			"BuddyPictureSetToGenericByUser": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ChatWindowControllerUnifiedChatListWidth": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "284",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ChatWindowControllerUnifiedIsVisible": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DaemonConnectionWaitTime": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "15",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DidCheckForDuplicateChats": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DidMarkGroupPhotosAsUnpurgeable": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DidMigratePersonCentricIDs": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DidRegenerateGroupID63841559": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"HasPromptediMessageFTU": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"HasPromptedSMSRelay": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"imageBrowser.disableOpenGL": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KeepMessageForDays": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KeepMessagesVersionID": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NotifyAboutMessagesFromUnknownContacts": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PlaySoundsKey": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PMPrintingExpandedStateForPrint2": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SaveConversationsOnClose": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SODefaultTranscriptName": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "ChatTranscriptViewController",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UnifiedChatListViewControllerAutomaticallySortsChats": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UnifiedChatWindowControllerSelectionGUIDSet": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "SMS;-;22395",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WebIconDatabaseEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.icloud.searchpartyuseragent": DomainPrefs{
			"lastLaunchLocale": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "en",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Image_Capture": DomainPrefs{
			"com.apple.imagekit.cameraviewmode": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_Accessory_selectedTag": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "1000",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_Camera_preferPostPocessingApp": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_Camera_selectedPathType": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_Camera_selectedTag": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "1100",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_lastUsedDeviceIsRemote": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_lastUsedDeviceName": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Mike's iPhone8",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_Scanner_downloadURL": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "~/Pictures",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_Scanner_PostProcessApplication": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "/Applications/Preview.app",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IKC_sort_ascending": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IKC_sort_key": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "creationDate",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.imagecapture": DomainPrefs{
			"loggingLevel": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3073",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ImageCaptureService": DomainPrefs{
			"IK_Accessory_selectedTag": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "1000",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.imagent": DomainPrefs{
			"IMChatVocabularyUpdaterDidPerformInitialUpdateKey": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Setting.EnableReadReceipts": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Setting.GlobalReadReceiptsVersionID": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.imdpersistence.IMDPersistenceAgent": DomainPrefs{
			"LogAllIOErrors": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.imessage": DomainPrefs{
			"PreviewTranscodingQualityOnWiFiWasInitializedToDefaultValue": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PreviewTranscodingWasInitializedToDefaultValue": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.imessage.bag": DomainPrefs{
			"CacheTime": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "4385",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"URL": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "https://init-p01md.apple.com/bag?ix=2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.internetconfigpriv": DomainPrefs{
			"WWWHomePage": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "http://livepage.apple.com/",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.internetconfigspec": DomainPrefs{
			"WWWHomePage": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "http://livepage.apple.com/",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ipTelephony": DomainPrefs{
			"ImsLoggingEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.iTunes": DomainPrefs{
			"AutomaticDeviceBackupsDisabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.iTunes.WebKit2AsynchronousSpellCheckingEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.iTunes.WebKit2DeveloperExtrasEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.iTunes.WebKit2EnableInheritURIQueryComponent": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.iTunes.WebKit2ShouldPrintBackgrounds": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.iTunes.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.iTunes.WebKit2SuppressesIncrementalRendering": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.iTunes.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"debugAssertCategoriesEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DeviceBackupsDisabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"disableRadio": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"disableShareLibraryInfo": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"doesAccountArtistListHaveSharePermission": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dontAutomaticallySyncIPods": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"isAccountAdmin": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"isAccountEnrolledInAppleMusic": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"isAccountEnrolledInITunesMatch": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"log-PlayQueue-LastSelectedTab": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"log-push": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MacBuddyStoreID": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "mikeschinkel@gmail.com",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Store Apple ID": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "mikeschinkel@gmail.com",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storefront": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "143441-1,32",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.keyboard": DomainPrefs{
			"KeyboardWordOrSentenceTrackingForPFL": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.keychainaccess": DomainPrefs{
			"Item Preview Closed": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Keychain Item List Sort Descending": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Keychain Item List Sorting": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "modified-date",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Keychains List Closed": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Last Selected Category": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "All Items",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Last Selected Keychain": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "/System/Library/Keychains/SystemRootCertificates.keychain",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"VPNSSItemsChecked": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Version 10.15.7 (Build 19H1715)",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.loginwindow": DomainPrefs{
			"MiniBuddyLaunch": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"oneTimeSSMigrationComplete": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TALLogoutSavesState": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.madrid": DomainPrefs{
			"CKMOCAccountsMatch": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitCheckedMiCSwitchEligibilityOnImagentLaunch": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitEligibleToToggleMiCSwitch": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitIsEligibleForTruthZone": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitIsRemovedFromBackup": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitIsSyncing": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitSyncingEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"createdChatZone": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableCKSyncingV2": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"hasCompletedInitialCKSync": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCKFinishedFetchingAttachmentsFromCloudKit": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitAccountStatusKey": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitStartingDisableDevices": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitStartingEnabledSettingChange": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitStartingInitialSync": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitStartingPeriodicSync": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitSyncControllerSyncRecordTypeKey": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitSyncControllerSyncStateKey": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitSyncControllerSyncTypeKey": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitSyncPaused": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"initialSyncRecordHasBeenWritten": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"RequestPriorityRamp": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.mediaaccessibility.public": DomainPrefs{
			"MACaptionDisplayType": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.menuextra.battery": DomainPrefs{
			"ShowPercent": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "YES",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.menuextra.clock": DomainPrefs{
			"DateFormat": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "EEE MMM d  j:mm a",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Messages": DomainPrefs{
			"hasPerformedNewDeviceBringUpSync": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"hasTriedToEnableCKAndSyncAfterFirstImagentConnection": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.messages.facetime": DomainPrefs{
			"FaceTimeTwoTimeCallthroughEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.messages.nicknames": DomainPrefs{
			"MeCardSharingAudience": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MeCardSharingEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MeCardSharingImageForkedFromMeCard": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NicknameAppleIDAndiCloudAccountMatchAndExist": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NicknameInfoRequestedFromPeers": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NicknameScrutinyDoNotHandle": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.messages.pinning": DomainPrefs{
			"IMPinningPinConfigMigrationKey-v2-r2": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.MobileSMS": DomainPrefs{
			"BusinessChatPrivacyPageDisplayed": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CKMigratedAutoSpamReports26375262": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KeepMessageForDays": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"LegacyAppSidebarPersistedWidth": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "284",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NotifyAboutMessagesFromUnknownContacts": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PendingCleared": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PlaySoundsKey": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"QuickSaveHasBeenUsedBefore": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ReadReceiptSettingsConfirmed": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sForceSMSSpamFilteringCompleted": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sForceUnknownFilteringCompleted": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SidebarPersistedWidth": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "320",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TextFontSize": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "13",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TextSize": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "4",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.MobileSMSPreview": DomainPrefs{
			"IMPreviewGenerationMaxPxWidth": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "300",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMPreviewGenerationMinHeight": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "50",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMPreviewGenerationMinWidth": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "50",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMPreviewGenerationScreenScale": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Music": DomainPrefs{
			"automaticallyDownloadArtwork": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"cddbPrefsOK": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.Music.WebKit2AsynchronousSpellCheckingEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.Music.WebKit2EnableInheritURIQueryComponent": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.Music.WebKit2ShouldPrintBackgrounds": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.Music.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.Music.WebKit2SuppressesIncrementalRendering": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.Music.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"debugAssertCategoriesEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"doesStoreSupportCloudMusicLibrary": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dontAskForAllPlaylistItemRemoval": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dontShowRestrictionsPrefs": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dontWarnDownloadCloudPurchases": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"encoderName": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "AAC Encoder",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"eqEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"eqWindowHPos": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "1059",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"eqWindowVPos": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "617",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"gridViewSearchDoesNotSwitchView": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"haveAskedToCheckForAvailableDownloads": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"imported-eq-presets": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ITUserPrefsMigrated": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"log-PlayQueue-LastSelectedTab": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerHPos": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "1722",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniPlayerLargeArtVisible": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniPlayerQueueVisible": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerSnapMode": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerUserSetHeight": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "826",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerVPos": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "-958",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerWidthInDesignCoords": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "534",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"musicVideoPlaybackLocation": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"playbackIsFullscreen": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"preloadFilesIntoMemory": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"searchSavedTab": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"soundEnhancerAmount": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "191",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"soundEnhancerEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeMediaTypeFlags": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3422",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsCloudPurchases": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsPasswordSettings": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsPreviousPurchases": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsUPP": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"updateLevel": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "43",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userHasActivatedFullScreenVisualizer": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userMaxConcurrentDownloads": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userWantsPlaybackNotifications": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"videoWindowHPos": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "4025",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"videoWindowHSize": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3492",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"videoWindowVPos": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "739",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"videoWindowVSize": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "1964",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"volumeWSG": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "0.1204547",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ncprefs": DomainPrefs{
			"content_visibility": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sort_order": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.networkConnect": DomainPrefs{
			"VPNShowTime": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.NetworkUtility": DomainPrefs{
			"NUNetstatChoice": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NUPingChoice": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NUPortScanAddress": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "10.10.10.51",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NUPortScanRange": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NUSelectedTabItem": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "NUTraceroute",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NUTracerouteAddress": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "104.130.122.30",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NUWhoisSelectedServer": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "whois.internic.net",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.NewDeviceOutreach": DomainPrefs{
			"CachedWarrantyLocale": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "en_US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CachedWarrantyValidityDuration": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "28800",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"InitialOutreachActivityScheduled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PostUpgradeActivityCompleted": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PostUpgradeOSVersionKey": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "12.7.6",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.newscore": DomainPrefs{
			"com.apple.news.default_event_endpoint": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "https://news-events.apple.com/analyticseventsv2/async",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.news.notification_receipt_event_endpoint": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "https://news-notification-events.apple.com/analyticseventsv2/async",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.newsd.download.hasUnfinishedWork": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"force_refresh_user_segmentation": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"news_carplay_is_enabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"news_url_resolution_subscription_status": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"notificationArticleDiversificationSimilarityExpectationEnd": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "0.1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"notificationArticleDiversificationSimilarityExpectationStart": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "0.2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"notificationArticleDiversificationUniquePublisherExpectationSlope": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "0.75",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"notificationArticleDiversificationUniquePublisherExpectationYIntercept": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "0.25",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"notificationAssetPrefetchingRequiresWatch": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"notificationEnableAssetPrefetching": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"use_parsec_results_for_widget": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.newsd": DomainPrefs{
			"FCAppConfigurationBundleShortVersionKey": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "6.0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.notificationcenterui": DomainPrefs{
			"bannerTime": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "100",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.PhotoBooth": DomainPrefs{
			"AVDesktopPlaybackControlsControllerShowsDurationInsteadOfTimeRemainingDefaultsKey": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PBLaunchedAtLeastOnceOnLion": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserCameraUniqueIDPref": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "CC26373HVV3G1HNBN",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.photos.shareddefaults": DomainPrefs{
			"clearPurgeableResources": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CPLDefaultDownload": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Default",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DiskSpaceWasLow": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"downloadAndKeepOriginals": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"repushVideoAssetsMetadata": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Preferences": DomainPrefs{
			"UserDictionaryImportedSinceMaintenance": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.extensions.CredentialProvider": DomainPrefs{
			"userHasOrdered": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.extensions.FinderSync": DomainPrefs{
			"userHasOrdered": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.extensions.PhotoEditing": DomainPrefs{
			"userHasOrdered": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.extensions.PhotoProjects": DomainPrefs{
			"userHasOrdered": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.extensions.QuickLook": DomainPrefs{
			"userHasOrdered": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.extensions.ServicesWithUI": DomainPrefs{
			"userHasOrdered": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.extensions.SharedLinks": DomainPrefs{
			"userHasOrdered": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.extensions.ShareMenu": DomainPrefs{
			"userHasOrdered": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.preferences.extensions.XcodeSourceEditor": DomainPrefs{
			"userHasOrdered": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.print.custompresets.forprinter.Canon_MF5900_Series": DomainPrefs{
			"com.apple.print.lastPresetPref": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Default Settings",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.print.lastPresetPrefType": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.print.PrinterProxy": DomainPrefs{
			"IK_Scanner_downloadURL": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "~/Pictures",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_Scanner_PostProcessApplication": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "/System/Applications/Preview.app",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IK_Scanner_selectedTag": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "1000",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ProblemReporter": DomainPrefs{
			"ShowComments": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowDetails": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.PTPCamera": DomainPrefs{
			"FilenamePrefix": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "IMG_",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.PubSubAgent": DomainPrefs{
			"TigerMigrationCompleted": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.rapport": DomainPrefs{
			"familySyncedName": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "MacBookPro",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.remindd": DomainPrefs{
			"CloudConfigurationPath": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "/System/Library/PrivateFrameworks/ReminderKit.framework/Versions/A/Resources/CloudConfigurations/Normal.plist",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitAccountStatus": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"isDatabaseMigrated": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ThrottlingPolicyCurrentLevelIndex": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"useExtraneousAlarmBackOffThrottleInterval": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Safari": DomainPrefs{
			"cloudBookmarksMigrationEligibilityDataInvalidated": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.Safari.ContentPageGroupIdentifier.WebKit2DeveloperExtrasEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IncludeDebugMenu": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IncludeDevelopMenu": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IncludeInternalDebugMenu": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ResetCloudHistory": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowFullURLInSmartSearchField": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SuppressSearchSuggestions": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UniversalSearchEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WebKitDeveloperExtrasEnabledPreferenceKey": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WebKitHistoryAgeInDaysLimit": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "360",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WebKitHistoryItemLimit": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "9999",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WebKitInitialTimedLayoutDelay": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "0.25",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Safari.SafeBrowsing": DomainPrefs{
			"HasMigratedSafeBrowsingEnabledDefaults": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Safari.SandboxBroker": DomainPrefs{
			"DidMigrateDownloadFolderToSandbox": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DidMigrateDownloadMetadataToSandboxGroupContainer": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DidMigrateResourcesToSandbox": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowDevelopMenu": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.SafariCloudHistoryPushAgent": DomainPrefs{
			"AcknowledgedPushNotifications": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UnacknowledgedPushNotifications": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.SafariServices": DomainPrefs{
			"SearchProviderIdentifierMigratedToSystemPreference": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.screencapture": DomainPrefs{
			"captureDelay": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"last-selection-display": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"location": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "/Users/mikeschinkel/Desktop/Screenshots",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"style": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "display",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"video": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.screencaptureui": DomainPrefs{
			"com.apple.AnnotationKit.arrowHeadStyle": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AnnotationKit.brushStyle": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AnnotationKit.hasShadow": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AnnotationKit.highlightStyle": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "765200",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AnnotationKit.strokeIsDashed": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.AnnotationKit.strokeWidth": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.screensaver": DomainPrefs{
			"tokenRemovalAction": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.scriptmenu": DomainPrefs{
			"ScriptMenuEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.sharingd": DomainPrefs{
			"AutoUnlockPresentedBluetoothError": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AutoUnlockPresentedWiFiError": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AutoUnlockWatchCurrentlyInList": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AutoUnlockWatchExistedInUnlockList": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DiscoverableMode": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Contacts Only",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"HashManager-StoredDatabaseVersionKey": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"OneTimeAirDropReset": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"OneTimeAirDropReset2": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Siri": DomainPrefs{
			"StatusMenuVisible": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"VoiceTriggerUserEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.siri.shortcuts": DomainPrefs{
			"LegacyShortcutsZoneSubscriptionUnsubscribed": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WFDidUnconflictShortcuts": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WFServicesShortcutsMigrated": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.siri.VoiceShortcuts": DomainPrefs{
			"VCLSDataSequenceKey": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "62268",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.sms": DomainPrefs{
			"hasBeenApprovedForSMSRelay": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.SocialLayer": DomainPrefs{
			"SharedWithYouEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.sociallayerd.CloudKit.ckwriter": DomainPrefs{
			"accountAvailable": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"accountLastKnownUserRecordID": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "_4d72b079e6dd351d59dcdf0b9afd41d6",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"writerDone": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.SoftwareUpdate": DomainPrefs{
			"AutoUpdateMajorOSVersion": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "10.15",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.speech.recognition.AppleSpeechRecognition.prefs": DomainPrefs{
			"DictationIMCommandsWindowIsOpen": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DictationIMLocaleIdentifier": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "en_US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DictationIMMessageTracesSinceLastReport": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DictationIMPlaySoundUponRecognition": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DictationIMPresentedOfflineUpgradeSuggestion": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DictationIMUpgradedTo10_15": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DictationIMUpgradedTo10_16": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DictationIMUseOnlyOfflineDictation": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.SpeechRecognitionCore": DomainPrefs{
			"AllowAudioDucking": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Spotlight": DomainPrefs{
			"__NSEnableTSMDocumentWindowLevel": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"engagementCount-com.apple.mail": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"engagementCount-com.apple.Spotlight": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"engagementCount-com.apple.Spotlight.suggestions": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ModelName": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Canned",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"reloadShortcuts": DomainPref{
				Verified: Verified{},
				Type:     "unknown",
				Default:  "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"showedFTE": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userHasMovedWindow": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"windowHeight": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "430",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.stockholm": DomainPrefs{
			"PKWalletShouldAutomaticallyRegisterKey": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.suggestions": DomainPrefs{
			"findToShowMigrationPerformed": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"spToLearnMigrationPerformed": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SuggestionsAllowGeocode": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.systempreferences": DomainPrefs{
			"com.apple.preferences.accounts.outline.usersparent": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.preferences.sharing.selectedservice": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.SecurityPref.Privacy.LastSourceSelected": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DisableAutoLoginButtonIsHidden": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSFullScreenMenuItemEverywhere": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSQuitAlwaysKeepsWindows": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SecurityPrefTab": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "General",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowAllMode": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SoundTab": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "input",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.TelephonyUtilities": DomainPrefs{
			"FaceTimePhotosEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"relayCallingDisabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.Terminal": DomainPrefs{
			"Default Window Settings": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Basic",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"HasMigratedDefaults": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SecureKeyboardEntry": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Startup Window Settings": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Basic",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TTAppPreferences Selected Tab": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.TextEdit": DomainPrefs{
			"PlainTextEncoding": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "4",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PlainTextEncodingForWrite": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "4",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"RichText": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.textInput.keyboardServices.textReplacement": DomainPrefs{
			"_KSTRCloudKitMigratable": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"gmMigrationPercent": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"gmMigrationPercent1": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "999",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"internalMigrationPercent": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"internalMigrationPercent1": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "999",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"iOSMajorVersForCloudKitSync": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "11",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"iOSMinorVersForCloudKitSync": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"kDidMigrateToUUIDRecordName": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSCKDidSetupRecordZoneProd-TextReplacements": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSCKSubscriptionProd-TextReplacements": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSCloudKitMigrationDidComplete": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSDidMigrateToCloudKitOnCloud": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSDidPullLegacyEntriesFromPeers": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSDidPushAllLocalRecordsOnce-2": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSDidPushMigrationStatusOnce-2": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSLastKnownUserID-TextReplacements": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "_694927be690dcbcc4d7ebaccd1c628c5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSSampleShortcutWasImported_CK": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"osxMajorVersForCloudKitSync": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "10",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"osxMinorSubversionForCloudKitSync": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"osxMinorVersForCloudKitSync": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "13",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"seedMigrationPercent": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"seedMigrationPercent1": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "999",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TRCKSyncCountHalflifeInHours": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "10",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TRCKSyncMaxCountThreshold": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "200",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userDidFallInMigrationAllowBatch": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.TextInputMenu": DomainPrefs{
			"visible": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.tipsd": DomainPrefs{
			"TPSTipsAppInstalled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.TTY": DomainPrefs{
			"RTTContinuityRTTIsSupportedPreference": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"RTTSettingsVersionPreference": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TUIsRelayCallingEnabledPreference": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TUSupportsRelayCallingPreference": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.TV": DomainPrefs{
			"automaticallyDeleteVideoAssetsAfterWatching": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"automaticallyDownloadArtwork": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"cddbPrefsOK": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"checkForAvailableDownloads": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.TV.WebKit2AsynchronousSpellCheckingEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.TV.WebKit2EnableInheritURIQueryComponent": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.TV.WebKit2ShouldPrintBackgrounds": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.TV.WebKit2ShowsToolTipOverTruncatedText": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.TV.WebKit2SuppressesIncrementalRendering": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.TV.WebKit2UserInterfaceDirectionPolicy": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"debugAssertCategoriesEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dontAskForAllPlaylistItemRemoval": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"gridViewSearchDoesNotSwitchView": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"haveAskedToCheckForAvailableDownloads": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ITUserPrefsMigrated": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerHPos": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2723",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniPlayerLargeArtVisible": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniPlayerQueueVisible": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerSnapMode": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerUserSetHeight": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "771",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerVPos": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "821",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"miniplayerWidthInDesignCoords": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "288",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"musicVideoPlaybackLocation": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"playbackIsFullscreen": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"preloadFilesIntoMemory": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeMediaTypeFlags": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3422",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsCloudPurchases": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsPasswordSettings": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsPreviousPurchases": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsUPP": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"tabViewMode": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "22",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"updateLevel": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "43",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userHasActivatedFullScreenVisualizer": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userMaxConcurrentDownloads": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"userWantsPlaybackNotifications": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"videoWindowHPos": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "461",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"videoWindowHSize": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "1906",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"videoWindowVPos": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "139",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"videoWindowVSize": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "1072",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.UIKit": DomainPrefs{
			"hasAccessibilityBeenMigrated": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.universalaccess": DomainPrefs{
			"AssistiveControlType": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"closeViewDesiredZoomFactor": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "1.1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"closeViewHotkeysEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"closeViewScrollWheelToggle": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"closeViewSplitScreenRatio": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "0.2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"closeViewZoomDisplayID": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"closeViewZoomedIn": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"closeViewZoomFactor": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"closeViewZoomFactorBeforeTermination": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"contrast": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"differentiateWithoutColor": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dwellEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"dwellHideUITimeout": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "15",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"grayscale": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"headMouseEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"hoverTextEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"hoverTextIsAlwaysOn": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"hoverTextIsHoveringAndVisible": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"increaseContrast": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"keyboardAccessFocusRingTimeout": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "15",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastNightShiftActive": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastNightShiftEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"lastNightShiftMode": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"login": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"mouseDriver": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"mouseDriverCursorSize": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "1.44109",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"reduceTransparency": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"selectedTab": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "14",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sessionChange": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"slowKey": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"slowKeyDelay": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "250",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"stickyKey": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"switchAutoScanElementInterval": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "0.5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"switchAutoScanPanelInterval": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "0.5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"switchCoalescePressesDuration": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"switchFirstElementDelay": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"switchHideUITimeout": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "15",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"switchHoldBeforeRepeatDuration": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"switchMinimumPressDuration": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"switchOnOffKey": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"switchSweepingCursorSpeed": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"useStickyKeysShortcutKeys": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"virtualKeyboardOnOff": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"voiceOverOnOffKey": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"whiteOnBlack": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.UserAccountUpdater": DomainPrefs{
			"com.apple.HidLibraryFolderAlready": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.visualintelligence": DomainPrefs{
			"bypassAvailability": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"debugUI": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableAlbumsDomain": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableArtDomain": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableBooksDomain": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableCoarseClassification": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableLandmarkDomain": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableNatureDomain": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enablePetsDomain": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enablePhotosApp": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableQuickLook": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableSafariApp": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableScreenshots": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"firstTimeExperience": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"initialized": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sendLocationInfo": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sendOCRText": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.VoiceOver4.local": DomainPrefs{
			"SCRCUserDefaultsAllowAirPlay": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SCRCUserDefaultsPlannedShutdownSuccessful": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.VoiceOverTraining": DomainPrefs{
			"doNotShowSplashScreen": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.voicetrigger": DomainPrefs{
			"VoiceTrigger Enabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"GlobalPreferences": DomainPrefs{
			"646F6E7A_00000000_00000001_6E7A6361_696D6963": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AKLastIDMSEnvironment": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleActionOnDoubleClick": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Maximize",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleAntiAliasingThreshold": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "4",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleAquaColorVariant": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleEnableSwipeNavigateWithScrolls": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleKeyboardUIMode": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleLanguagesDidMigrate": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "12.7.6",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleLocale": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "en_US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleMeasurementUnits": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Inches",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleMetricUnits": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleMiniaturizeOnDoubleClick": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleShowAllExtensions": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleShowScrollBars": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Always",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleTemperatureUnit": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Fahrenheit",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleWindowTabbingMode": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "always",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AudioQuest inc. AudioQuest DragonFly": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Blue Microphones Yeti Stereo Microphone": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Canon MF5900 Series": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.mouse.scaling": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.sound.beep.flash": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.springing.delay": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "0.5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.springing.enabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.swipescrolldirection": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.trackpad.forceClick": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.trackpad.scaling": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ContextMenuGesture": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Fujitsu ScanSnap iX500": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"InitialKeyRepeat": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "15",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KeyRepeat": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NavPanelFileListModeForOpenMode": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NavPanelFileListModeForSaveMode": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSAutomaticCapitalizationEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSAutomaticDashSubstitutionEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSAutomaticPeriodSubstitutionEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSAutomaticQuoteSubstitutionEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSAutomaticSpellingCorrectionEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSAutomaticTextCompletionEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSAutomaticWindowAnimationsEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSLinguisticDataAssetsRequestLastInterval": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "86400",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"QLPanelAnimationDuration": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WebAutomaticSpellingCorrectionEnabled": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"WebKitDeveloperExtras": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Yubico YubiKey OTP+FIDO+CCID": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
	}
}
