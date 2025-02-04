package kvfilters

import (
	"errors"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/sliceconv"
	"github.com/mikeschinkel/prefsctl/stdlibex"
)

type QueryArgs struct {
	Filters        []Filter
	Groups         []Group
	Labels         []*Label
	KeepEmptyGroup bool
	OmitEmpty      bool
	OmitInvalid    bool
	OmitUnchanged  bool
}

func Query(args QueryArgs) (result []Group, err error) {
	var errs errutil.MultiErr
	var mustKeep, omit bool
	var newGroup Group
	var groupAdded bool
	var keyValues []KeyValue

	groups, err := QueryGroups(args)
	if err != nil {
		goto end
	}

	result = make([]Group, 0, len(groups))
	for _, g := range groups {
		keyValues = g.KeyValues()
		if len(keyValues) == 0 && args.KeepEmptyGroup {
			result = append(result, g.ShallowCopy())
			continue
		}
		groupAdded = false
		for _, kv := range keyValues {
			if args.Labels != nil && !stdlibex.SlicesIntersect(args.Labels, kv.Labels().labels) {
				continue
			}
			if args.OmitInvalid {
				if !g.Valid() {
					continue
				}
				if !kv.Valid() {
					continue
				}
			}
			if args.OmitEmpty && kv.Value() == "" {
				continue
			}
			mustKeep, err = MustKeepKeyValue2(args.Filters, kv, KeysValuesOrKeyValue...)
			if err != nil {
				errs.Add(errors.Join(err, kv.ErrorInfo()))
				continue
			}
			if !mustKeep {
				omit, err = OmitKeyValue2(args.Filters, kv, KeysValuesOrKeyValue...)
				if err != nil {
					errs.Add(errors.Join(err, kv.ErrorInfo()))
					continue
				}
				if omit {
					continue
				}
			}
			if groupAdded {
				print()
			} else {
				// Copy the group but w/o its key-values
				newGroup = g.ShallowCopy()
				// Add to our result
				result = append(result, newGroup)
				// Don't add another until group changes
				groupAdded = true
			}
			// Now add the key-value to the current new group
			//goland:noinspection GoDfaNilDereference
			newGroup.AddKeyValue(kv)
		}
	}
	// Eliminate any unused capacity
	result = sliceconv.Minimize(result)
end:
	return result, errs.Join(err)
}

func QueryGroups(args QueryArgs) (result []Group, err error) {
	var errs errutil.MultiErr
	var mustKeep, omit bool

	result = make([]Group, 0, len(args.Groups))
	for _, g := range args.Groups {
		mustKeep, err = MustKeepGroup(args.Filters, g.Name(), Groups)
		if err != nil {
			errs.Add(errors.Join(err, g.ErrorInfo()))
			continue
		}
		if !mustKeep {
			omit, err = OmitGroup(args.Filters, g.Name(), Groups)
			if err != nil {
				errs.Add(errors.Join(err, g.ErrorInfo()))
				continue
			}
			if omit {
				continue
			}
		}
		result = append(result, g)
	}
	// Eliminate any unused capacity
	result = sliceconv.Minimize(result)
	return result, errs.Join(err)
}
