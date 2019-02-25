package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)





func main() {
	fmt.Println("it is")
	if err:=sdl.Init(sdl.INIT_EVERYTHING);err!=nil {
		fmt.Println(err)
		return
	}
	window,err:= sdl.CreateWindow("sdl2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(boxWidth), int32(boxHeight), sdl.WINDOW_OPENGL)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer window.Destroy()
	renderer,err := sdl.CreateRenderer(window,-1,sdl.RENDERER_ACCELERATED)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	var tableTex *sdl.Texture
	tableTex = drawBackground(tableTex,renderer)
	defer tableTex.Destroy()


	player1,err:= newplayer(renderer)
	if err!=nil{
		fmt.Println(err)
		return
	}
	for{
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent(){
			switch event.(type){
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.Copy(tableTex,nil,nil)
		//renderer.Copy(ballTex,&sdl.Rect{X: 0, Y:0 ,W:800 , H:600 },&sdl.Rect{X: 200, Y:200 ,W:10, H:20 })
		player1.draw(renderer)
		player1.update()
		renderer.Present()
		sdl.Delay(16)
	}
}




































