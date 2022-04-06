package main

import "encoding/xml"

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

type User struct {
	XMLName xml.Name `xml:"user"`
	Id      int      `xml:"id"`
	Name    string   `xml:"name"`
	Phone   string   `xml:"phone"`
	Address string   `xml:"address"`
	Email   string   `xml:"email"`
}
