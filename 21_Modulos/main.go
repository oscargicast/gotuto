package main

import (
	"github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2"
	utils "github.com/oscargicast/gotuto/21_Modulos/utils"
)

func main() {
	hellomod.SayHello()
	hellomod2.SayHello("Oscar")
	utils.HelloWorld()
}
