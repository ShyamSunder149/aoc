mkdir "$1"
cd "$1"
cp "$(pwd)/../../boilerplate/aoc_bp.go" "main.go"
cp "$(pwd)/../../boilerplate/aoc_bp.go" "main2.go"
touch "input.txt"
