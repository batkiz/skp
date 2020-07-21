package commands

import (
	"fmt"
	"github.com/batkiz/skp/lib"
	. "github.com/batkiz/skp/utils"
	"io/ioutil"
	"regexp"
	"strings"
)

func Info(path string) {
	json, err := ioutil.ReadFile(path)

	HandleErr(err)



	app, err := lib.UnmarshalManifest(json)
	HandleErr(err)

	fmt.Println("Homepage:", *app.Homepage)

	if app.Description != nil {
		fmt.Println("Description:", *app.Description)
	}

	fmt.Println("Version:", app.Version)

	if app.License != nil {
		appLicense := *app.License
		var license string

		if appLicense.LicenseClass != nil {
			license = fmt.Sprintf("%v (%v)", appLicense.LicenseClass.Identifier, *appLicense.LicenseClass.URL)
		} else if match, err := regexp.MatchString("^((ht)|f)tps?://.*?", *appLicense.String); match {
			HandleErr(err)
			license = *appLicense.String
		} else if match, err := regexp.MatchString("[|,]", *appLicense.String); match {
			HandleErr(err)
			lics := strings.FieldsFunc(*appLicense.String, func(r rune) bool {
				return r == '|' || r == ','
			})

			var licsUrl []string
			for _, lic := range lics {
				licurl := fmt.Sprintf("https://spdx.org/licenses/%v.html", lic)
				licsUrl = append(licsUrl, licurl)
			}

			license = fmt.Sprintf("%v (%v)", *appLicense.String, strings.Join(licsUrl[:], ","))
		} else {
			license = fmt.Sprintf("%v (%v)",
				*appLicense.String,
				fmt.Sprintf("https://spdx.org/licenses/%v.html", *appLicense.String))
		}

		fmt.Println("License:", license)
	}
}
