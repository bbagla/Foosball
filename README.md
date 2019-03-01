# go-foosball

A simple game of multiplayer foosball written in Go. Created using gorilla/websockets .

## Get Started
Clone this repository and run

```
go build main.go
```
and
```
go build client.go
```
to build the server and client respectively.

Run the server using
```
./main -host= your host
```
and the client using
```
./client -host= your host -channel=0/1
```

### Prerequisites

* gorilla/websocket
* veandco/go-sdl2

Install using

```
go get github.com/gorilla/websocket
```
And

```
go get -v github.com/veandco/go-sdl2/{sdl,img}
```


## Built Using

* [gorilla/websocket](https://www.github.com/gorilla/websocket) - for bidirectional communication using websockets
* [go-sdl2](https://www.github.com/veandco/go-sdl2) - for rendering and graphics



## Authors

* **Bhavya Bagla** -  [bbagla](https://github.com/bbagla)
* **Supraj Bachawala**  - [youwishyouhadthishandle](https://github.com/youwishyouhadthishandle)
* **Chander Sekhar** -  [Chander-Shekhar](https://github.com/Chander-Shekhar)
* **Raniya Hameed**  - [raniyahameed](https://github.com/raniyahameed)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details


