package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_cpuinfo.h>
import "C"

// NumCPU returns the number of logical CPU cores, this is affected
// by hyperthreading
func NumCPU() int {
	return int(C.SDL_GetCPUCount())
}

// NumRam returns the amount of RAM installed in MB
//func NumRam() int {
//	return int(C.SDL_GetSystemRAM())
//}

type CPUInfo struct {
	Cores int
	// RamMB       int
	L1CacheLine int
	RDTSC       bool // Time Stamp Counter
	AltiVec     bool // PowerPC SIMD instruction set
	MMX         bool // Original x86 SIMD instructions
	Has3DNow    bool // MMX Extension by AMD
	SSE         bool // Streaming SIMD Extensions, extends MMX
	SSE2        bool // Extends SSE
	SSE3        bool // Extends SSE2
	SSE41       bool // Extends SSE3
	SSE42       bool // Extends SSE4, Text Instructions
	// AVX         bool // Advanced Vector Extensions, Sandy Bridge+ & Bulldozer+
}

func toBool(b C.SDL_bool) bool {
	return b == C.SDL_TRUE
}

// NewCPUInfo discovers the various parameters of the current CPU
// and sets them on the CPUInfo object it returns.
func NewCPUInfo() CPUInfo {
	return CPUInfo{
		Cores: NumCPU(),
		//RamMB:       NumRam(),
		L1CacheLine: int(C.SDL_GetCPUCacheLineSize()),
		RDTSC:       toBool(C.SDL_HasRDTSC()),
		AltiVec:     toBool(C.SDL_HasAltiVec()),
		MMX:         toBool(C.SDL_HasMMX()),
		Has3DNow:    toBool(C.SDL_Has3DNow()),
		SSE:         toBool(C.SDL_HasSSE()),
		SSE2:        toBool(C.SDL_HasSSE2()),
		SSE3:        toBool(C.SDL_HasSSE3()),
		SSE41:       toBool(C.SDL_HasSSE41()),
		SSE42:       toBool(C.SDL_HasSSE42()),
		//AVX:         toBool(C.SDL_HasAVX()),
	}
}
