package resource

const (
	NVIDIAVendorID      = "0x10de"
	NVIDIA              = "nvidia"
	HUAWEI              = "huawei"
	HUAWEIVendorID      = "0x19e5"
	HUAWEIRealVendorID  = "19e5"
	HUAWEI910BProductID = "0xd801"
	HUAWEI910ProductID  = "0xd802"
	HUAWEI310PProductID = "0xd500"
	KUNLUNVendorID      = "0x1d22"
	KUNLUNRealVendorID  = "1d22"
	MLUVendorID         = "0xcabc"
	MLURealVendorID     = "cabc"
	IXVendorID          = "0x1e3e"
	GPU                 = "gpu"
	NPU                 = "npu"
	XPU                 = "xpu"
	IX                  = "ix"
	MLU                 = "mlu"

	//The host root directory for the cnotainer
	HostRootDir = "/hostfs"

	ChangeRootCmd = "chroot"

	IXSmiCmd = "ixsmi"

	MLUCmd           = "cnmon"
	QueryMLUUUIDArgs = "-l"
	QueryMLUInfoArgs = "info"
	QueryMLUCard     = "-c"

	NvidiaSmiCmd      = "nvidia-smi"
	QueryGPUListGPU   = "--list-gpus"
	QueryGPUNameArgs  = "--query-gpu=gpu_name"
	QuerGPUFormatArgs = "--format=csv,noheader"
	QueryGPUIdArgs    = "--id=0"
	QueryGPUMemArgs   = "--query-gpu=memory.total"

	NPUSmiCmd             = "npu-smi"
	QueryNPUInfoArgs      = "info"
	QueryNPUMappingeArgs  = "-m"
	QueryNPUTopoArgs      = "-l"
	QueryNPUTypeArgs      = "-t"
	QueryNPUMemTypeArgs   = "memory"
	QueryNPUComTypeArgs   = "common"
	QueryNPUCpuNumCfgArgs = "cpu-num-cfg"
	QueryNPUCardIDArgs    = "-i"
	QueryNPUBoardArgs     = "board"
	QueryTopoArgs         = "topo"

	XPUSmiCmd                  = "xpu_smi"
	QueryXPUMachineReadbleArgs = "-m"
	QueryXPUDeivceIDQueryArgs  = "-d"
)
