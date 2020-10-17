<?php
$line = trim(fgets(STDIN));
$cols = explode(" ", $line);
$a = $cols[0];
$r = $cols[1];
$n = $cols[2];
for ($i = 1; $i < $n; $i++) {
    $a *= $r;
    if (1000000000 < $a) {
        echo "large\n";
        break;
    }
}
if ($a <= 1000000000) {
    echo $a."\n";
}
