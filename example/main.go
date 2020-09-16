package main

import (
	ulocal "arenainsight/util"
	"os"

	bnet "github.com/mrgreenturtle/bnetconnect/blizzardclient"
)

const tokenURL = "https://us.battle.net/oauth/token"

func main() {

	client := bnet.BnetClient{Client: bnet.CreateClient(os.Getenv("CID"), os.Getenv("CSID"), tokenURL)}
	// tophundred, topprofiles := ulocal.GetTopNumberArenaSpec(&client, 29, "3v3", "priest", "shadow", 10)
	// for i, entry := range tophundred.Entries {
	// 	fmt.Println(entry.Character.Name, entry.Rating, topprofiles[i].Equipment)
	// }
	ulocal.GetAllNumberArenaCasters(&client, 29, "3v3", 3)

}
