package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type asset struct {
	cards []int
	money int64
}

func start(conn net.Conn, client int) {
	host := asset{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 100}
	guest := asset{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 100}

	for i := 10; i > 0; i-- {
		ind1 := rand.Intn(i)
		ind2 := rand.Intn(i)
		hcard := host.cards[ind1]
		gcard := guest.cards[ind2]

		//fmt.Fprintf(conn, "Round [%d] - Your Money Remain: %d. You get the card [%d], how much do you bet\n", 10-i, guest.money, gcard)
		fmt.Fprintf(conn, "%d %d %d %d %d 0\n", 10-i, guest.money, host.money, gcard, hcard)

		host.cards = remove(host.cards, ind1)
		guest.cards = remove(guest.cards, ind2)

		for {
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil || message == "" {
				time.Sleep(1000 * time.Microsecond)
				continue
			}
			if num, err := strconv.ParseInt(message[0:len(message)-1], 10, 64); err == nil {
				if num > guest.money || num < 0 {
					fmt.Fprintf(conn, "Please give me a number between 0 and %d\n", guest.money)
					continue
				}

				if rand.Intn(10)%2 == 1 {
					num = 0
				}

				if hcard > gcard {
					guest.money -= num
					host.money += num
				} else {
					guest.money += num
					host.money -= num
				}
				fmt.Printf("Client[%d]: Host[%d] vs Client[%d]. Host Money Remain: %d\n", client, hcard, gcard, host.money)

				if guest.money <= 0 {
					defer conn.Close()
					fmt.Fprintf(conn, "%d %d %d %d %d 1\n", 10-i, guest.money, host.money, gcard, hcard)
					//conn.Write([]byte("Opps! You Lose!!!!!\n"))
					return
				}
				if host.money <= 0 {
					defer conn.Close()
					fmt.Fprintf(conn, "%d %d %d %d %d 1\n", 10-i, guest.money, host.money, gcard, hcard)
					//conn.Write([]byte("Congradulations! You Win!!!!!\n"))
					return
				}

				break
			} else {
				fmt.Fprintf(conn, "Please give me a number between 0 and %d\n", guest.money)
			}
		}
	}

	defer conn.Close()
	if host.money > guest.money {
		fmt.Fprintf(conn, "Opps! You Lose !!!!!\n")
	} else {
		fmt.Fprintf(conn, "Congradulations! You Win !!!!!\n")
	}
}

func remove(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
