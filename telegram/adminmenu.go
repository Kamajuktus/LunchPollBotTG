package telegram

import (
	tele "gopkg.in/tucnak/telebot.v2"
)

var f []string
var s []string
var a []string
var d []string

func (b *TeleBot) adminMenu() {
	b.Bot.Handle(&BtnNewMenu, func(m *tele.Message) {
		if !m.Private() {
			return
		}
		b.Bot.Send(m.Sender, "Введите первое блюдо.")
		b.Bot.Handle(tele.OnText, func(m *tele.Message) {
			f = append(f, m.Text)
			b.Bot.Send(m.Sender, "Введите второе блюдо.")
			b.Bot.Handle(tele.OnText, func(m *tele.Message) {
				s = append(s, m.Text)
				b.Bot.Send(m.Sender, "Введите апитайзер.")
				b.Bot.Handle(tele.OnText, func(m *tele.Message) {
					a = append(a, m.Text)
					b.Bot.Send(m.Sender, "Введите напиток.")
					b.Bot.Handle(tele.OnText, func(m *tele.Message) {
						d = append(d, m.Text)
						createMenu(f, s, a, d)
					})
				})
			})
		})

	})

}
