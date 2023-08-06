# [WIP] Bunkerbot-go [WIP]

This bot is still a Work-In-Progress and not all features are implemented yet.

The (rewritten) Discord Bot for the Bunkercoin Discord Server.

## How to use this

To compile this bot, you need the go tool installed.

### Optional Flags

* -a - Set a different Bunkercoin API URL (default: https://bkcexplorer.bunker.mt/api/)

### Compiling

To compile the bot, just run this command:

```
make
```

### Running the binary

If you successfully compiled the bot, run this command to start it:

```
BOT_TOKEN=<YOUR TOKEN HERE> ./bin/bunkerbot-go
```

### Developing

To be able to quickly run & compile the bot while developing, set the 'TOKEN' Variable at the top of the Makefile.

Afterwards, you can just run

```
make run
```

to compile & run the bot

## Available commands

* .help - Show a list of commands
* .bkcblocks - Show the number of blocks in the Bunkercoin blockchain
* .bkchashrate - Show the Bunkercoin network hashrate
* .bkcdiff - Show the Bunkercoin network difficulty
* .bkcinfo - Show all of the above with a nice embed

## Disclaimer

This bot uses the [discordgo library](https://github.com/bwmarrin/discordgo/)

The embed.go file to simplify creating embeds was made by [NecroForger](https://gist.github.com/Necroforger/8b0b70b1a69fa7828b8ad6387ebb3835)
