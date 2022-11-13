version=$(grep -oE '[0-9.]{5,}' internal/cmd/version/version.go)
mkdir -p               build/package/npm/{linux-x64,win32-x64}
cp propencrypt         build/package/npm/linux-x64/
cp propencrypt.exe     build/package/npm/win32-x64/
cp {README.md,LICENSE} build/package/npm/
cd                     build/package/npm/
chmod 755              */propencrypt*
npm version "${version}" --git-tag-version false
npm pack
