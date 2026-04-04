class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        st = []
        for t in tokens:
            if t == '*':
                left = st[-1]
                st.pop()
                right = st[-1]
                st.pop()
                st.append(left * right)
            elif t == '/':
                left = st[-1]
                st.pop()
                right = st[-1]
                st.pop()
                st.append(int(right / left))
            elif t == '+':
                left = st[-1]
                st.pop()
                right = st[-1]
                st.pop()
                st.append(left + right)
            elif t == '-':
                left = st[-1]
                st.pop()
                right = st[-1]
                st.pop()
                st.append(right - left)
            else: 
                st.append(int(t))
            print(f"token: {t}")
            print(f"stack: {st}")
        return int(st[0])
