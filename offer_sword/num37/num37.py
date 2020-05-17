# https://leetcode-cn.com/problems/xu-lie-hua-er-cha-shu-lcof/
# Definition for a binary tree node.
class TreeNode(object):
   def __init__(self, x):
       self.val = x
       self.left = None
       self.right = None


class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        l = []
        queue = []
        if root:
            queue.append(root)
        while queue:
            node = queue[0]
            queue = queue[1:]
            if node:
                queue.append(node.left)
                queue.append(node.right)
                l.append(node.val)
            else:
                l.append(None)
        return self.list2str(l)

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        l = self.str2list(data)
        root = None
        if not l:
            return root
        root = TreeNode(l[0])
        queue = [root]
        l = l[1:]
        while queue:
            node = queue[0]
            queue = queue[1:]
            left = self.toNode(l[0])
            right = self.toNode(l[1])
            l = l[2:]
            node.left = left
            node.right = right
            if left:
                queue.append(left)
            if right:
                queue.append(right)
        return root

    def toNode(self, v):
        if v is not None:
            return TreeNode(v)
        else:
            return None

    def list2str(self, l):
        return ','.join(self.intlist2strlist(l))

    def str2list(self, s):
        if s:
            return self.strlist2intlist(s.split(","))
        else:
            return []

    def intlist2strlist(self, intlist):
        strlist = []
        for v in intlist:
            strlist.append(str(v))
        return strlist

    def strlist2intlist(self, strlist):
        intlist = []
        for v in strlist:
            if v == str(None):
                intlist.append(None)
            else:
                intlist.append(int(v))
        return intlist
        

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))