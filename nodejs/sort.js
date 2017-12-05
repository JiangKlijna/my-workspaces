
let random = digit => parseInt(Math.random() * digit);

let array = function (size, digit) {
    if (size === undefined) size = 10;
    if (digit === undefined) digit = 10;
    let arr = new Array(size);
    for (var i = 0; i < size; i++) arr[i] = random(digit);
    return arr;
}

let test = function (Sort) {
    let arr = array(20, 1000);
    new Sort().sort(arr);
    console.log(`${Sort.name}\t${arr}`);
}

Array.prototype.toString = function () {
    var buf = ['['];
    for (var i = 0; i < this.length-1; i++)
        buf.push(this[i], ', ');
    buf.push(this[this.length-1], ']');
    return buf.join('');
}
//冒泡排序
let BubbleSort = function () {
    this.sort = function (arr) {
        for (var i = arr.length - 1; i >= 0; i--) {
            for (var j = 0; j < i; j++) {
                if (arr[j] > arr[j+1]) {
                    var temp = arr[j+1];
                    arr[j+1] = arr[j];
                    arr[j] = temp;
                }
            }
        }
    }
}
//选择排序
let SelectSort = function () {
    this.sort = function (arr) {
        for (var i = 0; i < arr.length; i++) {
            var min = i;
            for (var j = i + 1; j < arr.length; j++) {
                if (arr[min] > arr[j]) min = j;
            }
            if (i === min) continue;
            var temp = arr[i];
            arr[i] = arr[min];
            arr[min] = temp;
        }
    }
}
//插入排序
let InsertSort = function () {
    this.sort = function (arr) {
        for (var i = 1; i < arr.length; i++) {
            var j = i - 1;
            var tmp = arr[i];
            while (j >= 0 && arr[j] > tmp) {
                arr[j+1] = arr[j];
                j -= 1;
            }
            arr[j+1] = tmp;
        }
    }
}
//希尔排序
let ShellSort = function () {
    this.sort = function (arr) {
        var gap = parseInt(arr.length/2);
        while (1 <= gap) {
            for (var i = gap; i < arr.length; i++) {
                var j = i - gap;
                var tmp = arr[i];
                while (j >= 0 && arr[j] > tmp) {
                    arr[j+gap] = arr[j];
                    j -= gap;
                }
                arr[j+gap] = tmp;
            }
            gap = parseInt(gap/2);
        }
    }
}
//归并排序
let MergeSort = function () {
    this.sort = arr => sort(arr, 0, arr.length - 1);
    let sort = function (arr, l, r) {
        if (l >= r) return;
        var m = parseInt((l + r) / 2);
        sort(arr, l, m);
        sort(arr, m+1, r);
        merge(arr, l, m, r);
    }
    let merge = function (arr, l, m, r) {
        var tmp = new Array(r - l + 1);
        var i = l, j = m+1, k = 0;
        while (i <= m && j <= r) tmp[k++] = (arr[i] < arr[j]) ? arr[i++] : arr[j++];
        while (i <= m) tmp[k++] = arr[i++];
        while (j <= r) tmp[k++] = arr[j++];
        for (var i  = 0; i < tmp.length; i++) arr[l+i] = tmp[i];
    }
    this.merge = function (arr) {
        merge(arr, 0, parseInt(arr.length / 2), arr.length - 1)
    }
}
//快速排序
let QuickSort = function () {
    this.sort = arr => sort(arr, 0, arr.length - 1);
    let sort = function (arr, l, r) {
        if (l >= r) return;
        var m = partition(arr, l, r);
        sort(arr, l, m-1);
        sort(arr, m+1, r);
    }
    let partition = function (arr, l, r) {
        var m = l;
        for (var i = l; i < r; i++) {
            if (arr[i] < arr[r]) {
                var tmp = arr[i];
                arr[i] = arr[m];
                arr[m] = tmp;
                m++;
            }
        }
        var tmp = arr[r];
        arr[r] = arr[m];
        arr[m] = tmp;
        return m;
    }
}
//堆排序
let DeapSort = function () {
    this.sort = function (arr) {
        var n = arr.length - 1;
        for (var i = parseInt(n/2); i >= 0; i--) sift(arr, i, n)
        for (var i = n; i >= 1; i--) {
            var temp = arr[i];
            arr[i] = arr[0];
            arr[0] = temp;
            sift(arr, 0, i - 1);
        }
    }
    let sift = function (arr, l, r) {
        var i = l, j = l * 2 + 1, tmp = arr[l];
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
let RadixSort = function () {
    this.sort = function (arr) {
        var d = maxbit(arr), radix = 1;
        var tmp = new Array(arr.length);
        for (var i = 1; i <= d; i++) {
            var count = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
            for (var j = 0; j < arr.length; j++) count[parseInt(arr[j] / radix) % 10]++;
            for (var j = 1; j < 10; j++) count[j] += count[j - 1];
            for (var j = arr.length - 1; j >= 0; j--) tmp[--count[parseInt(arr[j] / radix) % 10]] = arr[j];
            for (var j = 0; j < arr.length; j++) arr[j] = tmp[j];
            radix *= 10;
        }
    }
    let maxbit = function (arr) {
        var d = 1, p = 10;
        for (var i of arr)
            while (i >= p) {
                p *= 10;
                ++d;
            }
        return d;
    }
}
// main
if (require.main === module) {
    [BubbleSort, SelectSort, InsertSort, ShellSort, MergeSort, QuickSort, DeapSort, RadixSort]
    .forEach(Sort => test(Sort));
}
