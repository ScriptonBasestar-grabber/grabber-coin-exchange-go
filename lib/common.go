package lib

func Err(e error) {
	if e != nil {
		panic(e)
	}
}
