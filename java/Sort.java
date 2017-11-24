
import java.util.Arrays;
import java.util.Random;

public class Sort {

    public static void main(String[] args) throws Exception {
        test(BubbleSort.class);
        test(SelectSort.class);
        test(InsertSort.class);
        test(ShellSort.class);
        test(MergeSort.class);
        test(QuickSort.class);
        test(HeapSort.class);
        test(RadixSort.class);
    }

    public static void test(Class<? extends Sorter> cls) throws Exception {
        int[] arr = array(20, 1000);
        cls.newInstance().sort(arr);
        System.out.format("%s\t%s\n", cls.getSimpleName(), Arrays.toString(arr));
    }

    public static int[] array() {
        return array(10, 10);
    }

    public static int[] array(int size) {
        return array(size, 10);
    }

    public static int[] array(int size, int digit) {
        Random r = new Random();
        int[] arr = new int[size];
        for (int i = 0; i < size; i++) arr[i] = r.nextInt(digit);
        return arr;
    }

    public interface Sorter {
        void sort(int[] arr);
    }

    //冒泡排序
    public static class BubbleSort implements Sorter {
        @Override
        public void sort(int[] arr) {
            for (int i = arr.length - 1; i >= 0; i--) {
                for (int j = 0; j < i; j++) {
                    if (arr[j] > arr[j + 1]) {
                        int temp = arr[j + 1];
                        arr[j + 1] = arr[j];
                        arr[j] = temp;
                    }
                }
            }
        }
    }

    //选择排序
    public static class SelectSort implements Sorter {
        @Override
        public void sort(int[] arr) {
            for (int i = 0; i < arr.length; i++) {
                int min = i;
                for (int j = i + 1; j < arr.length; j++) {
                    if (arr[j] < arr[min])
                        min = j;
                }
                if (min != i) {
                    int temp = arr[i];
                    arr[i] = arr[min];
                    arr[min] = temp;
                }
            }
        }
    }

    //插入排序
    public static class InsertSort implements Sorter {
        @Override
        public void sort(int[] arr) {
            for (int i = 1; i < arr.length; i++) {
                int j = i - 1;
                int tmp = arr[i];
                while (j >= 0 && tmp < arr[j]) {
                    arr[j + 1] = arr[j];
                    j--;
                }
                arr[j + 1] = tmp;
            }
        }
    }

    //希尔排序
    public static class ShellSort implements Sorter {
        @Override
        public void sort(int[] arr) {
            int gap = arr.length / 2;
            while (1 <= gap) {
                for (int i = gap; i < arr.length; i++) {
                    int j = i - gap;
                    int tmp = arr[i];
                    while (j >= 0 && tmp < arr[j]) {
                        arr[j + gap] = arr[j];
                        j -= gap;
                    }
                    arr[j + gap] = tmp;
                }
                gap /= 2;
            }
        }
    }

    //归并排序
    public static class MergeSort implements Sorter {
        @Override
        public void sort(int[] arr) {
            sort(arr, 0, arr.length - 1);
        }

        private void sort(int[] arr, int l, int r) {
            if (l >= r) return;
            int m = (l + r) / 2;
            sort(arr, l, m);
            sort(arr, m + 1, r);
            merge(arr, l, m, r);
        }

        private void merge(int[] arr, int l, int m, int r) {
            int[] tmp = new int[r - l + 1];
            int i = l, j = m + 1, k = 0;
            while (i <= m && j <= r) tmp[k++] = (arr[i] < arr[j]) ? arr[i++] : arr[j++];
            while (i <= m) tmp[k++] = arr[i++];
            while (j <= r) tmp[k++] = arr[j++];
            System.arraycopy(tmp, 0, arr, l, tmp.length);
        }
    }

    //快速排序
    public static class QuickSort implements Sorter {
        @Override
        public void sort(int[] arr) {
            sort(arr, 0, arr.length - 1);
        }

        private void sort(int[] arr, int l, int r) {
            if (l >= r) return;
            int m = partition(arr, l, r);
            sort(arr, l, m - 1);
            sort(arr, m + 1, r);
        }

        private int partition(int[] arr, int l, int r) {
            int m = l;
            for (int u = l; u < r; u++) {
                if (arr[u] <= arr[r]) {
                    int temp = arr[u];
                    arr[u] = arr[m];
                    arr[m] = temp;
                    m++;
                }
            }
            int temp = arr[r];
            arr[r] = arr[m];
            arr[m] = temp;
            return m;
        }
    }

    //堆排序
    public static class HeapSort implements Sorter {
        @Override
        public void sort(int[] arr) {
            int n = arr.length - 1;
            for (int i = (n / 2); i >= 0; i--) sift(arr, i, n);
            for (int i = n; i >= 1; i--) {
                int temp = arr[i];
                arr[i] = arr[0];
                arr[0] = temp;
                sift(arr, 0, i - 1);
            }
        }

        public void sift(int[] arr, int l, int r) {
            int i = l, j = l * 2 + 1, tmp = arr[l];
            while (j <= r) {
                if (j < r && arr[j] < arr[j + 1]) j++;
                if (tmp < arr[j]) {
                    arr[i] = arr[j];
                    i = j;
                    j = 2 * i + 1;
                } else break;
            }
            arr[i] = tmp;
        }
    }

    //基数排序
    public static class RadixSort implements Sorter {
        @Override
        public void sort(int[] arr) {
            int d = maxbit(arr), radix = 1;
            int[] tmp = new int[arr.length];
            for (int i = 1; i <= d; i++) {
                int[] count = new int[10];
                for (int j = 0; j < arr.length; j++) count[(arr[j] / radix) % 10]++;
                for (int j = 1; j < 10; j++) count[j] += count[j - 1];
                for (int j = arr.length - 1; j >= 0; j--) tmp[--count[(arr[j] / radix) % 10]] = arr[j];
                System.arraycopy(tmp, 0, arr, 0, arr.length);
                radix *= 10;
            }
        }

        public int maxbit(int[] arr) {
            int d = 1, p = 10;
            for (int i : arr)
                while (i >= p) {
                    p *= 10;
                    ++d;
                }
            return d;
        }
    }


}
