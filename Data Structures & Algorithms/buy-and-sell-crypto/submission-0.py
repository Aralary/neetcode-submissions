class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        i, j = 0, 1
        maxDelta = 0
        while j < len(prices):
            if prices[j] > prices[i]:
                maxDelta = max(maxDelta, prices[j] - prices[i])
            else:
                i = j
            j+= 1
        return maxDelta