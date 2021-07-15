package gphoto

// Binding sourced from: https://github.com/aqiank/go-gphoto2

// #cgo linux CFLAGS: -I/opt/shot-capture/include
// #include <gphoto2/gphoto2.h>
// #include <string.h>
import "C"

const (
	PORT_INFO_NONE            = C.GP_PORT_NONE
	PORT_INFO_SERIAL          = C.GP_PORT_SERIAL
	PORT_INFO_USB             = C.GP_PORT_USB
	PORT_INFO_DISK            = C.GP_PORT_DISK
	PORT_INFO_PTPIP           = C.GP_PORT_PTPIP
	PORT_INFO_USB_DISK_DIRECT = C.GP_PORT_USB_DISK_DIRECT
	PORT_INFO_USB_SCSI        = C.GP_PORT_USB_SCSI
	PORT_INFO_IP              = C.GP_PORT_IP
)

type PortInfoList struct {
	Ref  *C.GPPortInfoList
	Size int
}

type PortInfo struct {
	Ref  C.GPPortInfo
	Name string
}

func NewPortInfoList() (*PortInfoList, error) {
	list := &PortInfoList{}
	if ret := C.gp_port_info_list_new(&list.Ref); ret != PORT_RESULT_OK {
		return nil, AsPortResult(ret).Error()
	}
	return list, nil
}

// Load the port info list
func (list *PortInfoList) Load() error {
	if ret := C.gp_port_info_list_load(list.Ref); ret != PORT_RESULT_OK {
		return AsPortResult(ret).Error()
	}
	return nil
}

// Get the count of of ports in the list
func (list *PortInfoList) Count() error {
	count := C.gp_port_info_list_count(list.Ref)
	if count < PORT_RESULT_OK {
		return AsPortResult(count).Error()
	}
	list.Size = int(count)
	return nil
}

// Get the index of a port in the list from a path
func (list *PortInfoList) LookupPath(path string) (int, error) {
	index := C.gp_port_info_list_lookup_path(list.Ref, C.CString(path))
	if index < PORT_RESULT_OK {
		return -1, AsPortResult(index).Error()
	}
	return int(index), nil
}

// Get the indexed port info from the list
func (list *PortInfoList) GetInfo(index int) (*PortInfo, error) {
	info := &PortInfo{}
	if ret := C.gp_port_info_list_get_info(list.Ref, C.int(index), &info.Ref); ret != PORT_RESULT_OK {
		return nil, AsPortResult(ret).Error()
	}
	return info, nil
}

// Free resources for port info list
func (list *PortInfoList) Free() error {
	if ret := C.gp_port_info_list_free(list.Ref); ret != PORT_RESULT_OK {
		return AsPortResult(ret).Error()
	}
	return nil
}

// Get the name of the port info
func (info *PortInfo) GetName() error {
	var name *C.char
	if ret := C.gp_port_info_get_name(info.Ref, &name); ret != PORT_RESULT_OK {
		return AsPortResult(ret).Error()
	}
	info.Name = C.GoString(name)
	return nil
}
