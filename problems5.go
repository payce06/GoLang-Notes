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

// 7. Get the values of an object
func getValues(obj map[string]interface{}) []interface{} {
	var values []interface{}
	for _, value := range obj {
		values = append(values, value)
	}
	return values
}

// 8. Invert the keys and values of an object
func invertObject(obj map[string]interface{}) map[interface{}]string {
	inverted := make(map[interface{}]string)
	for key, value := range obj {
		inverted[value] = key
	}
	return inverted
}


// 9. Add a property to an object
func addProperty(obj map[string]interface{}, key string, value interface{}) map[string]interface{} {
	obj[key] = value
	return obj
}

// 10. Delete a property from an object
func deleteProperty(obj map[string]interface{}, key string) map[string]interface{} {
	delete(obj, key)
	return obj
}

// 11. Check if two objects are equal
func areObjectsEqual(obj1, obj2 map[string]interface{}) bool {
	if len(obj1) != len(obj2) {
		return false
	}
	for key, value := range obj1 {
		if obj2[key] != value {
			return false
		}
	}
	return true
}

// 12. Convert an object to an array of key-value pairs
func objectToPairs(obj map[string]interface{}) [][2]interface{} {
	var pairs [][2]interface{}
	for key, value := range obj {
		pairs = append(pairs, [2]interface{}{key, value})
	}
	return pairs
}

// 13. Sort an array of objects by a specific property
func sortObjects(arr []map[string]interface{}, prop string) []map[string]interface{} {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i][prop].(int) > arr[j][prop].(int) {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}