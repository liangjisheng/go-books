public class combination {
    public static final int COMBINATION_CNT = 2;
    public static void combination(int[] arr, int k, int[] select) {
        int selectNum = selectedNum(select);
        if (selectNum == COMBINATION_CNT) {
            for (int j = 0; j < select.length; j++) {
                if (select[j] == 1) {
                    System.out.print(arr[j]);
                }
            }
            System.out.print("\n");
        } else {
            if (k >= arr.length) {
                return;
            }
                   
            select[k] = 1;
            combination(arr, k+1, select);
            
            select[k] = 0;    
            combination(arr, k+1, select);
        }
    }

    public static int selectedNum(int[] select) {
        int res = 0;
        for (int i = 0; i < select.length; i++) {
            if (select[i] == 1) {
                res += select[i];
            }
        }
        return res;
    }

    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4};
        int[] select = {0, 0, 0, 0};    
        combination(arr, 0, select);
    }
}