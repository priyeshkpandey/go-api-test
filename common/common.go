package common

func CheckNoError(err error) bool {
	if err != nil {
		return false
	}
	return true
}
