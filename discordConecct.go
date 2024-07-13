package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func sendDiscordMessage(dg *discordgo.Session, channelID string, userData []Userdata) error {
	jst := time.FixedZone("JST", 9*60*60)
	now := time.Now().In(jst)
	oneDayAgo := now.AddDate(0, 0, -1)
	oneWeekAgo := now.AddDate(0, 0, -7)
	const layout = "2006/01/02"
	startDay := oneWeekAgo.Format(layout)
	endDay := oneDayAgo.Format(layout)

	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("🍑 週間ランキング結果 🍑\n%s ~ %s\n", startDay, endDay),
		Color: 0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   fmt.Sprintf("**1位  %s**", userData[0].Name),
				Value:  fmt.Sprintf("**コントリビューション数: %d**", userData[0].Contributions),
				Inline: false,
			},
			{
				Name:   fmt.Sprintf("**2位  %s**", userData[1].Name),
				Value:  fmt.Sprintf("**コントリビューション数: %d**", userData[1].Contributions),
				Inline: false,
			},
			{
				Name:   fmt.Sprintf("**3位  %s**", userData[2].Name),
				Value:  fmt.Sprintf("**コントリビューション数: %d**", userData[2].Contributions),
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "今週も一位目指して頑張ろう！！",
		},
	}

	_, err := dg.ChannelMessageSendEmbed(channelID, embed)
		err = pushRankingData(userData)
	return err
}
