cd "$(dirname "$0")/.."
mkdir -p test && cd test
npm init -y
npm i "$(ls ../propencrypt-*.tgz)"
echo 'const propencrypt = require("propencrypt"); propencrypt("-v");' > test.cjs
echo 'import propencrypt from "propencrypt";      propencrypt("-v");' > test.mjs
npx propencrypt    2>&1 | tee /dev/stderr | grep -q "Usage:"              || exit 1
npx propencrypt -v 2>&1 | tee /dev/stderr | grep -q "propencrypt version" || exit 2
node test.cjs      2>&1 | tee /dev/stderr | grep -q "propencrypt version" || exit 3
node test.mjs      2>&1 | tee /dev/stderr | grep -q "propencrypt version" || exit 4
cd ..
rm -r test
