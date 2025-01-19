package applemdm_test

import (
	"testing"

	"github.com/mikeschinkel/prefsctl/applemdm"
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
			pm := applemdm.NewProfileManifests()
			err := pm.Load()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListProfileManifests() error = %v, wantErr %v", err, tt.wantErr)
			}
			if pm.Files() == nil {
				t.Errorf("ListProfileManifests() profile manifests not nil")
			}
		})
	}
}
