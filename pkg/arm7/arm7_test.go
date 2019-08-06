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
	suite.cpu.registers.R0 = 0x5
	suite.cpu.registers.R13 = 0x15
	suite.cpu.registers.R15 = 0x20
	suite.cpu.registers.R13SupervisorMode = 0x25
	suite.cpu.registers.R13InterruptMode = 0x30
	suite.cpu.registers.Cpsr = 0x35
}

func TestArm7TestSuite(t *testing.T) {
	suite.Run(t, new(Arm7TestSuite))
}

func (suite *Arm7TestSuite) TestRegisterResetWhenNotBootingFromBIOS() {
	suite.cpu.registers.reset(false)

	assert.Equal(suite.T(), uint32(0x0), suite.cpu.registers.R0)
	assert.Equal(suite.T(), uint32(0x03007F00), suite.cpu.registers.R13)
	assert.Equal(suite.T(), uint32(0x8000000), suite.cpu.registers.R15)
	assert.Equal(suite.T(), uint32(0x03007FE0), suite.cpu.registers.R13SupervisorMode)
	assert.Equal(suite.T(), uint32(0x03007FA0), suite.cpu.registers.R13InterruptMode)
	assert.Equal(suite.T(), uint32(0x5F), suite.cpu.registers.Cpsr)
}

func (suite *Arm7TestSuite) TestRegisterResetWhenBootingFromBIOS() {
	suite.cpu.registers.reset(true)

	assert.Equal(suite.T(), uint32(0x0), suite.cpu.registers.R0)
	assert.Equal(suite.T(), uint32(0x0), suite.cpu.registers.R13)
	assert.Equal(suite.T(), uint32(0x0), suite.cpu.registers.R15)
	assert.Equal(suite.T(), uint32(0x0), suite.cpu.registers.R13SupervisorMode)
	assert.Equal(suite.T(), uint32(0x0), suite.cpu.registers.R13InterruptMode)
	assert.Equal(suite.T(), uint32(0xD3), suite.cpu.registers.Cpsr)
}

func (suite *Arm7TestSuite) TestGetGeneralPurposeRegister() {
	assert.Equal(suite.T(), uint32(0x5), suite.cpu.getRegister(0))
}

func (suite *Arm7TestSuite) TestGetGeneralPurposeRegisterWhenChangingCpuModes() {
	suite.cpu.cpuMode = SYS
	assert.Equal(suite.T(), uint32(0x15), suite.cpu.getRegister(13))

	suite.cpu.cpuMode = SVC
	assert.Equal(suite.T(), uint32(0x25), suite.cpu.getRegister(13))

	suite.cpu.cpuMode = IRQ
	assert.Equal(suite.T(), uint32(0x30), suite.cpu.getRegister(13))
}

func (suite *Arm7TestSuite) TestSetGeneralPurposeRegister() {
	suite.cpu.setRegister(0, 0x66)
	assert.Equal(suite.T(), uint32(0x66), suite.cpu.getRegister(0))
}

func (suite *Arm7TestSuite) TestSetGeneralPurposeRegisterWhenChangingCpuModes() {
	suite.cpu.cpuMode = SYS
	suite.cpu.setRegister(13, 0x66)

	suite.cpu.cpuMode = SVC
	suite.cpu.setRegister(13, 0x77)

	suite.cpu.cpuMode = IRQ
	suite.cpu.setRegister(13, 0x88)

	assert.Equal(suite.T(), uint32(0x66), suite.cpu.registers.R13)
	assert.Equal(suite.T(), uint32(0x77), suite.cpu.registers.R13SupervisorMode)
	assert.Equal(suite.T(), uint32(0x88), suite.cpu.registers.R13InterruptMode)
}
