package filters

import (
	"errors"
	"fmt"

	"github.com/mikeschinkel/prefsctl/logging"
	"github.com/mikeschinkel/prefsctl/macosutils"
)

var queryFiltersByOS = map[macosutils.Code][]Filter{
	macosutils.Monterey: MontereyQueryFilters(),
	macosutils.Sequoia:  SequoiaQueryFilters(),
}

func QueryFilters() (ff []Filter, err error) {
	var ok bool
	code, err := macosutils.VersionCode()
	if err != nil {
		goto end
	}
	ff, ok = queryFiltersByOS[code]
	if !ok {
		err = errors.Join(
			ErrInvalidFilterMatchCriteria,
			fmt.Errorf("%s=%s", logging.OSCodeLogArg, code),
		)
	}
end:
	return ff, nil
}

func GroupsQueryFilters() (ff []Filter, err error) {
	qfs, err := QueryFilters()
	if err != nil {
		goto end
	}
	ff = make([]Filter, 0, len(qfs))
	for _, f := range qfs {
		if f.Target() != Groups {
			continue
		}
		ff = append(ff, f)
	}
end:
	return ff, nil
}
