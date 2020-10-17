<?php
$line = trim(fgets(STDIN));
$cols = explode(" ", $line);
$n = $cols[0];
$m = $cols[1];
$q = $cols[2];
for($i = 0; $i < $n; $i++) {
    for($j = 0; $j < $m; $j++) {
        $scores[$i][$j] = 0;
    }
}
for($i = 1; $i <= $q; $i++) {
    $line = trim(fgets(STDIN));
    $cols = explode(" ", $line);
    if ($cols[0] == 1) {
        $a = 0;
        for($j = 0; $j < $m; $j++) {
            if ($scores[$cols[1] - 1][$j] == 1) {
                $b = $n;
                for($k = 0; $k < $n; $k++) {
                    $b -= $scores[$k][$j];
                }
                $a += $b;
            }
        }
        echo $a . "\n";
    } else {
        $scores[$cols[1] - 1][$cols[2] - 1] = 1;
    }
}
