func leastInterval(tasks []byte, n int) int {
    freq := make(map[byte]int)
    maxFreq := 0

    for _, t := range tasks {
        freq[t]++
        if freq[t] > maxFreq {
            maxFreq = freq[t]
        }
    }

    maxCount := 0
    for _, count := range freq {
        if count == maxFreq {
            maxCount++
        }
    }
    return max(len(tasks), (maxFreq-1)*(n+1)+maxCount)
}
