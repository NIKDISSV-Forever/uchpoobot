package main

import (
	"flag"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/nikdissv-forever/uchpoobot/imagesearch"
)

var token string

func init() {
	flag.StringVar(&token, "token", "", "Bot Token")
	flag.Parse()
}

var urlsCache []string
var pg = uint16(rand.Intn(0xFFFF + 1))

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Println(err)
		return
	}

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentMessageContent

	if err = dg.Open(); err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err = dg.Close(); err != nil {
			log.Println(err)
			return
		}
	}()
	log.Println("Bot is running")
	for {
		time.Sleep(math.MaxInt - 1)
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if len(urlsCache) == 0 {
		var err error
		urlsCache, err = imagesearch.Urls("учпучмак", pg)
		pg++
		if err != nil {
			log.Println(err)
		}
	}
	if len(urlsCache) == 0 {
		return
	}
	index := rand.Intn(len(urlsCache))
	url := urlsCache[index]
	urlsCache = append(urlsCache[:index], urlsCache[index+1:]...)
	get, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer get.Body.Close()
	splits := strings.Split(url, "/")
	if _, err = s.ChannelFileSend(m.ChannelID, splits[len(splits)-1], get.Body); err != nil {
		log.Println(err)
	}
}
