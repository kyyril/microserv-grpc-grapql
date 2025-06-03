package main

type Accpint struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Orders []Order `json:"orders"`
}