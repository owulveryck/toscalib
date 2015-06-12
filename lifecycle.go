package toscalib

// ToscaInterfacesNodeLifecycleStandard is a go interface for the standard normative lifecycle
type ToscaInterfacesNodeLifecycleStandard interface {
	Create()    // description: Standard lifecycle create operation.
	Configure() // description: Standard lifecycle configure operation.
	Start()     // description: Standard lifecycle start operation.
	Stop()      // description: Standard lifecycle stop operation.
	Delete()    //description: Standard lifecycle delete operation.
}
