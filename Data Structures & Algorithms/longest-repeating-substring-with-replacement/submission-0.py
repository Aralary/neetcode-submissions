class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        stat = {}
        res = 0
        max_cnt = 0
        l, r = 0, 0
        while r < len(s):
            stat[s[r]] = 1 + stat.get(s[r], 0)
            max_cnt = max(max_cnt, stat[s[r]])
            while r - l + 1 - max_cnt > k:
                stat[s[l]] -= 1
                l += 1
            res = max(r - l + 1, res)
            r += 1
        return res 