
object Sort {

    @JvmStatic
    fun main(args: Array<String>) {
        test(BubbleSort())
        test(SelectSort())
        test(InsertSort())
        test(ShellSort())
        test(MergeSort())
        test(QuickSort())
        test(HeapSort())
        test(RadixSort())
    }

    private fun test(s: Sorter) {
        array(digit = 1000, size = 20).let {
            val t0 = System.currentTimeMillis()
            s.sort(it)
            val t1 = System.currentTimeMillis()
            println("${s.javaClass.simpleName}\t${t1 - t0}\t${it.toList()}")
        }
    }

    @JvmOverloads fun array(size: Int = 10, digit: Int = 10): IntArray {
        return IntArray(size).apply {
            val r = java.util.Random()
            for (i in indices) this[i] = r.nextInt(digit)
        }
    }

    interface Sorter {
        fun sort(arr: IntArray)
    }

    //冒泡排序
    class BubbleSort : Sorter {
        override fun sort(arr: IntArray) {
            for (i in arr.indices) {
                for (j in i + 1..arr.size - 1) {
                    if (arr[i] > arr[j]) {
                        arr[i].let {
                            arr[i] = arr[j]
                            arr[j] = it
                        }
                    }
                }
            }
        }
    }

    //选择排序
    class SelectSort : Sorter {
        override fun sort(arr: IntArray) {
            for (i in arr.indices) {
                var min = i
                for (j in i + 1..arr.size - 1)
                    if (arr[j] < arr[min]) min = j
                if (min != i)
                    arr[i].let {
                        arr[i] = arr[min]
                        arr[min] = it
                    }
            }
        }
    }

    //插入排序
    class InsertSort : Sorter {
        override fun sort(arr: IntArray) {
            for (i in 1..arr.size - 1) {
                var j = i - 1
                val tmp = arr[i]
                while (j >= 0 && tmp < arr[j]) {
                    arr[j + 1] = arr[j]
                    j -= 1
                }
                arr[j + 1] = tmp
            }
        }
    }

    //希尔排序
    class ShellSort : Sorter {
        override fun sort(arr: IntArray) {
            var gap = arr.size / 2
            while (1 <= gap) {
                for (i in gap..arr.size - 1) {
                    var j = i - gap
                    val tmp = arr[i]
                    while (j >= 0 && tmp < arr[j]) {
                        arr[j + gap] = arr[j]
                        j -= gap
                    }
                    arr[j + gap] = tmp
                }
                gap /= 2
            }
        }
    }

    //归并排序
    class MergeSort : Sorter {
        override fun sort(arr: IntArray) {
            sort(arr, 0, arr.size - 1)
        }

        private fun sort(arr: IntArray, l: Int, r: Int) {
            if (l >= r) return
            val m = (l + r) / 2
            sort(arr, l, m)
            sort(arr, m + 1, r)
            merge(arr, l, m, r)
        }

        private fun merge(arr: IntArray, l: Int, m: Int, r: Int) {
            val tmp = IntArray(r - l + 1)
            var i = l
            var j = m + 1
            var k = 0
            while (i <= m && j <= r) tmp[k++] = if (arr[i] <= arr[j]) arr[i++] else arr[j++]
            while (i <= m) tmp[k++] = arr[i++]
            while (j <= r) tmp[k++] = arr[j++]
            System.arraycopy(tmp, 0, arr, l, tmp.size)
        }
    }

    //快速排序
    class QuickSort : Sorter {
        override fun sort(arr: IntArray) {
            sort(arr, 0, arr.size - 1)
        }

        private fun sort(arr: IntArray, l: Int, r: Int) {
            if (l >= r) return
            val m = partition(arr, l, r)
            sort(arr, l, m - 1)
            sort(arr, m + 1, r)
        }

        private fun partition(arr: IntArray, l: Int, r: Int): Int {
            var m = l
            for (u in l..r - 1) {
                if (arr[u] <= arr[r]) {
                    arr[u].let {
                        arr[u] = arr[m]
                        arr[m] = it
                        m++
                    }
                }
            }
            arr[m].let {
                arr[m] = arr[r]
                arr[r] = it
            }
            return m
        }
    }

    //堆排序
    class HeapSort : Sorter {
        override fun sort(arr: IntArray) {
            val n = arr.size - 1
            for (i in n / 2 downTo 0) sift(arr, i, n)
            for (i in n downTo 1) {
                arr[i].let {
                    arr[i] = arr[0]
                    arr[0] = it
                }
                sift(arr, 0, i - 1)
            }
        }

        private fun sift(arr: IntArray, l: Int, r: Int) {
            var i = l
            var j = 2 * i + 1
            val tmp = arr[l]
            while (j <= r) {
                if (j < r && arr[j] < arr[j + 1]) j++
                if (tmp < arr[j]) {
                    arr[i] = arr[j]
                    i = j
                    j = 2 * i + 1
                } else break
            }
            arr[i] = tmp
        }
    }


    //基数排序
    class RadixSort : Sorter {
        override fun sort(arr: IntArray) {
            val d = maxbit(arr)
            val tmp = IntArray(arr.size)
            var radix = 1
            for (i in 1..d) {
                val count = IntArray(10)
                for (j in arr.indices) count[(arr[j] / radix) % 10]++
                for (j in 1..9) count[j] += count[j - 1]
                for (j in arr.size - 1 downTo 0) tmp[--count[(arr[j] / radix) % 10]] = arr[j]
                System.arraycopy(tmp, 0, arr, 0, arr.size)
                radix *= 10
            }
        }

        private fun maxbit(arr: IntArray): Int {
            var d = 1
            var p = 10
            for (i in arr.indices)
                while (arr[i] >= p) {
                    p *= 10
                    ++d
                }
            return d
        }
    }

}
