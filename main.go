package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/hugolgst/rich-go/client"

	"fyne.io/fyne/v2/layout"
)

func main() {

	a := app.New()
	w := a.NewWindow("Watson")

	//Reading things
	BookToRead := widget.NewEntry()
	BookToRead.PlaceHolder = "The Book Your are Reading"
	BookGenre := widget.NewEntry()
	BookGenre.PlaceHolder = "The Genre of your Book"
	LinkToBook := widget.NewEntry()
	LinkToBook.PlaceHolder = "The Link To The Book"

	//Talking things
	TalkToWho := widget.NewEntry()
	TalkToWho.PlaceHolder = "Who Are you talking to?"

	WhereTalk := widget.NewEntry()
	WhereTalk.PlaceHolder = "Link to a Discord Server you are Talking on"

	//Browsing things

	Site := widget.NewEntry()
	Site.PlaceHolder = "Name of the Site You are Browsing"

	SiteURL := widget.NewEntry()
	SiteURL.PlaceHolder = "The Link To The Site"

	//Streaming things

	TitleOfStream := widget.NewEntry()
	TitleOfStream.PlaceHolder = "Title of your Stream"

	StreamLink := widget.NewEntry()
	StreamLink.PlaceHolder = "Link To your Stream"

	//Working Tthings

	TitleOfProject := widget.NewEntry()
	TitleOfProject.PlaceHolder = "Project Name"

	ProjectLink := widget.NewEntry()
	ProjectLink.PlaceHolder = "Link to Project"

	BrowsingPage := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),
		Site,
		SiteURL,
		layout.NewSpacer(),
	)

	TalkingPage := container.New(
		layout.NewVBoxLayout(),

		layout.NewSpacer(),
		TalkToWho,
		WhereTalk,
		layout.NewSpacer(),
	)

	ReadingPage := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),

		layout.NewSpacer(),
		BookToRead,
		BookGenre,
		LinkToBook,
		layout.NewSpacer(),
	)

	StramingPage := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),
		TitleOfStream,
		StreamLink,
		layout.NewSpacer(),
	)

	WorkingPage := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),
		TitleOfProject,
		ProjectLink,
		layout.NewSpacer(),
	)

	HomePage := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),
		widget.NewSelect([]string{"Talking", "Reading", "Streaming", "Browsing", "Working"}, func(value string) {
			switch value {

			case "Reading":
				StartReading(BookToRead.Text, BookGenre.Text, LinkToBook.Text)

			case "Talking":
				StartTalking(TalkToWho.Text, WhereTalk.Text)

			case "Browsing":
				StartBrowsing(Site.Text, SiteURL.Text)

			case "Streaming":
				StartStreaming(TitleOfStream.Text, StreamLink.Text)

			case "Working":
				StartWorking(TitleOfProject.Text, ProjectLink.Text)
			}

		}),
		layout.NewSpacer(),
	)

	tabs := container.NewAppTabs(
		container.NewTabItem("Home", HomePage),
		container.NewTabItem("Reading", ReadingPage),
		container.NewTabItem("Talking", TalkingPage),
		container.NewTabItem("Browsing", BrowsingPage),
		container.NewTabItem("Streaming", StramingPage),
		container.NewTabItem("Working", WorkingPage),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)

	w.Resize(fyne.NewSize(1920/2, 1080/2))

	w.ShowAndRun()

}

func StartReading(NameOfBook string, BookGenre string, LinkToBook string) {

	if LinkToBook == "" {
		err := client.Login("1151169281468416092")
		if err != nil {
			panic(err)
		}

		now := time.Now()
		err = client.SetActivity(client.Activity{

			Details:    NameOfBook,
			LargeImage: "open-book",
			LargeText:  "Genre: " + BookGenre,

			Timestamps: &client.Timestamps{
				Start: &now,
			},
		})
	} else {
		err := client.Login("1151169281468416092")
		if err != nil {
			panic(err)
		}

		now := time.Now()
		err = client.SetActivity(client.Activity{

			Details:    NameOfBook,
			LargeImage: "open-book",
			LargeText:  "Genre: " + BookGenre,

			Timestamps: &client.Timestamps{
				Start: &now,
			},
			Buttons: []*client.Button{
				&client.Button{
					Label: "Link To Book",
					Url:   LinkToBook,
				},
			},
		})
	}

}

func StartTalking(NameOfPartner string, DiscordLink string) {

	if DiscordLink == "" {
		err := client.Login("1151168383220449420")
		if err != nil {
			panic(err)
		}

		now := time.Now()
		err = client.SetActivity(client.Activity{

			Details:    "I'm talking with " + NameOfPartner,
			LargeImage: "talk",

			Timestamps: &client.Timestamps{
				Start: &now,
			},
		})
	} else {
		err := client.Login("1151168383220449420")
		if err != nil {
			panic(err)
		}

		now := time.Now()
		err = client.SetActivity(client.Activity{

			Details:    "I'm talking with " + NameOfPartner,
			LargeImage: "talk",

			Timestamps: &client.Timestamps{
				Start: &now,
			},
			Buttons: []*client.Button{
				&client.Button{
					Label: "Link To Book",
					Url:   DiscordLink,
				},
			},
		})
	}
}

func StartBrowsing(Site string, SiteURL string) {
	if SiteURL == "" {
		err := client.Login("1151167934887116931")
		if err != nil {
			panic(err)
		}

		now := time.Now()
		err = client.SetActivity(client.Activity{

			Details:    "Currently on " + Site,
			LargeImage: "search",

			Timestamps: &client.Timestamps{
				Start: &now,
			},
		})
	} else {
		err := client.Login("1151167934887116931")
		if err != nil {
			panic(err)
		}

		now := time.Now()
		err = client.SetActivity(client.Activity{

			Details:    "Currently on " + Site,
			LargeImage: "search",

			Timestamps: &client.Timestamps{
				Start: &now,
			},
			Buttons: []*client.Button{
				&client.Button{
					Label: "Visit Site",
					Url:   SiteURL,
				},
			},
		})
	}

}

func StartStreaming(Title string, URL string) {

	if URL == "" {
		err := client.Login("1151167472037286030")
		if err != nil {
			panic(err)
		}

		now := time.Now()
		err = client.SetActivity(client.Activity{
			Details:    Title,
			LargeImage: "stream",

			Timestamps: &client.Timestamps{
				Start: &now,
			},
		})
	} else {
		err := client.Login("1151167472037286030")
		if err != nil {
			panic(err)
		}

		now := time.Now()
		err = client.SetActivity(client.Activity{
			Details:    Title,
			LargeImage: "stream",

			Timestamps: &client.Timestamps{
				Start: &now,
			},
			Buttons: []*client.Button{
				&client.Button{
					Label: "Watch my Stream",
					Url:   URL,
				},
			},
		})
	}

}

func StartWorking(Title string, URL string) {

	if URL == "" {
		err := client.Login("1151165670634033193")
		if err != nil {
			panic(err)
		}

		now := time.Now()
		err = client.SetActivity(client.Activity{
			State: Title,

			LargeImage: "keyboard",

			Timestamps: &client.Timestamps{
				Start: &now,
			},
		})
	} else {
		err := client.Login("1151165670634033193")
		if err != nil {
			panic(err)
		}

		now := time.Now()
		err = client.SetActivity(client.Activity{
			State:      Title,
			LargeImage: "keyboard",

			Timestamps: &client.Timestamps{
				Start: &now,
			},
			Buttons: []*client.Button{
				&client.Button{
					Label: "Visit my Project",
					Url:   URL,
				},
			},
		})
	}

}
