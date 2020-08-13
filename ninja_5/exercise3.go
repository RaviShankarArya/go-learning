package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	four_wheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	s := sedan{
		vehicle: vehicle{doors: 4, color: "black"},
		luxury:  true,
	}

	fmt.Println(s)

	fmt.Println(s.luxury)
	fmt.Println(s.vehicle.color)
	fmt.Println(s.vehicle.doors)

	t := truck{
		vehicle:    vehicle{doors: 4, color: "red"},
		four_wheel: false,
	}
	fmt.Println(t)
	fmt.Println(t.four_wheel)
	fmt.Println(t.vehicle.color)
	fmt.Println(t.vehicle.doors)

}
