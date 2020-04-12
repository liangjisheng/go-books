package main

// 用递归方式反转链表

// java
// public static Node reverseList2(Node head){
// // 1.递归结束条件
// if (head == null || head.next == null) {
//          return head;
//      }
//      // 递归反转 子链表
//      Node newList = reverseList2(head.next);
//      // 改变 1，2节点的指向。
//         // 通过 head.next获取节点2
//         Node t1  = head.next;
//         // 让 2 的 next 指向 2
//         t1.next = head;
//         // 1 的 next 指向 null.
//        head.next = null;
//        // 把调整之后的链表返回。
//        return newList;
// }
