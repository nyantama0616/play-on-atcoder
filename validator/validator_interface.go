package validator

type IValidator interface {
	// Sample{num}がACになるか検証する
	Validate(num int) (bool, error)
}
