package retro

import (
	"context"
	"errors"
	"time"

	"github.com/kralamoure/retro/retrotyp"
)

var ErrNotFound = errors.New("not found")
var ErrAlreadyExists = errors.New("already exists")

var ErrGameServerHostAndPortAlreadyExist = errors.New("game server host and port already exist")
var ErrCharacterNameAndGameServerIdAlreadyExist = errors.New("character name and game server ID already exist")
var ErrTicketAccountIdAlreadyExists = errors.New("ticket account ID already exist")

type Repo interface {
	GameMaps(ctx context.Context) (map[int]GameMap, error)
	EffectTemplates(ctx context.Context) (map[int]EffectTemplate, error)
	ItemSets(ctx context.Context) (map[int]ItemSet, error)
	ItemTemplates(ctx context.Context) (map[int]ItemTemplate, error)
	NPCTemplates(ctx context.Context) (map[int]NPCTemplate, error)
	NPCDialogs(ctx context.Context) (map[int]NPCDialog, error)
	NPCResponses(ctx context.Context) (map[int]NPCResponse, error)
	Classes(ctx context.Context) (map[retrotyp.ClassId]Class, error)
	Spells(ctx context.Context) (map[int]Spell, error)
	MountTemplates(ctx context.Context) (map[int]MountTemplate, error)

	Triggers(ctx context.Context) (map[string]Trigger, error)
	TriggerByGameMapIdAndCellId(ctx context.Context, gameMapId, cellId int) (Trigger, error)

	CreateGameServer(ctx context.Context, gameServer GameServer) error
	GameServers(ctx context.Context) (map[int]GameServer, error)
	GameServer(ctx context.Context, id int) (GameServer, error)
	SetGameServerState(ctx context.Context, id int, state retrotyp.GameServerState) error

	CreateTicket(ctx context.Context, ticket Ticket) (string, error)
	DeleteTickets(ctx context.Context, before time.Time) (int, error)
	UseTicket(ctx context.Context, id string) (Ticket, error)
	Tickets(ctx context.Context) (map[string]Ticket, error)
	Ticket(ctx context.Context, id string) (Ticket, error)

	CreateCharacter(ctx context.Context, character Character) (int, error)
	UpdateCharacter(ctx context.Context, character Character) error
	DeleteCharacter(ctx context.Context, id int) error
	AllCharacters(ctx context.Context) (map[int]Character, error)
	AllCharactersByAccountId(ctx context.Context, accountId string) (map[int]Character, error)
	Characters(ctx context.Context, gameServerId int) (map[int]Character, error)
	CharactersByAccountId(ctx context.Context, gameServerId int, accountId string) (map[int]Character, error)
	CharactersByGameMapId(ctx context.Context, gameServerId int, gameMapId int) (map[int]Character, error)
	Character(ctx context.Context, id int) (Character, error)

	Markets(ctx context.Context, gameServerId int) (map[string]Market, error)

	CreateCharacterItem(ctx context.Context, item CharacterItem) (int, error)
	UpdateCharacterItem(ctx context.Context, item CharacterItem) error
	DeleteCharacterItem(ctx context.Context, id int) error
	CharacterItemsByCharacterId(ctx context.Context, characterId int) (map[int]CharacterItem, error)
	CharacterItem(ctx context.Context, id int) (CharacterItem, error)

	CreateMarketItem(ctx context.Context, item MarketItem) (int, error)
	DeleteMarketItem(ctx context.Context, id int) error
	MarketItems(ctx context.Context, gameServerId int) (map[int]MarketItem, error)
	MarketItemsByMarketId(ctx context.Context, marketId string) (map[int]MarketItem, error)
	MarketItem(ctx context.Context, id int) (MarketItem, error)

	NPCs(ctx context.Context, gameServerId int) (map[string]NPC, error)

	CreateMount(ctx context.Context, mount Mount) (int, error)
	UpdateMount(ctx context.Context, mount Mount) error
	DeleteMount(ctx context.Context, id int) error
	Mount(ctx context.Context, id int) (Mount, error)
	Mounts(ctx context.Context) (map[int]Mount, error)
	MountsByCharacterId(ctx context.Context, characterId int) (map[int]Mount, error)
}
