package macprefs

//type MacOSVersion struct {
//	Name         macosutils.Name
//	PrefDefaults []*PrefDefault
//}
//
//// tmplPrefs represents the structure passed to the template
//type tmplPrefs struct {
//	Name         string
//	Domains      PrefDomains
//	PrefDefaults [][]tmplPref
//}
//
//// tmplPref represents a preference formatted for template output
//type tmplPref PrefDefault
//
//// ShowValue returns true if there is no label "install-sets" — This means it is
//// either a true default or we do not yet know for sure, but we do not that it is
//// not determine exclusively by the person installing the OS.
//func (p tmplPref) ShowValue() bool {
//	return !slices.ContainsFunc(p.Labels, func(label *Label) bool {
//		return label.Value == SetupSets.Value
//	})
//}
//
//// GoCode generates Go code for preference defaults
//func (m *MacOSVersion) GoCode() (string, error) {
//	var tmpl = goMacOSPrefsFuncTemplate
//	// Group preferences by domain while preserving order
//	seen := make(map[DomainName]int)
//	var data tmplPrefs
//	data.Name = string(m.Name)
//
//	// First pass: collect domains in order of appearance
//	for _, pref := range m.PrefDefaults {
//		if _, exists := seen[pref.Domain]; !exists {
//			seen[pref.Domain] = len(data.Domains)
//			data.Domains = append(data.Domains, NewPrefsDomain(pref.Domain))
//			data.PrefDefaults = append(data.PrefDefaults, make([]tmplPref, 0))
//		}
//	}
//
//	// Second pass: populate preferences
//	for _, pref := range m.PrefDefaults {
//		idx := seen[pref.Domain]
//
//		tPref := tmplPref{
//			Name:   pref.Name,
//			Value:  pref.Value,
//			Labels: pref.Labels,
//		}
//
//		data.PrefDefaults[idx] = append(data.PrefDefaults[idx], tPref)
//	}
//
//	// Parse and execute template
//	t, err := template.New("prefDefaults").Parse(tmpl)
//	if err != nil {
//		return "", err
//	}
//
//	var buf bytes.Buffer
//	if err := t.Execute(&buf, data); err != nil {
//		return "", err
//	}
//
//	return buf.String(), nil
//}
//
//const goMacOSPrefsFuncTemplate = `func {{.Name}}PrefDefaults() macprefs.DomainPrefDefaults {
//	return macprefs.DomainPrefDefaults{
//		{{- range $i, $domain := .Domains}}
//		"{{$domain}}": []*macprefs.PrefDefault{
//			{{- range $pref := index $.PrefDefaults $i}}
//			{
//				Name:     "{{$pref.Name}}",
//				Type:     {{$pref.TypeName}},
//				{{- if $pref.ShowValue}}
//				Value:    "{{$pref.Value}}",
//				{{- end}}
//				Labels: []macprefs.Label{
//				{{- range $i, $label := .Labels}}
//					{{$label.GoProperty}},
//				{{- end}}
//				},
//			},
//			{{- end}}
//		},
//		{{- end}}
//	}
//}`
