package playlistcreator

import (
	"fmt"
)

type Track struct {
	artist string
	name string
	path string
}

func PrintTrack() {
	t := Track{"Hooded Fang", "Wasteland", ".."}
	fmt.Println("Track is: ", t)
}
