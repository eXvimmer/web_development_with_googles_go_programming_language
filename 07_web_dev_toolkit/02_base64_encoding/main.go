package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// God bless Jonathan Blow
	s := "You want to make a 2D game, I don't give a shit WHAT you USE to make your 2D game. STOP ASKING. JUST GO DO SOMETHING. Asking this kind of questions is a STALLING TACTIC wherein you're allowed to think of yourself as like a game developer person, without actually fucking doing anything. That's the function of asking that question; JUST GO DO SOMETHING. It doesn't matter. If you pick one and it's the wrong one, you can switch to the other one and then you'll have a better idea of why the other one was the wrong one and you'll have more experience and that'll be great."

	s64 := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println("encoded data: ", s64)

	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		panic("decoding failed")
	}
	fmt.Println("decoded data: ", string(bs))
}
