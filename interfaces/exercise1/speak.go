package main

func main() {
	var sp speaker

	var e english
	sp = e
	sp.speak()

	var c chinese
	sp = c
	sp.speak()

	sayHello(new(english))
	sayHello(&chinese{})
}

func sayHello(sp speaker) {
	sp.speak()
}
