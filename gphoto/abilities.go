package gphoto

// Binding sourced from: https://github.com/aqiank/go-gphoto2

// #cgo linux CFLAGS: -I/opt/shot-capture/include
// #include <gphoto2/gphoto2.h>
// #include <string.h>
import "C"

type CameraAbilitiesList struct {
	Ref     *C.CameraAbilitiesList
	Context *Context
}

type CameraAbilities struct {
	Ref C.CameraAbilities
}

func NewAbilitiesList() (*CameraAbilitiesList, error) {
	list := &CameraAbilitiesList{}
	if ret := C.gp_abilities_list_new(&list.Ref); ret != PORT_RESULT_OK {
		return nil, AsPortResult(ret).Error()
	}
	return list, nil
}

func (list *CameraAbilitiesList) Load() error {
	if ret := C.gp_abilities_list_load(list.Ref, list.Context.c()); ret != PORT_RESULT_OK {
		return AsPortResult(ret).Error()
	}
	return nil
}

func (list *CameraAbilitiesList) GetAbilities(index int) (*CameraAbilities, error) {
	abilities := &CameraAbilities{}
	if ret := C.gp_abilities_list_get_abilities(list.Ref, C.int(index), &abilities.Ref); ret != PORT_RESULT_OK {
		return nil, AsPortResult(ret).Error()
	}
	return abilities, nil
}

func (list *CameraAbilitiesList) LookupModel(model string) (int, error) {
	modelIndex := C.gp_abilities_list_lookup_model(list.Ref, C.CString(model))
	if modelIndex < PORT_RESULT_OK {
		return -1, AsPortResult(modelIndex).Error()
	}
	return int(modelIndex), nil
}
