class Solution:
    def partitionLabels(self, S):
        """
        :type S: str
        :rtype: List[int]
        """
        lastIndexes = [-1] * 26
        orda = ord('a')
        for i, c in enumerate(S):
            index = ord(c) - orda
            lastIndexes[index] = i
        start = 0
        result = []
        while start < len(S):
            end = start
            i = start
            last = None
            while i <= end:
                index = ord(S[i]) - orda
                last = lastIndexes[index]
                if end < last:
                    end = last
                i+=1
            result.append(end-start+1)
            start=end+1
        return result
    