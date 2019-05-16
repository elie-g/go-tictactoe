package messages

import (
    "fmt"
)

type Message string

//noinspection GoSnakeCaseUsage
const (
    MSG_GAME_WIN            Message = "Vous avez gagné!"
    MSG_GAME_LOST           Message = "Vous avez perdu!"
    MSG_NEW_GAME            Message = "Nouvelle partie?"
    MSG_BTN_RESUME_GAME     Message = "Continuer"
    MSG_BTN_PLAYER_1        Message = "Joueur 1"
    MSG_BTN_PLAYER_2        Message = "Joueur 2"
    MSG_BTN_EXIT_GAME       Message = "Quitter"
    MSG_BTN_ONLINE          Message = "En ligne"
    MSG_BTN_JOIN_ONLINE     Message = "Rejoindre"
    MSG_BTN_CREATE_ONLINE   Message = "Creer"
    MSG_BTN_RETURN          Message = "Retour"
    MSG_PLAYER_1_BOX_TITLE  Message = "Joueur 1"
    MSG_PLAYER_1_BOX        Message = "Nom du joueur 1:"
    MSG_PLAYER_2_BOX_TITLE  Message = "Joueur 2 (AI)"
    MSG_PLAYER_2_BOX        Message = "Nom du joueur 2 (AI):"
    MSG_NAME_TOO_LONG_TITLE Message = "Nom trop long"
    MSG_NAME_TOO_LONG       Message = "Le nom ne peut avoir plus de %d caractères."
    MSG_CHOOSE_PLAYER_TITLE Message = "Choisir un joueur"
    MSG_CHOOSE_PLAYER       Message = "Est-ce l'AI qui joue?"
    MSG_FIRST_PLAYER_TITLE  Message = "Qui est le premier joueur?"
    MSG_FIRST_PLAYER        Message = "Êtes-vous le premier à jouer?"
    MSG_NO_GAME_AVAILABLE   Message = "Aucun partie disponible"
)

func (msg Message) Str(params ...interface{}) string {
    return fmt.Sprintf(string(msg), params...)
}
