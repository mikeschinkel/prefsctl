package prefdefaults

//goland:noinspection SpellCheckingInspection
func montereyPrefDefaults() DomainDefaults {
	return DomainDefaults{
		"com.apple.dock": comAppleDock(),

		"com.apple.Accessibility": DomainPrefs{
			"ApplicationAccessibilityEnabled": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AutomationEnabled": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AXSClassicInvertColorsPreference": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"BrailleInputDeviceConnected": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CommandAndControlEnabled": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DarkenSystemColors": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DifferentiateWithoutColor": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"EnhancedBackgroundContrastEnabled": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FullKeyboardAccessEnabled": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FullKeyboardAccessFocusRingEnabled": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"GenericAccessibilityClientEnabled": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"InvertColorsEnabled": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KeyRepeatDelay": DomainPref{
				Descr:   "",
				Type:    "float",
				Default: "0.5",
				Labels: NewLabels(
					DefaultsSet,
					TypeVerified,
					UserManaged,
				),
			},
			"KeyRepeatEnabled": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ReduceMotionEnabled": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SpeakThisEnabled": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"VoiceOverTouchEnabled": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SelectedTab": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowCategory": DomainPref{
				Descr:   "",
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
				Descr:   "",
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
				Descr:   "",
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"home-sharing-enabled": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"home-sharing-group-id": DomainPref{
				Descr:   "",
				Type:    "unknown",
				Default: "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"home-sharing-sequence-id": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "100",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"photo-sharing-enabled": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"public-sharing-enabled": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"public-sharing-is-protected": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"shared-library-name": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"debugAssertCategoriesEnabled": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"imported-media-domains": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "15",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"podcasts-migrated": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeMediaTypeFlags": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "3422",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsPasswordSettings": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsPreviousPurchases": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"storeSupportsUPP": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "int",
				Default: "55",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseButtonMode": DomainPref{
				Descr:   "",
				Type:    "string",
				Default: "OneButton",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseHorizontalScroll": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseMomentumScroll": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseOneFingerDoubleTapGesture": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseTwoFingerDoubleTapGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseTwoFingerHorizSwipeGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseVerticalScroll": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserPreferences": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Clicking": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Dragging": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DragLock": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FirstClickThreshold": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ForceSuppressed": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SecondClickThreshold": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadCornerSecondaryClick": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFiveFingerPinchGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFourFingerHorizSwipeGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFourFingerPinchGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFourFingerVertSwipeGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadHandResting": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadHorizScroll": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadMomentumScroll": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadPinch": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadRightClick": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadRotate": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadScroll": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerDrag": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerHorizSwipeGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerTapGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerVertSwipeGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadTwoFingerDoubleTapGesture": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadTwoFingerFromRightEdgeSwipeGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"USBMouseStopsTrackpad": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserPreferences": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Voice Trigger Needs Sync": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomLockDevice-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenApp": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenApp-00000000-0000-0000-0000-000000000000:0": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenApp-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenApp-ask": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenURL": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenURL-00000000-0000-0000-0000-000000000000:0": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenURL-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomOpenURL-ask": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomScreenObservation-00000000-0000-0000-0000-000000000000:0": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"allowClassroomScreenObservation-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"forceUnpromptedRemoteScreenObservation-00000000-0000-0000-0000-000000000000:0": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"forceUnpromptedRemoteScreenObservation-00000000-0000-0000-0000-000000000000:0-ask": DomainPref{
				Descr:   "",
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
				Descr:   "",
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
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DisplayUseInvertedPolarity": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.driver.AppleBluetoothMultitouch.mouse": DomainPrefs{
			"MouseButtonDivision": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "55",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseButtonMode": DomainPref{
				Descr:   "",
				Type:    "string",
				Default: "OneButton",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseHorizontalScroll": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseMomentumScroll": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseOneFingerDoubleTapGesture": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseTwoFingerDoubleTapGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseTwoFingerHorizSwipeGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"MouseVerticalScroll": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserPreferences": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "bool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Dragging": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DragLock": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadCornerSecondaryClick": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFiveFingerPinchGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFourFingerHorizSwipeGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFourFingerPinchGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadFourFingerVertSwipeGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadHandResting": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadHorizScroll": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadMomentumScroll": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadPinch": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadRightClick": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadRotate": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadScroll": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerDrag": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerHorizSwipeGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerTapGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadThreeFingerVertSwipeGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadTwoFingerDoubleTapGesture": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TrackpadTwoFingerFromRightEdgeSwipeGesture": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"USBMouseStopsTrackpad": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"UserPreferences": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Button2": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Button3": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Button4": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Button4Click": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Button4Force": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ButtonDominance": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ScrollH": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ScrollS": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "4",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ScrollSSize": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "30",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ScrollV": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "int",
				Default: "3600",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"URL": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "string",
				Default: "Name",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXDesktopTouchBarUpgradedToTenTwelveOne": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudDriveDeclinedUpgrade": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudDriveDesktop": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudDriveDocuments": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudDriveEnabled": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudDriveFirstSyncDownComplete": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXICloudLoggedIn": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXPreferredViewStyle": DomainPref{
				Descr:   "",
				Type:    "string",
				Default: "Nlsv",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXQuickActionsConfigUpgradeLevel": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXSidebarUpgradedToTenFourteen": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXSidebarUpgradedToTenTen": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXToolbarUpgradedToTenEight": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXToolbarUpgradedToTenNine": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"FXToolbarUpgradedToTenSeven": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"RecentsArrangeGroupViewBy": DomainPref{
				Descr:   "",
				Type:    "string",
				Default: "Date Last Opened",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowExternalHardDrivesOnDesktop": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowHardDrivesOnDesktop": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowRemovableMediaOnDesktop": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowSidebar": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SidebarShowingiCloudDesktop": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SidebarShowingSignedIntoiCloud": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SidebarWidth": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "217",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TagsCloudSerialNumber": DomainPref{
				Descr:   "",
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
				Descr:   "",
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
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"BirthdayEventsGenerationLastLocale": DomainPref{
				Descr:   "",
				Type:    "string",
				Default: "en_US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CalForceTruthFileRestoreHashKey": DomainPref{
				Descr:   "",
				Type:    "string",
				Default: "16A1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"display birthdays calendar": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "int",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DidMarkGroupPhotosAsUnpurgeable": DomainPref{
				Descr:   "",
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
				Descr:   "",
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
				Descr:   "",
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
				Descr:   "",
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
				Descr:   "",
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
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PreviewTranscodingWasInitializedToDefaultValue": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "int",
				Default: "3600",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"URL": DomainPref{
				Descr:   "",
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
				Descr:   "",
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
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"oneTimeSSMigrationComplete": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TALLogoutReason": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitIsEligibleForTruthZone": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitIsRemovedFromBackup": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitSyncingEnabled": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitAccountStatusKey": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitStartingDisableDevices": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitStartingEnabledSettingChange": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitStartingInitialSync": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitStartingPeriodicSync": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitSyncControllerSyncRecordTypeKey": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitSyncControllerSyncStateKey": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitSyncControllerSyncTypeKey": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"IMCloudKitSyncPaused": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"initialSyncRecordHasBeenWritten": DomainPref{
				Descr:   "",
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
				Descr:   "",
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
				Descr:   "",
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
				Descr:   "",
				Type:    "int",
				Default: "2",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sort_order": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"PostUpgradeOSVersionKey": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"use_parsec_results_for_widget": DomainPref{
				Descr:   "",
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
				Descr:   "",
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
				Descr:   "",
				Type:    "string",
				Default: "/System/Library/PrivateFrameworks/ReminderKit.framework/Versions/A/Resources/CloudConfigurations/Normal.plist",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"CloudKitAccountStatus": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"isDatabaseMigrated": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"useExtraneousAlarmBackOffThrottleInterval": DomainPref{
				Descr:   "",
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
				Descr:   "",
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
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DidMigrateDownloadMetadataToSandboxGroupContainer": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DidMigrateResourcesToSandbox": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowDevelopMenu": DomainPref{
				Descr:   "",
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
				Descr:   "",
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
				Descr:   "",
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
				Descr:   "",
				Type:    "int",
				Default: "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"OneTimeAirDropReset2": DomainPref{
				Descr:   "",
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
				Descr:   "",
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
				Descr:   "",
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
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"engagementCount-com.apple.mail": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"engagementCount-com.apple.Spotlight": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"engagementCount-com.apple.Spotlight.suggestions": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"showedFTE": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"spToLearnMigrationPerformed": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SuggestionsAllowGeocode": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "intBool",
				Default: "8",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"DisableAutoLoginButtonIsHidden": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSFullScreenMenuItemEverywhere": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSQuitAlwaysKeepsWindows": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SecurityPrefTab": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "string",
				Default: "Basic",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"HasMigratedDefaults": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"SecureKeyboardEntry": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"Startup Window Settings": DomainPref{
				Descr:   "",
				Type:    "string",
				Default: "Basic",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TTAppPreferences Selected Tab": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KSSampleShortcutWasImported_CK": DomainPref{
				Descr:   "",
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
				Descr:   "",
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
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"TPSWaitingToShowWelcomeNotification": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"closeViewZoomDisplayID": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"closeViewZoomFactor": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"login": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sessionChange": DomainPref{
				Descr:   "",
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
				Descr:   "",
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
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"debugUI": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableAlbumsDomain": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableArtDomain": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableBooksDomain": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableCoarseClassification": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableLandmarkDomain": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableNatureDomain": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enablePetsDomain": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enablePhotosApp": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableQuickLook": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableSafariApp": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"enableScreenshots": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"firstTimeExperience": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"initialized": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sendLocationInfo": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"sendOCRText": DomainPref{
				Descr:   "",
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
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleAntiAliasingThreshold": DomainPref{
				Descr:   "",
				Type:    "int",
				Default: "4",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleLanguagesDidMigrate": DomainPref{
				Descr:   "",
				Type:    "string",
				Default: "12.7.6",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleLocale": DomainPref{
				Descr:   "",
				Type:    "string",
				Default: "en_US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleMiniaturizeOnDoubleClick": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.sound.beep.flash": DomainPref{
				Descr:   "",
				Type:    "intBool",
				Default: "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.springing.delay": DomainPref{
				Descr:   "",
				Type:    "float",
				Default: "0.5",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.springing.enabled": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.trackpad.forceClick": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSAutomaticCapitalizationEnabled": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSAutomaticPeriodSubstitutionEnabled": DomainPref{
				Descr:   "",
				Type:    "bool",
				Default: "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"NSLinguisticDataAssetsRequestLastInterval": DomainPref{
				Descr:   "",
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
