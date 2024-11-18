package main

import "C"
import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"prefsctl/macprefs"
)

func main() {
	domains, err := macprefs.GetPreferenceDomains()
	if err != nil {
		panic(err)
	}
	unsupported := make(map[string]struct{})
	for _, d := range domains {
		if !strings.HasPrefix(d.Name, "com.apple.") {
			continue
		}
		//fmt.Println(d.Name)
		prefs, err := d.Prefs()
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, pref := range prefs {
			if err := pref.Retrieve(); err != nil {
				if !errors.Is(err, macprefs.ErrUnsupportedType) {
					fmt.Printf("ERROR: %v [pref=%s/%s])\n",
						err,
						d.Name,
						pref.Key,
					)
					continue
				}
				unsupported[pref.Message()] = struct{}{}
				//fmt.Printf("WARNING: %s [pref=%s/%s]\n",
				//	pref.Message(),
				//	d.Name,
				//	pref.Key,
				//)
				continue
			}
			continue
			fmt.Printf("— %s:", pref.Key)
			// Check the type and get the value
			switch pref.Kind() {
			case reflect.Bool:
				fmt.Printf("%v\n", pref.Bool())
			case reflect.Int:
				fmt.Printf("%d\n", pref.Int())
			case reflect.String:
				fmt.Printf("%s\n", pref.String())
			default:
				fmt.Printf("%s; %s\n", pref.Message(), pref.Err())
			}

		}
	}
	for u := range unsupported {
		fmt.Println(u)
	}
}

//func main2() {
//	domain := "com.apple.dock"
//	key := "autohide"
//	//domain := "com.apple.finder"
//	//key := "ShowHardDrivesOnDesktop"
//
//	pref := NewPreference(domain, key)
//	if pref.Retrieve() != nil {
//		fmt.Println(pref.Err())
//	}
//	if err := pref.Retrieve(); err != nil {
//		log.Fatal(err)
//	}
//
//	// Check the type and get the value
//	switch pref.Kind() {
//	case reflect.Bool:
//
//		fmt.Printf("Boolean value: %v\n", pref.Bool())
//	case reflect.Int:
//
//		fmt.Printf("Integer value: %d\n", pref.Int())
//	case reflect.String:
//		fmt.Printf("String value: %s\n", pref.String())
//	default:
//
//	}
//
//	// Or get the raw string value regardless of type
//	fmt.Printf("Raw value: %s\n", pref.String())
//
//	// Check for errors at any time
//	if pref.Err() != nil {
//		log.Fatalf("An error occurred: %s; %v", pref.Message(), pref.Err())
//	}
//
//	switch pref.Kind() {
//	case reflect.Bool:
//		fmt.Printf("Boolean value: %v\n", pref.Bool())
//	case reflect.Int:
//		fmt.Printf("Integer value: %d\n", pref.Int())
//	case reflect.String:
//		fmt.Printf("String value: %s\n", pref.String())
//	case reflect.Invalid:
//		fallthrough
//	default:
//		fmt.Printf("Unexpected value: %s\n", pref.Message())
//	}
//
//}
