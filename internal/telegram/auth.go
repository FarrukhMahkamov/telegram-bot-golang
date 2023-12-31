package telegram

import (
	"context"
	"fmt"
	"strconv"
)

func (b *Bot) generateAuthorizationLink(chatID int64) (string, error) {
	redirectUrl := b.generateRedirectLink(chatID)

	requestToken, err := b.pocketClient.GetRequestToken(context.Background(), redirectUrl)
	if err != nil {
		return "", err
	}

	return b.pocketClient.GetAuthorizationURL(requestToken, redirectUrl)
}

func (b *Bot) generateRedirectLink(chatID int64) string {
	return fmt.Sprintf("%s?chat_id=%s", b.redirectUrl, strconv.FormatInt(chatID, 10))
}
