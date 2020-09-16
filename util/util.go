package arenainsight

import (
	arena "github.com/mrgreenturtle/bnetconnect/arenaleaderboard"
	bnet "github.com/mrgreenturtle/bnetconnect/blizzardclient"
	cp "github.com/mrgreenturtle/bnetconnect/characterprofile"
	"strings"
)

func GetArenaEntries(b *bnet.BnetClient, seasonID int, bracket string) arena.ArenaLeaderBoard {
	returnArenaEntriesOnly := arena.ArenaLeaderBoard{}
	temp := b.GenerateArenaLeaderBoard(seasonID, bracket)
	returnArenaEntriesOnly.Entries = temp.Entries
	return returnArenaEntriesOnly
}

func GetTopHundredArenaSpec(b *bnet.BnetClient, seasonID int, bracket string, class string, spec string) (arena.ArenaLeaderBoard, []cp.CharacterProfile) {

	topClass := arena.ArenaLeaderBoard{}
	topProfile := []cp.CharacterProfile{}
	tempAlb := GetArenaEntries(b, seasonID, bracket)
	tempAlb.Entries = tempAlb.Entries[0:100]
	for _, entry := range tempAlb.Entries {
		name := strings.ToLower(entry.Character.Name)
		realm := strings.ToLower(entry.Character.Realm.Slug)
		tempCp := b.GenerateCharacterProfile(name, realm)
		if strings.ToLower(tempCp.CharacterClass.Name) == class && strings.ToLower(tempCp.ActiveSpec.Name) == spec {
			topClass.Entries = append(topClass.Entries, entry)
			topProfile = append(topProfile, tempCp)
		}
	}
	return topClass, topProfile
}
