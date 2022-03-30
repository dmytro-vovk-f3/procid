//go:build darwin

package detector

/*
#include <stdlib.h>
#include <libproc.h>
*/
import "C"
import (
	"path/filepath"
	"strconv"
	"unsafe"

	"github.com/mitchellh/go-ps"
)

func FindPID(binaryName string) (string, error) {
	processes, err := ps.Processes()
	if err != nil {
		return "", nil
	}

	baseName := filepath.Base(binaryName)
	buf := C.malloc(C.PROC_PIDPATHINFO_MAXSIZE)
	defer C.free(unsafe.Pointer(buf))

	for i := range processes {
		if baseName != processes[i].Executable() {
			continue
		}

		ret := C.proc_pidpath(C.int(processes[i].Pid()), buf, C.PROC_PIDPATHINFO_MAXSIZE)
		if ret < 0 {
			continue
		}

		binaryFullName := string(C.GoBytes(unsafe.Pointer(buf), ret))
		if binaryFullName == binaryName {
			return strconv.Itoa(processes[i].Pid()), nil
		}
	}

	return "", nil
}
