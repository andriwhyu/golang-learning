package main

import "wire-providerset/wire"

func main() {
	msg := wire.InitializeDependency()
	msg.Print("Budi")
}
