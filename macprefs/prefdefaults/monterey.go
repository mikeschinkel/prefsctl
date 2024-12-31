package prefdefaults

//goland:noinspection SpellCheckingInspection
func montereyPrefDefaults() DomainDefaults {
	return DomainDefaults{
		"com.apple.Accessibility": DomainPrefs{
			"ApplicationAccessibilityEnabled": DomainPref{
				Type:    "intBool",
				Default: "0",
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
				Default: "0.5",
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
		},
		"com.apple.ActivityMonitor": DomainPrefs{
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
				Default: "102",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.AddressBook": DomainPrefs{
			"ABBirthDayVisible": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.amp.mediasharingd": DomainPrefs{
			"home-sharing-computer-id": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"home-sharing-enabled": DomainPref{
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"home-sharing-group-id": DomainPref{
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"home-sharing-sequence-id": DomainPref{
				Type:    "int",
				Default: "100",
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
			"shared-library-name": DomainPref{
				Type:    "string",
				Default: "Mike Schinkel’s Library",
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
			"imported-media-domains": DomainPref{
				Type:    "int",
				Default: "15",
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
			"storeMediaTypeFlags": DomainPref{
				Type:    "int",
				Default: "3422",
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
				Default: "OneButton",
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
			"Clicking": DomainPref{
				Type:    "bool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Dragging": DomainPref{
				Type:    "bool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DragLock": DomainPref{
				Type:    "bool",
				Default: "0",
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
				Default: "false",
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
				Default: "0",
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
				Default: "true",
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
		"com.apple.assistant": DomainPrefs{
			"MultiUser Shared Data Needs Sync": DomainPref{
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
				Default: "34",
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
		"com.apple.dock": DomainPrefs{
			"loc": DomainPref{
				Type:    "string",
				Default: "en_US:US",
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
			"wvous-br-corner": DomainPref{
				Type:    "int",
				Default: "14",
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
				Default: "OneButton",
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
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Dragging": DomainPref{
				Type:    "bool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DragLock": DomainPref{
				Type:    "bool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadCornerSecondaryClick": DomainPref{
				Type:    "int",
				Default: "0",
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
				Default: "true",
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
		"com.apple.facetime.bag": DomainPrefs{
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
			"FXArrangeGroupViewBy": DomainPref{
				Type:    "string",
				Default: "Name",
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
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudDriveFirstSyncDownComplete": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudLoggedIn": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXPreferredViewStyle": DomainPref{
				Type:    "string",
				Default: "Nlsv",
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
			"RecentsArrangeGroupViewBy": DomainPref{
				Type:    "string",
				Default: "Date Last Opened",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowExternalHardDrivesOnDesktop": DomainPref{
				Type:    "bool",
				Default: "true",
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
			"ShowRemovableMediaOnDesktop": DomainPref{
				Type:    "bool",
				Default: "true",
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
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SidebarWidth": DomainPref{
				Type:    "int",
				Default: "217",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TagsCloudSerialNumber": DomainPref{
				Type:    "int",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.FolderActionsDispatcher": DomainPrefs{
			"folderActionsEnabled": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.iCal": DomainPrefs{
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
			"CalForceTruthFileRestoreHashKey": DomainPref{
				Type:    "string",
				Default: "16A1",
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
		},
		"com.apple.iChat": DomainPrefs{
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
		},
		"com.apple.icloud.searchpartyuseragent": DomainPrefs{
			"lastLaunchLocale": DomainPref{
				Type:    "string",
				Default: "en",
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
		"com.apple.imagent": DomainPrefs{
			"IMChatVocabularyUpdaterDidPerformInitialUpdateKey": DomainPref{
				Type:    "bool",
				Default: "true",
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
				Default: "https://init-p01md.apple.com/bag?ix=2",
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
			"TALLogoutReason": DomainPref{
				Type:    "unknown",
				Default: "Shut Down",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.madrid": DomainPrefs{
			"CloudKitCheckedMiCSwitchEligibilityOnImagentLaunch": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitIsEligibleForTruthZone": DomainPref{
				Type:    "bool",
				Default: "false",
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
			"CloudKitSyncingEnabled": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitAccountStatusKey": DomainPref{
				Type:    "int",
				Default: "1",
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
		"com.apple.messages.nicknames": DomainPrefs{
			"NicknameScrutinyDoNotHandle": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.ncprefs": DomainPrefs{
			"content_visibility": DomainPref{
				Type:    "int",
				Default: "2",
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
		"com.apple.NewDeviceOutreach": DomainPrefs{
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
			"force_refresh_user_segmentation": DomainPref{
				Type:    "bool",
				Default: "false",
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
		"com.apple.ProblemReporter": DomainPrefs{
			"ShowComments": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.remindd": DomainPrefs{
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
				Default: "3",
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
			"useExtraneousAlarmBackOffThrottleInterval": DomainPref{
				Type:    "bool",
				Default: "true",
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
			"ShowDevelopMenu": DomainPref{
				Type:    "bool",
				Default: "false",
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
		"com.apple.sharingd": DomainPrefs{
			"HashManager-StoredDatabaseVersionKey": DomainPref{
				Type:    "int",
				Default: "3",
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
		"com.apple.siri.shortcuts": DomainPrefs{
			"WFDidUnconflictShortcuts": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.siri.VoiceShortcuts": DomainPrefs{
			"VCLSDataSequenceKey": DomainPref{
				Type:    "int",
				Default: "3913",
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
			"showedFTE": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.suggestions": DomainPrefs{
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
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.systempreferences": DomainPrefs{
			"com.apple.SecurityPref.Privacy.LastSourceSelected": DomainPref{
				Type:    "intBool",
				Default: "8",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DisableAutoLoginButtonIsHidden": DomainPref{
				Type:    "bool",
				Default: "false",
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
				Default: "Privacy",
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
		},
		"com.apple.textInput.keyboardServices.textReplacement": DomainPrefs{
			"KSDidPushMigrationStatusOnce-2": DomainPref{
				Type:    "bool",
				Default: "true",
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
		},
		"com.apple.TextInputMenu": DomainPrefs{
			"visible": DomainPref{
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.tipsd": DomainPrefs{
			"TPSTipsAppInstalled": DomainPref{
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TPSWaitingToShowWelcomeNotification": DomainPref{
				Type:    "unknown",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.universalaccess": DomainPrefs{
			"closeViewHotkeysEnabled": DomainPref{
				Type:    "bool",
				Default: "false",
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
			"closeViewZoomFactor": DomainPref{
				Type:    "intBool",
				Default: "1",
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
			"sessionChange": DomainPref{
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
		"GlobalPreferences": DomainPrefs{
			"AKLastIDMSEnvironment": DomainPref{
				Type:    "intBool",
				Default: "0",
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
			"AppleMiniaturizeOnDoubleClick": DomainPref{
				Type:    "bool",
				Default: "false",
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
			"com.apple.trackpad.forceClick": DomainPref{
				Type:    "bool",
				Default: "true",
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
			"NSAutomaticPeriodSubstitutionEnabled": DomainPref{
				Type:    "bool",
				Default: "true",
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
		},
	}
}
