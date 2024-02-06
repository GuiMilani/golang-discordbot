package bot

import (
	"fmt"
	"golang_discordbot/config"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Username != "Matheus Oliveira" {
		return
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, fetchAPIs())
}

func fetchAPIs() string {
	c1 := make(chan string)
	c2 := make(chan string)

	go func(ch chan string) {
			ch <- requestQuote("taylor")
		}(c1)

	go func(ch chan string) {
			ch <- requestQuote("kanye")
		}(c2)

	if(random75() == 1){
		return <-c1
	} else {
		return <-c2
	}
}

func random75() int {
	return random50() | random50()
}

func random50() int {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	return r.Intn(2) & 1
}

func requestQuote(artist string) string {
	const serverPort = 3333
	var requestURL string

	if artist == "kanye" {
		requestURL = fmt.Sprintf("http://localhost:%d/kanye", serverPort)
	} else {
		requestURL = fmt.Sprintf("http://localhost:%d/taylor", serverPort)
	}

	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	return string(resBody)
}