package main

func getIndex(id int) (int, bool) {
	for i, cert := range certificates {
		if cert.Id == id {
			return i, true
		}
	}
	return -1, false
}
