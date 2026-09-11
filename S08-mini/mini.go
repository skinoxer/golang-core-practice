package main

import (
	"fmt"
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
	fmt.Printf("Room List : %v", RoomList)
	input := ""
	for input != "exit" {
		fmt.Println("enter command  ")
		fmt.Println("1:Room Lists ")
		fmt.Println("2:Add Room ")
		fmt.Println("3:Reserve Room ")
		fmt.Scan(&input)
		switch input {
		case "1":
			getRoomList(RoomList)
		case "2":
			addRoom()
		case "3":
			reserveRoom()
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
func addRoom() {
	fmt.Println("please select type (single,double,vip) and how many bed you want like->doublee 2")
}

func reserveRoom() {}

func calculateRomePrice() {}

func getRoomFromInput() {}

func generateRooms() []Room {
	rooms := make([]Room, 0, 20)
	rooms = append(rooms, Room{Id: 1, Type: "single", BedCount: 1, Price: 200, Status: true})
	rooms = append(rooms, Room{Id: 2, Type: "single", BedCount: 1, Price: 200, Status: true})
	rooms = append(rooms, Room{Id: 3, Type: "double", BedCount: 2, Price: 400, Status: true})
	rooms = append(rooms, Room{Id: 4, Type: "vip", BedCount: 2, Price: 600, Status: true})
	rooms = append(rooms, Room{Id: 5, Type: "single", BedCount: 1, Price: 200, Status: true})
	return rooms

}
