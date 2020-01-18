package main

import (
	"math"
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

// as the lines are always vertical/horizontal, subtracting the x or y coordinate
func lineSegementLength(line lineSegment) int {
	xLength := abs(line.pointA.x - line.pointB.x)
	yLength := abs(line.pointA.y - line.pointB.y)

	if xLength > 0 {
		return xLength
	}

	return yLength
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

func lineSegmentPathLengthToIntersectionPoint(line []lineSegment, intersectionPoint point) int {
	pathSteps := 0
	for _, segment := range line {
		if onSegment(segment.pointA, intersectionPoint, segment.pointB) {
			newSegment := lineSegment{
				segment.pointA,
				intersectionPoint,
			}
			pathSteps += lineSegementLength(newSegment)
			break
		}

		// add distance from segement point a to b to pathSteps
		pathSteps += lineSegementLength(segment)
	}

	return pathSteps
}

func shortestPathToIntersection(wirePath string) int {
	wires := strings.Split(wirePath, "\n")

	wireASegments := constructLineSegmentsFromWire(strings.TrimSpace(wires[0]))
	wireBSegments := constructLineSegmentsFromWire(strings.TrimSpace(wires[1]))

	wireIntersections := findWireIntersectionPoints(wireASegments, wireBSegments)

	closestIntersectionDistance := math.MaxInt32 // hacky, whatever...

	for _, intersection := range wireIntersections {
		if intersection.x == 0 && intersection.y == 0 {
			continue
		}

		wireAPathLengthToIntersection := lineSegmentPathLengthToIntersectionPoint(wireASegments, intersection)
		wireBPathLengthToIntersection := lineSegmentPathLengthToIntersectionPoint(wireBSegments, intersection)

		pathDistanceToIntersection := wireAPathLengthToIntersection + wireBPathLengthToIntersection

		if closestIntersectionDistance > pathDistanceToIntersection {
			closestIntersectionDistance = pathDistanceToIntersection
		}
	}

	return closestIntersectionDistance
}

func closestIntersection(wirePath string) int {
	wires := strings.Split(wirePath, "\n")

	wireASegments := constructLineSegmentsFromWire(strings.TrimSpace(wires[0]))
	wireBSegments := constructLineSegmentsFromWire(strings.TrimSpace(wires[1]))

	wireIntersections := findWireIntersectionPoints(wireASegments, wireBSegments)

	closestIntersectionDistance := math.MaxInt32 // hacky, whatever...

	for _, intersection := range wireIntersections {
		if intersection.x == 0 && intersection.y == 0 {
			continue
		}
		manhattanDistance := abs(intersection.x) + abs(intersection.y)

		if closestIntersectionDistance > manhattanDistance {
			closestIntersectionDistance = manhattanDistance
		}
	}

	return closestIntersectionDistance
}

func findWireIntersectionPoints(wireA, wireB []lineSegment) []point {
	intersectionPoints := []point{}
	for _, segmentA := range wireA {
		for _, segmentB := range wireB {
			if doLineSegmentsIntersect(segmentA, segmentB) {
				intersectionPoints = append(intersectionPoints, lineLineIntersectionPoint(segmentA, segmentB))
			}
		}
	}

	return intersectionPoints
}

func lineLineIntersectionPoint(a, b lineSegment) point {
	a1 := a.pointB.y - a.pointA.y
	b1 := a.pointA.x - a.pointB.x
	c1 := a1*a.pointA.x + b1*a.pointA.y

	a2 := b.pointB.y - b.pointA.y
	b2 := b.pointA.x - b.pointB.x
	c2 := a2*b.pointA.x + b2*b.pointA.y

	determinant := a1*b2 - a2*b1

	if determinant == 0 {
		return point{0, 0}
	}

	x := (b2*c1 - b1*c2) / determinant
	y := (a1*c2 - a2*c1) / determinant

	return point{x, y}
}

func constructLineSegmentsFromWire(wire string) []lineSegment {
	lineSegments := []lineSegment{}
	wireVectors := strings.Split(wire, ",")

	previousPoint := point{0, 0}

	for _, vector := range wireVectors {
		newSegment := lineSegment{}
		newSegment.pointA = previousPoint

		direction := string(vector[0])
		distance, err := strconv.Atoi(string(vector[1:]))

		if err != nil {
			return []lineSegment{}
		}

		switch direction {
		case "U":
			newPoint := point{previousPoint.x, previousPoint.y + distance}
			lineSegments = append(lineSegments, lineSegment{
				pointA: previousPoint,
				pointB: newPoint,
			})
			previousPoint = newPoint
		case "R":
			newPoint := point{previousPoint.x + distance, previousPoint.y}
			lineSegments = append(lineSegments, lineSegment{
				pointA: previousPoint,
				pointB: newPoint,
			})
			previousPoint = newPoint
		case "D":
			newPoint := point{previousPoint.x, previousPoint.y - distance}
			lineSegments = append(lineSegments, lineSegment{
				pointA: previousPoint,
				pointB: newPoint,
			})
			previousPoint = newPoint
		case "L":
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

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// 0 for colinear
// 1 for clockwise
// 2 for counterclockwise
func threePointOrientation(a, b, c point) int {
	val := (b.y-a.y)*(c.x-b.x) - (b.x-a.x)*(c.y-b.y)

	if val == 0 {
		return 0
	}

	if val > 0 {
		return 1
	}

	return 2
}

// does b fall on segment a, c
func onSegment(a, b, c point) bool {
	if b.x <= max(a.x, c.x) && b.x >= min(a.x, c.x) && b.y <= max(a.y, c.y) && b.y >= min(a.y, c.y) {
		return true
	}
	return false
}

func doLineSegmentsIntersect(a, b lineSegment) bool {
	o1 := threePointOrientation(a.pointA, a.pointB, b.pointA)
	o2 := threePointOrientation(a.pointA, a.pointB, b.pointB)
	o3 := threePointOrientation(b.pointA, b.pointB, a.pointA)
	o4 := threePointOrientation(b.pointA, b.pointB, a.pointB)

	if o1 != o2 && o3 != o4 {
		return true
	}

	if o1 == 0 && onSegment(a.pointA, a.pointB, b.pointA) {
		return true
	}

	if o2 == 0 && onSegment(a.pointA, a.pointB, b.pointA) {
		return true
	}

	if o3 == 0 && onSegment(a.pointB, a.pointA, b.pointB) {
		return true
	}

	if o4 == 0 && onSegment(a.pointB, a.pointA, b.pointB) {
		return true
	}

	return false
}

func main() {

}

// for each line, form coordinate pairs for each line segement
// loop over each line segement, compare to all the other lines segements to check for intersections
// calculate shortest manhattan distance
