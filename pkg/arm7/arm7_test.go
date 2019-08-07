package arm7

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type Arm7TestSuite struct {
	suite.Suite
	cpu CPU
}

func (suite *Arm7TestSuite) SetupTest() {
	suite.cpu.Registers.sysRegisters.R0 = 0x5
	suite.cpu.Registers.sysRegisters.R13 = 0x15
	suite.cpu.Registers.sysRegisters.R15 = 0x20
	suite.cpu.Registers.svcRegisters.R13 = 0x25
	suite.cpu.Registers.irqRegisters.R13 = 0x30
	suite.cpu.Registers.sysRegisters.Cpsr = 0x35
}

func TestArm7TestSuite(t *testing.T) {
	suite.Run(t, new(Arm7TestSuite))
}

func (suite *Arm7TestSuite) TestRegisterResetWhenNotBootingFromBIOS() {
	suite.cpu.Registers.Reset(false)

	assert.Equal(suite.T(), uint32(0x0), suite.cpu.Registers.sysRegisters.R0)
	assert.Equal(suite.T(), uint32(0x03007F00), suite.cpu.Registers.sysRegisters.R13)
	assert.Equal(suite.T(), uint32(0x8000000), suite.cpu.Registers.sysRegisters.R15)
	assert.Equal(suite.T(), uint32(0x03007FE0), suite.cpu.Registers.svcRegisters.R13)
	assert.Equal(suite.T(), uint32(0x03007FA0), suite.cpu.Registers.irqRegisters.R13)
	assert.Equal(suite.T(), uint32(0x5F), suite.cpu.Registers.sysRegisters.Cpsr)
}

func (suite *Arm7TestSuite) TestRegisterResetWhenBootingFromBIOS() {
	suite.cpu.Registers.Reset(true)

	assert.Equal(suite.T(), uint32(0x0), suite.cpu.Registers.sysRegisters.R0)
	assert.Equal(suite.T(), uint32(0x0), suite.cpu.Registers.sysRegisters.R13)
	assert.Equal(suite.T(), uint32(0x0), suite.cpu.Registers.sysRegisters.R15)
	assert.Equal(suite.T(), uint32(0x0), suite.cpu.Registers.svcRegisters.R13)
	assert.Equal(suite.T(), uint32(0x0), suite.cpu.Registers.irqRegisters.R13)
	assert.Equal(suite.T(), uint32(0xD3), suite.cpu.Registers.sysRegisters.Cpsr)
}

func (suite *Arm7TestSuite) TestGetGeneralPurposeRegister() {
	suite.cpu.CPUMode = SYS
	assert.Equal(suite.T(), uint32(0x5), suite.cpu.getRegister(0))
}

func (suite *Arm7TestSuite) TestGetGeneralPurposeRegisterWhenChangingCPUModes() {
	suite.cpu.CPUMode = SYS
	assert.Equal(suite.T(), uint32(0x15), suite.cpu.getRegister(13))

	suite.cpu.CPUMode = SVC
	assert.Equal(suite.T(), uint32(0x25), suite.cpu.getRegister(13))

	suite.cpu.CPUMode = IRQ
	assert.Equal(suite.T(), uint32(0x30), suite.cpu.getRegister(13))
}

func (suite *Arm7TestSuite) TestSetGeneralPurposeRegister() {
	suite.cpu.setRegister(0, 0x66)
	assert.Equal(suite.T(), uint32(0x66), suite.cpu.getRegister(0))
}

func (suite *Arm7TestSuite) TestSetGeneralPurposeRegisterWhenChangingCPUModes() {
	suite.cpu.CPUMode = SYS
	suite.cpu.setRegister(13, 0x66)

	suite.cpu.CPUMode = SVC
	suite.cpu.setRegister(13, 0x77)

	suite.cpu.CPUMode = IRQ
	suite.cpu.setRegister(13, 0x88)

	assert.Equal(suite.T(), uint32(0x66), suite.cpu.Registers.sysRegisters.R13)
	assert.Equal(suite.T(), uint32(0x77), suite.cpu.Registers.svcRegisters.R13)
	assert.Equal(suite.T(), uint32(0x88), suite.cpu.Registers.irqRegisters.R13)
}

func (suite *Arm7TestSuite) TestBranchWithLink() {
	suite.cpu.Registers.Reset(false)

	suite.cpu.BranchWithLink([]byte{0x2E, 0x0, 0x0, 0xEA})
}
