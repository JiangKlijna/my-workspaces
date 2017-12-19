
fun test(sort_fun: ((IntArray) -> Unit)) =
        array(digit = 1000, size = 20).let {
            sort_fun(it)
            println("${sort_fun::class.java.simpleName.split("$").first()}\t${it.toList()}")
        }

fun array(size: Int = 10, digit: Int = 10): IntArray {
    val r = java.util.Random()
    val arr = IntArray(size)
    for (i in 0 until size) arr[i] = r.nextInt(digit)
    return arr
}

fun main(args: Array<String>) {
    test(BubbleSort)
    test(SelectSort)
    test(InsertSort)
    test(ShellSort)
    test(MergeSort)
    test(QuickSort)
    test(HeapSort)
    test(RadixSort)
}

//冒泡排序
val BubbleSort = fun(arr: IntArray) {
    for (i in arr.indices.reversed()) {
        for (j in 0 until i) {
            if (arr[j] > arr[j + 1]) {
                arr[j + 1].let {
                    arr[j + 1] = arr[j]
                    arr[j] = it
                }
            }
        }
    }
}

//选择排序
val SelectSort = fun(arr: IntArray) {
    for (i in arr.indices) {
        var min = i
        for (j in i + 1 until arr.size)
            if (arr[j] < arr[min]) min = j
        if (min != i)
            arr[i].let {
                arr[i] = arr[min]
                arr[min] = it
            }
    }
}

//插入排序
val InsertSort = fun(arr: IntArray) {
    for (i in 1 until arr.size) {
        var j = i - 1
        val tmp = arr[i]
        while (j >= 0 && tmp < arr[j]) {
            arr[j + 1] = arr[j]
            j -= 1
        }
        arr[j + 1] = tmp
    }
}

//希尔排序
val ShellSort = fun(arr: IntArray) {
    var gap = arr.size / 2
    while (1 <= gap) {
        for (i in gap until arr.size) {
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

//归并排序
val MergeSort = fun(arr: IntArray) {
    val merge = fun(arr: IntArray, l: Int, m: Int, r: Int) {
        val tmp = IntArray(r - l + 1)
        var i = l
        var j = m + 1
        var k = 0
        while (i <= m && j <= r) tmp[k++] = if (arr[i] <= arr[j]) arr[i++] else arr[j++]
        while (i <= m) tmp[k++] = arr[i++]
        while (j <= r) tmp[k++] = arr[j++]
        System.arraycopy(tmp, 0, arr, l, tmp.size)
    }

    fun sort(arr: IntArray, l: Int, r: Int) {
        if (l >= r) return
        val m = (l + r) / 2
        sort(arr, l, m)
        sort(arr, m + 1, r)
        merge(arr, l, m, r)
    }
    sort(arr, 0, arr.size - 1)
}

//快速排序
val QuickSort = fun(arr: IntArray) {
    val partition = fun(arr: IntArray, l: Int, r: Int): Int {
        var m = l
        for (u in l until r) {
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

    fun sort(arr: IntArray, l: Int, r: Int) {
        if (l >= r) return
        val m = partition(arr, l, r)
        sort(arr, l, m - 1)
        sort(arr, m + 1, r)
    }
    sort(arr, 0, arr.size - 1)
}

//堆排序
val HeapSort = fun(arr: IntArray) {
    val sift = fun(arr: IntArray, l: Int, r: Int) {
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

//基数排序
val RadixSort = fun(arr: IntArray) {
    val d = fun(arr: IntArray): Int {
        var d = 1
        var p = 10
        for (i in arr)
            while (i >= p) {
                p *= 10
                ++d
            }
        return d
    }(arr)
    var radix = 1
    val tmp = IntArray(arr.size)
    for (i in 1..d) {
        val count = IntArray(10)
        for (j in arr.indices) count[(arr[j] / radix) % 10]++
        for (j in 1..9) count[j] += count[j - 1]
        for (j in arr.size - 1 downTo 0) tmp[--count[(arr[j] / radix) % 10]] = arr[j]
        System.arraycopy(tmp, 0, arr, 0, arr.size)
        radix *= 10
    }
}
