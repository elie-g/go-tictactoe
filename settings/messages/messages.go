package message

type Message string

//noinspection GoSnakeCaseUsage
const (
    MSG_GAME_WIN           Message = "Vous avez gagné!"
    MSG_GAME_LOST          Message = "Vous avez perdu!"
    MSG_NEW_GAME           Message = "Nouvelle partie?"
    MSG_BTN_RESUME_GAME    Message = "Continuer"
    MSG_BTN_PLAYER_1       Message = "Joueur 1"
    MSG_BTN_PLAYER_2       Message = "Joueur 2 (AI)"
    MSG_BTN_EXIT_GAME      Message = "Quitter"
    MSG_PLAYER_1_BOX_TITLE Message = "Joueur 1"
    MSG_PLAYER_1_BOX       Message = "Nom du joueur 1:"
    MSG_PLAYER_2_BOX_TITLE Message = "Joueur 2 (AI)"
    MSG_PLAYER_2_BOX       Message = "Nom du joueur 2 (AI):"
)

func (msg Message) Str() string {
    return string(msg)
}
