import TreeNode from './tree-node'

/**
 * Insert at leaf positions is often preferred over inserting between 2 existing nodes. O(log n), assuming the tree
 * is balanced
 * i.e. if inserting 6 into this tree:
 *               5
 *              / \
 *            2   7
 *
 * You would end up with
 *               5
 *              / \
 *            2   7
 *               /
 *              6
 *
 * @param root
 * @param val
 * @returns {boolean|boolean|*}
 */
function insert(root, val) {
  if (root == null) {
    return new TreeNode(val);
  }

  if (val > root.val) {
    root.right = insert(root.right, val);
  } else  if (val < root.val) {
    root.left = insert(root.left, val);
  }
  return root;
}

// Remove a node and return the root of the BST.
function remove(root, val) {
  if (root == null) {
    return null;
  }
  if (val > root.val) {
    root.right = remove(root.right, val);
  } else if (val < root.val) {
    root.left = remove(root.left, val);
  } else {
    if (root.left == null) {
      return root.right;
    } else if (root.right == null) {
      return root.left;
    } else {
      let minNode = minValueNode(root.right);
      root.val = minNode.val;
      root.right = remove(root.right, minNode.val);
    }
  }
  return root;
}

// Return the minimum value node of the BST.
function minValueNode(root) {
  let curr = root;
  while(curr != null && curr.left != null) {
    curr = curr.left;
  }
  return curr;
}
