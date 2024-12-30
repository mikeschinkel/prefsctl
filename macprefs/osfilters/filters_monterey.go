package osfilters

import (
	"github.com/mikeschinkel/prefsctl/kvfilters"
)

//goland:noinspection SpellCheckingInspection
func MontereyQueryFilters() []kvfilters.Filter {
	return []kvfilters.Filter{
		//// DEBUGGING
		//kvfilters.OmitWhenGroupIsNotOneOf{
		//	"GlobalPreferences",
		//},
		kvfilters.OmitWhenGroupPrefixIsNotOneOf{
			"com.apple.",
		},
		//// TEMP for debugging
		//OmitWhenGroupIsNotOneOf{
		//	"GlobalPreferences",
		//},
		// Whitelisted user configuration props
		kvfilters.KeepWhenGroupIsOneOf{
			"GlobalPreferences",
		},
		kvfilters.KeepWhenKeyIsOneOf{
			"KeyRepeatInterval",
			"AutoUpdateMajorOSVersion",
			"TALLogoutSavesState",
			"DictationIMLocaleIdentifier",
			"PreferredVideoDeviceUID",
			"enableCKSyncingV2",
		},
		// Non-configurable domain prefixes
		kvfilters.OmitWhenGroupIsOneOf{
			"com.apple.security.",
			"com.apple.icloud.",
			"com.apple.cloud",
			"com.apple.ams",
			"com.apple.AdLib",
			"com.apple.Ad",
			"com.apple.preferences.extensions.",
			"com.apple.internal.",
		},
		// Non-configurable domain suffixes
		kvfilters.OmitWhenGroupSuffixIsOneOfIgnoringCase{
			".notbackedup",
			".daemon",
			".agent",
			".syncagent",
		},
		// Non-configurable preference keys
		kvfilters.OmitWhenKeySuffixIsOneOfIgnoringCase{
			"count",
			".lastselectedtab",
			".lastselection",
			"version",
			"state",
			"date",
			"cache",
			"identifier",
			"uid",
			"v2",
		},
		// Non-configurable preference keys
		kvfilters.OmitWhenKeyPrefixIsOneOf{
			"OSAStandardAdditions",
			"__uniquePageGroupID",
			"messageTracer",
		},
		// Non-configurable domains that contain
		kvfilters.OmitWhenGroupContainsOneOf{
			"cloudkit",
		},
		// Non-configurable prefs that contain
		kvfilters.OmitWhenKeyContainsOneOfIgnoringCase{
			".last",
		},
		// Only keep apple-specific prefs, at least currently
		kvfilters.OmitWhenGroupIsOneOf{
			// System Management & Security Domains
			"com.apple.security.cloudkeychainproxy3.keysToRegister",
			"com.apple.security.sosaccount",
			"com.apple.security.KCN",
			"com.apple.cloudpaird",
			"com.apple.cloudkeychain",
			"com.apple.security.pboxd",
			"com.apple.protectedcloudstorage.protectedcloudkeysyncing",
			// System Services & Agents
			"com.apple.bird",
			"com.apple.parsecd",
			"com.apple.xpc.activity2",
			"com.apple.identityservicesd",
			"com.apple.coreservices.useractivityd",
			"com.apple.diagnosticextensionsd",
			"com.apple.commcenter.callservices",
			"com.apple.amsengagementd",
			"com.apple.studentd",
			"com.apple.systemuiserver",
			"com.apple.IMCoreSpotlight",
			"com.apple.imdsmsrecordstore",
			"com.apple.messaging.watchdog",
			"com.apple.CallHistorySyncHelper",
			// Synchronization & CloudKit
			"com.apple.syncserver",
			"com.apple.cloudd",
			// Analytics & Metrics
			"com.apple.mmcs",
			"com.apple.appstored",
			"com.apple.AppleMediaServices",
			// Migration & State Management
			"com.apple.SetupAssistant",
			"com.apple.MobileSMS.CKDNDList",
			"com.apple.universalaccessAuthWarning",
			"com.apple.preferences.extensions.",
			// Device & Account Management
			"com.apple.iPod",
			"com.apple.accountsd",
			"com.apple.icloud.fmfd",
			// Store & Commerce
			"com.apple.itunesstored",
			"com.apple.commerce",
			"com.apple.AppStoreComponents",
			"com.apple.AdLib",
			"com.apple.AdPlatforms",
			// Debug & Development
			"com.apple.fileproviderd",
			"com.apple.internal.ck",
			"com.apple.networkresolutiond",
			// Temporary State Management
			"com.apple.stocks.account",
			"com.apple.commcenter.data",
			// Temporary State Management
			"com.apple.stocks.account",
			"com.apple.commcenter.data",
			// Version Control & Registration
			"com.apple.FontRegistry.user",
			"com.apple.GEO",
		},
		kvfilters.OmitWhenKeyPrefixIsOneOf{
			// System Paths
			"/System/",
			// Whatever this is
			"seed-numNotifications-",
			"seed-notificationDueDate",
			"seed-viewed",
		},
		// Agent/Daemon startup times
		kvfilters.OmitWhenKeyIsOneOf{
			"CKStartupTime",
		},
		kvfilters.OmitWhenKeyPrefixIsOneOf{
			// Objective-C Symbols
			"NSApplicationCrashOnExceptions",
			"NSWindowTabbingShoudShowTabBarKey",
			"NSSplitView",
			"NSStatusItem",
			"NSWindow",
			"NSNavPanel",
			"NSNavLastRoot",
			"NSTableView",
			"NSToolbar",
		},
		kvfilters.OmitWhenValueMatchedByOneOfRegexps{
			// Numbers with many digits
			`(-?\d{9,}|0\.\d{8,}|\d{7,}\.\d+)`,
			// Locations, Sizes
			`\{[0-9.]+,\s*[0-9.]+\}`,
			// Timestamps
			`\d{9,10}\.\d{5,7}`,
		},
		kvfilters.OmitWhenValueMatchedByOneOfRegexpsIgnoringCase{
			// UUIDs
			`[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}`,
			// Screen coordinates
			`\{\{-?\d+, -?\d+\}, \{-?\d+, -?\d+\}\}`,
		},
		//kvfilters.OmitWhenGroupPassedToFuncReturnsFalse{
		//	func(s string) bool {
		//		return s == ""
		//	},
		//},
	}
}
