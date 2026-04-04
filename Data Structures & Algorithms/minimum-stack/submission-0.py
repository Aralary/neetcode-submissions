class MinStack:

    def __init__(self):
        self.data = []
        self.min_prefix = []
        

    def push(self, val: int) -> None:
        self.data.append(val)
        if len(self.min_prefix) == 0:
            self.min_prefix.append(val)
        else:
            self.min_prefix.append(min(val, self.min_prefix[len(self.min_prefix) - 1]))

    def pop(self) -> None:
        if len(self.data) > 0:
            self.data.pop()
            self.min_prefix.pop()

    def top(self) -> int:
        return self.data[len(self.data)-1]

    def getMin(self) -> int:
        return self.min_prefix[len(self.min_prefix)-1]
        
