func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
    // Шаг 1: сортируем по timestamp
    n := len(username)
    indices := make([]int, n)
    for i := range indices { indices[i] = i }
    sort.Slice(indices, func(a, b int) bool {
        return timestamp[indices[a]] < timestamp[indices[b]]
    })

    // Шаг 2: строим map[user][]site (уже в хронологическом порядке)
    userSites := map[string][]string{}
    for _, idx := range indices {
        u := username[idx]
        userSites[u] = append(userSites[u], website[idx])
    }

    // Шаг 3: для каждого пользователя — уникальные тройки, потом глобальный счётчик
    patternCount := map[[3]string]int{}

    for _, sites := range userSites {
        m := len(sites)
        if m < 3 { continue }

        // собираем уникальные тройки этого пользователя
        seen := map[[3]string]bool{}
        for i := 0; i < m-2; i++ {
            for j := i + 1; j < m-1; j++ {
                for k := j + 1; k < m; k++ {
                    triple := [3]string{sites[i], sites[j], sites[k]}
                    seen[triple] = true
                }
            }
        }

        // инкрементируем глобальный счётчик (max +1 от одного пользователя)
        for triple := range seen {
            patternCount[triple]++
        }
    }

    // Шаг 4: находим паттерн с максимальным счётчиком, при ровном — лексикографически меньший
    best := [3]string{}
    bestScore := 0
    for triple, count := range patternCount {
        if (count > bestScore) || (count == bestScore && lessTriple(triple, best)) { 
            best = triple
            bestScore = count
        }
    }

    return []string{best[0], best[1], best[2]}
}

func lessTriple(a, b [3]string) bool {
    for i := 0; i < 3; i++ {
        if a[i] != b[i] {
            return a[i] < b[i]  // строки сравниваются через < без проблем [web:119]
        }
    }
    return false
}