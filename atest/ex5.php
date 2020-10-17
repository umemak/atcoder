<?php
$line = trim(fgets(STDIN));
$cols = explode(" ", $line);
$n = $cols[0];
$m = $cols[1];
$q = $cols[2];
for($i = 0; $i < $m; $i++) {
    $line = trim(fgets(STDIN));
    $cols = explode(" ", $line);
    $cons[$cols[0] - 1][] = $cols[1];
    $cons[$cols[1] - 1][] = $cols[0];
}
$line = trim(fgets(STDIN));
$colors = explode(" ", $line);
for($i = 0; $i < $q; $i++) {
    $line = trim(fgets(STDIN));
    $cols = explode(" ", $line);
    if ($cols[0] == 1) {
        $cc = $colors[$cols[1] - 1];
        echo $cc."\n";
        if (isset($cons[$cols[1] - 1])) {
            foreach($cons[$cols[1] - 1] as $to) {
                $colors[$to - 1] = $cc;
            }
        }
    } else {
        $cc = $colors[$cols[1] - 1];
        echo $cc."\n";
        $colors[$cols[1] - 1] = $cols[2];
    }
}
