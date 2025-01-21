package profilemanifests_test

import (
	"testing"

	"github.com/mikeschinkel/prefsctl/profilemanifests"
)

func Test_Load(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pms := profilemanifests.New()
			err := pms.Load()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListProfileManifests() error = %v, wantErr %v", err, tt.wantErr)
			}
			if pms.Files() == nil {
				t.Errorf("ListProfileManifests() profile manifests not nil")
			}
		})
	}
}
