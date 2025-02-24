package main

import (
	"fmt"
	"strings"
)

// 1. Merge two objects into one
func mergeObjects(obj1, obj2 map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})
	for key, value := range obj1 {
		merged[key] = value
	}
	for key, value := range obj2 {
		merged[key] = value
	}
	return merged
}