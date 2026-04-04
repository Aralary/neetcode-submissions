class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        k = 0
        delta = 0
        for i in range(9):
            tmpRowStat = {}
            tmpColStat = {}
            tmpSquareStat = {}
            if i in [3, 6]:
                k += 3
            delta = ( delta + 3 ) % 9
            for j in range(9):
                if board[i][j] in tmpRowStat:
                    return False
                if board[i][j] != '.':
                    tmpRowStat[board[i][j]] = 1

                if board[j][i] in tmpColStat:
                    return False
                if board[j][i] != '.':
                    tmpColStat[board[j][i]] = 1

                sq_i = k + int(j/3)
                sq_j = j % 3 + delta
                if board[sq_i][sq_j] in tmpSquareStat:
                    return False
                if board[sq_i][sq_j] != '.':
                    tmpSquareStat[board[sq_i][sq_j]] = 1


        return True
        