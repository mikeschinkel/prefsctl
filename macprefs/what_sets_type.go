package macprefs

type WhatSetsType string

const (
	UnknownWhatSets WhatSetsType = ""
	MacOSSets       WhatSetsType = "macos"
	InstallSets     WhatSetsType = "install"
	UnsureWhatSets  WhatSetsType = "unsure"
)
