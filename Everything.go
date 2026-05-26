package everything

import (
	"unsafe"
)

// Source: Everything.h -> Macro constants
const (
	EverythingSdkVersion                            uint32 = 2
	EverythingOk                                    uint32 = 0
	EverythingErrorMemory                           uint32 = 1
	EverythingErrorIpc                              uint32 = 2
	EverythingErrorRegisterclassex                  uint32 = 3
	EverythingErrorCreatewindow                     uint32 = 4
	EverythingErrorCreatethread                     uint32 = 5
	EverythingErrorInvalidindex                     uint32 = 6
	EverythingErrorInvalidcall                      uint32 = 7
	EverythingErrorInvalidrequest                   uint32 = 8
	EverythingErrorInvalidparameter                 uint32 = 9
	EverythingSortNameAscending                     uint32 = 1
	EverythingSortNameDescending                    uint32 = 2
	EverythingSortPathAscending                     uint32 = 3
	EverythingSortPathDescending                    uint32 = 4
	EverythingSortSizeAscending                     uint32 = 5
	EverythingSortSizeDescending                    uint32 = 6
	EverythingSortExtensionAscending                uint32 = 7
	EverythingSortExtensionDescending               uint32 = 8
	EverythingSortTypeNameAscending                 uint32 = 9
	EverythingSortTypeNameDescending                uint32 = 10
	EverythingSortDateCreatedAscending              uint32 = 11
	EverythingSortDateCreatedDescending             uint32 = 12
	EverythingSortDateModifiedAscending             uint32 = 13
	EverythingSortDateModifiedDescending            uint32 = 14
	EverythingSortAttributesAscending               uint32 = 15
	EverythingSortAttributesDescending              uint32 = 16
	EverythingSortFileListFilenameAscending         uint32 = 17
	EverythingSortFileListFilenameDescending        uint32 = 18
	EverythingSortRunCountAscending                 uint32 = 19
	EverythingSortRunCountDescending                uint32 = 20
	EverythingSortDateRecentlyChangedAscending      uint32 = 21
	EverythingSortDateRecentlyChangedDescending     uint32 = 22
	EverythingSortDateAccessedAscending             uint32 = 23
	EverythingSortDateAccessedDescending            uint32 = 24
	EverythingSortDateRunAscending                  uint32 = 25
	EverythingSortDateRunDescending                 uint32 = 26
	EverythingRequestFileName                       uint32 = 0x00000001
	EverythingRequestPath                           uint32 = 0x00000002
	EverythingRequestFullPathAndFileName            uint32 = 0x00000004
	EverythingRequestExtension                      uint32 = 0x00000008
	EverythingRequestSize                           uint32 = 0x00000010
	EverythingRequestDateCreated                    uint32 = 0x00000020
	EverythingRequestDateModified                   uint32 = 0x00000040
	EverythingRequestDateAccessed                   uint32 = 0x00000080
	EverythingRequestAttributes                     uint32 = 0x00000100
	EverythingRequestFileListFileName               uint32 = 0x00000200
	EverythingRequestRunCount                       uint32 = 0x00000400
	EverythingRequestDateRun                        uint32 = 0x00000800
	EverythingRequestDateRecentlyChanged            uint32 = 0x00001000
	EverythingRequestHighlightedFileName            uint32 = 0x00002000
	EverythingRequestHighlightedPath                uint32 = 0x00004000
	EverythingRequestHighlightedFullPathAndFileName uint32 = 0x00008000
	EverythingTargetMachineX86                      uint32 = 1
	EverythingTargetMachineX64                      uint32 = 2
	EverythingTargetMachineArm                      uint32 = 3
)

func (e *Everything) SetSearchW(LpString *uint16) {
	getProc("Everything_SetSearchW").Call(uintptr(unsafe.Pointer(LpString)))
}

func (e *Everything) SetSearchA(LpString *uint8) {
	getProc("Everything_SetSearchA").Call(uintptr(unsafe.Pointer(LpString)))
}

func (e *Everything) SetMatchPath(BEnable int32) {
	getProc("Everything_SetMatchPath").Call(uintptr(BEnable))
}

func (e *Everything) SetMatchCase(BEnable int32) {
	getProc("Everything_SetMatchCase").Call(uintptr(BEnable))
}

func (e *Everything) SetMatchWholeWord(BEnable int32) {
	getProc("Everything_SetMatchWholeWord").Call(uintptr(BEnable))
}

func (e *Everything) SetRegex(BEnable int32) {
	getProc("Everything_SetRegex").Call(uintptr(BEnable))
}

func (e *Everything) SetMax(DwMax uint32) {
	getProc("Everything_SetMax").Call(uintptr(DwMax))
}

func (e *Everything) SetOffset(DwOffset uint32) {
	getProc("Everything_SetOffset").Call(uintptr(DwOffset))
}

func (e *Everything) SetReplyWindow(HWnd uintptr) {
	getProc("Everything_SetReplyWindow").Call(HWnd)
}

func (e *Everything) SetReplyID(DwId uint32) {
	getProc("Everything_SetReplyID").Call(uintptr(DwId))
}

func (e *Everything) SetSort(DwSort uint32) {
	getProc("Everything_SetSort").Call(uintptr(DwSort))
}

func (e *Everything) SetRequestFlags(DwRequestFlags uint32) {
	getProc("Everything_SetRequestFlags").Call(uintptr(DwRequestFlags))
}

func (e *Everything) GetMatchPath() int32 {
	r1, _, _ := getProc("Everything_GetMatchPath").Call()
	return int32(r1)
}

func (e *Everything) GetMatchCase() int32 {
	r1, _, _ := getProc("Everything_GetMatchCase").Call()
	return int32(r1)
}

func (e *Everything) GetMatchWholeWord() int32 {
	r1, _, _ := getProc("Everything_GetMatchWholeWord").Call()
	return int32(r1)
}

func (e *Everything) GetRegex() int32 {
	r1, _, _ := getProc("Everything_GetRegex").Call()
	return int32(r1)
}

func (e *Everything) GetMax() uint32 {
	r1, _, _ := getProc("Everything_GetMax").Call()
	return uint32(r1)
}

func (e *Everything) GetOffset() uint32 {
	r1, _, _ := getProc("Everything_GetOffset").Call()
	return uint32(r1)
}

func (e *Everything) GetSearchA() *uint8 {
	r1, _, _ := getProc("Everything_GetSearchA").Call()
	return (*uint8)(unsafe.Pointer(r1))
}

func (e *Everything) GetSearchW() *uint16 {
	r1, _, _ := getProc("Everything_GetSearchW").Call()
	return (*uint16)(unsafe.Pointer(r1))
}

func (e *Everything) GetLastError() uint32 {
	r1, _, _ := getProc("Everything_GetLastError").Call()
	return uint32(r1)
}

func (e *Everything) GetReplyWindow() uintptr {
	r1, _, _ := getProc("Everything_GetReplyWindow").Call()
	return uintptr(r1)
}

func (e *Everything) GetReplyID() uint32 {
	r1, _, _ := getProc("Everything_GetReplyID").Call()
	return uint32(r1)
}

func (e *Everything) GetSort() uint32 {
	r1, _, _ := getProc("Everything_GetSort").Call()
	return uint32(r1)
}

func (e *Everything) GetRequestFlags() uint32 {
	r1, _, _ := getProc("Everything_GetRequestFlags").Call()
	return uint32(r1)
}

func (e *Everything) QueryA(BWait int32) int32 {
	r1, _, _ := getProc("Everything_QueryA").Call(uintptr(BWait))
	return int32(r1)
}

func (e *Everything) QueryW(BWait int32) int32 {
	r1, _, _ := getProc("Everything_QueryW").Call(uintptr(BWait))
	return int32(r1)
}

func (e *Everything) IsQueryReply(Message uint32, WParam uintptr, LParam uintptr, DwId uint32) int32 {
	r1, _, _ := getProc("Everything_IsQueryReply").Call(uintptr(Message), WParam, LParam, uintptr(DwId))
	return int32(r1)
}

func (e *Everything) SortResultsByPath() {
	getProc("Everything_SortResultsByPath").Call()
}

func (e *Everything) GetNumFileResults() uint32 {
	r1, _, _ := getProc("Everything_GetNumFileResults").Call()
	return uint32(r1)
}

func (e *Everything) GetNumFolderResults() uint32 {
	r1, _, _ := getProc("Everything_GetNumFolderResults").Call()
	return uint32(r1)
}

func (e *Everything) GetNumResults() uint32 {
	r1, _, _ := getProc("Everything_GetNumResults").Call()
	return uint32(r1)
}

func (e *Everything) GetTotFileResults() uint32 {
	r1, _, _ := getProc("Everything_GetTotFileResults").Call()
	return uint32(r1)
}

func (e *Everything) GetTotFolderResults() uint32 {
	r1, _, _ := getProc("Everything_GetTotFolderResults").Call()
	return uint32(r1)
}

func (e *Everything) GetTotResults() uint32 {
	r1, _, _ := getProc("Everything_GetTotResults").Call()
	return uint32(r1)
}

func (e *Everything) IsVolumeResult(DwIndex uint32) int32 {
	r1, _, _ := getProc("Everything_IsVolumeResult").Call(uintptr(DwIndex))
	return int32(r1)
}

func (e *Everything) IsFolderResult(DwIndex uint32) int32 {
	r1, _, _ := getProc("Everything_IsFolderResult").Call(uintptr(DwIndex))
	return int32(r1)
}

func (e *Everything) IsFileResult(DwIndex uint32) int32 {
	r1, _, _ := getProc("Everything_IsFileResult").Call(uintptr(DwIndex))
	return int32(r1)
}

func (e *Everything) GetResultFileNameW(DwIndex uint32) *uint16 {
	r1, _, _ := getProc("Everything_GetResultFileNameW").Call(uintptr(DwIndex))
	return (*uint16)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultFileNameA(DwIndex uint32) *uint8 {
	r1, _, _ := getProc("Everything_GetResultFileNameA").Call(uintptr(DwIndex))
	return (*uint8)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultPathW(DwIndex uint32) *uint16 {
	r1, _, _ := getProc("Everything_GetResultPathW").Call(uintptr(DwIndex))
	return (*uint16)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultPathA(DwIndex uint32) *uint8 {
	r1, _, _ := getProc("Everything_GetResultPathA").Call(uintptr(DwIndex))
	return (*uint8)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultFullPathNameA(DwIndex uint32, Buf *uint8, Bufsize uint32) uint32 {
	r1, _, _ := getProc("Everything_GetResultFullPathNameA").Call(uintptr(DwIndex), uintptr(unsafe.Pointer(Buf)), uintptr(Bufsize))
	return uint32(r1)
}

func (e *Everything) GetResultFullPathNameW(DwIndex uint32, Wbuf *uint16, Wbuf_size_in_wchars uint32) uint32 {
	r1, _, _ := getProc("Everything_GetResultFullPathNameW").Call(uintptr(DwIndex), uintptr(unsafe.Pointer(Wbuf)), uintptr(Wbuf_size_in_wchars))
	return uint32(r1)
}

func (e *Everything) GetResultListSort() uint32 {
	r1, _, _ := getProc("Everything_GetResultListSort").Call()
	return uint32(r1)
}

func (e *Everything) GetResultListRequestFlags() uint32 {
	r1, _, _ := getProc("Everything_GetResultListRequestFlags").Call()
	return uint32(r1)
}

func (e *Everything) GetResultExtensionW(DwIndex uint32) *uint16 {
	r1, _, _ := getProc("Everything_GetResultExtensionW").Call(uintptr(DwIndex))
	return (*uint16)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultExtensionA(DwIndex uint32) *uint8 {
	r1, _, _ := getProc("Everything_GetResultExtensionA").Call(uintptr(DwIndex))
	return (*uint8)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultSize(DwIndex uint32, LpSize *uintptr) int32 {
	r1, _, _ := getProc("Everything_GetResultSize").Call(uintptr(DwIndex), uintptr(unsafe.Pointer(LpSize)))
	return int32(r1)
}

func (e *Everything) GetResultDateCreated(DwIndex uint32, LpDateCreated *uintptr) int32 {
	r1, _, _ := getProc("Everything_GetResultDateCreated").Call(uintptr(DwIndex), uintptr(unsafe.Pointer(LpDateCreated)))
	return int32(r1)
}

func (e *Everything) GetResultDateModified(DwIndex uint32, LpDateModified *uintptr) int32 {
	r1, _, _ := getProc("Everything_GetResultDateModified").Call(uintptr(DwIndex), uintptr(unsafe.Pointer(LpDateModified)))
	return int32(r1)
}

func (e *Everything) GetResultDateAccessed(DwIndex uint32, LpDateAccessed *uintptr) int32 {
	r1, _, _ := getProc("Everything_GetResultDateAccessed").Call(uintptr(DwIndex), uintptr(unsafe.Pointer(LpDateAccessed)))
	return int32(r1)
}

func (e *Everything) GetResultAttributes(DwIndex uint32) uint32 {
	r1, _, _ := getProc("Everything_GetResultAttributes").Call(uintptr(DwIndex))
	return uint32(r1)
}

func (e *Everything) GetResultFileListFileNameW(DwIndex uint32) *uint16 {
	r1, _, _ := getProc("Everything_GetResultFileListFileNameW").Call(uintptr(DwIndex))
	return (*uint16)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultFileListFileNameA(DwIndex uint32) *uint8 {
	r1, _, _ := getProc("Everything_GetResultFileListFileNameA").Call(uintptr(DwIndex))
	return (*uint8)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultRunCount(DwIndex uint32) uint32 {
	r1, _, _ := getProc("Everything_GetResultRunCount").Call(uintptr(DwIndex))
	return uint32(r1)
}

func (e *Everything) GetResultDateRun(DwIndex uint32, LpDateRun *uintptr) int32 {
	r1, _, _ := getProc("Everything_GetResultDateRun").Call(uintptr(DwIndex), uintptr(unsafe.Pointer(LpDateRun)))
	return int32(r1)
}

func (e *Everything) GetResultDateRecentlyChanged(DwIndex uint32, LpDateRecentlyChanged *uintptr) int32 {
	r1, _, _ := getProc("Everything_GetResultDateRecentlyChanged").Call(uintptr(DwIndex), uintptr(unsafe.Pointer(LpDateRecentlyChanged)))
	return int32(r1)
}

func (e *Everything) GetResultHighlightedFileNameW(DwIndex uint32) *uint16 {
	r1, _, _ := getProc("Everything_GetResultHighlightedFileNameW").Call(uintptr(DwIndex))
	return (*uint16)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultHighlightedFileNameA(DwIndex uint32) *uint8 {
	r1, _, _ := getProc("Everything_GetResultHighlightedFileNameA").Call(uintptr(DwIndex))
	return (*uint8)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultHighlightedPathW(DwIndex uint32) *uint16 {
	r1, _, _ := getProc("Everything_GetResultHighlightedPathW").Call(uintptr(DwIndex))
	return (*uint16)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultHighlightedPathA(DwIndex uint32) *uint8 {
	r1, _, _ := getProc("Everything_GetResultHighlightedPathA").Call(uintptr(DwIndex))
	return (*uint8)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultHighlightedFullPathAndFileNameW(DwIndex uint32) *uint16 {
	r1, _, _ := getProc("Everything_GetResultHighlightedFullPathAndFileNameW").Call(uintptr(DwIndex))
	return (*uint16)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultHighlightedFullPathAndFileNameA(DwIndex uint32) *uint8 {
	r1, _, _ := getProc("Everything_GetResultHighlightedFullPathAndFileNameA").Call(uintptr(DwIndex))
	return (*uint8)(unsafe.Pointer(r1))
}

func (e *Everything) Reset() {
	getProc("Everything_Reset").Call()
}

func (e *Everything) CleanUp() {
	getProc("Everything_CleanUp").Call()
}

func (e *Everything) GetMajorVersion() uint32 {
	r1, _, _ := getProc("Everything_GetMajorVersion").Call()
	return uint32(r1)
}

func (e *Everything) GetMinorVersion() uint32 {
	r1, _, _ := getProc("Everything_GetMinorVersion").Call()
	return uint32(r1)
}

func (e *Everything) GetRevision() uint32 {
	r1, _, _ := getProc("Everything_GetRevision").Call()
	return uint32(r1)
}

func (e *Everything) GetBuildNumber() uint32 {
	r1, _, _ := getProc("Everything_GetBuildNumber").Call()
	return uint32(r1)
}

func (e *Everything) Exit() int32 {
	r1, _, _ := getProc("Everything_Exit").Call()
	return int32(r1)
}

func (e *Everything) MSIExitAndStopService(Msihandle unsafe.Pointer) uint32 {
	r1, _, _ := getProc("Everything_MSIExitAndStopService").Call(uintptr(Msihandle))
	return uint32(r1)
}

func (e *Everything) MSIStartService(Msihandle unsafe.Pointer) uint32 {
	r1, _, _ := getProc("Everything_MSIStartService").Call(uintptr(Msihandle))
	return uint32(r1)
}

func (e *Everything) IsDBLoaded() int32 {
	r1, _, _ := getProc("Everything_IsDBLoaded").Call()
	return int32(r1)
}

func (e *Everything) IsAdmin() int32 {
	r1, _, _ := getProc("Everything_IsAdmin").Call()
	return int32(r1)
}

func (e *Everything) IsAppData() int32 {
	r1, _, _ := getProc("Everything_IsAppData").Call()
	return int32(r1)
}

func (e *Everything) RebuildDB() int32 {
	r1, _, _ := getProc("Everything_RebuildDB").Call()
	return int32(r1)
}

func (e *Everything) UpdateAllFolderIndexes() int32 {
	r1, _, _ := getProc("Everything_UpdateAllFolderIndexes").Call()
	return int32(r1)
}

func (e *Everything) SaveDB() int32 {
	r1, _, _ := getProc("Everything_SaveDB").Call()
	return int32(r1)
}

func (e *Everything) SaveRunHistory() int32 {
	r1, _, _ := getProc("Everything_SaveRunHistory").Call()
	return int32(r1)
}

func (e *Everything) DeleteRunHistory() int32 {
	r1, _, _ := getProc("Everything_DeleteRunHistory").Call()
	return int32(r1)
}

func (e *Everything) GetTargetMachine() uint32 {
	r1, _, _ := getProc("Everything_GetTargetMachine").Call()
	return uint32(r1)
}

func (e *Everything) IsFastSort(SortType uint32) int32 {
	r1, _, _ := getProc("Everything_IsFastSort").Call(uintptr(SortType))
	return int32(r1)
}

func (e *Everything) IsFileInfoIndexed(FileInfoType uint32) int32 {
	r1, _, _ := getProc("Everything_IsFileInfoIndexed").Call(uintptr(FileInfoType))
	return int32(r1)
}

func (e *Everything) GetRunCountFromFileNameW(LpFileName *uint16) uint32 {
	r1, _, _ := getProc("Everything_GetRunCountFromFileNameW").Call(uintptr(unsafe.Pointer(LpFileName)))
	return uint32(r1)
}

func (e *Everything) GetRunCountFromFileNameA(LpFileName *uint8) uint32 {
	r1, _, _ := getProc("Everything_GetRunCountFromFileNameA").Call(uintptr(unsafe.Pointer(LpFileName)))
	return uint32(r1)
}

func (e *Everything) SetRunCountFromFileNameW(LpFileName *uint16, DwRunCount uint32) int32 {
	r1, _, _ := getProc("Everything_SetRunCountFromFileNameW").Call(uintptr(unsafe.Pointer(LpFileName)), uintptr(DwRunCount))
	return int32(r1)
}

func (e *Everything) SetRunCountFromFileNameA(LpFileName *uint8, DwRunCount uint32) int32 {
	r1, _, _ := getProc("Everything_SetRunCountFromFileNameA").Call(uintptr(unsafe.Pointer(LpFileName)), uintptr(DwRunCount))
	return int32(r1)
}

func (e *Everything) IncRunCountFromFileNameW(LpFileName *uint16) uint32 {
	r1, _, _ := getProc("Everything_IncRunCountFromFileNameW").Call(uintptr(unsafe.Pointer(LpFileName)))
	return uint32(r1)
}

func (e *Everything) IncRunCountFromFileNameA(LpFileName *uint8) uint32 {
	r1, _, _ := getProc("Everything_IncRunCountFromFileNameA").Call(uintptr(unsafe.Pointer(LpFileName)))
	return uint32(r1)
}
