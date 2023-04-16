class AVLNode {
public:
    int val;
    int height;
    AVLNode* left;
    AVLNode* right;
    AVLNode(int v) {
        val = v;
        height = 1;
        left = nullptr;
        right = nullptr;
    }
};

class AVLTree {
public:
    AVLNode* root;
    AVLTree() {
        root = nullptr;
    }
    int height(AVLNode* node) {
        if (node == nullptr) {
            return 0;
        }
        return node->height;
    }
    int balanceFactor(AVLNode* node) {
        if (node == nullptr) {
            return 0;
        }
        return height(node->left) - height(node->right);
    }
    void updateHeight(AVLNode* node) {
        node->height = 1 + max(height(node->left), height(node->right));
    }
    AVLNode* rotateRight(AVLNode* node) {
        AVLNode* newRoot = node->left;
        node->left = newRoot->right;
        newRoot->right = node;
        updateHeight(node);
        updateHeight(newRoot);
        return newRoot;
    }
    AVLNode* rotateLeft(AVLNode* node) {
        AVLNode* newRoot = node->right;
        node->right = newRoot->left;
        newRoot->left = node;
        updateHeight(node);
        updateHeight(newRoot);
        return newRoot;
    }
    AVLNode* findMin(AVLNode* node) {
        while (node->left != nullptr) {
            node = node->left;
        }
        return node;
    }
    AVLNode* remove(AVLNode* node, int val) {
        if (node == nullptr) {
            return node;
        }
        if (val < node->val) {
            node->left = remove(node->left, val);
        } else if (val > node->val) {
            node->right = remove(node->right, val);
        } else {
            if (node->left == nullptr || node->right == nullptr) {
                AVLNode* temp = node->left ? node->left : node->right;
                if (temp == nullptr) {
                    temp = node;
                    node = nullptr;
                } else {
                    *node = *temp;
                }
                delete temp;
            } else {
                AVLNode* temp = findMin(node->right);
                node->val = temp->val;
                node->right = remove(node->right, temp->val);
            }
        }
        if (node == nullptr) {
            return node;
        }
        updateHeight(node);
        int bf = balanceFactor(node);
        if (bf > 1 && balanceFactor(node->left) >= 0) {
            return rotateRight(node);
        }
        if (bf > 1 && balanceFactor(node->left) < 0) {
            node->left = rotateLeft(node->left);
            return rotateRight(node);
        }
        if (bf < -1 && balanceFactor(node->right) <= 0) {
            return rotateLeft(node);
        }
        if (bf < -1 && balanceFactor(node->right) > 0) {
            node->right = rotateRight(node->right);
            return rotateLeft(node);
        }
        return node;
    }
    void remove(int val) {
        root = remove(root, val);
    }
    void insert(int val) {
        root = insert(root, val);
    }
    AVLNode* insert(AVLNode* node, int val) {
        if (node == nullptr) {
            return new AVLNode(val);
        }
        if (val < node->val) {
            node->left = insert(node->left, val);
        } else if (val > node->val) {
            node->right = insert(node->right, val);
        } else {
            return node;
        }
        updateHeight(node);
        int bf = balanceFactor(node);
        if (bf > 1 && val < node->left->val) {
            return rotateRight(node);
        }
        if (bf < -1 && val > node->right->val) {
            return rotateLeft(node);
        }
        if (bf > 1 && val > node->left->val) {
            node->left = rotateLeft(node->left);
            return rotateRight(node);
        }
        if (bf < -1 && val < node->right->val) {
            node->right = rotateRight(node->right);
            return rotateLeft(node);
        }
        return node;
    }
    AVLNode* find(int val) {
        AVLNode* node = root;
        while (node != nullptr) {
            if (val < node->val) {
                node = node->left;
            } else if (val > node->val) {
                node = node->right;
            } else {
                return node;
            }
        }
        return nullptr;
    }
};
