package arenainsight

import (
	"fmt"
	"strings"

	arena "github.com/mrgreenturtle/bnetconnect/arenaleaderboard"
	bnet "github.com/mrgreenturtle/bnetconnect/blizzardclient"
	cp "github.com/mrgreenturtle/bnetconnect/characterprofile"
)

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
	go helperChannel(b, seasonID, bracket, "druid", "moonkin", amount, chamoonkin)
	go helperChannel(b, seasonID, bracket, "shaman", "elemental", amount, chaele)
	go helperChannel(b, seasonID, bracket, "warlock", "affliction", amount, chaaff)
	go helperChannel(b, seasonID, bracket, "warlock", "demonology", amount, chademo)
	go helperChannel(b, seasonID, bracket, "warlock", "destruction", amount, chadestro)
	select {
	case shadow := <-chashadow:
		fmt.Printf("%s", shadow[0].ActiveSpec.Name)
	case arcane := <-chaarcane:
		fmt.Printf("%s", arcane[0].ActiveSpec.Name)
	case fire := <-chafire:
		fmt.Printf("%s", fire[0].ActiveSpec.Name)
	case frost := <-chafrost:
		fmt.Printf("%s", frost[0].ActiveSpec.Name)
	case moonkin := <-chamoonkin:
		fmt.Printf("%s", moonkin[0].ActiveSpec.Name)
	case ele := <-chaele:
		fmt.Printf("%s", ele[0].ActiveSpec.Name)
	case aff := <-chaaff:
		fmt.Printf("%s", aff[0].ActiveSpec.Name)
	case demo := <-chademo:
		fmt.Printf("%s", demo[0].ActiveSpec.Name)
	case destro := <-chadestro:
		fmt.Printf("%s", destro[0].ActiveSpec.Name)
	}

}
