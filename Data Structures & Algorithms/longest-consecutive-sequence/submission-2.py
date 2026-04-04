class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        stat = {}
        for num in nums:
            stat[num] = 1
        starts = {}
        for num in nums:
            if num - 1 not in stat:
                starts[num] = 1
        res = 1
        for start in starts:
            tmp_len = 1
            tmp = start
            while tmp + 1 in stat:
                tmp_len += 1
                tmp += 1
            res = max(res, tmp_len)

        return res