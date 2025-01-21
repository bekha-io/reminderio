package main

import "gopkg.in/telebot.v4"

func SendMessage(b *telebot.Bot, recipients []int, text string) error {
	for _, recipient := range recipients {
		_, err := b.Send(telebot.ChatID(recipient), text)
		if err != nil {
			return err
		}
	}
	return nil
}