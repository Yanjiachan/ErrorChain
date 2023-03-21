package errorChain

type ErrorChain struct {
	Error error
}

func (ec *ErrorChain) CheckError(fs ...func(*ErrorChain) error) *ErrorChain {
	for _, f := range fs {
		ec.Error = f(ec)
		if ec.Error != nil {
			break
		}
	}
	return ec
}
