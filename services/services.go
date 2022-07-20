package services

func ValidateAccountNumber(accountNos string) bool {
	if len(accountNos) != 10 {
		return false
	}
	return true
}
