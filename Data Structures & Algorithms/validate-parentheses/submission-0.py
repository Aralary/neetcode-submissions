class Solution:
    def isValid(self, s: str) -> bool:
        st = []
        op = set(['(', '[', '{'])
        cl = set([')', ']', '}'])
        for sep in s:
            if sep in op:
                st.append(sep)
            elif sep in cl:
                if len(st) == 0:
                    return False
                match sep:
                    case ']':
                        if st[len(st)-1] != '[':
                            return False
                        else:
                            st.pop()
                    case ')':
                        if st[len(st)-1] != '(':
                            return False
                        else:
                            st.pop()
                    case '}':
                        if st[len(st)-1] != '{':
                            return False
                        else:
                            st.pop()
            else:
                return False

            print(st)
        return len(st) == 0
