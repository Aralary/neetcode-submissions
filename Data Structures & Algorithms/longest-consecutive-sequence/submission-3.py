class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        stat = set(nums)
        res = 0
        for num in stat:
            if num - 1 not in stat:
                tmp_len = 1
                tmp = num
                while tmp + 1 in stat:
                    tmp_len += 1
                    tmp += 1
                res = max(res, tmp_len)

        return res