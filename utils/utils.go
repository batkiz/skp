// utils provides functions which are independent of skp
package utils

import "log"

// HandleErr is a helper to decrease annoying `if err != nil {}`s
func HandleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}