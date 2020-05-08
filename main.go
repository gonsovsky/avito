package main

import (
	"newOne/grabber"
)

func main() {
	var page, _ = grabber.Grab("https://www.avito.ru/moskva/noutbuki/hp_dv6_metall_i7_2.3ghzx88gb750radeon_1901197465")
	page.Dump()
}
