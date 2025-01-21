package profilemanifests_test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/micromdm/plist"
	"github.com/mikeschinkel/prefsctl/profilemanifests"
	"github.com/mikeschinkel/prefsctl/stdlibex"
)

const (
	gitCloneDir          = "/tmp/profilemanifests-test"
	gitHubCloneURLPrefix = "https://github.com"
)

type (
	ProfileManifest = profilemanifests.ProfileManifest
	ManifestKey     = profilemanifests.ManifestKey
)

var (
	ProfileManifestsRepoURL = "https://github.com/ProfileManifests/ProfileManifests"
	ProfileManifestsDir     = fmt.Sprintf("%s/%s", gitCloneDir, filepath.Base(ProfileManifestsRepoURL))

	DevMgmtRepoURL = "https://github.com/apple/device-management/"
	DevMgmtDir     = fmt.Sprintf("%s/%s", gitCloneDir, filepath.Base(DevMgmtRepoURL))
)
var (
	LoadProfileManifest       = profilemanifests.LoadProfileManifest
	ErrFailedToDecodeManifest = profilemanifests.ErrFailedToDecodeManifest
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

func Test_DecodeProfileManifestWithMPFMRangeMinAsReal(t *testing.T) {
	const input = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>pfm_subkeys</key>
	<array>
		<dict>
			<key>pfm_range_min</key>
			<real>42</real>
			<key>pfm_type</key>
			<string>real</string>
		</dict>
	</array>
</dict>
</plist>`
	var err error

	bb := bytes.NewBufferString(input)
	decoder := plist.NewXMLDecoder(bb)
	pm := ProfileManifest{}
	err = decoder.Decode(&pm)
	if err != nil {
		t.Error(err)
		return
	}
	if pm.Subkeys == nil {
		t.Error(".Subkeys should not be nil")
		return
	}
	if len(pm.Subkeys) != 1 {
		t.Errorf("len(.Subkeys) should be 1, not %d", len(pm.Subkeys))
		return
	}
	rMin, ok := pm.Subkeys[0].RangeMinAny.(float64)
	if !ok {
		t.Errorf(".Subkeys[].RangeMin should type assert to 'float64', got '%T' instead", pm.Subkeys[0].RangeMinAny)
		return
	}
	if int(rMin) != 42 {
		t.Errorf(".Subkeys[0].RangeMin should be 42.000000, not %f", rMin)
		return
	}
}
func Test_DecodeNoteDeveloperProfileCreatorManifest(t *testing.T) {
	const input = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>pfm_description</key>
	<string>Developer: Note settings</string>
	<key>pfm_domain</key>
	<string>com.profilecreator.developer.note</string>
	<key>pfm_format_version</key>
	<integer>1</integer>
	<key>pfm_interaction</key>
	<string>exclusive</string>
	<key>pfm_last_modified</key>
	<date>2018-07-18T08:58:49Z</date>
	<key>pfm_platforms</key>
	<array>
		<string>iOS</string>
		<string>macOS</string>
		<string>tvOS</string>
	</array>
	<key>pfm_subkeys</key>
	<array>
		<dict>
			<key>pfm_default</key>
			<string>Configures Developer: Note settings</string>
			<key>pfm_description</key>
			<string>Description of the payload</string>
			<key>pfm_name</key>
			<string>PayloadDescription</string>
			<key>pfm_title</key>
			<string>Payload Description</string>
			<key>pfm_type</key>
			<string>string</string>
		</dict>
		<dict>
			<key>pfm_default</key>
			<string>Developer: Note</string>
			<key>pfm_description</key>
			<string>Name of the payload</string>
			<key>pfm_name</key>
			<string>PayloadDisplayName</string>
			<key>pfm_require</key>
			<string>always</string>
			<key>pfm_title</key>
			<string>Payload Display Name</string>
			<key>pfm_type</key>
			<string>string</string>
		</dict>
		<dict>
			<key>pfm_default</key>
			<string>com.profilecreator.developer.note</string>
			<key>pfm_description</key>
			<string>A unique identifier for the payload, dot-delimited.  Usually root PayloadIdentifier+subidentifier</string>
			<key>pfm_name</key>
			<string>PayloadIdentifier</string>
			<key>pfm_require</key>
			<string>always</string>
			<key>pfm_title</key>
			<string>Payload Identifier</string>
			<key>pfm_type</key>
			<string>string</string>
		</dict>
		<dict>
			<key>pfm_default</key>
			<string>com.profilecreator.developer.note</string>
			<key>pfm_description</key>
			<string>The type of the payload, a reverse dns string</string>
			<key>pfm_name</key>
			<string>PayloadType</string>
			<key>pfm_require</key>
			<string>always</string>
			<key>pfm_title</key>
			<string>Payload Type</string>
			<key>pfm_type</key>
			<string>string</string>
		</dict>
		<dict>
			<key>pfm_default</key>
			<string></string>
			<key>pfm_description</key>
			<string>Unique identifier for the payload (format 01234567-89AB-CDEF-0123-456789ABCDEF)</string>
			<key>pfm_format</key>
			<string>^[0-9A-Za-z]{8}-[0-9A-Za-z]{4}-[0-9A-Za-z]{4}-[0-9A-Za-z]{4}-[0-9A-Za-z]{12}$</string>
			<key>pfm_name</key>
			<string>PayloadUUID</string>
			<key>pfm_require</key>
			<string>always</string>
			<key>pfm_title</key>
			<string>Payload UUID</string>
			<key>pfm_type</key>
			<string>string</string>
		</dict>
		<dict>
			<key>pfm_default</key>
			<integer>1</integer>
			<key>pfm_description</key>
			<string>The version of the whole configuration profile.</string>
			<key>pfm_name</key>
			<string>PayloadVersion</string>
			<key>pfm_require</key>
			<string>always</string>
			<key>pfm_title</key>
			<string>Payload Version</string>
			<key>pfm_type</key>
			<string>integer</string>
		</dict>
		<dict>
			<key>pfm_description</key>
			<string>This value describes the issuing organization of the profile, as displayed to the user</string>
			<key>pfm_name</key>
			<string>PayloadOrganization</string>
			<key>pfm_title</key>
			<string>Payload Organization</string>
			<key>pfm_type</key>
			<string>string</string>
		</dict>
		
		<!-- Custom Keys -->
		
		<dict>
			<key>pfm_description</key>
			<string>TextField Note.</string>
			<key>pfm_enabled</key>
			<true/>
			<key>pfm_name</key>
			<string>Note01</string>
			<key>pfm_title</key>
			<string>Note 01: TextField</string>
			<key>pfm_type</key>
			<string>string</string>
			<key>pfm_note</key>
			<string>This is a short note</string>
		</dict>
	</array>
	<key>pfm_targets</key>
	<array>
		<string>user</string>
		<string>system</string>
	</array>
	<key>pfm_title</key>
	<string>Developer: Note</string>
	<key>pfm_unique</key>
	<true/>
	<key>pfm_version</key>
	<integer>1</integer>
</dict>
</plist>`
	var err error

	bb := bytes.NewBufferString(input)
	decoder := plist.NewXMLDecoder(bb)
	pm := ProfileManifest{}
	err = decoder.Decode(&pm)
	if err != nil {
		t.Error(err)
		return
	}
}

func Test_LoadProfileManifest(t *testing.T) {
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
			got, err := profilemanifests.LoadProfileManifestFromFile(filePath(tt.path))
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadProfileManifest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			{
				got, want := got.Description, tt.descr
				if want != got {
					t.Errorf("LoadProfileManifest() descr = %v, want %v", got, want)
				}
			}
			{
				got, want := len(got.Subkeys), tt.numSubkeys
				if got != want {
					t.Errorf("LoadProfileManifest() subkey count = %v, want %v", got, want)
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
			m := &ProfileManifest{
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
			m := &ProfileManifest{
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
			m := &ProfileManifest{
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
			m := &ProfileManifest{
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

func Test_SaveProfileManifest(t *testing.T) {
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
			m := &ProfileManifest{
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
			err := m.SaveProfileManifest(&w)
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
			m := &ProfileManifest{
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
			if err := m.SaveProfileManifestFile(tt.args.path); (err != nil) != tt.wantErr {
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
			m := &ProfileManifest{
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
