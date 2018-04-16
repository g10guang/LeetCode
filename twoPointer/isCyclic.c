/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
bool hasCycle(struct ListNode *head) {
    if (!head) {
        return false;
    }
    struct ListNode *p, *q;
    p = head;
    q = head->next;
    while (q && p != q) {
        p = p->next;
        q = q->next;
        if (!q) {
            return false;
        }
        q = q->next;
    }
    return q != 0;
}
