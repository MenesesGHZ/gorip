package fbrip

func includes(slice []string, v string) bool {
	for _, value := range slice {
		if value == v {
			return true
		}
	}
	return false
}
