package uiControl

import (
	"bufio"
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"net"
	"strconv"
	"strings"
	"time"
)

var conn net.Conn
var cardNumber1 string
var cardNumber2 string
var round string
var money1 string
var money2 string
var img1 *gtk.Image
var img2 *gtk.Image

var infoLabel *gtk.Label

type GameJoin struct {
	LabelInfo   *gtk.Label
	LabelRound  *gtk.Label
	LabelMoney  *gtk.Label
	LabelMoney1 *gtk.Label
	Builder     *gtk.Builder
	Img1        *gtk.Image
	Img2        *gtk.Image
}

type GameProcess struct {
	LabelInfo   *gtk.Label
	LabelRound  *gtk.Label
	LabelMoney  *gtk.Label
	LabelMoney1 *gtk.Label
	Img1        *gtk.Image
	Img2        *gtk.Image
	Betnumber   string
	OppLabel    *gtk.Label
	YouLabel    *gtk.Label
}

func Create(ctk *glib.CallbackContext) {
	arg := ctk.Data()
	infoLabel, ok := arg.(*gtk.Label)
	if ok {
		infoLabel.SetText("create successful")
		infoLabel.ModifyFontSize(15)
		infoLabel.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("white"))
	}
}

func Join(ctk *glib.CallbackContext) {
	arg := ctk.Data()
	data, ok := arg.(GameJoin)
	if ok {
		infoLabel = data.LabelInfo
		roundLabel := data.LabelRound
		moneyLabel := data.LabelMoney
		moneyLabel1 := data.LabelMoney1
		img1 = data.Img1
		img2 = data.Img2

		conn, _ = net.Dial("tcp", "127.0.0.1:8081")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		arr := strings.Fields(message)

		round = arr[0]
		money1 = arr[1]
		money2 = arr[2]
		cardNumber1 = arr[3]
		cardNumber2 = arr[4]
		ChangePic(img1, "back")
		ChangePic(img2, cardNumber2)

		infoLabel.SetText("Enjoy")
		infoLabel.ModifyFontSize(15)
		infoLabel.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("white"))
		roundLabel.SetText(round)
		moneyLabel.SetText(money1)
		moneyLabel1.SetText(money2)
	}
}

func Bet(ctk *glib.CallbackContext) {
	arg := ctk.Data()
	data, ok := arg.(GameProcess)
	if ok {
		//roundLabel := data.LabelRound
		//moneyLabel := data.LabelMoney
		time.Sleep(2 * time.Second)
		betNumber := data.Betnumber
		img1 := data.Img1
		img2 := data.Img2
		youLabel := data.YouLabel
		oppLabel := data.OppLabel
		youLabel.SetText("You bet " + betNumber)
		youLabel.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("white"))

		oppLabel.SetText("He bet " + betNumber)
		oppLabel.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("white"))

		// send betnum to server and get the result
		fmt.Fprintf(conn, betNumber+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		arr := strings.Fields(message)
		card1, _ := strconv.ParseInt(cardNumber1, 10, 64)
		card2, _ := strconv.ParseInt(cardNumber2, 10, 64)
		if card1 > card2 {
			infoLabel.SetText("You win")
			infoLabel.ModifyFontSize(15)
			infoLabel.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("white"))
		} else {
			infoLabel.SetText("You lose")
			infoLabel.ModifyFontSize(15)
			infoLabel.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("white"))
		}

		RoundEnd(img1, img2)
		end := arr[5]
		if end == "1" {
			infoLabel.SetText("Game end")
		}
		round = arr[0]

		if money1 == arr[1] {
			oppLabel.SetText("fold")
		}
		money1 = arr[1]
		money2 = arr[2]
		cardNumber1 = arr[3]
		cardNumber2 = arr[4]
	}
}

func NextButton(ctk *glib.CallbackContext) {
	arg := ctk.Data()
	data, ok := arg.(GameProcess)
	if ok {
		img1 := data.Img1
		img2 := data.Img2
		youLabel := data.YouLabel
		oppLabel := data.OppLabel
		roundLabel := data.LabelRound
		moneyLabel := data.LabelMoney
		moneyLabel1 := data.LabelMoney1
		ChangePic(img1, "back")
		ChangePic(img2, cardNumber2)
		youLabel.SetText("")
		oppLabel.SetText("")
		roundLabel.SetText(round)
		moneyLabel.SetText(money1)
		moneyLabel1.SetText(money2)
	}
}

func RoundEnd(img1 *gtk.Image, img2 *gtk.Image) {
	ChangePic(img1, cardNumber1)
	ChangePic(img2, cardNumber2)
}

func ChangePic(img *gtk.Image, changeTo string) {
	var w, h int
	img.GetSizeRequest(&w, &h)
	pix, _ := gdkpixbuf.NewPixbufFromFileAtScale("images/poker/"+changeTo+".png", w, h, false)
	img.SetFromPixbuf(pix)
	pix.Unref()
}
