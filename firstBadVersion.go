package main

import "fmt"

func isBadVersion(n int) bool {
	return n >= 6
}

func firstBadVersion(n int) int {
    lo := 0
    hi := n
    
    for lo < hi {
        mid := (lo + hi) / 2
        
        if isBadVersion(mid) && !isBadVersion(mid - 1) {
            return mid
        } else if !isBadVersion(mid) && isBadVersion(mid + 1) {
            return mid + 1
        } else if isBadVersion(mid) {
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }
    
    return -1
}

func main() {
	fmt.Println(firstBadVersion(10))
}