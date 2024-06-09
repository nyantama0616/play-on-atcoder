package validator

type IValidator interface {
	Validate(num int) (bool, error)
}
