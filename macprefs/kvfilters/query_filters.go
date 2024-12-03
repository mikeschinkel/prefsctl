package kvfilters

import (
	"errors"
	"fmt"
	"slices"

	"github.com/mikeschinkel/prefsctl/macosutils"
	"github.com/mikeschinkel/prefsctl/macprefs/logargs"
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
			fmt.Errorf("%s=%s", logargs.OSCodeLogArg, code),
		)
	}
end:
	return ff, nil
}

func QueryFiltersForTargets(targets ...Target) (ff []Filter, err error) {
	qfs, err := QueryFilters()
	if err != nil {
		goto end
	}
	ff = make([]Filter, 0, len(qfs))
	for _, f := range qfs {
		if !slices.Contains(targets, f.Target()) {
			continue
		}
		ff = append(ff, f)
	}
end:
	return ff, nil
}
