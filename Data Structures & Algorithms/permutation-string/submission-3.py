class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        if len(s2) < len(s1) :
            return False
        l, r = 0 , 0
        etalon = {}
        tmpStat = {}
        for i in range(len(s1)):
            etalon[s1[i]] = 1 + etalon.get(s1[i], 0)
            tmpStat[s2[i]] = 1 + tmpStat.get(s2[i], 0)
            r += 1
        r -= 1
        tmpStat[s2[r]] -= 1
        while r < len(s2):
            tmpStat[s2[r]] = 1 + tmpStat.get(s2[r], 0)
            # проверяем
            if len(tmpStat) == len(etalon):
                f = True
                for k in tmpStat:
                    if k not in etalon:
                        f = False
                        break
                    elif etalon[k] != tmpStat[k]:
                        f = False
                        break
                if f:
                    return True
            # убираем старую букву
            tmpStat[s2[l]] -= 1
            if tmpStat[s2[l]] == 0:
                tmpStat.pop(s2[l])
            l += 1
            r += 1

        return False