package src

import "time"

const playingFieldSize = 10

const Size int = playingFieldSize + 2 // because the borders remove 2
const MaxSnakeLength = Size * Size
const XOffset int = 20
const YOffset int = 2


var tickSpeed = 200 * time.Millisecond
var Fragments [MaxSnakeLength]Fragment
var GameBoard Board = CreateBoard()
var IsGameOver = false