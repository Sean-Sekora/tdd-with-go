package wordz

type GameRepository interface {
	Create(player Player, word Word) Game
	FetchForPlayer(player Player) Game
	Update(game Game)
}
