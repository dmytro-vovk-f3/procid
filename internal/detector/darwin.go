//go:build darwin

package detector

/*
#include <stdlib.h>
#include <libproc.h>
*/
import "C"
import (
	"github.com/mitchellh/go-ps"
	"path/filepath"
	"strconv"
	"unsafe"
)

func FindPID(name string) (string, error) {
	processes, err := ps.Processes()
	if err != nil {
		return "", nil
	}

	baseName := filepath.Base(name)
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
		if binaryFullName == name {
			return strconv.Itoa(processes[i].Pid()), nil
		}
	}

	return "", nil
}
