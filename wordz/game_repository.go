package wordz

type GameRepository interface {
	Create(game Game) Game
	FetchForPlayer(player Player) Game
	Update(game Game)
}
