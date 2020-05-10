# https://leetcode-cn.com/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/
# Definition for a Node.


class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def treeToDoublyList(self, root: 'Node') -> 'Node':
        head, tail = self.recurse(root)
        if head and tail:
            head.left = tail
            tail.right = head
        return head

    def recurse(self, root: 'Node'):
        if not root:
            return None, None
        leftHead, leftTail, rightHead, rightTail = None, None, None, None
        if root.right:
            rightHead, rightTail = self.recurse(root.right)
        if root.left:
            leftHead, leftTail = self.recurse(root.left)
        if leftTail is not None:
            leftTail.right = root
        if rightHead is not None:
            rightHead.left = root
        root.left = leftTail
        root.right = rightHead
        return leftHead or root, rightTail or root
