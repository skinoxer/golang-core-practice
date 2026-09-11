package main

import (
	"fmt"
	"strconv"
	// "strconv"
)

type Room struct {
	Id       int
	Type     string
	BedCount int
	Price    int
	Status   bool
}

func main() {
	RoomList := generateRooms()

	input := ""
	for input != "exit" {
		fmt.Println("enter command  ")
		fmt.Println("1:Room Lists ")
		fmt.Println("2:Add Room ")
		fmt.Scanln(&input)
		switch input {
		case "1":
			getRoomList(RoomList)
		case "2":
			addRoom(&RoomList)
		case "exit":
			fmt.Println("Exitning...")
			break
		default:
			println("invalid income")

		}
	}
}
func getRoomList(RoomList []Room) {
	for _, item := range RoomList {
		fmt.Printf("Room List : %v\n", item)

	}
}
func addRoom(RoomList *[]Room) {

	fmt.Println("what type or room do you want (single, double,vip)")

	Type := ""

	fmt.Scanln(&Type)

	if Type != "single" && Type != "double" && Type != "vip" {
		println("please enter single double or vip")
		return
	}

	fmt.Println("how many bed do you want (1,2,3)")
	input2 := ""

	fmt.Scanln(&input2)
	numberOfBed, err := strconv.Atoi(input2)

	if err != nil || numberOfBed <= 0 {
		fmt.Println("you need type number")
		return

	}
	lastId := 0
	for i := 0; i < len(*RoomList); i++ {
		lastId = (*RoomList)[i].Id
	}
	room := Room{Status: false}

	room.Id = lastId + 1
	room.Type = Type
	room.BedCount = numberOfBed
	room.Price = numberOfBed * 200
	*RoomList = append(*RoomList, room)

}

func generateRooms() []Room {
	rooms := make([]Room, 0, 20)
	rooms = append(rooms, Room{Id: 1, Type: "single", BedCount: 1, Price: 200, Status: true})
	rooms = append(rooms, Room{Id: 2, Type: "single", BedCount: 1, Price: 200, Status: true})
	rooms = append(rooms, Room{Id: 3, Type: "double", BedCount: 2, Price: 400, Status: true})
	rooms = append(rooms, Room{Id: 4, Type: "vip", BedCount: 2, Price: 600, Status: true})
	rooms = append(rooms, Room{Id: 5, Type: "single", BedCount: 1, Price: 200, Status: true})
	return rooms

}
