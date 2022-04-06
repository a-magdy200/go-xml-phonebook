package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
		//// Views Layout is the global layout for all template render until override on Render function.
		//ViewsLayout: "/",
	})
	app.Static("/static", "./static")
	app.Get("/new", func(c *fiber.Ctx) error {
		return c.Render("new", "")
	})
	app.Post("/new", func(c *fiber.Ctx) error {
		xmlFile, err := os.Open("users.xml")
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Successfully Opened users.xml")
		// defer the closing of our xmlFile so that we can parse it later on
		defer func(xmlFile *os.File) {
			err := xmlFile.Close()
			if err != nil {

			}
		}(xmlFile)

		// read our opened xmlFile as a byte array.
		byteValue, _ := ioutil.ReadAll(xmlFile)

		// we initialize our Users array
		var users Users
		// we unmarshal our byteArray which contains our
		// xmlFiles content into 'users' which we defined above
		err = xml.Unmarshal(byteValue, &users)
		if err != nil {
			return err
		}
		var user User
		err = c.BodyParser(&user)
		if err != nil {
			return err
		}
		if len(users.Users) > 0 {
			user.Id = users.Users[len(users.Users)-1].Id + 1
		} else {
			user.Id = 1
		}
		users.Users = append(users.Users, user)
		file, err := xml.MarshalIndent(users, "", "\t")
		if err != nil {
			return err

		}
		_ = ioutil.WriteFile("users.xml", file, 777)
		fmt.Println(user)
		//return c.JSON(users)
		return c.Redirect("/")
	})
	app.Get("/", func(c *fiber.Ctx) error {
		// Open our xmlFile
		xmlFile, err := os.Open("users.xml")
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Successfully Opened users.xml")
		// defer the closing of our xmlFile so that we can parse it later on
		defer func(xmlFile *os.File) {
			err := xmlFile.Close()
			if err != nil {

			}
		}(xmlFile)

		// read our opened xmlFile as a byte array.
		byteValue, _ := ioutil.ReadAll(xmlFile)

		// we initialize our Users array
		var users Users
		// we unmarshal our byteArray which contains our
		// xmlFiles content into 'users' which we defined above
		err = xml.Unmarshal(byteValue, &users)
		if err != nil {
			return err
		}

		query := c.Query("query")

		if query != "" {
			var searchUsers Users
			var user User
			for i := 0; i < len(users.Users); i++ {
				if strconv.Itoa(users.Users[i].Id) == query || users.Users[i].Name == query || users.Users[i].Email == query || users.Users[i].Phone == query || users.Users[i].Address == query {
					user = users.Users[i]
					searchUsers.Users = append(searchUsers.Users, user)
				}
			}
			return c.Render("all", searchUsers)
		}
		return c.Render("all", users)
		//return c.SendString(string(jsonString))
	})
	app.Get("/delete", func(c *fiber.Ctx) error {
		// Open our xmlFile
		xmlFile, err := os.Open("users.xml")
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Successfully Opened users.xml")
		// defer the closing of our xmlFile so that we can parse it later on
		defer func(xmlFile *os.File) {
			err := xmlFile.Close()
			if err != nil {

			}
		}(xmlFile)

		// read our opened xmlFile as a byte array.
		byteValue, _ := ioutil.ReadAll(xmlFile)

		// we initialize our Users array
		var users Users
		// we unmarshal our byteArray which contains our
		// xmlFiles content into 'users' which we defined above
		err = xml.Unmarshal(byteValue, &users)
		if err != nil {
			return err
		}
		var user User
		var newUsers Users
		id := c.Query("id")
		for i := 0; i < len(users.Users); i++ {
			if strconv.Itoa(users.Users[i].Id) != id {
				user = users.Users[i]
				newUsers.Users = append(newUsers.Users, user)
			}
		}
		file, err := xml.MarshalIndent(newUsers, "", "\t")
		if err != nil {
			return err
		}
		_ = ioutil.WriteFile("users.xml", file, 777)
		return c.Redirect("/")
	})
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
