#!/usr/bin/env node

const { spawnSync } = require("child_process");
const { join } = require("path");

const platforms = {
  linux: { x64: "linux-x64/propencrypt" },
  win32: { x64: "win32-x64/propencrypt.exe" }
};

const { platform, arch } = process;
const bin = platforms?.[platform]?.[arch];

if (!bin) {
  throw new Error(`Platform ${platform}-${arch} not supported`);
}

if (require.main === module) {
  process.exitCode = propencrypt(...process.argv.slice(2));
}

/**
 * Run propencrypt as a child process with the given arguments.
 * @param  {...string} args Command-line arguments to pass to the process
 * @returns {number} Process exit code
 */
function propencrypt(...args) {
  const result = spawnSync(join(__dirname, bin), args, { stdio: "inherit" });
  if (result.error) {
    throw result.error;
  }
  return result.status;
}

module.exports = propencrypt;
