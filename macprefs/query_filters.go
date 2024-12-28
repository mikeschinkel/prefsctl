package macprefs

import (
	"errors"
	"fmt"
	"slices"

	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/logargs"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macosutil/osversionfilters"
	"github.com/mikeschinkel/prefsctl/types"
)

var queryFiltersByOS = map[types.Code][]kvfilters.Filter{
	macosutil.Monterey: osversionfilters.MontereyQueryFilters(),
	macosutil.Sequoia:  osversionfilters.SequoiaQueryFilters(),
}

func QueryFilters() (ff []kvfilters.Filter, err error) {
	var ok bool
	code, err := macosutil.VersionCode()
	if err != nil {
		goto end
	}
	ff, ok = queryFiltersByOS[code]
	if !ok {
		err = errors.Join(
			kvfilters.ErrInvalidFilterMatchCriteria,
			fmt.Errorf("%s=%s", logargs.OSCodeLogArg, code),
		)
	}
end:
	return ff, nil
}

func QueryFiltersForTargets(targets ...kvfilters.Target) (ff []kvfilters.Filter, err error) {
	qfs, err := QueryFilters()
	if err != nil {
		goto end
	}
	ff = make([]kvfilters.Filter, 0, len(qfs))
	for _, f := range qfs {
		if !slices.Contains(targets, f.Target()) {
			continue
		}
		ff = append(ff, f)
	}
end:
	return ff, nil
}
