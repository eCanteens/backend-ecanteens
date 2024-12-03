package helpers

import (
	"fmt"
	"math/rand/v2"
	"regexp"
	"strings"
	"time"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
    snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
    snake  = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
    return strings.ToLower(snake)
}

func PointerTo[T any](s T) *T {
    return &s
}

func RandomElement[T any](nums []T) T {
	return nums[rand.IntN(len(nums))]
}

func Find[T any](slice []T, test func(*T) bool) (ret *T, idx int) {
    idx = -1
    for i, s := range slice {
        if test(&s) {
            ret = &s
            idx = i
            break
        }
    }
    return
}

func Map[T any, Y any](slice []T, test func(*T) *Y) []Y {
    var result []Y

    for _, s := range slice {
        ret := test(&s)
        if ret != nil {
            result = append(result, *ret)
        }
    }
    return result
}

func RemoveItem[T any](slice []T, test func(*T) bool) []T {
    for i, v := range slice {
        if test(&v) {
            return append(slice[:i], slice[i+1:]...)
        }
    }

    return slice
}

func GenerateTrxCode(userId uint) string {
    return fmt.Sprintf("EC-%d-%d", time.Now().Unix(), userId)
}

func RemoveDuplicates[T any, K comparable](items []T, keySelector func(*T) K, limit int) []T {
	uniqueMap := make(map[K]bool)
	var result []T

	for _, item := range items {
        if len(result) >= limit && limit != 0 {
            break
        }

		key := keySelector(&item)
		if !uniqueMap[key] {
			uniqueMap[key] = true
			result = append(result, item)
		}
	}

	return result
}