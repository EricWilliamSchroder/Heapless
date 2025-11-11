package src


type Part struct{
	X, Y int // x and y coordinates of the snake **PART**
	Head *Snake // it knows the the head
	Tail *Snake // it knows the tail
}


type Snake struct {
	root *Part
	// score so forth in here
}

func AddToSnake(newPart *Part, snake **Snake){
	if newPart == nil {
		// nothing we can do if caller didn't provide a pointer to *Snake
		return
	}
	if *snake == nil {
		// initialize the Snake and set the root
		*snake = &Snake{root: newPart}
		return
	}
	addPart(newPart, &(*snake).root)
}

func addPart(newPart *Part, part **Part){
	
}

