package main

type person struct{
	firstName string
	lastName string
}

func populatePerson() person{
	p := person{firstName: "Ang",lastName: "Bhatt"}
	return p

}
func (p *person) updatePerson(firstName string){
    (*p).firstName=firstName
}

func updatePersonWihoutPointer(firstname string,p person) person{
   p.firstName=firstname
   return p
}

