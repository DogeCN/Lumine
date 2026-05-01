//go:build linux && (amd64 || arm64 || loong64 || mips64 || mips64le || ppc64 || ppc64le || riscv64 || s390x || sparc64)

package lumine

func itou(n int) uint64 {
	return uint64(n)
}
