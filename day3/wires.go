package main

import (
	"fmt"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func (a point) Equal(b point) bool {
	if &a == &b {
		return true
	}

	if a.x != b.x || a.y != b.y {
		return false
	}

	return true
}

type lineSegment struct {
	pointA, pointB point
}

func (a lineSegment) findIntersection(b lineSegment) (point, error) {
	return point{0, 0}, nil
}

func (a lineSegment) Equal(b lineSegment) bool {
	if &a == &b {
		return true
	}

	if a.pointA != b.pointA || a.pointB != b.pointB {
		return false
	}

	return true
}

func closestIntersection(wirePath string) int {
	wires := strings.Split(wirePath, "\n")

	wireASegments := constructLineSegmentsFromWire(wires[0])
	wireBSegments := constructLineSegmentsFromWire(wires[1])

	fmt.Println(wireASegments, wireBSegments)

	return 0
}

func findWireIntersections(wireA, wireB []lineSegment) {
	for _, segmentA := range wireA {
		for _, segmentB := range wireB {

		}
	}
}

func constructLineSegmentsFromWire(wire string) []lineSegment {
	lineSegments := []lineSegment{}
	wireVectors := strings.Split(wire, ",")

	previousPoint := point{0, 0}

	for _, vector := range wireVectors {
		newSegment := lineSegment{}
		newSegment.pointA = previousPoint

		direction := vector[0]
		distance, err := strconv.Atoi(vector[0:])
		if err != nil {
			return []lineSegment{}
		}
		fmt.Printf("direction: '%v', distance: '%v'\n", direction, distance)
		switch direction {
		case 'U':
			newPoint := point{previousPoint.x, previousPoint.y + distance}
			lineSegments = append(lineSegments, lineSegment{
				pointA: previousPoint,
				pointB: newPoint,
			})
			previousPoint = newPoint
		case 'R':
			newPoint := point{previousPoint.x + distance, previousPoint.y}
			lineSegments = append(lineSegments, lineSegment{
				pointA: previousPoint,
				pointB: newPoint,
			})
			previousPoint = newPoint
		case 'D':
			newPoint := point{previousPoint.x, previousPoint.y - distance}
			lineSegments = append(lineSegments, lineSegment{
				pointA: previousPoint,
				pointB: newPoint,
			})
			previousPoint = newPoint
		case 'L':
			newPoint := point{previousPoint.x - distance, previousPoint.y}
			lineSegments = append(lineSegments, lineSegment{
				pointA: previousPoint,
				pointB: newPoint,
			})
			previousPoint = newPoint
		}
	}

	return lineSegments
}

func main() {

}

// for each line, form coordinate pairs for each line segement
// loop over each line segement, compare to all the other lines segements to check for intersections
// calculate shortest manhattan distance
