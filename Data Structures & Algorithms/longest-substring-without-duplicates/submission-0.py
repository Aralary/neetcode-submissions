class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        stat = {}
        l, r = 0, 0 
        res = 0
        tmpL = 0
        while r < len(s):
            if s[r] in stat and stat[s[r]] > 0:
                # двигаем l до тех пор, пока не уберутся повторы
                stat[s[l]] -= 1
                tmpL -= 1
                l += 1
            else:
                stat[s[r]] = 1
                tmpL += 1
                r += 1
            res = max(res, tmpL)
        return res