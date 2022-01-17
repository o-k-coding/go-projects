package maximize_person_distance

// For each "seat" in seats, a 1 represnets a person seated at seat i. A 0 represents an empty seat
// Alex wants to sit in a seat where he is the max distance from the next closest person
func MaximizePersonDistance(seats []int) int {
	maxDistance := 0
	lastOccupiedSeat := -1

	for i, seat := range seats {
		if seat == 1 {
			if lastOccupiedSeat != -1 {
				distanceBetween := (i - lastOccupiedSeat) / 2
				if distanceBetween > maxDistance {
					maxDistance = distanceBetween
				}
			} else {
				// No seat has been seen yet
				maxDistance = i
			}
			lastOccupiedSeat = i
		}
	}

	// If the last seat is not occupied, check the distance to the last occupied seat
	lastSeat := len(seats) - 1
	if seats[lastSeat] == 0 && (lastSeat - lastOccupiedSeat > maxDistance  ){
		maxDistance = lastSeat - lastOccupiedSeat
	}

	return maxDistance
}


// [1,0,0,0,1,0,1]
// [0,1,2,3,4,5,6]
// Multiple items,
//4 - 0, 6 - 4, max is 4 / 2 = 2 + the left index == index 2

// [1,0,0,0]
// [0,1,2,3]
// only 1 item, so find distance to each side and pick the larger
// 0 - 0 == 0, 3 - 0 = 3

// [1,0,0,1,0,0,1]
// [0,1,2,3,4,5,6]
// Multiple items,
//3 - 0, 6 - 3, max is 3 / 2 = floor(1.5) + the left index == index 2
