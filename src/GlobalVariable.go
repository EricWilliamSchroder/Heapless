package src

import "time"

const Size int = 30
const MaxSnakeLength = Size * Size
const XOffset int = 20
const YOffset int = 2
const UpdateSpeed = 500 * time.Millisecond // move every 150ms


var Fragments [MaxSnakeLength]Fragment
var GameBoard Board = CreateBoard()
var IsGameOver = false