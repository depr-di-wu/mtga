package mtga

import (
	"github.com/di-wu/mtga/thread/incoming/deck"
	"github.com/di-wu/mtga/thread/incoming/event"
	"github.com/di-wu/mtga/thread/incoming/front_door"
	"github.com/di-wu/mtga/thread/incoming/inventory"
	"github.com/di-wu/mtga/thread/incoming/mercantile"
	"github.com/di-wu/mtga/thread/incoming/mot_d"
	"github.com/di-wu/mtga/thread/incoming/progression"
	"github.com/di-wu/mtga/thread/incoming/quest"
	"os"
	"path/filepath"
	"testing"
)

func TestIncoming(t *testing.T) {
	filePath := filepath.Join(os.Getenv("APPDATA"), "..", "LocalLow", "Wizards Of The Coast", "MTGA", "output_log.txt")

	parser := Parser{}
	parser.Incoming.OnConnectionDetails(func(details front_door.ConnectionDetails) {})
	parser.Incoming.OnGetDeckLists(func(decks []deck.Deck) {})
	parser.Incoming.OnGetPreconDecks(func(decks []deck.PreconDeck) {})
	parser.Incoming.OnGetCatalogStatus(func(status inventory.CatalogStatus) {})
	parser.Incoming.OnGetActiveEvents(func(events []event.ActiveEvent) {})
	parser.Incoming.OnGetCombinedRankInfo(func(info event.CombinedRankInfo) {})
	parser.Incoming.OnGetSeasonAndRankDetail(func(detail event.SeasonRankAndDetail) {})
	parser.Incoming.OnGetFormats(func(formats []inventory.Format) {})
	parser.Incoming.OnGetPlayerArtSkins(func(skins inventory.PlayerArtSkins) {})
	parser.Incoming.OnGetPlayerCards(func(cards inventory.PlayerCards) {})
	parser.Incoming.OnGetPlayerInventory(func(inventory inventory.PlayerInventory) {})
	parser.Incoming.OnGetProductCatalog(func(catalog inventory.ProductCatalog) {})
	parser.Incoming.OnGetRewardSchedule(func(schedule inventory.RewardSchedule) {})
	parser.Incoming.OnGetMotD(func(d mot_d.MotD) {})
	parser.Incoming.OnGetAllTracks(func(tracks []progression.Track) {})
	parser.Incoming.OnGetPlayerProgress(func(progress progression.PlayerProgress) {})
	parser.Incoming.OnGetAllProducts(func(products []mercantile.Product) {})
	parser.Incoming.OnGetStoreStatus(func(status mercantile.StoreStatus) {})
	parser.Incoming.OnGetPlayerQuests(func(quests []quest.PlayerQuest) {})
	parser.Incoming.OnLogInfo(func(info []byte) {})

	tail, err := NewTail(filePath)
	if err != nil {
		t.Error(err)
	}

	for l := range tail.Logs() {
		parser.Parse(l)
	}
}
