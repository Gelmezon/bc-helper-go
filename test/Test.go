package main

import (
	"awesomeProject/model"
	"awesomeProject/serializer"
	"awesomeProject/service"
	"fmt"
)

func main() {
	model.DBInit()
}

func testCreate() {
	wear := model.BcWear{UpdateMemberNumber: "test", WearName: "sasasas"}
	var bcWearService service.BcWearService
	bcWearService.Create(&wear)
}

func testModify() {
	wear := model.BcWear{UpdateMemberNumber: "test", WearName: "vvvvvv"}
	wear.ID = 1
	var bcWearService service.BcWearService
	bcWearService.Modify(&wear)
}

func testRemove() {
	var bcWearService service.BcWearService
	bcWearService.Remove(1)
}

func testOneById() {
	var bcWearService service.BcWearService
	id := bcWearService.OneById(2)
	fmt.Println(id.WearName)
}

func testList() {
	var bcWearService service.BcWearService
	wears, _ := bcWearService.List(&model.BcWear{UpdateMemberNumber: "test"})

	for _, wear := range wears {
		fmt.Println(wear.WearName)
	}
}

func testPate() {
	var bcWearService service.BcWearService
	wears, count := bcWearService.Page(serializer.CreatePage(1, 10), &model.BcWear{UpdateMemberNumber: "test"})
	for _, wear := range wears {
		fmt.Println(wear.WearName)
	}
	fmt.Println(count)
}
