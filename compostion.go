package main

import "fmt"

// Go does support composition , which is more of embedding rather than inheritance, which similar to OOP languages.

type User struct{
	Id int 
	Name string
	Location string

}

type Player struct {
	Id int
	Name string
	Location string
	GameId int
}

// simplifying by composing the user and player struct.
type User2 struct{
	Id int 
	Name, Location string
}

type Player2 struct{
	*User2
	GameId int
}

// since our struct is composed of another struct, whatever method that is on the User2 struct will be available on the Player2 struct.

func (u *User2) Greetings() string{
	return fmt.Sprintf("Hello guys my name is %s and i'm from %s",u.Name,u.Location)
}

func NewPlayer(id int, name,location string,gameId int) *Player2{
	return &Player2{
		User2: &User2{id,name,location},
		GameId: gameId,
	}
}


func main(){

	p := Player{}
	p.Id = 001
	p.Name = "oldman in squid game"
	p.Location = "south korea"
	p.GameId = 1
	fmt.Printf("%+v",p)

	// initialize a variable of type player2.

	p2 := NewPlayer(005,"Tommy Bravo", "NYC",7)
	// p2.Id = 002
	// p2.Name = "oscar"
	// p2.Location = "virginia"
	// p2.GameId = 2
	// fmt.Printf("%+v",p2)

	fmt.Println(p2.Greetings())

	// another option is by using struct literals.
	p3 := Player2{
		&User2{
			Id: 003,
			Name: "childish gambino",
			Location: "Atlanta",	
		},
		3,
	}

	fmt.Printf("Id: %d, Name:%s, Location: %s, GameId : %d\n",
				p3.Id,p3.Name,p3.Location,p3.GameId)

		// this will set a field directly on a player struct.
		p3.Id = 11
		fmt.Printf("%+v",p3)

}