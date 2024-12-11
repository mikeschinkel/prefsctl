//go:build debug

package macprefs

func init() {
	pd := &PrefDomains{}
	pd.DebugString()
	p := &Pref{}
	p.DebugString()
}
