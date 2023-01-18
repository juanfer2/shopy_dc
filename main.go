package main

import (
	//"github.com/juanfer2/shopy_dc_golang/cmd"

	"github.com/juanfer2/shopy_dc_golang/cmd"
	"github.com/juanfer2/shopy_dc_golang/src/shared/infrastructure/config"
	//"github.com/juanfer2/shopy_dc_golang/src/shared/infrastructure/config"
)

// type User struct {
// 	Id       int    `validate:"number,min=1,max=1000"`
// 	Name     string `validate:"string,min=2,max=10"`
// 	Bio      string `validate:"string"`
// 	Email    string `validate:"email"`
// 	Username string `validate:"regex,expression=[0-9]"`
// }

// user := User{
// 	Id:       0,
// 	Name:     "superlongstring",
// 	Bio:      "",
// 	Email:    "foobar",
// 	Username: "sdfsdf",
// }
// fmt.Println("Errors:")
// for i, err := range utilities.ValidatorStruct(user) {
// 	fmt.Printf("\t%d. %s\n", i+1, err.Error())
// }

func main() {
	config.LoadEnv()
	cmd.Execute()
}
