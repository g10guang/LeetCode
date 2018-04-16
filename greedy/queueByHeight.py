# https://leetcode.com/problems/queue-reconstruction-by-height/description/

class Solution:
    def reconstructQueue(self, people):
        """
        :type people: List[List[int]]
        :rtype: List[List[int]]
        """
        if not people:
            return people
        length = len(people)
        i = length - 1
        l = sorted(people, key=lambda x: x[0] * 1000 + x[1])
        while i >= 0 and l[-1][0] == l[i][0]:
            i -= 1
        while i >= 0:
            pos = l[i][1]
            k = i - 1
            while k >= 0 and l[k][0] == l[i][0]:
                pos -= 1
                k -= 1
            for j in range(pos):
                t = i+j
                l[t+1], l[t] = l[t], l[t+1]
            i -= 1
        return l
