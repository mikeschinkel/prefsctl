package cobrautil

type Props interface {
	Props()
}

type propsParser interface {
	Parse(*Config) Props
}
type propsValidator interface {
	Validate(*Config) error
}

func ProcessProps[P Props](cfg *Config, props P) (_ P, err error) {
	var parser propsParser

	anyProps := any(props)
	validator, ok := anyProps.(propsValidator)
	if ok {
		// Only validate if it implements the method
		err = validator.Validate(cfg)
	}
	parser, ok = anyProps.(propsParser)
	if !ok {
		// No parser? Just return the passed-in props
		goto end
	}
	props = parser.Parse(cfg).(P)
end:
	return props, err
}
