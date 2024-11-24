Implements the idea from https://zerowidth.com/2013/a-visual-explanation-of-jump-point-search/

# Pseudo Code
```
Node {
    Position          // X, Y coordinates
    G, H, F           // Cost metrics: G (from start), H (heuristic to end), F (total cost)
    ParentDir         // Direction from parent node
    ForcedNeighborDir // Direction of a forced neighbor (if any)
    Parent            // Reference to parent node
}

findPathWithJPS(StartNode, EndNode, Matrix)
    OpenList = createHeap()          // Binary heap, pops node with lowest F value with log(N) speed
    OpenList.push(StartNode)         // Add StartNode to OpenList

    while OpenList is not empty
        CurrentNode = pop(OpenList)  // Get node with the lowest F value
        if CurrentNode.Position == EndNode.Position
            return reconstructPath(CurrentNode) // Build path by traversing Parent nodes

        Successors = findSuccessors(CurrentNode, EndNode, Matrix)
        for each Successor in Successors
            OpenList.push(Successor)

reconstructPath(Node)
    ResultList = []
    while Node != nil
        add Node to ResultList
        Node = Node.Parent
    reverse(ResultList)
    return ResultList

findSuccessors(CurrentNode, EndNode, Matrix)
    Successors = []
    X, Y = CurrentNode.Position

    for each Dir in getDirections(CurrentNode)  // Note: can be run in parallel
        DX, DY = Dir
        SX, SY, ForcedNeighborDir, found = jump(X+DX, Y+DY, DX, DY, EndNode, Matrix)
        if found
            Successor = new Node()
            Successor.Position = {SX, SY}
            Successor.G = CurrentNode.G + heuristic(CurrentNode.Position, Successor.Position)
            Successor.H = heuristic(Successor.Position, EndNode.Position)
            Successor.F = Successor.G + Successor.H
            Successor.ParentDir = Dir
            Successor.ForcedNeighborDir = ForcedNeighborDir
            Successor.Parent = CurrentNode
            add Successor to Successors
    return Successors

getDirections(Node)
    if Node.ParentDir == nil  // Starting node
        return [{0, -1}, {0, 1}, {-1, 0}, {1, 0}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}] // All 8 directions

    PDX, PDY = Node.ParentDir
    FNDX, FNDY = Node.ForcedNeighborDir

    if PDX != 0 and PDY != 0  // Diagonal movement
        return [{PDX, PDY}, {PDX, 0}, {0, PDY}, if ForcedNeighborDir exists then also {FNDX, FNDY}]
    else  // Horizontal or vertical movement
        return [{PDX, PDY}, if ForcedNeighborDir exists then also {FNDX, FNDY}]

jump(X, Y, DX, DY, EndNode, Matrix) returns X, Y, ForcedNeighnorDir, found
    while isWalkable(X, Y, Matrix)
        if {X, Y} == EndNode.Position
            return X, Y, nil, true

        // Diagonal movement: Check forced neighbors
        if DX != 0 and DY != 0
            // Try to find the forced neighbor pattern for diagonal move as described in the article 
            if !isWalkable(X, Y+DY*-1, Matrix) and isWalkable(X+DX, Y+DY*-1, Matrix)
                return X, Y, {DX, DY*-1}, true
            if !isWalkable(X+DX*-1, Y Matrix) and isWalkable(X+DX*-1, Y+DY, Matrix)
                return X, Y, {DX*-1, DY}, true

            // Recursive jump for horizontal and vertical continuations
            _, _, _, found1 = jump(X + DX, Y, DX, 0, EndNode, Matrix) // Only found is of interest
            _, _, _, found2 = jump(X, Y + DY, 0, DY, EndNode, Matrix)
            if (found1 == true) or (found2 == true)
                return X, Y, nil, true

        // Horizontal movement: Check forced neighbors
        else if DX != 0
            if !isWalkable(X, Y - 1, Matrix) and isWalkable(X + DX, Y - 1, Matrix)
                return X, Y, {DX, -1}, true
            if !isWalkable(X, Y + 1, Matrix) and isWalkable(X + DX, Y + 1, Matrix)
                return X, Y, {DX, 1}, true

        // Vertical movement: Check forced neighbors
        else if DY != 0
            if !isWalkable(X - 1, Y, Matrix) and isWalkable(X - 1, Y + DY, Matrix)
                return X, Y, {-1, DY}, true
            if !isWalkable(X + 1, Y, Matrix) and isWalkable(X + 1, Y + DY, Matrix)
                return X, Y, {1, DY}, true

        // Continue in the same direction
        X += DX
        Y += DY

    return 0, 0, nil, false


isWalkable(X, Y, Matrix)
	return X >= 0 and Y >= 0 and x < Matrix.size and y < Matrix[0].size and Matrix[x][y]


func heuristic(A, B, Matrix) returns int 
	// Using Manhattan distance as heuristic
	return math.Abs(A.X-B.X) + math.Abs(A.Y-B.Y)
```
