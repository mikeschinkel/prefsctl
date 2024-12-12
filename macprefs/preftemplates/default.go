package preftemplates

import (
	"strings"

	"github.com/mikeschinkel/prefsctl/kvfilters"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Default struct {
	Name     PrefName
	Value    string // raw string value for default
	Labels   *kvfilters.Labels
	Type     TypeName
	Verified Verified
	Domain   *Domain
}

func (d Default) TypeName() string {
	typ, _ := strings.CutSuffix(string(d.Type), "Type")
	return typ
}

type Verified []kvfilters.LabelName

func (vv Verified) String() string {
	sb := strings.Builder{}
	sb.WriteString("Verified{")
	titler := cases.Title(language.English)
	last := len(vv) - 1
	for i, labelName := range vv {
		sb.WriteString(titler.String(string(labelName)))
		if i < last {
			sb.WriteByte(',')
		}
	}
	sb.WriteByte('}')
	return sb.String()
}
