package cloudhsmiface

import (
	"github.com/awslabs/aws-sdk-go/service/cloudhsm"
)

type CloudHSMAPI interface {
	CreateHAPG(*cloudhsm.CreateHAPGInput) (*cloudhsm.CreateHAPGOutput, error)

	CreateHSM(*cloudhsm.CreateHSMInput) (*cloudhsm.CreateHSMOutput, error)

	CreateLunaClient(*cloudhsm.CreateLunaClientInput) (*cloudhsm.CreateLunaClientOutput, error)

	DeleteHAPG(*cloudhsm.DeleteHAPGInput) (*cloudhsm.DeleteHAPGOutput, error)

	DeleteHSM(*cloudhsm.DeleteHSMInput) (*cloudhsm.DeleteHSMOutput, error)

	DeleteLunaClient(*cloudhsm.DeleteLunaClientInput) (*cloudhsm.DeleteLunaClientOutput, error)

	DescribeHAPG(*cloudhsm.DescribeHAPGInput) (*cloudhsm.DescribeHAPGOutput, error)

	DescribeHSM(*cloudhsm.DescribeHSMInput) (*cloudhsm.DescribeHSMOutput, error)

	DescribeLunaClient(*cloudhsm.DescribeLunaClientInput) (*cloudhsm.DescribeLunaClientOutput, error)

	GetConfig(*cloudhsm.GetConfigInput) (*cloudhsm.GetConfigOutput, error)

	ListAvailableZones(*cloudhsm.ListAvailableZonesInput) (*cloudhsm.ListAvailableZonesOutput, error)

	ListHSMs(*cloudhsm.ListHSMsInput) (*cloudhsm.ListHSMsOutput, error)

	ListHapgs(*cloudhsm.ListHapgsInput) (*cloudhsm.ListHapgsOutput, error)

	ListLunaClients(*cloudhsm.ListLunaClientsInput) (*cloudhsm.ListLunaClientsOutput, error)

	ModifyHAPG(*cloudhsm.ModifyHAPGInput) (*cloudhsm.ModifyHAPGOutput, error)

	ModifyHSM(*cloudhsm.ModifyHSMInput) (*cloudhsm.ModifyHSMOutput, error)

	ModifyLunaClient(*cloudhsm.ModifyLunaClientInput) (*cloudhsm.ModifyLunaClientOutput, error)
}