sub fib {
    my $n = 12;
    my $a = 0;
    my $b = 1;

    for ($i = 0; $i < $n; $i++) {
        ($a, $b) = ($a+$b, $a);
    }
    print $a;
}

fib()