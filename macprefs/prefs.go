package macprefs

type Map map[PrefId]*Pref

var PrefsMap Map

func Initialize() {
	for id, pd := range PrefDefaultsMap() {
		// Create lookup for pref default
		PrefsMap[id] = &Pref{
			PrefDefault: pd,
			Kind:        ReflectKindFromPrefType(pd.Type),
		}
	}

}
