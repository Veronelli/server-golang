package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandlerRoot)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("POST", "/api", server.addMiddleware(HandleHome, ChechAuth(), Loggin()))
	server.Listen()
}
