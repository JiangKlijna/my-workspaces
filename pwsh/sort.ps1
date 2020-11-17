


function RandomArray {
    param ([int]$size = 10, [int]$digit = 10)
    return Get-Random -Maximum $digit -SetSeed (Get-Random) -Count $size
}

function ArrayToString {
    param ([Object[]]$arr)
    $str = "["
    for ($i = 0; $i -lt $arr.Count; $i++) {
        $o = $arr[$i]
        if ($i -eq $arr.Count - 1) {
            $str += "$o"
        }
        else {
            $str += "$o, "
        }
    }
    return $str + "]"
}

function Test {
    param ([Sorter]$sort)
    $arr = RandomArray 20 1000
    $sort.sort($arr)
    $res = ArrayToString($arr)
    Write-Output "$sort`t$res"
}

class Sorter {
    [void] sort($arr) { }
}

# 冒泡排序
class BubbleSort :Sorter {
    [void] sort($arr) {
        for ($i = $arr.length - 1; $i -ge 0; $i--) {
            for ($j = 0; $j -lt $i; $j++) {
                if ($arr[$j] -gt $arr[$j + 1]) {
                    $temp = $arr[$j + 1]
                    $arr[$j + 1] = $arr[$j]
                    $arr[$j] = $temp
                }
            }
        }
    }
}

# 选择排序
class SelectSort :Sorter {
    [void] sort($arr) {
        for ($i = 0; $i -lt $arr.length; $i++) {
            $min = $i
            for ($j = $i + 1; $j -lt $arr.length; $j++) {
                if ($arr[$min] -gt $arr[$j]) { $min = $j }
            }
            if ($i -eq $min) { continue }
            $temp = $arr[$i]
            $arr[$i] = $arr[$min]
            $arr[$min] = $temp
        }
    }
}

# 插入排序
class InsertSort :Sorter {
    [void] sort($arr) {
        for ($i = 1; $i -lt $arr.length; $i++) {
            $j = $i - 1;
            $tmp = $arr[$i];
            while ($j -ge 0 -and $arr[$j] -gt $tmp) {
                $arr[$j + 1] = $arr[$j];
                $j -= 1;
            }
            $arr[$j + 1] = $tmp;
        }
    }
}

# 希尔排序
class ShellSort :Sorter {
    [void] sort($arr) {
        $gap = [Math]::floor($arr.length / 2)
        while (1 -le $gap) {
            for ($i = $gap; $i -lt $arr.length; $i++) {
                $j = $i - $gap
                $tmp = $arr[$i]
                while ($j -ge 0 -and $arr[$j] -gt $tmp) {
                    $arr[$j + $gap] = $arr[$j]
                    $j -= $gap
                }
                $arr[$j + $gap] = $tmp
            }
            $gap = [Math]::floor($gap / 2)
        }
    }
}

# 归并排序
class MergeSort :Sorter {
    [void] sort($arr) {
        $this.realSort($arr, 0, $arr.Count - 1)
    }
    [void] realSort($arr, [int]$l, [int]$r) {
        if ($l -ge $r) { return }
        $m = [Math]::floor(($l + $r) / 2)
        $this.realSort($arr, $l, $m)
        $this.realSort($arr, $m + 1, $r)
        $this.merge($arr, $l, $m, $r)
    }
    [void] merge($arr, [int]$l, [int]$m, [int]$r) {
        $tmp = 0..($r - $l)
        $i = $l
        $j = $m + 1
        $k = 0
        while ($i -le $m -and $j -le $r) { $tmp[$k++] = ($arr[$i] -lt $arr[$j]) ? $arr[$i++] : $arr[$j++] }
        while ($i -le $m) { $tmp[$k++] = $arr[$i++] }
        while ($j -le $r) { $tmp[$k++] = $arr[$j++] }
        for ($i = 0; $i -lt $tmp.length; $i++) { 
            $arr[$l + $i] = $tmp[$i] 
        }
    }
}

# 快速排序
class QuickSort :Sorter {
    [void] sort($arr) {
        $this.realSort($arr, 0, $arr.Count - 1)
    }
    [void] realSort($arr, [int]$l, [int]$r) {
        if ($l -ge $r) { return }
        $m = $this.partition($arr, $l, $r)
        $this.realSort($arr, $l, $m - 1)
        $this.realSort($arr, $m + 1, $r)
    }
    [int] partition($arr, [int]$l, [int]$r) {
        $m = $l
        for ($i = $l; $i -lt $r; $i++) {
            if ($arr[$i] -lt $arr[$r]) {
                $tmp = $arr[$i];
                $arr[$i] = $arr[$m];
                $arr[$m] = $tmp;
                $m++;
            }
        }
        $tmp = $arr[$r]
        $arr[$r] = $arr[$m]
        $arr[$m] = $tmp
        return $m
    }
}

# 堆排序
class DeapSort :Sorter {
    [void] sort($arr) {
        $n = $arr.length - 1;
        for ($i = [Math]::floor($n / 2); $i -ge 0; $i--) { $this.sift($arr, $i, $n) }
        for ($i = $n; $i -ge 1; $i--) {
            $temp = $arr[$i]
            $arr[$i] = $arr[0]
            $arr[0] = $temp
            $this.sift($arr, 0, $i - 1)
        }    
    }
    [void] sift($arr, [int]$l, [int]$r) {
        $i = $l
        $j = $l * 2 + 1
        $tmp = $arr[$l]
        while ($j -le $r) {
            if ($j -lt $r -and $arr[$j] -lt $arr[$j + 1]) { $j++ }
            if ($tmp -lt $arr[$j]) {
                $arr[$i] = $arr[$j];
                $i = $j;
                $j = 2 * $i + 1;
            }
            else { break }
        }
        $arr[$i] = $tmp;
    }
}

# 基数排序
class RadixSort :Sorter {
    [void] sort($arr) {
        $d = $this.maxbit($arr)
        $radix = 1;
        $tmp = 0..($arr.length - 1)
        for ($i = 1; $i -le $d; $i++) {
            $count = 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
            for ($j = 0; $j -lt $arr.length; $j++) { $count[[Math]::floor($arr[$j] / $radix) % 10]++; }
            for ($j = 1; $j -lt 10; $j++) { $count[$j] += $count[$j - 1]; }
            for ($j = $arr.length - 1; $j -ge 0; $j--) { $tmp[--$count[[Math]::floor($arr[$j] / $radix) % 10]] = $arr[$j]; }
            for ($j = 0; $j -lt $arr.length; $j++) { $arr[$j] = $tmp[$j]; }
            $radix *= 10;
        }
    }
    [int] maxbit($arr) {
        $d = 1
        $p = 10
        foreach ($i in $arr) {
            while ($i -ge $p) {
                $p *= 10;
                $d++;
            }
        }
        return $d;
    }
}

foreach ($s in (
        [BubbleSort]::new(), (New-Object SelectSort), 
        [InsertSort]::new(), (New-Object ShellSort), 
        [MergeSort]::new(), (New-Object QuickSort),
        [DeapSort]::new(), (New-Object RadixSort))
) { Test($s) }
