package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type Message struct {
	Text string `json:"text"`
}
type Card struct {
	Rank string
	Suit string
}

type Response struct {
	HandType   string
	Rank       int64
	UniqueRank int64
	Hand       string
}

type Output struct {
	Responses []Response
}

var cardPower = map[string]int64{
	"2":  3,
	"3":  5,
	"4":  7,
	"5":  11,
	"6":  13,
	"7":  17,
	"8":  19,
	"9":  23,
	"10": 29,
	"11": 31,
	"12": 37,
	"13": 41,
	"1":  43,
}

var rankMapping = map[string]string{
	"A": "1",
	"T": "10",
	"J": "11",
	"Q": "12",
	"K": "13",
}
var rankMappingBack = map[string]string{
	"1":  "A",
	"10": "T",
	"11": "J",
	"12": "Q",
	"13": "K",
}

func main() {
	router := mux.NewRouter()
	router.Use(enableCORS)
	router.HandleFunc("/evaluate", postHandler)
	//http.HandleFunc("/evaluate", postHandler)
	fmt.Println("Server started at :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post Request Initialized...")
	var message Message
	var output Output
	var slices [][]Card
	// Parse JSON request body
	fmt.Println(message.Text)
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	//Parsing Input Card array...
	fmt.Println("Parsing input...")
	cards := parseInput(message.Text)
	for i := 0; i < len(cards); i += 5 {
		subarray := make([]Card, 0)
		subarray = append(subarray, cards[i:i+5]...)
		slices = append(slices, subarray)
	}
	fmt.Println("Evaluation started...")
	for i := 0; i < len(slices); i++ {
		var cardStrings []string
		for _, card := range slices[i] {
			cardStrings = append(cardStrings, fmt.Sprintf("%s%s", card.Rank, card.Suit))
		}
		combinedStr := strings.Join(cardStrings, ", ")
		response := FindBestPokerHandRank(slices[i], combinedStr)
		output.Responses = append(output.Responses, response)
	}
	fmt.Println("Evaluation Completed...")
	//Sending response back
	fmt.Println("Sending Evaluated Data...")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func parseInput(input string) []Card {
	input = strings.ToUpper(input)
	cardStrings := strings.Split(input, ",")
	if len(cardStrings)%5 != 0 {
		fmt.Println("Invalid card format:")
		return nil
	}
	var cards []Card

	for _, cardStr := range cardStrings {
		cardStr = strings.TrimSpace(cardStr)

		var rank, suit string
		switch len(cardStr) {
		case 2:
			rank, suit = string(cardStr[0]), string(cardStr[1])
		case 3:
			if cardStr[:2] == "10" {
				rank, suit = "10", string(cardStr[2])
			}
		}

		if rank == "" || suit == "" {
			fmt.Println("Invalid card format:", cardStr)
			return nil
		}

		cards = append(cards, Card{Rank: rank, Suit: suit})
	}

	return cards
}

func FindBestPokerHandRank(cards []Card, combinedStr string) Response {
	convertRanksInCards(cards)
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})
	uniqueRank := calculateMultiplication(cards)
	// Check for various poker hands in decreasing order of rank
	switch {
	case isRoyalFlush(cards):
		return Response{Rank: 1, UniqueRank: uniqueRank, HandType: "Royal Flush", Hand: combinedStr}
	case isStraightFlush(cards):
		if cards[0].Rank == "1" {
			uniqueRank = (uniqueRank * 2) / cardPower["1"]
		}
		return Response{Rank: 2, UniqueRank: uniqueRank, HandType: "Straight Flush", Hand: combinedStr}
	case isFourOfAKind(cards):
		return Response{Rank: 3, UniqueRank: uniqueRank, HandType: "Four Of A Kind", Hand: combinedStr}
	case isFullHouse(cards):
		return Response{Rank: 4, UniqueRank: uniqueRank, HandType: "Full House", Hand: combinedStr}
	case isFlush(cards):
		return Response{Rank: 5, UniqueRank: uniqueRank, HandType: "Flush", Hand: combinedStr}
	case isStraight(cards):
		if cards[0].Rank == "1" {
			uniqueRank = (uniqueRank * 2) / cardPower["1"]
		}
		return Response{Rank: 6, UniqueRank: uniqueRank, HandType: "Straight", Hand: combinedStr}
	case isThreeOfAKind(cards):
		return Response{Rank: 7, UniqueRank: uniqueRank, HandType: "Three Of A Kind", Hand: combinedStr}
	case isTwoPair(cards):
		return Response{Rank: 8, UniqueRank: uniqueRank, HandType: "Two Pair", Hand: combinedStr}
	case isPair(cards):
		return Response{Rank: 9, UniqueRank: uniqueRank, HandType: "One Pair", Hand: combinedStr}
	default:
		return Response{Rank: 10, UniqueRank: uniqueRank, HandType: "High Card", Hand: combinedStr} // High card
	}
}

// Implement functions to check different poker hands

func isStraightFlush(cards []Card) bool {
	// Check if the cards are both a straight and a flush
	return isFlush(cards) && isStraight(cards)
}

func isFourOfAKind(cards []Card) bool {
	// Check for four cards with the same rank
	return cards[0].Rank == cards[3].Rank || cards[1].Rank == cards[4].Rank
}

func isFullHouse(cards []Card) bool {
	// Check for three cards with the same rank and two cards with another rank
	return (cards[0].Rank == cards[2].Rank && cards[3].Rank == cards[4].Rank) ||
		(cards[0].Rank == cards[1].Rank && cards[2].Rank == cards[4].Rank)
}

func isFlush(cards []Card) bool {
	// Check if all cards have the same suit
	suit := cards[0].Suit
	for _, card := range cards {
		if card.Suit != suit {
			return false
		}
	}
	return true
}

func isRoyalFlush(cards []Card) bool {
	royalRanks := map[string]bool{
		"10": true,
		"11": true,
		"12": true,
		"13": true,
		"1":  true,
	}

	suit := cards[0].Suit
	for _, card := range cards {
		if card.Suit != suit || !royalRanks[card.Rank] {
			return false
		}
	}
	return true
}

func isStraight(cards []Card) bool {
	// Check if the ranks of the cards form a straight
	for i := 0; i < 4; i++ {
		num1, err1 := strconv.Atoi(cards[i+1].Rank)
		num2, err2 := strconv.Atoi(cards[i].Rank)
		if err1 != nil {
			panic("Invalid rank : " + cards[i+1].Rank)
		}
		if err2 != nil {
			panic("Invalid rank" + cards[i].Rank)
		}
		if num1-num2 != 1 {
			return checkRoyalStraight(cards)
		}
	}
	return true
}

func isThreeOfAKind(cards []Card) bool {
	// Check for three cards with the same rank
	return (cards[0].Rank == cards[2].Rank) || (cards[1].Rank == cards[3].Rank) || (cards[2].Rank == cards[4].Rank)
}

func isTwoPair(cards []Card) bool {
	// Check for two pairs of cards with the same rank
	return (cards[0].Rank == cards[1].Rank && cards[2].Rank == cards[3].Rank) ||
		(cards[0].Rank == cards[1].Rank && cards[3].Rank == cards[4].Rank) ||
		(cards[1].Rank == cards[2].Rank && cards[3].Rank == cards[4].Rank)
}

func isPair(cards []Card) bool {
	// Check for two cards with the same rank
	for i := 0; i < 4; i++ {
		if cards[i].Rank == cards[i+1].Rank {
			return true
		}
	}
	return false
}
func convertRank(card *Card) {
	if val, found := rankMapping[card.Rank]; found {
		card.Rank = val
	}
}
func convertRankBack(card *Card) {
	if val, found := rankMappingBack[card.Rank]; found {
		card.Rank = val
	}
}
func convertRanksInCards(cards []Card) {
	for i := range cards {
		convertRank(&cards[i])
	}
}
func covertRankBack(cards []Card) []Card {
	for i := range cards {
		convertRankBack(&cards[i])
	}
	return cards
}
func checkRoyalStraight(cards []Card) bool {
	check := ""
	for _, card := range cards {
		check += card.Rank
	}
	return check == "110111213"
}
func calculateMultiplication(cards []Card) int64 {
	multiplication := int64(1)
	for _, card := range cards {
		multiplication *= cardPower[card.Rank]
	}
	return multiplication
}
