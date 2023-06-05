package telegram

import (
	"errors"
	"myGoApp/clients/telegram"
	"myGoApp/events"
	"myGoApp/lib/e"
	"myGoApp/storage"
	"myGoApp/storage/files"
)

type TelProcessor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

type MessageMeta struct {
	ChatID   int
	Username string
}

type CallbackMeta struct {
	Data   string
	ChatID int
}

var (
	ErrUnknownEventType = errors.New("unknown event type")
	ErrUnknownMetaType  = errors.New("unknown mera type")
)

func New(client *telegram.Client, storage files.Storage) *TelProcessor {
	return &TelProcessor{
		tg:      client,
		storage: storage,
	}
}
func (p *TelProcessor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("can't get events", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, event(u))
	}

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func (p *TelProcessor) Process(event events.Event) error {
	switch event.Type {
	case events.CallbackQuery:
		return p.processMessage(event)
	case events.Message:
		return p.processMessage(event)
	default:
		return e.Wrap("can't process message", ErrUnknownEventType)
	}
}

func (p *TelProcessor) processMessage(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return e.Wrap("can't process message", err)
	}

	if err := p.processInputMessage(event.Text, meta.ChatID, meta.Username); err != nil {
		return e.Wrap("can't process message", err)
	}

	return nil
}

func meta(event events.Event) (MessageMeta, error) {
	res, ok := event.Meta.(MessageMeta)
	if !ok {
		return MessageMeta{}, e.Wrap("can't get meta", ErrUnknownMetaType)
	}

	return res, nil
}

func event(upd telegram.Update) events.Event {
	updType := fetchType(upd)
	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}
	if updType == events.CallbackQuery {
		res = events.Event{
			Type: updType,
			Text: upd.CallbackQuery.Data,
		}
		res.Meta = MessageMeta{
			Username: upd.CallbackQuery.Message.From.Username,
			ChatID:   upd.CallbackQuery.Message.Chat.ID}
	}
	if updType == events.Message {
		res.Meta = MessageMeta{
			ChatID:   upd.Message.Chat.ID,
			Username: upd.Message.From.Username}
	}
	return res
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}

	return upd.Message.Text
}
func fetchCallbackQuery(upd telegram.Update) string {
	if upd.CallbackQuery == nil {
		return ""
	}

	return upd.CallbackQuery.Data
}
func fetchType(upd telegram.Update) events.Type {

	if upd.Message == nil {
		return events.CallbackQuery
	}

	return events.Message
}
