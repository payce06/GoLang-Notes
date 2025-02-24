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

// 2. Count the number of properties in an object
func countProperties(obj map[string]interface{}) int {
	return len(obj)
}

// 3. Get an array of values for a specific property in an array of objects
func getPropertyValues(arr []map[string]interface{}, prop string) []interface{} {
	var values []interface{}
	for _, obj := range arr {
		if value, ok := obj[prop]; ok {
			values = append(values, value)
		}
	}
	return values
}

// 4. Check if an object contains a specific property
func containsProperty(obj map[string]interface{}, prop string) bool {
	_, exists := obj[prop]
	return exists
}

// 5. Deep clone an object
func deepClone(obj map[string]interface{}) map[string]interface{} {
	clone := make(map[string]interface{})
	for key, value := range obj {
		clone[key] = value
	}
	return clone
}

// 6. Get the keys of an object
func getKeys(obj map[string]interface{}) []string {
	var keys []string
	for key := range obj {
		keys = append(keys, key)
	}
	return keys
}