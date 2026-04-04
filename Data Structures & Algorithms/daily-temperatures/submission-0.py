class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        dp = [0] * len(temperatures)
        st = []
        j = 0
        for i in range(len(temperatures)):
            if len(st) == 0:
                st.append(i)
                j += 1
                continue
            if temperatures[i] > temperatures[st[-1]]:
                while len(st) > 0 and temperatures[i] > temperatures[st[-1]]:
                    dp[st[-1]] = i - st[-1]
                    st.pop()
            st.append(i)
            print(f"step {j}: stack = {st}")
            print(f"dp = {dp}")
            j += 1
        return dp
                