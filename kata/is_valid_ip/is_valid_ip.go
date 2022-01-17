package is_valid_ip

import (
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {
  // length between 7 and 15 inclusive
  // split on . and check each number between 0 and 255
	if len(ip) < 7 || len(ip) > 15 {
		return false
	}

	splitIp := strings.Split(ip, ".")

	if (len(splitIp) != 4) {
		return false
	}

	for _, octet := range splitIp {
		if (len(octet) > 1 && octet[0] == '0') {
			return false
		}
		octectInt, err := strconv.Atoi(octet)
		if (err != nil || octectInt < 0 || octectInt > 255) {
			return false
		}
	}

	return true
}
