//go:build linux && (arm || 386 || mips || mipsle || ppc)

package lumine

func itou(n int) uint32 {
	return uint32(n)
}
