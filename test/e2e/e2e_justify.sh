#!/bin/bash

# Audit Test Script

echo "════════════════════════════════════════════════════════════"
echo "TEST 1: Invalid flag (should show usage)"
echo "════════════════════════════════════════════════════════════"
echo "$ go run ./cmd --align right something standard"
go run ./cmd --align right something standard
echo ""

echo "════════════════════════════════════════════════════════════"
echo "TEST 2: --align=right left standard"
echo "════════════════════════════════════════════════════════════"
echo "$ go run ./cmd --align=right left standard"
go run ./cmd --align=right left standard
echo ""

echo "════════════════════════════════════════════════════════════"
echo "TEST 3: --align=left right standard"
echo "════════════════════════════════════════════════════════════"
echo "$ go run ./cmd --align=left right standard"
go run ./cmd --align=left right standard
echo ""

echo "════════════════════════════════════════════════════════════"
echo "TEST 4: --align=center hello shadow"
echo "════════════════════════════════════════════════════════════"
echo "$ go run ./cmd --align=center hello shadow"
go run ./cmd --align=center hello shadow
echo ""

echo "════════════════════════════════════════════════════════════"
echo "TEST 5: --align=justify \"1 Two 4\" shadow"
echo "════════════════════════════════════════════════════════════"
echo "$ go run ./cmd --align=justify \"1 Two 4\" shadow"
go run ./cmd --align=justify "1 Two 4" shadow
echo ""

echo "════════════════════════════════════════════════════════════"
echo "TEST 6: --align=right 23/32 standard"
echo "════════════════════════════════════════════════════════════"
echo "$ go run ./cmd --align=right 23/32 standard"
go run ./cmd --align=right 23/32 standard
echo ""

echo "════════════════════════════════════════════════════════════"
echo "TEST 7: --align=right ABCabc123 thinkertoy"
echo "════════════════════════════════════════════════════════════"
echo "$ go run ./cmd --align=right ABCabc123 thinkertoy"
go run ./cmd --align=right ABCabc123 thinkertoy
echo ""

echo "════════════════════════════════════════════════════════════"
echo "TEST 8: --align=center \"#\$%&\\\"\" thinkertoy"
echo "════════════════════════════════════════════════════════════"
echo "$ go run ./cmd --align=center \"#\$%&\\\"\" thinkertoy"
go run ./cmd --align=center "#\$%&\"" thinkertoy
echo ""

echo "════════════════════════════════════════════════════════════"
echo "TEST 9: --align=left \"23Hello World!\" standard"
echo "════════════════════════════════════════════════════════════"
echo "$ go run ./cmd --align=left \"23Hello World!\" standard"
go run ./cmd --align=left "23Hello World!" standard
echo ""

echo "════════════════════════════════════════════════════════════"
echo "TEST 10: --align=justify \"HELLO there HOW are YOU?!\" thinkertoy"
echo "════════════════════════════════════════════════════════════"
echo "$ go run ./cmd --align=justify \"HELLO there HOW are YOU?!\" thinkertoy"
go run ./cmd --align=justify "HELLO there HOW are YOU?!" thinkertoy
echo ""

echo "════════════════════════════════════════════════════════════"
echo "TEST 11: --align=right \"a -> A b -> B c -> C\" shadow"
echo "════════════════════════════════════════════════════════════"
echo "$ go run ./cmd --align=right \"a -> A b -> B c -> C\" shadow"
go run ./cmd --align=right "a -> A b -> B c -> C" shadow
echo ""

echo "════════════════════════════════════════════════════════════"
echo "✅ ALL TESTS COMPLETE"
echo "════════════════════════════════════════════════════════════"