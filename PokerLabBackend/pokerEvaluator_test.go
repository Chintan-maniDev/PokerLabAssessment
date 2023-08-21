package main

import "testing"

// Test Cases for basic Scenarios
func TestFindBestPokerHandRank_BasicScenarios(t *testing.T) {
	testCases := []struct {
		cards    []Card
		expected Response
	}{
		// Test case 1: Royal Flush
		{
			cards: []Card{
				{Rank: "1", Suit: "S"},
				{Rank: "13", Suit: "S"},
				{Rank: "12", Suit: "S"},
				{Rank: "11", Suit: "S"},
				{Rank: "10", Suit: "S"},
			},
			expected: Response{HandType: "Royal Flush", Rank: 1, UniqueRank: 58642669, Hand: ""},
		},
		// Test case 2: Straight Flush
		{
			cards: []Card{
				{Rank: "5", Suit: "H"},
				{Rank: "6", Suit: "H"},
				{Rank: "7", Suit: "H"},
				{Rank: "8", Suit: "H"},
				{Rank: "9", Suit: "H"},
			},
			expected: Response{HandType: "Straight Flush", Rank: 2, UniqueRank: 1062347, Hand: ""},
		},
		// Test case 3: Four Of A Kind
		{
			cards: []Card{
				{Rank: "9", Suit: "C"},
				{Rank: "9", Suit: "D"},
				{Rank: "9", Suit: "H"},
				{Rank: "9", Suit: "S"},
				{Rank: "13", Suit: "D"},
			},
			expected: Response{HandType: "Four Of A Kind", Rank: 3, UniqueRank: 11473481, Hand: ""},
		},
		// ... Add more test cases for other hands ...

		// Test case 10: High Card
		{
			cards: []Card{
				{Rank: "2", Suit: "C"},
				{Rank: "4", Suit: "H"},
				{Rank: "6", Suit: "D"},
				{Rank: "8", Suit: "S"},
				{Rank: "10", Suit: "S"},
			},
			expected: Response{HandType: "High Card", Rank: 10, UniqueRank: 150423, Hand: ""},
		},
	}

	for _, tc := range testCases {
		actualOutput := FindBestPokerHandRank(tc.cards, "")

		if actualOutput != tc.expected {
			t.Errorf("For cards %v, expected %v, but got %v", tc.cards, tc.expected, actualOutput)
		}
	}
}

// Test Cases for comparing hands
func TestFindBestPokerHandRank_CompareUniqueRank_Straight(t *testing.T) {
	InputHand1 := "2S,3D,4C,5C,6S"
	InputHand2 := "5S,6H,7D,8S,9S"
	expectedWinner := InputHand2

	testCases := []struct {
		input1   string
		input2   string
		expected string
	}{
		{
			input1:   InputHand1,
			input2:   InputHand2,
			expected: expectedWinner,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input1, func(t *testing.T) {
			hand1 := parseInput(testCase.input1)
			hand2 := parseInput(testCase.input2)

			actual1 := FindBestPokerHandRank(hand1, "")
			actual2 := FindBestPokerHandRank(hand2, "")

			if actual1.UniqueRank < actual2.UniqueRank {
				if testCase.expected != InputHand2 {
					t.Errorf("expected %s to win, but got %s", testCase.expected, InputHand2)
				}
			} else if actual1.UniqueRank > actual2.UniqueRank {
				if testCase.expected != InputHand1 {
					t.Errorf("expected %s to win, but got %s", testCase.expected, InputHand1)
				}
			} else {
				t.Errorf("hands have the same unique rank: %d", actual1.UniqueRank)
			}
		})
	}
}

func TestFindBestPokerHandRank_CompareUniqueRank_SameCardsDifferentSuits(t *testing.T) {
	InputHand1 := "AS,QS,KS,JS,TS"
	InputHand2 := "AH,QH,KH,JH,TH"
	expectedWinner := "Tie"

	testCases := []struct {
		input1   string
		input2   string
		expected string
	}{
		{
			input1:   InputHand1,
			input2:   InputHand2,
			expected: expectedWinner,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input1, func(t *testing.T) {
			hand1 := parseInput(testCase.input1)
			hand2 := parseInput(testCase.input2)

			actual1 := FindBestPokerHandRank(hand1, "")
			actual2 := FindBestPokerHandRank(hand2, "")

			if actual1.UniqueRank == actual2.UniqueRank {
				if testCase.expected != "Tie" {
					t.Errorf("expected a tie, but got %s", testCase.expected)
				}
			} else {
				t.Errorf("expected a tie, but got different unique ranks")
			}
		})
	}
}

func TestFindBestPokerHandRank_CompareUniqueRank_Flush(t *testing.T) {
	InputHand1 := "2S,4S,6S,8S,9S"
	InputHand2 := "1S,KS,8S,7S,2S"
	expectedWinner := InputHand2

	testCases := []struct {
		input1   string
		input2   string
		expected string
	}{
		{
			input1:   InputHand1,
			input2:   InputHand2,
			expected: expectedWinner,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input1, func(t *testing.T) {
			hand11 := parseInput(testCase.input1)
			hand22 := parseInput(testCase.input2)

			actual1 := FindBestPokerHandRank(hand11, "")
			actual2 := FindBestPokerHandRank(hand22, "")

			if actual1.UniqueRank < actual2.UniqueRank {
				if testCase.expected != InputHand2 {
					t.Errorf("expected %s to win, but got %s", testCase.expected, InputHand2)
				}
			} else if actual1.UniqueRank > actual2.UniqueRank {
				if testCase.expected != InputHand1 {
					t.Errorf("expected %s to win, but got %s", testCase.expected, InputHand1)
				}
			} else {
				t.Errorf("hands have the same unique rank: %d", actual1.UniqueRank)
			}
		})
	}
}
