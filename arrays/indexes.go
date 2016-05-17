package array_utils

import "gopkg.in/mgo.v2/bson"

// StringIndex Finds the index of the specified element.
func StringIndex(s string, a []string) int {
	for i, v := range a {
		if s == v {
			return i
		}
	}
	return -1
}

// IntIndex Finds the index of the specified element.
func IntIndex(n int, a []int) int {
	for i, v := range a {
		if n == v {
			return i
		}
	}
	return -1
}

// BsonIndex Finds the index of the specified element.
func BsonIndex(o bson.ObjectId, a []bson.ObjectId) int {
	for i, v := range a {
		if o == v {
			return i
		}
	}
	return -1
}
