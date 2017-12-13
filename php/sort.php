<?php
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

test(new BubbleSort());
test(new SelectSort());

?>
