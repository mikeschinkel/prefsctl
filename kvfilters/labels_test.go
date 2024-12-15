package kvfilters_test

import (
	"testing"

	"github.com/mikeschinkel/prefsctl/kvfilters"
)

type Label = kvfilters.Label
type Labels = kvfilters.Labels
type LabelName = kvfilters.LabelName
type LabelValue = kvfilters.LabelValue

var NewLabel = kvfilters.NewLabel
var NewLabels = kvfilters.NewLabels

const (
	Size  LabelName = "size"
	Color LabelName = "color"
)

var (
	Small  = NewLabel(Size, "small")
	Medium = NewLabel(Size, "medium")
	Large  = NewLabel(Size, "large")

	Red   = NewLabel(Color, "red")
	Green = NewLabel(Color, "green")
	Blue  = NewLabel(Color, "blue")
)

func TestLabels_ContainsAny(t *testing.T) {
	tests := []struct {
		name         string
		haveLabels   *Labels
		testLabels   []*Label
		wantContains bool
	}{
		{
			name:         "Small, Red has (one of) Red, Blue",
			haveLabels:   kvfilters.NewLabels(&Small, &Red),
			testLabels:   []*Label{&Red, &Blue},
			wantContains: true,
		},
		{
			name:         "Small, Red does not have Green",
			haveLabels:   kvfilters.NewLabels(&Small, &Red),
			testLabels:   []*Label{&Green},
			wantContains: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotContains := tt.haveLabels.ContainsAny(tt.testLabels...)
			if gotContains != tt.wantContains {
				t.Errorf("ContainsAny() = %v, want %v", gotContains, tt.wantContains)
			}
		})
	}
}

func TestLabelName_String(t *testing.T) {
	tests := []struct {
		name  string
		name1 LabelName
		want  string
	}{
		{name: "Foo Name String", name1: "Foo", want: "Foo"},
		{name: "Bar Name String", name1: "Bar", want: "Bar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.name1.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLabelValue_String(t *testing.T) {
	tests := []struct {
		name  string
		value LabelValue
		want  string
	}{
		{name: "Foo Value String", value: "Foo", want: "Foo"},
		{name: "Bar Value String", value: "Bar", want: "Bar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.value.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLabel_String(t *testing.T) {
	type fields struct {
		Name      LabelName
		Value     LabelValue
		IsUnknown bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Label Foo=Bar",
			fields: fields{
				Name:      "Foo",
				Value:     "Bar",
				IsUnknown: false,
			},
			want: "Foo=Bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Label{
				Name:      tt.fields.Name,
				Value:     tt.fields.Value,
				IsUnknown: tt.fields.IsUnknown,
			}
			if got := l.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLabels_GetNamedLabel(t *testing.T) {
	tests := []struct {
		name       string
		haveLabels *Labels
		labelName  LabelName
		wantLabel  *Label
	}{
		{
			name:       "GetNamedLabel for known 'color' label: Green",
			haveLabels: NewLabels(&Large, &Green),
			labelName:  Color,
			wantLabel:  &Green,
		},
		{
			name:       "GetNamedLabel for known 'size' label: Medium",
			haveLabels: NewLabels(&Medium, &Red),
			labelName:  Size,
			wantLabel:  &Medium,
		},
		{
			name:       "GetNamedLabel for 'unknown' label: nil",
			haveLabels: NewLabels(&Small, &Blue),
			labelName:  "unknown",
			wantLabel:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLabel := tt.haveLabels.GetNamedLabel(tt.labelName)
			if gotLabel != tt.wantLabel {
				t.Errorf("GetNamedLabel() = %v, want %v", gotLabel, tt.wantLabel)
			}
		})
	}
}
func TestLabels_HasNamedLabel(t *testing.T) {
	tests := []struct {
		name       string
		haveLabels *Labels
		labelName  LabelName
		want       bool
	}{
		{
			name:       "HasNamedLabel for green 'color' label: true",
			haveLabels: NewLabels(&Large, &Green),
			labelName:  Color,
			want:       true,
		},
		{
			name:       "HasNamedLabel for known medium 'size' label",
			haveLabels: NewLabels(&Medium, &Red),
			labelName:  Size,
			want:       true,
		},
		{
			name:       "HasNamedLabel for 'unknown' label: nil",
			haveLabels: NewLabels(&Small, &Blue),
			labelName:  "unknown",
			want:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.haveLabels.HasNamedLabel(tt.labelName)
			if got != tt.want {
				t.Errorf("HasNamedLabel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLabels_SetNamedLabel(t *testing.T) {
	labels := NewLabels(&Small)
	if !labels.HasNamedLabel(Size) {
		t.Errorf("HasNamedLabel(Size) expected to be true for 'small' 'size'; false instead")
	}
	if labels.HasNamedLabel(Color) {
		t.Errorf("HasNamedLabel(Color) expected to be false for 'color'; true instead")
	}
	labels.SetLabel(&Blue)
	if !labels.HasNamedLabel(Color) {
		t.Errorf("HasNamedLabel(Color) expected to be true for 'blue' 'color'; false instead")
	}
	labels.SetLabel(&Large)
	if !labels.HasNamedLabel(Size) {
		t.Errorf("HasNamedLabel(Size) expected to be true for 'large' 'size'; false instead")
	}
	if !labels.HasLabel(&Large) {
		t.Errorf("HasLabel(&Large) expected to be true; false instead")
	}
	if labels.HasLabel(&Small) {
		t.Errorf("HasLabel(&Small) expected to be false; true instead")
	}
	if !labels.Valid() {
		t.Errorf("Valid() expected to be true; false instead")
	}
}

func TestLabels_String(t *testing.T) {
	tests := []struct {
		name       string
		haveLabels *Labels
		want       string
	}{
		{
			name:       "Labels string",
			haveLabels: NewLabels(&Red, &Medium),
			want:       "color=red,size=medium",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.haveLabels.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLabels_LabelPtrs(t *testing.T) {
	tests := []struct {
		name       string
		haveLabels *Labels
		wantLabels *Labels
		wantMatch  bool
	}{
		{
			name:       "LabelPtrs Want Match",
			haveLabels: NewLabels(&Medium, &Red),
			wantLabels: NewLabels(&Red, &Medium),
			wantMatch:  true,
		},
		{
			name:       "LabelPtrs Want No Match",
			haveLabels: NewLabels(&Medium, &Red),
			wantLabels: NewLabels(&Blue, &Small),
			wantMatch:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLabels := NewLabels(tt.haveLabels.LabelPtrs()...)
			if tt.wantMatch && !gotLabels.Equivalent(tt.wantLabels) {
				t.Errorf("No match: LabelPtrs() = %v, want %v", gotLabels, tt.wantLabels)
			}
			if !tt.wantMatch && gotLabels.Equivalent(tt.wantLabels) {
				t.Errorf("Match: LabelPtrs() = %v, want not %v", gotLabels, tt.wantLabels)
			}
		})
	}
}
