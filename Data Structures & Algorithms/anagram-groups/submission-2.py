class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        alphabet = {
            'a': 0,
            'b': 1,
            'c': 2,
            'd': 3,
            'e': 4,
            'f': 5,
            'g': 6,
            'h': 7,
            'i': 8,
            'j': 9,
            'k': 10,
            'l': 11,
            'm': 12,
            'n': 13,
            'o': 14,
            'p': 15,
            'q': 16,
            'r': 17,
            's': 18,
            't': 19,
            'u': 20,
            'v': 21,
            'w': 22,
            'x': 23,
            'y': 24,
            'z': 25
        }
        stat = {}
        for s in strs:
            mask = [0]*26
            for sep in s:
                mask[alphabet[sep]] += 1
            str_mask = ",".join(map(str,mask))
            if str_mask in stat:
                stat[str_mask].append(s)
            else:
                stat[str_mask] = [s]
        print(stat)
        res = []
        for k in stat:
            res.append(stat[k])
        return res
        
        