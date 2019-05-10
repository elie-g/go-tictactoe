package message

type Message string

const (
    MSG_GAME_WIN  Message = "Vous avez gagné!"
    MSG_GAME_LOST Message = "Vous avez perdu!"
    MSG_NEW_GAME  Message = "Nouvelle partie?"
    MSG_YES       Message = "Oui"
    MSG_NO        Message = "Non"
)
