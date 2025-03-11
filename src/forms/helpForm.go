package forms

func HelpForm() {
	println("Usage: clh [options] [type]")
	println("Options:")
	println("  -debug     Run debug util")
	println("  -help      Show help")
	println("  -t string  Type of util to run")
	println("Types:")
	println("  -tree      Run tree util")
	println("  -color     Run color util")
	println("  -time      Run time util")
}
