class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        stat = {}
        for num in nums:
            if num not in stat:
                stat[num] = 0
            stat[num] += 1
        res = [[] for _ in range(len(nums) + 1)]
        for key in stat:
            res[stat[key]].append(key)
        ans = []
        print(res)
        for item in reversed(res):
            if len(item) == 0:
                continue
            ans.extend(item)
            if len(ans) >= k:
                break
        return ans
            