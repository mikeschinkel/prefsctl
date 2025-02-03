package yamlutil

import (
	"bytes"
	"errors"
	"strings"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/logargs"
	"gopkg.in/yaml.v3"
)

type Document string

func BuildDocument(object any) (yd Document, err error) {
	sb := strings.Builder{}
	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	err = enc.Encode(object)
	if err != nil {
		err = errors.Join(
			ErrFailedToYAMLEncodeObject,
			errutil.ErrArg(logargs.ObjectType, "%T", object),
			err,
		)
		goto end
	}
	err = enc.Close()
	if err != nil {
		err = errors.Join(ErrFailedToCloseYAMLEncoder, err)
		goto end
	}
	sb.WriteString("\n---\n")
	sb.WriteString(buf.String())
	sb.WriteByte('\n')
	yd = Document(sb.String())
end:
	return yd, err
}
