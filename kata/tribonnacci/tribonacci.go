package tribonacci

func sum(slice []float64) float64 {
    var total float64 = 0

    for _, num := range slice {
        total += num
    }

    return total
}

func Calculate(window []float64, results []float64, n int) []float64 {
    if (len(results) == n) {
        return results
    }
    nextNum := sum(window[:])
    results = append(results, nextNum)
    var nextWindow []float64 = window[1:3]
    return Calculate(append(nextWindow, nextNum), results, n)
}

func Tribonacci(signature [3]float64, n int) []float64 {
    if n < len(signature) {
        return signature[:n]
    }
    return Calculate(signature[:], signature[:], n)
}
