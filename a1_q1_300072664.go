//student name: Zhuoran Wang
//student number:300072664

package main

import (
	"fmt"
	"time"
)

// global
var prime = NewCategory("prime", 35)
var standard = NewCategory("standard", 25)
var special = NewCategory("special", 15)

type Play struct {
	name      string
	purchased []Ticket //
	showStart time.Time
	showEnd   time.Time
}

type Comedy struct {
	laughs float32
	deaths int32
	Play
}

func NewComedy() *Comedy {
	return &Comedy{0.2, 0, Play{"Tartuffe", make([]Ticket, 0),
		time.Date(2020, 3, 3, 16, 0, 0, 0, time.UTC),
		time.Date(2020, 3, 3, 17, 20, 0, 0, time.Local)}}
}

type Tragedy struct {
	laughs float32
	deaths int32
	Play
}

func NewTragedy() *Tragedy {
	return &Tragedy{0.0, 12, Play{"Macbeth", make([]Ticket, 0),
		time.Date(2020, 4, 16, 11, 30, 0, 0, time.UTC),
		time.Date(2020, 4, 16, 12, 30, 0, 0, time.Local)}}
}

type Show interface {
	getName() string
	getShowStart() time.Time
	getShowEnd() time.Time
	addPurchase(*Ticket) bool
	isNotPurchased(*Ticket) bool
}

type Seat struct {
	number int32
	row    int32
	cat    *Category
}

type Category struct {
	name string
	base float32
}

func NewCategory(name string, base float32) *Category {
	return &Category{name, base}
}

type Ticket struct {
	customer string
	s        *Seat
	show     *Show
}

type Theatre struct {
	seats []Seat
	shows []Show
}

func (c *Comedy) getName() string {
	return c.Play.name
}
func (t *Tragedy) getName() string {
	return t.Play.name
}
func (c *Comedy) getShowStart() time.Time {
	return c.Play.showStart
}
func (t *Tragedy) getShowStart() time.Time {
	return t.Play.showStart
}
func (c *Comedy) getShowEnd() time.Time {
	return c.Play.showEnd
}
func (t *Tragedy) getShowEnd() time.Time {
	return t.Play.showEnd
}
func (c *Comedy) addPurchase(t *Ticket) bool {
	for _, value := range c.Play.purchased {
		temp := *t
		if temp.s.row == value.s.row && temp.s.number == value.s.number {
			return false
		}
	}
	c.Play.purchased = append(c.Play.purchased, *t)
	return true
}
func (tr *Tragedy) addPurchase(t *Ticket) bool {
	for _, value := range tr.Play.purchased {
		temp := *t
		if temp.s.row == value.s.row && temp.s.number == value.s.number {
			return false
		}
	}
	tr.Play.purchased = append(tr.Play.purchased, *t)
	return true
}
func (c *Comedy) isNotPurchased(t *Ticket) bool {
	for _, value := range c.Play.purchased {
		temp := *t
		if temp.s.row == value.s.row && temp.s.number == value.s.number {
			return false
		}
	}
	return true
}
func (tr *Tragedy) isNotPurchased(t *Ticket) bool {
	for _, value := range tr.Play.purchased {
		temp := *t
		if temp.s.row == value.s.row && temp.s.number == value.s.number {
			return false
		}
	}
	return true
}

func NewSeat(seatNumber int32, row int32, cate *Category) *Seat {
	return &Seat{seatNumber, row, cate}
}

func NewTicket(name string, seat *Seat, show *Show) *Ticket {
	return &Ticket{name, seat, show}
}

func NewTheatre(n int32, shows []Show) *Theatre {
	return &Theatre{make([]Seat, n), shows}
}

func findPrimeSeat(show Show) (*Seat, bool) {
	for p := 1; p < 6; p++ {
		ticket := &Ticket{"agent", NewSeat(int32(p), 1, prime), &show}
		if show.isNotPurchased(ticket) {
			return NewSeat(int32(p), 1, prime), true
		}
	}
	return nil, false
}

func findStandardSeat(show Show) (*Seat, bool) {
	for r := 2; r <= 4; r++ {
		for p := 1; p < 6; p++ {
			ticket := &Ticket{"agent", NewSeat(int32(p), int32(r), standard), &show}
			if show.isNotPurchased(ticket) {
				return NewSeat(int32(p), int32(r), standard), true
			}
		}
	}
	return nil, false
}

func findSpecialSeat(show Show) (*Seat, bool) {
	for p := 1; p < 6; p++ {
		ticket := &Ticket{"agent", NewSeat(int32(p), 5, special), &show}
		if show.isNotPurchased(ticket) {
			return NewSeat(int32(p), 5, special), true
		}
	}
	return nil, false
}
func acceptAnotherSeat(answer string) bool {
	if answer == "Y" || answer == "y" || answer == "yes" {
		return true
	} else {
		return false
	}
}
func printPrice(cat *Category) {
	fmt.Println("This is a ", cat.name, " saet and the price is : $", cat.base)
}
func printSeat(seat *Seat) {
	fmt.Println("we've found a new seat for you, it is, [", seat.row, seat.number, "]")

}

func offer(customerName string, row int, show Show) (*Ticket, bool) {

	fmt.Println("Sorry, the seat that you selected sold out, but we find another seat for you, is that ok?")
	answer := ""
	fmt.Println("Please type Y[yes] or N[no]")
	fmt.Scanf("%s", &answer)
	if acceptAnotherSeat(answer) == false {
		fmt.Println("Sorry, we cannot help you")
		return nil, false
	}
	var seat *Seat
	var ok bool

	if row == 1 {
		seat, ok = findPrimeSeat(show)
		if !ok {
			seat, ok = findStandardSeat(show)
		}
		if !ok {
			seat, ok = findSpecialSeat(show)
		}
	} else if row == 5 {
		seat, ok = findSpecialSeat(show)
		if !ok {
			seat, ok = findPrimeSeat(show)
		}
		if !ok {
			seat, ok = findStandardSeat(show)
		}
	} else {
		seat, ok = findStandardSeat(show)
		if !ok {
			seat, ok = findPrimeSeat(show)
		}
		if !ok {
			seat, ok = findSpecialSeat(show)
		}
	}
	if ok {
		printSeat(seat)
		printPrice(seat.cat)
		fmt.Println("do you want to buy it?")
		answer = ""
		fmt.Println("Please type Y[yes] or N[no]")
		fmt.Scanf("%s", &answer)
		if acceptAnotherSeat(answer) == false {
			fmt.Println("Sorry, we cannot help you")
			return nil, false
		}
		return NewTicket(customerName, seat, &show), true
	}
	return nil, false

}

func main() {
	// set up two movie
	var comedy = NewComedy()
	comedy.Play.showStart = time.Date(2020, 3, 3, 19, 30, 0, 0, time.UTC)
	comedy.Play.showEnd = time.Date(2020, 3, 3, 22, 0, 0, 0, time.UTC)

	var tragedy = NewTragedy()
	tragedy.Play.showStart = time.Date(2020, 4, 10, 20, 0, 0, 0, time.UTC)
	tragedy.Play.showEnd = time.Date(2020, 4, 10, 23, 0, 0, 0, time.UTC)

	//ask getName
	var customerName = ""
	fmt.Println("Dear customer, can I have your name: ")
	fmt.Scanf("%s", &customerName)

	for {

		fmt.Println("do you want to sell tickets ? yes/no")
		//ask play
		result := ""
		fmt.Scanf("%s", &result)

		for {
			if result != "yes" && result != "no" {
				fmt.Println("Invalid! Try again!")
				fmt.Println("do you want to sell tickets ? yes/no")
				fmt.Scanf("%s", &result)
			} else {
				break
			}
		}

		if result == "no" {
			break
		}

		//ask play
		playName := ""
		fmt.Println("Welcome! Which play do you want to watch: ")
		fmt.Println("Now we have Comedy and tragedy")
		fmt.Scanf("%s", &playName)

		for {
			if playName != "Comedy" && playName != "tragedy" && playName != "comedy" && playName != "Tragedy" {
				fmt.Println("Invalid! Try again!")
				fmt.Println("Now we have Comedy and tragedy")
				fmt.Scanf("%s", &playName)
			} else {
				break
			}
		}
		var selected_show Show = nil
		if playName == "Comedy" || playName == "comedy" {
			selected_show = comedy
		}
		if playName == "Tragedy" || playName == "tragedy" {
			selected_show = tragedy
		}
		fmt.Println("Please select your seat:")
		fmt.Println("Note: the range of rows between 1 to 5")
		var seatRow int
		var seatcol int
		fmt.Printf("row: ")
		fmt.Scanf("%d", &seatRow)
		fmt.Printf("seat number: ")
		fmt.Scanf("%d", &seatcol)

		for {
			if seatRow < 1 || seatRow > 5 || seatcol < 1 || seatcol > 5 {
				fmt.Println("Please select your seat:")
				fmt.Println("Note: the range of rows between 1 to 5")
				fmt.Printf("row: ")
				fmt.Scanf("%d", &seatRow)
				fmt.Printf("seat number: ")
				fmt.Scanf("%d", &seatcol)
			} else {
				break
			}
		}
		var selected_cat *Category
		if seatRow == 1 {
			selected_cat = prime
		} else if seatRow == 5 {
			selected_cat = special
		} else {
			selected_cat = standard
		}

		var accepeted bool = false
		var needTicket = NewTicket(customerName, NewSeat(int32(seatcol), int32(seatRow), selected_cat), &selected_show)
		if selected_show.isNotPurchased(needTicket) {
			selected_show.addPurchase(needTicket)
			fmt.Println("You were select ", playName, ", your seat is [", seatRow, seatcol, "]")
			printPrice(selected_cat)
			fmt.Println("purchased successful")
		} else {
			fmt.Println("the seat your chose has been sold")
			needTicket, accepeted = offer(customerName, seatRow, selected_show)
			if accepeted {
				selected_show.addPurchase(needTicket)
			}
		}

	}

}
