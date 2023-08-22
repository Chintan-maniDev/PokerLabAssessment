# Poker Hand Evaluator Algorithm:

## Input Parsing:
  - The algorithm begins by parsing the input string containing comma-separated card representations.
  - Each card representation is split into its rank and suit components.
  - The algorithm handles different card formats (e.g., "10" as well as single-character ranks).
  - Sorting and Rank Conversion:
  - The parsed cards are then sorted based on their rank. This sorting is essential for checking poker hand combinations.
  - The ranks of the cards are converted from human-readable format (e.g., "A" for Ace) to internal numeric format (e.g., "1" for Ace).
## Checking for Poker Hands:
  - The algorithm sequentially checks for different poker hand combinations in decreasing order of rank.
  - It begins by checking for a royal flush, followed by a straight flush, four of a kind, full house, flush, straight, three of a kind, two pair, and finally a pair.
  - If none of these combinations are found, the algorithm determines the hand as a high card.
### Royal Flush Check:
  - The algorithm verifies if the cards have the same suit and rank combinations of 10, J, Q, K, and A.
  - If satisfied, it confirms a royal flush.
### Straight Flush Check:
  - The algorithm confirms a straight flush if the cards have both a straight and a flush.
  - It ensures the ranks form a consecutive sequence and that all cards share the same suit.
### Four of a Kind Check:
  - The algorithm identifies four cards with the same rank, indicating a four of a kind hand.
### Full House Check:
  - For a full house, the algorithm verifies three cards of one rank and two cards of another rank.
### Flush Check:
  - A flush is confirmed if all cards have the same suit.
### Straight Check:
  - The algorithm checks if the cards form a straight by ensuring their ranks are in a consecutive sequence.
### Three of a Kind Check:
  - A three of a kind hand has three cards of the same rank.
### Two Pair Check:
  - Two pair consists of two cards of one rank and two cards of another rank.
### One Pair Check:
  - A one pair hand has two cards of the same rank.
## Unique Rank Calculation:
  - For each poker hand combination, the algorithm calculates a unique rank by multiplying the powers of card ranks. This unique rank helps resolve ties between hands of the same type. To avoid collision, assigned prime numbers to each rank.
##  Response Formation:
  - The algorithm creates a response containing the hand type, rank, unique rank, and the combined string representation of the cards.
## Overall Complexity:
  - The algorithm processes each card only once and performs sorting, checking for poker hands, and calculating the unique rank. Each step operates in linear time with respect to the number of cards. Thus, the overall time complexity is O(N), where N is the number of cards in the hand.
## Conclusion:
  - The poker hand evaluator algorithm is well-structured, efficient, and capable of identifying different poker hand combinations accurately. It employs a systematic approach to parsing, sorting, and evaluating hands, ensuring accurate results for a variety of input scenarios.
