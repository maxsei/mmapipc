package common

/*
#include <sys/shm.h>
#include <sys/types.h>
#include <sys/ipc.h>
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>

int shm_attach(int shm_id, void *ret) {
	char* addr = shmat(shm_id, NULL, 0);
	if(addr == (char*)(-1)){
		return -1;
	}
	ret = addr;
	return 0;
}

*/
import "C"
import (
	"unsafe"
)

var key C.key_t

const ProjID = 69420

func init() {
	key := C.ftok(C.CString("./ipc/hello"), C.int(ProjID))
	if key == -1 {
		panic("ftok")
	}
}

func Key() C.key_t {
	return key
}

func Shmget() []byte {
	shmid := C.shmget(Key(), (C.size_t)(666), 0644|C.IPC_CREAT)
	if shmid == -1 {
		panic("shmget")
	}

	var data unsafe.Pointer
	/*
		data := C.shmat(shmid, C.NULL, 0)
		if data == (*C.char)(C.int(-1)) {
			panic("shmat")
		}
	*/
	if ret := int(C.shm_attach(shmid, data)); ret == -1 {
		panic("shm_attach")
	}
	return []byte{}
}
