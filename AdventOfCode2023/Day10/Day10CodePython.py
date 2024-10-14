# -*- coding: utf-8 -*-
"""
Created on Mon Jan  8 02:44:24 2024

@author: mdude
"""
def getPipes():
    with open("pipes.txt") as file:
        return file.read()

def spacePipes(pipes):
    """This function adds empty space between the pipes to make it so that we can recurse between pipes that 
    are not connected"""
    #Put spaces between characters so we can recurse between pipes
    #Puts filler spaces between rows
    spacedPipes = []
    for i, row in enumerate(pipes):
        spacedPipes.append([])

        #Calculate the length of the row to avoid extra filler spaces
        if len(row)*2 % 2 == 0:
            rowLength = (len(row) * 2) - 1
        else:
            rowLength = len(row) * 2

        for j in range(rowLength):
            if j % 2 == 0:
                spacedPipes[-1].append(pipes[i][j//2])
            else:
                if pipes[i][j//2] in ["-", "F", "L"] and pipes[i][(j//2) + 1] in ["-", "J", "7"]:
                    spacedPipes[-1].append("-")
                else:
                    spacedPipes[-1].append("#")

    tempSpacedPipes = spacedPipes
    spacedPipes = []
    for i in range(len(tempSpacedPipes) * 2):

        spacedPipes.append([])

        for j in range(len(tempSpacedPipes[i//2])):
            if i % 2 == 0:
                spacedPipes[-1].append(tempSpacedPipes[i//2][j])
            else:
                if tempSpacedPipes[i//2][j] in ["|", "F", "7"] and tempSpacedPipes[(i//2) + 1][j] in ["|", "J", "L"]:
                    spacedPipes[-1].append("|")
                else:
                    spacedPipes[-1].append("#")

    #Gets rid of additional row spacing at the end
    if len(set(spacedPipes[-1])) == 1:
        spacedPipes = spacedPipes[:-1]
    
    return spacedPipes

def isValidCoordinate(pipes, xCoordinate, yCoordinate):
    """Returns False if x or y coordinate is out of bounds otherwise return True"""
    if xCoordinate < 0 or xCoordinate > len(pipes)-1:
        return False
    if yCoordinate < 0 or yCoordinate > len(pipes[0])-1:
        return False
    return True

def traverseBoard(pipes, xCoordinate, yCoordinate):
    #Mark current point as an outside point
    pipes[xCoordinate][yCoordinate] = "O"
    
    #Checks to make sure the coordinate above the current coordinate is in bounds
    if isValidCoordinate(pipes, xCoordinate - 1, yCoordinate):
        #If the pipe is marked as either a filler space, or an inside point, call this function on that point
        if pipes[xCoordinate - 1][yCoordinate] in ["#", "I"]:
            pipes = traverseBoard(pipes, xCoordinate - 1, yCoordinate)
    
    #Checks to make sure the coordinate below the current coordinate is in bounds
    if isValidCoordinate(pipes, xCoordinate + 1, yCoordinate):
        #If the pipe is marked as either a filler space, or an inside point, call this function on that point
        if pipes[xCoordinate + 1][yCoordinate] in ["#", "I"]:
            pipes = traverseBoard(pipes, xCoordinate + 1, yCoordinate)
    
    #Checks to make sure the coordinate to the left of the current coordinate is in bounds
    if isValidCoordinate(pipes, xCoordinate, yCoordinate - 1):
        #If the pipe is marked as either a filler space, or an inside point, call this function on that point
        if pipes[xCoordinate][yCoordinate - 1] in ["#", "I"]:
            pipes = traverseBoard(pipes, xCoordinate, yCoordinate - 1)
    
    #Checks to make sure the coordinate to the right of the current coordinate is in bounds
    if isValidCoordinate(pipes, xCoordinate, yCoordinate + 1):
        #If the pipe is marked as either a filler space, or an inside point, call this function on that point
        if pipes[xCoordinate][yCoordinate + 1] in ["#", "I"]:
            pipes = traverseBoard(pipes, xCoordinate, yCoordinate + 1)
            
    return pipes
    

def printPipes(pipes):
    """Prints pipes, pretty much just used for testing purposes"""
    for i in range(len(pipes)):
        for j in range(len(pipes[i])):
            print(pipes[i][j], end="")
        print()

#Formatting
pipes = getPipes()
#TEST INPUTS
"""..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
.........."""
"""......................
..F----7F7F7F7F-7.....
..|F--7||||||||FJ.....
..||.FJ||||||||L7.....
.FJL7L7LJLJ||LJ.L-7...
.L--J.L7...LJS7F-7L7..
.....F-J..F7FJ|L7L7L7.
.....L7.F7||L7|.L7L7|.
......|FJLJ|FJ|F7|.LJ.
.....FJL-7.||.||||....
.....L---J.LJ.LJLJ....
......................"""
"""......................
.FF7FSF7F7F7F7F7F---7.
.L|LJ||||||||||||F--J.
.FL-7LJLJ||||||LJL-77.
.F--JF--7||LJLJIF7FJ-.
.L---JF-JLJIIIIFJLJJ7.
.|F|F-JF---7IIIL7L|7|.
.|FFJF7L7F-JF7IIL---7.
.7-L-JL7||F7|L7F-7F7|.
.L.L7LFJ|||||FJL7||LJ.
.L7JLJL-JLJLJL--JLJ.L.
......................"""
pipes = pipes.split("\n")
pipes = [list(x) for x in pipes]

#Find 'S' as the starting location:
for i, row in enumerate(pipes):
    for j, char in enumerate(row):
        if char == "S":
            sCoordinate = [i, j] #Keep track of the start Coordinate
            currentCoords = [i, j] #Keeps track of current coordinates
            break

#Find the next move based on the location of the start ("S")
#Check for valid path up (x coordinate - 1)
if pipes[currentCoords[0] - 1][currentCoords[1]] in ['|', 'F', '7']:
    currentCoords[0] -= 1
    previousDirection = "Up"
#Check for valid path down (x coordinate - 1) if there is no valid path up
elif pipes[currentCoords[0] + 1][currentCoords[1]] in ['|', 'L', 'J']:
    currentCoords[0] += 1
    previousDirection = "Down"
#The start has to be connected by exactly two sides, so if it's not up or down, it has to be left and right
else:
    pipes[1] -= 1
    previousDirection = "Left"
    
pathCoords = [[currentCoords[0], currentCoords[1]]]

#Loop through until we get back to "S"
while pipes[currentCoords[0]][currentCoords[1]] != "S":
    
    #Keep track of the current character
    currentChar = pipes[currentCoords[0]][currentCoords[1]]
    
    if currentChar == "|":
        #If the previous direction was up, that means it came from the bottom, so we need to move up
        if previousDirection == "Up":
            currentCoords[0] -= 1
            previousDirection = "Up"
        #This means if came from above, so we have to move down
        else:
            currentCoords[0] += 1
            previousDirection = "Down"
            
    elif currentChar == "-":
        #If we came from the left, that means we continue to move left
        if previousDirection == "Left":
            currentCoords[1] -= 1
            previousDirection = "Left"
        #This means if came from the right, so we have to keep moving to the right
        else:
            currentCoords[1] += 1
            previousDirection = "Right"
            
    elif currentChar == "L":
        #If coming from the left, move up
        if previousDirection == "Left":
            currentCoords[0] -= 1
            previousDirection = "Up"
        #If coming from up, move to the right
        else:
            currentCoords[1] += 1
            previousDirection = "Right"
            
    elif currentChar == "J":
        #If coming from right, move up
        if previousDirection == "Right":
            currentCoords[0] -= 1
            previousDirection = "Up"
        #If coming from up move to the left
        else:
            currentCoords[1] -= 1
            previousDirection = "Left"
            
    elif currentChar == "7":
        #If coming from the bottom, move left
        if previousDirection == "Up":
            currentCoords[1] -= 1
            previousDirection = "Left"
        #If coming from the right, move down
        else:
            currentCoords[0] += 1
            previousDirection = "Down"
            
    elif currentChar == "F":
        #If coming from the bottom, move right
        if previousDirection == "Up":
            currentCoords[1] += 1
            previousDirection = "Right"
        #If coming from right, move down
        else:
            currentCoords[0] += 1
            previousDirection = "Down"
    
    pathCoords.append([currentCoords[0], currentCoords[1]])
        
#Replace non path characters with "I"
for i, row in enumerate(pipes):
    for j in range(len(row)):
        if [i,j] not in pathCoords:
            pipes[i][j] = "I"
            
#Replace S with correct character
#This assumes S has to be not a '|' or a '-'
connectedBottom = pipes[sCoordinate[0] + 1][sCoordinate[1]] in ['|', 'L', 'J']
connectedTop = pipes[sCoordinate[0] - 1][sCoordinate[1]] in ['|', 'F', '7']
connectedRight = pipes[sCoordinate[0]][sCoordinate[1] + 1] in ['-', 'J', '7']
connectedLeft = pipes[sCoordinate[0]][sCoordinate[1] - 1] in ['-', 'F', 'L']
#Checks for F
if connectedRight and connectedBottom:
    pipes[sCoordinate[0]][sCoordinate[1]] = "F"
#Checks for L
if connectedRight and connectedTop:
    pipes[sCoordinate[0]][sCoordinate[1]] = "L"
#Checks for 7
if connectedLeft and connectedBottom:
    pipes[sCoordinate[0]][sCoordinate[1]] = "7"
#Checks for J
if connectedLeft and connectedTop:
    pipes[sCoordinate[0]][sCoordinate[1]] = "J"
    
spacedPipes = spacePipes(pipes)
spacedPipes = traverseBoard(spacedPipes, 0, 0)

#Number of points inside the the loop
numInsidePoints = 0
for row in spacedPipes:
    numInsidePoints += row.count("I")

print(f"Number of points inside the loop: {numInsidePoints}")
