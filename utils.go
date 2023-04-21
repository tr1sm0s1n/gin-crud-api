package main

import (
	"errors"
)

func getOne(id int) (*certificate, error) {
	for i, cert := range certificates {
		if cert.ID == id {
			return &certificates[i], nil
		}
	}
	return nil, errors.New("Invalid credential")
}

func updateOne(id int, updated certificate) (*certificate, error) {
	for i, cert := range certificates {
		if cert.ID == id {
			certificates[i] = updated
			return &certificates[i], nil
		}
	}
	return nil, errors.New("Invalid credential")
}

func getIndex(id int) (*int, error) {
	for i, cert := range certificates {
		if cert.ID == id {
			return &i, nil
		}
	}
	return nil, errors.New("Invalid credential")
}
