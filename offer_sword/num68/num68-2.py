# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def lowestCommonAncestor(self, root: TreeNode, p: TreeNode, q: TreeNode) -> TreeNode:
        self.p = p
        self.q = q
        self.p_path = []
        self.q_path = []
        self.p_flag = False
        self.q_flag = False
        self.stack = []
        self.dfs(root)
        common = None
        while len(self.p_path) > 0 and len(self.q_path) > 0 and self.p_path[0] == self.q_path[0]:
            common = self.p_path[0]
            self.p_path = self.p_path[1:]
            self.q_path = self.q_path[1:]
        return common

    def dfs(self, node: TreeNode):
        if self.p_flag and self.q_flag:
            return
        if node is None:
            return 
        self.stack.append(node)
        if node == self.p and self.p_flag is False:
            self.p_flag = True
            self.p_path = self.stack.copy()
        if node == self.q and self.q_flag is False:
            self.q_flag = True
            self.q_path = self.stack.copy()
        self.dfs(node.left)
        self.dfs(node.right)
        self.stack.pop()