cd "$(dirname "$0")/.."  || exit 1
mkdir -p test && cd test || exit 2
npm init -y
npm i "$(ls ../propencrypt-*.tgz)"
echo 'const propencrypt = require("propencrypt"); propencrypt("-v");' > test.cjs
echo 'import propencrypt from "propencrypt";      propencrypt("-v");' > test.mjs
npx propencrypt    2>&1 | tee /dev/stderr | grep -q "Usage:"              || exit 11
npx propencrypt -v 2>&1 | tee /dev/stderr | grep -q "propencrypt version" || exit 12
node test.cjs      2>&1 | tee /dev/stderr | grep -q "propencrypt version" || exit 13
node test.mjs      2>&1 | tee /dev/stderr | grep -q "propencrypt version" || exit 14
cd ..
rm -r test
