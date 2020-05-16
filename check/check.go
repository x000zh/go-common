package check

//Panic 如果有错误，panic
func Panic(e error) {
    if nil != e {
        panic(e)
    }
}
