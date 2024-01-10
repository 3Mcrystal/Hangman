package hangman

func ContainTable(s []string, car string) bool {
	for _, c := range s {
		if c == car {
			return true
		}
	}
	return false
}

func ContainString(s string, car string) bool {
	for _, c := range s {
		if string(c) == car {
			return true
		}
	}
	return false
}
