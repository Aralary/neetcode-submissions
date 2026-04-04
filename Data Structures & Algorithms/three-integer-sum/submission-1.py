class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        i, j, k = 0, len(nums) - 2, len(nums) - 1
        s_nums = sorted(nums)
        ans = []
        visited = set()
        while k > 1:
            while i < j:
                if s_nums[i] + s_nums[j] + s_nums[k] == 0:
                    tmplst = [s_nums[i], s_nums[j], s_nums[k]]
                    tmp = ",".join(map(str, tmplst))
                    if tmp not in visited:
                        ans.append(tmplst)
                        visited.add(tmp)
                    i += 1
                if s_nums[i] + s_nums[j] > -s_nums[k]:
                    j -= 1
                if s_nums[i] + s_nums[j] < -s_nums[k]:
                    i += 1
            k -= 1
            i = 0
            j = k - 1
        return ans