package helper

func PanicIfError(err error) {
	if err != nil {
		//LogError(err.Error(), err)
		panic(err)
	}
}
