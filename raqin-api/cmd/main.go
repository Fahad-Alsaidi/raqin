/*

هنا نقطة البدأ للبرنامج بأكمله

*/

package main

import (
	rest "raqin-api/api"
	"raqin-api/storage"
)

func main() {

	err := storage.Connect()
	if err != nil {
		panic(err)
	}
	rest.Init()
}
