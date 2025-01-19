package applemdm_test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/mikeschinkel/prefsctl/applemdm"
	"github.com/mikeschinkel/prefsctl/stdlibex"
)

const (
	gitCloneDir          = "/tmp/applemdm-test"
	gitHubCloneURLPrefix = "https://github.com"
)

type (
	PreferenceManifest = applemdm.PreferenceManifest
	ManifestKey        = applemdm.ManifestKey
)

var (
	ProfileManifestsRepoURL = "https://github.com/ProfileManifests/ProfileManifests"
	ProfileManifestsDir     = fmt.Sprintf("%s/%s", gitCloneDir, filepath.Base(ProfileManifestsRepoURL))

	DevMgmtRepoURL = "https://github.com/apple/device-management/"
	DevMgmtDir     = fmt.Sprintf("%s/%s", gitCloneDir, filepath.Base(DevMgmtRepoURL))

	LoadPreferenceManifest = applemdm.LoadPreferenceManifest
)

func filePath(path string) string {
	return fmt.Sprintf("%s/Manifests/%s", ProfileManifestsDir, path)
}
func TestMain(m *testing.M) {
	// Setup code before any tests run
	setup()

	// Run the tests
	code := m.Run()

	// Cleanup code after all tests complete
	cleanup()

	// Exit with the status code from the tests
	os.Exit(code)
}

func setup() {
	exists, err := stdlibex.DirExists(ProfileManifestsDir)
	if err != nil {
		panic(err)
	}
	var cmd *exec.Cmd
	if !exists {
		cmd = exec.Command("git", "clone", ProfileManifestsRepoURL, ProfileManifestsDir)
	} else {
		cmd = exec.Command("git", "-C", ProfileManifestsDir, "pull")
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "git command failed: %v\noutput: %s", err, output)
	}
}

func cleanup() {
}

func Test_LoadPreferenceManifest(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		numSubkeys int
		descr      string
		wantErr    bool
	}{
		{
			name:       "com.apple.dock.plist",
			path:       "ManifestsApple/com.apple.dock.plist",
			descr:      "Use this section to manage the Dock.",
			numSubkeys: 47,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := applemdm.LoadPreferenceManifestFromFile(filePath(tt.path))
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadPreferenceManifest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			{
				got, want := got.Description, tt.descr
				if want != got {
					t.Errorf("LoadPreferenceManifest() descr = %v, want %v", got, want)
				}
			}
			{
				got, want := len(got.Subkeys), tt.numSubkeys
				if got != want {
					t.Errorf("LoadPreferenceManifest() subkey count = %v, want %v", got, want)
				}
			}
		})
	}
}

func Test_GetBoolKey(t *testing.T) {
	type fields struct {
		Description      string
		DescriptionRef   string
		DocumentationURL string
		Domain           string
		FormatVersion    int
		Interaction      string
		LastModified     time.Time
		MacOSMin         string
		Platforms        []string
		Subkeys          []ManifestKey
		Targets          []string
		Title            string
		Unique           bool
		Version          int
	}
	type args struct {
		name string
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantKey          *ManifestKey
		wantDefaultValue bool
		wantErr          bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PreferenceManifest{
				Description:      tt.fields.Description,
				DescriptionRef:   tt.fields.DescriptionRef,
				DocumentationURL: tt.fields.DocumentationURL,
				Domain:           tt.fields.Domain,
				FormatVersion:    tt.fields.FormatVersion,
				Interaction:      tt.fields.Interaction,
				LastModified:     tt.fields.LastModified,
				MacOSMin:         tt.fields.MacOSMin,
				Platforms:        tt.fields.Platforms,
				Subkeys:          tt.fields.Subkeys,
				Targets:          tt.fields.Targets,
				Title:            tt.fields.Title,
				Unique:           tt.fields.Unique,
				Version:          tt.fields.Version,
			}
			gotKey, gotDefaultValue, err := m.GetBoolKey(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBoolKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotKey, tt.wantKey) {
				t.Errorf("GetBoolKey() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
			if gotDefaultValue != tt.wantDefaultValue {
				t.Errorf("GetBoolKey() gotDefaultValue = %v, want %v", gotDefaultValue, tt.wantDefaultValue)
			}
		})
	}
}

func Test_GetIntKey(t *testing.T) {
	type fields struct {
		Description      string
		DescriptionRef   string
		DocumentationURL string
		Domain           string
		FormatVersion    int
		Interaction      string
		LastModified     time.Time
		MacOSMin         string
		Platforms        []string
		Subkeys          []ManifestKey
		Targets          []string
		Title            string
		Unique           bool
		Version          int
	}
	type args struct {
		name string
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantKey          *ManifestKey
		wantDefaultValue int
		wantErr          bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PreferenceManifest{
				Description:      tt.fields.Description,
				DescriptionRef:   tt.fields.DescriptionRef,
				DocumentationURL: tt.fields.DocumentationURL,
				Domain:           tt.fields.Domain,
				FormatVersion:    tt.fields.FormatVersion,
				Interaction:      tt.fields.Interaction,
				LastModified:     tt.fields.LastModified,
				MacOSMin:         tt.fields.MacOSMin,
				Platforms:        tt.fields.Platforms,
				Subkeys:          tt.fields.Subkeys,
				Targets:          tt.fields.Targets,
				Title:            tt.fields.Title,
				Unique:           tt.fields.Unique,
				Version:          tt.fields.Version,
			}
			gotKey, gotDefaultValue, err := m.GetIntKey(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIntKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotKey, tt.wantKey) {
				t.Errorf("GetIntKey() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
			if gotDefaultValue != tt.wantDefaultValue {
				t.Errorf("GetIntKey() gotDefaultValue = %v, want %v", gotDefaultValue, tt.wantDefaultValue)
			}
		})
	}
}

func Test_GetKey(t *testing.T) {
	type fields struct {
		Description      string
		DescriptionRef   string
		DocumentationURL string
		Domain           string
		FormatVersion    int
		Interaction      string
		LastModified     time.Time
		MacOSMin         string
		Platforms        []string
		Subkeys          []ManifestKey
		Targets          []string
		Title            string
		Unique           bool
		Version          int
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantMk  *ManifestKey
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PreferenceManifest{
				Description:      tt.fields.Description,
				DescriptionRef:   tt.fields.DescriptionRef,
				DocumentationURL: tt.fields.DocumentationURL,
				Domain:           tt.fields.Domain,
				FormatVersion:    tt.fields.FormatVersion,
				Interaction:      tt.fields.Interaction,
				LastModified:     tt.fields.LastModified,
				MacOSMin:         tt.fields.MacOSMin,
				Platforms:        tt.fields.Platforms,
				Subkeys:          tt.fields.Subkeys,
				Targets:          tt.fields.Targets,
				Title:            tt.fields.Title,
				Unique:           tt.fields.Unique,
				Version:          tt.fields.Version,
			}
			gotMk, err := m.GetKey(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMk, tt.wantMk) {
				t.Errorf("GetKey() gotMk = %v, want %v", gotMk, tt.wantMk)
			}
		})
	}
}

func Test_GetStringKey(t *testing.T) {
	type fields struct {
		Description      string
		DescriptionRef   string
		DocumentationURL string
		Domain           string
		FormatVersion    int
		Interaction      string
		LastModified     time.Time
		MacOSMin         string
		Platforms        []string
		Subkeys          []ManifestKey
		Targets          []string
		Title            string
		Unique           bool
		Version          int
	}
	type args struct {
		name string
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantKey          *ManifestKey
		wantDefaultValue string
		wantErr          bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PreferenceManifest{
				Description:      tt.fields.Description,
				DescriptionRef:   tt.fields.DescriptionRef,
				DocumentationURL: tt.fields.DocumentationURL,
				Domain:           tt.fields.Domain,
				FormatVersion:    tt.fields.FormatVersion,
				Interaction:      tt.fields.Interaction,
				LastModified:     tt.fields.LastModified,
				MacOSMin:         tt.fields.MacOSMin,
				Platforms:        tt.fields.Platforms,
				Subkeys:          tt.fields.Subkeys,
				Targets:          tt.fields.Targets,
				Title:            tt.fields.Title,
				Unique:           tt.fields.Unique,
				Version:          tt.fields.Version,
			}
			gotKey, gotDefaultValue, err := m.GetStringKey(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStringKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotKey, tt.wantKey) {
				t.Errorf("GetStringKey() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
			if gotDefaultValue != tt.wantDefaultValue {
				t.Errorf("GetStringKey() gotDefaultValue = %v, want %v", gotDefaultValue, tt.wantDefaultValue)
			}
		})
	}
}

func Test_SavePreferenceManifest(t *testing.T) {
	type fields struct {
		Description      string
		DescriptionRef   string
		DocumentationURL string
		Domain           string
		FormatVersion    int
		Interaction      string
		LastModified     time.Time
		MacOSMin         string
		Platforms        []string
		Subkeys          []ManifestKey
		Targets          []string
		Title            string
		Unique           bool
		Version          int
	}
	tests := []struct {
		name    string
		fields  fields
		wantW   string
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PreferenceManifest{
				Description:      tt.fields.Description,
				DescriptionRef:   tt.fields.DescriptionRef,
				DocumentationURL: tt.fields.DocumentationURL,
				Domain:           tt.fields.Domain,
				FormatVersion:    tt.fields.FormatVersion,
				Interaction:      tt.fields.Interaction,
				LastModified:     tt.fields.LastModified,
				MacOSMin:         tt.fields.MacOSMin,
				Platforms:        tt.fields.Platforms,
				Subkeys:          tt.fields.Subkeys,
				Targets:          tt.fields.Targets,
				Title:            tt.fields.Title,
				Unique:           tt.fields.Unique,
				Version:          tt.fields.Version,
			}
			w := bytes.Buffer{}
			err := m.SavePreferenceManifest(&w)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Save() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_SaveToFile(t *testing.T) {
	type fields struct {
		Description      string
		DescriptionRef   string
		DocumentationURL string
		Domain           string
		FormatVersion    int
		Interaction      string
		LastModified     time.Time
		MacOSMin         string
		Platforms        []string
		Subkeys          []ManifestKey
		Targets          []string
		Title            string
		Unique           bool
		Version          int
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PreferenceManifest{
				Description:      tt.fields.Description,
				DescriptionRef:   tt.fields.DescriptionRef,
				DocumentationURL: tt.fields.DocumentationURL,
				Domain:           tt.fields.Domain,
				FormatVersion:    tt.fields.FormatVersion,
				Interaction:      tt.fields.Interaction,
				LastModified:     tt.fields.LastModified,
				MacOSMin:         tt.fields.MacOSMin,
				Platforms:        tt.fields.Platforms,
				Subkeys:          tt.fields.Subkeys,
				Targets:          tt.fields.Targets,
				Title:            tt.fields.Title,
				Unique:           tt.fields.Unique,
				Version:          tt.fields.Version,
			}
			if err := m.SavePreferenceManifestFile(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("SaveToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Validate(t *testing.T) {
	type fields struct {
		Description      string
		DescriptionRef   string
		DocumentationURL string
		Domain           string
		FormatVersion    int
		Interaction      string
		LastModified     time.Time
		MacOSMin         string
		Platforms        []string
		Subkeys          []ManifestKey
		Targets          []string
		Title            string
		Unique           bool
		Version          int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PreferenceManifest{
				Description:      tt.fields.Description,
				DescriptionRef:   tt.fields.DescriptionRef,
				DocumentationURL: tt.fields.DocumentationURL,
				Domain:           tt.fields.Domain,
				FormatVersion:    tt.fields.FormatVersion,
				Interaction:      tt.fields.Interaction,
				LastModified:     tt.fields.LastModified,
				MacOSMin:         tt.fields.MacOSMin,
				Platforms:        tt.fields.Platforms,
				Subkeys:          tt.fields.Subkeys,
				Targets:          tt.fields.Targets,
				Title:            tt.fields.Title,
				Unique:           tt.fields.Unique,
				Version:          tt.fields.Version,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
