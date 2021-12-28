package telego

import (
	objs "github.com/SakoDroid/telego/objects"
)

type MessageCopier struct {
	bot                     *Bot
	disableNotif            bool
	messageId, replyTo      int
	caption, parseMode      string
	captionEntities         []objs.MessageEntity
	allowSendingWihtouReply bool
	replyMarkup             objs.ReplyMarkup
}

/*Copies the given message from a user to another user. chatId is the user that message is being copied to and fromChatId is the user that message is being copied to.*/
func (mf *MessageCopier) CopyFromUserToUser(chatId, fromChatId int) (*objs.SendMethodsResult, error) {
	return mf.bot.apiInterface.CopyMessage(chatId, fromChatId, "", "", mf.messageId, mf.disableNotif, mf.caption, mf.parseMode, mf.replyTo, mf.allowSendingWihtouReply, mf.replyMarkup, mf.captionEntities)
}

/*Copies the given message from a user to a channel. chatId is the channel that message is being copied to and fromChatId is the user that message is being copied to.*/
func (mf *MessageCopier) CopyFromUserToChannel(chatId string, fromChatId int) (*objs.SendMethodsResult, error) {
	return mf.bot.apiInterface.CopyMessage(0, fromChatId, chatId, "", mf.messageId, mf.disableNotif, mf.caption, mf.parseMode, mf.replyTo, mf.allowSendingWihtouReply, mf.replyMarkup, mf.captionEntities)
}

/*Copies the given message from a channel to a user. chatId is the user that message is being copied to and fromChatId is the channel that message is being copied to.*/
func (mf *MessageCopier) CopyFromChannelToUser(chatId int, fromChatId string) (*objs.SendMethodsResult, error) {
	return mf.bot.apiInterface.CopyMessage(chatId, 0, "", fromChatId, mf.messageId, mf.disableNotif, mf.caption, mf.parseMode, mf.replyTo, mf.allowSendingWihtouReply, mf.replyMarkup, mf.captionEntities)
}

/*Copies the given message from a channel to another channel. chatId is the channel that message is being copied to and fromChatId is the channel that message is being copied to.*/
func (mf *MessageCopier) CopyFromChannelToChannel(chatId, fromChatId string) (*objs.SendMethodsResult, error) {
	return mf.bot.apiInterface.CopyMessage(0, 0, chatId, fromChatId, mf.messageId, mf.disableNotif, mf.caption, mf.parseMode, mf.replyTo, mf.allowSendingWihtouReply, mf.replyMarkup, mf.captionEntities)
}
