<?php
//https://tool.lu/coderunner/
// 获得一个随机数组
function random_array($size=10, $digit=10) {
    $arr = range(0, $size-1);
    for ($i=0; $i < $size; $i++) {
        $arr[$i] = rand(0, $digit-1);
    }
    return $arr;
}

function test($sorter) {
    $arr = random_array(20, 1000);
    $sorter->sort($arr);
    echo get_class($sorter).'    ';
    foreach ($arr as $k => $v) {
        echo ($k == 0 ? '[' : '').$v.($k == count($arr)-1 ? ']' : ', ');
    }
}

interface Sorter {
    public function sort(&$arr);
}
// 冒泡排序
class BubbleSort implements Sorter {
    public function sort(&$arr) {
        for ($i=count($arr)-1; $i >= 0; $i--) {
            for ($j=0; $j < $i; $j++) {
                if ($arr[$j] > $arr[$j+1]) {
                    $tmp = $arr[$j];
                    $arr[$j] = $arr[$j+1];
                    $arr[$j+1] = $tmp;
                }
            }
        }
    }
}
// 选择排序
class SelectSort implements Sorter {
    public function sort(&$arr) {
        for ($i = 0; $i < count($arr); $i++) {
            $min = $i;
            for ($j = $i + 1; $j < count($arr); $j++) {
                if ($arr[$j] < $arr[$min])
                    $min = $j;
            }
            if ($min != $i) {
                $temp = $arr[$i];
                $arr[$i] = $arr[$min];
                $arr[$min] = $temp;
            }
        }
    }
}
// 插入排序
class InsertSort implements Sorter {
    public function sort(&$arr) {
        for ($i = 1; $i < count($arr); $i++) {
            $j = $i - 1;
            $tmp = $arr[$i];
            while ($j >= 0 && $tmp < $arr[$j]) {
                $arr[$j + 1] = $arr[$j];
                $j--;
            }
            $arr[$j + 1] = $tmp;
        }
    }
}
// 希尔排序
class ShellSort implements Sorter {
    public function sort(&$arr) {
        $n = count($arr);
        $gap = (int)($n / 2);
        while (1 <= $gap) {
            for ($i = $gap; $i < $n; $i++) {
                $j = $i - $gap;
                $tmp = $arr[$i];
                while ($j >= 0 && $tmp < $arr[$j]) {
                    $arr[$j + $gap] = $arr[$j];
                    $j -= $gap;
                }
                $arr[$j + $gap] = $tmp;
            }
            $gap = (int)($gap / 2)
        }
    }
}

test(new BubbleSort());
test(new SelectSort());
test(new InsertSort());
test(new ShellSort());

?>
