package main

func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}

// open terminal and run nc 127.0.0.1 8888
