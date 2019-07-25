package arm7

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type Arm7TestSuite struct {
	suite.Suite
	registers registerSet
}

func (suite *Arm7TestSuite) SetupTest() {
	suite.registers.r0 = 0x5
	suite.registers.r13 = 0x15
	suite.registers.r15 = 0x20
	suite.registers.r13SupervisorMode = 0x25
	suite.registers.r13InterruptMode = 0x30
	suite.registers.cpsr = 0x35
}

func TestArm7TestSuite(t *testing.T) {
	suite.Run(t, new(Arm7TestSuite))
}

func (suite *Arm7TestSuite) TestRegisterResetWhenNotBootingFromBIOS() {
	suite.registers.reset(false)

	assert.Equal(suite.T(), uint32(0x0), suite.registers.r0)
	assert.Equal(suite.T(), uint32(0x03007F00), suite.registers.r13)
	assert.Equal(suite.T(), uint32(0x8000000), suite.registers.r15)
	assert.Equal(suite.T(), uint32(0x03007FE0), suite.registers.r13SupervisorMode)
	assert.Equal(suite.T(), uint32(0x03007FA0), suite.registers.r13InterruptMode)
	assert.Equal(suite.T(), uint32(0x5F), suite.registers.cpsr)
}

func (suite *Arm7TestSuite) TestRegisterResetWhenBootingFromBIOS() {
	suite.registers.reset(true)

	assert.Equal(suite.T(), uint32(0x0), suite.registers.r0)
	assert.Equal(suite.T(), uint32(0x0), suite.registers.r13)
	assert.Equal(suite.T(), uint32(0x0), suite.registers.r15)
	assert.Equal(suite.T(), uint32(0x0), suite.registers.r13SupervisorMode)
	assert.Equal(suite.T(), uint32(0x0), suite.registers.r13InterruptMode)
	assert.Equal(suite.T(), uint32(0xD3), suite.registers.cpsr)
}

func (suite *Arm7TestSuite) TestGetGeneralPurposeRegister() {
	assert.Equal(suite.T(), uint32(0x5), suite.registers.getRegister(0x0))
}
