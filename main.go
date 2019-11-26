package main

import (
	"bufio"
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/gtk"
	"net"
	"os"
	"project1/uiControl"
)

var conn net.Conn

func main() {
	// initilze gtk
	gtk.Init(&os.Args)

	// load glade
	builder := gtk.NewBuilder()
	builder.AddFromFile("ui.glade")

	// get widget from glade
	window := gtk.WindowFromObject(builder.GetObject("window1"))
	var i, j int
	window.Connect("configure_event", func() {
		window.GetSize(&i, &j)
	})
	// set up window background
	window.SetAppPaintable(true)
	window.Connect("expose-event", func() {
		painter := window.GetWindow().GetDrawable()
		gc := gdk.NewGC(painter)
		backGroundimg, _ := gdkpixbuf.NewPixbufFromFileAtScale("images/scene/table.png", i, j, false)
		painter.DrawPixbuf(gc, backGroundimg, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
		backGroundimg.Unref()

	})
	// two cards image config
	img1 := gtk.ImageFromObject(builder.GetObject("image1"))
	img2 := gtk.ImageFromObject(builder.GetObject("image2"))

	// set up labels
	roundLabel := gtk.LabelFromObject(builder.GetObject("round"))
	roundLabel.SetText("Round")
	roundnumLabel := gtk.LabelFromObject(builder.GetObject("roundnum"))
	moneyLabel := gtk.LabelFromObject(builder.GetObject("money"))
	moneyLabel.SetText("Money")
	moneynumLabel := gtk.LabelFromObject(builder.GetObject("moneynum"))
	infoLabel := gtk.LabelFromObject(builder.GetObject("info"))

	youLabel := gtk.LabelFromObject(builder.GetObject("youlabel"))
	oppLabel := gtk.LabelFromObject(builder.GetObject("opplabel"))

	// set up buttons
	bet10Button := gtk.ButtonFromObject(builder.GetObject("bet10"))
	bet20Button := gtk.ButtonFromObject(builder.GetObject("bet20"))
	bet50Button := gtk.ButtonFromObject(builder.GetObject("bet30"))
	createButton := gtk.ButtonFromObject(builder.GetObject("createButton"))
	joinButton := gtk.ButtonFromObject(builder.GetObject("joinButton"))
	nextButton := gtk.ButtonFromObject(builder.GetObject("nextRound"))
	foldButton := gtk.ButtonFromObject(builder.GetObject("foldButton"))

	process01 := uiControl.GameProcess{
		LabelInfo:  infoLabel,
		LabelRound: roundnumLabel,
		LabelMoney: moneynumLabel,
		Img1:       img1,
		Img2:       img2,
		Betnumber:  "0",
		YouLabel:   youLabel,
		OppLabel:   oppLabel,
	}
	foldButton.Clicked(uiControl.Bet, process01)

	process0 := uiControl.GameProcess{
		LabelInfo:  infoLabel,
		LabelRound: roundnumLabel,
		LabelMoney: moneynumLabel,
		Img1:       img1,
		Img2:       img2,
		Betnumber:  "",
		YouLabel:   youLabel,
		OppLabel:   oppLabel,
	}
	nextButton.Clicked(uiControl.NextButton, process0)

	// set up buttons actions
	process1 := uiControl.GameProcess{
		LabelInfo:  infoLabel,
		LabelRound: roundnumLabel,
		LabelMoney: moneynumLabel,
		Img1:       img1,
		Img2:       img2,
		Betnumber:  "10",
		YouLabel:   youLabel,
		OppLabel:   oppLabel,
	}
	process2 := uiControl.GameProcess{
		LabelInfo:  infoLabel,
		LabelRound: roundnumLabel,
		LabelMoney: moneynumLabel,
		Img1:       img1,
		Img2:       img2,
		Betnumber:  "20",
		YouLabel:   youLabel,
		OppLabel:   oppLabel,
	}
	process3 := uiControl.GameProcess{
		LabelInfo:  infoLabel,
		LabelRound: roundnumLabel,
		LabelMoney: moneynumLabel,
		Img1:       img1,
		Img2:       img2,
		Betnumber:  "50",
		YouLabel:   youLabel,
		OppLabel:   oppLabel,
	}

	bet10Button.Clicked(uiControl.Bet, process1)
	bet20Button.Clicked(uiControl.Bet, process2)
	bet50Button.Clicked(uiControl.Bet, process3)

	createButton.Clicked(uiControl.Create, infoLabel)

	joinInfo := uiControl.GameJoin{
		LabelInfo:  infoLabel,
		LabelRound: roundnumLabel,
		LabelMoney: moneynumLabel,
		Builder:    builder,
		Img1:       img1,
		Img2:       img2,
	}
	joinButton.Clicked(uiControl.Join, joinInfo)

	// show
	window.Show()
	// main event looping
	gtk.Main()
}

func show(t1 *gtk.Label) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			conn.Close()
			os.Exit(0)
		}
		fmt.Println(message)
		t1.SetText(message)
	}
}
