package filters

import (
	"reflect"
	"testing"
)

func TestQuery(t *testing.T) {
	tests := []struct {
		name       string
		args       QueryArgs
		wantResult []Group
		wantErr    bool
	}{
		{
			name: "Omit when key suffix is 'count'",
			args: QueryArgs{
				Filters: []Filter{
					OmitWhenKeySuffixIsOneOfIgnoringCase{
						"count",
					},
				},
				Groups: []Group{
					&group{
						name: "com.apple.accessibility.universalAccessAuthWarn",
						code: "com.apple.accessibility.universalAccessAuthWarn",
						keyValues: []KeyValue{
							&keyValue{
								key:   "ThirdPartyCount",
								value: "2",
							},
							&keyValue{
								key:   "GoogleAccount",
								value: "foobar",
							},
						},
						initialized: true,
					},
				},
			},
			wantResult: []Group{},
			wantErr:    false,
		},
		//{
		//	name: "Do not omit com.apple.dock/autohide",
		//	args: QueryArgs{
		//		Filters: MontereyQueryFilters(),
		//		Groups: []Group{
		//			&group{
		//				name: "com.apple.dock",
		//				code: "com.apple.dock",
		//				keyValues: []KeyValue{
		//					&keyValue{
		//						key:   "autohide",
		//						value: "true",
		//					},
		//				},
		//				initialized: true,
		//			},
		//		},
		//	},
		//	wantResult: []Group{},
		//	wantErr:    false,
		//},
		//{
		//	name: "Omit key seed-numNotifications-",
		//	args: QueryArgs{
		//		Filters: []Filter{
		//			OmitWhenKeyPrefixIsOneOf{
		//				"seed-numNotifications-",
		//			},
		//		},
		//		Groups: []Group{
		//			&group{
		//				name: "com.apple.tourist",
		//				code: "com.apple.tourist",
		//				keyValues: []KeyValue{
		//					&keyValue{
		//						key:   "seed-numNotifications-+trJt2VsTvK1yfPGwOySDw:",
		//						value: "foo bar",
		//					},
		//				},
		//				initialized: true,
		//			},
		//		},
		//	},
		//	wantResult: []Group{},
		//	wantErr:    false,
		//},
		//{
		//	name: "Keep key seed-numNotifications-",
		//	args: QueryArgs{
		//		Filters: []Filter{
		//			KeepWhenKeyPrefixIsOneOf{
		//				"seed-numNotifications-",
		//			},
		//		},
		//		Groups: []Group{
		//			&group{
		//				name: "com.apple.tourist",
		//				code: "com.apple.tourist",
		//				keyValues: []KeyValue{
		//					&keyValue{
		//						key:   "seed-numNotifications-+trJt2VsTvK1yfPGwOySDw:",
		//						value: "foo bar",
		//					},
		//				},
		//				//initialized: true,
		//			},
		//		},
		//	},
		//	wantResult: []Group{
		//		&group{
		//			name: "com.apple.tourist",
		//			code: "com.apple.tourist",
		//			keyValues: []KeyValue{
		//				&keyValue{
		//					key:   "seed-numNotifications-+trJt2VsTvK1yfPGwOySDw:",
		//					value: "foo bar",
		//				},
		//			},
		//			//initialized: true,
		//		},
		//	},
		//	wantErr: false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := Query(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Query() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
