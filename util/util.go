package arenainsight

import (
	"fmt"
	"strings"
	"sync"

	arena "github.com/mrgreenturtle/bnetconnect/arenaleaderboard"
	bnet "github.com/mrgreenturtle/bnetconnect/blizzardclient"
	cp "github.com/mrgreenturtle/bnetconnect/characterprofile"
)

var wg sync.WaitGroup

func GetArenaEntries(b *bnet.BnetClient, seasonID int, bracket string) arena.ArenaLeaderBoard {

	returnArenaEntriesOnly := arena.ArenaLeaderBoard{}
	temp := b.GenerateArenaLeaderBoard(seasonID, bracket)
	returnArenaEntriesOnly.Entries = temp.Entries
	return returnArenaEntriesOnly

}

func GetTopNumberArenaSpec(b *bnet.BnetClient, seasonID int, bracket string, class string, spec string, amount int) (arena.ArenaLeaderBoard, []cp.CharacterProfile) {

	topClass := arena.ArenaLeaderBoard{}
	topProfile := []cp.CharacterProfile{}
	tempAlb := GetArenaEntries(b, seasonID, bracket)
	tempAlb.Entries = tempAlb.Entries[0:1000]
	for _, entry := range tempAlb.Entries {
		name := strings.ToLower(entry.Character.Name)
		realm := strings.ToLower(entry.Character.Realm.Slug)
		tempCp := b.GenerateCharacterProfile(name, realm)
		if strings.ToLower(tempCp.CharacterClass.Name) == class && strings.ToLower(tempCp.ActiveSpec.Name) == spec {
			topClass.Entries = append(topClass.Entries, entry)
			topProfile = append(topProfile, tempCp)
			if len(topProfile) >= amount {
				break
			}
		}
	}
	return topClass, topProfile
}

func helperChannel(b *bnet.BnetClient, seasonID int, bracket string, class string, spec string, amount int, c chan []cp.CharacterProfile) {
	defer wg.Done()
	topProfile := []cp.CharacterProfile{}
	tempAlb := GetArenaEntries(b, seasonID, bracket)
	tempAlb.Entries = tempAlb.Entries[0:1000]
	for _, entry := range tempAlb.Entries {
		name := strings.ToLower(entry.Character.Name)
		realm := strings.ToLower(entry.Character.Realm.Slug)
		tempCp := b.GenerateCharacterProfile(name, realm)
		if strings.ToLower(tempCp.CharacterClass.Name) == class && strings.ToLower(tempCp.ActiveSpec.Name) == spec {
			topProfile = append(topProfile, tempCp)
			if len(topProfile) >= amount {
				break
			}
		}
	}
	c <- topProfile
}

func GetAllNumberArenaCasters(b *bnet.BnetClient, seasonID int, bracket string, amount int) {
	wg.Add(9)
	chashadow := make(chan []cp.CharacterProfile)
	chaarcane := make(chan []cp.CharacterProfile)
	chafire := make(chan []cp.CharacterProfile)
	chafrost := make(chan []cp.CharacterProfile)
	chamoonkin := make(chan []cp.CharacterProfile)
	chadestro := make(chan []cp.CharacterProfile)
	chaaff := make(chan []cp.CharacterProfile)
	chademo := make(chan []cp.CharacterProfile)
	chaele := make(chan []cp.CharacterProfile)

	go helperChannel(b, seasonID, bracket, "priest", "shadow", amount, chashadow)
	go helperChannel(b, seasonID, bracket, "mage", "arcane", amount, chaarcane)
	go helperChannel(b, seasonID, bracket, "mage", "fire", amount, chafire)
	go helperChannel(b, seasonID, bracket, "mage", "frost", amount, chafrost)
	go helperChannel(b, seasonID, bracket, "druid", "balance", amount, chamoonkin)
	go helperChannel(b, seasonID, bracket, "shaman", "elemental", amount, chaele)
	go helperChannel(b, seasonID, bracket, "warlock", "affliction", amount, chaaff)
	go helperChannel(b, seasonID, bracket, "warlock", "demonology", amount, chademo)
	go helperChannel(b, seasonID, bracket, "warlock", "destruction", amount, chadestro)
	shadow := <-chashadow
	arcane := <-chaarcane
	fire := <-chafire
	frost := <-chafrost
	moonkin := <-chamoonkin
	ele := <-chaele
	aff := <-chaaff
	demo := <-chademo
	destro := <-chadestro
	wg.Wait()
	fmt.Printf("%s \t %s\n", shadow[0].ActiveSpec.Name, shadow[0].Name)
	fmt.Printf("%s \t %s\n", arcane[0].ActiveSpec.Name, arcane[0].Name)
	fmt.Printf("%s \t %s\n", fire[0].ActiveSpec.Name, fire[0].Name)
	fmt.Printf("%s \t %s\n", frost[0].ActiveSpec.Name, frost[0].Name)
	fmt.Printf("%s \t %s\n", moonkin[0].ActiveSpec.Name, moonkin[0].Name)
	fmt.Printf("%s \t %s\n", ele[0].ActiveSpec.Name, ele[0].Name)
	fmt.Printf("%s \t %s\n", aff[0].ActiveSpec.Name, aff[0].Name)
	fmt.Printf("%s \t %s\n", demo[0].ActiveSpec.Name, demo[0].Name)
	fmt.Printf("%s \t %s\n", destro[0].ActiveSpec.Name, destro[0].Name)

}
