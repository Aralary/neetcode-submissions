class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        prefix = [1] * (len(nums) + 1)
        for i in range(len(nums)):
            prefix[i+1] = prefix[i] * nums[i]
        prefix = prefix[1:]
        postfix = [1] * (len(nums) + 1)
        i = len(nums) - 1
        while i > -1:
            postfix[i] = postfix[i+1] * nums[i]
            i -= 1
        postfix = postfix[:len(postfix)-1]
        res = [1] * len(nums)
        for i in range(len(nums)):
            if i == 0:
                res[i] = postfix[i+1]
            elif i == len(nums) - 1:
                res[i] = prefix[i-1]
            else:
                res[i] = postfix[i+1] * prefix[i-1]
        return res
        