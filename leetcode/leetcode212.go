package main

import "fmt"

type TireNode struct {
	Val   byte
	Nodes map[byte]*TireNode
	Word  string
}

type Question struct {
	MatchWords map[string]bool
}

func (q *Question) backtracking(board [][]byte, row int, col int, node *TireNode) {
	letter := board[row][col]
	curNode := node.Nodes[letter]

	if curNode.Word != "" {
		q.MatchWords[curNode.Word] = true
	}

	// visited
	board[row][col] = '#'

	neighbors := [][]int{
		[]int{0, 1},
		[]int{1, 0},
		[]int{0, -1},
		[]int{-1, 0},
	}

	for _, offset := range neighbors {
		newRow := row + offset[0]
		newCol := col + offset[1]
		if newRow < 0 || newCol < 0 || newRow >= len(board) || newCol >= len(board[0]) {
			continue
		}
		c := board[newRow][newCol]
		if _, ok := curNode.Nodes[c]; !ok {
			continue
		}
		q.backtracking(board, newRow, newCol, curNode)
	}

	board[row][col] = letter
}

func findWords(board [][]byte, words []string) []string {
	// 构造前缀树
	tree := &TireNode{
		Val:   byte('$'),
		Nodes: map[byte]*TireNode{},
		Word:  "",
	}
	for _, word := range words {
		curNode := tree
		for _, c := range []byte(word) {
			if _, ok := curNode.Nodes[c]; ok {
				curNode = curNode.Nodes[c]
			} else {
				newNode := &TireNode{
					Val:   c,
					Nodes: map[byte]*TireNode{},
					Word:  "",
				}
				curNode.Nodes[c] = newNode
				curNode = newNode
			}
		}
		curNode.Word = word
	}

	q := Question{
		MatchWords: map[string]bool{},
	}

	for row, line := range board {
		for col, c := range line {
			if _, ok := tree.Nodes[c]; ok {
				q.backtracking(board, row, col, tree)
			}
		}
	}

	res := []string{}
	for k, _ := range q.MatchWords {
		res = append(res, k)
	}

	return res
}

func runLeetcode212() {
	board := [][]byte{
		[]byte{'o', 'a', 'a', 'n'},
		[]byte{'e', 't', 'a', 'e'},
		[]byte{'i', 'h', 'k', 'r'},
		[]byte{'i', 'f', 'l', 'v'},
	}
	words := []string{
		"eat", "oath",
	}
	matched := findWords(board, words)
	fmt.Println(matched)
}
