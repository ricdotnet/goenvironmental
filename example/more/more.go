package more

import goenv "github.com/ricdotnet/goenvironmental"

func Hello() {
	println(goenv.Get("NUMBERS"))
}
