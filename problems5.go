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

// 14. Create an object from two arrays (keys and values)
func createObject(keys []string, values []interface{}) map[string]interface{} {
	obj := make(map[string]interface{})
	for i := 0; i < len(keys); i++ {
		obj[keys[i]] = values[i]
	}
	return obj
}

// 15. Get the entries of an object
func getEntries(obj map[string]interface{}) [][2]interface{} {
	var entries [][2]interface{}
	for key, value := range obj {
		entries = append(entries, [2]interface{}{key, value})
	}
	return entries
}

// 16. Check if an object is empty
func isObjectEmpty(obj map[string]interface{}) bool {
	return len(obj) == 0
}

// 17. Get common keys and values between two objects
func commonKeysValues(obj1, obj2 map[string]interface{}) map[string]interface{} {
	common := make(map[string]interface{})
	for key, value := range obj1 {
		if obj2[key] == value {
			common[key] = value
		}
	}
	return common
}

// 18. Find the keys with the highest value in an object
func keysWithHighestValue(obj map[string]interface{}) []string {
	maxValue := -1 << 63 // Minimum possible value
	var highestKeys []string
	for key, value := range obj {
		if val, ok := value.(int); ok {
			if val > maxValue {
				maxValue = val
				highestKeys = []string{key}
			} else if val == maxValue {
				highestKeys = append(highestKeys, key)
			}
		}
	}
	return highestKeys
}

// 19. Convert an object to a query string
func objectToQueryString(obj map[string]interface{}) string {
	var queryString []string
	for key, value := range obj {
		queryString = append(queryString, fmt.Sprintf("%s=%v", key, value))
	}
	return strings.Join(queryString, "&")
}

// 20. Get a nested value from an object using a key path
func getNestedValue(obj map[string]interface{}, keyPath string) interface{} {
	keys := strings.Split(keyPath, ".")
	for _, key := range keys {
		if subObj, ok := obj[key].(map[string]interface{}); ok {
			obj = subObj
		} else {
			return nil
		}
	}
	return obj
}

func main() {
	// Test functions

	// 1. Merge Objects
	fmt.Println(mergeObjects(map[string]interface{}{"a": 1, "b": 2}, map[string]interface{}{"b": 3, "c": 4})) // Output: map[a:1 b:3 c:4]

	// 2. Count Properties
	fmt.Println(countProperties(map[string]interface{}{"a": 1, "b": 2, "c": 3})) // Output: 3

	// 3. Get Property Values
	fmt.Println(getPropertyValues([]map[string]interface{}{{"a": 1}, {"a": 2}, {"b": 3}}, "a")) // Output: [1 2]

	// 4. Contains Property
	fmt.Println(containsProperty(map[string]interface{}{"a": 1, "b": 2}, "b")) // Output: true

	// 5. Deep Clone
	original := map[string]interface{}{"a": 1, "b": map[string]interface{}{"c": 2}}
	clone := deepClone(original)
	fmt.Println(clone) // Output: map[a:1 b:map[c:2]]

	// 6. Get Keys
	fmt.Println(getKeys(map[string]interface{}{"a": 1, "b": 2})) // Output: [a b]

	// 7. Get Values
	fmt.Println(getValues(map[string]interface{}{"a": 1, "b": 2})) // Output: [1 2]

	// 8. Invert Object
	fmt.Println(invertObject(map[string]interface{}{"a": 1, "b": 2})) // Output: map[1:a 2:b]

	// 9. Add Property
	fmt.Println(addProperty(map[string]interface{}{"a": 1}, "b", 2)) // Output: map[a:1 b:2]

	// 10. Delete Property
	fmt.Println(deleteProperty(map[string]interface{}{"a": 1, "b": 2}, "b")) // Output: map[a:1]

	// 11. Are Objects Equal
	fmt.Println(areObjectsEqual(map[string]interface{}{"a": 1, "b": 2}, map[string]interface{}{"a": 1, "b": 2})) // Output: true

	// 12. Object to Pairs
	fmt.Println(objectToPairs(map[string]interface{}{"a": 1, "b": 2})) // Output: [[a 1] [b 2]]