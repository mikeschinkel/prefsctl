package stdlibex

import (
	"encoding/json"
	"os"
)

func MarshalJSONFile(_ Context, file string, object any) error {
	content, err := json.MarshalIndent(object, "", "\t")
	if err != nil {
		goto end
	}
	err = os.WriteFile(file, content, os.ModePerm)
end:
	return err
}

func UnmarshalJSONFile(_ Context, filename string, object any) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		goto end
	}
	err = json.Unmarshal(content, &object)
	if err != nil {
		goto end
	}
end:
	return err
}
