
public class LongestIncreasingSubsequence{
    public static void main(String[] args) {
        int[] arr = new int[]{2,1,6,4,5,2,7,4};
        System.out.println(O_n2_one(arr));
        System.out.println(O_n2_two(arr));
        System.out.println(O_n_logn(arr));
        System.out.println(binarySearch(new int[]{0,1,3,6}, 0, 0, 3));
    }

    public static final int binarySearch(int[] is, int i, int l, int r) {
        while (l <= r) {
            int ave = (l + r) >> 1;
            int is_ave = is[ave];
            if (i < is_ave)
                r = ave - 1;
            else if (i > is_ave)
                l = ave + 1;
            else
                return ave;
        }
        return -1;
    }

    public static final int O_n2_one(int[] arr) {
        int[] h = new int[arr.length];
        h[0] = 1;
        int re_max = 1;
        for (int i = 1; i < arr.length; i++) {
            boolean is_min = true;
            int i_max = 1, arr_i = arr[i];
            for (int j = i - 1; j >= 0; j--) {
                if (arr[j] < arr_i) {
                    is_min = false;
                    if (i_max < h[j])
                        i_max = h[j];
                }
            }
            i_max = is_min ? 1 : i_max + 1;
            h[i] = i_max;
            if(i_max > re_max)
                re_max = i_max;
        }
        return re_max;
    }

    public static final int O_n2_two(int[] arr) {
        int[] h = new int[arr.length];
        h[0] = arr[0];
        int re_max = 1;
        for (int i = 1; i < arr.length; i++) {
            int arr_i = arr[i];
            if (arr_i > h[re_max - 1]) {
                h[re_max++] = arr_i;
                continue;
            }
            for (int j = 0; j < re_max; j++) {
                if (h[j] > arr_i) {
                    h[j] = arr_i;
                    break;
                }
            }
        }
        return re_max;
    }

    public static final int O_n_logn(int[] arr) {
        int[] h = new int[arr.length];
        h[0] = arr[0];
        int re_max = 1;
        for (int i = 1; i < arr.length; i++) {
            int arr_i = arr[i];
            int l = 0, r = re_max - 1;
            while (l <= r) {
                int ave = (l + r) >> 1;
                int ave_val = h[ave];
                if (arr_i > ave_val)
                    l = ave + 1;
                else if (arr_i < ave_val)
                    r = ave - 1;
                else
                    break;
            }
            if (arr_i > h[l])
                re_max++;
            h[l] = arr_i;
        }
        return re_max;
    }
}
