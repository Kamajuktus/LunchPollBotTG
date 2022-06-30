package telegram

import (
	tele "gopkg.in/tucnak/telebot.v2"
)

var cafeIndexes []int
var cafeName string

var found = false

func (b *TeleBot) Registration() {
	getCafe()
	b.Bot.Handle("/start", func(m *tele.Message) {
		if !m.Private() {
			return
		}

		b.Bot.Send(m.Sender, "Пожалуйста авторизуйтесь!", Auth)
	})

	b.Bot.Handle(&BtnAuthCafe, func(m *tele.Message) {
		b.Bot.Send(m.Sender, "Укажите название кафе")
		b.Bot.Handle(tele.OnText, func(m *tele.Message) {
			for index, name := range cafename {

				if name == m.Text {
					cafeIndexes = append(cafeIndexes, index)
					cafeName = name
					found = true
					b.Bot.Send(m.Sender, "Установите пароль")
					b.Bot.Handle(tele.OnText, func(m *tele.Message) {
						for index, name := range cafepass {
							if index == cafeIndexes[0] && name == m.Text {
								found = false
								b.Bot.Send(m.Sender, "Вы успешно авторизовались!", menuAdmin)

								b.adminMenu()
							}
						}
						if found {
							b.Bot.Send(m.Sender, "Неправильный пароль!")
						}
					})

				}

			}
			if !found {
				b.Bot.Send(m.Sender, "Неправильное название кафе!")
			}

		})
	})

	b.Bot.Handle(&BtnAuthCompany, func(m *tele.Message) {
		getCompany()
		b.Bot.Send(m.Sender, "Введите пароль")
		b.Bot.Handle(tele.OnText, func(m *tele.Message) {
			if adminpassword == m.Text {
				b.Bot.Send(m.Sender, "Вы успешно авторизовались!")
			} else {
				b.Bot.Send(m.Sender, "Неправильный пароль!")
			}

		})
	})

}
