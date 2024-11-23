package macprefs

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mikeschinkel/prefsctl/cliutil"
	"github.com/mikeschinkel/prefsctl/errutil"
	"gopkg.in/yaml.v3"
)

type GetArgs struct {
	Filename Filename
}

func Get(ctx Context, args GetArgs) (result cliutil.Result, err error) {
	var errs errutil.MultiErr
	var output []byte
	var domains []string
	var content []byte

	prefs := YAMLFile{}
	filePath := "prefs.yaml"

	// check if file already exists
	if _, err := os.Stat(filePath); err == nil {
		errs.Add(fmt.Errorf("File %s already exist.\n", filePath))
		goto end
	}

	// run 'defaults domains' command
	fmt.Printf("Reading defaults domains\n")
	output, err = exec.Command("defaults", "domains").Output()
	if err != nil {
		fmt.Println()
		errs.Add(fmt.Errorf("error getting domains: %w", err))
		goto end
	}

	// split domains
	domains = strings.Split(string(output), ", ")
	for _, domain := range domains {
		fmt.Printf("Running `defaults read %s`\n", domain)
		// run 'defaults read <domain>' command
		propertiesOutput, err := exec.Command("defaults", "read", domain).CombinedOutput()
		if err != nil {
			if bytes.Contains(propertiesOutput, []byte(" does not exist")) {
				continue
			}
			err = fmt.Errorf("%s: %s; %w", domain, string(propertiesOutput), err)
			fmt.Println(err)
			errs.Add(err)
			continue
		}

		// parse properties
		properties := strings.Split(string(propertiesOutput), "\n")
		propertyMap := make(map[string]string)
		for _, property := range properties {
			// split into key-value pair
			parts := strings.SplitN(property, " = ", 2)
			if len(parts) < 2 {
				continue
			}
			parts[0] = strings.TrimSpace(parts[0])
			propertyMap[parts[0]] = parts[1]
		}
		prefs.Preferences = append(prefs.Preferences, YAMLDomain{
			Name:       domain,
			Properties: propertyMap,
		})
	}

	content, err = yaml.Marshal(&prefs)
	if err != nil {
		errs.Add(fmt.Errorf("cannot marshal prefs; %w", err))
		goto end
	}

	err = os.WriteFile(filePath, content, 0644)
	if err != nil {
		errs.Add(fmt.Errorf("error writing file '%s'; %w", filePath, err))
		goto end
	}

	fmt.Printf("File '%s' successfully created.\n", filePath)
end:
	return result, errs.Err()
}
