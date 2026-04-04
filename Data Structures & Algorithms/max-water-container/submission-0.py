class Solution:
    def maxArea(self, heights: List[int]) -> int:
        left, right = 0, len(heights) - 1
        m = -1
        while left < right:
            tmpArea = min(heights[left], heights[right]) * (right - left)
            m = max(tmpArea, m)
            if heights[left] <= heights[right]:
                left += 1
            else:
                right -=1

        return m
