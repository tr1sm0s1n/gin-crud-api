package main

type certificate struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Course string `json:"course"`
	Grade  string `json:"grade"`
	Date   string `json:"date"`
}
