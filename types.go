package main

type Certificate struct {
	Id     int    `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Course string `json:"course" binding:"required"`
	Grade  string `json:"grade" binding:"required"`
	Date   string `json:"date" binding:"required"`
}
