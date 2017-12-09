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
    public function sort($arr);
}

class BubbleSort implements Sorter {

    public function sort($arr) {
        $n = count($arr);

    }
}

test(new BubbleSort());

?>
