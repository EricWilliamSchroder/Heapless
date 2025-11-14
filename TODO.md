# TODO

1. change the board to a struct instead of pritning directly thus allowing for better hit detection

2. Instead of hard coding the bytes for the pieces, add the byte code to the Part struct and have each piece carry its own byte code, thus allowing for easier addition of new pieces in the future, do this also for the board representation