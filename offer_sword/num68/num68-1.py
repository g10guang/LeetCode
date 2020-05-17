# https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof/
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        if root is None:
            return root
        if p is None or q is None:
            return root
        bigger, smaller = p, q
        if bigger.val < smaller.val:
            bigger, smaller = smaller, bigger
        return self.recurse(root, bigger, smaller)
    
    def recurse(self, root: 'TreeNode', bigger: 'TreeNode', smaller: 'TreeNode'):
        print('root={}'.format(root.val))
        if bigger.val >= root.val >= smaller.val:
            return root
        if root.val >= bigger.val:
            return self.recurse(root.left, bigger, smaller)
        if root.val <= smaller.val:
            return self.recurse(root.right, bigger, smaller)
        return None
        