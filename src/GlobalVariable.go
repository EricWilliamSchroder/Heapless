package src

const Size int = 30
const MaxSnakeLength = Size * Size
const XOffset int = 20
const YOffset int = 2


var Fragments [MaxSnakeLength]Fragment
var GameBoard Board = CreateBoard()
var IsGameOver = false