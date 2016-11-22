class Node(object):

    """Docstring for Node. """

    def __init__(self,val):
        """TODO: to be defined1. """
        self.left  = None
        self.right = None
        self.val = val

    def getHeight(self):
        left  = self.left.getHeight() if self.left else 0
        right = self.right.getHeight() if self.right else 0

        return left+1 if left > right else right+1

    def levelOrder(self):
        queue = []

        queue.append(self)
        ret = str(self.val)
        
        while len(queue) > 0:
            cnode = queue.pop(0)

            if cnode.left:
                queue.append(cnode.left)
                ret += " " + str(cnode.left.val)
            if cnode.right:
                queue.append(cnode.right)
                ret += " " + str(cnode.right.val)

        print(ret)

    def depthPath(self):
        stack = []
        cur = self
        while True:
            if cur:
                stack.append(cur)
                cur = cur.left
                continue

            if len(stack) == 0:
                break
            while True:
                if len(stack) == 0:
                    cur = None
                    break

                cur = stack.pop()
                
                if cur.right
                    cur = cur.right
                    break




