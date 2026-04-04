class Solution:

    def encode(self, strs: List[str]) -> str:
        res = ""
        for s in strs:
            res += (chr(len(s)) + s)
        return res

    def decode(self, s: str) -> List[str]:
        i = 0
        res = []
        while i < len(s):
            l = ord(s[i])
            tmps = s[i+1: i + 1 + l]
            res.append(tmps)
            i += l + 1
        return res