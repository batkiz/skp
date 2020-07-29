package commands

import (
	"fmt"
	"github.com/batkiz/skp/lib"
	. "github.com/batkiz/skp/utils"
	"regexp"
	"strings"
)

// Info is the base func for `scoop info <app>` command
func Info(path string) {
	name := ExtractName(path)

	fmt.Println("Name:", name)

	app := lib.ReadManifest(path)

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
			licenses := strings.FieldsFunc(*appLicense.String, func(r rune) bool {
				return r == '|' || r == ','
			})

			var licensesUrl []string
			for _, lic := range licenses {
				licenseUrl := fmt.Sprintf("https://spdx.org/licenses/%v.html", lic)
				licensesUrl = append(licensesUrl, licenseUrl)
			}

			license = fmt.Sprintf("%v (%v)", *appLicense.String, strings.Join(licensesUrl[:], ","))
		} else {
			license = fmt.Sprintf("%v (%v)",
				*appLicense.String,
				fmt.Sprintf("https://spdx.org/licenses/%v.html", *appLicense.String))
		}

		fmt.Println("License:", license)
	}

	if app.Bin != nil {
		fmt.Println("Binaries:")
		if app.Bin.String != nil {
			fmt.Printf("  %v\n", *app.Bin.String)
		} else if len(app.Bin.UnionArray) != 0 {
			var binaries []string

			for _, union := range app.Bin.UnionArray {
				binaries = append(binaries, *union.String)
				for _, s := range union.StringArray {
					binaries = append(binaries,s)
				}
			}

			s := "  "
			for _, binary := range binaries {
				s += fmt.Sprintf("%v ", binary)
			}
			fmt.Println(s)
		}
	}


}

// ExtractName extracts app name from a path or url
func ExtractName(path string) string {
	r := regexp.MustCompile(`.*[/\\](.*?)\.json`)

	name := r.FindStringSubmatch(path)

	return name[1]
}
