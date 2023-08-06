// Main package.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/bunkercoin/bunkerbot-go/bkc"

	"github.com/bwmarrin/discordgo"
)

var apiFlag = flag.String("a", "https://bkcexplorer.bunker.mt/api/", "The Bunkercoin API that should be used.")

// help string for the '.help' command
var helpCmd = `
# Bunkerbot Help:
* **.help** - show a list of commands
* **.bkcblocks** - show the number of blocks in the Bunkercoin Blockchain
* **.bkchashrate** - show the current network hashrate
* **.bkcdiff** - show the current network difficulty
* **.bkcinfo** - show chain information(Blockcount, hashrate, difficulty)
`

// Set the bot up.
func main() {
	flag.Parse()

	if os.Getenv("BOT_TOKEN") == "" {
		log.Fatal("BOT_TOKEN Environment Variable is not set!")
	}

	dg, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}

	// Wait here until CTRL-C or another term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

// Helper function that creates the chaininfo embed
func createChainEmbed(chaininfo bkc.ChainInfo) *discordgo.MessageEmbed {
	embed := NewEmbed().
		SetTitle("Bunkercoin Chain Information").
		SetDescription("Info about the Bunkercoin Blockchain").
		AddField("Number of blocks", fmt.Sprintf("**%d**", chaininfo.BlockCount)).
		AddField("Network Hashrate", fmt.Sprintf("**%.2f**", chaininfo.Hashrate/1000)).
		AddField("Network Difficulty", fmt.Sprintf("**%.2f**", chaininfo.Difficulty)).
		// SetImage("https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048").
		// SetThumbnail("https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048").
		SetColor(0x77c277).MessageEmbed

	return embed
}

// Handle MessageCreate events
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if cmd, ok := strings.CutPrefix(m.Content, "."); ok {
		switch cmd {
		case "":
			break
		case "help":
			s.ChannelMessageSend(m.ChannelID, helpCmd)
		case "bkcblocks":
			blockcount, err := bkc.GetBlockCount(*apiFlag)
			if err != nil {
				log.Println(err)
				s.ChannelMessageSend(m.ChannelID, "Error while getting the number of blocks!")
				break
			}
			s.ChannelMessageSend(m.ChannelID, "Number of blocks in the Bunkercoin Blockchain: "+strconv.Itoa(blockcount))
		case "bkchashrate":
			hashrate, err := bkc.GetHashrate(*apiFlag)
			if err != nil {
				log.Println(err)
				s.ChannelMessageSend(m.ChannelID, "Error while getting the network hashrate!")
				break
			}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Current network hashrate: %.2f KH/s", hashrate/1000))
		case "bkcdiff":
			diff, err := bkc.GetHashrate(*apiFlag)
			if err != nil {
				log.Println(err)
				s.ChannelMessageSend(m.ChannelID, "Error while getting the network difficulty!")
				break
			}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Current network difficulty: %.2f", diff))
		case "bkcinfo":
			chaininfo, err := bkc.GetChainInfo(*apiFlag)
			if err != nil {
				log.Println(err)
				s.ChannelMessageSend(m.ChannelID, "Error while getting Bunkercoin Chain Information!")
				break
			}
			s.ChannelMessageSendEmbed(m.ChannelID, createChainEmbed(chaininfo))
		default:
			s.ChannelMessageSend(m.ChannelID, "Invalid Command! Type .help to get a list of commands!")
		}
	}
}
