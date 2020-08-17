public class findKth {

    public static int findKth(int[] arr1, int[] arr2, int k) {
        if(arr1 == null || arr1.length < 1)
            return arr2[k-1];
        if(arr2 == null || arr2.length < 1)
            return arr1[k-1];
            
        return findKth(arr1, 0, arr1.length - 1, arr2, 0, arr2.length - 1, k - 1);
    }

    public static int findKth(int[] arr1, int l1, int r1, int[] arr2, int l2, int r2, int k) {
        if(l1 > r1)
            return arr2[l2 + k];
        if(l2 > r2)
            return arr1[l1 + k];
        if (k == 0)
            return Math.min(arr1[l1],arr2[l2]);
        int md1 = l1 + k/2 < r1 ? l1 + k/2 : r1;
        int md2 = l2 + k/2 < r2 ? l2 + k/2 : r2;
        if(arr1[md1] < arr2[md2])
            return findKth(arr1, md1 + 1, r1, arr2, l2, r2, k - k / 2 - 1);
        else if (arr1[md1] > arr2[md2])
            return findKth(arr1, l1, r1, arr2, md2 + 1, r2, k - k / 2 - 1);
        else
            return arr1[md1];
    }

    public static void main(String[] args) {
        // int[] arr1 = {8, 10, 23, 42, 54, 86, 90};
        // int[] arr2 = {3, 12, 50, 56, 65, 69, 75};
        // System.out.println(findKth(arr1, arr2, 5));

        int[] arr1 = {1, 2, 3};
        int[] arr2 = {0, 4, 5, 6, 7, 8};
        System.out.println(findKth(arr1, arr2, 7));
    }

}