package main

type person struct{
	firstName string
	lastName string
	address address
}

type address struct{
	addressLine string
	zipcode string
}

func populatePerson() person{
	p := person{firstName: "Ang",lastName: "Bhatt",address: address{addressLine: "848 parkland place",zipcode: "23059"}}
	return p

}
func (p *person) updatePerson(firstName string){
    (*p).firstName=firstName
}

func updatePersonWihoutPointer(firstname string,p person) person{
   p.firstName=firstname
   return p
}

