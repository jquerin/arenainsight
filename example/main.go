package main

import (
	ulocal "arenainsight/util"
	"fmt"
	bnet "github.com/mrgreenturtle/bnetconnect/blizzardclient"
	"os"
)

const tokenURL = "https://us.battle.net/oauth/token"

func main() {
	myArray := [100] int
	client := bnet.BnetClient{Client: bnet.CreateClient(os.Getenv("CID"), os.Getenv("CSID"), tokenURL)}
	tophundred, topprofiles := ulocal.GetTopHundredArenaSpec(&client, 29, "3v3", "paladin", "holy")
	for i, entry := range tophundred.Entries {
		fmt.Println(entry.Character.Name, entry.Rating, topprofiles[i].)
	}

}
