## GoIRCrelay
This tool acts as a relay between Telegram and IRC.

[![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

### Setup
```golang
conf := config{
	ircnick:    "gorelay", // Nickname of the bot in the IRC channel.
	channel:    "#go-eventirc-test", // IRC channel where the bot will set into. Be sure to be one of your own.
	ircserver:  "irc.freenode.net:7000", // IRC SSL server to connect to. Format <IP>:<PORT>
	tgbot_key:  "", // API key of the Telegram bot.
	tg_channel: , // Channel ID of the group you want to be using the bot. 
}
```

How does it work?
We just have to fill in the previous data so that the bot is able to connect to an IRC server and introduce the bot to a telegram group. If configured correctly, messages will be sent both ways allowing the IRC clients to receive the Telegram messages and the IRC users to communicate with the Telegram users.

Developed by <ineedblood> and released under the GPL 3 license.
