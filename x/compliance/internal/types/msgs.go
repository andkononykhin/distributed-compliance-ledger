package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

type MsgAddModelInfo struct {
	VID                      int16          `json:"vid"`
	PID                      int16          `json:"pid"`
	CID                      int16          `json:"cid,omitempty"`
	Name                     string         `json:"name"`
	Description              string         `json:"description"`
	SKU                      string         `json:"sku"`
	FirmwareVersion          string         `json:"firmware_version"`
	HardwareVersion          string         `json:"hardware_version"`
	Custom                   string         `json:"custom"`
	CertificateID            string         `json:"certificate_id,omitempty"`
	CertifiedDate            time.Time      `json:"certified_date,omitempty"`
	TisOrTrpTestingCompleted bool           `json:"tis_or_trp_testing_completed"`
	Signer                   sdk.AccAddress `json:"signer"`
}

func NewMsgAddModelInfo(vid int16, pid int16, cid int16, name string, description string, sku string,
	firmwareVersion string, hardwareVersion string, custom string, certificateID string, certifiedDate time.Time,
	tisOrTrpTestingCompleted bool, signer sdk.AccAddress) MsgAddModelInfo {
	return MsgAddModelInfo{
		VID:                      vid,
		PID:                      pid,
		CID:                      cid,
		Name:                     name,
		Description:              description,
		SKU:                      sku,
		FirmwareVersion:          firmwareVersion,
		HardwareVersion:          hardwareVersion,
		Custom:                   custom,
		CertificateID:            certificateID,
		CertifiedDate:            certifiedDate,
		TisOrTrpTestingCompleted: tisOrTrpTestingCompleted,
		Signer:                   signer,
	}
}

func (m MsgAddModelInfo) Route() string {
	return RouterKey
}

func (m MsgAddModelInfo) Type() string {
	return "add_model_info"
}

func (m MsgAddModelInfo) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress(m.Signer.String())
	}

	if m.VID == 0 ||
		m.PID == 0 ||
		len(m.Name) == 0 ||
		len(m.Description) == 0 ||
		len(m.SKU) == 0 ||
		len(m.FirmwareVersion) == 0 ||
		len(m.HardwareVersion) == 0 ||
		len(m.Custom) == 0 {
		return sdk.ErrUnknownRequest("VID, PID, Name, Description, SKU, FirmwareVersion, HardwareVersion and Custom  " +
			"cannot be empty")
	}

	return nil
}

func (m MsgAddModelInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAddModelInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

type MsgUpdateModelInfo struct {
	VID                         int16          `json:"vid"`
	PID                         int16          `json:"pid"`
	NewCID                      int16          `json:"cid,omitempty"`
	NewName                     string         `json:"new_name"`
	NewOwner                    sdk.AccAddress `json:"new_owner"`
	NewDescription              string         `json:"new_description"`
	NewSKU                      string         `json:"new_sku"`
	NewFirmwareVersion          string         `json:"new_firmware_version"`
	NewHardwareVersion          string         `json:"new_hardware_version"`
	NewCustom                   string         `json:"new_custom"`
	NewCertificateID            string         `json:"new_certificate_id,omitempty"`
	NewCertifiedDate            time.Time      `json:"new_certified_date,omitempty"`
	NewTisOrTrpTestingCompleted bool           `json:"new_tis_or_trp_testing_completed"`
	Signer                      sdk.AccAddress `json:"signer"`
}

func NewMsgUpdateModelInfo(vid int16, pid int16, newCid int16, newName string, newOwner sdk.AccAddress,
	newDescription string, newSKU string, newFirmwareVersion string, newHardwareVersion string, newCustom string,
	newCertificateID string, newCertifiedDate time.Time, newTisOrTrpTestingCompleted bool,
	signer sdk.AccAddress) MsgUpdateModelInfo {
	return MsgUpdateModelInfo{
		VID:                         vid,
		PID:                         pid,
		NewCID:                      newCid,
		NewName:                     newName,
		NewOwner:                    newOwner,
		NewDescription:              newDescription,
		NewSKU:                      newSKU,
		NewFirmwareVersion:          newFirmwareVersion,
		NewHardwareVersion:          newHardwareVersion,
		NewCustom:                   newCustom,
		NewCertificateID:            newCertificateID,
		NewCertifiedDate:            newCertifiedDate,
		NewTisOrTrpTestingCompleted: newTisOrTrpTestingCompleted,
		Signer:                      signer,
	}
}

func (m MsgUpdateModelInfo) Route() string {
	return RouterKey
}

func (m MsgUpdateModelInfo) Type() string {
	return "update_model_info"
}

func (m MsgUpdateModelInfo) ValidateBasic() sdk.Error {
	if m.NewOwner.Empty() {
		return sdk.ErrInvalidAddress(m.NewOwner.String())
	}

	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress(m.Signer.String())
	}

	if m.VID == 0 ||
		m.PID == 0 ||
		len(m.NewName) == 0 ||
		len(m.NewDescription) == 0 ||
		len(m.NewSKU) == 0 ||
		len(m.NewFirmwareVersion) == 0 ||
		len(m.NewHardwareVersion) == 0 ||
		len(m.NewCustom) == 0 {
		return sdk.ErrUnknownRequest("VID, PID, NewName, NewDescription, NewSKU, NewFirmwareVersion, NewHardwareVersion and NewCustom " +
			"cannot be empty")
	}

	return nil
}

func (m MsgUpdateModelInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgUpdateModelInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

type MsgDeleteModelInfo struct {
	VID    int16          `json:"vid"`
	PID    int16          `json:"pid"`
	Signer sdk.AccAddress `json:"signer"`
}

func NewMsgDeleteModelInfo(vid int16, pid int16, signer sdk.AccAddress) MsgDeleteModelInfo {
	return MsgDeleteModelInfo{
		VID:    vid,
		PID:    pid,
		Signer: signer,
	}
}

func (m MsgDeleteModelInfo) Route() string {
	return RouterKey
}

func (m MsgDeleteModelInfo) Type() string {
	return "delete_model_info"
}

func (m MsgDeleteModelInfo) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress(m.Signer.String())
	}

	if m.VID == 0 || m.PID == 0 {
		return sdk.ErrUnknownRequest("VID and PID cannot be empty")
	}

	return nil
}

func (m MsgDeleteModelInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgDeleteModelInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
