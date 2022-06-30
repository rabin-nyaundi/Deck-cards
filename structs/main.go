package main

import "fmt"

type ContactInfo struct {
	Email       string
	PhoneNumber string
	Address     string
}

type Person struct {
	FirstName string
	LastName  string
	Age       int64
	Contact   ContactInfo
}

func main() {
	p := Person{
		FirstName: "test",
		LastName:  "user",
		Age:       25,
		Contact: ContactInfo{
			Email:       "test@gmail.com",
			PhoneNumber: "0700304999",
			Address:     "00100 Nairobi",
		},
	}

	p.print()

	fmt.Println("\n\n*************\t")
	/**
	* second method
	 */

	var dev Person
	dev.Age = 12
	dev.FirstName = "Developer"
	dev.LastName = "Fullstack"
	dev.Contact = ContactInfo{
		Email:       "test@gmail.com",
		PhoneNumber: "0700304999",
		Address:     "00100 Nairobi",
	}

	dev.updateContactInfo(ContactInfo{
		Email:       "test@gmail.com",
		PhoneNumber: "08373",
		Address:     "shjgshdwh",
	})

	dev.print()

}

func (p *Person) updateContactInfo(newContactInfo ContactInfo) {
	p.Contact = newContactInfo
}

func (p Person) print() {
	fmt.Printf("\tPerson \n %+v\n\n", p)

	fmt.Printf("\n\tDev personal info \n %+v", p)
}
